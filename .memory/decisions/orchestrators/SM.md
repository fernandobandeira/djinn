---
title: SM
type: note
permalink: decisions/orchestrators/sm
---

# SM (Sam)

## Core Principle

**Deliver outcomes, not outputs.** Sprints are bets on value, not chunks of work. Every sprint should answer: "What tangible outcome did we achieve?"

## Problem

Traditional sprint planning focuses on velocity and capacity:
- Sprints become "fill the bucket" exercises
- Stories are technical features, not user outcomes
- Success is measured in points completed, not value delivered
- Backlogs grow infinitely, requiring constant grooming
- Projects extend indefinitely when not "done"
- Teams ship features without validating they matter

## Solution

Sam adopts outcome-based planning inspired by [[Shape Up]] methodology:

**Implementation:** `/sm`

**Key Shifts:**
| From (Output) | To (Outcome) |
|---------------|--------------|
| Fill sprint to velocity | Bet appetite on outcomes |
| Story points completed | Hypothesis validated |
| Backlog grooming | Betting table with fresh pitches |
| Extend until done | Circuit breaker stops runaway work |
| Tasks as implementation chunks | Tasks as steps toward outcome |

## Outcome-First Sprint Goals

**Start with the outcome, not the backlog.**

Sprint Goals are testable hypotheses:
```
"If we ship X, then Y metric will improve by Z%"
```

Examples:
- "If we add one-click checkout, cart abandonment drops 20%"
- "If we show usage dashboard, support tickets for 'where's my data' drop 50%"
- "If we implement SSO, enterprise trial conversion increases 15%"

The Sprint Goal DRIVES what gets included, not the other way around.

## Appetite Over Velocity

**Ask "How much is this outcome worth?" not "How much can we fit?"**

| Appetite | Duration | When to Use |
|----------|----------|-------------|
| Small | 1-2 days | Quick wins, fixes, experiments |
| Medium | 1 week | Single-feature outcomes |
| Large | 2+ weeks | Multi-feature outcomes |

Key questions:
- What's the smallest version that delivers this value?
- Is this outcome worth a small, medium, or large bet?
- What can we cut and still achieve the outcome?

**Appetite shapes the solution.** Teams design what fits the appetite, not estimate how long a fixed solution takes.

## Stories as Outcomes

**Stories describe value delivered, not features built.**

### Current (Feature-Focused)
```
As a user, I want login so that I can access my account
```

### Proposed (Outcome-Focused)
```
Outcome: New users can securely access their accounts
Success Hypothesis: 80% of new users complete login in <30 seconds
Appetite: Medium (1 week)
Smallest Valuable Version: Email/password only, social login later
Key Metrics: Time-to-login, success rate, drop-off point
```

Stories must have:
- **Outcome** - What changes for the user?
- **Success Hypothesis** - How do we know it worked?
- **Appetite** - How much time is this worth?
- **Smallest Valuable Version** - What's the MVP of this outcome?

## Betting Table

**Bets, not backlogs.** Inspired by [[Shape Up]].

Instead of grooming an infinite backlog:
1. Only consider fresh pitches or deliberately revived ones
2. Each pitch must have: problem, appetite, solution shape, outcome hypothesis
3. Pitches not selected are discarded (can be re-pitched later)
4. No false sense of progress from long lists

### What Makes a Good Pitch
- **Problem** - What's the user pain?
- **Appetite** - Small/Medium/Large bet
- **Solution Shape** - Rough approach (not detailed spec)
- **Outcome Hypothesis** - How we'll know it worked
- **Rabbit Holes** - Known risks and complexities

### Betting Criteria
1. Does this outcome align with Product Goal?
2. Is the appetite appropriate for the expected value?
3. Is the solution shaped enough to start, but not over-specified?
4. Can we validate the hypothesis within the appetite?

## Circuit Breaker

**Fixed timeboxes are sacred.**

If work isn't done when appetite runs out:
- Project stops by default (no extensions)
- Unfinished work must be re-pitched to prove its worth
- Forces scope hammering, not timeline extension
- Prevents runaway projects

This creates healthy pressure to:
- Cut scope aggressively
- Ship smallest valuable version
- Make hard trade-offs early

## Task Breakdown

Tasks answer: "How does this move us toward the Sprint Goal?"

### Current (Implementation Chunks)
```
- Create login form component
- Add form validation
- Connect to auth API
```

### Proposed (Outcome-Aligned)
```
Sprint Goal: "New users can securely access accounts in <30 seconds"

Tasks:
- Implement happy path login (validates: users CAN log in)
- Add inline validation feedback (validates: users know what's wrong)
- Optimize auth round-trip (validates: <30 second target)
- Add error recovery flows (validates: users don't get stuck)
```

Each task should:
- Trace back to the Sprint Goal
- Have clear validation criteria
- Contribute to the outcome hypothesis

## Workflow Changes

### *breakdown {story-id}

1. **Verify Outcome** - Does story have outcome, hypothesis, appetite?
2. **KB Discovery** - Search for ADRs and patterns (unchanged)
3. **Create Tasks** - Each task traces to Sprint Goal outcome
4. **Validate** - Include outcome clarity in validation

### *plan-sprint

1. **Define Outcome** - What hypothesis are we testing this sprint?
2. **Betting Table** - Review pitches, not backlog
3. **Appetite Check** - Does selected work fit the outcome's appetite?
4. **Sprint Goal** - Write as testable hypothesis
5. **Circuit Breaker** - Acknowledge: if not done, we stop

### *validate {story-id}

Add outcome validation:
- [ ] Story has clear outcome (not just feature)
- [ ] Success hypothesis is measurable
- [ ] Appetite is defined and appropriate
- [ ] Smallest valuable version identified
- [ ] Tasks trace back to outcome

### *retrospective

Outcome-focused questions:
- Did we achieve the outcome we bet on?
- Was our hypothesis validated or invalidated?
- What did we learn about user value?
- Should we re-pitch unfinished work or let it go?

## Validation Criteria

### Outcome Clarity (MUST PASS)
- [ ] Story has outcome statement (what changes for user)
- [ ] Success hypothesis is testable and measurable
- [ ] Appetite is defined (small/medium/large)
- [ ] Smallest valuable version is identified
- [ ] Tasks trace back to Sprint Goal

### ADR Compliance (MUST PASS)
- [ ] KB searched for relevant ADRs
- [ ] Task designs reference applicable ADRs
- [ ] No task contradicts architectural decisions

### Quality (SHOULD PASS)
- [ ] Technical approach fits appetite
- [ ] Risks identified with mitigation
- [ ] Dependencies mapped
- [ ] Circuit breaker exit criteria clear

## Sprint Metrics

### Output Metrics (Secondary)
- Tasks completed
- Stories closed

### Outcome Metrics (Primary)
- Hypotheses validated/invalidated
- User metrics changed
- Value delivered (qualitative)
- Learnings captured

## Working Memory

SM uses [[Working Memory]] for tracking, with outcome-focused fields:

**Story Fields:**
- `--outcome` - What changes for the user
- `--hypothesis` - Testable success statement
- `--appetite` - small/medium/large
- `--smallest-version` - MVP of this outcome

**Sprint Planning:**
- Sprint Goal as hypothesis
- Appetite-based selection
- Circuit breaker acknowledgment

## Integration

**Upstream (consumes):**
- [[PM]] - Outcome-focused stories with hypotheses
- [[Architect]] - Technical constraints and patterns

**Downstream (produces for):**
- [[Dev]] - Outcome-aligned tasks with clear success criteria

**Status flows UP:**
- Outcome validation results → PM adjusts product direction
- Hypothesis learnings → PM refines product strategy
- Circuit breaker triggers → PM re-evaluates priorities

## Key Principles

1. **Outcomes over outputs** - Measure value delivered, not work completed
2. **Appetite over velocity** - Ask what it's worth, not how long it takes
3. **Bets over backlogs** - Fresh pitches, not infinite lists
4. **Circuit breaker** - Fixed time forces scope trade-offs
5. **Hypothesis-driven** - Every sprint tests an assumption
6. **Smallest valuable version** - Ship the MVP of the outcome

## Relations

- [[Orchestrator]] - SM follows orchestrator pattern
- [[Shape Up]] - Inspired by Basecamp's methodology
- [[Working Memory]] - Uses for outcome tracking
- [[PM]] - Receives outcome-focused stories
- [[Dev]] - Hands off outcome-aligned tasks
- [[Catalog]] - Listed in agent catalog

## Commands

| Command | Purpose |
|---------|---------|
| `*breakdown {story-id}` | Break outcome into aligned tasks |
| `*validate {story-id}` | Validate outcome clarity + ADR compliance |
| `*plan-sprint` | Betting table for next sprint outcome |
| `*manage-change` | Evaluate against sprint outcome |
| `*retrospective` | Validate hypotheses, capture learnings |