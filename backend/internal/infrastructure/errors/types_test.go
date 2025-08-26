package errors

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidationError(t *testing.T) {
	tests := []struct {
		name     string
		err      *ValidationError
		expected string
	}{
		{
			name: "with field",
			err: &ValidationError{
				Field:   "email",
				Value:   "invalid",
				Message: "must be a valid email",
				Code:    "INVALID_EMAIL",
			},
			expected: "validation error on field email: must be a valid email",
		},
		{
			name: "without field",
			err: &ValidationError{
				Message: "validation failed",
				Code:    "VALIDATION_FAILED",
			},
			expected: "validation error: validation failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.err.Error())
		})
	}
}

func TestNotFoundError(t *testing.T) {
	tests := []struct {
		name     string
		err      *NotFoundError
		expected string
	}{
		{
			name: "with ID",
			err: &NotFoundError{
				Resource: "User",
				ID:       "123",
				Message:  "not found",
			},
			expected: "User with ID 123 not found",
		},
		{
			name: "without ID",
			err: &NotFoundError{
				Resource: "Account",
				Message:  "no active account",
			},
			expected: "Account not found: no active account",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.err.Error())
		})
	}
}

func TestUnauthorizedError(t *testing.T) {
	err := &UnauthorizedError{
		Reason:  "token_expired",
		Message: "authentication token has expired",
	}
	
	assert.Equal(t, "unauthorized: authentication token has expired", err.Error())
}

func TestExternalServiceError(t *testing.T) {
	baseErr := errors.New("connection refused")
	
	tests := []struct {
		name     string
		err      *ExternalServiceError
		expected string
		unwrap   error
	}{
		{
			name: "with status code",
			err: &ExternalServiceError{
				Service:    "PaymentAPI",
				Operation:  "charge",
				StatusCode: 503,
				Message:    "service unavailable",
			},
			expected: "external service error [PaymentAPI/charge]: service unavailable (status: 503)",
			unwrap:   nil,
		},
		{
			name: "without status code",
			err: &ExternalServiceError{
				Service:   "EmailService",
				Operation: "send",
				Message:   "connection timeout",
			},
			expected: "external service error [EmailService/send]: connection timeout",
			unwrap:   nil,
		},
		{
			name: "with wrapped error",
			err: &ExternalServiceError{
				Service:      "Database",
				Operation:    "query",
				Message:      "query failed",
				WrappedError: baseErr,
			},
			expected: "external service error [Database/query]: query failed",
			unwrap:   baseErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.err.Error())
			if tt.unwrap != nil {
				assert.Equal(t, tt.unwrap, tt.err.Unwrap())
			}
		})
	}
}

func TestRateLimitError(t *testing.T) {
	err := &RateLimitError{
		Resource:   "API",
		Limit:      100,
		Window:     time.Minute,
		RetryAfter: 30 * time.Second,
		Message:    "too many requests",
	}
	
	expected := fmt.Sprintf("rate limit exceeded for API: too many requests (retry after %v)", 30*time.Second)
	assert.Equal(t, expected, err.Error())
}

func TestConflictError(t *testing.T) {
	tests := []struct {
		name     string
		err      *ConflictError
		expected string
	}{
		{
			name: "with field",
			err: &ConflictError{
				Resource: "User",
				Field:    "email",
				Value:    "test@example.com",
				Message:  "already exists",
			},
			expected: "conflict on User.email: already exists",
		},
		{
			name: "without field",
			err: &ConflictError{
				Resource: "Transaction",
				Message:  "duplicate transaction ID",
			},
			expected: "conflict on Transaction: duplicate transaction ID",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.err.Error())
		})
	}
}

func TestInternalError(t *testing.T) {
	baseErr := errors.New("database connection failed")
	
	err := &InternalError{
		Operation: "SaveUser",
		Message:   "failed to save user",
		Wrapped:   baseErr,
	}
	
	assert.Equal(t, "internal error during SaveUser: failed to save user", err.Error())
	assert.Equal(t, baseErr, err.Unwrap())
}

func TestSentinelErrors(t *testing.T) {
	// Test that sentinel errors are properly defined
	sentinelErrors := []error{
		ErrInvalidCredentials,
		ErrTokenExpired,
		ErrTokenInvalid,
		ErrPermissionDenied,
		ErrUserNotFound,
		ErrResourceNotFound,
		ErrAlreadyExists,
		ErrInvalidID,
		ErrInvalidInput,
		ErrMissingField,
		ErrCircuitOpen,
		ErrCircuitHalfOpen,
		ErrRateLimited,
		ErrInternal,
		ErrBadRequest,
		ErrConflict,
		ErrTimeout,
	}
	
	for _, err := range sentinelErrors {
		require.NotNil(t, err)
		require.NotEmpty(t, err.Error())
	}
	
	// Test error comparison
	assert.True(t, errors.Is(ErrUserNotFound, ErrUserNotFound))
	assert.False(t, errors.Is(ErrUserNotFound, ErrResourceNotFound))
}