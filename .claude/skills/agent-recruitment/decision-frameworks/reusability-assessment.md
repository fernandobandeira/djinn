# Reusability Assessment Framework

## The Core Question

Before creating any capability, ask: **Who else would use this?**

This determines:
- Whether it's a skill or sub-agent
- Which tier/location it belongs to
- How to design its interface

## The Decision Flow

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

## Thinking vs Context Isolation

### Thinking Techniques → Skills

These teach HOW to analyze or approach problems:
- Root cause analysis (Five Whys, First Principles)
- Ideation methods (SCAMPER, Walt Disney)
- Perspective techniques (Six Hats, Stakeholder Roundtable)
- Challenge methods (Pre-mortem, Red Team)

**Characteristics:**
- Reusable mental models
- Guide the thinking process
- Work across domains
- Need to be done directly (require reasoning)

### Context Isolation → Sub-agents

These DO work that would flood main context:
- Market research (web searches, data gathering)
- Document generation (many file writes)
- Heavy data processing (many reads/transforms)

**Characteristics:**
- Heavy I/O that would pollute context
- Produce summarized outputs
- Process disposable, result matters
- DON'T need deep reasoning

## Skill Tiers

### Tier 1: Universal Skills

**Audience:** Almost any agent could use this

**Examples:**
| Skill | Used By |
|-------|---------|
| `root-cause` | PM, Dev, UX, Analyst, Architect |
| `ideation` | PM, UX, Analyst, Dev, Marketing |
| `devils-advocate` | PM, Dev, Architect, Analyst |
| `role-playing` | PM, UX, Analyst, any decision-maker |
| `teacher` | Any agent explaining concepts |

**Signals it's Tier 1:**
- Fundamental thinking technique
- Domain-agnostic
- Would feel "missing" if not available to all agents
- Classic methodology (Six Hats, Five Whys, SCAMPER)

### Tier 2: Domain Skills

**Audience:** 2-4 related agents in a domain cluster

**Examples:**
| Skill | Used By |
|-------|---------|
| `strategic-analysis` | Analyst, PM, Marketing |
| `user-research` | UX, PM, Analyst |

**Signals it's Tier 2:**
- Valuable but domain-adjacent
- Specific methodology cluster
- Clear set of agents that would use it
- Others might use occasionally but not core to their work

### When NOT to Make a Skill

**Embed instead if:**
- Only one agent would ever use it
- It's tightly coupled to a specific workflow
- It's more of a checklist than a thinking technique
- Creating it would add complexity without reuse benefit

## Sub-agent Placement

### All Sub-agents → `agents/shared/`

Since sub-agents are for **context isolation only** (not reasoning), they go in the shared folder. Multiple orchestrators might need similar I/O isolation.

| Sub-agent | Purpose |
|-----------|---------|
| `market-researcher` | Heavy web research |
| `competitive-analyzer` | Research + comparison |
| `knowledge-harvester` | External source harvesting |
| `diagram-generator` | Diagram creation |

**Key point:** Don't create agent-specific sub-agents for reasoning work. Do reasoning directly in the command/skill.

## Assessment Checklist

Before creating, answer these:

### For Skills
- [ ] Is this a thinking technique (not execution)?
- [ ] Would 2+ agents benefit from this?
- [ ] Is it a recognized methodology?
- [ ] Does it require reasoning (can't be delegated)?

**If all yes → Create skill (Tier 1 if 4+ agents, Tier 2 if 2-3)**

### For Sub-agents
- [ ] Is this for context isolation (heavy I/O)?
- [ ] Does it NOT need deep reasoning?
- [ ] Can it return a summary instead of raw data?
- [ ] Would the process flood main context?

**If all yes → Create sub-agent in `agents/shared/`**

## Examples

### Example 1: JTBD (Jobs-to-be-Done)

- Is it thinking or context isolation? **Thinking** (framework for understanding needs)
- Who would use it? **PM, UX, Analyst, Marketing** (4+ agents)
- Decision: **Tier 1 Skill** ✓ (already in `root-cause`)

### Example 2: Market Research

- Is it thinking or context isolation? **Context isolation** (heavy web searches)
- Does it need reasoning? **Some** (assessing sources)
- Would it flood context? **Yes** (many web fetches)
- Decision: **Sub-agent** → `agents/shared/market-researcher.md`

### Example 3: Constraint Validation

- Is it thinking or context isolation? **Thinking** (requires judgment)
- Does it need reasoning? **Yes** (calculating scores, identifying issues)
- Decision: **Do directly in skill** (not a sub-agent)

### Example 4: SWOT Analysis

- Is it thinking or context isolation? **Thinking** (strategic framework)
- Who would use it? **Analyst, PM, Marketing** (3 agents)
- Decision: **Tier 2 Skill** → `skills/strategic-analysis/`

### Example 5: Architecture Planning

- Is it thinking or context isolation? **Thinking** (complex reasoning)
- Does it need skill access? **Yes** (devils-advocate, strategic-analysis)
- Decision: **Do directly in command** (not a sub-agent)

## Integration with Type Selection

This framework runs AFTER type selection:

```
1. type-selection.md    →  Command, Skill, or Sub-agent?
2. reusability-assessment.md  →  Which tier? (for skills)
3. cookbook/{type}.md   →  How to build it?
```

## Common Mistakes

### Creating Sub-agents for Reasoning
```
BAD: Sub-agent for validation or planning
GOOD: Do reasoning directly (has skill access)
```

### Over-sharing
Creating skills that only one agent actually uses.
**Fix:** Start embedded, extract when second user emerges.

### Under-sharing
Duplicating thinking techniques across agents.
**Fix:** Extract to skill when you see duplication.

### Wrong Type
Making context isolation into skills (or vice versa).
**Fix:** Skills guide thinking, sub-agents isolate I/O.

### Premature Tier 1
Creating universal skills for niche techniques.
**Fix:** Start at Tier 2, promote to Tier 1 when usage proves universal.
