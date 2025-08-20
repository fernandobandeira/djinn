# ADR-20250820: File Storage for Data Imports and Processing

## Status
Accepted

## Context
While Djinn uses on-device OCR for receipt processing (images never leave the device), we still need robust server-side file storage for:
- CSV file imports (bank statements, transaction exports)
- Manual data imports (QIF, OFX, other financial formats)
- Bulk data processing workflows
- Temporary file processing before database storage
- User-initiated data exports
- Potential future needs for document storage

### Requirements
- Secure storage for uploaded CSV and financial data files
- Support for various financial data formats (CSV, QIF, OFX, JSON, XLSX)
- Temporary storage with automatic cleanup
- Compliance with financial regulations (GLBA, CCPA)
- Audit trail for all file operations
- Cost-effective storage solution for MVP
- Integration with data import/export workflows
- Malware protection for uploaded files
- User isolation and access control

### Constraints
- MVP budget: €3.79/month Hetzner VPS with 40GB storage
- Files are temporary (processing only)
- Must support various file formats
- Must maintain audit trail for compliance
- Storage should be automatically cleaned after processing
- No public access to uploaded files
- Must be S3-compatible for future migration

## Decision

### 1. Storage Architecture

#### MinIO Object Storage (MVP Phase)
```yaml
storage_architecture:
  platform: MinIO (S3-compatible)
  deployment: Docker container on Hetzner VPS
  
  bucket_structure:
    imports-pending: "Temporary uploads awaiting processing"
    imports-processing: "Files being processed"
    imports-archive: "Processed files (7-day retention)"
    exports-temp: "User-requested exports (24-hour retention)"
    
  security_features:
    - Server-side encryption (SSE-C)
    - Policy-based access control
    - Audit logging
    - Malware scanning via ClamAV integration
    - File type validation
    - Size limits enforcement
    
  lifecycle:
    upload: "Store in imports-pending with validation"
    scan: "ClamAV malware scan before acceptance"
    process: "Parse and import to database"
    archive: "Move to imports-archive after processing"
    expire: "Auto-delete after retention period"
    
  capacity:
    total_storage: 40GB (Hetzner CX22)
    system_reserved: 10GB
    minio_data: 15GB (reduced from 20GB since no receipt images)
    database: 10GB (increased for transaction data)
    application: 5GB
    
  resource_usage:
    minio_memory: ~400MB (reduced usage)
    clamav_memory: ~300MB
    available_ram: 3.3GB for app + PostgreSQL
```

#### MinIO Configuration
```yaml
# docker-compose.yml for MinIO + ClamAV
version: '3.8'

services:
  minio:
    image: minio/minio:latest
    container_name: djinn-minio
    command: server /data --console-address ":9001"
    environment:
      MINIO_ROOT_USER: ${MINIO_ROOT_USER}
      MINIO_ROOT_PASSWORD: ${MINIO_ROOT_PASSWORD}
      MINIO_BROWSER: "off"  # Disable web console for security
      MINIO_CACHE_DRIVES: "/cache"
      MINIO_CACHE_QUOTA: 80
      MINIO_CACHE_WATERMARK_LOW: 70
      MINIO_CACHE_WATERMARK_HIGH: 90
    volumes:
      - /var/djinn/minio/data:/data
      - /var/djinn/minio/cache:/cache
    ports:
      - "127.0.0.1:9000:9000"  # API only on localhost
    restart: unless-stopped
    mem_limit: 400m
    
  clamav:
    image: clamav/clamav:stable
    container_name: djinn-clamav
    volumes:
      - /var/djinn/clamav:/var/lib/clamav
    restart: unless-stopped
    mem_limit: 300m
```

### 2. Data Import Pipeline

#### Import Flow
```yaml
data_import_flow:
  1_file_upload:
    endpoint: POST /api/v1/imports/upload
    authentication: Firebase Auth token
    max_file_size: 50MB  # Larger for bulk imports
    accepted_formats: 
      - csv
      - qif
      - ofx
      - json
      - xlsx (with pandas processing)
    
  2_server_processing:
    - Validate file type and size
    - Generate unique import_id (UUIDv7)
    - Store in imports-pending bucket
    - Create import job record
    - Return import_id for tracking
    
  3_background_processing:
    - Move to imports-processing
    - Parse file based on format
    - Validate data schema
    - Detect duplicates
    - Apply user's categorization rules
    - Store in database (batched inserts)
    - Generate import report
    
  4_user_review:
    - Present import summary
    - Show detected duplicates
    - Allow bulk categorization
    - Confirm or rollback import
    
  5_completion:
    user_confirms:
      - Commit transaction batch
      - Move file to imports-archive
      - Schedule deletion (7 days)
    user_cancels:
      - Rollback database transaction
      - Delete import file immediately
```

#### Secure Import Pipeline
```go
// Secure file import with validation and processing
type SecureDataImporter struct {
    minioClient *minio.Client
    clamav      *clamav.Client
    maxFileSize int64  // 50MB for bulk imports
    processors  map[string]FileProcessor
}

type FileProcessor interface {
    Validate(io.Reader) error
    Parse(io.Reader) ([]Transaction, error)
    DetectFormat([]byte) bool
}

func (i *SecureDataImporter) ImportFile(ctx context.Context, userID string, file io.Reader, filename string) (*ImportResult, error) {
    // 1. Detect file format
    buffer := make([]byte, 8192)
    n, err := file.Read(buffer)
    if err != nil {
        return nil, fmt.Errorf("failed to read file header: %w", err)
    }
    
    var processor FileProcessor
    for format, proc := range i.processors {
        if proc.DetectFormat(buffer[:n]) {
            processor = proc
            break
        }
    }
    
    if processor == nil {
        return nil, fmt.Errorf("unsupported file format")
    }
    
    // 2. Create temp file for scanning
    tmpFile, err := os.CreateTemp("", "import-*")
    if err != nil {
        return nil, err
    }
    defer os.Remove(tmpFile.Name())
    
    // 3. Write to temp file and check size
    multiReader := io.MultiReader(bytes.NewReader(buffer[:n]), file)
    size, err := io.Copy(tmpFile, multiReader)
    if err != nil {
        return nil, err
    }
    
    if size > i.maxFileSize {
        return nil, fmt.Errorf("file too large: %d bytes", size)
    }
    
    // 4. Scan with ClamAV
    scanResult, err := i.clamav.ScanFile(tmpFile.Name())
    if err != nil {
        return nil, fmt.Errorf("malware scan failed: %w", err)
    }
    
    if scanResult.Status != "OK" {
        return nil, fmt.Errorf("malware detected: %s", scanResult.Description)
    }
    
    // 5. Upload to MinIO with encryption
    objectName := fmt.Sprintf("%s/%s/%s", userID, time.Now().Format("2006-01-02"), filename)
    _, err = i.minioClient.FPutObject(
        ctx,
        "imports-pending",
        objectName,
        tmpFile.Name(),
        minio.PutObjectOptions{
            ContentType: "application/octet-stream",
            UserMetadata: map[string]string{
                "user-id":    userID,
                "upload-time": time.Now().Format(time.RFC3339),
                "scan-status": "clean",
            },
        },
    )
    
    if err != nil {
        return nil, fmt.Errorf("failed to upload to storage: %w", err)
    }
    
    // 6. Process file
    tmpFile.Seek(0, 0)
    transactions, err := processor.Parse(tmpFile)
    if err != nil {
        return nil, fmt.Errorf("failed to parse file: %w", err)
    }
    
    return &ImportResult{
        ImportID:     generateImportID(),
        Transactions: transactions,
        FileLocation: objectName,
        Status:       "pending_review",
    }, nil
}
```

### 3. Security and Access Control

#### Access Patterns
```yaml
access_control:
  internal_only:
    description: "No public access to import files"
    implementation: "Files served only through authenticated API"
    
  user_access:
    pattern: "Users can only access their own imports"
    authentication: "Firebase Auth token required"
    authorization: "User ID from token must match file owner"
    
  api_endpoints:
    upload: "POST /api/v1/imports/upload"
    status: "GET /api/v1/imports/{import_id}/status"
    download: "GET /api/v1/exports/{export_id}/download"
    list: "GET /api/v1/imports"
    
  file_serving:
    method: "API generates presigned URLs"
    no_direct_urls: "Never expose filesystem paths"
    temporary_access: "Presigned URLs expire quickly"
```

#### MinIO Access Policies
```go
// User-specific bucket policy for isolation
func createUserPolicy(userID string) string {
    return fmt.Sprintf(`{
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Action": [
                    "s3:GetObject",
                    "s3:DeleteObject"
                ],
                "Resource": [
                    "arn:aws:s3:::imports-*/%s/*",
                    "arn:aws:s3:::exports-*/%s/*"
                ]
            },
            {
                "Effect": "Allow",
                "Action": ["s3:ListBucket"],
                "Resource": ["arn:aws:s3:::imports-*", "arn:aws:s3:::exports-*"],
                "Condition": {
                    "StringLike": {
                        "s3:prefix": ["%s/*"]
                    }
                }
            }
        ]
    }`, userID, userID, userID)
}

// Secure file download with presigned URL
func (h *ExportHandler) GetExport(w http.ResponseWriter, r *http.Request) {
    // 1. Authenticate user
    userID, err := h.auth.ValidateToken(r)
    if err != nil {
        http.Error(w, "Unauthorized", 401)
        return
    }
    
    // 2. Get export ID and verify ownership
    exportID := chi.URLParam(r, "exportID")
    export, err := h.db.GetExport(exportID)
    if err != nil || export.UserID != userID {
        http.Error(w, "Not found", 404)
        return
    }
    
    // 3. Generate presigned URL (valid for 1 hour)
    presignedURL, err := h.minioClient.PresignedGetObject(
        r.Context(),
        "exports-temp",
        export.ObjectName,
        1*time.Hour,
        url.Values{},
    )
    if err != nil {
        http.Error(w, "Internal error", 500)
        return
    }
    
    // 4. Return download URL
    json.NewEncoder(w).Encode(map[string]string{
        "download_url": presignedURL.String(),
        "expires_in": "3600",
    })
}
```

### 4. Retention and Lifecycle Management

#### MinIO Lifecycle Policies
```json
{
  "Rules": [
    {
      "ID": "delete-pending-imports",
      "Status": "Enabled",
      "Prefix": "",
      "Expiration": {
        "Days": 1
      }
    },
    {
      "ID": "delete-processing-imports",
      "Status": "Enabled",
      "Prefix": "",
      "Expiration": {
        "Hours": 6
      }
    },
    {
      "ID": "delete-archived-imports",
      "Status": "Enabled",
      "Prefix": "",
      "Expiration": {
        "Days": 7
      }
    },
    {
      "ID": "delete-temp-exports",
      "Status": "Enabled",
      "Prefix": "",
      "Expiration": {
        "Hours": 24
      }
    }
  ]
}
```

```go
// Set up lifecycle policies in MinIO
func setupLifecyclePolicies(client *minio.Client) error {
    policies := map[string]lifecycle.Configuration{
        "imports-pending": {
            Rules: []lifecycle.Rule{
                {
                    ID:     "delete-old-pending",
                    Status: "Enabled",
                    Expiration: lifecycle.Expiration{
                        Days: 1,
                    },
                },
            },
        },
        "imports-archive": {
            Rules: []lifecycle.Rule{
                {
                    ID:     "delete-old-archives",
                    Status: "Enabled",
                    Expiration: lifecycle.Expiration{
                        Days: 7,
                    },
                },
            },
        },
        "exports-temp": {
            Rules: []lifecycle.Rule{
                {
                    ID:     "delete-old-exports",
                    Status: "Enabled",
                    Expiration: lifecycle.Expiration{
                        Hours: 24,
                    },
                },
            },
        },
    }
    
    for bucket, policy := range policies {
        err := client.SetBucketLifecycle(context.Background(), bucket, &policy)
        if err != nil {
            return fmt.Errorf("failed to set lifecycle for %s: %w", bucket, err)
        }
    }
    
    return nil
}
```

### 5. Database Schema

```sql
-- Import job tracking
CREATE TABLE import_jobs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    
    -- File metadata
    original_filename TEXT NOT NULL,
    stored_path TEXT NOT NULL, -- MinIO object path
    file_size INTEGER NOT NULL,
    file_format TEXT NOT NULL,
    file_hash TEXT NOT NULL, -- For duplicate detection
    
    -- Processing status
    status TEXT NOT NULL CHECK (status IN ('pending', 'processing', 'review', 'completed', 'failed', 'cancelled')),
    
    -- Import results
    total_rows INTEGER,
    imported_count INTEGER,
    duplicate_count INTEGER,
    error_count INTEGER,
    import_summary JSONB,
    
    -- Timestamps
    uploaded_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    processing_started_at TIMESTAMPTZ,
    processing_completed_at TIMESTAMPTZ,
    user_confirmed_at TIMESTAMPTZ,
    expires_at TIMESTAMPTZ,
    
    -- Audit
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Export job tracking
CREATE TABLE export_jobs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    
    -- Export parameters
    export_type TEXT NOT NULL, -- 'transactions', 'full_backup', etc
    export_format TEXT NOT NULL, -- 'csv', 'json', 'qif', 'pdf'
    date_range_start DATE,
    date_range_end DATE,
    filters JSONB,
    
    -- File metadata
    stored_path TEXT, -- MinIO object path
    file_size INTEGER,
    download_count INTEGER DEFAULT 0,
    
    -- Status
    status TEXT NOT NULL CHECK (status IN ('pending', 'processing', 'completed', 'failed')),
    
    -- Timestamps
    requested_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMPTZ,
    expires_at TIMESTAMPTZ,
    last_downloaded_at TIMESTAMPTZ,
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Indexes
CREATE INDEX idx_import_jobs_user_id ON import_jobs(user_id);
CREATE INDEX idx_import_jobs_status ON import_jobs(status);
CREATE INDEX idx_import_jobs_expires_at ON import_jobs(expires_at) WHERE expires_at IS NOT NULL;
CREATE INDEX idx_export_jobs_user_id ON export_jobs(user_id);
CREATE INDEX idx_export_jobs_expires_at ON export_jobs(expires_at) WHERE expires_at IS NOT NULL;

-- Audit table
CREATE TABLE file_audit_log (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    action TEXT NOT NULL,
    file_type TEXT,
    file_name TEXT,
    details JSONB,
    ip_address INET,
    user_agent TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

### 6. Monitoring and Security

```yaml
minio_monitoring:
  metrics_endpoint: "http://localhost:9000/minio/v2/metrics/cluster"
  
  key_metrics:
    - minio_bucket_usage_total_bytes
    - minio_bucket_objects_total
    - minio_s3_requests_total
    - minio_s3_requests_errors_total
    - minio_inter_node_traffic_recv_bytes
    
  security_monitoring:
    audit_logs:
      enabled: true
      targets:
        - type: "webhook"
          endpoint: "http://localhost:8080/audit"
        - type: "postgresql"
          dsn: "${DATABASE_URL}"
      
    alerts:
      failed_auth_attempts:
        threshold: ">10 in 5 minutes"
        action: "Block IP, alert team"
        
      malware_detected:
        action: "Quarantine file, alert team, log incident"
        
      unusual_import_size:
        threshold: ">40MB single file"
        action: "Flag for review, notify user"
        
      bulk_imports:
        threshold: ">5 imports/hour from single user"
        action: "Rate limit, investigate"
        
  clamav_monitoring:
    database_updates: "Every 2 hours"
    scan_performance: "Track avg scan time"
    detection_rate: "Log all detections"
```

```go
// MinIO audit webhook handler
func (h *AuditHandler) HandleMinIOAudit(w http.ResponseWriter, r *http.Request) {
    var auditEntry MinIOAuditEntry
    if err := json.NewDecoder(r.Body).Decode(&auditEntry); err != nil {
        http.Error(w, "Bad request", 400)
        return
    }
    
    // Log to database for compliance
    h.db.LogAuditEntry(auditEntry)
    
    // Check for security events
    if auditEntry.API.Name == "GetObject" && auditEntry.API.Status == "Forbidden" {
        h.securityMonitor.RecordFailedAccess(auditEntry.RemoteHost)
    }
    
    // Monitor for unusual patterns
    if auditEntry.API.Name == "PutObject" {
        size := auditEntry.API.RequestHeader.Get("Content-Length")
        if size > 40*1024*1024 { // 40MB
            h.alerts.LargeFileUpload(auditEntry.UserInfo.UserName, size)
        }
    }
    
    w.WriteHeader(200)
}
```

### 7. Export Functionality

```yaml
data_export:
  supported_formats:
    csv:
      - Transactions with all fields
      - Categories summary
      - Budget vs actual report
      - Monthly/yearly statements
    json:
      - Full data backup
      - Selective export with filters
    qif:
      - Quicken Interchange Format
      - For migration to other tools
    pdf:
      - Formatted statements
      - Tax reports
      - Receipt summaries
    
  export_flow:
    1_request:
      endpoint: POST /api/v1/exports
      parameters:
        - type: transactions|categories|full_backup
        - format: csv|json|qif|pdf
        - date_range: optional
        - filters: optional category/merchant filters
    
    2_processing:
      - Create export job record
      - Generate file in background worker
      - Apply user's timezone and formatting
      - Compress if > 10MB
      - Encrypt sensitive data
    
    3_storage:
      - Store in exports-temp bucket
      - Set 24-hour expiration
      - Generate unique download token
    
    4_notification:
      - Send email/push with download link
      - Include expiration warning
    
    5_download:
      - Validate token and user
      - Generate presigned URL
      - Track download count
      - Auto-delete after expiry
```

### 8. Migration Path from MinIO to Cloud Storage

```yaml
future_migration:
  trigger: ">1000 active users or >100GB monthly processing"
  
  target_options:
    option_1:
      service: "Cloudflare R2"
      cost: "$15/TB/month (no egress fees)"
      migration: "MinIO Mirror or rclone sync"
      benefits: "No bandwidth charges, global edge"
      
    option_2:
      service: "Backblaze B2"
      cost: "$6/TB/month + $10/TB egress"
      migration: "S3 compatible, easy switch"
      benefits: "Lowest storage cost"
      
    option_3:
      service: "AWS S3 Intelligent-Tiering"
      cost: "$23/TB/month + monitoring"
      migration: "Native S3, full feature set"
      benefits: "Auto-tiering, enterprise features"
      
    option_4:
      service: "Distributed MinIO"
      deployment: "Multiple Hetzner servers"
      cost: "€7.59/server/month"
      benefits: "Stay self-hosted, full control"
      
  migration_strategy:
    preparation:
      - Enable MinIO gateway mode
      - Set up bucket replication
      - Configure tiering policies
      - Test with staging bucket
      
    execution:
      1_replicate: "Enable real-time replication to cloud"
      2_verify: "Validate all objects replicated"
      3_switch: "Update application endpoints"
      4_monitor: "Watch for issues for 48 hours"
      5_cleanup: "Decommission local MinIO"
      
    rollback_plan:
      - Keep local MinIO running for 7 days
      - Maintain bidirectional sync
      - Quick switch back if issues
```

```bash
# Migration script using MinIO client
#!/bin/bash

# Set up mirror to cloud storage (example with R2)
mc alias set local http://localhost:9000 $MINIO_ACCESS_KEY $MINIO_SECRET_KEY
mc alias set r2 https://r2.cloudflarestorage.com $R2_ACCESS_KEY $R2_SECRET_KEY

# Mirror buckets to cloud
mc mirror --watch local/imports-archive r2/imports-archive
mc mirror --watch local/exports-temp r2/exports-temp

# Verify migration
mc diff local/imports-archive r2/imports-archive

# Update application config
sed -i 's|http://localhost:9000|https://r2.cloudflarestorage.com|g' /etc/djinn/config.yaml

# Restart application
systemctl restart djinn-api
```

## Consequences

### Positive
- **S3 Compatibility**: MinIO provides full S3 API compatibility for easy migration
- **Enterprise Security**: Built-in encryption, access policies, and audit logging
- **Automated Lifecycle**: Policies handle retention and cleanup automatically
- **Malware Protection**: ClamAV integration prevents malicious file uploads
- **Cost Effective**: Runs on existing VPS with minimal resource usage
- **User Isolation**: Policy-based access control ensures data privacy
- **Format Flexibility**: Supports multiple import/export formats
- **Professional Grade**: Production-ready from day one

### Negative
- **Additional Complexity**: MinIO and ClamAV add operational overhead
- **Memory Usage**: ~700MB RAM for storage stack
- **Limited Capacity**: 15GB storage limit on MVP VPS
- **No Geo-Redundancy**: Single location storage
- **Learning Curve**: Team needs MinIO operational knowledge

### Risks and Mitigations
- **Risk**: MinIO service failure affects all import/export operations
  - **Mitigation**: Health checks, automatic restart, graceful degradation
  
- **Risk**: Storage exhaustion from large imports
  - **Mitigation**: Quotas per user, aggressive cleanup policies, monitoring
  
- **Risk**: ClamAV false positives blocking legitimate files
  - **Mitigation**: Whitelist known formats, manual review option
  
- **Risk**: Performance impact from scanning large files
  - **Mitigation**: Async scanning, queue management, timeout limits

## Alternatives Considered

### Option A: Direct Filesystem Storage
- **Description**: Store files directly on VPS filesystem
- **Pros**: Simpler, no additional services
- **Cons**: No S3 compatibility, path traversal risks, harder lifecycle management
- **Reason for not choosing**: Security risks and lack of features

### Option B: Cloud Storage from Start
- **Description**: Use S3/R2/B2 from day one
- **Pros**: Unlimited storage, managed service
- **Cons**: Additional cost ($5-20/month), external dependency, data sovereignty
- **Reason for not choosing**: MinIO on VPS is free and sufficient for MVP

### Option C: Database BLOB Storage
- **Description**: Store files in PostgreSQL
- **Pros**: Single system, transactional consistency
- **Cons**: Database bloat, poor performance, difficult streaming
- **Reason for not choosing**: Not suitable for file operations

### Option D: In-Memory Processing Only
- **Description**: Process files without storing them
- **Pros**: No storage needed, simpler architecture
- **Cons**: No retry capability, memory limitations, no audit trail
- **Reason for not choosing**: Poor user experience and reliability

## Implementation Notes

### MVP Implementation Steps
1. **Week 1**: Deploy MinIO and ClamAV containers
2. **Week 1**: Implement CSV import pipeline
3. **Week 2**: Add format detection and validation
4. **Week 2**: Implement duplicate detection
5. **Week 3**: Add export functionality
6. **Week 3**: Set up monitoring and alerts

### Security Checklist
- [ ] File type validation (MIME type and magic bytes)
- [ ] File size limits (50MB for imports, 100MB for exports)
- [ ] Virus scanning (ClamAV integration)
- [ ] MinIO access policies (user isolation)
- [ ] Audit logging (MinIO webhook + database)
- [ ] Encryption at rest (MinIO SSE)
- [ ] Rate limiting (5 imports/hour per user)
- [ ] Input validation (CSV injection prevention)
- [ ] Secure parsing (memory limits, timeout)
- [ ] GDPR compliance (data export, deletion)

## References
- ADR-20250819: Data Architecture and Schema Design
- ADR-20250819: Mobile Offline-First Synchronization (OCR on-device)
- ADR-20250819: Deployment Architecture (Hetzner VPS)
- ADR-20250819: Security Architecture
- MinIO Documentation: https://docs.min.io
- ClamAV Integration Guide
- GLBA/CCPA Compliance Requirements

## Decision Makers
- Author: Archie (System Architect)
- Reviewers: [Pending]
- Approver: [Pending]
- Date: 2025-08-20

## Revision History
- 2025-08-20: Initial draft for data import/export file storage (OCR remains on-device)