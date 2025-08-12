# Elicitation Methods Reference

## Core Universal Methods

### Expand or Contract for Audience
**Purpose**: Adjust content depth and detail for specific audiences
**Application**:
- Identify target audience expertise level
- Add technical details for experts
- Simplify and add examples for beginners
- Adjust terminology and jargon
- Include or remove implementation details

**Questions to Ask**:
- Who will consume this content?
- What's their technical level?
- What context do they have?
- What decisions will they make?

### Critique and Refine
**Purpose**: Systematic quality improvement
**Application**:
- Review for completeness
- Check logical flow
- Identify weak arguments
- Find missing evidence
- Strengthen conclusions
- Improve clarity

**Critique Framework**:
1. Is it complete?
2. Is it accurate?
3. Is it clear?
4. Is it actionable?
5. Is it well-supported?

### Identify Potential Risks
**Purpose**: Surface and address hidden risks
**Application**:
- Technical risks (performance, security, scalability)
- Business risks (cost, timeline, resources)
- User risks (adoption, usability, satisfaction)
- Integration risks (compatibility, dependencies)
- Operational risks (maintenance, support)

**Risk Assessment Matrix**:
| Risk | Probability | Impact | Mitigation |
|------|------------|--------|------------|
| [Risk] | H/M/L | H/M/L | [Strategy] |

### Assess Alignment with Goals
**Purpose**: Ensure work supports objectives
**Application**:
- Map to stated goals
- Check success criteria
- Verify value delivery
- Identify misalignments
- Suggest corrections

**Alignment Check**:
- Original goal: [State goal]
- Current output: [How it aligns]
- Gaps: [What's missing]
- Corrections: [Needed changes]

## Technical Analysis Methods

### Tree of Thoughts (ToT)
**Purpose**: Explore multiple solution paths systematically
**Application**:
1. Generate multiple approaches (3-5)
2. Develop each branch independently
3. Evaluate pros/cons of each
4. Combine best elements
5. Document decision rationale

**Tree Structure**:
```
Problem
├── Approach A
│   ├── Pros
│   ├── Cons
│   └── Viability
├── Approach B
│   ├── Pros
│   ├── Cons
│   └── Viability
└── Hybrid Solution
```

### First Principles Decomposition
**Purpose**: Strip away assumptions to find fundamental truths
**Application**:
1. List all assumptions
2. Question each assumption
3. Identify fundamental truths
4. Rebuild from basics
5. Find simpler solution

**Process**:
- Assumption: [State it]
- Challenge: [Why might it be false?]
- Fundamental: [What must be true?]
- Rebuild: [New approach]

### Chain of Density
**Purpose**: Optimize information density
**Application**:
- Start with verbose explanation
- Progressively compress
- Maintain critical information
- Balance detail and brevity
- Layer information appropriately

**Density Levels**:
1. Executive: 1 paragraph
2. Summary: 1 page
3. Detailed: Full section
4. Technical: Complete spec

### Technical Debt Analysis
**Purpose**: Identify and plan for technical compromises
**Application**:
- Identify shortcuts taken
- Assess long-term impact
- Calculate interest (ongoing cost)
- Prioritize repayment
- Document trade-offs

**Debt Categories**:
- Design debt
- Code debt
- Testing debt
- Documentation debt
- Infrastructure debt

## Strategic Analysis Methods

### Red Team vs Blue Team
**Purpose**: Stress-test through adversarial analysis
**Application**:
- **Red Team** (Attack):
  - Find vulnerabilities
  - Challenge assumptions
  - Exploit weaknesses
  - Break the system
  
- **Blue Team** (Defend):
  - Justify decisions
  - Strengthen weak points
  - Add safeguards
  - Prove resilience

### Pre-mortem Analysis
**Purpose**: Prevent failure by imagining it
**Application**:
1. Imagine total failure
2. Work backward to causes
3. Identify warning signs
4. Build prevention measures
5. Create contingency plans

**Pre-mortem Questions**:
- "It's [future date] and this failed. Why?"
- "What early warning signs did we miss?"
- "What could we have done differently?"

### SWOT Deep Dive
**Purpose**: Comprehensive strategic assessment
**Application**:
- **Strengths**: What advantages do we have?
- **Weaknesses**: What limitations exist?
- **Opportunities**: What could we leverage?
- **Threats**: What could derail us?

**SWOT Actions**:
- Leverage strengths
- Mitigate weaknesses
- Capture opportunities
- Defend against threats

### Porter's Five Forces
**Purpose**: Analyze competitive dynamics
**Application**:
1. **Competitive Rivalry**: Direct competition intensity
2. **Supplier Power**: Dependency on suppliers
3. **Buyer Power**: Customer negotiation leverage
4. **Substitute Threat**: Alternative solutions
5. **New Entry Threat**: Barrier to entry analysis

## Creative Enhancement Methods

### Innovation Tournament
**Purpose**: Generate and evaluate competing solutions
**Application**:
1. Generate 3-5 distinct solutions
2. Define evaluation criteria
3. Score each solution
4. Combine winning elements
5. Create hybrid approach

**Scoring Matrix**:
| Solution | Feasibility | Impact | Cost | Innovation | Total |
|----------|------------|--------|------|------------|-------|
| A | | | | | |
| B | | | | | |

### What If Scenarios
**Purpose**: Explore edge cases and possibilities
**Application**:
- What if scale 10x?
- What if constraints removed?
- What if different context?
- What if failure occurs?
- What if wildly successful?

**Scenario Planning**:
- Scenario: [Description]
- Implications: [What changes]
- Preparations: [What's needed]
- Opportunities: [What's possible]

### Analogical Extension
**Purpose**: Learn from similar solved problems
**Application**:
1. Find analogous situation
2. Map solution pattern
3. Identify principles
4. Adapt to current context
5. Test applicability

**Analogy Mapping**:
- Source domain: [Where solution exists]
- Target domain: [Current problem]
- Mapping: [How they connect]
- Adaptation: [How to apply]

### Constraint Removal
**Purpose**: Find flexibility by challenging limits
**Application**:
1. List all constraints
2. Remove one temporarily
3. Explore new solutions
4. Gradually reintroduce
5. Find compromise position

**Constraint Analysis**:
- Constraint: [What limits us]
- If removed: [What becomes possible]
- Partial relaxation: [Middle ground]
- Workaround: [Alternative approach]

## User-Centered Methods

### Stakeholder Roundtable
**Purpose**: Consider all perspectives
**Application**:
- List all stakeholders
- Define their interests
- Explore each viewpoint
- Identify conflicts
- Find balanced solution

**Stakeholder Matrix**:
| Stakeholder | Needs | Concerns | Influence | Priority |
|-------------|-------|----------|-----------|----------|

### Journey Mapping
**Purpose**: Understand end-to-end experience
**Application**:
1. Map current journey
2. Identify pain points
3. Find improvement opportunities
4. Design ideal journey
5. Plan transition

**Journey Stages**:
- Awareness
- Consideration
- Decision
- Onboarding
- Usage
- Advocacy

### Persona Analysis
**Purpose**: Ensure solution fits users
**Application**:
- Create detailed personas
- Test solution against each
- Identify gaps in coverage
- Adjust for inclusivity
- Validate assumptions

**Persona Template**:
- Name: [Persona identifier]
- Goals: [What they want]
- Frustrations: [Pain points]
- Needs: [Requirements]
- Scenario: [How they use solution]

### Jobs-to-be-Done
**Purpose**: Focus on user objectives
**Application**:
- Identify functional job
- Understand emotional job
- Consider social job
- Measure success criteria
- Optimize for job completion

**JTBD Framework**:
- When [situation]
- I want to [motivation]
- So I can [outcome]

## Analytical Methods

### Five Whys Deep Dive
**Purpose**: Reach root cause through questioning
**Application**:
1. State the problem
2. Ask why it occurs
3. For each answer, ask why again
4. Continue at least 5 times
5. Address root cause

**Example Chain**:
- Problem: Users abandon signup
- Why? Form is too long
- Why? We ask for too much info
- Why? We think we need it all
- Why? We haven't tested what's essential
- Why? We assumed rather than validated
- Root: Lack of user research

### Morphological Analysis
**Purpose**: Explore systematic combinations
**Application**:
1. Identify key parameters
2. List options for each
3. Create combination matrix
4. Explore unusual combinations
5. Evaluate feasibility

**Parameter Grid**:
| P1 Options | P2 Options | P3 Options | Result |
|------------|------------|------------|--------|

### Systems Thinking
**Purpose**: Understand interconnections and dynamics
**Application**:
- Map system components
- Identify relationships
- Find feedback loops
- Locate leverage points
- Predict ripple effects

**System Elements**:
- Inputs
- Processes
- Outputs
- Feedback loops
- Control mechanisms

### Meta-Analysis
**Purpose**: Analyze the analysis itself
**Application**:
- Question methodology
- Examine biases
- Check completeness
- Validate conclusions
- Consider alternatives

**Meta Questions**:
- Is our approach correct?
- What are we not seeing?
- What biases exist?
- What would others do?
- How confident are we?