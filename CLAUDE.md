# Djinn

Agent architecture project with reusable skills and shared sub-agents.

## Basic Memory Configuration

**Primary**: `djinn`

## Templates Configuration

**Location**: `templates/`

> **CRITICAL: NEVER manually read, write, or edit files in `.memory/`**
>
> The `.memory/` directory is managed exclusively by Basic Memory MCP.
> - **ALWAYS** use `mcp__basic-memory__*` tools for ALL memory operations
> - **NEVER** use Read, Write, Edit, Bash, or any file system tool on `.memory/`
> - **NEVER** use `cat`, `grep`, `find`, or any shell command on `.memory/`
>
> Manual file operations bypass indexing, break semantic links, and corrupt the knowledge graph.

### Usage

Always include `project` parameter in Basic Memory MCP calls:

```
mcp__basic-memory__search_notes(query="topic", project="djinn")

mcp__basic-memory__write_note(
    title="Note Title",
    content="Content with [[wikilinks]]",
    folder="decisions",
    project="djinn"
)

mcp__basic-memory__read_note(identifier="note-name", project="djinn")

mcp__basic-memory__recent_activity(project="djinn")
```

### Workflow
1. **Search first** - Always check what exists before creating
2. **Link notes** - Use [[wikilinks]] to connect related content
3. **Include project** - Every MCP call needs `project="djinn"`

### Folders

| Content | Folder |
|---------|--------|
| Architecture decisions | `decisions/` |
| Reusable patterns | `patterns/` |
| Research & analysis | `research/` |
| Brainstorming sessions | `sessions/` |
| Technical diagrams | `diagrams/` |

## Skills Architecture

Skills teach HOW to think. They're organized in tiers:

### Tier 1 (Universal)
- `root-cause` - Five Whys, First Principles, Jobs-to-be-Done
- `ideation` - SCAMPER, Walt Disney Method, Reverse Brainstorming
- `devils-advocate` - Red Team/Blue Team, Pre-mortem Analysis
- `role-playing` - Six Thinking Hats, Stakeholder Roundtable
- `teacher` - Socratic Dialogue, Feynman Technique

### Tier 2 (Domain)
- `strategic-analysis` - SWOT, Porter's Five Forces, Scenario Planning
- `user-research` - Journey Mapping, Interview Design

## Shared Sub-Agents

Delegate context-heavy work via Task tool to these in `.claude/agents/shared/`:

### Research (Heavy I/O)
- `market-researcher` - Market research via web search
- `competitive-analyzer` - Competitive analysis via web search
- `knowledge-harvester` - Harvest external sources into Basic Memory

### Visualization
- `diagram-generator` - Mermaid/PlantUML technical diagrams

## Key Principles

1. **Skills teach HOW to think. Sub-agents isolate context.**
2. **Memory First**: Always search before creating
3. **Link Everything**: Use [[wikilinks]] to connect notes
4. **Brownfield**: Build on existing knowledge, never recreate
5. **You Write to KB**: Sub-agents return synthesis to you; you handle all KB writes

## After Delegating to Sub-agents

When you delegate to a sub-agent (market-researcher, competitive-analyzer, diagram-generator, knowledge-harvester), they return synthesis - they don't write to Basic Memory.

**You are responsible for writing their results to KB.**

Sub-agent results include:
- `synthesized_content` - Markdown ready for storage
- `suggested_title` - Recommended note title
- `suggested_folder` - Where to store (research, diagrams, etc.)

After receiving results, write to Basic Memory:
```python
mcp__basic-memory__write_note(
    title=result.suggested_title,
    content=result.synthesized_content,
    folder=result.suggested_folder,
    project="djinn"
)
```

## Agent Creation

Use the `agent-recruitment` skill to create new agents. Check `.claude/skills/agent-recruitment/decision-frameworks/reusability-assessment.md` to determine if something should be a skill, shared sub-agent, or agent-specific.
