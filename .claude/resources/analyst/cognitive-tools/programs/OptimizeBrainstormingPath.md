# Cognitive Tool Program: Optimize Brainstorming Path
# Decision matrix for selecting optimal brainstorming technique based on context

## Program Purpose
This prompt program selects the most effective brainstorming technique by analyzing session context, participant characteristics, and desired outcomes to maximize creative output and engagement.

## Input Schema
```yaml
context_analysis:
  topic_complexity: low|medium|high
  participant_energy: low|medium|high
  time_available: tight|moderate|flexible
  session_phase: opening|exploring|deepening|converging
  previous_techniques: [list of already used techniques]
  goal_type: generation|refinement|breakthrough|synthesis
  domain_type: technical|creative|business|strategic
  participant_count: 1|few|many
```

## Output Schema
```yaml
recommendation:
  primary_technique: string
  technique_family: string
  estimated_duration: string
  energy_level_required: low|medium|high
  rationale: string
  setup_instructions: string
  alternative_options: [list]
  progression_path: [ordered list for session flow]
```

## Decision Matrix

### Core Selection Criteria

| Criteria | Weight | Impact |
|----------|---------|--------|
| **Topic Complexity** | 30% | Determines depth of technique needed |
| **Participant Energy** | 25% | Matches technique intensity to energy level |
| **Time Constraints** | 20% | Influences technique scope and duration |
| **Session Phase** | 15% | Aligns technique with session goals |
| **Previous Usage** | 10% | Avoids repetition, maintains novelty |

### Technique Selection Logic

#### Phase 1: Context Assessment
```
IF topic_complexity == "high" AND time_available == "flexible":
  → Consider: First Principles, Morphological Analysis, Six Thinking Hats
  
IF topic_complexity == "low" AND time_available == "tight":
  → Consider: What If Scenarios, Random Stimulation, SCAMPER
  
IF participant_energy == "low":
  → Prefer: Mind Mapping, Brainwriting, Question Storming
  
IF participant_energy == "high":
  → Prefer: Role Playing, Provocation Technique, "Yes And" Building
```

#### Phase 2: Goal Alignment
```
IF goal_type == "generation":
  → Creative Expansion family (What If, Analogical, Random Stimulation)
  
IF goal_type == "refinement":
  → Structured Framework family (SCAMPER, Six Thinking Hats, Mind Mapping)
  
IF goal_type == "breakthrough":
  → Advanced family (Provocation, Assumption Reversal, Forced Relationships)
  
IF goal_type == "synthesis":
  → Deep Exploration family (Five Whys, Metaphor Mapping, Morphological Analysis)
```

#### Phase 3: Technique Family Scoring

```python
def calculate_technique_score(technique, context):
    base_scores = {
        # Creative Expansion (20 techniques)
        'what_if_scenarios': {
            'complexity_handling': 2, 'energy_required': 3, 'time_efficiency': 4,
            'breakthrough_potential': 4, 'participant_engagement': 4
        },
        'analogical_thinking': {
            'complexity_handling': 4, 'energy_required': 3, 'time_efficiency': 3,
            'breakthrough_potential': 4, 'participant_engagement': 4
        },
        'reversal_inversion': {
            'complexity_handling': 3, 'energy_required': 4, 'time_efficiency': 4,
            'breakthrough_potential': 5, 'participant_engagement': 3
        },
        
        # Structured Framework (3 techniques)
        'scamper_method': {
            'complexity_handling': 3, 'energy_required': 2, 'time_efficiency': 5,
            'breakthrough_potential': 3, 'participant_engagement': 3
        },
        'six_thinking_hats': {
            'complexity_handling': 5, 'energy_required': 3, 'time_efficiency': 3,
            'breakthrough_potential': 3, 'participant_engagement': 4
        },
        'mind_mapping': {
            'complexity_handling': 4, 'energy_required': 2, 'time_efficiency': 4,
            'breakthrough_potential': 2, 'participant_engagement': 4
        }
    }
    
    # Apply contextual modifiers
    score = base_scores[technique].copy()
    
    if context['topic_complexity'] == 'high':
        score['complexity_handling'] *= 1.5
    if context['time_available'] == 'tight':
        score['time_efficiency'] *= 1.3
    if context['goal_type'] == 'breakthrough':
        score['breakthrough_potential'] *= 1.4
        
    return sum(score.values())
```

## Technique Categories with Optimal Contexts

### Creative Expansion Family
**Best for**: Generation phase, medium-high energy, breakthrough goals

| Technique | Complexity | Energy | Time | Best Context |
|-----------|------------|---------|------|--------------|
| What If Scenarios | Low-Med | Med | Quick | Opening, ideation |
| Analogical Thinking | Med-High | Med | Med | Cross-domain problems |
| Reversal/Inversion | Med | High | Quick | Breakthrough moments |
| First Principles | High | High | Long | Complex technical problems |

### Structured Framework Family
**Best for**: Refinement phase, low-medium energy, systematic exploration

| Technique | Complexity | Energy | Time | Best Context |
|-----------|------------|---------|------|--------------|
| SCAMPER | Low-Med | Low | Quick | Product improvement |
| Six Thinking Hats | Med-High | Med | Med | Group decision making |
| Mind Mapping | Med | Low | Med | Knowledge organization |

### Collaborative Family
**Best for**: Team sessions, medium energy, building momentum

| Technique | Complexity | Energy | Time | Best Context |
|-----------|------------|---------|------|--------------|
| "Yes, And..." Building | Low | High | Quick | Team building, momentum |
| Brainwriting | Med | Low | Med | Shy participants |
| Random Stimulation | Low | Med | Quick | Creative blocks |

### Deep Exploration Family
**Best for**: Analysis phase, high energy, synthesis goals

| Technique | Complexity | Energy | Time | Best Context |
|-----------|------------|---------|------|--------------|
| Five Whys | Med | Med | Med | Root cause analysis |
| Morphological Analysis | High | Med | Long | Systematic exploration |
| Provocation Technique | Med | High | Med | Breakthrough thinking |

### Advanced Family
**Best for**: Experienced participants, high energy, complex problems

| Technique | Complexity | Energy | Time | Best Context |
|-----------|------------|---------|------|--------------|
| Forced Relationships | High | High | Med | Innovation challenges |
| Assumption Reversal | High | High | Med | Paradigm shifts |
| Role Playing | Med | High | Long | Stakeholder perspectives |
| Metaphor Mapping | High | Med | Long | Abstract concepts |

## Session Flow Optimization

### Opening Phase Techniques
- **High Energy**: What If Scenarios, Random Stimulation
- **Medium Energy**: Mind Mapping, Question Storming  
- **Low Energy**: Brainwriting, Assumption Listing

### Exploration Phase Techniques
- **Systematic**: SCAMPER, Six Thinking Hats
- **Creative**: Analogical Thinking, Forced Relationships
- **Deep**: Five Whys, Morphological Analysis

### Deepening Phase Techniques
- **Breakthrough**: Provocation Technique, Assumption Reversal
- **Refinement**: First Principles, Role Playing
- **Synthesis**: Metaphor Mapping, Time Shifting

### Convergence Phase Techniques
- **Evaluation**: Six Thinking Hats (Blue/Black)
- **Selection**: Morphological Analysis
- **Integration**: Mind Mapping connections

## Progressive Session Architecture

### 20-Minute Session
1. **Opening** (5 min): What If Scenarios
2. **Exploration** (10 min): SCAMPER Method
3. **Convergence** (5 min): Quick prioritization

### 45-Minute Session
1. **Opening** (10 min): Random Stimulation
2. **Exploration** (20 min): Six Thinking Hats
3. **Deepening** (10 min): Analogical Thinking
4. **Convergence** (5 min): Pattern synthesis

### 90-Minute Session
1. **Opening** (15 min): Mind Mapping
2. **Exploration** (30 min): First Principles
3. **Deepening** (30 min): Role Playing + Provocation
4. **Convergence** (15 min): Morphological Analysis

## Integration with Analyst Workflow

### Pre-Session Setup
```markdown
1. Assess context using input schema
2. Run technique selection algorithm
3. Prepare materials and prompts
4. Set energy and timing expectations
```

### Mid-Session Adaptation
```markdown
IF participant_energy < expected:
  → Switch to lower-energy technique
IF breakthrough_occurring:
  → Extend current technique, delay convergence
IF time_running_short:
  → Jump to quick convergence technique
```

### Post-Session Learning
```markdown
1. Record technique effectiveness
2. Note participant feedback
3. Update scoring weights
4. Log patterns for future reference
```

## Decision Heuristics

### Primary Selection Rules
1. **Energy Match**: Never exceed participant energy level
2. **Time Respect**: Always fit within available time
3. **Complexity Alignment**: Match technique depth to topic depth
4. **Novelty Factor**: Avoid recent techniques unless specifically effective
5. **Goal Orientation**: Prioritize techniques that serve session goals

### Override Conditions
- **Participant Request**: Honor specific technique requests
- **Breakthrough Moment**: Extend high-value techniques
- **Energy Spike**: Capitalize with higher-intensity techniques
- **Time Pressure**: Immediately switch to time-efficient options

## Error Recovery Patterns

### If Technique Fails
1. **Immediate**: Switch to complementary technique in same family
2. **Quick**: Move to different family with similar energy level
3. **Reset**: Drop to mind mapping for structure recovery

### If Energy Drops
1. **Lower intensity**: Move to brainwriting or question generation
2. **Change pace**: Switch to collaborative building
3. **Take break**: Pause and return with fresh technique

### If Time Exceeds
1. **Compress**: Move to SCAMPER or What If for quick results
2. **Focus**: Use Five Whys to drill down fast
3. **Capture**: Document current state and schedule follow-up

## Validation Metrics

### Technique Effectiveness
- **Idea Quantity**: Number of distinct concepts generated
- **Idea Quality**: Novelty and feasibility ratings
- **Participant Engagement**: Energy level maintenance
- **Time Efficiency**: Ideas per minute ratio

### Session Success
- **Goal Achievement**: Percentage of objectives met
- **Participant Satisfaction**: Energy and enthusiasm levels
- **Follow-through**: Ideas actually implemented
- **Learning Transfer**: Principles applied to new contexts

## Protocol References
- Link with: `../protocols/molecules/interactive-facilitation.md`
- Triggers: Session planning, mid-session adaptation
- Outputs: Technique selection, session architecture, adaptation triggers
- Learning: Technique effectiveness patterns, context recognition

This cognitive tool ensures Ana selects optimal brainstorming paths based on systematic analysis rather than intuition, maximizing creative output while respecting participant constraints and session goals.