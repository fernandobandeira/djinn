package errors

import "errors"

// Sentinel errors for infrastructure/system-level error conditions
// Domain-specific errors should be defined in their respective domain packages
var (
	// Authentication/Authorization errors (infrastructure concern)
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrTokenExpired       = errors.New("token expired")
	ErrTokenInvalid       = errors.New("token invalid")
	ErrPermissionDenied   = errors.New("permission denied")
	ErrUnauthorized       = errors.New("unauthorized")

	// System/Infrastructure errors
	ErrRateLimited      = errors.New("rate limit exceeded")
	ErrTimeout          = errors.New("operation timeout")
	ErrInternal         = errors.New("internal server error")
	ErrServiceUnavailable = errors.New("service unavailable")
	
	// Request processing errors
	ErrBadRequest       = errors.New("bad request")
	ErrRequestTooLarge  = errors.New("request too large")
	ErrUnsupportedMediaType = errors.New("unsupported media type")
)