# Feedback Loop Analysis

Identify the circular patterns that drive system behavior over time. Understanding loops reveals why systems grow, stabilize, or collapse.

## When to Use

- System behavior is counterintuitive
- Interventions have unexpected results
- Growth is accelerating or stalling
- Need to predict system dynamics
- Looking for high-leverage interventions
- Time: 20-45 minutes

## Two Types of Feedback

### Reinforcing Loops (R)
**Amplify change** - deviation grows. Can drive growth OR decline.

```
     Revenue
        ▲
        │  +
        │
    Investment ◄───+ ─── Profit
        │
        │  +
        ▼
     Growth
```

"Success breeds success" - or "failure breeds failure"

**Characteristics**:
- Exponential behavior (accelerating change)
- "Virtuous cycles" or "vicious cycles"
- Unstable - small changes compound
- Eventually hit limits

**Examples**:
- Word of mouth: Users → Recommendations → New Users
- Compound interest: Principal → Interest → Principal
- Technical debt: Shortcuts → Complexity → More Shortcuts
- Network effects: Users → Value → More Users

### Balancing Loops (B)
**Resist change** - push toward equilibrium or goal.

```
     Actual State
          │
          │  comparison
          ▼
        Gap ──────► Action
          ▲              │
          │              │
          └─── - ────────┘
```

"Thermostat behavior" - system seeks a target

**Characteristics**:
- Oscillation or stabilization
- Goal-seeking behavior
- Homeostasis
- Can cause "fixes that fail"

**Examples**:
- Thermostat: Temperature → Gap from setpoint → Heating/Cooling
- Hunger: Food level → Hunger signal → Eating
- Hiring: Workload → Gap from capacity → Hiring
- Market: Price → Demand gap → Price adjustment

## The Process

### Step 1: Identify Variables (10 min)

List factors that change over time.

**Good variables**:
- Quantifiable (even roughly)
- Can increase or decrease
- Meaningful to the system

**Examples**:
- User count, revenue, cost
- Quality, morale, trust
- Inventory, backlog, debt
- Speed, efficiency, capacity

### Step 2: Map Causal Connections (15 min)

For each variable, ask: "What does this affect? What affects it?"

**Determine polarity**:
- **+ (Same direction)**: When A increases, B increases (all else equal)
- **- (Opposite direction)**: When A increases, B decreases

**Test with sentences**:
- "As [A] increases, [B] increases" → positive (+)
- "As [A] increases, [B] decreases" → negative (-)

### Step 3: Find the Loops (10 min)

Trace paths that return to their starting point.

**Determine loop type**:
- Count the negative (-) links in the loop
- **Even number of negatives** = Reinforcing (R)
- **Odd number of negatives** = Balancing (B)

**Example**:
```
A ──+──► B ──+──► C ──-──► A

Count: One negative (-)
Result: Balancing loop (B)
```

### Step 4: Name and Describe Loops (5 min)

Give each loop a memorable name that captures its behavior.

**Good loop names**:
- "The Growth Engine"
- "Death Spiral"
- "Quality Trap"
- "Success to the Successful"
- "Firefighting Syndrome"

Describe in one sentence what the loop does.

### Step 5: Identify Delays (5 min)

Where does cause-and-effect take time?

```
A ───//───► B  (delay marked with //)
```

Delays matter because:
- They cause overshoot and oscillation
- They make intervention timing critical
- They obscure cause-effect relationships

## Common Loop Patterns

### Growth Engine (R)
```
   Performance
        │
        │  +
        ▼
    Resources ◄───+ ─── Investment
        │
        │  +
        ▼
     Results
```
Success attracts investment, enabling more success.

### Limits to Growth (R + B)
```
     Growth                    Resource
        │                      Constraint
        │  +                       │
        ▼                          │  -
    Success ◄──────────────────────┘
        │
        │  + (R)
        ▼
    Effort
```
Growth eventually hits constraints that slow it.

### Shifting the Burden (B + B)
```
    Symptom                    Problem
        │                          │
        │  -                       │  -
        ▼                          ▼
   Quick Fix              Fundamental Fix
        │                          │
        │  +                       │  -
        ▼                          ▼
  Dependency ─────────────► Capability
```
Quick fixes undermine capability for real solutions.

### Tragedy of the Commons (R)
```
   Individual                Total
    Benefit ──+──► Activity ──+──► Resource
        ▲                          Depletion
        │                              │
        └──────────── - ───────────────┘
```
Individual incentives deplete shared resources.

### Escalation (R + R)
```
   Party A's                 Party B's
    Actions ──+──► Threat ──+──► Actions
        ▲                          │
        │                          │
        └──────────+ ──────────────┘
```
Each party's response triggers the other's escalation.

## Loop Analysis Template

```markdown
## Feedback Loop Analysis: [System Name]

**Date**:
**Analyzed by**:

### System Variables
| Variable | Description | Increases when... |
|----------|-------------|-------------------|
| | | |
| | | |

### Causal Links
| From | To | Polarity | Delay? | Notes |
|------|----| -------- |--------|-------|
| | | +/- | Y/N | |
| | | | | |

### Identified Loops

**Loop 1: [Name]**
- Type: R / B
- Path: A → B → C → A
- Behavior: [What this loop does]
- Strength: Strong / Moderate / Weak

**Loop 2: [Name]**
[Same structure]

### Loop Interactions
[How do loops interact? Which dominates when?]

### Key Delays
| Location | Duration | Impact |
|----------|----------|--------|
| | | |

### Leverage Points
Based on this analysis:
1. [Highest leverage intervention]
2.
3.

### Predicted Behavior
If current loops continue: [What happens?]
If we intervene at [point]: [Expected result]
```

## Tips for Better Loop Analysis

### Look for the counterintuitive
The most valuable loops are ones that explain surprising behavior.

### Watch for delays
Delays cause most "unexpected" outcomes. Map them explicitly.

### Find the dominant loop
Multiple loops exist; which one is currently winning?

### Check loop shifting
Different loops dominate at different times/scales.

### Validate with data
Loop diagrams are hypotheses. Test with real data.

## Common Mistakes

- **Confusing correlation and causation**: A and B move together ≠ A causes B
- **Missing polarity testing**: Always verify +/- with "as X increases..." test
- **Ignoring delays**: Delays are where systems fool us
- **Static thinking**: Loops are dynamic; draw behavior over time
- **Single loop focus**: Systems have multiple interacting loops
- **Analysis paralysis**: Don't map everything; focus on what drives outcomes
