# Check Existing Agents Task

## Purpose
Efficiently check for existing agents before creating new ones using direct file navigation.

## Workflow

### Step 1: Get Overview
```bash
# See full structure
tree .claude/ -L 2

# Quick stats
echo "Commands: $(ls .claude/commands/*.md 2>/dev/null | wc -l)"
echo "Sub-agents: $(ls .claude/agents/*.md 2>/dev/null | wc -l)"
```

### Step 2: Check for Name Conflicts
```bash
# Check if name exists (case-insensitive)
ls .claude/commands/ | grep -i "[proposed-name]"
ls .claude/agents/ | grep -i "[proposed-name]"
```

### Step 3: Search for Similar Functionality
```bash
# Search by keywords
grep -r "[main-keyword]" .claude/commands/
grep -r "[main-keyword]" .claude/agents/

# Search by purpose
grep -r "brainstorm\|ideate\|generate" .claude/
grep -r "review\|audit\|check" .claude/
grep -r "test\|validate\|verify" .claude/
```

### Step 4: Examine Similar Agents
If similar agents found:
```bash
# Quick preview
head -30 [similar-agent-file]

# Check their tools (for sub-agents)
grep "tools:" [similar-agent-file]

# Check their commands (for command agents)
grep "^\*" [similar-agent-file]
```

### Step 5: Decision Matrix

| Finding | Action |
|---------|--------|
| Exact name match | Abort - name taken |
| Same functionality | Abort - duplicate |
| Similar but different | Proceed - differentiate clearly |
| No conflicts | Proceed - create new |

## Quick Check Commands

### For Command Agents
```bash
# List all command agents
ls -1 .claude/commands/*.md | xargs -I {} basename {} .md

# Show their purposes
grep -h "^## Activation" -A 1 .claude/commands/*.md
```

### For Sub-Agents
```bash
# List all sub-agents with descriptions
grep -h "^description:" .claude/agents/*.md

# Show their tool usage
grep -h "^tools:" .claude/agents/*.md | sort -u
```

## Common Agent Categories

### Analysis/Research
- analyst
- market-researcher
- competitor-analyzer

### Architecture/Design
- architect
- system-designer
- pattern-documenter

### Code Quality
- code-reviewer
- security-auditor
- performance-analyzer

### Testing
- test-writer
- test-runner
- coverage-reporter

### Documentation
- doc-writer
- api-documenter
- readme-generator

### Utility
- formatter
- dependency-updater
- config-validator

## Duplicate Detection Patterns

### Check These Combinations
```bash
# Semantic duplicates
grep -r "code.*review\|review.*code" .claude/
grep -r "test.*generat\|generat.*test" .claude/
grep -r "document.*api\|api.*document" .claude/

# Tool overlap
grep "tools:.*Read.*Write.*Edit" .claude/agents/*.md
```

## Reporting Results

### If No Conflicts
"✅ No existing agents with name '[name]' or similar functionality found."

### If Name Conflict
"❌ Agent named '[name]' already exists at [path]"

### If Functionality Overlap
"⚠️ Similar agent found: [name] at [path]
Purpose: [their-purpose]
Consider: 
1. Using existing agent
2. Extending existing agent
3. Creating with different focus"

## Remember
- Always check both commands and sub-agents
- Consider semantic similarity, not just names
- Look at tool usage patterns
- Review existing agent purposes
- Direct file access is faster than KB for this