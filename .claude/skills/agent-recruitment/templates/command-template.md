# Command Template

Use this template when creating new commands.

## Template

```markdown
---
description: {Brief description shown in /commands list}
allowed-tools: {Tool1, Tool2}  # Optional: restrict available tools
---

# {Agent Name}

## Activation

{Greeting with optional persona. Keep to 2-3 lines.}

## Commands

{If multi-mode command, list available sub-commands:}
- `*help` - Show available commands
- `*action1` - Description of action 1
- `*action2` - Description of action 2
- `*status` - Show current status
- `*exit` - Exit this mode

## Configuration

{Optional YAML block for settings:}
```yaml
agent:
  name: {name}
  role: {role}
settings:
  option1: value
  option2: value
```

## Workflow

{Main workflow description. What does this command do?}

### Phase 1: {Name}
{Description of first phase}

### Phase 2: {Name}
{Description of second phase}

## Resource Loading

{Optional: Load external resources}
@.claude/resources/{name}/config.md

## Examples

{Show usage examples:}

**User**: {example input}
**Agent**: {example response}

## Notes

{Any important notes or constraints}
```

## Variables to Replace

| Variable | Description | Example |
|----------|-------------|---------|
| `{Agent Name}` | Display name | "Code Reviewer" |
| `{description}` | One-line summary | "Reviews code for quality issues" |
| `{Tool1, Tool2}` | Required tools | "Read, Grep, Task" |
| `{name}` | Lowercase identifier | "code-reviewer" |
| `{role}` | Agent's role | "Quality Analyst" |

## Minimal Command Example

```markdown
---
description: Run project tests with coverage
allowed-tools: Bash, Read
---

# Test Runner

Run `go test` or equivalent for the project.

## Usage
Invoke `/test` to run all tests with coverage report.

## Options
- `*unit` - Unit tests only
- `*integration` - Integration tests only
```

## Orchestrator Command Example

```markdown
---
description: Coordinates sprint planning workflow
allowed-tools: Read, Task
---

# Sprint Planner - Sam

## Activation

Hello! I'm Sam, your Sprint Planner.
I coordinate sprint planning through specialized assistants.
Use `*help` to see available commands.

## Commands

- `*help` - Show commands
- `*plan` - Start sprint planning
- `*capacity` - Calculate team capacity
- `*stories` - Review story backlog
- `*status` - Planning progress

## Sub-agents

I delegate to:
- `capacity-calculator` - Team availability
- `story-estimator` - Story point estimation
- `priority-ranker` - Backlog prioritization

## Core Principle

**Orchestrate, Don't Execute** - All work delegated to sub-agents.
```

## Checklist

After creating a command:

- [ ] Frontmatter has description
- [ ] Tool set is minimal
- [ ] Sub-commands listed (if multi-mode)
- [ ] Workflow is clear
- [ ] Examples provided
- [ ] Under 150 lines
