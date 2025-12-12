# Creating Commands

## When to Create a Command

Create a **command** when:
- User must explicitly invoke with `/command`
- Requires ongoing dialogue or state
- Has multiple modes/sub-commands (`*help`, `*status`)
- Entry point for a workflow
- Persona-based interaction (role-playing an expert)

## File Location

```
.claude/commands/{name}.md
```

## Command Structure

```markdown
---
description: Brief description shown in /commands list
allowed-tools: Tool1, Tool2  # Optional: restrict tools
---

# Agent Name

## Activation
[Greeting and persona introduction]

## Core Configuration (optional)
[YAML block with settings]

## Commands (if multi-mode)
- `*help` - Show available commands
- `*status` - Current state
- `*action` - Do something

## Interaction Protocol
[How to handle user requests]

## Resource Loading (optional)
@path/to/resource.md

## Examples
[Usage examples]
```

## Process

### 1. Gather Requirements

Before creating, clarify:
- What triggers this command? (`/name`)
- Does it need sub-commands?
- What tools does it need?
- Does it need resources?
- Should it have a persona?

### 2. Plan with agent-planner

Delegate to `agent-planner`:
```
Create a plan for command '{name}'.
Purpose: {purpose}
Expected interactions: {dialogue style}
Tools needed: {tool list}
Resources needed: {resource list}

Return structured plan to Rita.
```

### 3. Build with agent-builder

Provide specifications to `agent-builder`:
```
Build command at .claude/commands/{name}.md

Specifications:
- Name: {name}
- Description: {description}
- Tools: {tools}
- Has sub-commands: {yes/no}
- Persona: {persona description}
- Resources: {list}

Create all necessary files.
```

### 4. Validate

Run all three validators:
1. `resource-validator` - Files exist?
2. `constraint-validator` - Balance score?
3. `coherence-verifier` - Components integrate?

## Best Practices

### Keep Commands Focused
```
GOOD: /review-pr - Reviews pull requests
BAD: /code-tools - Does everything code-related
```

### Use Descriptive Names
```
GOOD: /backend-primer, /sprint-planner
BAD: /bp, /sp
```

### Minimal Tool Sets
```yaml
# Only include tools actually needed
allowed-tools: Read, Task  # For orchestrators
allowed-tools: Read, Write, Bash  # For implementers
```

### Progressive Resource Loading
```markdown
## Resource Loading

# Load only when needed:
IF planning phase:
  @resources/planning-guide.md
IF building phase:
  @resources/templates/
```

## Sub-Command Pattern

For commands with multiple modes:

```markdown
## Commands

### Primary Commands
- `*help` - Show this help
- `*start` - Begin workflow
- `*status` - Current progress

### Action Commands
- `*plan` - Create a plan
- `*build` - Execute plan
- `*validate` - Run checks

### Navigation
- `*back` - Previous step
- `*exit` - Exit command mode
```

## Orchestrator Pattern

Commands that coordinate sub-agents:

```markdown
## Orchestration

I delegate all work to specialized sub-agents:

### Planning Phase
Delegate to `agent-planner` for requirements analysis.

### Building Phase
Delegate to `agent-builder` for file creation.

### Validation Phase
Run validators in sequence:
1. `resource-validator`
2. `constraint-validator`
3. `coherence-verifier`

### Core Principle
**Orchestrate, Don't Execute** - Coordinate but don't implement.
```

## Example: Simple Command

```markdown
---
description: Run project tests with coverage report
allowed-tools: Bash, Read
---

# Test Runner

Run tests with coverage analysis.

## Usage

Simply invoke `/test` to run all tests.

Options:
- `*unit` - Unit tests only
- `*integration` - Integration tests only
- `*coverage` - Show coverage report

## Execution

1. Detect test framework (go test, pytest, jest)
2. Run appropriate command
3. Report results with summary
```

## Example: Orchestrator Command

```markdown
---
description: Scrum master workflow orchestrator
allowed-tools: Read, Task
---

# Scrum Master - Sarah

## Activation
Hello! I'm Sarah, your Scrum Master.
I coordinate sprint planning, tracking, and retrospectives.

## Commands
- `*help` - Available commands
- `*plan` - Sprint planning
- `*daily` - Daily standup format
- `*retro` - Retrospective facilitation

## Orchestration

All work delegated to sub-agents:
- `sprint-planner` - Planning ceremonies
- `story-creator` - Story writing
- `execution-tracker` - Progress tracking
- `retrospective-facilitator` - Retros

## Resource Loading
@resources/sm/sprint-templates.md
```

## Common Mistakes

### Too Many Tools
```
BAD: allowed-tools: Read, Write, Edit, Bash, Grep, Glob, Task, WebFetch
GOOD: allowed-tools: Read, Task  (for orchestrators)
```

### Missing Description
```
BAD: (no frontmatter)
GOOD:
---
description: Clear description of what this does
---
```

### Implementing Instead of Delegating
```
BAD: Command contains 500 lines of implementation
GOOD: Command delegates to sub-agents, stays under 100 lines
```

## Checklist

Before finalizing a command:

- [ ] Clear, descriptive name
- [ ] Frontmatter with description
- [ ] Minimal tool set
- [ ] Sub-commands documented (if any)
- [ ] Resources load correctly
- [ ] Validates with all three validators
- [ ] Under 150 lines (orchestrators can be slightly more)
