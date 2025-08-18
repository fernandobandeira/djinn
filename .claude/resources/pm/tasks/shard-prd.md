# Shard PRD Task

## Purpose
Split a monolithic PRD into searchable, modular components for better Knowledge Base integration and team collaboration.

## When to Shard
- PRD exceeds 500 lines
- Multiple epics defined
- Different team members need different sections
- KB search returns too much context

## Sharding Structure

### Main PRD Structure
```
/docs/requirements/
├── prd.md                    # Main index/overview
├── prd-executive-summary.md  # Executive summary
├── prd-user-research.md      # User research findings
├── prd-requirements.md        # Functional requirements
├── prd-technical.md          # Technical requirements
├── epics/
│   ├── epic-1-auth.md       # Epic 1: Authentication
│   ├── epic-2-core.md       # Epic 2: Core Features
│   ├── epic-3-admin.md      # Epic 3: Admin Panel
│   └── epic-index.md        # Epic overview
└── acceptance/
    ├── ac-epic-1.md         # Acceptance criteria Epic 1
    ├── ac-epic-2.md         # Acceptance criteria Epic 2
    └── ac-epic-3.md         # Acceptance criteria Epic 3
```

### Main PRD Index (prd.md)
```markdown
# Product Requirements Document

## Document Map
- [Executive Summary](./prd-executive-summary.md)
- [User Research](./prd-user-research.md)
- [Functional Requirements](./prd-requirements.md)
- [Technical Requirements](./prd-technical.md)
- [Epics Overview](./epics/epic-index.md)

## Quick Links
- Epic 1: [Authentication](./epics/epic-1-auth.md)
- Epic 2: [Core Features](./epics/epic-2-core.md)
- Epic 3: [Admin Panel](./epics/epic-3-admin.md)

## Version
- Version: 2.0
- Last Updated: {{date}}
- Status: {{status}}

## Search Tips
Use KB search with specific queries:
- `./.vector_db/kb search "authentication requirements" --collection documentation`
- `./.vector_db/kb search "epic 2 stories" --collection documentation`
- `./.vector_db/kb search "acceptance criteria admin" --collection documentation`
```

### Epic File Structure (epic-N-name.md)
```markdown
# Epic N: {{Epic Title}}

## Epic Overview
{{epic_description}}

## Business Value
{{business_justification}}

## User Stories
### Story N.1: {{story_title}}
**As a** {{role}}
**I want** {{action}}
**So that** {{benefit}}

**Priority**: {{priority}}
**Estimate**: {{points}}

### Story N.2: {{story_title}}
...

## Acceptance Criteria
[Detailed criteria in ac-epic-N.md](../acceptance/ac-epic-N.md)

## Dependencies
- Depends on: {{dependencies}}
- Blocks: {{blocked_items}}

## Technical Considerations
{{technical_notes}}

## Success Metrics
{{metrics}}
```

## Sharding Workflow

### 1. Analyze PRD Size
```bash
# Check line count
wc -l /docs/requirements/prd.md

# If > 500 lines, proceed with sharding
```

### 2. Create Directory Structure
```bash
mkdir -p /docs/requirements/epics
mkdir -p /docs/requirements/acceptance
```

### 3. Split Content
```bash
# Extract sections to separate files
# Keep cross-references via relative links
```

### 4. Create Index Files
- Main prd.md becomes navigation hub
- epic-index.md provides epic overview
- Each section self-contained but linked

### 5. Update Cross-References
Replace inline content with links:
```markdown
<!-- Before -->
## Authentication Requirements
[500 lines of content]

<!-- After -->
## Authentication Requirements
See [Epic 1: Authentication](./epics/epic-1-auth.md)
```

### 6. Index with KB
```bash
# Index all new files
./.vector_db/kb index --path /docs/requirements/

# Verify searchability
./.vector_db/kb search "authentication" --collection documentation
./.vector_db/kb search "epic 2" --collection documentation
```

## Benefits of Sharding

### For KB Search
- More precise search results
- Relevant context without noise
- Faster indexing
- Better semantic matching

### For Team Collaboration
- SM loads only relevant epics
- Dev searches specific stories
- QA finds acceptance criteria easily
- Parallel work on different epics

### For Maintenance
- Update sections independently
- Version control per component
- Clear ownership boundaries
- Easier review process

## Example KB Queries After Sharding

```bash
# Find specific epic
./.vector_db/kb search "epic authentication" --collection documentation

# Find acceptance criteria
./.vector_db/kb search "acceptance criteria payment" --collection documentation

# Find user research
./.vector_db/kb search "user personas" --collection documentation

# Find technical requirements
./.vector_db/kb search "API requirements" --collection documentation
```

## Anti-Patterns to Avoid
- Don't shard too granularly (< 100 lines per file)
- Don't lose context by over-splitting
- Don't break logical groupings
- Don't create circular references
- Don't duplicate content across shards