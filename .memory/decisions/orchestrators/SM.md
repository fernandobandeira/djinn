---
title: SM
type: note
permalink: decisions/orchestrators/sm
---

# SM (Sam)

## Core Principle
**Break stories into tasks, plan sprints.** PM creates stories; SM breaks them into implementation tasks and organizes sprints. Tasks inherit sprint context from their parent story.
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
Sam follows a breakdown-and-plan workflow:

1. **Discovery** - Find story from PM's epic in Working Memory
2. **Context** - Gather architecture docs, PRD, dependencies
3. **Breakdown** - Create tasks as children of the story (`*breakdown {story-id}`)
4. **Validation** - Auto-validate with [[Devils Advocate]] (`*validate {story-id}`)
5. **Sprint Planning** - Assign validated stories to sprint (`*plan-sprint`)
## Story Validation
Stories are validated with structured criteria:

**ADR Compliance (MUST PASS):**
- KB searched for relevant ADRs before task creation
- Each task's `--design` field cites applicable ADRs
- Task acceptance criteria include ADR compliance checks
- No task contradicts existing architectural decisions

**Critical (MUST PASS):**
- Clear "As a / I want / So that" format
- All acceptance criteria measurable and testable
- Tasks cover all acceptance criteria
- Task designs reference applicable ADRs by name
- Dev Notes provide complete technical context

**Quality (SHOULD PASS):**
- Technical claims verified against architecture
- Test scenarios clearly defined
- Dependencies explicitly mapped
- Risks and mitigation identified
- Patterns from KB followed consistently

**Scoring:**
- **GO** (>=80): All critical + ADR compliance pass, quality >=70%
- **CONDITIONAL** (60-79): All critical pass, quality 50-69%
- **NO-GO** (<60): Any critical/ADR fail or quality <50%
## Sprint Allocation

Sprints balance work types:
- Features: 60-70%
- Tech debt: 20-30%
- Buffer: 10%

Velocity calculated as 3-sprint rolling average.

## Working Memory
SM uses [[Working Memory]] for task and sprint tracking.

**What SM Does in Working Memory:**
- **Query stories** - Find stories from PM with no blockers
- **Break into tasks** - Create tasks as children of stories
- **Sprint labels** - Tag stories with `sprint:N` (tasks inherit from parent)
- **Action items** - Track retrospective action items

**Key Points:**
- PM creates stories; SM creates tasks under them
- Sprint labels go on stories, not tasks
- Tasks inherit sprint context from parent story
- **ADRs are law** - must search KB for ADRs BEFORE creating tasks
- Task `--design` fields MUST reference applicable ADRs by name

**Workflow:**
1. Query Working Memory for ready stories (no blockers)
2. Break stories into implementation tasks (`*breakdown`)
3. Validate story is sprint-ready (`*validate`)
4. Label stories for sprint assignment (`sprint:N`)
5. Create action items from retrospectives
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
- [[PM]] - Epics with stories (PM creates both)
- [[Architect]] - Technical architecture, constraints, patterns

**Downstream (produces for):**
- [[Dev]] - Stories with tasks ready for implementation

**The Handoff:**
- PM creates story with acceptance criteria
- SM breaks story into tasks (`*breakdown`)
- SM validates story is sprint-ready (`*validate`)
- SM assigns story to sprint (`sprint:N` label)
- Dev works on tasks, closes story when all done

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

## Commands
| Command | Purpose |
|---------|---------|
| `*breakdown {story-id}` | Break story into implementation tasks |
| `*validate {story-id}` | Validate story is sprint-ready |
| `*plan-sprint` | Plan next sprint, assign stories |
| `*manage-change` | Analyze sprint change impact |
| `*retrospective` | Facilitate sprint retrospective |