---
title: Working Memory
type: note
permalink: patterns/working-memory
---

# Working Memory Pattern

## What is Working Memory?

The operational work-tracking system that complements Knowledge Memory.

**Characteristics:**
- **Work items** - Epics, stories, tasks with status
- **Hierarchy** - Parent-child relationships (epic → story → task)
- **Dependencies** - Blocking relationships
- **Sprint organization** - Labels for grouping work
- **Multi-session** - Persists across conversations

## Core Principle

**Knowledge Memory stores WHAT and WHY. Working Memory tracks WHO and WHEN.**

- Knowledge Memory = Rich docs humans review
- Working Memory = Work status agents coordinate

## Work Item Types

| Type | Purpose | Example |
|------|---------|---------|
| **Epic** | Large feature container | "User Authentication System" |
| **Story/Feature** | Deliverable unit (2-4 hours) | "Login UI with validation" |
| **Task** | Implementation step | "Create login form component" |
| **Bug** | Defect to fix | "Login fails with special chars" |

## Relationships
### Hierarchy (Containment)

Use hierarchy to organize work items structurally:

| Relationship | Purpose | Example |
|--------------|---------|---------|
| **Epic → Story** | Epic contains stories | "Auth System" contains "Login UI" |
| **Story → Task** | Story contains tasks | "Login UI" contains "Create form component" |

Hierarchy answers: "What work belongs to what?"

### Dependencies (Blocking)

Use dependencies to control work ordering:

| Relationship | Purpose | Example |
|--------------|---------|---------|
| **Blocks** | Hard dependency | Story A must complete before Story B starts |
| **Discovered-from** | Traceability | Bug found while working on Task |

Dependencies answer: "What must complete before this can start?"

### Key Distinction

- **Hierarchy** = organizational structure (what contains what)
- **Dependencies** = execution order (what blocks what)

An epic doesn't "block" its stories - it *contains* them. A task doesn't "depend on" its story - it *belongs to* it.
## Operations

Working Memory supports these operations (implementation varies by tool):

### Query
- **Ready work** - Items with no blockers
- **Blocked work** - Items waiting on dependencies
- **By sprint** - Items tagged for a sprint
- **By status** - Open, in progress, closed

### Create
- **Epic** - Container with child stories
- **Story** - With description, acceptance criteria
- **Task** - With parent link
- **Discovered issue** - Linked to source work

### Update
- **Claim** - Mark as in progress
- **Complete** - Mark as done with reason
- **Add label** - Tag for sprint/category

## Integration with Orchestrators
| Orchestrator | Creates | Works On | Closes |
|--------------|---------|----------|--------|
| **PM** | Epics, Stories | - | Epics (when all stories done) |
| **SM** | Tasks (under stories) | Stories | - |
| **Dev** | Discovered issues only | Tasks | Tasks, Stories (when all tasks done) |

**The Flow:**
1. **PM** creates epic with stories (acceptance criteria defined)
2. **SM** breaks stories into tasks (`*breakdown`), validates, assigns to sprint
3. **Dev** works tasks (`*sprint` → `*pick` → `*next` → `*done`)
4. **Dev** closes story when all tasks complete
5. **SM** closes epic when all stories complete (or PM does)

**Sprint Labels:**
- Stories get sprint labels (`sprint:N`)
- Tasks inherit sprint context from parent story
- Dev queries stories by sprint, then works child tasks
## Graceful Degradation

Working Memory is optional. Without it:
- Orchestrators work with Knowledge Memory only
- No persistent task tracking across sessions
- Use TodoWrite for session-scoped progress

## Relations

- [[Memory]] - Knowledge Memory pattern
- [[Beads]] - Current implementation
- [[PM]] - Product Manager workflow
- [[SM]] - Scrum Master workflow
- [[Dev]] - Developer workflow