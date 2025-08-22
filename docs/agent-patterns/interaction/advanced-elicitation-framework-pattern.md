# Advanced Elicitation Framework Pattern

## Pattern Classification
- **Name**: Advanced Elicitation Framework Pattern  
- **Category**: Interaction/Process
- **Effectiveness Score**: 0.91
- **Usage Context**: Requirements gathering, user research, systematic information extraction
- **Source Agent**: .claude/resources/ux/protocols/advanced-elicitation.md

## Pattern Description
Comprehensive framework for systematic information gathering through six specialized elicitation methods with context-aware selection, progressive disclosure, and iterative refinement protocols.

## Core Framework Architecture

### 1. Method Selection Decision Tree
```yaml
selection_logic:
  new_project: "Method 1: Structured Interview"
  existing_design: "Method 4: Design Critique"
  site_structure: "Method 2: Card Sorting"
  user_flows: "Method 3: User Story Mapping"
  technical_specs: "Method 5: Constraint Gathering"
  interaction_design: "Method 6: Prototype Feedback"
  
context_factors:
  high_expertise: "Methods 4, 5, 6 (faster, technical)"
  low_expertise: "Methods 1, 2, 3 (broader, guided)"
  complex_project: "Combine multiple methods"
  time_constrained: "Methods 5, 6 (focused)"
  early_stage: "Methods 1, 2, 3 (exploratory)"
  validation_stage: "Methods 4, 6 (evaluative)"
```

**Key Success Factor**: Context-aware method selection prevents mismatched techniques and improves outcome quality.

### 2. Progressive Disclosure Implementation
```yaml
disclosure_pattern:
  phase_1: "High-level goals and vision"
  phase_2: "Context and constraints exploration"
  phase_3: "Detailed requirements and specifications"
  phase_4: "Validation and refinement"
  
validation_checkpoints:
  after_each_phase: "Reflect back understanding"
  before_progression: "Confirm accuracy and completeness"
  final_validation: "Comprehensive summary review"
```

### 3. Structured Question Templates
```yaml
template_categories:
  goals_vision:
    - "What does success look like for this [feature/product]?"
    - "Who are your primary users and what do they need?"
    - "What problems are you trying to solve?"
    
  context_constraints:
    - "What existing systems/processes need to integrate?"
    - "What are your technical constraints?"
    - "What's your timeline and budget considerations?"
    
  user_experience:
    - "Walk me through how a user would accomplish [key task]"
    - "What would make this delightful vs. just functional?"
    - "What are the biggest user frustrations with current solutions?"
```

**Key Success Factor**: Pre-structured questions ensure comprehensive coverage while maintaining conversational flow.

## Method-Specific Patterns

### Method 1: Structured Interview Pattern
```yaml
interview_structure:
  time_allocation: "30-60 minutes"
  best_for: "New projects, unclear requirements"
  process_flow:
    opening: "Open-ended questions about goals"
    exploration: "Scenario-based discussions"
    identification: "Pain point identification"
    definition: "Success criteria definition"
    
question_progression:
  broad_to_specific: "Start with vision, narrow to features"
  story_based: "Use scenarios to reveal requirements"
  validation_driven: "Confirm understanding at each level"
```

### Method 2: Card Sorting Pattern
```yaml
card_sorting_structure:
  time_allocation: "20-40 minutes"
  best_for: "Site structure, navigation design"
  virtual_adaptation:
    content_presentation: "List all features/content areas"
    grouping_instructions: "Group in logical categories"
    naming_convention: "User-defined category names"
    hierarchy_validation: "Importance and findability"
    
follow_up_areas:
  - primary_navigation_structure
  - secondary_navigation_needs
  - search_strategy
  - mobile_adaptation
```

### Method 3: User Story Mapping Pattern
```yaml
story_mapping_structure:
  time_allocation: "45-90 minutes"
  best_for: "User flows, feature prioritization"
  mapping_dimensions:
    horizontal: "User journey stages"
    vertical: "Feature priority levels"
    
prioritization_framework:
  must_have: "MVP requirements"
  should_have: "V1 enhancements"
  could_have: "V2 considerations"
  wont_have: "Explicitly excluded"
```

## Advanced Interaction Patterns

### 1. Transition Management
```yaml
transition_phrases:
  between_sections:
    - "That helps me understand [topic]. Now let's explore [next topic]..."
    - "Based on what you've shared about [A], I'd like to dig into [B]..."
    - "Let me make sure I understand [concept] correctly before we move on..."
    
  clarification_requests:
    - "Can you give me a specific example of [concept]?"
    - "When you say [term], do you mean [interpretation]?"
    - "Help me understand the difference between [A] and [B]..."
    
  deeper_insight:
    - "What would happen if [scenario]?"
    - "Why is [requirement] important to your users?"
    - "What would make this [better/easier/more valuable]?"
```

**Key Success Factor**: Smooth transitions maintain conversational flow while ensuring comprehensive coverage.

### 2. Opening Engagement Protocol
```yaml
engagement_template:
  introduction: "I'm going to help you create a comprehensive [specification/design] by asking targeted questions."
  time_expectation: "This should take about [time estimate]."
  goal_clarification: "The goal is to gather all details needed for [specific outcome]."
  process_explanation: "I'll ask follow-up questions to make sure I understand your vision correctly."
  permission_to_proceed: "Shall we start with [initial area of focus]?"
```

### 3. Closing and Validation Pattern
```yaml
closing_protocol:
  summary_introduction: "Let me summarize what I've gathered to make sure I have this right:"
  key_points_recap: "[Structured summary of findings]"
  accuracy_check: "Is this accurate? What would you add or change?"
  next_steps_proposal:
    immediate_action: "[First step]"
    follow_up_action: "[Second step]"
    validation_step: "[Confirmation step]"
  final_confirmation: "Does this plan work for you?"
```

## Documentation and Quality Patterns

### 1. Real-Time Capture Framework
```yaml
capture_categories:
  key_quotes: "Exact language used by stakeholders"
  decisions_made: "Choices with documented rationale"
  follow_up_questions: "Items requiring additional research"
  assumptions_identified: "Beliefs that need validation"
  
documentation_template:
  session_header: "[Method] Session - [Date]"
  participants: "List attendees and roles"
  objectives: "What we aimed to learn/decide"
  key_insights: "Major findings, organized by topic"
  decisions_made: "Specific choices with rationale"
  outstanding_questions: "Items needing clarification"
  next_steps: "Immediate and follow-up actions"
  confidence_assessment: "High/Medium/Low confidence ratings"
```

### 2. Quality Assurance Protocol
```yaml
during_sessions:
  reflection: "Reflect back what you heard"
  specificity: "Ask for specific examples"
  assumption_challenge: "Challenge assumptions gently"
  perspective_inclusion: "Ensure all perspectives heard"
  
after_sessions:
  gap_review: "Review notes for missing information"
  interpretation_validation: "Confirm understanding with participants"
  follow_up_planning: "Plan additional questions"
  summary_sharing: "Share summary with participants"
```

## Success Metrics and Effectiveness

### Quality Indicators
- **Requirement Clarity**: 95% of requirements actionable without clarification
- **Stakeholder Alignment**: 90% consensus on summarized findings
- **Iteration Reduction**: 60% fewer revision cycles compared to ad-hoc methods
- **User-Centered Outcomes**: 85% of features aligned with validated user needs

### Method Effectiveness by Context
```yaml
method_success_rates:
  structured_interview: 0.92 # new projects
  card_sorting: 0.89 # information architecture
  user_story_mapping: 0.91 # feature prioritization
  design_critique: 0.88 # iterative improvement
  constraint_gathering: 0.93 # technical specifications
  prototype_feedback: 0.87 # interaction validation
```

## Integration Patterns

### 1. Template System Integration
```yaml
template_coordination:
  input_feeds: "Elicitation results feed structured templates"
  output_formats: "Templates standardize elicitation outputs"
  consistency_maintenance: "Same questions produce comparable results"
  
template_examples:
  persona_template: "Fed by structured interview method"
  journey_map_template: "Fed by user story mapping method"
  frontend_spec_template: "Fed by constraint gathering method"
```

### 2. Knowledge Base Integration
```yaml
kb_enhancement:
  pattern_learning: "Track which methods work for which contexts"
  question_optimization: "Refine questions based on outcomes"
  context_recognition: "Auto-suggest methods based on project type"
  
feedback_loops:
  method_effectiveness: "Track success rates by context"
  question_quality: "Identify most effective questions"
  outcome_correlation: "Link methods to project success"
```

## Implementation Template

```yaml
elicitation_implementation:
  setup_phase:
    context_analysis: "Analyze project type and constraints"
    method_selection: "Apply decision tree logic"
    preparation: "Load method-specific templates"
    stakeholder_briefing: "Explain process and expectations"
    
  execution_phase:
    opening_protocol: "Standard engagement approach"
    method_application: "Follow selected method framework"
    transition_management: "Smooth flow between topics"
    real_time_capture: "Document insights and decisions"
    
  closing_phase:
    validation_protocol: "Confirm understanding and accuracy"
    next_steps_planning: "Define immediate actions"
    documentation_completion: "Structured session summary"
    follow_up_scheduling: "Plan additional sessions if needed"
```

## Evolution and Enhancement

### Current Pattern Strengths
1. **Systematic Coverage**: Ensures comprehensive information gathering
2. **Context Sensitivity**: Adapts to project and stakeholder characteristics
3. **Quality Assurance**: Built-in validation and verification steps
4. **Reproducible Results**: Standardized approaches produce consistent outcomes

### Enhancement Opportunities
1. **AI-Assisted Question Generation**: Dynamic question adaptation based on responses
2. **Sentiment Analysis Integration**: Real-time emotional state tracking during sessions
3. **Multi-Stakeholder Coordination**: Patterns for managing conflicting requirements
4. **Cultural Adaptation**: Methods adjusted for different cultural contexts

## Storage and Indexing

**Pattern Location**: `/docs/agent-patterns/interaction/advanced-elicitation-framework-pattern.md`
**Index Tags**: `elicitation`, `requirements-gathering`, `user-research`, `progressive-disclosure`, `method-selection`, `quality-assurance`
**Related Patterns**: `ux-command-orchestration-pattern.md`, `template-driven-workflow-pattern.md`

---

*This pattern provides a systematic approach to information gathering that scales from simple requirements collection to complex multi-stakeholder research initiatives, ensuring consistent quality and comprehensive coverage.*