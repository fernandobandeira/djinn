# ADR-20250120: Go Error Handling & slog Logging Strategy

## Status
Accepted

## Context

The Djinn personal finance application requires a consistent, comprehensive approach to error handling and logging across all Go services and components. Currently, the system lacks standardized error handling patterns, leading to:

- Inconsistent error responses across API endpoints
- Difficulty in debugging production issues due to inadequate logging
- Poor user experience when errors occur
- Challenges in monitoring system health and performance
- Security risks from exposing sensitive data in logs
- Inefficient error recovery mechanisms

Key forces driving this decision:
- **Technical**: Need for idiomatic Go error handling using standard library patterns
- **Business**: Requirement for reliable financial data processing with clear audit trails
- **Operational**: Need for effective monitoring and alerting for production issues
- **Regulatory**: Financial applications require comprehensive logging for compliance
- **Performance**: Modern Go 1.21+ slog provides structured logging with minimal overhead

## Decision

We will implement a **Go-Native Error Handling & slog Logging Strategy** that leverages Go's standard library patterns and the new structured logging capabilities.

### Core Components:

1. **Standard Library Error Handling with fmt.Errorf and error wrapping**
2. **Custom Error Types for Business Logic**
3. **slog Structured Logging with Context**
4. **Circuit Breaker Pattern for External Services**
5. **Middleware for Request Logging and Error Handling**

### Implementation Notes (from actual backend):
- **Simplified error wrapping**: Using `fmt.Errorf("failed to X: %w", err)` pattern consistently
- **Domain-specific errors**: Sentinel errors in domain layer (e.g., `ErrInvalidFirebaseUID`)
- **Service interfaces**: Error handling through interface boundaries (see ADR-20250127)
- **GraphQL error mapping**: Structured error responses with proper HTTP codes

### Error Classification:

```go
// Custom Error Types
type ValidationError struct {
    Field   string
    Value   interface{}
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation failed for field %s: %s", e.Field, e.Message)
}

type NotFoundError struct {
    Resource string
    ID       string
}

func (e NotFoundError) Error() string {
    return fmt.Sprintf("%s with ID %s not found", e.Resource, e.ID)
}

type UnauthorizedError struct {
    UserID string
    Action string
}

func (e UnauthorizedError) Error() string {
    return fmt.Sprintf("user %s not authorized for action: %s", e.UserID, e.Action)
}

type ExternalServiceError struct {
    Service string
    Err     error
}

func (e ExternalServiceError) Error() string {
    return fmt.Sprintf("external service %s failed: %v", e.Service, e.Err)
}

func (e ExternalServiceError) Unwrap() error {
    return e.Err
}

// Sentinel errors for common cases
var (
    ErrUserNotFound      = errors.New("user not found")
    ErrInvalidCredentials = errors.New("invalid credentials")
    ErrInvalidID         = errors.New("invalid ID format")
    ErrCircuitOpen       = errors.New("circuit breaker is open")
    ErrRateLimited       = errors.New("rate limit exceeded")
)
```

### slog Configuration:

```go
// Logger setup for different environments
func SetupLogger() *slog.Logger {
    var handler slog.Handler
    
    if os.Getenv("ENVIRONMENT") == "production" {
        handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
            Level:     slog.LevelInfo,
            AddSource: true,
        })
    } else {
        handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
            Level:     slog.LevelDebug,
            AddSource: true,
        })
    }
    
    logger := slog.New(handler)
    
    // Set as default logger
    slog.SetDefault(logger)
    
    return logger
}

// Context-aware logger with correlation ID
func LoggerWithCorrelationID(ctx context.Context, logger *slog.Logger) *slog.Logger {
    // Use centralized correlation ID utility from infrastructure/context
    if correlationID := ctxutil.GetCorrelationID(ctx); correlationID != "" {
        return logger.With(slog.String("correlation_id", correlationID))
    }
    return logger
}

// Service-specific logger
func NewServiceLogger(serviceName string) *slog.Logger {
    return slog.Default().With(
        slog.String("service", serviceName),
        slog.String("version", os.Getenv("APP_VERSION")),
    )
}
```

### HTTP Error Handling:

```go
// HTTP error response format
type ErrorResponse struct {
    Error struct {
        Code         string                 `json:"code"`
        Message      string                 `json:"message"`
        Details      map[string]interface{} `json:"details,omitempty"`
        CorrelationID string                `json:"correlation_id"`
        Timestamp    time.Time              `json:"timestamp"`
    } `json:"error"`
}

// HTTP error handler
func HandleHTTPError(w http.ResponseWriter, r *http.Request, err error, logger *slog.Logger) {
    correlationID := r.Header.Get("X-Correlation-ID")
    if correlationID == "" {
        correlationID = uuid.New().String()
    }
    
    reqLogger := logger.With(
        slog.String("correlation_id", correlationID),
        slog.String("method", r.Method),
        slog.String("path", r.URL.Path),
        slog.String("remote_addr", r.RemoteAddr),
    )
    
    var statusCode int
    var code, message string
    
    switch {
    case errors.Is(err, ErrUserNotFound):
        statusCode = http.StatusNotFound
        code = "USER_NOT_FOUND"
        message = "User not found"
        reqLogger.Info("user not found", slog.Any("error", err))
        
    case errors.Is(err, ErrInvalidCredentials):
        statusCode = http.StatusUnauthorized
        code = "INVALID_CREDENTIALS"
        message = "Invalid credentials"
        reqLogger.Warn("authentication failed", slog.Any("error", err))
        
    case errors.As(err, &ValidationError{}):
        statusCode = http.StatusBadRequest
        code = "VALIDATION_ERROR"
        message = "Validation failed"
        reqLogger.Warn("validation error", slog.Any("error", err))
        
    case errors.As(err, &ExternalServiceError{}):
        statusCode = http.StatusBadGateway
        code = "EXTERNAL_SERVICE_ERROR"
        message = "External service unavailable"
        reqLogger.Error("external service error", slog.Any("error", err))
        
    default:
        statusCode = http.StatusInternalServerError
        code = "INTERNAL_ERROR"
        message = "Internal server error"
        reqLogger.Error("internal server error", slog.Any("error", err))
    }
    
    response := ErrorResponse{
        Error: struct {
            Code         string                 `json:"code"`
            Message      string                 `json:"message"`
            Details      map[string]interface{} `json:"details,omitempty"`
            CorrelationID string                `json:"correlation_id"`
            Timestamp    time.Time              `json:"timestamp"`
        }{
            Code:         code,
            Message:      message,
            CorrelationID: correlationID,
            Timestamp:    time.Now().UTC(),
        },
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("X-Correlation-ID", correlationID)
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(response)
}
```

### gRPC Error Handling:

```go
import (
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func HandleGRPCError(ctx context.Context, err error, logger *slog.Logger) error {
    reqLogger := LoggerWithCorrelationID(ctx, logger)
    
    switch {
    case errors.Is(err, ErrUserNotFound):
        reqLogger.Info("user not found", slog.Any("error", err))
        return status.Error(codes.NotFound, "user not found")
        
    case errors.Is(err, ErrInvalidID):
        reqLogger.Warn("invalid user ID", slog.Any("error", err))
        return status.Error(codes.InvalidArgument, "invalid user ID")
        
    case errors.As(err, &ValidationError{}):
        reqLogger.Warn("validation error", slog.Any("error", err))
        return status.Error(codes.InvalidArgument, err.Error())
        
    case errors.As(err, &UnauthorizedError{}):
        reqLogger.Warn("unauthorized access", slog.Any("error", err))
        return status.Error(codes.PermissionDenied, "access denied")
        
    case errors.As(err, &ExternalServiceError{}):
        reqLogger.Error("external service error", slog.Any("error", err))
        return status.Error(codes.Unavailable, "service temporarily unavailable")
        
    default:
        reqLogger.Error("internal error", slog.Any("error", err))
        return status.Error(codes.Internal, "internal server error")
    }
}
```

### Circuit Breaker Pattern:

```go
type State int

const (
    Closed State = iota
    Open
    HalfOpen
)

type CircuitBreaker struct {
    maxFailures  int
    resetTimeout time.Duration
    failures     int
    lastFailTime time.Time
    state        State
    mu           sync.RWMutex
    logger       *slog.Logger
}

func NewCircuitBreaker(maxFailures int, resetTimeout time.Duration, logger *slog.Logger) *CircuitBreaker {
    return &CircuitBreaker{
        maxFailures:  maxFailures,
        resetTimeout: resetTimeout,
        state:        Closed,
        logger:       logger,
    }
}

func (cb *CircuitBreaker) Call(fn func() error) error {
    cb.mu.Lock()
    defer cb.mu.Unlock()
    
    if cb.state == Open {
        if time.Since(cb.lastFailTime) > cb.resetTimeout {
            cb.state = HalfOpen
            cb.logger.Info("circuit breaker transitioning to half-open")
        } else {
            return ErrCircuitOpen
        }
    }
    
    err := fn()
    if err != nil {
        cb.failures++
        cb.lastFailTime = time.Now()
        
        cb.logger.Warn("circuit breaker failure",
            slog.Int("failures", cb.failures),
            slog.Int("max_failures", cb.maxFailures),
            slog.Any("error", err),
        )
        
        if cb.failures >= cb.maxFailures {
            cb.state = Open
            cb.logger.Error("circuit breaker opened",
                slog.Time("will_retry_at", time.Now().Add(cb.resetTimeout)),
            )
        }
        return err
    }
    
    if cb.state == HalfOpen {
        cb.state = Closed
        cb.failures = 0
        cb.logger.Info("circuit breaker closed")
    }
    
    return nil
}
```

### Request Logging Middleware:

```go
type responseWriter struct {
    http.ResponseWriter
    statusCode int
    body       []byte
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.statusCode = code
    rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
    rw.body = b
    return rw.ResponseWriter.Write(b)
}

func RequestLogger(logger *slog.Logger) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            start := time.Now()
            
            // Generate correlation ID
            correlationID := r.Header.Get("X-Correlation-ID")
            if correlationID == "" {
                correlationID = uuid.New().String()
            }
            
            // Add to context
            ctx := ctxutil.WithCorrelationID(r.Context(), correlationID)
            r = r.WithContext(ctx)
            
            // Create request logger
            reqLogger := logger.With(
                slog.String("correlation_id", correlationID),
                slog.String("method", r.Method),
                slog.String("path", r.URL.Path),
                slog.String("remote_addr", r.RemoteAddr),
                slog.String("user_agent", r.UserAgent()),
            )
            
            // Wrap response writer
            wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
            
            // Log request start
            reqLogger.Info("request started")
            
            // Process request
            next.ServeHTTP(wrapped, r)
            
            // Log request completion
            duration := time.Since(start)
            reqLogger.Info("request completed",
                slog.Int("status", wrapped.statusCode),
                slog.Duration("duration", duration),
                slog.Int("response_size", len(wrapped.body)),
            )
            
            // Log slow requests
            if duration > 5*time.Second {
                reqLogger.Warn("slow request detected",
                    slog.Duration("duration", duration),
                )
            }
        })
    }
}
```

### Error Wrapping and Context:

```go
// Error wrapping with context
func ProcessUser(ctx context.Context, userID string) error {
    user, err := getUserFromDB(ctx, userID)
    if err != nil {
        return fmt.Errorf("failed to get user %s from database: %w", userID, err)
    }
    
    if err := validateUser(user); err != nil {
        return fmt.Errorf("user %s validation failed: %w", userID, err)
    }
    
    if err := updateUserProfile(ctx, user); err != nil {
        return fmt.Errorf("failed to update profile for user %s: %w", userID, err)
    }
    
    return nil
}

// Concurrent error handling with errgroup
import "golang.org/x/sync/errgroup"

func ProcessItems(ctx context.Context, items []Item, logger *slog.Logger) error {
    var g errgroup.Group
    
    for _, item := range items {
        item := item // capture loop variable
        g.Go(func() error {
            start := time.Now()
            if err := processItem(ctx, item); err != nil {
                logger.Error("failed to process item",
                    slog.String("item_id", item.ID),
                    slog.Any("error", err),
                    slog.Duration("duration", time.Since(start)),
                )
                return fmt.Errorf("processing item %s: %w", item.ID, err)
            }
            
            logger.Debug("item processed successfully",
                slog.String("item_id", item.ID),
                slog.Duration("duration", time.Since(start)),
            )
            return nil
        })
    }
    
    if err := g.Wait(); err != nil {
        logger.Error("failed to process items",
            slog.Any("error", err),
            slog.Int("item_count", len(items)),
        )
        return fmt.Errorf("failed to process items: %w", err)
    }
    
    logger.Info("all items processed successfully",
        slog.Int("item_count", len(items)),
    )
    return nil
}
```

## Consequences

### Positive

- **Go Idioms**: Uses standard library patterns familiar to Go developers
- **Performance**: slog provides structured logging with minimal allocation overhead
- **Type Safety**: Compile-time error checking with custom error types
- **Observability**: Rich structured logs enable effective monitoring and debugging
- **Reliability**: Circuit breakers and error wrapping improve system resilience
- **Security**: Structured approach prevents sensitive data exposure in logs
- **Compliance**: Comprehensive audit trails support regulatory requirements
- **Developer Experience**: Consistent patterns reduce cognitive load

### Negative

- **Learning Curve**: Developers must learn slog and error handling patterns
- **Verbosity**: Go error handling requires explicit error checking
- **Storage Requirements**: Structured logs may increase storage needs
- **Migration Effort**: Existing error handling needs to be updated

### Risks

- **Performance Impact**: Excessive logging could affect application performance
- **Log Volume**: High-frequency operations may generate large log volumes
- **Error Fatigue**: Too many wrapped errors could make debugging harder
- **Correlation ID Management**: Missing correlation IDs reduce traceability

## Alternatives Considered

### Option A: pkg/errors Package
**Description**: Use the popular pkg/errors library for error handling
**Pros**: 
- Rich error context and stack traces
- Well-established in Go community
- Good debugging capabilities

**Cons**: 
- Deprecated in favor of standard library
- Additional dependency
- Not compatible with Go 1.13+ error wrapping

**Reason for not choosing**: Standard library provides equivalent functionality

### Option B: zerolog for Logging  
**Description**: Use zerolog for high-performance structured logging
**Pros**:
- Extremely fast performance
- Zero allocation logging
- Rich feature set

**Cons**:
- Not part of standard library
- Less familiar to developers
- slog is the official Go solution

**Reason for not choosing**: slog is the standard and provides sufficient performance

### Option C: Structured Error Responses with Custom Types
**Description**: Create elaborate error type hierarchies for all error scenarios
**Pros**:
- Very specific error handling
- Rich error metadata
- Type-safe error handling

**Cons**:
- Overly complex for most use cases
- Maintenance overhead
- Can lead to over-engineering

**Reason for not choosing**: Simple custom types provide better balance

### Option D: External Error Tracking Only
**Description**: Rely entirely on external services (Sentry, Rollbar) for error handling
**Pros**:
- Proven, feature-rich solutions
- Advanced error tracking capabilities
- Reduced development effort

**Cons**:
- Vendor lock-in and costs
- Data privacy concerns with financial data
- Limited customization for Go patterns

**Reason for not choosing**: Financial data sensitivity requires internal control

## Implementation Notes

### Phase 1: Core Infrastructure (Weeks 1-2)
- Set up slog configuration for different environments
- Implement custom error types for common business scenarios
- Create HTTP and gRPC error handling utilities
- Establish correlation ID middleware

### Phase 2: Advanced Patterns (Weeks 3-4)
- Implement circuit breaker pattern for external services
- Add request logging middleware with performance monitoring
- Create error wrapping utilities and context management
- Set up basic monitoring dashboards

### Phase 3: Integration & Testing (Weeks 5-6)
- Integrate with existing services gradually
- Add comprehensive error handling tests
- Performance test logging overhead
- Create developer documentation and examples

### Phase 4: Monitoring & Optimization (Weeks 7-8)
- Set up log aggregation and alerting
- Optimize log levels and filtering rules
- Integrate with OpenTelemetry for distributed tracing
- Conduct load testing and fine-tuning

### Migration Strategy:
1. Implement new error handling alongside existing code
2. Update service by service to use new patterns
3. Maintain backward compatibility during transition
4. Remove legacy error handling after migration completion

### Success Metrics:
- Reduced mean time to resolution (MTTR) for production issues
- Improved error clarity in customer support tickets
- Better system uptime and reliability metrics
- Developer satisfaction with debugging capabilities
- Performance benchmarks within acceptable ranges

## References

- [Related ADR: ADR-20250120-monitoring-observability.md](/docs/architecture/adrs/ADR-20250120-monitoring-observability.md)
- [Related ADR: ADR-20250819-api-design-standards.md](/docs/architecture/adrs/ADR-20250819-api-design-standards.md) 
- [Related ADR: ADR-20250819-security-architecture.md](/docs/architecture/adrs/ADR-20250819-security-architecture.md)
- [Go Error Handling Best Practices](https://go.dev/blog/error-handling-and-go)
- [slog Package Documentation](https://pkg.go.dev/log/slog)
- [Go 1.13 Error Wrapping](https://go.dev/blog/go1.13-errors)
- [Circuit Breaker Pattern - Martin Fowler](https://martinfowler.com/bliki/CircuitBreaker.html)
- [Structured Logging with slog](https://betterstack.com/community/guides/logging/logging-in-go/)
- [Go Concurrency Patterns](https://go.dev/blog/pipelines)

## Decision Makers

- Author: Archie (System Architect)
- Reviewers: [To be assigned]
- Approver: [To be assigned]
- Date: 2025-01-20

## Implementation Update (August 2025)

### GraphQL-Specific Error Handling

As the application uses GraphQL as its primary API layer (not HTTP/gRPC), we've implemented a specialized GraphQL error presenter:

```go
// GraphQL Error Presenter with Extensions
func GraphQLErrorPresenter(logger *slog.Logger) graphql.ErrorPresenterFunc {
    return func(ctx context.Context, err error) *gqlerror.Error {
        correlationID := ctxutil.GetCorrelationID(ctx)
        extensions := map[string]interface{}{
            "correlationId": correlationID,
        }
        
        // Map custom error types to GraphQL extensions
        var valErr *ValidationError
        if errors.As(err, &valErr) {
            extensions["code"] = string(CodeValidationFailed)
            extensions["field"] = valErr.Field
            if valErr.Value != nil {
                extensions["value"] = valErr.Value
            }
            // ... detailed error mapping
        }
        
        // Return structured GraphQL error
        gqlErr.Extensions = extensions
        return gqlErr
    }
}
```

### Domain-Specific Error Patterns

We've implemented a domain-driven approach where domain packages define their own errors that extend infrastructure error types:

```go
// Domain-specific errors in internal/domain/user/errors.go
package user

import "github.com/fernandobandeira/djinn/backend/internal/infrastructure/errors"

// Domain-specific error codes
const (
    CodeInvalidFirebaseUID = "USER_INVALID_FIREBASE_UID"
    CodeInvalidEmail       = "USER_INVALID_EMAIL"
    CodeInvalidName        = "USER_INVALID_NAME"
)

// Domain-specific errors using infrastructure error types
var (
    ErrUserNotFound      = &errors.NotFoundError{Resource: "user"}
    ErrUserAlreadyExists = &errors.ConflictError{Resource: "user", Message: "user already exists"}
    
    ErrInvalidFirebaseUID = &errors.ValidationError{
        Field:   "FirebaseUID",
        Message: "invalid Firebase UID",
        Code:    CodeInvalidFirebaseUID,
    }
    
    ErrInvalidEmail = &errors.ValidationError{
        Field:   "Email",
        Message: "invalid email address",
        Code:    CodeInvalidEmail,
    }
)
```

### Error Code Mapping System

Instead of HTTP/gRPC handlers, we've implemented a centralized error code mapper:

```go
// ErrorCodeMapper maps errors to standardized codes
type ErrorCodeMapper struct{}

func (m *ErrorCodeMapper) MapError(err error) ErrorCode {
    // Check for typed errors first
    var valErr *ValidationError
    if errors.As(err, &valErr) {
        return CodeValidationFailed
    }
    
    // Check for sentinel errors
    switch {
    case errors.Is(err, ErrInvalidCredentials):
        return CodeInvalidCredentials
    case errors.Is(err, ErrUserNotFound):
        return CodeUserNotFound
    // ... comprehensive error mapping
    }
}
```

### Architecture Benefits

1. **Domain Independence**: Each domain package defines its business-specific errors
2. **Infrastructure Reuse**: Common error types are shared across domains  
3. **GraphQL Integration**: Errors are properly formatted for GraphQL responses
4. **Type Safety**: Domain errors use infrastructure types for consistency
5. **Correlation Tracking**: All errors include correlation IDs for tracing

### Updated Implementation Status

- ✅ **GraphQL Error Presenter**: Implemented with structured error extensions
- ✅ **Domain Error Patterns**: User domain demonstrates pattern for other domains
- ✅ **Error Code Mapping**: Centralized mapping system replacing HTTP/gRPC handlers
- ✅ **Correlation ID Support**: Consistent tracking across all error scenarios
- ✅ **Test Coverage**: Comprehensive tests for error infrastructure

### Lessons Learned

1. **GraphQL-First Design**: Specialized error handling provides better developer experience
2. **Domain Autonomy**: Domain packages should own their error definitions
3. **Infrastructure Consistency**: Shared error types ensure consistent behavior
4. **Extensible Design**: New domains can easily adopt the established patterns
5. **Clean Architecture Principles**: Efficient type switching outperforms pseudo-strategy patterns
6. **Code Elimination**: Removing unused HTTP/gRPC handlers simplified the architecture significantly

### Final Implementation Architecture

The final implementation consists of:

```go
// Core error mapping with O(1) type switching
func MapError(err error, ctx context.Context, logger *slog.Logger) *ErrorMapping {
    switch e := err.(type) {
    case *ValidationError:
        return &ErrorMapping{...}
    case *AuthenticationError:
        code := mapAuthenticationReason(e.Reason) // Context-aware
        return &ErrorMapping{...}
    default:
        // Handle sentinel errors efficiently
        if mapping := mapSentinelError(err, baseExtensions); mapping != nil {
            return mapping
        }
        return &ErrorMapping{Code: CodeUnknown, ...}
    }
}

// Clean GraphQL presenter (38 lines)
func GraphQLErrorPresenter(logger *slog.Logger) graphql.ErrorPresenterFunc {
    return func(ctx context.Context, err error) *gqlerror.Error {
        mapping := MapError(err, ctx, logger)
        gqlErr.Message = mapping.Message
        gqlErr.Extensions = mapping.Extensions
        logger.Log(ctx, mapping.LogLevel, "GraphQL error", ...)
        return gqlErr
    }
}
```

**Performance Characteristics:**
- **O(1) error mapping** via efficient type switching
- **Single memory allocation** per error via pre-sized extension maps
- **38-line GraphQL presenter** (reduced from 250+ lines)
- **Zero runtime strategy lookup overhead**

## Revision History

- 2025-01-20: Initial draft created with TypeScript/Node.js patterns
- 2025-01-20: Complete rewrite with Go-specific patterns and slog integration
- 2025-08-27: Implementation update with GraphQL-specific handling and domain patterns
- 2025-08-27: Standardized correlation ID handling with centralized `internal/infrastructure/context` utility
- 2025-08-27: **Architecture cleanup** - Eliminated redundant HTTP/gRPC handlers, implemented efficient type switching, removed pseudo-strategy pattern in favor of O(1) dispatch