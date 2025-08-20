# Low-Cost MVP Hosting Comparison for Go + PostgreSQL

## Executive Summary
Comprehensive analysis of ultra-low-cost hosting options for MVP deployment, focusing on Go backend applications with PostgreSQL databases and no-sleep requirements.

## Date
2025-08-19

## Top Recommendations

### Best Overall Value: Hetzner Cloud CX22
- **Cost**: €3.79/month (~$4.20)
- **Specs**: 2 vCPU, 4GB RAM, 40GB SSD
- **Pros**: No sleep, excellent performance, full control
- **Cons**: Requires DevOps knowledge, manual scaling

### Best Free Option: Fly.io + Supabase
- **Cost**: $0 (with limitations)
- **Pros**: Generous free tiers, minimal cold starts
- **Cons**: Database pauses after 1 week, limited resources

### Best Developer Experience: Railway
- **Cost**: $5-20/month (usage-based)
- **Pros**: One-click deploy, no cold starts, managed PostgreSQL
- **Cons**: Can get expensive with scale

## Detailed Platform Analysis

### VPS Providers (Always-On, No Sleep)

#### Hetzner Cloud
- **CX22**: €3.79/month (2 vCPU, 4GB RAM, 40GB SSD)
- **CX32**: €7.59/month (4 vCPU, 8GB RAM, 80GB SSD)
- **Benefits**: German quality, excellent network, no bandwidth limits
- **Go Deployment**: Docker, systemd, or binary deployment
- **PostgreSQL**: Self-hosted with full control

#### Contabo
- **VPS S**: €5.99/month (4 vCPU, 8GB RAM, 50GB NVMe)
- **Benefits**: Very high resources for price
- **Concerns**: Overselling, mixed reliability reports

#### DigitalOcean
- **Basic Droplet**: $6/month (1 vCPU, 1GB RAM, 25GB SSD)
- **Benefits**: $200 free credit, excellent documentation
- **Go Deployment**: App Platform available ($5/month minimum)

#### Vultr
- **Regular Performance**: $6/month (1 vCPU, 1GB RAM, 25GB SSD)
- **Benefits**: $100 free credit, global locations
- **Special**: High-frequency compute available

#### Linode (Akamai)
- **Nanode**: $5/month (1 vCPU, 1GB RAM, 25GB SSD)
- **Benefits**: $100 trial credit, reliable service
- **Go Support**: Excellent documentation

### PaaS Platforms

#### Railway
- **Pricing**: Usage-based, ~$5-20/month for small apps
- **Free Tier**: $5 trial credit
- **No Sleep**: Services stay warm with any usage
- **PostgreSQL**: Managed, included in pricing
- **Go Deploy**: Nixpacks auto-detection, excellent

#### Fly.io
- **Free Tier**: 3 shared VMs, 3GB storage
- **No Sleep**: Scales to zero but fast wake (100-300ms)
- **PostgreSQL**: Via separate providers or self-hosted
- **Go Deploy**: Excellent Go support, global deployment

#### Render
- **Free Tier**: Available but sleeps after 15 min
- **Paid**: $7/month for always-on
- **PostgreSQL**: Free tier (expires after 90 days)
- **Go Deploy**: Native Go support

#### Koyeb
- **Free Tier**: $5.50 worth of resources
- **Features**: Global deployment, built-in CI/CD
- **Sleep**: Scales to zero on free tier
- **PostgreSQL**: Available as managed service

#### Northflank
- **Free Tier**: Limited but available
- **Pricing**: $6/month for basic services
- **Features**: Multi-cloud, good developer experience

### Database-Specific Options

#### Supabase
- **Free Tier**: 500MB database, 2 projects
- **Sleep**: Pauses after 1 week inactivity
- **Features**: Built-in auth, realtime, storage
- **Cost**: $25/month for no-pause

#### Neon
- **Free Tier**: 512MB storage, scales to zero
- **Features**: Branching, serverless PostgreSQL
- **Wake Time**: 500ms cold start
- **Cost**: $19/month for always-on

#### ElephantSQL
- **Free Tier**: 20MB (very limited)
- **Tiny Turtle**: $5/month for 1GB
- **Features**: Managed PostgreSQL, reliable

#### Aiven
- **No Free Tier**: Starts at $19/month
- **Features**: Professional PostgreSQL hosting
- **Benefits**: High availability, backups

### Container Platforms

#### Google Cloud Run
- **Free Tier**: 2M requests, 360K GB-seconds
- **Sleep**: Yes, but fast cold starts
- **PostgreSQL**: Cloud SQL ($7.67/month minimum)
- **Go Support**: Excellent

#### AWS Free Tier
- **EC2**: 750 hours/month for 12 months
- **RDS**: 750 hours PostgreSQL for 12 months
- **Limitations**: Complex setup, expires after 1 year

## Cost Comparison Table

| Platform | Monthly Cost | Always-On | PostgreSQL | Setup Complexity |
|----------|-------------|-----------|------------|------------------|
| Hetzner CX22 | €3.79 | ✅ | Self-hosted | Medium |
| Railway | $5-20 | ✅ | Managed | Very Low |
| Fly.io | $0 | ⚠️ | External | Low |
| Render Free | $0 | ❌ | Limited | Low |
| Render Paid | $7 | ✅ | Available | Low |
| DO + Managed DB | $18 | ✅ | Managed | Medium |
| Supabase | $0 | ⚠️ | Managed | Very Low |
| GCP Cloud Run | $0-10 | ❌ | External | Medium |

## Deployment Patterns for Go

### VPS Deployment
```bash
# Binary deployment
scp djinn-api user@server:/opt/djinn/
ssh user@server 'systemctl restart djinn'

# Docker deployment
docker build -t djinn-api .
docker save djinn-api | ssh user@server docker load
ssh user@server 'docker-compose up -d'
```

### PaaS Deployment
```bash
# Railway
railway up

# Fly.io
fly deploy

# Render
git push origin main  # Auto-deploys
```

## Recommendations by Use Case

### Just Testing/Learning
- **Choice**: Fly.io + Supabase (free)
- **Rationale**: Zero cost, easy setup

### Serious MVP Testing
- **Choice**: Hetzner CX22 VPS
- **Rationale**: Best performance per dollar, no restrictions

### Want Simplicity
- **Choice**: Railway
- **Rationale**: Excellent DX, scales automatically

### Need Geographic Distribution
- **Choice**: Fly.io (paid)
- **Rationale**: Edge deployment capabilities

## Migration Considerations

### From VPS to PaaS
- Containerize application
- Externalize configuration
- Set up CI/CD pipeline
- Migrate database (dump/restore)

### From PaaS to VPS
- Set up server infrastructure
- Configure monitoring
- Implement backup strategy
- Handle SSL/security

## Security Considerations

All platforms analyzed provide:
- SSL/TLS certificates (free via Let's Encrypt)
- DDoS protection (varies by tier)
- Firewall capabilities
- Environment variable management

## Conclusion

For ultra-low-cost MVP deployment:
1. **Hetzner VPS** offers best value for always-on services
2. **Railway** provides best developer experience
3. **Fly.io + Supabase** maximizes free tier usage

Choose based on your technical expertise and specific requirements.