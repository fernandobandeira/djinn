# Sub-agent Template

Use this template when creating new sub-agents.

## Template

```yaml
---
name: {agent-name}
description: {IMPORTANT if proactive} {what this agent does}
tools: {Tool1, Tool2, Tool3}
model: {haiku | sonnet | opus}
---
```

```markdown
# {Agent Purpose}

{One-line description of what this agent does.}

## Instructions

{Step-by-step process this agent follows:}

1. {Step 1}
2. {Step 2}
3. {Step 3}

## Input

{What this agent receives from caller:}

```yaml
agent_name: string
path: string
context: object  # Optional additional context
```

## Output Format

{Structured output returned to caller:}

```yaml
status: success | failure
result:
  field1: type
  field2: type
issues:
  - severity: critical | high | medium | low
    description: string
    location: string
recommendations:
  - string
```

## Rules

{Specific rules this agent follows:}

- Rule 1
- Rule 2
- Rule 3

## Basic Memory Protocol

**Sub-agents do NOT write to Basic Memory.** Return structured data to the orchestrator.

The orchestrator handles all KB operations:
- Decides what to save
- Applies project configuration from CLAUDE.md
- Controls formatting and linking
- Writes to Basic Memory with proper project parameter

If the sub-agent needs to suggest content for storage, include it in the output:
```yaml
suggested_content:
  title: "Note Title"
  folder: "research"
  content: "Markdown content..."
```

## Examples

{Example inputs and outputs:}

**Input**:
```yaml
agent_name: "test-validator"
path: ".claude/agents/test-validator.md"
```

**Output**:
```yaml
status: success
result:
  score: 8.2
issues: []
recommendations:
  - "Consider reducing tool count"
```
```

## Variables to Replace

| Variable | Description | Example |
|----------|-------------|---------|
| `{agent-name}` | Lowercase, hyphenated | "code-reviewer" |
| `{Agent Purpose}` | Clear purpose statement | "Code Quality Reviewer" |
| `{IMPORTANT}` | Include for proactive trigger | "IMPORTANT reviews..." |
| `{Tool1, Tool2}` | Minimal tool set | "Read, Grep, Glob" |
| `{haiku\|sonnet\|opus}` | Model selection | "sonnet" |

## Model Selection Guide

| Task Type | Model | Rationale |
|-----------|-------|-----------|
| Validators | haiku | Fast, rule-based |
| Formatters | haiku | Mechanical |
| Reviewers | sonnet | Judgment needed |
| Analyzers | sonnet | Pattern matching |
| Planners | opus | Complex reasoning |
| Architects | opus | Strategic decisions |

## Proactive vs Explicit

```yaml
# Proactive - triggers automatically on relevant context
description: IMPORTANT reviews code after changes

# Explicit - only called via Task tool
description: Analyzes architecture when requested
```

## Validator Example (Haiku)

```yaml
---
name: syntax-validator
description: IMPORTANT validates syntax in configuration files
tools: Read, Bash
model: haiku
---
```

```markdown
# Syntax Validator

Validate syntax in YAML, JSON, and TOML files.

## Instructions

1. Identify file type from extension
2. Run appropriate parser
3. Report syntax errors with line numbers

## Output Format

```yaml
status: pass | fail
files_checked: int
errors:
  - file: string
    line: int
    message: string
```

## Rules

- Check all files in scope
- Report exact line numbers
- Provide fix suggestions when possible
```

## Reviewer Example (Sonnet)

```yaml
---
name: security-reviewer
description: IMPORTANT reviews code for security vulnerabilities
tools: Read, Grep, Glob
model: sonnet
---
```

```markdown
# Security Reviewer

Review code for security vulnerabilities.

## Instructions

1. Scan for common vulnerability patterns
2. Check authentication/authorization
3. Identify data exposure risks
4. Review input validation

## Checklist

- [ ] SQL injection
- [ ] XSS vulnerabilities
- [ ] CSRF protection
- [ ] Auth bypass
- [ ] Sensitive data exposure
- [ ] Input validation

## Output Format

```yaml
overall_risk: critical | high | medium | low | none
vulnerabilities:
  - severity: critical | high | medium | low
    type: string
    file: string
    line: int
    description: string
    remediation: string
summary: string
```
```

## Planner Example (Opus)

```yaml
---
name: architecture-planner
description: IMPORTANT plans system architecture for complex features
tools: Read, Grep, Glob, WebFetch
model: opus
---
```

```markdown
# Architecture Planner

Design system architecture for complex features.

## Instructions

1. Understand requirements thoroughly
2. Research existing patterns
3. Evaluate architectural options
4. Recommend optimal approach
5. Document decision rationale

## Thinking Level

For complex architecture decisions, ultrathink about:
- Trade-offs between approaches
- Long-term maintainability
- Integration with existing systems

## Output Format

```yaml
recommendation:
  approach: string
  rationale: string
options_considered:
  - name: string
    pros: [string]
    cons: [string]
components:
  - name: string
    responsibility: string
    interfaces: [string]
risks:
  - description: string
    mitigation: string
```
```

## Checklist

After creating a sub-agent:

- [ ] Name is lowercase, hyphenated
- [ ] Description is clear (with IMPORTANT if proactive)
- [ ] Tool set is minimal (no Write for KB purposes)
- [ ] Model is appropriate for task
- [ ] Output format is structured
- [ ] Returns data to orchestrator (no direct KB writes)
- [ ] Rules/instructions are clear
- [ ] Examples provided
- [ ] Under 100 lines
