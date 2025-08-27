# Domain-Driven Design (DDD) Service Layer Placement in Clean Architecture

## Research Date
2025-08-27

## Executive Summary

Domain services and application services serve distinct purposes in Clean Architecture and DDD. Domain services contain business logic that doesn't naturally fit within a single aggregate, while application services orchestrate use cases and coordinate domain operations.

## Service Layer Architecture in Clean Architecture

### 1. Domain Services (Domain Layer)

**Location**: `internal/domain/{aggregate}/service.go`

**Purpose**: Contains business logic that:
- Spans multiple aggregates
- Represents domain concepts that aren't entities or value objects  
- Encapsulates complex business rules
- Doesn't fit naturally within a single aggregate root

**Example from current codebase**:
```go
// internal/domain/user/service.go
type Service struct {
    repo   Repository
    logger *slog.Logger
}

func (s *Service) CreateUser(ctx context.Context, firebaseUID, email, name string, profileImageURL *string) (*User, error) {
    // Business logic: check if user already exists
    exists, err := s.repo.Exists(ctx, firebaseUID)
    if err != nil {
        return nil, fmt.Errorf("failed to check user existence: %w", err)
    }
    if exists {
        return nil, ErrUserAlreadyExists
    }

    // Create domain entity with validation
    user, err := NewUser(firebaseUID, email, name, profileImageURL)
    if err != nil {
        return nil, err
    }

    // Persist to repository
    if err := s.repo.Create(ctx, user); err != nil {
        return nil, fmt.Errorf("failed to create user: %w", err)
    }

    return user, nil
}
```

### 2. Application Services (Application Layer)

**Location**: `internal/application/command/{aggregate}/{command}.go` or `internal/application/query/{aggregate}/{query}.go`

**Purpose**: 
- Orchestrate use cases
- Handle application concerns (logging, validation, transactions)
- Convert between DTOs and domain objects
- Coordinate multiple domain services
- Manage cross-cutting concerns

**Example from current codebase**:
```go
// internal/application/command/user/create_user.go
type CreateUserHandler struct {
    userService *user.Service  // Domain service
    logger      *slog.Logger
}

func (h *CreateUserHandler) Handle(ctx context.Context, input dto.CreateUserInput) (*dto.UserDTO, error) {
    // Application layer responsibilities:
    // 1. Input validation (DTO level)
    // 2. Call domain service
    // 3. Convert result to DTO
    // 4. Handle application-level errors
    
    domainUser, err := h.userService.CreateUser(
        ctx,
        input.FirebaseUID,
        input.Email,
        input.Name,
        input.ProfileImageURL,
    )
    if err != nil {
        h.logger.Error("Failed to create user", "error", err)
        return nil, fmt.Errorf("failed to create user: %w", err)
    }

    // Convert domain object to DTO
    result := dto.ToDTO(domainUser)
    return result, nil
}
```

### 3. GraphQL Resolver → Application Layer → Domain Layer Flow

**Recommended Flow**:
1. **GraphQL Resolver** (Presentation Layer): `internal/graph/resolver/user.go`
2. **Application Service** (Use Case): `internal/application/command/user/create_user.go`
3. **Domain Service** (Business Logic): `internal/domain/user/service.go`
4. **Repository** (Data Access): `internal/adapter/repository/postgres/user.go`

**Example Flow**:
```go
// 1. GraphQL Resolver
func (r *mutationResolver) CreateUser(ctx context.Context, input graph.CreateUserInput) (*graph.User, error) {
    // Convert GraphQL input to DTO
    dto := dto.CreateUserInput{
        FirebaseUID:      input.FirebaseUID,
        Email:           input.Email,
        Name:            input.Name,
        ProfileImageURL: input.ProfileImageURL,
    }
    
    // Call application service
    user, err := r.createUserHandler.Handle(ctx, dto)
    if err != nil {
        return nil, err
    }
    
    // Convert DTO to GraphQL response
    return toGraphQLUser(user), nil
}

// 2. Application Service
func (h *CreateUserHandler) Handle(ctx context.Context, input dto.CreateUserInput) (*dto.UserDTO, error) {
    // Orchestrate the use case
    domainUser, err := h.userService.CreateUser(ctx, input.FirebaseUID, input.Email, input.Name, input.ProfileImageURL)
    if err != nil {
        return nil, err
    }
    return dto.ToDTO(domainUser), nil
}

// 3. Domain Service
func (s *Service) CreateUser(ctx context.Context, firebaseUID, email, name string, profileImageURL *string) (*User, error) {
    // Business logic and domain rules
    // ... (as shown in previous example)
}
```

## Best Practices for Service Interfaces and Implementations

### 1. Domain Service Patterns

```go
// Domain service interface (if needed for testing)
type UserDomainService interface {
    CreateUser(ctx context.Context, firebaseUID, email, name string, profileImageURL *string) (*User, error)
    ValidateUserUniqueness(ctx context.Context, email string) error
    CalculateUserStats(ctx context.Context, userID uuid.UUID) (*UserStats, error)
}

// Implementation
type Service struct {
    repo   Repository
    logger *slog.Logger
}

func NewService(repo Repository, logger *slog.Logger) *Service {
    return &Service{
        repo:   repo,
        logger: logger.With(slog.String("service", "user")),
    }
}
```

### 2. Application Service Patterns

```go
// Application service interface
type CreateUserUseCase interface {
    Handle(ctx context.Context, input dto.CreateUserInput) (*dto.UserDTO, error)
}

// Command/Query separation
type UserCommandHandlers struct {
    CreateUser *CreateUserHandler
    UpdateUser *UpdateUserHandler
    DeleteUser *DeleteUserHandler
}

type UserQueryHandlers struct {
    GetUser     *GetUserHandler
    ListUsers   *ListUsersHandler
    SearchUsers *SearchUsersHandler
}
```

### 3. Dependency Injection with Wire

```go
// Domain layer wire.go
var DomainSet = wire.NewSet(
    NewService,
    wire.Bind(new(UserDomainService), new(*Service)),
)

// Application layer wire.go
var ApplicationSet = wire.NewSet(
    NewCreateUserHandler,
    NewUpdateUserHandler,
    wire.Bind(new(CreateUserUseCase), new(*CreateUserHandler)),
)
```

## Common Anti-Patterns to Avoid

### 1. **Anemic Domain Services**
❌ **Wrong**: Domain service that only forwards to repository
```go
func (s *Service) CreateUser(ctx context.Context, user *User) error {
    return s.repo.Create(ctx, user)  // No business logic!
}
```

✅ **Correct**: Domain service with business logic
```go
func (s *Service) CreateUser(ctx context.Context, firebaseUID, email, name string) (*User, error) {
    // Validate business rules
    if err := s.validateEmailUniqueness(ctx, email); err != nil {
        return nil, err
    }
    
    // Apply business logic
    user, err := NewUser(firebaseUID, email, name, nil)
    if err != nil {
        return nil, err
    }
    
    // Persist
    return user, s.repo.Create(ctx, user)
}
```

### 2. **Fat Application Services**
❌ **Wrong**: Application service with business logic
```go
func (h *CreateUserHandler) Handle(ctx context.Context, input dto.CreateUserInput) (*dto.UserDTO, error) {
    // Business logic in application layer - WRONG!
    if input.Email == "" || !strings.Contains(input.Email, "@") {
        return nil, errors.New("invalid email")
    }
    
    // More business logic...
    if len(input.Name) < 2 {
        return nil, errors.New("name too short")
    }
    
    // Direct repository access - WRONG!
    exists, _ := h.userRepo.GetByEmail(ctx, input.Email)
    if exists != nil {
        return nil, errors.New("user exists")
    }
}
```

✅ **Correct**: Application service as orchestrator
```go
func (h *CreateUserHandler) Handle(ctx context.Context, input dto.CreateUserInput) (*dto.UserDTO, error) {
    // Delegate business logic to domain service
    domainUser, err := h.userService.CreateUser(ctx, input.FirebaseUID, input.Email, input.Name, input.ProfileImageURL)
    if err != nil {
        return nil, err
    }
    
    // Convert to DTO
    return dto.ToDTO(domainUser), nil
}
```

### 3. **Direct Repository Access from Application Layer**
❌ **Wrong**: Application service calling repository directly
```go
func (h *CreateUserHandler) Handle(ctx context.Context, input dto.CreateUserInput) (*dto.UserDTO, error) {
    user := &domain.User{...}
    err := h.userRepo.Create(ctx, user)  // Bypasses domain logic!
    return dto.ToDTO(user), err
}
```

✅ **Correct**: Application service calls domain service
```go
func (h *CreateUserHandler) Handle(ctx context.Context, input dto.CreateUserInput) (*dto.UserDTO, error) {
    user, err := h.userService.CreateUser(ctx, input.FirebaseUID, input.Email, input.Name, input.ProfileImageURL)
    return dto.ToDTO(user), err
}
```

### 4. **God Services**
❌ **Wrong**: Single service handling multiple aggregates
```go
type FinanceService struct {
    userRepo        UserRepository
    accountRepo     AccountRepository
    transactionRepo TransactionRepository
    receiptRepo     ReceiptRepository
}

func (s *FinanceService) ProcessEverything(ctx context.Context) error {
    // Handles users, accounts, transactions, receipts - TOO MUCH!
}
```

✅ **Correct**: Separate services per aggregate
```go
type UserService struct { repo UserRepository }
type AccountService struct { repo AccountRepository }  
type TransactionService struct { repo TransactionRepository }
type ReceiptService struct { repo ReceiptRepository }
```

## Language-Specific Examples

### Go Implementation (Current Project)
```go
// Domain Service
type UserService struct {
    repo   UserRepository
    logger *slog.Logger
}

func (s *UserService) CreateUser(ctx context.Context, firebaseUID, email, name string) (*User, error) {
    // Business logic
    exists, err := s.repo.Exists(ctx, firebaseUID)
    if err != nil {
        return nil, err
    }
    if exists {
        return nil, ErrUserAlreadyExists
    }
    
    user, err := NewUser(firebaseUID, email, name, nil)
    if err != nil {
        return nil, err
    }
    
    return user, s.repo.Create(ctx, user)
}

// Application Service
type CreateUserHandler struct {
    userService *UserService
    logger      *slog.Logger
}

func (h *CreateUserHandler) Handle(ctx context.Context, input CreateUserInput) (*UserDTO, error) {
    user, err := h.userService.CreateUser(ctx, input.FirebaseUID, input.Email, input.Name)
    if err != nil {
        return nil, err
    }
    return ToUserDTO(user), nil
}
```

### Java Implementation
```java
// Domain Service
@Service
public class UserDomainService {
    private final UserRepository userRepository;
    
    public User createUser(String firebaseUID, String email, String name) {
        if (userRepository.existsByFirebaseUID(firebaseUID)) {
            throw new UserAlreadyExistsException();
        }
        
        User user = new User(firebaseUID, email, name);
        return userRepository.save(user);
    }
}

// Application Service
@Service
public class CreateUserUseCase {
    private final UserDomainService userDomainService;
    
    @Transactional
    public UserDTO handle(CreateUserCommand command) {
        User user = userDomainService.createUser(
            command.getFirebaseUID(),
            command.getEmail(), 
            command.getName()
        );
        return UserDTO.fromDomain(user);
    }
}
```

### C# Implementation
```csharp
// Domain Service
public class UserDomainService
{
    private readonly IUserRepository _userRepository;
    
    public async Task<User> CreateUserAsync(string firebaseUID, string email, string name)
    {
        var exists = await _userRepository.ExistsAsync(firebaseUID);
        if (exists)
        {
            throw new UserAlreadyExistsException();
        }
        
        var user = new User(firebaseUID, email, name);
        await _userRepository.CreateAsync(user);
        return user;
    }
}

// Application Service
public class CreateUserUseCase
{
    private readonly UserDomainService _userDomainService;
    
    public async Task<UserDto> HandleAsync(CreateUserCommand command)
    {
        var user = await _userDomainService.CreateUserAsync(
            command.FirebaseUID,
            command.Email,
            command.Name
        );
        
        return UserDto.FromDomain(user);
    }
}
```

## Service Layer Organization Summary

### Clean Architecture Layers (Bottom to Top)

1. **Domain Layer** (`internal/domain/`)
   - **Entities**: Core business objects
   - **Value Objects**: Immutable domain concepts  
   - **Domain Services**: Business logic spanning multiple entities
   - **Repository Interfaces**: Data access contracts
   - **Domain Events**: Business event definitions

2. **Application Layer** (`internal/application/`)
   - **Command Handlers**: Write operations (CQRS commands)
   - **Query Handlers**: Read operations (CQRS queries)
   - **DTOs**: Data transfer objects for application boundaries
   - **Use Cases**: Application-specific business logic
   - **Application Services**: Orchestrate domain services

3. **Infrastructure/Adapter Layer** (`internal/adapter/`)
   - **Repository Implementations**: Data access implementations
   - **External Service Clients**: Third-party integrations
   - **Message Publishers/Subscribers**: Event handling

4. **Presentation Layer** (`internal/graph/`)
   - **GraphQL Resolvers**: API endpoint handlers
   - **REST Controllers**: HTTP endpoint handlers (if applicable)
   - **Middleware**: Cross-cutting concerns (auth, logging, etc.)

### Key Principles

1. **Dependency Direction**: Always inward (Presentation → Application → Domain)
2. **Interface Segregation**: Small, focused service interfaces
3. **Single Responsibility**: Each service has one reason to change
4. **Dependency Inversion**: Depend on abstractions, not concretions
5. **Separation of Concerns**: Clear boundaries between layers

## Domain Services vs Application Services: Decision Matrix

| Concern | Domain Service | Application Service |
|---------|----------------|-------------------|
| **Business Logic** | ✅ Complex domain rules | ❌ No business logic |
| **Entity Validation** | ✅ Domain entity validation | ❌ Only DTO validation |
| **Cross-Aggregate Operations** | ✅ When domain-relevant | ✅ Orchestration only |
| **Repository Access** | ✅ Direct access allowed | ❌ Via domain services only |
| **Transaction Management** | ❌ Infrastructure concern | ✅ Application boundary |
| **DTO Conversion** | ❌ No DTOs in domain | ✅ Primary responsibility |
| **Logging/Monitoring** | ✅ Domain events/errors | ✅ Application events |
| **External Service Calls** | ❌ Via interfaces only | ✅ Via interfaces/adapters |

## Service Placement Guidelines

### When to Use Domain Services
- Business logic spans multiple aggregates
- Complex domain calculations or rules
- Domain invariant enforcement
- Rich domain behavior that doesn't fit in entities
- Domain-specific algorithms or processes

### When to Use Application Services
- Orchestrating multiple domain operations
- Managing application workflows
- Handling cross-cutting concerns (transactions, logging)
- Converting between domain objects and DTOs
- Coordinating external service calls
- Implementing use cases

### Current Project Analysis

Looking at the current Djinn project structure:

**Domain Service** (`internal/domain/user/service.go`):
- ✅ Contains business logic (user existence check)
- ✅ Enforces domain rules (user validation)
- ✅ Direct repository access for domain operations
- ✅ Proper error handling with domain errors

**Application Service** (`internal/application/command/user/create_user.go`):
- ✅ Orchestrates use case (create user workflow)
- ✅ Handles DTO conversion
- ✅ Delegates business logic to domain service
- ✅ Manages application concerns (logging)

**GraphQL Integration**:
- ✅ Proper layer separation maintained
- ✅ GraphQL resolver → Application → Domain → Repository flow
- ✅ Clear responsibility boundaries

## Advanced Patterns

### 1. Multi-Aggregate Domain Services

```go
// Domain service coordinating multiple aggregates
type TransactionCategorizationService struct {
    transactionRepo TransactionRepository
    categoryRepo    CategoryRepository
    ruleRepo        CategorizationRuleRepository
    logger         *slog.Logger
}

func (s *TransactionCategorizationService) CategorizeTransaction(
    ctx context.Context, 
    transaction *Transaction,
) (*Category, error) {
    // Complex business logic involving multiple aggregates
    rules, err := s.ruleRepo.GetActiveRules(ctx)
    if err != nil {
        return nil, err
    }
    
    for _, rule := range rules {
        if rule.Matches(transaction) {
            category, err := s.categoryRepo.GetByID(ctx, rule.CategoryID)
            if err != nil {
                return nil, err
            }
            
            // Apply domain logic
            transaction.SetCategory(category)
            return category, nil
        }
    }
    
    // Default categorization logic
    return s.getDefaultCategory(ctx, transaction)
}
```

### 2. Event-Driven Domain Services

```go
// Domain service publishing domain events
type UserService struct {
    repo      UserRepository
    publisher DomainEventPublisher
    logger    *slog.Logger
}

func (s *UserService) CreateUser(ctx context.Context, firebaseUID, email, name string) (*User, error) {
    user, err := NewUser(firebaseUID, email, name, nil)
    if err != nil {
        return nil, err
    }
    
    if err := s.repo.Create(ctx, user); err != nil {
        return nil, err
    }
    
    // Publish domain event
    event := &UserCreatedEvent{
        UserID:      user.ID,
        FirebaseUID: user.FirebaseUID,
        Email:       user.Email,
        OccurredAt:  time.Now(),
    }
    
    if err := s.publisher.Publish(ctx, event); err != nil {
        s.logger.Error("Failed to publish user created event", "error", err)
        // Decision: Continue or rollback depends on business requirements
    }
    
    return user, nil
}
```

### 3. Specification Pattern in Domain Services

```go
// Domain service using specification pattern
type UserSpecification interface {
    IsSatisfiedBy(user *User) bool
}

type ActiveUserSpecification struct{}

func (s *ActiveUserSpecification) IsSatisfiedBy(user *User) bool {
    return user.Status == UserStatusActive && !user.IsDeleted()
}

type UserService struct {
    repo UserRepository
    spec UserSpecification
}

func (s *UserService) GetActiveUsers(ctx context.Context) ([]*User, error) {
    allUsers, err := s.repo.GetAll(ctx)
    if err != nil {
        return nil, err
    }
    
    var activeUsers []*User
    for _, user := range allUsers {
        if s.spec.IsSatisfiedBy(user) {
            activeUsers = append(activeUsers, user)
        }
    }
    
    return activeUsers, nil
}
```

## Testing Strategies

### Domain Service Testing

```go
func TestUserService_CreateUser(t *testing.T) {
    // Arrange
    mockRepo := &MockUserRepository{}
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
    service := NewUserService(mockRepo, logger)
    
    ctx := context.Background()
    firebaseUID := "firebase123"
    email := "test@example.com"
    name := "Test User"
    
    mockRepo.On("Exists", ctx, firebaseUID).Return(false, nil)
    mockRepo.On("Create", ctx, mock.AnythingOfType("*User")).Return(nil)
    
    // Act
    user, err := service.CreateUser(ctx, firebaseUID, email, name, nil)
    
    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, user)
    assert.Equal(t, firebaseUID, user.FirebaseUID)
    assert.Equal(t, email, user.Email)
    assert.Equal(t, name, user.Name)
    
    mockRepo.AssertExpectations(t)
}

func TestUserService_CreateUser_AlreadyExists(t *testing.T) {
    // Test business logic: user already exists
    mockRepo := &MockUserRepository{}
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
    service := NewUserService(mockRepo, logger)
    
    ctx := context.Background()
    firebaseUID := "firebase123"
    
    mockRepo.On("Exists", ctx, firebaseUID).Return(true, nil)
    
    // Act
    user, err := service.CreateUser(ctx, firebaseUID, "test@example.com", "Test", nil)
    
    // Assert
    assert.Error(t, err)
    assert.Nil(t, user)
    assert.ErrorIs(t, err, ErrUserAlreadyExists)
    
    // Verify Create was never called
    mockRepo.AssertNotCalled(t, "Create")
}
```

### Application Service Testing

```go
func TestCreateUserHandler_Handle(t *testing.T) {
    // Arrange
    mockUserService := &MockUserDomainService{}
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
    handler := NewCreateUserHandler(mockUserService, logger)
    
    ctx := context.Background()
    input := dto.CreateUserInput{
        FirebaseUID: "firebase123",
        Email:       "test@example.com",
        Name:        "Test User",
    }
    
    expectedUser := &User{
        ID:          uuid.New(),
        FirebaseUID: input.FirebaseUID,
        Email:       input.Email,
        Name:        input.Name,
    }
    
    mockUserService.On("CreateUser", ctx, input.FirebaseUID, input.Email, input.Name, (*string)(nil)).
        Return(expectedUser, nil)
    
    // Act
    result, err := handler.Handle(ctx, input)
    
    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, expectedUser.ID, result.ID)
    assert.Equal(t, expectedUser.FirebaseUID, result.FirebaseUID)
    assert.Equal(t, expectedUser.Email, result.Email)
    assert.Equal(t, expectedUser.Name, result.Name)
    
    mockUserService.AssertExpectations(t)
}
```

## Conclusion

Domain services belong in the domain layer and contain business logic that doesn't fit within individual aggregates. Application services belong in the application layer and orchestrate use cases by coordinating domain services. The current Djinn project structure correctly implements these patterns with proper separation of concerns.

### Key Takeaways

1. **Domain Services**: Business logic, domain rules, aggregate coordination
2. **Application Services**: Use case orchestration, DTO conversion, cross-cutting concerns
3. **Layer Flow**: GraphQL → Application → Domain → Repository
4. **Avoid Anti-patterns**: Anemic domain services, fat application services, direct repository access
5. **Testing**: Focus on business logic in domain service tests, orchestration in application service tests

The architecture ensures maintainability, testability, and clear separation of concerns while following DDD and Clean Architecture principles.