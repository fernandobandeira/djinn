# User Story: {STORY-ID} - {Title}

## Metadata
- **Story ID:** {STORY-ID}
- **Epic:** {Epic reference}
- **Author:** {author}
- **Created:** {YYYY-MM-DD}
- **Last Updated:** {YYYY-MM-DD}
- **Status:** Draft | Approved | In Progress | Review | Done
- **Priority:** Must Have | Should Have | Could Have | Won't Have
- **Estimated Effort:** {story points or hours}

## User Story

**As a** {user role}
**I want** {feature/capability}
**So that** {business value/outcome}

## Acceptance Criteria

### AC-1: {Criterion Title}
**Given** {precondition}
**When** {action}
**Then** {expected result}

- Measurable: {specific metric or validation}
- Test approach: {how to verify}

### AC-2: {Criterion Title}
**Given** {precondition}
**When** {action}
**Then** {expected result}

- Measurable: {specific metric or validation}
- Test approach: {how to verify}

## Tasks and Subtasks

### Task 1: {Task Description} (AC-1)
- **1.1**: {Subtask description}
  - Estimated: {time}
- **1.2**: {Subtask description}
  - Estimated: {time}

### Task 2: {Task Description} (AC-1, AC-2)
- **2.1**: {Subtask description}
  - Estimated: {time}
- **2.2**: {Subtask description}
  - Estimated: {time}

## Dev Notes

### Architecture Context
[Source: architecture/{file}.md#section]

{Relevant architecture decisions and patterns}

### Code Examples
```{language}
// Example implementation pattern
{code snippet}
```

### File Locations
- {file path}: {purpose}
- {file path}: {purpose}

### Integration Points
- {External system/API}: {description}
- {Internal component}: {description}

### Configuration
```yaml
# Required configuration
{config example}
```

## Testing Requirements

### Unit Tests
- [ ] {Test case 1}
- [ ] {Test case 2}

### Integration Tests
- [ ] {Test scenario 1}
- [ ] {Test scenario 2}

### Edge Cases
- [ ] {Edge case 1}
- [ ] {Edge case 2}

## Definition of Done

### Code Quality
- [ ] Code follows project style guide
- [ ] All linting checks pass
- [ ] Code reviewed and approved

### Testing
- [ ] Unit tests written and passing
- [ ] Integration tests passing
- [ ] Edge cases covered

### Documentation
- [ ] Code comments where needed
- [ ] API documentation updated
- [ ] README updated if applicable

### Deployment
- [ ] Configuration documented
- [ ] Deployment instructions clear
- [ ] Rollback procedure defined

## Dependencies

### Blocking Dependencies
- {Story/task that must complete first}

### Non-blocking Dependencies
- {Story/task that can proceed in parallel}

### External Dependencies
- {External service/API}

## Risks and Mitigation

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| {Risk description} | Low/Med/High | Low/Med/High | {Mitigation strategy} |

## Rollback Strategy

### Triggers
- {Condition that triggers rollback}

### Procedure
1. {Rollback step 1}
2. {Rollback step 2}

## Change Log

| Date | Author | Change |
|------|--------|--------|
| {YYYY-MM-DD} | {name} | Initial creation |
