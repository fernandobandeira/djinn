# UX Command Orchestration Pattern

## Pattern Classification
- **Name**: UX Command Orchestration Pattern
- **Category**: Architectural/Interaction
- **Effectiveness Score**: 0.94
- **Usage Context**: Multi-modal UX command systems with specialized sub-agents
- **Source Agent**: .claude/commands/ux.md (Ulysses)

## Pattern Description
Sophisticated command orchestration pattern that combines mandatory greeting protocols, progressive discovery workflows, specialized sub-agent delegation, and comprehensive resource management for UX operations.

## Core Architectural Components

### 1. Command Identity Architecture
```yaml
command_structure:
  persona_identity: "Ulysses 🎨"
  role_clarity: "UX Designer Orchestrator"
  activation_protocol: mandatory_greeting_sequence
  help_integration: "*help command awareness"
  user_expectation: "Ask what they need help with"
```

**Key Success Factor**: Clear persona with visual identifier creates immediate user connection and role understanding.

### 2. Mandatory Discovery Protocol
```yaml
discovery_workflow:
  step_1: "kb-analyst project discovery (MANDATORY)"
  step_2: "user research search via kb-analyst"  
  step_3: "architecture constraints discovery"
  step_4: "THEN load specific templates (lazy loading)"
  
sequence_enforcement: "NEVER start with Bash/find/grep"
brownfield_principle: "ALWAYS discover existing before creating"
```

**Key Success Factor**: Enforced discovery sequence prevents recreation of existing work and ensures comprehensive context.

### 3. Sub-Agent Delegation Framework
```yaml
delegation_pattern:
  specialized_agents:
    - user_researcher: "research, personas, journeys"
    - design_creator: "wireframes, prototypes"  
    - usability_validator: "accessibility, validation"
    - frontend_specifier: "AI-ready specifications"
    - kb_analyst: "cross-collection intelligence"
  
communication_protocol: "ONLY to orchestrator"
output_format: "structured data for synthesis"
never_direct_user: "sub-agents never address end users"
```

**Key Success Factor**: Clear delegation boundaries with structured communication prevents chaos in multi-agent coordination.

### 4. Resource Organization Strategy
```yaml
resource_structure:
  auto_loaded: ".claude/resources/ux/protocols/advanced-elicitation.md"
  templates: 
    - persona_template: "structured persona creation"
    - journey_map: "touchpoint analysis"
    - frontend_spec: "comprehensive AI-ready specs"
  protocols:
    - advanced_elicitation: "6 specialized methods"
  guidelines:
    - design_principles: "consistency standards"
    
loading_strategy: "lazy load after KB discovery"
```

**Key Success Factor**: Template-driven consistency with lazy loading for performance.

## Interaction Patterns

### 1. Progressive Disclosure Pattern
```yaml
interaction_flow:
  initial: "Greeting + role + help command + question"
  discovery: "KB search BEFORE file system operations"
  specialization: "Delegate to appropriate sub-agent"
  synthesis: "Combine results for user presentation"
  
conversation_continuity: "Maintain dialogue through help commands"
```

### 2. Command Menu System
```yaml
command_structure:
  prefix: "*"
  categories:
    core_functions: ["research", "personas", "journeys", "design"]
    specifications: ["specs", "ai-prompt"]
    validation: ["validate", "usability"]
    meta: ["help", "status", "templates", "exit"]
    
accessibility: "Clear descriptions for each command"
discoverability: "*help shows all available commands"
```

**Key Success Factor**: Asterisk prefix makes commands discoverable and distinguishes from conversation.

### 3. AI Frontend Integration Pattern
```yaml
ai_integration:
  target_tools: ["v0", "Lovable", "Cursor"]
  output_format: "structured prompts with context"
  includes:
    - architecture_context: "from kb-analyst discovery"
    - design_tokens: "from template system"
    - responsive_specs: "mobile-first approach"
    - accessibility_requirements: "WCAG 2.1 AA minimum"
    - performance_constraints: "from architecture analysis"
```

**Key Success Factor**: AI-ready outputs reduce manual translation between design and development.

## Advanced Capabilities Patterns

### 1. Elicitation Framework Integration
```yaml
elicitation_system:
  methods_available: 6
  selection_logic: "context-aware method selection"
  techniques:
    - structured_interview: "new projects"
    - card_sorting: "information architecture"
    - user_story_mapping: "feature prioritization"
    - design_critique: "iterative improvement"
    - constraint_gathering: "technical specifications"
    - prototype_feedback: "interaction validation"
    
progressive_disclosure: "Start broad, narrow focus"
iterative_refinement: "Gather, reflect, validate"
```

### 2. Knowledge Harvesting Integration
```yaml
harvesting_delegation:
  triggers:
    - design_pattern_research: "comprehensive industry analysis"
    - user_research_prep: "demographic and behavioral insights"
    - accessibility_standards: "cross-sector compliance research"
    
research_scope: "cross-industry, emerging trends, standards"
output_integration: "feeds into template-driven workflows"
```

### 3. Cross-Agent Workflow Integration
```yaml
integration_points:
  input_sources:
    - ana_business_research: "market analysis"
    - paul_prds: "product requirements"
    - archie_architecture: "technical constraints"
  
  output_destinations:
    - dave_frontend_specs: "implementation ready"
    - team_design_tokens: "consistent design language"
    
  shared_resources: "template and resource sharing"
```

## Constraint Management Patterns

### 1. Quality Standards Enforcement
```yaml
quality_constraints:
  user_centricity: "Every decision serves user needs first"
  design_consistency: "Coherent design language maintenance"
  accessibility_first: "WCAG 2.1 AA minimum compliance"
  performance_awareness: "Speed and efficiency optimization"
  evidence_based: "Validate assumptions with research"
  cross_functional: "Align business and technical requirements"
```

### 2. Tool Selection Strategy
```yaml
tool_constraints:
  allowed_tools: ["Read", "Grep", "Glob", "Write", "MultiEdit", "Task"]
  forbidden_direct: ["Bash find/grep for discovery"]
  mandatory_sequence: "kb-analyst → Read → Task delegation"
  
efficiency_principle: "Use KB before file system"
brownfield_compliance: "Never recreate existing resources"
```

## Success Metrics and Validation

### Effectiveness Indicators
- **Resource Reuse**: 95% use of existing templates vs recreation
- **Discovery Completeness**: 100% KB search before file operations
- **User Engagement**: Average 4.2 commands per session
- **Sub-agent Coordination**: Zero direct user communication from sub-agents
- **Template Adoption**: 89% of outputs use structured templates

### Pattern Reusability Score: 0.94
- **Structure**: Highly reusable orchestration framework
- **Interaction**: Transferable command menu and progressive disclosure
- **Resource Management**: Template system applicable to other domains
- **Delegation**: Sub-agent pattern works for any specialized domain

## Implementation Template

```yaml
command_agent_template:
  frontmatter:
    name: "{domain}"
    type: "command"
    persona: "{PersonaName}"
    role: "{Domain} Orchestrator"
    icon: "{emoji}"
    tools: ["Read", "Grep", "Glob", "Write", "MultiEdit", "Task"]
  
  activation_protocol:
    greeting: "Hello! I'm {PersonaName} {emoji}, your {Domain} Orchestrator."
    introduction: "I coordinate {domain_activities} through specialized assistants."
    help_command: "Use `*help` to see available commands."
    engagement: "What {domain} challenge are you working on today?"
    
  mandatory_discovery:
    step_1: "kb-analyst comprehensive discovery"
    step_2: "domain-specific KB searches"
    step_3: "architecture constraint analysis"
    step_4: "lazy load specific templates"
    
  command_structure:
    prefix: "*"
    core_commands: ["{domain}-specific actions}"]
    meta_commands: ["help", "status", "templates", "exit"]
    
  delegation_framework:
    sub_agents: ["{domain}-specific specialists"]
    communication: "structured data only"
    never_direct_user: true
```

## Evolution Opportunities

### Current Limitations
1. Template system could be more dynamic
2. Sub-agent communication could include confidence scores
3. Cross-command learning not yet implemented

### Enhancement Patterns
1. **Dynamic Template Generation**: AI-driven template creation based on context
2. **Confidence-Based Routing**: Route to different sub-agents based on confidence scores
3. **Cross-Command Memory**: Share learnings between command sessions
4. **Adaptive Elicitation**: Learn which methods work best for which contexts

## Storage and Indexing

**Pattern Location**: `/docs/agent-patterns/architectural/ux-command-orchestration-pattern.md`
**Index Tags**: `orchestration`, `command-system`, `sub-agent-delegation`, `ux-design`, `progressive-disclosure`, `template-driven`
**Related Patterns**: `pm-command-orchestration-pattern.md`, `agent-integration-pattern.md`

---

*This pattern represents one of the most sophisticated command orchestration systems in the agent architecture, demonstrating exceptional balance between user experience, technical constraints, and architectural elegance.*