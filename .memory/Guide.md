---
title: Guide
type: note
permalink: guide
tags:
- djinn
- guide
- installation
- extension
---

# Getting Started with Djinn

Complete guide to installing, using, and extending Djinn.

## Prerequisites

### Basic Memory MCP (Required)

Djinn uses [Basic Memory](https://github.com/basicmachines-co/basic-memory) for knowledge management. Install it first:

```bash
# Install Basic Memory
uv tool install basic-memory

# Add MCP to Claude Code
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

# Copy to Claude Code config
cp -r ~/.djinn/.claude/* ~/.claude/
```

### Option 2: Project Template

Use Djinn as a starting point for a new project:

```bash
# Clone as project template
git clone https://github.com/your-org/djinn.git my-project
cd my-project
```

## Per-Project Setup

Each project needs Basic Memory initialized:

```bash
cd your-project

# Initialize Basic Memory
basic-memory project add "$(basename $PWD)" ./.memory
basic-memory project default "$(basename $PWD)"

# Create folder structure
mkdir -p .memory/{decisions,patterns,research,context,sessions,diagrams}
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
| **Command** | You need a role persona with specific behavior and boundaries |
| **Skill** | You have a thinking technique multiple agents could use |
| **Sub-agent** | You need to isolate heavy I/O from the main conversation |

**Use `/recruiter` to create new components.** Rita knows the file formats, handles structure, and guides you through type selection. See [[architecture]] for design principles.

## Knowledge Management

Djinn stores project knowledge in `.memory/` using Basic Memory. This is your shared knowledge base - **review it regularly**.

| Content | Folder |
|---------|--------|
| Decisions (ADRs) | `.memory/decisions/` |
| Patterns | `.memory/patterns/` |
| Research | `.memory/research/` |
| Brainstorming | `.memory/sessions/` |
| Diagrams | `.memory/diagrams/` |

**You own this knowledge.** AI generates drafts, but you review and refine. Check research for accuracy, validate decisions, update patterns as you learn. The knowledge base serves both you and AI - keep it trustworthy.

**Tip:** Use [Obsidian](https://obsidian.md) to browse `.memory/`. The markdown files and `[[wikilinks]]` work perfectly with it.

## Relations

- [[project]] - What Djinn is and why
- [[architecture]] - Design principles and rules
