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
- **Asks before saving** - Always asks user before writing to memory (read is automatic, write is opt-in)
- **Handles all storage** - Sub-agents return synthesis, orchestrators write to KB

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

## Flow

1. **Guide the user** - Multi-phase workflows with approval gates
2. **Search memory first** - Always check existing content before creating
3. **Use skills for thinking** - Devils Advocate, Ideation, Root Cause, etc.
4. **Delegate heavy I/O** - Use sub-agents for research that would flood context
5. **Control all storage** - Sub-agents return synthesis, orchestrators decide what to save
6. **Ask before saving** - Always ask user before writing to memory (prevents polluting project docs)
7. **Apply configuration** - Use correct project/scope for memory writes
8. **Link content** - Add wikilinks to connect related notes

## When to Create an Orchestrator

- Workflow requires multiple phases
- User needs guided multi-step process
- Different phases have distinct purposes
- Need to coordinate skills and sub-agents

## Why Orchestrators Handle Storage

Memory is the most important artifact - docs are the source of truth. Orchestrators are best positioned to write meaningful content because they:

- Have full conversation context and user intent
- Can apply reasoning and skills to refine content
- Understand how new content relates to existing notes
- Can ensure quality standards before committing to memory

Sub-agents return synthesis. Orchestrators decide what's worth keeping and write it properly.

## Relations

- [[Architecture]] - Djinn architecture overview
- [[Skill]] - Skills pattern (orchestrators use skills for reasoning)
- [[Sub-agent]] - Sub-agent pattern (orchestrators delegate for heavy I/O)
- [[Catalog]] - All orchestrators listed