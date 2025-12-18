---
title: Skill
type: note
permalink: patterns/skill
---

# Skill Pattern

## What is a Skill?

A skill is a **thinking technique** that auto-activates based on conversation context. Skills teach HOW to think about problems.

**Characteristics:**
- **Auto-activates** - Triggers on keywords/context, not explicit invocation
- **Teaches thinking** - Provides structured methods, not just information
- **Has techniques** - Contains specific, actionable methods (not just concepts)
- **Reusable** - Multiple agents benefit from the same skill

## Skill Structure

Every skill note should follow this structure:

```markdown
# {Skill Name}

**Tier:** Foundational | Universal | Domain

## Core Principle
**Bold memorable statement.** Brief explanation of the key insight.

## Problem
What goes wrong without this skill.

## Solution
**Techniques:**
- **Technique 1** - Brief description
- **Technique 2** - Brief description

**Auto-activates when:** trigger words/phrases

## Supported By (if applicable)
Links to foundational skills this skill builds on.

## Why It Matters
Bullet points on value.

## Used By
Which orchestrators use this skill.

## Relations
Links to Architecture, Catalog, Skill pattern.
```

## Skill Tiers

| Tier | What it is | Examples |
|------|------------|----------|
| **Foundational** | Building blocks other skills compose | Role Playing, Devils Advocate |
| **Universal** | Most agents use, may compose foundational | Ideation, Root Cause, Teacher |
| **Domain** | Specific to a cluster of 2-3 agents | Strategic Analysis, User Research, Agent Recruitment |

### Foundational Skills

Building blocks that other skills compose. These enable core thinking modes:

- [[Role Playing]] - Enables perspective-shifting (stepping into other viewpoints)
- [[Devils Advocate]] - Enables challenging/critiquing (finding weaknesses)

### Universal Skills

Used by most agents. May compose foundational skills:

- [[Ideation]] - Creative idea generation (composes Role Playing + Devils Advocate)
- [[Root Cause]] - Deep problem analysis (composes Role Playing + Devils Advocate)
- [[Teacher]] - Knowledge transfer (independent)

### Domain Skills

Specific to a cluster of related agents:

- [[Strategic Analysis]] - Strategic frameworks (composes Role Playing + Devils Advocate)
- [[User Research]] - User understanding (composes Role Playing)
- [[Agent Recruitment]] - Agent creation (meta-skill for building Djinn)

## Skill Interconnections

```
              FOUNDATIONAL
    ┌─────────────────────────────────┐
    │  Role Playing   Devils Advocate │
    │  (perspectives)   (challenge)   │
    └───────────┬─────────────────────┘
                │ compose into
    ┌───────────┴───────────┐
    │                       │
    ▼       UNIVERSAL       ▼
┌──────────┐ ┌──────────┐ ┌──────────┐
│ Ideation │ │Root Cause│ │ Teacher  │
│          │ │          │ │(independ)│
└──────────┘ └──────────┘ └──────────┘

               DOMAIN
┌──────────┐ ┌──────────┐ ┌──────────┐
│Strategic │ │  User    │ │  Agent   │
│ Analysis │ │ Research │ │Recruitmt │
└──────────┘ └──────────┘ └──────────┘
```

## Valid Skill Test

Before creating a skill, verify:

1. **Has structured techniques?** - Not just concepts, but methods with steps
2. **Teaches thinking?** - Not execution or information retrieval
3. **Multiple users?** - 2+ agents would benefit
4. **Recognized methodology?** - Based on established practices

**If NO to any → Don't create a skill**

Example that failed: Systems Thinking had concepts (feedback loops, leverage points) but no structured techniques with clear steps.

## Relations

- [[Architecture]] - Skills are part of the architecture layers
- [[Orchestrator]] - Orchestrators use skills for reasoning
- [[Catalog]] - All skills listed in catalog