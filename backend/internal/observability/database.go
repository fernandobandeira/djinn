package observability

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// DatabaseTracer wraps database operations with OpenTelemetry tracing
type DatabaseTracer struct {
	tracer trace.Tracer
}

// NewDatabaseTracer creates a new database tracer
func NewDatabaseTracer() *DatabaseTracer {
	return &DatabaseTracer{
		tracer: GetTracer("database"),
	}
}

// TraceQuery traces a database query
func (dt *DatabaseTracer) TraceQuery(ctx context.Context, query string, args []interface{}, fn func(context.Context) error) error {
	// Start span for database operation
	ctx, span := dt.tracer.Start(ctx, "db.query",
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			attribute.String("db.system", "postgresql"),
			attribute.String("db.statement", truncateQuery(query, 1000)),
			attribute.Int("db.args_count", len(args)),
		),
	)
	defer span.End()
	
	// Record timing
	start := time.Now()
	
	// Execute the query
	err := fn(ctx)
	
	// Record duration
	duration := time.Since(start)
	span.SetAttributes(
		attribute.Float64("db.query_duration_ms", float64(duration.Milliseconds())),
	)
	
	// Handle error
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		
		// Add additional context for specific errors
		switch err {
		case pgx.ErrNoRows:
			span.SetAttributes(attribute.Bool("db.no_rows", true))
		case sql.ErrNoRows:
			span.SetAttributes(attribute.Bool("db.no_rows", true))
		default:
			if pgErr, ok := err.(*pgconn.PgError); ok {
				span.SetAttributes(
					attribute.String("db.error_code", pgErr.Code),
					attribute.String("db.error_message", pgErr.Message),
					attribute.String("db.error_severity", pgErr.Severity),
				)
			}
		}
	} else {
		span.SetStatus(codes.Ok, "")
	}
	
	return err
}

// TraceTransaction traces a database transaction
func (dt *DatabaseTracer) TraceTransaction(ctx context.Context, fn func(context.Context) error) error {
	// Start span for transaction
	ctx, span := dt.tracer.Start(ctx, "db.transaction",
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			attribute.String("db.system", "postgresql"),
			attribute.String("db.operation", "transaction"),
		),
	)
	defer span.End()
	
	// Record timing
	start := time.Now()
	
	// Execute the transaction
	err := fn(ctx)
	
	// Record duration
	duration := time.Since(start)
	span.SetAttributes(
		attribute.Float64("db.transaction_duration_ms", float64(duration.Milliseconds())),
	)
	
	// Handle error
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetAttributes(attribute.Bool("db.transaction_rollback", true))
	} else {
		span.SetStatus(codes.Ok, "")
		span.SetAttributes(attribute.Bool("db.transaction_commit", true))
	}
	
	return err
}

// TraceBatch traces a batch database operation
func (dt *DatabaseTracer) TraceBatch(ctx context.Context, batchSize int, fn func(context.Context) error) error {
	// Start span for batch operation
	ctx, span := dt.tracer.Start(ctx, "db.batch",
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			attribute.String("db.system", "postgresql"),
			attribute.String("db.operation", "batch"),
			attribute.Int("db.batch_size", batchSize),
		),
	)
	defer span.End()
	
	// Record timing
	start := time.Now()
	
	// Execute the batch
	err := fn(ctx)
	
	// Record duration
	duration := time.Since(start)
	span.SetAttributes(
		attribute.Float64("db.batch_duration_ms", float64(duration.Milliseconds())),
	)
	
	// Handle error
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	} else {
		span.SetStatus(codes.Ok, "")
	}
	
	return err
}

// TraceConnectionPool traces connection pool metrics
func (dt *DatabaseTracer) TraceConnectionPool(ctx context.Context, stats *pgxpool.Stat) {
	if stats == nil {
		return
	}
	
	// Create a span for recording pool metrics
	_, span := dt.tracer.Start(ctx, "db.pool.stats",
		trace.WithSpanKind(trace.SpanKindInternal),
		trace.WithAttributes(
			attribute.Int64("db.pool.acquired_conns", int64(stats.AcquiredConns())),
			attribute.Int64("db.pool.idle_conns", int64(stats.IdleConns())),
			attribute.Int64("db.pool.total_conns", int64(stats.TotalConns())),
			attribute.Int64("db.pool.max_conns", int64(stats.MaxConns())),
			attribute.Int64("db.pool.acquire_count", stats.AcquireCount()),
			attribute.Float64("db.pool.acquire_duration_ms", 
				float64(stats.AcquireDuration().Milliseconds())),
		),
	)
	defer span.End()
}

// truncateQuery truncates a query string to a maximum length
func truncateQuery(query string, maxLength int) string {
	if len(query) <= maxLength {
		return query
	}
	return fmt.Sprintf("%s... (truncated)", query[:maxLength])
}