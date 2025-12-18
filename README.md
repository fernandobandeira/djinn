# Djinn

AI agent personas for Claude Code that help you brainstorm, research, analyze, and build.

## Prerequisites

### Required: Basic Memory MCP

Djinn uses [Basic Memory](https://basicmemory.io) for knowledge management. Install it before using Djinn:

```bash
# Install Basic Memory
uv tool install basic-memory

# Add MCP to Claude Code
claude mcp add basic-memory -- uvx basic-memory mcp

# Verify installation
claude mcp list
```

## Installation

### Option 1: Clone to your home directory (Recommended)

```bash
# Clone Djinn
git clone https://github.com/your-org/djinn.git ~/.djinn

# Copy to Claude Code config
cp -r ~/.djinn/.claude/* ~/.claude/
```

### Option 2: Use as project template

```bash
# Clone as starting point for a new project
git clone https://github.com/your-org/djinn.git my-project
cd my-project
```

## Per-Project Setup

Each project using Djinn needs Basic Memory initialized:

```bash
cd your-project

# Initialize Basic Memory for this project
basic-memory project add "$(basename $PWD)" ./.memory
basic-memory project default "$(basename $PWD)"

# Create folder structure
mkdir -p .memory/{decisions,patterns,research,context,sessions,diagrams}

# (Optional) Create initial project note
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

## Quick Start

```bash
# Type a slash command to activate an agent
/analyst       # Start brainstorming or research with Ana
/architect     # Design systems and create ADRs with Archie
/recruiter     # Create new agents with Rita
```

Once activated, type `*help` to see what the agent can do.

Skills auto-activate based on context - just mention "brainstorm", "root cause", "SWOT", etc.

## What would you like to do?

- **[Brainstorm & Research](#-ana---analyst)** - Market research, competitive analysis, ideation sessions
- **[Design Systems](#-archie---architect)** - System architecture, ADRs, technical diagrams
- **[Create New Agents](#-rita---recruiter)** - Build custom commands, skills, and sub-agents
- **[Use Thinking Skills](#skills)** - Root cause analysis, ideation, strategic analysis
- **[Understand the Architecture](#architecture)** - How agents are organized

## Meet Your Agents

### Ana - Analyst

Your strategic thinking partner for research and ideation.

**What Ana does:**
- Facilitates brainstorming sessions with proven techniques
- Conducts market research and competitive analysis
- Creates project briefs and documentation
- Applies structured thinking via skills (Six Hats, Five Whys, SWOT)

**Quick commands:**
```
*brainstorm {topic}     # Interactive ideation session
*research {topic}       # Market research document
*analyze-competition    # Competitive landscape analysis
*create-brief           # Project brief with guided questions
*swot                   # SWOT analysis (via strategic-analysis skill)
*five-whys              # Root cause analysis (via root-cause skill)
```

Ana works iteratively - she'll ask clarifying questions and offer numbered options at each step.

---

### Archie - Architect

Your system architecture partner for technical design decisions.

**What Archie does:**
- Designs system architectures with multiple options and trade-offs
- Creates Architecture Decision Records (ADRs)
- Reviews existing architectures for improvements
- Generates technical diagrams (Mermaid/PlantUML)
- Documents patterns, RFCs, and runbooks

**Quick commands:**
```
*design-system {scope}   # Multi-option architecture design
*create-adr {topic}      # Architecture Decision Record
*review-architecture     # Comprehensive architecture review
*diagram {type}          # Generate system/flow/component diagrams
*create-pattern {name}   # Document architectural pattern
```

Archie presents options with trade-offs and waits for your approval before proceeding.

---

### Rita - Recruiter

Creates new Claude Code agents using systematic decomposition.

**What Rita does:**
- Plans agent architecture with proper decomposition
- Builds commands, skills, and sub-agents
- Validates agents for coherence and balance
- Assesses reusability (shared vs agent-specific)

**Quick commands:**
```
*recruit {name}    # Full agent creation workflow
*plan {name}       # Design agent architecture
*validate          # Run all validation checks
```

Rita also auto-activates when you mention "create agent", "build agent", or "new command".

---

## Skills

Skills auto-activate based on context and provide thinking techniques any agent can use.

### Tier 1: Universal Skills
| Skill | Techniques | Trigger phrases |
|-------|------------|-----------------|
| `research` | KB Search, Harvesting, Evaluation | "search for", "find existing" |
| `root-cause` | Five Whys, First Principles, JTBD | "why", "root cause", "first principles" |
| `ideation` | SCAMPER, Walt Disney, Reverse Brainstorm | "brainstorm", "generate ideas" |
| `devils-advocate` | Red Team, Pre-mortem | "challenge this", "what could go wrong" |
| `role-playing` | Six Hats, Stakeholder, Persona | "perspectives", "what would X think" |
| `teacher` | Socratic, Feynman, Problem-Based | "explain", "teach me" |

### Tier 2: Domain Skills
| Skill | Techniques | Trigger phrases |
|-------|------------|-----------------|
| `strategic-analysis` | SWOT, Porter's Five Forces, Scenarios | "SWOT", "competitive forces" |
| `user-research` | Journey Mapping, Interviews, Surveys | "journey map", "user research" |
| `systems-thinking` | Systems Mapping, Feedback Loops, Tech Debt | "system dynamics", "feedback loop" |

---

## Knowledge Management

All project knowledge is stored in `.memory/` using Basic Memory:

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

**Key principle**: Always search before creating. Use [[wikilinks]] to connect notes.

---

## Architecture

```
                         SKILLS (auto-activate)
  research | root-cause | ideation | devils-advocate | role-playing | teacher
  strategic-analysis | user-research | systems-thinking
────────────────────────────────────────────────────────────────────────────
                              You
────────────────────────────────────────────────────────────────────────────
          ┌──────────────────────┼──────────────────────┐
          ▼                      ▼                      ▼
    ┌──────────┐          ┌───────────┐          ┌──────────┐
    │ /analyst │          │/architect │          │/recruiter│
    │   (Ana)  │          │ (Archie)  │          │  (Rita)  │
    └────┬─────┘          └─────┬─────┘          └────┬─────┘
         │                      │                     │
         └──────────────────────┼─────────────────────┘
                                │
                    ┌───────────┴───────────┐
                    ▼                       ▼
            SHARED SUB-AGENTS         AGENT-SPECIFIC
            (any agent can call)      SUB-AGENTS
            ├── market-researcher     (future)
            ├── competitive-analyzer
            ├── insight-synthesizer
            ├── documentation-gen
            ├── diagram-generator
            └── research-architect
```

**Key Principle:** Skills teach HOW to think. Sub-agents DO work.

**Agent Types:**
| Type | How it works | Example |
|------|--------------|---------|
| Command | You type `/name` to activate | `/analyst`, `/architect` |
| Skill | Auto-activates on context | "brainstorm" triggers ideation |
| Sub-agent | Called by parent agents | market-researcher, diagram-generator |

## Project Structure

```
.claude/
├── commands/           # /analyst, /architect, /recruiter
├── skills/             # Auto-activating thinking techniques
│   ├── research/       # KB search, harvesting, evaluation
│   ├── root-cause/     # Five Whys, First Principles, JTBD
│   ├── ideation/       # SCAMPER, Walt Disney, Reverse
│   ├── strategic-analysis/  # SWOT, Porter's, Scenarios
│   └── ...
├── agents/
│   └── shared/         # Sub-agents any command can call
└── resources/          # Templates, protocols, checklists
```

## Creating New Agents

Just ask! Say "create an agent" or use `/recruiter` and Rita will guide you through:

1. **Type selection** - Command vs Skill vs Sub-agent
2. **Reusability assessment** - Shared vs agent-specific
3. **Architecture planning** - Decomposition and tool selection
4. **Building** - File creation with proper structure
5. **Validation** - Coherence, constraints, and resource checks

---

## Git Integration

Add to your `.gitignore`:

```gitignore
# Basic Memory - keep everything (all Markdown)
# .memory/ is fully tracked

# Optional: ignore any generated caches
.memory/.cache/
```

Your entire project knowledge travels with the repo - clone and you have full context.

---

*Want more agents? Use `/recruiter` to build your own.*
