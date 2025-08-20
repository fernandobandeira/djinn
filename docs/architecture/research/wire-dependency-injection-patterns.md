# Architecture Research: Wire Dependency Injection Patterns for Go

## Research Type
Pattern Comparison

## Executive Summary
Google Wire provides compile-time dependency injection for Go applications with zero runtime overhead, excellent clean architecture support, and robust testing patterns. Recommended as the primary DI solution for the Djinn application due to its type safety, performance characteristics, and alignment with hexagonal architecture principles.

## Evaluation Context
- **Date**: 2025-08-20
- **Author**: Archie (System Architect)
- **Related ADRs**: [ADR-20250820-code-organization-module-structure](../adrs/ADR-20250820-code-organization-module-structure.md)
- **Business Context**: Need for maintainable, testable dependency management in a growing financial application
- **Technical Context**: Go backend with clean architecture, GraphQL API, multiple infrastructure dependencies

## Options Evaluated

### Option 1: Google Wire (Compile-Time DI)
- **Description**: Code generation-based dependency injection that resolves dependencies at compile time
- **Strengths**:
  - Zero runtime overhead - generates plain Go code
  - Compile-time type safety and error detection
  - Excellent IDE support and debugging experience
  - Clean separation of injection logic from business code
  - Strong interface binding support for testing
- **Weaknesses**:
  - Requires code generation step in build process
  - Learning curve for provider patterns
  - Build tag management complexity
- **Technical Fit**: High - Perfect alignment with clean architecture and Go idioms
- **Maturity**: Production-ready (Google-maintained, widely adopted)
- **Community/Support**: Strong (Google backing, active maintenance)

### Option 2: Uber Dig/Fx (Runtime DI)
- **Description**: Reflection-based runtime dependency injection framework
- **Strengths**:
  - No code generation required
  - Dynamic dependency resolution
  - Lifecycle management hooks
  - Module system for organizing dependencies
- **Weaknesses**:
  - Runtime reflection overhead
  - Runtime panics for misconfiguration
  - Less transparent dependency graph
  - Harder to debug injection issues
- **Technical Fit**: Medium - Works but adds runtime complexity
- **Maturity**: Production-ready (Uber-maintained)
- **Community/Support**: Strong (Uber backing)

### Option 3: Manual Dependency Injection
- **Description**: Hand-written dependency wiring without frameworks
- **Strengths**:
  - Complete control and transparency
  - No external dependencies
  - No learning curve for Go developers
  - Easiest to debug
- **Weaknesses**:
  - Significant boilerplate code
  - Error-prone manual wiring
  - Difficult to maintain as application grows
  - Testing setup complexity
- **Technical Fit**: Low - Doesn't scale well for complex applications
- **Maturity**: N/A (Pattern, not framework)
- **Community/Support**: N/A

## Evaluation Criteria

| Criterion | Weight | Wire | Dig/Fx | Manual |
|-----------|--------|------|--------|--------|
| Performance | High | 5 | 3 | 5 |
| Scalability | High | 5 | 4 | 2 |
| Maintainability | High | 5 | 4 | 2 |
| Learning Curve | Medium | 3 | 3 | 5 |
| Testing Support | High | 5 | 4 | 3 |
| Integration Complexity | Medium | 4 | 5 | 3 |
| Type Safety | High | 5 | 3 | 5 |
| **Total Weighted Score** | | **32** | **26** | **25** |

## Deep Dive Analysis

### Performance Characteristics
- **Wire**: Zero runtime overhead, all resolution at compile-time
- **Dig/Fx**: ~5-10% overhead for reflection-based injection
- **Manual**: Same as Wire but with more code to maintain

### Architecture Implications

**Wire with Clean Architecture:**
```go
// Domain layer - pure business logic
package domain
var DomainSet = wire.NewSet(
    NewUserService,
    NewTransactionProcessor,
)

// Application layer - use cases
package application
var ApplicationSet = wire.NewSet(
    NewCreateUserHandler,
    NewProcessReceiptHandler,
    wire.Bind(new(UserRepository), new(*PostgresUserRepo)),
)

// Infrastructure layer - external services
package infrastructure
var InfrastructureSet = wire.NewSet(
    NewPostgresDB,
    NewRedisCache,
    NewS3Storage,
)
```

### Best Practices Discovered

1. **Provider Organization**:
   - Group providers by architectural layer
   - Export provider sets for reusability
   - Use interface binding for dependency inversion

2. **Testing Patterns**:
   ```go
   //+build wireinject
   
   func InitializeTestApp(db *sql.DB) (*App, error) {
       wire.Build(
           TestProviderSet,
           wire.Value(db), // Inject test database
           NewApp,
       )
       return nil, nil
   }
   ```

3. **Error Handling**:
   - Wire automatically propagates errors from providers
   - Use cleanup functions for resource management
   - Provider functions should return errors for initialization failures

4. **Advanced Patterns**:
   - Conditional providers using build tags
   - Provider overrides for different environments
   - Multi-injector patterns for microservices

### Risk Assessment
| Risk | Probability | Impact | Mitigation |
|------|------------|--------|------------|
| Circular dependencies | Medium | High | Wire detects at compile-time, refactor architecture |
| Build process complexity | Low | Medium | Automate with Makefile, document process |
| Learning curve for team | Medium | Low | Provide examples, pair programming |
| Provider maintenance overhead | Low | Low | Good organization patterns, regular reviews |

## Recommendation

### Recommended Option: Google Wire
Wire is the recommended dependency injection solution for the Djinn application based on comprehensive evaluation.

### Rationale
1. **Performance Critical**: Zero runtime overhead essential for financial calculations
2. **Type Safety**: Compile-time validation prevents runtime DI failures
3. **Clean Architecture Alignment**: Perfect support for hexagonal architecture patterns
4. **Testing Excellence**: Superior mocking and test injection capabilities
5. **Debugging Simplicity**: Generated code is plain Go, easy to understand and debug

### Alternative Scenarios
- **If team prefers runtime DI**: Consider Uber Fx for its lifecycle management
- **If project remains small (<10 dependencies)**: Manual injection might suffice
- **If CI/CD cannot support code generation**: Reconsider Dig/Fx

## Implementation Roadmap

### Prerequisites
- [x] Install Wire: `go install github.com/google/wire/cmd/wire@latest`
- [ ] Add Wire to project tools: `tools.go`
- [ ] Configure Makefile targets for wire generation
- [ ] Set up build tags properly

### Implementation Steps
1. Create provider functions for existing types
2. Organize providers into layer-specific sets
3. Create main injector in `cmd/api/wire.go`
4. Generate wire code: `wire ./...`
5. Create test injectors with mock providers
6. Document provider patterns for team

### Success Metrics
- **Build Time**: <5 seconds for wire generation
- **Test Coverage**: 100% of providers tested
- **Dependency Clarity**: All dependencies explicit in provider sets
- **Developer Velocity**: Reduced boilerplate by 70%

## Code Examples

### Layer-Based Provider Sets
```go
// internal/domain/wire.go
package domain

import "github.com/google/wire"

var DomainSet = wire.NewSet(
    NewUserService,
    NewBudgetCalculator,
    NewReceiptProcessor,
)

// internal/infrastructure/wire.go
package infrastructure

import "github.com/google/wire"

var InfrastructureSet = wire.NewSet(
    NewPostgresConnection,
    NewRedisClient,
    repositories.RepositorySet,
    wire.Bind(new(domain.UserRepository), new(*PostgresUserRepository)),
)
```

### Test Injection Pattern
```go
//+build wireinject

package main

import "github.com/google/wire"

func InitializeTestServer(cfg *TestConfig) (*Server, error) {
    wire.Build(
        TestProviderSet,
        MockRepositorySet,
        NewServer,
    )
    return nil, nil
}
```

### Main Application Injector
```go
//+build wireinject

package main

import (
    "github.com/google/wire"
    "djinn/internal/domain"
    "djinn/internal/infrastructure"
    "djinn/internal/application"
)

func InitializeApp(cfg *Config) (*App, error) {
    wire.Build(
        domain.DomainSet,
        infrastructure.InfrastructureSet,
        application.ApplicationSet,
        NewGraphQLServer,
        NewApp,
    )
    return nil, nil
}
```

## References
- [Google Wire Official Documentation](https://github.com/google/wire)
- [Wire User Guide](https://github.com/google/wire/blob/main/docs/guide.md)
- [Wire Best Practices](https://github.com/google/wire/blob/main/docs/best-practices.md)
- [Compile-Time DI in Go](https://blog.golang.org/compile-time-dependency-injection-with-go-cloud-wire)
- [Wire vs Other DI Frameworks](https://medium.com/@dche423/golang-dependency-injection-with-wire-74f81cd222f6)
- [Clean Architecture with Wire](https://github.com/bxcodec/go-clean-arch)

## Appendix

### Wire Command Reference
```bash
# Generate wire code
wire ./...

# Check for errors without generating
wire check ./...

# Show dependency graph
wire show ./...

# Generate with verbose output
wire -v ./...
```

### Common Wire Patterns
1. **Value Providers**: `wire.Value(config)`
2. **Interface Binding**: `wire.Bind(new(Interface), new(*Implementation))`
3. **Provider Sets**: `wire.NewSet(providers...)`
4. **Struct Providers**: `wire.Struct(new(AppConfig), "*")`
5. **Field Providers**: `wire.FieldsOf(new(Config), "Database")`

---
*This research informs architecture decisions and may lead to formal ADRs*