# Creating Skills

## When to Create a Skill

Create a **skill** when:
- Should auto-activate when user context matches
- Requires multi-file resources (progressive disclosure)
- Domain expertise reusable across conversations
- Includes executable scripts or tools
- User shouldn't need to remember explicit commands

## File Location

```
.claude/skills/{skill-name}/
├── SKILL.md           # Required: entry point
├── cookbook/          # Optional: workflow recipes
├── templates/         # Optional: reusable templates
├── scripts/           # Optional: executable tools
└── reference/         # Optional: documentation
```

## SKILL.md Structure

```yaml
---
name: skill-name
description: What this skill does. Use when [trigger phrases].
  Include keywords users might say that should activate this skill.
allowed-tools: Tool1, Tool2  # Optional: restrict tools
---
```

```markdown
# Skill Title

Brief description of the skill's purpose.

## Quick Start
[Minimal instructions to get started]

## When to Load What

### Scenario A
Read [cookbook/scenario-a.md](./cookbook/scenario-a.md)

### Scenario B
Read [cookbook/scenario-b.md](./cookbook/scenario-b.md)

## Key Concepts
[Essential information always needed]

## References
For details: See [reference/](./reference/)
```

## Process

### 1. Identify Trigger Phrases

The `description` field is CRITICAL for auto-discovery:

```yaml
# BAD: Too vague
description: Helps with documents

# GOOD: Specific triggers
description: Process PDF documents including text extraction,
  form filling, and merging. Use when user mentions PDF, forms,
  document extraction, or needs to work with PDF files.
```

### 2. Design Progressive Disclosure

Map what loads when:

```
Tier 1 (Always): name + description in system prompt
Tier 2 (On match): SKILL.md body
Tier 3 (On demand): cookbook/, templates/, reference/
```

### 3. Plan Architecture

Design the skill structure:
```yaml
name: {skill-name}
type: skill
location: .claude/skills/{name}/SKILL.md
description: {with trigger phrases}
tools: {minimal set}
progressive_disclosure:
  tier1: name + description
  tier2: SKILL.md body
  tier3: cookbook files
cookbook_files: {list}
templates: {list}
```

### 4. Build Files

Create using [templates/skill-template.md](../templates/skill-template.md):
- SKILL.md with trigger phrases
- cookbook/ files for workflows
- templates/ for reusable patterns
- Proper cross-references

### 5. Validate

Follow [cookbook/validation-workflow.md](./validation-workflow.md):
- Resource validation - All files exist?
- Constraint validation - Balance appropriate?
- Coherence verification - Files reference correctly?

## Best Practices

### Specific Descriptions with Triggers

```yaml
# Include WHAT and WHEN
description: Create and manage Claude Code agents (commands, skills,
  sub-agents). Use when user mentions "create agent", "build agent",
  "recruit", "new command", "design skill", or discusses agent
  architecture and decomposition.
```

### Progressive Disclosure

```markdown
## Quick Start
[30 lines max - essential only]

## Detailed Workflows
For specific scenarios, load:
- [cookbook/scenario-a.md](./cookbook/scenario-a.md)
- [cookbook/scenario-b.md](./cookbook/scenario-b.md)

## Reference
For API details: [reference/api.md](./reference/api.md)
```

### Relative Links

Always use relative paths in skills:
```markdown
GOOD: [cookbook/file.md](./cookbook/file.md)
BAD: [file.md](/absolute/path/to/file.md)
```

### Tool Restrictions

```yaml
# Read-only skill
allowed-tools: Read, Grep, Glob

# Full capability skill
allowed-tools: Read, Write, Bash, Grep, Glob
```

## Cookbook Pattern

Organize workflows into separate files:

```
cookbook/
├── getting-started.md     # First-time setup
├── common-workflows.md    # Frequent operations
├── advanced-usage.md      # Complex scenarios
└── troubleshooting.md     # Problem solving
```

In SKILL.md:
```markdown
## Workflows

**New to this?**
Read [cookbook/getting-started.md](./cookbook/getting-started.md)

**Common tasks?**
Read [cookbook/common-workflows.md](./cookbook/common-workflows.md)

**Having issues?**
Read [cookbook/troubleshooting.md](./cookbook/troubleshooting.md)
```

## Scripts in Skills

Skills can include executable scripts:

```
scripts/
├── extract_text.py
├── validate_format.sh
└── transform_data.js
```

Reference in SKILL.md:
```markdown
## Tools

Extract text from PDF:
```bash
python .claude/skills/pdf-processing/scripts/extract_text.py input.pdf
```
```

## Example: Simple Skill

```yaml
---
name: go-testing
description: Go testing patterns and best practices. Use when writing
  Go tests, discussing test coverage, or troubleshooting test failures.
allowed-tools: Read, Grep, Bash
---
```

```markdown
# Go Testing Patterns

## Quick Start

Run tests:
```bash
go test ./... -v
```

With coverage:
```bash
go test ./... -cover -coverprofile=coverage.out
```

## Patterns

### Table-Driven Tests
[See cookbook/table-driven.md](./cookbook/table-driven.md)

### Mocking
[See cookbook/mocking.md](./cookbook/mocking.md)

### Integration Tests
[See cookbook/integration.md](./cookbook/integration.md)
```

## Example: Complex Skill with Scripts

```yaml
---
name: pdf-processing
description: Process PDF documents - extract text, fill forms, merge files.
  Use when working with PDF files, forms, or document extraction.
allowed-tools: Read, Write, Bash
---
```

```markdown
# PDF Processing

## Quick Start

Extract text:
```python
python .claude/skills/pdf-processing/scripts/extract.py doc.pdf
```

## Workflows

### Text Extraction
[cookbook/extraction.md](./cookbook/extraction.md)

### Form Filling
[cookbook/forms.md](./cookbook/forms.md)

### Merging Documents
[cookbook/merging.md](./cookbook/merging.md)

## Requirements
```bash
pip install pypdf pdfplumber
```
```

## Hybrid Pattern: Skill + Command

When you need BOTH auto-discovery AND explicit invocation:

```
.claude/
├── skills/
│   └── agent-recruitment/
│       └── SKILL.md           # Auto-discovers
└── commands/
    └── recruiter.md           # Explicit /recruiter
        # Contains: @.claude/skills/agent-recruitment/SKILL.md
```

The command loads the skill, giving users both options.

## Common Mistakes

### Vague Description
```
BAD: description: Helps with stuff
GOOD: description: [specific triggers and use cases]
```

### Everything in SKILL.md
```
BAD: 500-line SKILL.md with all information
GOOD: 50-line SKILL.md with links to cookbook/
```

### Absolute Paths
```
BAD: Read /home/user/.claude/skills/x/file.md
GOOD: Read [./cookbook/file.md](./cookbook/file.md)
```

### No allowed-tools
```
BAD: (no tool restriction - skill has all tools)
GOOD: allowed-tools: Read, Grep  # Minimal set
```

## Checklist

Before finalizing a skill:

- [ ] Descriptive name (lowercase, hyphens)
- [ ] Description with trigger phrases
- [ ] SKILL.md under 100 lines
- [ ] Progressive disclosure structure
- [ ] Relative links work correctly
- [ ] Tool restrictions appropriate
- [ ] All referenced files exist
- [ ] Scripts are executable (if any)
- [ ] Validates successfully
