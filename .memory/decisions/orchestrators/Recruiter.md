---
title: Recruiter
type: note
permalink: decisions/orchestrators/recruiter
---

# Recruiter (Rita)

## Core Principle

**Use an agent to help users build agents.** Encapsulate complexity behind a conversational interface—users describe what they need, Rita handles structure.

## Problem

Creating agents requires learning complex internals:
- Type selection (command vs skill vs sub-agent)
- Reusability assessment (skill tiers, shared vs embedded)
- Validation requirements (resource, constraint, coherence)
- Template structures and file locations

Users shouldn't need to understand these details to extend Djinn.

## Solution

Rita is a **meta-agent** - an agent that helps users build agents. She guides users through systematic creation without exposing internals.

**Implementation:** `/recruiter`

**What Rita Creates:**
- **Orchestrators** - Workflow personas like Ana, Archie
- **Skills** - Thinking techniques that auto-activate
- **Sub-agents** - Context isolation for heavy I/O

**What Rita Does:**
- Determines the right type for what user wants to build
- Assesses reusability to place components correctly
- Creates components using proven templates
- Validates before completion

**Skill Rita Uses:**
- [[Agent Recruitment]] - Contains the detailed frameworks (type selection, orchestrator design, reusability assessment, validation)

## Architecture

Recruiter follows a thin orchestrator + thick skill pattern:
- **Recruiter command** - Lightweight orchestrator (entry point)
- **[[Agent Recruitment]] skill** - Heavy lifting (frameworks, templates, validation)

This separation keeps the orchestrator focused on user interaction while the skill contains reusable agent creation logic.

## Workflow

Rita follows a systematic creation workflow:

1. **Search memory** - Check existing patterns/decisions
2. **Determine type** - Command, Skill, or Sub-agent?
3. **Assess reusability** - Skill (thinking) vs Sub-agent (context isolation)
4. **Plan architecture** - Design components
5. **Build files** - Create using templates
6. **Validate** - Check resources, constraints
7. **Offer to document** - Ask if user wants decision saved

## File Locations

| Agent Type | Location |
|------------|----------|
| Orchestrators | `.claude/commands/{name}.md` |
| Skills | `.claude/skills/{name}/SKILL.md` |
| Sub-agents | `.claude/agents/shared/{name}.md` |

## Why It Matters

- **Lowers barrier** - Users extend Djinn without learning internals
- **Ensures consistency** - All created agents follow patterns
- **Captures best practices** - Decision frameworks encoded in one place
- **Self-extending** - Djinn grows systematically without documentation burden

## Relations

- [[Architecture]] - Part of Djinn orchestrator layer
- [[Catalog]] - Listed in component index
- [[Agent Recruitment]] - The skill Rita loads for agent creation
- [[Orchestrator]] - Rita follows this pattern (guides users, uses skills, delegates to sub-agents for heavy I/O)