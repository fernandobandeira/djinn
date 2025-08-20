# ADR-20250120: Monitoring and Observability Strategy

## Status
Accepted

## Context
Djinn requires a comprehensive monitoring and observability strategy that balances:
- Cost optimization for MVP (targeting €3.79-20/month infrastructure)
- VPS deployment on Hetzner Cloud requiring infrastructure monitoring
- Mobile-first architecture with Flutter frontend
- PostgreSQL database monitoring needs
- User behavior analytics and error tracking
- Compliance requirements for financial data
- Small team operations (5-8 people)
- Ability to scale monitoring as user base grows

### Current Infrastructure Context
Based on ADR-20250819 (Deployment Architecture):
- **Primary Platform**: Hetzner VPS (CX22: 2 vCPU, 4GB RAM, 40GB SSD)
- **Backend**: Go API server with PostgreSQL
- **Mobile App**: Flutter with Firebase Auth
- **Feature Flags**: PostHog (already planned)
- **Budget**: €3.79-20/month for MVP phase

### Monitoring Requirements
1. **Infrastructure Monitoring**
   - Server health (CPU, memory, disk, network)
   - Database performance and connections
   - API response times and throughput
   - Service availability (99.9% uptime target)

2. **Application Monitoring**
   - Error tracking and debugging
   - API endpoint performance
   - Background job processing
   - Receipt processing pipeline

3. **User Analytics**
   - User behavior and journey tracking
   - Feature adoption metrics
   - Session recordings for debugging
   - Conversion funnel analysis

4. **Security & Compliance**
   - Authentication attempts and failures
   - Data access audit logging
   - Suspicious activity detection
   - PCI DSS compliance readiness

## Decision

### Hybrid Monitoring Architecture
Adopt a hybrid approach leveraging PostHog's strengths in user analytics while implementing lightweight infrastructure monitoring for VPS operations, comprehensive crash reporting, and user behavior tracking.

#### 1. PostHog for User-Centric Monitoring (Free Tier)

```yaml
posthog_configuration:
  purpose: "User analytics and frontend monitoring"
  cost: "$0/month (free tier)"
  
  capabilities_used:
    analytics:
      limit: "1M events/month"
      use_cases:
        - User behavior tracking
        - Feature adoption metrics
        - Conversion funnel analysis
        - Custom event tracking
        
    session_replay:
      limit: "5K recordings/month"
      use_cases:
        - Bug reproduction
        - UX optimization
        - Customer support
        - Error context
        
    error_tracking:
      limit: "100K exceptions/month"
      use_cases:
        - Frontend JavaScript errors
        - Mobile app crashes
        - API error patterns
        - Stack trace analysis
        
    feature_flags:
      limit: "1M requests/month"
      use_cases:
        - Progressive rollouts
        - A/B testing
        - Kill switches
        - User segmentation
        
  implementation:
    frontend: "PostHog Flutter SDK"
    backend: "PostHog Go SDK for custom events"
    deployment: "PostHog Cloud (SaaS)"
```

#### 2. Infrastructure Monitoring Stack (Self-Hosted)

```yaml
infrastructure_monitoring:
  cost: "€0/month (self-hosted on existing VPS)"
  
  components:
    metrics_collection:
      tool: "Prometheus"
      deployment: "Native binary with systemd"
      memory_usage: "~50MB"
      storage: "/var/lib/prometheus (~500MB for 7 days retention)"
      
    visualization:
      tool: "Grafana"
      deployment: "Native binary with systemd"
      memory_usage: "~100MB"
      storage: "/var/lib/grafana (database and dashboards)"
      access: "Behind Nginx with basic auth"
      
    node_metrics:
      tool: "Node Exporter"
      deployment: "Native binary with systemd"
      memory_usage: "~10MB"
      
    postgres_metrics:
      tool: "postgres_exporter"
      deployment: "Native binary with systemd"
      memory_usage: "~20MB"
      
  monitoring_targets:
    - System metrics (CPU, memory, disk, network)
    - PostgreSQL performance and connections
    - Go application custom metrics
    - Nginx request metrics
    - Service health via systemd
```

#### 3. Application Performance Monitoring

```yaml
application_monitoring:
  approach: "Custom instrumentation with OpenTelemetry"
  
  go_implementation:
    metrics:
      - Request latency histograms
      - Request counter by endpoint
      - Active database connections
      - Background job queue depth
      - Receipt processing duration
      
    traces:
      storage: "/var/log/djinn/traces (JSON files with rotation)"
      sampling: "1% of requests"
      retention: "24 hours"
      
    logs:
      format: "Structured JSON with slog"
      storage: "/var/log/djinn/app.log"
      rotation: "Daily with 7-day retention via logrotate"
      level: "INFO in production"
```

#### 4. Analytics & User Behavior Tracking

```yaml
analytics_strategy:
  purpose: "Deep understanding of user behavior and feature adoption"
  
  user_journey_tracking:
    implementation:
      - Session recording with PostHog
      - Custom event tracking for key actions
      - Funnel analysis for conversion flows
      - Cohort analysis for user segments
    
    key_events:
      authentication:
        - "user_signup_started"
        - "user_signup_completed"
        - "user_login"
        - "user_logout"
        - "password_reset_requested"
      
      receipt_processing:
        - "receipt_upload_started"
        - "receipt_captured"
        - "receipt_processed"
        - "receipt_categorized"
        - "receipt_manual_correction"
      
      financial_insights:
        - "dashboard_viewed"
        - "report_generated"
        - "budget_created"
        - "goal_set"
        - "alert_configured"
      
      engagement:
        - "feature_discovered"
        - "feature_first_use"
        - "feature_repeated_use"
        - "help_accessed"
        - "settings_changed"
  
  behavioral_metrics:
    user_engagement:
      - Daily Active Users (DAU)
      - Weekly Active Users (WAU)
      - Monthly Active Users (MAU)
      - Session duration
      - Sessions per user
      - Feature adoption rate
    
    business_metrics:
      - Receipt processing rate
      - Categorization accuracy
      - Time to first value
      - User retention (D1, D7, D30)
      - Churn rate
      - Customer lifetime value (CLV)
    
    performance_metrics:
      - Page load times
      - API response times
      - Error rates by feature
      - Crash-free rate
      - Network failure rate

  custom_properties:
    user_attributes:
      - account_type: "free|premium|trial"
      - signup_source: "organic|referral|campaign"
      - device_type: "ios|android|web"
      - app_version: "semantic_version"
      - locale: "language_code"
    
    session_context:
      - network_type: "wifi|cellular|offline"
      - battery_level: "percentage"
      - device_memory: "MB"
      - screen_resolution: "width x height"
```

#### 5. Crash Reporting & Diagnostics

```yaml
crash_reporting:
  purpose: "Comprehensive crash detection, analysis, and recovery"
  
  mobile_crash_handling:
    flutter_implementation:
      tools:
        - "Flutter built-in error handling"
        - "PostHog error tracking integration"
        - "Custom crash reporter with slog patterns"
      
      crash_capture:
        fatal_errors:
          - Flutter framework crashes
          - Dart unhandled exceptions
          - Native platform crashes (iOS/Android)
        
        non_fatal_errors:
          - API request failures (ref: error-handling-ADR)
          - Widget rendering errors
          - Navigation errors
          - State management errors
    
    crash_context:
      device_info:
        - OS version
        - Device model
        - Available memory
        - Storage space
        - Battery state
      
      app_state:
        - Current screen/route
        - User actions before crash
        - Network connectivity
        - Background/foreground state
        - Feature flags active
      
      user_context:
        - User ID (anonymized)
        - Session ID
        - Account type
        - App version
        - Locale settings
  
  backend_crash_handling:
    go_implementation:
      # Integration with error-handling-ADR patterns
      panic_recovery:
        - Structured panic logging with slog
        - Graceful degradation
        - Circuit breaker activation
        - Correlation ID tracking
      
      error_aggregation:
        - Group similar errors
        - Track error frequency
        - Identify error patterns
        - Alert on error spikes
  
  diagnostic_data:
    breadcrumbs:
      - User interactions
      - Network requests
      - Navigation events
      - System events
      - Custom checkpoints
    
    memory_diagnostics:
      - Memory usage trends
      - Memory leak detection
      - Large allocation tracking
      - GC pressure monitoring
    
    performance_diagnostics:
      - Slow frame rendering
      - Jank detection
      - Network latency
      - Database query times
      - Background job duration
  
  crash_recovery:
    client_side:
      - Automatic crash report submission
      - Offline crash queue
      - Retry with exponential backoff
      - User feedback collection
    
    server_side:
      - Crash report processing pipeline
      - Automatic issue creation
      - Duplicate detection
      - Severity classification
      - Team notification
```

#### 6. Alerting Strategy

```yaml
alerting:
  critical_alerts:
    delivery: "Email + Telegram bot"
    conditions:
      - "Server CPU > 90% for 5 minutes"
      - "Memory usage > 85%"
      - "Disk usage > 80%"
      - "Database connections > 80"
      - "API error rate > 5%"
      - "Service down for > 1 minute"
      
  monitoring_alerts:
    delivery: "PostHog dashboard + daily summary"
    conditions:
      - "Daily active users drop > 20%"
      - "Feature flag errors"
      - "Session replay quota approaching"
      
  implementation:
    uptime: "Uptime Robot (free tier)"
    custom: "Prometheus AlertManager"
    notifications: "Simple Go webhook handler"
```

### Implementation Architecture

#### Native Installation Script for Monitoring Stack

```bash
#!/bin/bash
# install-monitoring.sh - Install monitoring stack without Docker

set -e

# Create monitoring user
sudo useradd --no-create-home --shell /bin/false prometheus || true
sudo useradd --no-create-home --shell /bin/false node_exporter || true

# Create directories
sudo mkdir -p /etc/prometheus /var/lib/prometheus /var/lib/grafana
sudo chown prometheus:prometheus /etc/prometheus /var/lib/prometheus

# Install Prometheus
PROM_VERSION="2.48.0"
wget https://github.com/prometheus/prometheus/releases/download/v${PROM_VERSION}/prometheus-${PROM_VERSION}.linux-amd64.tar.gz
tar xvf prometheus-${PROM_VERSION}.linux-amd64.tar.gz
sudo cp prometheus-${PROM_VERSION}.linux-amd64/prometheus /usr/local/bin/
sudo cp prometheus-${PROM_VERSION}.linux-amd64/promtool /usr/local/bin/
sudo chown prometheus:prometheus /usr/local/bin/prometheus /usr/local/bin/promtool
rm -rf prometheus-${PROM_VERSION}.linux-amd64*

# Install Node Exporter
NODE_VERSION="1.7.0"
wget https://github.com/prometheus/node_exporter/releases/download/v${NODE_VERSION}/node_exporter-${NODE_VERSION}.linux-amd64.tar.gz
tar xvf node_exporter-${NODE_VERSION}.linux-amd64.tar.gz
sudo cp node_exporter-${NODE_VERSION}.linux-amd64/node_exporter /usr/local/bin/
sudo chown node_exporter:node_exporter /usr/local/bin/node_exporter
rm -rf node_exporter-${NODE_VERSION}.linux-amd64*

# Install Postgres Exporter
PG_VERSION="0.15.0"
wget https://github.com/prometheus-community/postgres_exporter/releases/download/v${PG_VERSION}/postgres_exporter-${PG_VERSION}.linux-amd64.tar.gz
tar xvf postgres_exporter-${PG_VERSION}.linux-amd64.tar.gz
sudo cp postgres_exporter-${PG_VERSION}.linux-amd64/postgres_exporter /usr/local/bin/
rm -rf postgres_exporter-${PG_VERSION}.linux-amd64*

# Install Grafana
sudo apt-get install -y software-properties-common
sudo add-apt-repository "deb https://packages.grafana.com/oss/deb stable main"
wget -q -O - https://packages.grafana.com/gpg.key | sudo apt-key add -
sudo apt-get update
sudo apt-get install -y grafana

# Configure systemd services
cat <<EOF | sudo tee /etc/systemd/system/prometheus.service
[Unit]
Description=Prometheus
Wants=network-online.target
After=network-online.target

[Service]
User=prometheus
Group=prometheus
Type=simple
ExecStart=/usr/local/bin/prometheus \
    --config.file /etc/prometheus/prometheus.yml \
    --storage.tsdb.path /var/lib/prometheus/ \
    --storage.tsdb.retention.time=7d \
    --storage.tsdb.retention.size=500MB \
    --web.console.templates=/etc/prometheus/consoles \
    --web.console.libraries=/etc/prometheus/console_libraries

[Install]
WantedBy=multi-user.target
EOF

cat <<EOF | sudo tee /etc/systemd/system/node_exporter.service
[Unit]
Description=Node Exporter
Wants=network-online.target
After=network-online.target

[Service]
User=node_exporter
Group=node_exporter
Type=simple
ExecStart=/usr/local/bin/node_exporter

[Install]
WantedBy=multi-user.target
EOF

cat <<EOF | sudo tee /etc/systemd/system/postgres_exporter.service
[Unit]
Description=Postgres Exporter
Wants=network-online.target
After=network-online.target

[Service]
User=prometheus
Group=prometheus
Type=simple
Environment="DATA_SOURCE_NAME=postgresql://monitoring:password@localhost:5432/djinn_production?sslmode=disable"
ExecStart=/usr/local/bin/postgres_exporter

[Install]
WantedBy=multi-user.target
EOF

# Enable and start services
sudo systemctl daemon-reload
sudo systemctl enable prometheus node_exporter postgres_exporter grafana-server
sudo systemctl start prometheus node_exporter postgres_exporter grafana-server

echo "Monitoring stack installed successfully!"
echo "Prometheus: http://localhost:9090"
echo "Grafana: http://localhost:3000"
echo "Configure Nginx reverse proxy for external access"
```

#### Go Application Instrumentation

```go
package monitoring

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "github.com/posthog/posthog-go"
    "log/slog"
    "net/http"
    "time"
)

var (
    // Prometheus metrics
    httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
        Name: "djinn_http_duration_seconds",
        Help: "Duration of HTTP requests.",
    }, []string{"method", "route", "status"})
    
    httpRequests = promauto.NewCounterVec(prometheus.CounterOpts{
        Name: "djinn_http_requests_total",
        Help: "Total number of HTTP requests.",
    }, []string{"method", "route", "status"})
    
    dbConnections = promauto.NewGauge(prometheus.GaugeOpts{
        Name: "djinn_db_connections_active",
        Help: "Number of active database connections.",
    })
    
    receiptProcessing = promauto.NewHistogram(prometheus.HistogramOpts{
        Name: "djinn_receipt_processing_seconds",
        Help: "Time taken to process receipts.",
        Buckets: []float64{0.1, 0.5, 1, 2, 5, 10},
    })
    
    // PostHog client
    posthogClient posthog.Client
)

func InitMonitoring(posthogAPIKey string) {
    // Initialize PostHog
    posthogClient = posthog.New(posthogAPIKey)
    
    // Expose Prometheus metrics endpoint
    http.Handle("/metrics", promhttp.Handler())
    
    // Start metrics server
    go func() {
        if err := http.ListenAndServe(":9091", nil); err != nil {
            slog.Error("Failed to start metrics server", "error", err)
        }
    }()
}

// HTTPMiddleware tracks HTTP metrics
func HTTPMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // Wrap response writer to capture status
        wrapped := &responseWriter{ResponseWriter: w, statusCode: 200}
        
        // Process request
        next.ServeHTTP(wrapped, r)
        
        // Record metrics
        duration := time.Since(start).Seconds()
        status := fmt.Sprintf("%d", wrapped.statusCode)
        
        httpDuration.WithLabelValues(r.Method, r.URL.Path, status).Observe(duration)
        httpRequests.WithLabelValues(r.Method, r.URL.Path, status).Inc()
        
        // Track in PostHog for user analytics
        if userID := r.Context().Value("userID"); userID != nil {
            posthogClient.Capture(posthog.Capture{
                DistinctId: userID.(string),
                Event: "api_request",
                Properties: posthog.NewProperties().
                    Set("method", r.Method).
                    Set("path", r.URL.Path).
                    Set("status", wrapped.statusCode).
                    Set("duration_ms", duration*1000),
            })
        }
        
        // Log slow requests
        if duration > 1.0 {
            slog.Warn("Slow request detected",
                "method", r.Method,
                "path", r.URL.Path,
                "duration", duration,
                "status", wrapped.statusCode,
            )
        }
    })
}

// TrackReceiptProcessing tracks receipt processing metrics
func TrackReceiptProcessing(userID string, duration time.Duration, success bool) {
    receiptProcessing.Observe(duration.Seconds())
    
    // Track in PostHog
    posthogClient.Capture(posthog.Capture{
        DistinctId: userID,
        Event: "receipt_processed",
        Properties: posthog.NewProperties().
            Set("duration_ms", duration.Milliseconds()).
            Set("success", success),
    })
    
    // Log if processing was slow
    if duration > 5*time.Second {
        slog.Warn("Slow receipt processing",
            "user_id", userID,
            "duration", duration,
            "success", success,
        )
    }
}

// TrackError tracks errors in both Prometheus and PostHog
func TrackError(userID, errorType, errorMessage string, context map[string]interface{}) {
    // Log error
    slog.Error("Application error",
        "user_id", userID,
        "error_type", errorType,
        "error_message", errorMessage,
        "context", context,
    )
    
    // Track in PostHog for analysis
    props := posthog.NewProperties().
        Set("error_type", errorType).
        Set("error_message", errorMessage)
    
    for k, v := range context {
        props.Set(k, v)
    }
    
    posthogClient.Capture(posthog.Capture{
        DistinctId: userID,
        Event: "error_occurred",
        Properties: props,
    })
}
```

#### Prometheus Configuration

```yaml
# prometheus.yml
global:
  scrape_interval: 15s
  evaluation_interval: 15s

alerting:
  alertmanagers:
    - static_configs:
        - targets: []

rule_files:
  - "alerts.yml"

scrape_configs:
  - job_name: 'prometheus'
    static_configs:
      - targets: ['localhost:9090']

  - job_name: 'node'
    static_configs:
      - targets: ['localhost:9100']

  - job_name: 'postgres'
    static_configs:
      - targets: ['localhost:9187']

  - job_name: 'djinn-api'
    static_configs:
      - targets: ['localhost:9091']
    metrics_path: '/metrics'

  - job_name: 'nginx'
    static_configs:
      - targets: ['localhost:9113']
```

#### Alert Rules Configuration

```yaml
# alerts.yml
groups:
  - name: djinn_alerts
    interval: 30s
    rules:
      - alert: HighCPUUsage
        expr: 100 - (avg(irate(node_cpu_seconds_total{mode="idle"}[5m])) * 100) > 90
        for: 5m
        labels:
          severity: critical
        annotations:
          summary: "High CPU usage detected"
          description: "CPU usage is above 90% (current value: {{ $value }}%)"

      - alert: HighMemoryUsage
        expr: (1 - (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes)) * 100 > 85
        for: 5m
        labels:
          severity: warning
        annotations:
          summary: "High memory usage detected"
          description: "Memory usage is above 85% (current value: {{ $value }}%)"

      - alert: DiskSpaceRunningOut
        expr: (node_filesystem_avail_bytes{mountpoint="/"} / node_filesystem_size_bytes{mountpoint="/"}) * 100 < 20
        for: 5m
        labels:
          severity: critical
        annotations:
          summary: "Disk space running out"
          description: "Less than 20% disk space remaining"

      - alert: DatabaseConnectionsHigh
        expr: pg_stat_database_numbackends{datname="djinn_production"} > 80
        for: 2m
        labels:
          severity: warning
        annotations:
          summary: "High number of database connections"
          description: "Database connections exceed 80 (current: {{ $value }})"

      - alert: APIErrorRateHigh
        expr: rate(djinn_http_requests_total{status=~"5.."}[5m]) > 0.05
        for: 5m
        labels:
          severity: critical
        annotations:
          summary: "High API error rate"
          description: "API error rate exceeds 5% (current: {{ $value }})"

      - alert: APIResponseSlow
        expr: histogram_quantile(0.95, rate(djinn_http_duration_seconds_bucket[5m])) > 1
        for: 5m
        labels:
          severity: warning
        annotations:
          summary: "Slow API response times"
          description: "95th percentile response time exceeds 1 second"
```

### Monitoring Dashboard Strategy

#### Key Dashboards to Implement

1. **System Overview Dashboard**
   - Server resource utilization
   - Service health status
   - Error rates and response times
   - Active users and sessions

2. **API Performance Dashboard**
   - Request rate by endpoint
   - Response time distribution
   - Error rates by endpoint
   - Slow query log

3. **Database Dashboard**
   - Connection pool status
   - Query performance
   - Table sizes and growth
   - Replication lag (when applicable)

4. **User Analytics Dashboard (PostHog)**
   - Daily/Weekly/Monthly active users
   - Feature adoption rates
   - User journey funnels
   - Session replay insights

5. **Business Metrics Dashboard**
   - Receipt processing volume
   - Transaction categorization accuracy
   - User engagement metrics
   - Feature flag rollout status

### Cost Analysis

```yaml
cost_breakdown:
  mvp_phase:
    posthog: "$0/month (free tier)"
    infrastructure_monitoring: "$0/month (self-hosted)"
    uptime_monitoring: "$0/month (Uptime Robot free)"
    additional_resources:
      cpu_overhead: "~5% of VPS capacity"
      memory_overhead: "~200MB RAM"
      disk_overhead: "~500MB storage"
    total: "$0/month additional cost"
    
  growth_phase_10k_users:
    posthog: "$0-450/month (depends on usage)"
    grafana_cloud: "$50/month (optional)"
    sentry: "$26/month (optional)"
    datadog: "$0 (stay self-hosted)"
    total: "$0-526/month"
    
  scale_phase_50k_users:
    posthog: "$450-2000/month"
    monitoring_infrastructure: "$200-500/month"
    incident_management: "$100/month"
    total: "$750-2600/month"
    
  recommendation: "Stay on free tier through MVP, evaluate paid tools at 10K users"
```

### Migration Path

```yaml
monitoring_evolution:
  mvp_0_1k_users:
    stack: "PostHog free + self-hosted Prometheus/Grafana"
    focus: "Basic monitoring and user analytics"
    
  early_growth_1k_10k:
    additions:
      - "Sentry for error tracking (if needed)"
      - "More sophisticated alerting"
      - "Custom dashboards"
    
  validated_10k_50k:
    evaluate:
      - "PostHog paid tier for advanced features"
      - "Grafana Cloud for managed monitoring"
      - "Dedicated APM solution"
    
  scale_50k_plus:
    enterprise_options:
      - "Datadog or New Relic for full APM"
      - "PostHog Enterprise for compliance"
      - "Dedicated observability team"
```

## Consequences

### Positive
- **Zero Additional Cost**: Entire monitoring stack runs on existing VPS infrastructure
- **No Docker Overhead**: Native binaries use less resources than containers
- **Comprehensive Coverage**: Combines user analytics with infrastructure monitoring
- **Scalable Architecture**: Can grow from free tier to enterprise solutions
- **Developer Friendly**: Simple instrumentation with familiar tools
- **Privacy Compliant**: Self-hosted infrastructure monitoring keeps data in-house
- **Quick Implementation**: Can be deployed in 1-2 days
- **Low Resource Usage**: ~200MB total RAM for monitoring stack

### Negative
- **Manual Setup Required**: Self-hosted monitoring needs configuration
- **Limited APM Features**: No distributed tracing or advanced APM
- **Maintenance Overhead**: Team must maintain monitoring infrastructure
- **Basic Alerting**: Simple alerting compared to enterprise solutions
- **No Machine Learning**: Missing anomaly detection and predictive alerts

### Risks
- **Resource Contention**: Monitoring stack could impact application performance
- **Data Loss**: Self-hosted metrics lost if VPS fails
- **Alert Fatigue**: Poor alert tuning could overwhelm team
- **Scaling Challenges**: Self-hosted solution may not scale smoothly
- **PostHog Limits**: Could hit free tier limits with rapid growth

### Mitigation Strategies
1. **Resource Monitoring**: Set strict memory limits for monitoring containers
2. **Backup Strategy**: Regular metric data exports to object storage
3. **Alert Tuning**: Start with few critical alerts, expand gradually
4. **Capacity Planning**: Monitor PostHog usage, plan for paid tier
5. **Documentation**: Maintain runbooks for common monitoring scenarios

## Alternatives Considered

### Option A: PostHog Only
- **Description**: Use PostHog for all monitoring needs
- **Pros**: Single tool, no additional setup, generous free tier
- **Cons**: No infrastructure monitoring, no APM capabilities
- **Verdict**: Insufficient for VPS deployment monitoring

### Option B: Full Datadog Stack
- **Description**: Datadog for infrastructure, APM, and logs
- **Pros**: Comprehensive solution, excellent features
- **Cons**: Expensive ($31/host/month minimum), overkill for MVP
- **Verdict**: Too expensive for current budget

### Option C: Elastic Stack (ELK)
- **Description**: Elasticsearch, Logstash, Kibana for observability
- **Pros**: Powerful search and analytics, good visualization
- **Cons**: Resource intensive (needs 2GB+ RAM), complex setup
- **Verdict**: Too heavy for small VPS

### Option D: Cloud Provider Monitoring
- **Description**: Use Hetzner Cloud monitoring features
- **Pros**: Integrated with infrastructure, simple setup
- **Cons**: Very basic metrics, no application monitoring
- **Verdict**: Insufficient for production needs

## Implementation Checklist

### Week 1: Foundation
- [ ] Deploy Prometheus and Grafana containers
- [ ] Configure node_exporter for system metrics
- [ ] Set up PostHog account and get API keys
- [ ] Implement basic Go instrumentation

### Week 2: Integration
- [ ] Add PostHog SDK to Flutter app
- [ ] Configure postgres_exporter
- [ ] Create initial Grafana dashboards
- [ ] Set up Prometheus alerting rules

### Week 3: Refinement
- [ ] Tune alert thresholds
- [ ] Implement custom business metrics
- [ ] Configure log rotation
- [ ] Set up Uptime Robot for external monitoring

### Week 4: Documentation
- [ ] Create monitoring runbook
- [ ] Document alert response procedures
- [ ] Train team on dashboard usage
- [ ] Establish on-call rotation (if needed)

## References
- [PostHog Documentation](https://posthog.com/docs)
- [Prometheus Best Practices](https://prometheus.io/docs/practices/)
- [Grafana VPS Deployment Guide](https://grafana.com/docs/grafana/latest/setup-grafana/)
- [OpenTelemetry Go SDK](https://opentelemetry.io/docs/languages/go/)
- ADR-20250819: Deployment Architecture
- ADR-20250812: Personal Finance Tech Stack Selection

## Decision Makers
- Author: Ana (Business Analyst)
- Technical Review: [Pending]
- Approval: [Pending]
- Date: 2025-01-20

## Revision History
- 2025-01-20: Initial comprehensive draft with PostHog analysis and hybrid monitoring approach