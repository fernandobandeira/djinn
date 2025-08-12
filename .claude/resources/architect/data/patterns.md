# Architectural Patterns Reference

## System Architecture Patterns

### Monolithic Architecture
**When to Use:**
- Small to medium applications
- Tight deadlines
- Limited team size
- Simple deployment requirements

**Trade-offs:**
- ✅ Simple to develop and deploy
- ✅ Easy debugging and testing
- ✅ Good performance (no network calls)
- ❌ Scaling limitations
- ❌ Technology lock-in
- ❌ Large codebase complexity

### Microservices Architecture
**When to Use:**
- Large, complex applications
- Multiple teams
- Different scaling requirements per component
- Need for technology diversity

**Trade-offs:**
- ✅ Independent deployment and scaling
- ✅ Technology flexibility
- ✅ Fault isolation
- ❌ Operational complexity
- ❌ Network latency
- ❌ Data consistency challenges

### Modular Monolith
**When to Use:**
- Want microservices benefits without complexity
- Clear module boundaries needed
- Future migration path to microservices
- Medium-sized teams

**Trade-offs:**
- ✅ Simpler than microservices
- ✅ Clear boundaries
- ✅ Easier refactoring
- ❌ Still monolithic deployment
- ❌ Shared database
- ❌ Module discipline required

### Serverless Architecture
**When to Use:**
- Event-driven workloads
- Variable traffic patterns
- Rapid prototyping
- Cost optimization focus

**Trade-offs:**
- ✅ No server management
- ✅ Automatic scaling
- ✅ Pay per use
- ❌ Vendor lock-in
- ❌ Cold starts
- ❌ Limited execution time

## Communication Patterns

### Request-Response (Synchronous)
**Use Cases:**
- Real-time interactions
- Simple workflows
- Low latency requirements

**Implementation:**
- REST APIs
- GraphQL
- gRPC

### Event-Driven (Asynchronous)
**Use Cases:**
- Decoupled systems
- High throughput
- Complex workflows
- Eventual consistency acceptable

**Implementation:**
- Message queues (RabbitMQ, SQS)
- Event streams (Kafka, Kinesis)
- Webhooks

### Hybrid Communication
**Use Cases:**
- Mix of real-time and batch
- Different consistency requirements
- Gradual migration

**Strategy:**
- Sync for user-facing
- Async for background
- Event sourcing for audit

## Data Architecture Patterns

### Shared Database
**When to Use:**
- Monolithic applications
- Strong consistency required
- Simple transactions
- Small team

**Considerations:**
- Schema coordination
- Scaling limitations
- Coupling risk

### Database per Service
**When to Use:**
- Microservices architecture
- Service autonomy needed
- Different data requirements
- Independent scaling

**Considerations:**
- Data consistency
- Cross-service queries
- Distributed transactions

### CQRS (Command Query Responsibility Segregation)
**When to Use:**
- Different read/write patterns
- Complex queries
- High read load
- Event sourcing

**Implementation:**
```
Commands (Write) → Command DB
                 ↓
              Events
                 ↓
Queries (Read) ← Query DB
```

### Event Sourcing
**When to Use:**
- Audit requirements
- Time travel queries
- Complex domain logic
- Event-driven architecture

**Considerations:**
- Event store design
- Snapshot strategy
- Event versioning
- GDPR compliance

## API Design Patterns

### REST (Representational State Transfer)
**Principles:**
- Resource-based
- Stateless
- Uniform interface
- HTTP verbs

**Best Practices:**
```
GET    /api/v1/users          # List
POST   /api/v1/users          # Create
GET    /api/v1/users/{id}     # Read
PUT    /api/v1/users/{id}     # Update
DELETE /api/v1/users/{id}     # Delete
```

### GraphQL
**When to Use:**
- Complex data requirements
- Mobile applications
- Multiple client types
- Rapid iteration

**Benefits:**
- Single endpoint
- Precise data fetching
- Strong typing
- Self-documenting

### gRPC
**When to Use:**
- Microservices communication
- High performance needs
- Streaming requirements
- Strong contracts

**Features:**
- Protocol buffers
- HTTP/2
- Bidirectional streaming
- Code generation

### API Gateway Pattern
**Purpose:**
- Single entry point
- Cross-cutting concerns
- Request routing
- Protocol translation

**Responsibilities:**
- Authentication
- Rate limiting
- Request/response transformation
- Load balancing
- Caching

## Resilience Patterns

### Circuit Breaker
**Purpose:** Prevent cascading failures

```python
states = ["CLOSED", "OPEN", "HALF_OPEN"]
if failures > threshold:
    state = "OPEN"
    return cached_response
```

### Retry with Exponential Backoff
**Purpose:** Handle transient failures

```python
delay = initial_delay
for attempt in range(max_retries):
    try:
        return make_request()
    except:
        sleep(delay)
        delay *= 2
```

### Bulkhead Pattern
**Purpose:** Isolate failures

```
Service A → [Pool 1] → Database
Service B → [Pool 2] → Database
Service C → [Pool 3] → Database
```

### Timeout Pattern
**Purpose:** Fail fast

```python
with timeout(seconds=5):
    response = make_request()
```

## Deployment Patterns

### Blue-Green Deployment
**Process:**
1. Deploy to green environment
2. Test green environment
3. Switch traffic to green
4. Keep blue as rollback

**Benefits:**
- Zero downtime
- Easy rollback
- Full testing

### Canary Deployment
**Process:**
1. Deploy to small subset
2. Monitor metrics
3. Gradually increase traffic
4. Full rollout or rollback

**Benefits:**
- Risk mitigation
- Real user testing
- Gradual rollout

### Rolling Deployment
**Process:**
1. Update one instance
2. Health check
3. Update next instance
4. Repeat until complete

**Benefits:**
- No additional infrastructure
- Gradual update
- Continuous availability

### Feature Flags
**Purpose:** Decouple deployment from release

```python
if feature_flag("new_feature", user):
    return new_implementation()
else:
    return old_implementation()
```

## Security Patterns

### Defense in Depth
**Layers:**
1. Network security (firewall, VPC)
2. Application security (WAF)
3. Authentication/Authorization
4. Data encryption
5. Audit logging

### Zero Trust Architecture
**Principles:**
- Never trust, always verify
- Least privilege access
- Assume breach
- Verify explicitly

### OAuth 2.0 / OpenID Connect
**Flow Types:**
- Authorization Code (web apps)
- Implicit (SPA) - deprecated
- Client Credentials (M2M)
- PKCE (mobile/SPA)

### API Security
**Best Practices:**
- API keys for identification
- JWT for authentication
- Rate limiting
- Input validation
- HTTPS only

## Caching Patterns

### Cache-Aside (Lazy Loading)
```python
def get_data(key):
    data = cache.get(key)
    if not data:
        data = database.get(key)
        cache.set(key, data)
    return data
```

### Write-Through
```python
def save_data(key, value):
    cache.set(key, value)
    database.save(key, value)
```

### Write-Behind (Write-Back)
```python
def save_data(key, value):
    cache.set(key, value)
    queue_for_write(key, value)  # Async
```

### Refresh-Ahead
```python
def refresh_cache():
    for key in frequently_used:
        if near_expiry(key):
            data = database.get(key)
            cache.set(key, data)
```

## Scalability Patterns

### Horizontal Scaling (Scale Out)
**Approach:** Add more machines

**Considerations:**
- Load balancing
- Session management
- Data consistency
- Service discovery

### Vertical Scaling (Scale Up)
**Approach:** Add more power

**Limitations:**
- Hardware limits
- Single point of failure
- Cost efficiency
- Downtime for upgrades

### Database Sharding
**Strategies:**
- Range-based (by ID range)
- Hash-based (by hash key)
- Geographic (by location)
- Functional (by feature)

### Read Replicas
**Purpose:** Scale read operations

```
Master (Writes)
    ↓
Replica 1 (Reads)
Replica 2 (Reads)
Replica 3 (Reads)
```

## Message Queue Patterns

### Point-to-Point
**Use Case:** Task distribution

```
Producer → Queue → Consumer
```

### Publish-Subscribe
**Use Case:** Event broadcasting

```
Producer → Topic → Subscriber 1
                 → Subscriber 2
                 → Subscriber 3
```

### Request-Reply
**Use Case:** Async request/response

```
Client → Request Queue → Server
      ← Reply Queue ←
```

### Priority Queue
**Use Case:** Task prioritization

```
High Priority    → [][][]
Medium Priority  → [][][]
Low Priority     → [][][]
```

## Anti-Patterns to Avoid

### God Object
**Problem:** Single class doing everything
**Solution:** Split responsibilities

### Chatty APIs
**Problem:** Too many small requests
**Solution:** Batch operations, GraphQL

### Distributed Monolith
**Problem:** Tightly coupled microservices
**Solution:** Proper boundaries, async communication

### Shared Database
**Problem:** Services sharing database
**Solution:** Database per service, API communication

### Premature Optimization
**Problem:** Optimizing without metrics
**Solution:** Measure first, optimize later

## Pattern Selection Guide

### Questions to Ask:
1. What are the scalability requirements?
2. What's the team size and expertise?
3. What are the consistency requirements?
4. What's the expected traffic pattern?
5. What are the latency requirements?
6. What's the deployment complexity tolerance?
7. What are the cost constraints?

### Decision Matrix:
| Requirement | Monolith | Microservices | Serverless |
|------------|----------|---------------|------------|
| Simple | ✅ | ❌ | ✅ |
| Scale | ❌ | ✅ | ✅ |
| Cost | ✅ | ❌ | ✅* |
| Speed | ✅ | ❌ | ✅ |

*Depends on usage pattern

## Remember
- Patterns are tools, not rules
- Context matters more than trends
- Start simple, evolve as needed
- Document pattern decisions as ADRs
- Consider trade-offs honestly