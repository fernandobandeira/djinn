package logging

import (
	"context"
	"log/slog"

	ctxutil "github.com/fernandobandeira/djinn/backend/internal/infrastructure/context"
)

// ServiceLogger wraps slog.Logger with service-specific metadata
type ServiceLogger struct {
	*slog.Logger
	service string
	version string
}

// NewServiceLogger creates a logger with service metadata
func NewServiceLogger(base *slog.Logger, service, version string) *ServiceLogger {
	logger := base.With(
		"service", service,
		"version", version,
	)
	
	return &ServiceLogger{
		Logger:  logger,
		service: service,
		version: version,
	}
}

// WithCorrelationID adds correlation ID to the logger context
func (l *ServiceLogger) WithCorrelationID(ctx context.Context) *ServiceLogger {
	correlationID := ctxutil.GetCorrelationID(ctx)
	if correlationID == "" {
		return l
	}
	
	return &ServiceLogger{
		Logger:  l.Logger.With("correlation_id", correlationID),
		service: l.service,
		version: l.version,
	}
}

// LoggerWithCorrelationID creates a logger with correlation ID from context
func LoggerWithCorrelationID(ctx context.Context, logger *slog.Logger) *slog.Logger {
	correlationID := ctxutil.GetCorrelationID(ctx)
	if correlationID == "" {
		return logger
	}
	return logger.With("correlation_id", correlationID)
}

// WithOperation adds operation context to the logger
func (l *ServiceLogger) WithOperation(operation string) *ServiceLogger {
	return &ServiceLogger{
		Logger:  l.Logger.With("operation", operation),
		service: l.service,
		version: l.version,
	}
}

// WithUserID adds user ID to the logger context
func (l *ServiceLogger) WithUserID(userID string) *ServiceLogger {
	return &ServiceLogger{
		Logger:  l.Logger.With("user_id", userID),
		service: l.service,
		version: l.version,
	}
}

// WithError adds error details to the logger context
func (l *ServiceLogger) WithError(err error) *ServiceLogger {
	if err == nil {
		return l
	}
	
	return &ServiceLogger{
		Logger:  l.Logger.With("error", err.Error(), "error_type", getErrorType(err)),
		service: l.service,
		version: l.version,
	}
}

// InfoContext logs at info level with context
func (l *ServiceLogger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.WithCorrelationID(ctx).Info(msg, args...)
}

// WarnContext logs at warn level with context
func (l *ServiceLogger) WarnContext(ctx context.Context, msg string, args ...any) {
	l.WithCorrelationID(ctx).Warn(msg, args...)
}

// ErrorContext logs at error level with context
func (l *ServiceLogger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.WithCorrelationID(ctx).Error(msg, args...)
}

// DebugContext logs at debug level with context
func (l *ServiceLogger) DebugContext(ctx context.Context, msg string, args ...any) {
	l.WithCorrelationID(ctx).Debug(msg, args...)
}


// getErrorType returns the type name of an error
func getErrorType(err error) string {
	if err == nil {
		return ""
	}
	
	// Try to get the underlying type name
	switch err.(type) {
	case interface{ Unwrap() error }:
		if unwrapped := err.(interface{ Unwrap() error }).Unwrap(); unwrapped != nil {
			return getErrorType(unwrapped)
		}
	}
	
	// Return the type name
	return err.Error()
}