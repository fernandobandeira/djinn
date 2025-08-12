# Runbook: [Service/System Name]

## Overview
**Service:** [Name]  
**Owner:** [Team/Person]  
**Last Updated:** [YYYY-MM-DD]  
**Criticality:** Low | Medium | High | Critical  

### Service Description
<!-- What does this service do? -->

### Architecture Diagram
<!-- Link to or embed architecture diagram -->

## Quick Links
- **Dashboard:** [Link]
- **Logs:** [Link]
- **Metrics:** [Link]
- **Source Code:** [Link]
- **Documentation:** [Link]
- **On-Call Schedule:** [Link]

## Key Metrics & SLIs
| Metric | Target | Alert Threshold | Dashboard |
|--------|--------|-----------------|-----------|
| Availability | 99.9% | < 99.5% | [Link] |
| Latency P99 | < 200ms | > 500ms | [Link] |
| Error Rate | < 0.1% | > 1% | [Link] |

## Common Issues & Solutions

### Issue: [High Latency]
**Symptoms:**
- P99 latency > 500ms
- User complaints about slow response

**Diagnosis:**
```bash
# Check current latency
curl -w "@curl-format.txt" https://api.example.com/health

# Check database connections
psql -c "SELECT count(*) FROM pg_stat_activity;"
```

**Resolution:**
1. Check database connection pool
2. Verify cache hit rates
3. Scale up if needed

**Root Cause Prevention:**
- Implement connection pooling limits
- Add monitoring for connection count

### Issue: [Service Down]
**Symptoms:**
- Health check failing
- 5xx errors

**Diagnosis:**
```bash
# Check service status
systemctl status service-name

# Check recent logs
journalctl -u service-name -n 100
```

**Resolution:**
1. Restart service: `systemctl restart service-name`
2. Check disk space: `df -h`
3. Verify configuration: `cat /etc/service/config.yml`

## Deployment Procedures

### Normal Deployment
```bash
# 1. Run tests
make test

# 2. Build
make build

# 3. Deploy to staging
make deploy-staging

# 4. Verify staging
curl https://staging.example.com/health

# 5. Deploy to production
make deploy-prod
```

### Rollback Procedure
```bash
# Get previous version
kubectl rollout history deployment/service-name

# Rollback to previous
kubectl rollout undo deployment/service-name

# Or rollback to specific revision
kubectl rollout undo deployment/service-name --to-revision=2
```

## Emergency Procedures

### Complete Outage
1. **Notify:** Page on-call engineer
2. **Assess:** Check all regions/zones
3. **Communicate:** Update status page
4. **Diagnose:** Follow troubleshooting steps
5. **Fix:** Apply emergency patch if needed
6. **Verify:** Run smoke tests
7. **Post-mortem:** Schedule within 48 hours

### Data Corruption
1. **Stop writes:** Enable read-only mode
2. **Backup current state:** `pg_dump ...`
3. **Identify corruption scope**
4. **Restore from backup if needed**
5. **Verify data integrity**
6. **Re-enable writes**

## Backup & Recovery

### Backup Schedule
- **Database:** Daily at 02:00 UTC
- **Configuration:** On every change
- **Logs:** Retained for 30 days

### Recovery Procedures
```bash
# Restore database from backup
pg_restore -d dbname backup_file.dump

# Restore configuration
aws s3 cp s3://backups/config/latest.tar.gz .
tar -xzf latest.tar.gz
```

### Recovery Time Objectives
- **RTO:** 4 hours
- **RPO:** 24 hours

## Scaling Procedures

### Horizontal Scaling
```bash
# Scale up replicas
kubectl scale deployment service-name --replicas=10

# Autoscaling
kubectl autoscale deployment service-name --min=3 --max=10 --cpu-percent=80
```

### Vertical Scaling
```bash
# Update resource limits
kubectl set resources deployment service-name -c=app --limits=cpu=2,memory=4Gi
```

## Monitoring & Alerting

### Key Alerts
| Alert | Condition | Action | Runbook Section |
|-------|-----------|--------|-----------------|
| High Error Rate | > 1% for 5 min | Check logs | Common Issues |
| Low Disk Space | < 10% free | Clean logs/tmp | Emergency |
| High Memory | > 90% for 10 min | Restart/Scale | Scaling |

### Health Checks
```bash
# Application health
curl https://api.example.com/health

# Database health
psql -c "SELECT 1;"

# Cache health
redis-cli ping
```

## Dependencies

### Upstream Dependencies
- **Database:** PostgreSQL 13
- **Cache:** Redis 6
- **Queue:** RabbitMQ

### Downstream Dependencies
- **Service A:** Uses our API for user data
- **Service B:** Subscribes to our events

## Security Procedures

### Secret Rotation
```bash
# Rotate database password
./scripts/rotate-db-password.sh

# Update API keys
kubectl create secret generic api-keys --from-file=keys.json
```

### Security Incident Response
1. Isolate affected systems
2. Preserve evidence
3. Notify security team
4. Apply patches
5. Document in incident report

## Contacts

### Escalation Path
1. **Primary On-Call:** [PagerDuty]
2. **Secondary On-Call:** [PagerDuty]
3. **Team Lead:** [Name] - [Email]
4. **Director:** [Name] - [Email]

### External Contacts
- **Cloud Provider Support:** [Case URL]
- **Database Vendor:** [Support Number]

## Appendix

### Useful Queries
```sql
-- Top slow queries
SELECT query, mean_exec_time 
FROM pg_stat_statements 
ORDER BY mean_exec_time DESC 
LIMIT 10;
```

### Configuration Files
- Production: `/etc/service/prod.yml`
- Staging: `/etc/service/staging.yml`

### Related Documentation
- [Architecture Decision Records]
- [Service Design Document]
- [API Documentation]