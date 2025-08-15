# Cognitive Tool Program: Assess Understanding
# Structured reasoning program to evaluate learner's current understanding state

## Program Purpose
This prompt program guides systematic assessment of learner understanding to determine appropriate teaching constraints and next steps.

## Input Schema
```yaml
learner_state:
  topic: Current subject being learned
  responses: Recent learner responses and explanations
  questions_asked: Types of questions learner is asking
  performance: Success on practice problems or tasks
  confidence_indicators: Self-assessment and behavioral cues
  prior_sessions: Previous understanding levels and progress
```

## Assessment Framework

### Step 1: Analyze Explanation Quality
```
IF learner can explain concept in own words clearly:
  → CONCEPTUAL UNDERSTANDING PRESENT
ELSE IF learner uses memorized phrases/jargon:
  → SURFACE-LEVEL UNDERSTANDING
ELSE IF learner struggles to explain:
  → INSUFFICIENT UNDERSTANDING
```

### Step 2: Evaluate Transfer Ability
```
IF learner applies concept to novel contexts successfully:
  → DEEP UNDERSTANDING
ELSE IF learner handles slight variations:
  → DEVELOPING UNDERSTANDING  
ELSE IF learner only repeats examples:
  → PROCEDURAL UNDERSTANDING ONLY
```

### Step 3: Check Pattern Recognition
```
IF learner identifies underlying patterns across examples:
  → STRONG PATTERN RECOGNITION
ELSE IF learner recognizes patterns with guidance:
  → EMERGING PATTERN RECOGNITION
ELSE IF learner focuses on surface features:
  → LIMITED PATTERN RECOGNITION
```

### Step 4: Assess Constraint Understanding
```
IF learner knows when/why concepts don't apply:
  → BOUNDARY UNDERSTANDING PRESENT
ELSE IF learner overgeneralizes concepts:
  → BOUNDARY UNDERSTANDING WEAK
ELSE IF learner uncertain about limitations:
  → BOUNDARY UNDERSTANDING ABSENT
```

## Understanding Levels Matrix

| Dimension | Novice | Developing | Proficient | Expert |
|-----------|---------|------------|------------|---------|
| **Explanation** | Memorized terms | Own words, basic | Clear, accurate | Multiple perspectives |
| **Transfer** | None | Slight variations | Novel contexts | Creative applications |
| **Patterns** | Surface features | Simple patterns | Complex patterns | Pattern creation |
| **Boundaries** | Unaware | Some awareness | Clear boundaries | Boundary reasoning |
| **Questions** | What/how | Why basic | Why complex | What-if advanced |

## Diagnostic Questions by Level

### Surface Understanding Indicators
- "What is [concept]?" → Gives definition
- "How do you do [procedure]?" → Lists steps
- Cannot explain why steps work
- Struggles with variations

### Developing Understanding Indicators  
- "Why does this work?" → Partial explanation
- "What would happen if...?" → Some prediction
- Connects to some prior knowledge
- Handles guided variations

### Deep Understanding Indicators
- "Explain this to a beginner" → Clear, simple explanation
- "When would this not work?" → Identifies limitations
- "How is this like [other concept]?" → Makes connections
- "Design a new example" → Creates novel applications

## Assessment Protocol

### Initial Assessment Questions
1. **Background Check**: "What do you already know about [topic]?"
2. **Confidence Check**: "How confident do you feel about [concept]?"
3. **Connection Check**: "What does this remind you of?"
4. **Application Check**: "Where might you use this?"

### Understanding Probe Questions
1. **Explanation Test**: "Explain [concept] in your own words"
2. **Analogy Test**: "What is this similar to?"
3. **Prediction Test**: "What would happen if we changed [parameter]?"
4. **Boundary Test**: "When would this approach fail?"

### Transfer Assessment Tasks
1. **Near Transfer**: Same concept, different context
2. **Far Transfer**: Same principles, different domain
3. **Creative Transfer**: Novel application creation
4. **Teaching Transfer**: Explain to imaginary learner

## Constraint-Based Assessment Indicators

### Well-Constrained Understanding
- Consistent explanations across contexts
- Appropriate application boundaries
- Creative but valid extensions
- Metacognitive awareness present

### Under-Constrained Understanding  
- Overgeneralization to inappropriate contexts
- Missing boundary awareness
- Inconsistent application
- Overconfidence in explanations

### Over-Constrained Understanding
- Limited to exact examples taught
- Rigid application patterns
- Fear of variations or extensions
- Underconfidence despite competence

## Program Output Template

```markdown
## Understanding Assessment

### Current Understanding Level: {Novice|Developing|Proficient|Expert}

### Dimension Analysis:
- **Explanation Quality**: {Surface|Developing|Clear|Sophisticated}
- **Transfer Ability**: {None|Limited|Good|Creative}
- **Pattern Recognition**: {Weak|Emerging|Strong|Expert}
- **Boundary Awareness**: {Absent|Partial|Clear|Sophisticated}

### Constraint State: {Under-constrained|Well-constrained|Over-constrained}

### Evidence:
- {Specific examples from learner responses}
- {Performance on assessment tasks}
- {Quality of questions asked}
- {Confidence and metacognitive indicators}

### Next Steps:
- **Immediate Focus**: {What needs attention first}
- **Methodology Recommendation**: {Which approach to use}
- **Constraint Adjustments**: {How to modify teaching constraints}
- **Zelda Coordination**: {Whether knowledge capture is ready}

### Learning Trajectory:
- **Strengths to Build On**: {Areas of solid understanding}
- **Gaps to Address**: {Missing pieces}
- **Extensions Ready**: {Areas for advancement}
```

## Integration with Constraint Architecture

### Atomic Constraint Assessment
- Dialogue patterns: Is learner engaging appropriately?
- Questioning syntax: Are responses showing depth?
- Understanding validation: Which validation methods work?
- Pacing control: Is cognitive load appropriate?
- Metacognitive triggers: Is self-awareness developing?

### Molecular Protocol Selection
Based on assessment, recommend:
- **Feynman Method**: If explanation quality needs work
- **Socratic Dialogue**: If pattern recognition needs development
- **Elaborative Interrogation**: If transfer ability is weak
- **Problem-Based Learning**: If application skills need practice

### Cellular Memory Updates
- Update learning progression tracking
- Note misconception patterns revealed
- Record effective assessment approaches
- Update adaptive strategy preferences

### Organ Coordination
- Trigger Zelda if understanding breakthrough achieved
- Coordinate with knowledge base for related concepts
- Share assessment results with other learning tools

## Edge Cases and Considerations

### Confident but Incorrect
- Learner sounds confident but has misconceptions
- **Response**: Gentle probing with counterexamples
- **Avoid**: Direct contradiction or embarrassment

### Correct but Insecure
- Learner understands but lacks confidence
- **Response**: Validation and confidence building
- **Focus**: Reinforcement of correct understanding

### Partial Understanding
- Learner has pieces but not complete picture
- **Response**: Fill gaps systematically
- **Approach**: Build on what's correct

### Domain Transfer Issues
- Understands in one context but not others
- **Response**: Explicit connection making
- **Method**: Guided transfer practice

This program ensures systematic, evidence-based understanding assessment that informs constraint-based teaching decisions.