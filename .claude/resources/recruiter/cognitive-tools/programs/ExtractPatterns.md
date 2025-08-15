# Cognitive Tool Program: Extract Patterns
# Automated pattern extraction from successful agents
# Based on cognitive tools framework and pattern recognition from Zettelkasten

## Program Purpose
This cognitive tool systematically extracts reusable patterns from successful agents, enabling Rita to build a growing library of proven design patterns.

## Input Schema
```yaml
extraction_request:
  target_agents: "List of agents to analyze"
  pattern_types: "Types to extract (tool/interaction/trigger/constraint)"
  success_threshold: "Minimum success rate to consider"
  generalization_level: "How abstract to make patterns"
```

## Pattern Recognition Process

### Step 1: Agent Success Analysis
```markdown
FOR each agent:
  EVALUATE success metrics:
    - Usage frequency
    - Rework rate
    - User satisfaction
    - Constraint balance
    - Delegation success
  
  IF meets threshold:
    ADD to extraction candidates
```

### Step 2: Structure Analysis
```markdown
EXTRACT structural elements:
  - Configuration patterns
  - Tool combinations
  - Trigger mechanisms
  - Interaction flows
  - Output formats
  - Constraint arrangements
```

### Step 3: Pattern Identification
```markdown
IDENTIFY recurring patterns:
  - Common tool sets
  - Repeated interaction styles
  - Successful trigger formulas
  - Effective constraint balances
```

### Step 4: Pattern Generalization
```markdown
GENERALIZE to reusable form:
  - Remove agent-specific details
  - Create template variables
  - Define applicability conditions
  - Document success factors
```

## Pattern Types

### Tool Patterns
```yaml
extraction_focus:
  - Minimal tool sets
  - Tool combinations
  - Tool-purpose alignment
  - Efficiency patterns

example_extraction:
  pattern: "read_analyze"
  tools: [Read]
  when: "Only analyzing existing content"
  success_rate: 0.92
  found_in: [doc-analyzer, code-reviewer, pattern-finder]
```

### Interaction Patterns
```yaml
extraction_focus:
  - Dialogue structures
  - User engagement methods
  - Information presentation
  - Decision flows

example_extraction:
  pattern: "progressive_revelation"
  structure: "Start simple → Add detail → Full complexity"
  when: "Teaching or explaining complex topics"
  success_rate: 0.88
  found_in: [tina, concept-explainer, tutorial-guide]
```

### Trigger Patterns
```yaml
extraction_focus:
  - Activation conditions
  - Keyword effectiveness
  - Context sensitivity
  - Delegation triggers

example_extraction:
  pattern: "proactive_on_change"
  trigger: "IMPORTANT ... after {event}"
  when: "Automated response needed"
  success_rate: 0.90
  found_in: [test-runner, linter, security-checker]
```

### Constraint Patterns
```yaml
extraction_focus:
  - Balance strategies
  - Scope definitions
  - Boundary enforcement
  - Flexibility points

example_extraction:
  pattern: "flexible_core"
  structure: "Rigid core + Flexible periphery"
  when: "Need consistency with adaptability"
  success_rate: 0.87
  found_in: [ana, archie, market-researcher]
```

## Extraction Algorithm

### Phase 1: Collection
```python
def collect_agents(criteria):
    agents = []
    for agent in all_agents:
        metrics = evaluate_agent(agent)
        if metrics.success_rate > criteria.threshold:
            agents.append({
                'agent': agent,
                'metrics': metrics,
                'structure': parse_structure(agent)
            })
    return agents
```

### Phase 2: Analysis
```python
def analyze_patterns(agents):
    patterns = {
        'tools': {},
        'interactions': {},
        'triggers': {},
        'constraints': {}
    }
    
    for agent in agents:
        for pattern_type in patterns:
            found = extract_type_patterns(agent, pattern_type)
            merge_patterns(patterns[pattern_type], found)
    
    return patterns
```

### Phase 3: Validation
```python
def validate_pattern(pattern):
    # Test with variations
    test_score = 0
    for test in generate_tests(pattern):
        if pattern_works(test):
            test_score += 1
    
    pattern.confidence = test_score / num_tests
    return pattern.confidence > 0.8
```

## Pattern Extraction Examples

### Example 1: Tool Pattern Extraction
```markdown
## Analyzing: Multiple Test Runners

### Agents Analyzed:
- test-runner (success: 92%)
- ci-validator (success: 89%)
- test-monitor (success: 90%)

### Common Structure Found:
- Tools: Always Read + Bash
- Never includes Write/Edit
- Purpose: Execute and analyze

### Extracted Pattern:
**Name**: test_executor
**Tools**: [Read, Bash]
**Purpose**: "Execute tests and analyze results"
**Constraint**: "Read-only file access"
**Success**: 90% average
**Reusability**: High
```

### Example 2: Interaction Pattern Extraction
```markdown
## Analyzing: Dialogue Agents

### Agents Analyzed:
- ana (success: 95%)
- tina (success: 93%)
- market-researcher (success: 88%)

### Common Structure Found:
- Present numbered options
- Wait for selection
- Progressive depth
- Maintain context

### Extracted Pattern:
**Name**: structured_dialogue
**Structure**:
  1. Present 3-5 options
  2. User selects
  3. Dive deeper on selection
  4. Maintain thread
**Success**: 92% average
**Reusability**: Very High
```

## Pattern Quality Metrics

### Extraction Quality
```yaml
metrics:
  pattern_confidence:
    calculation: "Success rate across instances"
    threshold: 0.85
    current_avg: 0.89
  
  pattern_coverage:
    calculation: "Agents using pattern / applicable agents"
    threshold: 0.60
    current_avg: 0.72
  
  pattern_stability:
    calculation: "Unchanged over time period"
    threshold: "3 months"
    current_avg: "4.5 months"
```

### Pattern Effectiveness
```yaml
metrics:
  reuse_rate:
    calculation: "Times pattern reused / opportunities"
    target: 0.70
    current: 0.75
  
  success_transfer:
    calculation: "Success in new agents / attempts"
    target: 0.85
    current: 0.88
  
  evolution_resistance:
    calculation: "Patterns needing modification"
    target: < 0.20
    current: 0.15
```

## Output Format

### Extracted Pattern Template
```yaml
pattern:
  id: "auto_generated_id"
  name: "descriptive_name"
  category: "tool|interaction|trigger|constraint"
  
  structure:
    core_elements: []
    optional_elements: []
    constraints: []
  
  metrics:
    found_in: []  # Source agents
    success_rate: 0.00
    confidence: 0.00
    instances: 0
  
  applicability:
    use_when: ""
    avoid_when: ""
    best_for: ""
  
  examples:
    - agent: "name"
      usage: "how it's used"
      outcome: "result"
```

## Integration with Rita

### Automatic Extraction Triggers
```yaml
triggers:
  - New agent achieves 90% success
  - Pattern used successfully 5+ times
  - Weekly extraction cycle
  - User requests extraction
```

### Storage and Indexing
```bash
# Store extracted patterns
save_to: /docs/agent-patterns/extracted/{category}/

# Index for searching
./.vector_db/kb index --path /docs/agent-patterns/extracted/

# Update pattern evolution cell
update_cell('pattern-evolution.yaml', new_patterns)
```

### Pattern Library Management
```python
def manage_pattern_library():
    # Promote successful patterns
    for pattern in experimental_patterns:
        if pattern.success_rate > 0.85:
            promote_to_validated(pattern)
    
    # Archive outdated patterns
    for pattern in validated_patterns:
        if pattern.last_used > "6 months ago":
            move_to_archive(pattern)
    
    # Track pattern evolution
    for pattern in all_patterns:
        track_mutations(pattern)
```

## Success Indicators

### Extraction Success
- Patterns extracted: 5+ per week
- Pattern quality: 85%+ confidence
- Reuse rate: 70%+ of patterns
- Time to extract: < 2 minutes per agent

### Library Growth
- New patterns/month: 15-20
- Active patterns: 50+
- Deprecated rate: < 10%/quarter
- Coverage: 80%+ of use cases

## Program Output Example

```markdown
## Pattern Extraction Report

### Agents Analyzed: 5
- test-runner (92% success)
- code-reviewer (89% success)
- doc-analyzer (91% success)
- security-scanner (88% success)
- format-checker (90% success)

### Patterns Extracted: 3

#### Pattern 1: minimal_analyzer
- **Category**: Tool Selection
- **Structure**: [Read] only
- **Found in**: 4/5 agents
- **Confidence**: 91%
- **Recommendation**: Add to validated library

#### Pattern 2: proactive_checker
- **Category**: Trigger
- **Structure**: "IMPORTANT ... after changes"
- **Found in**: 5/5 agents
- **Confidence**: 94%
- **Recommendation**: High-value pattern

#### Pattern 3: structured_report
- **Category**: Output
- **Structure**: Summary → Details → Actions
- **Found in**: 3/5 agents
- **Confidence**: 87%
- **Recommendation**: Consider for validation

### Actions Taken:
- Stored 3 patterns in library
- Indexed in vector database
- Updated pattern evolution tracking
```

This cognitive tool enables Rita to continuously learn from successful agents, building an ever-growing library of proven patterns that improve future agent creation.