# Task: Create Advanced PRD

## Purpose
Create a comprehensive Product Requirements Document using advanced elicitation techniques and interactive YAML-driven templates. This PRD will serve as the definitive guide for architecture, design, and development teams.

## Prerequisites
- Project Brief completed (`/docs/requirements/project-brief.md`)
- Market research available
- User research conducted
- Stakeholder alignment achieved

## Workflow

### 1. Foundation Setup
```bash
# Verify project brief exists
[ -f "/docs/requirements/project-brief.md" ] && echo "Brief found" || echo "Create brief first"

# Load existing research
./.vector_db/kb search "market research" --collection documentation
./.vector_db/kb search "user research" --collection documentation
./.vector_db/kb search "technical constraints" --collection documentation
```

**Foundation Check:**
- [ ] Project brief available and complete
- [ ] Market research findings accessible
- [ ] User research insights available
- [ ] Technical preferences documented
- [ ] Stakeholder requirements clear

### 2. Template-Driven Creation
Use advanced YAML template: `.claude/resources/pm/templates/prd-advanced.yaml`

**Workflow Mode:**
- **Interactive**: Section-by-section with elicitation
- **YOLO**: Complete draft generation

### 3. Section Development

#### Goals & Background Context
- Extract from project brief
- Validate and refine goals
- Document context and rationale
- Maintain change log

#### Requirements Definition
**Functional Requirements:**
- Use prefix FR (FR1, FR2, etc.)
- Clear, testable statements
- User-focused descriptions
- Priority assignment

**Non-Functional Requirements:**
- Use prefix NFR (NFR1, NFR2, etc.)
- Performance criteria
- Security requirements
- Scalability needs
- Compliance requirements

#### UI/UX Design Goals
- Overall UX vision
- Key interaction paradigms
- Core screens identification
- Accessibility requirements (WCAG AA/AAA)
- Branding guidelines
- Target platform specifications

#### Technical Assumptions
- Repository structure (Monorepo/Polyrepo)
- Service architecture (Monolith/Microservices/Serverless)
- Testing approach (Unit/Integration/Full Pyramid)
- Technology stack preferences
- Infrastructure requirements

#### Epic List Creation
**Epic Principles:**
- Sequential delivery order
- End-to-end functionality
- Deployable increments
- Foundation first (Epic 1)
- Value-driven prioritization

**Epic Structure:**
```
Epic 1: Foundation & Infrastructure
- Project setup, CI/CD, core services
- Initial functionality (health check, basic pages)

Epic 2: Core Business Logic
- Primary domain objects
- Essential business rules
- Basic CRUD operations

Epic 3: User Workflows
- Complete user journeys
- Integration between components
- Advanced functionality

Epic 4: Analytics & Optimization
- Reporting capabilities
- Performance optimization
- Advanced features
```

#### Epic Detailed Planning
For each epic:
- **Goal Statement**: 2-3 sentences on objective and value
- **Story Breakdown**: Vertical slices of functionality
- **Story Sizing**: 2-4 hour AI agent sessions
- **Acceptance Criteria**: Clear, testable conditions
- **Dependencies**: Prerequisites and integrations

### 4. Advanced Elicitation Techniques

#### Requirements Elicitation
- **Expand/Contract**: Adjust detail for audience
- **Critique and Refine**: Review for gaps and improvements
- **Risk Assessment**: Identify potential issues
- **Stakeholder Perspective**: Multiple viewpoint analysis

#### Scope Validation
- **MVP Challenge**: "What's truly minimum?"
- **Value vs Effort**: Prioritization matrix
- **YAGNI Application**: Remove unnecessary features
- **Competitive Analysis**: Differentiation focus

#### Technical Validation
- **Feasibility Check**: Engineering review
- **Architecture Impact**: System design implications
- **Performance Requirements**: Load and response targets
- **Security Assessment**: Threat model review

### 5. Quality Assurance

#### PM Checklist Execution
Run comprehensive validation: `.claude/resources/pm/checklists/pm-validation-checklist.md`

**Validation Categories:**
- Problem definition & context
- MVP scope definition
- User experience requirements
- Functional requirements
- Non-functional requirements
- Epic & story structure
- Technical guidance
- Stakeholder alignment
- Risk management
- Readiness assessment

#### Stakeholder Review
- [ ] Business stakeholders approve
- [ ] Technical leads validate feasibility
- [ ] UX team confirms design direction
- [ ] Security team approves requirements
- [ ] Operations team understands deployment needs

### 6. Finalization & Handoff

**Output Management:**
- Save to: `/docs/requirements/prd.md`
- Update knowledge base:
  ```bash
  ./.vector_db/kb index --path /docs/requirements/prd.md
  ```

**Handoff Documentation:**
- **UX Handoff**: Design requirements and constraints
- **Architecture Handoff**: Technical specifications and constraints
- **Development Handoff**: Ready for epic breakdown and sprint planning

## Elicitation Method Application

### Interactive Sections
When `elicit: true` in template:
1. Present drafted content
2. Provide detailed rationale
3. Offer 9 elicitation options:
   - "Proceed to next section"
   - 8 methods from elicitation-methods.md
4. Wait for user response
5. Apply selected method
6. Iterate until satisfactory

### Advanced Techniques
- **Tree of Thoughts**: Complex requirement breakdown
- **Multi-Persona**: Stakeholder perspective shifts
- **Risk-Opportunity**: Business impact analysis
- **Technical Deep Dive**: Architecture implications
- **User Journey Mapping**: End-to-end experience validation

## Quality Gates

### Completeness Check
- [ ] All template sections completed
- [ ] Requirements traceable to business goals
- [ ] Stories sized for AI agent execution
- [ ] Acceptance criteria testable
- [ ] Dependencies mapped
- [ ] Risks identified and mitigated

### Validation Results
- Pass: 90%+ checklist items complete
- Partial: 60-89% complete, requires refinement
- Fail: <60% complete, significant work needed

### Readiness Criteria
- [ ] Problem clearly defined with evidence
- [ ] Solution approach validated
- [ ] MVP scope minimal but viable
- [ ] Technical approach feasible
- [ ] Success metrics measurable
- [ ] Team alignment achieved

## Success Criteria
- [ ] Comprehensive requirements documented
- [ ] Epic structure logically sequenced
- [ ] Stories appropriately sized
- [ ] Technical constraints clear
- [ ] Stakeholder approval obtained
- [ ] Ready for architecture and design phases
- [ ] Foundation for sprint planning established

## Integration Points
- **Archie (Architect)**: Technical requirements and constraints
- **Ulysses (UX)**: Design requirements and user experience
- **Sam (SM)**: Sprint planning and story management
- **Ana (Analyst)**: Market validation and competitive positioning
- **Dave (Developer)**: Technical feasibility and implementation approach

## Next Steps
After PRD completion:
1. **Architecture Design**: Archie creates technical architecture
2. **UX Design**: Ulysses creates wireframes and prototypes
3. **Sprint Planning**: Sam breaks down epics into sprints
4. **Development Planning**: Dave prepares technical approach