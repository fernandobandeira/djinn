---
title: Claude Code Guide
type: note
permalink: reference/claude-code-guide
tags:
- djinn
- claude-code
- installation
- guide
---

# Claude Code Guide

Guide to installing and using Djinn on Claude Code (Anthropic's CLI).

For the conceptual framework, see [[Architecture]]. For detailed syntax and conventions, see [[Claude Code]].

---

## Part 1: Core Setup (One-Time)

These steps configure your global environment. Do this once.

### 1.1 Install Basic Memory

[Basic Memory](https://github.com/basicmachines-co/basic-memory) provides knowledge management.

```bash
# Install Basic Memory
uv tool install basic-memory

# Verify installation
basic-memory --version
```

### 1.2 Install Beads CLI

[Beads](https://github.com/steveyegge/beads) provides work tracking (epics, stories, tasks).

```bash
# Install Beads
go install github.com/steveyegge/beads/cmd/bd@latest

# Verify installation
bd --version
```

### 1.3 Install Djinn Framework

Clone Djinn and copy to your global Claude config:

```bash
# Clone Djinn
git clone git@github.com:fernandobandeira/djinn.git ~/.djinn

# Copy implementation to global Claude config
cp -r ~/.djinn/.claude/* ~/.claude/
```

**Templates** stay at `~/.djinn/templates/` and are referenced by path.

---

## Part 2: Per-Project Setup

Run these steps for **every new project** using Djinn.

### 2.1 Add Basic Memory MCP

```bash
cd your-project

# Add MCP to this project
claude mcp add basic-memory -- uvx basic-memory mcp

# Verify
claude mcp list
```

### 2.2 Create Basic Memory Project

```bash
# Register project with Basic Memory
basic-memory project add "$(basename $PWD)" ./.memory

# Create folder structure
mkdir -p .memory/{decisions,patterns,research,context,sessions,diagrams}
```

### 2.3 Initialize Beads

```bash
# Initialize work tracking
bd init --quiet

# Remove generated files (orchestrators are our source of truth)
rm -f AGENTS.md @AGENTS.md
```

### 2.4 Copy CLAUDE.md Template

```bash
# Copy template
cp ~/.djinn/templates/CLAUDE.md ./CLAUDE.md

# Edit and set your project name
# Change: **Primary**: `your-project-name`
```

### 2.5 Optional: Create Project Note

```bash
cat > .memory/project.md << 'EOF'
---
title: Project
type: note
permalink: project
---

## Vision
[Your project vision here]

## Goals
- Goal 1
- Goal 2

## Relations
EOF
```

---

## Using Djinn

### Activating Agents

Type a slash command to activate an agent persona:

```
/analyst       # Ana - brainstorming, research, strategic analysis
/architect     # Archie - system design, ADRs, diagrams
/recruiter     # Rita - create new agents
```

Once activated, type `*help` to see available commands.

**One orchestrator per chat.** Start a new conversation when switching between orchestrators.

```
Chat 1: /analyst → explore idea, create brief → end session
Chat 2: /architect → review brief, create ADRs → end session
Chat 3: /pm → synthesize into epics → end session
```

Why? Each orchestrator maintains persona and context throughout the session. Mixing them in one chat causes confusion. Your work persists in memory across sessions - nothing is lost.

### Using Skills

Skills auto-activate based on context. Just mention the trigger:

| Say this... | Activates... | Techniques |
|-------------|--------------|------------|
| "brainstorm", "generate ideas" | ideation | SCAMPER, Walt Disney, Reverse |
| "why", "root cause" | root-cause | Five Whys, First Principles, JTBD |
| "challenge this", "what could go wrong" | devils-advocate | Pre-mortem, Red Team |
| "perspectives", "what would X think" | role-playing | Six Hats, Stakeholder |
| "SWOT", "competitive forces" | strategic-analysis | SWOT, Porter's Five Forces |

### Workflow Example

```
You: /analyst
Ana: I'm Ana, your Business Analyst. Type *help for commands...

You: *brainstorm new feature ideas
Ana: [Asks setup questions, then facilitates session using ideation skill]

You: Now let's do a SWOT analysis
[strategic-analysis skill auto-activates]
Ana: [Walks through SWOT framework]

You: *research the market for this feature
Ana: [Delegates to market-researcher sub-agent, returns summary]
```

---

## Extending Djinn

| Create a... | When... |
|-------------|---------|
| **Orchestrator** | You need a role persona with specific behavior and boundaries |
| **Skill** | You have a thinking technique multiple agents could use |
| **Sub-agent** | You need to isolate heavy I/O from the main conversation |

**Use `/recruiter` to create new components.** Rita knows the file formats, handles structure, and guides you through type selection. See [[Architecture]] for design principles.

---

## Knowledge Management

Your `.memory/` is a shared knowledge base that serves both you and AI.

| Content | Folder |
|---------|--------|
| Decisions (ADRs) | `.memory/decisions/` |
| Patterns | `.memory/patterns/` |
| Research | `.memory/research/` |
| Brainstorming | `.memory/sessions/` |
| Diagrams | `.memory/diagrams/` |

### Obsidian for Human Review

Use [Obsidian](https://obsidian.md) to browse and review `.memory/`:

- **Graph view** - See how notes connect via wikilinks
- **Search** - Full-text search across all knowledge
- **Edit** - Refine AI-generated drafts directly
- **Tags** - Filter by topic or type

The memory is designed to be human-readable. AI generates drafts; you review, validate, and refine. Your review is what makes memory trustworthy.

### Access Through Tools Only

**Critical:** Never manually edit `.memory/` files with code editors or shell commands. Always use memory tools (search, read, write, edit). Manual edits bypass indexing and break semantic links.

See [[Memory]] pattern for the full philosophy.

---

## Work Tracking

Working Memory provides persistent work tracking across sessions. See [[Working Memory]] pattern and [[Beads]] for CLI reference.

| Use Working Memory | Use TodoWrite |
|-------------------|---------------|
| Multi-session work | Session progress display |
| Epic → Story → Task hierarchy | Quick 1-3 item checklist |
| Blocking dependencies | Simple sequential work |
| Sprint planning | Immediate visibility |

### Orchestrator Integration

- **PM**: Creates epics with stories (`bd create -t epic`)
- **SM**: Tracks sprint tasks, organizes sprints (`bd ready`, `bd label`)
- **Dev**: Claims tasks, tracks discovered issues (`bd update`, `bd close`)

---

## Recommended Workflow

**Docs are the source of truth. Implementation is ephemeral.**

| Layer | Contains | Truth Status |
|-------|----------|--------------|
| **Memory** (`.memory/`) | Decisions, principles, patterns | **Source of truth** |
| **Implementation** (`.claude/`) | Code, prompts, configs | Derived from memory |

1. **Document first** - Before building, capture key decisions in memory
2. **Memory captures WHY** - Principles, rationale, criteria
3. **Implementation captures HOW** - Details derived from memory
4. **Human reviews memory** - Spend thought energy on docs, not implementation
5. **Refactor from memory** - When changing implementation, reference docs for guidance

See [[Memory]] pattern for the complete docs-first philosophy.

## Relations

- [[Architecture]] - Design principles and rules
- [[Memory]] - Docs-first philosophy
- [[Claude Code]] - Detailed syntax and conventions
- [[Working Memory]] - Work tracking pattern
- [[Beads]] - CLI reference