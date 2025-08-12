# Create Deep Research Prompt Task

## Purpose
Generate comprehensive research prompts that define clear objectives, methodologies, and expected deliverables for systematic investigation of complex topics.

## Research Type Selection

### Present Research Focus Options
Ask the user to select their research focus (1-9):

1. **Product Validation Research**
   - Validate product hypotheses and market fit
   - Test assumptions about user needs
   - Assess technical and business feasibility
   - Identify risks and mitigation strategies

2. **Market Opportunity Research**
   - Analyze market size and growth potential
   - Identify segments and dynamics
   - Assess market entry strategies
   - Evaluate timing and readiness

3. **User & Customer Research**
   - Deep dive into personas and behaviors
   - Understand jobs-to-be-done
   - Map customer journeys
   - Analyze willingness to pay

4. **Competitive Intelligence Research**
   - Detailed competitor analysis
   - Feature and capability comparisons
   - Business model analysis
   - Identify competitive advantages

5. **Technology & Innovation Research**
   - Assess technology trends
   - Evaluate technical approaches
   - Identify emerging technologies
   - Analyze build vs buy vs partner

6. **Industry & Ecosystem Research**
   - Map value chains and dynamics
   - Identify key players
   - Analyze regulatory factors
   - Understand partnership opportunities

7. **Strategic Options Research**
   - Evaluate strategic directions
   - Assess business model alternatives
   - Analyze go-to-market strategies
   - Consider expansion paths

8. **Risk & Feasibility Research**
   - Identify and assess risk factors
   - Evaluate implementation challenges
   - Analyze resource requirements
   - Consider regulatory implications

9. **Custom Research Focus**
   - User-defined objectives
   - Specialized domain investigation

## Input Processing

### Context Gathering
Based on available inputs:

**If Project Brief provided:**
- Extract key concepts and goals
- Identify target users
- Note constraints and preferences
- Highlight uncertainties

**If Brainstorming Results provided:**
- Synthesize main themes
- Identify areas needing validation
- Extract hypotheses to test
- Note creative directions

**If Market Research provided:**
- Build on opportunities
- Deepen specific insights
- Validate initial findings
- Explore adjacent possibilities

**If Starting Fresh:**
- Gather essential context
- Define problem space
- Clarify objectives
- Establish success criteria

## Research Prompt Development

### Collaborate with user to develop:

#### A. Research Objectives
- Primary research goal
- Key decisions to inform
- Success criteria
- Constraints and boundaries

#### B. Research Questions

**Primary Questions (Must Answer):**
- Central questions requiring answers
- Priority ranking
- Dependencies between questions

**Secondary Questions (Nice to Have):**
- Context-building questions
- Additional insights
- Future considerations

#### C. Research Methodology

**Data Collection:**
- Secondary research sources
- Primary research approaches
- Data quality requirements
- Source credibility criteria

**Analysis Frameworks:**
- Specific frameworks to apply
- Comparison criteria
- Evaluation methodologies
- Synthesis approaches

#### D. Output Requirements

**Format Specifications:**
- Executive summary requirements
- Detailed findings structure
- Visual presentations
- Supporting documentation

**Key Deliverables:**
- Must-have sections
- Decision-support elements
- Action-oriented recommendations
- Risk documentation

## Research Prompt Template

```markdown
# Deep Research Prompt: [Topic]

## Research Objective
[Clear statement of what this research aims to achieve]

## Background Context
[Relevant information from inputs]

## Research Questions

### Primary Questions (Must Answer)
1. [Specific, actionable question]
2. [Specific, actionable question]
3. [Specific, actionable question]

### Secondary Questions (Nice to Have)
1. [Supporting question]
2. [Supporting question]

## Research Methodology

### Information Sources
- [Specific source types and priorities]
- [Credibility requirements]

### Analysis Frameworks
- [Framework 1]: [How to apply]
- [Framework 2]: [How to apply]

### Data Requirements
- Quality: [Standards]
- Recency: [Requirements]
- Credibility: [Criteria]

## Expected Deliverables

### Executive Summary (1-2 pages)
- Key findings and insights
- Critical implications
- Recommended actions
- Risk assessment

### Detailed Analysis
[Sections based on research type]

### Supporting Materials
- Data tables
- Comparison matrices
- Source documentation
- Visual diagrams

## Success Criteria
[How to evaluate if research achieved objectives]

## Timeline and Priority
- Phase 1: [What and when]
- Phase 2: [What and when]
- Critical deadlines: [If any]

## Integration Points
- How findings feed into next phases
- Review stakeholders
- Validation approach
- Future research needs
```

## Review and Refinement Process

### 1. Present Complete Prompt
- Show full research prompt
- Explain key elements
- Highlight assumptions

### 2. Gather Feedback
Ask user:
- Are objectives clear and correct?
- Do questions address all concerns?
- Is scope appropriate?
- Are output requirements sufficient?

### 3. Refine as Needed
- Incorporate feedback
- Adjust scope or focus
- Add missing elements
- Clarify ambiguities

## Execution Guidance

### Usage Options:
1. **AI Research**: Use with AI research assistants
2. **Human Research**: Framework for manual research
3. **Hybrid Approach**: Combine AI and human research

### Knowledge Base Integration:
```bash
# Save research prompt
./kb index --path ./docs/research/

# Search for similar research
./kb search "[topic] research" --collection documentation
```

## Best Practices

- Be specific rather than general in questions
- Balance comprehensiveness with focus
- Consider current state and future implications
- Document assumptions and limitations
- Plan for iterative refinement
- Include measurable success criteria
- Specify output formats clearly

## Example Research Prompts

### Example 1: Market Validation
"Research the market opportunity for an AI-powered personal finance assistant targeting millennials, focusing on market size, competitive landscape, user willingness to pay, and regulatory considerations."

### Example 2: Technical Feasibility
"Investigate the feasibility of implementing real-time collaborative editing using CRDTs, comparing different approaches, performance implications, and integration challenges with existing architecture."

### Example 3: User Research
"Analyze how software developers currently manage technical debt, identifying pain points, existing solutions, decision criteria for tooling, and opportunities for innovation."

## Remember
- Quality of prompt determines quality of insights
- Specific questions yield actionable answers
- Consider multiple perspectives
- Plan for validation and iteration