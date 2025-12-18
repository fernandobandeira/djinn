# Djinn

AI agent personas for Claude Code that help you brainstorm, research, analyze, and build.

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
| **Commands** | WHO - Role personas with boundaries | /analyst (Ana), /architect (Archie) |
| **Skills** | HOW - Thinking techniques that auto-activate | SCAMPER, Five Whys, Six Hats |
| **Sub-agents** | ISOLATE - Heavy I/O without polluting context | market-researcher, diagram-generator |

Type `/architect` and Claude becomes Archie - thinking like an architect, triggering systems-thinking skills, delegating research to sub-agents while keeping reasoning clean.

## Philosophy

**You stay in the loop.** Djinn isn't about vibe coding or blind delegation. AI accelerates your work - drafting research, applying frameworks, generating options - but you remain the decision maker.

- **AI generates, you review** - Research, plans, and decisions are drafts until you validate them
- **Shared knowledge base** - `.memory/` serves both you and AI; keep it accurate
- **Structured collaboration** - Personas and skills make AI predictable, not autonomous

## Quick Start

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

## Usage

```
/analyst       # Ana - brainstorming, research, strategic analysis
/architect     # Archie - system design, ADRs, diagrams
/recruiter     # Rita - create new agents
```

Once activated, type `*help` to see available commands.

Skills auto-activate based on context - just mention "brainstorm", "root cause", "SWOT", etc.

## Documentation

Full documentation is in `.memory/`:

- **[Project](.memory/Project.md)** - Vision, goals, and full problem/solution context
- **[Architecture](.memory/Architecture.md)** - Core concepts, design rules, how to extend
- **[Guide](.memory/Guide.md)** - Installation details, usage, knowledge management

Use [Obsidian](https://obsidian.md) to browse the docs - they use `[[wikilinks]]` for navigation.

## License

MIT
