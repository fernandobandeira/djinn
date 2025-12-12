# Model Selection Framework

## Available Models

| Model | Speed | Cost | Best For |
|-------|-------|------|----------|
| **Haiku** | Fast | Low | Simple I/O tasks, context isolation |
| **Sonnet** | Medium | Medium | Research, analysis sub-agents |
| **Opus** | Slow | High | DON'T use for sub-agents (do reasoning directly) |

## Model Selection by Agent Type

### Commands
Commands inherit the user's current model. No explicit selection needed.
Commands do reasoning work directly.

### Skills
Skills also inherit. No model field in SKILL.md frontmatter.
Skills do reasoning work directly.

### Sub-agents
Sub-agents SHOULD specify model. But remember: **sub-agents are for context isolation only**, not for reasoning.

```yaml
---
name: my-subagent
model: haiku  # or sonnet
---
```

## Key Insight

**If you need opus-level reasoning, don't use a sub-agent.**

Sub-agents can't call skills and have limited reasoning. If the task requires deep reasoning, do it directly in the command/skill.

## Decision Framework

```
What kind of task is this sub-agent performing?
│
├─► Heavy I/O, formatting, simple output?
│   └─► HAIKU
│       Examples: doc generators, diagram generators
│
├─► Research, data gathering, analysis?
│   └─► SONNET
│       Examples: market researcher, competitive analyzer
│
└─► Complex reasoning, architecture, decisions?
    └─► DON'T USE SUB-AGENT
        Do this work directly in command/skill
```

## Detailed Guidelines

### Use HAIKU When

| Scenario | Rationale |
|----------|-----------|
| Document generation | Formatting, no reasoning |
| Diagram generation | Template-based output |
| File operations | Mechanical, predictable |
| Simple transformations | No complex reasoning |
| High-volume I/O | Cost efficiency |

**Examples**:
- `documentation-generator` - Formats docs from input
- `diagram-generator` - Creates Mermaid/PlantUML
- Simple data formatters and transformers

### Use SONNET When

| Scenario | Rationale |
|----------|-----------|
| Web research | Needs to assess sources |
| Competitive analysis | Some judgment required |
| Data summarization | Must identify key points |
| Heavy reading + summary | Context isolation needed |

**Examples**:
- `market-researcher` - Researches and summarizes
- `competitive-analyzer` - Gathers and compares
- `insight-synthesizer` - Summarizes findings

### DON'T Use Sub-agents When

| Scenario | Do Instead |
|----------|------------|
| Validation | Do directly in skill |
| Architecture decisions | Do directly with skills |
| Planning | Do directly with skills |
| Anything needing follow-up questions | Do directly |
| Complex reasoning | Do directly |

## Context Size Considerations

| Model | Context Window | Implication |
|-------|----------------|-------------|
| Haiku | 200K | Sufficient for I/O tasks |
| Sonnet | 200K | Sufficient for research |
| Opus | 200K | Use directly, not in sub-agent |

## Cost-Performance Tradeoffs

### Heavy I/O (Context Isolation)
```
Generating 10 documentation files:
- Haiku: Fast, cheap, perfect for formatting
- Sonnet: Overkill for doc generation
→ Choose HAIKU
```

### Research (Context Isolation)
```
Researching 20 web pages and summarizing:
- Haiku: May miss nuances
- Sonnet: Good at identifying key points
→ Choose SONNET
```

### Reasoning (Don't Use Sub-agent)
```
Validating agent constraints:
- Any sub-agent: Can't reason deeply
→ Do directly in skill/command
```

## Anti-Patterns

### Using Opus for Sub-agents
```
BAD: model: opus for any sub-agent
GOOD: If you need opus reasoning, do it directly
```

### Using Sub-agents for Reasoning
```
BAD: Sub-agent for validation, planning, architecture
GOOD: Do reasoning work directly (access to skills)
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
│  HAIKU   →  Document generation, diagrams         │
│             Simple I/O, formatting tasks          │
│                                                    │
│  SONNET  →  Research, competitive analysis        │
│             Data gathering that needs judgment    │
│                                                    │
│  OPUS    →  DON'T USE FOR SUB-AGENTS             │
│             Do reasoning work directly instead    │
│                                                    │
├────────────────────────────────────────────────────┤
│  KEY PRINCIPLE:                                    │
│  Sub-agents are for CONTEXT ISOLATION only        │
│  Reasoning work → do directly in command/skill    │
└────────────────────────────────────────────────────┘
```
