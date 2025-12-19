---
title: PM
type: note
permalink: decisions/orchestrators/pm
---

# PM (Paul)

## Core Principle

**Synthesize, Don't Duplicate.** PM aggregates findings from [[Analyst]], [[Architect]], and [[UX]] teams into unified product artifacts. Uses existing research rather than generating new analysis.

## Problem

Product teams need unified artifacts that:
- Synthesize insights from multiple specialized teams
- Translate research into actionable requirements
- Create clear handoffs for execution (to Scrum Master)
- Track stakeholder alignment and changes

## Solution

An orchestrator that:
- Consumes upstream research (market, competitive, user, technical)
- Uses skills for structured thinking (not sub-agents for reasoning)
- Produces PRDs, roadmaps, and epics
- Hands off to [[SM]] (Sam) for sprint planning and execution

## Commands

| Command | Purpose |
|---------|---------|
| `*create-brief` | Aggregate all findings into project brief |
| `*create-prd` | Create Product Requirements Document |
| `*create-roadmap` | Create product roadmap (NOW/NEXT/LATER) |
| `*create-epic` | Create single epic with stories for SM |
| `*stakeholder-update` | Generate status update |
| `*change-assessment` | Analyze scope change impact |

## Skills Used

- [[Strategic Analysis]] - Roadmap prioritization, trade-offs
- [[User Research]] - User story validation
- [[Root Cause]] - True requirements (JTBD)
- [[Ideation]] - Feature brainstorming
- [[Devils Advocate]] - Scope challenge, MVP validation
- [[Role Playing]] - Stakeholder perspectives

## Sub-agents Used

Only for heavy I/O when gaps exist:
- [[Market Researcher]] - Market context
- [[Competitive Analyzer]] - Competitive positioning
- [[Knowledge Harvester]] - External requirements

## Workflow

Paul follows a synthesis-first workflow:

1. **Search KB** - Always aggregate existing research first
2. **Identify Gaps** - Note what's missing vs what exists
3. **Delegate if needed** - Sub-agents only for gaps
4. **Synthesize** - Combine all sources into unified artifact
5. **Validate** - Challenge with [[Devils Advocate]]
6. **Store** - Save with user permission

## Working Memory

PM uses [[Working Memory]] for persistent epic and story tracking.

**What PM Creates in Working Memory:**
- **Epics** - Container for related stories, with description from PRD
- **Stories** - Child items with acceptance criteria, sized for 2-4 hour sessions
- **Dependencies** - Blocking relationships between stories

**Workflow:**
1. Create PRD in Knowledge Memory (rich documentation)
2. Create epic in Working Memory with description from PRD
3. Create stories as children with acceptance criteria
4. Map blocking dependencies between stories

**SM Handoff:** Epic in Working Memory with linked stories. SM queries ready stories for sprint planning.

## Storage Structure

| Content | Location |
|---------|----------|
| Project briefs | Knowledge Memory `research/product/` |
| Stakeholder updates | Knowledge Memory `research/product/` |
| PRDs, roadmaps | Knowledge Memory `requirements/` |
| Epics | Working Memory (type: epic) |
| Stories | Working Memory (type: feature) |

## SM Handoff

Epics include "Ready for Sprint Planning" status with:
- All stories defined with acceptance criteria
- Dependencies mapped between stories
- Technical context from [[Architect]]
- Notes for [[SM|Scrum Master]]

## Status Updates

Status flows UP to [[Analyst]]:

- **Epic completion** - When SM closes epics, update roadmap status
- **Roadmap blockers** - Assess impact on product goals
- **Pivot signals** - Flag when assumptions invalidated by implementation learnings

## Integration

### Upstream (PM Consumes)

| Source | Artifacts | When to Use |
|--------|-----------|-------------|
| [[Analyst]] | Market research, competitive analysis, project briefs | Market context and validation |
| [[Architect]] | ADRs, technical constraints, system designs | Technical feasibility |
| [[UX]] | Personas, journey maps, frontend specs | User understanding and UI requirements |

**Important:** Always search memory for existing upstream artifacts before creating PRDs. Reference existing research, link to it, and build on it. This avoids duplication and ensures alignment.

### Downstream (PM Produces)
- [[SM]] (Sam) - Epics with stories, acceptance criteria for sprint planning

## Templates

Location: `{templates}/pm/` (configurable per [[Templates]] pattern)

| Template | Purpose | Key Sections |
|----------|---------|--------------|
| prd-template.md | Product Requirements Document | Problem, Context, Requirements, User Stories, MVP Scope |
| roadmap-template.md | NOW/NEXT/LATER strategic roadmap | Vision, Prioritization, Dependencies, Review Cadence |
| stakeholder-update.md | Status update for business stakeholders | Progress, Metrics, Risks, Decisions Needed |

Note: Epics and stories live in [[Working Memory]], not as separate documents.

See [[Templates]] pattern for the platform-agnostic template approach.

## Relations

- [[Orchestrator]] - Follows orchestrator pattern
- [[Working Memory]] - Uses for epic/story tracking
- [[SM]] - PM hands off epics to SM for execution
- [[Analyst]] - Consumes research from Analyst
- [[Architect]] - Consumes ADRs and constraints from Architect
- [[UX]] - Consumes personas, journey maps, frontend specs from UX
- [[Catalog]] - Listed in component catalog
- [[Templates]] - Uses PM templates
