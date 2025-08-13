# Advanced Elicitation Task

## Purpose
Provide intelligent, context-aware elicitation methods to refine and enhance content through multiple analytical perspectives. This task enables deeper exploration and quality improvement through structured techniques.

## Usage Scenarios

### Scenario 1: Document Section Refinement
After outputting any document section:
1. Provide brief context summary
2. Present 9 contextually-relevant methods
3. User selects by number (0-8) or 9 to proceed
4. Apply method and re-offer choices

### Scenario 2: General Content Enhancement
User requests elicitation on any content:
- Analyze context type and complexity
- Select 9 most relevant methods
- Execute same numbered selection process

## Intelligent Method Selection

### Context Analysis Framework

Before presenting options, analyze:

**Content Dimensions:**
- Type: Technical spec, user story, architecture, requirements, analysis
- Complexity: Simple, moderate, complex, critical
- Audience: Technical, business, mixed, external
- Stage: Draft, review, final, iteration
- Risk Level: Low, medium, high impact

**Selection Strategy:**

1. **Core Methods** (Always include 3-4):
   - Expand/Contract for Audience
   - Critique and Refine
   - Identify Risks
   - Validate Alignment

2. **Context-Specific** (Select 4-5 based on analysis):
   - Technical → Tree of Thoughts, First Principles
   - Strategic → Red Team, SWOT, Pre-mortem
   - Creative → Innovation Tournament, What If
   - User-Facing → Stakeholder Perspectives, Journey Mapping

3. **Option 9**: Always "Proceed / No Further Actions"

## Elicitation Methods Library

### Core Methods (Universal Application)

#### Expand or Contract for Audience
- Adjust detail level for specific audience
- Add technical depth or simplify
- Include examples or remove jargon
- Tailor tone and terminology

#### Critique and Refine
- Systematic quality review
- Identify gaps and weaknesses
- Suggest improvements
- Strengthen arguments

#### Identify Potential Risks
- Surface hidden risks
- Assess probability and impact
- Propose mitigations
- Create contingency plans

#### Assess Alignment with Goals
- Check against stated objectives
- Verify success criteria coverage
- Identify misalignments
- Suggest corrections

### Technical Methods

#### Tree of Thoughts (ToT)
- Branch into multiple solution paths
- Evaluate each branch systematically
- Compare and combine best elements
- Document decision rationale

#### First Principles Decomposition
- Break down to fundamental truths
- Question every assumption
- Rebuild from basics
- Identify unnecessary complexity

#### Chain of Density
- Progressive information compression
- Layer detail appropriately
- Optimize information density
- Balance completeness and clarity

#### Technical Debt Analysis
- Identify shortcuts taken
- Assess long-term implications
- Prioritize remediation
- Document trade-offs

### Strategic Methods

#### Red Team vs Blue Team
- Attack assumptions (Red)
- Defend approach (Blue)
- Find vulnerabilities
- Strengthen position

#### Pre-mortem Analysis
- Imagine failure scenario
- Work backward to causes
- Identify prevention measures
- Build resilience

#### SWOT Deep Dive
- Strengths: What advantages?
- Weaknesses: What limitations?
- Opportunities: What possibilities?
- Threats: What challenges?

#### Porter's Five Forces
- Competitive rivalry
- Supplier power
- Buyer power
- Threat of substitution
- Threat of new entry

### Creative Methods

#### Innovation Tournament
- Generate competing solutions
- Score against criteria
- Combine best features
- Select winner

#### What If Scenarios
- Explore edge cases
- Test boundaries
- Challenge constraints
- Discover opportunities

#### Analogical Extension
- Find similar solved problems
- Map solution patterns
- Adapt to current context
- Extract principles

#### Constraint Removal
- Remove one constraint
- Explore new possibilities
- Gradually reintroduce
- Find flexibility

### User-Focused Methods

#### Stakeholder Roundtable
- View from each stakeholder
- Identify conflicts
- Find common ground
- Balance interests

#### Journey Mapping
- Map user flow
- Identify pain points
- Find opportunities
- Optimize experience

#### Persona Analysis
- Create detailed personas
- Test against each persona
- Identify gaps
- Ensure coverage

#### Jobs-to-be-Done
- What job does this solve?
- How well does it solve it?
- What's missing?
- What's unnecessary?

### Analytical Methods

#### Five Whys Deep Dive
- Start with surface issue
- Ask why repeatedly
- Reach root cause
- Address fundamentally

#### Morphological Analysis
- Identify parameters
- List parameter options
- Explore combinations
- Find optimal configuration

#### Systems Thinking
- Map interconnections
- Identify feedback loops
- Find leverage points
- Predict ripple effects

#### Meta-Analysis
- Step back to overview
- Compare to alternatives
- Assess approach itself
- Question methodology

## Execution Framework

### Step 1: Context Assessment
```yaml
context_analysis:
  content_type: [identify]
  complexity: [assess]
  stakeholders: [list]
  risks: [evaluate]
  stage: [determine]
```

### Step 2: Method Selection
Based on context, select 9 methods:
- 3-4 core methods (always relevant)
- 4-5 context-specific methods
- Balance analytical styles

### Step 3: Presentation Format
```
**Advanced Elicitation Options**
Choose a method (0-8) or 9 to proceed:

0. [Method Name] - [Brief description]
1. [Method Name] - [Brief description]
2. [Method Name] - [Brief description]
3. [Method Name] - [Brief description]
4. [Method Name] - [Brief description]
5. [Method Name] - [Brief description]
6. [Method Name] - [Brief description]
7. [Method Name] - [Brief description]
8. [Method Name] - [Brief description]
9. Proceed / No Further Actions

Your selection:
```

### Step 4: Method Application

**Execution Process:**
1. Acknowledge selection
2. Briefly explain method
3. Apply to specific content
4. Provide concrete results
5. Offer to integrate changes
6. Re-present options

**Application Guidelines:**
- Be specific and actionable
- Reference actual content
- Provide examples
- Keep focused and concise
- Build on previous refinements

### Step 5: Integration
After each elicitation:
- Show how content would change
- Highlight improvements
- Track cumulative enhancements
- Maintain version history

## Context-Specific Templates

### For Technical Documentation
```
0. Add implementation examples
1. Include error handling scenarios
2. Expand API specifications
3. Add performance considerations
4. Include security analysis
5. Tree of Thoughts for architecture
6. First Principles simplification
7. Technical debt assessment
8. Identify integration risks
9. Proceed
```

### For Strategic Planning
```
0. SWOT analysis
1. Red Team challenge
2. Pre-mortem failure analysis
3. Stakeholder impact assessment
4. Market dynamics evaluation
5. Competitive positioning
6. Resource optimization
7. Timeline stress testing
8. Success metrics validation
9. Proceed
```

### For User Stories
```
0. Persona perspective check
1. Journey mapping
2. Edge case exploration
3. Acceptance criteria expansion
4. Technical feasibility
5. Dependencies identification
6. Story splitting options
7. Value proposition clarity
8. Implementation hints
9. Proceed
```

## Knowledge Base Integration

### Before Elicitation:
```bash
# Check for similar refinements
./.vector_db/kb search "elicitation [topic]" --collection documentation

# Find relevant patterns
./.vector_db/kb search "[content type]" --collection architecture
```

### After Elicitation:
```bash
# Save refined version
./.vector_db/kb index --path ./docs/refined/[filename]

# Document elicitation patterns
./.vector_db/kb index --mode docs
```

## Quality Principles

### Selection Intelligence
- Match methods to content type
- Consider user's goals
- Balance depth and breadth
- Vary analytical angles

### Application Excellence
- Concrete, specific feedback
- Clear improvement paths
- Preserve original intent
- Build incrementally

### User Experience
- Clear option descriptions
- Quick to understand
- Easy to select
- Valuable results

## Example Interactions

### Technical Specification Elicitation
```
Ana: I've completed the API specification section. Let me offer elicitation options to enhance it:

**Advanced Elicitation Options**
Choose a method (0-8) or 9 to proceed:

0. Add error handling examples
1. Include rate limiting details
2. Tree of Thoughts for endpoint design
3. Security vulnerability analysis
4. Performance optimization opportunities
5. API versioning strategy
6. Integration test scenarios
7. SDK implementation hints
8. Breaking change assessment
9. Proceed to next section

Your selection: 3

Applying Security Vulnerability Analysis...

[Detailed security analysis with specific recommendations]

Would you like to apply another elicitation method?
[Re-presents options]
```

### Strategic Document Elicitation
```
Ana: The market analysis is complete. Here are refinement options:

**Advanced Elicitation Options**
Choose a method (0-8) or 9 to proceed:

0. Expand with sensitivity analysis
1. Deep dive competitor comparison
2. Red Team market assumptions
3. Porter's Five Forces application
4. Customer segment personas
5. Pre-mortem on go-to-market
6. Adjacent market opportunities
7. Disruption scenario planning
8. Success metrics validation
9. Proceed to next section

Your selection:
```

## Success Metrics
- Relevant method selection (>90% applicable)
- Clear value delivery per elicitation
- Progressive quality improvement
- User engagement maintenance
- Actionable refinements

## Remember
- Analyze context before selecting
- Offer most relevant methods
- Apply methods concretely
- Build on previous refinements
- Keep user in control
- Document improvements