---
title: Templates
type: note
permalink: patterns/templates
---

# Templates Pattern

## Core Principle

**Templates are platform-agnostic reference patterns.** They define the structure of artifacts (PRDs, ADRs, Epics) independent of which AI platform implements them.

## Why Separate Templates?

```
djinn/
├── templates/           # Platform-agnostic (reusable anywhere)
│   ├── CLAUDE.md       # Project setup template
│   ├── pm/             # PM artifact structures
│   ├── analyst/        # Analyst artifact structures
│   ├── architect/      # Architect artifact structures
│   └── sm/             # SM artifact structures
├── .claude/            # Claude Code implementation (platform-specific)
└── .memory/            # Framework knowledge
```

**Separation enables:**
1. **Reusability** - Same templates work with Claude Code, Cursor, custom agents
2. **Clear ownership** - Templates define WHAT, implementation defines HOW
3. **Future-proofing** - New platforms adopt existing templates
4. **Independent evolution** - Template structure vs implementation can change separately

## Template Categories

### PM Templates (`templates/pm/`)

| Template | Purpose |
|----------|---------|
| prd-template.md | Product Requirements Document - synthesizes research into requirements |
| epic-template.md | Epic with stories - sized for SM handoff (2-4 hour stories) |
| roadmap-template.md | NOW/NEXT/LATER strategic roadmap |
| stakeholder-update.md | Status update for business stakeholders |

### Analyst Templates (`templates/analyst/`)

| Template | Purpose |
|----------|---------|
| project-brief.md | Project brief structure - objectives, scope, timeline, risks |
| brainstorming-output.md | Structured brainstorming session output |

### Architect Templates (`templates/architect/`)

| Template | Purpose |
|----------|---------|
| adr-template.md | Architecture Decision Record - context, decision, consequences |
| pattern-template.md | Reusable pattern documentation |
| rfc-template.md | Request for Comments - proposals for discussion |
| runbook-template.md | Operational runbook for services |

### SM Templates (`templates/sm/`)

| Template | Purpose |
|----------|---------|
| story-template.md | User story with acceptance criteria, tasks, dev notes |
| sprint-template.md | Sprint plan with goals, velocity, story allocation |
| retrospective-template.md | Sprint retrospective with metrics, analysis, actions |

### Setup Templates (`templates/`)

| Template | Purpose |
|----------|---------|
| CLAUDE.md | Project setup template with Basic Memory configuration |

## Configuration

Templates path is configured in CLAUDE.md:

```markdown
## Templates Configuration

**Location**: `templates/`
```

**Options:**
- `templates/` - Local project templates (default)
- `~/.djinn/templates/` - Global templates (for global install)

Local templates override global ones, enabling per-project customization.

## How Orchestrators Use Templates

1. **Read path** from CLAUDE.md `Templates Configuration`
2. **Load template** from `{path}/{orchestrator}/`
3. **Fill structure** with user-specific content
4. **Store output** in `.memory/` via Basic Memory

```markdown
## Resources

**Templates**: `templates/pm/`  (or path from CLAUDE.md)
- prd-template.md - Product Requirements Document structure
- epic-template.md - Epic with stories for SM handoff
```

## Installation Scenarios

### Per-Project (Default)
```
my-project/
├── templates/          # Local templates (copied or customized)
├── .claude/            # From djinn
├── .memory/
└── CLAUDE.md           # Location: templates/
```

### Global Install
```
~/.djinn/
└── templates/          # Global templates

~/.claude/              # Implementation (copied from djinn)

my-project/
├── .memory/
└── CLAUDE.md           # Location: ~/.djinn/templates/
```

### Hybrid (Global + Local Overrides)
```
~/.djinn/templates/     # Global defaults

my-project/
├── templates/          # Local overrides (takes precedence)
├── .memory/
└── CLAUDE.md           # Location: templates/
```

## Template vs Checklist

| Type | Location | Purpose |
|------|----------|---------|
| Template | `templates/` | Structure for creating artifacts |
| Checklist | `.claude/resources/` | Operational verification during workflows |

Templates are filled to create something new. Checklists verify existing work.

## Extending Templates

When creating new orchestrators:
1. Create `templates/{orchestrator}/` folder
2. Add templates for each artifact type
3. Document templates in orchestrator's decision note
4. Reference from orchestrator command

## Relations

- [[Orchestrator]] - Orchestrators load and fill templates
- [[PM]] - Uses PM templates
- [[Analyst]] - Uses Analyst templates  
- [[Architect]] - Uses Architect templates
- [[SM]] - Uses SM templates
- [[Claude Code Implementation]] - Platform-specific implementation