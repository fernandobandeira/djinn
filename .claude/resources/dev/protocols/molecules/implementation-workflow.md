# Implementation Workflow Protocol

## Overview
Comprehensive multi-phase workflow for story implementation from intake to deployment.

## Workflow Phases

```yaml
workflow:
  name: implementation_workflow
  version: "1.0"
  phases:
    - intake
    - analysis
    - planning
    - implementation
    - testing
    - review
    - optimization
    - deployment
```

## Phase 1: Story Intake

### Inputs
- Story file from SM: `/docs/requirements/stories/[epic].[story].story.md`
- Architecture context embedded in story
- Acceptance criteria and test requirements

### Process
```bash
# Load story context
THEN load /docs/requirements/stories/[epic].[story].story.md

# Extract key elements
- User story statement
- Acceptance criteria
- Technical requirements
- Security considerations
- Previous story insights
```

### Outputs
- Parsed story requirements
- Identified dependencies
- Risk assessment

## Phase 2: Analysis

### Technical Analysis
```yaml
analysis_checklist:
  - component_impact: Identify affected components
  - dependency_check: External and internal dependencies
  - complexity_score: Calculate using complexity-estimator
  - pattern_match: Find similar implementations
  - risk_factors: Security, performance, scalability
```

### Knowledge Base Search
```bash
# Search for relevant patterns
./.vector_db/kb search "[feature_type] implementation" --collection code
./.vector_db/kb search "[technology] pattern" --collection architecture
./.vector_db/kb search "similar feature" --collection documentation
```

## Phase 3: Planning

### Implementation Plan Structure
```markdown
## Implementation Plan: [Story Title]

### Approach
- Technical strategy
- Design patterns to use
- Architecture alignment

### Components
1. **Frontend**
   - Components to create/modify
   - State management approach
   - UI/UX considerations

2. **Backend**
   - API endpoints
   - Business logic
   - Data models

3. **Database**
   - Schema changes
   - Migrations
   - Indexes

### Task Breakdown
- [ ] Task 1: [Description] (Est: Xh)
  - [ ] Subtask 1.1
  - [ ] Subtask 1.2
- [ ] Task 2: [Description] (Est: Xh)
  - [ ] Subtask 2.1
  - [ ] Subtask 2.2

### Timeline
- Total estimate: X story points
- Development: X days
- Testing: X days
- Review: X days
```

## Phase 4: Implementation

### TDD Approach
```yaml
tdd_cycle:
  red:
    - Write failing test for AC
    - Verify test fails correctly
    - Commit test
  green:
    - Write minimal code to pass
    - Run test suite
    - Verify passing
  refactor:
    - Improve code quality
    - Maintain test green
    - Apply patterns
```

### Code Organization
```bash
# Create feature structure
src/
└── features/
    └── [feature_name]/
        ├── models/
        ├── services/
        ├── controllers/
        ├── views/
        └── tests/
```

### Progress Tracking
```yaml
implementation_state:
  current_task: "Task 1.2"
  completed_tasks:
    - "Task 1.1"
  progress: 25
  blockers: []
  commits:
    - hash: "abc123"
      message: "feat: Add user authentication"
```

## Phase 5: Testing

### Test Levels
```yaml
test_pyramid:
  unit:
    coverage_target: 80
    focus: "Business logic, utilities"
    
  integration:
    coverage_target: 60
    focus: "API endpoints, database"
    
  e2e:
    coverage_target: "Critical paths"
    focus: "User journeys"
    
  performance:
    targets:
      response_time: "< 200ms"
      throughput: "> 1000 req/s"
```

### Test Execution
```bash
# Run test suite
npm test

# Check coverage
npm run test:coverage

# Run specific test level
npm run test:unit
npm run test:integration
npm run test:e2e
```

## Phase 6: Review (MANDATORY QA-REVIEWER DELEGATION)

### CRITICAL: This phase is MANDATORY and AUTOMATIC
**MUST delegate to qa-reviewer - NO EXCEPTIONS**

### Automatic QA-Reviewer Invocation
```yaml
mandatory_review:
  trigger: "ALWAYS after Phase 5 (Testing)"
  delegate_to: qa-reviewer
  instruction: "IMPORTANT actively critiques code and finds issues"
  iteration_limit: 5
  approval_required: true
```

### Review Process
1. **AUTOMATIC**: Delegate to qa-reviewer for critique
2. Receive critique_report with issues
3. Apply all critical/high severity fixes
4. Re-delegate to qa-reviewer for validation
5. Repeat until approved or max iterations

### Self-Review Checklist (Pre-QA)
- [ ] Code follows standards
- [ ] All tests passing
- [ ] Documentation updated
- [ ] No console logs
- [ ] Error handling complete
- [ ] Security considerations addressed

### Peer Review Request
```markdown
## Review Request: [Story Title]

### Changes Summary
- What was implemented
- Key decisions made
- Patterns used

### Testing
- Test coverage: X%
- Test types: unit, integration
- Manual testing completed

### Areas of Focus
- Specific areas needing attention
- Complex logic sections
- Performance considerations

### Checklist
- [ ] Functionality matches requirements
- [ ] Code quality acceptable
- [ ] Tests comprehensive
- [ ] Documentation sufficient
```

## Phase 7: Optimization

### Performance Analysis
```yaml
optimization_areas:
  database:
    - Query optimization
    - Index usage
    - N+1 prevention
    
  frontend:
    - Bundle size
    - Render performance
    - Asset optimization
    
  backend:
    - Response time
    - Memory usage
    - Caching strategy
```

### Optimization Techniques
```bash
# Profile application
npm run profile

# Analyze bundle
npm run analyze

# Database query analysis
npm run db:analyze
```

## Phase 8: Deployment

### Pre-Deployment Checklist
- [ ] All tests passing
- [ ] Code review approved
- [ ] Documentation complete
- [ ] Migration scripts ready
- [ ] Feature flags configured
- [ ] Monitoring setup

### Deployment Process
```yaml
deployment:
  staging:
    - Deploy to staging
    - Run smoke tests
    - Verify functionality
    
  production:
    - Create deployment PR
    - Get approval
    - Deploy with rollback plan
    - Monitor metrics
    
  post_deployment:
    - Verify in production
    - Monitor error rates
    - Check performance metrics
```

## Workflow State Management

### State Structure
```yaml
workflow_state:
  story_id: "2.1"
  current_phase: "implementation"
  phase_status:
    intake: "completed"
    analysis: "completed"
    planning: "completed"
    implementation: "in_progress"
    testing: "pending"
    review: "pending"
    optimization: "pending"
    deployment: "pending"
    
  metrics:
    start_time: "2024-01-10T09:00:00Z"
    story_points: 5
    actual_hours: 12
    test_coverage: 0
    
  artifacts:
    plan: "/docs/plans/2.1-plan.md"
    tests: "/src/features/auth/tests/"
    documentation: "/docs/features/auth.md"
```

## Quality Gates

### Phase Transitions
Each phase must meet criteria before proceeding:

```yaml
quality_gates:
  analysis_to_planning:
    - Dependencies identified
    - Complexity estimated
    - Risks assessed
    
  planning_to_implementation:
    - Plan reviewed
    - Tasks broken down
    - Estimates provided
    
  implementation_to_testing:
    - Code complete
    - Unit tests written
    - Linting passed
    
  testing_to_review:
    - All tests passing
    - Coverage target met
    - Integration tested
    
  review_to_optimization:
    - **qa-reviewer approval_status: approved (MANDATORY)**
    - All critical issues fixed
    - Review feedback addressed
    - Approval received
    
  optimization_to_deployment:
    - Performance targets met
    - Security scan passed
    - Documentation complete
```

## Integration Points

### Collaboration with Other Agents
```yaml
integrations:
  sm_agent:
    - Receive story specifications
    - Report progress updates
    - Request clarifications
    
  architect_agent:
    - Consult on design patterns
    - Verify architecture alignment
    - Request architecture updates
    
  qa_reviewer:
    - **MANDATORY: Perform critical code review after implementation**
    - Generate test scenarios
    - Validate test coverage
    - Find and report issues with severity levels
    - Provide actionable fixes
    - Block or approve implementation
    
  kb_analyst:
    - Search for patterns
    - Find similar implementations
    - Store new patterns
```

## Error Recovery

### Common Issues and Recovery
```yaml
error_handling:
  blocked_by_dependency:
    action: "Identify alternative approach or escalate"
    
  failing_tests:
    action: "Debug, fix, or modify approach"
    
  performance_issues:
    action: "Profile, optimize, or redesign"
    
  review_rejection:
    action: "Address feedback and resubmit"
```

## Metrics and Reporting

### Progress Metrics
```yaml
metrics:
  velocity:
    - Story points completed
    - Actual vs estimated time
    
  quality:
    - Test coverage percentage
    - Defect density
    - Code review iterations
    
  efficiency:
    - Time per story point
    - Rework percentage
    - Automation coverage
```

## Continuous Improvement

### Retrospective Points
- What went well
- What could improve
- Patterns to document
- Tools to enhance
- Process optimizations