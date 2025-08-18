# Problem Decomposer - Cognitive Tool for Teacher

## Purpose
Systematically break down complex learning objectives into atomic, digestible steps with clear dependencies and checkpoints.

## Decomposition Process

### Step 1: Parse Learning Objective
Identify and extract the core components of what needs to be learned.

```yaml
objective_parsing:
  core_concepts:
    identification:
      - "What is the main concept to learn?"
      - "What are the supporting concepts?"
      - "What are the boundaries of this topic?"
    
    extraction_method:
      - Identify nouns as potential concepts
      - Identify verbs as potential skills
      - Identify adjectives as concept properties
      - Map relationships between concepts
  
  prerequisites:
    discovery:
      - "What must be known before this?"
      - "What foundational knowledge is assumed?"
      - "What skills are required?"
    
    validation:
      - Check if learner has prerequisites
      - Identify gaps in foundation
      - Plan prerequisite coverage if needed
  
  success_criteria:
    definition:
      - "What demonstrates understanding?"
      - "What can learner do after learning?"
      - "How do we measure comprehension?"
    
    measurability:
      - Observable behaviors
      - Testable outcomes
      - Application examples
  
  knowledge_mapping:
    existing_knowledge:
      - Search Zettelkasten for related notes
      - Find previous learning sessions
      - Identify connection points
    
    knowledge_gaps:
      - What's new to learn
      - What needs reinforcement
      - What misconceptions to address
```

### Step 2: Generate Solution Path
Create a structured learning journey with clear progression.

```python
class LearningPathGenerator:
    def __init__(self, objective):
        self.objective = objective
        self.atomic_steps = []
        self.dependencies = {}
        self.checkpoints = []
        
    def break_into_atoms(self):
        """Break objective into smallest teachable units"""
        
        # Decomposition strategy based on content type
        if self.is_procedural():
            return self.procedural_decomposition()
        elif self.is_conceptual():
            return self.conceptual_decomposition()
        elif self.is_problem_solving():
            return self.problem_solving_decomposition()
        else:
            return self.hybrid_decomposition()
    
    def procedural_decomposition(self):
        """For step-by-step processes"""
        steps = []
        
        # 1. Identify the overall process
        process = self.extract_process()
        
        # 2. Break into sequential steps
        for i, step in enumerate(process):
            atomic_step = {
                "id": f"step_{i}",
                "content": step,
                "type": "procedure",
                "duration": self.estimate_duration(step),
                "prerequisites": [f"step_{i-1}"] if i > 0 else [],
                "verification": self.create_verification(step)
            }
            steps.append(atomic_step)
        
        return steps
    
    def conceptual_decomposition(self):
        """For understanding concepts"""
        atoms = []
        
        # 1. Core concept first
        core = {
            "id": "core_concept",
            "content": self.extract_core_concept(),
            "type": "concept",
            "duration": 5,
            "prerequisites": [],
            "verification": "Can explain in own words"
        }
        atoms.append(core)
        
        # 2. Properties and characteristics
        for prop in self.extract_properties():
            atoms.append({
                "id": f"property_{prop}",
                "content": prop,
                "type": "property",
                "duration": 3,
                "prerequisites": ["core_concept"],
                "verification": f"Can identify {prop}"
            })
        
        # 3. Examples and non-examples
        for example in self.generate_examples():
            atoms.append({
                "id": f"example_{example['id']}",
                "content": example,
                "type": "example",
                "duration": 2,
                "prerequisites": ["core_concept"],
                "verification": "Can recognize pattern"
            })
        
        return atoms
    
    def problem_solving_decomposition(self):
        """For problem-solving skills"""
        steps = []
        
        # 1. Problem understanding
        steps.append({
            "id": "understand_problem",
            "content": "Learn to parse and understand problem statement",
            "type": "skill",
            "duration": 5,
            "prerequisites": [],
            "verification": "Can restate problem clearly"
        })
        
        # 2. Strategy selection
        steps.append({
            "id": "select_strategy",
            "content": "Learn applicable solution strategies",
            "type": "skill",
            "duration": 10,
            "prerequisites": ["understand_problem"],
            "verification": "Can identify appropriate approach"
        })
        
        # 3. Solution execution
        steps.append({
            "id": "execute_solution",
            "content": "Apply strategy to solve",
            "type": "skill",
            "duration": 15,
            "prerequisites": ["select_strategy"],
            "verification": "Can solve similar problems"
        })
        
        # 4. Verification
        steps.append({
            "id": "verify_solution",
            "content": "Check and validate answer",
            "type": "skill",
            "duration": 5,
            "prerequisites": ["execute_solution"],
            "verification": "Can validate correctness"
        })
        
        return steps
    
    def sequence_by_dependency(self, atoms):
        """Order atoms by dependencies"""
        sequenced = []
        completed = set()
        
        while len(sequenced) < len(atoms):
            for atom in atoms:
                if atom["id"] in completed:
                    continue
                    
                # Check if all prerequisites are completed
                if all(prereq in completed for prereq in atom["prerequisites"]):
                    sequenced.append(atom)
                    completed.add(atom["id"])
        
        return sequenced
    
    def estimate_time(self):
        """Estimate total learning time"""
        total = sum(step["duration"] for step in self.atomic_steps)
        
        # Add buffer for transitions and reviews
        total_with_buffer = total * 1.3
        
        return {
            "minimum": total,
            "recommended": total_with_buffer,
            "with_practice": total_with_buffer * 1.5
        }
    
    def define_checkpoints(self):
        """Create verification checkpoints"""
        checkpoints = []
        
        # After each major concept group
        concept_groups = self.group_by_concept()
        
        for group in concept_groups:
            checkpoint = {
                "after_steps": [s["id"] for s in group],
                "verification": {
                    "knowledge_check": self.create_knowledge_check(group),
                    "application_task": self.create_application(group),
                    "success_criteria": self.define_success(group)
                },
                "remediation": self.plan_remediation(group)
            }
            checkpoints.append(checkpoint)
        
        return checkpoints
```

### Step 3: Progressive Verification
Check understanding at each step and adjust path.

```yaml
progressive_verification:
  check_types:
    immediate:
      description: "Quick comprehension check"
      methods:
        - "Can you explain this back to me?"
        - "What would happen if...?"
        - "Show me an example"
      timing: "After each atomic step"
    
    checkpoint:
      description: "Deeper understanding check"
      methods:
        - "Solve this problem"
        - "Find the error in this example"
        - "Connect to previous concept"
      timing: "After concept groups"
    
    synthesis:
      description: "Integration check"
      methods:
        - "Combine concepts to solve"
        - "Explain the big picture"
        - "Teach it back"
      timing: "After major sections"
  
  adjustment_triggers:
    struggling:
      indicators:
        - "Multiple incorrect attempts"
        - "Confusion expressions"
        - "Asks for clarification repeatedly"
      actions:
        - "Slow down pace"
        - "Add more examples"
        - "Break into smaller steps"
        - "Review prerequisites"
    
    mastering_quickly:
      indicators:
        - "Immediate correct responses"
        - "Asks advanced questions"
        - "Makes connections independently"
      actions:
        - "Speed up pace"
        - "Skip redundant examples"
        - "Add challenge problems"
        - "Introduce extensions"
    
    misconception_detected:
      indicators:
        - "Consistent error pattern"
        - "Wrong mental model"
        - "Overgeneralization"
      actions:
        - "Address directly"
        - "Provide counter-examples"
        - "Rebuild from foundation"
        - "Document in Zettelkasten"
  
  celebration_points:
    breakthrough:
      - "Aha moment detected"
      - "Difficult concept mastered"
      - "Connection made independently"
    
    milestone:
      - "Chapter completed"
      - "Skill demonstrated"
      - "Problem solved independently"
```

## Integration with Learning Session

```python
def execute_decomposed_learning(objective):
    """Execute a learning session with decomposition"""
    
    # 1. Decompose the objective
    decomposer = ProblemDecomposer(objective)
    learning_path = decomposer.generate_path()
    
    # 2. Check prerequisites
    missing_prereqs = decomposer.check_prerequisites()
    if missing_prereqs:
        # Add prerequisite learning to path
        learning_path = decomposer.prepend_prerequisites(missing_prereqs) + learning_path
    
    # 3. Execute learning path
    for step in learning_path:
        # Teach the step
        teach_atomic_concept(step)
        
        # Verify understanding
        understanding = verify_comprehension(step)
        
        # Adjust if needed
        if understanding < 0.7:
            # Add reinforcement
            add_reinforcement_step(step)
        elif understanding > 0.95:
            # Can skip similar steps
            learning_path = optimize_remaining_path(learning_path, step)
        
        # Document in Zettelkasten
        create_learning_note(step, understanding)
    
    # 4. Final synthesis
    synthesis = synthesize_learning(learning_path)
    
    # 5. Schedule review
    schedule_spaced_review(learning_path, synthesis)
    
    return {
        "path_taken": learning_path,
        "understanding_level": calculate_overall_understanding(),
        "time_taken": sum(step.duration for step in learning_path),
        "notes_created": get_session_notes(),
        "review_schedule": get_review_plan()
    }
```

## Example Decomposition

```yaml
example_objective: "Learn how recursion works in programming"

decomposed_path:
  1_prerequisites:
    - id: "prereq_functions"
      content: "Understanding functions and return values"
      duration: 5
      verification: "Can write and call a simple function"
  
  2_core_concept:
    - id: "recursion_definition"
      content: "A function calling itself"
      duration: 3
      verification: "Can identify recursive call in code"
  
  3_components:
    - id: "base_case"
      content: "Stopping condition"
      duration: 5
      verification: "Can identify and write base cases"
    
    - id: "recursive_case"
      content: "Self-call with modified input"
      duration: 5
      verification: "Can write recursive call correctly"
  
  4_simple_examples:
    - id: "factorial"
      content: "Computing factorial recursively"
      duration: 10
      verification: "Can trace factorial(5)"
    
    - id: "countdown"
      content: "Simple countdown function"
      duration: 5
      verification: "Can write countdown function"
  
  5_visualization:
    - id: "call_stack"
      content: "Understanding the call stack"
      duration: 10
      verification: "Can draw call stack for factorial(3)"
  
  6_complex_examples:
    - id: "fibonacci"
      content: "Fibonacci sequence"
      duration: 15
      verification: "Can implement and optimize"
    
    - id: "tree_traversal"
      content: "Recursive tree navigation"
      duration: 20
      verification: "Can traverse binary tree"
  
  7_problem_solving:
    - id: "when_to_use"
      content: "Recognizing recursive problems"
      duration: 10
      verification: "Can identify good use cases"
    
    - id: "conversion"
      content: "Converting iteration to recursion"
      duration: 15
      verification: "Can convert loop to recursion"
  
  checkpoints:
    - after: ["base_case", "recursive_case"]
      check: "Write a recursive function to sum array"
    
    - after: ["factorial", "countdown"]
      check: "Debug broken recursive function"
    
    - after: ["tree_traversal"]
      check: "Solve new recursive problem"

estimated_time:
  minimum: 90 minutes
  recommended: 120 minutes
  with_practice: 150 minutes
```

## Benefits

1. **Systematic Coverage**: No important concepts missed
2. **Adaptive Pacing**: Adjusts to learner's speed
3. **Clear Progress**: Visible checkpoints and achievements
4. **Efficient Learning**: Skip what's known, focus on gaps
5. **Better Retention**: Structured approach aids memory
6. **Struggle Detection**: Early intervention when stuck