# Performance and Scalability Architecture

## Context
Performance architecture focusing on database-centric progressive enhancement, optimized for Hetzner VPS constraints (4GB RAM) with sub-200ms API response targets.

```mermaid
graph TD
    subgraph "Client Layer"
        A[API Client] --> |Optimized Requests| B[GraphQL API]
    end

    subgraph "API Layer"
        B --> |DataLoader Caching| C{Request-Scoped Caching}
        C --> |Elimination of N+1 Queries| D[DataLoader Service]
    end

    subgraph "Database Layer"
        D --> |PgBouncer Connection Pooling| E[(PostgreSQL)]
        E <--> |System Data Caching| F[Redis Cache]
    end

    subgraph "Monitoring"
        G[Prometheus] --> H[Grafana Dashboards]
        I[Distributed Tracing] --> J[Jaeger UI]
    end

    subgraph "Infrastructure"
        K[Hetzner VPS] 
        K --> |4GB RAM Constraint| L[Resource Optimization]
    end

    classDef core fill:#4A90E2,color:white;
    classDef data fill:#F5A623,color:black;
    classDef monitoring fill:#27AE60,color:white;

    class A,B,C,D core
    class E,F data
    class G,H,I,J monitoring
    class K,L fill:#E74C3C,color:white
```

## Key Performance Strategies
- **DataLoader Pattern**: Eliminates N+1 queries through batched, cached data retrieval
- **Connection Pooling**: PgBouncer managing 25 PostgreSQL connections
- **Caching Strategy**: 
  - Request-scoped caching 
  - Redis ONLY for system-level data
- **Response Target**: < 200ms API responses
- **Infrastructure**: Optimized for 4GB RAM Hetzner VPS

## Monitoring Components
- Prometheus for metrics collection
- Grafana for performance dashboards
- Jaeger for distributed tracing

## Constraints and Optimizations
- Minimal external caching
- Database-centric performance model
- Progressive enhancement approach