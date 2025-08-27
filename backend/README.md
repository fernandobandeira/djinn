# Djinn Backend API

A high-performance GraphQL API server built with Go, featuring clean architecture, interface-based dependency injection, and production-ready infrastructure components.

## 🚀 Quick Start

### Prerequisites
- Go 1.21+
- Docker & Docker Compose
- PostgreSQL 15+ (or use Docker)
- Make

### Running the Server

```bash
# Quick start (requires database already running)
make run

# Development mode (runs migrations and all code generation)
make run-dev

# Production build
make build
./bin/server
```

## 📋 Available Make Commands

```bash
make help              # Show all available commands
make run               # Run the API server
make run-dev           # Run with all code generation and migrations
make build             # Build server binary
make test              # Run all tests
make test-coverage     # Run tests with coverage report
make generate          # Generate all code (Wire, sqlc, mocks, dataloaders)
make migrate-up        # Apply database migrations
make migrate-down      # Reset database
make lint              # Run linters
```

## 🏗️ Architecture

### Clean Architecture Layers
```
┌─────────────────────────────────────────┐
│          Presentation Layer             │
│         (GraphQL Resolvers)             │
├─────────────────────────────────────────┤
│          Application Layer              │
│     (Command & Query Handlers)          │
├─────────────────────────────────────────┤
│           Domain Layer                  │
│    (Business Logic & Interfaces)        │
├─────────────────────────────────────────┤
│        Infrastructure Layer             │
│   (Database, External Services)         │
└─────────────────────────────────────────┘
```

### Key Features
- **Interface-Based DI**: All services exposed as interfaces for testability
- **Wire Dependency Injection**: Compile-time DI with Google Wire
- **Graceful Shutdown**: Coordinated component lifecycle management
- **Connection Pooling**: PgBouncer integration for efficient database connections
- **Observability**: OpenTelemetry tracing, Prometheus metrics
- **Testing**: 100% handler coverage with interface mocks

## 🔌 API Endpoints

- **GraphQL API**: `http://localhost:8080/graphql`
- **GraphQL Playground**: `http://localhost:8080/playground`
- **Health Check**: `http://localhost:8080/health`
- **Metrics**: `http://localhost:9091/metrics`

## 🐘 Database Architecture

### Connection Flow
```
API Server → PgBouncer (Pool) → PostgreSQL
     ↑            ↑                 ↑
  App Pool    Transaction Pool   Database
  (100s)         (10-20)          (5-10)
```

### PgBouncer Configuration

The application uses PgBouncer for connection pooling:
- **Development**: Optional, direct connection works
- **Production**: Recommended for connection efficiency

## 🚢 Production Deployment Architecture

### PgBouncer Deployment Strategies

#### 1. **Sidecar Pattern (Recommended for Kubernetes)**
```yaml
Pod:
  - API Container (your service)
  - PgBouncer Container (sidecar)
  
Benefits:
- 1 PgBouncer per pod/replica
- Network locality (localhost connection)
- Scales automatically with replicas
- Resource isolation
```

**Answer to your questions:**
- **Per Service**: YES - Each microservice gets its own PgBouncer
- **Per Replica**: YES - Each replica has its own PgBouncer sidecar
- **Per Node**: NO - Sidecars are per pod, not per node

#### 2. **Centralized PgBouncer Pool**
```
All Services → PgBouncer Cluster (3-5 instances) → PostgreSQL
                      ↑
                 Load Balanced
```

**Pros**: Simpler management, shared configuration
**Cons**: Network hop, single point of failure

#### 3. **Hybrid Approach (Best for Large Scale)**
```
Service Type A ─→ PgBouncer Pool A ─→┐
Service Type B ─→ PgBouncer Pool B ─→├─→ PostgreSQL Primary
Service Type C ─→ PgBouncer Pool C ─→┘
                                      └─→ PostgreSQL Read Replicas
```

### Recommended Production Setup

For **40 microservices** with **multiple replicas** on **5 nodes**:

```yaml
# Each microservice deployment
apiVersion: apps/v1
kind: Deployment
metadata:
  name: service-a
spec:
  replicas: 3
  template:
    spec:
      containers:
      - name: api
        image: your-service:latest
        env:
        - name: PGBOUNCER_URL
          value: "localhost:6432"  # Sidecar connection
      
      - name: pgbouncer
        image: pgbouncer/pgbouncer:latest
        env:
        - name: DATABASES_HOST
          value: "postgres-primary.db.svc.cluster.local"
        - name: DATABASES_PORT
          value: "5432"
        - name: POOL_MODE
          value: "transaction"
        - name: DEFAULT_POOL_SIZE
          value: "25"
        - name: MAX_CLIENT_CONN
          value: "100"
```

### PgBouncer Pool Sizing Formula

```
Per Service Instance:
- default_pool_size = 25 (to PostgreSQL)
- max_client_conn = 100 (from your app)

Total PostgreSQL Connections:
- 40 services × 3 replicas × 25 connections = 3000 connections
- With PgBouncer: ~200-300 actual PostgreSQL connections (90% reduction)
```

### Environment Variables

```bash
# Connection Configuration
DATABASE_URL=postgres://user:pass@postgres:5432/dbname
PGBOUNCER_URL=postgres://user:pass@localhost:6432/dbname  # Sidecar

# Pool Configuration (per instance)
DB_MAX_CONNECTIONS=25      # Max connections to PgBouncer
DB_MIN_CONNECTIONS=5       # Min connections to maintain
DB_MAX_CONN_LIFETIME=1h    # Connection lifetime
DB_MAX_CONN_IDLE_TIME=30m  # Idle timeout
```

## 🧪 Testing

```bash
# Unit tests
make test

# Integration tests (requires Docker)
make test-integration

# Full test suite with coverage
make test-coverage-full

# Stress tests
make test-stress
```

## 📊 Monitoring

The application exposes metrics and traces:
- **Prometheus Metrics**: Port 9091
- **OpenTelemetry Traces**: Port 4318
- **Health Endpoint**: `/health`

## 🔧 Development

### Code Generation

The project uses several code generation tools:

```bash
# Generate everything
make generate

# Individual generators
make wire-generate       # Wire dependency injection
make sqlc-generate       # Type-safe SQL
make mockery-generate    # Interface mocks
make dataloader-generate # GraphQL dataloaders
```

### Project Structure

```
backend/
├── cmd/
│   └── server/          # Application entrypoint
├── internal/
│   ├── domain/          # Business logic (interfaces)
│   ├── application/     # Use cases (command/query)
│   ├── infrastructure/  # External concerns
│   └── graph/           # GraphQL layer
├── migrations/          # Database migrations
└── test/               # Integration tests
```

## 🐳 Docker Compose Setup

```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f api

# Stop all services
docker-compose down
```

## 📚 Architecture Decision Records

See `/docs/architecture/adrs/` for detailed architectural decisions:
- ADR-20250127: Interface-Based Dependency Injection
- ADR-20250827: Graceful Shutdown Patterns
- ADR-20250820: Code Organization & Module Structure
- ADR-20250820: Testing Strategy

## 🤝 Contributing

1. Follow the interface-based architecture
2. Write tests for new functionality
3. Update ADRs for architectural changes
4. Run `make lint` before committing

## 📄 License

[Your License Here]