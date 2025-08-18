# Task: Create Project Brief

## Purpose
Create a comprehensive Project Brief that serves as the foundation for all product development activities. The brief should capture the essential context, goals, and constraints before moving to detailed PRD creation.

## Prerequisites
- Initial product concept or idea
- Basic understanding of target market
- Stakeholder input gathered
- Problem identification completed

## Workflow

### 1. Context Gathering
```bash
# Search for existing research
./.vector_db/kb search "market research" --collection documentation
./.vector_db/kb search "user research" --collection documentation
./.vector_db/kb search "competitive analysis" --collection documentation
```

**Collect Information:**
- [ ] Existing market research
- [ ] User research findings
- [ ] Competitive analysis data
- [ ] Business context and constraints
- [ ] Stakeholder requirements

### 2. Template Selection
Use interactive YAML template: `.claude/resources/pm/templates/project-brief.yaml`

**Mode Selection:**
- **Interactive Mode**: Work through each section with stakeholder input
- **YOLO Mode**: Generate complete draft for review

### 3. Section-by-Section Creation

#### Executive Summary
- Product concept (1-2 sentences)
- Primary problem being solved
- Target market identification
- Key value proposition

#### Problem Statement
- Current state and pain points
- Quantified impact of problem
- Why existing solutions fall short
- Urgency and business importance

#### Proposed Solution
- Core concept and approach
- Key differentiators
- Success factors
- High-level product vision

#### Target Users
- Primary user segment profile
- Secondary segments (if applicable)
- User behaviors and workflows
- Needs and pain points
- Goals and objectives

#### Goals & Success Metrics
- Business objectives (SMART)
- User success metrics
- Key Performance Indicators
- Measurement approach

#### MVP Scope
- Core features (must-have)
- Out-of-scope items
- Success criteria for MVP
- Validation approach

#### Post-MVP Vision
- Phase 2 features
- Long-term vision (1-2 years)
- Expansion opportunities

#### Technical Considerations
- Platform requirements
- Technology preferences
- Architecture considerations
- Integration needs

#### Constraints & Assumptions
- Budget limitations
- Timeline constraints
- Resource availability
- Technical constraints
- Key assumptions

#### Risks & Open Questions
- Business risks
- Technical risks
- Market risks
- Open questions
- Research needs

### 4. Validation & Review

**Quality Check:**
- [ ] Problem clearly defined and validated
- [ ] Solution approach sound
- [ ] Target users specific
- [ ] MVP scope truly minimal
- [ ] Success metrics measurable
- [ ] Constraints realistic
- [ ] Risks identified

**Stakeholder Review:**
- [ ] Business stakeholders approve
- [ ] Technical feasibility confirmed
- [ ] Market assumptions validated
- [ ] Resource allocation approved

### 5. Finalization

**Output Management:**
- Save to: `/docs/requirements/project-brief.md`
- Update knowledge base:
  ```bash
  ./.vector_db/kb index --path /docs/requirements/project-brief.md
  ```

**Handoff Preparation:**
- Brief summary for PRD creation
- Key requirements highlighted
- Open questions documented
- Next steps defined

## Elicitation Techniques

### For Problem Definition
- Five Whys analysis
- Problem tree analysis
- Impact assessment
- Root cause analysis

### For Solution Design
- Solution brainstorming
- Competitive benchmarking
- Technology scanning
- Feasibility assessment

### For Scope Definition
- MoSCoW prioritization
- Value vs. Effort analysis
- Kano model analysis
- MVP canvas

### For Risk Management
- SWOT analysis
- Risk assessment matrix
- Scenario planning
- Mitigation brainstorming

## Success Criteria
- [ ] Problem clearly articulated with evidence
- [ ] Solution approach sound and differentiated
- [ ] Target users specifically defined
- [ ] MVP scope minimal but viable
- [ ] Success metrics quantifiable
- [ ] Technical approach feasible
- [ ] Constraints and risks documented
- [ ] Stakeholder alignment achieved
- [ ] Foundation ready for PRD creation

## Common Pitfalls
- **Too broad scope**: Keep MVP truly minimal
- **Vague metrics**: Make success criteria specific
- **Undefined users**: Be specific about target segments
- **Missing constraints**: Document all limitations
- **Weak differentiation**: Clearly articulate unique value
- **Unvalidated assumptions**: Challenge key assumptions

## Integration Points
- **Ana (Analyst)**: Leverage market research and competitive analysis
- **Ulysses (UX)**: Incorporate user research and journey insights
- **Archie (Architect)**: Validate technical feasibility and constraints
- **Sam (SM)**: Understand delivery constraints and team capacity

## Next Steps
After completion:
1. **PRD Creation**: Use brief as foundation for detailed requirements
2. **Architecture Planning**: Provide technical context for system design
3. **UX Research**: Guide user research and design activities
4. **Sprint Planning**: Inform initial epic and story creation