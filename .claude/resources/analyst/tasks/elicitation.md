# Advanced Elicitation Task

## Overview
This task enables Ana to systematically extract tacit knowledge, uncover hidden requirements, and refine understanding through strategic questioning and probing techniques.

## Usage Scenarios

### Scenario 1: Document Refinement
User provides document, Ana elicits:
- Missing details
- Unclear specifications
- Hidden assumptions
- Edge cases
- Success criteria

### Scenario 2: Content Enhancement
Enrich existing content with:
- Deeper context
- Examples and scenarios
- User perspectives
- Technical details
- Strategic rationale

### Scenario 3: Requirement Discovery
Starting from rough concept, elicit:
- True needs vs stated wants
- Constraints and trade-offs
- Stakeholder perspectives
- Success metrics
- Risk factors

## Intelligent Method Selection Framework

### Step 1: Analyze the Context

#### Content Type Assessment
- **Technical**: APIs, architecture, code, infrastructure
- **Strategic**: Vision, goals, business case, roadmap
- **Process**: Workflows, procedures, operations
- **User-Focused**: Features, experiences, journeys
- **Creative**: Concepts, designs, innovations

#### Complexity Level
- **Simple**: Clear scope, well-defined, few dependencies
- **Moderate**: Some ambiguity, multiple stakeholders, interconnected
- **Complex**: High ambiguity, many unknowns, systemic issues

#### Audience Needs
- **Technical team**: Precision, completeness, edge cases
- **Business stakeholders**: Value, impact, ROI
- **End users**: Usability, benefits, workflows
- **Mixed audience**: Multiple perspectives needed

#### Project Stage
- **Discovery**: Broad exploration, many questions
- **Definition**: Narrowing scope, specific details
- **Refinement**: Gap filling, validation
- **Optimization**: Edge cases, improvements

#### Risk Profile
- **High Risk**: Critical systems, compliance, security
- **Moderate Risk**: Important but not critical
- **Low Risk**: Experimental, flexible requirements

### Step 2: Select Elicitation Methods

## Elicitation Methods Library

### Core Methods (Always Available)

**1. Open-Ended Probing**
- When: Need exploration, understanding context
- Questions: "Tell me more about...", "What led to...", "How does...work?"
- Output: Broad understanding, context, narrative

**2. Clarification Questions**
- When: Ambiguity, vague statements, assumptions
- Questions: "What exactly do you mean by...", "Can you give an example of..."
- Output: Precise definitions, concrete examples

**3. Five Whys**
- When: Need root causes, deeper understanding
- Questions: "Why is that important?" repeated 5 times
- Output: Underlying motivations, core requirements

**4. Scenario Analysis**
- When: Need concrete examples, edge cases
- Questions: "Walk me through what happens when...", "What if...occurs?"
- Output: Detailed workflows, exception handling

**5. Assumption Surfacing**
- When: Implicit knowledge, unstated constraints
- Questions: "What are you assuming about...", "What must be true for..."
- Output: Hidden assumptions, constraints, dependencies

### Technical Elicitation Methods

**6. Edge Case Exploration**
- When: Technical specifications, APIs, systems
- Questions: "What happens when...[extreme condition]", "How does it handle..."
- Output: Error handling, boundaries, limits

**7. Integration Probing**
- When: Systems, APIs, architecture
- Questions: "How does this connect to...", "What data flows between..."
- Output: Dependencies, interfaces, data contracts

**8. Performance Requirements**
- When: Technical systems, user experiences
- Questions: "How fast must...", "What volume of...", "How many concurrent..."
- Output: SLAs, scalability needs, performance targets

**9. Technical Constraints**
- When: Implementation planning, architecture
- Questions: "What technologies must/cannot be used...", "What are the platform limitations..."
- Output: Technology constraints, compatibility requirements

### Strategic Elicitation Methods

**10. Value Probing**
- When: Business cases, prioritization
- Questions: "What value does this deliver?", "How does this impact..."
- Output: Business value, ROI, strategic alignment

**11. Stakeholder Mapping**
- When: Projects with multiple stakeholders
- Questions: "Who else is affected by...", "Who needs to approve..."
- Output: Stakeholder matrix, concerns, requirements

**12. Trade-off Analysis**
- When: Conflicting requirements, resource constraints
- Questions: "If you had to choose between X and Y...", "What's the minimum acceptable..."
- Output: Priorities, must-haves vs nice-to-haves

**13. Success Criteria Definition**
- When: Project goals, feature completion
- Questions: "How will you know this is successful?", "What would failure look like?"
- Output: Measurable success criteria, acceptance criteria

### Creative Elicitation Methods

**14. Analogical Thinking**
- When: Novel problems, innovative solutions
- Questions: "What's this similar to?", "How do [other domains] handle..."
- Output: Analogies, transferable patterns, inspiration

**15. Reversal Questions**
- When: Stuck, need new perspectives
- Questions: "What would make this fail?", "How could we make this worse?"
- Output: Risks to avoid, constraints to challenge

**16. Future Casting**
- When: Long-term planning, evolution
- Questions: "How might this evolve in 2 years?", "What will users expect eventually?"
- Output: Future requirements, extensibility needs

### User-Focused Methods

**17. User Story Expansion**
- When: User features, experiences
- Questions: "As a [user], I want to...", "Why would a user need..."
- Output: User stories, acceptance criteria, user value

**18. Journey Mapping**
- When: User workflows, experiences
- Questions: "Walk me through the user's journey from...", "What happens before/after..."
- Output: End-to-end workflows, touchpoints, pain points

**19. Pain Point Analysis**
- When: Problem discovery, improvements
- Questions: "What's frustrating about...", "Where do users get stuck..."
- Output: Problems to solve, improvement opportunities

### Analytical Methods

**20. Gap Analysis**
- When: Assessing completeness
- Questions: "What's missing from...", "What haven't we covered..."
- Output: Missing elements, incomplete areas

**21. Dependency Mapping**
- When: Complex systems, projects
- Questions: "What must happen before...", "What depends on..."
- Output: Dependency graph, critical path, blockers

**22. Risk Elicitation**
- When: Project planning, critical systems
- Questions: "What could go wrong...", "What's the biggest risk..."
- Output: Risk register, mitigation strategies

## Execution Framework

### Phase 1: Context Assessment (30 seconds)

Quickly evaluate:
1. What type of content/problem is this?
2. What's the complexity level?
3. Who's the audience?
4. What stage is this project?
5. What's the risk level?

### Phase 2: Method Selection (30 seconds)

Choose 3-5 methods based on assessment:

**For Technical + Complex + High Risk:**
- Edge Case Exploration
- Integration Probing
- Technical Constraints
- Risk Elicitation
- Scenario Analysis

**For Strategic + Moderate + Business Audience:**
- Value Probing
- Stakeholder Mapping
- Trade-off Analysis
- Success Criteria Definition
- Gap Analysis

**For User-Focused + Simple + Low Risk:**
- User Story Expansion
- Journey Mapping
- Scenario Analysis
- Clarification Questions
- Pain Point Analysis

**For Creative + Complex + Mixed Audience:**
- Analogical Thinking
- Scenario Analysis
- Future Casting
- Assumption Surfacing
- Open-Ended Probing

### Phase 3: Question Presentation Format

Present questions in organized batches:

```markdown
## [Category Name] - [Method Used]

I'd like to understand [specific aspect] better. Here are [X] questions:

1. [Question designed to elicit specific type of information]
2. [Question building on first]
3. [Question exploring edge case or alternative]
...

**Why these questions matter**: [Brief explanation of what insights these will provide]
```

### Phase 4: Application

Execute in layers:

**Layer 1: Broad Understanding (3-5 questions)**
- Open-ended probing
- Context gathering
- High-level clarification

**Layer 2: Specific Details (5-8 questions)**
- Scenario analysis
- Technical probing
- Requirement specifics

**Layer 3: Edge Cases & Validation (3-5 questions)**
- Edge cases
- Assumptions
- Gap analysis

**Layer 4: Synthesis & Confirmation (2-3 questions)**
- Trade-offs
- Success criteria
- Final validation

### Phase 5: Integration

After eliciting information:

1. **Synthesize**: Combine original content with elicited knowledge
2. **Validate**: Confirm understanding is correct
3. **Document**: Create enriched, comprehensive output
4. **Highlight**: Show what was added/clarified
5. **Iterate**: Offer another round if needed

## Context-Specific Templates

### Template A: Technical Specification Enhancement

```markdown
## Technical Clarity Questions

**Architecture & Design** (Integration Probing)
1. [How does this component integrate with X?]
2. [What's the data flow between Y and Z?]
3. [What are the key interfaces/contracts?]

**Edge Cases & Error Handling** (Edge Case Exploration)
4. [What happens when input exceeds limits?]
5. [How should the system handle failures in X?]
6. [What are the timeout/retry strategies?]

**Performance & Scale** (Performance Requirements)
7. [What's the expected transaction volume?]
8. [What are the latency requirements?]
9. [How should this scale?]

**Constraints & Dependencies** (Technical Constraints)
10. [What technology constraints exist?]
11. [What are the critical dependencies?]
12. [What security/compliance requirements apply?]
```

### Template B: Strategic Document Enhancement

```markdown
## Strategic Alignment Questions

**Value & Impact** (Value Probing)
1. [What business outcomes does this enable?]
2. [How does this align with company strategy?]
3. [What's the expected ROI or impact?]

**Stakeholders & Governance** (Stakeholder Mapping)
4. [Who are the key stakeholders?]
5. [Who needs to approve decisions?]
6. [What are each stakeholder's main concerns?]

**Trade-offs & Priorities** (Trade-off Analysis)
7. [What are the must-haves vs nice-to-haves?]
8. [If resources are limited, what takes priority?]
9. [What compromises are acceptable?]

**Success Definition** (Success Criteria)
10. [How will we measure success?]
11. [What does "done" look like?]
12. [What would constitute failure?]
```

### Template C: User Story Enhancement

```markdown
## User Experience Deep Dive

**User Context** (Journey Mapping)
1. [Walk me through the user's full journey]
2. [What happens before they reach this feature?]
3. [What's their next step after completing this?]

**User Needs** (User Story Expansion)
4. [Why does the user need this capability?]
5. [What problem does this solve for them?]
6. [What alternatives do they use today?]

**User Pain Points** (Pain Point Analysis)
7. [What frustrates users about the current approach?]
8. [Where do users typically get stuck?]
9. [What would delight users?]

**Scenarios & Examples** (Scenario Analysis)
10. [Give me a specific example of a user doing this]
11. [What are the most common use cases?]
12. [What are the edge cases users encounter?]
```

## Quality Principles

### Question Quality:
- **Specific**: Not "Tell me about X" but "How does X handle scenario Y?"
- **Purposeful**: Each question targets specific knowledge gap
- **Layered**: Questions build on each other
- **Neutral**: Avoid leading or biased questions
- **Actionable**: Answers should inform decisions

### Facilitation Style:
- **Professional**: Consultant helping extract expertise
- **Curious**: Genuine interest in understanding
- **Systematic**: Organized, logical flow
- **Adaptive**: Adjust based on responses
- **Synthesizing**: Connect dots, identify patterns

### Red Flags:
- Asking questions for the sake of asking
- Overwhelming with too many questions at once
- Not adapting to user's response style
- Missing obvious gaps
- Repeating information user already provided

## Success Metrics

### For User:
- Feels heard and understood
- Previously tacit knowledge is now explicit
- Document/understanding is significantly enhanced
- Effort feels worthwhile
- Would engage in this process again

### For Output:
- 50%+ more detail than original
- Ambiguities resolved
- Edge cases covered
- Clear success criteria
- Actionable and complete

### For Process:
- Right questions for the context
- Efficient (not exhaustive)
- Logical flow maintained
- User engaged throughout
- Clear value delivered

## Advanced Techniques

### Progressive Refinement:
1. Start broad, narrow progressively
2. Use answers to shape next questions
3. Circle back to earlier topics with new insights
4. Build mental model as you go

### Pattern Recognition:
- Notice what user emphasizes
- Identify repeated themes
- Spot contradictions
- Recognize unstated assumptions

### Meta-Elicitation:
- "How confident are you about X?"
- "What don't you know yet?"
- "Where are you most uncertain?"
- "What would you most like clarity on?"

### Silence & Space:
- Don't rush to fill silence
- Give user time to think
- Allow for "I don't know" responses
- Create space for deeper reflection

## Closing & Documentation

After elicitation, deliver:

```markdown
# Enhanced [Document Title]

## Original Content
[Original content with preserved formatting]

## Key Clarifications & Additions

### [Category 1]
- **Clarified**: [What was ambiguous and is now clear]
- **Added**: [New information elicited]
- **Examples**: [Concrete scenarios provided]

### [Category 2]
...

## Integrated Enhanced Version
[Original content enriched with elicited information, clearly formatted]

## Outstanding Questions
[Any remaining gaps or areas needing further input]

## Recommended Next Steps
1. [Action item]
2. [Action item]
```

## Adaptation Guidelines

### For Reluctant Responders:
- Fewer questions per batch (3-4 max)
- Explain value of each question
- Share what you'll do with answers
- Make it feel collaborative, not interrogative

### For Verbose Responders:
- Guide back to specific questions
- Summarize and validate understanding
- Use closed questions occasionally
- Help organize their thoughts

### For Expert Users:
- More technical/advanced questions
- Challenge assumptions respectfully
- Explore edge cases deeply
- Match their sophistication level

### For Novice Users:
- Provide more context for why asking
- Use simpler language
- Offer examples with questions
- Be patient with "I don't know" responses
