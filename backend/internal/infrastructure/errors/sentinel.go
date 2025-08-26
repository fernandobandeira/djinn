package errors

import "errors"

// Sentinel errors for common error conditions
var (
	// Authentication/Authorization errors
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrTokenExpired       = errors.New("token expired")
	ErrTokenInvalid       = errors.New("token invalid")
	ErrPermissionDenied   = errors.New("permission denied")

	// Resource errors
	ErrUserNotFound     = errors.New("user not found")
	ErrResourceNotFound = errors.New("resource not found")
	ErrAlreadyExists    = errors.New("resource already exists")

	// Validation errors
	ErrInvalidID    = errors.New("invalid ID format")
	ErrInvalidInput = errors.New("invalid input")
	ErrMissingField = errors.New("required field missing")

	// Circuit breaker errors
	ErrCircuitOpen     = errors.New("circuit breaker is open")
	ErrCircuitHalfOpen = errors.New("circuit breaker is half-open")

	// Rate limiting errors
	ErrRateLimited = errors.New("rate limit exceeded")

	// Generic errors
	ErrInternal   = errors.New("internal server error")
	ErrBadRequest = errors.New("bad request")
	ErrConflict   = errors.New("resource conflict")
	ErrTimeout    = errors.New("operation timeout")
)