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

## Documentation

Full documentation is in `.memory/` (use [Obsidian](https://obsidian.md) to browse - they use `[[wikilinks]]`):

| Doc | Content |
|-----|---------|
| **[Project](.memory/Project.md)** | Vision, goals, philosophy |
| **[Architecture](.memory/Architecture.md)** | Core concepts, design rules |

### Patterns

| Pattern | Description |
|---------|-------------|
| **[Orchestrator](.memory/patterns/Orchestrator.md)** | Workflow personas that guide users |
| **[Skill](.memory/patterns/Skill.md)** | Thinking techniques that auto-activate |
| **[Sub-agent](.memory/patterns/Sub-agent.md)** | Context-isolated workers for heavy I/O |
| **[Memory](.memory/patterns/Memory.md)** | Docs-first knowledge management |

---

## Implementations

### Claude Code

Djinn is currently implemented for [Claude Code](https://claude.ai/code) (Anthropic's CLI).

#### Quick Start

```bash
# 1. Install Basic Memory
uv tool install basic-memory
claude mcp add basic-memory -- uvx basic-memory mcp

# 2. Clone Djinn
git clone https://github.com/your-org/djinn.git ~/.djinn
cp -r ~/.djinn/.claude/* ~/.claude/

# 3. Initialize your project
cd your-project
cp ~/.djinn/templates/CLAUDE.md ./CLAUDE.md  # Copy and customize
basic-memory project add "$(basename $PWD)" ./.memory
mkdir -p .memory/{decisions,patterns,research,context,sessions,diagrams}

# 4. Edit CLAUDE.md - set your project name:
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

- **[Claude Code Guide](.memory/Claude%20Code%20Guide.md)** - Installation and usage
- **[Claude Code Implementation](.memory/Claude%20Code%20Implementation.md)** - Syntax, file structure, conventions

### Other Platforms

Want Djinn on another platform? The patterns are platform-agnostic.

- **Request an implementation** - [Open an issue](https://github.com/your-org/djinn/issues/new) describing your platform
- **Build it yourself** - PRs welcome! See `.memory/` for the conceptual patterns to implement

## License

MIT
