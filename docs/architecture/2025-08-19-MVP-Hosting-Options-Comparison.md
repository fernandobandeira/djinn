# MVP Hosting Options Comparison 2025: Go/Golang + PostgreSQL

## Executive Summary

This comprehensive analysis evaluates hosting platforms suitable for MVP deployment with Go backends and PostgreSQL databases, focusing on generous free tiers and low-cost options ($0-20/month). The research covers VPS providers, PaaS platforms, database hosting, and container services.

## Key Findings

### Best Overall Options for MVP Testing

1. **Hetzner Cloud** - Best VPS value with reliable always-on service
2. **Railway** - Best PaaS with genuine usage-based pricing
3. **Fly.io** - Excellent free tier and Go deployment support
4. **Supabase + Neon** - Best PostgreSQL free tiers

## Detailed Platform Analysis

### VPS Providers (Always-On, No Sleep)

#### 1. Hetzner Cloud ⭐⭐⭐⭐⭐
- **Pricing**: CX22 (2 vCPU, 4GB RAM, 40GB SSD) €3.79/month ($4.59)
- **Free Tier**: None, but extremely low cost
- **Sleep/Cold Start**: Never - always-on VPS
- **Go Support**: Excellent - full control over environment
- **PostgreSQL**: Self-hosted, full control
- **Pros**:
  - Best price/performance ratio
  - European data centers (GDPR compliant)
  - Excellent network (20TB traffic included)
  - Simple hourly billing
  - Full root access
- **Cons**: 
  - No free tier
  - Requires server management
  - Self-managed database backups

#### 2. Vultr Cloud Compute
- **Pricing**: Regular Performance 1 vCPU, 1GB RAM, 25GB SSD) $6/month
- **Free Tier**: None
- **Sleep/Cold Start**: Never - always-on VPS
- **Go Support**: Excellent - full control
- **PostgreSQL**: Self-hosted
- **Pros**:
  - Global data centers (32 regions)
  - High-performance NVMe storage
  - Good API and automation tools
- **Cons**:
  - Higher cost than Hetzner
  - No free tier
  - Requires server management

#### 3. Linode (Akamai)
- **Pricing**: Nanode 1GB $5/month
- **Free Tier**: $100 credit for 60 days (new accounts)
- **Sleep/Cold Start**: Never - always-on VPS
- **Go Support**: Excellent
- **PostgreSQL**: Self-hosted
- **Pros**:
  - Generous trial credit
  - Good documentation
  - Reliable infrastructure
- **Cons**:
  - Higher base cost
  - Limited free tier duration

### PaaS Platforms (Potential Sleep Issues)

#### 1. Railway ⭐⭐⭐⭐⭐
- **Pricing**: $5/month minimum (includes $5 usage), $20/month Pro
- **Free Tier**: $5 trial credits (30 days)
- **Sleep/Cold Start**: None - pay-per-use keeps services running
- **Go Support**: Excellent - native Go deployment
- **PostgreSQL**: Managed PostgreSQL included
- **Pros**:
  - True usage-based pricing
  - No cold starts when paying for usage
  - Excellent Go integration
  - Managed PostgreSQL included
  - Simple deployment from Git
  - Up to 32GB RAM/32 vCPU on Pro
- **Cons**:
  - Limited free trial
  - Still relatively new platform

#### 2. Fly.io ⭐⭐⭐⭐
- **Pricing**: $0 free tier, then pay-per-use
- **Free Tier**: 
  - 3x 256MB shared vCPU instances
  - 3GB storage
  - 160GB outbound traffic
- **Sleep/Cold Start**: Minimal - services stay warm with traffic
- **Go Support**: Excellent - optimized for Go deployments
- **PostgreSQL**: Separate PostgreSQL app or external service
- **Pros**:
  - Generous free tier
  - Excellent Go deployment tools
  - Global edge deployment
  - No sleep on free tier with activity
  - Strong documentation
- **Cons**:
  - PostgreSQL requires separate setup
  - Complex pricing model
  - Can become expensive with scale

#### 3. Koyeb ⭐⭐⭐⭐
- **Pricing**: Starter $0/month, Pro $29/month
- **Free Tier**: 
  - 1 web service (512MB RAM)
  - 1 PostgreSQL database (5h free monthly)
  - Scale-to-zero
- **Sleep/Cold Start**: Yes - scales to zero when inactive
- **Go Support**: Good - container-based deployment
- **PostgreSQL**: Serverless PostgreSQL included
- **Pros**:
  - True free tier
  - Managed PostgreSQL
  - Autoscaling
  - Good for intermittent testing
- **Cons**:
  - Cold starts affect always-on requirement
  - Limited free PostgreSQL hours
  - Newer platform with less ecosystem

### Database Hosting Options

#### 1. Supabase ⭐⭐⭐⭐⭐
- **Pricing**: Free tier, Pro $25/month
- **Free Tier**:
  - Up to 500MB database
  - Up to 2GB bandwidth
  - 50,000 monthly active users
  - 7-day log retention
- **Sleep/Cold Start**: Database pauses after 1 week inactivity
- **Pros**:
  - Generous free tier
  - Full PostgreSQL compatibility
  - Excellent tooling and dashboard
  - Real-time subscriptions
  - Built-in authentication
- **Cons**:
  - Database pausing on free tier
  - Storage limitations

#### 2. Neon ⭐⭐⭐⭐
- **Pricing**: Free tier, Pro $19/month
- **Free Tier**:
  - 512MB storage
  - Compute scales to zero
  - 1 concurrent connection
- **Sleep/Cold Start**: Yes - compute scales to zero
- **Pros**:
  - True serverless PostgreSQL
  - Branch databases for development
  - Fast cold starts
- **Cons**:
  - Very limited free storage
  - Single connection limit on free tier
  - Cold starts impact performance

#### 3. Aiven
- **Pricing**: Starts around $19/month for PostgreSQL
- **Free Tier**: 30-day trial
- **Sleep/Cold Start**: Never - always-on managed service
- **Pros**:
  - Professional-grade service
  - Multiple cloud regions
  - Always-on reliability
- **Cons**:
  - No meaningful free tier
  - Higher cost for MVP testing

### Container Platforms

#### 1. Google Cloud Run ⭐⭐⭐
- **Pricing**: 
  - 2 million requests/month free
  - 400,000 GB-seconds/month free
  - 200,000 vCPU-seconds/month free
- **Free Tier**: Very generous compute limits
- **Sleep/Cold Start**: Yes - scales to zero between requests
- **Go Support**: Excellent
- **PostgreSQL**: Requires external database
- **Pros**:
  - Extremely generous free tier
  - Excellent Go support
  - Global deployment
  - Only pay for actual usage
- **Cons**:
  - Cold starts impact user experience
  - Requires external database
  - Complex for persistent connections

#### 2. AWS Free Tier
- **Pricing**: 12 months free tier
- **Free Tier**:
  - EC2: 750 hours/month t2.micro
  - RDS: 750 hours/month db.t2.micro PostgreSQL
- **Sleep/Cold Start**: EC2 never sleeps, RDS configurable
- **Go Support**: Excellent
- **PostgreSQL**: RDS PostgreSQL included
- **Pros**:
  - Complete infrastructure included
  - Full PostgreSQL instance
  - Professional-grade services
- **Cons**:
  - Complex pricing after free tier
  - 12-month limit
  - Steep learning curve

## Recommended Combinations for MVP Testing

### Option 1: Budget Professional (€8-12/month)
- **Compute**: Hetzner Cloud CX22 (€3.79/month)
- **Database**: Self-hosted PostgreSQL on same instance
- **Total**: €3.79/month (~$4.59)
- **Best for**: MVP with predictable traffic, full control needed

### Option 2: Managed Convenience ($25-30/month)
- **Compute**: Railway Pro ($20/month with $20 included usage)
- **Database**: Built-in PostgreSQL
- **Total**: $20/month minimum
- **Best for**: Focus on development, not infrastructure

### Option 3: Maximum Free Tier (Free for 2-3 months)
- **Compute**: Fly.io free tier
- **Database**: Supabase free tier
- **Total**: $0/month for limited usage
- **Best for**: Early prototyping, proof of concept

### Option 4: Hybrid Approach ($5-10/month)
- **Compute**: Hetzner Cloud CX22 (€3.79/month)
- **Database**: Supabase free tier or Neon free tier
- **Total**: €3.79/month + potential database costs
- **Best for**: Reliable compute with managed database

## Go Deployment Considerations

### VPS Deployment
```bash
# Typical Go deployment on VPS
GOOS=linux GOARCH=amd64 go build -o myapp
scp myapp user@vps:/opt/myapp/
ssh user@vps "systemctl restart myapp"
```

### Container Platforms
```dockerfile
# Multi-stage Docker build for Go
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
```

## Performance and Reliability Assessment

### Always-On Services (No Cold Starts)
1. **VPS Providers**: Hetzner, Vultr, Linode
2. **Managed Services**: Aiven PostgreSQL
3. **PaaS with Usage**: Railway (when paying for usage)

### Services with Cold Starts
1. **Serverless**: Google Cloud Run, Koyeb
2. **Free Tiers**: Fly.io (minimal), Supabase (database pausing)
3. **Auto-scaling**: Neon PostgreSQL

## Cost Projections by Usage

### Low Traffic MVP (< 1000 requests/day)
- **Hetzner + Self-hosted DB**: €3.79/month
- **Railway**: $5-8/month
- **Fly.io + Supabase**: $0-5/month

### Medium Traffic MVP (10k requests/day)
- **Hetzner + Self-hosted DB**: €3.79-6.80/month
- **Railway**: $15-25/month
- **Fly.io + Supabase**: $10-20/month

### High Traffic MVP (100k requests/day)
- **Hetzner CX32**: €6.80/month
- **Railway**: $30-50/month
- **Requires paid database tiers**: +$19-25/month

## Final Recommendations by Use Case

### For Learning/Prototyping
- **Primary**: Fly.io free tier + Supabase free tier
- **Backup**: Google Cloud Run + Neon free tier

### For Serious MVP Testing
- **Primary**: Hetzner Cloud CX22 + self-hosted PostgreSQL
- **Alternative**: Railway Pro plan

### For Professional Development
- **Primary**: Railway Pro + managed PostgreSQL
- **Enterprise**: Hetzner + Aiven PostgreSQL

### For Maximum Uptime on Budget
- **Primary**: Hetzner Cloud CX22 with self-hosted PostgreSQL
- **Monitoring**: Railway for staging environment

## Conclusion

**For MVP deployment prioritizing reliability and cost-effectiveness**, Hetzner Cloud CX22 with self-hosted PostgreSQL offers the best value at €3.79/month with no cold starts and excellent performance.

**For developer productivity and managed services**, Railway provides the best experience with true usage-based pricing and no cold starts when services are actively used.

**For maximum free tier testing**, the combination of Fly.io compute with Supabase PostgreSQL offers several months of free development and testing capability.

The choice depends on your priorities: cost optimization (Hetzner), developer experience (Railway), or maximizing free usage (Fly.io + Supabase).