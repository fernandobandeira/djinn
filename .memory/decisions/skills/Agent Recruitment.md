---
title: Agent Recruitment
type: note
permalink: decisions/skills/agent-recruitment
---

# Agent Recruitment

**Tier:** Domain

## Core Principle

**Skills teach HOW to think. Sub-agents isolate context. Orchestrators coordinate both.**

This guides all type selection:
- Need thinking technique → **Skill**
- Need context isolation (heavy I/O) → **Sub-agent**
- Need explicit invocation with workflow → **Command/Orchestrator**

Orchestrators do reasoning work directly (using skills) and only delegate to sub-agents for context isolation.

## Problem

Creating Djinn components is inconsistent and error-prone:
- Users don't know which type to create (command/orchestrator vs skill vs sub-agent)
- Orchestrators are complex - coordinating skills and sub-agents requires careful design
- Reusability decisions are made ad-hoc (skill tiers, shared vs embedded)
- Validation is skipped or incomplete
- Created components don't follow established patterns

## Solution

Encapsulated frameworks for systematic creation of all Djinn component types.

**Creates:**
- **Commands/Orchestrators** - Workflow personas that coordinate skills and sub-agents
- **Skills** - Thinking techniques that auto-activate on context
- **Sub-agents** - Context isolation for heavy I/O

**Decision Frameworks:**
- **Type Selection** - Command vs Skill vs Sub-agent decision tree
- **Orchestrator Design** - How to decompose workflows, when to delegate
- **Reusability Assessment** - Tier 1 vs Tier 2 vs embedded
- **Model Selection** - haiku vs sonnet for sub-agents
- **Validation Workflow** - Resource, constraint, coherence checks

**Auto-activates when:** User mentions "create agent", "build agent", "recruit", "new command", "new orchestrator", "design skill", "make a sub-agent", or agent modification requests

## Why It Matters

- **Meta-skill** - Skill for building skills/agents
- **Encapsulates complexity** - Users don't learn internals
- **Ensures consistency** - All agents follow patterns
- **Self-extending** - Djinn grows systematically

## Used By

- [[Recruiter]] - Rita loads this skill for all agent creation work

## Relations

- [[Architecture]] - Part of Tier 2 skill layer
- [[Catalog]] - Listed in component index
- [[Recruiter]] - The command that loads this skill
- [[Orchestrator]] - Key pattern for orchestrator design (sub-agents return synthesis, orchestrators write to KB)
- [[Skill]] - Follows skill pattern