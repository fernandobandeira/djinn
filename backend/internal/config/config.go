package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// Server configuration
	Port         int           `envconfig:"PORT" default:"8080"`
	ReadTimeout  time.Duration `envconfig:"READ_TIMEOUT" default:"15s"`
	WriteTimeout time.Duration `envconfig:"WRITE_TIMEOUT" default:"15s"`
	IdleTimeout  time.Duration `envconfig:"IDLE_TIMEOUT" default:"60s"`

	// Database configuration
	DatabaseURL  string `envconfig:"DATABASE_URL" required:"true"`
	PgBouncerURL string `envconfig:"PGBOUNCER_URL"`

	// Logging
	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`
	LogJSON  bool   `envconfig:"LOG_JSON" default:"false"`

	// GraphQL
	EnablePlayground bool `envconfig:"ENABLE_PLAYGROUND" default:"true"`
	QueryComplexity  int  `envconfig:"QUERY_COMPLEXITY" default:"1000"`

	// Monitoring
	OTELEndpoint     string `envconfig:"OTEL_EXPORTER_OTLP_ENDPOINT"`
	MetricsEnabled   bool   `envconfig:"METRICS_ENABLED" default:"true"`
	TracingEnabled   bool   `envconfig:"TRACING_ENABLED" default:"true"`
	ServiceName      string `envconfig:"SERVICE_NAME" default:"djinn-api"`
	ServiceVersion   string `envconfig:"SERVICE_VERSION" default:"1.0.0"`
	Environment      string `envconfig:"ENVIRONMENT" default:"development"`
}

func Load() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	// Use PgBouncer URL if available, otherwise fallback to direct database URL
	if cfg.PgBouncerURL == "" {
		cfg.PgBouncerURL = cfg.DatabaseURL
	}

	return &cfg, nil
}