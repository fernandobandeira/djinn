# Session Coherence Verification Protocol

## Purpose
Ensure smooth learning transitions and maintain momentum throughout teaching sessions.

## Coherence Verification Process

```python
def verify_session_coherence(current_topic, previous_topic, learner_state):
    """Ensure smooth learning transitions"""
    
    checks = {
        "prerequisite_coverage": check_prerequisites_met(current_topic, learner_state),
        "difficulty_progression": check_difficulty_appropriate(current_topic, previous_topic),
        "concept_connection": check_concepts_linked(current_topic, previous_topic),
        "momentum_maintained": check_engagement_level(learner_state),
        "cognitive_load": check_cognitive_load_balanced(learner_state)
    }
    
    coherence_score = calculate_coherence_score(checks)
    
    if not all(checks.values()):
        adjustments = suggest_bridge_content(checks, current_topic, previous_topic)
        return {"coherent": False, "adjustments": adjustments, "score": coherence_score}
    
    return {"coherent": True, "score": coherence_score}
```

## Coherence Dimensions

### 1. Prerequisite Coverage
Ensure all required knowledge is in place before proceeding.

```yaml
prerequisite_verification:
  check_methods:
    explicit_check:
      - "Quick review question about prerequisite"
      - "Can you remind me what X means?"
      - "Show me how you'd approach Y"
    
    implicit_observation:
      - "Watch for confusion indicators"
      - "Note questions that reveal gaps"
      - "Observe problem-solving approach"
    
    knowledge_probe:
      - "What if we changed this part?"
      - "Why do you think that works?"
      - "What would happen without X?"
  
  gap_detection:
    indicators:
      - "Confusion about terminology"
      - "Incorrect assumptions"
      - "Missing mental models"
      - "Skipping important steps"
    
    severity_levels:
      critical: "Cannot proceed without this"
      important: "Will struggle without this"
      helpful: "Would benefit from knowing"
  
  remediation:
    critical_gaps:
      - "Stop and teach prerequisite"
      - "Provide minimum viable knowledge"
      - "Reschedule topic for later"
    
    important_gaps:
      - "Quick review of concept"
      - "Provide reference material"
      - "Weave into current teaching"
    
    helpful_gaps:
      - "Note for future sessions"
      - "Provide optional resources"
      - "Mention in passing"
```

### 2. Difficulty Progression
Ensure appropriate challenge level progression.

```python
class DifficultyProgression:
    def __init__(self):
        self.difficulty_scale = {
            "trivial": 1,
            "easy": 2,
            "moderate": 3,
            "challenging": 4,
            "difficult": 5,
            "expert": 6
        }
        
        self.optimal_progression = 0.5  # Half-step increases
        self.max_jump = 1.5  # Maximum difficulty jump
        
    def check_progression(self, current_difficulty, previous_difficulty):
        """Check if difficulty progression is appropriate"""
        
        jump = current_difficulty - previous_difficulty
        
        if jump < 0:
            return {
                "appropriate": True,
                "type": "review",
                "note": "Reviewing or reinforcing"
            }
        elif jump <= self.optimal_progression:
            return {
                "appropriate": True,
                "type": "optimal",
                "note": "Perfect progression"
            }
        elif jump <= self.max_jump:
            return {
                "appropriate": True,
                "type": "stretch",
                "note": "Challenging but manageable"
            }
        else:
            return {
                "appropriate": False,
                "type": "gap",
                "note": "Too large a jump",
                "suggestion": self.suggest_bridge(previous_difficulty, current_difficulty)
            }
    
    def suggest_bridge(self, from_level, to_level):
        """Suggest intermediate content to bridge gap"""
        
        bridges = []
        current = from_level
        
        while current < to_level - self.optimal_progression:
            current += self.optimal_progression
            bridges.append({
                "difficulty": current,
                "content": self.generate_bridge_content(current)
            })
        
        return bridges
    
    def adjust_for_learner(self, base_difficulty, learner_performance):
        """Adjust difficulty based on learner performance"""
        
        if learner_performance > 0.9:
            # Doing very well, can increase difficulty
            return min(base_difficulty + 0.5, 6)
        elif learner_performance > 0.7:
            # Doing well, maintain pace
            return base_difficulty
        elif learner_performance > 0.5:
            # Struggling slightly, slow down
            return max(base_difficulty - 0.25, 1)
        else:
            # Struggling significantly, step back
            return max(base_difficulty - 0.5, 1)
```

### 3. Concept Connection
Ensure concepts build on each other logically.

```yaml
concept_linking:
  connection_types:
    builds_on:
      description: "New concept extends previous"
      transition_phrases:
        - "Building on what we just learned..."
        - "This extends our previous concept..."
        - "Remember X? Now we'll see how..."
    
    relates_to:
      description: "Concepts are related but independent"
      transition_phrases:
        - "This is similar to..."
        - "You might notice parallels with..."
        - "Like X, this also..."
    
    contrasts_with:
      description: "Concepts are different/opposite"
      transition_phrases:
        - "Unlike what we saw before..."
        - "In contrast to X..."
        - "While X does this, Y does..."
    
    applies:
      description: "Practical application of concept"
      transition_phrases:
        - "Let's apply what we learned..."
        - "Now we can use this to..."
        - "This helps us solve..."
  
  connection_strength:
    strong:
      - "Direct dependency"
      - "Same category"
      - "Sequential steps"
    
    moderate:
      - "Same domain"
      - "Shared principles"
      - "Common patterns"
    
    weak:
      - "Different domain"
      - "Tangential relation"
      - "Optional connection"
  
  bridging_strategies:
    no_connection:
      - "Acknowledge topic change"
      - "Provide clear transition"
      - "Reset mental context"
    
    weak_connection:
      - "Point out subtle links"
      - "Use analogies"
      - "Reference shared principles"
    
    strong_connection:
      - "Build directly"
      - "Reference previous examples"
      - "Extend mental models"
```

### 4. Momentum Maintenance
Keep engagement and energy levels appropriate.

```python
class MomentumTracker:
    def __init__(self):
        self.engagement_indicators = {
            "high": ["asking questions", "making connections", "eager responses"],
            "moderate": ["following along", "occasional questions", "completing tasks"],
            "low": ["delayed responses", "confusion", "distraction"],
            "lost": ["no response", "frustration", "giving up"]
        }
        
        self.momentum_score = 1.0  # Start at full momentum
        self.decay_rate = 0.05  # Natural momentum decay
        
    def check_momentum(self, learner_signals):
        """Assess current momentum level"""
        
        # Detect engagement level
        engagement = self.detect_engagement(learner_signals)
        
        # Update momentum score
        if engagement == "high":
            self.momentum_score = min(1.0, self.momentum_score + 0.1)
        elif engagement == "moderate":
            self.momentum_score -= self.decay_rate
        elif engagement == "low":
            self.momentum_score -= self.decay_rate * 2
        else:  # lost
            self.momentum_score -= self.decay_rate * 4
        
        return {
            "score": self.momentum_score,
            "level": self.get_level(),
            "action_needed": self.momentum_score < 0.6
        }
    
    def boost_momentum(self):
        """Strategies to restore momentum"""
        
        if self.momentum_score > 0.7:
            return ["Continue with current pace"]
        elif self.momentum_score > 0.5:
            return [
                "Add interactive element",
                "Use engaging example",
                "Quick success opportunity"
            ]
        elif self.momentum_score > 0.3:
            return [
                "Take a short break",
                "Switch to easier content",
                "Celebrate recent progress",
                "Change teaching modality"
            ]
        else:
            return [
                "Major reset needed",
                "Review and consolidate",
                "End session positively",
                "Schedule fresh start"
            ]
```

### 5. Cognitive Load Balance
Ensure learner isn't overwhelmed or under-stimulated.

```yaml
cognitive_load_management:
  load_types:
    intrinsic:
      description: "Complexity of material itself"
      indicators:
        - "Number of new concepts"
        - "Abstraction level"
        - "Interconnections"
      
    extraneous:
      description: "How material is presented"
      indicators:
        - "Clarity of explanation"
        - "Quality of examples"
        - "Presentation format"
    
    germane:
      description: "Effort to build schemas"
      indicators:
        - "Making connections"
        - "Pattern recognition"
        - "Transfer to long-term memory"
  
  load_assessment:
    optimal:
      total_load: "60-80% capacity"
      distribution:
        intrinsic: "40-50%"
        extraneous: "< 20%"
        germane: "20-30%"
    
    overloaded:
      indicators:
        - "Confusion"
        - "Slower responses"
        - "Missing connections"
        - "Frustration"
      
      interventions:
        - "Simplify presentation"
        - "Break into smaller chunks"
        - "Remove non-essential info"
        - "Provide worked examples"
    
    underloaded:
      indicators:
        - "Boredom"
        - "Distraction"
        - "Impatience"
      
      interventions:
        - "Increase pace"
        - "Add complexity"
        - "Introduce challenges"
        - "Encourage exploration"
```

## Session Flow Optimization

```python
def optimize_session_flow(session_plan, learner_profile):
    """Optimize session for maximum coherence"""
    
    optimizer = SessionOptimizer(learner_profile)
    
    # 1. Sequence topics optimally
    optimized_sequence = optimizer.sequence_topics(session_plan.topics)
    
    # 2. Insert bridges where needed
    with_bridges = optimizer.add_bridges(optimized_sequence)
    
    # 3. Add momentum boosters
    with_boosters = optimizer.add_momentum_points(with_bridges)
    
    # 4. Balance cognitive load
    balanced = optimizer.balance_cognitive_load(with_boosters)
    
    # 5. Add checkpoints
    final_plan = optimizer.add_checkpoints(balanced)
    
    return final_plan

class SessionOptimizer:
    def sequence_topics(self, topics):
        """Order topics for best flow"""
        
        # Build dependency graph
        graph = self.build_dependency_graph(topics)
        
        # Topological sort with optimization
        sequence = self.topological_sort_optimized(graph)
        
        return sequence
    
    def add_bridges(self, sequence):
        """Add bridging content between topics"""
        
        bridged = []
        for i, topic in enumerate(sequence):
            bridged.append(topic)
            
            if i < len(sequence) - 1:
                next_topic = sequence[i + 1]
                connection = self.assess_connection(topic, next_topic)
                
                if connection["strength"] < 0.5:
                    bridge = self.create_bridge(topic, next_topic)
                    bridged.append(bridge)
        
        return bridged
    
    def add_momentum_points(self, sequence):
        """Insert engagement boosters"""
        
        with_boosters = []
        time_since_booster = 0
        
        for item in sequence:
            with_boosters.append(item)
            time_since_booster += item.estimated_duration
            
            if time_since_booster > 15:  # Every 15 minutes
                booster = self.create_booster(item.topic)
                with_boosters.append(booster)
                time_since_booster = 0
        
        return with_boosters
```

## Coherence Report Template

```markdown
## Session Coherence Report

### Overall Coherence Score: [X.X/10]

### Dimension Scores:
- Prerequisites Coverage: [✓/✗] [Score]
- Difficulty Progression: [✓/✗] [Score]
- Concept Connections: [✓/✗] [Score]
- Momentum Maintained: [✓/✗] [Score]
- Cognitive Load Balanced: [✓/✗] [Score]

### Identified Issues:
1. [Issue description]
   - Impact: [High/Medium/Low]
   - Suggested fix: [Intervention]

### Optimization Suggestions:
- [Suggestion 1]
- [Suggestion 2]

### Bridge Content Needed:
- Between [Topic A] and [Topic B]: [Bridge description]

### Momentum Boost Points:
- After [Topic]: [Booster activity]

### Cognitive Load Adjustments:
- [Adjustment needed]
```

## Integration with Teacher Workflow

```python
# During session planning
def plan_coherent_session(learning_objective):
    # Decompose objective
    decomposed = decompose_objective(learning_objective)
    
    # Check coherence
    coherence = verify_session_coherence(decomposed)
    
    if not coherence["coherent"]:
        # Apply suggested adjustments
        decomposed = apply_adjustments(decomposed, coherence["adjustments"])
    
    return decomposed

# During session execution
def teach_with_coherence_monitoring():
    for topic in session_topics:
        # Check coherence before teaching
        pre_check = verify_topic_coherence(topic, previous_topic, learner_state)
        
        if not pre_check["coherent"]:
            # Add bridge content
            teach_bridge(pre_check["bridge"])
        
        # Teach topic
        teach_topic(topic)
        
        # Monitor coherence during teaching
        if detect_coherence_break():
            apply_real_time_adjustment()
```