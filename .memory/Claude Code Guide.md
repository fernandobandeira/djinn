---
title: Claude Code Guide
type: note
permalink: claude-code-guide
tags:
- djinn
- claude-code
- installation
- guide
---

# Claude Code Guide

Guide to installing and using Djinn on Claude Code (Anthropic's CLI).

For the conceptual framework, see [[Architecture]]. For detailed syntax and conventions, see [[Claude Code Implementation]].

## Prerequisites

### Basic Memory MCP (Required)

Djinn uses [Basic Memory](https://github.com/basicmachines-co/basic-memory) for knowledge management. Install it first:

```bash
# Install Basic Memory
uv tool install basic-memory

# Add MCP configuration
claude mcp add basic-memory -- uvx basic-memory mcp

# Verify installation
claude mcp list
```

## Installation

### Option 1: Clone to Home Directory (Recommended)

Use Djinn across all your projects:

```bash
# Clone Djinn
git clone https://github.com/your-org/djinn.git ~/.djinn

# Copy to framework config
cp -r ~/.djinn/.claude/* ~/.claude/
```

### Option 2: Project Template

Use Djinn as a starting point for a new project:

```bash
# Clone as project template
git clone https://github.com/your-org/djinn.git my-project
cd my-project
```

## Basic Memory Setup

### Step 1: Register Your Project(s)

Choose your setup based on your needs:

**Per-project KB** (most common - each project has its own knowledge):
```bash
cd your-project
basic-memory project add "$(basename $PWD)" ./.memory
mkdir -p .memory/{decisions,patterns,research,context,sessions,diagrams}
```

**Shared KB** (team knowledge across multiple projects):
```bash
basic-memory project add "company-kb" ~/Documents/shared-memory
mkdir -p ~/Documents/shared-memory/{decisions,patterns,research}
```

**Hybrid** (both project-specific and shared):
```bash
# Register per-project KB
basic-memory project add "myproject" ./.memory
mkdir -p .memory/{decisions,patterns,research,context,sessions,diagrams}

# Register shared KB (one-time team setup)
basic-memory project add "company-kb" ~/Documents/shared-memory
```

### Step 2: Create CLAUDE.md

Copy the template and customize:

```bash
cp templates/CLAUDE.md ./CLAUDE.md
# Edit CLAUDE.md and set your project name
```

**Per-project only:**
```markdown
**Primary**: `your-project-name`
```

**Hybrid setup (project + team KB):**
```markdown
**Primary**: `your-project-name`
**Shared**: `company-kb`
```

### Optional: Create Project Note

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

## Using Djinn

### Activating Agents

Type a slash command to activate an agent persona:

```
/analyst       # Ana - brainstorming, research, strategic analysis
/architect     # Archie - system design, ADRs, diagrams
/recruiter     # Rita - create new agents
```

Once activated, type `*help` to see available commands.

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

## Extending Djinn

| Create a... | When... |
|-------------|---------|
| **Orchestrator** | You need a role persona with specific behavior and boundaries |
| **Skill** | You have a thinking technique multiple agents could use |
| **Sub-agent** | You need to isolate heavy I/O from the main conversation |

**Use `/recruiter` to create new components.** Rita knows the file formats, handles structure, and guides you through type selection. See [[Architecture]] for design principles.

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

## Recommended Workflow

**Docs are the source of truth. Implementation is ephemeral.**

| Layer | Contains | Truth Status |
|-------|----------|--------------|
| **Memory** (`.memory/`) | Decisions, principles, patterns | **Source of truth** |
| **Implementation** (`.claude/`) | Code, prompts, configs | Derived from memory |

When using Djinn, follow this pattern:

1. **Document first** - Before building, capture key decisions in memory
2. **Memory captures WHY** - Principles, rationale, criteria
3. **Implementation captures HOW** - Details derived from memory
4. **Human reviews memory** - Spend thought energy on docs, not implementation
5. **Refactor from memory** - When changing implementation, reference docs for guidance

This means:
- Implementation can be regenerated if docs are clear
- New team members understand WHY by reading docs
- AI makes consistent decisions by following documented principles

See [[Memory]] pattern for the complete docs-first philosophy.

## Relations

- [[Project]] - What Djinn is and why
- [[Architecture]] - Design principles and rules
- [[Memory]] - Docs-first philosophy
- [[Claude Code Implementation]] - Detailed syntax and conventions
