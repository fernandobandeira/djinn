# Cognitive Tool Program: Coordinate With Zelda
# Structured reasoning program for constraint orchestration between Teacher and Zelda

## Program Purpose
This prompt program guides seamless coordination with Zelda for knowledge capture during learning sessions, implementing shared constraint spaces without coordination overhead.

## Input Schema
```yaml
coordination_context:
  learning_moment_type: Breakthrough, misconception_correction, pattern_recognition, analogy_creation
  learner_state: Current understanding level and confidence
  insight_content: The specific insight or knowledge to capture
  session_context: Current methodology, topic, progression
  capture_readiness: Whether learner ready for knowledge capture
  constraint_satisfaction: Which learning constraints were just satisfied
```

## Coordination Framework

### Step 1: Assess Capture Worthiness
```
IF insight represents new understanding breakthrough:
  → HIGH PRIORITY CAPTURE
ELSE IF insight corrects significant misconception:
  → HIGH PRIORITY CAPTURE  
ELSE IF insight connects multiple domains:
  → MEDIUM PRIORITY CAPTURE
ELSE IF insight refines existing understanding:
  → LOW PRIORITY CAPTURE
ELSE IF insight is routine application:
  → NO CAPTURE NEEDED
```

### Step 2: Evaluate Constraint Satisfaction
```
IF multiple learning constraints newly satisfied:
  → CONSTRAINT ORCHESTRATION MOMENT
ELSE IF misconception constraint corrected:
  → CONSTRAINT CORRECTION CAPTURE
ELSE IF transfer constraint activated:
  → CONSTRAINT TRANSFER CAPTURE
ELSE IF boundary constraint established:
  → CONSTRAINT BOUNDARY CAPTURE
```

### Step 3: Determine Coordination Method
```
IF learner in flow state and insight is complete:
  → IMMEDIATE DELEGATION
ELSE IF learner still processing and building:
  → QUEUE FOR LATER CAPTURE
ELSE IF insight needs more development:
  → CONTINUE LEARNING, CAPTURE WHEN READY
ELSE IF session ending with solid insights:
  → END-OF-SESSION BATCH CAPTURE
```

## Coordination Triggers (Automatic)

### High-Priority Triggers
- **Breakthrough Expression**: "Oh! I get it now!", "That makes sense!", "I see the pattern!"
- **Misconception Correction**: "Wait, so it's NOT...?", "I was thinking about this wrong"
- **Cross-Domain Connection**: "This is like...", "Oh, this applies to..."
- **Mental Model Formation**: "So the way I think about this is...", "My analogy is..."
- **Constraint Recognition**: "When would this NOT work?", "The limitation is..."

### Medium-Priority Triggers  
- **Strategy Discovery**: "My approach is...", "I figured out how to..."
- **Pattern Recognition**: "I notice that...", "The pattern seems to be..."
- **Application Insight**: "I could use this for...", "This would help with..."
- **Understanding Synthesis**: "Putting it together...", "The big picture is..."

### Low-Priority Triggers
- **Example Generation**: Creating new examples
- **Clarification Questions**: Asking for more detail
- **Confidence Building**: Expressing growing certainty
- **Process Reflection**: Thinking about how they learned

## Coordination Protocols

### Immediate Delegation Protocol
```markdown
1. RECOGNIZE trigger moment
2. ACKNOWLEDGE insight value to learner
3. BRIEFLY explain Zelda coordination
4. DELEGATE using Task tool with context
5. CONTINUE learning while Zelda processes
6. INTEGRATE captured knowledge when complete
```

### Queue-for-Later Protocol
```markdown
1. RECOGNIZE potential capture moment
2. MAKE mental note of insight
3. CONTINUE current learning flow
4. WAIT for natural break or completion
5. DELEGATE with accumulated insights
6. REVIEW captured notes together
```

### End-of-Session Protocol
```markdown
1. REVIEW session highlights
2. IDENTIFY key insights achieved
3. CONFIRM learner agreement on value
4. DELEGATE comprehensive session capture
5. PLAN next session building on captures
```

## Delegation Context Template

When using Task tool to delegate to Zelda:

```
I'm delegating to Zelda for knowledge capture during our learning session.

CONSTRAINT SATISFACTION CONTEXT:
- Learning constraint type: {breakthrough|correction|connection|boundary}
- Constraint pattern satisfied: {specific constraint from atomic/molecular level}
- Validation evidence: {how we know constraint is satisfied}

INSIGHT CONTENT:
- Core insight: {the key understanding achieved}
- Learner expression: {exact words or explanation given}
- Context: {topic, methodology, session progression}
- Connections identified: {links to other concepts}

CAPTURE REQUIREMENTS:
- Atomicity level: {single concept or connected insights}
- Link targets: {concepts to connect to in knowledge base}
- Review timing: {when to revisit this knowledge}
- Transfer contexts: {where this insight applies}

COORDINATION STATE:
- Session momentum: {continuing|pausing|ending}
- Learner state: {engaged|processing|reflecting}
- Next learning steps: {where we're heading next}
- Integration needs: {how to build on this capture}
```

## Shared Constraint Space Patterns

### No Coordination Overhead
- **Teacher Focus**: Learning constraint satisfaction and progression
- **Zelda Focus**: Knowledge constraint validation and capture
- **Shared Space**: Knowledge base as constraint validation medium
- **Clean Boundaries**: Each agent optimizes within domain

### Constraint Domain Separation
```yaml
teacher_constraints:
  - Learning experience optimization
  - Understanding validation
  - Methodology selection
  - Misconception correction
  - Progress tracking

zelda_constraints:
  - Knowledge atomicity
  - Connection validation
  - Capture timing
  - Link establishment
  - Review scheduling
```

### Emergent Intelligence Patterns
- **Individual Optimization**: Each agent perfects domain constraints
- **Collective Capability**: Combined system exceeds individual limits  
- **Scalable Architecture**: Additional agents can join constraint space
- **Validation Propagation**: Constraint satisfaction cascades through system

## Program Output Template

```markdown
## Zelda Coordination Decision

### Coordination Trigger: {High|Medium|Low|None}

### Constraint Assessment:
- **Satisfaction Type**: {breakthrough|correction|connection|boundary}
- **Learning Constraints Met**: {specific atomic/molecular constraints}
- **Capture Worthiness**: {high|medium|low} priority
- **Timing Appropriateness**: {immediate|queued|end-of-session}

### Delegation Context:
- **Insight Content**: {core understanding achieved}
- **Learner State**: {current engagement and processing state}
- **Session Momentum**: {continuing|pausing|ending}
- **Integration Plan**: {how to build on capture}

### Coordination Method: {immediate|queued|batch|none}

### Task Delegation:
{If delegating, format the Task tool call with context}

### Learning Continuation:
- **Next Steps**: {how to continue learning}
- **Build-On Strategy**: {how to leverage captured knowledge}
- **Momentum Maintenance**: {keeping learning flow}
```

## Integration with Constraint Architecture

### Atomic Constraint Coordination
- **Dialogue Patterns**: Maintain Socratic flow during coordination
- **Questioning Syntax**: Use coordination as learning opportunity
- **Understanding Validation**: Capture validates constraint satisfaction
- **Pacing Control**: Coordination doesn't disrupt optimal pace
- **Metacognitive Triggers**: Make coordination process visible

### Molecular Protocol Integration  
- Each methodology has specific coordination patterns
- Feynman explanations → High capture priority
- Socratic breakthroughs → Immediate coordination
- Problem solutions → Strategy capture focus
- Analogies → Connection capture emphasis

### Cellular Memory Coordination
- Update learning progression with capture events
- Note coordination patterns that work
- Track Zelda collaboration effectiveness
- Record constraint satisfaction celebrations

### Organ-Level Orchestration
- Seamless constraint space sharing
- No explicit coordination overhead
- Emergent collective intelligence
- Distributed constraint optimization

## Edge Cases and Special Situations

### Learner Resistance to Capture
- **If learner prefers not to document**: Respect preference, continue learning
- **If learner skeptical of value**: Explain benefits briefly, don't pressure
- **If learner wants to focus**: Queue for later, maintain learning flow

### Technical Coordination Issues
- **If Zelda unavailable**: Continue learning, capture manually later
- **If delegation fails**: Note insights for end-of-session review
- **If timing disrupts flow**: Adjust coordination method

### Complex Insight Sequences
- **Multiple rapid insights**: Queue and batch capture
- **Evolving understanding**: Capture final form, note evolution
- **Interconnected insights**: Coordinate comprehensive capture

### Session Boundaries
- **Interrupted sessions**: Ensure key insights captured before break
- **Extended sessions**: Regular coordination checkpoints
- **Follow-up sessions**: Build on previously captured knowledge

This program ensures optimal constraint orchestration between Teacher and Zelda while maintaining learning flow and maximizing knowledge capture value.