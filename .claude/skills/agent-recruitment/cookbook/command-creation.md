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
.claude/orchestrators/{name}.md
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

### 2. Plan Architecture

Design the command structure:
```yaml
name: {name}
type: command
location: .claude/orchestrators/{name}.md
description: {clear description}
tools: {minimal set}
has_subcommands: {yes/no}
persona: {description if applicable}
resources: {list of resource files}
```

### 3. Build Files

Create using [templates/command-template.md](../templates/command-template.md):
- Apply specifications
- Create all necessary files
- Use relative paths for resources

### 4. Validate

Follow [cookbook/validation-workflow.md](./validation-workflow.md):
- Resource validation - All files exist?
- Constraint validation - Score 8.0-8.5?
- Coherence verification - Components integrate?

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
allowed-tools: Read, Grep, Glob  # For reviewers
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

## Skill Loading Pattern

Commands can load skills for comprehensive capabilities:

```markdown
## Skill Loading

This command loads the {skill-name} skill.

Read the skill:
@.claude/skills/{skill-name}/SKILL.md
```

This gives the command access to all skill cookbooks and resources.

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

## Example: Command with Skill

```markdown
---
description: Orchestrates agent creation workflows
allowed-tools: Read, Write, Grep, Glob, Bash, MultiEdit
---

# Rita - Recruiter

## Activation
Hello! I'm Rita, the Recruiter.
I create agents using best practices.

## Skill Loading

@.claude/skills/agent-recruitment/SKILL.md

## Commands
- `*help` - Available commands
- `*recruit {name}` - Full creation workflow
- `*plan {name}` - Plan architecture
- `*build` - Build from plan
- `*validate` - Run validations

## Core Principle
**Do Work Directly** - Skills do reasoning. Only spawn sub-agents for context isolation.
```

## Common Mistakes

### Too Many Tools
```
BAD: allowed-tools: Read, Write, Edit, Bash, Grep, Glob, Task, WebFetch
GOOD: allowed-tools: Read, Grep, Glob  (for reviewers)
```

### Missing Description
```
BAD: (no frontmatter)
GOOD:
---
description: Clear description of what this does
---
```

### Over-engineering
```
BAD: Command with 10 sub-agents for simple task
GOOD: Command does work directly, spawns sub-agents only for context isolation
```

## Checklist

Before finalizing a command:

- [ ] Clear, descriptive name
- [ ] Frontmatter with description
- [ ] Minimal tool set
- [ ] Sub-commands documented (if any)
- [ ] Resources load correctly
- [ ] Validates successfully
- [ ] Under 150 lines (load skills for more)
