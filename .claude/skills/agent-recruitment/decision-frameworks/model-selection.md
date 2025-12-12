# Model Selection Framework

## Available Models

| Model | Speed | Cost | Best For |
|-------|-------|------|----------|
| **Haiku** | Fast | Low | Simple tasks, validators, quick checks |
| **Sonnet** | Medium | Medium | Balanced analysis, most agents |
| **Opus** | Slow | High | Complex reasoning, architecture, planning |

## Model Selection by Agent Type

### Commands
Commands inherit the user's current model. No explicit selection needed.

### Skills
Skills also inherit. No model field in SKILL.md frontmatter.

### Sub-agents
Sub-agents SHOULD specify model for predictable behavior:

```yaml
---
name: my-subagent
model: haiku  # or sonnet, opus
---
```

## Decision Framework

```
What kind of task does the sub-agent perform?
│
├─► Simple, mechanical, rule-based?
│   └─► HAIKU
│       Examples: validators, formatters, simple extractors
│
├─► Analytical, requires judgment?
│   └─► SONNET
│       Examples: reviewers, analyzers, planners (simple)
│
└─► Complex reasoning, architecture, critical decisions?
    └─► OPUS
        Examples: architects, complex planners, strategists
```

## Detailed Guidelines

### Use HAIKU When

| Scenario | Rationale |
|----------|-----------|
| Validation checks | Rule-based, fast feedback |
| File operations | Mechanical, predictable |
| Simple transformations | No complex reasoning |
| Quick lookups | Speed > depth |
| High-volume operations | Cost efficiency |

**Examples**:
- `resource-validator` - Checks if files exist
- `constraint-validator` - Validates score ranges
- `coherence-verifier` - Simple consistency checks
- `file-formatter` - Applies formatting rules

### Use SONNET When

| Scenario | Rationale |
|----------|-----------|
| Code review | Needs judgment but not deep architecture |
| Pattern matching | Some inference required |
| Documentation | Quality matters, not critical |
| Standard planning | Follows known patterns |
| Most sub-agents | Good balance of speed/quality |

**Examples**:
- `backend-qa-reviewer` - Reviews code quality
- `pattern-extractor` - Identifies patterns
- `agent-planner` (simple agents) - Plans straightforward agents
- `documentation-generator` - Creates docs

### Use OPUS When

| Scenario | Rationale |
|----------|-----------|
| Architecture decisions | Long-term impact |
| Complex decomposition | Multiple tradeoffs |
| Novel problem solving | No existing patterns |
| Critical business logic | Errors are costly |
| Strategic planning | Requires deep reasoning |

**Examples**:
- `architecture-analyst` - Deep system analysis
- `agent-planner` (orchestrators) - Complex agent design
- `system-designer` - Creates system architectures
- `strategic-advisor` - Business decisions

## Context Size Considerations

| Model | Context Window | Implication |
|-------|----------------|-------------|
| Haiku | 200K | Sufficient for most tasks |
| Sonnet | 200K | Sufficient for most tasks |
| Opus | 200K | Same, but better at using large context |

Choose model based on **reasoning needs**, not context size.

## Cost-Performance Tradeoffs

### High Volume Scenarios
```
Processing 100 files for validation:
- Haiku: Fast, cheap, good enough
- Sonnet: Slower, 5x cost, marginal benefit
- Opus: Much slower, 15x cost, overkill
→ Choose HAIKU
```

### Critical Decision Scenarios
```
Designing system architecture:
- Haiku: Fast, might miss nuances
- Sonnet: Good analysis, may miss edge cases
- Opus: Thorough, catches subtleties
→ Choose OPUS
```

### Balanced Scenarios
```
Reviewing a pull request:
- Haiku: Too shallow for code review
- Sonnet: Good depth, reasonable speed
- Opus: Overkill for most PRs
→ Choose SONNET
```

## Model Escalation Pattern

Start with faster/cheaper, escalate if needed:

```
1. Try HAIKU first
   │
   ├─► Success? → Done
   │
   └─► Insufficient? → Escalate to SONNET
       │
       ├─► Success? → Done
       │
       └─► Still insufficient? → Escalate to OPUS
```

**When to use this pattern**:
- Uncertain about complexity
- Cost-sensitive environments
- Iterative refinement workflows

## Anti-Patterns

### Using Opus for Everything
```
BAD: All sub-agents use opus "to be safe"
GOOD: Match model to task complexity
```

### Using Haiku for Complex Tasks
```
BAD: Architecture planning with haiku
GOOD: Use opus for architectural decisions
```

### Not Specifying Model
```
BAD: Sub-agent without model field (unpredictable)
GOOD: Always specify model for sub-agents
```

## Quick Reference

```
┌────────────────────────────────────────────────────┐
│              MODEL SELECTION GUIDE                 │
├────────────────────────────────────────────────────┤
│                                                    │
│  HAIKU   →  Validators, formatters, lookups       │
│             Fast, cheap, mechanical tasks          │
│                                                    │
│  SONNET  →  Reviewers, analyzers, most agents     │
│             Balanced speed and quality             │
│                                                    │
│  OPUS    →  Architects, strategists, planners     │
│             Complex reasoning, critical decisions  │
│                                                    │
├────────────────────────────────────────────────────┤
│  DEFAULT: sonnet (when uncertain)                  │
│  ESCALATE: haiku → sonnet → opus (if needed)      │
└────────────────────────────────────────────────────┘
```
