# Checklist Template

## Standard Checklist Structure

```markdown
# [Purpose] Checklist

## Overview
[Brief description of what this checklist validates]

## Context
**When to use**: [Situation or trigger]
**Who uses this**: [Target user/role]
**Time required**: [Estimate]

## Pre-Requirements
Must have before starting:
- [ ] [Prerequisite 1]
- [ ] [Prerequisite 2]
- [ ] [Prerequisite 3]

## Main Checklist

### Section 1: [Category Name]
- [ ] [Specific check item]
  - **Why**: [Reason this matters]
  - **How**: [Quick verification method]
- [ ] [Specific check item]
  - **Why**: [Reason this matters]
  - **How**: [Quick verification method]
- [ ] [Specific check item]
  - **Why**: [Reason this matters]
  - **How**: [Quick verification method]

### Section 2: [Category Name]
- [ ] [Specific check item]
- [ ] [Specific check item]
- [ ] [Specific check item]

### Section 3: [Category Name]
- [ ] [Specific check item]
- [ ] [Specific check item]
- [ ] [Specific check item]

## Critical Items
**MUST PASS** - Blocking issues:
- [ ] ⚠️ [Critical check 1]
- [ ] ⚠️ [Critical check 2]
- [ ] ⚠️ [Critical check 3]

## Nice to Have
Optional but recommended:
- [ ] 💡 [Enhancement 1]
- [ ] 💡 [Enhancement 2]

## Common Issues

### Issue: [Problem Description]
- **Check**: [What to look for]
- **Fix**: [How to resolve]

### Issue: [Problem Description]
- **Check**: [What to look for]
- **Fix**: [How to resolve]

## Completion Criteria
Checklist is complete when:
- ✅ All critical items pass
- ✅ [Percentage]% of main items checked
- ✅ [Specific outcome achieved]

## Follow-Up Actions
After completion:
1. [Next step]
2. [Documentation to update]
3. [Stakeholder to notify]

## References
- [Related checklist]
- [Documentation link]
- [Best practices guide]
```

## Checklist Types and Variations

### Quality Review Checklist
```markdown
# [Component] Quality Review Checklist

## Review Scope
**Component**: [What's being reviewed]
**Version**: [Version/commit]
**Reviewer**: [Name/role]
**Date**: [Review date]

## Code Quality
- [ ] Follows coding standards
- [ ] No obvious bugs
- [ ] Error handling present
- [ ] Comments where needed
- [ ] No dead code
- [ ] Consistent naming

## Functionality
- [ ] Meets requirements
- [ ] Edge cases handled
- [ ] Input validation
- [ ] Output correct format
- [ ] Performance acceptable

## Documentation
- [ ] README updated
- [ ] API documented
- [ ] Examples provided
- [ ] Changelog updated

## Testing
- [ ] Unit tests pass
- [ ] Integration tests pass
- [ ] Coverage adequate
- [ ] Manual testing done

## Security
- [ ] No hardcoded secrets
- [ ] Input sanitization
- [ ] Authentication proper
- [ ] Authorization checked

## Rating
Overall: [ ] Approved | [ ] Needs Work | [ ] Rejected
```

### Launch Readiness Checklist
```markdown
# [Feature] Launch Readiness Checklist

## Pre-Launch

### Development Complete
- [ ] All features implemented
- [ ] Code reviewed
- [ ] Tests passing
- [ ] Documentation complete

### Infrastructure Ready
- [ ] Environments configured
- [ ] Monitoring setup
- [ ] Backups configured
- [ ] Scaling tested

### Operations Prepared
- [ ] Runbooks created
- [ ] Team trained
- [ ] Support ready
- [ ] Rollback plan

## Launch Day
- [ ] Final testing
- [ ] Stakeholders notified
- [ ] Monitoring active
- [ ] Team available

## Post-Launch
- [ ] Metrics tracking
- [ ] User feedback collected
- [ ] Issues addressed
- [ ] Retrospective scheduled
```

### Validation Checklist
```markdown
# [Process] Validation Checklist

## Input Validation
- [ ] Required fields present
- [ ] Data types correct
- [ ] Ranges valid
- [ ] Formats matched

## Process Validation
- [ ] Steps executed in order
- [ ] Intermediate results correct
- [ ] No errors encountered
- [ ] Performance acceptable

## Output Validation
- [ ] All outputs generated
- [ ] Format correct
- [ ] Content accurate
- [ ] Location correct

## Success Confirmed
- [ ] Expected outcome achieved
- [ ] No side effects
- [ ] Stakeholders satisfied
```

### Migration Checklist
```markdown
# [System] Migration Checklist

## Pre-Migration
- [ ] Backup created
- [ ] Rollback plan documented
- [ ] Dependencies identified
- [ ] Downtime scheduled
- [ ] Team notified

## During Migration
- [ ] Services stopped
- [ ] Data transferred
- [ ] Schema updated
- [ ] Configuration applied
- [ ] Services restarted

## Post-Migration
- [ ] Functionality verified
- [ ] Data integrity checked
- [ ] Performance validated
- [ ] Monitoring confirmed
- [ ] Documentation updated

## Sign-off
- [ ] Technical lead approval
- [ ] Business owner approval
- [ ] Old system decommissioned
```

## Checklist Design Guidelines

### Characteristics of Good Checklists
1. **Specific**: Clear, unambiguous items
2. **Actionable**: Each item is verifiable
3. **Organized**: Logical grouping
4. **Prioritized**: Critical items marked
5. **Complete**: Covers all aspects
6. **Concise**: Not overwhelming

### Item Writing Tips
#### Good ✅
- "Verify API key is configured in .env"
- "Confirm all tests pass with 80%+ coverage"
- "Check error messages are user-friendly"

#### Bad ❌
- "Testing done" (too vague)
- "Everything works" (not measurable)
- "Code is good" (subjective)

### Grouping Strategies
- **By Phase**: Planning → Development → Testing → Deployment
- **By Component**: Frontend → Backend → Database → Infrastructure
- **By Priority**: Critical → Important → Nice-to-have
- **By Responsibility**: Developer → Reviewer → Operations

## Interactive Elements

### Progress Tracking
```markdown
## Progress: [23/30] - 77% Complete

### ✅ Completed (23)
### ⏳ In Progress (3)
### ❌ Blocked (2)
### ⏸️ Pending (2)
```

### Status Indicators
- 🟢 Complete
- 🟡 In Progress
- 🔴 Blocked
- ⚪ Not Started
- 🔵 Not Applicable

### Priority Markers
- 🚨 Critical (Blocker)
- ⚠️ High Priority
- 📌 Medium Priority
- 💡 Nice to Have

## Usage Notes

### When to Create a Checklist
- Repetitive validation needed
- Quality assurance required
- Complex process with many steps
- Compliance verification
- Launch/deployment validation

### Checklist vs Task
- **Checklist**: Validation and verification
- **Task**: Execution and creation

### Maintenance
- Review quarterly
- Update based on incidents
- Incorporate learnings
- Version if significant changes

## Remember
- Make items specific and measurable
- Group related checks together
- Mark critical items clearly
- Include verification methods
- Keep checklist focused
- Test with actual users
- Iterate based on feedback