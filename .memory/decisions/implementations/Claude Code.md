---
title: Claude Code
type: note
permalink: decisions/implementations/claude-code
---

# Claude Code

Implementation of Djinn patterns on Claude Code (Anthropic's CLI).

## Context

Djinn patterns are platform-agnostic. This documents how we implement them specifically on Claude Code, including syntax, conventions, and rationale.

This document serves as:
- Reference for understanding Claude Code-specific syntax
- Guide for creating new agents correctly
- Documentation of implementation rationale
- Foundation for potential ports to other AI platforms

---

## Memory Layer

### Basic Memory MCP

**Choice:** Use Basic Memory MCP for knowledge management

**Rationale:**
- Semantic search and wikilinks
- Obsidian-compatible markdown for human review
- Project isolation via project parameter
- Tool-based access prevents manual file corruption
- Graph-based relationships between notes

**Configuration:** In CLAUDE.md, specify project name:
```markdown
**Primary**: `your-project-name`
**Shared**: `company-kb`  # Optional for team knowledge
```

**Why Specify Project Name:**
- Enables multi-project support (different KBs for different codebases)
- Sub-agents can't read CLAUDE.md, so orchestrators must apply config
- Explicit project prevents cross-contamination

**Access Pattern:** All memory operations through MCP tools:
- `mcp__basic-memory__search_notes` - Find existing content
- `mcp__basic-memory__read_note` - Read specific note
- `mcp__basic-memory__write_note` - Create new note
- `mcp__basic-memory__edit_note` - Modify existing note
- `mcp__basic-memory__build_context` - Build context from memory URI

**Critical Rule:** Never manually read/write/edit `.memory/` files. Manual edits bypass indexing and corrupt the knowledge graph.

---

## Working Memory Layer

Implements [[Working Memory]] pattern for persistent work tracking.

### Beads CLI

**Choice:** Use Beads CLI (`bd`) for operational work tracking

**Rationale:**
- Git-backed (`.beads/issues.jsonl`) - version controlled, merge-safe
- Agent-optimized with JSON output
- Dependency graph with ready-work calculation
- Hash-based IDs prevent merge conflicts
- No external service required

**Setup:**
```bash
# Install beads (one-time)
go install github.com/steveyegge/beads/cmd/bd@latest

# Initialize in project
bd init --quiet
```

**Concept Mapping:**

| Working Memory Concept | Beads Implementation |
|-----------------------|---------------------|
| Epic | `bd create -t epic` |
| Story/Feature | `bd create -t feature` |
| Task | `bd create -t task` |
| Bug | `bd create -t bug` |
| Parent-child | `--deps parent-child:{id}` |
| Blocks | `--deps blocks:{id}` |
| Discovered-from | `--deps discovered-from:{id}` |
| Ready work | `bd ready --json` |
| Sprint label | `bd label add {id} sprint-N` |
| Claim | `bd update {id} --status in_progress` |
| Complete | `bd close {id} --reason "..."` |

**Common Operations:**

```bash
# Create epic with stories
bd create "Epic: Auth System" -t epic -p 1 -d "Full description" --json
bd create "Story: Login UI" -t feature --deps parent-child:{epic-id} --json

# Plan sprint
bd ready --type feature --json           # Find ready stories
bd label add {story-id} sprint-1         # Assign to sprint
bd create "Task: Form" -t task --deps parent-child:{story-id} --json

# Implement
bd ready --limit 1 --json                # Get next task
bd update {id} --status in_progress      # Claim it
bd create "Found: Bug" -t bug --deps discovered-from:{id} --json
bd close {id} --reason "Implemented"     # Complete

# Query
bd blocked --json                        # What's stuck
bd dep tree {epic-id}                    # Hierarchy view
bd list --label sprint-1 --json          # Sprint items
```

**Graceful Degradation:** Working Memory is optional. If beads not initialized, orchestrators use TodoWrite for session-scoped tracking.

**Critical Rule:** Orchestrators reference [[Working Memory]] pattern concepts, not beads commands directly. This keeps orchestrators portable.

---

## Agent Layer

### Commands (Orchestrators)

**Location:** `.claude/commands/{name}.md`

**Invocation:** User types `/{name}` (e.g., `/analyst`, `/architect`)

**Session Rule:** One orchestrator per chat. Start a new conversation when switching orchestrators. Each maintains persona and context throughout the session - mixing them causes confusion. Memory persists across sessions via Basic Memory.

**Frontmatter:**
```yaml
---
description: Brief description shown in /commands list
allowed-tools: Tool1, Tool2, Tool3  # Optional restriction
---
```

**Key Sections:**
- `## Activation` - Greeting/persona introduction
- `## Commands` - Sub-commands with `*` prefix (e.g., `*help`, `*status`)
- `## Workflow` - Multi-phase process description
- `## Basic Memory Protocol` - How to handle KB operations

**Sub-command Convention:** Use `*command` syntax for modes within orchestrators. This creates multi-mode commands where users can switch between different actions.

### Skills (Thinking Techniques)

**Location:** `.claude/skills/{name}/SKILL.md` + `cookbook/*.md`

**Invocation:** Auto-activates on context match based on description keywords

**Frontmatter:**
```yaml
---
name: skill-name
description: What this skill does. Use when {trigger phrases}.
  Keywords: {keyword1}, {keyword2}, {keyword3}.
allowed-tools: Tool1, Tool2  # Optional
---
```

**Structure:**
```
.claude/skills/{name}/
├── SKILL.md              # Main entry, overview, routing
├── cookbook/             # Detailed technique guides
│   ├── technique1.md
│   └── technique2.md
├── templates/            # Optional: reusable templates
├── decision-frameworks/  # Optional: decision guides
└── reference/            # Optional: detailed docs
```

**Description is Critical:** The description field determines auto-discovery. Include trigger phrases and keywords that should activate this skill.

**Progressive Disclosure:** SKILL.md routes to cookbook files. Load only what's needed for the current task.

### Sub-agents (Context Isolation)

**Location:** `.claude/agents/shared/{name}.md`

**Invocation:** Via Task tool with `subagent_type` parameter

**Frontmatter:**
```yaml
---
name: agent-name
description: What this agent does. IMPORTANT if proactive trigger needed.
tools: Tool1, Tool2, Tool3  # Note: "tools" not "allowed-tools"
model: haiku | sonnet  # Required for sub-agents
---
```

**Key Differences from Commands/Skills:**
- Use `tools` (not `allowed-tools`)
- Must specify `model`
- Cannot call skills or reason deeply
- Return synthesis to orchestrator, don't write to KB

**Proactive Trigger:** Include `IMPORTANT` at start of description for auto-triggering based on context.

**Output Schema:** Sub-agents return structured synthesis:
```yaml
synthesized_content: |
  [markdown ready for storage]
suggested_title: "Note Title"
suggested_folder: "research"
key_findings: [list]
relations: [suggested wikilinks]
```

### Resources

**Location:**
- Shared: `.claude/resources/shared/{type}/{file}`
- Agent-specific: `.claude/resources/{agent}/{type}/{file}`

**Types:** `templates/`, `checklists/`

**Loading:** Use `@` syntax to load resources:
```markdown
@.claude/resources/analyst/templates/project-brief.md
```

---

## Model Selection

| Model | Speed | Cost | Use For |
|-------|-------|------|---------|
| **Haiku** | Fast | Low | Simple I/O, formatting, diagrams |
| **Sonnet** | Medium | Medium | Research, analysis, judgment needed |
| **Opus** | Slow | High | DON'T use for sub-agents - do reasoning directly |

**Key Insight:** If you need Opus-level reasoning, don't use a sub-agent. Do the work directly in the command/skill.

**Commands/Skills:** Inherit user's current model. No explicit selection needed.

**Sub-agents:** Must specify model in frontmatter.

**Decision Framework:**
```
What kind of task?
│
├─► Heavy I/O, formatting, simple output?
│   └─► HAIKU (cost-efficient)
│
├─► Research, data gathering, judgment needed?
│   └─► SONNET (balanced)
│
└─► Complex reasoning, architecture, decisions?
    └─► DON'T USE SUB-AGENT (do directly)
```

---

## Thinking Levels

Extended thinking budget via keywords in prompts:

| Level | Keywords | Tokens | Use When |
|-------|----------|--------|----------|
| **Standard** | (none) | ~2,000 | Simple, clear tasks |
| **Think** | "think" | 4,000 | 2-3 tools, some ambiguity |
| **Megathink** | "think hard", "think deeply" | 10,000 | 3-4 tools, multi-component |
| **Ultrathink** | "ultrathink", "think harder" | 32,000 | Orchestrators, architecture, decomposition |

**Always Ultrathink For:**
- Orchestrators with skill integration
- System decomposition decisions
- Architecture migrations
- Ambiguous requirements
- Writing to the knowledge base (docs are source of truth)

**Complexity Heuristics:**
| Factor | Standard | Think | Megathink | Ultrathink |
|--------|----------|-------|-----------|------------|
| Tools | 1-2 | 2-3 | 3-4 | 5+ |
| Skills needed | 0 | 0-1 | 1-2 | 3+ |
| Files | 1 | 2-3 | 4+ | Many |
| Clarity | Clear | Mostly | Some ambiguity | Ambiguous |

---

## Configuration Files

### CLAUDE.md (Project Instructions)

**Location:** Project root

**Purpose:**
- Automatically loaded by Claude Code at startup
- Contains project-specific instructions
- Configures memory project name(s)
- Lists available skills and sub-agents
- Documents project conventions

**Template:** We provide `templates/CLAUDE.md` for new projects.

**Why Template Exists:**
- Standardizes project setup across new projects
- Documents memory configuration pattern
- Includes critical access rules (never manual .memory edits)
- Easy copy-and-customize workflow
- Ensures consistent structure

**Template Usage:**
```bash
cp templates/CLAUDE.md ./CLAUDE.md
# Edit to set your project name and customize
```

### settings.json

**Location:** `.claude/settings.json`

**Purpose:** Permissions and tool allowlists

```json
{
  "permissions": {
    "allow": [
      "Skill(skill-name)",
      "Bash(command:*)"
    ]
  }
}
```

**Why Permissions:**
- Auto-approve certain tool uses without prompting
- Streamlines workflow for trusted operations
- Skills can be pre-approved for smoother activation

---

## Conventions

### File Naming
- Commands: lowercase, hyphenated (e.g., `code-reviewer.md`)
- Skills: lowercase, hyphenated folder (e.g., `root-cause/SKILL.md`)
- Sub-agents: lowercase, hyphenated (e.g., `market-researcher.md`)

### KB Write Responsibility
- Sub-agents return `synthesized_content`, `suggested_title`, `suggested_folder`
- Orchestrators receive synthesis and write to KB with proper project parameter
- This ensures consistent configuration application (sub-agents can't read CLAUDE.md)

### Progressive Disclosure
- SKILL.md contains overview and routing
- cookbook/ contains detailed technique guides
- Load only what's needed for the current task
- Reduces context usage, improves focus

### Wikilink Convention
Every note should include a `## Relations` section with wikilinks to related content.

---

## File Structure Summary

```
project-root/
├── CLAUDE.md                    # Project instructions (auto-loaded)
├── .memory/                     # Knowledge base (via MCP only!)
│   ├── decisions/               # Architecture decisions
│   ├── patterns/                # Reusable patterns
│   ├── research/                # Research & analysis
│   ├── sessions/                # Brainstorming sessions
│   └── diagrams/                # Technical diagrams
├── .claude/
│   ├── commands/                # Orchestrator personas
│   │   ├── analyst.md           # /analyst
│   │   ├── architect.md         # /architect
│   │   └── recruiter.md         # /recruiter
│   ├── skills/                  # Thinking techniques
│   │   └── {name}/
│   │       ├── SKILL.md
│   │       ├── cookbook/
│   │       ├── templates/
│   │       └── decision-frameworks/
│   ├── agents/
│   │   └── shared/              # Context-isolated workers
│   │       ├── market-researcher.md
│   │       ├── competitive-analyzer.md
│   │       └── knowledge-harvester.md
│   ├── resources/               # Templates, checklists
│   │   ├── shared/
│   │   └── {agent}/
│   │       ├── templates/
│   │       └── checklists/
│   └── settings.json            # Permissions
└── templates/
    └── CLAUDE.md                # Template for new projects
```

---

## Summary of Key Decisions

| Decision | Choice | Rationale |
|----------|--------|-----------|
| Memory backend | Basic Memory MCP | Semantic search, wikilinks, Obsidian-compatible |
| Project config | CLAUDE.md | Auto-loaded, project-specific settings |
| Orchestrators | Commands (`.claude/commands/`) | Explicit `/invocation`, multi-mode via `*` |
| Skills | SKILL.md + cookbook | Auto-discovery, progressive disclosure |
| Sub-agents | agents/shared/ | Context isolation, Task tool invocation |
| Model for sub-agents | Haiku/Sonnet only | Opus reasoning should be done directly |
| KB writes | Orchestrator responsibility | Sub-agents return synthesis |
| Template | templates/CLAUDE.md | Standardize new project setup |

## Relations

- [[Architecture]] - Platform-agnostic design patterns
- [[Memory]] - Docs-first philosophy and access rules
- [[Working Memory]] - Work tracking pattern (implemented via Beads)
- [[Beads]] - Working Memory implementation details
- [[Skill]] - Skill pattern details
- [[Sub-agent]] - Sub-agent pattern details
- [[Orchestrator]] - Orchestrator pattern details
- [[Project]] - Overall vision and goals