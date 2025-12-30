---
title: Catalog
type: note
permalink: reference/catalog
---

# Component Catalog

Reference for all Djinn components.

## Orchestrators

| Orchestrator | Persona | Purpose |
|--------------|---------|---------|
| [[Analyst]] | Ana | Research, analysis, brainstorming |
| [[Architect]] | Archie | System design, ADRs |
| [[PM]] | Paul | Product requirements, epics with stories |
| [[Recruiter]] | Rita | Create new agents |
| [[SM]] | Sam | Sprint planning, story breakdown into tasks |
| [[UX]] | Ulysses | User research, design, frontend specs |
| [[Dev]] | Dave | Task implementation, TDD, quality gates |

## Skills

### Foundational

Building blocks that other skills compose:

| Skill | Enables |
|-------|---------|
| [[Role Playing]] | Perspective-shifting |
| [[Devils Advocate]] | Challenging/critiquing |

### Universal

Used by most agents:

| Skill | Solves |
|-------|--------|
| [[Ideation]] | Unfocused brainstorming |
| [[Root Cause]] | Surface-level analysis |
| [[Teacher]] | Poor explanations |

### Domain

Cluster-specific (2-3 agents):

| Skill | Solves |
|-------|--------|
| [[Strategic Analysis]] | Unstructured decisions |
| [[User Research]] | Shallow user understanding |
| [[Agent Recruitment]] | Inconsistent creation |

## Sub-agents

### Foundational

Used by all orchestrators:

| Sub-agent | Isolates |
|-----------|----------|
| [[Knowledge Harvester]] | External source gathering |

### Domain

Analyst-specific:

| Sub-agent | Isolates |
|-----------|----------|
| [[Market Researcher]] | Market research |
| [[Competitive Analyzer]] | Competitor comparison |

---

---

## Claude Code Implementation

Djinn on Claude Code (Anthropic's CLI):

| Doc | Purpose |
|-----|---------|
| [[Claude Code Guide]] | Installation and usage |
| [[Claude Code]] | Implementation reference (syntax, conventions) |
| [[Beads]] | Working Memory implementation (beads CLI) |

---

## Understand the Design

Why each component type exists and how they relate:

**Patterns:**
- [[Orchestrator]] - Workflow personas that guide users
- [[Skill]] - Thinking techniques that auto-activate
- [[Sub-agent]] - Context isolation for heavy I/O
- [[Templates]] - Platform-agnostic artifact structures (separate, reusable)
- [[Checklists]] - Workflow verification (embedded in orchestrators)
- [[Memory]] - Docs-first philosophy (Knowledge Memory)
- [[Working Memory]] - Persistent work tracking (epics, stories, tasks)

## Build Something New

Use `/recruiter` - Rita guides you through creating orchestrators, skills, and sub-agents.

## Relations

- [[Architecture]] - Design principles
- [[Project]] - Vision and philosophy
