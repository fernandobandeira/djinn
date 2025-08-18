# Adaptive Pacing Algorithm

## Resource Loading
- **Load when**: Teacher needs to adjust pacing based on learner performance
- **Loading method**: Read tool
- **Dependencies**: None (self-contained algorithm)
- **Used by**: Teacher command agent during interactive sessions

## Purpose
Dynamically adjust teaching pace based on learner comprehension, engagement, and cognitive load to maintain optimal learning conditions.

## Core Algorithm

```python
class AdaptivePacer:
    def __init__(self):
        self.learner_state = {
            "comprehension_velocity": 0.0,  # Rate of understanding (-1 to 1)
            "optimal_challenge_zone": [0.7, 0.85],  # Sweet spot for learning
            "current_difficulty": 0.5,  # 0 to 1 scale
            "engagement_level": 0.8,  # 0 to 1 scale
            "cognitive_load": 0.5,  # 0 to 1 scale
            "fatigue_level": 0.0,  # 0 to 1 scale
            "confidence": 0.7,  # 0 to 1 scale
            "streak": 0,  # Consecutive successes/failures
        }
        
        self.pacing_params = {
            "base_speed": 1.0,  # Normal pace multiplier
            "current_speed": 1.0,  # Active pace multiplier
            "min_speed": 0.3,  # Slowest pace
            "max_speed": 2.0,  # Fastest pace
            "adjustment_sensitivity": 0.1,  # How quickly to adjust
            "smoothing_factor": 0.3  # Smooth out rapid changes
        }
        
        self.history = {
            "performance": [],  # Recent performance metrics
            "adjustments": [],  # Pacing adjustments made
            "concepts": [],  # Concepts covered with timing
            "struggles": [],  # Where learner struggled
            "breakthroughs": []  # Where learner excelled
        }
    
    def adjust_pace(self, performance_metric):
        """Main pacing adjustment based on performance"""
        
        # Update comprehension velocity
        self.update_comprehension_velocity(performance_metric)
        
        # Calculate optimal pace adjustment
        adjustment = self.calculate_adjustment(performance_metric)
        
        # Apply adjustment with smoothing
        self.apply_adjustment(adjustment)
        
        # Check for special conditions
        self.handle_special_conditions()
        
        # Record adjustment
        self.record_adjustment(adjustment, performance_metric)
        
        return self.get_pacing_recommendation()
    
    def update_comprehension_velocity(self, performance):
        """Track rate of understanding change"""
        
        if len(self.history["performance"]) > 0:
            previous = self.history["performance"][-1]
            velocity = performance - previous
            
            # Smooth velocity calculation
            self.learner_state["comprehension_velocity"] = (
                self.learner_state["comprehension_velocity"] * 0.7 + 
                velocity * 0.3
            )
        
        self.history["performance"].append(performance)
        
        # Keep history manageable
        if len(self.history["performance"]) > 20:
            self.history["performance"].pop(0)
    
    def calculate_adjustment(self, performance):
        """Calculate pace adjustment needed"""
        
        adjustment = 0.0
        
        # Performance-based adjustment
        if performance < self.learner_state["optimal_challenge_zone"][0]:
            # Below optimal zone - slow down
            deficit = self.learner_state["optimal_challenge_zone"][0] - performance
            adjustment = -deficit * self.pacing_params["adjustment_sensitivity"]
            
        elif performance > self.learner_state["optimal_challenge_zone"][1]:
            # Above optimal zone - speed up
            surplus = performance - self.learner_state["optimal_challenge_zone"][1]
            adjustment = surplus * self.pacing_params["adjustment_sensitivity"]
        
        # Velocity-based modification
        if self.learner_state["comprehension_velocity"] < -0.2:
            # Rapid decline - extra slowdown
            adjustment -= 0.1
        elif self.learner_state["comprehension_velocity"] > 0.2:
            # Rapid improvement - can speed up more
            adjustment += 0.05
        
        # Engagement-based modification
        if self.learner_state["engagement_level"] < 0.5:
            # Low engagement - change pace significantly
            if performance < 0.7:
                adjustment -= 0.15  # Slow down if struggling
            else:
                adjustment += 0.15  # Speed up if bored
        
        # Fatigue adjustment
        if self.learner_state["fatigue_level"] > 0.7:
            adjustment = min(adjustment, -0.1)  # Force slowdown when tired
        
        return adjustment
    
    def apply_adjustment(self, adjustment):
        """Apply pace adjustment with smoothing"""
        
        # Calculate new speed
        target_speed = self.pacing_params["current_speed"] + adjustment
        
        # Smooth the transition
        self.pacing_params["current_speed"] = (
            self.pacing_params["current_speed"] * (1 - self.pacing_params["smoothing_factor"]) +
            target_speed * self.pacing_params["smoothing_factor"]
        )
        
        # Apply bounds
        self.pacing_params["current_speed"] = max(
            self.pacing_params["min_speed"],
            min(self.pacing_params["max_speed"], self.pacing_params["current_speed"])
        )
    
    def handle_special_conditions(self):
        """Handle special learning conditions"""
        
        # Streak handling
        if self.learner_state["streak"] > 3:
            # Consistent success - increase challenge
            self.learner_state["current_difficulty"] = min(
                1.0, 
                self.learner_state["current_difficulty"] + 0.1
            )
            self.pacing_params["current_speed"] = min(
                self.pacing_params["max_speed"],
                self.pacing_params["current_speed"] * 1.1
            )
            
        elif self.learner_state["streak"] < -3:
            # Consistent struggle - reduce challenge
            self.learner_state["current_difficulty"] = max(
                0.0,
                self.learner_state["current_difficulty"] - 0.1
            )
            self.pacing_params["current_speed"] = max(
                self.pacing_params["min_speed"],
                self.pacing_params["current_speed"] * 0.9
            )
            # Reset streak to prevent spiral
            self.learner_state["streak"] = 0
        
        # Cognitive overload protection
        if self.learner_state["cognitive_load"] > 0.85:
            self.pacing_params["current_speed"] *= 0.8
            self.add_break_recommendation()
        
        # Confidence boost needed
        if self.learner_state["confidence"] < 0.4:
            self.add_easy_win_recommendation()
    
    def get_pacing_recommendation(self):
        """Generate specific pacing recommendations"""
        
        speed = self.pacing_params["current_speed"]
        
        recommendations = {
            "pace_multiplier": speed,
            "content_adjustments": [],
            "interaction_style": "",
            "break_needed": False,
            "scaffolding_level": ""
        }
        
        # Pace-specific recommendations
        if speed < 0.5:
            # Very slow pace
            recommendations["content_adjustments"] = [
                "Break concepts into smaller pieces",
                "Add more examples",
                "Use simpler language",
                "Increase repetition"
            ]
            recommendations["interaction_style"] = "patient_supportive"
            recommendations["scaffolding_level"] = "high"
            
        elif speed < 0.8:
            # Slow pace
            recommendations["content_adjustments"] = [
                "Add clarifying examples",
                "Check understanding frequently",
                "Provide hints before answers"
            ]
            recommendations["interaction_style"] = "encouraging"
            recommendations["scaffolding_level"] = "moderate"
            
        elif speed > 1.5:
            # Fast pace
            recommendations["content_adjustments"] = [
                "Skip basic examples",
                "Introduce advanced concepts",
                "Encourage exploration"
            ]
            recommendations["interaction_style"] = "challenging"
            recommendations["scaffolding_level"] = "minimal"
            
        else:
            # Normal pace
            recommendations["interaction_style"] = "balanced"
            recommendations["scaffolding_level"] = "adaptive"
        
        # Special conditions
        if self.learner_state["fatigue_level"] > 0.7:
            recommendations["break_needed"] = True
        
        return recommendations
```

## Advanced Pacing Strategies

### Micro-Adjustments
Real-time pace adjustments within a single concept.

```python
class MicroPacer:
    """Fine-grained pacing within topics"""
    
    def __init__(self, parent_pacer):
        self.parent = parent_pacer
        self.micro_state = {
            "current_segment": "",
            "segment_start_time": 0,
            "expected_duration": 0,
            "actual_duration": 0,
            "comprehension_signals": []
        }
    
    def monitor_segment(self, segment_name, expected_duration):
        """Monitor progress through a segment"""
        
        self.micro_state["current_segment"] = segment_name
        self.micro_state["expected_duration"] = expected_duration
        self.micro_state["segment_start_time"] = time.time()
    
    def check_comprehension_signal(self, signal_type, signal_value):
        """Process comprehension signals"""
        
        signals = {
            "question_asked": -0.2,  # May indicate confusion
            "quick_response": 0.3,  # Good understanding
            "hesitation": -0.1,  # Uncertainty
            "self_correction": 0.2,  # Good metacognition
            "confusion_expressed": -0.3,  # Direct feedback
            "connection_made": 0.4,  # Excellent understanding
            "repetition_needed": -0.2,  # Didn't stick first time
        }
        
        if signal_type in signals:
            self.micro_state["comprehension_signals"].append(
                (signal_type, signals[signal_type])
            )
            
            # Immediate adjustment if strong signal
            if abs(signals[signal_type]) > 0.25:
                self.apply_micro_adjustment(signals[signal_type])
    
    def apply_micro_adjustment(self, signal_strength):
        """Apply immediate pace adjustment"""
        
        if signal_strength < 0:
            # Slow down
            self.add_elaboration()
            self.increase_wait_time()
        else:
            # Can speed up
            self.reduce_redundancy()
            self.decrease_wait_time()
    
    def add_elaboration(self):
        """Add more explanation"""
        return {
            "action": "elaborate",
            "suggestions": [
                "Add another example",
                "Rephrase explanation",
                "Use visual aid",
                "Connect to known concept"
            ]
        }
    
    def reduce_redundancy(self):
        """Skip unnecessary repetition"""
        return {
            "action": "streamline",
            "suggestions": [
                "Skip next example",
                "Move to application",
                "Increase complexity"
            ]
        }
```

### Adaptive Content Delivery

```yaml
content_adaptation:
  slow_pace_modifications:
    presentation:
      - "Smaller chunks"
      - "More visuals"
      - "Concrete examples"
      - "Step-by-step walkthrough"
    
    language:
      - "Simpler vocabulary"
      - "Shorter sentences"
      - "More repetition"
      - "Explicit connections"
    
    interaction:
      - "Frequent check-ins"
      - "Guided practice"
      - "Hints before solutions"
      - "Celebrate small wins"
    
    scaffolding:
      - "Pre-teaching vocabulary"
      - "Worked examples"
      - "Partially completed problems"
      - "Think-aloud modeling"
  
  fast_pace_modifications:
    presentation:
      - "Larger conceptual chunks"
      - "Abstract representations"
      - "Challenging examples"
      - "Self-directed exploration"
    
    language:
      - "Technical terminology"
      - "Complex explanations"
      - "Minimal repetition"
      - "Implicit connections"
    
    interaction:
      - "Open-ended questions"
      - "Independent practice"
      - "Problem-solving focus"
      - "Peer teaching opportunities"
    
    enrichment:
      - "Extension problems"
      - "Cross-domain connections"
      - "Research opportunities"
      - "Creative applications"
```

### Fatigue Detection and Management

```python
class FatigueManager:
    def __init__(self):
        self.fatigue_indicators = {
            "response_time": [],
            "error_rate": [],
            "engagement_scores": [],
            "time_elapsed": 0,
            "concepts_covered": 0
        }
        
        self.break_recommendations = {
            "micro_break": 2,  # 2 minutes
            "short_break": 5,  # 5 minutes
            "long_break": 15,  # 15 minutes
            "session_end": -1  # End session
        }
    
    def assess_fatigue(self):
        """Calculate current fatigue level"""
        
        fatigue_score = 0.0
        
        # Time-based fatigue
        if self.fatigue_indicators["time_elapsed"] > 30:
            fatigue_score += 0.2
        if self.fatigue_indicators["time_elapsed"] > 60:
            fatigue_score += 0.3
        if self.fatigue_indicators["time_elapsed"] > 90:
            fatigue_score += 0.3
        
        # Performance-based fatigue
        if len(self.fatigue_indicators["error_rate"]) >= 5:
            recent_errors = self.fatigue_indicators["error_rate"][-5:]
            if sum(recent_errors) / len(recent_errors) > 0.4:
                fatigue_score += 0.2
        
        # Engagement-based fatigue
        if len(self.fatigue_indicators["engagement_scores"]) >= 5:
            recent_engagement = self.fatigue_indicators["engagement_scores"][-5:]
            if sum(recent_engagement) / len(recent_engagement) < 0.5:
                fatigue_score += 0.2
        
        return min(1.0, fatigue_score)
    
    def recommend_break(self, fatigue_level):
        """Recommend appropriate break based on fatigue"""
        
        if fatigue_level < 0.3:
            return None
        elif fatigue_level < 0.5:
            return {
                "type": "micro_break",
                "duration": self.break_recommendations["micro_break"],
                "activity": "Quick stretch or deep breaths"
            }
        elif fatigue_level < 0.7:
            return {
                "type": "short_break",
                "duration": self.break_recommendations["short_break"],
                "activity": "Walk around, get water, rest eyes"
            }
        elif fatigue_level < 0.9:
            return {
                "type": "long_break",
                "duration": self.break_recommendations["long_break"],
                "activity": "Proper break, snack, change of scenery"
            }
        else:
            return {
                "type": "session_end",
                "duration": self.break_recommendations["session_end"],
                "activity": "End session, resume when refreshed"
            }
```

## Integration with Teacher Workflow

```python
# Initialize pacer at session start
pacer = AdaptivePacer()
micro_pacer = MicroPacer(pacer)
fatigue_manager = FatigueManager()

# During teaching
def teach_with_adaptive_pacing(concept):
    # Get initial pace recommendation
    pace_config = pacer.get_pacing_recommendation()
    
    # Adapt content based on pace
    adapted_content = adapt_content(concept, pace_config)
    
    # Monitor micro-signals during teaching
    for segment in adapted_content.segments:
        micro_pacer.monitor_segment(segment.name, segment.expected_duration)
        
        # Teach segment
        response = teach_segment(segment, pace_config)
        
        # Process comprehension signals
        for signal in response.signals:
            micro_pacer.check_comprehension_signal(signal.type, signal.value)
        
        # Check fatigue
        fatigue = fatigue_manager.assess_fatigue()
        if fatigue > 0.5:
            break_rec = fatigue_manager.recommend_break(fatigue)
            if break_rec:
                take_break(break_rec)
    
    # Adjust pace for next concept
    performance = assess_performance(response)
    pacer.adjust_pace(performance)
```

## Benefits

1. **Optimal Challenge**: Maintains learning in the zone of proximal development
2. **Reduced Frustration**: Slows down when struggling
3. **Sustained Engagement**: Speeds up when material is too easy
4. **Fatigue Management**: Prevents cognitive overload
5. **Personalized Experience**: Adapts to individual learning patterns
6. **Better Retention**: Appropriate pacing improves memory formation