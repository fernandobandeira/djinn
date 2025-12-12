# Project Documentation Task

## Overview
This task guides Ana through comprehensive project documentation, from initial discovery to complete architectural documentation with diagrams, workflows, and technical specifications.

## Step 1: Initial Context Check

Ask user to provide context by selecting documentation type:

```markdown
What type of documentation would you like to create?

1. **New Project Documentation** - Documenting an existing project that lacks docs
2. **Architecture Documentation** - Technical architecture, system design, diagrams
3. **API Documentation** - API endpoints, contracts, integration guides
4. **Feature Documentation** - Specific feature or module deep-dive
5. **Process Documentation** - Workflows, procedures, operational guides

Please share:
- Documentation type (1-5)
- Project/repository location
- Any existing documentation
- Specific areas of focus
- Target audience
```

## Step 2: Scope Definition

Based on user response, clarify scope:

### For New Project Documentation:
- Full project overview or specific components?
- Technical depth level needed?
- Audience: developers, stakeholders, or both?
- Timeline: quick reference or comprehensive guide?

### For Architecture Documentation:
- Full system architecture or specific subsystems?
- Need diagrams (C4 model, flowcharts, sequence)?
- Focus on structure, data flow, or both?
- Include deployment/infrastructure?

### For API Documentation:
- REST, GraphQL, gRPC, or other?
- All endpoints or specific API subset?
- Need integration examples?
- Authentication/authorization details?

### For Feature Documentation:
- Feature functionality and usage?
- Technical implementation details?
- User workflows and scenarios?
- Integration with other features?

### For Process Documentation:
- Operational procedures?
- Development workflows?
- Deployment processes?
- Troubleshooting guides?

## Step 3: Project Analysis Phase

Execute systematic discovery:

### A. Structure Discovery

Analyze project structure:
```markdown
Let me analyze your project structure...

**Directories Found:**
- `/src` - Source code
- `/tests` - Test suites
- `/docs` - Existing documentation
- `/config` - Configuration files
- [etc.]

**Key Files Identified:**
- Entry points
- Configuration files
- Core modules
- Test files
```

### B. Technology Stack Analysis

Identify technologies:
```markdown
**Technology Stack Detected:**
- **Languages**: [Languages and versions]
- **Frameworks**: [Framework versions]
- **Database**: [Database type and version]
- **Infrastructure**: [Cloud services, containers, etc.]
- **Build Tools**: [Build system]
- **Dependencies**: [Key dependencies]
```

### C. Existing Documentation Review

Assess current state:
```markdown
**Documentation Audit:**
- README: [Exists? Quality? Coverage?]
- API Docs: [Status]
- Architecture Docs: [Status]
- Code Comments: [Coverage level]
- Inline Documentation: [Quality]

**Gaps Identified:**
1. [Missing documentation area]
2. [Outdated documentation]
3. [Unclear specifications]
```

### D. Pattern Analysis

Identify architectural patterns:
```markdown
**Architectural Patterns Found:**
- **Design Patterns**: [MVC, Repository, Factory, etc.]
- **Architecture Style**: [Microservices, Monolith, Layered, etc.]
- **Communication Patterns**: [REST, Events, Message Queue, etc.]
- **Data Patterns**: [ORM, Repository, CQRS, etc.]
```

## Step 4: Elicitation Questions

Present targeted questions based on analysis:

### Technical Understanding Questions
```markdown
## Technical Clarity Questions

**System Purpose & Scope**
1. What is the primary purpose of this system?
2. Who are the main users/consumers?
3. What problems does it solve?

**Architecture & Design**
4. Are there any architecture decisions you want documented?
5. What are the main components and how do they interact?
6. Are there any critical data flows?

**Technical Constraints**
7. What are the key technical constraints or requirements?
8. What are the performance/scale requirements?
9. Are there specific security or compliance requirements?

**Integration & Dependencies**
10. What external systems does this integrate with?
11. What are the critical dependencies?
12. How does data flow between systems?

**Deployment & Operations**
13. How is this deployed?
14. What's the operational model?
15. What monitoring/observability exists?
```

Adapt questions based on documentation type selected in Step 1.

## Step 5: Deep Analysis

After gathering information, perform:

### A. Code Analysis
- Read key source files
- Understand component relationships
- Identify design patterns
- Map data flows

### B. Configuration Analysis
- Review config files
- Understand environment setup
- Document configuration options
- Identify deployment configurations

### C. Test Analysis
- Review test structure
- Understand test coverage
- Document test scenarios
- Identify integration points

### D. Dependency Analysis
- Map external dependencies
- Understand integration points
- Document API contracts
- Identify version requirements

## Step 6: Documentation Generation

Create comprehensive documentation using this template:

```markdown
# [Project Name] Documentation

## Overview

### Purpose
[Clear statement of what this system does and why it exists]

### Key Features
- [Feature 1]
- [Feature 2]
- [Feature 3]

### Target Audience
[Who uses this system and how]

## Architecture

### System Architecture

[High-level architecture description]

**Architecture Diagram:**
[ASCII or Mermaid diagram showing main components]

### Key Components

#### [Component 1 Name]
- **Purpose**: [What this component does]
- **Responsibilities**: [Key responsibilities]
- **Location**: [File path or module]
- **Dependencies**: [What it depends on]
- **Dependents**: [What depends on it]

#### [Component 2 Name]
...

### Data Flow

[Description of how data flows through the system]

**Data Flow Diagram:**
[Diagram showing data flow]

### Technology Stack

- **Language**: [Language + version]
- **Framework**: [Framework + version]
- **Database**: [Database + version]
- **Infrastructure**: [Cloud/deployment platform]
- **Key Libraries**:
  - [Library 1]: [Purpose]
  - [Library 2]: [Purpose]

## Getting Started

### Prerequisites
- [Requirement 1]
- [Requirement 2]

### Installation

```bash
# Installation steps
[step 1]
[step 2]
```

### Configuration

[Configuration instructions]

**Environment Variables:**
- `VAR_NAME`: [Description]
- `VAR_NAME`: [Description]

**Configuration Files:**
- `config/file.yml`: [Purpose and key settings]

### Running Locally

```bash
# Commands to run locally
[command]
```

## API Documentation

### Endpoints

#### [Endpoint 1]
```
[HTTP_METHOD] /path/to/endpoint
```

**Purpose**: [What this endpoint does]

**Request:**
```json
{
  "field": "value"
}
```

**Response:**
```json
{
  "field": "value"
}
```

**Error Codes:**
- `400`: [Description]
- `404`: [Description]

### Authentication
[Authentication method and details]

## Component Deep Dives

### [Component Name]

**Overview**: [Detailed description]

**Key Classes/Modules:**
- `ClassName`: [Purpose and responsibilities]

**Workflows:**

1. **[Workflow Name]**
   - Step 1: [Description]
   - Step 2: [Description]

**Sequence Diagram:**
[Diagram showing interaction sequence]

## Data Model

### Entities

#### [Entity 1]
```
Entity Name
- field1: type [description]
- field2: type [description]
```

**Relationships:**
- [Relationship description]

### Database Schema

[Schema description or diagram]

## Integration Points

### [External System 1]
- **Purpose**: [Why integrated]
- **Type**: [REST, GraphQL, Message Queue, etc.]
- **Authentication**: [Auth method]
- **Data Exchanged**: [What data flows]
- **Error Handling**: [How failures are handled]

## Deployment

### Architecture

[Deployment architecture description]

**Deployment Diagram:**
[Infrastructure diagram]

### Environments

- **Development**: [Description]
- **Staging**: [Description]
- **Production**: [Description]

### Deployment Process

1. [Step 1]
2. [Step 2]
3. [Step 3]

### Infrastructure

- **Hosting**: [Platform]
- **Database**: [Database details]
- **Caching**: [Caching layer]
- **Storage**: [Storage solution]

## Operations

### Monitoring

- **Metrics**: [What's monitored]
- **Logging**: [Logging approach]
- **Alerts**: [Alert conditions]

### Troubleshooting

#### [Common Issue 1]
- **Symptoms**: [How it presents]
- **Cause**: [Root cause]
- **Resolution**: [How to fix]

### Maintenance

[Routine maintenance tasks]

## Development

### Project Structure

```
/project-root
  /src          - [Description]
  /tests        - [Description]
  /config       - [Description]
  /docs         - [Description]
```

### Development Workflow

1. [Step 1]
2. [Step 2]

### Testing

**Test Structure:**
- Unit Tests: [Location and approach]
- Integration Tests: [Location and approach]
- E2E Tests: [Location and approach]

**Running Tests:**
```bash
[test command]
```

### Code Standards

[Coding standards and conventions]

## Security

### Authentication & Authorization
[Auth approach]

### Data Protection
[How sensitive data is protected]

### Security Considerations
- [Security practice 1]
- [Security practice 2]

## Performance

### Performance Characteristics
- [Metric 1]: [Expected value]
- [Metric 2]: [Expected value]

### Optimization
[Performance optimization strategies]

### Scalability
[Scaling approach]

## Appendices

### Glossary
- **Term**: Definition

### References
- [Reference 1]
- [Reference 2]

### Change Log
- **[Date]**: [Changes]
```

## Step 7: Documentation Delivery

After generating documentation:

1. **Save Documentation**
   ```markdown
   I've created comprehensive documentation. Let me save this to your project:

   **Created Files:**
   - `/docs/architecture/README.md` - Main architecture documentation
   - `/docs/architecture/diagrams/` - System diagrams
   - `/docs/api/README.md` - API documentation
   - `/docs/getting-started.md` - Quick start guide
   ```

2. **Provide Overview**
   ```markdown
   ## Documentation Summary

   **What's Included:**
   - [Section 1]: [Description]
   - [Section 2]: [Description]

   **Key Highlights:**
   - [Highlight 1]
   - [Highlight 2]

   **Documentation Structure:**
   [Tree view of created documentation]
   ```

## Step 8: Refinement Options

Offer follow-up options:

```markdown
## Next Steps

Would you like me to:

1. **Expand Specific Sections** - Deep dive into particular components
2. **Add More Diagrams** - Create additional visualizations
3. **Create Developer Guides** - Add how-to guides for common tasks
4. **Add Examples** - Include code examples and usage scenarios
5. **Create API Guides** - Detailed API integration guides
6. **Update README** - Create/update project README
7. **Generate Onboarding Doc** - New developer onboarding guide
8. **Add Troubleshooting** - Expand troubleshooting section

Or if the documentation is complete, I can help you with something else!
```

## Documentation Focus Strategies

### For Large Projects:
1. Start with high-level overview
2. Create component summaries
3. Deep dive on request
4. Link detailed documentation
5. Use diagrams heavily

### For Complex Systems:
1. Break into subsystems
2. Document interactions clearly
3. Use sequence diagrams
4. Provide integration examples
5. Document decision rationale

### For API-Heavy Projects:
1. OpenAPI/Swagger specification
2. Example requests/responses
3. Integration scenarios
4. Error handling guide
5. Authentication details

### For Microservices:
1. Service catalog
2. Service interactions
3. Data flow between services
4. API contracts
5. Deployment topology

## Quality Checklist

Before delivering documentation, verify:

- [ ] Clear purpose and overview
- [ ] Architecture clearly explained
- [ ] Key components documented
- [ ] Getting started instructions work
- [ ] API endpoints documented (if applicable)
- [ ] Data model explained (if applicable)
- [ ] Deployment process documented
- [ ] Integration points covered
- [ ] Troubleshooting guidance included
- [ ] Code examples provided
- [ ] Diagrams included where helpful
- [ ] Technical terms defined
- [ ] Target audience appropriate
- [ ] Next steps clear

## Success Criteria

### Documentation Quality:
- Technically accurate
- Appropriately detailed
- Well-structured
- Includes diagrams
- Covers key workflows
- Actionable information

### User Experience:
- Easy to navigate
- Clear and concise
- Helpful examples
- Addresses common questions
- Enables self-service

### Completeness:
- All major components covered
- Integration points documented
- Operational aspects included
- Development workflow explained
- Troubleshooting guidance provided

## Common Pitfalls to Avoid

1. **Too Generic** - Documentation must be specific to this project
2. **Too Much Code** - Focus on concepts, not code dumps
3. **No Diagrams** - Visual aids are crucial for understanding
4. **Missing Context** - Explain why, not just what
5. **Outdated Info** - Verify accuracy during documentation
6. **No Examples** - Always include concrete examples
7. **Wrong Audience** - Match technical level to readers
8. **No Maintenance Plan** - Documentation must stay current
