---
name: AssessResearchNeed
type: cognitive-tool
version: 1.0
description: Research need assessment protocol for determining when to delegate to research sub-agents
triggers:
  - when_knowledge_gap_identified
  - when_unknown_information_required
  - when_market_analysis_needed
  - when_technical_feasibility_uncertain
  - when_competitive_landscape_unclear
---

# Assess Research Need

## Input Schema

```yaml
context:
  trigger_type: enum
    - Unknown Market Info
    - Technical Feasibility
    - Competitive Landscape
    - User Behavior
    - Emerging Tech
  current_knowledge: string  # What's known
  knowledge_gap: string  # What's missing
  impact: enum  # Critical, High, Medium, Low
  urgency: enum  # Immediate, Soon, Eventually, Nice-to-have
  research_scope: enum  # Narrow, Focused, Broad, Comprehensive
  stakeholder_preferences: string  # Research style preferences
  prior_research: string  # Previous research conducted
  industry_context: string  # Market/domain context
```

## Output Schema

```yaml
research_needed: boolean
priority: enum  # Critical, High, Medium, Low
research_type: array  # Types of research recommended
delegation_recommended: boolean
suggested_sub_agent: string  # If delegation recommended
research_scope: string
estimated_timeline: string  # Days or weeks
estimated_effort: string  # High/Medium/Low
confidence_in_assessment: float  # 0.0-1.0
rationale: string
next_steps: array  # Action items
```

## Research Triggers

### Unknown Market Information
- **Indicators**:
  - Market size uncertainty
  - Pricing strategy questions
  - Market segment feasibility
  - Customer acquisition channel uncertainty
  - Revenue model validation needed

- **Assessment Questions**:
  - How precise must the market information be?
  - Are there existing market reports?
  - Is this for business case or direction setting?

- **Recommended Research Type**: Market Analysis, Customer Research, Competitive Analysis

### Technical Feasibility
- **Indicators**:
  - Technology selection uncertain
  - Architecture viability unknown
  - Performance requirements unclear
  - Integration complexity unclear
  - Scalability concerns

- **Assessment Questions**:
  - Is this a proof-of-concept or production decision?
  - Are there existing technical assessments?
  - How critical is the decision?

- **Recommended Research Type**: Technical Investigation, Prototype Testing, Architecture Review

### Competitive Landscape
- **Indicators**:
  - Competitor capability uncertainty
  - Market position unclear
  - Feature parity questions
  - Pricing comparison needed
  - Strategic positioning uncertain

- **Assessment Questions**:
  - How competitive is the market?
  - What's our differentiation strategy?
  - How frequently does landscape change?

- **Recommended Research Type**: Competitive Analysis, Market Intelligence, Benchmarking

### User Behavior
- **Indicators**:
  - User needs unclear
  - Usage pattern uncertainty
  - Feature prioritization needed
  - User workflow understanding incomplete
  - Adoption risk unknown

- **Assessment Questions**:
  - Do we have user access?
  - How complex is the behavior?
  - How critical is user understanding?

- **Recommended Research Type**: User Research, Behavioral Analysis, Ethnographic Study

### Emerging Technology
- **Indicators**:
  - Technology maturity uncertain
  - Adoption readiness unknown
  - Integration patterns unclear
  - Industry standards undefined
  - Risk assessment needed

- **Assessment Questions**:
  - How new is the technology?
  - Is there existing documentation?
  - How applicable is it to our domain?

- **Recommended Research Type**: Technology Analysis, Industry Research, Expert Consultation

## Assessment Criteria Scoring

### Research Necessity Matrix

| Trigger Type | Critical | High | Medium | Low |
|--------------|----------|------|--------|-----|
| Unknown Market Info | Market size, viability | Pricing, segments | Trends, forecasts | Minor metrics |
| Technical Feasibility | Architecture, scalability | Technology selection | Integration approach | Performance tuning |
| Competitive Landscape | Strategic positioning | Feature parity | Pricing comparison | Minor competitors |
| User Behavior | Core workflows | Feature priority | Usage patterns | Minor preferences |
| Emerging Tech | Adoption viability | Standards clarity | Integration approach | Technical details |

### Impact Assessment
- **Critical**: Affects core business decision, go/no-go criteria, strategic direction
- **High**: Affects feature set, timeline, resource allocation
- **Medium**: Affects optimization, secondary decisions
- **Low**: Nice-to-have, enhancement decisions

### Urgency Mapping
- **Immediate**: Needed within 2-3 days
- **Soon**: Needed within 1-2 weeks
- **Eventually**: Needed within 3-4 weeks
- **Nice-to-have**: No strict deadline

### Knowledge Gap Severity Scale

| Level | Characteristics | Research Effort |
|-------|-----------------|-----------------|
| 1 - Minor gap | Clear scope, easily researchable | Low |
| 2 - Focused gap | Defined scope, moderate research | Medium |
| 3 - Significant gap | Broader scope, deeper research needed | High |
| 4 - Critical gap | Foundational gap, comprehensive research needed | Very High |

## Scoring Mechanism

### Research Priority Calculation

```
priority_score = (impact_weight × 0.35) +
                 (urgency_weight × 0.30) +
                 (knowledge_gap_severity × 0.20) +
                 (decision_criticality × 0.15)

Impact weights: Critical=4, High=3, Medium=2, Low=1
Urgency weights: Immediate=4, Soon=3, Eventually=2, Nice-to-have=1
Gap severity: 0.0-1.0 scale
Decision criticality: 0.0-1.0 scale
```

### Delegation Criteria

Research should be delegated to sub-agents when:

1. **Complexity Score > 6/10**
   - Multiple research methodologies needed
   - Significant data gathering required
   - Analysis requiring specialized expertise

2. **Scope Score > 6/10**
   - Broad research scope
   - Multiple stakeholder interviews needed
   - Extensive source gathering

3. **Timeline Allows**
   - Urgency not "Immediate"
   - Timeline at least 5+ days for focused research
   - Timeline at least 2+ weeks for comprehensive research

4. **Resource Availability**
   - Analyst has capacity to oversee
   - Research sub-agent available
   - Budget allocated

## Decision Logic

### Step 1: Research Necessity Check

```
IF knowledge_gap is well-defined AND
   current_knowledge is incomplete AND
   gap impacts decision THEN
  research_needed = true
ELSE
  research_needed = false
```

### Step 2: Priority Assessment

```
priority_score = weighted calculation (see Scoring Mechanism)

IF priority_score >= 3.0 THEN priority = High/Critical
ELSE IF priority_score >= 2.0 THEN priority = Medium
ELSE priority = Low
```

### Step 3: Research Type Determination

**Unknown Market Info** → Market Analysis + Customer Research
- If pricing uncertainty: Add Pricing Strategy Research
- If segment uncertainty: Add Customer Segmentation Research
- If CAC uncertainty: Add Go-to-Market Research

**Technical Feasibility** → Technology Assessment + Feasibility Study
- If architecture uncertain: Add Architecture Review
- If scalability concerns: Add Performance Analysis
- If integration complex: Add Integration Study

**Competitive Landscape** → Competitive Analysis + Market Intelligence
- If positioning unclear: Add Strategic Analysis
- If feature parity needed: Add Feature Benchmarking
- If pricing uncertain: Add Pricing Analysis

**User Behavior** → User Research + Behavioral Analysis
- If needs unclear: Add User Interviews
- If workflows unknown: Add Workflow Mapping
- If adoption risky: Add Adoption Analysis

**Emerging Tech** → Technology Research + Industry Analysis
- If standards undefined: Add Standards Research
- If maturity unclear: Add Maturity Assessment
- If adoption risky: Add Risk Analysis

### Step 4: Delegation Decision

```
complexity_score = assess(research_scope, methodology_count, analysis_depth)
timeline_buffer = available_time - minimum_required_time

IF complexity_score > 6 AND timeline_buffer > 3 days THEN
  delegation_recommended = true
  estimated_timeline = 5-21 days (depending on scope)
ELSE
  delegation_recommended = false
  handle_inline_with_analyst = true
```

### Step 5: Sub-Agent Selection

- **Market Research Sub-Agent**: Market Info, Competitive Landscape triggers
- **Technical Research Sub-Agent**: Technical Feasibility, Emerging Tech triggers
- **User Research Sub-Agent**: User Behavior, feature validation triggers
- **General Research Sub-Agent**: Cross-domain research, secondary triggers

## Implementation Steps

1. Identify research trigger based on context
2. Assess knowledge gap specificity and severity
3. Evaluate impact and urgency
4. Calculate priority score
5. Determine research types needed
6. Assess complexity and timeline feasibility
7. Make delegation decision
8. Select appropriate research methodologies
9. Provide timeline and effort estimates
10. Set success criteria for research completion
