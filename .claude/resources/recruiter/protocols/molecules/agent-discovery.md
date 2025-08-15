# Molecular Protocol: Agent Discovery & Duplication Prevention
# Combines search constraints to prevent duplicate agent creation

## Protocol Overview
This molecular protocol ensures no duplicate agents are created by combining discovery constraints with Rita's knowledge base search.

## Atomic Constraints Applied
1. **Naming uniqueness** - No duplicate names
2. **Functional uniqueness** - No duplicate purposes
3. **Knowledge reuse** - Learn from existing patterns

## Discovery Workflow

### Step 1: Search Rita's Knowledge Base
```bash
# Search for similar patterns
./.vector_db/kb search "[agent-name] [purpose]"
./.vector_db/kb search "[main-functionality]"
```

### Step 2: File System Discovery
```bash
# Overview
tree .claude/ -L 2

# Name check (case-insensitive)
ls .claude/commands/ | grep -i "[proposed-name]"
ls .claude/agents/ | grep -i "[proposed-name]"

# Function check
grep -r "[main-keyword]" .claude/commands/
grep -r "[main-keyword]" .claude/agents/
```

### Step 3: Pattern Analysis
If similar agents found:
```bash
# Examine structure
head -30 [similar-agent-file]

# Extract patterns
grep "tools:" [similar-agent-file]
grep "description:" [similar-agent-file]
grep "^\*" [similar-agent-file]  # Commands
```

### Step 4: Decision Matrix

| Finding | Action | Rationale |
|---------|--------|-----------|
| **Exact name match** | Abort | Name must be unique |
| **Similar purpose** | Review & Differentiate | May enhance or specialize |
| **Related domain** | Learn patterns | Reuse successful patterns |
| **No matches** | Proceed | Green light for creation |

## Constraint Validation
- [ ] Name is unique across both commands and agents
- [ ] Purpose is sufficiently differentiated
- [ ] Patterns extracted from similar agents
- [ ] Knowledge base consulted for best practices

## Pattern Extraction
When similar agents found, extract:
1. **Delegation patterns** - How they're triggered
2. **Tool combinations** - What tools they use
3. **Constraint balance** - Are they well-constrained?
4. **Success indicators** - Do they work well?

## Integration Points
- Search `/docs/agent-patterns/` for learned patterns
- Delegate to `recruiter-pattern-learner` for extraction
- Update Rita's KB with findings

## Output Schema
```yaml
discovery_result:
  name_conflicts: []
  similar_agents: 
    - name: string
      similarity: percentage
      patterns_to_learn: []
  recommendation: proceed|abort|differentiate
  patterns_extracted: []
```

This protocol ensures Rita never creates duplicate agents while learning from existing successes.