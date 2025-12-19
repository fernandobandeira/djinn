---
title: Sub-agent
type: note
permalink: patterns/sub-agent
---

# Sub-agent Pattern

## What is a Sub-agent?

A sub-agent is a **context-isolated worker** that handles heavy I/O tasks that would flood the main conversation context.

**Characteristics:**
- **One input, one output** - Receives a task, returns synthesis, then gone
- **No user interaction** - Cannot talk to the user or ask follow-up questions
- **Context isolation** - Runs separately, doesn't pollute main conversation
- **Heavy I/O** - Web searches, data gathering, external content processing
- **No reasoning** - Cannot call skills, cannot make decisions
- **Ephemeral** - The process is disposable, only the output matters

## Core Principle

**Isolate the noise, return the signal.** Sub-agents handle heavy I/O that would flood context, then synthesize findings for the orchestrator.

## When to Create a Sub-agent

- Heavy I/O that would flood main context
- Process is disposable, only synthesis matters
- No reasoning or skills needed
- No user interaction required

**If NO to any → Don't create a sub-agent, do it directly in the orchestrator**

See [[Catalog]] for current sub-agents.

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
- Memory/storage write access (orchestrator handles that)

Sub-agents should have:
- Web search and fetch (for research)
- File reading (for local content if needed)

## Relations

- [[Orchestrator]] - Orchestrators delegate to sub-agents, handle all KB writes
- [[Architecture]] - Sub-agents are part of the architecture layers
- [[Catalog]] - All sub-agents listed in catalog