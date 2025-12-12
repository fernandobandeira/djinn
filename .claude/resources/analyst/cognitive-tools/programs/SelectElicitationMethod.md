---
name: SelectElicitationMethod
type: cognitive-tool
version: 1.0
description: Decision matrix for selecting optimal elicitation technique based on context
triggers:
  - when_gathering_requirements
  - when_stakeholder_input_needed
  - when_feature_definition_required
---

# Select Elicitation Method

## Input Schema

```yaml
context:
  stakeholder_type: string  # Technical, Business, End-User, Mixed
  urgency: enum  # High, Medium, Low
  stakeholder_availability: enum  # Full, Partial, Limited
  domain_complexity: enum  # Simple, Moderate, Complex
  geographic_distribution: enum  # Collocated, Distributed, Global
  budget_constraints: enum  # None, Moderate, Tight
  timeline: integer  # Days available
  team_size: integer  # Number of stakeholders
  prior_context: string  # What's already known
  goal: string  # What information is needed
```

## Output Schema

```yaml
selected_technique: string
rationale: string
estimated_effort:
  preparation: integer  # Hours
  execution: integer  # Hours
  analysis: integer  # Hours
alternatives: array  # List of 2-3 alternatives
confidence_score: float  # 0.0-1.0
```

## Elicitation Techniques

### Interview
- **Best for**: Direct stakeholder engagement, deep exploration
- **Cost**: Medium-High
- **Timeline**: Medium (1-2 weeks)
- **Ideal team size**: 1-3 stakeholders
- **Characteristics**:
  - One-on-one or small group
  - Flexible, exploratory
  - Real-time clarification
  - High information density

### Survey
- **Best for**: Broad input gathering, preference validation
- **Cost**: Low-Medium
- **Timeline**: Short-Medium (1 week)
- **Ideal team size**: 5+ stakeholders
- **Characteristics**:
  - Structured questionnaire
  - Scalable to large groups
  - Asynchronous
  - Statistical analysis possible

### Focus Group
- **Best for**: Collective exploration, idea refinement
- **Cost**: Medium
- **Timeline**: Medium (2-3 weeks)
- **Ideal team size**: 6-12 stakeholders
- **Characteristics**:
  - Group dynamics
  - Creative exploration
  - Consensus building
  - Interactive discussion

### Delphi Method
- **Best for**: Expert consensus, prediction, complex estimation
- **Cost**: High
- **Timeline**: Long (3-4 weeks)
- **Ideal team size**: 5-15 expert stakeholders
- **Characteristics**:
  - Multiple rounds
  - Structured feedback
  - Anonymity
  - Convergence toward consensus

### Observation
- **Best for**: Current behavior understanding, workflow mapping
- **Cost**: Medium
- **Timeline**: Medium (2-4 weeks)
- **Ideal team size**: N/A
- **Characteristics**:
  - Field study
  - Natural context
  - Behavior focused
  - Minimal disruption

### Prototyping
- **Best for**: Concrete feedback, refinement, validation
- **Cost**: High (requires development)
- **Timeline**: Long (3-6 weeks)
- **Ideal team size**: 3-5 stakeholders
- **Characteristics**:
  - Tangible artifact
  - Experiential feedback
  - Iterative refinement
  - Visualization of possibilities

### Workshop
- **Best for**: Collaborative design, alignment, rapid ideation
- **Cost**: High
- **Timeline**: Short-Medium (1-2 days)
- **Ideal team size**: 8-20 stakeholders
- **Characteristics**:
  - Intensive, collocated
  - Cross-functional
  - Structured activities
  - Immediate consensus building

## Selection Criteria Scoring

### Stakeholder Type Alignment
| Technique | Technical | Business | End-User | Mixed |
|-----------|-----------|----------|----------|-------|
| Interview | 4 | 5 | 4 | 4 |
| Survey | 3 | 4 | 5 | 3 |
| Focus Group | 3 | 4 | 4 | 5 |
| Delphi | 5 | 4 | 2 | 3 |
| Observation | 5 | 2 | 5 | 3 |
| Prototyping | 5 | 4 | 5 | 4 |
| Workshop | 4 | 5 | 3 | 5 |

### Timeline Feasibility
| Technique | < 1 week | 1-2 weeks | 2-4 weeks | 4+ weeks |
|-----------|----------|-----------|-----------|----------|
| Interview | 3 | 5 | 4 | 3 |
| Survey | 5 | 4 | 3 | 2 |
| Focus Group | 2 | 4 | 5 | 3 |
| Delphi | 1 | 2 | 4 | 5 |
| Observation | 1 | 3 | 5 | 4 |
| Prototyping | 1 | 2 | 4 | 5 |
| Workshop | 5 | 4 | 2 | 1 |

### Availability Requirements
| Technique | Full | Partial | Limited |
|-----------|------|---------|---------|
| Interview | 5 | 4 | 3 |
| Survey | 5 | 5 | 4 |
| Focus Group | 4 | 3 | 2 |
| Delphi | 3 | 4 | 5 |
| Observation | 3 | 4 | 4 |
| Prototyping | 4 | 3 | 2 |
| Workshop | 2 | 3 | 1 |

### Cost-Effectiveness
| Technique | Low | Medium | High |
|-----------|-----|--------|------|
| Interview | 3 | 5 | 3 |
| Survey | 5 | 4 | 2 |
| Focus Group | 3 | 5 | 3 |
| Delphi | 2 | 3 | 5 |
| Observation | 3 | 5 | 3 |
| Prototyping | 1 | 3 | 5 |
| Workshop | 2 | 4 | 5 |

## Decision Logic

### Primary Selection Algorithm

1. **Timeline Constraint Check**
   - If timeline < 5 days: Prefer Survey, Workshop
   - If timeline 5-14 days: Prefer Interview, Survey, Focus Group, Workshop
   - If timeline 2-4 weeks: All techniques viable
   - If timeline 4+ weeks: Can include Delphi, Observation, Prototyping

2. **Team Size Filter**
   - If 1-3 stakeholders: Interview preferred
   - If 3-6 stakeholders: Interview, Prototyping viable
   - If 6-12 stakeholders: Focus Group, Workshop
   - If 12+ stakeholders: Survey, Delphi

3. **Availability Assessment**
   - Full availability: Interview, Focus Group, Workshop, Prototyping
   - Partial availability: Survey, Interview, Delphi, Observation
   - Limited availability: Survey, Delphi

4. **Goal Type Matching**
   - Deep exploration: Interview, Workshop
   - Broad validation: Survey, Focus Group
   - Expert consensus: Delphi
   - Behavior understanding: Observation
   - Tangible feedback: Prototyping

5. **Context-Specific Rules**
   - If domain_complexity = Complex AND stakeholder_type = Mixed: Focus Group or Workshop
   - If geographic_distribution = Global: Survey, Delphi (avoid Interview, Workshop)
   - If budget_constraints = Tight: Survey, Interview
   - If urgency = High: Survey, Workshop, Interview

### Weighting Formula

```
score = (timeline_score × 0.25) +
        (team_size_score × 0.20) +
        (availability_score × 0.20) +
        (goal_alignment_score × 0.20) +
        (cost_effectiveness_score × 0.15)
```

## Implementation Steps

1. Collect input context from stakeholder or project requirements
2. Apply primary selection algorithm
3. Score all viable techniques using weighting formula
4. Rank techniques by composite score
5. Select highest-scoring technique
6. Identify 2-3 alternatives
7. Provide confidence score based on context clarity
8. Calculate effort estimates based on team size and timeline
