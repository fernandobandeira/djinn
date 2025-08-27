package errors

import (
	"context"
	"errors"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMapError(t *testing.T) {
	logger := slog.Default()
	ctx := context.Background()

	tests := []struct {
		name           string
		err            error
		expectedCode   ErrorCode
		expectedLevel  slog.Level
		hasExtensions  []string
	}{
		{
			name:          "ValidationError",
			err:           &ValidationError{Field: "email", Message: "invalid email"},
			expectedCode:  CodeValidationFailed,
			expectedLevel: slog.LevelInfo,
			hasExtensions: []string{"code", "field", "correlationId", "timestamp"},
		},
		{
			name:          "NotFoundError",
			err:           &NotFoundError{Resource: "user", ID: "123"},
			expectedCode:  CodeNotFound,
			expectedLevel: slog.LevelInfo,
			hasExtensions: []string{"code", "resource", "id", "correlationId", "timestamp"},
		},
		{
			name:          "AuthenticationError - token_expired",
			err:           &AuthenticationError{Reason: "token_expired", Message: "token expired"},
			expectedCode:  CodeTokenExpired,
			expectedLevel: slog.LevelInfo,
			hasExtensions: []string{"code", "reason", "correlationId", "timestamp"},
		},
		{
			name:          "AuthenticationError - token_invalid",
			err:           &AuthenticationError{Reason: "token_invalid", Message: "invalid token"},
			expectedCode:  CodeInvalidCredentials,
			expectedLevel: slog.LevelInfo,
			hasExtensions: []string{"code", "reason", "correlationId", "timestamp"},
		},
		{
			name:          "AuthorizationError",
			err:           &AuthorizationError{Resource: "admin", Action: "access", Principal: "user123"},
			expectedCode:  CodePermissionDenied,
			expectedLevel: slog.LevelInfo,
			hasExtensions: []string{"code", "resource", "action", "principal", "correlationId", "timestamp"},
		},
		{
			name:          "ErrInvalidCredentials sentinel",
			err:           ErrInvalidCredentials,
			expectedCode:  CodeInvalidCredentials,
			expectedLevel: slog.LevelInfo,
			hasExtensions: []string{"code", "correlationId", "timestamp"},
		},
		{
			name:          "ErrBadRequest sentinel - sanitized message",
			err:           ErrBadRequest,
			expectedCode:  CodeInvalidInput,
			expectedLevel: slog.LevelInfo,
			hasExtensions: []string{"code", "correlationId", "timestamp"},
		},
		{
			name:          "Unknown error",
			err:           errors.New("unknown error"),
			expectedCode:  CodeUnknown,
			expectedLevel: slog.LevelError,
			hasExtensions: []string{"correlationId", "timestamp"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapping := MapError(tt.err, ctx, logger)
			
			require.NotNil(t, mapping)
			assert.Equal(t, tt.expectedCode, mapping.Code)
			assert.Equal(t, tt.expectedLevel, mapping.LogLevel)
			assert.NotEmpty(t, mapping.Message)
			assert.NotNil(t, mapping.Extensions)
			
			// Check expected extensions exist
			for _, ext := range tt.hasExtensions {
				assert.Contains(t, mapping.Extensions, ext, "Missing extension: %s", ext)
			}
			
			// Check standard extensions
			assert.Contains(t, mapping.Extensions, "correlationId")
			assert.NotEmpty(t, mapping.Extensions["timestamp"])
		})
	}
}

func TestMapError_MessageSanitization(t *testing.T) {
	logger := slog.Default()
	ctx := context.Background()

	// Test that bad request errors are sanitized
	mapping := MapError(ErrBadRequest, ctx, logger)
	assert.Equal(t, "Invalid request format", mapping.Message, "Bad request message should be sanitized")
	
	// Test that internal errors don't expose details
	internalErr := &InternalError{Operation: "secret_operation", Message: "sensitive details"}
	mapping = MapError(internalErr, ctx, logger)
	assert.Equal(t, "Internal server error", mapping.Message, "Internal error details should not be exposed")
}

func TestMapAuthenticationReason(t *testing.T) {
	tests := []struct {
		reason       string
		expectedCode ErrorCode
	}{
		{"token_expired", CodeTokenExpired},
		{"token_invalid", CodeInvalidCredentials},
		{"insufficient_privileges", CodePermissionDenied},
		{"other", CodeInvalidCredentials}, // default case
	}

	for _, tt := range tests {
		t.Run(tt.reason, func(t *testing.T) {
			code := mapAuthenticationReason(tt.reason)
			assert.Equal(t, tt.expectedCode, code)
		})
	}
}

func TestExtendMap(t *testing.T) {
	base := map[string]interface{}{
		"correlationId": "test-123",
		"timestamp":     "2023-01-01T00:00:00Z",
	}
	
	specific := map[string]interface{}{
		"code":  "VALIDATION_FAILED",
		"field": "email",
		"empty": "",    // Should be filtered out
		"nil":   nil,   // Should be filtered out
		"zero":  0,     // Should be included
	}
	
	result := extendMap(base, specific)
	
	// Check base extensions preserved
	assert.Equal(t, "test-123", result["correlationId"])
	assert.Equal(t, "2023-01-01T00:00:00Z", result["timestamp"])
	
	// Check specific extensions included
	assert.Equal(t, "VALIDATION_FAILED", result["code"])
	assert.Equal(t, "email", result["field"])
	assert.Equal(t, 0, result["zero"])
	
	// Check filtered out values
	assert.NotContains(t, result, "empty")
	assert.NotContains(t, result, "nil")
}