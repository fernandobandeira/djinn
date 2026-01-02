---
title: Orchestrator Skill Invocation Pattern
type: note
permalink: decisions/agents/orchestrator-skill-invocation-pattern
---

# Orchestrator Skill Invocation Pattern

## Context

Orchestrators need to invoke skills (thinking techniques) during their workflows. Skills auto-activate based on context triggers in their YAML `description` field, but this isn't reliable when the conversation is dominated by an orchestrator's persona.

## Decision

**Orchestrators must explicitly invoke skills using the Skill tool.**

When an orchestrator command needs skill-based thinking, the workflow must include an explicit instruction to use the Skill tool rather than just referencing the skill name.

## Pattern

### Before (Broken)
```markdown
### *brainstorm {topic}
1. Use ideation skill for brainstorming
```

This just documents that a skill exists - Claude Code doesn't know to invoke it.

### After (Working)
```markdown
### *brainstorm {topic}
1. **Invoke skill** - Use Skill tool with `skill: "ideation"`
2. **Skill facilitates** - Ideation skill runs the session
3. **Save output** - Offer to save results to memory
```

The explicit "Use Skill tool with..." instruction tells Claude Code to actually call the skill.

## Skill Tool Syntax

```
Use Skill tool with `skill: "skill-name"`
Use Skill tool with `skill: "skill-name", args: "technique-name"`
```

### Examples

| Need | Skill Tool Call |
|------|-----------------|
| Brainstorming | `skill: "ideation"` |
| Five Whys analysis | `skill: "root-cause", args: "five-whys"` |
| SWOT analysis | `skill: "strategic-analysis", args: "swot"` |
| Pre-mortem | `skill: "devils-advocate", args: "pre-mortem"` |
| Journey mapping | `skill: "user-research", args: "journey-map"` |
| Six Hats thinking | `skill: "role-playing", args: "six-hats"` |

## Workflow Structure

Each orchestrator command that uses skills should follow this pattern:

```markdown
### *command-name {param}
1. **Context** - Gather required context
2. **Invoke skill** - Use Skill tool with `skill: "...", args: "..."`
3. **Skill facilitates** - Skill runs the technique
4. **Save output** - Offer to save results to appropriate folder
```

## Applied To

This pattern was applied to all orchestrators:

- [[Analyst]] - brainstorm, five-whys, swot, pre-mortem, journey-map, six-hats, first-principles
- [[PM]] - create-brief, create-prd, create-roadmap, change-assessment
- [[UX]] - research, personas, journeys, interviews, surveys, design, validate, accessibility
- [[Architect]] - design-system, review-architecture
- [[SM]] - validate-story, plan-sprint, manage-change, retrospective
- [[Dev]] - review, debug

## Key Insight

**Orchestrators are dispatchers.** When a command matches a skill's domain, invoke that skill explicitly rather than hoping auto-activation triggers.

## Related

- [[Skill]] - How skills are structured
- [[Orchestrator]] - How orchestrators work
- [[Architecture]] - Overall framework design
