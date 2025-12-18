# Project Name

<!--
  Copy this to your project root as CLAUDE.md and customize.

  Setup:
  1. basic-memory project add "your-project" ./.memory
  2. mkdir -p .memory/{decisions,patterns,research,sessions,diagrams}
  3. Update **Primary** below with your project name
-->

## Basic Memory Configuration

**Primary**: `YOUR_PROJECT_NAME`

<!-- Optional: shared KB for team knowledge -->
<!-- **Shared**: `company-kb` -->

### Usage

Always include `project` parameter in Basic Memory MCP calls:

```
mcp__basic-memory__search_notes(query="topic", project="YOUR_PROJECT_NAME")

mcp__basic-memory__write_note(
    title="Note Title",
    content="Content with [[wikilinks]]",
    folder="decisions",
    project="YOUR_PROJECT_NAME"
)

mcp__basic-memory__read_note(identifier="note-name", project="YOUR_PROJECT_NAME")

mcp__basic-memory__recent_activity(project="YOUR_PROJECT_NAME")

mcp__basic-memory__build_context(url="memory://topic", project="YOUR_PROJECT_NAME")
```

### Workflow

1. **Search first** - Always check what exists before creating
2. **Link notes** - Use [[wikilinks]] to connect related content
3. **Include project** - Every MCP call needs the project parameter

### Folders

| Content | Folder |
|---------|--------|
| Architecture decisions | `decisions/` |
| Reusable patterns | `patterns/` |
| Research & analysis | `research/` |
| Brainstorming sessions | `sessions/` |
| Technical diagrams | `diagrams/` |

### Wikilinks

Every note should include a `## Relations` section:

```markdown
## Relations
- [[project]] - project context
- [[related-note]] - related content
```

<!-- Add your project-specific instructions below -->
