# ADR-20250819: Deployment Architecture

## Status
Accepted

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
- Initial budget: $500K-1M total, $50-100/month for MVP infrastructure
- Small team cannot manage complex Kubernetes deployments initially
- Must support both web and mobile platforms
- PostgreSQL with specific requirements (UUIDv7, RLS)
- Temporal workflow engine requirement
- Firebase Auth integration
- Need CI/CD from day one
- Must be able to migrate to enterprise platforms if needed

## Decision

### 1. Deployment Platform Strategy

#### MVP Phase (0-10K users, $50-100/month)
```yaml
deployment_architecture:
  frontend:
    platform: Vercel
    reason: "Best-in-class frontend hosting with security features"
    cost: "$20/month (Pro plan)"
    features:
      - Edge functions for API routes
      - Automatic HTTPS/SSL
      - DDoS protection included
      - Web Application Firewall
      - SOC 2 Type 2 compliant
      - Analytics included
      
  backend:
    platform: Railway
    reason: "Rapid deployment with built-in services"
    cost: "$20-40/month"
    features:
      - One-click PostgreSQL deployment
      - Automatic SSL certificates
      - Built-in Redis support
      - GitHub integration
      - Auto-scaling capability
      - Private networking
      
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

#### Growth Phase (10K-50K users, $100-300/month)
```yaml
scaling_adjustments:
  frontend:
    upgrade: "Vercel Pro with increased limits"
    additions:
      - CloudFlare CDN for global distribution
      - Image optimization pipeline
    
  backend:
    scaling: "Railway horizontal scaling"
    additions:
      - Multiple service replicas
      - Load balancer configuration
      - Dedicated Redis cluster
      
  database:
    upgrade: "Neon Scale plan"
    additions:
      - Read replicas for analytics
      - Connection pooling
      - Automated backup to S3
```

#### Scale Phase (50K-100K+ users, $300-800/month)
```yaml
enterprise_migration:
  trigger_points:
    - "Regulatory requirement for specific compliance"
    - "Need for multi-region deployment"
    - "Custom infrastructure requirements"
    
  migration_path:
    frontend: "Vercel Enterprise or AWS CloudFront"
    backend: "AWS ECS or GCP Cloud Run"
    database: "AWS RDS or GCP Cloud SQL"
    orchestration: "Temporal Cloud Enterprise"
```

### 2. Infrastructure as Code

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

### 7. Mobile App Deployment

#### Flutter App Distribution
```yaml
mobile_deployment:
  ios:
    build_service: "Codemagic"
    cost: "$0 (500 build minutes free)"
    distribution:
      beta: "TestFlight"
      production: "App Store"
    ci_integration: "GitHub Actions"
    
  android:
    build_service: "Codemagic"
    distribution:
      beta: "Firebase App Distribution"
      production: "Google Play Store"
    signing: "Google Play App Signing"
    
  over_the_air_updates:
    provider: "CodePush (React Native) or custom"
    limitations: "Asset updates only, no native code"
    
  crash_reporting:
    provider: "Firebase Crashlytics"
    cost: "Free"
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

## Alternatives Considered

### Option A: AWS from Day One
- **Description**: Full AWS deployment
- **Pros**: Enterprise-ready, all compliance, infinite scale
- **Cons**: Complex, expensive ($500+/month), steep learning curve
- **Reason for not choosing**: Over-engineered for MVP, too expensive

### Option B: Kubernetes on DigitalOcean
- **Description**: Self-managed Kubernetes cluster
- **Pros**: Full control, portable, cost-effective at scale
- **Cons**: Complex operations, requires DevOps expertise
- **Reason for not choosing**: Team too small to manage Kubernetes

### Option C: Serverless Everything
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