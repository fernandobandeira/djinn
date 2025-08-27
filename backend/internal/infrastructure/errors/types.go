package errors

import (
	"fmt"
	"time"
)

// ValidationError represents validation errors with field-specific details
type ValidationError struct {
	Field   string `json:"field"`
	Value   any    `json:"value,omitempty"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (e *ValidationError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("validation error on field %s: %s", e.Field, e.Message)
	}
	return fmt.Sprintf("validation error: %s", e.Message)
}

// NotFoundError represents resource not found errors
type NotFoundError struct {
	Resource string `json:"resource"`
	ID       string `json:"id,omitempty"`
	Message  string `json:"message"`
}

func (e *NotFoundError) Error() string {
	if e.ID != "" {
		return fmt.Sprintf("%s with ID %s not found", e.Resource, e.ID)
	}
	return fmt.Sprintf("%s not found: %s", e.Resource, e.Message)
}

// UnauthorizedError represents authentication/authorization failures
type UnauthorizedError struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("unauthorized: %s", e.Message)
}

// ExternalServiceError represents failures from external dependencies
type ExternalServiceError struct {
	Service      string        `json:"service"`
	Operation    string        `json:"operation"`
	StatusCode   int           `json:"status_code,omitempty"`
	Message      string        `json:"message"`
	RetryAfter   time.Duration `json:"retry_after,omitempty"`
	WrappedError error         `json:"-"`
}

func (e *ExternalServiceError) Error() string {
	msg := fmt.Sprintf("external service error [%s/%s]: %s", e.Service, e.Operation, e.Message)
	if e.StatusCode > 0 {
		msg = fmt.Sprintf("%s (status: %d)", msg, e.StatusCode)
	}
	return msg
}

func (e *ExternalServiceError) Unwrap() error {
	return e.WrappedError
}

// RateLimitError represents rate limiting errors
type RateLimitError struct {
	Resource   string        `json:"resource"`
	Limit      int           `json:"limit"`
	Window     time.Duration `json:"window"`
	RetryAfter time.Duration `json:"retry_after"`
	Message    string        `json:"message"`
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("rate limit exceeded for %s: %s (retry after %v)", e.Resource, e.Message, e.RetryAfter)
}

// ConflictError represents resource conflict errors
type ConflictError struct {
	Resource string `json:"resource"`
	Field    string `json:"field,omitempty"`
	Value    any    `json:"value,omitempty"`
	Message  string `json:"message"`
}

func (e *ConflictError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("conflict on %s.%s: %s", e.Resource, e.Field, e.Message)
	}
	return fmt.Sprintf("conflict on %s: %s", e.Resource, e.Message)
}

// InternalError represents internal server errors
type InternalError struct {
	Operation string `json:"operation"`
	Message   string `json:"message"`
	Wrapped   error  `json:"-"`
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("internal error during %s: %s", e.Operation, e.Message)
}

func (e *InternalError) Unwrap() error {
	return e.Wrapped
}

// AuthenticationError represents authentication failures
type AuthenticationError struct {
	Method  string `json:"method,omitempty"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

func (e *AuthenticationError) Error() string {
	if e.Method != "" {
		return fmt.Sprintf("authentication failed [%s]: %s", e.Method, e.Message)
	}
	return fmt.Sprintf("authentication failed: %s", e.Message)
}

// AuthorizationError represents authorization/permission failures
type AuthorizationError struct {
	Resource  string `json:"resource"`
	Action    string `json:"action"`
	Principal string `json:"principal,omitempty"`
	Message   string `json:"message"`
}

func (e *AuthorizationError) Error() string {
	if e.Principal != "" {
		return fmt.Sprintf("authorization denied for %s to %s %s: %s", e.Principal, e.Action, e.Resource, e.Message)
	}
	return fmt.Sprintf("authorization denied to %s %s: %s", e.Action, e.Resource, e.Message)
}