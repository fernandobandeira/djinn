---
title: Knowledge Harvester
type: note
permalink: decisions/subagents/knowledge-harvester
---

# Knowledge Harvester

**Type:** Sub-agent (Context Isolation)

## Core Principle

**Harvest the noise, distill the knowledge.** Process external sources and documentation so only structured insights return to the orchestrator.

## Problem

External source gathering pollutes conversation with raw data:
- Harvesting documentation, articles, examples requires many fetches
- Raw web content overwhelms context
- Valuable insights buried in noise
- Synthesis work competes with main reasoning

## Solution

Isolated sub-agent that harvests external sources and returns synthesis.

**What It Does:**
- Receives topic and search queries
- Searches for relevant external sources
- Fetches and extracts key content
- Returns synthesized knowledge (not raw content)

**Output Format:**
```yaml
status: success | failure
synthesized_content: string  # Ready for KB storage
suggested_title: string
suggested_folder: research
sources_harvested: [list of URLs]
key_concepts: [list]
patterns: [list]
```

## Why It Matters

- **Context isolation** - Raw web content stays out
- **Only insights return** - Synthesized, structured knowledge
- **Heavy I/O handled** - Many fetches don't pollute context
- **Orchestrator writes** - Calling agent saves to KB

## Used By

- [[Analyst]] - Harvesting market research, industry articles, documentation
- [[Architect]] - Harvesting technical docs, best practices, architectural patterns
- [[Recruiter]] - Harvesting agent pattern examples, framework documentation

## Relations

- [[Architecture]] - Part of sub-agent layer
- [[Catalog]] - Listed in component index
- [[Sub-agent]] - Follows sub-agent pattern
- [[Orchestrator]] - Returns synthesis, orchestrator writes to KB