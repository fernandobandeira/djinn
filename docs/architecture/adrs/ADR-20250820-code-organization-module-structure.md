# ADR-20250820: Code Organization & Module Structure

## Status
Accepted

## Context

The Djinn personal finance application requires a well-structured codebase that supports:
- **Clean/Hexagonal Architecture**: Domain-driven design with clear layer separation
- **Multi-Technology Stack**: Go backend (GraphQL API with gqlgen) + Flutter frontend
- **Team Scalability**: Small team initially, but architecture must support growth
- **Type Safety**: End-to-end type safety from database to UI
- **Testability**: Unit, integration, and end-to-end testing strategies
- **Feature Velocity**: Rapid development while maintaining code quality
- **Maintainability**: Clear module boundaries and dependency management

### Current Architecture Context
From existing ADRs, the system uses:
- **Backend**: Go with gqlgen (GraphQL), sqlc (type-safe SQL), PostgreSQL, Temporal workflows
- **Frontend**: Flutter with Ferry (GraphQL client), Drift (local database), Riverpod (state management)
- **Architecture Style**: Clean/Hexagonal with domain-driven design
- **Key Features**: Client-side receipt scanning (Google ML Kit), offline-first mobile, AI assistance, financial workflows

### Constraints
- **Clean Architecture**: Domain logic must be independent of frameworks
- **Framework Isolation**: Infrastructure concerns separated from business logic
- **Circular Dependencies**: Must prevent circular imports between modules
- **Code Generation**: Support for sqlc, gqlgen, dataloaden, Ferry code generation
- **Testing**: Architecture must facilitate comprehensive testing at all levels
- **Deployment**: Single Go binary for backend, standard Flutter build for mobile
- **Internal Package Protection**: Generated code and internal logic must be protected from external imports

## Decision

### 1. Backend Go Module Structure (Clean Architecture)

#### Generated Code Organization

All generated code follows a consistent naming pattern:
- `internal/infrastructure/persistence/postgres/generated/` - sqlc generated database code
- `internal/graph/generated/` - gqlgen generated GraphQL code
- `*/wire_gen.go` - Wire generated dependency injection code

This consistent `generated/` subdirectory pattern makes it immediately clear which code is tool-generated versus hand-written.

#### Dependency Injection with Google Wire

Google Wire will serve as the compile-time dependency injection framework for the Go backend, providing:
- **Compile-time DI**: Code generation eliminates runtime reflection overhead
- **Type Safety**: Wire validates dependency graphs at compile time
- **Clear Dependencies**: Explicit provider functions make dependencies visible
- **Testing Support**: Easy mock injection through interface-based providers
- **Build Integration**: Seamless integration with existing code generation workflow

**Wire Best Practices (from research):**
- **Provider Organization**: Group providers by architectural layer, export as sets
- **Interface Binding**: Use `wire.Bind` for dependency inversion and testability
- **Error Propagation**: Leverage Wire's automatic error handling in providers
- **Build Constraints**: Separate wire.go files with `//+build wireinject` tag
- **Provider Naming**: Use `New` prefix consistently for constructor functions

#### Root Directory Structure
```
/djinn-backend/
├── cmd/                          # Application entry points
│   ├── server/                   # HTTP/GraphQL server
│   │   ├── main.go               # Server entry point
│   │   ├── wire.go               # Wire injector definition
│   │   └── wire_gen.go           # Wire generated code
│   ├── worker/                   # Temporal worker
│   │   ├── main.go               # Worker entry point
│   │   ├── wire.go               # Wire injector definition
│   │   └── wire_gen.go           # Wire generated code
│   ├── migrate/                  # Database migration tool
│   └── seed/                     # Data seeding tool
├── internal/                     # Private application code
│   ├── adapter/                  # Infrastructure layer (adapters)
│   ├── application/              # Application layer (use cases)
│   ├── domain/                   # Domain layer (business logic)
│   ├── infrastructure/           # Infrastructure layer (frameworks)
│   ├── database/                 # Database infrastructure and generated code
│   │   ├── generated/            # sqlc generated code (internal access only)
│   │   ├── mocks/                # Database mocks for testing
│   │   └── connection.go         # Database connection management
│   ├── graph/                    # GraphQL layer (presentation)
│   │   ├── generated/            # gqlgen generated code
│   │   ├── dataloader/           # DataLoader implementations
│   │   ├── resolver/             # GraphQL resolvers
│   │   └── schema/               # GraphQL schema files
│   └── shared/                   # Shared utilities and types
├── sqlc/                         # sqlc configuration and queries
│   ├── queries/                  # SQL query files
│   └── sqlc.yaml                 # sqlc configuration
├── migrations/                   # Database migrations (Atlas HCL or SQL)
├── scripts/                      # Build and development scripts
├── test/                         # Integration and E2E tests
└── tools/                        # Development tools
```

#### Domain Layer (Core Business Logic)
```
internal/domain/
├── account/                      # Account aggregate
│   ├── account.go               # Account entity and value objects
│   ├── repository.go            # Account repository interface
│   ├── service.go               # Account domain service
│   ├── events.go                # Domain events
│   └── wire.go                  # Wire provider set for account domain
├── transaction/                  # Transaction aggregate
│   ├── transaction.go           # Transaction entity
│   ├── repository.go            # Transaction repository interface
│   ├── service.go               # Transaction domain service
│   ├── categorization.go        # Categorization rules
│   ├── matching.go              # Receipt matching logic
│   └── wire.go                  # Wire provider set for transaction domain
├── budget/                       # Budget aggregate
│   ├── budget.go                # Budget entity
│   ├── repository.go            # Budget repository interface
│   ├── service.go               # Budget domain service
│   ├── period.go                # Budget period calculations
│   └── wire.go                  # Wire provider set for budget domain
├── receipt/                      # Receipt aggregate
│   ├── receipt.go               # Receipt entity
│   ├── repository.go            # Receipt repository interface
│   ├── service.go               # Receipt domain service
│   ├── receipt_data.go          # Receipt data value objects
│   ├── matching.go              # Transaction matching logic
│   └── wire.go                  # Wire provider set for receipt domain
├── user/                         # User aggregate
│   ├── user.go                  # User entity
│   ├── repository.go            # User repository interface
│   ├── service.go               # User domain service
│   └── wire.go                  # Wire provider set for user domain
├── shared/                       # Shared domain concepts
│   ├── money.go                 # Money value object
│   ├── currency.go              # Currency handling
│   ├── errors.go                # Domain-specific errors
│   ├── events.go                # Domain event interfaces
│   └── specification.go         # Specification pattern
└── temporal/                     # Temporal workflow definitions
    ├── workflows/               # Workflow definitions
    ├── activities/              # Activity definitions
    └── interfaces.go            # Workflow interfaces
```

#### Application Layer (Use Cases/Services)
```
internal/application/
├── wire.go                       # Wire provider set for application layer
├── command/                      # Command handlers (writes)
│   ├── account/                 # Account commands
│   │   ├── create_account.go
│   │   ├── update_account.go
│   │   └── sync_account.go
│   ├── transaction/             # Transaction commands
│   │   ├── import_transactions.go
│   │   ├── categorize_transaction.go
│   │   └── match_receipt.go
│   ├── budget/                  # Budget commands
│   │   ├── create_budget.go
│   │   ├── update_budget.go
│   │   └── calculate_budget.go
│   └── receipt/                 # Receipt commands
│       ├── process_receipt.go
│       ├── match_receipt.go
│       └── verify_items.go
├── query/                        # Query handlers (reads)
│   ├── account/                 # Account queries
│   │   ├── get_account.go
│   │   ├── list_accounts.go
│   │   └── account_analytics.go
│   ├── transaction/             # Transaction queries
│   │   ├── get_transaction.go
│   │   ├── list_transactions.go
│   │   └── transaction_analytics.go
│   ├── budget/                  # Budget queries
│   │   ├── get_budget.go
│   │   ├── list_budgets.go
│   │   └── budget_analytics.go
│   └── receipt/                 # Receipt queries
│       ├── get_receipt.go
│       ├── list_receipts.go
│       └── unmatched_receipts.go
├── dto/                          # Data Transfer Objects
│   ├── account.go
│   ├── transaction.go
│   ├── budget.go
│   └── receipt.go
├── interfaces/                   # Application service interfaces
│   ├── repositories.go          # Repository interfaces
│   ├── services.go              # External service interfaces
│   └── events.go                # Event handling interfaces
└── workflow/                     # Temporal workflow implementations
    ├── import/                  # Import workflows
    ├── categorization/          # Categorization workflows
    └── synchronization/         # Sync workflows
```

#### Adapter Layer (Infrastructure Adapters)
```
internal/adapter/
├── wire.go                       # Wire provider set for adapter layer
├── repository/                   # Database adapters
│   ├── postgres/                # PostgreSQL implementations
│   │   ├── account.go           # Account repository
│   │   ├── transaction.go       # Transaction repository
│   │   ├── budget.go            # Budget repository
│   │   ├── receipt.go           # Receipt repository
│   │   ├── user.go              # User repository
│   │   └── wire.go              # Wire provider set for postgres adapters
│   └── memory/                  # In-memory implementations (testing)
│       ├── account.go
│       ├── transaction.go
│       └── user.go
├── external/                     # External service adapters
│   ├── plaid/                   # Plaid API client
│   │   ├── client.go
│   │   ├── accounts.go
│   │   ├── transactions.go
│   │   └── institutions.go
│   ├── firebase/                # Firebase Auth client
│   │   ├── client.go
│   │   └── auth.go
│   └── temporal/                # Temporal client adapter
│       ├── client.go
│       ├── worker.go
│       └── workflow_runner.go
├── messaging/                    # Event/message adapters
│   ├── temporal/                # Temporal for events
│   └── local/                   # Local event bus
└── cache/                        # Caching adapters
    ├── redis/                   # Redis cache
    └── memory/                  # In-memory cache
```

#### Infrastructure Layer (Framework Integration)
```
internal/infrastructure/
├── wire.go                       # Wire provider set for infrastructure layer
├── config/                       # Configuration management
│   ├── config.go                # Configuration structure
│   ├── database.go              # Database configuration
│   ├── auth.go                  # Authentication configuration
│   └── temporal.go              # Temporal configuration
├── database/                     # Database infrastructure
│   ├── connection.go            # Database connection setup
│   ├── migration.go             # Migration runner
│   ├── transaction.go           # Transaction management
│   └── health.go                # Health check
├── auth/                         # Authentication infrastructure
│   ├── firebase.go              # Firebase Auth setup
│   ├── middleware.go            # Auth middleware
│   └── context.go               # User context handling
├── logging/                      # Logging infrastructure
│   ├── logger.go                # Logger setup
│   ├── middleware.go            # Logging middleware
│   └── structured.go            # Structured logging
├── monitoring/                   # Monitoring infrastructure
│   ├── metrics.go               # Prometheus metrics
│   ├── tracing.go               # OpenTelemetry tracing
│   └── health.go                # Health checks
├── temporal/                     # Temporal infrastructure
│   ├── client.go                # Temporal client setup
│   ├── worker.go                # Worker setup
│   └── namespace.go             # Namespace configuration
└── server/                       # HTTP server infrastructure
    ├── server.go                # Server setup
    ├── middleware.go            # HTTP middleware
    ├── routes.go                # Route configuration
    └── graceful.go              # Graceful shutdown
```

#### GraphQL Presentation Layer

The GraphQL layer is placed within `internal/graph/` to prevent external packages from directly accessing resolvers and generated code, enforcing API access only through the defined GraphQL endpoint.

```
internal/graph/
├── wire.go                       # Wire provider set for GraphQL layer
├── schema/                       # GraphQL schema definitions
│   ├── schema.graphql           # Root schema
│   ├── user.graphql             # User types
│   ├── account.graphql          # Account types
│   ├── transaction.graphql      # Transaction types
│   ├── budget.graphql           # Budget types
│   ├── receipt.graphql          # Receipt types
│   ├── scalars.graphql          # Custom scalars
│   └── directives.graphql       # Custom directives
├── resolver/                     # GraphQL resolvers
│   ├── resolver.go              # Root resolver
│   ├── user.go                  # User resolvers
│   ├── account.go               # Account resolvers
│   ├── transaction.go           # Transaction resolvers
│   ├── budget.go                # Budget resolvers
│   ├── receipt.go               # Receipt resolvers
│   └── mutation.go              # Mutation resolvers
├── dataloader/                   # DataLoader implementations
│   ├── loaders.go               # Loader setup
│   ├── user.go                  # User loader
│   ├── account.go               # Account loader
│   ├── transaction.go           # Transaction loader
│   └── category.go              # Category loader
├── directive/                    # Custom directives
│   ├── auth.go                  # Authentication directive
│   ├── validate.go              # Validation directive
│   └── rate_limit.go            # Rate limiting directive
├── middleware/                   # GraphQL middleware
│   ├── auth.go                  # Authentication
│   ├── logging.go               # Request logging
│   ├── metrics.go               # Metrics collection
│   └── complexity.go            # Query complexity
└── generated/                    # Generated code (gqlgen)
    ├── generated.go             # gqlgen generated code
    └── models_gen.go            # Generated models
```

#### Wire Integration Patterns

##### Provider Sets by Layer (Research-Validated Best Practices)
```go
// internal/domain/account/wire.go
//+build !wireinject

package account

import "github.com/google/wire"

// DomainSet provides all account domain services
// Group providers by aggregate for clear boundaries
var DomainSet = wire.NewSet(
    NewService,                      // Account domain service
    NewCategorizationService,         // Categorization logic
    wire.Bind(new(ServiceInterface), new(*Service)),
)

// NewService creates a new account domain service
// Wire automatically handles error propagation
// Follows error handling patterns from ADR-20250120
func NewService(repo Repository, logger *slog.Logger) (*Service, error) {
    if repo == nil {
        return nil, errors.New("repository is required")
    }
    if logger == nil {
        return nil, errors.New("logger is required")
    }
    
    return &Service{
        repo:   repo,
        logger: logger.With(slog.String("service", "account")),
    }, nil
}
```

##### Application Layer Providers
```go
// internal/application/wire.go
package application

import (
    "github.com/google/wire"
    "github.com/djinn/internal/domain/account"
    "github.com/djinn/internal/domain/transaction"
)

// ApplicationSet provides all application layer services
var ApplicationSet = wire.NewSet(
    // Command handlers
    command.NewAccountCommandHandler,
    command.NewTransactionCommandHandler,
    command.NewBudgetCommandHandler,
    
    // Query handlers
    query.NewAccountQueryHandler,
    query.NewTransactionQueryHandler,
    query.NewBudgetQueryHandler,
    
    // Bind interfaces
    wire.Bind(new(command.AccountCommandHandlerInterface), new(*command.AccountCommandHandler)),
    wire.Bind(new(query.AccountQueryHandlerInterface), new(*query.AccountQueryHandler)),
)
```

##### Infrastructure Providers
```go
// internal/infrastructure/repository/postgres/wire.go
package postgres

import (
    "github.com/google/wire"
    "github.com/djinn/internal/domain/account"
    "github.com/djinn/internal/domain/transaction"
)

// RepositorySet provides all PostgreSQL repository implementations
var RepositorySet = wire.NewSet(
    NewAccountRepository,
    NewTransactionRepository,
    NewBudgetRepository,
    NewReceiptRepository,
    NewUserRepository,
    
    // Bind to domain interfaces
    wire.Bind(new(account.RepositoryInterface), new(*AccountRepository)),
    wire.Bind(new(transaction.RepositoryInterface), new(*TransactionRepository)),
)

// NewAccountRepository creates a PostgreSQL account repository
func NewAccountRepository(db *sql.DB) *AccountRepository {
    return &AccountRepository{db: db}
}
```

##### Main Injector Setup (Production Pattern)
```go
// cmd/server/wire.go
//go:build wireinject
// +build wireinject

package main

import (
    "context"
    "database/sql"
    "github.com/google/wire"
    "github.com/djinn/internal/infrastructure/repository/postgres"
    "github.com/djinn/internal/application"
    "github.com/djinn/internal/domain/account"
    "github.com/djinn/internal/domain/transaction"
    "github.com/djinn/internal/infrastructure"
    "github.com/djinn/graph"
)

// ServerSet combines all layer provider sets
// Organized by architectural layer for clarity
var ServerSet = wire.NewSet(
    // Infrastructure (lowest layer)
    infrastructure.InfrastructureSet,
    
    // Adapters (repository implementations)
    postgres.RepositorySet,
    
    // Domain (business logic)
    account.DomainSet,
    transaction.DomainSet,
    budget.DomainSet,
    receipt.DomainSet,
    user.DomainSet,
    
    // Application
    application.ApplicationSet,
    
    // Presentation
    graph.GraphQLSet,
    
    // Server
    NewServer,
)

// InitializeServer creates a fully configured server
// Integrates error handling and logging from ADR-20250120
func InitializeServer(ctx context.Context, config *Config) (*Server, func(), error) {
    wire.Build(
        ServerSet,
        logger.LoggerSet,      // slog infrastructure (ADR-20250120)
        errors.ErrorSet,       // Custom error types and handlers
        middleware.MiddlewareSet, // Request logging & error middleware
    )
    return &Server{}, nil, nil
}
```

##### Generated Injector Usage
```go
// cmd/server/main.go
package main

import (
    "context"
    "log"
    "os"
)

func main() {
    ctx := context.Background()
    
    // Load configuration
    config, err := LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }
    
    // Initialize server using Wire (includes cleanup function)
    server, cleanup, err := InitializeServer(ctx, config)
    defer cleanup() // Ensure proper resource cleanup
    if err != nil {
        log.Fatalf("Failed to initialize server: %v", err)
    }
    
    // Start server
    if err := server.Start(ctx); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}
```

##### Testing with Wire (Mock Injection Pattern)
```go
// internal/infrastructure/repository/postgres/wire_test.go
//+build wireinject

package postgres_test

import (
    "testing"
    "github.com/google/wire"
    "github.com/djinn/internal/domain/account"
    "github.com/djinn/internal/test/mocks"
)

// TestSet provides test-specific dependencies with mocks
var TestSet = wire.NewSet(
    NewTestDB,                       // Test database connection
    postgres.NewAccountRepository,
    // Bind mock implementations for testing
    wire.Bind(new(account.RepositoryInterface), new(*postgres.AccountRepository)),
)

// MockSet provides pure mock dependencies (no DB)
var MockSet = wire.NewSet(
    mocks.NewMockAccountRepository,
    wire.Bind(new(account.RepositoryInterface), new(*mocks.MockAccountRepository)),
)

// InitializeTestAccountRepository creates real repository with test DB
func InitializeTestAccountRepository(t *testing.T) (account.RepositoryInterface, func(), error) {
    wire.Build(TestSet, NewTestDB)
    return nil, nil, nil
}

// InitializeMockAccountRepository creates mock repository for unit tests
func InitializeMockAccountRepository(t *testing.T) (account.RepositoryInterface, error) {
    wire.Build(MockSet)
    return nil, nil
}
```

### 2. Frontend Flutter Module Structure (Feature-First)

#### Root Directory Structure
```
/djinn-mobile/
├── lib/
│   ├── main.dart                # Application entry point
│   ├── app/                     # App-level configuration
│   ├── core/                    # Core utilities and shared code
│   ├── features/                # Feature modules
│   ├── shared/                  # Shared UI components and utilities
│   └── generated/               # Generated code (Ferry, etc.)
├── test/                        # Tests
├── integration_test/            # Integration tests
├── assets/                      # Static assets
├── android/                     # Android-specific code
├── ios/                         # iOS-specific code
└── web/                         # Web-specific code (future)
```

#### Feature-First Organization
```
lib/features/
├── authentication/              # Authentication feature
│   ├── data/                   # Data layer
│   │   ├── datasources/        # Remote/local data sources
│   │   ├── models/             # Data models
│   │   └── repositories/       # Repository implementations
│   ├── domain/                 # Domain layer
│   │   ├── entities/           # Domain entities
│   │   ├── repositories/       # Repository interfaces
│   │   └── usecases/           # Use cases
│   └── presentation/           # Presentation layer
│       ├── pages/              # Screens/pages
│       ├── widgets/            # Feature-specific widgets
│       ├── providers/          # Riverpod providers
│       └── controllers/        # State controllers
├── accounts/                    # Account management
│   ├── data/
│   │   ├── datasources/
│   │   │   ├── account_remote_datasource.dart
│   │   │   └── account_local_datasource.dart
│   │   ├── models/
│   │   │   ├── account_model.dart
│   │   │   └── institution_model.dart
│   │   └── repositories/
│   │       └── account_repository_impl.dart
│   ├── domain/
│   │   ├── entities/
│   │   │   ├── account.dart
│   │   │   └── institution.dart
│   │   ├── repositories/
│   │   │   └── account_repository.dart
│   │   └── usecases/
│   │       ├── get_accounts.dart
│   │       ├── sync_accounts.dart
│   │       └── connect_plaid_account.dart
│   └── presentation/
│       ├── pages/
│       │   ├── accounts_page.dart
│       │   ├── account_details_page.dart
│       │   └── connect_account_page.dart
│       ├── widgets/
│       │   ├── account_card.dart
│       │   ├── balance_chart.dart
│       │   └── sync_status_indicator.dart
│       ├── providers/
│       │   ├── accounts_provider.dart
│       │   └── account_sync_provider.dart
│       └── controllers/
│           └── account_controller.dart
├── transactions/                # Transaction management
│   ├── data/
│   │   ├── datasources/
│   │   │   ├── transaction_remote_datasource.dart
│   │   │   └── transaction_local_datasource.dart
│   │   ├── models/
│   │   │   ├── transaction_model.dart
│   │   │   └── category_model.dart
│   │   └── repositories/
│   │       └── transaction_repository_impl.dart
│   ├── domain/
│   │   ├── entities/
│   │   │   ├── transaction.dart
│   │   │   ├── category.dart
│   │   │   └── money.dart
│   │   ├── repositories/
│   │   │   └── transaction_repository.dart
│   │   └── usecases/
│   │       ├── get_transactions.dart
│   │       ├── categorize_transaction.dart
│   │       └── search_transactions.dart
│   └── presentation/
│       ├── pages/
│       │   ├── transactions_page.dart
│       │   ├── transaction_details_page.dart
│       │   └── categorize_page.dart
│       ├── widgets/
│       │   ├── transaction_list.dart
│       │   ├── transaction_card.dart
│       │   ├── category_selector.dart
│       │   └── amount_formatter.dart
│       ├── providers/
│       │   ├── transactions_provider.dart
│       │   └── categories_provider.dart
│       └── controllers/
│           └── transaction_controller.dart
├── budgets/                     # Budget management
│   ├── data/
│   ├── domain/
│   └── presentation/
├── receipts/                    # Receipt management
│   ├── data/
│   │   ├── datasources/
│   │   │   ├── receipt_remote_datasource.dart
│   │   │   ├── receipt_local_datasource.dart
│   │   │   └── ml_kit_scanner.dart  # Google ML Kit OCR
│   │   ├── models/
│   │   │   ├── receipt_model.dart
│   │   │   └── receipt_item_model.dart
│   │   └── repositories/
│   │       └── receipt_repository_impl.dart
│   ├── domain/
│   │   ├── entities/
│   │   │   ├── receipt.dart
│   │   │   └── receipt_item.dart
│   │   ├── repositories/
│   │   │   └── receipt_repository.dart
│   │   └── usecases/
│   │       ├── scan_receipt.dart
│   │       ├── process_scanned_data.dart
│   │       └── match_receipt.dart
│   └── presentation/
│       ├── pages/
│       │   ├── receipt_scanner_page.dart
│       │   ├── receipt_details_page.dart
│       │   └── receipt_matching_page.dart
│       ├── widgets/
│       │   ├── camera_preview.dart
│       │   ├── receipt_item_list.dart
│       │   └── matching_suggestions.dart
│       ├── providers/
│       │   ├── receipt_provider.dart
│       │   └── camera_provider.dart
│       └── controllers/
│           └── receipt_controller.dart
└── analytics/                   # Analytics and reporting
    ├── data/
    ├── domain/
    └── presentation/
```

#### Core and Shared Modules
```
lib/core/
├── constants/                   # App constants
│   ├── app_constants.dart
│   ├── api_constants.dart
│   └── ui_constants.dart
├── error/                       # Error handling
│   ├── exceptions.dart
│   ├── failures.dart
│   └── error_handler.dart
├── network/                     # Network layer
│   ├── graphql_client.dart     # Ferry GraphQL client
│   ├── network_info.dart       # Network connectivity
│   └── interceptors.dart       # Request/response interceptors
├── database/                    # Local database
│   ├── drift_database.dart     # Drift database setup
│   ├── tables/                 # Database tables
│   └── migrations/             # Database migrations
├── utils/                       # Utilities
│   ├── date_utils.dart
│   ├── money_utils.dart
│   ├── validation_utils.dart
│   └── image_utils.dart
├── services/                    # Core services
│   ├── auth_service.dart       # Authentication service
│   ├── storage_service.dart    # Local storage
│   ├── sync_service.dart       # Data synchronization
│   └── ml_kit_service.dart     # Google ML Kit text recognition
└── dependency_injection/        # DI setup
    ├── injection_container.dart
    └── providers.dart

lib/shared/
├── widgets/                     # Reusable UI components
│   ├── buttons/
│   │   ├── primary_button.dart
│   │   └── secondary_button.dart
│   ├── inputs/
│   │   ├── text_field.dart
│   │   ├── amount_input.dart
│   │   └── date_picker.dart
│   ├── cards/
│   │   ├── info_card.dart
│   │   └── stat_card.dart
│   ├── lists/
│   │   ├── infinite_list.dart
│   │   └── searchable_list.dart
│   └── navigation/
│       ├── bottom_nav_bar.dart
│       └── app_drawer.dart
├── themes/                      # App theming
│   ├── app_theme.dart
│   ├── colors.dart
│   ├── typography.dart
│   └── spacing.dart
├── extensions/                  # Dart extensions
│   ├── context_extensions.dart
│   ├── date_extensions.dart
│   └── money_extensions.dart
└── providers/                   # Global providers
    ├── theme_provider.dart
    ├── locale_provider.dart
    └── app_state_provider.dart
```

### 3. Package Naming Conventions

#### Go Backend Packages
```go
// Domain packages - singular, descriptive nouns
package account    // ✓ Good
package transaction // ✓ Good
package budget     // ✓ Good

// Avoid plurals in package names
package accounts   // ✗ Avoid
package transactions // ✗ Avoid

// Service packages - what they do
package categorization // ✓ Good
package matching      // ✓ Good
package synchronization // ✓ Good

// Infrastructure packages - technology/protocol
package postgres   // ✓ Good
package temporal   // ✓ Good
package firebase   // ✓ Good
```

#### Flutter/Dart Packages
```dart
// Feature packages - plural nouns
lib/features/accounts/     // ✓ Good
lib/features/transactions/ // ✓ Good
lib/features/budgets/      // ✓ Good

// Core packages - singular nouns
lib/core/network/         // ✓ Good
lib/core/database/        // ✓ Good
lib/core/utils/           // ✓ Good

// Shared packages - descriptive
lib/shared/widgets/       // ✓ Good
lib/shared/themes/        // ✓ Good
lib/shared/extensions/    // ✓ Good
```

### 4. Import Rules and Dependency Management

#### Go Import Rules
```go
// Import order (enforced by goimports)
// 1. Standard library
import (
    "context"
    "fmt"
    "time"
)

// 2. External dependencies
import (
    "github.com/99designs/gqlgen/graphql"
    "github.com/google/uuid"
)

// 3. Internal packages (domain first, then infrastructure)
import (
    "github.com/djinn/internal/domain/account"
    "github.com/djinn/internal/domain/transaction"
    "github.com/djinn/internal/application/command"
    "github.com/djinn/internal/infrastructure/repository/postgres"
)

// Dependency direction rules:
// Domain → NOTHING (pure business logic)
// Application → Domain only
// Adapter → Domain + Application
// Infrastructure → Any
// GraphQL → Application + Domain (not adapters directly)
```

#### Flutter Import Rules
```dart
// Import order (enforced by flutter_lints)
// 1. Dart/Flutter framework
import 'dart:async';
import 'package:flutter/material.dart';

// 2. External packages
import 'package:riverpod/riverpod.dart';
import 'package:ferry/ferry.dart';

// 3. Internal imports (relative paths)
import '../../../core/error/failures.dart';
import '../../domain/entities/account.dart';
import '../providers/account_provider.dart';

// Dependency direction rules:
// Domain → Core only
// Data → Domain + Core
// Presentation → Domain + Data + Core + Shared
```

### 5. Testing Structure

#### Go Backend Testing
```
test/
├── unit/                        # Unit tests (mirror source structure)
│   ├── domain/
│   │   ├── account/
│   │   ├── transaction/
│   │   └── budget/
│   ├── application/
│   │   ├── command/
│   │   └── query/
│   └── adapter/
│       └── repository/
├── integration/                 # Integration tests
│   ├── api/                    # GraphQL API tests
│   ├── database/               # Database integration
│   ├── temporal/               # Workflow tests
│   └── external/               # External service tests
├── e2e/                        # End-to-end tests
│   ├── scenarios/              # User scenarios
│   └── fixtures/               # Test data
├── performance/                # Performance tests
│   ├── load/                   # Load testing
│   └── stress/                 # Stress testing
└── testdata/                   # Test fixtures and data
    ├── sql/                    # SQL test data
    ├── json/                   # JSON fixtures
    └── csv/                    # CSV import tests
```

#### Flutter Frontend Testing
```
test/
├── unit/                       # Unit tests
│   ├── core/
│   ├── features/
│   │   ├── accounts/
│   │   │   ├── data/
│   │   │   ├── domain/
│   │   │   └── presentation/
│   │   └── transactions/
│   └── shared/
├── widget/                     # Widget tests
│   ├── features/
│   └── shared/
└── integration/                # Integration tests
    ├── app_test.dart
    └── feature_flows/

integration_test/
├── app_integration_test.dart   # Full app integration
├── auth_flow_test.dart         # Authentication flow
├── transaction_flow_test.dart  # Transaction management
└── offline_sync_test.dart      # Offline synchronization
```

### 6. Code Generation Integration

#### Backend Code Generation
```makefile
# Makefile for code generation
.PHONY: generate
generate: generate-wire generate-sqlc generate-gqlgen generate-dataloaders

.PHONY: generate-wire
generate-wire:
	@echo "Generating Wire dependency injection..."
	@wire ./cmd/server/
	@wire ./cmd/worker/

.PHONY: generate-sqlc
generate-sqlc:
	@echo "Generating database code..."
	@sqlc generate  # Outputs to internal/infrastructure/persistence/postgres/generated/

.PHONY: generate-gqlgen  
generate-gqlgen:
	@echo "Generating GraphQL code..."
	@go run github.com/99designs/gqlgen generate

.PHONY: generate-dataloaders
generate-dataloaders:
	@echo "Generating dataloaders..."
	@go generate ./graph/dataloader/...

# Development workflow
.PHONY: dev
dev: generate
	@air -c .air.toml

# Wire-specific commands
.PHONY: wire-check
wire-check:
	@echo "Checking Wire dependencies..."
	@wire check ./...

.PHONY: wire-graph
wire-graph:
	@echo "Showing Wire dependency graph..."
	@wire show ./cmd/server/

# Clean generated files
.PHONY: clean-wire
clean-wire:
	@echo "Cleaning Wire generated files..."
	@find . -name "wire_gen.go" -delete
	@echo "Generating dependency graph visualization..."
	@wire show ./cmd/server/ | dot -Tpng -o dependency-graph.png
```

#### Frontend Code Generation
```yaml
# pubspec.yaml
dev_dependencies:
  build_runner: ^2.4.7
  ferry_generator: ^0.8.0
  drift_dev: ^2.14.0
  json_annotation: ^4.8.1
  json_serializable: ^6.7.1

# Generation commands
scripts:
  generate: "dart run build_runner build --delete-conflicting-outputs"
  watch: "dart run build_runner watch --delete-conflicting-outputs"
```

### 7. Error Handling Integration with Wire

#### Provider Error Handling Patterns

Following the error handling strategy from ADR-20250120, all Wire providers must properly handle and propagate errors:

```go
// internal/infrastructure/logger/wire.go
package logger

import (
    "log/slog"
    "os"
    "github.com/google/wire"
)

// LoggerSet provides logging infrastructure
var LoggerSet = wire.NewSet(
    NewLogger,
    NewServiceLoggers,
)

// NewLogger creates the root logger following ADR-20250120
func NewLogger(cfg *Config) (*slog.Logger, error) {
    var handler slog.Handler
    
    if cfg.Environment == "production" {
        handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
            Level:     slog.LevelInfo,
            AddSource: true,
        })
    } else {
        handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
            Level:     slog.LevelDebug,
            AddSource: true,
        })
    }
    
    logger := slog.New(handler)
    slog.SetDefault(logger)
    
    return logger, nil
}

// ServiceLoggers provides domain-specific loggers
type ServiceLoggers struct {
    AccountLogger     *slog.Logger
    TransactionLogger *slog.Logger
    ReceiptLogger     *slog.Logger
}

func NewServiceLoggers(logger *slog.Logger) *ServiceLoggers {
    return &ServiceLoggers{
        AccountLogger:     logger.With(slog.String("service", "account")),
        TransactionLogger: logger.With(slog.String("service", "transaction")),
        ReceiptLogger:     logger.With(slog.String("service", "receipt")),
    }
}
```

#### Error Types Provider

```go
// internal/domain/errors/wire.go
package errors

import "github.com/google/wire"

// ErrorSet provides custom error types
var ErrorSet = wire.NewSet(
    NewErrorHandler,
    NewCircuitBreaker,
)

// ErrorHandler centralizes error handling logic
type ErrorHandler struct {
    logger *slog.Logger
}

func NewErrorHandler(logger *slog.Logger) *ErrorHandler {
    return &ErrorHandler{
        logger: logger.With(slog.String("component", "error_handler")),
    }
}

// HandleDomainError processes domain-specific errors
func (h *ErrorHandler) HandleDomainError(err error) error {
    switch e := err.(type) {
    case ValidationError:
        h.logger.Warn("validation error", 
            slog.String("field", e.Field),
            slog.Any("value", e.Value))
        return e
    case NotFoundError:
        h.logger.Info("resource not found",
            slog.String("resource", e.Resource),
            slog.String("id", e.ID))
        return e
    default:
        h.logger.Error("unexpected error", slog.Any("error", err))
        return fmt.Errorf("internal error: %w", err)
    }
}
```

#### Repository with Error Handling

```go
// internal/infrastructure/repository/postgres/account_repository.go
package postgres

type AccountRepository struct {
    db          *sql.DB
    logger      *slog.Logger
    errHandler  *errors.ErrorHandler
}

func NewAccountRepository(db *sql.DB, logger *slog.Logger, errHandler *errors.ErrorHandler) (*AccountRepository, error) {
    if db == nil {
        return nil, errors.New("database connection required")
    }
    
    return &AccountRepository{
        db:         db,
        logger:     logger.With(slog.String("repository", "account")),
        errHandler: errHandler,
    }, nil
}

func (r *AccountRepository) GetByID(ctx context.Context, id string) (*domain.Account, error) {
    r.logger.Debug("fetching account", slog.String("id", id))
    
    var account domain.Account
    err := r.db.QueryRowContext(ctx, "SELECT * FROM accounts WHERE id = $1", id).Scan(&account)
    
    if err == sql.ErrNoRows {
        return nil, errors.NotFoundError{Resource: "account", ID: id}
    }
    
    if err != nil {
        r.logger.Error("database error", 
            slog.String("operation", "GetByID"),
            slog.String("id", id),
            slog.Any("error", err))
        return nil, fmt.Errorf("failed to fetch account %s: %w", id, err)
    }
    
    return &account, nil
}
```

#### Middleware Integration with Wire

```go
// internal/infrastructure/middleware/wire.go
package middleware

import (
    "github.com/google/wire"
    "log/slog"
)

// MiddlewareSet provides HTTP middleware components
var MiddlewareSet = wire.NewSet(
    NewRequestLogger,
    NewErrorHandler,
    NewCircuitBreaker,
)

// RequestLogger middleware from ADR-20250120
func NewRequestLogger(logger *slog.Logger) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Implementation from ADR-20250120
            correlationID := r.Header.Get("X-Correlation-ID")
            if correlationID == "" {
                correlationID = uuid.New().String()
            }
            
            ctx := context.WithValue(r.Context(), "correlation_id", correlationID)
            reqLogger := logger.With(
                slog.String("correlation_id", correlationID),
                slog.String("method", r.Method),
                slog.String("path", r.URL.Path),
            )
            
            reqLogger.Info("request started")
            next.ServeHTTP(w, r.WithContext(ctx))
            reqLogger.Info("request completed")
        })
    }
}

// GraphQL resolver with error handling
type Resolver struct {
    logger      *slog.Logger
    errHandler  *errors.ErrorHandler
    circuitBreaker *CircuitBreaker
}

func NewResolver(logger *slog.Logger, errHandler *errors.ErrorHandler, cb *CircuitBreaker) *Resolver {
    return &Resolver{
        logger:         logger.With(slog.String("component", "graphql")),
        errHandler:     errHandler,
        circuitBreaker: cb,
    }
}
```

### 8. Wire Common Pitfalls and Solutions (Research-Based)

#### Common Pitfalls

1. **Circular Dependencies**
   - **Problem**: Wire detects circular dependencies at compile time
   - **Solution**: Use interfaces to break dependency cycles, restructure into layers

2. **Missing Providers**
   - **Problem**: "No provider found" errors during wire generation
   - **Solution**: Ensure all dependencies are provided in wire sets

3. **Build Tag Confusion**
   - **Problem**: wire.go and wire_gen.go conflicts
   - **Solution**: Always use `//+build wireinject` on wire.go files

4. **Provider Cleanup**
   - **Problem**: Resources not properly cleaned up
   - **Solution**: Use cleanup functions in providers that return them

#### Best Practices Summary

```go
// tools.go - Ensure Wire is tracked as dependency
//+build tools

package tools

import _ "github.com/google/wire/cmd/wire"

// .gitignore - Never commit generated files
**/wire_gen.go

// wire.go - Always use build constraints
//+build wireinject

// Provider functions - Return errors for proper handling
func NewService(db *sql.DB) (*Service, error) {
    if db == nil {
        return nil, errors.New("database required")
    }
    return &Service{db: db}, nil
}

// Cleanup functions - For resource management
func NewDatabase(cfg *Config) (*sql.DB, func(), error) {
    db, err := sql.Open("postgres", cfg.DSN)
    if err != nil {
        return nil, nil, err
    }
    cleanup := func() { db.Close() }
    return db, cleanup, nil
}
```

## Consequences

### Positive
- **Clear Architecture**: Clean separation of concerns with well-defined layers
- **Testability**: Architecture facilitates comprehensive testing at all levels
- **Error Handling Integration**: Seamless integration with slog and custom error types (ADR-20250120)
- **Type Safety**: End-to-end type safety from database to UI
- **Scalability**: Module structure supports team growth and feature expansion
- **Maintainability**: Clear boundaries prevent spaghetti code and technical debt
- **Code Generation**: Integrated tooling reduces boilerplate and errors
- **Feature Velocity**: Feature-first organization enables rapid development
- **Dependency Management**: Clear import rules prevent circular dependencies
- **Compile-time DI**: Wire eliminates runtime reflection overhead and provides type safety
- **Dependency Visibility**: Wire provider functions make dependencies explicit and traceable
- **Testing Simplification**: Wire enables easy mock injection through interface binding
- **Build Integration**: Wire generation integrates seamlessly with existing code generation workflow
- **Graph Visualization**: Wire can generate dependency graphs for architecture documentation
- **Zero Runtime Cost**: All dependency injection happens at compile time

### Negative
- **Initial Complexity**: Architecture may seem over-engineered for MVP
- **Learning Curve**: Team needs to understand clean architecture principles and Wire concepts
- **Code Generation Setup**: Initial tooling setup requires configuration for multiple tools
- **File Count**: More files and directories to navigate initially, including Wire provider files
- **Abstraction Overhead**: Multiple layers may slow simple changes
- **Wire Learning Curve**: Developers need to understand Wire provider patterns and build tags
- **Build Tag Management**: Requires careful management of build tags for Wire generation

### Risks
- **Over-Architecture**: Risk of over-engineering simple features
- **Inconsistent Application**: Team may not follow architectural patterns consistently
- **Circular Dependencies**: Improper imports could create circular dependencies
- **Code Generation Breaks**: Generated code changes could break builds
- **Testing Complexity**: Multiple test types may be overwhelming initially
- **Wire Provider Complexity**: Complex provider sets could become difficult to maintain
- **Build Tag Conflicts**: Incorrect build tags could cause Wire generation issues
- **Interface Binding Errors**: Incorrect interface bindings could cause runtime failures

## Alternatives Considered

### Option A: Monolithic Package Structure
- **Description**: Single large package per technology with minimal separation
- **Pros**: Simpler initially, faster to prototype, fewer files
- **Cons**: Becomes unmaintainable as codebase grows, difficult to test, tight coupling
- **Reason for not choosing**: Doesn't scale with team or feature growth

### Option B: Microservices Architecture
- **Description**: Split backend into multiple services by domain
- **Pros**: Independent deployment, team autonomy, technology diversity
- **Cons**: Distributed system complexity, data consistency challenges, overhead for small team
- **Reason for not choosing**: Unnecessary complexity for current team size and requirements

### Option C: Framework-First Organization (Flutter)
- **Description**: Organize by framework concepts (pages, widgets, providers)
- **Pros**: Familiar to Flutter developers, matches framework structure
- **Cons**: Features scattered across directories, difficult to find related code
- **Reason for not choosing**: Feature-first organization better for maintainability

## Implementation Notes

### Migration Strategy
1. **Week 1-2**: Set up basic directory structure and package organization
2. **Week 3-4**: Implement core domain entities and repository interfaces
3. **Week 5-6**: Add application layer use cases and command/query handlers
4. **Week 7-8**: Implement adapter layer and infrastructure components
5. **Week 9-10**: Set up GraphQL layer and data loaders
6. **Week 11-12**: Complete testing structure and CI/CD integration

### Development Workflow
```bash
# Backend development
make generate              # Generate all code
make test                 # Run all tests
make lint                 # Run linters
make build                # Build application

# Frontend development
flutter packages get      # Get dependencies
dart run build_runner build  # Generate code
flutter test             # Run tests
flutter analyze          # Static analysis
flutter build apk        # Build application
```

### Code Quality Gates
- **Unit Test Coverage**: Minimum 80% for domain and application layers
- **Integration Tests**: All API endpoints and database operations
- **Linting**: Zero linting errors (enforced by CI)
- **Type Safety**: All code must pass strict type checking
- **Import Rules**: Automated checking of import order and circular dependencies

### Monitoring and Success Metrics
- **Build Times**: Backend build < 30 seconds, Flutter build < 2 minutes
- **Test Execution**: All tests run in < 5 minutes
- **Code Quality**: Maintain technical debt ratio < 5%
- **Developer Velocity**: Features delivered without architectural refactoring
- **Onboarding Time**: New developers productive within 1 week

## References
- [Clean Architecture by Robert Martin](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Hexagonal Architecture Pattern](https://alistair.cockburn.us/hexagonal-architecture/)
- [Go Project Layout](https://github.com/golang-standards/project-layout)
- [Flutter App Architecture Guide](https://docs.flutter.dev/development/data-and-backend/state-mgmt/options)
- [Domain-Driven Design](https://martinfowler.com/tags/domain%20driven%20design.html)
- [Google Wire Documentation](https://github.com/google/wire)
- [Wire User Guide](https://github.com/google/wire/blob/main/docs/guide.md)
- [Wire Best Practices](https://github.com/google/wire/blob/main/docs/best-practices.md)
- [Compile-time Dependency Injection in Go](https://blog.golang.org/wire)
- [Wire DI Patterns Research](../research/wire-dependency-injection-patterns.md)
- [ADR-20250120: Error Handling & Logging Strategy](ADR-20250120-error-handling-logging-strategy.md)
- ADR-20250812: Personal Finance Tech Stack Selection
- ADR-20250819: Data Architecture and Schema Design
- ADR-20250819: API Design Standards
- ADR-20250820: Testing Strategy

## Decision Makers
- Author: Archie (System Architect)
- Reviewers: [Pending]
- Approver: [Pending]  
- Date: 2025-08-20

## Revision History
- 2025-08-20: Initial comprehensive draft with clean architecture focus
- 2025-08-26: Updated database structure to reflect actual implementation with improved internal package organization