# ADR-20250819: Security Architecture

## Status
Proposed

## Context
Djinn handles sensitive financial data requiring bank-grade security:
- Personal financial information (PII) and transaction history
- Bank account credentials and balances
- Receipt images potentially containing sensitive information
- Budget and spending patterns revealing lifestyle details
- Multi-user system requiring strict data isolation
- Mobile app requiring secure offline/online synchronization
- Regulatory compliance (GLBA, CCPA/CPRA, PCI DSS for future payments)
- Target market includes privacy-conscious ex-Mint users

### Constraints
- Firebase Auth already chosen for authentication (tech stack ADR)
- PostgreSQL with Row Level Security capabilities
- Mobile-first architecture with offline support
- Limited initial budget (need cost-effective solutions)
- Small team (cannot manage complex security infrastructure)
- Must maintain user trust as core value proposition

## Decision

### 1. Authentication Architecture

#### Primary Authentication: Firebase Auth
```typescript
// Authentication flow
interface AuthenticationFlow {
  // Primary methods
  methods: {
    email: "Password with email verification required",
    google: "OAuth 2.0 via Firebase",
    apple: "Sign in with Apple via Firebase", 
    biometric: "FaceID/TouchID/Fingerprint for mobile"
  },
  
  // Multi-factor authentication
  mfa: {
    required: "For premium users and sensitive operations",
    methods: ["SMS", "TOTP via authenticator apps"],
    triggers: ["Large transfers", "Account deletion", "Export data"]
  },
  
  // Session management
  sessions: {
    web: "1 hour with refresh tokens (30 days)",
    mobile: "7 days with biometric re-auth",
    api: "Short-lived JWT (15 minutes)"
  }
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
        
        // Check token freshness for sensitive operations
        if isSensitiveOperation(r) && tokenAge(decoded) > 5*time.Minute {
            return unauthorized(w, "Token too old for this operation")
        }
        
        // Set user context for RLS
        ctx = context.WithValue(ctx, "user_id", decoded.UID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

### 2. Authorization Model

#### Role-Based Access Control (RBAC) with Attribute-Based Refinements
```sql
-- User roles and permissions
CREATE TABLE roles (
    id UUID PRIMARY KEY,
    name TEXT UNIQUE NOT NULL, -- free_user, premium_user, premium_plus_user, admin
    permissions JSONB NOT NULL DEFAULT '[]'
);

CREATE TABLE user_roles (
    user_id UUID REFERENCES users(id),
    role_id UUID REFERENCES roles(id),
    granted_at TIMESTAMPTZ DEFAULT NOW(),
    granted_by UUID REFERENCES users(id),
    expires_at TIMESTAMPTZ,
    PRIMARY KEY (user_id, role_id)
);

-- Feature flags for gradual rollout
CREATE TABLE feature_flags (
    id UUID PRIMARY KEY,
    feature_name TEXT UNIQUE NOT NULL,
    enabled_for_roles TEXT[] DEFAULT '{}',
    enabled_for_users UUID[] DEFAULT '{}',
    percentage_rollout INTEGER DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT NOW()
);
```

#### GraphQL Authorization Directives
```graphql
directive @auth(requires: Role = USER) on FIELD_DEFINITION
directive @owner on FIELD_DEFINITION  
directive @premium on FIELD_DEFINITION
directive @rateLimit(max: Int!, window: String!) on FIELD_DEFINITION

type Query {
  # Public endpoints
  publicStats: PublicStats! 
  
  # Authenticated endpoints
  me: User! @auth
  myAccounts: [Account!]! @auth @owner
  
  # Premium features
  advancedAnalytics: Analytics! @auth @premium
  ocrReceipt(id: ID!): Receipt! @auth @premium @rateLimit(max: 100, window: "1h")
  
  # Admin only
  systemMetrics: SystemMetrics! @auth(requires: ADMIN)
}
```

### 3. Data Encryption Strategy

#### Encryption Layers
```sql
-- Application-level encryption for sensitive fields
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Key rotation table
CREATE TABLE encryption_keys (
    id UUID PRIMARY KEY,
    key_version INTEGER NOT NULL,
    key_data BYTEA NOT NULL, -- Encrypted with master key from KMS
    algorithm TEXT DEFAULT 'AES-256-GCM',
    created_at TIMESTAMPTZ DEFAULT NOW(),
    rotated_at TIMESTAMPTZ,
    is_active BOOLEAN DEFAULT TRUE
);

-- Encrypted fields implementation
CREATE OR REPLACE FUNCTION encrypt_sensitive(
    plain_text TEXT,
    field_type TEXT -- 'pii', 'financial', 'credential'
) RETURNS JSONB AS $$
DECLARE
    current_key RECORD;
    encrypted_data BYTEA;
    iv BYTEA;
BEGIN
    -- Get active key for field type
    SELECT * INTO current_key 
    FROM encryption_keys 
    WHERE is_active = TRUE 
    ORDER BY key_version DESC 
    LIMIT 1;
    
    -- Generate IV
    iv := gen_random_bytes(16);
    
    -- Encrypt with AES-256-GCM
    encrypted_data := pgp_sym_encrypt_bytea(
        plain_text::bytea,
        current_key.key_data,
        'compress-algo=0, cipher-algo=aes256'
    );
    
    RETURN jsonb_build_object(
        'ciphertext', encode(encrypted_data, 'base64'),
        'iv', encode(iv, 'base64'),
        'key_version', current_key.key_version,
        'encrypted_at', NOW()
    );
END;
$$ LANGUAGE plpgsql SECURITY DEFINER;
```

#### Data Classification and Encryption Requirements
| Data Type | Classification | Encryption | Storage | Retention |
|-----------|---------------|------------|---------|-----------|
| Passwords | Critical | Firebase managed | Never stored | N/A |
| Bank Credentials | Critical | Plaid token vault | Token reference only | Until disconnected |
| Account Numbers | High | AES-256 at rest | Masked (last 4) | 7 years |
| Transaction Data | Medium | TLS + DB encryption | PostgreSQL | 7 years |
| Receipt Images | Medium | S3 SSE-S3 | S3 with signed URLs | 7 years |
| Email Addresses | High | Hashed for lookup | Encrypted at rest | Until deletion |
| User Preferences | Low | DB encryption | JSONB | Until deletion |

### 4. API Security

#### Rate Limiting Strategy
```go
// Rate limiter configuration
type RateLimitConfig struct {
    // Per-user limits
    FreeUserLimits: Limits{
        RequestsPerMinute: 60,
        OCRPerDay: 10,
        ExportsPerDay: 3,
    },
    PremiumUserLimits: Limits{
        RequestsPerMinute: 300,
        OCRPerDay: 100,
        ExportsPerDay: 50,
    },
    
    // Per-endpoint limits
    EndpointLimits: map[string]Limit{
        "/graphql": {Rate: 100, Burst: 200},
        "/upload":  {Rate: 10, Burst: 20},
        "/export":  {Rate: 5, Burst: 10},
    },
    
    // Global limits
    GlobalLimits: Limits{
        RequestsPerSecond: 1000,
        ConcurrentConnections: 5000,
    }
}

// Implementation using Redis
func RateLimitMiddleware(config RateLimitConfig) func(http.Handler) http.Handler {
    limiter := redis_rate.NewLimiter(redisClient)
    
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            userID := getUserID(r.Context())
            endpoint := r.URL.Path
            
            // Check user limit
            userKey := fmt.Sprintf("rate:user:%s", userID)
            userLimit := getUserLimit(userID, config)
            
            res, err := limiter.Allow(r.Context(), userKey, userLimit)
            if err != nil || res.Allowed == 0 {
                http.Error(w, "Rate limit exceeded", 429)
                w.Header().Set("X-RateLimit-Retry-After", res.RetryAfter.String())
                return
            }
            
            // Set rate limit headers
            w.Header().Set("X-RateLimit-Limit", strconv.Itoa(userLimit.Rate))
            w.Header().Set("X-RateLimit-Remaining", strconv.Itoa(res.Remaining))
            w.Header().Set("X-RateLimit-Reset", res.ResetAfter.String())
            
            next.ServeHTTP(w, r)
        })
    }
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
FROM golang:1.21-alpine AS builder

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
-- Security audit events
CREATE TABLE security_events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_type TEXT NOT NULL, -- login_failed, permission_denied, suspicious_activity
    severity TEXT NOT NULL CHECK (severity IN ('info', 'warning', 'critical')),
    user_id UUID REFERENCES users(id),
    ip_address INET,
    user_agent TEXT,
    details JSONB NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX idx_security_events_severity ON security_events(severity, created_at DESC);
CREATE INDEX idx_security_events_user ON security_events(user_id, created_at DESC);

-- Anomaly detection triggers
CREATE OR REPLACE FUNCTION detect_anomalies() RETURNS TRIGGER AS $$
BEGIN
    -- Multiple failed login attempts
    IF NEW.event_type = 'login_failed' THEN
        PERFORM pg_notify('security_alert', json_build_object(
            'type', 'multiple_failed_logins',
            'user_id', NEW.user_id,
            'count', (
                SELECT COUNT(*) 
                FROM security_events 
                WHERE user_id = NEW.user_id 
                AND event_type = 'login_failed'
                AND created_at > NOW() - INTERVAL '10 minutes'
            )
        )::text);
    END IF;
    
    -- Unusual access patterns
    IF NEW.event_type = 'data_export' THEN
        PERFORM pg_notify('security_alert', json_build_object(
            'type', 'unusual_export',
            'user_id', NEW.user_id,
            'details', NEW.details
        )::text);
    END IF;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
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

#### Regulatory Compliance Checklist
| Regulation | Requirement | Implementation |
|------------|------------|---------------|
| GLBA | Financial data protection | Encryption, access controls, audit logs |
| CCPA/CPRA | Privacy rights | Data export, deletion, consent management |
| PCI DSS | Payment card security | No card storage, use tokenization |
| SOC 2 Type II | Security controls | Monitoring, incident response, vendor management |
| GDPR | EU data protection | Privacy by design, data minimization, consent |

## Consequences

### Positive
- **User Trust**: Bank-grade security builds confidence in the platform
- **Compliance Ready**: Architecture supports current and future regulations
- **Scalable Security**: Role-based system grows with user base
- **Defense in Depth**: Multiple security layers prevent single point of failure
- **Audit Trail**: Complete security event logging for forensics
- **Privacy First**: Strong encryption and data isolation
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

### Option C: Passwordless Only
- **Description**: Magic links and biometric only
- **Pros**: More secure, no password management
- **Cons**: Email dependency, user education needed, recovery complex
- **Reason for not choosing**: Users expect password option, especially from Mint

## Implementation Notes

### Migration Strategy
1. **Phase 1**: Implement Firebase Auth with email/password
2. **Phase 2**: Add OAuth providers (Google, Apple)
3. **Phase 3**: Enable MFA for premium users
4. **Phase 4**: Implement encryption for sensitive fields
5. **Phase 5**: Add rate limiting and monitoring
6. **Phase 6**: Security audit and penetration testing

### Testing Approach
- Unit tests for all security functions
- Integration tests for auth flows
- Penetration testing before launch
- Security scanning in CI/CD pipeline
- Chaos engineering for incident response

### Monitoring and Success Metrics
- Failed login attempts < 1% legitimate users
- Zero security breaches
- MFA adoption > 30% premium users
- Security scan findings: 0 critical, < 5 high
- Mean time to detect (MTTD) < 5 minutes
- Mean time to respond (MTTR) < 30 minutes

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