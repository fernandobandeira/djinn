# ADR-20250819: Security Architecture

## Status
Accepted

## Context
Djinn handles personal financial data requiring appropriate security measures:
- Personal financial information (PII) and aggregated transaction history
- Plaid access tokens for read-only bank connections
- Receipt images potentially containing sensitive information
- Budget and spending patterns revealing lifestyle details
- Personal finance app with single-user data access (users only see own data)
- Mobile app requiring secure offline/online synchronization
- Privacy compliance (CCPA/CPRA for US users, GDPR for EU)
- Target market includes privacy-conscious ex-Mint users seeking data aggregation
- **Physical device security risk**: Unauthorized access if phone is stolen or borrowed

### Constraints
- Firebase Auth already chosen for authentication (tech stack ADR)
- PostgreSQL with Row Level Security capabilities
- Mobile-only architecture with offline support (no web app)
- Limited initial budget (need cost-effective solutions)
- Small team (cannot manage complex security infrastructure)
- Must maintain user trust as core value proposition
- No payment processing (read-only financial data aggregation)
- Users only access their own data (no cross-user authorization needed)

## Decision

### 1. Authentication Architecture

#### Primary Authentication: Firebase Auth (Mobile OAuth-Only)
```typescript
// Mobile-only OAuth authentication flow with biometric protection
interface AuthenticationFlow {
  // Initial authentication (only required for first login or after logout)
  initialAuth: {
    google: "OAuth 2.0 via Firebase",
    apple: "Sign in with Apple via Firebase"
  },
  
  // Biometric authentication (primary day-to-day access method)
  biometricAuth: {
    // Required scenarios
    alwaysRequired: [
      "App launch from closed state",
      "App returns from background after 5+ minutes",
      "Viewing sensitive data (full account numbers)",
      "Data export operations",
      "Deleting data or accounts"
    ],
    
    // Implementation
    ios: {
      primary: "Face ID",
      fallback: "Touch ID",
      lastResort: "Device passcode"
    },
    android: {
      primary: "BiometricPrompt API (fingerprint/face)",
      fallback: "Device PIN/pattern"
    },
    
    // User flow
    userExperience: {
      appLaunch: "Biometric prompt immediately on open",
      backgroundReturn: "Biometric if > 5 min, else seamless",
      timeout: "Biometric re-auth, NOT Google/Apple re-login",
      failure: "After 3 failed attempts, require device passcode"
    },
    
    // Security rationale
    threatModel: {
      scenario: "Phone stolen/borrowed while unlocked",
      protection: "Biometric gate prevents financial data access",
      compliance: "Meets financial app industry standards",
      userTrust: "Similar to banking apps users already trust"
    }
  },
  
  // Session management
  sessions: {
    firebaseToken: "30 days (refreshed automatically)",
    biometricGate: "Controls app access, not token validity",
    backgroundTimer: "5 minutes (resets on foreground)",
    
    // Key point: Firebase token remains valid even when biometric times out
    // User just needs biometric to unlock app, not to re-authenticate
    tokenPersistence: "Secure storage, survives app restarts"
  }
}

// Example flow
const appStateTransitions = {
  firstTimeUser: "Google/Apple OAuth → Store token → Enable biometric",
  dailyUse: "Open app → Face ID → Access granted",
  backgroundShort: "Background 2 min → Foreground → No auth needed",
  backgroundLong: "Background 10 min → Foreground → Face ID → Access granted",
  appClosed: "Force close → Open later → Face ID → Access granted",
  tokenExpired: "After 30 days → Google/Apple re-auth → Enable biometric"
}
```

#### JWT Token Strategy
```go
// Token structure
type Claims struct {
    jwt.RegisteredClaims
    UserID         string   `json:"uid"`
    SubscriptionTier string `json:"tier"`
    Permissions    []string `json:"perms"`
    DeviceID       string   `json:"did"`
    SessionID      string   `json:"sid"`
}

// Token validation middleware
func ValidateToken(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := extractToken(r)
        
        // Verify Firebase token
        decoded, err := firebaseAuth.VerifyIDToken(ctx, token)
        if err != nil {
            return unauthorized(w, "Invalid token")
        }
        
        // No sensitive operation checks needed for read-only data aggregation
        
        // Set user context for RLS
        ctx = context.WithValue(ctx, "user_id", decoded.UID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

### 2. Authorization Model

#### Simple Subscription-Based Authorization
```sql
-- Users table with subscription tier
CREATE TABLE users (
    id UUID PRIMARY KEY,
    firebase_uid TEXT UNIQUE NOT NULL,
    email TEXT,
    subscription_tier TEXT DEFAULT 'free' CHECK (subscription_tier IN ('free', 'premium')),
    subscription_expires_at TIMESTAMPTZ, -- NULL for free users
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- No feature flags table needed - using PostHog or similar service

-- No database functions - all logic in application layer
```

#### Feature Flag Management (PostHog)
```go
// Feature flag implementation using PostHog
package features

import (
    "github.com/posthog/posthog-go"
)

type FeatureFlags struct {
    client posthog.Client
}

func NewFeatureFlags() *FeatureFlags {
    // PostHog initialization
    client := posthog.New(os.Getenv("POSTHOG_API_KEY"))
    return &FeatureFlags{client: client}
}

// Check if feature is enabled for user
func (f *FeatureFlags) IsEnabled(userID string, feature string) bool {
    // PostHog handles:
    // - Percentage rollouts
    // - User targeting
    // - A/B testing
    // - Analytics tracking
    
    isEnabled, err := f.client.IsFeatureEnabled(
        posthog.FeatureFlagPayload{
            Key: feature,
            DistinctId: userID,
            Properties: posthog.Properties{
                "subscription_tier": getUserTier(userID),
            },
        },
    )
    
    if err != nil {
        // Default to disabled on error
        return false
    }
    
    return isEnabled
}

// Example usage in GraphQL resolver
func (r *Resolver) GetAdvancedAnalytics(ctx context.Context) (*Analytics, error) {
    userID := ctx.Value("user_id").(string)
    
    // Check feature flag for gradual rollout
    if !features.IsEnabled(userID, "advanced_analytics_v2") {
        return r.getLegacyAnalytics(ctx) // Fallback to old version
    }
    
    return r.getNewAnalytics(ctx)
}
```

#### Why PostHog?
- **Free tier**: 1M feature flag requests/month (plenty for MVP)
- **Analytics included**: Track feature adoption automatically
- **A/B testing**: Built-in experimentation capabilities
- **Easy integration**: SDKs for Go, Flutter, and web
- **Self-host option**: Can migrate to self-hosted if needed
- **Alternative options**: Flagsmith (50k requests free), Unleash (self-hosted free)
```

#### GraphQL Authorization Directives (Zero Trust)
```graphql
# Zero Trust: Authentication required by default
# Only @public endpoints are accessible without auth
directive @public on FIELD_DEFINITION
directive @premium on FIELD_DEFINITION

type Query {
  # Public endpoints must be explicitly marked
  health: String! @public
  mobileAppConfig: AppConfig! @public # Firebase config, feature flags
  
  # All other endpoints require authentication by default
  me: User! # Auth required (default)
  myAccounts: [Account!]! # Auth required (default)
  myTransactions: TransactionConnection! # Auth required (default)
  
  # Premium features (auth implicit, premium check added)
  advancedAnalytics: Analytics! @premium
  ocrReceipt(id: ID!): Receipt! @premium
  
  # Rate limiting handled at API Gateway level
}
```

### 3. Data Encryption Strategy (Go-Based)

#### Encryption Approach
```go
// Application-level encryption in Go 1.25 using AES-GCM
// Key stored in environment variable or cloud KMS, never in database

package encryption

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "os"
)

type EncryptionService struct {
    key []byte // Loaded from environment/KMS at startup
}

func NewEncryptionService() (*EncryptionService, error) {
    // Key from environment variable (production uses KMS)
    keyBase64 := os.Getenv("ENCRYPTION_KEY")
    if keyBase64 == "" {
        // In production, fetch from AWS KMS or GCP Secret Manager
        keyBase64 = fetchFromKMS()
    }
    
    key, err := base64.StdEncoding.DecodeString(keyBase64)
    if err != nil {
        return nil, err
    }
    
    return &EncryptionService{key: key}, nil
}

// Encrypt Plaid tokens before storing in database
func (e *EncryptionService) EncryptToken(plaintext string) (string, error) {
    block, err := aes.NewCipher(e.key)
    if err != nil {
        return "", err
    }
    
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }
    
    nonce := make([]byte, gcm.NonceSize())
    if _, err = rand.Read(nonce); err != nil {
        return "", err
    }
    
    ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
    return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt tokens when needed for Plaid API calls
func (e *EncryptionService) DecryptToken(ciphertext string) (string, error) {
    // Implementation similar to encrypt, but in reverse
    // ...
}
```

```sql
-- Database just stores encrypted tokens as text
CREATE TABLE secure_credentials (
    user_id UUID PRIMARY KEY REFERENCES users(id),
    plaid_access_token TEXT, -- Already encrypted by Go app
    encrypted_at TIMESTAMPTZ DEFAULT NOW(),
    key_version INTEGER DEFAULT 1
);
```

#### Key Management Strategy
- **Development**: AES-256 key in environment variable
- **Production**: Key stored in cloud KMS (AWS KMS or GCP Secret Manager)
- **Key Rotation**: Quarterly rotation with versioning support
- **No pgcrypto**: Avoids storing keys in database or PostgreSQL config
```

#### Data Classification and Encryption Requirements (Simplified)
| Data Type | Classification | Encryption | Storage | Retention |
|-----------|---------------|------------|---------|-----------|
| OAuth Tokens | Critical | Firebase managed | Never stored | N/A |
| Plaid Tokens | Critical | pgcrypto encryption | PostgreSQL | Until disconnected |
| Account Numbers | Medium | DB encryption + masking | Display last 4 only | 7 years |
| Transaction Data | Medium | TLS + DB encryption | PostgreSQL | 7 years |
| Receipt OCR Data | Low | DB encryption | PostgreSQL (JSON) | 7 years |
| Email Addresses | Low | DB encryption | PostgreSQL | Until deletion |
| User Preferences | Low | DB encryption | JSONB | Until deletion |

### 4. API Security

#### API Protection and Rate Limiting
```go
// Simplified rate limiting at API Gateway level
type APIProtectionConfig struct {
    // Public endpoints (whitelist for mobile app)
    PublicEndpoints: []string{
        "/health",
        "/version",
        "/mobile/config", // App configuration endpoint
        "/.well-known/*", // Firebase auth discovery
    },
    
    // All other endpoints require Firebase auth token
    RequireAuth: true,
    
    // Simple tier-based rate limiting
    RateLimits: map[string]int{
        "free_user":    100, // requests per minute
        "premium_user": 500, // requests per minute
    },
    
    // Feature-specific limits tracked in Redis
    FeatureLimits: map[string]map[string]int{
        "ocr_processing": {
            "free_user":    10,  // per day
            "premium_user": 100, // per day
        },
        "data_export": {
            "free_user":    3,  // per day
            "premium_user": 20, // per day
        },
    }
}

// Feature limit implementation using Redis
func CheckFeatureLimit(userID string, feature string, tier string) (bool, int) {
    key := fmt.Sprintf("limit:%s:%s:%s", feature, userID, today())
    
    // Get current usage from Redis
    count, _ := redis.Get(key).Int()
    limit := FeatureLimits[feature][tier]
    
    if count >= limit {
        return false, 0 // Limit exceeded
    }
    
    // Increment counter with 24h expiration
    redis.Incr(key)
    redis.Expire(key, 24*time.Hour)
    
    return true, limit - count - 1 // Allowed, return remaining
}

// GraphQL resolver example
func (r *Resolver) ProcessOCR(ctx context.Context, receiptID string) (*Receipt, error) {
    userID := ctx.Value("user_id").(string)
    tier := getUserTier(userID)
    
    // Check feature limit
    allowed, remaining := CheckFeatureLimit(userID, "ocr_processing", tier)
    if !allowed {
        return nil, fmt.Errorf("daily OCR limit reached for %s tier", tier)
    }
    
    // Process OCR...
    receipt := processReceipt(receiptID)
    
    // Add remaining count to response headers
    ctx = context.WithValue(ctx, "X-Feature-Remaining", remaining)
    
    return receipt, nil
}

// Zero Trust authentication middleware - auth required by default
func ZeroTrustAuthMiddleware() func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // In GraphQL, check if the field has @public directive
            // This is handled by gqlgen directive implementation
            
            // For REST endpoints, only allow specific public paths
            publicPaths := []string{"/health", "/version", "/mobile/config"}
            for _, path := range publicPaths {
                if r.URL.Path == path {
                    next.ServeHTTP(w, r)
                    return
                }
            }
            
            // Default: Authentication required for everything else
            token := extractFirebaseToken(r)
            if token == "" {
                http.Error(w, "Authentication required (zero trust)", 401)
                return
            }
            
            // Verify Firebase token
            decoded, err := firebaseAuth.VerifyIDToken(r.Context(), token)
            if err != nil {
                http.Error(w, "Invalid token", 401)
                return
            }
            
            // Log sensitive operations
            if isSensitiveOperation(r.URL.Path) {
                logSecurityEvent("sensitive_access", decoded.UID)
            }
            
            // Set user context for RLS
            ctx := context.WithValue(r.Context(), "user_id", decoded.UID)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

// GraphQL directive implementations
func PublicDirective(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
    // @public directive allows bypassing auth
    return next(ctx)
}

func PremiumDirective(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
    // Check if user has premium subscription
    userID := ctx.Value("user_id").(string)
    
    // Use application service to check premium status
    userService := ctx.Value("user_service").(*UserService)
    if !userService.IsPremium(userID) {
        return nil, fmt.Errorf("premium subscription required")
    }
    
    return next(ctx)
}

// Application layer premium check
type UserService struct {
    db *sql.DB
}

func (u *UserService) IsPremium(userID string) bool {
    var tier string
    var expiresAt *time.Time
    
    err := u.db.QueryRow(`
        SELECT subscription_tier, subscription_expires_at 
        FROM users WHERE firebase_uid = $1
    `, userID).Scan(&tier, &expiresAt)
    
    if err != nil {
        return false
    }
    
    // Premium if tier is premium and not expired
    return tier == "premium" && (expiresAt == nil || expiresAt.After(time.Now()))
}

func AuthRequired(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
    // Default behavior - auth required unless @public is present
    userID := ctx.Value("user_id")
    if userID == nil {
        return nil, fmt.Errorf("authentication required")
    }
    return next(ctx)
}
```

#### Input Validation and Sanitization
```go
// GraphQL query complexity limiting
func QueryComplexityAnalyzer() *gqlcomplexity.Analyzer {
    return gqlcomplexity.NewAnalyzer(schema, gqlcomplexity.Config{
        MaxComplexity: map[string]int{
            "free_user": 100,
            "premium_user": 500,
            "admin": 1000,
        },
        Variables: map[string]interface{}{
            "first": 20, // Default pagination limit
        },
        ScalarCost: 1,
        ObjectCost: 2,
        ListFactor: 10,
        DepthLimit: 7,
    })
}

// Input sanitization
func SanitizeInput(input string, fieldType string) (string, error) {
    // Remove null bytes
    input = strings.ReplaceAll(input, "\x00", "")
    
    switch fieldType {
    case "email":
        return validateEmail(input)
    case "currency":
        return validateCurrency(input)
    case "description":
        return sanitizeHTML(input)
    case "amount":
        return validateAmount(input)
    default:
        return escapeSpecialChars(input), nil
    }
}
```

### 5. Infrastructure Security

#### Network Security
```yaml
# Network segmentation
networks:
  public_subnet:
    components: ["load_balancer", "cdn"]
    ingress: ["443", "80->443 redirect"]
    
  app_subnet:
    components: ["api_servers", "graphql"]
    ingress: ["from:public_subnet:443"]
    egress: ["to:data_subnet:5432", "to:cache_subnet:6379"]
    
  data_subnet:
    components: ["postgresql", "redis"]
    ingress: ["from:app_subnet:5432,6379"]
    egress: ["to:backup_subnet:443"]
    
  backup_subnet:
    components: ["s3_backup", "snapshots"]
    ingress: ["from:data_subnet:443"]
```

#### Secrets Management
```yaml
# AWS Secrets Manager / GCP Secret Manager configuration
secrets:
  storage:
    provider: "aws_secrets_manager" # or gcp_secret_manager
    kms_key: "arn:aws:kms:region:account:key/uuid"
    rotation: "enabled"
    
  categories:
    database:
      - name: "djinn/prod/db/master"
        rotation_days: 90
        
    api_keys:
      - name: "djinn/prod/api/plaid"
        rotation_days: 180
      - name: "djinn/prod/api/openai"
        rotation_days: 365
        
    certificates:
      - name: "djinn/prod/certs/tls"
        rotation_days: 60
        
  access_patterns:
    runtime: "IAM role based"
    development: "Temporary credentials via AWS SSO"
    ci_cd: "GitHub OIDC provider"
```

#### Container Security
```dockerfile
# Secure container configuration
FROM golang:1.25-alpine AS builder

# Run as non-root user
RUN adduser -D -g '' appuser

# Security updates
RUN apk update && apk upgrade && apk add --no-cache ca-certificates

# Build with security flags
ARG CGO_ENABLED=0
ARG GOFLAGS="-buildmode=pie -trimpath -ldflags=-linkmode=external -extldflags=-static"

# Final stage
FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Run as non-root
USER appuser

# Read-only root filesystem
# Security headers via reverse proxy
```

### 6. Security Monitoring and Incident Response

#### Security Event Monitoring
```sql
-- Security audit events (Firebase handles auth, we track API abuse)
CREATE TABLE security_events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_type TEXT NOT NULL, -- invalid_token, wrong_app_token, rate_limit_exceeded, suspicious_activity
    severity TEXT NOT NULL CHECK (severity IN ('info', 'warning', 'critical')),
    user_id UUID REFERENCES users(id), -- NULL for unauthenticated attempts
    ip_address INET,
    user_agent TEXT,
    details JSONB NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX idx_security_events_severity ON security_events(severity, created_at DESC);
CREATE INDEX idx_security_events_ip ON security_events(ip_address, created_at DESC);

```

#### Anomaly Detection (Application Layer)
```go
// Anomaly detection in Go application
type SecurityMonitor struct {
    db *sql.DB
    alertService *AlertService
}

// Check for API abuse patterns after logging security event
func (s *SecurityMonitor) CheckForAnomalies(event SecurityEvent) {
    switch event.Type {
    case "invalid_token", "wrong_app_token":
        // Check for token abuse from same IP
        var count int
        s.db.QueryRow(`
            SELECT COUNT(*) FROM security_events 
            WHERE ip_address = $1 
            AND event_type IN ('invalid_token', 'wrong_app_token')
            AND created_at > NOW() - INTERVAL '10 minutes'
        `, event.IPAddress).Scan(&count)
        
        if count > 10 {
            s.alertService.SendAlert("potential_token_abuse", event.IPAddress)
        }
        
    case "data_export":
        // Check for excessive exports
        if event.UserID != nil {
            var count int
            s.db.QueryRow(`
                SELECT COUNT(*) FROM security_events 
                WHERE user_id = $1 
                AND event_type = 'data_export'
                AND created_at > NOW() - INTERVAL '1 hour'
            `, event.UserID).Scan(&count)
            
            if count > 5 {
                s.alertService.SendAlert("excessive_exports", *event.UserID)
            }
        }
    }
}
```

#### Incident Response Plan
```yaml
incident_response:
  severity_levels:
    P0_critical:
      description: "Data breach or system compromise"
      response_time: "15 minutes"
      team: ["security_lead", "cto", "legal"]
      
    P1_high:
      description: "Active attack or vulnerability"
      response_time: "1 hour"
      team: ["security_lead", "engineering"]
      
    P2_medium:
      description: "Suspicious activity or policy violation"
      response_time: "4 hours"
      team: ["security_engineer"]
      
  playbooks:
    data_breach:
      steps:
        1: "Isolate affected systems"
        2: "Preserve evidence and logs"
        3: "Assess scope and impact"
        4: "Notify legal and compliance"
        5: "Begin user notification process"
        6: "Implement containment measures"
        7: "Document timeline and actions"
        
    account_takeover:
      steps:
        1: "Lock affected account"
        2: "Invalidate all sessions"
        3: "Notify user via verified channel"
        4: "Review recent transactions"
        5: "Force password reset"
        6: "Enable MFA requirement"
```

### 7. Compliance and Privacy

#### GDPR/CCPA Implementation
```go
// Privacy rights implementation
type PrivacyService struct {
    db *sql.DB
    encryption *EncryptionService
}

// Right to access (data export)
func (ps *PrivacyService) ExportUserData(userID string) (*UserDataExport, error) {
    export := &UserDataExport{
        RequestedAt: time.Now(),
        UserID: userID,
    }
    
    // Collect all user data
    export.Profile = ps.getUserProfile(userID)
    export.Accounts = ps.getUserAccounts(userID)
    export.Transactions = ps.getUserTransactions(userID)
    export.Categories = ps.getUserCategories(userID)
    export.Receipts = ps.getUserReceipts(userID)
    export.Preferences = ps.getUserPreferences(userID)
    
    // Generate secure download link
    export.DownloadURL = ps.generateSecureURL(export)
    
    // Audit log
    ps.logDataExport(userID)
    
    return export, nil
}

// Right to erasure (deletion)
func (ps *PrivacyService) DeleteUserData(userID string, confirmation string) error {
    // Verify deletion request
    if !ps.verifyDeletionRequest(userID, confirmation) {
        return ErrInvalidDeletionRequest
    }
    
    tx, err := ps.db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()
    
    // Anonymize financial records (keep for regulatory compliance)
    ps.anonymizeFinancialRecords(tx, userID)
    
    // Delete personal data
    ps.deletePersonalData(tx, userID)
    
    // Delete receipts and images
    ps.deleteUserReceipts(tx, userID)
    
    // Mark account as deleted
    ps.markAccountDeleted(tx, userID)
    
    // Schedule data purge from backups
    ps.schedulePurgeFromBackups(userID)
    
    return tx.Commit()
}

// Cookie consent and tracking
func (ps *PrivacyService) UpdateConsent(userID string, consent ConsentUpdate) error {
    // Store consent preferences
    _, err := ps.db.Exec(`
        INSERT INTO user_consent (user_id, consent_type, granted, ip_address, timestamp)
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT (user_id, consent_type) 
        DO UPDATE SET granted = $3, ip_address = $4, timestamp = $5
    `, userID, consent.Type, consent.Granted, consent.IPAddress, time.Now())
    
    return err
}
```

#### Privacy Compliance Requirements
| Regulation | Requirement | Implementation |
|------------|------------|---------------|
| CCPA/CPRA | Privacy rights | Data export, deletion, minimal collection |
| GDPR | EU data protection | Privacy by design, data minimization, consent |
| Privacy-First | User trust | Transparent data handling, user control |

## Consequences

### Positive
- **User Trust**: Bank-grade security builds confidence in the platform
- **Physical Security**: Biometric protection against device theft/unauthorized access
- **Compliance Ready**: Architecture supports current and future regulations
- **Scalable Security**: Simple tier-based system grows with user base
- **Defense in Depth**: Multiple security layers prevent single point of failure
- **Audit Trail**: Security event logging for API abuse detection
- **Privacy First**: Strong encryption for sensitive tokens and data isolation
- **Cost Effective**: Leverages managed services (Firebase, KMS) to reduce overhead

### Negative
- **Complexity**: Multiple security layers add development complexity
- **Performance Impact**: Encryption/decryption adds latency
- **Key Management**: Rotation and recovery procedures needed
- **User Experience**: Security measures may add friction (MFA, re-auth)
- **Cost**: Security services and compliance audits are expensive

### Risks
- **Firebase Dependency**: Authentication tied to third-party service
- **Key Compromise**: Master key breach would be catastrophic
- **Insider Threat**: Admin access needs careful monitoring
- **Zero-Day Vulnerabilities**: Unknown security flaws in dependencies
- **Social Engineering**: Users may be tricked despite technical controls

## Alternatives Considered

### Option A: Roll Custom Authentication
- **Description**: Build authentication from scratch
- **Pros**: Full control, no vendor lock-in
- **Cons**: Complex, time-consuming, security risks, maintenance burden
- **Reason for not choosing**: Small team cannot maintain secure auth system

### Option B: AWS Cognito
- **Description**: Use AWS managed authentication
- **Pros**: Well-integrated with AWS, feature-rich
- **Cons**: More complex than Firebase, AWS lock-in, steeper learning curve
- **Reason for not choosing**: Firebase offers better mobile SDKs and simpler integration

### Option C: Email/Password Authentication
- **Description**: Traditional email and password authentication
- **Pros**: Familiar to users, no third-party dependency
- **Cons**: Password management complexity, reset flows, breach risk
- **Reason for not choosing**: OAuth provides better security without password risks

## Implementation Notes

### Implementation Notes for Mobile Biometric Flow

#### Flutter Implementation
```dart
class AuthService {
  final FirebaseAuth _firebaseAuth = FirebaseAuth.instance;
  final LocalAuthentication _localAuth = LocalAuthentication();
  
  // Check if user needs authentication
  Future<AuthState> checkAuthRequired() async {
    // Check if Firebase token exists and is valid
    final user = _firebaseAuth.currentUser;
    if (user == null) {
      return AuthState.needsOAuthLogin; // First time or logged out
    }
    
    // Token exists, check if biometric needed
    final lastBackground = await getLastBackgroundTime();
    final now = DateTime.now();
    
    if (lastBackground == null || 
        now.difference(lastBackground).inMinutes > 5) {
      return AuthState.needsBiometric; // Need biometric unlock
    }
    
    return AuthState.authenticated; // Good to go
  }
  
  // Biometric authentication
  Future<bool> authenticateWithBiometric() async {
    try {
      final isAuthenticated = await _localAuth.authenticate(
        localizedReason: 'Authenticate to access your financial data',
        options: AuthenticationOptions(
          biometricOnly: false, // Allow passcode fallback
          stickyAuth: true, // Don't minimize on background
        ),
      );
      
      if (isAuthenticated) {
        await updateLastAuthTime();
        return true;
      }
    } catch (e) {
      // Biometric not available, fall back to device passcode
    }
    return false;
  }
  
  // Google/Apple sign in (only for initial auth or re-auth after 30 days)
  Future<void> signInWithGoogle() async {
    // Only called for first-time users or token expiration
    final googleUser = await GoogleSignIn().signIn();
    final googleAuth = await googleUser?.authentication;
    final credential = GoogleAuthProvider.credential(
      accessToken: googleAuth?.accessToken,
      idToken: googleAuth?.idToken,
    );
    
    await _firebaseAuth.signInWithCredential(credential);
    await enableBiometricAuth(); // Prompt to enable biometric
  }
}
```

### Migration Strategy
1. **Phase 1**: Implement Firebase Auth with OAuth only (Google, Apple) for mobile
2. **Phase 2**: Integrate biometric authentication with proper session handling
3. **Phase 3**: Implement PostgreSQL RLS and basic encryption
4. **Phase 4**: Add API gateway with rate limiting for mobile clients
5. **Phase 5**: Implement privacy features (data export, deletion)
6. **Phase 6**: Security audit focused on mobile app data aggregation

### Testing Approach
- Unit tests for all security functions
- Integration tests for auth flows
- Penetration testing before launch
- Security scanning in CI/CD pipeline
- Chaos engineering for incident response

### Monitoring and Success Metrics
- Invalid token attempts < 0.1% of API calls
- Zero security breaches
- OAuth adoption: 100% (only auth method)
- Biometric adoption > 90% of devices that support it
- Wrong app token attempts < 10 per day (tokens from other apps)
- Security scan findings: 0 critical, < 3 high
- User data isolation: 100% effective (verified via testing)
- Privacy compliance: Full CCPA/GDPR feature implementation

## References
- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
- [NIST Cybersecurity Framework](https://www.nist.gov/cyberframework)
- [Firebase Security Best Practices](https://firebase.google.com/docs/rules/basics)
- [PostgreSQL Row Level Security](https://www.postgresql.org/docs/current/ddl-rowsecurity.html)
- [GLBA Safeguards Rule](https://www.ftc.gov/business-guidance/resources/ftc-safeguards-rule-what-your-business-needs-know)
- ADR-20250812: Personal Finance Tech Stack Selection
- ADR-20250819: Data Architecture and Schema Design

## Decision Makers
- Author: Archie (System Architect)
- Reviewers: [Pending]
- Approver: [Pending]
- Date: 2025-08-19

## Revision History
- 2025-08-19: Initial comprehensive draft