# Advanced Elicitation Protocol

## Purpose
Comprehensive elicitation framework for gathering UX requirements, user insights, and design specifications through structured interaction methods.

## When to Use
- Creating new frontend specifications
- Gathering user persona requirements
- Defining information architecture
- Collecting accessibility needs
- Establishing design system requirements

## Core Principles

### 1. Progressive Disclosure
Start broad, then narrow focus based on responses
- Begin with high-level goals
- Drill down into specifics
- Validate understanding at each level

### 2. Context-Aware Method Selection
Choose elicitation technique based on:
- User's domain expertise
- Project complexity
- Available information
- Time constraints

### 3. Iterative Refinement
- Gather initial requirements
- Reflect back understanding
- Refine and validate
- Document decisions

## Elicitation Methods Framework

### Method 1: Structured Interview
**Best for**: New projects, unclear requirements
**Time**: 30-60 minutes
**Process**:
1. Open-ended questions about goals
2. Scenario-based discussions
3. Pain point identification
4. Success criteria definition

**Question Templates**:
```
Goals & Vision:
- "What does success look like for this [feature/product]?"
- "Who are your primary users and what do they need?"
- "What problems are you trying to solve?"

Context & Constraints:
- "What existing systems/processes need to integrate?"
- "What are your technical constraints?"
- "What's your timeline and budget considerations?"

User Experience:
- "Walk me through how a user would accomplish [key task]"
- "What would make this delightful vs. just functional?"
- "What are the biggest user frustrations with current solutions?"
```

### Method 2: Card Sorting (Information Architecture)
**Best for**: Site structure, navigation design
**Time**: 20-40 minutes
**Process**:
1. Present content/feature cards
2. Ask for logical groupings
3. Discuss naming conventions
4. Validate hierarchy

**Template**:
```
Virtual Card Sorting Session:

Content Areas to Organize:
[List all features/content areas]

Instructions:
1. "Group these items in a way that makes sense to you"
2. "What would you call each group?"
3. "Which items are most important to users?"
4. "How would users expect to find [specific item]?"

Follow-up:
- Primary navigation structure
- Secondary navigation needs
- Search strategy
- Mobile adaptation
```

### Method 3: User Story Mapping
**Best for**: User flows, feature prioritization
**Time**: 45-90 minutes
**Process**:
1. Identify user types
2. Map user journey stages
3. Detail activities at each stage
4. Prioritize features by impact

**Template**:
```
User Story Mapping Session:

User Types:
[List primary personas]

Journey Stages:
[Awareness → Evaluation → Purchase → Onboarding → Usage → Support]

For each stage, detail:
- User goals
- Required features
- Pain points to avoid
- Success metrics

Prioritization:
- Must have (MVP)
- Should have (V1)
- Could have (V2)
- Won't have (not planned)
```

### Method 4: Design Critique & Feedback
**Best for**: Existing designs, iterative improvement
**Time**: 30-45 minutes
**Process**:
1. Present design concepts
2. Structured feedback collection
3. Alternative exploration
4. Decision documentation

**Framework**:
```
Design Review Protocol:

What's Working:
- "What aspects support user goals?"
- "What feels intuitive or expected?"
- "What creates positive emotional response?"

What's Not Working:
- "Where might users get confused?"
- "What feels unnecessarily complex?"
- "What's missing for task completion?"

Alternative Approaches:
- "How else might we solve [specific problem]?"
- "What if we prioritized [different aspect]?"
- "What would [specific user type] expect here?"
```

### Method 5: Constraint & Requirements Gathering
**Best for**: Technical specifications, accessibility requirements
**Time**: 20-30 minutes
**Process**:
1. Identify mandatory requirements
2. Explore trade-offs
3. Document decisions
4. Plan validation

**Template**:
```
Requirements & Constraints:

Technical Constraints:
- Browser support requirements
- Performance requirements
- Integration requirements
- Security requirements

Business Constraints:
- Timeline limitations
- Budget constraints
- Brand guidelines
- Legal/compliance needs

User Constraints:
- Accessibility requirements
- Device/platform usage
- Bandwidth limitations
- Technical literacy levels

Trade-off Discussions:
- "If we had to choose between [A] and [B]..."
- "What's the minimum viable experience?"
- "Where can we compromise without impacting core goals?"
```

### Method 6: Rapid Prototype Feedback
**Best for**: Interaction design, usability validation
**Time**: 15-30 minutes
**Process**:
1. Present low-fidelity concepts
2. Task-based walkthroughs
3. Immediate feedback collection
4. Iteration planning

**Protocol**:
```
Prototype Testing:

Task Scenarios:
[List realistic user tasks]

For each task:
1. "Show me how you would [accomplish task]"
2. "What do you expect to happen when you [action]?"
3. "How would you recover if [error scenario]?"
4. "What would make this easier/faster?"

Specific Elements:
- Navigation clarity
- Form usability
- Error handling
- Mobile interactions
```

## Method Selection Guide

### Quick Decision Tree:
```
Is this a new project or feature?
├─ Yes: Start with Method 1 (Structured Interview)
└─ No: Is there existing design?
   ├─ Yes: Use Method 4 (Design Critique)
   └─ No: What's the primary need?
      ├─ Site structure: Method 2 (Card Sorting)
      ├─ User flows: Method 3 (User Story Mapping)
      ├─ Technical specs: Method 5 (Requirements)
      └─ Interaction design: Method 6 (Prototype Feedback)
```

### Context Factors:
- **High expertise user**: Methods 4, 5, 6 (faster, more technical)
- **Low expertise user**: Methods 1, 2, 3 (broader, more guided)
- **Complex project**: Combine multiple methods
- **Time constrained**: Methods 5, 6 (focused, specific)
- **Early stage**: Methods 1, 2, 3 (exploratory)
- **Validation stage**: Methods 4, 6 (evaluative)

## Interaction Patterns

### Opening Engagement
```
"I'm going to help you create a comprehensive [specification/design] by asking targeted questions. This should take about [time estimate].

The goal is to gather all the details needed for [specific outcome]. I'll ask follow-up questions to make sure I understand your vision correctly.

Shall we start with [initial area of focus]?"
```

### Transition Phrases
```
Between sections:
- "That helps me understand [topic]. Now let's explore [next topic]..."
- "Based on what you've shared about [A], I'd like to dig into [B]..."
- "Let me make sure I understand [concept] correctly before we move on..."

For clarification:
- "Can you give me a specific example of [concept]?"
- "When you say [term], do you mean [interpretation]?"
- "Help me understand the difference between [A] and [B]..."

For deeper insight:
- "What would happen if [scenario]?"
- "Why is [requirement] important to your users?"
- "What would make this [better/easier/more valuable]?"
```

### Closing & Validation
```
"Let me summarize what I've gathered to make sure I have this right:

[Key points summary]

Is this accurate? What would you add or change?

Based on this, the next steps are:
1. [Immediate action]
2. [Follow-up action]
3. [Validation step]

Does this plan work for you?"
```

## Documentation During Elicitation

### Real-time Capture
- Key quotes (exact language used)
- Decisions made and rationale
- Questions that need follow-up
- Assumptions to validate

### Immediate Summary
After each session, document:
- Primary insights
- Outstanding questions
- Next actions
- Confidence level on each requirement

### Template for Session Notes:
```
# [Method] Session - [Date]

## Participants
- [List attendees and roles]

## Objectives
- [What we aimed to learn/decide]

## Key Insights
- [Major findings, organized by topic]

## Decisions Made
- [Specific choices with rationale]

## Outstanding Questions
- [Items needing clarification/research]

## Next Steps
- [Immediate actions]
- [Follow-up sessions needed]
- [Validation activities]

## Confidence Assessment
- High confidence: [items we're sure about]
- Medium confidence: [items needing validation]
- Low confidence: [items needing more research]
```

## Quality Assurance

### During Sessions
- Reflect back what you heard
- Ask for specific examples
- Challenge assumptions gently
- Ensure all perspectives are heard

### After Sessions
- Review notes for gaps
- Validate interpretations
- Plan follow-up questions
- Share summary with participants

### Success Metrics
- Clear, actionable requirements
- Stakeholder alignment
- Reduced iteration cycles
- User-centered outcomes

## Integration with UX Workflow

### Pre-Elicitation
1. Review existing documentation
2. Identify key stakeholders
3. Prepare method-specific materials
4. Set clear session objectives

### During Elicitation
1. Follow selected method framework
2. Adapt based on responses
3. Document in real-time
4. Validate understanding

### Post-Elicitation
1. Synthesize findings
2. Create specifications
3. Share for validation
4. Plan next steps

This protocol ensures comprehensive, systematic gathering of UX requirements while maintaining efficiency and user focus.