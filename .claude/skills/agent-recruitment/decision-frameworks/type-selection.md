# Agent Type Selection Framework

## The Core Decision Tree

```
                    What are you creating?
                           │
         ┌─────────────────┼─────────────────┐
         ▼                 ▼                 ▼
    User invokes      Auto-activates    Context isolation
    explicitly?       on context?        needed?
         │                 │                 │
         ▼                 ▼                 ▼
      COMMAND            SKILL           SUB-AGENT
```

## Detailed Decision Flow

```
START: What capability are you building?
│
├─► Does the user need to EXPLICITLY invoke this?
│   │
│   ├─► YES: Will there be ongoing dialogue/state?
│   │   │
│   │   ├─► YES ────────────────────────► COMMAND
│   │   │   (Workflows, personas, orchestrators)
│   │   │
│   │   └─► NO: Is it a simple, quick task?
│   │       │
│   │       ├─► YES ────────────────────► COMMAND
│   │       │   (Shortcuts, templates)
│   │       │
│   │       └─► NO: Complex domain expertise?
│   │           │
│   │           └─► Consider SKILL + COMMAND alias
│   │
│   └─► NO: Should it auto-activate on context?
│       │
│       ├─► YES: Does it have multi-file resources?
│       │   │
│       │   ├─► YES ────────────────────► SKILL
│       │   │   (Progressive disclosure)
│       │   │
│       │   └─► NO ─────────────────────► SKILL
│       │       (Simpler skill)
│       │
│       └─► NO: Heavy I/O that would flood context?
│           │
│           ├─► YES ────────────────────► SUB-AGENT
│           │   (Context isolation)
│           │
│           └─► NO: Can run in parallel?
│               │
│               └─► YES ────────────────► SUB-AGENT
│                   (Parallel execution)
```

## Quick Reference Matrix

| Signal | Command | Skill | Sub-agent |
|--------|:-------:|:-----:|:---------:|
| User explicitly invokes with `/` | YES | - | - |
| Auto-activates on context match | - | YES | - |
| Maintains dialogue/state | YES | - | - |
| Heavy I/O (would flood context) | - | - | YES |
| Parallel task execution | - | - | YES |
| Multi-file resources/scripts | possible | YES | - |
| Reusable domain expertise | possible | YES | - |
| Needs reasoning/skill access | YES | YES | NO |

## Key Insight: Sub-agents for Context Isolation

**Sub-agents can't reason deeply and can't call skills.**

Use sub-agents ONLY when:
- Heavy I/O would flood main context (research, many web fetches)
- Tasks can run independently in parallel
- Process is disposable, only output matters

Do NOT use sub-agents for:
- Validation (needs reasoning)
- Architecture decisions (needs skills + context)
- Anything interactive (can't ask follow-ups)
- Any work requiring deep reasoning

## Examples by Type

### COMMAND Examples
- `/dev` - Development workflow with ongoing dialogue
- `/recruiter` - Agent creation with skills
- `/review-pr` - Quick shortcut for PR review
- `/architect` - Architecture design persona

**Characteristics**:
- User types `/command` to start
- Can have `*subcommands` for multi-mode
- Maintains conversation context
- Can call skills for reasoning

### SKILL Examples
- `pdf-processing` - Auto-activates when user works with PDFs
- `agent-recruitment` - Auto-activates on "create agent" phrases
- `go-patterns` - Domain expertise for Go development
- `ideation` - Brainstorming techniques

**Characteristics**:
- Triggers on description match
- Progressive disclosure (loads files as needed)
- Can include scripts and tools
- Teaches HOW to think

### SUB-AGENT Examples
- `market-researcher` - Heavy web research, returns summary
- `diagram-generator` - Generates diagrams in parallel
- `documentation-generator` - Heavy doc generation

**Characteristics**:
- Has `tools` field (not `allowed-tools`)
- Has `model` field (haiku/sonnet)
- Returns summarized results
- Context isolation ONLY

## The Hybrid Pattern

Sometimes you need BOTH auto-discovery AND explicit invocation:

```
.claude/
├── skills/
│   └── agent-recruitment/      # Auto-discovers
│       └── SKILL.md
└── commands/
    └── recruiter.md            # Explicit /recruiter
        (loads the skill)
```

**When to use hybrid**:
- Complex capability that should auto-activate
- But ALSO needs explicit workflow entry point
- User might say "create an agent" (skill triggers)
- Or explicitly type `/recruiter` (command triggers)

## Model Selection by Type

| Type | Default Model | Notes |
|------|---------------|-------|
| Command | (inherits) | Does work directly |
| Skill | (inherits) | Does work directly |
| Sub-agent | haiku/sonnet | For context isolation only |

## Tool Selection Guidelines

### Commands
```yaml
allowed-tools: [tools needed for the domain]
# Example: Read, Write, Grep, Glob, Bash, MultiEdit
```

### Skills
```yaml
allowed-tools: [tools needed for the domain]
# Example: Read, Write, Bash, Grep, Glob
```

### Sub-agents
```yaml
tools: [specific tools for isolated I/O]
# Example: Read, WebFetch, WebSearch (for research)
# Example: Read, Write (for document generation)
```

## Red Flags: Wrong Type Selected

### Should be COMMAND, not Skill
- User expected to type something to start
- Needs `*subcommands` or modes
- Maintains workflow state across turns

### Should be SKILL, not Command
- Domain expertise that applies across contexts
- Multiple reference files needed
- Should auto-activate ("when I work with X...")

### Should be doing work DIRECTLY, not Sub-agent
- Needs reasoning or judgment
- Needs access to skills
- Needs to ask follow-up questions
- Validation, planning, architecture decisions

### Appropriate for SUB-AGENT
- Heavy I/O (many web fetches, file reads)
- Can run in parallel with other work
- Process disposable, only summary matters

## Decision Checklist

Before finalizing type, verify:

- [ ] **Invocation method** clear (user vs auto)
- [ ] **Reasoning needs** understood (direct vs isolated)
- [ ] **Tool set** appropriate for the type
- [ ] **Model selection** justified (if sub-agent)
- [ ] **File location** correct for type
- [ ] **Progressive disclosure** considered (if complex)

## After Type Selection: Assess Reusability

Once you've determined Command, Skill, or Sub-agent, assess reusability:

**Read [reusability-assessment.md](./reusability-assessment.md)** for the full framework.

### Quick Reference

**For Skills - Which Tier?**

| Tier | Audience | Signal |
|------|----------|--------|
| **Tier 1: Universal** | 4+ agents | Fundamental thinking technique, domain-agnostic |
| **Tier 2: Domain** | 2-3 agents | Domain-adjacent, clear agent cluster |
| **Embed** | 1 agent | Tightly coupled, not worth extracting |

**For Sub-agents**

Sub-agents go in `agents/shared/` since they're for context isolation and multiple orchestrators might need similar I/O isolation.

### The Key Questions

1. **Is this THINKING or CONTEXT ISOLATION?**
   - Thinking (how to analyze) → Skill
   - Context isolation (heavy I/O) → Sub-agent

2. **Who else would use this?**
   - Many agents → Tier 1 skill
   - Few agents → Tier 2 skill
   - One agent → Embed in command/skill

## When Uncertain: ULTRATHINK

If the decision isn't clear after this framework:

1. List the specific requirements
2. Map each to type characteristics
3. Identify conflicts or edge cases
4. ULTRATHINK about the tradeoffs
5. Consider hybrid pattern if needed
6. **Assess reusability** - who else would use this?
