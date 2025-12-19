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

## Core Principle

**Teach HOW to think, not WHAT to do.** Skills guide reasoning through structured techniques. They stay in the main conversation (never delegated) because thinking requires full context.

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

| Tier | What it is |
|------|------------|
| **Foundational** | Building blocks other skills compose |
| **Universal** | Most agents use, may compose foundational |
| **Domain** | Specific to a cluster of 2-3 agents |

See [[Catalog]] for current skills.

## When to Create a Skill

- Has structured techniques (not just concepts, but methods with steps)
- Teaches thinking (not execution or information retrieval)
- 2+ agents would benefit
- Based on recognized methodology

**If NO to any → Don't create a skill**

## Relations

- [[Architecture]] - Skills are part of the architecture layers
- [[Orchestrator]] - Orchestrators use skills for reasoning
- [[Catalog]] - All skills listed in catalog