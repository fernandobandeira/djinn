# ADR-20250819: Deployment Architecture

## Status
Proposed (Updated with Ultra-Low-Cost MVP Options)

## Context
Djinn requires a deployment architecture that balances:
- Cost optimization for MVP ($50-100/month infrastructure budget)
- Ability to scale to 100K+ users without major re-architecture
- High availability for financial data (99.9% uptime target)
- Security and compliance requirements (GLBA, CCPA, future PCI DSS)
- Developer productivity with small team (5-8 people)
- Mobile app deployment and updates
- Fast global response times (< 200ms)
- Disaster recovery and backup strategies
- Monitoring and observability without breaking the budget

### Constraints
- Initial budget: $500K-1M total, **targeting $0-20/month for initial MVP testing**
- Small team cannot manage complex Kubernetes deployments initially
- Must support both web and mobile platforms
- PostgreSQL with specific requirements (UUIDv7, RLS)
- Temporal workflow engine requirement (can be deferred for initial MVP)
- Firebase Auth integration
- Need CI/CD from day one
- Must be able to migrate to enterprise platforms if needed
- **No sleep/cold start issues for core API services**

## Decision

### 1. Deployment Platform Strategy

#### Ultra-Low-Cost MVP Testing Phase (0-1K users, €3.79-20/month)

**PRIMARY RECOMMENDATION: Hetzner VPS (Best Value for Mobile-Only Backend)**
```yaml
deployment_architecture:
  backend:
    platform: Hetzner Cloud CX22
    cost: "€3.79/month (~$4.20/month)"
    
    what_you_get:
      - 2 vCPU AMD EPYC processors
      - 4GB RAM
      - 40GB NVMe SSD (includes file storage for receipts)
      - 20TB bandwidth (essentially unlimited)
      - Full root access
      - Run complete backend stack:
        - PostgreSQL database
        - Go API server
        - Redis cache
        - Receipt image storage (local filesystem)
        - Nginx for API gateway
        - Background job processors
        - Push notification workers
    
  mobile_app:
    platform: Flutter
    distribution:
      ios: "App Store via TestFlight"
      android: "Google Play Store"
    build_service: "GitHub Actions (free)"
    
  why_its_best:
    - Single server runs entire backend
    - No frontend hosting needed (mobile-only)
    - 40GB storage included for receipts
    - Complete control over API and data
    - No cold starts ever
    - Perfect for API-first architecture
    
  perfect_for:
    - Mobile-first financial apps
    - Teams comfortable with basic DevOps
    - Cost-conscious startups
    - Apps needing file storage
    - Full backend control
```

**Option B: Railway (Best for Zero DevOps)**
```yaml
deployment_architecture:
  platform: Railway
  cost: "$5-20/month (usage-based)"
  
  what_you_get:
    - One-click GitHub deployments
    - Managed PostgreSQL with backups
    - Automatic SSL/HTTPS
    - Built-in monitoring
    - Zero-downtime deployments
    
  limitations:
    - No built-in file storage (need external solution)
    - Higher cost than VPS
    - Need separate receipt storage solution
    
  perfect_for:
    - Teams with no DevOps expertise
    - Rapid prototyping
    - When external storage is acceptable

**Option C: Hybrid Free-Tier Solution (For Absolute Minimum Cost)**
```yaml
deployment_architecture:
  backend:
    platform: Fly.io
    cost: "$0 (free tier)"
    limits:
      - 3 shared-cpu-1x VMs
      - 256MB RAM per VM
      - 3GB persistent storage
      - Minimal cold starts
      
  database:
    platform: Supabase
    cost: "$0 (free tier)"
    limits:
      - 500MB database
      - Pauses after 1 week inactivity
      - 2GB bandwidth
      
  mobile_app:
    platform: Flutter
    distribution: "TestFlight + Google Play"
    cost: "$0 (except store fees)"
```


#### MVP Phase (0-10K users, €10-50/month)
```yaml
deployment_architecture:
  backend:
    platform: Hetzner Cloud CX32
    reason: "Scaled VPS for growing user base"
    cost: "€7.59/month (~$8.50/month)"
    specs:
      - 4 vCPU, 8GB RAM, 80GB SSD
      - Handles 10K+ users easily
      - 80GB storage for more receipts
      - Same stack as MVP phase
      
  mobile_app:
    platform: Flutter
    distribution:
      ios: "App Store (production)"
      android: "Google Play Store (production)"
    analytics: "Firebase Analytics (free tier)"
    crash_reporting: "Firebase Crashlytics (free)"
    push_notifications: "Firebase Cloud Messaging (free)"
      
  database:
    platform: Neon
    reason: "Serverless PostgreSQL with branching"
    cost: "$19/month (Launch plan)"
    features:
      - Serverless auto-scaling
      - Database branching for testing
      - Point-in-time recovery
      - Read replicas
      - PostgreSQL 16 support
      
  workflow_engine:
    platform: Temporal Cloud
    reason: "Managed Temporal for financial workflows"
    cost: "$25/month (Starter)"
    features:
      - Fully managed service
      - 99.99% SLA
      - Built-in monitoring
      - Automatic upgrades
```

#### Comparison Matrix for MVP Options

| Criteria | Railway | Hetzner VPS (Recommended) | Fly.io + Supabase | Render |
|----------|---------|----------------------------|-------------------|--------|
| **Monthly Cost** | $5-20 | €3.79 (~$4) | $0 | $0-7 |
| **Setup Time** | 5 minutes | 2-4 hours | 30 minutes | 20 minutes |
| **No Sleep/Cold Starts** | ✅ Never | ✅ Never | ⚠️ Minimal | ❌ Sleeps |
| **PostgreSQL** | ✅ Managed + Backups | Self-hosted | Managed (pauses) | Limited free |
| **Performance** | Good | Excellent | Good | Fair |
| **Scalability** | ✅ Automatic | Manual | Auto | Auto |
| **DevOps Required** | ❌ None | ✅ Yes | Minimal | Minimal |
| **Monitoring** | ✅ Built-in | Self-setup | Basic | Basic |
| **Time to Production** | Immediate | 1-2 days | Few hours | Few hours |
| **Team Collaboration** | ✅ Excellent | Manual setup | Good | Good |
| **Rollback Support** | ✅ One-click | Manual | Manual | Manual |

#### Hetzner VPS Scaling Strategy

```yaml
hetzner_scaling_path:
  mvp_phase:
    users: "0-1K"
    server: "CX22 (2 vCPU, 4GB RAM, 40GB SSD)"
    cost: "€3.79/month"
    setup: "Single server runs everything"
    
  early_growth:
    users: "1K-10K"
    server: "CX32 (4 vCPU, 8GB RAM, 80GB SSD)"
    cost: "€7.59/month"
    changes:
      - More CPU/RAM for growing load
      - Larger storage for receipts
      - Add automated backups
      
  validated_product:
    users: "10K-50K"
    architecture: "Multi-server setup"
    cost: "€20-50/month"
    changes:
      - Separate database server (CX32)
      - 2x API servers (CX22)
      - Hetzner Load Balancer
      - Dedicated backup volume
      
  enterprise_ready:
    users: "50K+"
    decision: "Evaluate managed cloud"
    triggers:
      - Need auto-scaling
      - Multi-region requirements
      - Compliance requirements
      - Team lacks DevOps bandwidth
```

#### Growth Phase (50K+ users, €50-200/month)
```yaml
scaling_adjustments:
  backend:
    option_1: "Multiple Hetzner servers with load balancer"
    option_2: "Move to managed cloud (AWS/GCP)"
    additions:
      - Dedicated database server
      - Redis cluster
      - Load balancer (Hetzner LB or HAProxy)
    
  mobile_scaling:
    considerations:
      - App Store optimization
      - Regional CDN for API responses
      - Push notification infrastructure
      
  database:
    upgrade: "Neon Scale plan"
    additions:
      - Read replicas for analytics
      - Connection pooling
      - Automated backup to S3
```

#### Scale Phase (50K-100K+ users, €200-500/month)
```yaml
enterprise_migration:
  trigger_points:
    - "Regulatory requirement for specific compliance"
    - "Need for multi-region deployment"
    - "Custom infrastructure requirements"
    
  migration_path:
    backend: "AWS ECS or GCP Cloud Run"
    database: "AWS RDS or GCP Cloud SQL"
    file_storage: "S3 or Cloud Storage"
    orchestration: "Temporal Cloud Enterprise"
    mobile_backend: "Consider Firebase for some services"
```

### 2. Infrastructure as Code

#### Hetzner VPS Setup (Primary Approach)
```bash
#!/bin/bash
# Ultra-simple VPS setup script for Go + PostgreSQL

# Update system
apt-get update && apt-get upgrade -y

# Install Docker and Docker Compose
curl -fsSL https://get.docker.com | sh
apt-get install -y docker-compose

# Install PostgreSQL 16
sh -c 'echo "deb https://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add -
apt-get update && apt-get install -y postgresql-16

# Configure PostgreSQL for production
cp /etc/postgresql/16/main/postgresql.conf /etc/postgresql/16/main/postgresql.conf.backup
cat >> /etc/postgresql/16/main/postgresql.conf <<EOF
# Performance tuning for 4GB RAM VPS
shared_buffers = 1GB
effective_cache_size = 3GB
maintenance_work_mem = 256MB
wal_buffers = 16MB
default_statistics_target = 100
random_page_cost = 1.1
effective_io_concurrency = 200
min_wal_size = 1GB
max_wal_size = 4GB
max_worker_processes = 2
max_parallel_workers_per_gather = 1
max_parallel_workers = 2
EOF

# Setup SSL with Let's Encrypt
apt-get install -y certbot python3-certbot-nginx nginx

# Create deployment directory
mkdir -p /opt/djinn
cd /opt/djinn

# Docker Compose configuration
cat > docker-compose.yml <<'EOF'
version: '3.8'

services:
  api:
    image: djinn-api:latest
    restart: always
    ports:
      - "8080:8080"
    environment:
      DATABASE_URL: postgres://djinn:${DB_PASSWORD}@localhost:5432/djinn_production
      REDIS_URL: redis://localhost:6379
    network_mode: host
    volumes:
      - ./config:/app/config:ro
    logging:
      driver: "json-file"
      options:
        max-size: "10m"
        max-file: "3"
        
  redis:
    image: redis:7-alpine
    restart: always
    ports:
      - "127.0.0.1:6379:6379"
    volumes:
      - redis_data:/data
    command: redis-server --appendonly yes --maxmemory 256mb --maxmemory-policy allkeys-lru

volumes:
  redis_data:
EOF

echo "VPS setup complete! Configure nginx and deploy your app."
```

#### Terraform Configuration
```hcl
# MVP Infrastructure Definition
terraform {
  required_providers {
    vercel = {
      source  = "vercel/vercel"
      version = "~> 0.16"
    }
    railway = {
      source  = "railway/railway"
      version = "~> 0.2"
    }
    neon = {
      source  = "neon/neon"
      version = "~> 0.1"
    }
  }
}

# Vercel Frontend
resource "vercel_project" "djinn_web" {
  name      = "djinn-web"
  framework = "nextjs"
  
  environment_variables = {
    NEXT_PUBLIC_API_URL = railway_service.api.domain
    NEXT_PUBLIC_FIREBASE_CONFIG = var.firebase_config
  }
  
  domains = ["app.djinn.finance"]
}

# Railway Backend
resource "railway_project" "djinn_backend" {
  name = "djinn-backend"
}

resource "railway_service" "api" {
  project_id = railway_project.djinn_backend.id
  name       = "djinn-api"
  
  source = {
    repo = "github.com/djinn/backend"
    branch = "main"
  }
  
  environment_variables = {
    DATABASE_URL = neon_database.djinn.connection_string
    REDIS_URL = railway_redis.cache.url
    TEMPORAL_ADDRESS = var.temporal_cloud_endpoint
  }
}

# Neon Database
resource "neon_project" "djinn" {
  name       = "djinn-db"
  region     = "us-east-1"
  postgres_version = 16
}

resource "neon_database" "djinn" {
  project_id = neon_project.djinn.id
  name       = "djinn_production"
  owner      = "djinn_app"
}

resource "neon_branch" "staging" {
  project_id = neon_project.djinn.id
  name       = "staging"
  parent_id  = neon_project.djinn.main_branch_id
}
```

### 3. CI/CD Pipeline

#### GitHub Actions Workflow
```yaml
name: Deploy Djinn

on:
  push:
    branches: [main, staging]
  pull_request:
    branches: [main]

env:
  VERCEL_ORG_ID: ${{ secrets.VERCEL_ORG_ID }}
  VERCEL_PROJECT_ID: ${{ secrets.VERCEL_PROJECT_ID }}

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Setup Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
          
      - name: Run Tests
        run: |
          go test -v -race -coverprofile=coverage.out ./...
          go tool cover -html=coverage.out -o coverage.html
          
      - name: Run Security Scan
        uses: securego/gosec@master
        with:
          args: ./...
          
      - name: SQLc Generate Check
        run: |
          sqlc generate
          git diff --exit-code
  
  deploy-frontend:
    needs: test
    if: github.ref == 'refs/heads/main'
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Deploy to Vercel
        run: |
          npm install -g vercel
          vercel pull --yes --environment=production
          vercel build --prod
          vercel deploy --prebuilt --prod
          
  deploy-backend:
    needs: test
    if: github.ref == 'refs/heads/main'
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Deploy to Railway
        env:
          RAILWAY_TOKEN: ${{ secrets.RAILWAY_TOKEN }}
        run: |
          npm install -g @railway/cli
          railway up --service djinn-api
          
  database-migration:
    needs: [deploy-backend]
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Run Migrations
        env:
          DATABASE_URL: ${{ secrets.DATABASE_URL }}
        run: |
          go install github.com/pressly/goose/v3/cmd/goose@latest
          goose -dir ./migrations postgres "$DATABASE_URL" up
```

### 4. Container Strategy

#### Docker Configuration
```dockerfile
# Multi-stage build for minimal image size
FROM golang:1.21-alpine AS builder

# Security: Run as non-root
RUN adduser -D -g '' appuser

# Install dependencies
RUN apk add --no-cache ca-certificates git

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build application
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags='-w -s -extldflags "-static"' \
    -o djinn-api ./cmd/api

# Final stage - minimal image
FROM scratch

# Import from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app/djinn-api /djinn-api

# Run as non-root
USER appuser

EXPOSE 8080
ENTRYPOINT ["/djinn-api"]
```

#### Docker Compose for Local Development
```yaml
version: '3.8'

services:
  postgres:
    image: postgres:16-alpine
    environment:
      POSTGRES_DB: djinn_dev
      POSTGRES_USER: djinn
      POSTGRES_PASSWORD: localdev
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./migrations:/docker-entrypoint-initdb.d
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U djinn"]
      interval: 5s
      timeout: 5s
      retries: 5

  redis:
    image: redis:7-alpine
    command: redis-server --appendonly yes
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data

  temporal:
    image: temporalio/auto-setup:latest
    ports:
      - "7233:7233"
    environment:
      - DB=postgresql
      - DB_PORT=5432
      - POSTGRES_USER=djinn
      - POSTGRES_PWD=localdev
      - POSTGRES_SEEDS=postgres
    depends_on:
      postgres:
        condition: service_healthy

  api:
    build: .
    ports:
      - "8080:8080"
    environment:
      DATABASE_URL: postgres://djinn:localdev@postgres:5432/djinn_dev
      REDIS_URL: redis://redis:6379
      TEMPORAL_ADDRESS: temporal:7233
      FIREBASE_EMULATOR: http://firebase:9099
    depends_on:
      - postgres
      - redis
      - temporal
    volumes:
      - ./:/app
    command: air # Hot reload for development

volumes:
  postgres_data:
  redis_data:
```

### 5. Database Deployment Strategy

#### Production Database Configuration
```sql
-- Neon-specific optimizations
ALTER SYSTEM SET max_connections = 100;
ALTER SYSTEM SET shared_buffers = '256MB';
ALTER SYSTEM SET effective_cache_size = '1GB';
ALTER SYSTEM SET maintenance_work_mem = '64MB';
ALTER SYSTEM SET checkpoint_completion_target = 0.9;
ALTER SYSTEM SET wal_buffers = '16MB';
ALTER SYSTEM SET default_statistics_target = 100;
ALTER SYSTEM SET random_page_cost = 1.1;

-- Connection pooling with PgBouncer (Railway/Neon provided)
-- Pool mode: transaction
-- Max client connections: 100
-- Default pool size: 25
```

#### Backup and Recovery Strategy
```yaml
backup_strategy:
  continuous:
    provider: "Neon built-in PITR"
    retention: "7 days"
    recovery_time_objective: "< 1 hour"
    recovery_point_objective: "< 5 minutes"
    
  daily_snapshots:
    destination: "AWS S3"
    retention: "30 days"
    encryption: "AES-256"
    schedule: "02:00 UTC"
    
  monthly_archives:
    destination: "AWS Glacier"
    retention: "7 years" # Regulatory requirement
    schedule: "First Sunday 03:00 UTC"
    
  testing:
    frequency: "Quarterly"
    procedure: "Full restore to staging"
```

### 6. Monitoring and Observability

#### Monitoring Stack (Budget-Conscious)
```yaml
monitoring:
  application:
    provider: Railway Metrics
    cost: "Included"
    metrics:
      - Request rate
      - Response time
      - Error rate
      - CPU/Memory usage
      
  database:
    provider: Neon Monitoring
    cost: "Included"
    metrics:
      - Query performance
      - Connection pool status
      - Replication lag
      - Storage usage
      
  advanced: # When budget allows
    provider: Grafana Cloud
    cost: "$50/month"
    features:
      - Custom dashboards
      - Alert manager
      - Log aggregation
      - Distributed tracing
      
  error_tracking:
    provider: Sentry
    cost: "Free tier (5K events/month)"
    features:
      - Error grouping
      - Performance monitoring
      - Release tracking
```

#### Observability Implementation
```go
// OpenTelemetry setup
import (
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/otlp/otlptrace"
    "go.opentelemetry.io/otel/sdk/trace"
)

func initTracing() {
    exporter, _ := otlptrace.New(
        context.Background(),
        otlptrace.WithEndpoint(os.Getenv("OTEL_ENDPOINT")),
    )
    
    tp := trace.NewTracerProvider(
        trace.WithBatcher(exporter),
        trace.WithResource(resource.NewWithAttributes(
            semconv.ServiceNameKey.String("djinn-api"),
            semconv.ServiceVersionKey.String(version),
        )),
    )
    
    otel.SetTracerProvider(tp)
}

// Structured logging with slog
func initLogging() {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
        Level: slog.LevelInfo,
        AddSource: true,
    }))
    
    slog.SetDefault(logger)
}

// Custom metrics
func initMetrics() {
    meter := otel.Meter("djinn-api")
    
    requestCounter, _ := meter.Int64Counter("requests_total")
    requestDuration, _ := meter.Float64Histogram("request_duration_seconds")
    activeUsers, _ := meter.Int64ObservableGauge("active_users")
}
```

### 7. Mobile App Deployment (Flutter-Only Architecture)

#### Flutter App Distribution
```yaml
mobile_deployment:
  development:
    ios: "TestFlight for beta testing"
    android: "Google Play Internal Testing"
    backend: "Hetzner VPS staging server"
    
  production:
    ios:
      store: "App Store"
      cost: "$99/year developer account"
      review_time: "24-48 hours typical"
      
    android:
      store: "Google Play Store"
      cost: "$25 one-time fee"
      review_time: "2-3 hours typical"
      
  build_pipeline:
    ci_cd: "GitHub Actions (free for public repos)"
    signing: "Fastlane for automation"
    
  backend_connection:
    api_endpoint: "https://api.djinn.finance"
    authentication: "Firebase Auth"
    certificate_pinning: "For security"
    
  advantages_of_mobile_only:
    - No web hosting costs
    - No frontend CDN needed
    - Simplified architecture
    - Better security (no web surface)
    - Native performance
    - Biometric authentication
    - Direct camera access for receipts
```

### 8. Security Hardening

#### Production Security Checklist
```yaml
security_measures:
  network:
    - SSL/TLS everywhere (Let's Encrypt)
    - WAF enabled (Vercel/CloudFlare)
    - DDoS protection (Platform provided)
    - Private networking for services
    
  secrets:
    - GitHub Secrets for CI/CD
    - Railway/Vercel environment variables
    - Temporal Cloud secret store
    - No secrets in code or Docker images
    
  access_control:
    - MFA for all admin accounts
    - Service-to-service authentication
    - Database connection over SSL
    - VPN for production access (when scaled)
    
  compliance:
    - HTTPS only
    - Security headers (CSP, HSTS, etc.)
    - Regular dependency updates
    - Security scanning in CI/CD
```

### 9. Disaster Recovery Plan

#### RTO and RPO Targets
```yaml
disaster_recovery:
  targets:
    recovery_time_objective: "4 hours"
    recovery_point_objective: "1 hour"
    
  scenarios:
    database_failure:
      detection: "< 5 minutes"
      action: "Failover to read replica"
      recovery: "Restore from PITR"
      
    region_outage:
      detection: "< 1 minute"
      action: "DNS failover to backup region"
      recovery: "Full service restoration"
      
    data_corruption:
      detection: "< 1 hour"
      action: "Stop writes, assess damage"
      recovery: "PITR to before corruption"
      
    security_breach:
      detection: "ASAP"
      action: "Isolate, preserve evidence"
      recovery: "Clean rebuild from backup"
```

### 10. Cost Optimization Strategies

#### Mobile-First MVP Strategy
```yaml
cost_progression:
  testing_phase:
    duration: "1-3 months"
    target: "€3.79/month"
    approach:
      primary: "Hetzner CX22 for complete backend"
      alternative: "Fly.io + Supabase free tiers"
    features:
      - Basic API functionality
      - PostgreSQL database
      - Simple frontend on Vercel
      - GitHub Actions CI/CD
    limitations_accepted:
      - Manual scaling
      - Basic monitoring only
      - Single region deployment
      - Self-managed backups
      
  validation_phase:
    duration: "3-6 months"
    target: "€7.59/month"
    approach:
      primary: "Hetzner CX32 upgrade"
      alternative: "Add monitoring and backups"
    additions:
      - Managed database backups
      - Basic monitoring (Uptimerobot free)
      - CDN integration
      - Error tracking (Sentry free)
      
  growth_ready:
    duration: "6+ months"
    target: "€20-100/month"
    approach: "Multiple Hetzner servers or managed cloud"
    trigger: "Product-market fit validated"
    
  scale_transition:
    duration: "When exceeding 50K users"
    target: "Optimize for scale not cost"
    approach: "Migrate to AWS/GCP for auto-scaling"
    trigger: "Need for multi-region or compliance"
```

#### Hetzner VPS Cost Breakdown (Mobile Backend)
```yaml
hetzner_costs:
  mvp_phase:
    server: "€3.79/month (CX22)"
    backup: "€0.76/month (optional)"
    total: "€3.79-4.55/month"
    
  growth_phase:
    server: "€7.59/month (CX32)"
    backup: "€1.52/month"
    monitoring: "€0 (self-hosted)"
    total: "€9.11/month"
    
  scale_phase:
    option_1_multiple_vps:
      web_servers: "2x CX22 = €7.58/month"
      db_server: "CX32 = €7.59/month"
      load_balancer: "€5.39/month"
      total: "€20.56/month"
    
    option_2_larger_servers:
      api_server: "CX42 = €15.49/month"
      db_server: "CX42 = €15.49/month"
      total: "€30.98/month"
    
  mobile_specific_savings:
    no_web_hosting: "Save $20-50/month"
    no_cdn: "Save $10-50/month"
    no_frontend_build: "Save CI/CD minutes"
    simplified_stack: "Reduce complexity 50%"
```

#### Progressive Cost Management
```yaml
cost_optimization:
  mvp_phase:
    target: "$50-100/month"
    strategies:
      - Use free tiers maximally
      - Shared resources (database, Redis)
      - Platform-native monitoring
      - Aggressive caching
      
  growth_phase:
    target: "$100-300/month"
    strategies:
      - Reserved capacity discounts
      - Optimize database queries
      - CDN for static assets
      - Implement rate limiting
      
  scale_phase:
    target: "< $10 per 1000 users"
    strategies:
      - Negotiate enterprise pricing
      - Multi-year commitments
      - Self-host where beneficial
      - Regional deployment optimization
      
  monitoring:
    - Weekly cost reviews
    - Alert on 20% increase
    - Monthly optimization audit
    - Quarterly vendor review
```

## Consequences

### Positive
- **Low Initial Cost**: MVP deployable for $50-100/month
- **Quick Time to Market**: Platform services reduce setup time by 80%
- **Scalability Path**: Clear migration path to enterprise platforms
- **Developer Productivity**: Managed services free team to focus on features
- **Security**: Platform-provided security features meet compliance needs
- **Flexibility**: Can switch providers without major code changes
- **Monitoring**: Built-in observability without additional cost

### Negative
- **Vendor Lock-in**: Some platform-specific features used
- **Limited Customization**: Platform constraints may limit options
- **Migration Complexity**: Moving to enterprise platforms requires effort
- **Cost Scaling**: Per-seat pricing can become expensive at scale
- **Regional Limitations**: Some platforms have limited regions

### Risks
- **Platform Outages**: Dependency on third-party availability
- **Price Changes**: Platforms may increase pricing
- **Feature Deprecation**: Platform features may be discontinued
- **Compliance Gaps**: Platform may not meet future compliance needs
- **Performance Limits**: Platform constraints may impact performance

## Migration Path from Hetzner to Cloud

### When to Consider Migration
```yaml
migration_triggers:
  - users: "> 50,000 active users"
  - requirements: "Multi-region deployment needed"
  - compliance: "Specific cloud compliance (SOC2, etc)"
  - team: "Need managed services to reduce DevOps"
  - scaling: "Unpredictable traffic spikes"
  
migration_options:
  stay_on_hetzner:
    when:
      - "Predictable growth"
      - "Cost is primary concern"
      - "Team has DevOps skills"
      - "Single region is sufficient"
    scaling:
      - "Add more servers (€3.79-30 each)"
      - "Use Hetzner Load Balancer (€5.39/month)"
      - "Add Hetzner Volumes for storage"
    
  migrate_to_cloud:
    aws_option:
      target: "ECS Fargate + RDS"
      cost: "$200-500/month"
      benefits: "Auto-scaling, managed services"
      
    gcp_option:
      target: "Cloud Run + Cloud SQL"
      cost: "$150-400/month"
      benefits: "Serverless, automatic scaling"
      
  hybrid_approach:
    keep_on_hetzner: "Database and file storage"
    move_to_cloud: "API servers for auto-scaling"
    benefits: "Balance cost and flexibility"
```

## Alternatives Considered

### Option A: AWS from Day One
- **Description**: Full AWS deployment
- **Pros**: Enterprise-ready, all compliance, infinite scale
- **Cons**: Complex, expensive ($500+/month), steep learning curve
- **Reason for not choosing**: Over-engineered for MVP, too expensive

### Option B: Kubernetes on Hetzner
- **Description**: Self-managed K3s cluster on Hetzner
- **Pros**: Full control, very cost-effective, cloud-native
- **Cons**: Complex operations, requires Kubernetes expertise
- **Reason for not choosing**: Overkill for MVP, adds unnecessary complexity

### Option C: Serverless Everything
- **Description**: AWS Lambda/Vercel Functions only
- **Pros**: True pay-per-use, infinite scale
- **Cons**: Vendor lock-in, cold starts, complex local development
- **Reason for not choosing**: Temporal workflows need persistent compute

### Option D: DigitalOcean Droplets
- **Description**: Similar to Hetzner but US-based
- **Pros**: $200 free credits, good documentation, predictable pricing
- **Cons**: More expensive than Hetzner (€6/month for 1GB RAM)
- **Reason for not choosing**: Hetzner offers better value (4GB RAM for €3.79)

### Option E: Self-Hosted on Raspberry Pi
- **Description**: Home hosting on ARM hardware
- **Pros**: One-time hardware cost (~$100), full control
- **Cons**: Reliability issues, ISP restrictions, no SLA
- **Reason for not choosing**: Not suitable for production financial services
- **Description**: AWS Lambda/Vercel Functions only
- **Pros**: True pay-per-use, infinite scale
- **Cons**: Vendor lock-in, cold starts, complex local development
- **Reason for not choosing**: Temporal workflows need persistent compute

## Implementation Notes

### Migration Strategy
1. **Week 1**: Set up Vercel and Railway projects
2. **Week 2**: Configure Neon database and migrations
3. **Week 3**: Implement CI/CD pipeline
4. **Week 4**: Set up monitoring and alerts
5. **Week 5**: Security hardening and testing
6. **Week 6**: Production deployment and validation

### Testing Approach
- Infrastructure testing with Terraform
- Deployment pipeline testing
- Disaster recovery drills
- Load testing on staging
- Security scanning

### Monitoring and Success Metrics
- Deployment frequency > 5/week
- Deployment failure rate < 5%
- Mean time to recovery < 1 hour
- Infrastructure cost < $100/month for MVP
- Page load time < 2 seconds globally
- API response time p95 < 200ms
- Uptime > 99.9%

## References
- [Vercel Pricing and Features](https://vercel.com/pricing)
- [Railway Documentation](https://docs.railway.app)
- [Neon Serverless Postgres](https://neon.tech/docs)
- [Temporal Cloud](https://temporal.io/cloud)
- [12 Factor App Methodology](https://12factor.net)
- ADR-20250812: Personal Finance Tech Stack Selection
- ADR-20250819: Data Architecture and Schema Design
- ADR-20250819: Security Architecture
- ADR-20250819: API Design Standards

## Decision Makers
- Author: Archie (System Architect)
- Reviewers: [Pending]
- Approver: [Pending]
- Date: 2025-08-19

## Revision History
- 2025-08-19: Initial comprehensive draft with cost optimization focus
- 2025-08-19: Added ultra-low-cost MVP options ($0-20/month) with VPS and free-tier strategies