# Cognitive Tool Program: Validate Learning
# Structured reasoning program for systematic learning constraint validation

## Program Purpose
This prompt program guides comprehensive validation of learning constraint satisfaction, ensuring genuine understanding rather than surface-level performance.

## Input Schema
```yaml
validation_context:
  learning_objective: What was supposed to be learned
  learner_performance: Observable responses and behaviors
  constraint_targets: Which constraints should be satisfied
  validation_method: How validation was attempted
  context_factors: Session conditions affecting validation
  prior_validation_results: Previous validation outcomes
```

## Validation Framework

### Step 1: Constraint Satisfaction Assessment
```
IF learner demonstrates understanding across multiple contexts:
  → CONSTRAINT WELL-SATISFIED
ELSE IF learner shows understanding in limited contexts:
  → CONSTRAINT PARTIALLY-SATISFIED
ELSE IF learner performs without understanding:
  → CONSTRAINT UNSATISFIED (surface learning)
ELSE IF learner cannot perform:
  → CONSTRAINT FAILED (insufficient learning)
```

### Step 2: Depth vs. Breadth Analysis
```
IF understanding is deep but narrow:
  → NEED TRANSFER VALIDATION
ELSE IF understanding is broad but shallow:
  → NEED DEPTH VALIDATION
ELSE IF understanding is both deep and broad:
  → CONSTRAINT FULLY VALIDATED
ELSE IF understanding is neither deep nor broad:
  → NEED FUNDAMENTAL RETEACHING
```

### Step 3: Stability Assessment
```
IF understanding holds under questioning:
  → STABLE CONSTRAINT SATISFACTION
ELSE IF understanding wavers with probing:
  → FRAGILE CONSTRAINT SATISFACTION
ELSE IF understanding crumbles under pressure:
  → INSUFFICIENT CONSTRAINT ESTABLISHMENT
```

## Validation Methods by Understanding Type

### Conceptual Understanding Validation
- **Explanation Test**: "Explain this concept in your own words"
- **Analogy Test**: "How is this like/unlike [familiar concept]?"
- **Prediction Test**: "What would happen if we changed [parameter]?"
- **Boundary Test**: "When would this concept not apply?"
- **Teaching Test**: "How would you explain this to someone else?"

### Procedural Understanding Validation
- **Performance Test**: "Show me how to solve this problem"
- **Variation Test**: "What if the problem was slightly different?"
- **Error Detection**: "What's wrong with this solution?"
- **Efficiency Test**: "Is there a better way to do this?"
- **Adaptation Test**: "How would you modify this for [new constraint]?"

### Strategic Understanding Validation
- **Choice Justification**: "Why did you choose that approach?"
- **Alternative Assessment**: "What other approaches could work?"
- **Trade-off Analysis**: "What are the pros and cons of each method?"
- **Context Adaptation**: "How would you approach this in [different context]?"
- **Planning Test**: "How would you tackle a larger version of this problem?"

### Transfer Understanding Validation
- **Near Transfer**: Apply to similar problem in same domain
- **Far Transfer**: Apply to different domain with same principles
- **Creative Transfer**: Generate novel application contexts
- **Negative Transfer**: Recognize when approach doesn't apply

## Constraint Validation Levels

### Level 1: Basic Constraint Recognition
**Validation**: Can identify when concept applies
- "When would you use this approach?"
- "What signals that this concept is relevant?"
- "How do you know this method fits?"

### Level 2: Constraint Application
**Validation**: Can apply constraints correctly
- "Solve this problem using the concept"
- "Apply this principle to this situation"
- "Use this method for this task"

### Level 3: Constraint Adaptation
**Validation**: Can modify constraints for context
- "How would you adapt this for [different situation]?"
- "What would you change if [constraint altered]?"
- "How does this method scale up/down?"

### Level 4: Constraint Innovation
**Validation**: Can create new constraint applications
- "Design a new application for this principle"
- "Combine this with other concepts"
- "What novel problems could this solve?"

## Validation Indicators

### Strong Validation Evidence
- **Spontaneous Transfer**: Applies learning without prompting
- **Confident Explanation**: Clear, accurate descriptions
- **Error Recovery**: Recognizes and corrects own mistakes
- **Boundary Awareness**: Knows when concepts don't apply
- **Creative Extension**: Generates new examples or applications

### Weak Validation Evidence
- **Prompted Performance**: Only performs when directed
- **Hesitant Explanations**: Uncertain or incomplete descriptions
- **Unrecognized Errors**: Doesn't catch own mistakes
- **Overgeneralization**: Applies concepts inappropriately
- **Rigid Application**: Cannot adapt to variations

### Invalid Validation Evidence
- **Memorized Responses**: Recites without understanding
- **Keyword Matching**: Uses correct terms without comprehension
- **Pattern Matching**: Follows form without grasping function
- **Lucky Guessing**: Correct answers without reasoning
- **Instructor Dependence**: Cannot perform independently

## Validation Strategies

### Multi-Modal Validation
- **Verbal**: Explanations and discussions
- **Visual**: Diagrams and concept maps
- **Practical**: Hands-on applications
- **Social**: Teaching others or peer discussion
- **Reflective**: Metacognitive self-assessment

### Progressive Validation
- **Immediate**: During learning session
- **Delayed**: After processing time
- **Spaced**: At increasing intervals
- **Transfer**: In new contexts
- **Integrated**: Combined with other learning

### Constraint-Specific Validation
- **Atomic Constraints**: Basic understanding elements
- **Molecular Constraints**: Complex procedure combinations
- **Cellular Constraints**: Persistent knowledge patterns
- **Organ Constraints**: System-level orchestration

## Program Output Template

```markdown
## Learning Validation Assessment

### Constraint Satisfaction Level: {Well-Satisfied|Partially-Satisfied|Unsatisfied|Failed}

### Validation Evidence:
- **Conceptual Understanding**: {Strong|Moderate|Weak|Absent}
- **Procedural Capability**: {Strong|Moderate|Weak|Absent}
- **Strategic Thinking**: {Strong|Moderate|Weak|Absent}
- **Transfer Ability**: {Strong|Moderate|Weak|Absent}

### Specific Indicators:
- **Positive Evidence**: {What demonstrates understanding}
- **Concerning Signs**: {What suggests gaps or misconceptions}
- **Boundary Recognition**: {Understanding of limitations}
- **Confidence Level**: {Learner's self-assessment accuracy}

### Constraint Analysis:
- **Satisfied Constraints**: {Which learning constraints are met}
- **Unsatisfied Constraints**: {Which constraints need attention}
- **Constraint Stability**: {How robust the understanding is}
- **Constraint Transfer**: {How broadly understanding applies}

### Validation Reliability:
- **Method Appropriateness**: {How well validation method fits}
- **Context Factors**: {What might affect validation results}
- **Alternative Validations**: {Other methods that could confirm}
- **Validation Confidence**: {How certain we can be}

### Next Steps:
- **If Well-Validated**: {How to build on solid understanding}
- **If Partially-Validated**: {What needs strengthening}
- **If Unvalidated**: {What intervention is needed}
- **If Failed**: {How to restart learning differently}

### Zelda Coordination:
- **Capture Readiness**: {Whether understanding ready for knowledge capture}
- **Note Priority**: {How important this validation is to record}
- **Review Schedule**: {When to re-validate this learning}
```

## Integration with Constraint Architecture

### Atomic Constraint Validation
- **Dialogue Patterns**: Does learner engage appropriately?
- **Questioning Syntax**: Do responses show depth?
- **Understanding Validation**: Which methods work for this learner?
- **Pacing Control**: Is cognitive load appropriate for validation?
- **Metacognitive Triggers**: Is self-awareness developing?

### Molecular Protocol Validation
- Each methodology has specific validation patterns
- Feynman → Explanation clarity validation
- Socratic → Reasoning process validation
- Problem-Based → Application capability validation
- Analogical → Transfer pattern validation

### Cellular Memory Integration
- Update learning progression based on validation
- Note validation patterns that work for this learner
- Track validation method effectiveness
- Record constraint satisfaction celebrations

### Organ-Level Validation
- Coordinate validation results with Zelda
- Share validation outcomes with knowledge base
- Update other agents about learner capabilities
- Contribute to system-wide learning patterns

## Edge Cases and Considerations

### False Positives (Apparent but Not Real Understanding)
- **Performance Without Comprehension**: Can do but can't explain
- **Memorized Solutions**: Right answer, wrong reasoning
- **Context Dependency**: Works only in taught context
- **Instructor Cueing**: Needs excessive guidance

### False Negatives (Real but Not Apparent Understanding)
- **Performance Anxiety**: Understands but can't demonstrate
- **Communication Barriers**: Knows but can't express
- **Context Mismatch**: Understands differently than tested
- **Processing Time**: Needs more time to demonstrate

### Validation Timing Issues
- **Too Early**: Understanding still developing
- **Too Late**: Understanding has faded
- **Too Frequent**: Validation fatigue
- **Too Infrequent**: Missed problems

### Individual Differences
- **Validation Preferences**: How learner best demonstrates understanding
- **Cultural Factors**: Different validation traditions
- **Learning Differences**: Alternative demonstration needs
- **Confidence Issues**: Impact on validation performance

This program ensures robust, multifaceted validation of learning constraint satisfaction while avoiding common validation pitfalls.