# Insight Synthesis Flow Protocol

## Purpose
Extract and synthesize patterns during analysis sessions to provide coherent insights from complex discussions.

## Activation Triggers
- Multiple related ideas emerge during discussion
- Pattern recognition needed across conversation threads
- Session reaching conclusion phase and requires synthesis
- User asks for "what does this all mean?"
- Information overload requires structure

## Process Flow

### 1. Pattern Detection
```yaml
action: Monitor conversation for:
  - Recurring themes
  - Contradictory viewpoints
  - Emerging frameworks
  - User preference patterns
```

### 2. Delegation to Insight-Synthesizer
```yaml
delegate:
  sub_agent: insight-synthesizer
  context:
    - session_transcript: full conversation
    - key_points: extracted highlights
    - user_goals: stated objectives
    - patterns_detected: preliminary insights
```

### 3. Pattern Presentation
```yaml
format:
  structured_insights:
    - Core themes identified
    - Key tensions and trade-offs
    - Emerging opportunities
    - Recommended next steps
  visual_aids:
    - Mind maps for complex relationships
    - Decision trees for choice points
    - Priority matrices for options
```

### 4. Feedback Integration
```yaml
refinement_loop:
  - Present initial synthesis
  - Gather user reactions
  - Identify gaps or misunderstandings
  - Refine and re-present
  - Confirm accuracy before proceeding
```

## Success Criteria
- User expresses "aha moments"
- Complex information becomes actionable
- Clear next steps emerge
- User can articulate insights in their own words

## Integration Points
- Triggers after deep analysis sessions
- Connects to decision-making protocols
- Feeds into action planning workflows