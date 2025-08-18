# User Persona Template

## Purpose
Structured template for creating detailed user personas that inform design decisions and development priorities.

## Usage Instructions
1. Conduct user research first (interviews, surveys, analytics)
2. Use elicitation methods to gather stakeholder input
3. Fill template with validated data, not assumptions
4. Create 2-4 primary personas maximum
5. Reference personas throughout design process

---

# {{persona_name}} - {{role_title}}

## Overview
**Age**: {{age_range}}  
**Location**: {{location}}  
**Occupation**: {{occupation}}  
**Experience Level**: {{experience_level}}

*"{{persona_quote}}"*

## Demographics & Context

### Background
{{background_description}}

### Technical Proficiency
- **Overall Tech Comfort**: {{tech_comfort_level}}
- **Platform Usage**: {{primary_platforms}}
- **Device Preferences**: {{device_preferences}}
- **Internet Usage**: {{internet_usage_patterns}}

### Professional Context
- **Industry**: {{industry}}
- **Company Size**: {{company_size}}
- **Role Responsibilities**: {{role_responsibilities}}
- **Tools Used**: {{professional_tools}}

## Goals & Motivations

### Primary Goals
{{#each primary_goals}}
- **{{goal}}**: {{description}}
{{/each}}

### Success Metrics
{{#each success_metrics}}
- {{metric}}: {{target_outcome}}
{{/each}}

### Motivations
{{#each motivations}}
- **{{motivation_type}}**: {{description}}
{{/each}}

## Pain Points & Frustrations

### Current Challenges
{{#each pain_points}}
- **{{pain_point}}**: {{description}}
  - Impact: {{impact_level}}
  - Frequency: {{frequency}}
{{/each}}

### Barriers to Success
{{#each barriers}}
- {{barrier}}: {{barrier_description}}
{{/each}}

### Technology Frustrations
{{technology_frustrations}}

## Behaviors & Preferences

### Usage Patterns
- **Frequency**: {{usage_frequency}}
- **Duration**: {{typical_session_length}}
- **Time of Day**: {{preferred_usage_times}}
- **Context**: {{usage_context}}

### Interaction Preferences
- **Navigation Style**: {{navigation_preference}}
- **Information Processing**: {{info_processing_style}}
- **Learning Style**: {{learning_preference}}
- **Communication Style**: {{communication_preference}}

### Device & Platform Behavior
{{#each device_behaviors}}
- **{{device}}**: {{usage_pattern}}
{{/each}}

## Scenario & Use Cases

### Primary Use Case
**Scenario**: {{primary_scenario}}

**Context**: {{scenario_context}}

**Steps**:
{{#each primary_steps}}
{{@index}}. {{step}}
{{/each}}

**Desired Outcome**: {{primary_outcome}}

### Secondary Use Cases
{{#each secondary_scenarios}}
#### {{scenario_name}}
**Context**: {{context}}
**Key Actions**: {{key_actions}}
**Success Criteria**: {{success_criteria}}
{{/each}}

## Needs & Requirements

### Functional Needs
{{#each functional_needs}}
- **{{need}}**: {{description}}
  - Priority: {{priority_level}}
  - Frequency: {{usage_frequency}}
{{/each}}

### Emotional Needs
{{#each emotional_needs}}
- **{{need}}**: {{description}}
{{/each}}

### Accessibility Needs
{{#each accessibility_needs}}
- **{{need}}**: {{accommodation_required}}
{{/each}}

### Information Needs
{{#each information_needs}}
- **{{info_type}}**: {{when_needed}} | {{format_preference}}
{{/each}}

## Design Implications

### UI/UX Preferences
- **Visual Style**: {{visual_preferences}}
- **Complexity Level**: {{complexity_tolerance}}
- **Guidance Needs**: {{guidance_preferences}}
- **Error Recovery**: {{error_recovery_preferences}}

### Content Requirements
- **Language Level**: {{language_complexity}}
- **Detail Level**: {{detail_preference}}
- **Media Types**: {{preferred_media}}
- **Help Format**: {{help_format_preference}}

### Interaction Design
- **Input Methods**: {{preferred_input_methods}}
- **Feedback Needs**: {{feedback_preferences}}
- **Confirmation Needs**: {{confirmation_requirements}}
- **Shortcuts**: {{shortcut_usage}}

## Influencers & Decision Making

### Decision Process
{{decision_making_process}}

### Influence Network
{{#each influencers}}
- **{{influencer_type}}**: {{influence_description}}
{{/each}}

### Information Sources
{{#each info_sources}}
- {{source}}: {{trust_level}} | {{usage_context}}
{{/each}}

## Quotes & Insights

### Key Quotes
{{#each quotes}}
*"{{quote}}"* - {{context}}
{{/each}}

### Behavioral Insights
{{#each insights}}
- {{insight}}
{{/each}}

## Persona Validation

### Research Sources
{{#each research_sources}}
- **{{source_type}}**: {{sample_size}} | {{date}} | {{key_findings}}
{{/each}}

### Confidence Level
- **Demographics**: {{demo_confidence}}/10
- **Goals & Motivations**: {{goals_confidence}}/10
- **Pain Points**: {{pain_confidence}}/10
- **Behaviors**: {{behavior_confidence}}/10

### Validation Methods
{{#each validation_methods}}
- {{method}}: {{results_summary}}
{{/each}}

## Anti-Persona Notes

### Who This Is NOT
{{anti_persona_description}}

### Common Misconceptions
{{#each misconceptions}}
- **Myth**: {{myth}}
- **Reality**: {{reality}}
{{/each}}

---

## Persona Usage Guidelines

### When to Reference
- Feature prioritization decisions
- Design concept evaluation
- User flow creation
- Content strategy development
- Accessibility requirement planning

### How to Apply
1. **Feature Planning**: "Would {{persona_name}} find this valuable?"
2. **Design Decisions**: "How would {{persona_name}} expect this to work?"
3. **Content Creation**: "What language would resonate with {{persona_name}}?"
4. **Testing**: "Does this solve {{persona_name}}'s key pain points?"

### Update Schedule
- **Review**: {{review_frequency}}
- **Update**: {{update_frequency}}
- **Validation**: {{validation_frequency}}

---

**Created**: {{creation_date}}  
**Last Updated**: {{last_updated}}  
**Version**: {{version}}  
**Research Team**: {{research_team}}  
**Stakeholder Approval**: {{approval_status}}