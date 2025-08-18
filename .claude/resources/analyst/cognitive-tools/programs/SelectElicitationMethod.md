---
name: SelectElicitationMethod
purpose: Choose optimal elicitation technique based on context
version: 1.0
last_updated: 2025-08-18

# Decision Matrix for Elicitation Techniques

## Input Parameters
- current_context: string
- user_responses: list
- session_goals: list
- techniques_used: list

## Output Structure
- recommended_technique: string
- rationale: string
- alternative_options: list

## Elicitation Technique Decision Matrix

### Technique Options
1. Interview
2. Survey
3. Focus Group
4. Delphi Method
5. Observation
6. Prototyping
7. Workshop

### Selection Criteria
- Complexity of Information
- Depth of Insight Required
- Participant Availability
- Time Constraints
- Previous Techniques Used

## Decision Logic

### Scoring Mechanism
- Context Alignment: 0-5 points
- Goal Relevance: 0-5 points
- Participant Engagement: 0-5 points
- Information Depth: 0-5 points

### Technique Scoring Example
```python
def score_technique(technique, context):
    base_scores = {
        'Interview': {'complexity_handling': 4, 'depth': 5, 'time_investment': 3},
        'Survey': {'complexity_handling': 2, 'depth': 3, 'time_investment': 5},
        # ... other techniques
    }
    
    # Dynamic scoring based on context
    if 'technical_domain' in context:
        base_scores['Interview']['depth'] += 2
    
    return calculate_weighted_score(base_scores[technique])
```

## Molecular Protocols Integration
- Link with: `.../protocols/elicitation-workflow.md`
- Trigger sub-agents for complex scenarios
- Log technique selection for future learning

## Example Usage Scenario
```yaml
input:
  current_context: Enterprise software requirements gathering
  user_responses: []
  session_goals: 
    - Understand user pain points
    - Identify feature priorities
  techniques_used: ['initial survey']

output:
  recommended_technique: Focus Group
  rationale: >
    Initial survey provided broad insights. 
    Focus group will allow deep-dive discussion 
    with key stakeholders.
  alternative_options: 
    - Structured Interview
    - Prototyping Workshop
```

## Cognitive Decision Heuristics
1. Avoid repeating recently used techniques
2. Match technique to complexity of domain
3. Prioritize techniques that maximize participant engagement
4. Consider time and resource constraints

## Performance Tracking
- Record technique effectiveness
- Update scoring mechanism periodically
- Learn from past elicitation sessions

## Constraints and Limitations
- Not suitable for highly sensitive or confidential topics
- Requires skilled facilitation
- Dependent on participant availability and willingness

## Metaprotocol References
- `.../protocols/molecular/elicitation-flow.md`
- `.../protocols/atomic/technique-selection.yaml`
