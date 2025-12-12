# Technical Debt Analysis

Systematically identify, quantify, and prioritize technical debt. Make informed decisions about when and what to pay down.

## When to Use

- Velocity is declining despite team growth
- Simple changes take increasingly long
- New developers struggle to ramp up
- Bugs keep appearing in certain areas
- Planning tech debt paydown work
- Time: 1-3 hours for initial assessment

## What is Technical Debt?

Technical debt is the implied cost of choosing an expedient solution now over a better approach that would take longer.

**Like financial debt**:
- Principal: The original shortcut taken
- Interest: Ongoing cost of working around it
- Payment: Effort to fix it properly

**Unlike financial debt**:
- Interest rate isn't fixed (can compound rapidly)
- Principal can grow (debt attracts more debt)
- May never need to repay (if code isn't touched)

## Types of Technical Debt

### Deliberate vs. Inadvertent

| Type | Description | Example |
|------|-------------|---------|
| **Deliberate** | Conscious tradeoff for speed | "Ship now, refactor later" |
| **Inadvertent** | Didn't know better | Junior dev's first architecture |

### Prudent vs. Reckless

| Type | Description | Example |
|------|-------------|---------|
| **Prudent** | Know the cost, worth it | "We must ship by Friday" |
| **Reckless** | Didn't consider cost | "What's a design pattern?" |

### The Quadrant

```
                   Deliberate
                       │
    ┌──────────────────┼──────────────────┐
    │                  │                  │
    │  "We don't have  │  "We must ship   │
    │   time for       │   now and deal   │
    │   design"        │   with it later" │
    │                  │                  │
Reckless ─────────────┼──────────────Prudent
    │                  │                  │
    │  "What's         │  "Now we know    │
    │   layering?"     │   how we should  │
    │                  │   have done it"  │
    │                  │                  │
    └──────────────────┼──────────────────┘
                       │
                  Inadvertent
```

**Goal**: Be in the Prudent quadrants. Deliberate-Prudent is acceptable; Inadvertent-Prudent means learning (good).

## The Process

### Step 1: Inventory Debt (30-60 min)

List all known technical debt.

**Sources to check**:
- Team members' heads
- Code comments (TODO, FIXME, HACK)
- Architecture decision records
- Post-mortem documents
- Areas with high bug density
- Modules that take long to change

**For each item, capture**:
- Description: What is the debt?
- Location: Where in the codebase?
- Origin: When/why was it created?
- Known impact: What problems does it cause?

### Step 2: Categorize Debt (15 min)

Group debt by type:

| Category | Examples |
|----------|----------|
| **Code quality** | Duplication, complexity, poor naming |
| **Architecture** | Wrong patterns, tight coupling, scaling issues |
| **Testing** | Missing tests, flaky tests, slow tests |
| **Documentation** | Missing/outdated docs, tribal knowledge |
| **Infrastructure** | Outdated dependencies, manual processes |
| **Security** | Vulnerabilities, outdated libraries |

### Step 3: Assess Interest Rate (20 min)

How much ongoing cost does each debt item cause?

**Interest indicators**:
- Frequency of touching this code
- Developer time wasted working around it
- Bugs caused by this debt
- Onboarding time for new developers
- Incidents/outages related to it

**Interest rate scale**:
| Rate | Description | Urgency |
|------|-------------|---------|
| **High** | Daily pain, blocking work | Address soon |
| **Medium** | Weekly friction, slows features | Plan to address |
| **Low** | Rarely encountered | Address opportunistically |
| **Zero** | Code never touched | Leave it (for now) |

### Step 4: Estimate Payoff Effort (20 min)

How much work to properly fix each debt?

**T-shirt sizing**:
| Size | Effort | Example |
|------|--------|---------|
| **S** | < 1 day | Rename confusing variable, add missing test |
| **M** | 1-3 days | Refactor a module, update a dependency |
| **L** | 1-2 weeks | Redesign a component, major refactoring |
| **XL** | > 2 weeks | Architecture change, rewrite |

### Step 5: Calculate Priority (15 min)

**Priority = Interest Rate / Payoff Effort**

High interest + low effort = High priority
Low interest + high effort = Low priority

**Priority matrix**:

```
             High Interest
                  │
    ┌─────────────┼─────────────┐
    │             │             │
    │   DO NEXT   │   PLAN FOR  │
    │ (high ROI)  │  (but hard) │
    │             │             │
Low Effort ───────┼───────High Effort
    │             │             │
    │ QUICK WINS  │   PROBABLY  │
    │ (why not?)  │   IGNORE    │
    │             │             │
    └─────────────┼─────────────┘
                  │
             Low Interest
```

### Step 6: Create Paydown Plan (15 min)

**Strategies**:

| Strategy | When to Use |
|----------|-------------|
| **Continuous paydown** | Dedicate % of each sprint (e.g., 20%) |
| **Debt sprints** | Periodic focused cleanup (e.g., quarterly) |
| **Boy scout rule** | Improve code when you touch it |
| **Targeted campaigns** | Address specific debt category |

**Planning questions**:
- How much capacity to allocate to debt?
- Which items to address first?
- How to track progress?
- When to reassess?

## Technical Debt Template

```markdown
## Technical Debt Assessment: [System/Service]

**Date**:
**Assessed by**:
**Total items identified**: [N]

### Debt Inventory

| ID | Description | Location | Category | Interest | Effort | Priority |
|----|-------------|----------|----------|----------|--------|----------|
| TD-1 | | | | H/M/L | S/M/L/XL | |
| TD-2 | | | | | | |

### Category Summary

| Category | Count | High Priority |
|----------|-------|---------------|
| Code quality | | |
| Architecture | | |
| Testing | | |
| Infrastructure | | |
| Documentation | | |

### Top Priority Items

1. **[ID]: [Description]**
   - Interest: [Impact details]
   - Effort: [What's involved]
   - Payoff: [What improves when fixed]

2. **[ID]: [Description]**
   [Same structure]

3. **[ID]: [Description]**
   [Same structure]

### Paydown Plan

**Strategy**: [Continuous / Sprint / Campaign]
**Capacity allocation**: [X% of sprint / X days per quarter]

**Next sprint**:
- [ ] TD-X:
- [ ] TD-Y:

**Next quarter**:
- [ ] TD-Z:

### Metrics to Track
- [ ] Debt item count over time
- [ ] Velocity trends
- [ ] Bug density in debt areas
- [ ] Developer satisfaction
```

## Tips for Managing Technical Debt

### Make it visible
Track debt formally. What's not measured doesn't get managed.

### Accept that some debt is okay
Prudent debt enables speed. Zero debt often means over-engineering.

### Pay continuously
Deferred payment compounds. Regular small payments beat rare big ones.

### Fix debt where you work
Boy scout rule: Leave code better than you found it.

### Don't let perfect block good
80% fix now beats 100% fix never.

### Communicate in business terms
"This slows feature delivery by 30%" > "We have technical debt"

## Common Mistakes

- **Ignoring debt until crisis**: Compounds until it blocks all progress
- **Gold plating as debt paydown**: "While we're here, let's add..." isn't paydown
- **Paying wrong debt first**: Low-interest debt that's easy to fix isn't high priority
- **No visibility**: Team knows about debt but leadership doesn't
- **Debt amnesia**: Creating same debt repeatedly
- **All or nothing**: Either ignoring debt or trying to fix everything at once
- **Not tracking impact**: Paydown should improve something measurable
