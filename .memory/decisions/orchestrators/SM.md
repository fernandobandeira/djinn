---
title: SM
type: note
permalink: decisions/orchestrators/sm
---

# SM (Sam)

## Core Principle

**Do Work Directly.** Sam coordinates sprint workflows using skills for reasoning. No dedicated sub-agents - all thinking work happens directly with skill support.

## Problem

Scrum Master workflows involve multiple phases: story creation, validation, sprint planning, change management, retrospectives. Each requires structured thinking and quality gates.

**Without a dedicated orchestrator:**
- Story quality varies without consistent validation
- Sprint planning lacks strategic prioritization
- Changes disrupt without proper impact analysis
- Retrospectives miss root causes

## Solution

**Techniques:**
- **Story validation** with `devils-advocate` (pre-mortem, red team)
- **Sprint planning** with `strategic-analysis` (SWOT for prioritization)
- **Change analysis** with `strategic-analysis` (scenario planning)
- **Retrospectives** with `root-cause` (Five Whys)

**Characteristics:**
- Strict "Do Work Directly" - no reasoning sub-agents
- Uses existing Tier 1/2 skills for all analysis
- Templates for artifact creation (stories, sprints, retros)
- Validation criteria embedded in orchestrator

## Templates

Per [[Templates]] pattern, SM uses:

| Template | Purpose |
|----------|---------|
| `templates/sm/story-template.md` | User story creation |
| `templates/sm/sprint-template.md` | Sprint plan creation |
| `templates/sm/retrospective-template.md` | Retrospective format |

## Commands

- `*create-story` - Create next story from epic (uses story-template)
- `*validate-story {id}` - Validate with embedded criteria + devils-advocate
- `*plan-sprint` - Plan sprint with strategic-analysis
- `*manage-change` - Analyze changes with strategic-analysis
- `*retrospective` - Facilitate retro with root-cause skill
- `*status` - Sprint metrics and progress
- `*help` / `*exit` - Standard orchestrator commands

## Architectural Decision

**Why no SM-specific sub-agents?**

The original djinn-app SM had 6 sub-agents (story-creator, story-validator, change-analyzer, sprint-planner, retrospective-facilitator, execution-tracker). We eliminated all of them because:

1. **They did reasoning work** - Per framework principles, reasoning should be done directly
2. **They needed skill access** - Sub-agents can't call skills
3. **They needed user interaction** - Sub-agents can't ask follow-up questions
4. **Context wasn't a problem** - None required heavy I/O that would flood context

**What we use instead:**
- Skills for structured thinking (root-cause, devils-advocate, strategic-analysis)
- `knowledge-harvester` (shared) for agile methodology research if needed
- Embedded validation criteria in the orchestrator

## Integration

**Upstream (consumes):**
- [[PM]] - Epics with stories, priorities, acceptance criteria
- [[Architect]] - Technical architecture, constraints, patterns

**Downstream (produces for):**
- Dev agents - Stories ready for implementation

## Why It Matters

- **Consistent quality** - Embedded validation ensures all stories meet criteria
- **Strategic planning** - Skills provide structured prioritization
- **Continuous improvement** - Root cause analysis drives real change
- **Framework alignment** - Follows "Do Work Directly" principle

## Relations

- [[Orchestrator]] - SM follows orchestrator pattern
- [[Templates]] - Uses templates for artifact creation
- [[PM]] - Receives epics from PM
- [[Architect]] - Consumes architecture context
- [[Catalog]] - Listed in agent catalog
