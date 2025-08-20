# ADR-20250819: File Storage and Receipt Image Handling

## Status
Proposed

## Context
Djinn needs to handle receipt images for credit card transaction matching and expense tracking. Based on our architecture decisions:
- Users upload receipt images from mobile devices
- Receipts are used for transaction verification and matching
- Users have local copies on their devices
- Receipt storage is temporary (confirmation-based lifecycle)
- No public CDN access needed for receipt images
- Compliance requirements for financial document storage
- Using Hetzner VPS with 40GB local storage for MVP

### Requirements
- Secure storage of receipt images
- OCR-friendly image formats and processing
- Compliance with financial regulations (GLBA, CCPA)
- Temporary storage with user-controlled retention
- No public access to receipt images
- Audit trail for all receipt operations
- Cost-effective storage solution
- Integration with transaction matching workflow

### Constraints
- MVP budget: €3.79/month Hetzner VPS with 40GB storage
- No CDN needed (internal access only)
- Must support mobile upload (iOS/Android)
- Images should be processable for OCR
- Must maintain audit trail for compliance
- Storage should be easily purgeable after confirmation

## Decision

### 1. Storage Architecture

#### MinIO Object Storage (MVP Phase)
```yaml
storage_architecture:
  platform: MinIO (S3-compatible)
  deployment: Docker container on Hetzner VPS
  
  bucket_structure:
    receipts-pending: "Temporary uploads awaiting processing"
    receipts-processed: "OCR-processed receipts"
    receipts-archive: "Confirmed receipts (30-day retention)"
    
  security_features:
    - Server-side encryption (SSE-C)
    - Policy-based access control
    - Audit logging
    - Malware scanning via ClamAV integration
    - File type validation
    - Size limits enforcement
    
  lifecycle:
    upload: "Store in receipts-pending with validation"
    scan: "ClamAV malware scan before acceptance"
    process: "OCR and move to receipts-processed"
    confirm: "User confirms → move to receipts-archive"
    expire: "Auto-delete after 30 days via lifecycle policy"
    
  capacity:
    total_storage: 40GB (Hetzner CX22)
    system_reserved: 10GB
    minio_data: 20GB
    database: 5GB
    application: 5GB
    
  resource_usage:
    minio_memory: ~500MB
    clamav_memory: ~300MB
    available_ram: 3.2GB for app + PostgreSQL
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
    mem_limit: 512m
    
  clamav:
    image: clamav/clamav:stable
    container_name: djinn-clamav
    volumes:
      - /var/djinn/clamav:/var/lib/clamav
    restart: unless-stopped
    mem_limit: 384m
```

### 2. Upload and Processing Pipeline

#### Upload Flow
```yaml
receipt_upload_flow:
  1_mobile_upload:
    endpoint: POST /api/v1/receipts/upload
    authentication: Firebase Auth token
    max_file_size: 10MB
    accepted_formats: [jpeg, jpg, png, heic, pdf]
    
  2_server_processing:
    - Validate file type and size
    - Generate unique upload_id (UUIDv7)
    - Compress/optimize image
    - Store in pending directory
    - Create database record
    - Return upload_id to client
    
  3_background_processing:
    - OCR extraction (async)
    - Merchant detection
    - Amount extraction
    - Date extraction
    - Store extracted data
    
  4_matching_workflow:
    - Match with pending transactions
    - Present matches to user
    - Wait for user confirmation
    
  5_confirmation:
    user_confirms:
      - Link receipt to transaction
      - Move to archived
      - Schedule deletion (30 days)
    user_rejects:
      - Delete receipt immediately
      - Clear database records
```

#### Secure Upload Pipeline with MinIO
```go
// Secure receipt upload with validation and scanning
type SecureReceiptUploader struct {
    minioClient *minio.Client
    clamav      *clamav.Client
    maxFileSize int64  // 10MB
}

func (u *SecureReceiptUploader) Upload(ctx context.Context, userID string, file io.Reader, filename string) error {
    // 1. Validate file type
    buffer := make([]byte, 512)
    n, err := file.Read(buffer)
    if err != nil {
        return fmt.Errorf("failed to read file header: %w", err)
    }
    
    contentType := http.DetectContentType(buffer[:n])
    if !isAllowedImageType(contentType) {
        return fmt.Errorf("invalid file type: %s", contentType)
    }
    
    // 2. Create temp file for scanning
    tmpFile, err := ioutil.TempFile("", "receipt-*")
    if err != nil {
        return err
    }
    defer os.Remove(tmpFile.Name())
    
    // 3. Write to temp file and check size
    multiReader := io.MultiReader(bytes.NewReader(buffer[:n]), file)
    size, err := io.Copy(tmpFile, multiReader)
    if err != nil {
        return err
    }
    
    if size > u.maxFileSize {
        return fmt.Errorf("file too large: %d bytes", size)
    }
    
    // 4. Scan with ClamAV
    scanResult, err := u.clamav.ScanFile(tmpFile.Name())
    if err != nil {
        return fmt.Errorf("malware scan failed: %w", err)
    }
    
    if scanResult.Status != "OK" {
        return fmt.Errorf("malware detected: %s", scanResult.Description)
    }
    
    // 5. Upload to MinIO with encryption
    objectName := fmt.Sprintf("%s/%s/%s", userID, time.Now().Format("2006-01-02"), filename)
    
    tmpFile.Seek(0, 0)
    _, err = u.minioClient.PutObject(ctx, "receipts-pending", objectName, tmpFile, size,
        minio.PutObjectOptions{
            ContentType: contentType,
            UserMetadata: map[string]string{
                "user-id":      userID,
                "upload-time":  time.Now().UTC().Format(time.RFC3339),
                "scan-status":  "clean",
            },
        })
    
    return err
}

// Allowed image types for receipts
func isAllowedImageType(contentType string) bool {
    allowed := []string{
        "image/jpeg",
        "image/png",
        "image/heic",
        "image/heif",
        "application/pdf",
    }
    
    for _, t := range allowed {
        if contentType == t {
            return true
        }
    }
    return false
}
```

### 3. Security and Access Control

#### Access Patterns
```yaml
access_control:
  internal_only:
    description: "No public CDN access"
    implementation: "Files served only through authenticated API"
    
  user_access:
    pattern: "Users can only access their own receipts"
    authentication: "Firebase Auth token required"
    authorization: "User ID from token must match receipt owner"
    
  api_endpoints:
    upload: "POST /api/v1/receipts/upload"
    retrieve: "GET /api/v1/receipts/{receipt_id}"
    delete: "DELETE /api/v1/receipts/{receipt_id}"
    list: "GET /api/v1/receipts"
    
  file_serving:
    method: "API streams file directly"
    no_direct_urls: "Never expose filesystem paths"
    temporary_access: "No signed URLs needed (internal only)"
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
                    "arn:aws:s3:::receipts-*/%s/*"
                ]
            },
            {
                "Effect": "Allow",
                "Action": ["s3:ListBucket"],
                "Resource": ["arn:aws:s3:::receipts-*"],
                "Condition": {
                    "StringLike": {
                        "s3:prefix": ["%s/*"]
                    }
                }
            }
        ]
    }`, userID, userID)
}

// Secure receipt retrieval from MinIO
func (h *ReceiptHandler) GetReceipt(w http.ResponseWriter, r *http.Request) {
    // 1. Authenticate user
    userID, err := h.auth.ValidateToken(r)
    if err != nil {
        http.Error(w, "Unauthorized", 401)
        return
    }
    
    // 2. Get receipt ID and verify ownership
    receiptID := chi.URLParam(r, "receiptID")
    receipt, err := h.db.GetReceipt(receiptID)
    if err != nil || receipt.UserID != userID {
        http.Error(w, "Not found", 404)
        return
    }
    
    // 3. Generate presigned URL (valid for 5 minutes)
    presignedURL, err := h.minioClient.PresignedGetObject(
        r.Context(),
        receipt.Bucket,
        receipt.ObjectName,
        5*time.Minute,
        url.Values{},
    )
    if err != nil {
        http.Error(w, "Internal error", 500)
        return
    }
    
    // 4. Redirect to presigned URL (or proxy the content)
    http.Redirect(w, r, presignedURL.String(), http.StatusTemporaryRedirect)
}
```

### 4. Retention and Lifecycle Management

#### MinIO Lifecycle Policies
```json
{
  "Rules": [
    {
      "ID": "delete-pending-receipts",
      "Status": "Enabled",
      "Prefix": "",
      "Expiration": {
        "Days": 1
      },
      "NoncurrentVersionExpiration": {
        "NoncurrentDays": 1
      }
    },
    {
      "ID": "delete-archived-receipts",
      "Status": "Enabled",
      "Prefix": "",
      "Expiration": {
        "Days": 30
      }
    }
  ]
}
```

```go
// Set up lifecycle policies in MinIO
func setupLifecyclePolicies(client *minio.Client) error {
    // Pending receipts - delete after 24 hours
    pendingPolicy := lifecycle.NewConfiguration()
    pendingPolicy.Rules = []lifecycle.Rule{
        {
            ID:     "delete-old-pending",
            Status: "Enabled",
            Expiration: lifecycle.Expiration{
                Days: 1,
            },
        },
    }
    
    err := client.SetBucketLifecycle(context.Background(), "receipts-pending", pendingPolicy)
    if err != nil {
        return fmt.Errorf("failed to set pending lifecycle: %w", err)
    }
    
    // Archive receipts - delete after 30 days
    archivePolicy := lifecycle.NewConfiguration()
    archivePolicy.Rules = []lifecycle.Rule{
        {
            ID:     "delete-old-archives",
            Status: "Enabled",
            Expiration: lifecycle.Expiration{
                Days: 30,
            },
        },
    }
    
    return client.SetBucketLifecycle(context.Background(), "receipts-archive", archivePolicy)
}
```

#### Cleanup Implementation
```go
// Scheduled cleanup job
func (c *ReceiptCleaner) RunDaily() error {
    // 1. Delete expired pending receipts
    expiredPending := time.Now().Add(-24 * time.Hour)
    pending, err := c.db.GetPendingReceiptsOlderThan(expiredPending)
    for _, receipt := range pending {
        os.Remove(receipt.Path)
        c.db.DeleteReceipt(receipt.ID)
        c.audit.Log("Deleted expired pending receipt", receipt.ID)
    }
    
    // 2. Delete confirmed receipts older than 30 days
    expiredConfirmed := time.Now().Add(-30 * 24 * time.Hour)
    confirmed, err := c.db.GetConfirmedReceiptsOlderThan(expiredConfirmed)
    for _, receipt := range confirmed {
        os.Remove(receipt.Path)
        c.db.DeleteReceipt(receipt.ID)
        c.audit.Log("Deleted expired confirmed receipt", receipt.ID)
    }
    
    // 3. Clean orphaned files
    c.cleanOrphanedFiles()
    
    return nil
}
```

### 5. Database Schema

```sql
-- Receipt metadata storage
CREATE TABLE receipts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    transaction_id UUID REFERENCES transactions(id),
    
    -- File metadata
    original_filename TEXT NOT NULL,
    stored_path TEXT NOT NULL, -- Internal path, never exposed
    file_size INTEGER NOT NULL,
    mime_type TEXT NOT NULL,
    image_hash TEXT NOT NULL, -- For duplicate detection
    
    -- Processing status
    status TEXT NOT NULL CHECK (status IN ('pending', 'processing', 'matched', 'confirmed', 'rejected')),
    
    -- OCR results
    ocr_text TEXT,
    extracted_amount DECIMAL(10,2),
    extracted_merchant TEXT,
    extracted_date DATE,
    ocr_confidence FLOAT,
    
    -- Timestamps
    uploaded_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    processed_at TIMESTAMPTZ,
    confirmed_at TIMESTAMPTZ,
    expires_at TIMESTAMPTZ,
    
    -- Audit
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Indexes
CREATE INDEX idx_receipts_user_id ON receipts(user_id);
CREATE INDEX idx_receipts_status ON receipts(status);
CREATE INDEX idx_receipts_expires_at ON receipts(expires_at) WHERE expires_at IS NOT NULL;
CREATE INDEX idx_receipts_transaction_id ON receipts(transaction_id) WHERE transaction_id IS NOT NULL;

-- Audit table
CREATE TABLE receipt_audit_log (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    receipt_id UUID NOT NULL,
    user_id UUID NOT NULL,
    action TEXT NOT NULL,
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
        
      unusual_access_pattern:
        threshold: ">100 requests/minute from single user"
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
    
    w.WriteHeader(200)
}
```

### 7. Migration Path from MinIO to Cloud Storage

```yaml
future_migration:
  trigger: ">1000 active users or >30GB storage"
  
  target_options:
    option_1:
      service: "Cloudflare R2"
      cost: "$15/TB/month (no egress fees)"
      migration: "MinIO Mirror or rclone sync"
      
    option_2:
      service: "Backblaze B2"
      cost: "$6/TB/month + egress"
      migration: "S3 compatible, easy switch"
      
    option_3:
      service: "Distributed MinIO"
      deployment: "Multiple Hetzner servers"
      cost: "€7.59/server/month"
      benefit: "Stay self-hosted"
      
  migration_strategy:
    preparation:
      - Set up MinIO gateway mode
      - Configure tiering policies
      - Test with staging bucket
      
    execution:
      1_replicate: "Enable bucket replication to cloud"
      2_verify: "Validate all objects replicated"
      3_switch: "Update application endpoints"
      4_monitor: "Watch for issues"
      5_cleanup: "Decommission local MinIO"
```

```bash
# Migration script using MinIO client
#!/bin/bash

# Set up mirror to cloud storage
mc alias set local http://localhost:9000 $MINIO_ACCESS_KEY $MINIO_SECRET_KEY
mc alias set r2 https://r2.cloudflarestorage.com $R2_ACCESS_KEY $R2_SECRET_KEY

# Mirror buckets to cloud
mc mirror --watch local/receipts-archive r2/receipts-archive
mc mirror --watch local/receipts-processed r2/receipts-processed

# Verify migration
mc diff local/receipts-archive r2/receipts-archive
```

## Consequences

### Positive
- **Enhanced security**: MinIO provides enterprise-grade security features
- **S3 compatibility**: Easy migration to cloud storage later
- **Built-in features**: Encryption, versioning, lifecycle policies included
- **Malware protection**: ClamAV integration prevents malicious uploads
- **Audit compliance**: Comprehensive logging for regulatory requirements
- **User isolation**: Policy-based access control per user
- **Automated cleanup**: Lifecycle policies handle retention automatically
- **Professional grade**: Production-ready from day one

### Negative
- **Additional complexity**: MinIO and ClamAV add operational overhead
- **Memory usage**: MinIO + ClamAV use ~800MB RAM
- **Limited capacity**: Still bound by 20GB VPS storage
- **No geo-redundancy**: Single location storage
- **Learning curve**: Team needs to understand MinIO operations

### Risks
- **MinIO availability**: Service failure affects all file operations
- **Storage exhaustion**: Need monitoring and alerts
- **ClamAV false positives**: May reject legitimate receipts
- **Performance impact**: Scanning adds latency to uploads
- **Update management**: Need to keep MinIO and ClamAV updated

## Alternatives Considered

### Option A: Direct Filesystem Storage
- **Description**: Store files directly on VPS filesystem
- **Pros**: Simpler, no additional services
- **Cons**: Security risks, no built-in features, path traversal vulnerabilities
- **Reason for not choosing**: Too many security risks for financial data

### Option B: External Cloud Storage from Start
- **Description**: Use S3/R2/B2 from day one
- **Pros**: Unlimited storage, managed service
- **Cons**: Additional cost ($5-20/month), external dependency
- **Reason for not choosing**: MinIO on VPS is free and sufficient for MVP

### Option C: SeaweedFS Instead of MinIO
- **Description**: Lightweight alternative to MinIO
- **Pros**: Lower memory usage, better for small files
- **Cons**: Less S3 compatibility, smaller ecosystem
- **Reason for not choosing**: MinIO's S3 compatibility valuable for future migration

## Implementation Notes

### MVP Implementation Steps
1. **Week 1**: Set up directory structure and permissions
2. **Week 1**: Implement upload API and file processing
3. **Week 2**: Add OCR processing pipeline
4. **Week 2**: Implement retention and cleanup jobs
5. **Week 3**: Add monitoring and alerts
6. **Week 3**: Test with real receipt images

### Security Checklist
- [ ] File type validation (MIME type and magic bytes)
- [ ] File size limits (10MB max)
- [ ] Virus scanning (ClamAV integration)
- [ ] MinIO access policies (user isolation)
- [ ] Audit logging (MinIO audit webhook)
- [ ] Encryption at rest (MinIO SSE)
- [ ] Rate limiting (API gateway level)
- [ ] Backup encryption (encrypted snapshots)
- [ ] Security updates (automated MinIO updates)
- [ ] Penetration testing (before production)

## References
- ADR-20250819: Deployment Architecture (Hetzner VPS decision)
- ADR-20250819: Security Architecture
- ADR-20250819: Data Architecture and Schema Design
- GLBA Compliance Requirements
- CCPA Data Retention Guidelines
- OCR Best Practices for Financial Documents

## Decision Makers
- Author: Archie (System Architect)
- Reviewers: [Pending]
- Approver: [Pending]
- Date: 2025-08-19

## Revision History
- 2025-08-19: Initial draft focusing on local VPS storage without CDN