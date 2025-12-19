# Djinn

A framework for building AI agent personas that help you brainstorm, research, analyze, and build.

## The Problem

AI assistance is powerful but unstructured. Most interactions are ad-hoc: prompt, respond, forget. No framework for sustained success.

- **No structure** - AI assistance is inconsistent without clear patterns
- **No standards** - Teams reinvent workflows instead of building on proven patterns
- **No methodology** - Proven thinking techniques (SCAMPER, Five Whys) aren't integrated
- **No memory** - Insights disappear; your 100th conversation isn't smarter than your first
- **No boundaries** - AI tries to be everything instead of focused expertise

## The Solution

A framework with three layers:

| Layer | Purpose | Examples |
|-------|---------|----------|
| **Orchestrators** | WHO - Role personas with boundaries | Analyst, Architect, Recruiter |
| **Skills** | HOW - Thinking techniques that auto-activate | SCAMPER, Five Whys, Six Hats |
| **Sub-agents** | ISOLATE - Heavy I/O without polluting context | market-researcher, knowledge-harvester |

Invoke an orchestrator and the AI becomes that persona - thinking with that role's perspective, using relevant skills, delegating research to sub-agents while keeping reasoning clean.

## Philosophy

**You stay in the loop.** Djinn isn't about vibe coding or blind delegation. AI accelerates your work - drafting research, applying frameworks, generating options - but you remain the decision maker.

- **AI generates, you review** - Research, plans, and decisions are drafts until you validate them
- **Shared knowledge base** - Memory serves both you and AI; keep it accurate
- **Structured collaboration** - Personas and skills make AI predictable, not autonomous

**Docs are the source of truth.** Your memory defines WHAT and WHY. Implementation is ephemeral and can change as long as it follows the docs.

---

## Quick Start: What Do You Want to Do?

| Goal | Command | Orchestrator |
|------|---------|--------------|
| Validate an idea | `/analyst` | Ana |
| Design architecture | `/architect` | Archie |
| Research users | `/ux` | Ulysses |
| Plan product/epics | `/pm` | Paul |
| Create stories | `/sm` | Sam |
| Implement code | `/dev` | Dave |
| Create new agents | `/recruiter` | Rita |

See **[Orchestrator Workflow](.memory/diagrams/Orchestrator%20Workflow.md)** for the full guide.

## Documentation

Full documentation is in `.memory/` (use [Obsidian](https://obsidian.md) to browse):

### Getting Started

| Doc | What It Covers |
|-----|----------------|
| **[Orchestrator Workflow](.memory/diagrams/Orchestrator%20Workflow.md)** | How to use the orchestrators (start here) |
| **[Catalog](.memory/reference/Catalog.md)** | All orchestrators, skills, and sub-agents |
| **[Architecture](.memory/Architecture.md)** | Design principles, extending the framework |

### Patterns

| Pattern | Description |
|---------|-------------|
| **[Orchestrator](.memory/patterns/Orchestrator.md)** | Workflow personas that guide users |
| **[Skill](.memory/patterns/Skill.md)** | Thinking techniques that auto-activate |
| **[Sub-agent](.memory/patterns/Sub-agent.md)** | Context-isolated workers for heavy I/O |

---

## Implementations

### Claude Code

Djinn is currently implemented for [Claude Code](https://claude.ai/code) (Anthropic's CLI).

#### Quick Start

```bash
# 1. Install Basic Memory (knowledge management)
uv tool install basic-memory
claude mcp add basic-memory -- uvx basic-memory mcp

# 2. Install Beads (work tracking)
go install github.com/steveyegge/beads/cmd/bd@latest

# 3. Clone Djinn
git clone https://github.com/your-org/djinn.git ~/.djinn
cp -r ~/.djinn/.claude/* ~/.claude/

# 4. Initialize your project
cd your-project
cp ~/.djinn/templates/CLAUDE.md ./CLAUDE.md  # Copy and customize
basic-memory project add "$(basename $PWD)" ./.memory
mkdir -p .memory/{decisions,patterns,research,context,sessions,diagrams}
bd init --quiet  # Initialize work tracking

# 5. Edit CLAUDE.md - set your project name:
# **Primary**: `your-project-name`
```

#### Usage

```
/analyst       # Ana - brainstorming, research, strategic analysis
/architect     # Archie - system design, ADRs, diagrams
/recruiter     # Rita - create new agents
```

Once activated, type `*help` to see available commands.

Skills auto-activate based on context - just mention "brainstorm", "root cause", "SWOT", etc.

#### Docs

- **[Claude Code Guide](.memory/reference/Claude%20Code%20Guide.md)** - Installation and usage
- **[Claude Code](.memory/decisions/implementations/Claude%20Code.md)** - Syntax, file structure, conventions
- **[Beads](.memory/decisions/implementations/Beads.md)** - Work tracking CLI

### Other Platforms

Want Djinn on another platform? The patterns are platform-agnostic.

- **Request an implementation** - [Open an issue](https://github.com/your-org/djinn/issues/new) describing your platform
- **Build it yourself** - PRs welcome! See `.memory/` for the conceptual patterns to implement

## License

MIT
