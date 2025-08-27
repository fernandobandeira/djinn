package errors

import (
	"time"
)

// ErrorCode represents standardized error codes
type ErrorCode string

const (
	// Validation errors (400)
	CodeValidationFailed ErrorCode = "VALIDATION_FAILED"
	CodeInvalidInput     ErrorCode = "INVALID_INPUT"
	CodeMissingField     ErrorCode = "MISSING_FIELD"

	// Authentication/Authorization (401/403)
	CodeUnauthorized      ErrorCode = "UNAUTHORIZED"
	CodeInvalidCredentials ErrorCode = "INVALID_CREDENTIALS"
	CodeTokenExpired      ErrorCode = "TOKEN_EXPIRED"
	CodePermissionDenied  ErrorCode = "PERMISSION_DENIED"

	// Not Found (404)
	CodeNotFound     ErrorCode = "NOT_FOUND"
	CodeUserNotFound ErrorCode = "USER_NOT_FOUND"

	// Conflict (409)
	CodeConflict       ErrorCode = "CONFLICT"
	CodeAlreadyExists  ErrorCode = "ALREADY_EXISTS"

	// Rate Limiting (429)
	CodeRateLimited ErrorCode = "RATE_LIMITED"

	// External Service (502/503)
	CodeExternalServiceError ErrorCode = "EXTERNAL_SERVICE_ERROR"
	CodeServiceUnavailable  ErrorCode = "SERVICE_UNAVAILABLE"

	// Internal (500)
	CodeInternalError ErrorCode = "INTERNAL_ERROR"
	CodeUnknown      ErrorCode = "UNKNOWN_ERROR"
)

// ErrorDetail represents a single error detail
type ErrorDetail struct {
	Field   string      `json:"field,omitempty"`
	Value   interface{} `json:"value,omitempty"`
	Message string      `json:"message"`
}

// ErrorResponse represents the standardized HTTP error response
type ErrorResponse struct {
	Error struct {
		Code          ErrorCode      `json:"code"`
		Message       string         `json:"message"`
		Details       []ErrorDetail  `json:"details,omitempty"`
		CorrelationID string         `json:"correlation_id"`
		Timestamp     time.Time      `json:"timestamp"`
		RetryAfter    *time.Duration `json:"retry_after,omitempty"`
	} `json:"error"`
}

// NewErrorResponse creates a new error response
func NewErrorResponse(code ErrorCode, message string, correlationID string) *ErrorResponse {
	resp := &ErrorResponse{}
	resp.Error.Code = code
	resp.Error.Message = message
	resp.Error.CorrelationID = correlationID
	resp.Error.Timestamp = time.Now().UTC()
	return resp
}

// AddDetail adds an error detail
func (r *ErrorResponse) AddDetail(field, message string, value interface{}) *ErrorResponse {
	r.Error.Details = append(r.Error.Details, ErrorDetail{
		Field:   field,
		Message: message,
		Value:   value,
	})
	return r
}

// SetRetryAfter sets the retry after duration
func (r *ErrorResponse) SetRetryAfter(duration time.Duration) *ErrorResponse {
	r.Error.RetryAfter = &duration
	return r
}