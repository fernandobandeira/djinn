# Workflow Templates for Architecture Orchestration

## Single-Agent Workflows

### ADR Creation Workflow
```yaml
workflow_name: "create_adr"
description: "Create Architecture Decision Record"
sub_agent: "adr-manager"

steps:
  1. validate_input:
      - decision_context_provided: true
      - problem_statement_clear: true
      - stakeholders_identified: true
  
  2. delegate_to_adr_manager:
      prompt_template: |
        Create ADR for: {decision_topic}
        Context: {problem_context}
        Stakeholders: {stakeholders}
        Decision deadline: {timeline}
  
  3. quality_check:
      - template_sections_complete: true
      - alternatives_documented: true
      - consequences_clear: true
  
  4. finalize:
      - save_to_docs: "/docs/architecture/adrs/"
      - update_knowledge_base: true
      - notify_stakeholders: true
```

### System Design Workflow
```yaml
workflow_name: "design_system"
description: "Complete system architecture design"
sub_agent: "system-designer"

steps:
  1. requirements_gathering:
      - functional_requirements: documented
      - non_functional_requirements: documented
      - constraints: identified
  
  2. delegate_to_system_designer:
      prompt_template: |
        Design system architecture for: {system_name}
        Requirements: {requirements}
        Constraints: {constraints}
        Timeline: {deadline}
  
  3. quality_check:
      - multiple_options_provided: true
      - tradeoffs_analyzed: true
      - migration_path_defined: true
  
  4. finalize:
      - create_architecture_diagrams: true
      - document_decisions_as_adrs: true
      - update_pattern_library: true
```

## Multi-Agent Workflows

### Comprehensive Architecture Review
```yaml
workflow_name: "comprehensive_architecture_review"
description: "Full architecture assessment with multiple perspectives"
orchestration_type: "sequential_with_synthesis"

phase_1_analysis:
  sub_agent: "architecture-reviewer"
  task: "Initial architecture analysis and issue identification"
  prompt_template: |
    Analyze current architecture for: {system_name}
    Focus areas: {focus_areas}
    Business context: {business_context}
  
  expected_outputs:
    - current_state_assessment
    - critical_issues_list
    - improvement_priorities
    - recommended_patterns

phase_2_visualization:
  sub_agent: "diagram-generator"
  task: "Create current state and proposed architecture diagrams"
  dependencies: [phase_1_analysis]
  prompt_template: |
    Create architecture diagrams based on analysis:
    Current state issues: {phase_1.critical_issues}
    Proposed improvements: {phase_1.recommendations}
  
  expected_outputs:
    - current_state_diagram
    - proposed_architecture_diagram
    - migration_flow_diagram

phase_3_decision_documentation:
  sub_agent: "adr-manager"
  task: "Document key architectural decisions"
  dependencies: [phase_1_analysis, phase_2_visualization]
  prompt_template: |
    Create ADRs for architectural decisions:
    Key decisions: {phase_1.decisions}
    Supporting diagrams: {phase_2.diagrams}
  
  expected_outputs:
    - architectural_decision_records
    - decision_rationale_documentation

phase_4_pattern_integration:
  sub_agent: "pattern-librarian"
  task: "Update patterns and create operational documentation"
  dependencies: [phase_1_analysis, phase_3_decision_documentation]
  prompt_template: |
    Update architectural patterns based on review:
    New patterns identified: {phase_1.patterns}
    Decisions made: {phase_3.decisions}
  
  expected_outputs:
    - updated_pattern_library
    - operational_runbooks
    - rfc_documents

synthesis:
  description: "Combine all phase outputs into comprehensive deliverable"
  inputs: [phase_1, phase_2, phase_3, phase_4]
  output_format: "comprehensive_architecture_review_report"
```

### New System Architecture Design
```yaml
workflow_name: "new_system_design"
description: "End-to-end new system architecture design"
orchestration_type: "iterative_refinement"

iteration_1_initial_design:
  sub_agent: "system-designer"
  task: "Create initial architecture options"
  prompt_template: |
    Design new system architecture for: {system_requirements}
    Business constraints: {constraints}
    Technology preferences: {tech_preferences}

iteration_2_pattern_integration:
  sub_agent: "pattern-librarian"
  task: "Identify applicable patterns and standards"
  dependencies: [iteration_1_initial_design]
  prompt_template: |
    Identify patterns applicable to design:
    Proposed architecture: {iteration_1.architecture}
    System characteristics: {iteration_1.characteristics}

iteration_3_design_refinement:
  sub_agent: "system-designer"
  task: "Refine design with patterns and standards"
  dependencies: [iteration_1_initial_design, iteration_2_pattern_integration]
  prompt_template: |
    Refine architecture design incorporating:
    Initial design: {iteration_1.options}
    Applicable patterns: {iteration_2.patterns}
    Standards requirements: {iteration_2.standards}

iteration_4_visualization:
  sub_agent: "diagram-generator"
  task: "Create comprehensive architecture diagrams"
  dependencies: [iteration_3_design_refinement]
  prompt_template: |
    Create diagrams for refined architecture:
    Final architecture: {iteration_3.final_design}
    Key components: {iteration_3.components}

iteration_5_documentation:
  sub_agent: "adr-manager"
  task: "Document key design decisions"
  dependencies: [iteration_3_design_refinement, iteration_4_visualization]
  prompt_template: |
    Document architectural decisions:
    Design rationale: {iteration_3.rationale}
    Key tradeoffs: {iteration_3.tradeoffs}

final_synthesis:
  description: "Create complete system design package"
  inputs: [iteration_3, iteration_4, iteration_5]
  output_format: "system_design_specification"
```

## Workflow Execution Patterns

### Sequential Execution
```yaml
execution_pattern: "sequential"
description: "Tasks execute one after another, each building on previous results"

characteristics:
  - predictable_order: true
  - dependency_management: explicit
  - error_handling: stop_on_failure
  - parallel_processing: false

use_cases:
  - architecture_review: "Each phase builds on previous analysis"
  - system_design: "Design → Documentation → Validation"
  - adr_creation: "Analysis → Decision → Documentation"
```

### Parallel Execution
```yaml
execution_pattern: "parallel"
description: "Independent tasks execute simultaneously"

characteristics:
  - concurrent_processing: true
  - faster_completion: true
  - resource_intensive: true
  - synchronization_required: true

use_cases:
  - diagram_generation: "Multiple diagram types for same system"
  - standards_validation: "Multiple quality checks simultaneously"
  - pattern_analysis: "Different pattern categories in parallel"
```

### Iterative Refinement
```yaml
execution_pattern: "iterative"
description: "Progressive improvement through multiple cycles"

characteristics:
  - quality_improvement: incremental
  - feedback_incorporation: continuous
  - flexibility: high
  - time_investment: higher

use_cases:
  - complex_system_design: "Refine architecture through iterations"
  - pattern_development: "Evolve patterns based on usage"
  - quality_improvement: "Continuous architecture refinement"
```

## Error Handling and Recovery

### Workflow Failure Patterns
```yaml
failure_handling:
  sub_agent_timeout:
    action: "retry_with_simplified_prompt"
    max_retries: 2
    escalation: "manual_intervention"
  
  quality_gate_failure:
    action: "request_additional_context"
    retry_with_context: true
    escalation: "workflow_modification"
  
  dependency_failure:
    action: "skip_dependent_tasks"
    alternative_path: "simplified_workflow"
    escalation: "user_notification"
  
  resource_unavailable:
    action: "use_fallback_resources"
    fallback_strategy: "generic_templates"
    escalation: "defer_workflow"
```

### Recovery Strategies
- **Checkpoint Creation**: Save intermediate results for recovery
- **Graceful Degradation**: Deliver partial results when full workflow fails
- **Alternative Workflows**: Switch to simpler workflows when complex ones fail
- **Manual Intervention**: Escalate to human expert when automation fails

## Workflow Monitoring and Metrics

### Performance Metrics
- **Execution Time**: Track workflow completion times
- **Success Rate**: Monitor workflow completion rates
- **Quality Scores**: Measure output quality consistently
- **User Satisfaction**: Collect feedback on workflow results

### Optimization Opportunities
- **Bottleneck Identification**: Find slow steps in workflows
- **Resource Utilization**: Optimize sub-agent usage
- **Quality Improvement**: Enhance prompt effectiveness
- **Workflow Simplification**: Remove unnecessary complexity