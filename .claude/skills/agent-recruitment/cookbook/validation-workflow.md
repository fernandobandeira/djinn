# Validation Workflow

## Overview

Every agent creation MUST pass validation before completion. Validation checks three aspects directly:

1. **Resource Validation** - Files exist and load correctly
2. **Constraint Validation** - Balance is optimal (8.0-8.5 score)
3. **Coherence Verification** - Components integrate properly

## Validation Flow

```
Agent Created
      │
      ▼
┌──────────────────┐
│ 1. RESOURCE      │
│    VALIDATION    │
│    Files exist?  │
└──────────────────┘
      │
      ├── FAIL → Fix missing files → Retry
      │
      ▼ PASS
┌──────────────────┐
│ 2. CONSTRAINT    │
│    VALIDATION    │
│    Balance ok?   │
└──────────────────┘
      │
      ├── FAIL → Adjust constraints → Retry
      │
      ▼ PASS
┌──────────────────┐
│ 3. COHERENCE     │
│    VERIFICATION  │
│    Aligned?      │
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

---

## 1. Resource Validation

Check all referenced files exist and are syntactically correct.

### Reference Patterns to Check

| Pattern | Example | Resolution |
|---------|---------|------------|
| `@` references | `@.claude/resources/file.md` | Direct path |
| Markdown links | `[text](./path/file.md)` | Relative to agent |
| Skill cookbook | `[cookbook/file.md](./cookbook/file.md)` | Relative to skill root |

### Resolution Rules

```
From Orchestrators (.claude/orchestrators/):
  @resources/file.md → .claude/resources/file.md

From Skills (.claude/skills/{name}/):
  ./cookbook/file.md → .claude/skills/{name}/cookbook/file.md

From Sub-agents (.claude/agents/{parent}/):
  @../shared/common.md → .claude/agents/shared/common.md
```

### Validation Checklist

**File Existence**:
- [ ] All `@path/file.md` references resolve
- [ ] All relative links `[text](./path)` exist
- [ ] All skill cookbook files exist
- [ ] All template references exist

**Syntax Validation**:
- [ ] YAML frontmatter parses correctly
- [ ] Markdown syntax is valid (no broken formatting)
- [ ] JSON/YAML data files parse without errors

**Path Validation**:
- [ ] Paths use correct relative format
- [ ] No absolute paths (should be relative to .claude/)
- [ ] Parent directories exist

### How to Validate Resources

1. Read the agent file
2. Extract all reference patterns using regex:
   - `@[^\s]+\.md` for @ references
   - `\[.*?\]\(\.?/[^)]+\)` for markdown links
3. Resolve each path based on agent location
4. Check file existence with Glob or Read
5. For YAML files, attempt to parse frontmatter

### Common Resource Issues

| Issue | Example | Fix |
|-------|---------|-----|
| Missing cookbook | `cookbook/advanced.md` not found | Create the file |
| Invalid YAML | Frontmatter parse error at line 5 | Fix YAML syntax |
| Broken link | `./templates/missing.md` → 404 | Update link or create file |
| Wrong path format | `/absolute/path.md` | Use relative `./path.md` |

### Pass Criteria

✅ All referenced files exist
✅ No syntax errors in any file
✅ All paths resolve correctly

---

## 2. Constraint Validation

Verify agent constraints are balanced for optimal performance.

### Constraint Score Formula

```
Score = Base(5.0)
      + ToolBalance(0-2)      # Appropriate tool count
      + ModelFit(0-1.5)       # Model matches complexity
      + ScopeClarity(0-1.5)   # Clear, focused purpose
      + Integration(0-1)      # Proper integration

Target: 8.0 - 8.5 (optimal)
Warning: < 7.5 or > 9.0
Fail: < 7.0 or > 9.5
```

### Scoring Criteria

#### Tool Balance (0-2 points)

| Tool Count | Agent Type | Score |
|------------|------------|-------|
| 1-2 | Validator | 2.0 |
| 2-3 | Reviewer | 2.0 |
| 3-4 | Builder | 1.5 |
| 5+ | Any | 0.5 |
| Too few for task | Any | 0.5 |

**Appropriate tool sets by role**:
- Validators: Read, Grep, Glob (3 max)
- Reviewers: Read, Grep, Glob (3 max)
- Builders: Read, Write, MultiEdit, Bash (4 max)
- Analyzers: Read, Grep, Glob, WebFetch (4 max)
- Orchestrators: Read, Task (2 typical)

#### Model Fit (0-1.5 points)

| Task Type | Expected Model | Score |
|-----------|----------------|-------|
| Simple/mechanical | haiku | 1.5 |
| Analysis/judgment | sonnet | 1.5 |
| Architecture/strategy | opus | 1.5 |
| Mismatch | any | 0.5 |

**Mismatch examples**:
- Opus for simple file validation → overkill
- Haiku for complex architecture decisions → underpowered

#### Scope Clarity (0-1.5 points)

| Description Quality | Score |
|--------------------|-------|
| Single, clear purpose | 1.5 |
| Focused but broad | 1.0 |
| Vague or multi-purpose | 0.5 |

**Good**: "Validates Go test coverage meets 80% threshold"
**Bad**: "Helps with code stuff"

#### Integration (0-1 point)

| Integration Quality | Score |
|--------------------|-------|
| Proper paths, all references valid | 1.0 |
| Minor issues (cosmetic) | 0.5 |
| Broken references | 0.0 |

### How to Calculate Score

1. Read agent frontmatter for tools and model
2. Count tools, check against agent type
3. Evaluate model against task complexity
4. Assess description clarity
5. Check integration points
6. Sum scores

### Example Calculations

**Well-balanced validator**:
```
Base:         5.0
Tool Balance: 2.0 (Read, Grep, Glob = 3 tools, appropriate)
Model Fit:    1.5 (haiku for validation, correct)
Scope:        1.3 (clear but could be more specific)
Integration:  0.5 (minor path issue)
─────────────────
Total:        8.3 ✅ PASS
```

**Over-constrained builder**:
```
Base:         5.0
Tool Balance: 0.5 (7 tools, excessive)
Model Fit:    0.5 (opus for simple task, overkill)
Scope:        1.0 (broad but focused)
Integration:  0.5 (some issues)
─────────────────
Total:        6.5 ❌ FAIL
```

### Common Constraint Issues

| Issue | Current | Suggested |
|-------|---------|-----------|
| Too many tools | 7 tools for reviewer | Reduce to 3: Read, Grep, Glob |
| Wrong model | Opus for validation | Use haiku |
| Vague scope | "Helps with code" | "Validates TypeScript types" |
| Over-engineering | Handles 10 cases | Focus on 2-3 core cases |

### Pass Criteria

✅ Score in 8.0-8.5 range
✅ No critical issues (tool/model mismatch)
✅ Clear, focused purpose

---

## 3. Coherence Verification

Verify all components of an agent integrate correctly.

### Coherence Scoring

```
Coherence Score = 1.0
  - 0.1 per naming issue
  - 0.15 per reference issue
  - 0.1 per documentation mismatch
  - 0.2 per missing component
  - 0.3 per circular dependency

Pass threshold: 0.85
Warning threshold: 0.70
```

### Coherence Checks

#### Reference Consistency
- [ ] Sub-agents reference correct parent paths
- [ ] Resources use consistent naming across files
- [ ] Cookbook files cross-reference correctly
- [ ] Templates match actual file structure

#### Naming Conventions
- [ ] Agent name in frontmatter matches filename
- [ ] Lowercase with hyphens (kebab-case)
- [ ] Consistent naming across all references
- [ ] Parent/child naming aligned

#### Documentation Alignment
- [ ] Description matches actual behavior
- [ ] Documented commands match implementation
- [ ] Output format documented correctly
- [ ] Examples are accurate and runnable

#### Integration Points
- [ ] All referenced sub-agents exist (for orchestrators)
- [ ] All skill cookbook files exist
- [ ] All resource paths resolve
- [ ] No circular dependencies

### How to Verify Coherence

1. Read the agent file
2. Extract all references (sub-agents, resources, cookbooks)
3. For each reference:
   - Check file exists
   - Check naming matches
   - Check cross-references are bidirectional
4. Compare documentation to implementation
5. Calculate coherence score

### Coherence Issue Types

| Type | Example | Deduction |
|------|---------|-----------|
| Naming | `name: agent-planner` but file is `planner.md` | -0.1 |
| Reference | Orchestrator references missing sub-agent | -0.15 |
| Documentation | `*patterns` command not implemented | -0.1 |
| Missing | Required cookbook file doesn't exist | -0.2 |
| Circular | A → B → A dependency | -0.3 |

### Common Coherence Issues

**Name Mismatch**:
```
Location: .claude/agents/recruiter/planner.md
Expected: name: planner
Actual:   name: agent-planner
Fix: Change name in frontmatter to match filename
```

**Missing Sub-agent**:
```
Location: orchestrator.md:45
Expected: .claude/agents/parent/helper.md exists
Actual:   File not found
Fix: Create helper.md or remove reference
```

**Documentation Drift**:
```
Location: SKILL.md:30
Expected: *coverage command documented
Actual:   No implementation for *coverage
Fix: Implement command or remove from docs
```

**Inconsistent Reference**:
```
Component A: recruiter.md references "agent-planner"
Component B: planner.md has name: "planner"
Fix: Align naming - use "planner" everywhere
```

### Integration Tree Verification

For orchestrators, verify complete chain:

```
Orchestrator
├── Sub-agent 1 exists
│   └── Has correct tools for its purpose
├── Sub-agent 2 exists
│   └── Has correct model for its complexity
└── All resources exist
    └── All paths resolve correctly
```

### Pass Criteria

✅ Coherence score > 0.85
✅ No critical inconsistencies
✅ All components exist and integrate

---

## Running Validation

### Manual Validation (`*validate`)

Perform all three validations in sequence:

1. **Resource Validation**
   - Extract all references from agent file
   - Check each file exists
   - Validate syntax where applicable
   - Report missing/invalid files

2. **Constraint Validation**
   - Extract tools, model, description
   - Calculate score using formula
   - Identify issues
   - Report score and recommendations

3. **Coherence Verification**
   - Check cross-references
   - Verify naming consistency
   - Compare docs to implementation
   - Report coherence score and issues

### Automatic Validation

After `*build` completes, validation runs automatically.

---

## Validation Report Format

```markdown
## Validation Report for {agent_name}

### Resource Validation
Status: ✅ PASS | ❌ FAIL
Files Checked: {count}
Missing: {list or "None"}
Invalid Syntax: {list or "None"}

### Constraint Validation
Status: ✅ PASS | ⚠️ WARNING | ❌ FAIL
Score: {score}/10 (target: 8.0-8.5)
Breakdown:
  - Tool Balance: {score}
  - Model Fit: {score}
  - Scope Clarity: {score}
  - Integration: {score}
Issues: {list or "None"}

### Coherence Verification
Status: ✅ PASS | ❌ FAIL
Coherence: {score}% (target: >85%)
Issues: {list or "None"}
Inconsistencies: {list or "None"}

### Overall
{ALL PASS: "✅ Agent validated successfully!"}
{ANY FAIL: "❌ Validation failed. Issues to fix:"}
{list of required fixes}
```

---

## Handling Failures

### Resource Failures

**Missing file**:
```
Issue: cookbook/advanced.md not found
Fix: Create the file with appropriate content
```

**Invalid syntax**:
```
Issue: config.yaml parse error at line 5
Fix: Correct YAML syntax
```

### Constraint Failures

**Score too low**:
```
Issue: Score 7.2 - tool set too broad
Fix: Remove unnecessary tools
Current: Read, Write, Edit, Bash, Grep, Glob, Task (7)
Suggested: Read, Write, Task (3)
```

**Score too high**:
```
Issue: Score 9.3 - may be over-constrained
Fix: Verify agent can accomplish its purpose with current constraints
```

### Coherence Failures

**Reference mismatch**:
```
Issue: Sub-agent name doesn't match file
Fix: Align name in frontmatter with filename
```

**Missing component**:
```
Issue: Referenced helper.md doesn't exist
Fix: Create file or update reference
```

---

## Re-validation

After fixing issues:

1. Run only the failed validation(s) first
2. If pass, run remaining validations
3. If all pass, validation complete
4. If fail again, address new issues

---

## Quality Gates Summary

| Gate | Criteria | Action on Fail |
|------|----------|----------------|
| Resources | All files exist, valid syntax | Create missing, fix syntax |
| Constraints | Score 8.0-8.5 | Adjust tools/model/scope |
| Coherence | Score > 85% | Fix naming/references |

---

## Final Checklist

Before marking agent complete:

- [ ] Resource validation passed (all files exist)
- [ ] Constraint validation passed (score 8.0-8.5)
- [ ] Coherence verification passed (>85%)
- [ ] All critical issues addressed
- [ ] No warnings left unexplained
- [ ] Validation report generated
