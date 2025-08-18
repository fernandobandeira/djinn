---
name: AssessResearchNeed
purpose: Determine when to delegate to research sub-agents
version: 1.0
last_updated: 2025-08-18

# Research Need Assessment Protocol

## Input Parameters
- discussion_context: string
- information_gaps: list
- user_questions: list

## Output Structure
- research_needed: boolean
- sub_agent: string
- research_scope: string

## Research Delegation Decision Matrix

### Research Triggers
1. Unknown Market Information
2. Technical Feasibility Questions
3. Competitive Landscape Analysis
4. User Behavior Insights
5. Emerging Technology Assessment

### Assessment Criteria
- Criticality of Information
- Complexity of Research Required
- Potential Impact on Decision
- Time to Resolve

## Decision Logic

### Scoring Mechanism
- Information Significance: 0-5 points
- Complexity of Research: 0-5 points
- Potential Decision Impact: 0-5 points

### Research Delegation Example
```python
def assess_research_need(context, information_gaps):
    research_score = calculate_research_score(context, information_gaps)
    
    if research_score > THRESHOLD:
        return {
            'research_needed': True,
            'sub_agent': select_optimal_research_agent(context),
            'research_scope': define_research_boundaries()
        }
    return {'research_needed': False}
```

## Molecular Protocols Integration
- Link with: `.../protocols/research-workflow.md`
- Trigger specialized research sub-agents
- Log research requests for knowledge base

## Example Usage Scenario
```yaml
input:
  discussion_context: Enterprise AI adoption strategy
  information_gaps: 
    - Current AI vendor landscape
    - Implementation cost ranges
  user_questions: 
    - What are top AI vendors for enterprise?

output:
  research_needed: true
  sub_agent: market_research
  research_scope: >
    Comprehensive analysis of enterprise AI vendors, 
    focusing on implementation costs, feature comparison, 
    and market penetration
```

## Cognitive Decision Heuristics
1. Prioritize research for high-impact decisions
2. Consider cost of research vs. potential value
3. Leverage existing knowledge base first
4. Set clear research boundaries

## Performance Tracking
- Record research request outcomes
- Measure time and quality of research
- Update research delegation strategies

## Constraints and Limitations
- Research time and resource constraints
- Potential bias in research selection
- Depth of available information

## Metaprotocol References
- `.../protocols/molecular/research-delegation.md`
- `.../protocols/atomic/research-triggers.yaml`
