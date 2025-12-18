---
title: Competitive Analyzer
type: note
permalink: decisions/subagents/competitive-analyzer
---

# Competitive Analyzer

**Type:** Sub-agent (Context Isolation)

## Core Principle

**Research the competition, return the battlefield map.** Handle the I/O overhead of multi-competitor analysis so the orchestrator can focus on strategic implications.

## Problem

Competitive analysis requires many web fetches that pollute context:
- Researching multiple competitors requires extensive I/O
- Raw competitive data overwhelms conversation
- Comparison tables need many data points gathered
- Analysis gets buried in research noise

## Solution

Isolated sub-agent that gathers competitive data and returns synthesis.

**What It Does:**
- Receives competitor list and analysis criteria
- Researches each competitor
- Compares on specified dimensions
- Returns synthesized analysis (not raw data)

**Output Format:**
```yaml
status: success | partial | failure
synthesized_content: string  # Ready for KB storage
suggested_title: string
suggested_folder: research
competitors: [comparison data]
```

## Why It Matters

- **Context isolation** - Main conversation stays focused
- **Heavy I/O separated** - Multiple competitor research isolated
- **Structured comparison** - Returns organized analysis
- **Orchestrator writes** - Ana decides what to save to KB

## Used By

- [[Analyst]] - Competitive analysis via `*analyze-competition` command

## Relations

- [[Architecture]] - Part of sub-agent layer
- [[Catalog]] - Listed in component index
- [[Sub-agent]] - Follows sub-agent pattern
- [[Orchestrator]] - Returns synthesis, orchestrator writes to KB