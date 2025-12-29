---
title: Beads
type: note
permalink: decisions/implementations/beads
---

# Beads

Implementation of [[Working Memory]] using Beads CLI.

## Context

Working Memory pattern requires persistent work tracking. Beads is a git-backed issue tracker optimized for AI agents.

## Decision

Use beads CLI (`bd`) to implement Working Memory.

## Knowledge vs Working Memory

**Knowledge Memory (Basic Memory)** - Rich documentation humans review:
- PRDs, architecture decisions, research
- Sprint plans (optional rich docs)
- Retrospective insights
- Stored in `.memory/` folders

**Working Memory (beads)** - Work items agents coordinate:
- Epics, stories, tasks, bugs
- Status, priority, dependencies
- Sprints (via labels: `sprint-N`)
- Stored in `.beads/`

**Principle**: Work items live in beads. Documentation lives in Basic Memory. Don't duplicate.

**AGENTS.md Removed**: Beads generates AGENTS.md via `bd onboard`, but we remove it because:
1. Quick reference duplicates orchestrator docs
2. Session completion belongs in orchestrator workflows
3. One source of truth per concept

## Mapping
| Working Memory Concept | Beads Implementation |
|-----------------------|---------------------|
| Epic | `bd create -t epic` |
| Story/Feature | `bd create -t feature` |
| Task | `bd create -t task` |
| Bug | `bd create -t bug` |
| **Hierarchy** | `--parent {id}` (child belongs to parent) |
| **Blocks** | `--deps blocks:{id}` (A depends on B) |
| **Discovered-from** | `--deps discovered-from:{id}` |
| Ready work | `bd ready --json` |
| Sprint label | `bd label add {id} sprint-N` |
| Claim | `bd update {id} --status in_progress` |
| Complete | `bd close {id} --reason "..."` |

### Hierarchy vs Dependencies

**Hierarchy** (`--parent`) - Organizational containment:
- Story belongs to Epic: `bd create "Story" -t feature --parent {epic-id}`
- Task belongs to Story: `bd create "Task" -t task --parent {story-id}`

**Dependencies** (`--deps`) - Execution order:
- B blocks A: `bd dep add {A-id} {B-id} --type blocks`
- Bug discovered from Task: `--deps discovered-from:{task-id}`
## Setup

```bash
# Install beads (one-time)
go install github.com/steveyegge/beads/cmd/bd@latest

# Initialize in project
bd init --quiet

# Remove generated agent files (orchestrators are our source of truth)
rm -f AGENTS.md @AGENTS.md
```

## Issue Fields
Beads issues support rich structured data:

| Field | Flag | Purpose |
|-------|------|---------|
| `title` | positional | Issue name (required) |
| `description` | `-d` / `--description` | What and why - full context |
| `acceptance` | `--acceptance` | Testable completion criteria |
| `design` | `--design` | Technical approach, constraints |
| `status` | `--status` | open, in_progress, blocked, closed |
| `priority` | `-p` | 0-4 (0=critical, 4=backlog) |
| `assignee` | `-a` | Who's working on it |
| `labels` | `-l` | Sprint tags, categories |
| `estimate` | `-e` | Time estimate in minutes |

### Rich Issue Creation

Always populate description, acceptance, and design fields:

```bash
bd create "Story: Login UI" -t feature --parent {epic-id} -p 1 \
  -d "As a user, I want to log in so I can access my account. Entry point for auth." \
  --design "LoginForm component with useAuth hook. Redirect to dashboard on success." \
  --acceptance "Given valid credentials, user redirects to dashboard
Given invalid credentials, error displays without reload" \
  --json
```

This provides context for anyone (human or agent) picking up the work.
## Common Workflows
### PM: Create Epic with Stories

```bash
# Create epic with full context
bd create "Epic: User Authentication" -t epic -p 1 \
  -d "Complete auth system for secure access. Required for MVP - blocks user features." \
  --design "JWT with refresh rotation. OAuth2 ready for future social login." \
  --acceptance "- Users can register and log in
- Password reset works end-to-end
- Protected routes reject unauthenticated requests" \
  --json

# Create stories as children (use --parent for hierarchy)
bd create "Story: Login UI" -t feature --parent {epic-id} -p 1 \
  -d "As a user, I want to log in so I can access my account." \
  --design "LoginForm with useAuth hook. Redirect to dashboard on success." \
  --acceptance "Given valid credentials, redirect to dashboard
Given invalid credentials, show error without reload" \
  --json
```

### SM: Break Story into Tasks

```bash
# Get story details
bd show {story-id} --json

# Create tasks as children (use --parent for hierarchy)
bd create "Task: Create login form component" -t task --parent {story-id} -p 1 \
  -d "Build LoginForm with email/password fields and validation states." \
  --design "Formik + Yup validation. Follow AuthCard layout pattern." \
  --acceptance "- Form renders with email/password fields
- Validation runs on blur
- Submit disabled during API call" \
  --json

# Add blocking dependency if needed (API task needs form first)
bd dep add {api-task-id} {form-task-id} --type blocks
```

### Dev: Implement and Track Discoveries

```bash
# Get next ready task
bd ready --limit 1 --json

# Claim it
bd update {id} --status in_progress

# Track discovered bug with context
bd create "Bug: Login fails with special chars" -t bug \
  --deps discovered-from:{current-task-id} -p 2 \
  -d "Passwords with '&' or '+' fail. URL encoding issue in API call." \
  --design "Fix: encodeURIComponent on password before sending." \
  --acceptance "- Password 'test&123' works
- Password 'test+456' works
- No regression on standard passwords" \
  --json

# Complete task
bd close {id} --reason "Implemented and tested"
```

### Query Commands

```bash
bd ready --json              # Ready work (no blockers)
bd blocked --json            # Blocked issues
bd list --label sprint-N     # Sprint items
bd dep tree {epic-id}        # Hierarchy view
```
### PM: Create Epic

```bash
# Create epic with description
bd create "Epic: {title}" -t epic -p {priority} -d "{description}" --json

# Create stories as children
bd create "Story: {title}" -t feature --deps parent-child:{epic-id} -p {priority} --json
```

### SM: Plan Sprint

```bash
# Find ready stories
bd ready --type feature --json

# Assign to sprint
bd label add {story-id} sprint-N

# Break story into tasks
bd create "Task: {title}" -t task --deps parent-child:{story-id} -p {priority} --json
```

### Dev: Implement

```bash
# Get next task
bd ready --limit 1 --json

# Claim it
bd update {id} --status in_progress

# Track discovered issue
bd create "Found: {issue}" -t bug --deps discovered-from:{id} -p {priority} --json

# Complete
bd close {id} --reason "Implemented and tested"
```

### Query Commands

```bash
# All ready work (no blockers)
bd ready --json

# Blocked issues
bd blocked --json

# Sprint items
bd list --label sprint-N --json

# Epic tree (hierarchy + dependencies)
bd dep tree {epic-id}

# Search
bd list --status open --priority 1 --json
```

## Status Flow

```
open → in_progress → closed
  ↓         ↓
blocked ←───┘
```

## Dependencies
### Dependency Types

| Type | Purpose | When to Use |
|------|---------|-------------|
| `blocks` | Hard dependency | B can't start until A done |
| `discovered-from` | Traceability | Bug/task found while working |
| `related` | Soft link | Context, no blocking |

### Hierarchy vs Dependencies

**Hierarchy** is NOT a dependency. Use `--parent` during creation:
- Story under Epic: `--parent {epic-id}`
- Task under Story: `--parent {story-id}`

**Dependencies** control execution order. Use `--deps` or `bd dep add`:
- Story B blocked by Story A: `bd dep add {B-id} {A-id} --type blocks`
- Bug found from Task: `--deps discovered-from:{task-id}`

### Dependency Semantics

`bd dep add A B --type blocks` means:
- "A depends on B"
- "B blocks A"
- "B must complete before A can start"
## Git Integration

Beads stores data in `.beads/issues.jsonl` (git-versioned). This provides:
- Version history for all changes
- Merge-safe JSONL format
- Automatic sync via daemon
- No external service required

## Relations

- [[Working Memory]] - Concept pattern this implements
- [[Claude Code]] - Parent implementation doc
- [[PM]] - Product Manager workflow
- [[SM]] - Scrum Master workflow
- [[Dev]] - Developer workflow


## Session Completion ("Landing the Plane")

When ending a work session, orchestrators MUST complete all steps. Work is NOT complete until `git push` succeeds.

### Mandatory Workflow

1. **File issues for remaining work** - Create issues for anything needing follow-up
2. **Run quality gates** (if code changed) - Tests, linters, builds
3. **Update issue status** - Close finished work, update in-progress items
4. **Sync and push**:
   ```bash
   git pull --rebase
   bd sync
   git push
   git status  # MUST show "up to date with origin"
   ```
5. **Clean up** - Clear stashes, prune remote branches
6. **Verify** - All changes committed AND pushed
7. **Hand off** - Provide context for next session

### Critical Rules

- Work is NOT complete until `git push` succeeds
- NEVER stop before pushing - that leaves work stranded locally
- NEVER say "ready to push when you are" - agent MUST push
- If push fails, resolve and retry until it succeeds

### Where This Lives

This session completion discipline is embedded in each orchestrator that produces artifacts (Dev, SM, PM), not in a separate AGENTS.md file. Rationale: orchestrators are the source of truth for their workflows.

The beads `bd onboard` command generates AGENTS.md, but we remove it because:
1. Quick reference duplicates orchestrator docs
2. Session completion belongs in orchestrator workflows
3. One source of truth per concept