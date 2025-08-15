# Cognitive Tool Program: Adapt Instruction
# Structured reasoning program for real-time instructional constraint adaptation

## Program Purpose
This prompt program guides dynamic adaptation of teaching constraints based on real-time learner responses, ensuring optimal learning constraint satisfaction.

## Input Schema
```yaml
adaptation_context:
  current_instruction_state: What approach is being used
  learner_response_pattern: How learner is responding
  constraint_satisfaction_level: How well constraints are being met
  adaptation_triggers: What indicates need for change
  available_adaptations: What modifications are possible
  session_context: Time, energy, progress factors
```

## Adaptation Framework

### Step 1: Detect Adaptation Signals
```
IF learner showing confusion despite explanation:
  → COMPREHENSION ADAPTATION NEEDED
ELSE IF learner disengaged or bored:
  → ENGAGEMENT ADAPTATION NEEDED
ELSE IF learner overwhelmed or frustrated:
  → COGNITIVE LOAD ADAPTATION NEEDED
ELSE IF learner advancing faster than expected:
  → ACCELERATION ADAPTATION NEEDED
ELSE IF constraint satisfaction declining:
  → CONSTRAINT REALIGNMENT NEEDED
```

### Step 2: Diagnose Root Cause
```
IF adaptation signal is comprehension issue:
  → Analyze: abstraction level, prerequisite gaps, explanation clarity
ELSE IF adaptation signal is engagement issue:
  → Analyze: relevance, challenge level, interaction style
ELSE IF adaptation signal is cognitive overload:
  → Analyze: information density, pace, complexity
ELSE IF adaptation signal is under-challenge:
  → Analyze: depth level, extension opportunities, autonomy
```

### Step 3: Select Adaptation Strategy
```
IF root cause is conceptual:
  → REPRESENTATION ADAPTATION (analogies, visuals, examples)
ELSE IF root cause is motivational:
  → ENGAGEMENT ADAPTATION (relevance, interaction, choice)
ELSE IF root cause is cognitive:
  → LOAD ADAPTATION (chunking, pacing, scaffolding)
ELSE IF root cause is instructional:
  → METHODOLOGY ADAPTATION (different approach entirely)
```

## Adaptation Signal Detection

### Comprehension Signals
**Confusion Indicators**:
- Longer response times without deeper thinking
- Repeated questions about same concept
- Contradictory statements
- "I don't understand" expressions
- Off-topic responses

**Response**: Simplify, re-explain, use different approach

### Engagement Signals
**Disengagement Indicators**:
- Minimal responses
- Clock-watching behavior
- Passive participation
- "Do I have to?" attitudes
- Attention drifting

**Response**: Increase relevance, add interaction, show value

### Cognitive Load Signals
**Overload Indicators**:
- Information processing delays
- Increased error rates
- Expressions of overwhelm
- Requests to slow down
- Attention fragmentation

**Response**: Reduce complexity, chunk information, slow pace

### Challenge Level Signals
**Under-challenge Indicators**:
- Quick, automatic responses
- Requests for "something harder"
- Completion well before expected
- Signs of boredom
- Off-task creativity

**Response**: Increase complexity, add depth, accelerate pace

## Adaptation Strategies

### Representation Adaptations
**From Abstract to Concrete**:
- Add specific examples
- Use visual representations
- Provide hands-on activities
- Create physical analogies

**From Verbal to Visual**:
- Draw diagrams or concept maps
- Use animations or simulations
- Create visual metaphors
- Provide graphic organizers

**From Complex to Simple**:
- Break into smaller steps
- Remove non-essential details
- Use familiar vocabulary
- Start with basics

### Interaction Adaptations
**From Passive to Active**:
- Ask more questions
- Require learner participation
- Create hands-on activities
- Use collaborative elements

**From Individual to Social**:
- Encourage teaching others
- Use peer discussion
- Create group activities
- Share different perspectives

**From Guided to Independent**:
- Reduce scaffolding gradually
- Encourage self-direction
- Provide choices
- Support exploration

### Cognitive Load Adaptations
**Reduce Extraneous Load**:
- Remove unnecessary information
- Simplify presentation format
- Eliminate distractions
- Focus on core concepts

**Optimize Intrinsic Load**:
- Sequence from simple to complex
- Build on prior knowledge
- Use familiar contexts
- Connect to existing schemas

**Support Germane Load**:
- Encourage pattern recognition
- Promote schema construction
- Support meaning-making
- Foster deep processing

### Methodology Adaptations
**Switch Approaches**:
- From explanatory to discovery
- From theoretical to practical
- From systematic to exploratory
- From individual to collaborative

**Combine Methods**:
- Use multiple representations
- Blend different approaches
- Sequence complementary methods
- Layer simple and complex

## Real-Time Adaptation Protocol

### Step 1: Recognition (1-2 seconds)
- Notice adaptation signal
- Identify pattern type
- Assess severity level
- Consider context factors

### Step 2: Diagnosis (3-5 seconds)  
- Analyze likely causes
- Consider individual factors
- Review what's been tried
- Assess adaptation options

### Step 3: Decision (2-3 seconds)
- Select adaptation strategy
- Plan implementation approach
- Consider timing factors
- Prepare transition method

### Step 4: Implementation (immediate)
- Make adaptation smoothly
- Monitor learner response
- Adjust as needed
- Maintain learning flow

### Step 5: Evaluation (ongoing)
- Track adaptation effectiveness
- Note learner response changes
- Prepare further adaptations
- Update adaptation patterns

## Program Output Template

```markdown
## Instructional Adaptation Decision

### Adaptation Trigger: {Comprehension|Engagement|Cognitive Load|Challenge Level}

### Signal Analysis:
- **Observed Indicators**: {Specific behaviors or responses noticed}
- **Pattern Type**: {What type of difficulty or opportunity}
- **Severity Level**: {How urgent the need for adaptation}
- **Context Factors**: {Session conditions affecting situation}

### Root Cause Diagnosis:
- **Primary Cause**: {Most likely reason for adaptation need}
- **Contributing Factors**: {Other elements affecting situation}
- **Individual Considerations**: {Learner-specific factors}
- **Environmental Factors**: {Context conditions}

### Adaptation Strategy: {Representation|Interaction|Cognitive Load|Methodology}

### Implementation Plan:
- **Specific Changes**: {Exactly what will be modified}
- **Transition Method**: {How to make change smoothly}
- **Timing Approach**: {When and how quickly to adapt}
- **Monitoring Focus**: {What to watch for after adaptation}

### Constraint Adjustments:
- **Atomic Constraints**: {Which basic constraints need modification}
- **Molecular Protocols**: {How protocols should change}
- **Cellular Memory**: {What patterns to update}
- **Organ Coordination**: {How adaptation affects system}

### Success Indicators:
- **Immediate**: {Signs adaptation is working within minutes}
- **Short-term**: {Evidence within current session}
- **Learning Impact**: {Effect on constraint satisfaction}

### Fallback Options:
- **If adaptation ineffective**: {Alternative approaches to try}
- **If over-correction**: {How to dial back adaptation}
- **If learner resistance**: {How to address concerns}
```

## Integration with Constraint Architecture

### Atomic Constraint Adaptation
- **Dialogue Patterns**: Adjust questioning style and interaction
- **Questioning Syntax**: Modify question complexity and type
- **Understanding Validation**: Adapt validation methods
- **Pacing Control**: Adjust cognitive load and timing
- **Metacognitive Triggers**: Modify awareness-building approach

### Molecular Protocol Flexibility
- Switch between methodologies as needed
- Combine protocol elements creatively
- Adapt protocol steps to learner needs
- Maintain protocol integrity while being flexible

### Cellular Memory Updates
- Record adaptation patterns that work
- Note individual adaptation preferences
- Track adaptation timing effectiveness
- Update learner-specific adaptation strategies

### Organ-Level Adaptation
- Coordinate adaptations with Zelda
- Share adaptation insights with knowledge base
- Maintain system-wide learning from adaptations
- Support emergent adaptation intelligence

## Adaptation Patterns by Learner Type

### Visual Learners
- **Adaptations**: More diagrams, concept maps, visual analogies
- **Signals**: "Can you show me?" requests, better response to visuals
- **Constraints**: Emphasize visual representation patterns

### Auditory Learners  
- **Adaptations**: More discussion, verbal explanation, think-aloud
- **Signals**: Better engagement with verbal methods
- **Constraints**: Emphasize dialogue and verbal reasoning patterns

### Kinesthetic Learners
- **Adaptations**: More hands-on activities, movement, manipulation
- **Signals**: Restlessness, better engagement with active methods
- **Constraints**: Emphasize experiential and active patterns

### Analytical Learners
- **Adaptations**: More systematic approaches, logical sequences
- **Signals**: Requests for step-by-step processes
- **Constraints**: Emphasize structured and sequential patterns

### Intuitive Learners
- **Adaptations**: More big-picture connections, pattern exploration
- **Signals**: Better response to holistic approaches
- **Constraints**: Emphasize connection and insight patterns

## Common Adaptation Scenarios

### Scenario 1: Explanation Not Landing
- **Signal**: Blank stares, "I don't get it"
- **Diagnosis**: Abstraction level too high
- **Adaptation**: Use concrete examples, analogies
- **Validation**: Check understanding with simpler cases

### Scenario 2: Learner Bored/Disengaged
- **Signal**: Minimal participation, distraction
- **Diagnosis**: Under-challenge or irrelevance
- **Adaptation**: Increase complexity, show relevance
- **Validation**: Monitor engagement and energy

### Scenario 3: Information Overload
- **Signal**: Overwhelm expressions, performance decline
- **Diagnosis**: Too much too fast
- **Adaptation**: Chunk smaller, slow pace
- **Validation**: Check retention and confidence

### Scenario 4: Rapid Advancement
- **Signal**: Quick mastery, requests for more
- **Diagnosis**: Under-challenge
- **Adaptation**: Accelerate, add depth/breadth
- **Validation**: Ensure solid foundation maintained

This program enables responsive, learner-centered instruction that maintains optimal learning constraint satisfaction through dynamic adaptation.