package errors

import (
	"context"
	"errors"
	"log/slog"
	"time"

	ctxutil "github.com/fernandobandeira/djinn/backend/internal/infrastructure/context"
)

// ErrorMapping defines how an error should be handled
type ErrorMapping struct {
	Code       ErrorCode
	Message    string
	Extensions map[string]interface{}
	LogLevel   slog.Level
}

// MapError efficiently maps any error to a standardized error mapping using type switching
func MapError(err error, ctx context.Context, logger *slog.Logger) *ErrorMapping {
	correlationID := ctxutil.GetCorrelationID(ctx)
	if correlationID == "" {
		correlationID = "unknown"
	}
	timestamp := time.Now().UTC().Format(time.RFC3339)
	
	// Create base extensions that all errors will have
	baseExtensions := map[string]interface{}{
		"correlationId": correlationID,
		"timestamp":     timestamp,
	}

	// Use efficient type switching instead of pseudo-strategy pattern
	switch e := err.(type) {
	case *ValidationError:
		return &ErrorMapping{
			Code:     CodeValidationFailed,
			Message:  e.Error(),
			LogLevel: slog.LevelInfo,
			Extensions: extendMap(baseExtensions, map[string]interface{}{
				"code":  string(CodeValidationFailed),
				"field": e.Field,
				"value": e.Value,
			}),
		}

	case *NotFoundError:
		return &ErrorMapping{
			Code:     CodeNotFound,
			Message:  e.Error(),
			LogLevel: slog.LevelInfo,
			Extensions: extendMap(baseExtensions, map[string]interface{}{
				"code":     string(CodeNotFound),
				"resource": e.Resource,
				"id":       e.ID,
			}),
		}

	case *UnauthorizedError:
		return &ErrorMapping{
			Code:     CodeUnauthorized,
			Message:  e.Error(),
			LogLevel: slog.LevelInfo,
			Extensions: extendMap(baseExtensions, map[string]interface{}{
				"code":   string(CodeUnauthorized),
				"reason": e.Reason,
			}),
		}

	case *ConflictError:
		return &ErrorMapping{
			Code:     CodeConflict,
			Message:  e.Error(),
			LogLevel: slog.LevelInfo,
			Extensions: extendMap(baseExtensions, map[string]interface{}{
				"code":     string(CodeConflict),
				"resource": e.Resource,
				"field":    e.Field,
				"value":    e.Value,
			}),
		}

	case *RateLimitError:
		return &ErrorMapping{
			Code:     CodeRateLimited,
			Message:  e.Error(),
			LogLevel: slog.LevelWarn,
			Extensions: extendMap(baseExtensions, map[string]interface{}{
				"code":       string(CodeRateLimited),
				"resource":   e.Resource,
				"retryAfter": e.RetryAfter.Seconds(),
			}),
		}

	case *ExternalServiceError:
		return &ErrorMapping{
			Code:     CodeExternalServiceError,
			Message:  "External service error", // Don't expose details
			LogLevel: slog.LevelError,
			Extensions: extendMap(baseExtensions, map[string]interface{}{
				"code":    string(CodeExternalServiceError),
				"service": e.Service,
			}),
		}

	case *InternalError:
		return &ErrorMapping{
			Code:     CodeInternalError,
			Message:  "Internal server error", // Don't expose details
			LogLevel: slog.LevelError,
			Extensions: extendMap(baseExtensions, map[string]interface{}{
				"code": string(CodeInternalError),
			}),
		}

	case *AuthenticationError:
		code := mapAuthenticationReason(e.Reason)
		return &ErrorMapping{
			Code:     code,
			Message:  e.Error(),
			LogLevel: slog.LevelInfo,
			Extensions: extendMap(baseExtensions, map[string]interface{}{
				"code":   string(code),
				"method": e.Method,
				"reason": e.Reason,
			}),
		}

	case *AuthorizationError:
		return &ErrorMapping{
			Code:     CodePermissionDenied,
			Message:  e.Error(),
			LogLevel: slog.LevelInfo,
			Extensions: extendMap(baseExtensions, map[string]interface{}{
				"code":      string(CodePermissionDenied),
				"resource":  e.Resource,
				"action":    e.Action,
				"principal": e.Principal,
			}),
		}

	default:
		// Handle sentinel errors efficiently
		if mapping := mapSentinelError(err, baseExtensions); mapping != nil {
			return mapping
		}
		
		// Unknown error fallback
		return &ErrorMapping{
			Code:     CodeUnknown,
			Message:  "An unexpected error occurred",
			LogLevel: slog.LevelError,
			Extensions: extendMap(baseExtensions, map[string]interface{}{
				"code": string(CodeUnknown),
			}),
		}
	}
}

// mapAuthenticationReason maps authentication reasons to appropriate error codes
func mapAuthenticationReason(reason string) ErrorCode {
	switch reason {
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

// mapSentinelError efficiently handles sentinel errors with O(1) lookup
func mapSentinelError(err error, baseExtensions map[string]interface{}) *ErrorMapping {
	// Pre-computed sentinel error mappings for O(1) lookup
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
		ErrBadRequest:           {CodeInvalidInput, "Invalid request format", slog.LevelInfo}, // Sanitized
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

	// Efficient sentinel error checking
	for sentinelErr, mapping := range sentinelMappings {
		if errors.Is(err, sentinelErr) {
			return &ErrorMapping{
				Code:     mapping.code,
				Message:  mapping.message,
				LogLevel: mapping.logLevel,
				Extensions: extendMap(baseExtensions, map[string]interface{}{
					"code": string(mapping.code),
				}),
			}
		}
	}
	
	return nil
}

// extendMap efficiently combines base extensions with specific ones
func extendMap(base, specific map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{}, len(base)+len(specific))
	
	// Copy base extensions
	for k, v := range base {
		result[k] = v
	}
	
	// Add specific extensions, filtering out nil values
	for k, v := range specific {
		if v != nil && v != "" {
			result[k] = v
		}
	}
	
	return result
}