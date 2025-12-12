# Validation Workflow

## Overview

Every agent creation MUST pass validation before completion. Validation uses three specialized sub-agents that check different aspects of quality.

## The Three Validators

| Validator | Purpose | Model | Key Checks |
|-----------|---------|-------|------------|
| `resource-validator` | Files exist and load | haiku | File existence, syntax, references |
| `constraint-validator` | Balance is optimal | haiku | 8.0-8.5 score, tool count, complexity |
| `coherence-verifier` | Components integrate | haiku | Cross-references, consistency |

## Validation Flow

```
Agent Created
      │
      ▼
┌──────────────────┐
│ resource-validator│
│   Files exist?    │
└──────────────────┘
      │
      ├── FAIL → Fix missing files → Retry
      │
      ▼ PASS
┌──────────────────┐
│constraint-validator│
│  Balance optimal? │
└──────────────────┘
      │
      ├── FAIL → Adjust constraints → Retry
      │
      ▼ PASS
┌──────────────────┐
│ coherence-verifier│
│ Components align? │
└──────────────────┘
      │
      ├── FAIL → Fix inconsistencies → Retry
      │
      ▼ PASS
┌──────────────────┐
│  VALIDATION      │
│  COMPLETE        │
└──────────────────┘
```

## Running Validation

### Manual Validation

When user requests `*validate`:

```markdown
Run all three validators in sequence:

1. Delegate to resource-validator
2. Delegate to constraint-validator
3. Delegate to coherence-verifier
4. Compile results
5. Report to user
```

### Automatic Validation

After `*build` completes, validation runs automatically.

## Validator Details

### 1. Resource Validator

**Purpose**: Verify all referenced files exist and are syntactically correct.

**Checks**:
- [ ] All `@path/file.md` references resolve
- [ ] Linked files `[text](./path)` exist
- [ ] YAML frontmatter parses correctly
- [ ] Markdown syntax is valid
- [ ] Script files are executable (if applicable)

**Delegation**:
```
Task(
  subagent_type="resource-validator",
  description="Validate resource files",
  prompt="""
  Validation request from Rita:

  Agent: {agent_name}
  Path: {agent_path}

  Check all resource references and file existence.

  Return:
  - status: pass | fail
  - files_checked: [list]
  - missing: [list with details]
  - invalid_syntax: [list with details]
  """
)
```

**Pass Criteria**: All files exist, no syntax errors.

### 2. Constraint Validator

**Purpose**: Verify agent constraints are balanced for optimal performance.

**Checks**:
- [ ] Constraint score in 8.0-8.5 range (optimal)
- [ ] Tool count appropriate for agent type
- [ ] Model selection justified
- [ ] Complexity matches purpose
- [ ] No over-engineering

**Constraint Score Formula**:
```
Score = Base(5.0)
      + ToolBalance(0-2)
      + ModelFit(0-1.5)
      + ScopeClarity(0-1.5)
      + Integration(0-1)

Optimal Range: 8.0 - 8.5
Warning: < 7.5 or > 9.0
Fail: < 7.0 or > 9.5
```

**Delegation**:
```
Task(
  subagent_type="constraint-validator",
  description="Validate constraint balance",
  prompt="""
  Validation request from Rita:

  Agent: {agent_name}
  Path: {agent_path}
  Type: {command | skill | subagent}

  Calculate constraint balance score.

  Return:
  - score: float
  - breakdown: {tool_balance, model_fit, scope, integration}
  - issues: [list]
  - recommendations: [list]
  """
)
```

**Pass Criteria**: Score 8.0-8.5, no critical issues.

### 3. Coherence Verifier

**Purpose**: Verify all components integrate correctly.

**Checks**:
- [ ] Sub-agents reference correct paths
- [ ] Resource files used consistently
- [ ] No circular dependencies
- [ ] Naming conventions followed
- [ ] Documentation matches implementation

**Delegation**:
```
Task(
  subagent_type="coherence-verifier",
  description="Verify component coherence",
  prompt="""
  Validation request from Rita:

  Agent: {agent_name}
  Path: {agent_path}
  Components: {list of related files}

  Check cross-component coherence.

  Return:
  - coherence_score: float (0-1)
  - issues: [list with severity]
  - inconsistencies: [list]
  - recommendations: [list]
  """
)
```

**Pass Criteria**: Coherence score > 0.85, no critical inconsistencies.

## Compiling Results

After all validators run, compile a report:

```markdown
## Validation Report for {agent_name}

### Resource Validation
Status: ✅ PASS | ❌ FAIL
Files Checked: {count}
Issues: {count}

### Constraint Validation
Status: ✅ PASS | ❌ FAIL
Score: {score}/10
Issues: {count}

### Coherence Verification
Status: ✅ PASS | ❌ FAIL
Coherence: {score}%
Issues: {count}

### Overall
{ALL PASS: "Agent validated successfully!"}
{ANY FAIL: "Validation failed. Issues to fix:"}
```

## Handling Failures

### Resource Failures

```markdown
Missing file: .claude/resources/x/template.md

Fix: Create the missing file with agent-builder
```

### Constraint Failures

```markdown
Score too low (7.2): Tool set too broad

Fix: Remove unnecessary tools:
- Current: Read, Write, Edit, Bash, Grep, Glob, Task
- Suggested: Read, Write, Task
```

### Coherence Failures

```markdown
Inconsistency: Sub-agent references non-existent resource

Fix: Update reference or create missing resource
```

## Re-validation

After fixing issues:

1. Run only the failed validator(s)
2. If pass, run remaining validators
3. If all pass, validation complete
4. If fail again, report new issues

## Validation Bypass

**NEVER** skip validation. If validation seems wrong:

1. Review the validator's reasoning
2. If validator is incorrect, fix the validator
3. If agent is incorrect, fix the agent
4. Re-run validation

## Quality Gates

| Gate | Criteria | Action on Fail |
|------|----------|----------------|
| Resources | All files exist | Create missing |
| Constraints | Score 8.0-8.5 | Adjust tools/scope |
| Coherence | Score > 85% | Fix inconsistencies |

## Checklist

Before marking agent complete:

- [ ] Resource validator passed
- [ ] Constraint validator passed (8.0-8.5)
- [ ] Coherence verifier passed (>85%)
- [ ] All issues addressed
- [ ] No warnings ignored
- [ ] Report generated
