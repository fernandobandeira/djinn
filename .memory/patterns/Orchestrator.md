---
title: Orchestrator
type: note
permalink: patterns/orchestrator
---

# Orchestrator Pattern

## What is an Orchestrator?

An orchestrator is a **persona that guides users** to solve a particular problem.

**Characteristics:**
- **Has a persona** - A name and identity (Ana the Analyst, Archie the Architect, Rita the Recruiter)
- **Guides the user** - Helps users through a multi-phase workflow
- **Solves a specific problem** - Each orchestrator has a clear domain
- **Coordinates resources** - Uses skills for thinking, delegates to sub-agents for heavy I/O
- **Does reasoning directly** - Never delegates thinking work
- **Handles all storage** - Sub-agents return synthesis, orchestrators write to KB

**Implementation:** In Claude Code, orchestrators are `.claude/commands/{name}.md` files invoked via `/name`

## Core Principle

**Do Work Directly. Delegate only for context isolation.**

```
┌─────────────────────────────────────────┐
│            ORCHESTRATOR                 │
│  • Guides user through workflow         │
│  • Does reasoning work directly         │
│  • Uses skills for thinking frameworks  │
│  • Writes all content to KB             │
└─────────────────────────────────────────┘
         │               │
    (reasoning)    (context isolation)
         │               │
         ▼               ▼
    ┌─────────┐    ┌─────────────┐
    │ Skills  │    │ Sub-agents  │
    │ (HOW)   │    │ (heavy I/O) │
    └─────────┘    └─────────────┘
```

## Orchestrator Responsibilities

1. **Guide the user** - Multi-phase workflows with approval gates
2. **Search KB first** - Always check Basic Memory before creating
3. **Use skills for thinking** - Devils Advocate, Ideation, Root Cause, etc.
4. **Delegate heavy I/O** - Use sub-agents for research that would flood context
5. **Control all storage** - Sub-agents return synthesis, orchestrators decide what to save
6. **Apply configuration** - Use project parameter from CLAUDE.md for KB writes
7. **Link content** - Add [[wikilinks]] to connect related notes

## When to Create an Orchestrator

- Workflow requires multiple phases
- User needs guided multi-step process
- Different phases have distinct purposes
- Need to coordinate skills and sub-agents

## Orchestrator Flow

```
1. Receive user request
2. Search Basic Memory for existing content
3. Use skills for reasoning/analysis
4. Delegate to sub-agent if heavy I/O needed
5. Receive synthesis from sub-agent
6. Review and optionally transform
7. Write to Basic Memory with correct project
8. Return confirmation to user
```

## Why Orchestrators Handle Storage

Sub-agents execute in isolated contexts and cannot access CLAUDE.md configuration. If sub-agents write directly:
- Wrong project parameters
- Inconsistent note structures
- No orchestrator oversight on what to save

**Solution:** Sub-agents return synthesis. Orchestrators write with correct config.

See [[Sub-agent]] pattern for sub-agent details.

## Current Orchestrators

| Orchestrator | Persona | Domain |
|--------------|---------|--------|
| [[Analyst]] | Ana | Research, assumptions, briefs |
| [[Architect]] | Archie | Architecture, ADRs, patterns |
| [[Recruiter]] | Rita | Agent creation |

## Relations

- [[Architecture]] - Djinn architecture overview
- [[Skill]] - Skills pattern (orchestrators use skills for reasoning)
- [[Sub-agent]] - Sub-agent pattern (orchestrators delegate for heavy I/O)
- [[Catalog]] - All orchestrators listed