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

## Workflow

Sam follows a story-driven workflow:

1. **Discovery** - Find next story from epic, check existing stories
2. **Context** - Gather architecture docs, PRD, dependencies
3. **Creation** - Use story template, fill from epic content
4. **Validation** - Auto-validate with [[Devils Advocate]]
5. **Storage** - Save if approved with GO decision

## Story Validation

Stories are validated with structured criteria:

**Critical (MUST PASS):**
- Clear "As a / I want / So that" format
- All acceptance criteria measurable and testable
- Tasks cover all acceptance criteria
- Dev Notes provide complete technical context

**Quality (SHOULD PASS):**
- Technical claims verified against architecture
- Test scenarios clearly defined
- Dependencies explicitly mapped
- Risks and mitigation identified

**Scoring:**
- **GO** (>=80): All critical pass, quality >=70%
- **CONDITIONAL** (60-79): All critical pass, quality 50-69%
- **NO-GO** (<60): Any critical fail or quality <50%

## Sprint Allocation

Sprints balance work types:
- Features: 60-70%
- Tech debt: 20-30%
- Buffer: 10%

Velocity calculated as 3-sprint rolling average.

## Working Memory

SM uses [[Working Memory]] for task and sprint tracking.

**What SM Uses Working Memory For:**
- **Query ready stories** - Find stories with no blockers from PM
- **Break into tasks** - Create tasks as children of stories
- **Sprint labels** - Tag stories/tasks for sprint-N
- **Action items** - Track retrospective action items

**Workflow:**
1. Query Working Memory for ready stories (no blockers)
2. Break stories into implementation tasks (parent-child)
3. Label selected items for sprint assignment
4. Create action items from retrospectives

**Retrospective Split:**
- Action items → Working Memory (tracked tasks)
- Lessons learned → Knowledge Memory (`research/retrospectives/`)

## Storage Structure

| Content | Location |
|---------|----------|
| Stories | Working Memory (type: feature) |
| Tasks | Working Memory (type: task, parent-child) |
| Sprints | Working Memory (labels: sprint-N) |
| Action items | Working Memory (type: task) |
| Retrospective insights | Knowledge Memory `research/retrospectives/` |

## Templates

Per [[Templates]] pattern, SM uses:

| Template | Purpose |
|----------|---------|
| `templates/sm/retrospective-template.md` | Retrospective insights format |

Note: Stories and sprints live in [[Working Memory]], not as separate documents.

## Why It Matters

- **Quality handoff** - Dev receives validated, ready stories
- **Strategic planning** - Skills provide structured prioritization
- **Continuous improvement** - Root cause analysis drives real change
- **Story as contract** - Clear specification prevents implementation drift

## Status Updates

Status flows UP to [[PM]]:

- **Story completion** - When Dev closes stories, check epic progress
- **Epic completion** - Close epic when all stories done
- **Sprint blockers** - Escalate blockers affecting sprint goals
- **Velocity data** - Helps PM refine estimates

## Integration

**Upstream (consumes):**
- [[PM]] - Epics with stories, priorities, acceptance criteria
- [[Architect]] - Technical architecture, constraints, patterns

**Downstream (produces for):**
- [[Dev]] - Validated stories ready for implementation

**Status flows UP:**
- Epic completion → PM tracks roadmap
- Blockers → PM adjusts priorities

## Relations

- [[Orchestrator]] - SM follows orchestrator pattern
- [[Working Memory]] - Uses for task/sprint tracking
- [[Templates]] - Uses templates for artifact creation
- [[PM]] - Receives epics from PM
- [[Dev]] - Hands off validated stories to Dev
- [[Catalog]] - Listed in agent catalog