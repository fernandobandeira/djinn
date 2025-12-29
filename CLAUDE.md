# Djinn

Agent framework for going from idea to code using orchestrator personas.

**Core principle:** Skills teach HOW to think. Sub-agents isolate context.

## Orchestrators

| Goal | Command | Orchestrator |
|------|---------|--------------|
| Validate an idea | `/analyst` | Ana |
| Design architecture | `/architect` | Archie |
| Research users | `/ux` | Ulysses |
| Plan product/epics | `/pm` | Paul |
| Create stories | `/sm` | Sam |
| Implement code | `/dev` | Dave |
| Create new agents | `/recruiter` | Rita |

**One orchestrator per chat.** Start a new conversation when switching orchestrators.

## The Workflow

Analyst → Architect/UX → PM → SM → Dev

1. **Discovery** - `/analyst` creates validated brief
2. **Context** - `/architect` + `/ux` add technical and user context
3. **Planning** - `/pm` synthesizes into epics
4. **Execution** - `/sm` creates stories, `/dev` implements

Status flows UP: Dev → SM → PM → Analyst (for pivots)

See `.memory/diagrams/Orchestrator Workflow.md` for full guide.

## Basic Memory Configuration

**Primary**: `djinn`

**Templates**: `templates/`

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

### Operations

Choose the right operation for the task:

| Operation | When to Use |
|-----------|-------------|
| `write_note` | Create new notes or completely rewrite existing ones |
| `edit_note` | Modify part of an existing note (append, prepend, find_replace, replace_section) |
| `move_note` | Rename a note or relocate it to a different folder |

**Prefer `edit_note` over delete + recreate** when updating existing content:

```python
# Append new section to existing note
mcp__basic-memory__edit_note(
    identifier="note-name",
    operation="append",
    content="\n\n## New Section\nAdditional content...",
    project="djinn"
)

# Replace a specific section
mcp__basic-memory__edit_note(
    identifier="note-name",
    operation="replace_section",
    section="## Section Name",
    content="## Section Name\nUpdated content...",
    project="djinn"
)

# Rename or move a note
mcp__basic-memory__move_note(
    identifier="old-name",
    destination_path="new-folder/new-name.md",
    project="djinn"
)
```

### Folders

| Content | Folder |
|---------|--------|
| Architecture decisions | `decisions/` |
| Reusable patterns | `patterns/` |
| Research & analysis | `research/` |
| Brainstorming sessions | `sessions/` |
| Technical diagrams | `diagrams/` |

## Skills

Skills teach HOW to think. They're organized in tiers. See `patterns/skill` for the pattern.

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

Delegate context-heavy work via Task tool. See `patterns/sub-agent` for the pattern.

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

## Extend the Framework

Use `/recruiter` to create new orchestrators, skills, or sub-agents. Rita guides you through the process.

## Documentation

Key docs in Basic Memory (`.memory/`):

| Doc | Content | Identifier |
|-----|---------|------------|
| Orchestrator Workflow | How to use the workflow | `diagrams/orchestrator-workflow` |
| Catalog | All orchestrators, skills, sub-agents | `reference/catalog` |
| Architecture | Design principles, extending | `architecture` |

### Patterns

| Pattern | What it defines | Identifier |
|---------|-----------------|------------|
| Orchestrator | Workflow personas that guide users | `patterns/orchestrator` |
| Skill | Thinking techniques that auto-activate | `patterns/skill` |
| Sub-agent | Context-isolated workers for heavy I/O | `patterns/sub-agent` |
| Templates | Platform-agnostic artifact structures | `patterns/templates` |
| Checklists | Workflow verification (embedded) | `patterns/checklists` |
| Memory | Docs-first philosophy | `patterns/memory` |

Read with: `mcp__basic-memory__read_note(identifier="...", project="djinn")`
