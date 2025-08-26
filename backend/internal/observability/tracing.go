package observability

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

// TracingConfig holds configuration for OpenTelemetry tracing
type TracingConfig struct {
	ServiceName    string
	ServiceVersion string
	Environment    string
	OTLPEndpoint   string
	Enabled        bool
	SampleRate     float64
}

// InitTracer initializes OpenTelemetry tracing
func InitTracer(ctx context.Context, config TracingConfig) (func(context.Context) error, error) {
	if !config.Enabled {
		// Return no-op shutdown function if tracing is disabled
		return func(ctx context.Context) error { return nil }, nil
	}

	// Create resource with service information
	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(config.ServiceName),
			semconv.ServiceVersion(config.ServiceVersion),
			attribute.String("environment", config.Environment),
			attribute.String("service.namespace", "djinn"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	// Create OTLP exporter
	var exporter *otlptrace.Exporter
	if config.OTLPEndpoint != "" {
		client := otlptracehttp.NewClient(
			otlptracehttp.WithEndpoint(config.OTLPEndpoint),
			otlptracehttp.WithInsecure(), // Use TLS in production
			otlptracehttp.WithTimeout(30*time.Second),
		)
		
		exporter, err = otlptrace.New(ctx, client)
		if err != nil {
			return nil, fmt.Errorf("failed to create OTLP exporter: %w", err)
		}
	}

	// Create sampler with configured rate
	sampler := sdktrace.TraceIDRatioBased(config.SampleRate)

	// Create trace provider
	var tp *sdktrace.TracerProvider
	if exporter != nil {
		tp = sdktrace.NewTracerProvider(
			sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(res),
			sdktrace.WithSampler(sampler),
		)
	} else {
		// No exporter configured, use a simple provider for local development
		tp = sdktrace.NewTracerProvider(
			sdktrace.WithResource(res),
			sdktrace.WithSampler(sampler),
		)
	}

	// Set global tracer provider
	otel.SetTracerProvider(tp)

	// Set global propagator for distributed tracing
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	// Return shutdown function
	return func(ctx context.Context) error {
		// Ensure all spans are flushed
		shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		
		if err := tp.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("failed to shutdown tracer provider: %w", err)
		}
		return nil
	}, nil
}

// GetTracer returns a tracer for the given component
func GetTracer(component string) trace.Tracer {
	return otel.Tracer(
		fmt.Sprintf("github.com/fernandobandeira/djinn/backend/%s", component),
		trace.WithSchemaURL(semconv.SchemaURL),
	)
}

// StartSpan starts a new span with common attributes
func StartSpan(ctx context.Context, tracer trace.Tracer, name string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return tracer.Start(ctx, name, opts...)
}

// AddSpanAttributes adds attributes to the current span
func AddSpanAttributes(ctx context.Context, attrs ...attribute.KeyValue) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attrs...)
}

// RecordError records an error on the current span
func RecordError(ctx context.Context, err error) {
	if err == nil {
		return
	}
	span := trace.SpanFromContext(ctx)
	span.RecordError(err)
}

// SetSpanStatus sets the status of the current span
func SetSpanStatus(ctx context.Context, code codes.Code, description string) {
	span := trace.SpanFromContext(ctx)
	span.SetStatus(code, description)
}