---
name: systems-thinking
description: Understand complex systems through interconnections, feedback loops, and emergent behavior. Use when user wants "systems diagram", "feedback loop", "leverage points", "technical debt", "system dynamics", "emergent behavior", "interconnections", or needs to understand how parts of a system influence each other.
allowed-tools: Read
---

# Systems Thinking - Complexity Navigator

Understand how complex systems behave through interconnections, feedback loops, and emergent properties. See the forest AND the trees.

## Quick Start

1. Identify the system you're trying to understand
2. Select technique based on what aspect you need to analyze
3. Map the system and find intervention points

## Technique Selection

| Need | Use | Why |
|------|-----|-----|
| Visualize system structure | Systems Mapping | See components and relationships |
| Understand dynamic behavior | Feedback Loop Analysis | Find what drives growth or stability |
| Identify technical burden | Technical Debt Analysis | Quantify and prioritize debt payoff |

**Default**: Start with Systems Mapping, then identify feedback loops.

## Techniques

### Systems Mapping
Visualize the components, relationships, and boundaries of a system. Make the invisible visible.

Read [cookbook/systems-mapping.md](./cookbook/systems-mapping.md)

### Feedback Loop Analysis
Identify reinforcing (growth) and balancing (stabilizing) loops that drive system behavior over time.

Read [cookbook/feedback-loops.md](./cookbook/feedback-loops.md)

### Technical Debt Analysis
Systematically identify, quantify, and prioritize technical debt. Make informed decisions about when to pay it down.

Read [cookbook/technical-debt.md](./cookbook/technical-debt.md)

## Core Principles

1. **Interconnections matter** - Components interact; changing one affects others
2. **Structure drives behavior** - System structure determines outcomes more than individual actors
3. **Delays cause oscillation** - Time lags create unexpected dynamics
4. **Feedback loops dominate** - Self-reinforcing and self-limiting cycles shape behavior
5. **Leverage points vary** - Small changes in the right place create big effects

## Systems Thinking Concepts

### Stocks and Flows
**Stocks**: Accumulations (inventory, cash, users, technical debt)
**Flows**: Rates of change (sales rate, churn rate, code commits)

### Feedback Loops
**Reinforcing (R)**: Amplifies change (growth or decline)
**Balancing (B)**: Resists change (stabilizes toward goal)

### Delays
Time between cause and effect. Often where systems "surprise" us.

### Leverage Points
Places where small interventions create large change (from least to most powerful):
12. Numbers (subsidies, taxes)
11. Buffer sizes
10. Stock-and-flow structures
9. Delays
8. Balancing feedback loops
7. Reinforcing feedback loops
6. Information flows
5. Rules (incentives, constraints)
4. Power to self-organize
3. System goals
2. Mindset/paradigm
1. Power to transcend paradigms

## When Systems Thinking Works Best

- Problems that keep recurring despite fixes
- Unintended consequences of interventions
- Complex situations with multiple stakeholders
- Dynamic systems that change over time
- When "obvious" solutions make things worse

## Warning Signs You Need This

- "We fixed it but the problem came back"
- "That solution created three new problems"
- "Everyone is optimizing their part but overall performance is declining"
- "We can't figure out why this keeps happening"
- Short-term fixes undermining long-term health
