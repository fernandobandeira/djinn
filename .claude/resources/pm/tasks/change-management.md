# Task: Change Management & Course Correction

## Purpose
Guide structured response to change triggers using systematic analysis. Handle mid-sprint changes, scope adjustments, and requirement updates while maintaining project integrity and stakeholder alignment.

## When to Use
- Requirements change mid-sprint
- New market information affects scope
- Technical constraints discovered
- Stakeholder priorities shift
- Resource availability changes
- Timeline pressures emerge

## Prerequisites
- Current PRD and epic definitions
- Understanding of change trigger
- Access to project artifacts
- Stakeholder availability for decisions

## Workflow

### 1. Change Assessment & Setup

**Initial Setup:**
- [ ] Confirm change trigger and context
- [ ] Gather relevant project artifacts
- [ ] Identify affected stakeholders
- [ ] Set interaction mode (Incremental/YOLO)

**Artifact Collection:**
```bash
# Gather current project state
./.vector_db/kb search "PRD" --collection documentation
./.vector_db/kb search "epic" --collection documentation
./.vector_db/kb search "architecture" --collection architecture
./.vector_db/kb search "sprint" --collection documentation
```

### 2. Systematic Change Analysis

#### Change Context Analysis
- **Trigger Source**: What initiated this change?
- **Urgency Level**: How quickly must this be addressed?
- **Scope of Impact**: Which areas are affected?
- **Stakeholder Impact**: Who is affected and how?

#### Impact Assessment Matrix
| Area | Impact Level | Description | Mitigation Required |
|------|-------------|-------------|-------------------|
| PRD Requirements | High/Med/Low | Specific impacts | Yes/No |
| Epic Structure | High/Med/Low | Specific impacts | Yes/No |
| Current Sprint | High/Med/Low | Specific impacts | Yes/No |
| Architecture | High/Med/Low | Specific impacts | Yes/No |
| Timeline | High/Med/Low | Specific impacts | Yes/No |
| Resources | High/Med/Low | Specific impacts | Yes/No |

#### Epic & Story Impact Analysis
For each affected epic:
- **Current Status**: In Progress/Planned/Complete
- **Stories Affected**: List specific stories
- **Dependency Impact**: Upstream/downstream effects
- **Effort Change**: Increase/decrease/neutral
- **Priority Change**: Higher/lower/same

### 3. Solution Path Evaluation

#### Option Analysis
**Option 1: Scope Adjustment**
- Reduce current scope
- Move features to future phases
- Maintain timeline
- Impact on value delivery

**Option 2: Timeline Extension**
- Maintain full scope
- Extend delivery dates
- Resource implications
- Stakeholder acceptance

**Option 3: Resource Augmentation**
- Add team members
- Bring in specialists
- Budget implications
- Onboarding time

**Option 4: Technical Pivot**
- Change implementation approach
- Reduce technical complexity
- Architectural implications
- Risk assessment

#### Decision Criteria
- **Business Value**: Impact on core objectives
- **Technical Feasibility**: Can we execute this?
- **Resource Availability**: Do we have capacity?
- **Timeline Constraints**: Does this fit schedule?
- **Risk Level**: What could go wrong?
- **Stakeholder Acceptance**: Will stakeholders approve?

### 4. Recommended Path Forward

Based on analysis, recommend specific approach:

#### High-Level Strategy
- **Primary Approach**: Chosen option with rationale
- **Fallback Option**: If primary fails
- **Success Criteria**: How to measure success
- **Risk Mitigation**: Key risks and responses

#### Specific Actions Required
1. **PRD Updates**: Specific sections to modify
2. **Epic Changes**: Additions/removals/modifications
3. **Story Adjustments**: New/modified/removed stories
4. **Architecture Impact**: Technical changes needed
5. **Timeline Adjustments**: New milestones/deadlines

### 5. Change Proposal Creation

#### Sprint Change Proposal Document
```markdown
# Sprint Change Proposal

## Change Summary
- **Trigger**: What caused this change
- **Impact**: High-level effects
- **Recommended Action**: Chosen approach

## Analysis Results
- **Scope Impact**: Detailed scope changes
- **Timeline Impact**: Schedule adjustments
- **Resource Impact**: Team/budget changes
- **Risk Assessment**: Key risks and mitigation

## Specific Proposed Edits

### PRD Changes
- Section X: Change from [old] to [new]
- Section Y: Add new requirement [details]

### Epic Modifications
- Epic 1: [specific changes]
- Epic 2: [specific changes]

### Story Updates
- Story A.1: [modifications]
- Story B.2: [new acceptance criteria]

### Architecture Updates
- Component X: [changes needed]
- Integration Y: [new requirements]

## Implementation Plan
1. **Immediate Actions**: What to do now
2. **Sprint Adjustments**: Current sprint changes
3. **Future Planning**: Longer-term implications
4. **Communication**: Who to inform when

## Success Metrics
- How to measure change success
- Validation criteria
- Review checkpoints
```

### 6. Stakeholder Review & Approval

#### Review Process
- **Technical Review**: Engineering team validation
- **Business Review**: Product/business stakeholder approval
- **Design Review**: UX team impact assessment
- **Operations Review**: Deployment/support implications

#### Approval Criteria
- [ ] Technical feasibility confirmed
- [ ] Business value maintained/improved
- [ ] Resource allocation approved
- [ ] Timeline acceptable to stakeholders
- [ ] Risk level acceptable
- [ ] Communication plan clear

### 7. Implementation & Handoff

#### Change Implementation
Based on approved proposal:
- [ ] Update PRD sections
- [ ] Modify epic definitions
- [ ] Adjust story acceptance criteria
- [ ] Update architecture documents
- [ ] Communicate changes to team

#### Handoff Scenarios
**For Minor Changes (Direct Implementation)**
- Update project artifacts
- Communicate to development team
- Monitor implementation

**For Major Changes (Replanning Required)**
- Hand off to PM for fundamental replanning
- Provide change proposal as input
- Schedule replanning sessions
- Coordinate with Architect for technical replanning

## Change Types & Responses

### Scope Changes
- **Addition**: New features/requirements
- **Removal**: Cut features/requirements
- **Modification**: Change existing features

### Resource Changes
- **Team Changes**: Members added/removed
- **Skill Changes**: New expertise needed
- **Budget Changes**: Funding adjustments

### Timeline Changes
- **Acceleration**: Faster delivery needed
- **Extension**: More time available
- **Milestone Shifts**: Date changes

### Technical Changes
- **Platform Changes**: Technology shifts
- **Architecture Changes**: Structural modifications
- **Integration Changes**: External system updates

## Elicitation Techniques for Change

### Change Impact Analysis
- **Ripple Effect Mapping**: Trace change impacts
- **Stakeholder Impact Grid**: Who's affected how
- **Risk Assessment Matrix**: What could go wrong
- **Option Evaluation**: Compare alternatives

### Solution Generation
- **Brainstorming**: Generate multiple options
- **Constraint Relaxation**: Question assumptions
- **Trade-off Analysis**: Evaluate compromises
- **Creative Problem Solving**: Think outside box

### Decision Support
- **Decision Tree**: Map decision points
- **Pros/Cons Analysis**: Evaluate options
- **Impact vs Effort**: Prioritize changes
- **ROI Analysis**: Calculate value

## Success Criteria
- [ ] Change properly analyzed
- [ ] Impact fully understood
- [ ] Solution path validated
- [ ] Stakeholder agreement achieved
- [ ] Implementation plan clear
- [ ] Risks identified and mitigated
- [ ] Project integrity maintained

## Common Pitfalls
- **Knee-jerk Reactions**: Rushing to solutions
- **Incomplete Analysis**: Missing impact areas
- **Poor Communication**: Not involving stakeholders
- **Scope Creep**: Adding unnecessary changes
- **Risk Blindness**: Ignoring potential problems
- **Change Paralysis**: Over-analyzing without action

## Integration Points
- **Sam (SM)**: Sprint planning and story management
- **Archie (Architect)**: Technical impact assessment
- **Ulysses (UX)**: Design impact evaluation
- **Ana (Analyst)**: Business impact analysis
- **Dave (Developer)**: Implementation feasibility