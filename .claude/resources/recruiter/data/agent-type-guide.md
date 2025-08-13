# Agent Type Selection Guide

## CRITICAL DISTINCTION

### 🎯 Sub-Agents (One-Shot Execution)
- **Execution**: Single run, produces final output
- **Interaction**: NONE - completes task and returns result
- **Location**: `.claude/agents/`
- **Activation**: Automatic delegation by Claude
- **Use Cases**: 
  - Code review → Returns review report
  - Test generation → Creates tests and exits
  - Documentation → Writes docs and completes
  - Analysis → Provides analysis and done

### 💬 Command Agents (Interactive)
- **Execution**: Multi-turn conversation
- **Interaction**: Back-and-forth dialogue with user
- **Location**: `.claude/commands/`
- **Activation**: User types slash command (e.g., `/analyst`)
- **Use Cases**:
  - Brainstorming sessions
  - Step-by-step guidance
  - Iterative refinement
  - Exploration and research

## ALWAYS ASK FIRST

When user says "create an agent", ALWAYS ask:

```
I need to understand what type of agent you want to create:

1. **Sub-Agent** (One-Shot)
   - Performs a specific task automatically
   - Returns a final result/report
   - No interaction with user
   - Example: "Review this code" → Returns review

2. **Command Agent** (Interactive)
   - Engages in dialogue with user
   - Multiple back-and-forth exchanges
   - Guides through processes
   - Example: "/analyst" → Facilitates brainstorming

Which type would you like to create?
```

## Decision Matrix

| Criteria | Sub-Agent | Command Agent |
|----------|-----------|---------------|
| User interaction needed? | No → Sub-Agent | Yes → Command |
| Multiple steps with user input? | No → Sub-Agent | Yes → Command |
| Needs refinement/iteration? | No → Sub-Agent | Yes → Command |
| Single deliverable output? | Yes → Sub-Agent | No → Command |
| Facilitation/guidance? | No → Sub-Agent | Yes → Command |

## Examples to Clarify

### Good Sub-Agent Candidates:
- "Review this PR for security issues" → One-shot review
- "Generate tests for this module" → Creates tests, done
- "Format this code" → Formats and returns
- "Analyze dependencies" → Returns analysis report

### Good Command Agent Candidates:
- "Help me brainstorm features" → Interactive session
- "Guide me through architecture" → Step-by-step process
- "Let's refactor this code" → Iterative refinement
- "Research market options" → Exploration with user

## Tool Selection by Type

### Sub-Agents (Minimal Tools):
```yaml
# Read-only analysis
tools: Read, Grep, Glob

# Code modification
tools: Read, Edit, Write

# With execution
tools: Read, Edit, Bash
```

### Command Agents (Full Capabilities):
```yaml
# Full Claude Code access
# Tools specified in interaction
# Access to all available tools
# KB integration
# Resource file loading
```

## Syntax Differences

### Sub-Agent Header:
```yaml
---
name: security-reviewer
description: Use proactively for security review of code changes
tools: Read, Grep, Glob
model: sonnet
---
```

### Command Agent Header:
```markdown
# Security Analyst Agent - Sam

## Activation
You are Sam, the Security Analyst. Your role is to guide users through security assessments.
```

## Remember
- ALWAYS ask: Sub-agent or Command?
- Sub-agents = One-shot task execution
- Commands = Interactive dialogue
- Don't assume - clarify first
- Match type to use case