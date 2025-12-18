---
title: Sub-agent
type: note
permalink: patterns/sub-agent
---

# Sub-agent Pattern

## What is a Sub-agent?

A sub-agent is a **context-isolated worker** that handles heavy I/O tasks that would flood the main conversation context.

**Characteristics:**
- **Context isolation** - Runs separately, doesn't pollute main conversation
- **Heavy I/O** - Web searches, data gathering, external content processing
- **Returns synthesis** - Summarizes findings, doesn't return raw data
- **No reasoning** - Cannot call skills, cannot make decisions
- **Disposable process** - The work is throwaway, only the output matters

## Core Principle

**Isolate the noise, return the signal.** Sub-agents handle heavy I/O that would flood context, then synthesize findings for the orchestrator.

## When to Create a Sub-agent

Create a sub-agent when ALL of these are true:

1. **Heavy I/O** - Multiple web searches, large data gathering
2. **Context would flood** - Raw results would overwhelm conversation
3. **Process is disposable** - Only the synthesis matters
4. **No reasoning needed** - Doesn't need skills or decision-making

## When NOT to Create a Sub-agent

**Don't create a sub-agent if:**
- No heavy I/O (just code generation → orchestrator does it directly)
- Needs reasoning or skills (must stay in orchestrator)
- Context handoff would lose accuracy
- Interactive (needs follow-up questions)

**Example that failed:** Diagram Generator - no heavy I/O, just code generation. Context handoff loses accuracy about what diagram to create. Orchestrator should do it directly.

## Sub-agent Structure

Every sub-agent note should include:

```markdown
# {Sub-agent Name}

## Core Principle
**Bold statement about what this sub-agent does.** Why context isolation helps.

## Problem
What heavy I/O challenge this sub-agent solves.

## Solution
How it isolates the work and what synthesis it returns.

## Output Schema
What the sub-agent returns to the orchestrator.

## Used By
Which orchestrators delegate to this sub-agent.

## Relations
Links to patterns and related components.
```

## Output Schema

Sub-agents return structured synthesis:

```yaml
result:
  synthesized_content: string    # Markdown ready for storage
  suggested_title: string        # Recommended note title
  suggested_folder: string       # Recommended folder
  key_findings: [list]           # Structured summary
  relations: [list]              # Suggested [[wikilinks]]
  confidence: high|medium|low
```

## Tool Restrictions

Sub-agents should NOT have:
- `Write` tool (orchestrator handles storage)
- MCP Basic Memory tools (orchestrator handles KB)

Sub-agents should have:
- `WebSearch`, `WebFetch` (for research)
- `Read`, `Grep`, `Glob` (for file access if needed)

## Current Sub-agents

### Foundational (all orchestrators)

| Sub-agent | Core Purpose |
|-----------|--------------|
| [[Knowledge Harvester]] | Extract and synthesize external content (docs, articles, patterns) |

### Domain-specific (Analyst)

| Sub-agent | Core Purpose |
|-----------|--------------|
| [[Market Researcher]] | Synthesize market research from web sources |
| [[Competitive Analyzer]] | Synthesize competitive landscape analysis |

## Relations

- [[Orchestrator]] - Orchestrators delegate to sub-agents, handle all KB writes
- [[Architecture]] - Sub-agents are part of the architecture layers
- [[Catalog]] - All sub-agents listed in catalog