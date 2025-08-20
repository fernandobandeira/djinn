# ADR-20250820: Testing Strategy and Standards

## Status
Proposed

## Context

The Djinn personal finance application requires a comprehensive testing strategy that covers both the Go backend API and Flutter mobile application. Based on our existing technology decisions:

- **Backend**: Go with GraphQL (gqlgen), PostgreSQL, Firebase Auth
- **Mobile**: Flutter with Drift database for offline-first functionality
- **Financial Requirements**: High accuracy for monetary calculations, regulatory compliance
- **Architecture**: Offline-first with eventual consistency, complex synchronization logic

Key forces at play:
- Need for high reliability in financial calculations (zero tolerance for monetary errors)
- Complex offline synchronization requiring thorough integration testing
- Small team requiring efficient testing practices with good ROI
- MVP phase requiring balance between speed and quality
- Regulatory requirements for financial data handling

Current state from existing ADRs:
- ADR-20250812 mentions: "standard library + go-cmp + rapid (property testing)" for Go
- Security ADR requires: penetration testing, security scanning in CI/CD
- Multiple ADRs mention testing needs but lack comprehensive strategy

## Decision

We will adopt a **pragmatic testing trophy strategy** with the following standards:

### Overall Testing Philosophy
- **Test Trophy Model**: Emphasize integration tests (most bang for buck)
- **TDD for Critical Paths**: Apply TDD selectively for business logic and financial calculations
- **Shift-Left Security**: Security testing integrated from the start
- **Automated Everything**: All tests must run in CI/CD

### Go Backend Testing Stack

#### Testing Frameworks
```yaml
unit_testing:
  framework: testify/suite + assert + require
  rationale: "Superior assertions, better error messages than stdlib"
  
mocking:
  primary: testify/mock
  generator: mockery (v2)
  rationale: "Balance of simplicity and power, good gqlgen integration"
  
integration_testing:
  database: testcontainers-go
  api: httptest + gqlgen/client
  rationale: "Real PostgreSQL testing, accurate GraphQL testing"
  
property_testing:
  framework: rapid (as decided in ADR-20250812)
  use_cases: "Financial calculations, money handling, date logic"
```

#### Testing Patterns
```go
// Standard: AAA Pattern with Table-Driven Tests
func TestTransactionCalculation(t *testing.T) {
    tests := []struct {
        name     string
        // Arrange
        input    Transaction
        // Assert
        expected Money
        wantErr  bool
    }{
        {
            name: "positive transaction adds to balance",
            input: Transaction{Amount: 1000, Type: CREDIT},
            expected: Money{AmountMinor: 1000, Currency: "USD"},
        },
        // More test cases...
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Act
            result, err := CalculateBalance(tt.input)
            
            // Assert
            if tt.wantErr {
                require.Error(t, err)
                return
            }
            require.NoError(t, err)
            assert.Equal(t, tt.expected, result)
        })
    }
}
```

### Flutter Mobile Testing Stack

#### Testing Frameworks
```yaml
unit_testing:
  framework: flutter_test (built-in)
  patterns: "AAA with group/test structure"
  
widget_testing:
  framework: flutter_test
  golden_testing: golden_toolkit
  rationale: "Visual regression testing for UI consistency"
  
mocking:
  framework: mocktail
  rationale: "Null-safe, no code generation, cleaner syntax"
  
integration_testing:
  framework: patrol
  rationale: "6x faster than flutter_driver, native UI interaction"
  
state_testing:
  bloc: bloc_test
  drift: drift_test_utils
  rationale: "Specialized testing for our state management"
```

#### Testing Patterns
```dart
// Widget Testing with Golden Files
testWidgets('Transaction list displays correctly', (tester) async {
  // Arrange
  final transactions = [
    Transaction(amount: Money(1000, 'USD'), merchant: 'Coffee Shop'),
  ];
  
  // Act
  await tester.pumpWidget(
    MaterialApp(
      home: TransactionList(transactions: transactions),
    ),
  );
  
  // Assert
  await expectLater(
    find.byType(TransactionList),
    matchesGoldenFile('goldens/transaction_list.png'),
  );
});

// Offline Database Testing
test('sync queue handles conflicts correctly', () async {
  // Arrange
  final db = DriftTestDatabase();
  final queue = SyncQueue(db);
  
  // Act
  await queue.add(localTransaction);
  await queue.add(conflictingTransaction);
  
  // Assert
  final resolved = await queue.resolveConflicts();
  expect(resolved.length, equals(1));
  expect(resolved.first.timestamp, 
    greaterThan(localTransaction.timestamp));
});
```

### Firebase Testing Strategy

Since most API endpoints are protected by Firebase Auth, we need comprehensive testing for authentication and authorization flows.

#### Firebase Local Emulator Suite Setup
```yaml
firebase_emulators:
  auth: 
    port: 9099
    enabled: true
  firestore:
    port: 8080
    enabled: false  # Not using Firestore, only Auth
  hosting:
    enabled: false
    
environment_setup:
  local:
    emulator: always
    seed_users: test fixtures
  ci:
    emulator: always
    parallel_safe: true
  staging:
    emulator: never
    use_test_project: true
```

#### Go Backend - Firebase Auth Testing

```go
// Test helpers for Firebase Auth
package auth_test

import (
    "context"
    "os"
    "testing"
    firebase "firebase.google.com/go/v4"
    "firebase.google.com/go/v4/auth"
    "github.com/stretchr/testify/suite"
)

type AuthTestSuite struct {
    suite.Suite
    app    *firebase.App
    auth   *auth.Client
    testUsers map[string]*auth.UserRecord
}

func (s *AuthTestSuite) SetupSuite() {
    // Use emulator in tests
    os.Setenv("FIREBASE_AUTH_EMULATOR_HOST", "localhost:9099")
    
    // Initialize Firebase app for testing
    app, err := firebase.NewApp(context.Background(), nil)
    s.Require().NoError(err)
    s.app = app
    
    authClient, err := app.Auth(context.Background())
    s.Require().NoError(err)
    s.auth = authClient
    
    // Create test users
    s.seedTestUsers()
}

func (s *AuthTestSuite) seedTestUsers() {
    s.testUsers = make(map[string]*auth.UserRecord)
    
    // Create standard test user
    user, err := s.auth.CreateUser(context.Background(), &auth.UserToCreate{}.
        Email("test@example.com").
        Password("testpass123").
        EmailVerified(true))
    s.Require().NoError(err)
    s.testUsers["standard"] = user
    
    // Create admin test user with custom claims
    adminUser, err := s.auth.CreateUser(context.Background(), &auth.UserToCreate{}.
        Email("admin@example.com").
        Password("adminpass123"))
    s.Require().NoError(err)
    
    // Set admin claims
    claims := map[string]interface{}{
        "admin": true,
        "role": "admin",
    }
    err = s.auth.SetCustomUserClaims(context.Background(), adminUser.UID, claims)
    s.Require().NoError(err)
    s.testUsers["admin"] = adminUser
}

// Mock authenticated GraphQL context
func (s *AuthTestSuite) TestProtectedEndpoint() {
    // Create custom token for testing
    token, err := s.auth.CustomToken(context.Background(), s.testUsers["standard"].UID)
    s.Require().NoError(err)
    
    // Test GraphQL resolver with auth context
    ctx := context.WithValue(context.Background(), "auth_token", token)
    
    // Your GraphQL test here
    result, err := resolver.GetUserTransactions(ctx)
    s.Require().NoError(err)
    s.Require().NotEmpty(result)
}

// Test authorization rules
func (s *AuthTestSuite) TestAuthorizationRules() {
    tests := []struct {
        name      string
        user      *auth.UserRecord
        endpoint  string
        shouldAllow bool
    }{
        {
            name:      "admin can access admin endpoints",
            user:      s.testUsers["admin"],
            endpoint:  "/admin/users",
            shouldAllow: true,
        },
        {
            name:      "standard user cannot access admin endpoints",
            user:      s.testUsers["standard"],
            endpoint:  "/admin/users",
            shouldAllow: false,
        },
        {
            name:      "authenticated user can access own data",
            user:      s.testUsers["standard"],
            endpoint:  "/users/self",
            shouldAllow: true,
        },
    }
    
    for _, tt := range tests {
        s.Run(tt.name, func() {
            // Test authorization logic
            allowed := checkAuthorization(tt.user, tt.endpoint)
            s.Equal(tt.shouldAllow, allowed)
        })
    }
}
```

#### Flutter - Firebase Auth Mocking

```dart
// Firebase Auth test configuration
import 'package:firebase_auth/firebase_auth.dart';
import 'package:firebase_auth_mocks/firebase_auth_mocks.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';

// Use firebase_auth_mocks package for better testing
class AuthTestHelper {
  static MockFirebaseAuth createMockAuth({
    MockUser? mockUser,
    bool isSignedIn = false,
  }) {
    final user = mockUser ?? 
      MockUser(
        uid: 'test-uid-123',
        email: 'test@example.com',
        displayName: 'Test User',
        isEmailVerified: true,
      );
    
    return MockFirebaseAuth(
      signedIn: isSignedIn,
      mockUser: user,
    );
  }
}

// Repository testing with mocked Firebase Auth
class AuthRepositoryTest {
  group('AuthRepository', () {
    late MockFirebaseAuth mockAuth;
    late AuthRepository repository;
    
    setUp(() {
      mockAuth = AuthTestHelper.createMockAuth();
      repository = AuthRepository(firebaseAuth: mockAuth);
    });
    
    test('should sign in with Google OAuth', () async {
      // Arrange - Mock will handle OAuth flow
      
      // Act
      final user = await repository.signInWithGoogle();
      
      // Assert
      expect(user, isNotNull);
      expect(user?.email, equals('test@example.com'));
      expect(mockAuth.authStateChanges(), emits(isA<User>()));
    });
    
    test('should handle auth state changes', () async {
      // Arrange
      final mockAuth = AuthTestHelper.createMockAuth(isSignedIn: false);
      final repository = AuthRepository(firebaseAuth: mockAuth);
      
      // Act & Assert - Stream testing
      expectLater(
        repository.authStateChanges,
        emitsInOrder([
          null,  // Initially signed out
          isA<User>(),  // After sign in
          null,  // After sign out
        ]),
      );
      
      await mockAuth.signInWithEmailAndPassword(
        email: 'test@example.com',
        password: 'password',
      );
      await mockAuth.signOut();
    });
  });
}

// Widget testing with authenticated context
testWidgets('should show user data when authenticated', (tester) async {
  // Arrange
  final mockAuth = AuthTestHelper.createMockAuth(isSignedIn: true);
  
  await tester.pumpWidget(
    MaterialApp(
      home: Provider<FirebaseAuth>.value(
        value: mockAuth,
        child: const HomeScreen(),
      ),
    ),
  );
  
  // Act
  await tester.pumpAndSettle();
  
  // Assert
  expect(find.text('Welcome, Test User'), findsOneWidget);
  expect(find.text('Sign In'), findsNothing);
});
```

#### Integration Testing with Firebase Emulator

```yaml
# test/integration/firebase_emulator_test.dart
import 'package:integration_test/integration_test.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:firebase_auth/firebase_auth.dart';

void main() {
  IntegrationTestWidgetsFlutterBinding.ensureInitialized();
  
  setUpAll(() async {
    await Firebase.initializeApp();
    
    // Connect to emulator
    await FirebaseAuth.instance.useAuthEmulator('localhost', 9099);
    
    // Clear any existing auth state
    await FirebaseAuth.instance.signOut();
  });
  
  testWidgets('full auth flow with emulator', (tester) async {
    // Create test user in emulator
    final email = 'integration_${DateTime.now().millisecondsSinceEpoch}@test.com';
    
    // Test sign up
    final credential = await FirebaseAuth.instance
      .createUserWithEmailAndPassword(
        email: email,
        password: 'testpass123',
      );
    
    expect(credential.user, isNotNull);
    expect(credential.user?.email, equals(email));
    
    // Test sign out
    await FirebaseAuth.instance.signOut();
    expect(FirebaseAuth.instance.currentUser, isNull);
    
    // Test sign in
    final signInCredential = await FirebaseAuth.instance
      .signInWithEmailAndPassword(
        email: email,
        password: 'testpass123',
      );
    
    expect(signInCredential.user?.email, equals(email));
  });
}
```

#### CI/CD Firebase Testing

```yaml
# GitHub Actions with Firebase Emulator
name: Test with Firebase
on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Setup Node for Firebase
        uses: actions/setup-node@v3
        with:
          node-version: '18'
          
      - name: Install Firebase CLI
        run: npm install -g firebase-tools
        
      - name: Start Firebase Emulator
        run: |
          firebase emulators:start \
            --only auth \
            --project demo-test \
            --import ./test/fixtures/auth-export \
            --export-on-exit &
          sleep 5  # Wait for emulator to start
          
      - name: Run Go Tests
        env:
          FIREBASE_AUTH_EMULATOR_HOST: localhost:9099
        run: |
          go test -v -race ./... -tags=integration
          
      - name: Run Flutter Tests  
        run: |
          flutter test
          flutter test integration_test --dart-define=USE_FIREBASE_EMULATOR=true
```

#### Test Data Management

```go
// Test fixtures for consistent test data
package fixtures

type TestUser struct {
    UID      string
    Email    string
    Password string
    Claims   map[string]interface{}
}

var TestUsers = []TestUser{
    {
        UID:      "test-user-1",
        Email:    "user1@test.com",
        Password: "password123",
        Claims:   map[string]interface{}{"role": "user"},
    },
    {
        UID:      "test-admin-1",
        Email:    "admin@test.com", 
        Password: "admin123",
        Claims:   map[string]interface{}{"role": "admin", "admin": true},
    },
    {
        UID:      "test-premium-1",
        Email:    "premium@test.com",
        Password: "premium123",
        Claims:   map[string]interface{}{"role": "user", "subscription": "premium"},
    },
}

// Helper to seed test users
func SeedTestUsers(authClient *auth.Client) error {
    for _, user := range TestUsers {
        _, err := authClient.CreateUser(context.Background(), 
            (&auth.UserToCreate{}).
                UID(user.UID).
                Email(user.Email).
                Password(user.Password))
        if err != nil && !strings.Contains(err.Error(), "already exists") {
            return err
        }
        
        if user.Claims != nil {
            authClient.SetCustomUserClaims(context.Background(), user.UID, user.Claims)
        }
    }
    return nil
}
```

#### GraphQL Protected Endpoint Testing

```go
// Testing protected GraphQL resolvers with Firebase Auth
package graph_test

import (
    "context"
    "testing"
    "github.com/99designs/gqlgen/client"
    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestProtectedGraphQLEndpoints(t *testing.T) {
    // Setup test server with auth middleware
    srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
        Resolvers: &graph.Resolver{},
    }))
    
    // Create test client
    c := client.New(srv)
    
    t.Run("unauthenticated request should fail", func(t *testing.T) {
        var resp struct {
            GetTransactions []Transaction
        }
        
        err := c.Post(`
            query {
                getTransactions {
                    id
                    amount
                }
            }
        `, &resp)
        
        require.Error(t, err)
        assert.Contains(t, err.Error(), "unauthorized")
    })
    
    t.Run("authenticated request should succeed", func(t *testing.T) {
        // Create authenticated context
        ctx := context.WithValue(context.Background(), "user", &auth.UserInfo{
            UID:   "test-user-123",
            Email: "test@example.com",
        })
        
        var resp struct {
            GetTransactions []Transaction
        }
        
        // Use authenticated context
        err := c.Post(`
            query {
                getTransactions {
                    id
                    amount
                }
            }
        `, &resp, client.WithContext(ctx))
        
        require.NoError(t, err)
        assert.NotEmpty(t, resp.GetTransactions)
    })
    
    t.Run("user can only access own data", func(t *testing.T) {
        // Test data isolation between users
        user1Ctx := context.WithValue(context.Background(), "user", &auth.UserInfo{
            UID: "user-1",
        })
        
        user2Ctx := context.WithValue(context.Background(), "user", &auth.UserInfo{
            UID: "user-2",
        })
        
        var user1Resp, user2Resp struct {
            GetUserProfile Profile
        }
        
        // User 1 request
        err := c.Post(`
            query { 
                getUserProfile { id, email } 
            }
        `, &user1Resp, client.WithContext(user1Ctx))
        require.NoError(t, err)
        
        // User 2 request
        err = c.Post(`
            query { 
                getUserProfile { id, email }
            }
        `, &user2Resp, client.WithContext(user2Ctx))
        require.NoError(t, err)
        
        // Verify data isolation
        assert.NotEqual(t, user1Resp.GetUserProfile.ID, user2Resp.GetUserProfile.ID)
    })
}

// Middleware testing helper
func withAuthContext(uid, email string) client.Option {
    return func(r *http.Request) {
        ctx := context.WithValue(r.Context(), "auth", &auth.Claims{
            UID:   uid,
            Email: email,
        })
        *r = *r.WithContext(ctx)
    }
}

// Test with custom claims
func TestAdminEndpoints(t *testing.T) {
    c := client.New(handler.NewDefaultServer(schema))
    
    t.Run("admin can access admin endpoints", func(t *testing.T) {
        var resp struct {
            AdminGetAllUsers []User
        }
        
        err := c.Post(`
            query {
                adminGetAllUsers {
                    id
                    email
                }
            }
        `, &resp, withAuthContext("admin-1", "admin@example.com", map[string]interface{}{
            "admin": true,
        }))
        
        require.NoError(t, err)
        assert.NotEmpty(t, resp.AdminGetAllUsers)
    })
}
```

### Testing Standards and Patterns

#### 1. Test Naming Convention
```
// Go: TestFunctionName_StateUnderTest_ExpectedBehavior
TestCalculateBalance_NegativeAmount_ReturnsError

// Flutter: 'should [expected behavior] when [condition]'
test('should return error when amount is negative', () {});
```

#### 2. Test Coverage Requirements
```yaml
minimum_coverage:
  critical_paths: 90%  # Financial calculations, auth, sync
  business_logic: 80%  # Core features
  ui_components: 70%   # Flutter widgets
  utilities: 60%       # Helper functions
  
excluded_from_coverage:
  - generated_code    # gqlgen, mockery, freezed
  - migrations        # Database migrations
  - config_files      # Setup and configuration
```

#### 3. Test Pyramid Distribution
```yaml
go_backend:
  unit: 60%        # Fast, focused tests
  integration: 30%  # API and database tests
  e2e: 10%         # Full system tests
  
flutter_mobile:
  unit: 40%        # Business logic
  widget: 40%      # UI component tests
  integration: 20%  # Full app flows
```

#### 4. TDD Application Guidelines

**MUST use TDD for:**
- Financial calculations and money handling
- Transaction processing logic
- Synchronization conflict resolution
- Authentication and authorization
- Data validation and sanitization

**SHOULD use TDD for:**
- Complex business rules
- API endpoint handlers
- State management logic
- Critical user workflows

**MAY skip TDD for:**
- Simple CRUD operations
- UI layout and styling
- External service integrations
- Prototype features
- Documentation and configs

### Testing Environments

```yaml
environments:
  local:
    database: testcontainers (PostgreSQL 16)
    services: Docker Compose
    data: Fixtures and factories
    
  ci:
    database: testcontainers
    parallel: true (go test -parallel 4)
    cache: Go modules, Flutter packages
    
  staging:
    database: Separate PostgreSQL instance
    data: Anonymized production subset
    monitoring: Full observability
    
  production:
    testing: Synthetic monitoring only
    rollback: Blue-green deployment ready
```

### Security Testing Requirements

```yaml
static_analysis:
  go:
    - gosec      # Security linter
    - nancy      # Dependency vulnerabilities
    - trivy      # Container scanning
  
  flutter:
    - flutter analyze --fatal-warnings
    - dependency vulnerability scanning
    
dynamic_testing:
  api:
    - OWASP ZAP scan (weekly)
    - Rate limiting tests
    - Auth bypass attempts
    
  mobile:
    - Certificate pinning validation
    - Local storage encryption verification
    - Jailbreak/root detection testing
```

### Performance Testing Standards

```yaml
benchmarks:
  go:
    - Run on every PR for critical paths
    - Compare against baseline
    - Fail if >10% regression
    
  flutter:
    - Profile app startup time
    - Monitor frame rendering (60fps target)
    - Memory leak detection
    
load_testing:
  tool: k6
  scenarios:
    - 100 concurrent users (MVP)
    - 1000 concurrent users (growth)
    - Spike testing (2x normal load)
```

## Consequences

### Positive
- **High Confidence**: Comprehensive testing catches bugs early
- **Fast Feedback**: Emphasis on integration tests provides quick validation
- **Maintainable**: Clear patterns and standards reduce confusion
- **Cost Effective**: Test trophy approach optimizes testing ROI
- **Security First**: Integrated security testing prevents vulnerabilities
- **Documentation**: Tests serve as living documentation
- **Refactoring Safety**: Good coverage enables confident refactoring

### Negative
- **Initial Setup Time**: Testcontainers and patrol require configuration
- **Learning Curve**: Team needs to learn property-based testing
- **CI/CD Complexity**: Parallel testing and containers add complexity
- **Test Maintenance**: Golden files and snapshots require updates
- **Slower CI**: Integration tests take longer than unit tests

### Risks
- **Flaky Tests**: Integration tests may have timing issues
- **Golden File Drift**: UI changes require golden file regeneration
- **Container Resources**: Testcontainers need Docker availability
- **Mock Synchronization**: Mocks may drift from actual implementations

## Alternatives Considered

### Option A: Minimal Testing (Unit Tests Only)
- **Pros**: Fast execution, simple setup, low maintenance
- **Cons**: Misses integration issues, poor confidence, more production bugs
- **Reason for not choosing**: Financial application requires high reliability

### Option B: Full Test Pyramid (70% Unit, 20% Integration, 10% E2E)
- **Pros**: Traditional approach, well-understood, good isolation
- **Cons**: Many mocks, false confidence, misses integration issues
- **Reason for not choosing**: Modern tools make integration testing more valuable

### Option C: BDD with Cucumber/Gherkin
- **Pros**: Business-readable tests, living documentation, stakeholder involvement
- **Cons**: Extra abstraction layer, slower execution, maintenance overhead
- **Reason for not choosing**: Small team doesn't need the communication overhead

### Option D: Testing in Production Only
- **Pros**: Real-world testing, no test environment needed, fast development
- **Cons**: Customer impact, debugging difficulty, regulatory issues
- **Reason for not choosing**: Unacceptable risk for financial application

## Implementation Notes

### Phase 1: Foundation (Weeks 1-2)
1. Set up testcontainers for Go backend
2. Configure mockery for mock generation
3. Set up patrol for Flutter integration tests
4. Establish golden file workflow
5. Create test data factories

### Phase 2: Critical Path Coverage (Weeks 3-4)
1. Add property tests for money calculations
2. Write integration tests for sync logic
3. Create golden tests for key UI screens
4. Add security scanning to CI/CD

### Phase 3: Comprehensive Coverage (Weeks 5-6)
1. Reach 80% coverage on business logic
2. Add load testing scenarios
3. Implement mutation testing trial
4. Create test documentation

### Migration Strategy
```bash
# Go Backend Migration
1. Keep existing tests (stdlib + go-cmp)
2. New tests use testify
3. Gradually migrate critical tests
4. Add testcontainers for new integration tests

# Flutter Migration  
1. New widgets use golden tests
2. Replace mockito with mocktail in new tests
3. Add patrol for new integration tests
4. Keep existing working tests
```

### CI/CD Integration
```yaml
# GitHub Actions Workflow
name: Test Suite
on: [push, pull_request]

jobs:
  go-tests:
    runs-on: ubuntu-latest
    services:
      postgres:
        image: postgres:16
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
      - run: go install github.com/vektra/mockery/v2
      - run: go generate ./...
      - run: go test -race -coverprofile=coverage.out ./...
      - run: go tool cover -html=coverage.out -o coverage.html
      
  flutter-tests:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: subosito/flutter-action@v2
      - run: flutter pub get
      - run: flutter analyze --fatal-warnings
      - run: flutter test --coverage
      - run: flutter test --update-goldens
```

### Success Metrics
- **Test Execution Time**: <10 minutes for full suite
- **Coverage**: Meet minimum thresholds per component
- **Flakiness**: <1% flaky test rate
- **Bug Escape Rate**: <5% bugs found in production
- **MTTR**: <1 hour for test-caught issues

## References

### Related ADRs
- ADR-20250812: Personal Finance Tech Stack (mentions testing tools)
- ADR-20250819: Security Architecture (security testing requirements)
- ADR-20250819: Mobile Offline-First Synchronization (sync testing needs)
- ADR-20250819: Data Architecture (data integrity testing)

### External Documentation
- [Go Testing Best Practices](https://github.com/golang/go/wiki/TestComments)
- [Flutter Testing Documentation](https://docs.flutter.dev/testing)
- [Testcontainers Go](https://golang.testcontainers.org/)
- [Patrol Testing Framework](https://patrol.leancode.co/)
- [Test Trophy Strategy](https://kentcdodds.com/blog/the-testing-trophy-and-testing-classifications)
- [Property-Based Testing in Go](https://github.com/flyingmutant/rapid)

### Research Documents
- `/docs/architecture/research/go-testing-frameworks-comparison-2025.md`
- `/docs/architecture/research/flutter-testing-frameworks-comparison-2025.md`
- `/docs/architecture/patterns/testing-patterns-comprehensive-comparison.md`

## Decision Makers
- Author: Archie (System Architect)
- Reviewers: [Development Team]
- Approver: [CTO/Technical Lead]
- Date: 2025-08-20

## Revision History
- 2025-08-20: Initial draft based on comprehensive testing research
- 2025-08-20: Incorporated existing ADR constraints and decisions