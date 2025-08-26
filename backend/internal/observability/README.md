# Observability Package

## Overview
This package provides comprehensive observability capabilities for the Djinn backend using OpenTelemetry for distributed tracing and metrics collection.

## Features

### Distributed Tracing
- **OpenTelemetry Integration**: Full OpenTelemetry support with OTLP export
- **HTTP Request Tracing**: Automatic tracing for all HTTP requests
- **GraphQL Operation Tracing**: Specialized tracing for GraphQL queries/mutations
- **Database Query Tracing**: Detailed tracing of database operations
- **Trace Context Propagation**: Support for distributed trace context

### Configuration

#### Environment Variables
```bash
# Enable/disable tracing (default: true)
TRACING_ENABLED=true

# OpenTelemetry collector endpoint
OTEL_EXPORTER_OTLP_ENDPOINT=localhost:4317

# Sampling rate (0.0 to 1.0, default: 1.0 for 100% sampling)
TRACING_SAMPLE_RATE=1.0

# Service identification
SERVICE_NAME=djinn-api
SERVICE_VERSION=1.0.0
ENVIRONMENT=development
```

### Usage

#### HTTP Middleware
The tracing middleware is automatically added to all HTTP requests when tracing is enabled:

```go
if s.config.TracingEnabled {
    s.router.Use(middleware.TracingMiddleware())
    s.router.Use(middleware.TracingGraphQLMiddleware())
}
```

#### Database Operations
Database operations are automatically traced when using the DatabaseTracer:

```go
tracer := observability.NewDatabaseTracer()
err := tracer.TraceQuery(ctx, query, args, func(ctx context.Context) error {
    // Execute database query
    return db.Query(ctx, query, args...)
})
```

#### Custom Spans
Create custom spans for specific operations:

```go
tracer := observability.GetTracer("my-component")
ctx, span := observability.StartSpan(ctx, tracer, "operation-name")
defer span.End()

// Add attributes
observability.AddSpanAttributes(ctx,
    attribute.String("key", "value"),
    attribute.Int("count", 42),
)

// Record errors
if err != nil {
    observability.RecordError(ctx, err)
}
```

### Trace Attributes

#### HTTP Traces
- `http.method`: HTTP method (GET, POST, etc.)
- `http.target`: Full request URL
- `http.route`: URL path
- `http.status_code`: Response status code
- `http.request_content_length`: Request body size
- `http.response_content_length`: Response body size
- `http.request.duration_ms`: Request duration in milliseconds
- `request.id`: Request correlation ID

#### Database Traces
- `db.system`: Database system (postgresql)
- `db.statement`: SQL query (truncated to 1000 chars)
- `db.args_count`: Number of query arguments
- `db.query_duration_ms`: Query execution time
- `db.error_code`: PostgreSQL error code (on error)
- `db.transaction_commit/rollback`: Transaction outcome

#### GraphQL Traces
- `graphql.endpoint`: GraphQL endpoint path
- `graphql.transport`: Transport type (HTTP POST)

### Performance Considerations

#### Sampling
In production, consider reducing the sampling rate to control overhead:
- Development: `TRACING_SAMPLE_RATE=1.0` (100% sampling)
- Staging: `TRACING_SAMPLE_RATE=0.1` (10% sampling)
- Production: `TRACING_SAMPLE_RATE=0.01` (1% sampling)

#### Batch Export
Traces are batched before export to reduce network overhead.

### Integration with Monitoring Stack

#### Local Development
For local development, traces can be viewed using:
1. **Jaeger**: `docker run -p 16686:16686 -p 4317:4317 jaegertracing/all-in-one`
2. **Grafana Tempo**: Part of the Grafana stack

#### Production
Traces are exported to the configured OTLP endpoint which can be:
- **Grafana Cloud**: Managed Tempo instance
- **Self-hosted**: Tempo or Jaeger on Hetzner VPS
- **PostHog**: Integrated with user analytics

### Troubleshooting

#### No Traces Appearing
1. Check `TRACING_ENABLED=true`
2. Verify OTLP endpoint is accessible
3. Check application logs for tracing errors
4. Ensure sampling rate > 0

#### High Memory Usage
1. Reduce sampling rate
2. Check for trace span leaks (spans not ended)
3. Verify batch export is working

#### Missing Trace Context
1. Ensure TracingMiddleware is registered
2. Check trace propagation headers are not filtered
3. Verify context is passed correctly