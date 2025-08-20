# ADR-20250120: Performance & Scalability Architecture

## Status
Accepted

## Context
Djinn requires a performance and scalability architecture that balances:
- **Simplicity**: Small team (5-8 people) needs maintainable solutions
- **Performance**: Sub-200ms API response times for financial operations
- **Scalability**: Support growth from MVP to 100K+ users
- **Cost-effectiveness**: Optimize for low operational costs at each stage
- **Mobile-first**: Optimize for mobile app performance and battery life
- **Financial requirements**: ACID compliance for transactions, audit trails
- **Deployment constraints**: Building on Hetzner VPS infrastructure (see ADR-20250819-deployment-architecture)

### Current State
- Go backend with chi router
- PostgreSQL 16 with UUIDv7
- Redis for caching and sessions
- Mobile-only Flutter app (no web frontend)
- Target: €3.79-50/month infrastructure for MVP

### Performance Requirements
| Metric | MVP (0-1K users) | Growth (1K-10K) | Scale (10K-100K+) |
|--------|------------------|-----------------|-------------------|
| API Response (p95) | < 500ms | < 200ms | < 100ms |
| Database Queries | < 100ms | < 50ms | < 30ms |
| Concurrent Users | 50 | 500 | 5,000 |
| Requests/Second | 10 | 100 | 1,000 |
| Uptime | 99% | 99.5% | 99.9% |

## Decision

### Core Strategy: Progressive Enhancement with Database-Centric Architecture

We will adopt a **simplicity-first approach** that starts with database optimization and progressively adds caching and scaling layers only when metrics indicate they're needed.

### 1. Database Performance Foundation

#### PostgreSQL Optimization Strategy
```yaml
database_optimization:
  connection_management:
    pooling:
      tool: "PgBouncer in transaction mode"
      pool_size: 25  # For 4GB RAM VPS
      max_connections: 100
      benefits:
        - Reduces connection overhead by 90%
        - Prevents connection exhaustion
        - Transaction-level pooling for Go's short-lived connections
    
  query_optimization:
    indexing_strategy:
      primary: "UUIDv7 for time-ordered clustering"
      composite: "Multi-column indexes for common queries"
      partial: "For filtered queries (e.g., active users)"
      covering: "Include columns to avoid heap lookups"
    
    patterns:
      - Use prepared statements (sqlc)
      - Batch operations where possible
      - EXPLAIN ANALYZE on slow queries
      - Query result size limits
```

#### Database Configuration (Hetzner CX22 - 4GB RAM)
```sql
-- Performance tuning for 4GB RAM VPS
ALTER SYSTEM SET shared_buffers = '1GB';
ALTER SYSTEM SET effective_cache_size = '3GB';
ALTER SYSTEM SET maintenance_work_mem = '256MB';
ALTER SYSTEM SET work_mem = '16MB';  -- Per operation memory

-- Connection pooling settings
ALTER SYSTEM SET max_connections = 100;

-- Write performance
ALTER SYSTEM SET wal_buffers = '16MB';
ALTER SYSTEM SET checkpoint_completion_target = 0.9;
ALTER SYSTEM SET min_wal_size = '1GB';
ALTER SYSTEM SET max_wal_size = '4GB';

-- Query planning
ALTER SYSTEM SET random_page_cost = 1.1;  -- SSD optimization
ALTER SYSTEM SET effective_io_concurrency = 200;  -- NVMe SSD
ALTER SYSTEM SET default_statistics_target = 100;

-- Parallel query execution (conservative for small VPS)
ALTER SYSTEM SET max_worker_processes = 2;
ALTER SYSTEM SET max_parallel_workers_per_gather = 1;
ALTER SYSTEM SET max_parallel_workers = 2;
```

### 2. Application-Level Performance

#### Go Backend Optimization
```go
// Connection pool configuration
type DatabaseConfig struct {
    MaxOpenConns    int           // 25 for PgBouncer
    MaxIdleConns    int           // 5
    ConnMaxLifetime time.Duration // 1 hour
    ConnMaxIdleTime time.Duration // 10 minutes
}

// Query optimization patterns
type PerformancePatterns struct {
    // 1. Use context with timeout for all queries
    QueryTimeout: 5 * time.Second,
    
    // 2. Implement query result limits
    DefaultLimit: 100,
    MaxLimit:    1000,
    
    // 3. Use prepared statements via sqlc
    UsePreparedStatements: true,
    
    // 4. Batch operations
    BatchSize: 100,
}

// Request handling optimization
func OptimizedMiddleware() func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // 1. Request ID for tracing
            requestID := uuid.New().String()
            ctx := context.WithValue(r.Context(), "request_id", requestID)
            
            // 2. Timeout for entire request
            ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
            defer cancel()
            
            // 3. Response compression
            w.Header().Set("Content-Encoding", "gzip")
            gz := gzip.NewWriter(w)
            defer gz.Close()
            
            // 4. Add performance headers
            w.Header().Set("X-Request-ID", requestID)
            w.Header().Set("Cache-Control", "private, max-age=0")
            
            next.ServeHTTP(gzipResponseWriter{gz, w}, r.WithContext(ctx))
        })
    }
}
```

### 3. Caching Strategy (Minimal & Strategic)

#### DataLoader Pattern (Primary Caching Strategy)
```go
// Using dataloaden for type-safe code generation
// This eliminates N+1 queries without cache invalidation complexity
// Install: go get github.com/vektah/dataloaden

//go:generate dataloaden UserLoader string *model.User
//go:generate dataloaden AccountLoader string *model.Account
//go:generate dataloaden CategoryLoader string *model.Category
//go:generate dataloaden TransactionLoader string *model.Transaction

// Generated loaders with full type safety
type DataLoaders struct {
    UserLoader       *UserLoader       // Generated by dataloaden
    AccountLoader    *AccountLoader    // Generated by dataloaden
    CategoryLoader   *CategoryLoader   // Generated by dataloaden
    TransactionLoader *TransactionLoader // Generated by dataloaden
}

// Initialize loaders with batch functions
func NewDataLoaders(db *sqlc.Queries) *DataLoaders {
    return &DataLoaders{
        UserLoader: NewUserLoader(UserLoaderConfig{
            MaxBatch: 100,
            Wait:     2 * time.Millisecond,
            Fetch: func(keys []string) ([]*model.User, []error) {
                return batchUsers(db, keys)
            },
        }),
        AccountLoader: NewAccountLoader(AccountLoaderConfig{
            MaxBatch: 100,
            Wait:     2 * time.Millisecond,
            Fetch: func(keys []string) ([]*model.Account, []error) {
                return batchAccounts(db, keys)
            },
        }),
        // ... other loaders initialized similarly
    }
}

// Middleware to attach DataLoaders to request context
func DataLoaderMiddleware(db *sqlc.Queries) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            loaders := NewDataLoaders(db)
            ctx := context.WithValue(r.Context(), "loaders", loaders)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}
```

#### Thin System Data Cache (Static Content Only)
```go
// Cache ONLY truly static system data that changes rarely
type SystemCache struct {
    data sync.Map  // Thread-safe map
    ttl  time.Duration
}

// System endpoints that benefit from caching
var systemCacheTargets = map[string]time.Duration{
    "/api/system/institutions":     24 * time.Hour,  // Bank/institution list
    "/api/system/categories":       24 * time.Hour,  // Default categories
    "/api/system/currencies":       24 * time.Hour,  // Supported currencies
    "/api/system/countries":        7 * 24 * time.Hour,  // Country list
    "/api/system/timezones":        7 * 24 * time.Hour,  // Timezone data
    "/api/system/config/mobile":    1 * time.Hour,   // Mobile app config
}

// Simple cache implementation
func (c *SystemCache) Get(key string) (interface{}, bool) {
    if val, ok := c.data.Load(key); ok {
        if entry, ok := val.(*CacheEntry); ok {
            if time.Now().Before(entry.ExpiresAt) {
                return entry.Data, true
            }
            c.data.Delete(key)
        }
    }
    return nil, false
}

// Middleware for system endpoint caching
func SystemCacheMiddleware(cache *SystemCache) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Only cache GET requests to system endpoints
            if r.Method != "GET" || !strings.HasPrefix(r.URL.Path, "/api/system/") {
                next.ServeHTTP(w, r)
                return
            }
            
            // Check cache
            if data, found := cache.Get(r.URL.Path); found {
                w.Header().Set("X-Cache", "HIT")
                w.Header().Set("Content-Type", "application/json")
                json.NewEncoder(w).Encode(data)
                return
            }
            
            // Cache miss - capture response
            rec := httptest.NewRecorder()
            next.ServeHTTP(rec, r)
            
            // Cache successful responses
            if rec.Code == http.StatusOK {
                var data interface{}
                json.Unmarshal(rec.Body.Bytes(), &data)
                if ttl, ok := systemCacheTargets[r.URL.Path]; ok {
                    cache.Set(r.URL.Path, data, ttl)
                }
            }
            
            // Copy response to client
            for k, v := range rec.Header() {
                w.Header()[k] = v
            }
            w.Header().Set("X-Cache", "MISS")
            w.WriteHeader(rec.Code)
            w.Write(rec.Body.Bytes())
        })
    }
}
```

#### Why This Approach Works
```yaml
dataloader_benefits:
  automatic_batching:
    - Combines multiple queries into single batch
    - Reduces database round trips by 10x
    - No manual batching code needed
    
  request_scoped_caching:
    - Cache lives only for request duration
    - No cache invalidation problems
    - Perfect consistency within request
    - Zero memory leak risk
    
  n_plus_one_elimination:
    - GraphQL resolver friendly
    - Transparent to application code
    - Works with any data fetching pattern

system_cache_benefits:
  minimal_complexity:
    - Only cache truly static data
    - Simple TTL-based expiration
    - No invalidation logic needed
    
  high_impact:
    - System endpoints called by every app launch
    - Reduces database load significantly
    - Improves app startup time
    
  examples:
    institutions_endpoint:
      calls_per_day: "10,000 (once per app launch)"
      database_queries_saved: "10,000/day"
      response_time_improvement: "200ms → 5ms"
    
    categories_endpoint:
      calls_per_day: "5,000"
      database_queries_saved: "5,000/day"
      response_time_improvement: "150ms → 5ms"
```

#### Redis Addition Triggers (Only When Needed)
```yaml
add_redis_when:
  conditions:
    - "Session management across multiple servers"
    - "Rate limiting across multiple servers"
    - "Distributed locks needed"
    - "Cache needs to survive server restart"
    - "DataLoader insufficient for query patterns"
  
  probably_not_needed_because:
    - "DataLoaders handle most caching needs"
    - "Single server handles 10K+ users"
    - "System cache handles static content"
    - "Database is fast with proper indexes"
```

### 4. API Performance Patterns

#### Rate Limiting (Simple but Effective)
```go
// In-memory rate limiter for MVP
type RateLimiter struct {
    requests map[string][]time.Time
    mu       sync.Mutex
    limit    int           // Requests per window
    window   time.Duration // Time window
}

func (rl *RateLimiter) Allow(clientID string) bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()
    
    now := time.Now()
    windowStart := now.Add(-rl.window)
    
    // Clean old requests
    requests := rl.requests[clientID]
    validRequests := []time.Time{}
    for _, t := range requests {
        if t.After(windowStart) {
            validRequests = append(validRequests, t)
        }
    }
    
    if len(validRequests) >= rl.limit {
        return false
    }
    
    rl.requests[clientID] = append(validRequests, now)
    return true
}

// Progressive rate limits
var rateLimits = map[string]RateLimit{
    "anonymous":    {10, time.Minute},   // 10 req/min
    "authenticated": {100, time.Minute},  // 100 req/min
    "premium":      {1000, time.Minute},  // 1000 req/min
}
```

#### Response Optimization
```yaml
response_optimization:
  compression:
    algorithm: "gzip"
    min_size: 1024  # Don't compress tiny responses
    level: 5        # Balance speed/compression
    
  pagination:
    default_size: 20
    max_size: 100
    cursor_based: true  # More efficient than offset
    
  field_filtering:
    strategy: "GraphQL-like field selection"
    example: "?fields=id,name,balance"
    
  etag_support:
    enabled: true
    cache_duration: "5 minutes"
```

### 5. Mobile-Specific Optimizations

#### API Design for Mobile
```yaml
mobile_optimizations:
  batching:
    description: "Combine multiple requests"
    endpoint: "POST /api/batch"
    max_operations: 10
    benefits:
      - Reduces round trips
      - Better battery life
      - Lower latency
    
  delta_sync:
    description: "Send only changes"
    implementation:
      - Version vectors
      - Last-modified timestamps
      - Incremental updates
    
  adaptive_quality:
    network_detection: true
    responses:
      slow_network: "Essential data only"
      fast_network: "Full data with prefetch"
    
  push_instead_of_poll:
    technology: "Firebase Cloud Messaging"
    events:
      - "Transaction received"
      - "Budget exceeded"
      - "Bill reminder"
```

### 6. Scaling Strategy (When Needed)

#### Vertical Scaling Path (Simplest)
```yaml
vertical_scaling:
  phase_1_mvp:
    server: "Hetzner CX22"
    specs: "2 vCPU, 4GB RAM"
    capacity: "0-1K users"
    cost: "€3.79/month"
    
  phase_2_growth:
    server: "Hetzner CX32"
    specs: "4 vCPU, 8GB RAM"
    capacity: "1K-10K users"
    cost: "€7.59/month"
    upgrade: "10 minute downtime"
    
  phase_3_scale:
    server: "Hetzner CX42"
    specs: "8 vCPU, 16GB RAM"
    capacity: "10K-50K users"
    cost: "€15.49/month"
    
  phase_4_distributed:
    trigger: "> 50K users or vertical limits"
    approach: "Move to horizontal scaling"
```

#### Horizontal Scaling Path (When Vertical Limits Hit)
```yaml
horizontal_scaling:
  triggers:
    - "CPU consistently > 80%"
    - "Memory > 90%"
    - "Response times > SLA"
    - "Need for zero-downtime deployments"
    
  progression:
    step_1:
      action: "Separate database server"
      servers:
        - "API: CX22 (€3.79)"
        - "DB: CX32 (€7.59)"
      total_cost: "€11.38/month"
      
    step_2:
      action: "Add read replica"
      servers:
        - "API: CX22"
        - "DB Primary: CX32"
        - "DB Replica: CX22"
      total_cost: "€15.17/month"
      
    step_3:
      action: "Multiple API servers + Load Balancer"
      servers:
        - "2x API: CX22"
        - "DB Primary: CX32"
        - "DB Replica: CX22"
        - "Load Balancer: €5.39"
      total_cost: "€24.35/month"
```

### 7. Monitoring & Performance Tracking

#### Essential Metrics (Simple Stack)
```yaml
monitoring:
  application_metrics:
    tool: "Prometheus + Grafana (self-hosted)"
    key_metrics:
      - request_duration_seconds
      - request_total
      - error_rate
      - active_connections
      - memory_usage
      - cpu_usage
    
  database_metrics:
    tool: "pg_stat_statements"
    track:
      - Slow queries (> 100ms)
      - Connection pool usage
      - Cache hit ratio (> 95% target)
      - Index usage
      - Table bloat
    
  business_metrics:
    custom_tracking:
      - Active users per hour
      - Transactions per minute
      - API calls per user
      - Feature usage patterns
      
  alerting:
    critical:
      - "API response > 1s"
      - "Error rate > 1%"
      - "Database connections > 80"
      - "Disk usage > 80%"
    
    warning:
      - "API response > 500ms"
      - "Memory usage > 70%"
      - "Cache hit ratio < 90%"
```

#### Performance Testing
```bash
# Simple load testing with vegeta
echo "GET https://api.djinn.finance/health" | vegeta attack \
  -duration=30s \
  -rate=100/s | vegeta report

# Database query analysis
cat > analyze_queries.sql << EOF
SELECT 
  query,
  calls,
  mean_exec_time,
  total_exec_time,
  (total_exec_time / sum(total_exec_time) OVER ()) * 100 as percentage
FROM pg_stat_statements
ORDER BY total_exec_time DESC
LIMIT 20;
EOF
```

### 8. Progressive Enhancement Triggers

```yaml
enhancement_triggers:
  add_redis_cache:
    when:
      - "Database CPU > 60% consistently"
      - "Same queries repeated > 100/minute"
      - "Response times > 200ms p95"
    
  add_read_replica:
    when:
      - "Read/write ratio > 10:1"
      - "Database CPU > 80%"
      - "Backup impact on performance"
    
  add_cdn:
    when:
      - "Static content > 50% of bandwidth"
      - "Global users with latency > 200ms"
      - "Cost justified by scale"
    
  add_queue_system:
    when:
      - "Long-running operations > 5 seconds"
      - "Timeout issues with synchronous operations"
      - "Need for retry logic"
    
  move_to_microservices:
    when:
      - "Team size > 20 developers"
      - "Clear bounded contexts"
      - "Different scaling needs per service"
    probably: "Never for this application"
```

## Consequences

### Positive
- **Simplicity**: Single server can handle 10K+ users with optimization
- **Cost-effective**: Stay under €20/month for most growth phases
- **Maintainable**: Small team can manage without DevOps expertise
- **Progressive**: Add complexity only when metrics justify it
- **Debuggable**: Simple architecture is easier to troubleshoot
- **Fast deployment**: Changes deploy in seconds, not minutes
- **Strong consistency**: No cache invalidation nightmares

### Negative
- **Vertical scaling limits**: Eventually hit CPU/RAM ceiling
- **Single point of failure**: One server = one failure point
- **Manual scaling**: Requires planning and brief downtime
- **Limited geographic distribution**: Single region initially
- **No automatic failover**: Manual intervention for failures

### Risks
- **Sudden viral growth**: May overwhelm single server
  - *Mitigation*: Have scaling runbook ready, use feature flags to limit access
- **Database becomes bottleneck**: All load on single PostgreSQL
  - *Mitigation*: Monitor metrics, prepare read replica setup
- **Technical debt**: Simple solutions may need refactoring later
  - *Mitigation*: Clean architecture from start, making changes easier
- **Monitoring blind spots**: Self-hosted monitoring can fail with server
  - *Mitigation*: External uptime monitoring (UptimeRobot free tier)

## Alternatives Considered

### Option A: Multi-Layer Caching Architecture
- **Description**: Application cache → Redis → CDN from day one
- **Pros**: Maximum performance, sub-100ms globally, handles viral growth
- **Cons**: Complex cache invalidation, 3x operational overhead, €50+/month
- **Reason for not choosing**: Over-engineered for MVP, adds unnecessary complexity

### Option B: Event-Driven CQRS Architecture  
- **Description**: Separate read/write models with event sourcing
- **Pros**: Perfect audit trail, unlimited read scaling, event replay capability
- **Cons**: High complexity, eventual consistency challenges, requires expertise
- **Reason for not choosing**: Massive complexity increase for uncertain benefits

### Option C: Serverless Architecture
- **Description**: API Gateway → Lambda → DynamoDB
- **Pros**: True auto-scaling, pay per request, zero maintenance
- **Cons**: Vendor lock-in, cold starts, expensive at scale, complex local development
- **Reason for not choosing**: Cold starts unacceptable for financial app, PostgreSQL requirement

### Option D: Kubernetes from Start
- **Description**: K3s on Hetzner with auto-scaling
- **Pros**: Cloud-native, portable, industry standard
- **Cons**: Operational complexity, expertise required, overkill for small team
- **Reason for not choosing**: Team size doesn't justify Kubernetes complexity

## Implementation Notes

### Migration Strategy
1. **Week 1**: Implement database optimizations and connection pooling
2. **Week 2**: Add application-level performance monitoring
3. **Week 3**: Implement simple in-memory caching for hot paths
4. **Week 4**: Add rate limiting and response optimization
5. **Week 5**: Performance testing and query optimization
6. **Week 6**: Documentation and runbooks for scaling

### Testing Approach
```yaml
performance_testing:
  unit_level:
    - Query execution time tests
    - Cache hit/miss ratio tests
    - Rate limiter accuracy tests
    
  integration_level:
    - API endpoint latency tests
    - Database connection pool tests
    - End-to-end transaction tests
    
  load_testing:
    tools: ["vegeta", "k6", "pgbench"]
    scenarios:
      - Normal load (100 req/s)
      - Peak load (1000 req/s)
      - Sustained load (500 req/s for 1 hour)
      - Spike test (0 to 2000 req/s)
    
  chaos_testing:
    scenarios:
      - Database connection loss
      - Redis cache failure
      - High memory pressure
      - Network latency injection
```

### Monitoring and Success Metrics
```yaml
success_metrics:
  performance:
    - API p95 latency < 200ms
    - Database query p95 < 50ms
    - Cache hit ratio > 80%
    - Error rate < 0.1%
    
  scalability:
    - Support 100 concurrent users on CX22
    - Support 1000 concurrent users on CX32
    - Database CPU < 60% under normal load
    - Memory usage < 80% under peak load
    
  operational:
    - Deployment time < 1 minute
    - Time to scale vertically < 10 minutes
    - Incident detection < 1 minute
    - Incident resolution < 1 hour
    
  cost:
    - Infrastructure cost < €20/month until 10K users
    - Cost per user < €0.002/month
    - Monitoring cost: €0 (self-hosted)
```

### Scaling Runbook
```bash
#!/bin/bash
# Emergency scaling procedure

# 1. Check current metrics
echo "=== Current System Status ==="
ssh djinn-prod "top -bn1 | head -5"
ssh djinn-prod "df -h"
ssh djinn-prod "psql -c 'SELECT count(*) FROM pg_stat_activity;'"

# 2. Backup current data
echo "=== Creating Backup ==="
ssh djinn-prod "pg_dump djinn_prod | gzip > backup-$(date +%Y%m%d).sql.gz"

# 3. Resize Hetzner server (via CLI or web console)
echo "=== Resize Instructions ==="
echo "1. Go to Hetzner Cloud Console"
echo "2. Select server > Resize"
echo "3. Choose new size (CX32 or CX42)"
echo "4. Enable 'Create snapshot before resize'"
echo "5. Click 'Resize now'"
echo "6. Wait 5-10 minutes for completion"

# 4. Verify new resources
echo "=== Verify New Resources ==="
ssh djinn-prod "free -h"
ssh djinn-prod "nproc"

# 5. Adjust PostgreSQL settings for new RAM
ssh djinn-prod "sudo -u postgres psql -c 'ALTER SYSTEM SET shared_buffers = \"2GB\";'"
ssh djinn-prod "sudo systemctl restart postgresql"

# 6. Update application configuration if needed
ssh djinn-prod "vim /opt/djinn/.env"
ssh djinn-prod "sudo systemctl restart djinn-api"

# 7. Verify application health
curl -f https://api.djinn.finance/health || echo "Health check failed!"
```

## References
- [PostgreSQL Performance Tuning](https://wiki.postgresql.org/wiki/Tuning_Your_PostgreSQL_Server)
- [PgBouncer Configuration Guide](https://www.pgbouncer.org/config.html)
- [Go Database Connection Best Practices](https://golang.org/doc/database/sql)
- [Hetzner Cloud Scaling Documentation](https://docs.hetzner.cloud/cloud/servers/resize)
- [The Twelve-Factor App](https://12factor.net)
- ADR-20250819: Deployment Architecture
- ADR-20250819: Data Architecture and Schema Design
- ADR-20250819: API Design Standards
- ADR-20250120: Monitoring and Observability

## Decision Makers
- Author: Archie (System Architect)
- Reviewers: [Pending]
- Approver: [Pending]
- Date: 2025-01-20

## Revision History
- 2025-01-20: Initial draft focusing on simplicity and progressive enhancement