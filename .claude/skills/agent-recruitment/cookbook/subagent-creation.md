# Creating Sub-agents

## When to Create a Sub-agent

Create a **sub-agent** when:
- Task needs isolated context window
- Work should run in parallel
- Results returned to calling agent
- Different tool permissions needed
- Proactive triggering desired (IMPORTANT keyword)
- Heavy computation that shouldn't pollute main context

## File Location

```
.claude/agents/{parent-command}/{name}.md    # Under a command
.claude/agents/shared/{name}.md              # Shared across commands
```

## Sub-agent Structure

```yaml
---
name: agent-name
description: IMPORTANT what this agent does (IMPORTANT for proactive)
tools: Tool1, Tool2, Tool3
model: haiku  # or sonnet, opus
---
```

```markdown
# Agent Purpose

[Clear statement of what this agent does]

## Instructions

[Step-by-step process]

## Output Format

[What to return to caller]

## Examples

[Sample inputs and outputs]
```

## Process

### 1. Determine Placement

```
Is this sub-agent specific to one command?
├── YES → .claude/agents/{command}/{name}.md
└── NO → .claude/agents/shared/{name}.md
```

### 2. Select Model

Use [decision-frameworks/model-selection.md](../decision-frameworks/model-selection.md):

| Task Type | Model |
|-----------|-------|
| Validators, formatters | haiku |
| Reviewers, analyzers | sonnet |
| Architects, planners | opus |

### 3. Plan with agent-planner

```
Create a plan for sub-agent '{name}'.
Purpose: {purpose}
Parent command: {command} or shared
Tools needed: {minimal set}
Model: {haiku/sonnet/opus}
Proactive: {yes/no - needs IMPORTANT?}

Return structured plan to Rita.
```

### 4. Build with agent-builder

```
Build sub-agent at .claude/agents/{parent}/{name}.md

Specifications:
- Name: {name}
- Description: {description}
- Tools: {tools}
- Model: {model}
- Returns: {output format}

Create the file.
```

### 5. Validate

Run validators:
1. `resource-validator` - File exists?
2. `constraint-validator` - Balance score?
3. `coherence-verifier` - Integrates with parent?

## Key Fields

### name
```yaml
name: code-reviewer  # Lowercase, hyphens
```

### description
```yaml
# Standard (called explicitly via Task tool)
description: Reviews code for quality issues

# Proactive (auto-triggers on relevant context)
description: IMPORTANT reviews code after significant changes
```

The `IMPORTANT` keyword enables proactive triggering.

### tools
```yaml
tools: Read, Grep, Glob  # Minimal set for the task
```

Only include tools the agent actually needs.

### model
```yaml
model: haiku   # Fast, cheap - validators, formatters
model: sonnet  # Balanced - reviewers, analyzers
model: opus    # Powerful - architects, complex reasoning
```

## Best Practices

### Minimal Tool Sets

```yaml
# Reviewer (read-only)
tools: Read, Grep, Glob

# Builder (needs to write)
tools: Read, Write, MultiEdit, Bash

# Analyzer (may need search)
tools: Read, Grep, Glob, WebSearch
```

### Clear Output Format

```markdown
## Output Format

Return to calling agent:
```yaml
status: success | failure
score: float (0-10)
issues:
  - severity: high | medium | low
    description: string
    location: string
recommendations:
  - string
```
```

### Single Responsibility

```
GOOD: constraint-validator - Only validates constraints
BAD: validator - Validates constraints AND resources AND coherence
```

### Proactive vs Explicit

```yaml
# Proactive - triggers automatically
description: IMPORTANT reviews backend code after changes

# Explicit - only when delegated
description: Analyzes architecture on request
```

## Example: Simple Validator (Haiku)

```yaml
---
name: resource-validator
description: IMPORTANT validates resource references and file existence
tools: Read, Grep, Glob, Bash
model: haiku
---
```

```markdown
# Resource Validator

Verify all referenced resources exist and are accessible.

## Process

1. Extract all resource references (@path, links)
2. Check each file exists
3. Verify load syntax is correct
4. Report missing or invalid references

## Output Format

```yaml
status: pass | fail
files_checked: int
missing:
  - path: string
    referenced_from: string
valid:
  - path: string
```

## Validation Rules

- `@path/file.md` must exist
- Relative links must resolve
- YAML frontmatter must parse
- Markdown links must have targets
```

## Example: Analyzer (Sonnet)

```yaml
---
name: code-reviewer
description: IMPORTANT reviews code for quality, security, and best practices
tools: Read, Grep, Glob
model: sonnet
---
```

```markdown
# Code Reviewer

Review code changes for quality issues.

## Review Checklist

1. **Security** - Injection, auth, data exposure
2. **Quality** - Clean code, DRY, SOLID
3. **Performance** - Efficiency, resource usage
4. **Testing** - Coverage, edge cases
5. **Documentation** - Comments, docs

## Process

1. Read changed files
2. Identify patterns and anti-patterns
3. Check against best practices
4. Generate actionable feedback

## Output Format

```yaml
overall_score: float (1-10)
issues:
  - severity: critical | high | medium | low
    category: security | quality | performance | testing | docs
    file: string
    line: int
    description: string
    suggestion: string
summary: string
```
```

## Example: Planner (Opus)

```yaml
---
name: agent-planner
description: IMPORTANT plans agent architecture with systematic decomposition
tools: Read, Grep, Glob, WebFetch, Task
model: opus
---
```

```markdown
# Agent Planner

Design agent architecture with proper decomposition.

## Planning Process

1. **Gather Requirements**
   - Purpose and goals
   - User interactions expected
   - Integration points

2. **Determine Type**
   - Command vs Skill vs Sub-agent
   - Apply type-selection framework

3. **Design Architecture**
   - Tool selection (minimal)
   - Model selection (appropriate)
   - Resource structure

4. **Plan Decomposition** (if needed)
   - Identify sub-agents
   - Define interfaces
   - Map delegation flow

## Output Format

```yaml
agent_type: command | skill | subagent
name: string
location: string
complexity: simple | moderate | complex
tools: [list]
model: haiku | sonnet | opus  # if subagent
sub_agents: [list]  # if orchestrator
resources: [list]
rationale: string
```

## Thinking Levels

- Simple agent: Standard analysis
- Complex agent: Think hard
- Orchestrator: Ultrathink about decomposition
```

## Integration with Orchestrators

Sub-agents are called via Task tool:

```markdown
## In Orchestrator (command/skill)

### Delegation Pattern

For planning:
```
Task(
  subagent_type="agent-planner",
  description="Plan agent architecture",
  prompt="Analyze requirements for {name}..."
)
```

For validation:
```
Task(
  subagent_type="constraint-validator",
  description="Validate constraints",
  prompt="Check constraint balance for..."
)
```
```

## Common Mistakes

### Too Many Tools
```yaml
BAD: tools: Read, Write, Edit, MultiEdit, Bash, Grep, Glob, Task, WebFetch
GOOD: tools: Read, Grep, Glob  # Only what's needed
```

### Wrong Model
```yaml
BAD: model: opus  # For simple validator
GOOD: model: haiku  # Validators are fast, rule-based
```

### Missing IMPORTANT
```yaml
BAD: description: reviews code  # Won't trigger proactively
GOOD: description: IMPORTANT reviews code  # Triggers on relevant context
```

### Vague Output Format
```
BAD: "Return the results"
GOOD: Structured YAML with specific fields
```

## Checklist

Before finalizing a sub-agent:

- [ ] Clear, descriptive name
- [ ] Description includes IMPORTANT (if proactive)
- [ ] Minimal tool set
- [ ] Appropriate model selection
- [ ] Clear output format documented
- [ ] Single responsibility
- [ ] Correct file location
- [ ] Validates with all three validators
