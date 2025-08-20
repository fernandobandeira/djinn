# ADR-20250812: Personal Finance App Tech Stack Selection

## Status
Accepted

## Context
Building a personal finance application requiring:
- Strong contracts and correctness for financial calculations
- Offline-first mobile experience with reliable sync
- Complex financial workflows (imports, categorization, budgeting)
- Multi-user support with proper isolation
- High data integrity for monetary values
- Scalable architecture that can evolve with features

### Constraints
- Small team requiring high developer productivity
- Need for type safety across the entire stack
- Must handle financial regulations and data precision
- Mobile-first but may expand to web
- Limited initial budget (optimize for proven, cost-effective solutions)

## Decision

### Core Technology Stack

#### Backend
- **Language**: Go (type safety, performance, simple deployment)
- **HTTP Framework**: chi (lightweight, stdlib-compatible)
- **GraphQL**: gqlgen (schema-first, type-safe codegen)
- **Database**: PostgreSQL 16+ (ACID, native UUIDv7 support)
- **Schema Management**: Atlas HCL-first (declarative, safe migrations)
- **Query Builder**: sqlc with pgx/v5 driver (compile-time SQL verification)
- **Workflow Engine**: Temporal (durable execution, perfect for financial workflows)
- **Authentication**: Firebase Auth with OAuth only (Google/Apple Sign-In, no email/password)
- **Logging**: log/slog (stdlib, Temporal native integration)
- **UUID**: gofrs/uuid/v5 (UUIDv7 support)
- **Config**: envconfig (12-factor app compliant)
- **Validation**: gqlgen directives + domain layer
- **DataLoader**: dataloaden (type-safe code generation)
- **Testing**: standard library + go-cmp + rapid (property testing)

#### Mobile (Flutter)
- **Framework**: Flutter (single codebase, strong typing)
- **GraphQL Client**: Ferry (code generation, normalized cache)
- **Local Database**: Drift (SQLite wrapper with type safety and migrations)
- **Cache Persistence**: Ferry HiveStore (GraphQL normalized cache) + Drift (structured data)
- **State Management**: Riverpod 2.0 (reactive, less boilerplate)
- **Navigation**: go_router (official, declarative)
- **HTTP Client**: built-in http package (Ferry handles GraphQL)
- **UUID**: uuid package with v7 support

#### Infrastructure & Operations
- **Observability**: OpenTelemetry + Prometheus + slog
- **Caching**: Redis with go-redis/v9 (optional, for rate limiting)
- **CSV Storage**: Local filesystem or S3/GCS (for CSV imports/exports only)
- **CSV Parsing**: encoding/csv (stdlib sufficient)

### Key Architectural Decisions

#### 1. Data Integrity
- **UUIDv7 for all IDs**: Time-ordered, better index performance, natural audit trail
- **Money as BIGINT minor units**: Never use floating point for money
- **Import deduplication**: External IDs prevent duplicate transaction imports
- **Idempotency keys**: Client-generated UUIDv7 for all mutations

#### 2. API Design
- **Polling over subscriptions**: Start simple, add WebSockets later if needed
- **GraphQL with persisted queries**: Type safety, reduced bandwidth
- **Row-level multi-tenancy**: user_id on every table

#### 3. Integrations
- **Manual CSV import for v1**: Defer bank API connectors to v2
- **Store native currency**: Display-only conversion using ECB rates
- **Temporal for all async operations**: Guaranteed completion with observability

#### 4. Database Strategy
```sql
-- All tables use UUIDv7 primary keys
CREATE TABLE accounts (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    user_id UUID NOT NULL,
    name TEXT NOT NULL,
    currency CHAR(3) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Transactions stored as minor units
CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    user_id UUID NOT NULL,
    account_id UUID NOT NULL REFERENCES accounts(id),
    external_id TEXT UNIQUE,  -- Plaid/bank transaction ID
    amount_minor BIGINT NOT NULL,  -- Cents, pence, etc
    currency CHAR(3) NOT NULL,
    occurred_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Idempotency with UUIDv7
CREATE TABLE idempotency_keys (
    key UUID PRIMARY KEY,  -- UUIDv7 from client
    user_id UUID NOT NULL,
    operation TEXT NOT NULL,
    result JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

## Consequences

### Positive
- **Type Safety End-to-End**: sqlc → gqlgen → Ferry → Drift ensures compile-time correctness
- **Schema Evolution**: Atlas HCL provides safe, reviewable migrations
- **Workflow Reliability**: Temporal guarantees completion of financial operations
- **Developer Velocity**: Code generation reduces boilerplate and bugs
- **Operational Simplicity**: Go single binary, minimal runtime dependencies
- **Cost Effective**: Mostly open source, Firebase Auth has generous free tier
- **Index Performance**: UUIDv7 provides 30% better insert performance than UUIDv4
- **Debugging**: UUIDv7 timestamps embedded in IDs aid troubleshooting
- **Standard Library**: Using slog, encoding/csv reduces dependencies

### Negative
- **Learning Curve**: Team needs to learn Atlas HCL, Temporal concepts
- **Initial Setup**: Multiple codegen tools require build pipeline setup
- **Flutter Web**: Would require additional work if web becomes priority
- **GraphQL Complexity**: Adds layer vs REST, but Ferry makes it worthwhile
- **PostgreSQL 16+**: Requires newer version for native UUIDv7 support

### Risks
- **Temporal Operational Overhead**: Requires separate cluster
  - *Mitigation*: Use Temporal Cloud initially ($25/mo)
- **Atlas Lock-in**: Less common than Flyway/Liquibase
  - *Mitigation*: Can export to SQL, tool is open source
- **Firebase Auth Vendor Lock**: Tied to Google
  - *Mitigation*: Standard JWT tokens, can migrate to Auth0/Supabase

## Alternatives Considered

### Backend Platform
- **Node.js + TypeScript + Prisma**: Single language, but weaker type safety for finance
- **Hasura + PostgreSQL**: Instant GraphQL, but less control over business logic
- **Rejected**: Go's compile-time safety and performance better for financial domain

### Mobile Framework
- **React Native + Apollo**: Larger ecosystem, but two codebases and weaker types
- **Native iOS/Android**: Best performance, but 2x development effort
- **Rejected**: Flutter's single codebase and Dart's type system preferred

### UUID Strategy
- **UUIDv4**: More common, but random = index fragmentation
- **BIGSERIAL**: Best performance, but no distributed generation
- **Rejected**: UUIDv7 balances global uniqueness with index performance

### Logging
- **zerolog**: Slightly faster, but non-standard API
- **Rejected**: slog is stdlib, integrates better with Temporal

## Implementation Notes

### Critical Implementation Details
```go
// Always use UUIDv7
import "github.com/gofrs/uuid/v5"

func NewID() string {
    id, _ := uuid.NewV7()
    return id.String()
}

// Money handling
type Money struct {
    MinorUnits int64  // NEVER float
    Currency   string // ISO 4217
}

// Idempotency for all mutations
type Mutation {
    CreateTransaction(input: Input!, idempotencyKey: ID!): Transaction!
}
```

### Migration Strategy
1. **Week 1-2**: Core schema + basic CRUD
2. **Week 3-4**: Import workflows via Temporal
3. **Week 5-6**: Categorization + rules engine
4. **Week 7-8**: Budgeting + reporting
5. **Week 9**: Multi-currency display

### Rollback Plan
- Database: Atlas down migrations tested in staging
- API: GraphQL schema versioning, deprecation before removal
- Mobile: Feature flags for gradual rollout

## References
- [UUIDv7 RFC 9562](https://www.rfc-editor.org/rfc/rfc9562.html#name-uuid-version-7)
- [Atlas HCL Documentation](https://atlasgo.io/atlas-schema/hcl)
- [sqlc + pgx Documentation](https://docs.sqlc.dev/en/latest/howto/ddl.html)
- [Temporal Best Practices](https://docs.temporal.io/best-practices)
- [Ferry Architecture](https://ferrygraphql.com/architecture)
- [Monzo's Architecture](https://monzo.com/blog/2016/09/19/building-a-modern-bank-backend)