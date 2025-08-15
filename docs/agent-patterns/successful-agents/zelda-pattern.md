# Pattern: Zelda - Well-Constrained Knowledge Capture Agent

## Context
- Agent Type: Sub-agent
- Use Case: Automatic knowledge capture during learning sessions
- Constraint Level: Well-balanced

## Pattern Structure

### Frontmatter
```yaml
---
name: zelda
description: "Zettelkasten knowledge capture specialist. Proactively identifies learning moments, creates atomic notes, builds knowledge graphs, and enables spaced repetition during learning sessions."
tools: Read, Write, MultiEdit, Grep, LS, Bash
---
```

## Why This Works

### Constraint Balance
- **Clear Trigger**: "Proactively identifies learning moments" - specific but flexible
- **Defined Scope**: Zettelkasten method - bounded but adaptable
- **Minimal Tools**: Just what's needed for file operations
- **Specific Actions**: Creates notes, builds graphs, enables repetition

### Effective Elements
1. **Description Length**: Comprehensive but concise
2. **Keywords**: "Proactively" ensures automatic delegation
3. **Tool Selection**: Read/Write for notes, Grep/LS for searching, Bash for git
4. **Clear Purpose**: Knowledge capture specialist role

## When to Use This Pattern
- Creating specialized domain experts
- Agents that work alongside users
- Knowledge management tasks
- Learning support systems

## Constraint Score: 8/10
- Purpose clarity: 9/10
- Trigger specificity: 8/10
- Output definition: 8/10
- Scope boundaries: 8/10
- Tool minimalism: 7/10

## Extracted From
- Source: .claude/agents/zelda.md
- Date: 2025-08-15
- Success Indicator: Actively used in Teacher-Zelda architecture