# Djinn

AI agent personas for Claude Code that help you brainstorm, research, analyze, and build.

## What would you like to do?

- **[Brainstorm & Research](#-ana---analyst)** - Market research, competitive analysis, ideation sessions
- **[Design Systems](#-archie---architect)** - System architecture, ADRs, technical diagrams
- **[Create New Agents](#-rita---recruiter)** - Build custom commands, skills, and sub-agents
- **[Use Thinking Skills](#skills)** - Root cause analysis, ideation, strategic analysis
- **[Understand the Architecture](#architecture)** - How agents are organized

## Quick Start

```bash
# Type a slash command to activate an agent
/analyst       # Start brainstorming or research with Ana
/architect     # Design systems and create ADRs with Archie
/recruiter     # Create new agents with Rita
```

Once activated, type `*help` to see what the agent can do.

Skills auto-activate based on context - just mention "brainstorm", "root cause", "SWOT", etc.

## Meet Your Agents

### 📊 Ana - Analyst

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

Archie presents options with trade-offs and waits for your approval before proceeding.

---

### 🎯 Rita - Recruiter

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

## Architecture

```
┌─────────────────────────────────────────────────────────────────────┐
│                         SKILLS (auto-activate)                       │
│  root-cause │ ideation │ devils-advocate │ role-playing │ teacher   │
│  strategic-analysis │ user-research │ systems-thinking              │
├─────────────────────────────────────────────────────────────────────┤
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
         │                      │                     │
         └──────────────────────┼─────────────────────┘
                                │
                    ┌───────────┴───────────┐
                    ▼                       ▼
            SHARED SUB-AGENTS         AGENT-SPECIFIC
            (any agent can call)      SUB-AGENTS
            ├── market-researcher     ├── system-designer
            ├── competitive-analyzer  ├── adr-manager
            ├── insight-synthesizer   ├── agent-planner
            ├── documentation-gen     └── agent-builder
            └── research-architect
```

**Key Principle:** Skills teach HOW to think. Sub-agents DO work.

**Agent Types:**
| Type | How it works | Example |
|------|--------------|---------|
| Command | You type `/name` to activate | `/analyst`, `/architect` |
| Skill | Auto-activates on context | "brainstorm" triggers ideation |
| Sub-agent | Called by parent agents | market-researcher, system-designer |

## Project Structure

```
.claude/
├── commands/           # /analyst, /architect, /recruiter
├── skills/             # Auto-activating thinking techniques
│   ├── root-cause/     # Five Whys, First Principles, JTBD
│   ├── ideation/       # SCAMPER, Walt Disney, Reverse
│   ├── strategic-analysis/  # SWOT, Porter's, Scenarios
│   └── ...
├── agents/
│   ├── shared/         # Sub-agents any command can call
│   ├── architect/      # Archie's specialists
│   └── recruiter/      # Rita's builders
└── resources/          # Templates, protocols, data
```

## Creating New Agents

Just ask! Say "create an agent" or use `/recruiter` and Rita will guide you through:

1. **Type selection** - Command vs Skill vs Sub-agent
2. **Reusability assessment** - Shared vs agent-specific
3. **Architecture planning** - Decomposition and tool selection
4. **Building** - File creation with proper structure
5. **Validation** - Coherence, constraints, and resource checks

---

*Want more agents? Use `/recruiter` to build your own.*
