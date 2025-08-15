# Molecular Protocol: Constraint Negotiation
# Implements bidirectional constraint negotiation for resolving conflicts
# Based on organ-level orchestration insights from Zettelkasten

## Protocol Overview
This protocol handles constraint conflicts through systematic negotiation, ensuring agents achieve optimal constraint balance without manual intervention.

## Negotiation Framework

### Phase 1: Conflict Detection
```yaml
detection_triggers:
  - Multiple constraints apply to same aspect
  - Constraints have opposing requirements
  - Resource constraints conflict with functional needs
  - Scope boundaries overlap with tool capabilities
```

### Phase 2: Constraint Analysis
Analyze each conflicting constraint:

```markdown
For each constraint:
1. **Priority Level**: Safety > Core Function > Efficiency > UX > Nice-to-have
2. **Flexibility Score**: How much can this constraint bend? (0-10)
3. **Dependencies**: What other constraints depend on this?
4. **Impact Radius**: How many aspects does this affect?
```

### Phase 3: Negotiation Strategies

#### Strategy 1: Hierarchical Resolution
When constraints have different priorities:
```yaml
process:
  1. Identify priority levels
  2. Higher priority wins
  3. Lower priority adapts or is removed
  4. Document decision rationale

example:
  conflict: "Safety vs Performance"
  resolution: "Safety wins, accept performance impact"
```

#### Strategy 2: Mutual Adaptation
When constraints have equal priority:
```yaml
process:
  1. Find compromise point
  2. Both constraints adjust
  3. Verify core requirements still met
  4. Test compatibility

example:
  conflict: "Scope breadth vs Tool minimalism"
  resolution: "Moderate scope with essential tools only"
```

#### Strategy 3: Creative Synthesis
When direct resolution isn't possible:
```yaml
process:
  1. Reframe the problem
  2. Find alternative approach
  3. May involve delegation or composition
  4. Validate solution meets original needs

example:
  conflict: "Need write access but safety prevents it"
  resolution: "Delegate to specialized sub-agent with write permissions"
```

#### Strategy 4: Temporal Sequencing
When constraints conflict only when simultaneous:
```yaml
process:
  1. Identify temporal dependencies
  2. Sequence constraint application
  3. Apply in phases
  4. Ensure end state satisfies all

example:
  conflict: "Validation requires structure that transformation removes"
  resolution: "Validate first, then transform"
```

## Negotiation Protocol Steps

### Step 1: Gather Constraints
```python
# Pseudo-code for constraint gathering
constraints = {
    'atomic': load_atomic_constraints(),
    'molecular': load_molecular_protocols(),
    'cellular': load_memory_constraints(),
    'external': get_user_requirements()
}
```

### Step 2: Identify Conflicts
```python
conflicts = []
for c1 in constraints:
    for c2 in constraints:
        if conflicts_with(c1, c2):
            conflicts.append({
                'constraint_1': c1,
                'constraint_2': c2,
                'type': classify_conflict(c1, c2)
            })
```

### Step 3: Apply Negotiation
```python
for conflict in conflicts:
    strategy = select_strategy(conflict)
    resolution = apply_strategy(strategy, conflict)
    validate_resolution(resolution)
    document_outcome(resolution)
```

### Step 4: Validate Outcome
```python
final_constraints = apply_resolutions(constraints, resolutions)
score = constraint_analyzer.evaluate(final_constraints)
if score < 7.0:
    trigger_rebalancing()
```

## Conflict Resolution Examples

### Example 1: Tool-Scope Conflict
**Conflict**: Agent needs global search but should only have local tools
```yaml
analysis:
  - Global search: Priority 3 (Efficiency), Flexibility 8
  - Local tools: Priority 2 (Core Function), Flexibility 3

resolution:
  strategy: Hierarchical Resolution
  outcome: Keep local tools, find alternative search approach
  implementation: Use filtered grep instead of web search
```

### Example 2: Trigger Overlap
**Conflict**: Two agents trigger on same conditions
```yaml
analysis:
  - Agent A trigger: Priority 2, Flexibility 6
  - Agent B trigger: Priority 2, Flexibility 7

resolution:
  strategy: Mutual Adaptation
  outcome: Add distinguishing keywords to both
  implementation: 
    - Agent A: Requires "analyze" keyword
    - Agent B: Requires "review" keyword
```

### Example 3: Output Format Rigidity
**Conflict**: Strict format requirement but variable input
```yaml
analysis:
  - Strict format: Priority 3, Flexibility 2
  - Handle variations: Priority 2, Flexibility 9

resolution:
  strategy: Creative Synthesis
  outcome: Use adaptive formatter
  implementation: Create format template with optional sections
```

## Integration with Rita's Memory Cells

### Reading Conflict History
```bash
# Load past resolutions for similar conflicts
past_conflicts = load_cell('constraint-conflicts.yaml')
similar = find_similar_conflicts(current_conflict, past_conflicts)
if similar:
    apply_learned_resolution(similar.resolution)
```

### Updating Conflict Memory
```bash
# Store successful resolution
update_cell('constraint-conflicts.yaml', {
    'conflict': current_conflict,
    'resolution': applied_resolution,
    'outcome': measured_effectiveness
})
```

### Pattern Evolution
```bash
# If resolution creates new pattern
if is_novel_resolution(resolution):
    update_cell('pattern-evolution.yaml', {
        'new_pattern': extract_pattern(resolution),
        'parent': original_pattern,
        'improvement': calculate_improvement()
    })
```

## Negotiation Metrics

### Success Indicators
- Constraint score: 7.0-8.5 (well-balanced)
- No critical conflicts remaining
- All safety constraints preserved
- Core functionality maintained
- Improved efficiency metrics

### Failure Indicators
- Constraint score < 6.0 or > 9.0
- Unresolved conflicts
- Safety constraints compromised
- Core function degraded
- User intervention required

## Automated Negotiation Rules

### Always Automatically Resolve
1. **Tool Redundancy**: Remove unused tools
2. **Missing Triggers**: Add IMPORTANT keyword
3. **Vague Descriptions**: Add specific context
4. **Format Conflicts**: Default to flexible format

### Require Human Review
1. **Safety Compromises**: Never auto-resolve
2. **Core Function Changes**: Need approval
3. **Multi-Agent Coordination**: Complex interactions
4. **Novel Conflict Types**: Unknown patterns

## Learning Integration

### Pattern Extraction
After successful negotiation:
1. Extract resolution pattern
2. Generalize for reuse
3. Store in pattern repository
4. Index in vector database

### Continuous Improvement
Track negotiation effectiveness:
- Success rate by strategy
- Time to resolution
- Recurrence of conflicts
- Pattern reuse frequency

## Example Negotiation Session

```markdown
## Constraint Negotiation for: test-runner

### Detected Conflicts:
1. Tools: Has Write but only needs Read and Bash
2. Trigger: Could overlap with code-reviewer
3. Output: Format not specified

### Negotiation Process:

#### Conflict 1: Tool Redundancy
- Strategy: Hierarchical (Efficiency > Nice-to-have)
- Resolution: Remove Write tool
- Outcome: Reduced from 3 to 2 tools

#### Conflict 2: Trigger Overlap
- Strategy: Mutual Adaptation
- Resolution: Add "test" keyword requirement
- Outcome: Clear trigger separation

#### Conflict 3: Output Format
- Strategy: Creative Synthesis
- Resolution: Use structured but flexible format
- Outcome: Consistent yet adaptable output

### Final Constraint Score: 8.2 (Well-Balanced)
```

This protocol ensures Rita can automatically resolve most constraint conflicts while maintaining agent quality and learning from each resolution.