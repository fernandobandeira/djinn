# Dev-QA Review Iteration Protocol

## 🚨 MANDATORY ENFORCEMENT RULES
1. **ALWAYS call qa-reviewer after implementation** - NO EXCEPTIONS
2. **ALWAYS call qa-reviewer after applying fixes** - EVERY TIME
3. **NEVER skip qa-reviewer validation** - It's MANDATORY
4. **Continue loop until qa-reviewer approves** - No shortcuts

## Overview
This protocol defines the MANDATORY internal quality assurance workflow for the dev agent, establishing a systematic approach to code review, critique, and iterative improvement.

**CRITICAL**: The review loop is NOT optional. Every implementation MUST go through qa-reviewer critique and approval.

## Phases of Review Iteration

### 1. Initial Implementation Phase
- Trigger: Completion of initial code implementation
- Action: Automatically delegate to qa-reviewer
- Artifact: `implementation_draft`

### 2. QA Critique Phase
- Trigger: Delegation to qa-reviewer
- Action: 
  - Comprehensive code analysis
  - Generate `critique_report`
  - Categorize issues:
    - Critical (must fix)
    - Important (should fix)
    - Minor (optional)

### 3. Fix Application Phase
- Trigger: Receiving `critique_report`
- Logic:
  ```yaml
  fix_strategy:
    critical_issues:
      - auto_apply: true
      - require_confirmation: false
    important_issues:
      - auto_apply: true
      - require_confirmation: true
    minor_issues:
      - auto_apply: false
      - require_human_review: true
  ```
- **MANDATORY NEXT STEP**: After applying ANY fixes → MUST call qa-reviewer again

### 4. Iteration Control
- Max Iterations: 5
- Iteration Tracking:
  ```yaml
  iteration_state:
    current_iteration: 0
    max_iterations: 5
    issues_resolved: []
    critical_fixes_applied: 0
    improvement_score: 0.0
  ```

### 5. Re-Validation Phase (MANDATORY)
- **Trigger**: ALWAYS after fixes applied (NO EXCEPTIONS)
- **MANDATORY Actions**: 
  - **MUST re-run qa-reviewer** (delegate again with "IMPORTANT actively critiques code")
  - Compare new `critique_report` with previous
  - Assess improvement
  - **REPEAT CYCLE**: If new issues found → Apply fixes → Call qa-reviewer AGAIN
- **CRITICAL**: This phase is NOT optional - ALWAYS re-validate after fixes

## Workflow Execution Example

```markdown
# Workflow Execution (MANDATORY LOOP)

## Iteration 1
- Initial Implementation ✓
- **MANDATORY**: Call qa-reviewer
- QA Critique Received
  - 3 Critical Issues
  - 2 Important Issues
- Fixes Applied: 
  - Critical Issues: Automatic
  - Important Issues: Confirmed
- **MANDATORY**: Call qa-reviewer again

## Iteration 2
- Re-submitted to QA (AUTOMATIC)
- New Critique:
  - 1 Critical Issue Remains
  - 1 New Important Issue Discovered
- Fixes Applied
- **MANDATORY**: Call qa-reviewer again

## Iteration 3
- Re-submitted to QA (AUTOMATIC)
- New Critique:
  - No critical issues
  - Minor improvements only
- Status: APPROVED ✓

**RULE**: After EVERY fix → MUST call qa-reviewer
```

## Delegation Triggers
- `@qa-trigger`: Initiate QA review
- `@fix-critical`: Apply critical fixes
- `@review-complete`: Stop iteration loop

## State Tracking Metrics
- `improvement_score`: Weighted reduction in code issues
- `critical_fix_ratio`: Percentage of critical issues resolved
- `iteration_count`: Number of review cycles

## Termination Conditions
1. qa-reviewer returns `approval_status: approved`
2. Max iterations reached (5) - escalate if not approved
3. All critical issues resolved AND no new critical issues found

**IMPORTANT**: Even when terminating, the LAST action must be a qa-reviewer validation

## Integration Hooks
- Integrates with existing dev workflow
- Connects to project's quality management system
- Logs all iterations for retrospective analysis

## Error Handling
- If max iterations reached without resolution:
  1. Escalate to senior developer
  2. Trigger manual code review
  3. Provide comprehensive issue report