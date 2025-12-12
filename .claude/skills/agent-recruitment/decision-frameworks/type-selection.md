# Agent Type Selection Framework

## The Core Decision Tree

```
                    What are you creating?
                           │
         ┌─────────────────┼─────────────────┐
         ▼                 ▼                 ▼
    User invokes      Auto-activates    Called by other
    explicitly?       on context?        agents only?
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
│       ├─► YES: Does it need isolated context?
│       │   │
│       │   ├─► YES ────────────────────► SUB-AGENT
│       │   │   (With IMPORTANT keyword for proactive)
│       │   │
│       │   └─► NO: Does it have multi-file resources?
│       │       │
│       │       ├─► YES ────────────────► SKILL
│       │       │   (Progressive disclosure)
│       │       │
│       │       └─► NO ─────────────────► SKILL or SUB-AGENT
│       │           (Preference: SKILL)
│       │
│       └─► NO: Is it called by other agents only?
│           │
│           └─► YES ────────────────────► SUB-AGENT
│               (Explicit Task tool delegation)
```

## Quick Reference Matrix

| Signal | Command | Skill | Sub-agent |
|--------|:-------:|:-----:|:---------:|
| User explicitly invokes with `/` | YES | - | - |
| Auto-activates on context match | - | YES | partial* |
| Maintains dialogue/state | YES | - | - |
| Needs isolated context window | - | - | YES |
| Returns results to caller | - | - | YES |
| Multi-file resources/scripts | possible | YES | - |
| Reusable domain expertise | possible | YES | - |
| Parallel task execution | - | - | YES |
| Proactive triggering | - | - | YES** |

*Sub-agents can auto-trigger via `IMPORTANT` keyword in description
**Requires `IMPORTANT` in description field

## Examples by Type

### COMMAND Examples
- `/dev` - Development workflow with ongoing dialogue
- `/recruiter` - Agent creation orchestrator
- `/review-pr` - Quick shortcut for PR review
- `/architect` - Architecture design persona

**Characteristics**:
- User types `/command` to start
- Can have `*subcommands` for multi-mode
- Maintains conversation context
- Often orchestrates sub-agents

### SKILL Examples
- `pdf-processing` - Auto-activates when user works with PDFs
- `agent-recruitment` - Auto-activates on "create agent" phrases
- `go-patterns` - Domain expertise for Go development
- `testing-strategies` - Auto-loads when discussing tests

**Characteristics**:
- Triggers on description match
- Progressive disclosure (loads files as needed)
- Can include scripts and tools
- Stateless (no conversation memory)

### SUB-AGENT Examples
- `backend-qa-reviewer` - IMPORTANT reviews Go code after changes
- `agent-planner` - Called by Rita for planning
- `constraint-validator` - Called for validation checks
- `pattern-extractor` - Called to learn from success

**Characteristics**:
- Has `tools` field (not `allowed-tools`)
- Has `model` field (haiku/sonnet/opus)
- Returns structured results
- Isolated context window

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

| Type | Default Model | When to Change |
|------|---------------|----------------|
| Command | (inherits) | Complex orchestration → sonnet/opus |
| Skill | (inherits) | N/A (skills don't specify model) |
| Sub-agent | haiku | Complex analysis → sonnet, Critical decisions → opus |

## Tool Selection Guidelines

### Commands
```yaml
allowed-tools: [minimal set for orchestration]
# Often just: Read, Task (for delegation)
```

### Skills
```yaml
allowed-tools: [tools needed for the domain]
# Example: Read, Write, Bash, Grep
```

### Sub-agents
```yaml
tools: [specific tools for the task]
# Be restrictive - only what's needed
# Example: Read, Grep, Glob (for reviewers)
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

### Should be SUB-AGENT, not Command
- Only called by other agents, never by users
- Needs isolated context (large file analysis)
- Returns structured data to caller

### Should be SUB-AGENT, not Skill
- Needs to run in parallel with other work
- Must not pollute main conversation context
- Heavy computation or long-running task

## Decision Checklist

Before finalizing type, verify:

- [ ] **Invocation method** clear (user vs auto vs delegated)
- [ ] **Context needs** understood (shared vs isolated)
- [ ] **Tool set** appropriate for the type
- [ ] **Model selection** justified (if sub-agent)
- [ ] **File location** correct for type
- [ ] **Progressive disclosure** considered (if complex)

## When Uncertain: ULTRATHINK

If the decision isn't clear after this framework:

1. List the specific requirements
2. Map each to type characteristics
3. Identify conflicts or edge cases
4. ULTRATHINK about the tradeoffs
5. Consider hybrid pattern if needed
