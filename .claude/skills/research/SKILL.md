---
name: research
description: Discover and gather knowledge effectively using Basic Memory semantic search and external harvesting. Use when user wants to "search for", "find existing", "what do we know about", "research", "look up", "check if we have", or needs to discover existing knowledge before creating something new.
allowed-tools: Read, Grep, Glob, WebFetch, WebSearch
---

# Research - Knowledge Discovery Skill

Discover existing knowledge before creating new content. Search the knowledge base semantically, harvest external sources, and synthesize findings.

## Quick Start

1. **Always search first** before creating anything
2. Select technique based on what you need
3. Evaluate sources and synthesize findings

## Technique Selection

| Need | Use | Why |
|------|-----|-----|
| Find existing internal knowledge | KB Search | Semantic search across project notes |
| Gather external documentation | Knowledge Harvesting | Fetch and store external sources |
| Assess what you found | Source Evaluation | Determine relevance and quality |

**Default**: Always start with KB Search. Only harvest externally if internal knowledge is insufficient.

## Techniques

### KB Search
Semantic search across all project notes using Basic Memory.

Read [cookbook/kb-search.md](./cookbook/kb-search.md)

### Knowledge Harvesting
Extract and store knowledge from external sources into Basic Memory.

Read [cookbook/knowledge-harvesting.md](./cookbook/knowledge-harvesting.md)

### Source Evaluation
Assess relevance, quality, and applicability of discovered knowledge.

Read [cookbook/source-evaluation.md](./cookbook/source-evaluation.md)

## Core Principles

1. **Search before create** - Always check what exists first
2. **Internal first** - Prefer existing project knowledge over external
3. **Semantic over keyword** - Use meaning-based queries, not just keywords
4. **Link everything** - Use [[wikilinks]] to connect related notes
5. **Read full notes** - Search returns previews; read full docs for context

## Basic Memory Commands Reference

```
# Search notes semantically
mcp__basic-memory__search_notes(query="your query here")

# Read a specific note
mcp__basic-memory__read_note(permalink="note-permalink")

# Write a new note (with links)
mcp__basic-memory__write_note(
    title="Note Title",
    content="Content with [[links]] to other notes",
    folder="decisions"  # optional subfolder
)

# See recent activity
mcp__basic-memory__recent_activity()

# Get project canvas (overview)
mcp__basic-memory__canvas()
```

## Project Knowledge Structure

Notes are stored in `.memory/` with this structure:

```
.memory/
├── project.md              # Vision, goals, overview
├── architecture.md         # System design
├── decisions/              # ADRs and key decisions
│   └── auth-strategy.md
├── patterns/               # Documented patterns
│   └── error-handling.md
├── research/               # Research outputs
│   └── market-analysis.md
├── context/                # Current state
│   └── current-focus.md
└── sessions/               # Brainstorming sessions
```

## When Research Works Best

- Before creating ADRs, patterns, or documentation
- When exploring a new problem space
- When you need prior art or examples
- Before brainstorming (know what's been tried)
- When debugging (search for similar issues)

## The Research Mindset

**Brownfield thinking**: Assume relevant knowledge already exists. Your job is to find it, build on it, and extend it - not recreate it.

Questions to ask:
- "What do we already know about this?"
- "Has this been solved before?"
- "What patterns exist?"
- "What decisions were already made?"
