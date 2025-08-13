# Knowledge Base Instructions for AI Agents

## Overview
The knowledge base is a unified vector database that intelligently organizes and retrieves all project knowledge. It uses ChromaDB with semantic search to find relevant information across different document types and domains.

## Collection Categories

### 1. `architecture` Collection
**What's Indexed Here:**
- Architecture Decision Records (ADRs) in `/docs/architecture/adrs/`
- System design documents
- Technical architecture diagrams
- Infrastructure documentation
- Pattern documentation in `/docs/architecture/patterns/`

**How to Search:**
```bash
# Find architectural decisions
./.vector_db/kb search "database selection" --collection architecture
./.vector_db/kb search "ADR microservices" --collection architecture
./.vector_db/kb search "system design" --collection architecture
```

### 2. `documentation` Collection
**What's Indexed Here:**
- Brainstorming session results
- Market research documents
- Competitive analysis
- Project briefs
- Business requirements (PRDs)
- User stories and epics
- Meeting notes

**How to Search:**
```bash
# Find business decisions and research
./.vector_db/kb search "brainstorm features" --collection documentation
./.vector_db/kb search "market analysis" --collection documentation
./.vector_db/kb search "competitor pricing" --collection documentation
./.vector_db/kb search "user requirements" --collection documentation
```

### 3. `code` Collection
**What's Indexed Here:**
- Source code with extracted docstrings
- Function and class definitions
- Implementation patterns
- TODOs and FIXMEs
- Code comments

**How to Search:**
```bash
# Find code implementations
./.vector_db/kb search "authentication implementation" --collection code
./.vector_db/kb search "TODO" --collection code
./.vector_db/kb search "class UserService" --collection code
```

### 4. `api` Collection
**What's Indexed Here:**
- API endpoint definitions
- OpenAPI/Swagger specifications
- GraphQL schemas
- API documentation
- Request/response examples

**How to Search:**
```bash
# Find API information
./.vector_db/kb search "POST /users" --collection api
./.vector_db/kb search "authentication endpoint" --collection api
./.vector_db/kb search "GraphQL mutation" --collection api
```

### 5. `tests` Collection
**What's Indexed Here:**
- Test files and test cases
- Test documentation
- Test patterns and helpers

**How to Search:**
```bash
# Find test information
./.vector_db/kb search "unit test user" --collection tests
./.vector_db/kb search "integration test" --collection tests
```

### 6. `config` Collection
**What's Indexed Here:**
- Configuration files
- Environment settings
- Deployment configurations
- CI/CD pipelines

**How to Search:**
```bash
# Find configuration
./.vector_db/kb search "database config" --collection config
./.vector_db/kb search "environment variables" --collection config
```

## Cross-Domain Search Patterns

### For Architects (Archie)
```bash
# Before designing anything, understand what exists
./.vector_db/kb search "current architecture"  # Search all collections
./.vector_db/kb search "technical debt" --collection code
./.vector_db/kb search "ADR" --collection architecture
./.vector_db/kb search "system constraints" --collection documentation

# Find related business context
./.vector_db/kb search "requirements" --collection documentation
./.vector_db/kb search "user stories" --collection documentation
```

### For Analysts (Ana)
```bash
# Before brainstorming, check existing ideas
./.vector_db/kb search "brainstorm" --collection documentation
./.vector_db/kb search "feature ideas" --collection documentation

# Before market research, check existing analysis
./.vector_db/kb search "market" --collection documentation
./.vector_db/kb search "competitor" --collection documentation

# Find technical constraints for analysis
./.vector_db/kb search "limitations" --collection architecture
./.vector_db/kb search "technology stack" --collection architecture
```

### For Product Managers
```bash
# Find all product decisions
./.vector_db/kb search "PRD" --collection documentation
./.vector_db/kb search "user story" --collection documentation
./.vector_db/kb search "requirements" --collection documentation

# Check technical feasibility
./.vector_db/kb search "feasibility" --collection architecture
./.vector_db/kb search "constraints" --collection architecture
```

## Search Best Practices

### 1. Start Broad, Then Narrow
```bash
# First: broad search across all collections
./.vector_db/kb search "authentication"

# Then: narrow to specific collection
./.vector_db/kb search "authentication" --collection architecture
./.vector_db/kb search "authentication" --collection code
```

### 2. Use Multiple Search Terms
```bash
# Try different variations
./.vector_db/kb search "auth"
./.vector_db/kb search "authentication"
./.vector_db/kb search "login"
./.vector_db/kb search "OAuth"
```

### 3. Search Before Creating
**ALWAYS search before creating new documents:**
```bash
# Before creating an ADR
./.vector_db/kb search "[decision topic]" --collection architecture

# Before brainstorming
./.vector_db/kb search "[feature area]" --collection documentation

# Before writing code
./.vector_db/kb search "[functionality]" --collection code
```

### 4. Cross-Reference Collections
```bash
# Find business context for technical decisions
./.vector_db/kb search "payment" --collection documentation  # Business requirements
./.vector_db/kb search "payment" --collection architecture   # Technical design
./.vector_db/kb search "payment" --collection code          # Implementation

# Find technical context for business decisions
./.vector_db/kb search "scalability" --collection documentation  # Business needs
./.vector_db/kb search "scalability" --collection architecture   # How we scale
```

## Indexing Guidelines

### When to Index
- After creating any document in `/docs/`
- After significant code changes
- After brainstorming sessions
- After architectural decisions

### How to Index
```bash
# Index everything (smart - only changed files)
./.vector_db/kb index

# Index specific paths
./.vector_db/kb index --path ./docs/architecture/adrs/
./.vector_db/kb index --path ./docs/brainstorming/
./.vector_db/kb index --path ./docs/analysis/

# Force re-index (when needed)
./.vector_db/kb index --force
```

### What Gets Auto-Categorized
The KB automatically categorizes based on:
- **File location**: `/docs/architecture/` → architecture collection
- **File name**: `ADR-*.md` → architecture collection
- **Content type**: Brainstorming results → documentation collection
- **File extension**: `.py`, `.js`, `.ts` → code collection

## Query Patterns by Role

### Business Analysis Queries
```bash
# Find all business context
./.vector_db/kb search "business" --collection documentation
./.vector_db/kb search "market" --collection documentation
./.vector_db/kb search "user needs" --collection documentation
./.vector_db/kb search "competitor" --collection documentation
```

### Technical Architecture Queries
```bash
# Find all technical decisions
./.vector_db/kb search "architecture" --collection architecture
./.vector_db/kb search "design pattern" --collection architecture
./.vector_db/kb search "technology choice" --collection architecture
./.vector_db/kb search "infrastructure" --collection architecture
```

### Implementation Queries
```bash
# Find code patterns
./.vector_db/kb search "pattern" --collection code
./.vector_db/kb search "implementation" --collection code
./.vector_db/kb search "TODO" --collection code
./.vector_db/kb search "FIXME" --collection code
```

## Integration with Agents

### Ana (Analyst) Should:
1. Search documentation before any analysis
2. Search for previous brainstorms on topic
3. Search for market research already done
4. Index all new findings

### Archie (Architect) Should:
1. Search architecture collection first
2. Review all ADRs before decisions
3. Search code for current implementation
4. Check documentation for requirements

### All Agents Should:
1. Search before creating
2. Index after creating
3. Cross-reference collections
4. Build on existing knowledge

## Remember
- The KB is cumulative - it grows smarter over time
- Always search before creating new content
- Index regularly to keep knowledge current
- Use collection filters for targeted searches
- Cross-reference for complete context