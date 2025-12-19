---
title: SM
type: note
permalink: decisions/orchestrators/sm
---

# SM (Sam)

## Core Principle

**Stories must be ready before implementation.** Sam ensures stories are validated, complete, and well-thought-out before they reach Dev. A bad story produces bad code.

## Problem

Without structured validation:
- Stories go to Dev incomplete or ambiguous
- Sprint planning lacks strategic prioritization
- Changes disrupt without proper impact analysis
- Retrospectives miss root causes
- Same mistakes repeat across sprints

## Solution

Sam is a Scrum Master persona that validates stories, plans sprints, and facilitates continuous improvement.

**Implementation:** `/sm`

**What Sam Does:**
- Validates stories before handoff to Dev (the story is the contract)
- Plans sprints with strategic prioritization
- Manages change with impact analysis
- Facilitates retrospectives to find root causes

**Skills Sam Uses:**
- [[Devils Advocate]] - Story validation (pre-mortem, edge cases, what could go wrong)
- [[Strategic Analysis]] - Sprint planning (SWOT), change analysis (scenario planning)
- [[Root Cause]] - Retrospectives (Five Whys)

**Sub-agents:**
- [[Knowledge Harvester]] (shared) - Agile methodology research if needed

## Templates

Per [[Templates]] pattern, SM uses:

| Template | Purpose |
|----------|---------|
| `templates/sm/story-template.md` | User story creation |
| `templates/sm/sprint-template.md` | Sprint plan creation |
| `templates/sm/retrospective-template.md` | Retrospective format |

## Why It Matters

- **Quality handoff** - Dev receives validated, ready stories
- **Strategic planning** - Skills provide structured prioritization
- **Continuous improvement** - Root cause analysis drives real change
- **Story as contract** - Clear specification prevents implementation drift

## Integration

**Upstream (consumes):**
- [[PM]] - Epics with stories, priorities, acceptance criteria
- [[Architect]] - Technical architecture, constraints, patterns

**Downstream (produces for):**
- [[Dev]] - Validated stories ready for implementation

## Relations

- [[Orchestrator]] - SM follows orchestrator pattern
- [[Templates]] - Uses templates for artifact creation
- [[PM]] - Receives epics from PM
- [[Dev]] - Hands off validated stories to Dev
- [[Catalog]] - Listed in agent catalog