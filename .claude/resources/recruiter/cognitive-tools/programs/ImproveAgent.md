# Cognitive Tool Program: Improve Agent
# Recursive improvement program for iteratively refining agent definitions
# Based on recursive prompting and constraint optimization from Zettelkasten

## Program Purpose
This cognitive tool applies recursive improvement cycles to agent definitions, progressively optimizing constraint balance and effectiveness.

## Input Schema
```yaml
agent_definition:
  current_state: "Current agent configuration"
  constraint_score: "Current balance score (0-10)"
  identified_issues: "List of problems found"
  success_metrics: "Current performance data"
  improvement_goals: "What to optimize for"
```

## Improvement Process

### Step 1: Diagnostic Assessment
```markdown
ANALYZE current agent:
- Constraint balance: {under/over/well}
- Tool efficiency: {count and necessity}
- Trigger clarity: {ambiguous/clear}
- Output definition: {vague/specific}
- Scope boundaries: {unclear/defined}
```

### Step 2: Identify Improvement Vectors
```markdown
For each issue found:
1. Classify severity (critical/major/minor)
2. Determine improvement approach
3. Estimate impact on constraint score
4. Check for side effects
```

### Step 3: Apply Improvements
```markdown
ORDERED BY PRIORITY:
1. Fix critical issues first (safety, core function)
2. Address major issues (efficiency, clarity)
3. Polish minor issues (UX, nice-to-haves)
```

## Improvement Strategies

### For Under-Constrained Agents
```yaml
strategies:
  add_triggers:
    when: "No clear activation condition"
    action: "Add IMPORTANT or Proactively keyword"
    impact: "+1.5 constraint score"
  
  define_output:
    when: "Output format undefined"
    action: "Specify structure and format"
    impact: "+1.0 constraint score"
  
  clarify_scope:
    when: "Boundaries unclear"
    action: "Add explicit scope limits"
    impact: "+1.2 constraint score"
  
  minimize_tools:
    when: "Too many tools for purpose"
    action: "Remove unnecessary tools"
    impact: "+0.8 constraint score"
```

### For Over-Constrained Agents
```yaml
strategies:
  generalize_conditions:
    when: "Too specific requirements"
    action: "Abstract to patterns"
    impact: "-1.0 constraint score"
  
  add_flexibility:
    when: "Rigid format requirements"
    action: "Allow variations"
    impact: "-0.8 constraint score"
  
  broaden_applicability:
    when: "Narrow use case"
    action: "Expand scope slightly"
    impact: "-1.2 constraint score"
  
  simplify_logic:
    when: "Complex nested conditions"
    action: "Flatten and simplify"
    impact: "-0.7 constraint score"
```

### For Well-Balanced Agents
```yaml
strategies:
  optimize_efficiency:
    when: "Room for optimization"
    action: "Streamline without breaking balance"
    impact: "±0.2 constraint score"
  
  enhance_clarity:
    when: "Could be clearer"
    action: "Improve descriptions"
    impact: "±0.1 constraint score"
  
  add_examples:
    when: "Benefits from examples"
    action: "Add usage examples"
    impact: "No score change"
```

## Recursive Improvement Cycle

### Iteration 1: Major Fixes
```markdown
FOCUS: Critical issues
- Fix safety violations
- Restore core functionality
- Remove obvious problems
TARGET: Move score toward 5-6 range
```

### Iteration 2: Balance Tuning
```markdown
FOCUS: Constraint balance
- Adjust constraint levels
- Optimize tool selection
- Clarify boundaries
TARGET: Move score toward 7-8 range
```

### Iteration 3: Polish
```markdown
FOCUS: Refinement
- Enhance descriptions
- Add helpful examples
- Optimize for elegance
TARGET: Achieve 8-8.5 score
```

### Iteration 4: Validation
```markdown
FOCUS: Testing
- Verify improvements
- Check for regressions
- Measure effectiveness
TARGET: Confirm optimal state
```

## Improvement Examples

### Example 1: Under-Constrained to Balanced
```yaml
before:
  name: helper
  description: "Assists with tasks"
  tools: Read, Write, Edit, Bash
  score: 3.5

iteration_1:
  description: "Assists with code review tasks"
  tools: Read, Edit, Bash
  score: 5.0
  changes: ["Clarified purpose", "Reduced tools"]

iteration_2:
  description: "IMPORTANT reviews code for issues after changes"
  tools: Read, Bash
  score: 7.0
  changes: ["Added trigger", "Minimized tools"]

iteration_3:
  description: "IMPORTANT reviews code for security and quality issues after changes"
  tools: Read, Bash
  score: 8.2
  changes: ["Specified review focus"]

final:
  status: "Well-balanced"
  improvement: "+4.7 score"
  iterations: 3
```

### Example 2: Over-Constrained to Balanced
```yaml
before:
  name: json-validator
  description: "MUST validate only config.json with exactly 5 keys in order"
  tools: Read
  score: 9.5

iteration_1:
  description: "Validates JSON configuration files with specific structure"
  tools: Read
  score: 8.5
  changes: ["Generalized file requirement"]

iteration_2:
  description: "Validates JSON configuration files for required keys"
  tools: Read
  score: 8.0
  changes: ["Removed ordering requirement"]

final:
  status: "Well-balanced"
  improvement: "-1.5 score (positive)"
  iterations: 2
```

## Decision Matrix

| Current Score | Primary Strategy | Expected Iterations | Target Score |
|--------------|------------------|-------------------|--------------|
| 0-3 | Add major constraints | 3-4 | 7-8 |
| 3-5 | Add specific constraints | 2-3 | 7-8 |
| 5-7 | Fine-tune balance | 1-2 | 7.5-8.5 |
| 7-8.5 | Polish only | 0-1 | 8-8.5 |
| 8.5-10 | Relax constraints | 1-2 | 7.5-8.5 |

## Integration with Rita's Systems

### Load Current State
```python
# Get agent's current state
current = load_agent_definition(agent_name)
score = constraint_analyzer.evaluate(current)
issues = constraint_analyzer.diagnose(current)
```

### Apply Improvements
```python
# Iterative improvement
improved = current
for iteration in range(MAX_ITERATIONS):
    improved = apply_improvement_strategy(improved, issues)
    new_score = constraint_analyzer.evaluate(improved)
    
    if abs(new_score - TARGET_SCORE) < 0.5:
        break
    
    issues = constraint_analyzer.diagnose(improved)
```

### Track Progress
```python
# Record improvement history
history = {
    'agent': agent_name,
    'iterations': iteration_count,
    'initial_score': initial_score,
    'final_score': final_score,
    'strategies_applied': strategies,
    'time_taken': elapsed_time
}
store_in_pattern_evolution(history)
```

## Success Metrics

### Improvement Effectiveness
- Average score improvement: +2.5 points
- Iterations needed: 2-3 typical
- Success rate: 92% reach target
- Time per iteration: < 30 seconds

### Quality Indicators
- Agents stay balanced after improvement
- No regression in core functionality
- Improved delegation success rate
- Reduced rework requests

## Termination Conditions

### Stop Improving When:
1. Score reaches 8.0-8.5 range
2. No improvements found in iteration
3. Maximum iterations (5) reached
4. Score oscillating without progress
5. User intervention requested

## Output Template

```markdown
## Agent Improvement Report

### Initial State
- Score: {initial_score}
- Issues: {issue_list}

### Improvements Applied
**Iteration 1**: {changes_1}
- New Score: {score_1}

**Iteration 2**: {changes_2}
- New Score: {score_2}

### Final State
- Score: {final_score}
- Status: {under/over/well-balanced}
- Total Improvement: {score_change}

### Recommendations
{any additional suggestions}
```

This cognitive tool ensures Rita can automatically improve agent definitions through systematic, recursive refinement, achieving optimal constraint balance efficiently.