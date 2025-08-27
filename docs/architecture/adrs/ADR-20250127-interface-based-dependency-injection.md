# ADR-20250127: Interface-Based Dependency Injection

## Status
Accepted

## Context
During implementation of the backend services, we discovered that concrete type dependencies create significant testing challenges and violate SOLID principles, particularly the Dependency Inversion Principle. The initial architecture defined layers (domain, application, infrastructure) but didn't enforce strict interface boundaries.

Key issues identified:
- Concrete dependencies make unit testing difficult, requiring complex test setup
- Direct coupling between layers violates clean architecture principles
- Circular dependencies emerged when concrete types were used across layers
- Parallel development was blocked when teams depended on concrete implementations
- Mock generation and maintenance was cumbersome without clear interfaces

## Decision
We will enforce interface-based dependency injection throughout the entire codebase:

### 1. Domain Layer Interfaces
All domain services MUST expose interfaces that define their contracts:
```go
// Domain layer defines the interface
type ServiceInterface interface {
    CreateUser(ctx context.Context, firebaseUID, email, name string, profileImageURL *string) (*User, error)
    GetUser(ctx context.Context, id uuid.UUID) (*User, error)
    UpdateUser(ctx context.Context, id uuid.UUID, email *string, name *string, profileImageURL *string) (*User, error)
    DeleteUser(ctx context.Context, id uuid.UUID) error
}
```

### 2. Application Layer Dependencies
Application handlers MUST depend only on domain interfaces:
```go
type CreateUserHandler struct {
    userService user.ServiceInterface  // Interface, not concrete type
    logger      *slog.Logger
}
```

### 3. Repository Pattern with Interfaces
All repositories MUST be defined as interfaces in the domain layer:
```go
// Domain defines what it needs
type RepositoryInterface interface {
    Create(ctx context.Context, user *User) error
    GetByID(ctx context.Context, id uuid.UUID) (*User, error)
    Update(ctx context.Context, user *User) error
    Delete(ctx context.Context, id uuid.UUID) error
}
```

### 4. Infrastructure Layer Implementation
Infrastructure provides concrete implementations that satisfy domain interfaces:
```go
// Infrastructure implements domain interface
type PostgresUserRepository struct {
    db *sql.DB
}

func (r *PostgresUserRepository) Create(ctx context.Context, user *User) error {
    // Implementation details
}
```

### 5. Wire Dependency Injection
Google Wire binds concrete implementations to interfaces at compile time:
```go
var DomainSet = wire.NewSet(
    NewService,
    wire.Bind(new(ServiceInterface), new(*Service)),
)

var RepositorySet = wire.NewSet(
    NewPostgresUserRepository,
    wire.Bind(new(user.RepositoryInterface), new(*PostgresUserRepository)),
)
```

### 6. Test Mocking Strategy
Mock implementations are generated for all interfaces:
```go
// Mock generation via mockery
//go:generate mockery --name ServiceInterface --output ./mocks --filename mock_user_service.go

// Test usage
mockService := &mocks.MockUserService{}
mockService.On("CreateUser", mock.Anything, ...).Return(expectedUser, nil)
```

## Consequences

### Positive
- **Enhanced Testability**: Unit tests can easily mock dependencies without complex setup
- **Clean Architecture Boundaries**: Clear separation between layers with explicit contracts
- **Elimination of Circular Dependencies**: Interfaces break circular dependency chains
- **Parallel Development**: Teams can work independently using interface contracts
- **Flexibility**: Easy to swap implementations (e.g., PostgreSQL to MongoDB)
- **Compile-Time Safety**: Wire ensures all dependencies are satisfied at compile time
- **Better Documentation**: Interfaces serve as clear API documentation

### Negative
- **Initial Overhead**: Requires defining interfaces upfront
- **More Files**: Separate interface definitions increase file count
- **Learning Curve**: Developers must understand dependency inversion principle
- **Mock Maintenance**: Generated mocks need regeneration when interfaces change

## Implementation Examples

### Current Implementation Structure
```
internal/
├── domain/
│   └── user/
│       ├── service_interface.go    # Interface definition
│       ├── service.go              # Concrete implementation
│       ├── repository.go           # Repository interface
│       └── wire.go                 # Wire provider set
├── application/
│   └── command/
│       └── user/
│           ├── create_user.go      # Depends on ServiceInterface
│           └── create_user_test.go # Uses mock service
└── infrastructure/
    └── repository/
        └── postgres/
            └── user_repository.go   # Implements RepositoryInterface
```

### Testing Pattern
```go
func TestCreateUserHandler_Handle(t *testing.T) {
    mockService := &mocks.MockUserService{}
    handler := NewCreateUserHandler(mockService, logger)
    
    mockService.On("CreateUser", ...).Return(expectedUser, nil)
    
    result, err := handler.Handle(ctx, input)
    
    mockService.AssertExpectations(t)
}
```

## References
- [Clean Architecture by Robert C. Martin](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Dependency Inversion Principle](https://en.wikipedia.org/wiki/Dependency_inversion_principle)
- [Google Wire Documentation](https://github.com/google/wire)
- ADR-20250820: Code Organization & Module Structure (defines layer separation)
- ADR-20250820: Testing Strategy (defines mock strategy)

## Related Patterns
- Repository Pattern
- Hexagonal Architecture (Ports and Adapters)
- Dependency Injection
- Interface Segregation Principle