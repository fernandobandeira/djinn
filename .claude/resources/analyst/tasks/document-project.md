# Document Project Task

## Purpose
Generate comprehensive documentation for existing projects, optimized for AI agents and human developers. Creates structured reference materials that capture the current state, patterns, constraints, and architectural decisions of any codebase.

## Task Workflow

### Step 1: Initial Context Check

**CRITICAL:** First, determine documentation scope:

1. **Check for Requirements/PRD**
   - If exists: Focus on relevant areas only
   - If not: Ask user for focus or document comprehensively

2. **Check Knowledge Base**
   ```bash
   # Search for existing documentation
   ./kb search "architecture" --collection documentation
   ./kb search "system design" --collection architecture
   ```

3. **Determine Documentation Type**
   Ask user:
   ```
   What type of documentation do you need?

   1. Complete system documentation (comprehensive)
   2. Feature-focused documentation (specific enhancement)
   3. Architecture overview (high-level)
   4. API documentation (endpoints and contracts)
   5. Development guide (setup and workflows)
   
   Please select (1-5):
   ```

### Step 2: Scope Definition

#### If Requirements Exist:
```
I found requirements/PRD for [enhancement]. I'll focus documentation on:
- Affected modules and services
- Integration points
- Relevant patterns and conventions
- Impact areas for the planned changes

This keeps documentation lean and relevant. Shall I proceed?
```

#### If No Requirements:
```
To create focused documentation, I recommend:

1. **Create requirements first** - Define what you're building
2. **Provide existing docs** - Share any requirements/specs
3. **Describe the focus** - Brief description of planned work
4. **Document everything** - Comprehensive but potentially excessive

Please select (1-4):
```

### Step 3: Project Analysis

#### 3.1 Structure Discovery
```bash
# Examine project structure
ls -la
tree -L 2 -I 'node_modules|.git|dist|build'

# Identify project type
find . -name "package.json" -o -name "requirements.txt" -o -name "Cargo.toml"

# Check for configuration files
ls -la *.config.* *.json *.yaml *.toml
```

#### 3.2 Technology Stack Identification
- **Package managers**: package.json, requirements.txt, Cargo.toml, go.mod
- **Build tools**: Makefile, webpack.config.js, rollup.config.js
- **Frameworks**: Identify from dependencies
- **Databases**: Check for migrations, schemas, connection configs
- **Testing**: Identify test runners and frameworks

#### 3.3 Documentation Review
```bash
# Find existing documentation
find . -name "README*" -o -name "*.md" | grep -E "(doc|guide|api)"

# Check for API specs
find . -name "*.yaml" -o -name "*.json" | grep -E "(swagger|openapi|api)"
```

#### 3.4 Code Pattern Analysis
Sample key files to understand:
- Naming conventions
- Code organization
- Design patterns
- Error handling
- Authentication/authorization
- Data flow patterns

### Step 4: User Elicitation

Ask targeted questions:

```
Before I analyze your codebase, please help me understand:

1. **Project Purpose**: What does this system do?
2. **Critical Areas**: Which parts are most complex/important?
3. **Tech Debt**: Any known issues or workarounds?
4. **Conventions**: Special patterns or practices used?
5. **Future Plans**: What changes are planned?

[Await responses]
```

### Step 5: Deep Analysis

#### 5.1 Entry Points
- Main files (index.*, main.*, app.*)
- Route definitions
- Command handlers
- API endpoints

#### 5.2 Architecture Mapping
- Service boundaries
- Data flow
- External integrations
- Authentication flow
- State management

#### 5.3 Business Logic
- Core domain models
- Business rules
- Validation logic
- Transaction handling

#### 5.4 Infrastructure
- Database schemas
- Caching strategies
- Queue systems
- Deployment configuration

### Step 6: Documentation Generation

## Output Template

```markdown
# [Project Name] Architecture Documentation

## Executive Summary
Brief overview of the system, its purpose, and current state.

## Quick Start Guide

### Prerequisites
- Required software and versions
- Environment setup
- Access requirements

### Local Development
\`\`\`bash
# Clone repository
git clone [repo]

# Install dependencies
[package manager] install

# Configure environment
cp .env.example .env
# Edit .env with your settings

# Run development server
[start command]
\`\`\`

## System Architecture

### High-Level Overview
[Architecture diagram or ASCII art]

### Technology Stack
| Category | Technology | Version | Notes |
|----------|------------|---------|-------|
| Runtime | Node.js | 18.x | |
| Framework | Express | 4.x | |
| Database | PostgreSQL | 14 | |
| Cache | Redis | 7.x | |

### Project Structure
\`\`\`
project-root/
├── src/
│   ├── controllers/     # Request handlers
│   ├── services/        # Business logic
│   ├── models/          # Data models
│   ├── middlewares/     # Express middleware
│   └── utils/           # Utilities
├── tests/               # Test suites
├── docs/                # Documentation
└── config/              # Configuration
\`\`\`

## Core Components

### [Component Name]
**Purpose**: What it does
**Location**: Where to find it
**Key Files**: 
- \`src/services/[name].js\` - Main logic
- \`src/models/[name].js\` - Data model

**Patterns Used**:
- [Pattern description]

**Dependencies**:
- Internal: [List]
- External: [List]

[Repeat for each major component]

## Data Architecture

### Database Schema
Reference to actual schema files or ERD

### Data Flow
1. Request arrives at controller
2. Controller validates input
3. Service processes business logic
4. Repository handles database operations
5. Response returned to client

## API Documentation

### Authentication
- Method: JWT/OAuth/Basic
- Token location: Header/Cookie
- Refresh strategy: [Description]

### Endpoints
Reference to OpenAPI spec or summary:

| Method | Path | Purpose | Auth Required |
|--------|------|---------|---------------|
| GET | /api/users | List users | Yes |
| POST | /api/users | Create user | Yes |

## Integration Points

### External Services
| Service | Purpose | Integration Type | Documentation |
|---------|---------|------------------|---------------|
| Stripe | Payments | REST API | [Link] |
| SendGrid | Email | SDK | [Link] |

### Internal Services
- Service communication patterns
- Message queues
- Event systems

## Development Workflows

### Git Workflow
- Branching strategy
- Commit conventions
- PR process

### Testing Strategy
- Unit tests: [Framework, coverage]
- Integration tests: [Approach]
- E2E tests: [Tools]

### Deployment
- Environments: Dev, Staging, Prod
- CI/CD pipeline
- Deployment process

## Technical Debt & Known Issues

### Critical Issues
1. **[Issue Name]**
   - Description: [What's wrong]
   - Impact: [Why it matters]
   - Workaround: [Current solution]
   - Fix plan: [If any]

### Technical Debt
- [Area]: [Description and impact]

### Constraints
- [Constraint]: [Reason and implications]

## Conventions & Patterns

### Coding Standards
- Style guide: [ESLint/Prettier config]
- Naming conventions
- File organization

### Design Patterns
- Repository pattern for data access
- Service layer for business logic
- [Other patterns]

### Error Handling
- Error types and codes
- Logging strategy
- User-facing messages

## Performance Considerations

### Optimization Strategies
- Caching: [Redis for X]
- Database: [Indexes, query optimization]
- API: [Rate limiting, pagination]

### Monitoring
- Metrics collected
- Alerting thresholds
- Performance baselines

## Security

### Authentication & Authorization
- Auth method and flow
- Permission system
- Session management

### Security Measures
- Input validation
- SQL injection prevention
- XSS protection
- CSRF tokens

## Appendix

### Useful Commands
\`\`\`bash
# Development
npm run dev

# Testing
npm test
npm run test:coverage

# Database
npm run migrate
npm run seed

# Production
npm run build
npm start
\`\`\`

### Environment Variables
| Variable | Purpose | Example |
|----------|---------|---------|
| DATABASE_URL | Database connection | postgresql://... |
| JWT_SECRET | Token signing | [random string] |

### Troubleshooting
Common issues and solutions

### Resources
- [Internal Wiki]
- [API Documentation]
- [Architecture Decisions]
```

### Step 7: Knowledge Base Integration

After documentation generation:

```bash
# Index the documentation
./kb index --path ./docs/architecture.md

# Tag for easy retrieval
./kb index --path ./docs --mode docs

# Update statistics
./kb stats
```

### Step 8: Refinement Options

After initial documentation, offer elicitation:

```
Documentation draft complete. Would you like to:

1. Apply advanced elicitation for refinement
2. Focus on specific sections
3. Add custom sections
4. Generate specialized views (API-only, Setup-only)
5. Export in different format
6. Save and index in knowledge base

Select option (1-6) or press Enter to finish:
```

## Focus Strategies

### For Enhancement Documentation
Focus on:
- Affected modules only
- Integration points
- Current implementation
- Change impact areas
- Migration considerations

### For Comprehensive Documentation
Cover:
- All system components
- Complete architecture
- Every integration
- Full API surface
- All workflows

### For API Documentation
Emphasize:
- Endpoint specifications
- Request/response formats
- Authentication details
- Rate limiting
- Error responses
- Examples

## Quality Checklist

Before finalizing:
- [ ] Accurate file paths
- [ ] Current versions
- [ ] Working commands
- [ ] Real constraints documented
- [ ] Technical debt acknowledged
- [ ] Patterns identified
- [ ] Knowledge base indexed

## Success Criteria

- Documentation reflects ACTUAL state (not ideal)
- Enables immediate productivity
- Answers common questions
- Identifies pain points
- Provides clear navigation
- Integrated with knowledge base

## Output Management

### File Locations
- `/docs/architecture/` - System documentation
- `/docs/api/` - API specifications
- `/docs/guides/` - Developer guides
- Knowledge base - Indexed and searchable

### Naming Conventions
- `architecture-[date].md` - System docs
- `api-spec-[version].md` - API docs
- `dev-guide-[topic].md` - Guides
- `analysis-[feature].md` - Feature analysis

## Remember

- Document REALITY, not aspirations
- Include workarounds and debt
- Reference actual files
- Keep it actionable
- Index in knowledge base
- Update regularly