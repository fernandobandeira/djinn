# Djinn

Agent architecture project with reusable skills and shared sub-agents.

## Prerequisites

**Required MCP**: Basic Memory must be installed and configured.

```bash
# Install Basic Memory
uv tool install basic-memory

# Add MCP to Claude Code
claude mcp add basic-memory -- uvx basic-memory mcp
```

## Project Initialization

For each project using Djinn, initialize Basic Memory:

```bash
cd your-project
basic-memory project add "$(basename $PWD)" ./.memory
basic-memory project default "$(basename $PWD)"
mkdir -p .memory/{decisions,patterns,research,context,sessions,diagrams}
```

## Knowledge Base System

All project knowledge is stored in `.memory/` using Basic Memory with [[wikilinks]].

### Memory-First Workflow
1. **Search first** before creating anything
2. Use [[wikilinks]] to connect related notes
3. Store decisions, patterns, research in appropriate folders

### Basic Memory Commands
```
# Search notes
mcp__basic-memory__search_notes(query="your search")

# Read a note
mcp__basic-memory__read_note(permalink="note-name")

# Write a note
mcp__basic-memory__write_note(
    title="Note Title",
    content="Content with [[links]]",
    folder="decisions"
)

# Recent activity
mcp__basic-memory__recent_activity()

# Project overview
mcp__basic-memory__canvas()
```

### Knowledge Structure
```
.memory/
├── project.md              # Vision, goals, overview
├── architecture.md         # System design
├── decisions/              # ADRs and key decisions
├── patterns/               # Documented patterns
├── research/               # Research outputs
├── context/                # Current state
├── sessions/               # Brainstorming sessions
└── diagrams/               # Technical diagrams
```

### Search Before Create
- Before creating an ADR: `search_notes("decision topic")`
- Before creating a pattern: `search_notes("pattern name")`
- Before brainstorming: `search_notes("topic ideas")`

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
- `systems-thinking` - Systems Mapping, Feedback Loops

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

## Agent Creation

Use the `agent-recruitment` skill to create new agents. Check `.claude/skills/agent-recruitment/decision-frameworks/reusability-assessment.md` to determine if something should be a skill, shared sub-agent, or agent-specific.
