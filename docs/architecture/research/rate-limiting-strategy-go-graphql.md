# ADR-20250820: Rate Limiting Strategy for Go GraphQL APIs

## Status
Proposed

## Context
We need to implement comprehensive rate limiting for our Go GraphQL API to prevent abuse, ensure fair usage, and support different user tiers (free vs premium). The system must handle distributed environments with multiple backend instances while maintaining low latency and being cost-effective for MVP scale (100K users).

### Requirements
- Support per-user rate limiting with Firebase Auth integration
- Handle different tiers (free vs premium users) with different quotas
- Work with distributed systems (multiple backend instances)
- Support GraphQL-specific features (query complexity, depth limiting)
- Low latency overhead (<1ms per request)
- Cost-effective for MVP scale
- Monitoring and alerting capabilities

## Decision
After comprehensive research into Go rate limiting libraries and patterns, we recommend implementing a **hybrid approach** combining:

1. **Primary**: Ulule Limiter with Redis backend for distributed rate limiting
2. **Secondary**: GraphQL query complexity analysis for query-based limiting
3. **Fallback**: In-memory golang.org/x/time/rate for Redis failures

## Analysis

### Go Rate Limiting Libraries Comparison

#### 1. golang.org/x/time/rate (Standard Library)
**Strengths:**
- Part of Go's extended standard library
- Token bucket algorithm with efficient implementation
- Excellent performance for single-instance scenarios
- Simple API: `rate.NewLimiter(rate.Limit(r), b)`
- Zero external dependencies

**Weaknesses:**
- In-memory only - doesn't support distributed rate limiting
- No built-in persistence across restarts
- No user-based rate limiting out of the box

**Best for:** Single-instance applications or as a fallback mechanism

#### 2. Uber-go/ratelimit
**Strengths:**
- Leaky bucket implementation (smooth rate limiting)
- Very simple API: `ratelimit.New(100)` for 100 ops/second
- High performance with minimal allocations
- Battle-tested at Uber scale

**Weaknesses:**
- In-memory only - not suitable for distributed systems
- Fixed rate limiting (no burst handling)
- No user-based differentiation

**Best for:** High-performance single-node services with consistent rate requirements

#### 3. Ulule Limiter ⭐ (Recommended)
**Strengths:**
- Multiple storage backends (Redis, Memory, etc.)
- Built-in middleware for popular frameworks
- User-based rate limiting with flexible key generation
- Support for different time windows (second, minute, hour, day)
- Excellent documentation and examples
- Active maintenance and community

**Features:**
```go
// Flexible rate configuration
rate := limiter.Rate{
    Period: 1 * time.Hour,
    Limit:  1000,
}

// Multiple storage options
redisStore := redis.NewStore(client)
memoryStore := memory.NewStore()

// User-based limiting with custom key generation
instance := limiter.New(store, rate, 
    limiter.WithClientIPHeader("X-Real-IP"),
    limiter.WithKeyGenerator(func(c *gin.Context) string {
        return getUserID(c) // Custom user identification
    }))
```

**Cost Analysis:**
- Redis memory usage: ~100 bytes per active user
- For 100K users: ~10MB Redis memory
- Redis hosting cost: $5-20/month (depending on provider)

#### 4. juju/ratelimit
**Strengths:**
- Token bucket with customizable fill rates
- Good performance characteristics

**Weaknesses:**
- Less actively maintained
- Limited documentation
- In-memory only

### GraphQL-Specific Rate Limiting Strategies

#### Query Complexity Analysis
GraphQL queries can vary dramatically in computational cost. A depth-limited approach:

```go
type ComplexityConfig struct {
    MaxDepth:      10
    MaxComplexity: 1000
    FieldWeights: map[string]int{
        "users":     10,
        "user":      1,
        "posts":     5,
        "comments":  2,
    }
}
```

#### Query Cost Analysis Patterns
1. **Static Analysis**: Pre-calculate query costs
2. **Dynamic Analysis**: Track actual execution time
3. **Hybrid**: Combine both approaches

### Distributed Rate Limiting Approaches

#### Redis-Based Implementation
**Sliding Window Counter Pattern:**
```lua
-- Redis Lua script for atomic rate limiting
local key = KEYS[1]
local window = tonumber(ARGV[1])
local limit = tonumber(ARGV[2])
local current = redis.call('INCR', key)
if current == 1 then
    redis.call('EXPIRE', key, window)
end
return {current, limit}
```

**Performance Characteristics:**
- Latency: ~0.1-0.5ms for local Redis, ~1-2ms for hosted Redis
- Throughput: 100K+ ops/second per Redis instance
- Memory efficient: O(active_users) space complexity

#### Database-Backed Rate Limiting
**Pros:**
- Persistent across all restarts
- Can leverage existing database infrastructure
- Complex query capabilities

**Cons:**
- Higher latency (5-50ms typical)
- Database load and scaling concerns
- More complex implementation

#### In-Memory with Synchronization
**Pros:**
- Lowest latency (<0.1ms)
- No external dependencies

**Cons:**
- Not suitable for distributed systems
- Lost state on restarts
- Scaling limitations

### User Tier Management

#### Firebase Auth Integration
```go
type UserTier struct {
    TierID      string
    RequestsPerHour int
    QueryComplexityLimit int
    Features    []string
}

func getUserTier(ctx context.Context, userID string) (*UserTier, error) {
    // Firebase custom claims or Firestore lookup
    token, err := firebase.VerifyIDToken(ctx, userID)
    if err != nil {
        return nil, err
    }
    
    if tier, ok := token.Claims["tier"].(string); ok {
        return getTierConfig(tier), nil
    }
    return getDefaultTier(), nil
}
```

#### Tier-Based Rate Limiting
```go
// Rate limiting configuration per tier
var TierConfigs = map[string]limiter.Rate{
    "free": {
        Period: 1 * time.Hour,
        Limit:  100,
    },
    "premium": {
        Period: 1 * time.Hour,
        Limit:  10000,
    },
    "enterprise": {
        Period: 1 * time.Hour,
        Limit:  100000,
    },
}
```

### Implementation Architecture

#### Recommended Stack
```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   GraphQL       │    │   Rate Limiter   │    │    Redis        │
│   Middleware    │◄──►│   (Ulule)        │◄──►│    Cluster      │
└─────────────────┘    └──────────────────┘    └─────────────────┘
         │                       │                       │
         ▼                       ▼                       ▼
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   Query         │    │   Fallback       │    │   Monitoring    │
│   Complexity    │    │   (golang/rate)  │    │   & Alerts      │
└─────────────────┘    └──────────────────┘    └─────────────────┘
```

#### gqlgen Integration
```go
func rateLimitMiddleware(limiterInstance *limiter.Limiter) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID := getUserIDFromContext(c)
        userTier := getUserTier(c, userID)
        
        // Dynamic rate based on user tier
        dynamicRate := getTierRate(userTier)
        
        // Check rate limit
        context, err := limiterInstance.Get(c, userID)
        if err != nil {
            c.JSON(500, gin.H{"error": "Rate limiter error"})
            c.Abort()
            return
        }
        
        if context.Reached {
            c.Header("X-RateLimit-Limit", strconv.FormatInt(context.Limit, 10))
            c.Header("X-RateLimit-Remaining", strconv.FormatInt(context.Remaining, 10))
            c.Header("X-RateLimit-Reset", strconv.FormatInt(context.Reset, 10))
            c.JSON(429, gin.H{"error": "Rate limit exceeded"})
            c.Abort()
            return
        }
        
        c.Next()
    }
}
```

### Monitoring and Observability

#### Key Metrics
1. **Rate Limit Metrics:**
   - Requests allowed/blocked per user/tier
   - Rate limit utilization by tier
   - Average time to rate limit exhaustion

2. **Performance Metrics:**
   - Rate limiter latency (p50, p95, p99)
   - Redis response time
   - Fallback activation rate

3. **Business Metrics:**
   - Tier upgrade conversion rate
   - Rate limit impact on user engagement
   - Cost per rate-limited request

#### Alerting Strategy
```yaml
alerts:
  - name: HighRateLimitUsage
    condition: rate_limit_utilization > 0.8
    for: 5m
    
  - name: RateLimiterLatency
    condition: rate_limiter_p95_latency > 10ms
    for: 2m
    
  - name: RedisUnavailable
    condition: redis_connection_failures > 5
    for: 1m
```

### Cost Analysis for MVP Scale (100K Users)

#### Redis Infrastructure Costs
- **Memory Requirements:** ~10MB for active user data
- **Hosted Redis Options:**
  - AWS ElastiCache: $15-30/month (t3.micro)
  - Redis Cloud: $5-15/month (30MB plan)
  - DigitalOcean Managed Redis: $15/month

#### Performance Projections
- **Expected Load:** 10K concurrent users, 100 requests/user/hour
- **Peak RPS:** ~300 requests/second
- **Rate Limiter Overhead:** <0.5ms per request
- **Redis Ops:** ~300 ops/second (well within limits)

### Implementation Phases

#### Phase 1: Basic Rate Limiting (Week 1-2)
- Implement Ulule limiter with Redis backend
- Basic user identification via Firebase Auth
- Simple tier-based rate limiting (free/premium)
- Basic monitoring

#### Phase 2: GraphQL Integration (Week 3-4)
- Query complexity analysis
- GraphQL-specific rate limiting
- Enhanced error handling and user feedback

#### Phase 3: Advanced Features (Week 5-6)
- Sophisticated monitoring and alerting
- Rate limit analytics dashboard
- Performance optimization
- Load testing and tuning

#### Phase 4: Production Hardening (Week 7-8)
- Comprehensive error handling
- Graceful degradation strategies
- Documentation and runbooks

## Consequences

### Positive
- Scalable solution supporting distributed architecture
- Flexible tier-based rate limiting for business model
- Low latency impact on API performance
- Cost-effective for MVP scale
- GraphQL-aware rate limiting capabilities
- Built-in monitoring and observability

### Negative
- Additional complexity with Redis dependency
- Need for Redis infrastructure management
- Potential single point of failure (mitigated by fallback)
- Operational overhead for monitoring and maintenance

### Risks and Mitigations
1. **Redis Failure:** Fallback to in-memory rate limiting
2. **Network Latency:** Local Redis deployment option
3. **Cost Scaling:** Monitoring and alerting for usage spikes
4. **Complexity:** Comprehensive documentation and testing

## Related Decisions
- ADR-20250815: Authentication and Authorization Strategy
- ADR-20250818: Monitoring and Observability Strategy
- ADR-20250819: Performance and Scalability Strategy

## References
- [Ulule Limiter Documentation](https://github.com/ulule/limiter)
- [golang.org/x/time/rate Package](https://pkg.go.dev/golang.org/x/time/rate)
- [GraphQL Query Complexity Analysis](https://graphql.org/learn/validation/)
- [Redis Rate Limiting Patterns](https://redis.io/docs/latest/develop/use/patterns/distributed-locks/)
- [GitHub API Rate Limiting Best Practices](https://docs.github.com/en/rest/using-the-rest-api/rate-limits-for-the-rest-api)