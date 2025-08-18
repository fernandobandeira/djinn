# Task: Create Brownfield Epic

## Purpose
Create a focused epic for brownfield enhancements to existing systems. This task handles isolated features or modifications that can be completed within 1-3 stories without requiring full PRD and Architecture documentation.

## When to Use This Task

### Use Brownfield Epic When:
- Enhancement requires 1-3 coordinated stories
- Some design work is needed
- Multiple integration points involved
- Moderate complexity but manageable scope
- Risk to existing system is low-moderate

### Use Single Story Task When:
- Enhancement is single story only
- No design work required
- Follows existing patterns exactly
- Minimal integration complexity

### Use Full PRD Process When:
- Multiple coordinated epics needed
- Significant architectural changes required
- Major integration work involved
- High risk to existing system

## Prerequisites
- Understanding of existing system
- Clear enhancement requirements
- Access to current codebase
- Stakeholder input on requirements

## Workflow

### 1. Existing System Analysis

**System Context Assessment:**
```bash
# Search for existing documentation
./.vector_db/kb search "architecture" --collection architecture
./.vector_db/kb search "current system" --collection documentation
./.vector_db/kb search "integration" --collection code
```

**Required Understanding:**
- [ ] Current system functionality relevant to enhancement
- [ ] Technology stack in affected areas
- [ ] Existing integration patterns
- [ ] Current architecture constraints
- [ ] Performance characteristics
- [ ] Security model

**Enhancement Scope Assessment:**
- [ ] Specific enhancement clearly defined
- [ ] Integration points identified
- [ ] Impact boundaries understood
- [ ] Success criteria established
- [ ] Compatibility requirements clear

### 2. Epic Structure Creation

#### Epic Title Format
`{{Enhancement Name}} - Brownfield Enhancement`

#### Epic Goal (2-3 sentences)
Describe what the epic accomplishes and why it adds value to the existing system.

#### Epic Description Template
```markdown
## Existing System Context
- **Current Functionality**: Brief description of relevant existing features
- **Technology Stack**: Relevant technologies in use
- **Integration Points**: Where new work connects to existing system
- **Architecture Pattern**: Current patterns being followed

## Enhancement Details
- **What's Being Added/Changed**: Clear description of modifications
- **Integration Approach**: How it connects to existing system
- **Value Proposition**: Why this enhancement matters
- **Success Criteria**: Measurable outcomes

## Compatibility Requirements
- [ ] Existing APIs remain unchanged or backward compatible
- [ ] Database schema changes are additive/backward compatible
- [ ] UI changes follow existing design system
- [ ] Performance impact is minimal
- [ ] Security model is maintained

## Risk Assessment
- **Primary Risk**: Main risk to existing system stability
- **Mitigation Strategy**: How risk will be addressed
- **Rollback Plan**: How to safely undo changes
- **Monitoring**: How to detect issues early
```

### 3. Story Breakdown (1-3 Stories)

#### Story Sequencing Principles
- **Story 1**: Often setup/foundation work
- **Story 2**: Core functionality implementation
- **Story 3**: Integration/polish/optimization

#### Story Template
```markdown
### Story {{epic}}.{{number}}: {{Title}}

**User Story:**
As a {{user_type}},
I want {{enhancement_capability}},
So that {{value_received}}.

**Brownfield Context:**
- Integrates with: {{existing_component}}
- Follows pattern: {{existing_pattern}}
- Maintains: {{existing_behavior}}

**Acceptance Criteria:**
1. {{Primary functional requirement}}
2. {{Integration requirement}}
3. {{Compatibility requirement}}
4. Existing {{relevant_functionality}} continues unchanged
5. New functionality follows {{existing_pattern}} pattern
6. No regression in existing functionality verified

**Definition of Done:**
- [ ] Functional requirements met
- [ ] Integration verified
- [ ] Existing functionality regression tested
- [ ] Performance impact assessed
- [ ] Documentation updated
- [ ] Code follows existing standards
```

### 4. Integration Analysis

#### Integration Points Assessment
For each integration point:
- **Type**: API, Database, UI, Service, etc.
- **Complexity**: Low/Medium/High
- **Risk**: Minimal/Moderate/High
- **Testing Strategy**: How to validate integration

#### Compatibility Matrix
| Component | Change Type | Backward Compatible | Risk Level | Mitigation |
|-----------|-------------|-------------------|------------|------------|
| API | Addition | Yes | Low | API versioning |
| Database | Schema addition | Yes | Low | Migration script |
| UI | New component | Yes | Low | Design system |

### 5. Risk Management

#### Risk Assessment Framework
**Technical Risks:**
- Integration complexity
- Performance impact
- Data consistency
- Security implications

**Business Risks:**
- User experience disruption
- Feature interaction conflicts
- Support complexity
- Timeline pressure

#### Mitigation Strategies
- **Feature Flags**: Gradual rollout capability
- **Database Migrations**: Safe schema changes
- **API Versioning**: Backward compatibility
- **Monitoring**: Early issue detection
- **Testing Strategy**: Comprehensive validation

### 6. Definition of Done

#### Epic-Level Criteria
- [ ] All stories completed and accepted
- [ ] Integration testing complete with existing system
- [ ] Performance impact within acceptable limits
- [ ] Security review passed (if applicable)
- [ ] Documentation updated appropriately
- [ ] Rollback plan tested and validated
- [ ] Stakeholder acceptance obtained
- [ ] No regression in existing functionality

#### Quality Gates
- Code coverage maintains existing standards
- Performance benchmarks not degraded
- Security scan shows no new vulnerabilities
- Integration tests pass completely

### 7. Validation Checklist

#### Scope Validation
- [ ] Epic can be completed in 1-3 stories
- [ ] No architectural documentation required
- [ ] Enhancement follows existing patterns
- [ ] Integration complexity is manageable
- [ ] Risk to existing system is acceptable

#### Requirements Completeness
- [ ] Enhancement requirements are clear
- [ ] Integration approach is defined
- [ ] Compatibility requirements specified
- [ ] Success criteria are measurable
- [ ] Rollback approach is feasible

#### Team Readiness
- [ ] Team understands existing system area
- [ ] Required skills are available
- [ ] Testing approach is clear
- [ ] Documentation is accessible

### 8. Handoff Documentation

#### For Story Manager (Sam)
```markdown
**Brownfield Epic Handoff:**

"Please develop detailed user stories for this brownfield epic. Key considerations:

- **Existing System**: Running on {{technology_stack}}
- **Integration Points**: {{list_key_integrations}}
- **Existing Patterns**: {{relevant_patterns_to_follow}}
- **Critical Compatibility**: {{key_compatibility_requirements}}
- **Story Sizing**: Each story must be completable in single development session
- **Verification**: Include verification that existing functionality remains intact

The epic should maintain system integrity while delivering {{epic_goal}}."
```

#### For Development Team (Dave)
- Architecture patterns to follow
- Integration guidelines
- Testing requirements
- Performance expectations
- Security considerations

## Success Criteria
- [ ] Enhancement scope clearly defined and appropriately sized
- [ ] Integration approach respects existing architecture
- [ ] Risk to existing functionality minimized
- [ ] Stories logically sequenced for safe implementation
- [ ] Compatibility requirements clearly specified
- [ ] Rollback plan feasible and documented
- [ ] Team ready to implement with confidence

## Common Pitfalls
- **Scope Creep**: Adding "while we're here" features
- **Pattern Deviation**: Not following existing system patterns
- **Insufficient Testing**: Not validating existing functionality
- **Poor Integration**: Creating tight coupling or breaking abstractions
- **Rollback Ignorance**: Not planning for failure scenarios
- **Documentation Debt**: Not updating relevant documentation

## Escalation Criteria
Escalate to full brownfield PRD process if:
- Scope grows beyond 3 stories
- Architectural changes become necessary
- Integration complexity exceeds expectations
- Risk assessment reveals high impact potential
- Multiple system areas become involved

## Integration Points
- **Archie (Architect)**: Architecture pattern guidance
- **Dave (Developer)**: Technical feasibility assessment
- **Sam (SM)**: Story breakdown and sprint planning
- **QA (Reviewer)**: Testing strategy and validation approach

## Templates and Tools
- Epic template: `.claude/resources/pm/templates/epic-template.md`
- Validation checklist: `.claude/resources/pm/checklists/pm-validation-checklist.md`
- Change management: `.claude/resources/pm/tasks/change-management.md`