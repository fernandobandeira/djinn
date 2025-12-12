# Djinn

AI agent personas for Claude Code that help you brainstorm, research, analyze, and build.

## What would you like to do?

- **[Brainstorm & Research](#-ana---analyst)** - Market research, competitive analysis, ideation sessions
- **[Design Systems](#-archie---architect)** - System architecture, ADRs, technical diagrams
- **[Create New Agents](#-rita---recruiter)** - Build custom commands, skills, and sub-agents
- **[Understand the Architecture](#architecture)** - How agents are organized

## Quick Start

```bash
# Type a slash command to activate an agent
/analyst       # Start brainstorming or research with Ana
/architect     # Design systems and create ADRs with Archie
/recruiter     # Create new agents with Rita
```

Once activated, type `*help` to see what the agent can do.

## Meet Your Agents

### 📊 Ana - Analyst

Your strategic thinking partner for research and ideation.

**What Ana does:**
- Facilitates brainstorming sessions with 20+ proven techniques
- Conducts market research and competitive analysis
- Creates project briefs and documentation
- Applies structured thinking (Six Hats, Five Whys, SWOT)

**Quick commands:**
```
*brainstorm {topic}     # Interactive ideation session
*research {topic}       # Market research document
*analyze-competition    # Competitive landscape analysis
*create-brief           # Project brief with guided questions
```

Ana works iteratively - she'll ask clarifying questions and offer numbered options at each step.

---

### 🏗️ Archie - Architect

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

Archie presents options with trade-offs and waits for your approval before proceeding to the next phase.

---

### 🎯 Rita - Recruiter

Creates new Claude Code agents using systematic decomposition.

**What Rita does:**
- Plans agent architecture with proper decomposition
- Builds commands, skills, and sub-agents
- Validates agents for coherence and balance
- Extracts patterns from successful agents

**Quick commands:**
```
*recruit {name}    # Full agent creation workflow
*plan {name}       # Design agent architecture
*validate          # Run all validation checks
```

Rita also auto-activates when you mention "create agent", "build agent", or "new command".

---

## Architecture

```
┌─────────────────────────────────────────────────────────────────────┐
│                              You                                     │
└────────────────────────────────┬────────────────────────────────────┘
                                 │
          ┌──────────────────────┼──────────────────────┐
          ▼                      ▼                      ▼
    ┌──────────┐          ┌───────────┐          ┌──────────┐
    │ /analyst │          │/architect │          │/recruiter│
    │   (Ana)  │          │ (Archie)  │          │  (Rita)  │
    └────┬─────┘          └─────┬─────┘          └────┬─────┘
         │                      │                     │
    ┌────┴────┐            ┌────┴────┐           ┌────┴────┐
    ▼         ▼            ▼         ▼           ▼         ▼
 market   competitive   system    adr-       agent-    agent-
researcher  analyzer   designer  manager    planner   builder
    │         │            │         │           │         │
    └────┬────┘            └────┬────┘           └────┬────┘
         ▼                      ▼                     ▼
   Research Docs          ADRs & Diagrams       New Agent Files
```

**Agent Types:**
| Type | How it works | Example |
|------|--------------|---------|
| Command | You type `/name` to activate | `/analyst`, `/architect`, `/recruiter` |
| Skill | Auto-activates on context | "create an agent" triggers Rita |
| Sub-agent | Called by parent agents | system-designer, agent-planner |

## Project Structure

```
.claude/
├── commands/      # /analyst, /architect, /recruiter
├── skills/        # Auto-activating capabilities
├── agents/        # Sub-agents by parent
│   ├── analyst/   # Ana's specialists
│   ├── architect/ # Archie's designers
│   └── recruiter/ # Rita's builders
└── resources/     # Templates, protocols, data
```

## Creating New Agents

Just ask! Say "create an agent" or use `/recruiter` and Rita will guide you through:

1. **Type selection** - Command vs Skill vs Sub-agent
2. **Architecture planning** - Decomposition and tool selection
3. **Building** - File creation with proper structure
4. **Validation** - Coherence, constraints, and resource checks

---

*Want more agents? Use `/recruiter` to build your own.*
