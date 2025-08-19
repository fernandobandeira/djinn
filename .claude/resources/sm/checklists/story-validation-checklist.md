# Story Validation Checklist (BMAD-METHOD Enhanced)

## Critical Requirements (MUST PASS for GO Decision)

### 1. Template Completeness ✓
- [ ] All required sections present from story template
- [ ] No placeholder text remaining ({{variables}}, TBD, etc.)
- [ ] Epic reference clearly specified
- [ ] Story ID follows naming convention
- [ ] Status field properly set

### 2. Story Definition ✓
- [ ] Clear "As a... I want... So that..." format
- [ ] Role is specific and valid
- [ ] Action is implementable
- [ ] Benefit is measurable

### 3. Acceptance Criteria ✓
- [ ] All AC are specific and measurable
- [ ] Each AC is testable
- [ ] No ambiguous terms ("fast", "easy", "good")
- [ ] Edge cases considered
- [ ] Success definition clear for each AC

### 4. Tasks and Subtasks ✓
- [ ] Tasks cover all acceptance criteria
- [ ] Logical sequence and dependencies clear
- [ ] Each task is actionable
- [ ] Task granularity appropriate (not too large/small)
- [ ] AC references included where applicable

### 5. Dev Notes Completeness ✓
- [ ] All technical context included
- [ ] File paths and locations specified
- [ ] Architecture patterns referenced
- [ ] No external document lookups needed
- [ ] Source tree information included if relevant

## Quality Requirements (SHOULD PASS for High Score)

### 6. Anti-Hallucination ✓
- [ ] All technical claims verified against source docs
- [ ] No invented libraries or frameworks
- [ ] Architecture alignment confirmed
- [ ] References to actual project files only
- [ ] No assumptions about implementation

### 7. Testing Instructions ✓
- [ ] Test approach clearly defined
- [ ] Unit test requirements specified
- [ ] Integration test scenarios included
- [ ] Edge cases and error handling covered
- [ ] Test data requirements identified

### 8. Security Considerations ✓
- [ ] Authentication/authorization addressed
- [ ] Input validation requirements specified
- [ ] Data protection measures defined
- [ ] Common vulnerabilities considered
- [ ] Compliance requirements noted (if applicable)

### 9. UI/Frontend Specifications (if applicable) ✓
- [ ] Component specifications detailed
- [ ] User interaction flows defined
- [ ] Styling/design guidance provided
- [ ] Responsive design considered
- [ ] Accessibility requirements addressed

### 10. Implementation Readiness ✓
- [ ] Story is self-contained
- [ ] No missing prerequisites
- [ ] Technical complexity appropriate
- [ ] Estimated effort reasonable
- [ ] No blocking dependencies

## Scoring Guidelines

### GO Decision (Ready for Development)
- All Critical Requirements: PASS
- Quality Score: ≥ 70%
- No blocking issues identified
- Dev Notes section comprehensive

### CONDITIONAL GO (Minor improvements needed)
- All Critical Requirements: PASS
- Quality Score: 50-69%
- Minor issues that can be clarified during development
- Dev Notes mostly complete

### NO-GO Decision (Requires revision)
- Any Critical Requirement: FAIL
- Quality Score: < 50%
- Blocking issues identified
- Dev Notes insufficient for implementation

## Validation Report Format

```yaml
validation_result:
  decision: GO | CONDITIONAL_GO | NO-GO
  readiness_score: 0-100
  critical_requirements: 
    passed: X/5
    failed: []
  quality_requirements:
    passed: X/5
    warnings: []
  blocking_issues: []
  recommendations: []
  dev_readiness: LOW | MEDIUM | HIGH
```

## Common Failure Patterns

1. **Incomplete Dev Notes**: Missing critical technical context
2. **Vague Acceptance Criteria**: Not specific or measurable
3. **Task Sequence Issues**: Dependencies not clear
4. **Template Violations**: Missing required sections
5. **Hallucinated Content**: Unverified technical claims
6. **Missing Test Strategy**: No clear testing approach
7. **Security Gaps**: Common vulnerabilities not addressed
8. **UI Ambiguity**: Frontend requirements unclear
9. **External Dependencies**: Requires reading other docs
10. **Scope Creep**: Story too large or complex

## Improvement Priority

1. **CRITICAL**: Fix immediately (blocks development)
2. **HIGH**: Should fix before development
3. **MEDIUM**: Can be clarified during development
4. **LOW**: Nice to have improvements
5. **FUTURE**: Consider for next iteration