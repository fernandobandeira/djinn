package errors

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
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

func TestAuthenticationError(t *testing.T) {
	tests := []struct {
		name     string
		err      *AuthenticationError
		expected string
	}{
		{
			name: "with method",
			err: &AuthenticationError{
				Method:  "Bearer",
				Reason:  "token_expired",
				Message: "JWT token has expired",
			},
			expected: "authentication failed [Bearer]: JWT token has expired",
		},
		{
			name: "without method",
			err: &AuthenticationError{
				Reason:  "invalid_credentials",
				Message: "Invalid username or password",
			},
			expected: "authentication failed: Invalid username or password",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.err.Error())
		})
	}
}

func TestAuthorizationError(t *testing.T) {
	tests := []struct {
		name     string
		err      *AuthorizationError
		expected string
	}{
		{
			name: "with principal",
			err: &AuthorizationError{
				Resource:  "user",
				Action:    "delete",
				Principal: "user123",
				Message:   "insufficient permissions",
			},
			expected: "authorization denied for user123 to delete user: insufficient permissions",
		},
		{
			name: "without principal",
			err: &AuthorizationError{
				Resource: "admin",
				Action:   "access",
				Message:  "admin access required",
			},
			expected: "authorization denied to access admin: admin access required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.err.Error())
		})
	}
}

func TestMapError_ErrorCodeMapping(t *testing.T) {
	logger := slog.Default()
	ctx := context.Background()

	tests := []struct {
		name     string
		err      error
		expected ErrorCode
	}{
		{
			name:     "ValidationError",
			err:      &ValidationError{Field: "email", Message: "invalid"},
			expected: CodeValidationFailed,
		},
		{
			name:     "NotFoundError",
			err:      &NotFoundError{Resource: "user", ID: "123"},
			expected: CodeNotFound,
		},
		{
			name:     "AuthenticationError with token_expired",
			err:      &AuthenticationError{Reason: "token_expired", Message: "token expired"},
			expected: CodeTokenExpired,
		},
		{
			name:     "AuthenticationError with token_invalid",
			err:      &AuthenticationError{Reason: "token_invalid", Message: "invalid token"},
			expected: CodeInvalidCredentials,
		},
		{
			name:     "AuthenticationError with insufficient_privileges",
			err:      &AuthenticationError{Reason: "insufficient_privileges", Message: "not enough privileges"},
			expected: CodePermissionDenied,
		},
		{
			name:     "AuthenticationError with default reason",
			err:      &AuthenticationError{Reason: "other", Message: "other error"},
			expected: CodeInvalidCredentials,
		},
		{
			name:     "AuthorizationError",
			err:      &AuthorizationError{Resource: "admin", Action: "access", Message: "denied"},
			expected: CodePermissionDenied,
		},
		{
			name:     "ErrInvalidCredentials sentinel",
			err:      ErrInvalidCredentials,
			expected: CodeInvalidCredentials,
		},
		{
			name:     "ErrTokenExpired sentinel",
			err:      ErrTokenExpired,
			expected: CodeTokenExpired,
		},
		{
			name:     "ErrUserNotFound sentinel",
			err:      ErrUserNotFound,
			expected: CodeUserNotFound,
		},
		{
			name:     "Unknown error",
			err:      errors.New("unknown error"),
			expected: CodeUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapping := MapError(tt.err, ctx, logger)
			assert.Equal(t, tt.expected, mapping.Code)
		})
	}
}