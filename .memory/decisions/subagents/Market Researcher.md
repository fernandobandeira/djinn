---
title: Market Researcher
type: note
permalink: decisions/subagents/market-researcher
---

# Market Researcher

**Type:** Sub-agent (Context Isolation)

## Core Principle

**Drown in the research, surface with insights.** Handle the heavy I/O of market research so the orchestrator stays focused on analysis and decisions.

## Problem

Heavy web research floods main conversation context:
- Many web fetches pollute the context window
- Raw research data overwhelms the conversation
- Main reasoning gets buried in I/O noise
- User loses track of the analytical thread

## Solution

Isolated sub-agent that performs research and returns synthesis.

**What It Does:**
- Receives research topic and focus areas
- Searches web for relevant market data
- Fetches and analyzes key sources
- Returns summarized findings (not raw data)

**Output Format:**
```yaml
status: success | partial | failure
synthesized_content: string  # Ready for KB storage
suggested_title: string
suggested_folder: research
key_findings: [list]
```

## Why It Matters

- **Context isolation** - Main conversation stays clean
- **Heavy I/O handled** - Many fetches don't pollute context
- **Synthesis returned** - Only insights come back, not raw data
- **Orchestrator writes** - Ana decides what to save to KB

## Used By

- [[Analyst]] - Market research via `*research` command

## Relations

- [[Architecture]] - Part of sub-agent layer
- [[Catalog]] - Listed in component index
- [[Sub-agent]] - Follows sub-agent pattern
- [[Orchestrator]] - Returns synthesis, orchestrator writes to KB