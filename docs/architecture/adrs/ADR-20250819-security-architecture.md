# ADR-20250819: Security Architecture

## Status
Accepted

## Context
Djinn handles personal financial data requiring appropriate security measures:
- Personal financial information (PII) and aggregated transaction history
- Plaid access tokens for read-only bank connections
- Receipt OCR data and imported CSV files containing transaction information
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

##### Feature Flag Naming Conventions
```go
// Standardized feature flag naming patterns for consistency
const (
    // Kill Switches - Emergency feature disabling (prefix: kill_)
    KillSwitch_OCRProcessing = "kill_ocr_processing"
    KillSwitch_AIInsights = "kill_ai_insights"
    KillSwitch_ThirdPartySync = "kill_third_party_sync"
    
    // Feature Rollouts - New feature releases (prefix: feature_)
    Feature_AdvancedAnalyticsV2 = "feature_advanced_analytics_v2"
    Feature_MultiCurrency = "feature_multi_currency"
    Feature_BudgetForecasting = "feature_budget_forecasting"
    
    // Experiments - A/B testing (prefix: exp_)
    Experiment_CheckoutFlow = "exp_checkout_flow"
    Experiment_OnboardingV2 = "exp_onboarding_v2"
    
    // Operations - Operational flags (prefix: ops_)
    Ops_MaintenanceMode = "ops_maintenance_mode"
    Ops_ReadOnlyMode = "ops_read_only_mode"
    Ops_DebugLogging = "ops_debug_logging"
    
    // Limits - Rate limiting and quotas (prefix: limit_)
    Limit_MaxReceiptsPerDay = "limit_max_receipts_per_day"
    Limit_APIRateLimit = "limit_api_rate_limit"
)
```

##### Core Feature Flag Implementation
```go
// Enhanced feature flag implementation with kill switches and progressive rollout
package features

import (
    "context"
    "log/slog"
    "sync"
    "time"
    "github.com/posthog/posthog-go"
)

type FeatureFlags struct {
    client          posthog.Client
    cache           sync.Map
    cacheTTL        time.Duration
    killSwitches    map[string]bool
    fallbackValues  map[string]interface{}
    logger          *slog.Logger
}

func NewFeatureFlags(apiKey string, logger *slog.Logger) *FeatureFlags {
    client := posthog.New(apiKey)
    
    return &FeatureFlags{
        client:     client,
        cacheTTL:   5 * time.Minute,
        logger:     logger,
        killSwitches: make(map[string]bool),
        fallbackValues: map[string]interface{}{
            // Critical feature fallbacks
            KillSwitch_OCRProcessing: false,  // Default: enabled
            KillSwitch_AIInsights: false,     // Default: enabled
            Feature_AdvancedAnalyticsV2: false, // Default: use v1
            Ops_MaintenanceMode: false,       // Default: normal operation
        },
    }
}

// IsEnabled with caching and kill switch support
func (f *FeatureFlags) IsEnabled(ctx context.Context, userID string, feature string) bool {
    // Priority 1: Check kill switches (immediate response)
    if f.isKillSwitchActive(feature) {
        f.logger.Warn("Feature disabled by kill switch",
            "feature", feature,
            "user_id", userID,
        )
        return false
    }
    
    // Priority 2: Check cached value
    if cached, ok := f.getCachedValue(feature, userID); ok {
        return cached.(bool)
    }
    
    // Priority 3: Fetch from PostHog with timeout
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()
    
    isEnabled, err := f.client.IsFeatureEnabled(
        posthog.FeatureFlagPayload{
            Key: feature,
            DistinctId: userID,
            Properties: f.getUserProperties(ctx, userID),
        },
    )
    
    if err != nil {
        f.logger.Error("Feature flag evaluation failed",
            "feature", feature,
            "user_id", userID,
            "error", err,
        )
        // Return safe fallback value
        return f.getFallbackValue(feature)
    }
    
    // Cache the result
    f.setCachedValue(feature, userID, isEnabled)
    
    return isEnabled
}

// Progressive rollout helper
func (f *FeatureFlags) GetRolloutPercentage(feature string) int {
    variant := f.client.GetFeatureFlag(
        posthog.FeatureFlagPayload{
            Key: feature + "_rollout_percentage",
        },
    )
    
    if percentage, ok := variant.(float64); ok {
        return int(percentage)
    }
    return 0
}
```

##### Progressive Rollout Strategy
```go
// Progressive rollout implementation with stages
type RolloutStage struct {
    Name       string
    Percentage int
    Criteria   map[string]interface{}
    StartDate  time.Time
    Duration   time.Duration
}

type ProgressiveRollout struct {
    flags  *FeatureFlags
    stages []RolloutStage
    logger *slog.Logger
}

func NewProgressiveRollout(flags *FeatureFlags, logger *slog.Logger) *ProgressiveRollout {
    return &ProgressiveRollout{
        flags:  flags,
        logger: logger,
        stages: []RolloutStage{
            {
                Name:       "internal_testing",
                Percentage: 0,
                Criteria:   map[string]interface{}{"email_domain": "company.com"},
                Duration:   72 * time.Hour,
            },
            {
                Name:       "beta_users",
                Percentage: 1,
                Criteria:   map[string]interface{}{"user_segment": "beta"},
                Duration:   7 * 24 * time.Hour,
            },
            {
                Name:       "canary_release",
                Percentage: 5,
                Duration:   3 * 24 * time.Hour,
            },
            {
                Name:       "gradual_rollout_10",
                Percentage: 10,
                Duration:   2 * 24 * time.Hour,
            },
            {
                Name:       "gradual_rollout_25",
                Percentage: 25,
                Duration:   2 * 24 * time.Hour,
            },
            {
                Name:       "gradual_rollout_50",
                Percentage: 50,
                Duration:   3 * 24 * time.Hour,
            },
            {
                Name:       "gradual_rollout_75",
                Percentage: 75,
                Duration:   3 * 24 * time.Hour,
            },
            {
                Name:       "general_availability",
                Percentage: 100,
                Duration:   0, // Permanent
            },
        },
    }
}

// Execute progressive rollout with monitoring
func (pr *ProgressiveRollout) ExecuteRollout(ctx context.Context, feature string) error {
    for _, stage := range pr.stages {
        pr.logger.Info("Starting rollout stage",
            "feature", feature,
            "stage", stage.Name,
            "percentage", stage.Percentage,
        )
        
        // Update PostHog configuration
        if err := pr.updateRolloutPercentage(feature, stage); err != nil {
            return pr.handleRolloutError(feature, stage, err)
        }
        
        // Monitor metrics during stage
        if err := pr.monitorStage(ctx, feature, stage); err != nil {
            // Automatic rollback on error
            return pr.rollback(feature, stage, err)
        }
        
        // Wait for stage duration
        if stage.Duration > 0 {
            select {
            case <-time.After(stage.Duration):
                continue
            case <-ctx.Done():
                return ctx.Err()
            }
        }
    }
    
    return nil
}
```

##### Kill Switch Implementation
```go
// Kill switch manager for emergency feature disabling
type KillSwitchManager struct {
    flags         *FeatureFlags
    logger        *slog.Logger
    notifier      AlertNotifier
    activeKills   sync.Map
}

func (ksm *KillSwitchManager) ActivateKillSwitch(feature string, reason string) error {
    ksm.logger.Error("KILL SWITCH ACTIVATED",
        "feature", feature,
        "reason", reason,
        "timestamp", time.Now(),
    )
    
    // Immediately disable in PostHog
    err := ksm.flags.client.ReloadFeatureFlags()
    if err != nil {
        // Even if PostHog fails, activate locally
        ksm.logger.Error("PostHog update failed, using local kill switch", "error", err)
    }
    
    // Set local kill switch (immediate effect)
    ksm.activeKills.Store(feature, true)
    ksm.flags.killSwitches[feature] = true
    
    // Notify operations team
    ksm.notifier.SendAlert(Alert{
        Severity: "CRITICAL",
        Title:    fmt.Sprintf("Kill Switch Activated: %s", feature),
        Message:  reason,
        Feature:  feature,
    })
    
    // Log to audit trail
    ksm.logKillSwitchEvent(feature, reason, "ACTIVATED")
    
    return nil
}

// Automatic kill switch based on error rates
func (ksm *KillSwitchManager) MonitorErrorRates(ctx context.Context) {
    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            for feature, threshold := range ksm.getMonitoredFeatures() {
                errorRate := ksm.calculateErrorRate(feature)
                
                if errorRate > threshold {
                    ksm.ActivateKillSwitch(feature, 
                        fmt.Sprintf("Error rate exceeded threshold: %.2f%% > %.2f%%", 
                            errorRate, threshold))
                }
            }
        case <-ctx.Done():
            return
        }
    }
}
```

##### GraphQL Integration Example
```go
// GraphQL resolver with feature flags and kill switches
func (r *Resolver) ProcessReceipt(ctx context.Context, input ReceiptInput) (*Receipt, error) {
    userID := ctx.Value("user_id").(string)
    
    // Check kill switch first
    if r.features.IsEnabled(ctx, userID, KillSwitch_OCRProcessing) {
        return nil, fmt.Errorf("receipt processing is temporarily disabled")
    }
    
    // Check feature flag for new OCR engine
    if r.features.IsEnabled(ctx, userID, Feature_OCREngineV2) {
        return r.processReceiptV2(ctx, input)
    }
    
    // Fallback to stable version
    return r.processReceiptV1(ctx, input)
}

// A/B testing example
func (r *Resolver) GetDashboard(ctx context.Context) (*Dashboard, error) {
    userID := ctx.Value("user_id").(string)
    
    // Get experiment variant
    variant := r.features.GetVariant(ctx, userID, Experiment_DashboardLayout)
    
    switch variant {
    case "control":
        return r.getClassicDashboard(ctx)
    case "variant_a":
        return r.getModernDashboard(ctx)
    case "variant_b":
        return r.getMinimalDashboard(ctx)
    default:
        return r.getClassicDashboard(ctx)
    }
}
```

##### GraphQL Directive Implementation (Minimal)
```go
// Custom GraphQL directives - only essential security controls
package directives

import (
    "context"
    "fmt"
    "github.com/99designs/gqlgen/graphql"
)

// @feature directive - checks PostHog feature flag
// This is the primary way to control feature rollout
func Feature(ctx context.Context, obj interface{}, next graphql.Resolver, flag string) (interface{}, error) {
    userID := ctx.Value("user_id").(string)
    features := ctx.Value("features").(*FeatureFlags)
    
    // Check PostHog feature flag (includes kill switch logic internally)
    if !features.IsEnabled(ctx, userID, flag) {
        return nil, fmt.Errorf("feature '%s' is not enabled for your account", flag)
    }
    
    return next(ctx)
}

// @premium directive - checks subscription tier
// Simple business logic check at the API level
func Premium(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
    userID := ctx.Value("user_id").(string)
    
    // Check user's subscription tier
    user, err := getUserFromDB(ctx, userID)
    if err != nil {
        return nil, err
    }
    
    if user.SubscriptionTier != "premium" && user.SubscriptionTier != "enterprise" {
        return nil, fmt.Errorf("this feature requires a premium subscription")
    }
    
    return next(ctx)
}

// Note: @rateLimit directive will be handled in a separate Rate Limiting ADR
```

##### GraphQL Schema with Security Directives
```graphql
# Directive definitions - focused on security and feature control
directive @feature(flag: String!) on FIELD_DEFINITION
directive @premium on FIELD_DEFINITION  

type Query {
  # Standard features - always available
  me: User!
  accounts: [Account!]!
  transactions: [Transaction!]!
  
  # Feature-flagged functionality (controlled by PostHog)
  # Kill switches are handled internally via PostHog flags
  aiInsights: AIInsights @feature(flag: "feature_ai_insights")
  budgetForecasting: Forecast @feature(flag: "feature_budget_forecasting")
  
  # Premium features (subscription check)
  advancedAnalytics: Analytics @premium
  unlimitedAccounts: [Account!]! @premium
  
  # Experimental features (PostHog handles the rollout logic)
  experimentalOCR: OCRResult @feature(flag: "exp_ocr_v2")
  
  # Regular endpoints (kill switches and rate limiting handled internally)
  processReceipt(input: ReceiptInput): Receipt 
  generateReport: Report
  exportData: DataExport
}

type Mutation {
  # User preferences handled internally, not via directives
  updatePreferences(input: PreferencesInput!): User!
  
  # Feature usage tracking
  trackFeatureUsage(feature: String!): Boolean!
}
```

##### Internal Kill Switch & Opt-In Handling
```go
// Kill switches and opt-in logic handled in resolver, not as directives
func (r *Resolver) ProcessReceipt(ctx context.Context, input ReceiptInput) (*Receipt, error) {
    userID := ctx.Value("user_id").(string)
    
    // Kill switch check (internal, not exposed in GraphQL schema)
    if r.features.IsEnabled(ctx, "system", KillSwitch_OCRProcessing) {
        return nil, fmt.Errorf("receipt processing is temporarily unavailable")
    }
    
    // User opt-in check (if needed, handled internally)
    user, _ := r.getUserByID(ctx, userID)
    if user.BetaFeaturesEnabled {
        // Use experimental version if opted in AND feature flag enabled
        if r.features.IsEnabled(ctx, userID, Feature_OCREngineV2) {
            return r.processReceiptV2(ctx, input)
        }
    }
    
    // Default to stable version
    return r.processReceiptV1(ctx, input)
}

// Client handles UI-level opt-in preferences
// The client can check user.preferences.betaFeatures and show/hide UI accordingly
// No need for GraphQL directives for this
```

##### Directive Registration in GraphQL Server
```go
// server.go - Wire up directives with gqlgen
package main

import (
    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
)

func NewGraphQLServer(features *FeatureFlags, rateLimiter *RateLimiter) *handler.Server {
    config := generated.Config{
        Resolvers: &resolver{
            features: features,
        },
    }
    
    // Register custom directives
    config.Directives.Feature = directives.Feature
    config.Directives.RequiresOptIn = directives.RequiresOptIn
    config.Directives.KillSwitch = directives.KillSwitch
    config.Directives.Premium = directives.Premium
    config.Directives.RateLimit = directives.RateLimit
    
    srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))
    
    // Add middleware to inject dependencies into context
    srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
        ctx = context.WithValue(ctx, "features", features)
        ctx = context.WithValue(ctx, "rateLimiter", rateLimiter)
        return next(ctx)
    })
    
    return srv
}
```

##### User Opt-In Management
```go
// User preferences stored in database, not PostHog
type UserPreferences struct {
    UserID              string
    BetaFeaturesEnabled bool      // User explicitly opted into beta
    EmailNotifications  bool      
    DataCollection      bool      // GDPR compliance
    MarketingEmails     bool
    UpdatedAt           time.Time
}

// Opt-in mutation resolver
func (r *mutationResolver) UpdateBetaPreferences(ctx context.Context, enableBeta bool) (*User, error) {
    userID := ctx.Value("user_id").(string)
    
    // Update user preferences in database
    err := r.db.UpdateUserPreferences(ctx, userID, map[string]interface{}{
        "beta_features_enabled": enableBeta,
        "updated_at": time.Now(),
    })
    
    if err != nil {
        return nil, err
    }
    
    // Track opt-in event in PostHog for analytics
    r.features.client.Capture(posthog.Capture{
        DistinctId: userID,
        Event: "beta_features_toggled",
        Properties: posthog.NewProperties().
            Set("enabled", enableBeta).
            Set("timestamp", time.Now()),
    })
    
    return r.getUserByID(ctx, userID)
}
```

#### Why PostHog?
- **Free tier**: 1M feature flag requests/month (plenty for MVP)
- **Analytics included**: Track feature adoption automatically
- **A/B testing**: Built-in experimentation capabilities
- **Easy integration**: SDKs for Go, Flutter, and web
- **Self-host option**: Can migrate to self-hosted if needed
- **Kill switches**: Instant feature disabling for emergencies
- **Progressive rollout**: Gradual feature deployment with monitoring
- **Alternative options**: Flagsmith (50k requests free), Unleash (self-hosted free)

#### Feature Flag Best Practices
```yaml
conventions:
  naming:
    - Use lowercase with underscores
    - Prefix with category (kill_, feature_, exp_, ops_, limit_)
    - Be descriptive but concise
    - Include version suffix for iterations (_v2, _v3)
  
  lifecycle:
    - Create flag before feature development
    - Test with internal users first
    - Progressive rollout with monitoring
    - Clean up after 100% rollout + 30 days
    - Document in ADR for significant features
  
  safety:
    - Always provide fallback values
    - Implement timeout for flag evaluation (3s max)
    - Cache flags locally with TTL
    - Monitor error rates and auto-disable if needed
    - Log all flag evaluations for audit trail
  
  rollout_stages:
    1_internal: "email contains @company.com"
    2_beta: "user_segment = beta OR organization_tier = enterprise"
    3_canary: "5% random users"
    4_gradual: "10% -> 25% -> 50% -> 75% -> 100%"
    5_general: "100% with monitoring for 7 days"
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
| Imported Files | Medium | MinIO encryption | MinIO (S3) | 7 days |
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