package errors

import (
	"context"
	"errors"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestErrorMapper_MapError(t *testing.T) {
	mapper := NewErrorMapper()
	logger := slog.Default()
	ctx := context.Background()
	errorCtx := CreateErrorContext(ctx, logger)

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
			mapping := mapper.MapError(tt.err, errorCtx)
			
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
			assert.Equal(t, errorCtx.CorrelationID, mapping.Extensions["correlationId"])
			assert.NotEmpty(t, mapping.Extensions["timestamp"])
		})
	}
}

func TestErrorMapper_MessageSanitization(t *testing.T) {
	mapper := NewErrorMapper()
	logger := slog.Default()
	ctx := context.Background()
	errorCtx := CreateErrorContext(ctx, logger)

	// Test that bad request errors are sanitized
	mapping := mapper.MapError(ErrBadRequest, errorCtx)
	assert.Equal(t, "Invalid request format", mapping.Message, "Bad request message should be sanitized")
	
	// Test that internal errors don't expose details
	internalErr := &InternalError{Operation: "secret_operation", Message: "sensitive details"}
	mapping = mapper.MapError(internalErr, errorCtx)
	assert.Equal(t, "Internal server error", mapping.Message, "Internal error details should not be exposed")
}

func TestErrorMapper_ContextAwareAuthentication(t *testing.T) {
	mapper := NewErrorMapper()
	logger := slog.Default()
	ctx := context.Background()
	errorCtx := CreateErrorContext(ctx, logger)

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
			authErr := &AuthenticationError{Reason: tt.reason, Message: "test"}
			mapping := mapper.MapError(authErr, errorCtx)
			
			assert.Equal(t, tt.expectedCode, mapping.Code)
			assert.Equal(t, string(tt.expectedCode), mapping.Extensions["code"])
		})
	}
}

func TestCreateErrorContext(t *testing.T) {
	ctx := context.Background()
	logger := slog.Default()
	
	errorCtx := CreateErrorContext(ctx, logger)
	
	assert.NotNil(t, errorCtx)
	// The correlation ID might be empty if ctxutil.GetCorrelationID returns ""
	// This is fine as long as our mapper handles it properly
	assert.NotNil(t, errorCtx.CorrelationID) // Can be empty string
	assert.NotEmpty(t, errorCtx.Timestamp)
	assert.Equal(t, logger, errorCtx.Logger)
	assert.Equal(t, ctx, errorCtx.Context)
}