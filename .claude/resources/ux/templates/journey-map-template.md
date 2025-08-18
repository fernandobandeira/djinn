# User Journey Map Template

## Purpose
Comprehensive template for mapping user experiences across touchpoints, identifying opportunities for improvement and optimization.

## Usage Instructions
1. Start with a specific persona and scenario
2. Define journey scope and boundaries
3. Map current state first, then future state
4. Involve stakeholders in mapping sessions
5. Validate with real user data

---

# {{journey_name}} - {{persona_name}}

## Journey Overview

### Scenario Context
**Journey Type**: {{journey_type}}  
**Persona**: {{persona_name}}  
**Goal**: {{primary_goal}}  
**Timeframe**: {{journey_duration}}  
**Frequency**: {{frequency}}

### Journey Scope
**Start Point**: {{journey_start}}  
**End Point**: {{journey_end}}  
**In Scope**: {{scope_included}}  
**Out of Scope**: {{scope_excluded}}

## Journey Phases

{{#each journey_phases}}
### Phase {{@index}}: {{phase_name}}

#### Phase Overview
**Duration**: {{phase_duration}}  
**Primary Goal**: {{phase_goal}}  
**Context**: {{phase_context}}

#### User Actions
{{#each user_actions}}
{{@index}}. {{action}}
   - **Method**: {{method}}
   - **Duration**: {{duration}}
   - **Frequency**: {{frequency}}
{{/each}}

#### Touchpoints
{{#each touchpoints}}
- **{{touchpoint_name}}** ({{channel}})
  - Purpose: {{purpose}}
  - Quality: {{quality_rating}}/10
  - Issues: {{issues}}
{{/each}}

#### Thoughts & Feelings
**Thoughts**: {{thoughts}}  
**Emotions**: {{emotions}}  
**Sentiment**: {{sentiment_level}}

#### Pain Points
{{#each pain_points}}
- **{{pain_point}}**
  - Severity: {{severity}}/10
  - Impact: {{impact}}
  - Root Cause: {{root_cause}}
{{/each}}

#### Opportunities
{{#each opportunities}}
- **{{opportunity}}**
  - Potential Impact: {{impact_level}}
  - Effort Required: {{effort_level}}
  - Priority: {{priority}}
{{/each}}

#### Success Metrics
{{#each success_metrics}}
- {{metric}}: {{current_value}} → {{target_value}}
{{/each}}

{{/each}}

## Journey Map Visualization

### High-Level Journey Flow
```mermaid
graph LR
{{#each journey_phases}}
    {{@key}}[{{phase_name}}]
    {{#unless @last}}{{@key}} --> {{next_phase_key}}{{/unless}}
{{/each}}
```

### Emotional Journey
```mermaid
graph LR
    subgraph "Emotional Journey"
    {{#each journey_phases}}
        {{@key}}["{{phase_name}}<br/>{{sentiment_level}}"]
        {{#unless @last}}{{@key}} --> {{next_phase_key}}{{/unless}}
    {{/each}}
    end
```

### Touchpoint Map
| Phase | Touchpoints | Channel | Quality | Issues |
|-------|-------------|---------|---------|--------|
{{#each journey_phases}}
{{#each touchpoints}}
| {{../phase_name}} | {{touchpoint_name}} | {{channel}} | {{quality_rating}}/10 | {{issues}} |
{{/each}}
{{/each}}

## Cross-Phase Analysis

### Overall Journey Metrics
- **Total Duration**: {{total_duration}}
- **Touchpoint Count**: {{touchpoint_count}}
- **Average Satisfaction**: {{average_satisfaction}}/10
- **Completion Rate**: {{completion_rate}}%
- **Drop-off Points**: {{drop_off_points}}

### Emotion Curve
{{emotion_curve_description}}

### Critical Moments
{{#each critical_moments}}
#### {{moment_name}}
**Phase**: {{phase}}  
**Description**: {{description}}  
**Impact**: {{impact}}  
**Current Experience**: {{current_rating}}/10  
**Improvement Opportunity**: {{improvement_potential}}
{{/each}}

### Pain Point Summary
| Pain Point | Phase | Severity | Frequency | Impact | Solution Priority |
|------------|-------|----------|-----------|--------|-------------------|
{{#each all_pain_points}}
| {{pain_point}} | {{phase}} | {{severity}}/10 | {{frequency}} | {{impact}} | {{priority}} |
{{/each}}

### Opportunity Summary
| Opportunity | Phase | Impact | Effort | ROI | Priority |
|-------------|-------|--------|--------|-----|----------|
{{#each all_opportunities}}
| {{opportunity}} | {{phase}} | {{impact}} | {{effort}} | {{roi}} | {{priority}} |
{{/each}}

## Supporting Elements

### Channels & Touchpoints
{{#each channels}}
#### {{channel_name}}
**Type**: {{channel_type}}  
**Ownership**: {{channel_ownership}}  
**Current Performance**: {{performance_rating}}/10  
**Key Issues**: {{key_issues}}  
**Improvement Potential**: {{improvement_potential}}
{{/each}}

### Backstage Processes
{{#each backstage_processes}}
#### {{process_name}}
**Phase**: {{related_phase}}  
**Description**: {{process_description}}  
**Current State**: {{current_state}}  
**Issues**: {{process_issues}}  
**Improvement Needs**: {{improvement_needs}}
{{/each}}

### Technology & Systems
{{#each systems}}
#### {{system_name}}
**Purpose**: {{system_purpose}}  
**User Facing**: {{user_facing}}  
**Performance**: {{system_performance}}/10  
**Issues**: {{system_issues}}  
**Dependencies**: {{dependencies}}
{{/each}}

## Future State Vision

### Ideal Experience
{{ideal_experience_description}}

### Target Improvements
{{#each target_improvements}}
#### {{improvement_area}}
**Current State**: {{current_state}}  
**Future State**: {{future_state}}  
**Success Metrics**: {{success_metrics}}  
**Timeline**: {{timeline}}
{{/each}}

### Enhanced Journey Flow
```mermaid
graph LR
    subgraph "Future State Journey"
    {{#each future_journey_phases}}
        {{@key}}[{{phase_name}}<br/>{{target_sentiment}}]
        {{#unless @last}}{{@key}} --> {{next_phase_key}}{{/unless}}
    {{/each}}
    end
```

## Implementation Roadmap

### Immediate Actions (0-3 months)
{{#each immediate_actions}}
- **{{action}}**
  - Owner: {{owner}}
  - Impact: {{impact}}
  - Effort: {{effort}}
{{/each}}

### Short-term Initiatives (3-6 months)
{{#each short_term_initiatives}}
- **{{initiative}}**
  - Owner: {{owner}}
  - Dependencies: {{dependencies}}
  - Success Metrics: {{metrics}}
{{/each}}

### Long-term Vision (6+ months)
{{#each long_term_vision}}
- **{{vision_item}}**
  - Strategic Impact: {{strategic_impact}}
  - Resource Requirements: {{resources}}
  - Timeline: {{timeline}}
{{/each}}

## Measurement & Validation

### Key Performance Indicators
{{#each kpis}}
- **{{kpi_name}}**
  - Current: {{current_value}}
  - Target: {{target_value}}
  - Measurement Method: {{measurement_method}}
  - Frequency: {{measurement_frequency}}
{{/each}}

### User Feedback Collection
{{#each feedback_methods}}
- **{{method}}**: {{description}}
  - Frequency: {{frequency}}
  - Sample Size: {{sample_size}}
  - Current Score: {{current_score}}
{{/each}}

### Analytics & Data Sources
{{#each data_sources}}
- **{{source_name}}**: {{data_type}}
  - Current Access: {{access_level}}
  - Data Quality: {{quality_rating}}/10
  - Usage: {{usage_description}}
{{/each}}

## Stakeholder Perspectives

### Internal Stakeholders
{{#each internal_stakeholders}}
#### {{stakeholder_role}}
**Current Pain Points**: {{pain_points}}  
**Success Criteria**: {{success_criteria}}  
**Resource Constraints**: {{constraints}}  
**Priorities**: {{priorities}}
{{/each}}

### External Partners
{{#each external_partners}}
#### {{partner_name}}
**Role in Journey**: {{role}}  
**Current Relationship**: {{relationship_quality}}  
**Integration Points**: {{integration_points}}  
**Improvement Opportunities**: {{opportunities}}
{{/each}}

## Research & Validation

### Research Methods Used
{{#each research_methods}}
- **{{method}}**: {{description}}
  - Sample Size: {{sample_size}}
  - Date: {{date}}
  - Key Findings: {{findings}}
  - Confidence Level: {{confidence}}/10
{{/each}}

### Validation Status
{{#each validation_items}}
- **{{item}}**: {{status}}
  - Evidence: {{evidence}}
  - Confidence: {{confidence_level}}
  - Next Steps: {{next_steps}}
{{/each}}

### Assumptions & Hypotheses
{{#each assumptions}}
- **{{assumption}}**
  - Impact if Wrong: {{impact}}
  - Validation Method: {{validation_method}}
  - Timeline: {{validation_timeline}}
{{/each}}

## Journey Variants

### Alternative Scenarios
{{#each alternative_scenarios}}
#### {{scenario_name}}
**Trigger**: {{trigger}}  
**Key Differences**: {{differences}}  
**Frequency**: {{frequency}}  
**Special Considerations**: {{considerations}}
{{/each}}

### Edge Cases
{{#each edge_cases}}
#### {{edge_case}}
**Probability**: {{probability}}  
**Impact**: {{impact}}  
**Current Handling**: {{current_handling}}  
**Improvement Needed**: {{improvement_needed}}
{{/each}}

## Appendices

### Detailed User Quotes
{{#each user_quotes}}
*"{{quote}}"* - {{user_profile}} ({{phase}})
{{/each}}

### Supporting Research
{{#each research_documents}}
- **{{document_title}}**: {{link}}
  - Date: {{date}}
  - Key Insights: {{insights}}
{{/each}}

### Change Log
| Date | Version | Changes | Author |
|------|---------|---------|--------|
| {{date}} | {{version}} | Initial journey map | {{author}} |

---

**Map Created**: {{creation_date}}  
**Last Updated**: {{last_updated}}  
**Version**: {{version}}  
**Research Team**: {{research_team}}  
**Stakeholder Sign-off**: {{signoff_status}}

---

*This journey map should be reviewed and updated regularly as user behavior and business processes evolve. Validate insights with real user data and stakeholder feedback.*