# Security Layers

```mermaid
graph TB
    subgraph "Client"
        App[Mobile App]
        Token[Firebase ID Token]
        Upload[File Upload]
        Privacy[Data Privacy Controls]
    end
    
    subgraph "API Layer"
        MW[Auth Middleware]
        Val[Input Validation]
        RL[Rate Limiter]
        QC[Query Complexity Limiter]
        Quota[API Quota Enforcement]
        Scan[Malware Scanner]
        ErrH[Structured Error Handler]
    end
    
    subgraph "Business Layer"
        Auth[Authorization]
        Audit[Audit Logger]
        Idemp[Idempotency]
        OCR[OCR Data Boundary]
        Sanitize[Data Sanitization]
    end
    
    subgraph "Data Layer"
        RLS[Row Level Security]
        Encrypt[Encryption at Rest]
        Backup[PITR Backups]
        Mask[Sensitive Data Masking]
    end
    
    App --> Token
    Token --> MW
    MW --> Val
    Val --> RL
    RL --> QC
    QC --> Quota
    Quota --> Auth
    Auth --> Audit
    Auth --> Idemp
    Upload --> Scan
    Scan --> OCR
    OCR --> Sanitize
    Sanitize --> ErrH
    ErrH --> RLS
    RLS --> Encrypt
    Encrypt --> Backup
    Encrypt --> Mask
    
    MW -.-> |Verify| FB[Firebase Admin SDK]
    Audit -.-> |Stream| SIEM[SIEM/Logging]
    ErrH -.-> |Log| Telemetry[Error Telemetry]
```

## Security Layer Components

### Client Layer
- **Firebase ID Token**: Authentication token
- **Data Privacy Controls**: User-level privacy settings
- **File Upload**: Secure file submission interface

### API Layer
- **Auth Middleware**: Token validation
- **Input Validation**: Request schema and type checking
- **Rate Limiter**: Prevents abuse via request throttling
- **Query Complexity Limiter**: Protects against resource-intensive queries
- **API Quota Enforcement**: Tiered access control
- **Malware Scanner**: File upload security
- **Structured Error Handler**: Controlled error responses

### Business Layer
- **Authorization**: Role-based access control
- **Audit Logger**: Comprehensive activity tracking
- **Idempotency**: Prevents duplicate transactions
- **OCR Data Boundary**: Isolation of OCR processing
- **Data Sanitization**: Removes potentially harmful content

### Data Layer
- **Row Level Security**: Fine-grained data access
- **Encryption at Rest**: Data protection
- **Point-in-Time Recovery Backups**: Data resilience
- **Sensitive Data Masking**: Protects confidential information