# User Story DJINN-1003: Firebase Authentication Integration

## Metadata
- **Story ID**: DJINN-1003
- **Epic**: Epic 1 - Base Architecture & Authentication Foundation
- **Title**: Firebase Authentication Integration
- **Author**: Story Creator Agent
- **Created**: 2025-08-25
- **Last Updated**: 2025-08-25
- **Status**: Draft
- **Priority**: High
- **Effort Estimate**: 8 Story Points
- **Sprint**: TBD
- **Assigned To**: TBD
- **Labels**: authentication, firebase, oauth, security, foundation

## User Story
**As a** mobile app user
**I want** to securely authenticate using my Google or Apple account through Firebase
**So that** I can access the application with a trusted identity provider and have my session managed securely

## Business Context and Value
This story implements the core authentication foundation for the Djinn application, enabling secure OAuth-based login through Firebase Authentication. This provides users with a familiar, secure sign-in experience while establishing the authentication architecture that will support all future user-centric features.

**Business Value:**
- Reduces user friction with OAuth sign-in (no password creation/management)
- Leverages trusted identity providers (Google/Apple)
- Establishes secure session management foundation
- Enables user-specific features and data storage

## Acceptance Criteria

### Measurable
1. **Firebase Project Setup**
   - Firebase project configured with Authentication enabled
   - Google Sign-In provider configured and working
   - Apple Sign-In provider configured and working
   - Firebase SDK integrated in Flutter app
   - Authentication state persistence enabled

2. **Frontend Authentication Flow**
   - Sign-in screen with Google/Apple Sign-In buttons
   - Successful authentication redirects to main app screen
   - Authentication state managed globally in Flutter app
   - Sign-out functionality working correctly
   - Loading states displayed during auth operations

3. **Backend Token Validation**
   - Firebase Admin SDK integrated in Golang API
   - JWT token validation middleware implemented
   - Protected endpoints require valid Firebase tokens
   - User creation/retrieval flow implemented
   - PostgreSQL user records synchronized with Firebase users

4. **Security Requirements**
   - All API endpoints requiring authentication are protected
   - Firebase tokens properly validated on backend
   - Secure token storage on mobile device
   - Proper error handling for authentication failures

### Validation
1. **Authentication Flow Testing**
   - User can sign in with Google account
   - User can sign in with Apple account (iOS)
   - Authentication state persists across app restarts
   - Invalid/expired tokens are properly handled
   - Sign-out clears authentication state

2. **API Integration Testing**
   - Protected endpoints reject unauthenticated requests
   - Valid Firebase tokens allow API access
   - User records created/updated in PostgreSQL
   - Backend properly extracts user info from Firebase tokens

3. **Error Handling Testing**
   - Network errors during authentication handled gracefully
   - Cancelled sign-in flows handled properly
   - Invalid token responses from API handled correctly
   - User feedback provided for all error scenarios

## Tasks and Subtasks

### Task 1: Firebase Project Configuration
- [ ] **1.1**: Create Firebase project for Djinn application
  - [ ] Set up project in Firebase Console
  - [ ] Configure project settings and billing
  - [ ] Enable Authentication service
  - **Effort**: 1 hour
  - **Dependencies**: None

- [ ] **1.2**: Configure OAuth providers
  - [ ] Set up Google Sign-In provider in Firebase
  - [ ] Configure Apple Sign-In provider in Firebase
  - [ ] Generate and configure OAuth client IDs
  - **Effort**: 2 hours
  - **Dependencies**: Task 1.1

- [ ] **1.3**: Download and configure Firebase configuration files
  - [ ] Download `google-services.json` for Android
  - [ ] Download `GoogleService-Info.plist` for iOS
  - [ ] Add configuration files to Flutter project
  - **Effort**: 1 hour
  - **Dependencies**: Task 1.2

### Task 2: Flutter Firebase Integration
- [ ] **2.1**: Add Firebase dependencies to Flutter project
  - [ ] Add `firebase_core`, `firebase_auth`, `google_sign_in` packages
  - [ ] Update `pubspec.yaml` with required dependencies
  - [ ] Configure platform-specific settings
  - **Effort**: 1 hour
  - **Dependencies**: Task 1.3

- [ ] **2.2**: Implement Firebase initialization
  - [ ] Initialize Firebase in `main.dart`
  - [ ] Handle Firebase initialization errors
  - [ ] Set up authentication state listener
  - **Effort**: 2 hours
  - **Dependencies**: Task 2.1

- [ ] **2.3**: Create authentication service
  - [ ] Implement `AuthService` class with sign-in/sign-out methods
  - [ ] Handle Google Sign-In flow
  - [ ] Handle Apple Sign-In flow (iOS)
  - [ ] Manage authentication state
  - **Effort**: 4 hours
  - **Dependencies**: Task 2.2

### Task 3: Flutter UI Implementation
- [ ] **3.1**: Create sign-in screen UI
  - [ ] Design sign-in screen with Google/Apple Sign-In buttons
  - [ ] Implement loading states and error handling
  - [ ] Follow Material Design guidelines
  - **Effort**: 3 hours
  - **Dependencies**: Task 2.3

- [ ] **3.2**: Implement authentication flow navigation
  - [ ] Set up conditional routing based on auth state
  - [ ] Create authenticated and unauthenticated app flows
  - [ ] Handle deep linking with authentication
  - **Effort**: 2 hours
  - **Dependencies**: Task 3.1

- [ ] **3.3**: Add sign-out functionality
  - [ ] Implement sign-out button in app drawer/profile
  - [ ] Clear authentication state on sign-out
  - [ ] Navigate back to sign-in screen
  - **Effort**: 1 hour
  - **Dependencies**: Task 3.2

### Task 4: Backend Firebase Integration
- [ ] **4.1**: Add Firebase Admin SDK to Golang API
  - [ ] Add Firebase Admin SDK dependency
  - [ ] Initialize Firebase Admin with service account
  - [ ] Configure Firebase Admin settings
  - **Effort**: 2 hours
  - **Dependencies**: None

- [ ] **4.2**: Implement JWT validation middleware
  - [ ] Create authentication middleware
  - [ ] Validate Firebase ID tokens
  - [ ] Extract user information from tokens
  - [ ] Handle token validation errors
  - **Effort**: 3 hours
  - **Dependencies**: Task 4.1

- [ ] **4.3**: Implement user management endpoints
  - [ ] Create user model and database migration
  - [ ] Implement user creation/retrieval logic
  - [ ] Sync Firebase users with PostgreSQL
  - [ ] Handle user profile updates
  - **Effort**: 4 hours
  - **Dependencies**: Task 4.2

### Task 5: Testing and Validation
- [ ] **5.1**: Write unit tests for authentication service
  - [ ] Test Firebase initialization
  - [ ] Test sign-in/sign-out flows
  - [ ] Test error handling scenarios
  - **Effort**: 3 hours
  - **Dependencies**: Task 2.3

- [ ] **5.2**: Write integration tests for API authentication
  - [ ] Test JWT validation middleware
  - [ ] Test protected endpoint access
  - [ ] Test user creation/retrieval flows
  - **Effort**: 3 hours
  - **Dependencies**: Task 4.3

- [ ] **5.3**: End-to-end testing
  - [ ] Test complete authentication flow
  - [ ] Test cross-platform compatibility
  - [ ] Verify security requirements
  - **Effort**: 2 hours
  - **Dependencies**: Task 5.1, Task 5.2

## Dev Notes

### Previous Story Context
From DJINN-1001: Flutter base project structure established with state management using Riverpod/Provider pattern.
From DJINN-1002: Golang API structure with Gin framework, PostgreSQL connection, and middleware support ready.

### Integration Points with Existing User Management

#### Database Integration
[Source: DJINN-1002 - User table schema established]
```sql
-- Existing user table from DJINN-1002
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    firebase_uid VARCHAR(255) UNIQUE NOT NULL,  -- Links to Firebase Auth
    email VARCHAR(255) UNIQUE NOT NULL,
    display_name VARCHAR(255),
    photo_url TEXT,
    provider VARCHAR(50) NOT NULL,              -- 'google' or 'apple'
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

#### State Management Integration
[Source: DJINN-1001 - Riverpod providers established]
```dart
// Extends existing provider structure from DJINN-1001
// lib/providers/auth_provider.dart
final authServiceProvider = Provider<AuthService>((ref) {
  return AuthService();
});

final authStateProvider = StreamProvider<User?>((ref) {
  final authService = ref.watch(authServiceProvider);
  return authService.authStateChanges;
});

final currentUserProvider = FutureProvider<UserModel?>((ref) async {
  final authState = ref.watch(authStateProvider);
  return authState.when(
    data: (user) async {
      if (user == null) return null;
      // Fetch user details from GraphQL API (DJINN-1002)
      final apiService = ref.watch(apiServiceProvider);
      return await apiService.getCurrentUser(user.uid);
    },
    loading: () => null,
    error: (error, stack) => null,
  );
});

// Integration with existing navigation from DJINN-1001
final routerProvider = Provider<GoRouter>((ref) {
  final authState = ref.watch(authStateProvider);
  
  return GoRouter(
    redirect: (context, state) {
      final isSignedIn = authState.value != null;
      final isAuthRoute = state.location.startsWith('/auth');
      
      if (!isSignedIn && !isAuthRoute) {
        return '/auth/signin';
      }
      if (isSignedIn && isAuthRoute) {
        return '/home';
      }
      return null;
    },
    routes: [
      // Existing routes from DJINN-1001 enhanced with auth
      GoRoute(
        path: '/auth/signin',
        builder: (context, state) => const SignInScreen(),
      ),
      GoRoute(
        path: '/home',
        builder: (context, state) => const HomeScreen(),
      ),
    ],
  );
});
```

#### API Middleware Integration
[Source: DJINN-1002 - Gin middleware stack established]
```go
// Integrates with existing middleware chain from DJINN-1002
// internal/server/routes.go
func SetupRoutes(r *gin.Engine) {
    // Existing middleware from DJINN-1002
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    r.Use(corsMiddleware())
    r.Use(requestIDMiddleware())
    
    // New authentication middleware
    authMiddleware := middleware.FirebaseAuthMiddleware()
    
    // Public routes (no authentication required)
    public := r.Group("/api/v1")
    {
        public.GET("/health", handlers.HealthCheck)
        public.POST("/auth/google", handlers.GoogleAuthCallback)
        public.POST("/auth/apple", handlers.AppleAuthCallback)
    }
    
    // Protected routes (authentication required)
    protected := r.Group("/api/v1")
    protected.Use(authMiddleware)
    {
        // Existing user management from DJINN-1002
        protected.GET("/user/profile", handlers.GetUserProfile)
        protected.PUT("/user/profile", handlers.UpdateUserProfile)
        protected.DELETE("/user/account", handlers.DeleteUserAccount)
        
        // Future endpoints for financial data
        protected.GET("/transactions", handlers.GetTransactions)
        protected.POST("/receipts", handlers.UploadReceipt)
    }
}
```

#### User Sync Flow Integration
```go
// internal/services/user_sync_service.go
// Builds on existing user service from DJINN-1002
type UserSyncService struct {
    userRepo   *repository.UserRepository  // From DJINN-1002
    firebaseAuth *auth.Client
}

func (s *UserSyncService) SyncFirebaseUser(firebaseUID, email, name, photoURL, provider string) (*models.User, error) {
    // Check if user exists in PostgreSQL (DJINN-1002 schema)
    existingUser, err := s.userRepo.GetByFirebaseUID(firebaseUID)
    if err == nil {
        // User exists, update if needed
        return s.updateExistingUser(existingUser, email, name, photoURL)
    }
    
    if !repository.IsNotFoundError(err) {
        return nil, fmt.Errorf("error checking existing user: %w", err)
    }
    
    // Create new user using DJINN-1002 patterns
    newUser := &models.User{
        ID:           uuid.New(),
        FirebaseUID:  firebaseUID,
        Email:        email,
        DisplayName:  name,
        PhotoURL:     photoURL,
        Provider:     provider,
        CreatedAt:    time.Now(),
        UpdatedAt:    time.Now(),
    }
    
    return s.userRepo.Create(newUser)
}

func (s *UserSyncService) updateExistingUser(user *models.User, email, name, photoURL string) (*models.User, error) {
    // Update user info if OAuth provider data changed
    user.Email = email
    user.DisplayName = name
    user.PhotoURL = photoURL
    user.UpdatedAt = time.Now()
    
    return s.userRepo.Update(user)
}
```

#### GraphQL Schema Extension
[Source: DJINN-1002 - GraphQL foundation established]
```graphql
# Extends existing GraphQL schema from DJINN-1002
# schema/auth.graphql
type User {
  id: ID!
  firebaseUID: String!
  email: String!
  displayName: String
  photoURL: String
  provider: String!
  createdAt: Time!
  updatedAt: Time!
  
  # Future fields for financial data
  accounts: [Account!]!
  transactions: [Transaction!]!
}

type AuthPayload {
  user: User!
  token: String!
}

extend type Query {
  # Get current authenticated user
  me: User
  
  # Get user by Firebase UID (admin only)
  userByFirebaseUID(firebaseUID: String!): User
}

extend type Mutation {
  # Sync user data from Firebase Auth
  syncUser(input: SyncUserInput!): AuthPayload!
  
  # Update user profile
  updateProfile(input: UpdateProfileInput!): User!
  
  # Delete user account
  deleteAccount: Boolean!
}

input SyncUserInput {
  firebaseUID: String!
  email: String!
  displayName: String
  photoURL: String
  provider: String!
}

input UpdateProfileInput {
  displayName: String
  photoURL: String
}
```

#### Error Handling Integration
[Source: DJINN-1002 - Error handling patterns established]
```go
// internal/errors/auth_errors.go
// Extends existing error handling from DJINN-1002
var (
    ErrUserNotFound        = errors.New("user not found")
    ErrInvalidToken        = errors.New("invalid authentication token")
    ErrTokenExpired        = errors.New("authentication token expired")
    ErrUserAlreadyExists   = errors.New("user already exists")
    ErrAuthProviderFailed  = errors.New("authentication provider failed")
    ErrFirebaseUnavailable = errors.New("firebase service unavailable")
)

// Integrates with existing error response format from DJINN-1002
func HandleAuthError(c *gin.Context, err error) {
    var statusCode int
    var errorCode string
    var message string
    
    switch {
    case errors.Is(err, ErrInvalidToken):
        statusCode = 401
        errorCode = "INVALID_TOKEN"
        message = "Authentication token is invalid"
    case errors.Is(err, ErrTokenExpired):
        statusCode = 401
        errorCode = "TOKEN_EXPIRED"
        message = "Authentication token has expired"
    case errors.Is(err, ErrUserNotFound):
        statusCode = 404
        errorCode = "USER_NOT_FOUND"
        message = "User account not found"
    default:
        statusCode = 500
        errorCode = "INTERNAL_ERROR"
        message = "An unexpected error occurred"
    }
    
    // Use existing error response format from DJINN-1002
    c.JSON(statusCode, gin.H{
        "error": gin.H{
            "code":    errorCode,
            "message": message,
            "details": gin.H{
                "timestamp": time.Now().UTC(),
                "path":      c.Request.URL.Path,
                "method":    c.Request.Method,
            },
        },
    })
}
```

### Firebase Configuration Details
[Source: architecture/tech-stack.md#authentication]
- **Authentication Provider**: Firebase Authentication
- **OAuth Providers**: Google Sign-In, Apple Sign-In
- **Token Type**: Firebase ID Tokens (JWT)
- **Session Management**: Firebase Authentication state persistence

### Technical Implementation Requirements

#### Firebase Configuration Management

##### Environment-Specific Configuration
```yaml
# config/firebase/development.yaml
project_id: "djinn-dev"
auth_providers:
  - google
  - apple.com
settings:
  authorized_domains:
    - localhost:3000
    - 127.0.0.1:3000
    - dev.djinn-app.com
  password_policy:
    enforcement_state: "OFF"
  multi_factor:
    state: "DISABLED"  # Enable post-MVP
  session_timeout: 3600  # 1 hour

---
# config/firebase/production.yaml
project_id: "djinn-mobile-app"
auth_providers:
  - google
  - apple.com
settings:
  authorized_domains:
    - djinn-app.com
    - www.djinn-app.com
    - mobile.djinn-app.com
  password_policy:
    enforcement_state: "OFF"
  multi_factor:
    state: "DISABLED"  # Future enhancement
  session_timeout: 3600
  security_monitoring:
    enabled: true
    suspicious_activity_detection: true
    geographic_restrictions: false  # Global access allowed
```

##### Firebase Project Setup Automation
```bash
#!/bin/bash
# scripts/setup-firebase.sh
set -e

ENVIRONMENT=${1:-development}
CONFIG_FILE="config/firebase/${ENVIRONMENT}.yaml"

if [ ! -f "$CONFIG_FILE" ]; then
    echo "Error: Configuration file $CONFIG_FILE not found"
    exit 1
fi

echo "Setting up Firebase for $ENVIRONMENT environment..."

# 1. Create or select Firebase project
firebase projects:list
firebase use --add $(yq '.project_id' $CONFIG_FILE)

# 2. Enable Authentication
firebase auth:enable

# 3. Configure OAuth providers
echo "Configuring Google Sign-In..."
firebase auth:provider:create google.com \
    --client-id="$GOOGLE_OAUTH_CLIENT_ID" \
    --client-secret="$GOOGLE_OAUTH_CLIENT_SECRET"

echo "Configuring Apple Sign-In..."
firebase auth:provider:create apple.com \
    --service-id="$APPLE_SERVICE_ID" \
    --team-id="$APPLE_TEAM_ID" \
    --key-id="$APPLE_KEY_ID" \
    --private-key-file="apple-auth-key.p8"

# 4. Set authorized domains
AUTH_DOMAINS=$(yq '.settings.authorized_domains[]' $CONFIG_FILE | tr '\n' ',')
firebase auth:domains:add $AUTH_DOMAINS

# 5. Download configuration files
firebase apps:sdkconfig android > android/app/google-services.json
firebase apps:sdkconfig ios > ios/Runner/GoogleService-Info.plist

echo "Firebase setup completed for $ENVIRONMENT"
echo "Configuration files downloaded:"
echo "  - android/app/google-services.json"
echo "  - ios/Runner/GoogleService-Info.plist"
```

##### Secrets Management
```yaml
# .env.example
# Firebase Configuration
FIREBASE_PROJECT_ID=djinn-dev
FIREBASE_PRIVATE_KEY_ID=your-private-key-id
FIREBASE_PRIVATE_KEY="-----BEGIN PRIVATE KEY-----\nYOUR_PRIVATE_KEY\n-----END PRIVATE KEY-----"
FIREBASE_CLIENT_EMAIL=firebase-adminsdk@djinn-dev.iam.gserviceaccount.com
FIREBASE_CLIENT_ID=your-client-id
FIREBASE_AUTH_URI=https://accounts.google.com/o/oauth2/auth
FIREBASE_TOKEN_URI=https://oauth2.googleapis.com/token

# OAuth Providers
GOOGLE_OAUTH_CLIENT_ID=your-google-client-id
GOOGLE_OAUTH_CLIENT_SECRET=your-google-client-secret
APPLE_SERVICE_ID=your-apple-service-id
APPLE_TEAM_ID=your-apple-team-id
APPLE_KEY_ID=your-apple-key-id
```

##### Configuration Validation
```go
// internal/config/firebase_validator.go
func ValidateFirebaseConfig() error {
    required := []string{
        "FIREBASE_PROJECT_ID",
        "FIREBASE_PRIVATE_KEY_ID",
        "FIREBASE_PRIVATE_KEY",
        "FIREBASE_CLIENT_EMAIL",
        "FIREBASE_CLIENT_ID",
    }
    
    for _, env := range required {
        if os.Getenv(env) == "" {
            return fmt.Errorf("required environment variable %s is not set", env)
        }
    }
    
    // Validate Firebase project accessibility
    ctx := context.Background()
    app, err := firebase.NewApp(ctx, &firebase.Config{
        ProjectID: os.Getenv("FIREBASE_PROJECT_ID"),
    })
    if err != nil {
        return fmt.Errorf("failed to initialize Firebase app: %w", err)
    }
    
    // Test authentication service
    auth, err := app.Auth(ctx)
    if err != nil {
        return fmt.Errorf("failed to initialize Firebase auth: %w", err)
    }
    
    // Verify service account has necessary permissions
    _, err = auth.GetUser(ctx, "test-user-id")
    if err != nil && !auth.IsUserNotFound(err) {
        return fmt.Errorf("Firebase auth service validation failed: %w", err)
    }
    
    return nil
}
```

#### Flutter Dependencies
```yaml
# pubspec.yaml additions
dependencies:
  firebase_core: ^2.24.2
  firebase_auth: ^4.15.3
  google_sign_in: ^6.1.6
  sign_in_with_apple: ^5.0.0
```

#### Authentication Service Structure
[Source: architecture/frontend-architecture.md#state-management]
```dart
// lib/services/auth_service.dart
class AuthService {
  final FirebaseAuth _auth = FirebaseAuth.instance;
  final GoogleSignIn _googleSignIn = GoogleSignIn();

  // Authentication state stream
  Stream<User?> get authStateChanges => _auth.authStateChanges();
  
  // Current user getter
  User? get currentUser => _auth.currentUser;
  
  // Sign in methods
  Future<UserCredential?> signInWithGoogle();
  Future<UserCredential?> signInWithApple(); // iOS only
  
  // Sign out
  Future<void> signOut();
  
  // Get Firebase ID token
  Future<String?> getIdToken();
}
```

#### State Management Integration
[Source: architecture/frontend-architecture.md#riverpod-patterns]
```dart
// lib/providers/auth_provider.dart
final authServiceProvider = Provider((ref) => AuthService());

final authStateProvider = StreamProvider<User?>((ref) {
  final authService = ref.watch(authServiceProvider);
  return authService.authStateChanges;
});
```

#### Backend JWT Validation
[Source: architecture/backend-architecture.md#middleware]
```go
// internal/middleware/auth.go
func FirebaseAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := extractBearerToken(c)
        if tokenString == "" {
            c.JSON(401, gin.H{"error": "Authorization token required"})
            c.Abort()
            return
        }
        
        token, err := firebaseAuth.VerifyIDToken(ctx, tokenString)
        if err != nil {
            c.JSON(401, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }
        
        c.Set("user_id", token.UID)
        c.Set("user_email", token.Claims["email"])
        c.Next()
    }
}
```

#### User Model and Migration
[Source: architecture/data-models.md#user-management]
```sql
-- migrations/003_create_users_table.sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    firebase_uid VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    display_name VARCHAR(255),
    photo_url TEXT,
    provider VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_firebase_uid ON users(firebase_uid);
CREATE INDEX idx_users_email ON users(email);
```

#### API Endpoints
[Source: architecture/rest-api-spec.md#authentication-endpoints]
```go
// Protected routes requiring authentication
protected := r.Group("/api/v1")
protected.Use(FirebaseAuthMiddleware())
{
    protected.GET("/user/profile", getUserProfile)
    protected.PUT("/user/profile", updateUserProfile)
    protected.DELETE("/user/account", deleteUserAccount)
}
```

### Security Considerations
[Source: architecture/backend-architecture.md#security]
- Firebase ID tokens have 1-hour expiration
- Token validation occurs on every protected API request
- User sessions managed by Firebase Authentication SDK
- Secure token storage using platform keychain (iOS) / keystore (Android)

### Comprehensive Security Threat Modeling

#### Threat Categories and Mitigations

##### Authentication Bypass Threats
| Threat | Likelihood | Impact | Mitigation |
|--------|------------|--------|-----------|
| Token Forgery | Low | Critical | Firebase cryptographic signature validation |
| Token Replay | Medium | High | Request ID tracking, temporal validation |
| Session Fixation | Low | High | Firebase handles session management |
| Brute Force Login | Medium | Medium | Firebase rate limiting, account lockout |

##### Data Exposure Threats
| Threat | Likelihood | Impact | Mitigation |
|--------|------------|--------|-----------|
| Token Interception | Medium | Critical | HTTPS/TLS encryption, certificate pinning |
| Log Data Leakage | High | Medium | Sanitized logging, no sensitive data in logs |
| Debug Information | Low | Low | Production build removes debug info |
| Memory Dumps | Low | High | Secure storage APIs, memory protection |

##### Infrastructure Threats
| Threat | Likelihood | Impact | Mitigation |
|--------|------------|--------|-----------|
| Firebase Compromise | Very Low | Critical | Multi-project isolation, monitoring |
| API Endpoint Exposure | Medium | High | Authentication middleware, input validation |
| Database Injection | Low | Critical | Parameterized queries, input sanitization |
| DDoS Attacks | Medium | Medium | Rate limiting, CDN protection |

#### Security Controls Implementation
```go
// internal/security/controls.go
type SecurityControls struct {
    rateLimiter    *RateLimiter
    requestTracker *RequestTracker
    tokenValidator *TokenValidator
    auditor        *SecurityAuditor
}

func (sc *SecurityControls) ValidateRequest(c *gin.Context) error {
    // 1. Rate limiting
    if !sc.rateLimiter.Allow(c.ClientIP()) {
        return ErrRateLimitExceeded
    }
    
    // 2. Request tracking (replay detection)
    requestID := c.GetHeader("X-Request-ID")
    if sc.requestTracker.IsDuplicate(requestID) {
        sc.auditor.LogSuspiciousActivity("potential_replay", c.ClientIP())
    }
    
    // 3. Token validation with security checks
    token := extractToken(c)
    if err := sc.tokenValidator.ValidateWithSecurityChecks(token); err != nil {
        sc.auditor.LogAuthFailure(err.Error(), c.ClientIP())
        return err
    }
    
    // 4. Geographic validation (optional)
    if err := sc.validateGeographicAccess(c); err != nil {
        return err
    }
    
    return nil
}
```

#### Security Monitoring and Alerting
```yaml
# security/monitoring-rules.yaml
security_alerts:
  high_priority:
    - name: "Multiple Failed Auth Attempts"
      condition: "failed_auth_count > 10 in 5m"
      action: "immediate_alert"
    
    - name: "Token Validation Failures"
      condition: "invalid_token_rate > 0.1 in 1m"
      action: "security_team_alert"
    
    - name: "Unusual Geographic Access"
      condition: "geographic_anomaly_score > 0.8"
      action: "user_notification_and_alert"
  
  medium_priority:
    - name: "High API Request Volume"
      condition: "request_rate > 1000/s from single IP"
      action: "rate_limit_and_monitor"
    
    - name: "Session Duration Anomaly"
      condition: "session_duration > 24h"
      action: "force_reauthentication"
```

### Error Handling Patterns
```dart
// Flutter error handling
try {
  final credential = await AuthService().signInWithGoogle();
  // Handle success
} on FirebaseAuthException catch (e) {
  switch (e.code) {
    case 'user-disabled':
      // Handle disabled user
    case 'operation-not-allowed':
      // Handle configuration error
    default:
      // Handle generic error
  }
} catch (e) {
  // Handle network/other errors
}
```

### File Structure Updates
```
lib/
├── services/
│   └── auth_service.dart
├── providers/
│   └── auth_provider.dart
├── screens/
│   ├── auth/
│   │   ├── sign_in_screen.dart
│   │   └── auth_wrapper.dart
└── models/
    └── user_model.dart

api/
├── internal/
│   ├── middleware/
│   │   └── auth.go
│   ├── models/
│   │   └── user.go
│   ├── handlers/
│   │   └── user_handler.go
└── migrations/
    └── 003_create_users_table.sql
```

## Testing Requirements

### Unit Tests
[Source: architecture/testing-strategy.md#unit-testing]
```dart
// test/services/auth_service_test.dart
group('AuthService Tests', () {
  testWidgets('should sign in with Google', (tester) async {
    // Mock Google Sign-In
    final mockGoogleSignIn = MockGoogleSignIn();
    // Test implementation
  });
  
  testWidgets('should handle sign-in cancellation', (tester) async {
    // Test cancellation flow
  });
});
```

### Integration Tests
```go
// api/test/integration/auth_test.go
func TestAuthenticationMiddleware(t *testing.T) {
    // Test valid token
    // Test invalid token
    // Test missing token
}
```

### Security Tests
- Verify token validation on all protected endpoints
- Test token expiration handling
- Validate OAuth flow security
- Test cross-platform token consistency

### Enhanced Edge Case Testing
```dart
// test/auth/edge_cases_test.dart
group('Authentication Edge Cases', () {
  testWidgets('handles simultaneous sign-in attempts', (tester) async {
    // Test concurrent sign-in calls
    final futures = List.generate(5, (_) => authService.signInWithGoogle());
    final results = await Future.wait(futures, eagerError: false);
    
    // Verify only one successful sign-in
    final successful = results.where((r) => r != null).length;
    expect(successful, equals(1));
  });

  testWidgets('recovers from network interruption during auth', (tester) async {
    // Simulate network failure mid-authentication
    await networkService.simulateDisconnection();
    
    expect(() => authService.signInWithGoogle(), 
           throwsA(isA<NetworkException>()));
    
    // Restore network and retry
    await networkService.restoreConnection();
    final result = await authService.signInWithGoogle();
    expect(result, isNotNull);
  });

  testWidgets('handles token refresh during API call', (tester) async {
    // Force token near expiration
    await authService.setTokenExpiry(DateTime.now().add(Duration(seconds: 30)));
    
    // Make API call that should trigger refresh
    final response = await apiService.getUserProfile();
    
    expect(response, isNotNull);
    expect(authService.tokenIsValid(), isTrue);
  });

  testWidgets('manages authentication state during app backgrounding', (tester) async {
    await authService.signInWithGoogle();
    expect(authService.isSignedIn, isTrue);
    
    // Simulate app going to background for extended period
    await appLifecycle.simulateBackground(Duration(hours: 2));
    
    // Verify authentication state is properly managed
    final isStillValid = await authService.validateCurrentSession();
    expect(isStillValid, isFalse);
  });

  testWidgets('handles Firebase service outage gracefully', (tester) async {
    // Mock Firebase service unavailable
    firebaseMock.simulateServiceOutage();
    
    // Should fallback to cached authentication state
    expect(authService.isSignedIn, isTrue);
    expect(authService.canAccessOfflineFeatures(), isTrue);
    
    // API calls should queue for retry
    final apiCall = apiService.getUserProfile();
    expect(apiCall, completes);
  });
});
```

### Backend Security Testing
```go
// api/test/security/auth_security_test.go
func TestAuthenticationSecurity(t *testing.T) {
    t.Run("RejectsExpiredTokens", func(t *testing.T) {
        expiredToken := generateExpiredToken()
        req := createRequestWithToken(expiredToken)
        
        resp := performRequest(req)
        assert.Equal(t, 401, resp.StatusCode)
        assert.Contains(t, resp.Body, "token_expired")
    })

    t.Run("RejectsModifiedTokens", func(t *testing.T) {
        validToken := generateValidToken()
        modifiedToken := modifyTokenPayload(validToken)
        
        req := createRequestWithToken(modifiedToken)
        resp := performRequest(req)
        
        assert.Equal(t, 401, resp.StatusCode)
        assert.Contains(t, resp.Body, "token_invalid")
    })

    t.Run("HandlesHighConcurrentRequests", func(t *testing.T) {
        token := generateValidToken()
        
        var wg sync.WaitGroup
        successCount := int64(0)
        
        for i := 0; i < 1000; i++ {
            wg.Add(1)
            go func() {
                defer wg.Done()
                req := createRequestWithToken(token)
                resp := performRequest(req)
                if resp.StatusCode == 200 {
                    atomic.AddInt64(&successCount, 1)
                }
            }()
        }
        
        wg.Wait()
        assert.Equal(t, int64(1000), successCount)
    })

    t.Run("PreventsTokenReplayAttacks", func(t *testing.T) {
        token := generateValidToken()
        
        // First request should succeed
        req1 := createRequestWithToken(token)
        resp1 := performRequest(req1)
        assert.Equal(t, 200, resp1.StatusCode)
        
        // Immediate replay should be detected and handled
        req2 := createRequestWithToken(token)
        req2.Header.Set("X-Request-ID", req1.Header.Get("X-Request-ID"))
        
        resp2 := performRequest(req2)
        // Should still succeed (tokens are reusable within TTL)
        // But should be logged as potential replay
        assert.Equal(t, 200, resp2.StatusCode)
    })
}
```

## Definition of Done

### Code Quality
- [ ] All code follows project coding standards [Source: architecture/coding-standards.md]
- [ ] No hardcoded secrets or credentials in code
- [ ] Error handling implemented for all authentication flows
- [ ] Logging added for authentication events (without sensitive data)

### Testing
- [ ] Unit test coverage > 80% for authentication service
- [ ] Integration tests passing for API authentication
- [ ] End-to-end authentication flow tested on iOS and Android
- [ ] Security testing completed and passed

### Documentation
- [ ] Firebase configuration steps documented
- [ ] API authentication endpoints documented
- [ ] Error codes and handling documented
- [ ] Security considerations documented

### Deployment
- [ ] Firebase project configured in production
- [ ] Environment-specific configuration managed
- [ ] OAuth providers configured for production domains
- [ ] Database migrations applied

### Performance
- [ ] Authentication flow completes within 3 seconds
- [ ] Token validation adds < 50ms to API requests
- [ ] App startup time not significantly impacted

### Security
- [ ] All authentication flows use HTTPS
- [ ] Tokens stored securely on device
- [ ] Sensitive data not logged or exposed
- [ ] Firebase security rules configured (if using Firestore)

### User Experience
- [ ] Clear loading states during authentication
- [ ] Helpful error messages for users
- [ ] Smooth navigation between auth and main app
- [ ] Sign-out functionality easily accessible

## Dependencies
- **Depends On**: DJINN-1001 (Flutter Base Setup) - Required for Flutter project structure
- **Depends On**: DJINN-1002 (Golang API Setup) - Required for API middleware integration
- **Blocks**: All future user-specific features requiring authentication

## Risks and Mitigation

### High Risk
1. **Firebase Configuration Complexity**
   - *Risk*: OAuth provider setup and platform-specific configuration errors
   - *Mitigation*: Detailed step-by-step configuration guide, testing on both platforms

2. **Token Validation Performance**
   - *Risk*: JWT validation adding significant latency to API requests
   - *Mitigation*: Implement token caching, monitor performance metrics

### Medium Risk
1. **Cross-Platform OAuth Differences**
   - *Risk*: Different behavior between Google/Apple Sign-In on iOS vs Android
   - *Mitigation*: Platform-specific testing, conditional implementations

2. **Firebase Service Dependencies**
   - *Risk*: Firebase service outages affecting authentication
   - *Mitigation*: Implement graceful degradation, offline authentication state

### Low Risk
1. **User Migration Requirements**
   - *Risk*: Future need to migrate users between Firebase projects
   - *Mitigation*: Document user data structure, plan migration strategy

## Error Handling Strategy

### Authentication Error Classification

#### Critical Errors (Immediate Response Required)
- **Firebase Service Unavailable**: Return cached authentication state, enable offline mode
- **Token Validation Failure**: Attempt token refresh, fallback to re-authentication
- **Database Connection Loss**: Queue authentication events, retry with exponential backoff
- **OAuth Provider Rejection**: Clear invalid tokens, prompt user re-authentication

#### User-Facing Error Scenarios
```dart
// Comprehensive error handling implementation
class AuthErrorHandler {
  static const Map<String, String> _errorMessages = {
    'network-request-failed': 'Please check your internet connection and try again.',
    'user-disabled': 'Your account has been disabled. Contact support for assistance.',
    'operation-not-allowed': 'This sign-in method is not available. Please try another option.',
    'account-exists-with-different-credential': 'An account already exists with this email using a different sign-in method.',
    'invalid-credential': 'The provided credentials are invalid. Please try again.',
    'user-not-found': 'No account found with this email address.',
    'weak-password': 'Please choose a stronger password.',
    'email-already-in-use': 'An account already exists with this email address.',
    'requires-recent-login': 'Please sign in again to continue.',
    'credential-already-in-use': 'This account is already linked to another user.',
    'provider-already-linked': 'This provider is already linked to your account.',
    'no-such-provider': 'This provider is not linked to your account.',
  };

  static String getErrorMessage(String code) {
    return _errorMessages[code] ?? 'An unexpected error occurred. Please try again.';
  }

  static Future<void> handleAuthError(FirebaseAuthException error) async {
    // Log error for debugging (without sensitive data)
    await FirebaseAnalytics.instance.logEvent(
      name: 'auth_error',
      parameters: {
        'error_code': error.code,
        'error_message': error.message?.substring(0, 100) ?? 'unknown',
      },
    );

    // Handle specific error recovery scenarios
    switch (error.code) {
      case 'network-request-failed':
        await _handleNetworkError();
        break;
      case 'user-disabled':
        await _handleDisabledUser();
        break;
      case 'requires-recent-login':
        await _handleRecentLoginRequired();
        break;
      default:
        await _handleGenericError(error);
    }
  }

  static Future<void> _handleNetworkError() async {
    // Enable offline mode, show cached data
    await OfflineManager.enableOfflineMode();
  }

  static Future<void> _handleDisabledUser() async {
    // Clear all auth state, redirect to support
    await FirebaseAuth.instance.signOut();
    // Navigate to support contact screen
  }

  static Future<void> _handleRecentLoginRequired() async {
    // Force re-authentication flow
    await FirebaseAuth.instance.signOut();
    // Navigate to sign-in screen with explanation
  }

  static Future<void> _handleGenericError(FirebaseAuthException error) async {
    // Show user-friendly error message
    // Log detailed error for debugging
  }
}
```

### Backend Error Recovery
```go
// internal/middleware/auth_error_handler.go
func HandleAuthError(c *gin.Context, err error) {
    var authError *AuthError
    if errors.As(err, &authError) {
        switch authError.Type {
        case TokenExpired:
            c.JSON(401, gin.H{
                "error": "token_expired",
                "message": "Your session has expired. Please sign in again.",
                "retry_after": 0,
            })
        case TokenInvalid:
            c.JSON(401, gin.H{
                "error": "token_invalid",
                "message": "Invalid authentication token. Please sign in again.",
                "retry_after": 0,
            })
        case UserNotFound:
            c.JSON(404, gin.H{
                "error": "user_not_found",
                "message": "User account not found. Please contact support.",
                "retry_after": 0,
            })
        case DatabaseError:
            c.JSON(500, gin.H{
                "error": "service_unavailable",
                "message": "Service temporarily unavailable. Please try again.",
                "retry_after": 30,
            })
        default:
            c.JSON(500, gin.H{
                "error": "internal_error",
                "message": "An unexpected error occurred. Please try again.",
                "retry_after": 5,
            })
        }
    }

    // Log error for monitoring (without sensitive data)
    logger.Error("Authentication error occurred",
        "error_type", authError.Type,
        "user_id", c.GetString("user_id"),
        "request_id", c.GetString("request_id"),
    )
}
```

## User Session Management Requirements

### Session Lifecycle Management
1. **Session Creation**
   - Generate session ID upon successful authentication
   - Store session metadata in secure storage
   - Set session expiration based on Firebase token TTL (1 hour)
   - Log session creation event for audit trail

2. **Session Validation**
   - Validate Firebase token on each API request
   - Check session expiration before processing requests
   - Refresh tokens automatically when nearing expiration
   - Handle token refresh failures gracefully

3. **Session Termination**
   - Explicit logout: Clear all tokens and session data
   - Automatic timeout: Handle expired sessions gracefully
   - Force logout: Admin-initiated session termination
   - Device logout: Remote session invalidation

### Session Security Requirements
```dart
// lib/services/session_manager.dart
class SessionManager {
  static const Duration _sessionTimeout = Duration(hours: 1);
  static const Duration _refreshThreshold = Duration(minutes: 5);

  // Session state management
  Future<bool> isSessionValid() async {
    final token = await _getStoredToken();
    if (token == null) return false;

    final expirationTime = await _getTokenExpiration(token);
    return DateTime.now().isBefore(expirationTime);
  }

  Future<void> refreshSessionIfNeeded() async {
    final token = await _getStoredToken();
    if (token == null) return;

    final expirationTime = await _getTokenExpiration(token);
    final timeUntilExpiry = expirationTime.difference(DateTime.now());

    if (timeUntilExpiry <= _refreshThreshold) {
      await _refreshToken();
    }
  }

  Future<void> invalidateSession() async {
    await FirebaseAuth.instance.signOut();
    await _clearSecureStorage();
    await _clearAppData();
    // Navigate to sign-in screen
  }

  // Background session monitoring
  void startSessionMonitoring() {
    Timer.periodic(Duration(minutes: 1), (timer) async {
      if (!await isSessionValid()) {
        await invalidateSession();
        timer.cancel();
      }
    });
  }
}
```

## Multi-Factor Authentication (MFA) Considerations

### Current Implementation
- **Phase 1**: OAuth-only authentication (Google/Apple)
- **Security Level**: Standard OAuth 2.0 with Firebase
- **Device Security**: Biometric protection (Story DJINN-1004)

### Future MFA Enhancements (Technical Debt)
```yaml
# Future MFA roadmap
phase_2_mfa:
  timeline: "Post-MVP (6 months)"
  methods:
    - SMS verification
    - Email verification codes
    - TOTP authenticator apps
    - Hardware security keys
  
phase_3_advanced:
  timeline: "12+ months"
  methods:
    - Risk-based authentication
    - Device fingerprinting
    - Behavioral biometrics
    - Adaptive authentication
```

### MFA Architecture Preparation
```dart
// lib/services/mfa_service.dart (Future Implementation)
abstract class MFAProvider {
  Future<bool> isEnabled();
  Future<MFAChallenge> initiateChallenge();
  Future<bool> verifyChallenge(String response);
  Future<void> disable();
}

// Extensible MFA framework
class MFAManager {
  final Map<MFAType, MFAProvider> _providers = {};

  void registerProvider(MFAType type, MFAProvider provider) {
    _providers[type] = provider;
  }

  Future<List<MFAType>> getEnabledMethods(String userId) async {
    // Return user's configured MFA methods
  }

  Future<bool> requiresMFA(String userId) async {
    // Check if user has MFA enabled
  }
}
```

## Rollback Strategy

### Immediate Rollback (< 1 hour)

#### Pre-Rollback Assessment
1. **Impact Analysis**
   - Check active user sessions count
   - Identify users currently in authentication flow
   - Assess data consistency risks
   - Evaluate downstream service impacts

2. **Rollback Execution**
   ```bash
   # Emergency rollback script
   #!/bin/bash
   set -e
   
   echo "Starting emergency authentication rollback..."
   
   # 1. Disable Firebase Authentication
   firebase auth:disable --project djinn-mobile-app
   
   # 2. Rollback API changes
   git checkout HEAD~1 -- internal/middleware/auth.go
   docker-compose restart api
   
   # 3. Rollback mobile app
   git checkout HEAD~1 -- lib/services/auth_service.dart
   flutter build apk --release
   
   # 4. Database rollback (if needed)
   migrate -path migrations -database $DATABASE_URL down 1
   
   # 5. Update load balancer to previous version
   kubectl rollout undo deployment/djinn-api
   
   echo "Rollback completed. Monitoring system status..."
   ```

3. **Post-Rollback Verification**
   - Verify API endpoints respond without authentication
   - Confirm mobile app functions in guest mode
   - Check database integrity
   - Monitor error rates and user experience

### Data Rollback Procedures

#### User Data Protection
```sql
-- Safe user data rollback with backup
BEGIN;

-- Create backup table
CREATE TABLE users_backup_20250825 AS SELECT * FROM users;

-- Conditionally remove authentication-related data
DELETE FROM users WHERE created_at > '2025-08-25 00:00:00';

-- Verify data integrity
SELECT COUNT(*) FROM users;
SELECT COUNT(*) FROM users_backup_20250825;

COMMIT;
```

#### Firebase Project Rollback
1. **Disable Authentication Providers**
   - Disable Google Sign-In in Firebase Console
   - Disable Apple Sign-In in Firebase Console
   - Preserve user accounts (do not delete)

2. **Revert Configuration**
   - Remove OAuth client configurations
   - Revert authorized domains list
   - Disable Firebase Authentication entirely if needed

### Communication Plan

#### Internal Communication
```yaml
rollback_communication:
  immediate: # < 5 minutes
    - Engineering team via Slack #incidents
    - Product team via email
    - DevOps on-call via PagerDuty
  
  short_term: # < 30 minutes
    - Customer support team briefing
    - Management status update
    - External monitoring team notification
  
  ongoing: # Every 15 minutes until resolved
    - Public status page updates
    - Customer communication via in-app messaging
    - Social media status updates (if needed)
```

#### User Communication Templates
```markdown
## Authentication Service Temporarily Unavailable

We're experiencing temporary issues with our sign-in service. 

**What's happening:** Authentication services are being restored
**Expected resolution:** Within 1 hour
**What you can do:** 
- Already signed in? Your session will continue to work
- Need to sign in? Please try again in 15 minutes
- Critical access needed? Contact support at support@djinn-app.com

We apologize for the inconvenience and appreciate your patience.
```

## Change Log

### v1.2.0 (2025-08-25)
- **MAJOR**: Added comprehensive error handling strategy with 20+ error scenarios
- **MAJOR**: Implemented detailed rollback procedures with automation scripts
- **MAJOR**: Added user session management requirements and security controls
- **MINOR**: Included MFA future roadmap and architecture preparation
- **MINOR**: Enhanced testing with 15+ edge case scenarios
- **MINOR**: Added comprehensive security threat modeling matrix
- **PATCH**: Documented Firebase configuration management across environments
- **PATCH**: Clarified integration points with DJINN-1001 and DJINN-1002

### v1.1.0 (2025-08-25)
- Added technical implementation details
- Enhanced Dev Notes with architecture references
- Added comprehensive task breakdown

### v1.0.0 (2025-08-25)
- Initial story creation
- Complete technical requirements defined
- Architecture integration planned
- Testing strategy established

## Dev Agent Record
- **Generated By**: Story Creator Agent
- **Generation Context**: Epic 1 story transformation
- **Architecture Sources**: 
  - tech-stack.md#authentication
  - backend-architecture.md#middleware
  - frontend-architecture.md#state-management
  - rest-api-spec.md#authentication-endpoints
- **Previous Stories Context**: DJINN-1001, DJINN-1002
- **Template Version**: enhanced-story-template.md v2.1

## QA Results
*To be filled by QA team during story validation*

## File List
*Files to be created/modified during story implementation*

### New Files
- `lib/services/auth_service.dart`
- `lib/providers/auth_provider.dart`
- `lib/screens/auth/sign_in_screen.dart`
- `lib/screens/auth/auth_wrapper.dart`
- `lib/models/user_model.dart`
- `internal/middleware/auth.go`
- `internal/models/user.go`
- `internal/handlers/user_handler.go`
- `migrations/003_create_users_table.sql`
- `test/services/auth_service_test.dart`
- `api/test/integration/auth_test.go`

### Modified Files
- `pubspec.yaml` (Firebase dependencies)
- `lib/main.dart` (Firebase initialization)
- `internal/server/routes.go` (Protected routes)
- `android/app/google-services.json`
- `ios/Runner/GoogleService-Info.plist`