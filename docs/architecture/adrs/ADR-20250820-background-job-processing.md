# ADR-20250820: Background Job Processing with Self-Hosted Temporal

## Status
Accepted

## Context
Djinn requires background job processing for various asynchronous operations that should not block the main API response cycle. These include:
- CSV file imports and data processing
- OCR metadata processing from mobile clients
- Transaction categorization and enrichment
- Bank account synchronization via Plaid
- Export file generation
- Notification dispatch
- Data cleanup and maintenance tasks
- Analytics aggregation
- Offline sync queue processing

### Requirements
- Reliable job execution with retry capabilities
- Job prioritization and scheduling
- Monitoring and observability
- Error handling and compensation logic
- Idempotency guarantees
- Support for long-running operations
- Complex multi-step workflows with user interactions
- Workflow versioning for zero-downtime deployments
- Visibility into workflow state and history

### Constraints
- Small team without dedicated DevOps
- Need production-grade reliability from day one
- Must handle financial data with consistency guarantees
- Mobile offline-first architecture requires sync queue processing
- User confirmation workflows for imports
- Limited budget (€3.79/month Hetzner VPS)

## Decision

We will use **self-hosted Temporal on our Hetzner VPS** for all background job processing, providing a unified, reliable, and observable system for both simple and complex workflows at zero additional cost.

### Why Self-Hosted Temporal

```yaml
justification:
  cost_efficiency:
    - Zero additional monthly cost
    - Runs on existing Hetzner VPS
    - 4GB RAM is sufficient for MVP scale
    
  reliability:
    - Guaranteed execution with automatic retries
    - Durable execution state survives crashes
    - Built-in compensation and rollback
    
  developer_experience:
    - Workflow as code (type-safe Go)
    - Excellent testing framework
    - Time travel debugging
    - Workflow versioning
    
  operations:
    - Built-in UI for monitoring
    - Comprehensive metrics
    - No separate queue infrastructure
    
  control:
    - Full control over configuration
    - No vendor lock-in
    - Can tune for our specific needs
```

### Resource Planning for Hetzner CX22 (4GB RAM)

```yaml
resource_allocation:
  total_ram: 4GB
  
  allocation:
    postgresql: 1GB      # Database for app + Temporal
    temporal_server: 600MB
    temporal_ui: 200MB
    go_api: 800MB
    go_worker: 400MB
    minio: 400MB
    clamav: 300MB
    redis: 100MB        # For rate limiting
    system: 200MB
    
  optimization_strategies:
    - Use single PostgreSQL for both app and Temporal
    - Run Temporal with minimal history retention (7 days)
    - Single worker process with controlled concurrency
    - Disable Temporal metrics service (use built-in UI only)
```

### Temporal Deployment with Docker Compose

```yaml
# docker-compose.yml for complete Djinn stack
version: '3.8'

services:
  # PostgreSQL - shared by app and Temporal
  postgres:
    image: postgres:16-alpine
    container_name: djinn-postgres
    environment:
      POSTGRES_USER: ${DB_USER}
      POSTGRES_PASSWORD: ${DB_PASSWORD}
      POSTGRES_MULTIPLE_DATABASES: djinn,temporal
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./scripts/create-multiple-dbs.sh:/docker-entrypoint-initdb.d/create-multiple-dbs.sh
    ports:
      - "127.0.0.1:5432:5432"
    restart: unless-stopped
    mem_limit: 1g
    
  # Temporal Server (minimal setup)
  temporal:
    image: temporalio/auto-setup:1.22.4
    container_name: djinn-temporal
    depends_on:
      - postgres
    environment:
      - DB=postgresql
      - DB_PORT=5432
      - POSTGRES_USER=${DB_USER}
      - POSTGRES_PWD=${DB_PASSWORD}
      - POSTGRES_SEEDS=postgres
      - DYNAMIC_CONFIG_FILE_PATH=config/dynamicconfig.yaml
      - ENABLE_ES=false  # Disable Elasticsearch
      - PROMETHEUS_ENDPOINT=0.0.0.0:9090
      - NUM_HISTORY_SHARDS=1  # Minimal for MVP
      - SERVICES=frontend,matching,history,worker  # Core services only
    volumes:
      - ./temporal/dynamicconfig.yaml:/etc/temporal/config/dynamicconfig.yaml
    ports:
      - "127.0.0.1:7233:7233"  # Temporal gRPC
    restart: unless-stopped
    mem_limit: 600m
    
  # Temporal Web UI
  temporal-ui:
    image: temporalio/ui:2.21.3
    container_name: djinn-temporal-ui
    depends_on:
      - temporal
    environment:
      - TEMPORAL_ADDRESS=temporal:7233
      - TEMPORAL_CORS_ORIGINS=http://localhost:3000
    ports:
      - "127.0.0.1:8080:8080"
    restart: unless-stopped
    mem_limit: 200m
    
  # MinIO for file storage
  minio:
    image: minio/minio:latest
    container_name: djinn-minio
    command: server /data --console-address ":9001"
    environment:
      MINIO_ROOT_USER: ${MINIO_ROOT_USER}
      MINIO_ROOT_PASSWORD: ${MINIO_ROOT_PASSWORD}
      MINIO_BROWSER: "off"
    volumes:
      - minio_data:/data
    ports:
      - "127.0.0.1:9000:9000"
    restart: unless-stopped
    mem_limit: 400m
    
  # ClamAV for malware scanning
  clamav:
    image: clamav/clamav:stable-alpine
    container_name: djinn-clamav
    volumes:
      - clamav_data:/var/lib/clamav
    restart: unless-stopped
    mem_limit: 300m
    
  # Redis for rate limiting
  redis:
    image: redis:7-alpine
    container_name: djinn-redis
    command: redis-server --maxmemory 100mb --maxmemory-policy allkeys-lru
    ports:
      - "127.0.0.1:6379:6379"
    restart: unless-stopped
    mem_limit: 100m
    
  # Djinn API
  djinn-api:
    build:
      context: .
      dockerfile: Dockerfile.api
    container_name: djinn-api
    depends_on:
      - postgres
      - redis
      - temporal
    environment:
      DATABASE_URL: postgresql://${DB_USER}:${DB_PASSWORD}@postgres:5432/djinn
      REDIS_URL: redis://redis:6379
      TEMPORAL_ADDRESS: temporal:7233
      MINIO_ENDPOINT: minio:9000
    ports:
      - "127.0.0.1:4000:4000"
    restart: unless-stopped
    mem_limit: 800m
    
  # Djinn Worker
  djinn-worker:
    build:
      context: .
      dockerfile: Dockerfile.worker
    container_name: djinn-worker
    depends_on:
      - postgres
      - temporal
      - minio
    environment:
      DATABASE_URL: postgresql://${DB_USER}:${DB_PASSWORD}@postgres:5432/djinn
      TEMPORAL_ADDRESS: temporal:7233
      MINIO_ENDPOINT: minio:9000
      WORKER_CONCURRENCY: 5  # Limit concurrent workflows
    restart: unless-stopped
    mem_limit: 400m

volumes:
  postgres_data:
  minio_data:
  clamav_data:

networks:
  default:
    name: djinn-network
```

### Temporal Dynamic Configuration (Optimized for Low Memory)

```yaml
# temporal/dynamicconfig.yaml
system.forceSearchAttributesCacheRefreshOnRead:
  - value: false

history.persistenceMaxQPS:
  - value: 100  # Limit QPS to reduce load

history.persistenceGlobalMaxQPS:
  - value: 100

history.historyMgrNumConns:
  - value: 2  # Minimal connections

history.defaultActivityRetryPolicy:
  - value:
      initialIntervalInSeconds: 1
      maximumIntervalInSeconds: 100
      backoffCoefficient: 2.0
      maximumAttempts: 3

history.defaultWorkflowExecutionTimeout:
  - value: "24h"

history.maximumHistoryPageSize:
  - value: 500  # Reduce page size

# Retention - keep history for 7 days only
history.historyArchival:
  - value: false  # Disable archival for MVP

system.advancedVisibilityWritingMode:
  - value: "off"  # Disable advanced visibility to save resources

# Worker settings
worker.taskQueueActivitiesPerSecond:
  - value: 10  # Limit activity rate

worker.pollerCount:
  - value: 2  # Minimal pollers

# Disable unused features
system.enableActivityEagerExecution:
  - value: false

system.enableParentClosePolicyWorker:
  - value: false
```

### Simplified Temporal Client Setup

```go
package temporal

import (
    "fmt"
    
    "go.temporal.io/sdk/client"
    "go.temporal.io/sdk/worker"
)

type Config struct {
    Address    string // localhost:7233 for self-hosted
    Namespace  string // "default" for self-hosted
    WorkerName string
}

func NewClient(cfg Config) (client.Client, error) {
    // Simple connection for self-hosted Temporal (no TLS needed internally)
    return client.Dial(client.Options{
        HostPort:  cfg.Address,
        Namespace: cfg.Namespace,
    })
}

func NewWorker(c client.Client, taskQueue string) worker.Worker {
    w := worker.New(c, taskQueue, worker.Options{
        MaxConcurrentActivityExecutionSize: 5,     // Reduced for low memory
        MaxConcurrentWorkflowTaskExecutionSize: 5, // Reduced for low memory
        MaxConcurrentLocalActivityExecutionSize: 2,
        WorkerLocalActivitiesPerSecond: 10,
    })
    
    // Register workflows
    w.RegisterWorkflow(CSVImportWorkflow)
    w.RegisterWorkflow(TransactionEnrichmentWorkflow)
    w.RegisterWorkflow(ExportGenerationWorkflow)
    w.RegisterWorkflow(PlaidSyncWorkflow)
    w.RegisterWorkflow(SyncQueueWorkflow)
    
    // Register activities
    w.RegisterActivity(&FileActivities{})
    w.RegisterActivity(&DatabaseActivities{})
    w.RegisterActivity(&EnrichmentActivities{})
    
    return w
}
```

### Database Schema Setup Script

```bash
#!/bin/bash
# scripts/create-multiple-dbs.sh
# Creates both djinn and temporal databases in PostgreSQL

set -e
set -u

function create_database() {
    local database=$1
    echo "Creating database '$database'"
    psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
        CREATE DATABASE $database;
        GRANT ALL PRIVILEGES ON DATABASE $database TO $POSTGRES_USER;
EOSQL
}

create_database "djinn"
create_database "temporal"

# Create temporal schema
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" -d temporal <<-EOSQL
    CREATE SCHEMA IF NOT EXISTS temporal;
    CREATE SCHEMA IF NOT EXISTS temporal_visibility;
EOSQL
```

### Core Workflows (Same as before, but with resource optimizations)

#### CSV Import Workflow (with memory considerations)
```go
func CSVImportWorkflow(ctx workflow.Context, input CSVImportInput) (*CSVImportResult, error) {
    logger := workflow.GetLogger(ctx)
    logger.Info("Starting CSV import", "userID", input.UserID, "file", input.FileURL)
    
    // Configure activity options with resource limits
    activityOptions := workflow.ActivityOptions{
        StartToCloseTimeout: 5 * time.Minute,
        HeartbeatTimeout: 30 * time.Second,  // Add heartbeat for long operations
        RetryPolicy: &temporal.RetryPolicy{
            InitialInterval: time.Second,
            BackoffCoefficient: 2.0,
            MaximumInterval: time.Minute,
            MaximumAttempts: 3,
        },
    }
    ctx = workflow.WithActivityOptions(ctx, activityOptions)
    
    result := &CSVImportResult{
        ImportJobID: input.ImportJobID,
        Status:     "processing",
    }
    
    // Process in chunks for large files to avoid memory issues
    const chunkSize = 1000
    
    // Download and parse file
    var fileData FileData
    err := workflow.ExecuteActivity(ctx, DownloadFileActivity, input.FileURL).Get(ctx, &fileData)
    if err != nil {
        return result, err
    }
    
    // Parse file in chunks
    var allTransactions []Transaction
    for offset := 0; offset < fileData.RowCount; offset += chunkSize {
        var chunk []Transaction
        err := workflow.ExecuteActivity(ctx, ParseCSVChunkActivity, 
            ParseChunkInput{
                FileURL: input.FileURL,
                Offset:  offset,
                Limit:   chunkSize,
            }).Get(ctx, &chunk)
        if err != nil {
            return result, err
        }
        allTransactions = append(allTransactions, chunk...)
        
        // Add progress tracking
        workflow.RecordHeartbeat(ctx, fmt.Sprintf("Processed %d/%d rows", offset+len(chunk), fileData.RowCount))
    }
    
    // Rest of the workflow remains the same...
    // (duplicate detection, categorization, user confirmation, etc.)
    
    return result, nil
}
```

### Deployment Steps for Hetzner VPS

```bash
#!/bin/bash
# deploy.sh - Deploy Djinn with self-hosted Temporal on Hetzner VPS

# 1. Install Docker and Docker Compose
apt update && apt upgrade -y
apt install -y docker.io docker-compose git

# 2. Clone repository
git clone https://github.com/yourusername/djinn.git /opt/djinn
cd /opt/djinn

# 3. Create environment file
cat > .env <<EOF
DB_USER=djinn
DB_PASSWORD=$(openssl rand -base64 32)
MINIO_ROOT_USER=djinn
MINIO_ROOT_PASSWORD=$(openssl rand -base64 32)
TEMPORAL_NAMESPACE=default
TEMPORAL_ADDRESS=temporal:7233
EOF

# 4. Create necessary directories
mkdir -p temporal
mkdir -p scripts

# 5. Start services
docker-compose up -d

# 6. Wait for Temporal to be ready
echo "Waiting for Temporal to start..."
sleep 30

# 7. Create default namespace (if needed)
docker exec djinn-temporal temporal operator namespace create default

# 8. Setup Nginx reverse proxy for Temporal UI
apt install -y nginx
cat > /etc/nginx/sites-available/temporal-ui <<EOF
server {
    listen 80;
    server_name temporal.yourdomain.com;
    
    location / {
        proxy_pass http://127.0.0.1:8080;
        proxy_set_header Host \$host;
        proxy_set_header X-Real-IP \$remote_addr;
        proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto \$scheme;
    }
}
EOF

ln -s /etc/nginx/sites-available/temporal-ui /etc/nginx/sites-enabled/
nginx -t && systemctl reload nginx

echo "Deployment complete!"
echo "Temporal UI available at http://temporal.yourdomain.com"
echo "API available at http://localhost:4000"
```

### Monitoring and Resource Management

```yaml
monitoring:
  prometheus_metrics:
    # Expose metrics on :9090 but don't run separate Prometheus
    temporal_metrics:
      - workflow_start_total
      - workflow_completed_total
      - activity_execution_failed
      - task_queue_depth
    
  health_checks:
    api_health: "http://localhost:4000/health"
    temporal_health: "http://localhost:7233/health"
    worker_health: "Check worker logs"
    
  resource_monitoring:
    # Monitor with simple bash scripts
    memory_check: |
      docker stats --no-stream --format "table {{.Container}}\t{{.MemUsage}}"
    
    disk_check: |
      df -h /var/lib/docker
      du -sh /var/lib/docker/volumes/*
    
  alerts:
    # Use free monitoring like UptimeRobot
    - API endpoint monitoring
    - Disk space > 80%
    - Memory usage > 90%
```

### Scaling Strategy

```yaml
scaling_path:
  mvp_phase:
    users: "0-100"
    setup: "Single Hetzner CX22"
    temporal: "Self-hosted with minimal config"
    
  growth_phase_1:
    users: "100-1000"
    upgrade: "Hetzner CX32 (8GB RAM)"
    changes:
      - Increase Temporal history shards to 4
      - Add more worker concurrency
      - Enable Temporal metrics
    
  growth_phase_2:
    users: "1000+"
    options:
      option_1:
        name: "Temporal Cloud"
        cost: "$25/month"
        benefit: "Managed, no ops overhead"
      option_2:
        name: "Dedicated Temporal server"
        setup: "Second Hetzner CX22 for Temporal"
        cost: "€7.58/month total"
        benefit: "More resources, better isolation"
```

## Consequences

### Positive
- **Zero additional cost**: Temporal runs on existing VPS
- **Production-ready from day one**: Enterprise-grade reliability
- **Excellent developer experience**: Workflow as code, type safety
- **Built-in observability**: UI and metrics included
- **No vendor lock-in**: Fully self-hosted solution
- **Complete control**: Can tune for specific needs
- **Single server simplicity**: Everything on one VPS for MVP

### Negative
- **Resource constraints**: Limited by 4GB RAM
- **Operational overhead**: Need to manage Temporal ourselves
- **No built-in HA**: Single point of failure
- **Limited scale**: Will need to migrate as we grow
- **Backup complexity**: Need to backup Temporal state

### Risks and Mitigations
- **Risk**: Running out of memory
  - **Mitigation**: Careful resource limits, monitoring, upgrade path ready
  
- **Risk**: Temporal database grows too large
  - **Mitigation**: 7-day retention, regular cleanup, monitor disk usage
  
- **Risk**: Single server failure
  - **Mitigation**: Regular backups, quick redeploy scripts, eventual HA setup

## Alternatives Considered

### 1. Temporal Cloud
- **Pros**: Fully managed, no ops overhead, HA built-in
- **Cons**: $25/month cost, external dependency
- **Rejected**: Can self-host for free initially, migrate later if needed

### 2. PostgreSQL Queue
- **Pros**: Even simpler, less resource usage
- **Cons**: No workflow features, build everything ourselves
- **Rejected**: Temporal provides massive value even self-hosted

### 3. Separate Temporal VPS
- **Pros**: Better resource isolation, easier scaling
- **Cons**: Another €3.79/month, more complex networking
- **Rejected**: Not needed for MVP scale

## Implementation Notes

### Week 1: Setup and Deployment
- [ ] Deploy Docker Compose stack on Hetzner
- [ ] Configure Temporal with minimal resources
- [ ] Implement CSV import workflow
- [ ] Implement export generation workflow
- [ ] Setup monitoring scripts

### Week 2: Core Workflows
- [ ] Implement transaction enrichment
- [ ] Implement sync queue processing
- [ ] Add user confirmation signals
- [ ] GraphQL mutations for workflows
- [ ] Resource usage optimization

### Week 3: Production Hardening
- [ ] Backup strategy for Temporal
- [ ] Monitoring and alerting
- [ ] Performance testing within limits
- [ ] Documentation
- [ ] Upgrade path planning

### Configuration
```yaml
# .env for self-hosted setup
TEMPORAL_NAMESPACE=default
TEMPORAL_ADDRESS=localhost:7233  # Internal connection
TEMPORAL_TASK_QUEUE=djinn-main
WORKER_CONCURRENCY=5
ACTIVITY_CONCURRENCY=5
DB_POOL_SIZE=10
TEMPORAL_RETENTION_DAYS=7
```

## References
- [Temporal Self-Hosting Guide](https://docs.temporal.io/self-hosted-guide)
- [Temporal Docker Compose](https://github.com/temporalio/docker-compose)
- [Temporal Resource Requirements](https://docs.temporal.io/self-hosted-guide/production-checklist)
- ADR-20250820-file-storage-data-imports: MinIO file storage
- ADR-20250819-deployment-architecture: Hetzner VPS setup
- ADR-20250812-personal-finance-tech-stack: Tech stack decisions

## Decision Makers
- Author: Archie (System Architect)
- Reviewers: [Pending]
- Approver: [Pending]
- Date: 2025-08-20

## Revision History
- 2025-08-20: Initial draft with Temporal from day one
- 2025-08-20: Updated to self-hosted Temporal on Hetzner VPS