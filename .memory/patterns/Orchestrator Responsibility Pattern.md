---
title: Orchestrator Responsibility Pattern
type: note
permalink: patterns/orchestrator-responsibility-pattern
---

# Orchestrator Responsibility Pattern

## Problem

Sub-agents execute in isolated contexts and cannot access orchestrator configuration (like CLAUDE.md's Basic Memory project setting). If sub-agents write directly to storage, they may:
- Use wrong project parameters
- Create inconsistent note structures
- Bypass orchestrator's decision on what to save
- Duplicate storage with different formats

## Solution

**Sub-agents return synthesized content. Orchestrators handle all storage.**

### Sub-agent Responsibilities
- Execute heavy I/O (web search, data gathering)
- Synthesize and structure findings
- Return structured output with suggested storage metadata

### Orchestrator Responsibilities
- Decide what to save (may filter/transform sub-agent output)
- Apply project configuration from CLAUDE.md
- Control note formatting and [[wikilinks]]
- Write to Basic Memory with correct project parameter
- Link new notes to related content

## Implementation

### Sub-agent Output Schema
```yaml
result:
  synthesized_content: string    # Markdown content ready for storage
  suggested_title: string        # Recommended note title
  suggested_folder: string       # Recommended folder (research, diagrams, etc.)
  key_findings: [list]           # Structured summary
  relations: [list]              # Suggested [[wikilinks]]
  confidence: high|medium|low
```

### Sub-agent Tool Restrictions
- **DO NOT** include `Write` tool in sub-agent frontmatter
- **DO NOT** include MCP Basic Memory tools
- **DO** include only I/O tools needed for research: `WebSearch`, `WebFetch`, `Read`, `Grep`, `Glob`

### Orchestrator Flow
```
1. Orchestrator receives user request
2. Search Basic Memory for existing content
3. Delegate to sub-agent with context
4. Receive synthesized result
5. Review and optionally transform
6. Write to Basic Memory with project parameter
7. Return confirmation to user
```

## Example

### Bad (Sub-agent writes directly)
```yaml
# sub-agent
tools: Read, Write, WebSearch
---
# In process...
mcp__basic-memory__write_note(
    title="Research",
    folder="research"
    # Missing project parameter!
)
```

### Good (Orchestrator writes)
```yaml
# sub-agent
tools: Read, WebSearch
---
# Returns:
synthesized_content: |
  ## Findings
  [content]
suggested_title: "Market Research: Topic"
suggested_folder: "research"
```

```yaml
# orchestrator
# After receiving sub-agent result:
mcp__basic-memory__write_note(
    title=result.suggested_title,
    content=result.synthesized_content,
    folder=result.suggested_folder,
    project="djinn"  # From CLAUDE.md config
)
```

## Affected Components

### Sub-agents (must NOT write)
- [[market-researcher]] - returns research synthesis
- [[competitive-analyzer]] - returns analysis synthesis
- [[diagram-generator]] - returns diagram code/markdown
- [[knowledge-harvester]] - returns harvested content

### Orchestrators (MUST write)
- [[analyst]] - writes research, sessions, briefs
- [[architect]] - writes ADRs, patterns, diagrams

## Relations
- [[architecture]] - Djinn architecture overview
- [[agent-recruitment]] - agent creation skill
