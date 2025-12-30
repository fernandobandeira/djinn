---
title: Dev
type: note
permalink: decisions/orchestrators/dev
---

# Dev (Dave)

## Core Principle
**Work on tasks only. Close story when all tasks done.** Dev implements tasks created by SM, not stories directly. The story's acceptance criteria validate the work; tasks are the implementation steps.
## Problem

Implementation quality varies without clear specification:
- Developers start coding without a validated story
- Tests written after code (or not at all)
- Code review done ad-hoc without structured critique
- Implementation drifts from requirements

## Solution

Dave is a Developer persona that implements validated stories with TDD discipline.

**Implementation:** `/dev`

**What Dave Does:**
- Receives validated stories from [[SM]] - the story is the contract
- TDD-driven implementation (test first, then code)
- Uses skills for code critique and debugging
- Searches KB for ADRs and patterns before implementation

**Skills Dave Uses:**
- [[Devils Advocate]] - Code critique, pre-mortem on implementations, edge case analysis
- [[Root Cause]] - Debugging, Five Whys for failures, First Principles for design

**Sub-agents:**
- [[Knowledge Harvester]] (shared) - Research for libraries, frameworks, patterns

## Workflow
Dave follows a task-focused workflow:

1. **Sprint** - View current sprint's stories (`*sprint`)
2. **Pick** - Claim a story to work on (`*pick {story-id}`)
3. **Next** - Claim next ready task from story (`*next`)
4. **TDD Cycle** - Red (failing test) → Green (minimal code) → Refactor
5. **Done** - Complete task, get next or close story (`*done`)
6. **Review** - Self-review using [[Devils Advocate]]

**Key Points:**
- Dev works on tasks, not stories directly
- Tasks are created by SM; Dev only creates discovered issues
- When all tasks complete, Dev closes the story
- Sprint labels are on stories; tasks inherit context
## Quality Gates

Dev uses embedded checklists for quality assurance:

- **Complexity Estimation** - Assess scope, technical, and risk factors before implementation
- **Implementation Quality** - Code quality, architecture compliance, test coverage, security

## TDD Discipline

The TDD cycle follows Red-Green-Refactor:

- **Red** - Write failing test that describes expected behavior
- **Green** - Write minimal code to make the test pass
- **Refactor** - Clean up while keeping tests green

## Working Memory
Dev uses [[Working Memory]] for task tracking and discovery logging.

**What Dev Does in Working Memory:**
- **Query sprint** - Find stories in current sprint
- **Pick story** - Choose story to work on, see its tasks
- **Claim task** - Mark task as in_progress
- **Track discoveries** - Log bugs/issues found during implementation
- **Complete task** - Mark done with reason
- **Close story** - When all tasks complete

**Key Points:**
- Dev never creates stories or tasks (except discovered issues)
- Tasks are created by SM under stories
- Sprint context comes from story label, not task label

**Commands:**
| Command | Purpose |
|---------|---------|
| `*sprint` | Show current sprint's stories |
| `*pick {story-id}` | Claim story, show its tasks |
| `*next` | Claim next ready task |
| `*done` | Complete task, prompt for next or close story |
| `*test` | TDD cycle |
| `*implement` | Continue implementation |
| `*review` | Code review with devils-advocate |
| `*validate` | Validate against acceptance criteria |
## Storage Structure

| Content | Location |
|---------|----------|
| Stories | Working Memory (from SM) |
| Tasks | Working Memory (type: task) |
| Discovered issues | Working Memory (type: bug, discovered-from link) |
| Implementation notes | Knowledge Memory `decisions/` |
| Code changes | Codebase directly |

## Why It Matters

- **Story-driven** - Implementation follows validated specification
- **TDD discipline** - Tests first prevents fragile code
- **Structured critique** - Skills provide rigorous review framework
- **Clear contract** - SM validates, Dev implements

## Status Updates

Status flows UP to [[SM]]:

- **Task completion** - Mark task done in [[Working Memory]]
- **Story completion** - Mark story done when all tasks complete
- **Blockers** - Flag blocked status and create blocking issue

## Integration
**Upstream (consumes):**
- [[SM]] - Stories with tasks (SM creates both from PM's stories)
- [[Architect]] - ADRs, technical constraints, patterns

**Downstream (produces for):**
- Users - Working, tested code
- [[SM]] - Closed tasks and stories for progress tracking

**The Handoff:**
- SM breaks story into tasks, validates, assigns to sprint
- Dev queries stories by sprint label
- Dev picks story, works its tasks
- Dev closes story when all tasks complete
- Status flows up to SM

**Status flows UP:**
- Task completion → SM sees progress
- Story completion → SM tracks epic progress
- Blockers → SM escalates if affecting sprint
## Relations

- [[Orchestrator]] - Dev follows orchestrator pattern
- [[Working Memory]] - Uses for task tracking and discoveries
- [[SM]] - Receives validated stories from SM
- [[Architect]] - Consumes architecture context
- [[Catalog]] - Listed in agent catalog