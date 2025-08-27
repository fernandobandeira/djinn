package errors

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/vektah/gqlparser/v2/gqlerror"
	ctxutil "github.com/fernandobandeira/djinn/backend/internal/infrastructure/context"
)

// ErrorContext holds contextual information for error handling
type ErrorContext struct {
	CorrelationID string
	Timestamp     time.Time
	Logger        *slog.Logger
	Context       context.Context
}

// ErrorMapping defines how an error should be handled
type ErrorMapping struct {
	Code       ErrorCode
	Message    string
	Extensions map[string]interface{}
	LogLevel   slog.Level
}

// ErrorMapper handles error mapping with a clean strategy pattern
type ErrorMapper struct {
	strategies map[string]func(error, *ErrorContext) *ErrorMapping
}

// NewErrorMapper creates a new error mapper with all strategies registered
func NewErrorMapper() *ErrorMapper {
	mapper := &ErrorMapper{
		strategies: make(map[string]func(error, *ErrorContext) *ErrorMapping),
	}
	
	// Register all error type strategies
	mapper.registerStrategies()
	return mapper
}

// MapError maps any error to a standardized error mapping
func (m *ErrorMapper) MapError(err error, ctx *ErrorContext) *ErrorMapping {
	// Try typed error strategies first
	for _, strategy := range m.strategies {
		if mapping := strategy(err, ctx); mapping != nil {
			// Add standard extensions
			mapping.Extensions["correlationId"] = ctx.CorrelationID
			mapping.Extensions["timestamp"] = ctx.Timestamp.Format(time.RFC3339)
			return mapping
		}
	}
	
	// Default mapping for unknown errors
	return &ErrorMapping{
		Code:    CodeUnknown,
		Message: "An unexpected error occurred",
		Extensions: map[string]interface{}{
			"correlationId": ctx.CorrelationID,
			"timestamp":     ctx.Timestamp.Format(time.RFC3339),
		},
		LogLevel: slog.LevelError,
	}
}

// registerStrategies registers all error handling strategies
func (m *ErrorMapper) registerStrategies() {
	m.strategies["validation"] = m.handleValidationError
	m.strategies["notfound"] = m.handleNotFoundError
	m.strategies["unauthorized"] = m.handleUnauthorizedError
	m.strategies["conflict"] = m.handleConflictError
	m.strategies["ratelimit"] = m.handleRateLimitError
	m.strategies["external"] = m.handleExternalServiceError
	m.strategies["internal"] = m.handleInternalError
	m.strategies["authentication"] = m.handleAuthenticationError
	m.strategies["authorization"] = m.handleAuthorizationError
	m.strategies["sentinel"] = m.handleSentinelErrors
}

// Validation error strategy
func (m *ErrorMapper) handleValidationError(err error, ctx *ErrorContext) *ErrorMapping {
	var valErr *ValidationError
	if !errors.As(err, &valErr) {
		return nil
	}
	
	extensions := map[string]interface{}{
		"code":  string(CodeValidationFailed),
		"field": valErr.Field,
	}
	if valErr.Value != nil {
		extensions["value"] = valErr.Value
	}
	
	return &ErrorMapping{
		Code:       CodeValidationFailed,
		Message:    valErr.Error(),
		Extensions: extensions,
		LogLevel:   slog.LevelInfo,
	}
}

// Not found error strategy
func (m *ErrorMapper) handleNotFoundError(err error, ctx *ErrorContext) *ErrorMapping {
	var notFoundErr *NotFoundError
	if !errors.As(err, &notFoundErr) {
		return nil
	}
	
	extensions := map[string]interface{}{
		"code":     string(CodeNotFound),
		"resource": notFoundErr.Resource,
	}
	if notFoundErr.ID != "" {
		extensions["id"] = notFoundErr.ID
	}
	
	return &ErrorMapping{
		Code:       CodeNotFound,
		Message:    notFoundErr.Error(),
		Extensions: extensions,
		LogLevel:   slog.LevelInfo,
	}
}

// Unauthorized error strategy
func (m *ErrorMapper) handleUnauthorizedError(err error, ctx *ErrorContext) *ErrorMapping {
	var unauthorizedErr *UnauthorizedError
	if !errors.As(err, &unauthorizedErr) {
		return nil
	}
	
	return &ErrorMapping{
		Code:    CodeUnauthorized,
		Message: unauthorizedErr.Error(),
		Extensions: map[string]interface{}{
			"code":   string(CodeUnauthorized),
			"reason": unauthorizedErr.Reason,
		},
		LogLevel: slog.LevelInfo,
	}
}

// Conflict error strategy
func (m *ErrorMapper) handleConflictError(err error, ctx *ErrorContext) *ErrorMapping {
	var conflictErr *ConflictError
	if !errors.As(err, &conflictErr) {
		return nil
	}
	
	extensions := map[string]interface{}{
		"code":     string(CodeConflict),
		"resource": conflictErr.Resource,
	}
	if conflictErr.Field != "" {
		extensions["field"] = conflictErr.Field
	}
	if conflictErr.Value != nil {
		extensions["value"] = conflictErr.Value
	}
	
	return &ErrorMapping{
		Code:       CodeConflict,
		Message:    conflictErr.Error(),
		Extensions: extensions,
		LogLevel:   slog.LevelInfo,
	}
}

// Rate limit error strategy
func (m *ErrorMapper) handleRateLimitError(err error, ctx *ErrorContext) *ErrorMapping {
	var rateLimitErr *RateLimitError
	if !errors.As(err, &rateLimitErr) {
		return nil
	}
	
	return &ErrorMapping{
		Code:    CodeRateLimited,
		Message: rateLimitErr.Error(),
		Extensions: map[string]interface{}{
			"code":       string(CodeRateLimited),
			"resource":   rateLimitErr.Resource,
			"retryAfter": rateLimitErr.RetryAfter.Seconds(),
		},
		LogLevel: slog.LevelWarn,
	}
}

// External service error strategy
func (m *ErrorMapper) handleExternalServiceError(err error, ctx *ErrorContext) *ErrorMapping {
	var extErr *ExternalServiceError
	if !errors.As(err, &extErr) {
		return nil
	}
	
	extensions := map[string]interface{}{
		"code":    string(CodeExternalServiceError),
		"service": extErr.Service,
	}
	if extErr.RetryAfter > 0 {
		extensions["retryAfter"] = extErr.RetryAfter.Seconds()
	}
	
	return &ErrorMapping{
		Code:       CodeExternalServiceError,
		Message:    "External service error",
		Extensions: extensions,
		LogLevel:   slog.LevelError,
	}
}

// Internal error strategy
func (m *ErrorMapper) handleInternalError(err error, ctx *ErrorContext) *ErrorMapping {
	var internalErr *InternalError
	if !errors.As(err, &internalErr) {
		return nil
	}
	
	return &ErrorMapping{
		Code:    CodeInternalError,
		Message: "Internal server error", // Don't expose internal details
		Extensions: map[string]interface{}{
			"code": string(CodeInternalError),
		},
		LogLevel: slog.LevelError,
	}
}

// Authentication error strategy
func (m *ErrorMapper) handleAuthenticationError(err error, ctx *ErrorContext) *ErrorMapping {
	var authErr *AuthenticationError
	if !errors.As(err, &authErr) {
		return nil
	}
	
	// Context-aware error code mapping
	var errorCode ErrorCode
	switch authErr.Reason {
	case "token_expired":
		errorCode = CodeTokenExpired
	case "token_invalid":
		errorCode = CodeInvalidCredentials
	case "insufficient_privileges":
		errorCode = CodePermissionDenied
	default:
		errorCode = CodeInvalidCredentials
	}
	
	extensions := map[string]interface{}{
		"code":   string(errorCode),
		"reason": authErr.Reason,
	}
	if authErr.Method != "" {
		extensions["method"] = authErr.Method
	}
	
	return &ErrorMapping{
		Code:       errorCode,
		Message:    authErr.Error(),
		Extensions: extensions,
		LogLevel:   slog.LevelInfo,
	}
}

// Authorization error strategy
func (m *ErrorMapper) handleAuthorizationError(err error, ctx *ErrorContext) *ErrorMapping {
	var authzErr *AuthorizationError
	if !errors.As(err, &authzErr) {
		return nil
	}
	
	extensions := map[string]interface{}{
		"code":     string(CodePermissionDenied),
		"resource": authzErr.Resource,
		"action":   authzErr.Action,
	}
	if authzErr.Principal != "" {
		extensions["principal"] = authzErr.Principal
	}
	
	return &ErrorMapping{
		Code:       CodePermissionDenied,
		Message:    authzErr.Error(),
		Extensions: extensions,
		LogLevel:   slog.LevelInfo,
	}
}

// Sentinel errors strategy
func (m *ErrorMapper) handleSentinelErrors(err error, ctx *ErrorContext) *ErrorMapping {
	sentinelMappings := map[error]struct {
		code     ErrorCode
		message  string
		logLevel slog.Level
	}{
		ErrInvalidCredentials:   {CodeInvalidCredentials, "Invalid credentials", slog.LevelInfo},
		ErrTokenExpired:         {CodeTokenExpired, "Token has expired", slog.LevelInfo},
		ErrTokenInvalid:         {CodeInvalidCredentials, "Invalid token", slog.LevelInfo},
		ErrUnauthorized:         {CodeUnauthorized, "Unauthorized", slog.LevelInfo},
		ErrPermissionDenied:     {CodePermissionDenied, "Permission denied", slog.LevelInfo},
		ErrRateLimited:          {CodeRateLimited, "Rate limit exceeded", slog.LevelWarn},
		ErrServiceUnavailable:   {CodeServiceUnavailable, "Service temporarily unavailable", slog.LevelError},
		ErrBadRequest:           {CodeInvalidInput, "Invalid request format", slog.LevelInfo}, // Sanitized message
		ErrUserNotFound:         {CodeUserNotFound, "User not found", slog.LevelInfo},
		ErrResourceNotFound:     {CodeNotFound, "Resource not found", slog.LevelInfo},
		ErrAlreadyExists:        {CodeAlreadyExists, "Resource already exists", slog.LevelInfo},
		ErrInvalidID:            {CodeInvalidInput, "Invalid ID format", slog.LevelInfo},
		ErrInvalidInput:         {CodeInvalidInput, "Invalid input", slog.LevelInfo},
		ErrMissingField:         {CodeMissingField, "Required field missing", slog.LevelInfo},
		ErrCircuitOpen:          {CodeServiceUnavailable, "Service temporarily unavailable", slog.LevelError},
		ErrCircuitHalfOpen:      {CodeServiceUnavailable, "Service temporarily unavailable", slog.LevelError},
		ErrConflict:             {CodeConflict, "Resource conflict", slog.LevelInfo},
		ErrTimeout:              {CodeInternalError, "Request timeout", slog.LevelError},
		ErrInternal:             {CodeInternalError, "Internal server error", slog.LevelError},
	}
	
	for sentinelErr, mapping := range sentinelMappings {
		if errors.Is(err, sentinelErr) {
			return &ErrorMapping{
				Code:    mapping.code,
				Message: mapping.message,
				Extensions: map[string]interface{}{
					"code": string(mapping.code),
				},
				LogLevel: mapping.logLevel,
			}
		}
	}
	
	return nil
}

// CreateErrorContext creates an error context from a Go context
func CreateErrorContext(ctx context.Context, logger *slog.Logger) *ErrorContext {
	return &ErrorContext{
		CorrelationID: ctxutil.GetCorrelationID(ctx),
		Timestamp:     time.Now().UTC(),
		Logger:        logger,
		Context:       ctx,
	}
}