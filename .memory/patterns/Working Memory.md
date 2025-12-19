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

| Relationship | Purpose | Example |
|--------------|---------|---------|
| **Parent-child** | Hierarchy | Epic → Stories → Tasks |
| **Blocks** | Hard dependency | Story A blocks Story B |
| **Discovered-from** | Traceability | Bug found while working on Task |

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

| Orchestrator | Uses Working Memory For |
|--------------|------------------------|
| **PM** | Create epics with stories, track completion |
| **SM** | Break stories into tasks, organize sprints |
| **Dev** | Claim tasks, track progress, log discoveries |

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