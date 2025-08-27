# User Story: DJINN-1002 - Golang GraphQL API Foundation with Database Setup

## Metadata
- **Story ID**: DJINN-1002
- **Title**: Golang GraphQL API Foundation with Database Setup
- **Epic**: Epic 1: Base Architecture & Authentication Foundation
- **Author**: Story Creator (SM Agent)
- **Created**: 2025-01-25
- **Last Modified**: 2025-01-25
- **Status**: Draft
- **Priority**: Must Have
- **Effort Estimate**: 10 hours
- **Sprint**: TBD
- **Assignee**: TBD

## User Story
**As a** mobile app  
**I want** a secure GraphQL API endpoint with proper database foundation  
**So that** I can communicate with the backend and persist user data securely  

## Business Value
This story establishes the foundational backend infrastructure for the Djinn personal finance application, enabling secure data persistence and API communication between the Flutter mobile app and the backend services.

## Acceptance Criteria

### Measurable Criteria
1. **Go Project Structure**
   - [ ] Project follows ADR-20250820-code-organization standards
   - [ ] Directory structure matches unified-project-structure specifications
   - **Validation**: Code review confirms adherence to Go project layout standards

2. **Database Foundation**
   - [ ] PostgreSQL 16 database configured with UUIDv7 support
   - [ ] Atlas migrations setup with reversible migration capability
   - [ ] User table created with specified schema (id, firebase_uid, email, name, profile_image_url, timestamps)
   - **Validation**: Database schema matches requirements, migrations run successfully

3. **Data Access Layer**
   - [ ] sqlc configured with pgx/v5 driver for type-safe queries
   - [ ] User CRUD operations implemented and tested
   - **Validation**: Generated queries compile without errors, CRUD operations tested

4. **GraphQL API Layer**
   - [ ] Chi HTTP framework configured with middleware stack
   - [ ] gqlgen GraphQL server setup with schema-first approach
   - [ ] User GraphQL schema and resolvers (getUser, createUser mutations)
   - [ ] DataLoader pattern implemented for User queries
   - **Validation**: GraphQL playground accessible, queries execute successfully

5. **Performance & Infrastructure**
   - [ ] PgBouncer connection pooling configured
   - [ ] Health check endpoint responds at /health
   - [ ] OpenTelemetry tracing configured per ADR-20250120-monitoring
   - **Validation**: Connection pooling metrics available, health checks pass, traces captured

6. **Configuration & Environment**
   - [ ] Environment configuration with envconfig
   - [ ] Docker containerization for local development
   - **Validation**: Environment variables loaded correctly, Docker containers start successfully

### Validation Criteria
- Database migrations execute without errors
- GraphQL playground serves at /graphql endpoint
- User queries/mutations work end-to-end
- DataLoader prevents N+1 query problems
- Health check endpoint returns 200 OK
- Unit tests achieve 90%+ coverage
- Integration tests verify database connectivity
- Code passes linting and formatting checks

## Tasks and Subtasks

### 1. Project Structure Setup (2 hours) ✅
- [x] **1.1** Initialize Go project with proper module structure
  - Dependencies: None
  - Effort: 30 minutes
- [x] **1.2** Create directory structure per ADR-20250820-code-organization
  - Dependencies: 1.1
  - Effort: 30 minutes
- [x] **1.3** Setup development Docker containers
  - Dependencies: 1.2
  - Effort: 1 hour

### 2. Database Foundation (3 hours) ✅
- [x] **2.1** Configure PostgreSQL 16 with UUIDv7 extension
  - Dependencies: 1.3
  - Effort: 1 hour
- [x] **2.2** Setup Atlas migration framework
  - Dependencies: 2.1
  - Effort: 45 minutes
- [x] **2.3** Create initial user table migration
  - Dependencies: 2.2
  - Effort: 45 minutes
- [x] **2.4** Configure PgBouncer connection pooling
  - Dependencies: 2.1
  - Effort: 30 minutes

### 3. Data Access Layer (2 hours) ✅
- [x] **3.1** Configure sqlc with pgx/v5 driver
  - Dependencies: 2.3
  - Effort: 45 minutes
- [x] **3.2** Implement User CRUD queries
  - Dependencies: 3.1
  - Effort: 1 hour
- [x] **3.3** Generate and validate type-safe query functions
  - Dependencies: 3.2
  - Effort: 15 minutes

### 4. GraphQL API Layer (2.5 hours) ✅
- [x] **4.1** Setup Chi HTTP framework with middleware
  - Dependencies: 1.2
  - Effort: 45 minutes
- [x] **4.2** Configure gqlgen with schema-first approach
  - Dependencies: 4.1
  - Effort: 1 hour
- [x] **4.3** Implement User GraphQL schema and resolvers
  - Dependencies: 3.3, 4.2
  - Effort: 45 minutes

### 5. Performance Optimization (1 hour) ✅
- [x] **5.1** Implement DataLoader pattern for User queries
  - Dependencies: 4.3
  - Effort: 1 hour

### 6. Infrastructure & Monitoring (1.5 hours) ✅
- [x] **6.1** Configure OpenTelemetry tracing
  - Dependencies: 4.1
  - Effort: 45 minutes
- [x] **6.2** Implement health check endpoint
  - Dependencies: 4.1
  - Effort: 30 minutes
- [x] **6.3** Setup environment configuration with envconfig
  - Dependencies: 1.2
  - Effort: 15 minutes

## Dev Notes

### Architecture Context
This story implements the backend foundation described in:
- [Source: architecture/backend-architecture.md#api-layer] - GraphQL API design patterns
- [Source: architecture/database-schema.md#user-table] - User table schema requirements
- [Source: architecture/tech-stack.md#backend-stack] - Go, PostgreSQL, GraphQL technology choices
- [Source: architecture/adrs/ADR-20250820-code-organization.md] - Go project structure standards

### Database Schema
```sql
-- User table following database-schema.md specifications
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v7(),
    firebase_uid VARCHAR(128) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    profile_image_url TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_users_firebase_uid ON users(firebase_uid);
CREATE INDEX idx_users_email ON users(email);
```
[Source: architecture/database-schema.md#user-table]

### GraphQL Schema
```graphql
type User {
  id: ID!
  firebaseUid: String!
  email: String!
  name: String!
  profileImageUrl: String
  createdAt: Time!
  updatedAt: Time!
}

type Query {
  user(id: ID!): User
  userByFirebaseUid(firebaseUid: String!): User
}

type Mutation {
  createUser(input: CreateUserInput!): User!
  updateUser(id: ID!, input: UpdateUserInput!): User!
}
```

### Implementation Details

#### Project Structure
Following [Source: architecture/unified-project-structure.md#backend-structure]:
```
backend/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── config/
│   ├── database/
│   ├── graph/
│   │   ├── generated/
│   │   ├── model/
│   │   ├── resolver/
│   │   └── schema/
│   ├── middleware/
│   └── service/
├── migrations/
├── pkg/
└── docker-compose.yml
```

#### Technology Integration
- **Chi Router**: HTTP middleware and routing [Source: architecture/tech-stack.md#http-framework]
- **gqlgen**: Schema-first GraphQL server generation [Source: architecture/backend-architecture.md#graphql-layer]
- **sqlc**: Type-safe SQL query generation [Source: architecture/tech-stack.md#database-orm]
- **Atlas**: Database migration management [Source: architecture/database-schema.md#migrations]
- **pgx/v5**: High-performance PostgreSQL driver [Source: architecture/tech-stack.md#database-driver]

#### Error Handling
Implement error handling patterns from [Source: architecture/adrs/ADR-20250120-error-handling.md]:
- Structured error responses in GraphQL
- Error logging with context
- Proper HTTP status codes for health checks

#### Monitoring Integration
Configure OpenTelemetry following [Source: architecture/adrs/ADR-20250120-monitoring.md]:
- HTTP request tracing
- Database query tracing
- GraphQL operation tracing
- Custom metrics for API performance

### Configuration Requirements
Environment variables required:
```
DATABASE_URL=postgres://user:pass@localhost:5432/djinn
PGBOUNCER_URL=postgres://user:pass@localhost:6543/djinn
PORT=8080
LOG_LEVEL=info
OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4317
```

### File Locations
Key files to be created/modified:
- `/backend/cmd/api/main.go` - Application entry point
- `/backend/internal/config/config.go` - Configuration management
- `/backend/internal/infrastructure/persistence/postgres/connection.go` - Database connection setup
- `/backend/internal/graph/schema/user.graphql` - User GraphQL schema
- `/backend/internal/graph/resolver/user.go` - User resolvers
- `/backend/migrations/001_create_users_table.sql` - Initial migration
- `/backend/sqlc/queries/users.sql` - User CRUD queries
- `/backend/docker-compose.yml` - Local development environment

## Testing Requirements

### Unit Tests (90%+ coverage target)
- Database connection and migration tests
- sqlc generated query tests
- GraphQL resolver unit tests
- Middleware unit tests
- Configuration loading tests

### Integration Tests
- End-to-end GraphQL query/mutation tests
- Database connectivity tests
- Health check endpoint tests
- DataLoader performance tests (N+1 prevention)

### Performance Tests
- GraphQL query performance benchmarks
- Database connection pooling tests
- Memory usage tests for DataLoader

### Security Tests
- Input validation tests
- SQL injection prevention tests (via sqlc)
- GraphQL query depth limiting tests

## Definition of Done

### Code Quality
- [ ] Code follows Go formatting standards (gofmt, goimports)
- [ ] All linting checks pass (golangci-lint)
- [ ] Unit tests achieve 90%+ coverage
- [ ] Integration tests pass
- [ ] Code review completed and approved

### Functionality
- [ ] Database migrations execute successfully
- [ ] GraphQL playground accessible at /graphql
- [ ] User CRUD operations work end-to-end
- [ ] DataLoader prevents N+1 queries (verified via logging)
- [ ] Health check endpoint returns proper status

### Performance
- [ ] API response times under 100ms for simple queries
- [ ] Database connection pooling working (max connections respected)
- [ ] Memory usage stable under load testing

### Documentation
- [ ] API schema documented in GraphQL playground
- [ ] Database schema documented in migration files
- [ ] Environment variables documented in README
- [ ] Architecture decisions captured in code comments

### Deployment
- [ ] Docker containers build successfully
- [ ] Local development environment starts with docker-compose
- [ ] Environment configuration works across environments
- [ ] OpenTelemetry tracing data captured

### Security
- [ ] No hardcoded credentials in code
- [ ] Environment-based configuration
- [ ] Input validation implemented
- [ ] Error messages don't leak sensitive information

## Dependencies
- **Blocking**: None (foundational story)
- **Related**: Future authentication stories will build on this foundation
- **External**: PostgreSQL 16, Docker environment

## Risks and Mitigation

### Technical Risks
1. **PostgreSQL UUIDv7 Compatibility**
   - Risk: UUIDv7 not available in older PostgreSQL versions
   - Mitigation: Verify PostgreSQL 16+ in deployment, fallback to UUIDv4 if needed

2. **GraphQL N+1 Query Problem**
   - Risk: DataLoader implementation complexity
   - Mitigation: Use proven dataloaden library, implement comprehensive testing

3. **Database Migration Complexity**
   - Risk: Atlas learning curve and migration failures
   - Mitigation: Start with simple migrations, comprehensive rollback testing

### Performance Risks
1. **Connection Pool Exhaustion**
   - Risk: Too few database connections under load
   - Mitigation: PgBouncer configuration with proper pool sizing

2. **GraphQL Query Complexity**
   - Risk: Expensive queries impacting performance
   - Mitigation: Implement query complexity limits, monitoring

### Security Risks
1. **SQL Injection via GraphQL**
   - Risk: Dynamic query construction vulnerabilities
   - Mitigation: Use sqlc for type-safe queries, input validation

## Rollback Strategy

### Database Rollback
1. **Atlas Migration Rollback**
   ```bash
   atlas migrate down --env local
   ```
2. **Manual Database Restoration**
   - Restore from backup if migration rollback fails
   - Verify data integrity after rollback

### Application Rollback
1. **Container Rollback**
   - Revert to previous Docker image
   - Update docker-compose.yml to previous version
2. **Configuration Rollback**
   - Restore previous environment configuration
   - Verify health checks pass after rollback

### Validation After Rollback
- [ ] Database schema matches expected state
- [ ] Application starts successfully
- [ ] Health checks pass
- [ ] No data corruption detected

## Change Log
| Date | Author | Change | Reason |
|------|--------|---------|--------|
| 2025-01-25 | Story Creator | Initial story creation | Epic 1 requirement |

## Dev Agent Record
- **Generated by**: SM Agent -> Story Creator Sub-Agent
- **Template Used**: `.claude/resources/sm/templates/story-template.md`
- **Architecture Sources**: 
  - backend-architecture.md
  - database-schema.md  
  - tech-stack.md
  - unified-project-structure.md
  - ADR-20250820-code-organization.md
  - ADR-20250120-monitoring.md
  - ADR-20250120-error-handling.md
- **Quality Patterns Applied**: acceptance-criteria-patterns.yaml, task-structure-guide.yaml
- **Context**: Epic 1 foundational backend implementation

## QA Results
*[To be populated during QA review phase]*
- [ ] Story validation checklist completed
- [ ] Acceptance criteria reviewed and validated
- [ ] Technical approach approved by architect
- [ ] Effort estimates reviewed by team

## File List
*[To be populated during implementation]*
### Created Files
- [ ] `/backend/cmd/api/main.go`
- [ ] `/backend/internal/config/config.go`
- [ ] `/backend/internal/infrastructure/persistence/postgres/connection.go`
- [ ] `/backend/internal/graph/schema/user.graphql`
- [ ] `/backend/internal/graph/resolver/user.go`
- [ ] `/backend/migrations/001_create_users_table.sql`
- [ ] `/backend/sqlc/queries/users.sql`
- [ ] `/backend/docker-compose.yml`

### Modified Files
- [ ] `/backend/go.mod` (dependencies)
- [ ] `/backend/go.sum` (dependency checksums)