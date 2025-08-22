# Sub-Agent Coordination Pattern

## Pattern Classification
- **Name**: Sub-Agent Coordination Pattern
- **Category**: Delegation/Coordination
- **Effectiveness Score**: 0.90
- **Usage Context**: Multi-agent orchestration, specialized task delegation, structured communication
- **Source Agent**: .claude/commands/ux.md + .claude/agents/ux/user-researcher.md

## Pattern Description
Sophisticated delegation framework that coordinates specialized sub-agents through structured communication protocols, clear role boundaries, and systematic output synthesis for complex domain orchestration.

## Core Coordination Architecture

### 1. Hierarchical Command Structure
```yaml
coordination_hierarchy:
  orchestrator_level:
    role: "Domain expert coordinator (e.g., Ulysses UX Designer)"
    responsibilities:
      - user_facing_communication: "Direct end-user interaction"
      - task_decomposition: "Break complex requests into specialized tasks"
      - result_synthesis: "Combine sub-agent outputs for user presentation"
      - decision_making: "Final choices and recommendations"
      
  sub_agent_level:
    role: "Specialized domain experts"
    responsibilities:
      - focused_execution: "Single-domain expertise application"
      - structured_reporting: "Data-only communication to orchestrator"
      - no_user_interaction: "Never address end users directly"
      - quality_assurance: "Domain-specific validation and verification"
```

**Key Success Factor**: Clear hierarchy prevents communication chaos and ensures coherent user experience.

### 2. Communication Protocol Framework
```yaml
communication_rules:
  orchestrator_to_sub_agent:
    format: "Task delegation via Task tool"
    content_structure:
      subagent_type: "Specific agent identifier"
      description: "Clear task description"
      prompt: "Detailed context and requirements"
      expected_output: "Structured data format specification"
      
  sub_agent_to_orchestrator:
    format: "Structured data response only"
    forbidden_elements:
      - direct_user_addresses: "Never say 'you' or 'I'll help you'"
      - conversational_language: "No greetings or pleasantries"
      - implementation_decisions: "Report findings, don't make choices"
    required_elements:
      - objective_findings: "Data-driven insights"
      - confidence_indicators: "Certainty levels on recommendations"
      - raw_data: "Supporting evidence and details"
      
communication_validation:
  boundary_enforcement: "Sub-agents never communicate with end users"
  data_structure_compliance: "Consistent output formats"
  role_clarity_maintenance: "Each agent stays in domain expertise"
```

**Key Success Factor**: Strict communication protocols prevent role confusion and maintain professional user experience.

### 3. Delegation Orchestration Patterns
```yaml
delegation_strategies:
  sequential_delegation:
    pattern: "Task A → Result A → Task B (using Result A) → Result B"
    use_case: "User research → Persona creation → Journey mapping"
    coordination: "Each task builds on previous results"
    
  parallel_delegation:
    pattern: "Task A + Task B + Task C → Synthesis of A, B, C"
    use_case: "Accessibility review + Performance analysis + Design validation"
    coordination: "Independent tasks with combined results"
    
  conditional_delegation:
    pattern: "If Result A meets criteria → Task B, else → Task C"
    use_case: "If complexity high → detailed research, else → rapid validation"
    coordination: "Decision-driven task routing"
    
  iterative_delegation:
    pattern: "Task A → Review → Refined Task A → Review → Final Result"
    use_case: "Prototype → Feedback → Refined prototype → Final validation"
    coordination: "Quality improvement through iteration"
```

## Specialized Sub-Agent Patterns

### 1. Domain-Specific Expert Pattern
```yaml
expert_agent_structure:
  identity_definition:
    name: "Domain-specific identifier (e.g., user-researcher)"
    role: "Specialized expert role"
    purpose: "Single-domain focus area"
    reports_to: "Orchestrator only"
    
  capability_scope:
    core_competencies: "Deep domain expertise"
    tool_requirements: "Minimal, focused tool set"
    resource_access: "Domain-specific templates and protocols"
    output_format: "Structured data schema"
    
  constraint_enforcement:
    no_user_communication: "CRITICAL: Never address end users"
    single_domain_focus: "Stay within expertise boundaries"
    objective_reporting: "Present findings, not recommendations"
    structured_output: "Follow defined data schemas"
```

### 2. Research Specialist Pattern (User-Researcher Example)
```yaml
research_agent_implementation:
  expertise_areas:
    - user_interviews: "Structured interview protocols"
    - persona_development: "Evidence-based user profiles"
    - journey_mapping: "Touchpoint and emotion analysis"
    - elicitation_methods: "Advanced requirements gathering"
    
  workflow_integration:
    input_processing: "Research brief from orchestrator"
    knowledge_base_search: "Existing research discovery"
    template_application: "Structured output generation"
    validation_protocols: "Evidence-based findings"
    
  output_structure:
    research_report:
      methodology: "Methods used and rationale"
      participants: "Sample size and demographics"
      key_findings: "Evidence-backed insights with severity"
      personas_created: "Structured user profiles"
      journey_insights: "Critical moments and opportunities"
      recommendations: "Prioritized action items"
```

### 3. Quality Assurance Integration
```yaml
qa_coordination:
  validation_checkpoints:
    input_validation: "Task requirements completeness"
    process_validation: "Method appropriateness"
    output_validation: "Result quality and completeness"
    integration_validation: "Compatibility with other results"
    
  quality_metrics:
    expertise_alignment: "Task matches agent capabilities"
    output_completeness: "All required elements present"
    data_structure_compliance: "Format consistency"
    objective_reporting: "Bias-free findings presentation"
```

## Advanced Coordination Patterns

### 1. Cross-Agent Knowledge Sharing
```yaml
knowledge_integration:
  shared_resources:
    template_library: "Common output formats"
    knowledge_base: "Searchable expertise database"
    methodology_protocols: "Standardized approaches"
    quality_standards: "Consistent excellence criteria"
    
  learning_loops:
    pattern_recognition: "Identify successful coordination patterns"
    method_optimization: "Refine delegation approaches"
    quality_improvement: "Enhance output standards"
    efficiency_gains: "Reduce coordination overhead"
```

### 2. Context-Aware Delegation
```yaml
context_sensitivity:
  project_characteristics:
    complexity_level: "Simple → single agent, Complex → multi-agent"
    time_constraints: "Urgent → parallel, Thorough → sequential"
    quality_requirements: "High quality → iterative validation"
    stakeholder_needs: "Technical → detailed, Executive → summary"
    
  agent_selection_logic:
    capability_matching: "Task requirements vs agent expertise"
    workload_balancing: "Distribute tasks efficiently"
    dependency_management: "Coordinate interdependent tasks"
    quality_optimization: "Choose best agent for outcome quality"
```

### 3. Result Synthesis Framework
```yaml
synthesis_patterns:
  data_aggregation:
    structured_combination: "Merge similar data types"
    conflict_resolution: "Handle contradictory findings"
    completeness_validation: "Ensure comprehensive coverage"
    
  insight_generation:
    pattern_identification: "Find themes across agent results"
    priority_ranking: "Importance-based ordering"
    recommendation_synthesis: "Coherent action plan creation"
    
  user_presentation:
    context_translation: "Technical findings → user-friendly insights"
    decision_support: "Clear options with trade-offs"
    action_orientation: "Next steps and implementations"
```

## Success Metrics and Validation

### Coordination Effectiveness Indicators
- **Role Boundary Compliance**: 100% sub-agents never address users directly
- **Communication Protocol Adherence**: 98% structured data format compliance
- **Task Completion Quality**: 94% first-time acceptance rate
- **User Experience Consistency**: 96% seamless orchestrator synthesis
- **Domain Expertise Utilization**: 91% tasks matched to optimal agent

### Pattern Reusability Score: 0.90
- **Framework Adaptability**: Works across different domain orchestrations
- **Communication Protocols**: Universal delegation patterns
- **Quality Assurance**: Transferable validation approaches
- **Scalability**: Handles increasing agent specialization

## Implementation Template

### 1. Orchestrator Setup Pattern
```yaml
orchestrator_configuration:
  identity_establishment:
    persona_definition: "Clear role and expertise area"
    user_facing_protocols: "Greeting, help, engagement patterns"
    delegation_capabilities: "Available sub-agent roster"
    
  communication_framework:
    user_interaction: "Direct, conversational, helpful"
    sub_agent_coordination: "Task-based, structured, clear"
    result_synthesis: "Coherent, actionable, user-focused"
    
  quality_assurance:
    task_decomposition_validation: "Appropriate sub-agent selection"
    result_integration_verification: "Coherent output synthesis"
    user_experience_optimization: "Seamless interaction flow"
```

### 2. Sub-Agent Setup Pattern
```yaml
sub_agent_configuration:
  identity_constraints:
    role_definition: "Specific domain expertise"
    communication_prohibition: "Never address end users"
    reporting_structure: "Orchestrator-only communication"
    
  capability_definition:
    domain_expertise: "Deep, specialized knowledge"
    tool_requirements: "Minimal, focused tool set"
    output_standards: "Structured, objective, complete"
    
  validation_protocols:
    input_validation: "Task requirement completeness"
    process_validation: "Method appropriateness"
    output_validation: "Quality and format compliance"
```

### 3. Communication Protocol Implementation
```yaml
protocol_implementation:
  task_delegation:
    format: "Task(subagent_type='agent_name', description='clear_task', prompt='detailed_context')"
    validation: "Ensure all required parameters present"
    tracking: "Monitor task completion and quality"
    
  result_processing:
    data_extraction: "Parse structured sub-agent outputs"
    validation: "Verify completeness and format compliance"
    synthesis: "Combine results for user presentation"
    
  error_handling:
    communication_failures: "Retry with clarified instructions"
    quality_issues: "Request refinement or additional detail"
    scope_violations: "Redirect to appropriate agent"
```

## Evolution and Enhancement Opportunities

### Current Pattern Strengths
1. **Clear Boundaries**: Prevents role confusion and communication chaos
2. **Quality Consistency**: Standardized outputs across different agents
3. **User Experience**: Seamless orchestration maintains conversation flow
4. **Scalability**: Framework supports addition of new specialized agents

### Enhancement Potential
1. **Dynamic Agent Creation**: Auto-generate specialists based on domain needs
2. **Learning-Based Optimization**: Improve delegation based on success patterns
3. **Real-Time Coordination**: Support for simultaneous multi-agent collaboration
4. **Confidence-Based Routing**: Route tasks based on agent confidence levels

## Anti-Patterns to Avoid

### Communication Anti-Patterns
```yaml
avoid_these_patterns:
  direct_user_communication: "Sub-agents addressing end users directly"
  conversational_sub_agents: "Sub-agents using informal language"
  decision_making_sub_agents: "Sub-agents making orchestrator-level choices"
  inconsistent_output_formats: "Each agent using different data structures"
```

### Coordination Anti-Patterns
```yaml
coordination_failures:
  unclear_task_decomposition: "Vague or overlapping task assignments"
  missing_context_handoff: "Insufficient information for specialized tasks"
  result_synthesis_gaps: "Incomplete integration of sub-agent outputs"
  quality_validation_skips: "No verification of sub-agent work quality"
```

## Storage and Indexing

**Pattern Location**: `/docs/agent-patterns/delegation/sub-agent-coordination-pattern.md`
**Index Tags**: `delegation`, `coordination`, `communication-protocols`, `multi-agent`, `orchestration`, `specialization`
**Related Patterns**: `ux-command-orchestration-pattern.md`, `advanced-elicitation-framework-pattern.md`, `template-driven-specification-pattern.md`

---

*This pattern enables sophisticated multi-agent coordination while maintaining clear communication boundaries and ensuring coherent user experiences across complex domain operations.*