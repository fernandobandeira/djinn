package errors

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleHTTPError handles errors and writes appropriate HTTP responses
func HandleHTTPError(ctx context.Context, w http.ResponseWriter, err error, logger *slog.Logger) {
	correlationID := getCorrelationID(ctx)
	
	var (
		statusCode int
		errResp    *ErrorResponse
	)

	// Check for specific error types
	switch {
	case errors.Is(err, ErrUserNotFound):
		statusCode = http.StatusNotFound
		errResp = NewErrorResponse(CodeUserNotFound, "User not found", correlationID)
		logger.InfoContext(ctx, "user not found",
			"correlation_id", correlationID,
			"error", err.Error())

	case errors.Is(err, ErrResourceNotFound):
		statusCode = http.StatusNotFound
		errResp = NewErrorResponse(CodeNotFound, "Resource not found", correlationID)
		logger.InfoContext(ctx, "resource not found",
			"correlation_id", correlationID,
			"error", err.Error())

	case errors.Is(err, ErrInvalidCredentials):
		statusCode = http.StatusUnauthorized
		errResp = NewErrorResponse(CodeInvalidCredentials, "Invalid credentials", correlationID)
		logger.InfoContext(ctx, "invalid credentials attempt",
			"correlation_id", correlationID)

	case errors.Is(err, ErrTokenExpired):
		statusCode = http.StatusUnauthorized
		errResp = NewErrorResponse(CodeTokenExpired, "Token has expired", correlationID)
		logger.InfoContext(ctx, "token expired",
			"correlation_id", correlationID)

	case errors.Is(err, ErrPermissionDenied):
		statusCode = http.StatusForbidden
		errResp = NewErrorResponse(CodePermissionDenied, "Permission denied", correlationID)
		logger.InfoContext(ctx, "permission denied",
			"correlation_id", correlationID)

	case errors.Is(err, ErrAlreadyExists):
		statusCode = http.StatusConflict
		errResp = NewErrorResponse(CodeAlreadyExists, "Resource already exists", correlationID)
		logger.InfoContext(ctx, "resource already exists",
			"correlation_id", correlationID,
			"error", err.Error())

	case errors.Is(err, ErrRateLimited):
		statusCode = http.StatusTooManyRequests
		errResp = NewErrorResponse(CodeRateLimited, "Rate limit exceeded", correlationID)
		logger.WarnContext(ctx, "rate limit exceeded",
			"correlation_id", correlationID)

	case errors.Is(err, ErrCircuitOpen):
		statusCode = http.StatusServiceUnavailable
		errResp = NewErrorResponse(CodeCircuitOpen, "Service temporarily unavailable", correlationID)
		logger.ErrorContext(ctx, "circuit breaker open",
			"correlation_id", correlationID)

	default:
		// Check for custom error types
		statusCode, errResp = handleCustomError(ctx, err, correlationID, logger)
	}

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Correlation-ID", correlationID)
	w.WriteHeader(statusCode)
	
	if err := json.NewEncoder(w).Encode(errResp); err != nil {
		logger.ErrorContext(ctx, "failed to encode error response",
			"correlation_id", correlationID,
			"error", err.Error())
	}
}

func handleCustomError(ctx context.Context, err error, correlationID string, logger *slog.Logger) (int, *ErrorResponse) {
	var valErr *ValidationError
	if errors.As(err, &valErr) {
		errResp := NewErrorResponse(CodeValidationFailed, "Validation failed", correlationID)
		errResp.AddDetail(valErr.Field, valErr.Message, valErr.Value)
		logger.InfoContext(ctx, "validation error",
			"correlation_id", correlationID,
			"field", valErr.Field,
			"error", valErr.Error())
		return http.StatusBadRequest, errResp
	}

	var notFoundErr *NotFoundError
	if errors.As(err, &notFoundErr) {
		errResp := NewErrorResponse(CodeNotFound, notFoundErr.Error(), correlationID)
		logger.InfoContext(ctx, "not found error",
			"correlation_id", correlationID,
			"resource", notFoundErr.Resource,
			"id", notFoundErr.ID)
		return http.StatusNotFound, errResp
	}

	var unauthorizedErr *UnauthorizedError
	if errors.As(err, &unauthorizedErr) {
		errResp := NewErrorResponse(CodeUnauthorized, unauthorizedErr.Error(), correlationID)
		logger.InfoContext(ctx, "unauthorized error",
			"correlation_id", correlationID,
			"reason", unauthorizedErr.Reason)
		return http.StatusUnauthorized, errResp
	}

	var extErr *ExternalServiceError
	if errors.As(err, &extErr) {
		errResp := NewErrorResponse(CodeExternalServiceError, "External service error", correlationID)
		if extErr.RetryAfter > 0 {
			errResp.SetRetryAfter(extErr.RetryAfter)
		}
		logger.ErrorContext(ctx, "external service error",
			"correlation_id", correlationID,
			"service", extErr.Service,
			"operation", extErr.Operation,
			"status_code", extErr.StatusCode,
			"error", extErr.Error())
		return http.StatusBadGateway, errResp
	}

	var rateLimitErr *RateLimitError
	if errors.As(err, &rateLimitErr) {
		errResp := NewErrorResponse(CodeRateLimited, rateLimitErr.Error(), correlationID)
		errResp.SetRetryAfter(rateLimitErr.RetryAfter)
		logger.WarnContext(ctx, "rate limit error",
			"correlation_id", correlationID,
			"resource", rateLimitErr.Resource,
			"limit", rateLimitErr.Limit)
		return http.StatusTooManyRequests, errResp
	}

	var conflictErr *ConflictError
	if errors.As(err, &conflictErr) {
		errResp := NewErrorResponse(CodeConflict, conflictErr.Error(), correlationID)
		if conflictErr.Field != "" {
			errResp.AddDetail(conflictErr.Field, conflictErr.Message, conflictErr.Value)
		}
		logger.InfoContext(ctx, "conflict error",
			"correlation_id", correlationID,
			"resource", conflictErr.Resource,
			"field", conflictErr.Field)
		return http.StatusConflict, errResp
	}

	var internalErr *InternalError
	if errors.As(err, &internalErr) {
		errResp := NewErrorResponse(CodeInternalError, "Internal server error", correlationID)
		logger.ErrorContext(ctx, "internal error",
			"correlation_id", correlationID,
			"operation", internalErr.Operation,
			"error", internalErr.Error())
		return http.StatusInternalServerError, errResp
	}

	// Default to internal server error
	errResp := NewErrorResponse(CodeUnknown, "An unexpected error occurred", correlationID)
	logger.ErrorContext(ctx, "unexpected error",
		"correlation_id", correlationID,
		"error", err.Error())
	return http.StatusInternalServerError, errResp
}

// HandleGRPCError converts application errors to gRPC status errors
func HandleGRPCError(ctx context.Context, err error, logger *slog.Logger) error {
	correlationID := getCorrelationID(ctx)

	switch {
	case errors.Is(err, ErrUserNotFound):
	case errors.Is(err, ErrResourceNotFound):
		logger.InfoContext(ctx, "grpc not found error",
			"correlation_id", correlationID,
			"error", err.Error())
		return status.Error(codes.NotFound, err.Error())

	case errors.Is(err, ErrInvalidCredentials):
		logger.InfoContext(ctx, "grpc unauthenticated error",
			"correlation_id", correlationID)
		return status.Error(codes.Unauthenticated, "Invalid credentials")

	case errors.Is(err, ErrPermissionDenied):
		logger.InfoContext(ctx, "grpc permission denied",
			"correlation_id", correlationID)
		return status.Error(codes.PermissionDenied, "Permission denied")

	case errors.Is(err, ErrInvalidInput):
	case errors.Is(err, ErrInvalidID):
		logger.InfoContext(ctx, "grpc invalid argument",
			"correlation_id", correlationID,
			"error", err.Error())
		return status.Error(codes.InvalidArgument, err.Error())

	case errors.Is(err, ErrAlreadyExists):
		logger.InfoContext(ctx, "grpc already exists",
			"correlation_id", correlationID,
			"error", err.Error())
		return status.Error(codes.AlreadyExists, err.Error())

	case errors.Is(err, ErrRateLimited):
		logger.WarnContext(ctx, "grpc rate limited",
			"correlation_id", correlationID)
		return status.Error(codes.ResourceExhausted, "Rate limit exceeded")

	case errors.Is(err, ErrCircuitOpen):
		logger.ErrorContext(ctx, "grpc circuit open",
			"correlation_id", correlationID)
		return status.Error(codes.Unavailable, "Service temporarily unavailable")
	}

	// Check custom error types
	var valErr *ValidationError
	if errors.As(err, &valErr) {
		logger.InfoContext(ctx, "grpc validation error",
			"correlation_id", correlationID,
			"field", valErr.Field)
		return status.Error(codes.InvalidArgument, valErr.Error())
	}

	var extErr *ExternalServiceError
	if errors.As(err, &extErr) {
		logger.ErrorContext(ctx, "grpc external service error",
			"correlation_id", correlationID,
			"service", extErr.Service)
		return status.Error(codes.Unavailable, "External service unavailable")
	}

	// Default to internal error
	logger.ErrorContext(ctx, "grpc internal error",
		"correlation_id", correlationID,
		"error", err.Error())
	return status.Error(codes.Internal, "Internal server error")
}

// getCorrelationID extracts correlation ID from context
func getCorrelationID(ctx context.Context) string {
	if id, ok := ctx.Value("correlation_id").(string); ok {
		return id
	}
	if id, ok := ctx.Value("request_id").(string); ok {
		return id
	}
	return "unknown"
}