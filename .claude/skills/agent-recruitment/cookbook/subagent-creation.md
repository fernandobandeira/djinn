# Creating Sub-agents

## When to Create a Sub-agent

Create a **sub-agent** ONLY for context isolation:
- Task needs isolated context window (heavy I/O that would flood main context)
- Work should run in parallel (independent tasks)
- Results summarized back to calling agent (process disposable, output matters)

## When NOT to Create a Sub-agent

Do NOT create sub-agents for:
- Reasoning work (sub-agents can't reason deeply)
- Anything requiring skill access (sub-agents can't call skills)
- Interactive work (sub-agents can't ask follow-up questions)
- Architecture decisions (need full context)
- Validation (do directly in skill/command)

## File Location

```
.claude/agents/shared/{name}.md    # Shared across commands (typical)
```

## Sub-agent Structure

```yaml
---
name: agent-name
description: IMPORTANT what this agent does (IMPORTANT for proactive)
tools: Tool1, Tool2, Tool3
model: haiku  # or sonnet (rarely opus)
---
```

```markdown
# Agent Purpose

[Clear statement of what this agent does]

## Instructions

[Step-by-step process]

## Output Format

[What to return to caller - summaries, not raw data]

## Examples

[Sample inputs and outputs]
```

## Process

### 1. Verify Context Isolation Need

Ask: "Is this truly for context isolation?"

```
Does this task:
├── Generate heavy I/O (research, many file reads)? → Sub-agent
├── Run independently in parallel? → Sub-agent
├── Need to summarize large amounts of data? → Sub-agent
└── Require reasoning, skills, or interaction? → Do directly
```

### 2. Select Model

| Task Type | Model |
|-----------|-------|
| Heavy I/O, formatting | haiku |
| Research, analysis | sonnet |
| Complex reasoning | DON'T use sub-agent |

### 3. Plan Architecture

Design the sub-agent:
```yaml
name: {name}
type: subagent
location: .claude/agents/shared/{name}.md
purpose: {context isolation purpose}
tools: {minimal set}
model: {haiku/sonnet}
output: {summary format}
```

### 4. Build File

Create using [templates/subagent-template.md](../templates/subagent-template.md):
- Clear purpose
- Minimal tools
- Appropriate model
- Structured output format

### 5. Validate

Follow [cookbook/validation-workflow.md](./validation-workflow.md):
- Resource validation - File exists?
- Constraint validation - Balance score?
- Coherence verification - Integrates correctly?

## Key Fields

### name
```yaml
name: market-researcher  # Lowercase, hyphens
```

### description
```yaml
# Standard (called explicitly via Task tool)
description: Researches market data and returns summary

# Proactive (auto-triggers on relevant context)
description: IMPORTANT researches market data when competitive analysis needed
```

The `IMPORTANT` keyword enables proactive triggering.

### tools
```yaml
tools: Read, Grep, Glob, WebFetch  # Minimal set for the task
```

Only include tools the agent actually needs.

### model
```yaml
model: haiku   # Fast, cheap - I/O heavy, formatting
model: sonnet  # Balanced - research, data processing
# Don't use opus for sub-agents - if you need opus reasoning, do it directly
```

## Best Practices

### Focus on Context Isolation

```
GOOD: Sub-agent that harvests 50 web pages and returns a summary
GOOD: Sub-agent that generates Mermaid diagrams in parallel

BAD: Sub-agent that validates constraints (do directly)
BAD: Sub-agent that plans architecture (needs reasoning)
```

### Minimal Tool Sets

```yaml
# Research sub-agent
tools: Read, Grep, WebFetch, WebSearch

# Documentation generator
tools: Read, Write, Glob

# Diagram generator
tools: Read, Write
```

### Clear Output Format

Sub-agents should return summaries, not raw data:

```markdown
## Output Format

Return to calling agent:
```yaml
status: success | failure
summary: string  # Key findings
data:
  - item: string
    relevance: string
recommendations:
  - string
```
```

### Single Responsibility

```
GOOD: market-researcher - Only researches market data
GOOD: diagram-generator - Only generates diagrams

BAD: research-and-analyze - Does too much
```

## Example: Research Sub-agent (Sonnet)

```yaml
---
name: market-researcher
description: IMPORTANT Researches market data and competitive landscape
tools: Read, WebSearch, WebFetch, Grep
model: sonnet
---
```

```markdown
# Market Researcher

Research market data and return summarized findings.

## Purpose

Context isolation for heavy research that would flood main context.

## Instructions

1. Receive research topic and focus areas
2. Search for relevant market data
3. Fetch and analyze key sources
4. Summarize findings concisely
5. Return structured summary

## Output Format

```yaml
status: success | partial | failure
topic: string
sources_consulted: int
key_findings:
  - finding: string
    source: string
    confidence: high | medium | low
market_size: string
competitors:
  - name: string
    position: string
trends:
  - string
recommendations:
  - string
```

## Constraints

- Maximum 10 web fetches
- Summarize, don't dump raw content
- Focus on actionable insights
```

## Example: Documentation Generator (Haiku)

```yaml
---
name: documentation-generator
description: IMPORTANT Generates structured documentation from analysis
tools: Read, Write, Glob
model: haiku
---
```

```markdown
# Documentation Generator

Generate formatted documentation from structured input.

## Purpose

Context isolation for document generation that shouldn't pollute main context.

## Instructions

1. Receive documentation requirements
2. Read relevant source files
3. Generate formatted documentation
4. Write to specified location
5. Return summary of generated files

## Output Format

```yaml
status: success | failure
files_generated:
  - path: string
    sections: int
    word_count: int
errors:
  - file: string
    error: string
```

## Supported Formats

- Markdown documentation
- API reference
- User guides
- Technical specifications
```

## Example: Diagram Generator (Haiku)

```yaml
---
name: diagram-generator
description: IMPORTANT generates technical diagrams in Mermaid/PlantUML format
tools: Read, Write, Glob
model: haiku
---
```

```markdown
# Diagram Generator

Generate technical diagrams from specifications.

## Purpose

Context isolation for diagram generation. Returns diagram code, not process.

## Instructions

1. Receive diagram requirements
2. Read relevant context files
3. Generate diagram in specified format
4. Write to file or return code
5. Return summary

## Output Format

```yaml
status: success | failure
diagram_type: flowchart | sequence | class | erd
format: mermaid | plantuml
code: string  # The diagram code
file_path: string | null  # If written to file
```

## Supported Diagrams

- Mermaid: flowchart, sequence, class, ER
- PlantUML: all diagram types
```

## Integration Pattern

Sub-agents are called via Task tool:

```markdown
## Calling a Sub-agent

For market research (heavy I/O):
```
Task(
  subagent_type="market-researcher",
  description="Research market for X",
  prompt="Research the market for {topic}. Focus on..."
)
```

For parallel diagram generation:
```
Task(
  subagent_type="diagram-generator",
  description="Generate architecture diagram",
  prompt="Create a Mermaid diagram showing..."
)
```
```

## Common Mistakes

### Using Sub-agents for Reasoning
```
BAD: Sub-agent for architecture planning
GOOD: Do architecture planning directly (needs reasoning + skills)
```

### Too Many Tools
```yaml
BAD: tools: Read, Write, Edit, MultiEdit, Bash, Grep, Glob, Task, WebFetch
GOOD: tools: Read, WebFetch  # Only what's needed for research
```

### Wrong Model
```yaml
BAD: model: opus  # Sub-agents don't need opus reasoning
GOOD: model: haiku or sonnet  # Appropriate for I/O tasks
```

### Returning Raw Data
```
BAD: Return all 50 pages of fetched content
GOOD: Return a 10-line summary of key findings
```

## Checklist

Before finalizing a sub-agent:

- [ ] Truly for context isolation (not reasoning)
- [ ] Clear, descriptive name
- [ ] Description includes IMPORTANT (if proactive)
- [ ] Minimal tool set
- [ ] Appropriate model (haiku/sonnet, not opus)
- [ ] Output format returns summaries
- [ ] Single responsibility
- [ ] Correct file location (agents/shared/)
- [ ] Validates successfully
