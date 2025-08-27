package errors

import (
	"context"
	"errors"

	ctxutil "github.com/fernandobandeira/djinn/backend/internal/infrastructure/context"
)

// GetErrorCode maps an error to its corresponding error code
func GetErrorCode(err error) ErrorCode {
	// Check for typed errors first
	var valErr *ValidationError
	if errors.As(err, &valErr) {
		return CodeValidationFailed
	}

	var notFoundErr *NotFoundError
	if errors.As(err, &notFoundErr) {
		return CodeNotFound
	}

	var unauthorizedErr *UnauthorizedError
	if errors.As(err, &unauthorizedErr) {
		return CodeUnauthorized
	}

	var conflictErr *ConflictError
	if errors.As(err, &conflictErr) {
		return CodeConflict
	}

	var rateLimitErr *RateLimitError
	if errors.As(err, &rateLimitErr) {
		return CodeRateLimited
	}

	var extErr *ExternalServiceError
	if errors.As(err, &extErr) {
		return CodeExternalServiceError
	}

	var internalErr *InternalError
	if errors.As(err, &internalErr) {
		return CodeInternalError
	}

	var authErr *AuthenticationError
	if errors.As(err, &authErr) {
		switch authErr.Reason {
		case "token_expired":
			return CodeTokenExpired
		case "token_invalid":
			return CodeInvalidCredentials
		case "insufficient_privileges":
			return CodePermissionDenied
		default:
			return CodeInvalidCredentials
		}
	}

	var authzErr *AuthorizationError
	if errors.As(err, &authzErr) {
		return CodePermissionDenied
	}

	// Check for sentinel errors
	switch {
	case errors.Is(err, ErrInvalidCredentials):
		return CodeInvalidCredentials
	case errors.Is(err, ErrTokenExpired):
		return CodeTokenExpired
	case errors.Is(err, ErrTokenInvalid):
		return CodeInvalidCredentials
	case errors.Is(err, ErrUnauthorized):
		return CodeUnauthorized
	case errors.Is(err, ErrPermissionDenied):
		return CodePermissionDenied
	case errors.Is(err, ErrRateLimited):
		return CodeRateLimited
	case errors.Is(err, ErrServiceUnavailable):
		return CodeServiceUnavailable
	case errors.Is(err, ErrBadRequest):
		return CodeInvalidInput
	case errors.Is(err, ErrRequestTooLarge):
		return CodeInvalidInput
	case errors.Is(err, ErrUnsupportedMediaType):
		return CodeInvalidInput
	case errors.Is(err, ErrTimeout):
		return CodeInternalError
	case errors.Is(err, ErrInternal):
		return CodeInternalError
	case errors.Is(err, ErrUserNotFound):
		return CodeUserNotFound
	case errors.Is(err, ErrResourceNotFound):
		return CodeNotFound
	case errors.Is(err, ErrAlreadyExists):
		return CodeAlreadyExists
	case errors.Is(err, ErrInvalidID):
		return CodeInvalidInput
	case errors.Is(err, ErrInvalidInput):
		return CodeInvalidInput
	case errors.Is(err, ErrMissingField):
		return CodeMissingField
	case errors.Is(err, ErrCircuitOpen):
		return CodeServiceUnavailable
	case errors.Is(err, ErrCircuitHalfOpen):
		return CodeServiceUnavailable
	case errors.Is(err, ErrConflict):
		return CodeConflict
	default:
		return CodeUnknown
	}
}

// getCorrelationID extracts correlation ID from context
func getCorrelationID(ctx context.Context) string {
	if id := ctxutil.GetCorrelationID(ctx); id != "" {
		return id
	}
	return "unknown"
}