# Systems Mapping

Visualize the components, relationships, and boundaries of a system. Make the structure visible so you can understand and change it.

## When to Use

- Complex situation with many interacting parts
- Need shared understanding across team
- Trying to identify where to intervene
- Understanding unintended consequences
- Onboarding others to a complex domain
- Time: 30-90 minutes

## Types of Systems Maps

| Type | Best For | Shows |
|------|----------|-------|
| **Concept Map** | General understanding | Entities and relationships |
| **Causal Loop Diagram** | Dynamic behavior | Feedback loops and causation |
| **Stock-Flow Diagram** | Quantitative analysis | Accumulations and rates |
| **Influence Diagram** | Decision support | Factors affecting outcomes |

## The Process

### Step 1: Define the System Boundary (10 min)

What's inside vs. outside the system you're analyzing?

**Questions to ask**:
- What's the purpose/function of this system?
- What are the key outcomes we care about?
- What's clearly inside the system?
- What's clearly outside (environment)?
- What's on the boundary (interacts both ways)?

**Example boundary**:
```
┌─────────────────────────────────────────────┐
│              SYSTEM BOUNDARY                 │
│                                             │
│    Product Team    Engineering    Support   │
│                                             │
│    Users    Features    Infrastructure      │
│                                             │
└─────────────────────────────────────────────┘
        │                    │
        ▼                    ▼
   Market Trends        Competitors
   (Environment)        (Environment)
```

### Step 2: Identify Key Elements (15 min)

List the main components of the system.

**Element types**:
- **Actors**: People, teams, organizations
- **Resources**: Money, time, inventory, information
- **Processes**: Activities that transform inputs to outputs
- **Artifacts**: Products, documents, deliverables
- **Concepts**: Goals, values, rules

**Brainstorm freely**, then group and filter:
- Which elements are essential?
- Which significantly affect outcomes?
- Which do we have data/insight about?

### Step 3: Map Relationships (20 min)

How do elements influence each other?

**Relationship types**:
- **Causal**: A causes/affects B
- **Correlational**: A and B move together
- **Enabling**: A makes B possible
- **Constraining**: A limits B
- **Information**: A informs B

**Arrow notation**:
```
A ──────► B      A directly causes B
A ───+──► B      A increases B (positive relationship)
A ───-──► B      A decreases B (negative relationship)
A ~~~►    B      A influences B (with delay)
A ◄─────► B      A and B influence each other
```

### Step 4: Identify Feedback Loops (15 min)

Find circular patterns where effects loop back to causes.

**Reinforcing loop** (R): Change amplifies (growth or decline)
```
     Sales ──+──► Revenue
        ▲              │
        │      +       │
        └───── Marketing Budget
```
More sales → more revenue → bigger marketing budget → more sales

**Balancing loop** (B): Change is resisted (stabilizes)
```
    Workload ──+──► Stress
        ▲              │
        │      -       │
        └───── Productivity
```
More workload → more stress → less productivity → (eventually) less workload taken on

### Step 5: Add Context and Boundaries (10 min)

Complete the map with:
- System boundary line
- External factors that influence the system
- Key delays (where effects take time)
- Leverage points (where intervention might help)

## Mapping Notation

### Simple concept map
```
┌─────────┐     affects      ┌─────────┐
│ Element │ ───────────────► │ Element │
│    A    │                  │    B    │
└─────────┘                  └─────────┘
```

### Causal loop diagram
```
        Variable A
            │
            │ + (same direction)
            ▼
        Variable B
            │
            │ - (opposite direction)
            ▼
        Variable C
            │
            └───────────► back to A (feedback)
```

### Stock-flow diagram
```
            Inflow                    Outflow
              │                          │
              ▼                          ▼
         ═════════                  ═════════
         ║ Stock ║ ◄─────────────► ║       ║
         ║(accum)║                  ║       ║
         ═════════                  ═════════
```

## Systems Map Template

```markdown
## Systems Map: [System Name]

**Purpose**: [What this map is meant to illuminate]
**Scope**: [Boundary of what's included]
**Created**: [Date]

### Key Elements

**Actors**:
-
-

**Resources/Stocks**:
-
-

**Processes/Flows**:
-
-

### Key Relationships

| From | To | Relationship | Notes |
|------|----|--------------| ------|
| | | (+/-/info) | |
| | | | |

### Feedback Loops Identified

**Loop 1: [Name]** (R/B)
[Description of the loop]
[Elements involved: A → B → C → A]

**Loop 2: [Name]** (R/B)
[Description]

### Key Delays
-
-

### Leverage Points
1. [Highest leverage]
2.
3.

### Diagram
[ASCII or link to visual diagram]

### Insights
-
-
```

## Example: SaaS Product System

```
                    ┌─────────────────────────────────────────┐
                    │                                         │
   Market ─────────►│    ┌─────────┐      ┌──────────┐      │
   Trends           │    │Marketing│──+──►│  Leads   │      │
                    │    └─────────┘      └────┬─────┘      │
                    │                          │            │
                    │                    +     ▼            │
                    │                    ┌──────────┐       │
                    │    ┌───────────────│  Users   │       │
                    │    │               └────┬─────┘       │
                    │    │                    │             │
                    │    │ +            +     ▼             │
                    │    │              ┌──────────┐        │
                    │    │              │ Revenue  │────────┼──► Profit
                    │    │              └────┬─────┘        │
                    │    │                   │              │
                    │    │             +     ▼              │
                    │    │         ┌────────────────┐       │
                    │    └─────────│   Dev Team     │       │
                    │              └───────┬────────┘       │
                    │                      │                │
                    │                +     ▼                │
                    │              ┌───────────┐            │
                    │              │ Features  │─────┐      │
                    │              └───────────┘     │      │
                    │                    ▲           │      │
                    │                    └─────+─────┘      │
                    │                  (R1: Growth Loop)    │
                    └─────────────────────────────────────────┘

Feedback Loops:
R1 (Growth): Users → Revenue → Dev Team → Features → Users
B1 (Churn): Users → Support Load → Response Time → Satisfaction → Churn
```

## Tips for Better Systems Maps

### Start messy, refine later
First draft should be a brain dump. Clean it up afterward.

### Focus on what matters
Include elements that significantly affect outcomes. Not every detail.

### Name your loops
Give feedback loops memorable names. "The Success Trap" is better than "Loop 3."

### Validate with others
Your mental model might be wrong. Check with people who know the system.

### Keep it usable
A map too complex to understand isn't useful. Simplify or create multiple views.

### Update over time
Systems change. Maps should too.

## Common Mistakes

- **Too much detail**: Maps become unusable wall art
- **Missing feedback loops**: Linear thinking misses system dynamics
- **Unclear boundaries**: Everything is connected, but you need a scope
- **Static thinking**: Maps should show dynamics, not just structure
- **Solo mapping**: Best maps come from multiple perspectives
- **One and done**: Maps need updating as understanding grows
