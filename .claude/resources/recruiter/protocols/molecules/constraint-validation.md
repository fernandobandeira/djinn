# Molecular Protocol: Constraint Validation Chain
# Validates agent definitions against the complete constraint architecture

## Protocol Overview
This molecular protocol chains together validation checks to ensure agents are well-constrained.

## Constraint Diagnostic Framework
Based on the three-state model from constraint architecture theory:
- **Under-constrained**: Too vague, inconsistent results
- **Over-constrained**: Too rigid, fails on novel inputs  
- **Well-balanced**: Consistent core with adaptability

## Validation Chain

### Phase 1: Syntactic Validation
```yaml
checks:
  - frontmatter_valid: Check YAML parsability
  - required_fields: Verify name and description
  - tool_format: Ensure comma-separated tools
  - file_location: Verify correct directory
```

### Phase 2: Semantic Validation
```yaml
checks:
  - purpose_clarity: Is the purpose specific?
  - trigger_effectiveness: Will it delegate properly?
  - tool_appropriateness: Do tools match purpose?
  - output_definition: Is expected output clear?
```

### Phase 3: Constraint Balance Analysis

#### Under-Constrained Indicators
```yaml
symptoms:
  - description: "Generic or vague purpose"
    example: "Helps with stuff"
    fix: "Add specific actions and contexts"
    
  - description: "Missing trigger keywords"
    example: "Reviews code"
    fix: "Add IMPORTANT or Proactively"
    
  - description: "No output format specified"
    example: "Analyzes things"
    fix: "Define expected output structure"
    
  - description: "Unclear scope"
    example: "Does various tasks"
    fix: "Define specific boundaries"
```

#### Over-Constrained Indicators
```yaml
symptoms:
  - description: "Too many specific conditions"
    example: "ONLY works with Python 3.8.2 on Ubuntu 20.04"
    fix: "Generalize to broader applicability"
    
  - description: "Inflexible output requirements"
    example: "Must produce exactly 47 lines of JSON"
    fix: "Allow format flexibility"
    
  - description: "Excessive tool restrictions"
    example: "Can only read files named test.py"
    fix: "Broaden tool usage patterns"
    
  - description: "Overly complex rules"
    example: "17 nested conditions for simple task"
    fix: "Simplify constraint logic"
```

#### Well-Balanced Indicators
```yaml
characteristics:
  - Clear purpose with flexibility
  - Appropriate trigger conditions
  - Minimal but sufficient tools
  - Defined but adaptable output
  - Specific domain with room for variation
```

## Validation Workflow

### Step 1: Load Agent Definition
```python
# Pseudo-code for validation
agent = load_agent_file(path)
constraints = load_atomic_constraints()
```

### Step 2: Apply Constraint Stack
```python
results = {
  'syntax': validate_syntax(agent, constraints.syntax),
  'tools': validate_tools(agent, constraints.tools),
  'triggers': validate_triggers(agent, constraints.triggers),
  'location': validate_location(agent, constraints.location)
}
```

### Step 3: Diagnose Constraint Balance
```python
balance = diagnose_balance(agent)
if balance == 'under_constrained':
    suggest_additions()
elif balance == 'over_constrained':
    suggest_relaxations()
else:
    mark_as_well_balanced()
```

### Step 4: Generate Validation Report
```markdown
## Validation Report for {agent_name}

### Syntax Validation
- ✅ Valid YAML frontmatter
- ✅ Required fields present
- ✅ Tools properly formatted

### Semantic Validation  
- ✅ Clear purpose defined
- ⚠️ Trigger keywords could be stronger
- ✅ Tools appropriate for task

### Constraint Balance
**Diagnosis**: Slightly under-constrained

**Recommendations**:
1. Add IMPORTANT keyword to description
2. Define expected output format
3. Specify error handling approach

### Overall Score: 7/10
Agent is functional but could benefit from stronger constraints.
```

## Example Validations

### Example 1: Under-Constrained Agent
```yaml
---
name: helper
description: Assists with tasks
tools: Read, Write, Edit, Bash, WebSearch
---
```
**Issues**:
- Vague purpose
- No trigger keywords
- Too many tools for undefined scope

### Example 2: Well-Balanced Agent
```yaml
---
name: test-runner
description: IMPORTANT runs tests and fixes failures after code changes
tools: Read, Bash
model: sonnet
---
```
**Strengths**:
- Clear trigger and purpose
- Minimal appropriate tools
- Specific but flexible scope

### Example 3: Over-Constrained Agent
```yaml
---
name: json-validator
description: MUST validate only config.json files with exactly 5 keys
tools: Read
---
```
**Issues**:
- Too specific file requirement
- Inflexible structure expectation

## Constraint Tuning Guidelines

### To Fix Under-Constrained:
1. Add specific trigger keywords
2. Define clear success criteria
3. Specify output format
4. Narrow scope to specific domain
5. Add validation rules

### To Fix Over-Constrained:
1. Generalize file/format requirements
2. Allow output flexibility
3. Broaden applicable contexts
4. Reduce conditional complexity
5. Enable graceful degradation

## Integration with Rita's Knowledge Base
After validation, store patterns:
```bash
# Store successful patterns
./.vector_db/kb index --path /docs/agent-patterns/successful-agents/

# Store failure patterns
./.vector_db/kb index --path /docs/agent-patterns/failure-patterns/
```

This allows Rita to learn from both successes and failures over time.