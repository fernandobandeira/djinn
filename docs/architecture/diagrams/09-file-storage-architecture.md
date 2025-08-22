# File Storage and Data Import Architecture

## Context
Secure, compliant file storage for financial data imports with automated processing and lifecycle management.

## Architecture Diagram

```mermaid
graph TD
    %% User Interaction
    A[User Device] -->|Authenticated Upload| B{API Gateway}
    
    %% Security Layers
    B -->|Validate Token| C[Firebase Authentication]
    
    %% File Processing Pipeline
    B -->|File Upload| D[MinIO Object Storage]
    D -->|Move to Processing| E[Imports-Pending Bucket]
    
    %% Security and Validation
    E -->|File Scan| F[ClamAV Antivirus]
    F -->|Clean File| G[Imports-Processing Bucket]
    F -->|Malware Detected| H[Quarantine/Delete]
    
    %% Background Processing
    G -->|Parse File| I[Temporal Workers]
    I -->|Validate Data| J[Data Validation Service]
    J -->|No Duplicates| K[PostgreSQL Database]
    J -->|Duplicates Found| L[User Review Queue]
    
    %% Lifecycle Management
    D -->|Automatic Cleanup| M[Lifecycle Policies]
    M -->|Expire Pending Files| N[1 Day]
    M -->|Archive Processed| O[7 Days]
    
    %% User Notifications
    I -->|Processing Result| P[Notification Service]
    P -->|Email/Push| Q[User Notification]
    
    %% Compliance and Audit
    R[Audit Log Database] -->|Log Events| S{Security Monitoring}
    
    %% Styling
    classDef storage fill:#4A90E2,color:white;
    classDef security fill:#F5A623,color:white;
    classDef processing fill:#50E3C2,color:black;
    
    class D,E,G,M storage
    class B,C,F,H,S security
    class I,J,K,L,P processing
```

## Key Components

### 1. Storage Architecture
- **MinIO Object Storage**: S3-compatible storage
- **Buckets**:
  - `imports-pending`: Initial file upload
  - `imports-processing`: Active file processing
  - `imports-archive`: Processed file retention

### 2. Security Layers
- Firebase Authentication
- File type and size validation
- ClamAV malware scanning
- User-specific access policies

### 3. Processing Pipeline
- Temporal workers for asynchronous processing
- Data validation and deduplication
- Background import to PostgreSQL
- User review for complex imports

### 4. Lifecycle Management
- Automatic file expiration
  - Pending files: 1 day
  - Processed files: 7 days
- Audit logging for all file operations

## Compliance and Monitoring
- Comprehensive audit trail
- Security event monitoring
- GDPR and financial regulation compliance

## Deployment Context
- Deployed on Hetzner VPS
- Docker containerized (MinIO, ClamAV)
- Minimal resource footprint (~700MB RAM)

## Future Migration Path
Potential cloud storage providers:
1. Cloudflare R2
2. Backblaze B2
3. AWS S3 Intelligent-Tiering

## Risks and Mitigations
- Service failure: Automatic restart, health checks
- Storage exhaustion: User quotas, aggressive cleanup
- False-positive malware detection: Whitelisting, manual review
