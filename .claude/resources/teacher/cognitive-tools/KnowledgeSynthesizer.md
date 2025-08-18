# Knowledge Synthesizer

## Resource Loading
- **Load when**: End of teaching session, need to synthesize learnings
- **Loading method**: Read tool
- **Dependencies**: Zettelkasten sub-agents for note creation
- **Used by**: Teacher command agent, zettelkasten-synthesizer

## Purpose
Extract, organize, and synthesize knowledge from learning sessions to create structured understanding and actionable next steps.

## Synthesis Framework

### Extraction Phase
Identify and capture key concepts from the learning session.

```yaml
synthesis_process:
  extraction:
    key_concepts:
      identification:
        - "Core ideas taught"
        - "Supporting concepts"
        - "Practical applications"
        - "Common patterns"
      
      capture_methods:
        - "Definition extraction"
        - "Example collection"
        - "Relationship mapping"
        - "Property listing"
    
    breakthrough_moments:
      indicators:
        - "Sudden understanding"
        - "Connection made"
        - "Misconception corrected"
        - "Pattern recognized"
      
      documentation:
        - "What was the insight?"
        - "What triggered it?"
        - "How to replicate?"
        - "Why it matters"
    
    analogies_examples:
      effective_ones:
        - "Which resonated?"
        - "Why did they work?"
        - "Can be reused when?"
      
      ineffective_ones:
        - "Which confused?"
        - "Why did they fail?"
        - "How to improve?"
```

### Relationship Mapping
Connect concepts to build a knowledge network.

```python
class KnowledgeMapper:
    def __init__(self):
        self.knowledge_graph = {}
        self.relationship_types = [
            "depends_on",
            "enables",
            "similar_to",
            "contrasts_with",
            "applies_to",
            "composed_of",
            "generalizes",
            "specializes"
        ]
    
    def map_session_knowledge(self, session_content):
        """Create knowledge map from session"""
        
        # Extract concepts
        concepts = self.extract_concepts(session_content)
        
        # Map relationships
        for concept_a in concepts:
            self.knowledge_graph[concept_a] = {}
            
            for concept_b in concepts:
                if concept_a != concept_b:
                    relationship = self.identify_relationship(concept_a, concept_b)
                    if relationship:
                        self.knowledge_graph[concept_a][concept_b] = relationship
        
        return self.knowledge_graph
    
    def identify_relationship(self, concept_a, concept_b):
        """Determine relationship between concepts"""
        
        # Check various relationship patterns
        if self.is_prerequisite(concept_a, concept_b):
            return {"type": "depends_on", "strength": 0.9}
        
        if self.shares_properties(concept_a, concept_b):
            return {"type": "similar_to", "strength": 0.7}
        
        if self.is_application(concept_a, concept_b):
            return {"type": "applies_to", "strength": 0.8}
        
        # Continue for other relationship types...
        return None
    
    def find_knowledge_gaps(self):
        """Identify missing connections or concepts"""
        
        gaps = []
        
        # Check for orphaned concepts
        for concept in self.knowledge_graph:
            if not self.knowledge_graph[concept]:
                gaps.append({
                    "type": "orphaned",
                    "concept": concept,
                    "suggestion": "Connect to related concepts"
                })
        
        # Check for missing links
        expected_connections = self.predict_connections()
        actual_connections = self.get_all_connections()
        
        missing = expected_connections - actual_connections
        for connection in missing:
            gaps.append({
                "type": "missing_link",
                "from": connection[0],
                "to": connection[1],
                "suggestion": f"Explore relationship between {connection[0]} and {connection[1]}"
            })
        
        return gaps
```

### Synthesis Generation
Create comprehensive summaries and insights.

```python
class SynthesisGenerator:
    def __init__(self, session_data):
        self.session_data = session_data
        self.synthesis_components = []
    
    def generate_synthesis(self):
        """Create complete synthesis of learning session"""
        
        synthesis = {
            "session_summary": self.create_session_summary(),
            "key_learnings": self.extract_key_learnings(),
            "knowledge_map": self.build_knowledge_map(),
            "learning_roadmap": self.generate_roadmap(),
            "next_topics": self.suggest_next_topics(),
            "review_schedule": self.schedule_reviews(),
            "practice_exercises": self.create_practice_set()
        }
        
        return synthesis
    
    def create_session_summary(self):
        """Summarize what was learned"""
        
        summary = {
            "topic": self.session_data["topic"],
            "duration": self.session_data["duration"],
            "concepts_covered": len(self.session_data["concepts"]),
            "understanding_level": self.calculate_understanding(),
            "key_achievements": self.identify_achievements(),
            "struggle_points": self.identify_struggles()
        }
        
        # Generate narrative summary
        narrative = self.generate_narrative(summary)
        summary["narrative"] = narrative
        
        return summary
    
    def extract_key_learnings(self):
        """Extract the most important learnings"""
        
        learnings = []
        
        # Rank concepts by importance
        ranked_concepts = self.rank_by_importance(self.session_data["concepts"])
        
        for concept in ranked_concepts[:5]:  # Top 5 key learnings
            learning = {
                "concept": concept["name"],
                "understanding": concept["understanding_level"],
                "importance": concept["importance_score"],
                "key_insight": self.extract_insight(concept),
                "application": self.identify_application(concept),
                "remember_as": self.create_memory_hook(concept)
            }
            learnings.append(learning)
        
        return learnings
    
    def build_knowledge_map(self):
        """Create visual/structural knowledge representation"""
        
        return {
            "central_concept": self.identify_central_concept(),
            "supporting_concepts": self.identify_supporting_concepts(),
            "relationships": self.map_relationships(),
            "hierarchy": self.build_hierarchy(),
            "cross_connections": self.find_cross_connections()
        }
    
    def generate_roadmap(self):
        """Create learning path forward"""
        
        roadmap = {
            "immediate_next": self.get_immediate_next_steps(),
            "short_term": self.plan_week_ahead(),
            "medium_term": self.plan_month_ahead(),
            "long_term": self.plan_learning_journey(),
            "milestones": self.define_milestones()
        }
        
        return roadmap
    
    def suggest_next_topics(self):
        """Suggest what to learn next"""
        
        suggestions = []
        
        # Based on current progress
        if self.session_data["understanding_level"] > 0.8:
            # Ready for advanced topics
            suggestions.extend(self.get_advanced_topics())
        elif self.session_data["understanding_level"] > 0.6:
            # Ready for related topics
            suggestions.extend(self.get_related_topics())
        else:
            # Need reinforcement
            suggestions.extend(self.get_reinforcement_topics())
        
        # Based on gaps
        gaps = self.identify_knowledge_gaps()
        suggestions.extend(self.get_gap_filling_topics(gaps))
        
        # Based on interests
        interests = self.detect_interests()
        suggestions.extend(self.get_interest_based_topics(interests))
        
        return self.prioritize_suggestions(suggestions)
    
    def schedule_reviews(self):
        """Create spaced repetition schedule"""
        
        schedule = []
        
        # For each concept learned
        for concept in self.session_data["concepts"]:
            difficulty = concept["difficulty"]
            understanding = concept["understanding_level"]
            
            # Calculate review intervals
            intervals = self.calculate_spaced_intervals(difficulty, understanding)
            
            for i, interval in enumerate(intervals):
                schedule.append({
                    "concept": concept["name"],
                    "review_number": i + 1,
                    "scheduled_for": interval,
                    "review_type": self.determine_review_type(i),
                    "estimated_duration": self.estimate_review_duration(concept, i)
                })
        
        return sorted(schedule, key=lambda x: x["scheduled_for"])
    
    def calculate_spaced_intervals(self, difficulty, understanding):
        """Calculate optimal review intervals"""
        
        # Base intervals (in days)
        base_intervals = [1, 3, 7, 14, 30, 90]
        
        # Adjust based on difficulty and understanding
        adjustment_factor = (1.0 / difficulty) * understanding
        
        adjusted_intervals = []
        for interval in base_intervals:
            adjusted = int(interval * adjustment_factor)
            adjusted_intervals.append(max(1, adjusted))  # Minimum 1 day
        
        return adjusted_intervals
```

### Learning Roadmap Generation

```yaml
roadmap_structure:
  immediate_actions:
    within_24_hours:
      - "Review session notes"
      - "Complete practice exercises"
      - "Write summary in own words"
      - "Attempt challenge problem"
    
    within_48_hours:
      - "First spaced review"
      - "Apply to real problem"
      - "Teach someone else"
      - "Create personal examples"
  
  weekly_plan:
    day_1_2:
      - "Consolidation phase"
      - "Practice fundamentals"
      - "Address confusion points"
    
    day_3_4:
      - "Extension phase"
      - "Explore related concepts"
      - "Increase complexity"
    
    day_5_7:
      - "Application phase"
      - "Real-world problems"
      - "Creative exploration"
      - "Second review"
  
  monthly_progression:
    week_1:
      focus: "Mastery of current topic"
      activities:
        - "Deep practice"
        - "Edge case exploration"
        - "Pattern recognition"
    
    week_2:
      focus: "Integration with prior knowledge"
      activities:
        - "Connect to other topics"
        - "Build mental models"
        - "Create concept maps"
    
    week_3:
      focus: "Advanced applications"
      activities:
        - "Complex problems"
        - "Multi-concept integration"
        - "Creative solutions"
    
    week_4:
      focus: "Teaching and synthesis"
      activities:
        - "Create teaching materials"
        - "Write explanations"
        - "Help others learn"
```

### Integration with Zettelkasten

```python
def create_learning_notes(synthesis):
    """Create Zettelkasten notes from synthesis"""
    
    notes = []
    
    # Create permanent notes for key insights
    for learning in synthesis["key_learnings"]:
        note = {
            "type": "permanent",
            "title": learning["concept"],
            "content": learning["key_insight"],
            "connections": find_related_notes(learning["concept"]),
            "tags": generate_tags(learning),
            "review_schedule": learning.get("review_schedule", [])
        }
        notes.append(note)
    
    # Create hub note for session
    hub_note = {
        "type": "hub",
        "title": f"Learning Session: {synthesis['session_summary']['topic']}",
        "content": synthesis["session_summary"]["narrative"],
        "connected_notes": [n["title"] for n in notes],
        "learning_roadmap": synthesis["learning_roadmap"],
        "next_topics": synthesis["next_topics"]
    }
    notes.append(hub_note)
    
    # Create literature notes for sources
    for source in synthesis.get("sources", []):
        lit_note = {
            "type": "literature",
            "source": source["reference"],
            "key_points": source["key_points"],
            "my_thoughts": source["reflections"],
            "connections": source["related_concepts"]
        }
        notes.append(lit_note)
    
    return notes
```

## Example Synthesis Output

```yaml
example_synthesis:
  session_summary:
    topic: "Recursion in Programming"
    duration: "90 minutes"
    concepts_covered: 8
    understanding_level: 0.85
    narrative: |
      Successfully learned recursion fundamentals including base cases,
      recursive calls, and call stack visualization. Breakthrough moment
      when factorial example clicked. Some struggle with tree traversal
      complexity but foundation is solid.
  
  key_learnings:
    - concept: "Base Case"
      understanding: 0.95
      key_insight: "Every recursion needs a stopping condition"
      application: "Prevents infinite loops"
      remember_as: "The emergency brake of recursion"
    
    - concept: "Call Stack"
      understanding: 0.80
      key_insight: "Each call creates new scope on stack"
      application: "Understanding memory usage"
      remember_as: "Stack of plates - LIFO"
  
  learning_roadmap:
    immediate_next:
      - "Practice writing 3 recursive functions"
      - "Trace execution of fibonacci(5)"
      - "Convert one iterative solution to recursive"
    
    short_term:
      - "Learn tail recursion optimization"
      - "Explore tree traversal patterns"
      - "Understand recursion vs iteration trade-offs"
    
    medium_term:
      - "Dynamic programming with memoization"
      - "Backtracking algorithms"
      - "Recursive data structures"
  
  next_topics:
    - title: "Tree Traversal Algorithms"
      reason: "Natural progression from recursion"
      prerequisites_met: true
      estimated_time: "2 hours"
    
    - title: "Dynamic Programming"
      reason: "Optimization of recursive solutions"
      prerequisites_met: false
      missing: ["Memoization concept"]
      estimated_time: "4 hours"
  
  review_schedule:
    - concept: "Recursion basics"
      review_date: "Tomorrow"
      duration: "15 minutes"
      type: "Quick recall"
    
    - concept: "Call stack visualization"
      review_date: "In 3 days"
      duration: "20 minutes"
      type: "Practice problems"
    
    - concept: "Complex recursion"
      review_date: "In 1 week"
      duration: "30 minutes"
      type: "New applications"
```

## Benefits

1. **Structured Knowledge**: Organized understanding, not scattered facts
2. **Clear Progress Path**: Know exactly what to learn next
3. **Optimized Retention**: Spaced repetition built-in
4. **Gap Identification**: Find and fill knowledge holes
5. **Personalized Journey**: Adapted to individual learning patterns
6. **Zettelkasten Integration**: Permanent knowledge capture