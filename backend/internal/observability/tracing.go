package observability

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
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

// InitTracer initializes OpenTelemetry tracing.
// This sets up the global tracer provider that all auto-instrumentation libraries use.
// The contrib packages (otelhttp, otelpgx, otelgqlgen) will automatically use this provider.
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
		// Configure client options based on environment
		clientOpts := []otlptracehttp.Option{
			otlptracehttp.WithEndpoint(config.OTLPEndpoint),
			otlptracehttp.WithTimeout(30 * time.Second),
		}
		
		// Use insecure only in development
		if config.Environment == "development" || config.Environment == "local" {
			clientOpts = append(clientOpts, otlptracehttp.WithInsecure())
		}
		// In production, TLS is used by default
		
		client := otlptracehttp.NewClient(clientOpts...)
		
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