# ADR-20250120: Dependency Management & Security Scanning

## Status
Accepted

## Context
Djinn requires a comprehensive approach to managing dependencies and detecting security vulnerabilities:
- Financial application handling sensitive user data requires proactive security measures
- Go ecosystem with extensive third-party dependencies needs supply chain security
- Small team (5-8 people) needs automated security processes
- Budget constraints ($50-100/month infrastructure) require cost-effective solutions
- Compliance requirements (GLBA, CCPA) demand documented security practices
- Need to balance security rigor with development velocity
- Must detect and remediate vulnerabilities before they reach production
- Container deployments require image scanning capabilities

### Constraints
- Limited budget rules out expensive commercial security tools initially
- Small team cannot manage complex security infrastructure
- Already using GitHub for source control and CI/CD (per deployment ADR)
- Go 1.25+ with module system for dependency management
- Need SBOM generation for compliance and security auditing
- Must integrate with existing GitHub Actions workflow
- Security scanning must complete within 5 minutes to maintain developer flow
- Need both development-time and production-time security monitoring

## Decision

### 1. Dependency Management Strategy

#### Go Module Configuration
```go
// go.mod - Strict version pinning
module github.com/djinn/djinn

go 1.25

require (
    github.com/99designs/gqlgen v0.17.45
    github.com/plaid/plaid-go/v20 v20.0.0
    github.com/lib/pq v1.10.9
    github.com/redis/go-redis/v9 v9.5.1
    // Pin all dependencies to exact versions
)

// Security-critical replacements
replace (
    // Replace vulnerable versions immediately
    github.com/vulnerable/package => github.com/vulnerable/package v1.2.4
)

// Exclude known vulnerable versions
exclude (
    github.com/example/vulnerable v1.0.0
    github.com/example/vulnerable v1.0.1
)
```

#### Dependency Update Policy
```yaml
update_policy:
  security_updates:
    schedule: "immediately"
    auto_merge: true
    severity: ["critical", "high"]
    
  minor_updates:
    schedule: "weekly"
    auto_merge: false
    review_required: true
    
  major_updates:
    schedule: "monthly"
    auto_merge: false
    review_required: true
    testing_required: true
    
  monitoring:
    - Track dependency freshness
    - Alert on unmaintained packages (>1 year without updates)
    - Review transitive dependencies quarterly
```

### 2. GitHub-Native Security Stack (Free Tier)

#### Core Components
```yaml
security_stack:
  dependency_updates:
    tool: "GitHub Dependabot"
    cost: "$0 (free for all repos)"
    features:
      - Automated dependency PRs
      - Security update prioritization
      - Version update grouping
      - Custom schedules
      
  vulnerability_scanning:
    primary: "Trivy"
    cost: "$0 (open source)"
    features:
      - Go module scanning
      - Container image scanning
      - SBOM generation
      - CVE database updates
      
  go_security_analysis:
    tool: "gosec"
    cost: "$0 (open source)"
    features:
      - Static code analysis
      - CWE detection
      - Security anti-patterns
      
  license_scanning:
    tool: "Trivy"
    cost: "$0 (included)"
    features:
      - License detection
      - Compliance checking
      
  secret_scanning:
    tool: "gitleaks"
    cost: "$0 (open source)"
    features:
      - Pre-commit hooks
      - Historical scanning
      - Custom patterns
```

#### Dependabot Configuration
```yaml
# .github/dependabot.yml
version: 2
updates:
  # Go modules
  - package-ecosystem: "gomod"
    directory: "/"
    schedule:
      interval: "daily"
      time: "06:00"
      timezone: "America/Los_Angeles"
    reviewers:
      - "@djinn/security-team"
    assignees:
      - "@djinn/backend-team"
    labels:
      - "dependencies"
      - "go"
    commit-message:
      prefix: "deps"
      include: "scope"
    open-pull-requests-limit: 10
    groups:
      # Group minor and patch updates
      minor-and-patch:
        patterns:
          - "*"
        update-types:
          - "minor"
          - "patch"
        exclude-patterns:
          - "github.com/plaid/*"  # Financial APIs updated separately
          - "github.com/lib/pq"    # Database driver needs careful testing
      
      # Security updates in separate PRs for immediate attention
      security:
        patterns:
          - "*"
        update-types:
          - "security"
    
  # Docker base images
  - package-ecosystem: "docker"
    directory: "/"
    schedule:
      interval: "weekly"
    reviewers:
      - "@djinn/devops-team"
    
  # GitHub Actions
  - package-ecosystem: "github-actions"
    directory: "/"
    schedule:
      interval: "monthly"
```

### 3. CI/CD Security Pipeline

#### GitHub Actions Security Workflow
```yaml
# .github/workflows/security.yml
name: Security Pipeline
on:
  push:
    branches: [main, develop]
  pull_request:
    branches: [main]
  schedule:
    # Daily security scan at 2 AM UTC
    - cron: '0 2 * * *'

jobs:
  security-scan:
    runs-on: ubuntu-latest
    permissions:
      contents: read
      security-events: write
      
    steps:
    - name: Checkout code
      uses: actions/checkout@v4
      
    - name: Setup Go
      uses: actions/setup-go@v5
      with:
        go-version: '1.25'
        cache: true
    
    # Verify dependency integrity
    - name: Verify dependencies
      run: |
        go mod download
        go mod verify
        go list -m all | grep -E '^\s' | awk '{print $1}' > dependencies.txt
        echo "Found $(wc -l < dependencies.txt) dependencies"
    
    # Vulnerability scanning with Trivy
    - name: Run Trivy vulnerability scanner
      uses: aquasecurity/trivy-action@master
      with:
        scan-type: 'fs'
        scan-ref: '.'
        format: 'sarif'
        output: 'trivy-results.sarif'
        severity: 'CRITICAL,HIGH,MEDIUM'
        ignore-unfixed: false
        
    - name: Upload Trivy results to GitHub Security
      uses: github/codeql-action/upload-sarif@v3
      if: always()
      with:
        sarif_file: 'trivy-results.sarif'
        category: 'trivy'
    
    # Go security analysis with gosec
    - name: Run gosec security scanner
      uses: securego/gosec@master
      with:
        args: '-fmt sarif -out gosec-results.sarif -severity medium ./...'
        
    - name: Upload gosec results
      uses: github/codeql-action/upload-sarif@v3
      if: always()
      with:
        sarif_file: 'gosec-results.sarif'
        category: 'gosec'
    
    # License compliance check
    - name: Check licenses
      run: |
        go install github.com/google/go-licenses@latest
        go-licenses check ./... --disallowed_types=forbidden,restricted
    
    # SBOM generation
    - name: Generate SBOM
      uses: aquasecurity/trivy-action@master
      with:
        scan-type: 'fs'
        format: 'cyclonedx'
        output: 'sbom.json'
        
    - name: Upload SBOM
      uses: actions/upload-artifact@v4
      with:
        name: sbom
        path: sbom.json
        retention-days: 30
    
    # Secret scanning with gitleaks
    - name: Scan for secrets
      uses: gitleaks/gitleaks-action@v2
      env:
        GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
    
    # Fail on critical vulnerabilities
    - name: Check for critical issues
      run: |
        if trivy fs . --severity CRITICAL --exit-code 1; then
          echo "No critical vulnerabilities found"
        else
          echo "Critical vulnerabilities detected!"
          exit 1
        fi

  container-scan:
    runs-on: ubuntu-latest
    if: github.event_name == 'push' && github.ref == 'refs/heads/main'
    needs: security-scan
    
    steps:
    - name: Checkout code
      uses: actions/checkout@v4
      
    - name: Build container image
      run: docker build -t djinn:${{ github.sha }} .
      
    - name: Run Trivy container scan
      uses: aquasecurity/trivy-action@master
      with:
        image-ref: 'djinn:${{ github.sha }}'
        format: 'sarif'
        output: 'container-results.sarif'
        
    - name: Upload container scan results
      uses: github/codeql-action/upload-sarif@v3
      with:
        sarif_file: 'container-results.sarif'
        category: 'container'
```

### 4. Local Development Security

#### Pre-commit Hooks
```yaml
# .pre-commit-config.yaml
repos:
  - repo: https://github.com/zricethezav/gitleaks
    rev: v8.18.0
    hooks:
      - id: gitleaks
        
  - repo: https://github.com/securego/gosec
    rev: v2.18.2
    hooks:
      - id: gosec
        args: ['-severity', 'medium']
        
  - repo: local
    hooks:
      - id: go-mod-tidy
        name: go mod tidy
        entry: go mod tidy
        language: system
        pass_filenames: false
```

#### Developer Setup Script
```bash
#!/bin/bash
# scripts/setup-security.sh

echo "Setting up security tools..."

# Install security tools
go install github.com/securego/gosec/v2/cmd/gosec@latest
go install github.com/google/go-licenses@latest

# Install pre-commit
pip install pre-commit
pre-commit install

# Install Trivy
curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin

echo "Security tools installed successfully"
```

### 5. Security Monitoring and Alerting

#### Metrics and Dashboards
```yaml
security_metrics:
  vulnerability_metrics:
    - total_vulnerabilities_by_severity
    - mean_time_to_remediation
    - vulnerability_age_distribution
    - false_positive_rate
    
  dependency_metrics:
    - total_dependencies
    - outdated_dependencies_count
    - dependency_freshness_score
    - license_compliance_rate
    
  scanning_metrics:
    - scan_execution_time
    - scan_failure_rate
    - coverage_percentage
    
  alerting_rules:
    critical_vulnerability:
      condition: "severity == CRITICAL"
      action: "immediate notification"
      channels: ["slack", "email", "pagerduty"]
      
    high_vulnerability:
      condition: "severity == HIGH && age > 7 days"
      action: "escalation"
      channels: ["slack", "email"]
      
    license_violation:
      condition: "license in ['GPL', 'AGPL']"
      action: "block merge"
      channels: ["slack"]
```

#### Security Review Process
```yaml
security_review:
  automated_checks:
    - No critical vulnerabilities
    - No high vulnerabilities > 30 days old
    - All dependencies from approved sources
    - License compliance verified
    - No secrets detected
    - SBOM generated successfully
    
  manual_review_triggers:
    - New dependency added
    - Major version update
    - Security-sensitive code changes
    - Container base image change
    
  approval_requirements:
    security_updates:
      approvers: 1
      auto_merge_after: "24 hours"
      
    new_dependencies:
      approvers: 2
      security_review: required
      
    major_updates:
      approvers: 2
      testing: required
```

### 6. SBOM Management

#### SBOM Generation and Storage
```bash
# Generate SBOM in multiple formats
trivy fs . --format cyclonedx --output sbom-cyclonedx.json
trivy fs . --format spdx-json --output sbom-spdx.json

# Include in container image
COPY sbom-cyclonedx.json /app/sbom.json

# Store in artifact registry with each release
aws s3 cp sbom-cyclonedx.json s3://djinn-artifacts/sboms/v${VERSION}/
```

#### SBOM Usage
```yaml
sbom_usage:
  compliance:
    - Include in customer deliverables
    - Regulatory reporting
    - Security audits
    
  operational:
    - Vulnerability tracking
    - License analysis
    - Dependency inventory
    
  incident_response:
    - Quick impact assessment
    - Affected version identification
    - Patch verification
```

### 7. Upgrade Path to Advanced Security

#### Future Enhancements (When Budget Allows)
```yaml
phase_2_upgrades:
  github_advanced_security:
    cost: "$21/user/month"
    benefits:
      - CodeQL semantic analysis
      - Advanced secret scanning
      - Security overview dashboard
      - Dependency review in PRs
    trigger: "When team > 10 or revenue > $1M"
    
  snyk_integration:
    cost: "$70/user/month"
    benefits:
      - Automated fix PRs
      - IDE integration
      - Priority scoring
      - Container base image recommendations
    trigger: "When needing compliance certification"
    
  sigstore_signing:
    cost: "$0 (open source)"
    benefits:
      - Artifact signing
      - Supply chain attestation
      - SLSA compliance
    trigger: "When serving enterprise customers"
```

## Consequences

### Positive
- **Zero Initial Cost**: Entire security stack is free to start
- **GitHub Integration**: Single platform for code and security
- **Automated Updates**: Dependabot handles routine updates
- **Comprehensive Scanning**: Multiple tools reduce blind spots
- **Developer Friendly**: Security integrated into existing workflow
- **Compliance Ready**: SBOM generation for audits
- **Scalable**: Can add commercial tools as budget grows
- **Fast Feedback**: Results available in PR checks

### Negative
- **Multiple Tools**: Need to configure and maintain several scanners
- **No Professional Support**: Relying on community tools initially
- **Limited Customization**: Free tier has fewer configuration options
- **Manual Remediation**: No automated fix suggestions (unlike Snyk)
- **Learning Curve**: Team needs to understand security reports

### Risks
- **False Positives**: May require tuning to reduce noise
- **Scanner Gaps**: Single scanner might miss vulnerabilities
- **GitHub Dependency**: Tied to GitHub platform
- **Tool Maintenance**: Open source tools may change or be abandoned
- **Alert Fatigue**: Too many security warnings could be ignored

## Alternatives Considered

### Option A: Open Source Security Stack
- **Description**: Self-hosted Renovate + multiple scanners
- **Pros**: Complete control, no vendor lock-in, highly customizable
- **Cons**: High operational overhead, complex setup, needs dedicated infrastructure
- **Reason for not choosing**: Too complex for small team, higher operational cost

### Option B: Commercial Security Platform (Snyk)
- **Description**: Comprehensive commercial security solution
- **Pros**: Automated fixes, great UX, professional support, IDE integration
- **Cons**: Expensive ($70-200/dev/month), vendor lock-in
- **Reason for not choosing**: Exceeds budget constraints for MVP phase

### Option C: Minimal Security (Go tools only)
- **Description**: Only use Go's built-in security features
- **Pros**: Simplest approach, no additional tools
- **Cons**: Insufficient for financial application, no vulnerability scanning
- **Reason for not choosing**: Inadequate security for sensitive financial data

## Implementation Notes

### Migration Strategy
1. **Week 1**: Set up Dependabot configuration
2. **Week 2**: Implement GitHub Actions security workflow
3. **Week 3**: Configure pre-commit hooks for developers
4. **Week 4**: Establish security review process
5. **Month 2**: Fine-tune scanning rules and alerts
6. **Month 3**: Create security dashboard and metrics

### Success Metrics
- Mean time to patch: < 7 days for high/critical
- Dependency freshness: > 80% updated within 6 months
- Security scan time: < 5 minutes
- False positive rate: < 10%
- SBOM generation: 100% of releases
- Developer satisfaction: > 7/10 with security process

### Testing Approach
- Test security pipeline on feature branches
- Validate SBOM generation accuracy
- Verify vulnerability detection with known CVEs
- Test secret scanning with dummy credentials
- Measure performance impact on CI/CD

### Rollback Plan
- Security scanning is additive, can be disabled without breaking builds
- Dependabot can be turned off via GitHub settings
- Pre-commit hooks are optional for developers
- Can revert to manual security reviews if needed

## References
- [OWASP Go Security Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Go_Security_Cheat_Sheet.html)
- [Google Go Security Best Practices](https://go.dev/doc/security/best-practices)
- [NIST Cybersecurity Supply Chain Risk Management](https://www.nist.gov/cyberframework/supply-chain-risk-management)
- [GitHub Dependabot Documentation](https://docs.github.com/en/code-security/dependabot)
- [Trivy Documentation](https://aquasecurity.github.io/trivy/)
- [CycloneDX SBOM Specification](https://cyclonedx.org/)
- ADR-20250819: Security Architecture
- ADR-20250819: Deployment Architecture

## Decision Makers
- Author: Archie (System Architect)
- Reviewers: [Security Team, Backend Team]
- Approver: [CTO]
- Date: 2025-01-20

## Revision History
- 2025-01-20: Initial draft with GitHub-native free tier approach