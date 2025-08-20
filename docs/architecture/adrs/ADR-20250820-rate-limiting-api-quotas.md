# ADR-20250820: Rate Limiting & API Quotas

## Status
Accepted

## Context

We need to implement comprehensive rate limiting and API quota management for our Go GraphQL API to:
- Prevent API abuse and ensure fair usage across users
- Support different user tiers (free vs premium) with differentiated quotas
- Handle distributed environments with multiple backend instances
- Provide GraphQL-specific protections (query complexity, depth limiting)
- Maintain cost-effectiveness for MVP scale (100K users, 5-10% premium conversion)
- Ensure low latency overhead (<0.5ms per request)

### Current System Architecture
- Go backend with GraphQL (gqlgen framework)
- Firebase Auth for user authentication and authorization
- Distributed system with multiple backend instances
- Target scale: 100K users with 5-10% premium conversion
- Budget constraint: ~$500-1000/month infrastructure

### Business Requirements
1. **User Identification**: Leverage Firebase Auth tokens for user-based rate limiting
2. **Tier-Based Quotas**: Different limits for free and premium users
   - Free tier: 100 requests/hour
   - Premium tier: 1000 requests/hour
3. **GraphQL Protection**: Query complexity and depth limiting
4. **Distributed Support**: Consistent rate limiting across multiple backend instances
5. **Graceful Degradation**: Continue operation if Redis becomes unavailable
6. **Monitoring**: Comprehensive tracking and alerting capabilities
7. **Cost Optimization**: Effective solution within MVP budget constraints

### Technical Requirements
- Integration with gqlgen GraphQL middleware
- Redis-based distributed state management
- Firebase custom claims for tier identification
- Query complexity calculation and enforcement
- Standard rate limit headers in HTTP responses
- Prometheus metrics for monitoring and alerting

## Decision

We will implement a **hybrid rate limiting and API quota system** combining:

1. **Primary Solution**: Ulule Limiter with Redis backend for distributed rate limiting
2. **GraphQL Layer**: Query complexity analysis and depth limiting
3. **Fallback Mechanism**: golang.org/x/time/rate for Redis failure scenarios
4. **Monitoring Stack**: Prometheus metrics with comprehensive alerting

### Architecture Components

#### 1. Distributed Rate Limiting (Ulule Limiter + Redis)
```go
import (
    "github.com/ulule/limiter/v3"
    "github.com/ulule/limiter/v3/drivers/store/redis"
    "github.com/ulule/limiter/v3/drivers/middleware/gin"
    goredis "github.com/go-redis/redis/v8"
)

// Initialize Redis client
redisClient := goredis.NewClient(&goredis.Options{
    Addr:     "localhost:6379",
    Password: "", // set password
    DB:       0,  // use default DB
})

// Rate configuration per user tier
var TierConfigs = map[string]limiter.Rate{
    "free": {
        Period: 1 * time.Hour,
        Limit:  100,
    },
    "premium": {
        Period: 1 * time.Hour,
        Limit:  1000,
    },
}

// Create Redis store for Ulule Limiter
store, err := redis.NewStoreWithOptions(redisClient, limiter.StoreOptions{
    Prefix:   "rate_limit",
    MaxRetry: 3,
})
if err != nil {
    log.Fatal(err)
}

// Create limiter instance for a specific user tier
func createUserLimiter(userTier string) *limiter.Limiter {
    rate := TierConfigs[userTier]
    return limiter.New(store, rate)
}
```

#### 2. GraphQL Query Complexity Analysis
```go
type ComplexityLimits struct {
    MaxDepth:      10
    MaxComplexity: 1000
    FieldWeights: map[string]int{
        "users":         10,  // Expensive list queries
        "transactions":  5,   // Medium cost queries
        "user":          1,   // Single entity queries
    }
}

// Tier-specific complexity limits
var TierComplexityLimits = map[string]ComplexityLimits{
    "free":    {MaxDepth: 8, MaxComplexity: 500},
    "premium": {MaxDepth: 12, MaxComplexity: 2000},
}
```

#### 3. Firebase Auth Integration
```go
func getUserTierFromFirebase(ctx context.Context, userID string) (string, error) {
    token, err := firebase.VerifyIDToken(ctx, userID)
    if err != nil {
        return "free", err // Default to free tier on error
    }
    
    if tier, ok := token.Claims["subscription_tier"].(string); ok {
        return tier, nil
    }
    return "free", nil // Default tier
}
```

#### 4. Middleware Integration with Ulule Limiter
```go
import (
    "github.com/ulule/limiter/v3"
    mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
    "golang.org/x/time/rate" // For fallback
)

// In-memory fallback limiters per user
var fallbackLimiters = make(map[string]*rate.Limiter)
var fallbackMutex sync.RWMutex

func RateLimitMiddleware(store limiter.Store) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID := getUserIDFromAuth(c)
        userTier, err := getUserTierFromFirebase(c.Request.Context(), userID)
        if err != nil {
            logError("Failed to get user tier", err)
            userTier = "free" // Default to most restrictive
        }
        
        // Generate rate limit key for this user
        key := fmt.Sprintf("user:%s", userID)
        
        // Get tier configuration
        tierConfig := TierConfigs[userTier]
        
        // Create Ulule limiter instance for this request
        rateLimiter := limiter.New(store, tierConfig)
        
        // Create context with user key
        limiterContext, err := rateLimiter.Get(c.Request.Context(), key)
        if err != nil {
            // Fallback to in-memory rate limiting using golang.org/x/time/rate
            fallbackMutex.Lock()
            if _, exists := fallbackLimiters[key]; !exists {
                // Create new fallback limiter: tierConfig.Limit requests per tierConfig.Period
                limit := rate.Every(tierConfig.Period / time.Duration(tierConfig.Limit))
                fallbackLimiters[key] = rate.NewLimiter(limit, int(tierConfig.Limit))
            }
            fallbackLimiter := fallbackLimiters[key]
            fallbackMutex.Unlock()
            
            if !fallbackLimiter.Allow() {
                c.JSON(429, gin.H{
                    "error": "Rate limit exceeded (fallback mode)",
                    "tier": userTier,
                })
                c.Abort()
                return
            }
            // Continue with degraded service
            c.Header("X-RateLimit-Mode", "fallback")
        } else {
            // Set standard rate limit headers
            c.Header("X-RateLimit-Limit", strconv.FormatInt(limiterContext.Limit, 10))
            c.Header("X-RateLimit-Remaining", strconv.FormatInt(limiterContext.Remaining, 10))
            c.Header("X-RateLimit-Reset", strconv.FormatInt(limiterContext.Reset, 10))
            c.Header("X-RateLimit-Tier", userTier)
            
            if limiterContext.Reached {
                c.JSON(429, gin.H{
                    "error": "Rate limit exceeded",
                    "limit": limiterContext.Limit,
                    "reset": limiterContext.Reset,
                    "tier": userTier,
                })
                c.Abort()
                return
            }
        }
        
        c.Next()
    }
}
```

#### 5. GraphQL Complexity Middleware
```go
func ComplexityMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        userTier := c.GetHeader("X-User-Tier")
        complexityLimit := TierComplexityLimits[userTier]
        
        // Parse and analyze GraphQL query
        if isGraphQLRequest(c) {
            query := extractGraphQLQuery(c)
            complexity := calculateQueryComplexity(query, complexityLimit.FieldWeights)
            
            if complexity > complexityLimit.MaxComplexity {
                c.JSON(400, gin.H{
                    "error": "Query complexity exceeds limit",
                    "complexity": complexity,
                    "limit": complexityLimit.MaxComplexity,
                    "tier": userTier,
                })
                c.Abort()
                return
            }
            
            // Add complexity headers for monitoring
            c.Header("X-Query-Complexity", strconv.Itoa(complexity))
            c.Header("X-Complexity-Limit", strconv.Itoa(complexityLimit.MaxComplexity))
        }
        
        c.Next()
    }
}
```

### Implementation Strategy

#### Phase 1: Core Rate Limiting (Week 1-2)
- Deploy Redis infrastructure (managed service)
- Implement Ulule Limiter with basic user identification
- Add Firebase Auth integration for user tier detection
- Implement fallback mechanism with golang.org/x/time/rate
- Basic monitoring and logging

#### Phase 2: GraphQL Integration (Week 3)
- Implement query complexity analysis
- Add GraphQL-specific rate limiting
- Enhanced error responses with detailed quota information
- Rate limit headers implementation

#### Phase 3: Monitoring & Alerting (Week 4)
- Prometheus metrics integration
- Grafana dashboards for quota utilization
- Alerting rules for quota abuse and system health
- Performance optimization and load testing

### Cost Analysis

#### Infrastructure Costs (Monthly)
- **Redis Hosting**: $15-30/month (AWS ElastiCache t3.micro or DigitalOcean Managed Redis)
- **Monitoring**: $0 (Prometheus + Grafana self-hosted) or $20/month (managed)
- **Total Additional Cost**: $15-50/month

#### Performance Projections
- **Redis Memory Usage**: ~100 bytes per active user = 10MB for 100K users
- **Rate Limiter Latency**: <0.5ms per request (within requirement)
- **Redis Operations**: ~300 ops/second at peak (well within Redis limits)
- **Cost per Request**: ~$0.0000005 (negligible)

## Consequences

### Positive
- **Scalable Architecture**: Supports distributed deployment across multiple instances
- **Business Model Support**: Enables tier-based pricing with clear quota differentiation
- **Performance Optimized**: <0.5ms overhead meets performance requirements
- **Cost Effective**: $15-50/month additional cost fits within MVP budget
- **GraphQL Aware**: Protects against expensive query attacks
- **Flexible Configuration**: Easy adjustment of quotas and complexity limits
- **Comprehensive Monitoring**: Full visibility into usage patterns and system health
- **Graceful Degradation**: Continues operation during Redis outages

### Negative
- **Infrastructure Dependency**: Adds Redis as a critical system component
- **Operational Complexity**: Requires monitoring and maintenance of additional services
- **Network Latency**: Redis round-trip adds minimal but measurable latency
- **Configuration Management**: Need to manage tier configurations and complexity weights

### Risks

#### Technical Risks
1. **Redis Single Point of Failure**
   - *Mitigation*: Fallback to in-memory rate limiting with degraded functionality
   - *Monitoring*: Redis connection health alerts

2. **Network Latency to Redis**
   - *Mitigation*: Deploy Redis in same region/VPC as application servers
   - *Monitoring*: Track p95/p99 latency metrics

3. **Query Complexity Calculation Overhead**
   - *Mitigation*: Cache complexity calculations for repeated queries
   - *Monitoring*: Track complexity calculation time

#### Business Risks
1. **Quota Abuse Detection**
   - *Mitigation*: Implement graduated responses and user communication
   - *Monitoring*: Alert on unusual usage patterns

2. **Tier Migration Complexity**
   - *Mitigation*: Implement real-time tier updates via Firebase claims
   - *Testing*: Comprehensive tier transition scenarios

## Alternatives Considered

### Option A: Database-Backed Rate Limiting
- **Description**: Store rate limit counters in PostgreSQL
- **Pros**: Leverages existing database infrastructure, persistent across all restarts
- **Cons**: Higher latency (5-50ms), database load concerns, complex implementation
- **Reason for Rejection**: Latency requirements cannot be met

### Option B: In-Memory Only (golang.org/x/time/rate)
- **Description**: Use standard library rate limiting without distributed state
- **Pros**: Lowest latency, no external dependencies, simple implementation
- **Cons**: Not suitable for distributed systems, lost state on restarts
- **Reason for Rejection**: Distributed architecture requirement

### Option C: Third-Party API Rate Limiting Service
- **Description**: Use external service like Kong, Envoy, or cloud provider rate limiting
- **Pros**: Managed solution, minimal implementation effort
- **Cons**: Additional cost ($100+/month), external dependency, less customization
- **Reason for Rejection**: Cost exceeds budget constraints

## Implementation Notes

### Migration Strategy
1. **Gradual Rollout**: Start with monitoring-only mode to establish baselines
2. **Tier-Based Deployment**: Enable rate limiting for free tier first, then premium
3. **Feature Flags**: Use feature toggles for safe rollback capability
4. **Load Testing**: Comprehensive testing before production deployment

### Rollback Plan
1. **Immediate**: Disable rate limiting via feature flag
2. **Redis Failure**: Automatic fallback to in-memory limiting
3. **Performance Issues**: Adjust limits or disable complexity analysis
4. **Business Impact**: Emergency quota increases via configuration

### Testing Approach
- **Unit Tests**: Rate limiter logic, complexity calculation
- **Integration Tests**: Redis connectivity, Firebase Auth integration
- **Load Tests**: Performance under expected traffic patterns
- **Chaos Tests**: Redis failure scenarios, network partitions

### Success Metrics
- **Performance**: p95 latency < 0.5ms for rate limiting operations
- **Availability**: 99.9% uptime for rate limiting service
- **Business**: <1% false positive rate limiting (legitimate users blocked)
- **Cost**: Stay within $50/month additional infrastructure cost

## References

- [Ulule Limiter Documentation](https://github.com/ulule/limiter)
- [GraphQL Query Complexity Analysis](https://graphql.org/learn/validation/)
- [Firebase Auth Custom Claims](https://firebase.google.com/docs/auth/admin/custom-claims)
- [Redis Rate Limiting Patterns](https://redis.io/docs/latest/develop/use/patterns/distributed-locks/)
- [GitHub API Rate Limiting](https://docs.github.com/en/rest/using-the-rest-api/rate-limits-for-the-rest-api)
- [Stripe API Rate Limiting Best Practices](https://stripe.com/docs/rate-limits)

## Decision Makers

- **Author**: ADR Manager
- **Reviewers**: [Pending]
- **Approver**: [Pending]
- **Date**: 2025-08-20

## Revision History

- **2025-08-20**: Initial comprehensive ADR draft with research findings