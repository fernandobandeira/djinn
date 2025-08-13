# Create Sub-Agent Task

## Purpose
Guide the creation of Claude Code sub-agents with proper frontmatter configuration and delegation patterns.

## Sub-Agent vs Command Agent

### Sub-Agent Characteristics:
- Lives in `.claude/agents/` directory
- Has frontmatter configuration
- Delegatable by Claude automatically
- Specific model selection
- Defined tool set

### Command Agent Characteristics:
- Lives in `.claude/commands/` directory  
- Activated by slash command
- Interactive workflow
- Resource file structure
- Full Claude Code capabilities

## Creation Workflow

### Step 1: Gather Requirements

Ask the user:
1. **Agent Purpose**: What specific task will this sub-agent handle?
2. **Delegation Trigger**: When should Claude delegate to this agent?
3. **Required Tools**: What operations does it need to perform?
4. **Model Preference**: haiku (fast), sonnet (balanced), opus (powerful)?
5. **Proactive Behavior**: Should it activate automatically?

### Step 2: Generate Agent Name
- Use `kebab-case` format
- Be descriptive but concise
- Examples: `code-reviewer`, `test-writer`, `dependency-updater`

### Step 3: Craft Delegation Description

#### Patterns for Effective Delegation:
```yaml
# Proactive pattern
description: Use proactively when user asks to review code quality

# Specialist pattern  
description: Specialist for analyzing Python performance issues

# Task pattern
description: Handles database migration and schema updates

# Conditional pattern
description: Use when user needs API endpoint testing
```

### Step 4: Select Minimal Tools

#### Tool Selection Guide:
```yaml
# Code Analysis
tools: Read, Grep, Glob

# Code Modification
tools: Read, Edit, MultiEdit, Write

# Testing & Debugging
tools: Read, Edit, Bash, Grep

# Documentation
tools: Read, Write, WebFetch

# Full Development
tools: Read, Write, Edit, Bash, Grep, Glob
```

### Step 5: Choose Appropriate Model

```yaml
# Fast, simple tasks
model: haiku

# Balanced performance
model: sonnet  # (default)

# Complex reasoning
model: opus
```

### Step 6: Write System Prompt

#### Template Structure:
```markdown
---
name: [agent-name]
description: [delegation-trigger]
tools: [Tool1, Tool2, ...]
model: [haiku|sonnet|opus]
color: [red|blue|green|yellow|purple|orange|pink|cyan]
---

# Purpose

You are a [role definition]. Your primary responsibility is to [main task].

IMPORTANT: [Critical behavior or constraint]

## Core Responsibilities

1. [First responsibility]
2. [Second responsibility]
3. [Third responsibility]

## Instructions

When invoked, follow these steps:

1. **Analyze**: [First action]
2. **Identify**: [Second action]
3. **Execute**: [Third action]
4. **Report**: [Final action]

## Best Practices

- [Domain-specific best practice]
- [Quality guideline]
- [Performance consideration]

## Output Format

Provide your analysis/results in the following structure:
- **Summary**: [Brief overview]
- **Details**: [Specific findings]
- **Recommendations**: [Next steps]

## Constraints

- Do not [limitation 1]
- Always [requirement 1]
- Ensure [quality check]
```

### Step 7: Create the File

Write to: `.claude/agents/[agent-name].md`

### Step 8: Test Delegation

1. Create a test scenario
2. Verify automatic delegation works
3. Check tool access
4. Validate output format

### Step 9: Document in KB

```bash
# Index the new sub-agent
./.vector_db/kb index --path .claude/agents/[agent-name].md
```

## Common Sub-Agent Patterns

### Code Quality Reviewer
```yaml
name: code-reviewer
description: Use proactively after significant code changes
tools: Read, Grep, Glob
model: sonnet
```

### Test Generator
```yaml
name: test-writer
description: Specialist for creating unit and integration tests
tools: Read, Write, Edit, Grep
model: sonnet
```

### Documentation Writer
```yaml
name: doc-writer  
description: Use when user needs documentation created or updated
tools: Read, Write, Edit, WebFetch
model: haiku
```

### Dependency Manager
```yaml
name: dependency-manager
description: Handles package updates and dependency conflicts
tools: Read, Edit, Bash, Write
model: sonnet
```

## Delegation Keywords

### Proactive Triggers:
- "Use proactively when..."
- "Automatically handle..."
- "Invoke after..."

### Specialist Indicators:
- "Specialist for..."
- "Expert in..."
- "Dedicated to..."

### Task Descriptors:
- "Handles..."
- "Manages..."
- "Responsible for..."

## Quality Checklist

- [ ] Clear, single purpose
- [ ] Descriptive delegation trigger
- [ ] Minimal tool set
- [ ] Appropriate model selection
- [ ] Well-structured prompt
- [ ] Defined output format
- [ ] Best practices included
- [ ] Constraints specified
- [ ] File created in correct location
- [ ] Delegation tested
- [ ] Indexed in KB

## Remember
- Sub-agents are for delegation, not interaction
- Keep them focused and specialized
- Use minimal tools for efficiency
- Test delegation triggers
- Document in knowledge base