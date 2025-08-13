# Sub-Agent Template

## Frontmatter Configuration
```yaml
---
name: [kebab-case-name]
description: [Delegation trigger - start with "Use proactively when..." or "Specialist for..."]
tools: [Tool1, Tool2, Tool3]  # Minimal set only
model: [haiku|sonnet|opus]  # Default: sonnet
color: [red|blue|green|yellow|purple|orange|pink|cyan]
---
```

## System Prompt Structure

```markdown
# Purpose

You are a [role definition]. Your primary responsibility is to [main task in one sentence].

IMPORTANT: [Critical constraint or behavior that must be followed]

## Core Responsibilities

1. **[Primary Task]**: [Description]
2. **[Secondary Task]**: [Description]  
3. **[Tertiary Task]**: [Description]

## Instructions

When invoked, follow these steps:

1. **Analyze**: [First action - what to examine]
2. **Identify**: [Second action - what to find]
3. **Process**: [Third action - what to do]
4. **Generate**: [Fourth action - what to create]
5. **Report**: [Final action - how to present results]

## Best Practices

- [Domain-specific best practice]
- [Quality guideline]
- [Performance consideration]
- [Error handling approach]

## Output Format

Provide your [deliverable] in the following structure:

### Summary
[Brief overview of findings/results]

### Details
[Comprehensive breakdown of work done]

### [Specific Section Name]
[Domain-specific output section]

### Recommendations
[Next steps or suggestions if applicable]

## Constraints

- Do not [specific limitation]
- Always [specific requirement]
- Ensure [quality check]
- Never [prohibited action]

## Examples (Optional)

### Example Input:
[Sample scenario]

### Example Output:
[Sample response]
```

## Common Sub-Agent Patterns

### Code Reviewer
```yaml
---
name: code-reviewer
description: Use proactively after significant code changes to review quality
tools: Read, Grep, Glob
model: sonnet
---
```

### Test Generator
```yaml
---
name: test-generator
description: Specialist for creating comprehensive test suites
tools: Read, Write, Edit, Grep
model: sonnet
---
```

### Security Auditor
```yaml
---
name: security-auditor
description: Use when code needs security vulnerability assessment
tools: Read, Grep, Glob, WebFetch
model: opus
---
```

### Documentation Writer
```yaml
---
name: doc-writer
description: Handles API documentation and README generation
tools: Read, Write, Edit
model: haiku
---
```

### Dependency Analyzer
```yaml
---
name: dependency-analyzer
description: Analyzes and reports on package dependencies
tools: Read, Bash, Grep
model: haiku
---
```

## Delegation Trigger Patterns

### Proactive Patterns:
- "Use proactively when [condition]"
- "Use proactively after [event]"
- "Use proactively for [task]"

### Specialist Patterns:
- "Specialist for [domain]"
- "Expert in [area]"
- "Dedicated to [task]"

### Handler Patterns:
- "Handles [responsibility]"
- "Manages [process]"
- "Responsible for [area]"

## Tool Selection Guidelines

### Minimal Read-Only:
```yaml
tools: Read, Grep, Glob
```

### Basic Modification:
```yaml
tools: Read, Edit, Write
```

### With Execution:
```yaml
tools: Read, Edit, Bash
```

### Full Analysis:
```yaml
tools: Read, Grep, Glob, WebFetch
```

## Model Selection Guide

### haiku (Fast, Simple)
- Simple analysis tasks
- Format conversions
- Basic validations
- Quick reports

### sonnet (Balanced)
- Code review
- Test generation
- Documentation
- Most general tasks

### opus (Powerful)
- Complex reasoning
- Security analysis
- Architecture design
- Critical decisions

## Remember for Sub-Agents
- One-shot execution (no interaction)
- Returns final result only
- Minimal tool set for efficiency
- Clear delegation trigger
- Specific output format
- No user dialogue
- Complete task autonomously