# Direct Navigation Patterns for Agent Discovery

## Quick Agent Discovery Commands

### See Everything at Once
```bash
# Full agent structure
tree .claude/ -L 3

# Compact view
tree .claude/ -L 2 -d  # directories only
```

### List Specific Agent Types
```bash
# Command agents
ls -la .claude/commands/*.md

# Sub-agents
ls -la .claude/agents/*.md

# Agent resources
ls -la .claude/resources/
```

### Search for Patterns
```bash
# Find agents with specific keywords
grep -r "brainstorm" .claude/commands/
grep -r "review" .claude/agents/

# Find agents by capability
grep -r "tools:" .claude/agents/  # See tool usage
grep -r "*help" .claude/commands/  # Find command patterns

# Find specific agent types
grep -l "description: Use proactively" .claude/agents/*.md
grep -l "interactive" .claude/commands/*.md
```

### Quick Agent Preview
```bash
# See agent headers
head -20 .claude/commands/*.md
head -10 .claude/agents/*.md

# See frontmatter only
sed -n '/^---$/,/^---$/p' .claude/agents/*.md

# See agent names
grep "^name:" .claude/agents/*.md
grep "^# .* Agent" .claude/commands/*.md
```

### Check for Duplicates
```bash
# Check if agent name exists
ls .claude/commands/ | grep -i "analyst"
ls .claude/agents/ | grep -i "reviewer"

# Check for similar functionality
grep -r "security" .claude/
grep -r "test" .claude/
```

### Analyze Agent Structure
```bash
# Count agents
ls .claude/commands/*.md 2>/dev/null | wc -l
ls .claude/agents/*.md 2>/dev/null | wc -l

# See resource organization
tree .claude/resources/ -L 2

# Find largest agents
wc -l .claude/commands/*.md | sort -n
wc -l .claude/agents/*.md | sort -n
```

## Navigation Workflow

### 1. Initial Discovery
```bash
# Start with tree for overview
tree .claude/ -L 2

# Output example:
.claude/
├── agents/
│   ├── code-reviewer.md
│   └── test-writer.md
├── commands/
│   ├── analyst.md
│   ├── architect.md
│   └── recruiter.md
└── resources/
    ├── analyst/
    ├── architect/
    └── recruiter/
```

### 2. Check Specific Type
```bash
# If creating a command
ls .claude/commands/

# If creating a sub-agent
ls .claude/agents/
```

### 3. Search for Similar
```bash
# Search by keyword
grep -l "[keyword]" .claude/*/*.md

# Search by pattern
grep -r "tools:.*Read.*Write" .claude/agents/
```

### 4. Read Similar Agent
```bash
# Quick preview
head -50 .claude/commands/analyst.md

# Full read
cat .claude/agents/code-reviewer.md
```

## Common Patterns to Look For

### In Commands
- Activation section
- Commands with * prefix
- Resource file references
- Interactive protocols
- Task execution patterns

### In Sub-Agents
- Frontmatter with name, description, tools
- IMPORTANT directives
- One-shot execution pattern
- Output format specification
- No interaction loops

## Efficient Navigation Tips

### Use Glob Patterns
```bash
# All markdown files
ls .claude/**/*.md

# All task files
ls .claude/resources/*/tasks/*.md

# All templates
ls .claude/resources/*/templates/*.md
```

### Use Find for Complex Searches
```bash
# Find recent agents
find .claude/ -name "*.md" -mtime -7

# Find by size
find .claude/ -name "*.md" -size +100

# Find and preview
find .claude/ -name "*.md" -exec head -5 {} \;
```

### Create Aliases (Optional)
```bash
# Quick commands for common operations
alias lscmds='ls -la .claude/commands/'
alias lsagents='ls -la .claude/agents/'
alias treeagents='tree .claude/ -L 2'
```

## Remember
- Direct navigation is faster than KB for agents
- Use tree for structure overview
- Use ls for specific directories
- Use grep for pattern searching
- Use head/cat for quick reading
- Combine commands with pipes for power