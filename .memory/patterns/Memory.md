---
title: Memory
type: note
permalink: patterns/memory
---

# Memory Pattern

## What is Memory?

The shared knowledge base that serves BOTH human and AI.

**Characteristics:**
- **Dual audience** - Humans review in Obsidian, AI reads/writes via tools
- **Source of truth** - Docs define intent, implementation follows
- **Human-owned** - AI generates, human reviews and approves
- **Obsidian-compatible** - Markdown + wikilinks, viewable in Obsidian
- **Tool-accessed only** - Never manually edit files

## Core Principle

**Docs are the source of truth. Implementation is ephemeral.**

The memory defines WHAT and WHY. Implementation (code, configs, prompts) defines HOW. When implementation needs to change, the docs guide the change. When docs are clear, implementation can be regenerated.

## The Contract

| Role | Responsibility |
|------|----------------|
| **Human** | Reviews docs, ensures accuracy, owns decisions |
| **AI** | Maximizes doc quality, follows existing patterns, links everything |
| **Implementation** | Can change freely as long as it follows docs |

This creates a virtuous cycle:
1. AI drafts high-quality docs
2. Human reviews and refines
3. Docs become trusted source of truth
4. Implementation follows docs
5. Changes to implementation reference docs for guidance

## AI Doc Quality Standards

When writing to memory, AI must apply maximum effort:

1. **Search first** - Check what exists before creating
2. **Deep think** - Maximum effort on structure and clarity
3. **Link everything** - Use wikilinks to connect related content
4. **Follow patterns** - Match existing style and structure
5. **Serve both audiences** - Clear for human review, useful for AI reference
6. **Atomic concepts** - One idea per note, linked to related notes
7. **Evergreen** - Write for long-term value, not just current context

The goal: docs so clear and well-structured that:
- Humans can review them efficiently
- AI can reference them for consistent decisions
- New team members can understand the "why"
- Implementation can be regenerated from principles

## Access Rules

Memory must be accessed through tools, never directly.

**ALWAYS:**
- Use memory tools (search, read, write, edit, build_context)
- Let the tool handle indexing and semantic links

**NEVER:**
- Manually read/write/edit files in `.memory/`
- Use file system tools (Read, Write, Edit, Bash) on `.memory/`
- Use shell commands (cat, grep, find) on `.memory/`

**Why:** Manual edits bypass indexing, break semantic links, and corrupt the knowledge graph. The tools maintain the integrity of the knowledge base.

## Human Review Workflow

The human is the final authority on memory content.

```
┌─────────────────────────────────────────────┐
│  AI generates draft → writes to memory      │
└─────────────────────────────────────────────┘
                    ↓
┌─────────────────────────────────────────────┐
│  Human opens Obsidian (or similar tool)     │
│  • Browse wikilinks graph                   │
│  • Review new/changed notes                 │
│  • Validate accuracy and relevance          │
│  • Refine wording and structure             │
└─────────────────────────────────────────────┘
                    ↓
┌─────────────────────────────────────────────┐
│  Memory becomes trusted source of truth     │
│  • AI references for consistent decisions   │
│  • Implementation follows documented intent │
└─────────────────────────────────────────────┘
```

**Obsidian is recommended** for human review because:
- Native markdown support
- Visual wikilink graph navigation
- Full-text search across notes
- Tags and folder organization
- Works offline with local files

## Why This Matters

- **Compound knowledge** - Insights accumulate instead of disappearing after each conversation
- **Consistent decisions** - AI references past rationale, avoids contradicting earlier choices
- **Team alignment** - Anyone can read the "why" behind decisions without tribal knowledge
- **Regenerable implementation** - If docs are clear, code/configs can be recreated from principles
- **Reduced context burden** - AI searches memory instead of re-exploring codebase each time

## Relations

- [[Project]] - Docs-first philosophy is core to Djinn
- [[Architecture]] - Memory supports all architecture layers
- [[Orchestrator]] - Orchestrators are responsible for writing to memory
- [[Skill]] - Skills are documented patterns in memory
- [[Sub-agent]] - Sub-agents return synthesis for orchestrators to write