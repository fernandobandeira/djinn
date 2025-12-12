# Thinking Level Triggers for Agent Creation

## Overview

Claude Code supports thinking budget keywords that allocate more reasoning tokens for complex problems. This framework defines when to use each level during agent creation.

## Thinking Levels

| Level | Keywords | Tokens | Cost Impact |
|-------|----------|--------|-------------|
| **Standard** | (none) | ~2,000 | Baseline |
| **Think** | "think" | 4,000 | Low |
| **Megathink** | "think hard", "think deeply", "think a lot" | 10,000 | Medium |
| **Ultrathink** | "ultrathink", "think harder", "think very hard" | 31,999 | High |

## When to Use Each Level

### Standard Analysis (No Keyword)

**Use for**:
- Simple validators with clear rules
- Single-purpose sub-agents
- Clear, well-defined requirements
- Agents under 100 lines expected
- Copying/adapting existing patterns

**Examples**:
- "Create a sub-agent that validates JSON schema"
- "Make a command that runs tests"
- "Add a simple greeting responder"

### Think (4K Tokens)

**Use for**:
- Moderate complexity agents
- 2-3 tool combinations
- Some ambiguity in requirements
- Adapting patterns with modifications

**Trigger phrases**:
```
"think about how to structure this"
"think through the requirements"
```

**Examples**:
- "Create a code formatter agent, think about edge cases"
- "Design a file watcher, think about the events to handle"

### Megathink (10K Tokens)

**Use for**:
- Complex agents with 3-4 tools
- Skills with cookbook patterns
- Multi-component agents
- Integration with existing systems
- Unclear tool selection

**Trigger phrases**:
```
"think hard about this architecture"
"think deeply about the design"
"think more about the decomposition"
```

**Examples**:
- "Create a PR reviewer skill, think hard about what to check"
- "Design a deployment agent, think deeply about the stages"

### Ultrathink (32K Tokens)

**Use for**:
- Orchestrators coordinating sub-agents
- System decomposition decisions
- Migration from monolithic to modular
- Architecture-level decisions
- Ambiguous or conflicting requirements
- Novel agent patterns not seen before
- Critical business logic agents

**Trigger phrases**:
```
"ultrathink about the decomposition"
"think harder about the architecture"
"think very hard about how to split this"
```

**Examples**:
- "Design a scrum master orchestrator, ultrathink about sub-agent coordination"
- "Create an architect agent that designs systems, ultrathink about the workflow"
- "Migrate this 500-line command to skills, think harder about the split"

## Complexity Assessment Heuristics

### Tool Count
| Tools | Suggested Level |
|-------|-----------------|
| 1-2 | Standard |
| 3-4 | Megathink |
| 5+ | Ultrathink |

### Sub-agent Coordination
| Pattern | Suggested Level |
|---------|-----------------|
| No sub-agents | Standard/Think |
| 1-2 sub-agents | Megathink |
| 3+ sub-agents | Ultrathink |
| Orchestrator pattern | Ultrathink |

### Resource Requirements
| Resources | Suggested Level |
|-----------|-----------------|
| Single file | Standard |
| 2-3 files | Think |
| 4+ files | Megathink |
| Progressive disclosure | Megathink/Ultrathink |

### Uncertainty Level
| Clarity | Suggested Level |
|---------|-----------------|
| Crystal clear | Standard |
| Mostly clear | Think |
| Some ambiguity | Megathink |
| Highly ambiguous | Ultrathink |

## Encoding in Agent Instructions

### In SKILL.md
```markdown
## Thinking Protocol

Before planning any agent:
1. Assess complexity using the heuristics
2. Select appropriate thinking level
3. Apply trigger phrase if needed

For simple agents: proceed with standard analysis
For complex agents: "think hard about the architecture"
For orchestrators: "ultrathink about decomposition strategy"
```

### In Sub-agent Prompts
```markdown
When delegating to agent-planner:

"Analyze requirements for {agent_name}.
Complexity indicators: {tool_count} tools, {sub_agent_count} sub-agents.
{IF complex: "Think hard about" | IF orchestrator: "Ultrathink about"}
the optimal architecture and decomposition strategy."
```

## Cost-Benefit Guidelines

### When Ultrathink is Worth It
- Architecture decisions affecting multiple agents
- Decomposition that will be hard to change later
- Novel patterns without existing examples
- Critical production agents
- User explicitly requests deep analysis

### When to Avoid Ultrathink
- Simple CRUD-style agents
- Well-established patterns being reused
- Time-sensitive requests
- Iterative refinement (use standard, iterate)

## Anti-Patterns

### Over-Thinking
```
BAD: "ultrathink about this simple validator"
GOOD: Standard analysis for validators
```

### Under-Thinking
```
BAD: Standard analysis for orchestrator with 5 sub-agents
GOOD: "ultrathink about the coordination strategy"
```

### Thinking Without Action
```
BAD: Spending 32K tokens then not using insights
GOOD: Apply thinking results to concrete decisions
```

## Integration with Workflow

### Planning Phase
```
1. Gather requirements
2. Assess complexity → Select thinking level
3. Apply thinking to architecture decisions
4. Document decisions in plan
```

### Building Phase
```
1. Follow plan from thinking phase
2. Standard analysis usually sufficient
3. Think hard only if implementation surprises arise
```

### Validation Phase
```
1. Standard analysis for running validators
2. Think hard if validation reveals design issues
3. Ultrathink if major refactoring needed
```

## Quick Reference Card

```
┌─────────────────────────────────────────────────────┐
│           THINKING LEVEL QUICK GUIDE                │
├─────────────────────────────────────────────────────┤
│                                                     │
│  Standard    →  Simple, clear, <100 lines          │
│  Think       →  Moderate, 2-3 tools, some ambiguity│
│  Megathink   →  Complex, 3-4 tools, multi-component│
│  Ultrathink  →  Orchestrators, decomposition, novel│
│                                                     │
├─────────────────────────────────────────────────────┤
│  ALWAYS ULTRATHINK FOR:                             │
│  • Orchestrators with sub-agents                    │
│  • System decomposition                             │
│  • Architecture migrations                          │
│  • Ambiguous requirements                           │
└─────────────────────────────────────────────────────┘
```
