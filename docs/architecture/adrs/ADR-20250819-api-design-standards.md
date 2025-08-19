# ADR-20250819: API Design Standards

## Status
Accepted

## Context
Djinn requires a well-designed API that balances:
- Type safety across mobile and backend (Flutter + Go)
- Efficient data fetching for financial dashboards
- Offline-first mobile synchronization
- Cost optimization for MVP with ability to scale
- Real-time updates for critical financial data
- Complex queries (spending analytics, budget tracking)
- Receipt OCR processing and itemization workflows
- Multi-tenant data isolation
- Regulatory compliance for financial data access

### Constraints
- GraphQL with gqlgen chosen (tech stack ADR)
- Must support offline-first mobile architecture
- Limited MVP budget ($500K-1M total, ~$500-1000/month infrastructure)
- Small team requiring high developer productivity
- 100K users target with 5-10% premium conversion
- Sub-100ms response time for critical queries
- Must handle burst traffic (month-end, tax season)

## Decision

### 1. GraphQL Schema Design Principles

#### Schema-First Development
```graphql
# Core design principles
"""
Financial amounts are always in minor units (cents)
and include currency information
"""
type Money {
  "Amount in minor units (cents)"
  amountMinor: Int!
  "ISO 4217 currency code"
  currency: String!
  "Formatted for display (e.g., $1,234.56)"
  formatted: String!
}

"""
All lists use cursor-based pagination for cost efficiency
30-60% database cost reduction vs offset pagination
"""
type TransactionConnection {
  edges: [TransactionEdge!]!
  pageInfo: PageInfo!
  totalCount: Int @premium # Only for premium users
}

type TransactionEdge {
  cursor: String!
  node: Transaction!
}

type PageInfo {
  hasNextPage: Boolean!
  hasPreviousPage: Boolean!
  startCursor: String
  endCursor: String
}
```

#### Domain-Driven Types
```graphql
# Union types for better type safety
union TransactionResult = Transaction | TransactionError

type TransactionError {
  code: ErrorCode!
  message: String!
  field: String
}

enum ErrorCode {
  INSUFFICIENT_FUNDS
  DUPLICATE_TRANSACTION
  INVALID_CATEGORY
  ACCOUNT_LOCKED
  RATE_LIMITED
}

# Interfaces for shared behavior
interface Timestamped {
  createdAt: DateTime!
  updatedAt: DateTime!
}

interface Owned {
  user: User!
}

# Specific types inherit interfaces
type Transaction implements Timestamped & Owned {
  id: ID!
  amount: Money!
  description: String!
  user: User!
  createdAt: DateTime!
  updatedAt: DateTime!
}
```

### 2. Query Patterns and Optimization

#### Cost-Optimized Query Design
```graphql
type Query {
  # Efficient single-resource fetches
  transaction(id: ID!): Transaction @cacheControl(maxAge: 300)
  account(id: ID!): Account @cacheControl(maxAge: 60)
  
  # Paginated lists (cursor-based for 30% cost savings)
  transactions(
    first: Int = 20  # Default small page size
    after: String
    filter: TransactionFilter
    orderBy: TransactionOrder
  ): TransactionConnection!
  
  # Aggregated data (pre-computed for efficiency)
  monthlySpending(
    month: String!
    accountIds: [ID!]
  ): MonthlySpendingSummary! @cacheControl(maxAge: 3600)
  
  # Premium features with higher cost
  advancedAnalytics(
    dateRange: DateRange!
    groupBy: GroupBy!
  ): AnalyticsResult! @premium @complexity(points: 100)
}

# Filtering to reduce data transfer
input TransactionFilter {
  accountIds: [ID!]
  categoryIds: [ID!]
  amountRange: AmountRange
  dateRange: DateRange
  searchTerm: String @validate(maxLength: 100)
}
```

#### DataLoader Pattern for N+1 Prevention
```go
// Cost-efficient batch loading
type Loaders struct {
    UserLoader      *dataloader.Loader[string, *model.User]
    AccountLoader   *dataloader.Loader[string, *model.Account]
    CategoryLoader  *dataloader.Loader[string, *model.Category]
    TransactionLoader *dataloader.Loader[string, *model.Transaction]
}

// Batch function reduces database calls by 90%
func batchAccounts(ctx context.Context, accountIDs []string) ([]*model.Account, []error) {
    // Single query instead of N queries
    accounts, err := db.GetAccountsByIDs(ctx, accountIDs)
    if err != nil {
        return nil, []error{err}
    }
    
    // Map results back to original order
    accountMap := make(map[string]*model.Account)
    for _, account := range accounts {
        accountMap[account.ID] = account
    }
    
    results := make([]*model.Account, len(accountIDs))
    for i, id := range accountIDs {
        results[i] = accountMap[id]
    }
    
    return results, nil
}
```

### 3. Mutation Patterns

#### Idempotent Mutations with Client-Generated IDs
```graphql
type Mutation {
  # Client provides idempotency key
  createTransaction(
    input: CreateTransactionInput!
    idempotencyKey: ID! # Client-generated UUIDv7
  ): TransactionPayload!
  
  # Batch operations for efficiency
  categorizeTransactions(
    transactionIds: [ID!]!
    categoryId: ID!
  ): CategorizePayload! @validate(maxItems: 100)
  
  # Async operations for heavy processing
  processReceipt(
    image: Upload!
    transactionId: ID
  ): ProcessReceiptPayload!
}

# Consistent payload pattern
type TransactionPayload {
  success: Boolean!
  transaction: Transaction
  errors: [UserError!]
  userMessage: String # User-friendly message
}

type UserError {
  field: String!
  message: String!
  code: ErrorCode!
}
```

#### Optimistic Updates Support
```graphql
# Include predicted results for optimistic UI
type CreateTransactionInput {
  accountId: ID!
  amount: MoneyInput!
  description: String!
  categoryId: ID
  occurredAt: DateTime!
  
  # For optimistic updates
  clientMutationId: String
  predictedBalance: MoneyInput
}

type MoneyInput {
  amountMinor: Int!
  currency: String!
}
```

### 4. Real-time Subscriptions (Post-MVP)

#### Selective Subscriptions for Cost Control
```graphql
type Subscription {
  # Only for critical updates
  accountBalanceChanged(accountId: ID!): Account! @premium
  
  # Batch updates to reduce messages
  transactionProcessed(
    batchInterval: Int = 1000 # milliseconds
  ): [Transaction!]! @rateLimit(max: 10, window: "1m")
  
  # Long-running operations
  receiptProcessingStatus(receiptId: ID!): ReceiptStatus!
}
```

### 5. Error Handling Standards

#### Structured Error Responses
```json
{
  "errors": [
    {
      "message": "Insufficient funds for transaction",
      "extensions": {
        "code": "INSUFFICIENT_FUNDS",
        "severity": "ERROR",
        "retriable": false,
        "userMessage": "Your account doesn't have enough funds for this transaction",
        "details": {
          "available": 10000,
          "required": 15000,
          "currency": "USD"
        }
      }
    }
  ]
}
```

#### Error Categories and Handling
| Category | HTTP Status | Retry | User Action | Example |
|----------|------------|-------|-------------|---------|
| Validation | 400 | No | Fix input | Invalid amount |
| Authentication | 401 | No | Re-login | Token expired |
| Authorization | 403 | No | Upgrade plan | Premium feature |
| Rate Limit | 429 | Yes | Wait | Too many requests |
| Business Logic | 422 | No | Review | Insufficient funds |
| Server Error | 500 | Yes | Retry later | Database error |

### 6. Caching Strategy for Cost Reduction

#### Multi-Layer Caching (60-80% cost reduction)
```go
// Layer 1: Response caching (CDN/CloudFlare)
// @cacheControl directives in schema

// Layer 2: Application caching (Redis)
type CacheConfig struct {
    // Reference data (long TTL)
    Categories:    CacheTTL{Duration: 1 * time.Hour},
    Merchants:     CacheTTL{Duration: 1 * time.Hour},
    
    // User data (medium TTL)
    Accounts:      CacheTTL{Duration: 5 * time.Minute},
    Transactions:  CacheTTL{Duration: 1 * time.Minute},
    
    // Sensitive data (short TTL)
    Balances:      CacheTTL{Duration: 30 * time.Second},
    
    // Analytics (long TTL)
    MonthlyStats:  CacheTTL{Duration: 6 * time.Hour},
}

// Layer 3: Database query caching
// Prepared statements and connection pooling
```

#### Cache Invalidation Patterns
```go
// Smart invalidation on mutations
func (r *mutationResolver) CreateTransaction(ctx context.Context, input model.CreateTransactionInput) (*model.TransactionPayload, error) {
    // Create transaction
    tx, err := r.service.CreateTransaction(ctx, input)
    if err != nil {
        return nil, err
    }
    
    // Invalidate affected caches
    cacheKeys := []string{
        fmt.Sprintf("account:%s:balance", tx.AccountID),
        fmt.Sprintf("user:%s:transactions", tx.UserID),
        fmt.Sprintf("category:%s:total", tx.CategoryID),
    }
    
    r.cache.DeleteMulti(ctx, cacheKeys...)
    
    return &model.TransactionPayload{
        Success: true,
        Transaction: tx,
    }, nil
}
```

### 7. Rate Limiting and Query Complexity

#### Multi-Dimensional Rate Limiting (40% infrastructure cost reduction)
```go
// Configuration for different tiers
type RateLimitConfig struct {
    Free: UserLimits{
        RequestsPerMinute: 60,
        ComplexityPerMinute: 1000,
        MutationsPerHour: 100,
        OCRPerDay: 5,
    },
    Premium: UserLimits{
        RequestsPerMinute: 300,
        ComplexityPerMinute: 5000,
        MutationsPerHour: 1000,
        OCRPerDay: 100,
    },
}

// Query complexity calculation
func calculateComplexity(query string) int {
    complexity := 0
    
    // Field costs
    fieldCosts := map[string]int{
        "transaction": 1,
        "transactions": 10,
        "account": 1,
        "accounts": 5,
        "advancedAnalytics": 100,
        "processReceipt": 50,
    }
    
    // Add pagination multiplier
    // Add depth multiplier
    // Add field count
    
    return complexity
}
```

#### Progressive Rate Limiting
```go
// Gradually restrict abusive clients
func progressiveRateLimit(clientID string, violation int) RateLimit {
    penalties := []RateLimit{
        {Factor: 1.0},    // Normal
        {Factor: 0.5},    // 50% reduction
        {Factor: 0.25},   // 75% reduction
        {Factor: 0.1},    // 90% reduction
        {Factor: 0},      // Blocked
    }
    
    if violation >= len(penalties) {
        return penalties[len(penalties)-1]
    }
    
    return penalties[violation]
}
```

### 8. Pagination Standards

#### Cursor-Based Pagination (Relay Specification)
```graphql
# Consistent pagination across all collections
interface Connection {
  pageInfo: PageInfo!
  totalCount: Int
}

interface Edge {
  cursor: String!
  node: Node!
}

# Implementation example
type Query {
  transactions(
    # Forward pagination
    first: Int
    after: String
    
    # Backward pagination  
    last: Int
    before: String
    
    # Filtering
    filter: TransactionFilter
  ): TransactionConnection!
}
```

#### Cursor Implementation
```go
// Opaque cursor encoding
func encodeCursor(timestamp time.Time, id string) string {
    cursor := fmt.Sprintf("%d:%s", timestamp.UnixNano(), id)
    return base64.StdEncoding.EncodeToString([]byte(cursor))
}

func decodeCursor(encoded string) (time.Time, string, error) {
    decoded, err := base64.StdEncoding.DecodeString(encoded)
    if err != nil {
        return time.Time{}, "", err
    }
    
    parts := strings.Split(string(decoded), ":")
    if len(parts) != 2 {
        return time.Time{}, "", errors.New("invalid cursor format")
    }
    
    nano, _ := strconv.ParseInt(parts[0], 10, 64)
    timestamp := time.Unix(0, nano)
    
    return timestamp, parts[1], nil
}
```

### 9. Mobile-First Considerations

#### Offline Support Patterns
```graphql
# Sync-friendly mutations
type Mutation {
  # Batch sync for offline changes
  syncTransactions(
    transactions: [SyncTransactionInput!]!
  ): SyncTransactionPayload!
  
  # Conflict resolution
  resolveConflict(
    localVersion: ID!
    serverVersion: ID!
    resolution: ConflictResolution!
  ): Transaction!
}

input SyncTransactionInput {
  localId: ID!         # Client-side ID
  serverId: ID         # Server ID if exists
  version: Int!        # Version for conflict detection
  data: CreateTransactionInput!
  deviceId: String!
  syncedAt: DateTime!
}

enum ConflictResolution {
  KEEP_LOCAL
  KEEP_SERVER
  MERGE
}
```

#### Bandwidth Optimization
```graphql
# Field selection for mobile
type Query {
  # Minimal fields for list views
  transactionsList(
    first: Int!
    after: String
  ): TransactionConnection! 
    @fieldSelection(default: ["id", "amount", "description", "date"])
  
  # Full details on demand
  transactionDetails(id: ID!): Transaction!
}

# Progressive data loading
type Account {
  # Essential fields (always loaded)
  id: ID!
  name: String!
  balance: Money!
  
  # Extended fields (loaded on demand)
  transactions: TransactionConnection @defer
  analytics: AccountAnalytics @defer
  history: [BalanceHistory!] @defer
}
```

### 10. API Versioning Strategy

#### Schema Evolution Without Breaking Changes
```graphql
# Deprecation instead of removal
type Transaction {
  id: ID!
  amount: Money!
  
  # Deprecated field
  amountCents: Int @deprecated(reason: "Use amount.amountMinor")
  
  # New field with default
  tags: [String!]! @since(version: "1.1.0")
}

# Feature flags for gradual rollout
directive @feature(flag: String!) on FIELD_DEFINITION

type Query {
  # New feature behind flag
  aiInsights: AIInsights @feature(flag: "ai_insights") @premium
}
```

## Consequences

### Positive
- **Cost Efficiency**: 40-60% reduction in MVP infrastructure costs
- **Type Safety**: Strong typing across entire stack
- **Developer Experience**: Self-documenting API with GraphQL
- **Performance**: DataLoader pattern prevents N+1 queries
- **Flexibility**: GraphQL allows client-specific queries
- **Mobile Optimized**: Efficient bandwidth usage and offline support
- **Scalability**: Architecture supports 100K+ users
- **Cache Hit Rate**: 60-80% reduction in database load

### Negative
- **Complexity**: GraphQL adds learning curve for team
- **Query Complexity**: Need to manage and limit query complexity
- **Caching Challenges**: GraphQL caching more complex than REST
- **File Uploads**: GraphQL multipart uploads add complexity
- **Monitoring**: Requires specialized GraphQL monitoring tools

### Risks
- **Query Explosion**: Clients could create expensive queries
- **Over-fetching**: Poor query design could negate benefits
- **Security**: GraphQL introspection needs careful management
- **Rate Limiting**: Complex to implement per-query limits
- **Breaking Changes**: Schema evolution needs careful planning

## Alternatives Considered

### Option A: REST API
- **Description**: Traditional RESTful API design
- **Pros**: Simple, well-understood, better caching
- **Cons**: Over/under-fetching, multiple round trips, no type safety
- **Reason for not choosing**: GraphQL better for complex financial queries

### Option B: gRPC
- **Description**: Binary protocol with protobuf
- **Pros**: Very efficient, strong typing, streaming support
- **Cons**: Poor browser support, complex for mobile, no introspection
- **Reason for not choosing**: Not ideal for mobile-first architecture

### Option C: JSON-RPC
- **Description**: Simple RPC over HTTP
- **Pros**: Simple, lightweight, easy to implement
- **Cons**: No type safety, no standard tooling, limited ecosystem
- **Reason for not choosing**: Lacks features needed for complex app

## Implementation Notes

### Migration Strategy
1. **Week 1-2**: Define core GraphQL schema
2. **Week 3-4**: Implement authentication and basic queries
3. **Week 5-6**: Add mutations and data loaders
4. **Week 7-8**: Implement caching layers
5. **Week 9-10**: Add rate limiting and monitoring
6. **Week 11-12**: Performance optimization and testing

### Testing Approach
- Schema validation tests
- Resolver unit tests with mocks
- Integration tests for complete flows
- Load testing for query complexity
- Mobile client integration tests

### Monitoring and Success Metrics
- P95 response time < 100ms
- Cache hit rate > 60%
- Query complexity budget adherence > 95%
- API error rate < 0.1%
- Monthly API costs < $500 for MVP
- Client satisfaction score > 4.5/5

### Cost Optimization Roadmap
| Phase | Implementation | Cost Savings | Timeline |
|-------|---------------|--------------|----------|
| MVP | Basic caching, pagination | 40% | Launch |
| Growth | Advanced caching, CDN | 60% | Month 3 |
| Scale | Edge computing, optimization | 80% | Month 6 |

## References
- [GraphQL Best Practices](https://graphql.org/learn/best-practices/)
- [Relay Cursor Specification](https://relay.dev/graphql/connections.htm)
- [DataLoader Pattern](https://github.com/graphql/dataloader)
- [GraphQL Cost Analysis](https://github.com/pa-bru/graphql-cost-analysis)
- [gqlgen Documentation](https://gqlgen.com/)
- ADR-20250812: Personal Finance Tech Stack Selection
- ADR-20250819: Data Architecture and Schema Design
- ADR-20250819: Security Architecture

## Decision Makers
- Author: Archie (System Architect)
- Reviewers: [Pending]
- Approver: [Pending]
- Date: 2025-08-19

## Revision History
- 2025-08-19: Initial comprehensive draft with cost optimization focus