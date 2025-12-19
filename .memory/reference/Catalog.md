---
title: Catalog
type: note
permalink: reference/catalog
---

# Component Catalog

Quick reference to all Djinn components. Each links to its decision note documenting problem, solution, and rationale.

## Orchestrators

| Orchestrator | Persona | What it solves |
|--------------|---------|----------------|
| [[Analyst]] | Ana | Over-optimistic assumptions, unvalidated briefs |
| [[Architect]] | Archie | Misjudged architecture (over/under-engineering) |
| [[Recruiter]] | Rita | Complex agent creation internals |

*Implementation: Invoked via `/name`*

## Skills

### Foundational

Building blocks that other skills compose:

| Skill | What it enables |
|-------|-----------------|
| [[Role Playing]] | Perspective-shifting (stepping into other viewpoints) |
| [[Devils Advocate]] | Challenging/critiquing (finding weaknesses) |

### Universal

Used by most agents:

| Skill | What it solves |
|-------|----------------|
| [[Ideation]] | Unfocused brainstorming |
| [[Root Cause]] | Surface-level problem solving |
| [[Teacher]] | Poor concept explanation |

### Domain

Cluster-specific (2-3 agents):

| Skill | What it solves |
|-------|----------------|
| [[Strategic Analysis]] | Unstructured strategic decisions |
| [[User Research]] | Shallow user understanding |
| [[Agent Recruitment]] | Inconsistent component creation |

## Sub-agents (Context Isolation)

### Foundational

Used by all orchestrators:

| Sub-agent | What it isolates |
|-----------|------------------|
| [[Knowledge Harvester]] | External source gathering and synthesis |

### Domain

Analyst-specific:

| Sub-agent | What it isolates |
|-----------|------------------|
| [[Market Researcher]] | Heavy web research for market analysis |
| [[Competitive Analyzer]] | Multi-competitor research and comparison |

## Key Patterns

- [[Orchestrator]] - Personas that guide users, handle all KB writes
- [[Skill]] - Thinking techniques; 3 tiers (Foundational → Universal → Domain)
- [[Sub-agent]] - Context isolation for heavy I/O; returns synthesis

## Relations

- [[Architecture]] - Design principles
- [[Claude Code Guide]] - How to use these components
- [[Project]] - Why Djinn exists