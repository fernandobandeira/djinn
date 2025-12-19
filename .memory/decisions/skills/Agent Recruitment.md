---
title: Agent Recruitment
type: note
permalink: decisions/skills/agent-recruitment
---

# Agent Recruitment

**Tier:** Domain

## Core Principle

**Skills teach HOW to think. Sub-agents isolate context. Orchestrators coordinate both.**

This guides all type selection:
- Need thinking technique → **Skill**
- Need context isolation (heavy I/O) → **Sub-agent**
- Need explicit invocation with workflow → **Command/Orchestrator**

Orchestrators do reasoning work directly (using skills) and only delegate to sub-agents for context isolation.

## Problem

Creating Djinn components is inconsistent and error-prone:
- Users don't know which type to create (command/orchestrator vs skill vs sub-agent)
- Orchestrators are complex—coordinating skills and sub-agents requires careful design
- Reusability decisions are made ad-hoc (skill tiers, shared vs embedded)
- Validation is skipped or incomplete
- Created components don't follow established patterns

## Solution

Encapsulated frameworks for systematic creation of all Djinn component types.

**Creates:**
- **Commands/Orchestrators** - Workflow personas that coordinate skills and sub-agents
- **Skills** - Thinking techniques that auto-activate on context
- **Sub-agents** - Context isolation for heavy I/O

**Auto-activates when:** User mentions "create agent", "build agent", "recruit", "new command", "new orchestrator", "design skill", "make a sub-agent", or agent modification requests

---

## Frameworks

### Type Selection Framework

Determine whether to create a Command, Skill, or Sub-agent.

#### The Decision Tree

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

#### Quick Reference Matrix

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

#### Key Insight: Sub-agents for Context Isolation

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

#### Examples by Type

**COMMAND Examples:**
- `/dev` - Development workflow with ongoing dialogue
- `/recruiter` - Agent creation with skills
- `/architect` - Architecture design persona

Characteristics: User types `/command` to start, can have `*subcommands`, maintains conversation context, can call skills

**SKILL Examples:**
- `pdf-processing` - Auto-activates when user works with PDFs
- `agent-recruitment` - Auto-activates on "create agent" phrases
- `ideation` - Brainstorming techniques

Characteristics: Triggers on description match, progressive disclosure (loads files as needed), teaches HOW to think

**SUB-AGENT Examples:**
- `market-researcher` - Heavy web research, returns summary
- `diagram-generator` - Generates diagrams in parallel

Characteristics: Has `tools` field (not `allowed-tools`), has `model` field, returns summarized results, context isolation ONLY

#### The Hybrid Pattern

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

Use when: Complex capability that should auto-activate, but ALSO needs explicit workflow entry point.

#### Red Flags: Wrong Type Selected

- **Should be COMMAND, not Skill:** User expected to type something to start, needs `*subcommands`, maintains workflow state
- **Should be SKILL, not Command:** Domain expertise that applies across contexts, should auto-activate
- **Should be doing work DIRECTLY, not Sub-agent:** Needs reasoning, needs skill access, needs to ask follow-ups
- **Appropriate for SUB-AGENT:** Heavy I/O, can run in parallel, process disposable

---

### Reusability Assessment Framework

Determine the right tier for skills and placement for sub-agents.

#### The Core Question

Before creating any capability, ask: **Who else would use this?**

#### Decision Flow

```
                    What are you creating?
                           │
         ┌─────────────────┴─────────────────┐
         ▼                                   ▼
   THINKING technique                  CONTEXT ISOLATION
   (how to analyze)                    (heavy I/O)
         │                                   │
         ▼                                   ▼
      SKILL                              SUB-AGENT
         │                                   │
         ▼                                   ▼
   How many agents                    Shared location
   would use this?                    (agents/shared/)
         │
    ┌────┼────┐
    ▼    ▼    ▼
  Many  Some  One
    │    │    │
    ▼    ▼    ▼
 Tier1 Tier2 Embed?
```

#### Skill Tiers

**Tier 1: Universal Skills**

Audience: Almost any agent could use this (4+ agents)

| Skill | Used By |
|-------|---------|
| `root-cause` | PM, Dev, UX, Analyst, Architect |
| `ideation` | PM, UX, Analyst, Dev, Marketing |
| `devils-advocate` | PM, Dev, Architect, Analyst |
| `role-playing` | PM, UX, Analyst, any decision-maker |
| `teacher` | Any agent explaining concepts |

Signals it's Tier 1:
- Fundamental thinking technique
- Domain-agnostic
- Would feel "missing" if not available to all agents
- Classic methodology (Six Hats, Five Whys, SCAMPER)

**Tier 2: Domain Skills**

Audience: 2-4 related agents in a domain cluster

| Skill | Used By |
|-------|---------|
| `strategic-analysis` | Analyst, PM, Marketing |
| `user-research` | UX, PM, Analyst |

Signals it's Tier 2:
- Valuable but domain-adjacent
- Specific methodology cluster
- Clear set of agents that would use it

**When NOT to Make a Skill**

Embed instead if:
- Only one agent would ever use it
- It's tightly coupled to a specific workflow
- It's more of a checklist than a thinking technique
- Creating it would add complexity without reuse benefit

#### Assessment Checklist

**For Skills:**
- [ ] Is this a thinking technique (not execution)?
- [ ] Would 2+ agents benefit from this?
- [ ] Is it a recognized methodology?
- [ ] Does it require reasoning (can't be delegated)?

If all yes → Create skill (Tier 1 if 4+ agents, Tier 2 if 2-3)

**For Sub-agents:**
- [ ] Is this for context isolation (heavy I/O)?
- [ ] Does it NOT need deep reasoning?
- [ ] Can it return a summary instead of raw data?
- [ ] Would the process flood main context?

If all yes → Create sub-agent in `agents/shared/`

#### Common Mistakes

- **Creating Sub-agents for Reasoning:** BAD: Sub-agent for validation. GOOD: Do reasoning directly
- **Over-sharing:** Creating skills that only one agent uses. Fix: Start embedded, extract when second user emerges
- **Under-sharing:** Duplicating thinking techniques across agents. Fix: Extract to skill when you see duplication
- **Premature Tier 1:** Creating universal skills for niche techniques. Fix: Start at Tier 2, promote when usage proves universal

---

### Model Selection Framework

Choose the right model for sub-agents.

#### Available Models

| Model | Speed | Cost | Best For |
|-------|-------|------|----------|
| **Haiku** | Fast | Low | Simple I/O tasks, context isolation |
| **Sonnet** | Medium | Medium | Research, analysis sub-agents |
| **Opus** | Slow | High | DON'T use for sub-agents (do reasoning directly) |

#### Decision Framework

```
What kind of task is this sub-agent performing?
│
├─► Heavy I/O, formatting, simple output?
│   └─► HAIKU
│       Examples: doc generators, diagram generators
│
├─► Research, data gathering, analysis?
│   └─► SONNET
│       Examples: market researcher, competitive analyzer
│
└─► Complex reasoning, architecture, decisions?
    └─► DON'T USE SUB-AGENT
        Do this work directly in command/skill
```

#### Use HAIKU When

| Scenario | Rationale |
|----------|-----------|
| Document generation | Formatting, no reasoning |
| Diagram generation | Template-based output |
| File operations | Mechanical, predictable |
| Simple transformations | No complex reasoning |
| High-volume I/O | Cost efficiency |

#### Use SONNET When

| Scenario | Rationale |
|----------|-----------|
| Web research | Needs to assess sources |
| Competitive analysis | Some judgment required |
| Data summarization | Must identify key points |
| Heavy reading + summary | Context isolation needed |

#### DON'T Use Sub-agents When

| Scenario | Do Instead |
|----------|------------|
| Validation | Do directly in skill |
| Architecture decisions | Do directly with skills |
| Planning | Do directly with skills |
| Anything needing follow-up questions | Do directly |
| Complex reasoning | Do directly |

#### Key Insight

**If you need opus-level reasoning, don't use a sub-agent.**

Sub-agents can't call skills and have limited reasoning. If the task requires deep reasoning, do it directly in the command/skill.

#### Anti-Patterns

- **Using Opus for Sub-agents:** If you need opus reasoning, do it directly
- **Using Sub-agents for Reasoning:** Validation, planning, architecture → do directly
- **Not Specifying Model:** Always specify model for sub-agents

---

### Validation Workflow

Every agent creation MUST pass validation before completion.

#### Three Validation Gates

```
Agent Created
      │
      ▼
┌──────────────────┐
│ 1. RESOURCE      │
│    VALIDATION    │
│    Files exist?  │
└──────────────────┘
      │
      ├── FAIL → Fix missing files → Retry
      │
      ▼ PASS
┌──────────────────┐
│ 2. CONSTRAINT    │
│    VALIDATION    │
│    Balance ok?   │
└──────────────────┘
      │
      ├── FAIL → Adjust constraints → Retry
      │
      ▼ PASS
┌──────────────────┐
│ 3. COHERENCE     │
│    VERIFICATION  │
│    Aligned?      │
└──────────────────┘
      │
      ├── FAIL → Fix inconsistencies → Retry
      │
      ▼ PASS
┌──────────────────┐
│  VALIDATION      │
│  COMPLETE        │
└──────────────────┘
```

#### 1. Resource Validation

Check all referenced files exist and are syntactically correct.

**Reference Patterns to Check:**

| Pattern | Example | Resolution |
|---------|---------|------------|
| `@` references | `@.claude/resources/file.md` | Direct path |
| Markdown links | `[text](./path/file.md)` | Relative to agent |
| Skill cookbook | `[cookbook/file.md](./cookbook/file.md)` | Relative to skill root |

**Checklist:**
- [ ] All `@path/file.md` references resolve
- [ ] All relative links `[text](./path)` exist
- [ ] YAML frontmatter parses correctly
- [ ] No absolute paths (should be relative to .claude/)

**Pass Criteria:** All referenced files exist, no syntax errors, all paths resolve

#### 2. Constraint Validation

Verify agent constraints are balanced for optimal performance.

**Constraint Score Formula:**

```
Score = Base(5.0)
      + ToolBalance(0-2)      # Appropriate tool count
      + ModelFit(0-1.5)       # Model matches complexity
      + ScopeClarity(0-1.5)   # Clear, focused purpose
      + Integration(0-1)      # Proper integration

Target: 8.0 - 8.5 (optimal)
Warning: < 7.5 or > 9.0
Fail: < 7.0 or > 9.5
```

**Tool Balance Scoring:**

| Tool Count | Agent Type | Score |
|------------|------------|-------|
| 1-2 | Validator | 2.0 |
| 2-3 | Reviewer | 2.0 |
| 3-4 | Builder | 1.5 |
| 5+ | Any | 0.5 |

**Appropriate tool sets by role:**
- Validators: Read, Grep, Glob (3 max)
- Reviewers: Read, Grep, Glob (3 max)
- Builders: Read, Write, MultiEdit, Bash (4 max)
- Orchestrators: Read, Task (2 typical)

**Model Fit Scoring:**

| Task Type | Expected Model | Score |
|-----------|----------------|-------|
| Simple/mechanical | haiku | 1.5 |
| Analysis/judgment | sonnet | 1.5 |
| Mismatch | any | 0.5 |

**Pass Criteria:** Score in 8.0-8.5 range, no critical issues, clear purpose

#### 3. Coherence Verification

Verify all components integrate correctly.

**Coherence Scoring:**

```
Coherence Score = 1.0
  - 0.1 per naming issue
  - 0.15 per reference issue
  - 0.1 per documentation mismatch
  - 0.2 per missing component
  - 0.3 per circular dependency

Pass threshold: 0.85
```

**Coherence Checks:**
- [ ] Sub-agents reference correct parent paths
- [ ] Agent name in frontmatter matches filename
- [ ] Lowercase with hyphens (kebab-case)
- [ ] Description matches actual behavior
- [ ] All referenced sub-agents exist (for orchestrators)
- [ ] No circular dependencies

**Common Issues:**

| Type | Example | Deduction |
|------|---------|-----------|
| Naming | `name: agent-planner` but file is `planner.md` | -0.1 |
| Reference | Orchestrator references missing sub-agent | -0.15 |
| Documentation | `*patterns` command not implemented | -0.1 |
| Missing | Required cookbook file doesn't exist | -0.2 |
| Circular | A → B → A dependency | -0.3 |

**Pass Criteria:** Coherence score > 0.85, no critical inconsistencies, all components exist

#### Validation Report Format

```markdown
## Validation Report for {agent_name}

### Resource Validation
Status: ✅ PASS | ❌ FAIL
Files Checked: {count}
Missing: {list or "None"}

### Constraint Validation
Status: ✅ PASS | ⚠️ WARNING | ❌ FAIL
Score: {score}/10 (target: 8.0-8.5)

### Coherence Verification
Status: ✅ PASS | ❌ FAIL
Coherence: {score}% (target: >85%)

### Overall
{ALL PASS: "✅ Agent validated successfully!"}
{ANY FAIL: "❌ Validation failed. Issues to fix:"}
```

#### Quality Gates Summary

| Gate | Criteria | Action on Fail |
|------|----------|----------------|
| Resources | All files exist, valid syntax | Create missing, fix syntax |
| Constraints | Score 8.0-8.5 | Adjust tools/model/scope |
| Coherence | Score > 85% | Fix naming/references |

---

## Why It Matters

- **Meta-skill** - Skill for building skills/agents
- **Encapsulates complexity** - Users don't learn internals
- **Ensures consistency** - All agents follow patterns
- **Self-extending** - Djinn grows systematically

## Used By

- [[Recruiter]] - Rita loads this skill for all agent creation work

## Relations

- [[Architecture]] - Part of Tier 2 skill layer
- [[Catalog]] - Listed in component index
- [[Recruiter]] - The command that loads this skill
- [[Orchestrator]] - Key pattern for orchestrator design (sub-agents return synthesis, orchestrators write to KB)
- [[Skill]] - Follows skill pattern