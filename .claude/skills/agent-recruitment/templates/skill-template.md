# Skill Template

Use this template when creating new skills.

## Folder Structure

```
.claude/skills/{skill-name}/
├── SKILL.md              # Required: main entry
├── cookbook/             # Optional: workflow recipes
│   ├── getting-started.md
│   └── advanced.md
├── templates/            # Optional: reusable templates
├── scripts/              # Optional: executable tools
└── reference/            # Optional: detailed docs
```

## SKILL.md Template

```yaml
---
name: {skill-name}
description: {What this skill does}. Use when {trigger phrases that
  should activate this skill}. Include keywords like {keyword1},
  {keyword2}, {keyword3}.
allowed-tools: {Tool1, Tool2}  # Optional: restrict tools
---
```

```markdown
# {Skill Title}

{One-line description of skill purpose.}

## Quick Start

{Minimal instructions to get started - keep under 20 lines.}

## When to Load What

{Route to appropriate cookbook files:}

### {Scenario A}
Read [cookbook/scenario-a.md](./cookbook/scenario-a.md)

### {Scenario B}
Read [cookbook/scenario-b.md](./cookbook/scenario-b.md)

## Key Concepts

{Essential information always needed - keep brief.}

### {Concept 1}
{Brief explanation}

### {Concept 2}
{Brief explanation}

## Tools Available

{If skill includes scripts:}
```bash
# Example usage
python .claude/skills/{skill-name}/scripts/tool.py input
```

## References

For detailed documentation:
- [Reference Guide](./reference/guide.md)
- [API Documentation](./reference/api.md)
```

## Variables to Replace

| Variable | Description | Example |
|----------|-------------|---------|
| `{skill-name}` | Lowercase, hyphenated | "pdf-processing" |
| `{Skill Title}` | Display name | "PDF Processing" |
| `{description}` | What + when | "Process PDF files..." |
| `{trigger phrases}` | Activation keywords | "PDF, forms, extract" |
| `{Tool1, Tool2}` | Allowed tools | "Read, Bash" |

## Description Best Practices

The description is CRITICAL for auto-discovery:

```yaml
# BAD: Too vague
description: Helps with documents

# BAD: No triggers
description: Processes PDF files

# GOOD: Specific with triggers
description: Process PDF documents including text extraction,
  form filling, and document merging. Use when user mentions
  PDF, forms, document extraction, or needs to work with
  PDF files.
```

## Minimal Skill Example

```yaml
---
name: go-testing
description: Go testing patterns and best practices. Use when
  writing Go tests, discussing test coverage, or troubleshooting
  test failures.
allowed-tools: Read, Bash
---
```

```markdown
# Go Testing Patterns

## Quick Start

Run tests:
```bash
go test ./... -v -cover
```

## Patterns

### Table-Driven Tests
[cookbook/table-driven.md](./cookbook/table-driven.md)

### Mocking with Mockery
[cookbook/mocking.md](./cookbook/mocking.md)
```

## Complex Skill Example

```yaml
---
name: api-design
description: Design REST APIs with OpenAPI specifications.
  Use when designing APIs, creating OpenAPI specs, or
  discussing API best practices, endpoints, schemas.
allowed-tools: Read, Write, Bash
---
```

```markdown
# API Design

Design REST APIs following best practices.

## Quick Start

1. Define resources and endpoints
2. Create OpenAPI spec
3. Generate code with oapi-codegen

## Workflows

### New API Design
[cookbook/new-api.md](./cookbook/new-api.md)

### Extend Existing API
[cookbook/extend-api.md](./cookbook/extend-api.md)

### Generate Code
[cookbook/code-generation.md](./cookbook/code-generation.md)

## Tools

Generate Go server:
```bash
oapi-codegen -generate types,chi-server -package api spec.yaml > api.gen.go
```

## References
- [OpenAPI Best Practices](./reference/openapi.md)
- [REST Design Guide](./reference/rest.md)
```

## Cookbook File Template

```markdown
# {Workflow Name}

## Overview
{What this workflow accomplishes}

## Prerequisites
{What's needed before starting}

## Steps

### Step 1: {Name}
{Instructions}

### Step 2: {Name}
{Instructions}

## Examples

{Concrete examples}

## Common Issues

{Troubleshooting tips}
```

## Checklist

After creating a skill:

- [ ] SKILL.md has proper frontmatter
- [ ] Description includes trigger phrases
- [ ] Quick Start is under 20 lines
- [ ] Cookbook files use relative links
- [ ] All referenced files exist
- [ ] Tool restrictions appropriate
- [ ] Scripts are executable (if any)
