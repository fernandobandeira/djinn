# Adaptive Pacing Protocol for Teacher

## Purpose
Dynamically adjust learning pace based on learner's constraint satisfaction signals

## Constraint Satisfaction Indicators

### Positive Signals (Accelerate Pace)
- Quick comprehension responses (< 2 seconds)
- Correct predictions about outcomes
- Spontaneous pattern recognition
- Self-initiated connections to other concepts
- "This is easy" or "I get it" expressions
- Asking advanced follow-up questions

### Negative Signals (Decelerate Pace)
- Long pauses before responses (> 5 seconds)
- Incorrect predictions or misconceptions
- Requests for clarification
- "I'm confused" or "Wait, what?" expressions
- Repeated errors on similar concepts
- Cognitive overload indicators

## Pacing Adjustment Algorithm

```yaml
pace_levels:
  sprint: 
    chunk_size: large
    wait_time: minimal
    complexity: high
    examples: 1-2
  
  jog:
    chunk_size: medium
    wait_time: moderate
    complexity: medium
    examples: 2-3
  
  walk:
    chunk_size: small
    wait_time: extended
    complexity: low
    examples: 3-5
  
  pause:
    chunk_size: micro
    wait_time: unlimited
    complexity: fundamental
    examples: 5+
```

## Dynamic Adjustment Protocol

1. **Monitor constraint satisfaction rate**
   - Track successful concept validations per minute
   - Measure error rate in understanding checks
   - Analyze response latency patterns

2. **Calculate pace score**
   ```
   pace_score = (success_rate * 0.4) + 
                (response_speed * 0.3) + 
                (connection_depth * 0.3)
   ```

3. **Adjust dynamically**
   - Score > 0.8: Consider acceleration
   - Score 0.5-0.8: Maintain current pace
   - Score < 0.5: Consider deceleration
   - Score < 0.3: Pause and reassess

## Constraint Negotiation Triggers

### When to Negotiate Pace
- Learner explicitly requests different pace
- Consistent mismatch between pace and performance
- Time constraints conflict with depth requirements
- Energy level changes detected

### Negotiation Dialogue
```
"I notice you're [grasping concepts quickly/taking time to process].
Would you prefer to:
1. Speed up and cover more ground
2. Maintain this pace for comfort
3. Slow down for deeper understanding
4. Take a break and resume later"
```

## Integration Points

- Links to: `/resources/teacher/cells/memory/learning-progression.yaml`
- Updates: Session pacing history
- Triggers: Zettelkasten capture at optimal moments
- Validates: Understanding before progression