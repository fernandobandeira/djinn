# Creating Orchestrators

## What is an Orchestrator?

An orchestrator is a **command** that coordinates complex workflows. It does reasoning work directly and only uses **sub-agents for context isolation** when heavy I/O would pollute the main context.

## When to Create an Orchestrator

Create an orchestrator when:
- Workflow requires multiple phases
- User needs guided multi-step workflow
- Different phases have distinct purposes
- Complex state management across steps

## Key Principle

**Do Work Directly. Sub-agents for Context Isolation Only.**

- Reasoning, validation, planning → Do directly (access to skills)
- Heavy I/O, parallel tasks → Delegate to sub-agents (context isolation)

## Architecture Pattern

```
┌─────────────────────────────────────────┐
│         ORCHESTRATOR (Command)          │
│  • Receives user request                │
│  • Does reasoning work directly         │
│  • Uses skills for thinking frameworks  │
│  • Only delegates for context isolation │
└─────────────────────────────────────────┘
         │               │
    (reasoning)    (context isolation)
         │               │
         ▼               ▼
    ┌─────────┐    ┌─────────────┐
    │ Skills  │    │ Sub-agents  │
    │ (HOW)   │    │ (heavy I/O) │
    └─────────┘    └─────────────┘
```

## ULTRATHINK Required

Creating orchestrators is complex. **ALWAYS ultrathink** about:
- What work requires reasoning (do directly)
- What work is heavy I/O (delegate for context isolation)
- How phases connect
- What skills to leverage

## File Structure

```
.claude/
├── orchestrators/
│   └── {orchestrator}.md      # The orchestrator definition
└── skills/
    └── {domain}/              # Skills the orchestrator uses
        └── SKILL.md
```

## Orchestrator Structure

```yaml
---
description: Orchestrates {workflow} with domain expertise
allowed-tools: Read, Write, Grep, Glob, Bash, MultiEdit
---
```

```markdown
# {Name} Orchestrator

## Activation
[Greeting with persona]

## Skill Loading
@.claude/skills/{domain}/SKILL.md

## Core Principle
**Do Work Directly** - Use skills for reasoning. Sub-agents only for context isolation.

## Workflow Phases

### Phase 1: {Name}
[What to do directly using skill]

### Phase 2: {Name}
[What to do - delegate if heavy I/O]

### Phase 3: {Name}
[What to do directly]

## Commands
- `*help` - Show commands
- `*{phase1}` - Run phase 1
- `*{phase2}` - Run phase 2
- `*status` - Current progress
```

## Decomposition Framework

### Step 1: Identify Phases

Break workflow into distinct phases:
```
Example: Agent Creation
├── Requirements Phase (reasoning - do directly)
├── Planning Phase (reasoning - do directly, use skills)
├── Building Phase (file creation - do directly)
├── Validation Phase (reasoning - do directly)
└── Research Phase (heavy I/O - delegate if needed)
```

### Step 2: Classify Each Phase

For each phase, ask:

| Question | If Yes → |
|----------|----------|
| Requires reasoning? | Do directly |
| Needs skill access? | Do directly |
| Requires follow-up questions? | Do directly |
| Heavy I/O (many fetches/reads)? | Consider sub-agent |
| Can run in parallel? | Consider sub-agent |
| Process disposable, output matters? | Consider sub-agent |

### Step 3: Design Skill Integration

Map skills to phases:
```yaml
phases:
  requirements:
    work: Gather and analyze requirements
    approach: Direct (reasoning)
    skills: root-cause, ideation

  planning:
    work: Design architecture
    approach: Direct (reasoning)
    skills: systems-thinking, strategic-analysis

  research:
    work: Market/competitive research
    approach: Sub-agent (heavy I/O)
    sub_agent: market-researcher

  building:
    work: Create files
    approach: Direct
    skills: none (mechanical)

  validation:
    work: Check quality
    approach: Direct (reasoning)
    skills: devils-advocate
```

### Step 4: Define Interfaces

What does each phase produce?

```yaml
requirements_output:
  - user_needs: [string]
  - constraints: [string]
  - success_criteria: [string]

planning_output:
  - architecture: object
  - components: [object]
  - rationale: string

research_output:  # From sub-agent
  - summary: string
  - key_findings: [object]
  - recommendations: [string]

building_output:
  - files_created: [string]
  - status: success | failure

validation_output:
  - score: float
  - issues: [object]
  - passed: boolean
```

## Delegation Pattern

### What to Do Directly

```markdown
## Planning Phase

Use the ideation skill for brainstorming options.
Read [cookbook/brainstorming.md](.claude/skills/ideation/cookbook/brainstorming.md)

Apply SCAMPER technique to explore alternatives.

Use systems-thinking skill for architecture.
Read [cookbook/systems-mapping.md](.claude/skills/systems-thinking/cookbook/systems-mapping.md)
```

### What to Delegate (Context Isolation)

```markdown
## Research Phase (if heavy I/O needed)

Delegate to market-researcher sub-agent:
- Heavy web research would flood context
- Results are synthesized back to you

Task(
  subagent_type="market-researcher",
  description="Research market for X",
  prompt="Research {topic}. Return synthesis."
)

# After receiving results, YOU write to KB:
mcp__basic-memory__write_note(
    title=result.suggested_title,
    content=result.synthesized_content,
    folder=result.suggested_folder,
    project="<PRIMARY>"  # From CLAUDE.md
)
```

**Key**: Sub-agents return synthesis. You write to KB with correct project parameter.

## Example: Rita (Recruiter Orchestrator)

```markdown
---
description: Orchestrates agent creation with systematic decomposition
allowed-tools: Read, Write, Grep, Glob, Bash, MultiEdit
---

# Rita - Recruiter

## Activation
Hello! I'm Rita, the Recruiter.
I create agents using best practices.

## Skill Loading
@.claude/skills/agent-recruitment/SKILL.md

## Core Principle
**Do Work Directly** - Use skills for reasoning. Sub-agents only for context isolation.

## Workflow Phases

### *plan {name}
1. Gather requirements (direct - need interaction)
2. Apply type-selection framework (direct - reasoning)
3. Design architecture (direct - reasoning)
4. Present plan for approval

### *build
1. Verify plan approved (direct)
2. Create files using templates (direct)
3. Report created files

### *validate
1. Resource validation (direct - check files exist)
2. Constraint validation (direct - calculate score)
3. Coherence verification (direct - check consistency)
4. Report results

### *research (optional - heavy I/O)
If market/competitive research needed:
- Delegate to market-researcher (context isolation)
- Research would flood main context
- Return summary to inform planning

## Commands
- `*help` - Show commands
- `*recruit {name}` - Full workflow
- `*plan {name}` - Planning only
- `*build` - Building only
- `*validate` - Validation only
- `*status` - Current progress
```

## State Management

Orchestrators track minimal state in conversation:

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

### Reasoning Errors
```
If reasoning leads to unclear results:
  → Ask user for clarification
  → Apply different skill/technique
  → Retry with more context
```

### Building Errors
```
If file creation fails:
  → Report specific error
  → Ask user to resolve blockers
  → Retry after resolution
```

### Validation Errors
```
If validation fails:
  → Report issues clearly
  → Suggest fixes
  → Apply fixes directly
  → Re-validate
```

## Best Practices

### Keep Orchestrator Focused
```
GOOD: ~150 lines, loads skill for domain expertise
BAD: 500+ lines with embedded domain knowledge
```

### Use Skills for Thinking
```
GOOD: Load ideation skill, apply SCAMPER technique
BAD: Implement brainstorming logic in orchestrator
```

### Minimal Sub-agents
```
GOOD: Sub-agent only for heavy research (context isolation)
BAD: Sub-agents for validation, planning, building (reasoning work)
```

### User Approval Gates
```
Don't auto-proceed between major phases.
Get user confirmation:
- After planning: "Shall I proceed?"
- After building: "Ready to validate?"
- After validation issues: "Fix them?"
```

## Checklist

Before finalizing an orchestrator:

- [ ] Clear workflow phases identified
- [ ] Reasoning work done directly
- [ ] Sub-agents only for context isolation
- [ ] **Orchestrator writes sub-agent results to KB** (sub-agents return synthesis)
- [ ] Skills loaded for domain expertise
- [ ] Interfaces defined between phases
- [ ] State management minimal
- [ ] Error handling for each phase
- [ ] User approval gates in place
- [ ] Orchestrator under 150 lines (loads skills)
- [ ] Validates successfully
