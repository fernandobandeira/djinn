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
| Parent-child | `--deps parent-child:{id}` |
| Blocks | `--deps blocks:{id}` |
| Discovered-from | `--deps discovered-from:{id}` |
| Ready work | `bd ready --json` |
| Sprint label | `bd label add {id} sprint-N` |
| Claim | `bd update {id} --status in_progress` |
| Complete | `bd close {id} --reason "..."` |

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

Beads issues have these fields that map to Working Memory concepts:

| Field | Purpose |
|-------|---------|
| `title` | Issue name (required) |
| `description` | Full details, context |
| `acceptance_criteria` | What "done" means |
| `design` | Technical approach |
| `notes` | Additional context |
| `status` | open, in_progress, blocked, closed |
| `priority` | 0-4 (0=critical, 4=backlog) |
| `assignee` | Who's working on it |
| `labels` | Sprint tags, categories |

## Common Workflows

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

- `blocks` - Hard dependency (B can't start until A is done)
- `parent-child` - Hierarchy (epic → story → task)
- `discovered-from` - Traceability (bug found while working on task)
- `related` - Soft link for context

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