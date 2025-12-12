# Creating Orchestrators

## What is an Orchestrator?

An orchestrator is a **command** that coordinates multiple **sub-agents** to accomplish complex workflows. It follows the principle: **Orchestrate, Don't Execute**.

## When to Create an Orchestrator

Create an orchestrator when:
- Workflow requires multiple specialized capabilities
- Different phases need different expertise
- Parallel execution would be beneficial
- Complex state management across steps
- User needs guided multi-step workflow

## Architecture Pattern

```
┌─────────────────────────────────────────┐
│         ORCHESTRATOR (Command)          │
│  • Receives user request                │
│  • Coordinates workflow                 │
│  • Translates between user & sub-agents │
│  • Reports results                      │
└─────────────────────────────────────────┘
         │           │           │
         ▼           ▼           ▼
    ┌─────────┐ ┌─────────┐ ┌─────────┐
    │Planner  │ │Builder  │ │Validator│
    │(opus)   │ │(haiku)  │ │(haiku)  │
    └─────────┘ └─────────┘ └─────────┘
```

## ULTRATHINK Required

Creating orchestrators is complex. **ALWAYS ultrathink** about:
- How to decompose the workflow
- What sub-agents are needed
- How they communicate
- What each sub-agent returns
- How to aggregate results

## File Structure

```
.claude/
├── commands/
│   └── {orchestrator}.md      # The orchestrator command
└── agents/
    └── {orchestrator}/        # Sub-agents for this orchestrator
        ├── planner.md
        ├── builder.md
        └── validator.md
```

## Orchestrator Structure

```yaml
---
description: Orchestrates {workflow} through specialized sub-agents
allowed-tools: Read, Task  # Minimal - delegates everything
---
```

```markdown
# {Name} Orchestrator

## Activation
[Greeting with persona]

## Core Principle
**Orchestrate, Don't Execute** - All work delegated to sub-agents.

## Sub-agents
- `{name}-planner` - Planning phase
- `{name}-builder` - Building phase
- `{name}-validator` - Validation phase

## Workflow Phases

### Phase 1: Planning
Delegate to `{name}-planner`

### Phase 2: Building
Delegate to `{name}-builder`

### Phase 3: Validation
Delegate to validators

## Commands
- `*help` - Show commands
- `*plan` - Run planning
- `*build` - Run building
- `*validate` - Run validation
- `*status` - Current progress
```

## Decomposition Framework

### Step 1: Identify Phases

Break workflow into distinct phases:
```
Example: Agent Creation
├── Planning Phase (requirements, architecture)
├── Building Phase (file creation)
├── Validation Phase (quality checks)
└── Learning Phase (pattern extraction)
```

### Step 2: Define Sub-agent Responsibilities

Each phase becomes a sub-agent:
```yaml
agent-planner:
  responsibility: Analyze requirements, design architecture
  model: opus (complex reasoning)
  tools: Read, Grep, Glob, WebFetch

agent-builder:
  responsibility: Create files, generate resources
  model: haiku (mechanical tasks)
  tools: Read, Write, MultiEdit, Bash

constraint-validator:
  responsibility: Check constraint balance
  model: haiku (rule-based)
  tools: Read, Grep
```

### Step 3: Define Interfaces

What does each sub-agent receive and return?

```yaml
planner_interface:
  receives:
    - agent_name: string
    - user_description: string
    - context: object
  returns:
    - agent_type: command | skill | subagent
    - specifications: object
    - rationale: string

builder_interface:
  receives:
    - specifications: object (from planner)
  returns:
    - created_files: [string]
    - errors: [string]

validator_interface:
  receives:
    - agent_path: string
  returns:
    - score: float
    - issues: [object]
    - recommendations: [string]
```

### Step 4: Design Translation Layer

Orchestrator translates between user and sub-agents:

```markdown
## Translation Rules

1. **User → Sub-agent**: Prepare structured context
2. **Sub-agent → User**: Format results readably
3. **Never**: Pass raw user input to sub-agents
4. **Never**: Show raw sub-agent output to users
```

## Delegation Patterns

### Sequential Delegation
```
User Request
    │
    ▼
┌─────────┐     ┌─────────┐     ┌─────────┐
│ Phase 1 │────▶│ Phase 2 │────▶│ Phase 3 │
└─────────┘     └─────────┘     └─────────┘
    │               │               │
    ▼               ▼               ▼
  Result          Result          Result
    │               │               │
    └───────────────┴───────────────┘
                    │
                    ▼
              Final Report
```

### Parallel Delegation
```
User Request
    │
    ├────────────┬────────────┐
    ▼            ▼            ▼
┌─────────┐ ┌─────────┐ ┌─────────┐
│Validator│ │Validator│ │Validator│
│    1    │ │    2    │ │    3    │
└─────────┘ └─────────┘ └─────────┘
    │            │            │
    └────────────┴────────────┘
                 │
                 ▼
           Aggregated Report
```

### Conditional Delegation
```
User Request
    │
    ▼
┌──────────────┐
│ Type Check   │
└──────────────┘
    │
    ├─── command? ───▶ Command Builder
    │
    ├─── skill? ─────▶ Skill Builder
    │
    └─── subagent? ──▶ Sub-agent Builder
```

## Example: Rita (Recruiter Orchestrator)

```markdown
---
description: Orchestrates agent creation through specialized sub-agents
allowed-tools: Read, Task
---

# Rita - Recruiter Orchestrator

## Activation
Hello! I'm Rita, your Recruiter Orchestrator.
I coordinate agent creation through specialized assistants.

## Core Principle
**Orchestrate, Don't Execute**

## Sub-agents
- `agent-planner` - Requirements and architecture (opus)
- `agent-builder` - File creation (haiku)
- `pattern-extractor` - Learning (sonnet)
- `constraint-validator` - Balance checks (haiku)
- `resource-validator` - File verification (haiku)
- `coherence-verifier` - Integration checks (haiku)

## Workflow

### *plan {name}
1. Gather user requirements
2. Delegate to agent-planner
3. Present plan for approval

### *build
1. Verify plan approved
2. Delegate to agent-builder
3. Report created files

### *validate
1. Run resource-validator
2. Run constraint-validator
3. Run coherence-verifier
4. Aggregate results

## Commands
- `*help` - Show commands
- `*recruit {name}` - Full workflow
- `*plan {name}` - Planning only
- `*build` - Building only
- `*validate` - Validation only
- `*status` - Current progress
```

## State Management

Orchestrators need minimal state:

```yaml
workflow_state:
  current_agent: string | null
  current_phase: Planning | Building | Validating | Complete
  plan: object | null
  files_created: [string]
  validation_results: object | null
```

Track state in conversation context, not external storage.

## Error Handling

### Planning Errors
```
If planner fails:
  → Ask user for clarification
  → Retry with more context
```

### Building Errors
```
If builder fails:
  → Report specific error
  → Ask user to resolve blockers
  → Retry after resolution
```

### Validation Errors
```
If validation fails:
  → Report issues clearly
  → Suggest fixes
  → Delegate fix to builder
  → Re-validate
```

## Best Practices

### Keep Orchestrator Thin
```
GOOD: ~100 lines, mostly delegation
BAD: 500+ lines with embedded logic
```

### Minimal Tools
```yaml
allowed-tools: Read, Task  # Only needs to delegate
```

### Clear Phase Boundaries
```
Each phase:
- Has clear entry/exit criteria
- Receives specific inputs
- Returns specific outputs
```

### User Approval Gates
```
Don't auto-proceed between major phases.
Get user confirmation:
- After planning: "Shall I proceed?"
- After building: "Ready to validate?"
- After validation: "Issues found, fix them?"
```

## Checklist

Before finalizing an orchestrator:

- [ ] Clear decomposition into sub-agents
- [ ] Each sub-agent has single responsibility
- [ ] Interfaces defined (input/output)
- [ ] Translation layer documented
- [ ] State management minimal
- [ ] Error handling for each phase
- [ ] User approval gates in place
- [ ] Orchestrator under 150 lines
- [ ] Sub-agents in correct folder
- [ ] All validators pass
