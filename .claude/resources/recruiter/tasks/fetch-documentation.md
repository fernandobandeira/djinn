# Fetch Documentation Task

## Purpose
Ensure Rita has the latest Claude Code documentation for accurate agent/sub-agent creation.

## Documentation URLs to Fetch

### For Commands
- `https://docs.anthropic.com/en/docs/claude-code/settings` - Settings and configuration
- `https://docs.anthropic.com/en/docs/claude-code/slash-commands` - Command syntax
- `https://docs.anthropic.com/en/docs/claude-code/settings#tools-available-to-claude` - Available tools

### For Sub-Agents
- `https://docs.anthropic.com/en/docs/claude-code/sub-agents` - Sub-agent feature
- `https://docs.anthropic.com/en/docs/claude-code/sub-agents#configuration` - Configuration syntax
- `https://docs.anthropic.com/en/docs/claude-code/sub-agents#best-practices` - Best practices

## Workflow

### Step 1: Determine Agent Type
Ask user:
1. **Command Agent** - Slash command like /analyst
2. **Sub-Agent** - Delegatable agent with frontmatter

### Step 2: Fetch Relevant Documentation
```bash
# For Commands
WebFetch https://docs.anthropic.com/en/docs/claude-code/slash-commands "Extract command syntax and special directives"

# For Sub-Agents
WebFetch https://docs.anthropic.com/en/docs/claude-code/sub-agents "Extract frontmatter format and delegation patterns"
```

### Step 3: Extract Key Syntax Rules

#### Command Syntax:
- `@` prefix for file loading
- `*` prefix for internal commands
- Resource file references
- Lazy loading patterns

#### Sub-Agent Syntax:
- Frontmatter format (name, description, tools, model)
- `IMPORTANT:` directive
- `Proactively` keyword for auto-delegation
- Tool selection guidelines

### Step 4: Store in Knowledge Base
```bash
# Index fetched documentation
./.vector_db/kb index --path .claude/resources/recruiter/data/claude-docs/
```

## Special Syntax Reference

### Command File Patterns:
```markdown
# DO NOT use @ in the command file itself
# Reference files like this:
resource_files:
  tasks:
    example: .claude/resources/agent/tasks/example.md

# Then load when needed:
When user requests `*command`:
1. THEN load: `.claude/resources/agent/tasks/example.md`
```

### Sub-Agent Patterns:
```markdown
---
name: agent-name
description: Use proactively when... | Specialist for...
tools: Tool1, Tool2, Tool3
model: haiku | sonnet | opus
---

# Purpose
IMPORTANT: [Critical behavior directive]

## Instructions
Proactively [action when condition met]
```

## Documentation Cache Strategy

1. Cache fetched docs for session
2. Refresh if older than 24 hours
3. Store key patterns in KB
4. Reference during creation

## Remember
- Always fetch docs BEFORE creating
- Understand current syntax
- Know special keywords
- Apply latest best practices