# Agent Workflow and Context Flow Diagrams

## Overview
This document visualizes the complete workflow of how commands delegate to sub-agents, showing context flow and interaction patterns.

## 1. Master Command Architecture

```mermaid
graph TB
    subgraph "User Interface Layer"
        User[User Input]
    end
    
    subgraph "Command Layer (Entry Points)"
        analyst["/analyst<br/>Ana"]
        ux["/ux<br/>Ulysses"]
        pm["/pm<br/>Paul"]
        architect["/architect<br/>Archie"]
        sm["/sm<br/>Sam"]
        dev["/dev<br/>Dave"]
        teacher["/teacher<br/>Tina"]
        recruiter["/recruiter<br/>Rita"]
    end
    
    subgraph "Shared Sub-Agents (Available to All)"
        kb["kb-analyst<br/>Universal KB Operations"]
        qa["qa-reviewer<br/>Quality Assurance"]
        harvester["knowledge-harvester<br/>Documentation Harvesting"]
        plan_val["plan-validator<br/>Plan Validation"]
    end
    
    subgraph "Specialized Sub-Agents"
        subgraph "Analyst Sub-Agents"
            research_arch["research-architect"]
            doc_gen["documentation-generator"]
            competitive["competitive-analyzer"]
            market["market-researcher"]
            insight["insight-synthesizer"]
        end
        
        subgraph "UX Sub-Agents"
            user_research["user-researcher"]
            usability["usability-validator"]
            design["design-creator"]
            frontend["frontend-specifier"]
        end
        
        subgraph "PM Sub-Agents"
            product_strat["product-strategist"]
            stakeholder["stakeholder-coordinator"]
            change_coord["change-coordinator"]
            exec_track["execution-tracker"]
            doc_proc["document-processor"]
        end
        
        subgraph "Architect Sub-Agents"
            system_design["system-designer"]
            arch_review["architecture-reviewer"]
            adr["adr-manager"]
            pattern_lib["pattern-librarian"]
            diagram["diagram-generator"]
        end
        
        subgraph "SM Sub-Agents"
            story_create["story-creator"]
            story_valid["story-validator"]
            sprint_plan["sprint-planner"]
            retro["retrospective-facilitator"]
            change_analyze["change-analyzer"]
        end
        
        subgraph "Teacher Sub-Agents"
            zk_capture["zettelkasten-capture"]
            zk_synth["zettelkasten-synthesizer"]
            zk_relate["zettelkasten-relationship-mapper"]
            gap_analyze["learning-gap-analyzer"]
            misconception["misconception-detector"]
            path_plan["learning-path-planner"]
            difficulty["concept-difficulty-ranker"]
            assessment["assessment-generator"]
        end
        
        subgraph "Recruiter Sub-Agents"
            agent_plan["agent-planner"]
            agent_build["agent-builder"]
            pattern_extract["pattern-extractor"]
            arch_analyze["architecture-analyst"]
            resource_val["resource-validator"]
            constraint_val["constraint-validator"]
            coherence["coherence-verifier"]
        end
    end
    
    User --> analyst & ux & pm & architect & sm & dev & teacher & recruiter
    
    %% Shared agent connections (all commands can use)
    analyst & ux & pm & architect & sm & dev & teacher & recruiter -.-> kb & qa & harvester & plan_val
    
    %% Specific sub-agent connections
    analyst --> research_arch & doc_gen & competitive & market & insight
    ux --> user_research & usability & design & frontend
    pm --> product_strat & stakeholder & change_coord & exec_track & doc_proc
    architect --> system_design & arch_review & adr & pattern_lib & diagram
    sm --> story_create & story_valid & sprint_plan & retro & change_analyze
    teacher --> zk_capture & zk_synth & zk_relate & gap_analyze & misconception & path_plan & difficulty & assessment
    recruiter --> agent_plan & agent_build & pattern_extract & arch_analyze & resource_val & constraint_val & coherence
    
    style kb fill:#e1f5fe
    style qa fill:#e1f5fe
    style harvester fill:#e1f5fe
    style plan_val fill:#e1f5fe
```

## 2. Context Flow Pattern

```mermaid
sequenceDiagram
    participant User
    participant Command as Command Agent
    participant SubAgent as Sub-Agent
    participant Tools as Tools/APIs
    
    User->>Command: Request/Task
    Note over Command: Parse request<br/>Determine delegation
    
    Command->>Command: Prepare context
    Note over Command: Add agent identity<br/>Structure request<br/>Set expectations
    
    Command->>SubAgent: Task(subagent_type, prompt)
    Note over SubAgent: Process request<br/>Use tools<br/>Generate results
    
    SubAgent->>Tools: Read/Write/Search
    Tools-->>SubAgent: Results
    
    SubAgent-->>Command: Structured response
    Note over Command: Translate response<br/>Format for user
    
    Command-->>User: User-friendly result
```

## 3. Analyst (Ana) Workflow

```mermaid
graph LR
    subgraph "Ana's Commands"
        A1["*research {topic}"]
        A2["*brainstorm {topic}"]
        A3["*competitive {company}"]
        A4["*market {topic}"]
        A5["*synthesize"]
        A6["*document"]
    end
    
    subgraph "Delegation Pattern"
        A1 --> research_arch["research-architect<br/>Design methodology"]
        A1 --> harvester["knowledge-harvester<br/>Gather info"]
        
        A2 --> insight["insight-synthesizer<br/>Extract patterns"]
        
        A3 --> competitive["competitive-analyzer<br/>Analysis"]
        A3 --> harvester
        
        A4 --> market["market-researcher<br/>Research"]
        A4 --> harvester
        
        A5 --> insight
        
        A6 --> doc_gen["documentation-generator<br/>Create docs"]
    end
```

## 4. UX Designer (Ulysses) Workflow

```mermaid
graph LR
    subgraph "Ulysses's Commands"
        U1["*persona {type}"]
        U2["*journey {persona}"]
        U3["*wireframe {feature}"]
        U4["*prototype {design}"]
        U5["*design-system"]
        U6["*usability {design}"]
    end
    
    subgraph "Delegation Pattern"
        U1 --> user_research["user-researcher<br/>Create personas"]
        U1 --> harvester["knowledge-harvester<br/>Research users"]
        
        U2 --> user_research
        U2 --> design["design-creator<br/>Journey maps"]
        
        U3 --> design
        U3 --> frontend["frontend-specifier<br/>Specifications"]
        
        U4 --> design
        U4 --> frontend
        
        U5 --> design
        
        U6 --> usability["usability-validator<br/>Validate"]
        U6 --> qa["qa-reviewer<br/>Quality check"]
    end
```

## 5. Product Manager (Paul) Workflow

```mermaid
graph LR
    subgraph "Paul's Commands"
        P1["*prd {feature}"]
        P2["*roadmap"]
        P3["*stakeholder {update}"]
        P4["*prioritize"]
        P5["*track {sprint}"]
        P6["*pivot {reason}"]
    end
    
    subgraph "Delegation Pattern"
        P1 --> product_strat["product-strategist<br/>Create PRD"]
        P1 --> harvester["knowledge-harvester<br/>Research"]
        P1 --> plan_val["plan-validator<br/>Validate"]
        
        P2 --> product_strat
        P2 --> exec_track["execution-tracker<br/>Track progress"]
        
        P3 --> stakeholder["stakeholder-coordinator<br/>Communications"]
        
        P4 --> product_strat
        P4 --> doc_proc["document-processor<br/>Process docs"]
        
        P5 --> exec_track
        
        P6 --> change_coord["change-coordinator<br/>Manage change"]
        P6 --> stakeholder
    end
```

## 6. Architect (Archie) Workflow

```mermaid
graph LR
    subgraph "Archie's Commands"
        AR1["*design {system}"]
        AR2["*adr {decision}"]
        AR3["*review {architecture}"]
        AR4["*pattern {name}"]
        AR5["*diagram {type}"]
    end
    
    subgraph "Delegation Pattern"
        AR1 --> system_design["system-designer<br/>Design system"]
        AR1 --> harvester["knowledge-harvester<br/>Research patterns"]
        
        AR2 --> adr["adr-manager<br/>Manage ADRs"]
        
        AR3 --> arch_review["architecture-reviewer<br/>Review"]
        AR3 --> qa["qa-reviewer<br/>Quality check"]
        
        AR4 --> pattern_lib["pattern-librarian<br/>Manage patterns"]
        AR4 --> kb["kb-analyst<br/>Search KB"]
        
        AR5 --> diagram["diagram-generator<br/>Create diagrams"]
    end
```

## 7. Scrum Master (Sam) Workflow

```mermaid
graph LR
    subgraph "Sam's Commands"
        S1["*story {description}"]
        S2["*sprint {action}"]
        S3["*retrospective"]
        S4["*velocity"]
        S5["*backlog {action}"]
    end
    
    subgraph "Delegation Pattern"
        S1 --> story_create["story-creator<br/>Create stories"]
        S1 --> story_valid["story-validator<br/>Validate"]
        
        S2 --> sprint_plan["sprint-planner<br/>Plan sprint"]
        S2 --> exec_track["execution-tracker<br/>Track"]
        
        S3 --> retro["retrospective-facilitator<br/>Run retro"]
        
        S4 --> exec_track
        S4 --> change_analyze["change-analyzer<br/>Analyze"]
        
        S5 --> story_create
        S5 --> sprint_plan
    end
```

## 8. Teacher (Tina) Workflow

```mermaid
graph LR
    subgraph "Tina's Commands"
        T1["*explain {concept}"]
        T2["*capture {insight}"]
        T3["*synthesize {notes}"]
        T4["*relate {concepts}"]
        T5["*gaps"]
        T6["*misconceptions"]
        T7["*path {topic}"]
        T8["*assess {level}"]
    end
    
    subgraph "Delegation Pattern"
        T1 --> kb["kb-analyst<br/>Search knowledge"]
        T1 --> harvester["knowledge-harvester<br/>Research"]
        
        T2 --> zk_capture["zettelkasten-capture<br/>Capture note"]
        
        T3 --> zk_synth["zettelkasten-synthesizer<br/>Synthesize"]
        
        T4 --> zk_relate["zettelkasten-relationship-mapper<br/>Map relations"]
        
        T5 --> gap_analyze["learning-gap-analyzer<br/>Find gaps"]
        
        T6 --> misconception["misconception-detector<br/>Detect"]
        
        T7 --> path_plan["learning-path-planner<br/>Plan path"]
        T7 --> difficulty["concept-difficulty-ranker<br/>Rank difficulty"]
        
        T8 --> assessment["assessment-generator<br/>Generate"]
    end
```

## 9. Recruiter (Rita) Workflow

```mermaid
graph LR
    subgraph "Rita's Commands"
        R1["*recruit {name}"]
        R2["*plan {name}"]
        R3["*build"]
        R4["*validate"]
        R5["*analyze {agent}"]
        R6["*patterns"]
    end
    
    subgraph "Delegation Pattern"
        R1 --> R2
        R2 --> agent_plan["agent-planner<br/>Plan architecture"]
        
        R3 --> agent_build["agent-builder<br/>Build files"]
        
        R4 --> resource_val["resource-validator<br/>Check resources"]
        R4 --> constraint_val["constraint-validator<br/>Check constraints"]
        R4 --> coherence["coherence-verifier<br/>Check coherence"]
        
        R5 --> arch_analyze["architecture-analyst<br/>Analyze"]
        
        R6 --> pattern_extract["pattern-extractor<br/>Extract patterns"]
        R6 --> kb["kb-analyst<br/>Search patterns"]
    end
```

## 10. Developer (Dave) Workflow

```mermaid
graph LR
    subgraph "Dave's Direct Actions"
        D1["Implementation"]
        D2["TDD"]
        D3["Debugging"]
        D4["Refactoring"]
    end
    
    subgraph "Shared Sub-Agents"
        D1 --> qa["qa-reviewer<br/>Code review"]
        D1 --> kb["kb-analyst<br/>Search examples"]
        
        D2 --> qa
        
        D3 --> harvester["knowledge-harvester<br/>Research solutions"]
        
        D4 --> qa
        D4 --> plan_val["plan-validator<br/>Validate changes"]
    end
    
    Note1["Dave primarily works directly<br/>but uses shared agents<br/>for support"]
    
    D1 -.-> Note1
```

## 11. Context Preparation Examples

### Example 1: Ana delegates to market-researcher

```yaml
User Request: "Research the AI coding assistant market"

Ana's Context Preparation:
  agent_context: "You are being invoked by Ana (analyst command)"
  task: "Comprehensive market research on AI coding assistants"
  scope: "Market size, key players, trends, opportunities"
  format: "Structured market research report"
  return_to: "Ana for synthesis and presentation"
  
market-researcher receives:
  - Clear task definition
  - Scope boundaries
  - Expected output format
  - No direct user interaction needed
```

### Example 2: Rita delegates to agent-planner

```yaml
User Request: "*recruit code-reviewer"

Rita's Context Preparation:
  agent_context: "Request from Rita's orchestration workflow"
  task: "Analyze and plan agent architecture for: code-reviewer"
  requirements:
    - Determine agent type (command/subagent)
    - Analyze complexity and decomposition needs
    - Select appropriate tools and model
    - Calculate constraint balance score
  return_format: "Structured analysis (not for end user)"
  
agent-planner receives:
  - Specific analysis requirements
  - Clear output structure needed
  - Instructions to report to Rita only
```

## 12. Shared Agent Access Pattern

```mermaid
graph TB
    subgraph "Any Command Can Access"
        kb["kb-analyst"]
        qa["qa-reviewer"]
        harvester["knowledge-harvester"]
        plan_val["plan-validator"]
    end
    
    subgraph "Access Pattern"
        cmd["Command Agent"]
        cmd -->|"Search knowledge"| kb
        cmd -->|"Review quality"| qa
        cmd -->|"Research external"| harvester
        cmd -->|"Validate plans"| plan_val
    end
    
    subgraph "Context Includes"
        context["- Invoking agent identity<br/>- Specific search/review scope<br/>- Return format needed<br/>- Integration instructions"]
    end
    
    cmd --> context
    context --> kb & qa & harvester & plan_val
```

## 13. Decision Flow for Delegation

```mermaid
flowchart TD
    Start([User Request]) --> Parse[Parse Command]
    Parse --> IsShared{Needs Shared<br/>Sub-Agent?}
    
    IsShared -->|Yes| WhichShared{Which Type?}
    WhichShared -->|Knowledge| kb[Delegate to kb-analyst]
    WhichShared -->|Quality| qa[Delegate to qa-reviewer]
    WhichShared -->|Research| harvester[Delegate to knowledge-harvester]
    WhichShared -->|Validation| plan_val[Delegate to plan-validator]
    
    IsShared -->|No| IsSpecialized{Needs Specialized<br/>Sub-Agent?}
    
    IsSpecialized -->|Yes| SelectAgent[Select Appropriate<br/>Specialized Agent]
    SelectAgent --> PrepareContext[Prepare Context]
    
    IsSpecialized -->|No| DirectAction[Handle Directly<br/>in Command]
    
    kb & qa & harvester & plan_val --> PrepareContext
    PrepareContext --> Delegate[Invoke via Task Tool]
    Delegate --> ProcessResult[Process Response]
    ProcessResult --> FormatUser[Format for User]
    FormatUser --> End([Return to User])
    
    DirectAction --> End
```

## 14. Error Handling and Recovery

```mermaid
sequenceDiagram
    participant User
    participant Command
    participant SubAgent
    participant Recovery
    
    User->>Command: Request
    Command->>SubAgent: Delegate task
    
    alt Success
        SubAgent-->>Command: Result
        Command-->>User: Formatted response
    else SubAgent Error
        SubAgent-->>Command: Error/Exception
        Command->>Recovery: Determine recovery
        
        alt Can Retry
            Recovery-->>Command: Retry strategy
            Command->>SubAgent: Retry with adjustments
        else Need User Input
            Recovery-->>Command: Request clarification
            Command-->>User: Ask for clarification
        else Use Fallback
            Recovery-->>Command: Fallback agent
            Command->>Recovery: Invoke fallback
        end
    end
```

## 15. Cross-Command Product Development Lifecycle

```mermaid
graph TB
    subgraph "Phase 1: Discovery & Research"
        Start([New Product Idea]) --> Ana1["/analyst<br/>Market Research & Brief"]
        Ana1 --> |"Creates market brief"| Brief[Market Analysis<br/>Opportunity Brief]
        
        Brief --> Teacher1["/teacher<br/>Learn Domain"]
        Teacher1 --> |"Captures knowledge"| Knowledge[Domain Knowledge<br/>Zettelkasten Notes]
    end
    
    subgraph "Phase 2: User Research & Definition"
        Brief --> UX1["/ux<br/>User Research"]
        Knowledge --> UX1
        UX1 --> |"Creates personas"| Personas[User Personas<br/>Journey Maps]
        
        Personas --> PM1["/pm<br/>Product Strategy"]
        Brief --> PM1
        PM1 --> |"Creates PRD"| PRD[Product Requirements<br/>Document]
    end
    
    subgraph "Phase 3: Design & Architecture"
        PRD --> UX2["/ux<br/>UI/UX Design"]
        UX2 --> |"Creates designs"| Mocks[Wireframes<br/>Prototypes<br/>Design System]
        
        PRD --> Arch1["/architect<br/>System Design"]
        Arch1 --> |"Creates architecture"| ADR[Architecture Decisions<br/>System Design<br/>Tech Stack]
        
        Mocks --> UX3["/ux<br/>Usability Testing"]
        UX3 --> |"Validates design"| Validated[Validated Designs]
    end
    
    subgraph "Phase 4: Planning & Stories"
        Validated --> SM1["/sm<br/>Story Creation"]
        ADR --> SM1
        PRD --> SM1
        SM1 --> |"Creates stories"| Stories[User Stories<br/>Acceptance Criteria<br/>Sprint Backlog]
    end
    
    subgraph "Phase 5: Implementation"
        Stories --> Dev1["/dev<br/>Implementation"]
        ADR --> Dev1
        Validated --> Dev1
        Dev1 --> |"Builds features"| Code[Working Code<br/>Tests<br/>Documentation]
        
        Code --> Dev2["/dev<br/>Code Review"]
        Dev2 --> |"Quality assurance"| Reviewed[Reviewed Code]
    end
    
    subgraph "Phase 6: Iteration & Learning"
        Reviewed --> SM2["/sm<br/>Sprint Review"]
        SM2 --> |"Retrospective"| Retro[Lessons Learned<br/>Improvements]
        
        Retro --> PM2["/pm<br/>Track & Adjust"]
        PM2 --> |"Product metrics"| Metrics[Performance Data<br/>User Feedback]
        
        Metrics --> Ana2["/analyst<br/>Analyze Results"]
        Ana2 --> |"Insights"| Insights[Market Insights<br/>Next Opportunities]
        
        Insights --> Teacher2["/teacher<br/>Capture Learning"]
        Teacher2 --> |"Knowledge synthesis"| UpdatedKnowledge[Updated Knowledge Base]
    end
    
    subgraph "Meta: Agent Evolution"
        Retro --> Recruiter["/recruiter<br/>Create/Improve Agents"]
        Recruiter --> |"New capabilities"| NewAgents[Enhanced Agents<br/>New Commands]
        NewAgents -.-> |"Improves process"| Start
    end
    
    %% Feedback loops
    Metrics -.-> |"Pivot decision"| PM1
    Insights -.-> |"New opportunity"| Start
    UpdatedKnowledge -.-> |"Domain expertise"| Ana1
    
    style Start fill:#e8f5e9
    style Code fill:#e3f2fd
    style NewAgents fill:#fff3e0
```

## 16. Detailed Cross-Command Workflow Sequence

```mermaid
sequenceDiagram
    participant User
    participant Analyst as /analyst (Ana)
    participant Teacher as /teacher (Tina)
    participant UX as /ux (Ulysses)
    participant PM as /pm (Paul)
    participant Architect as /architect (Archie)
    participant SM as /sm (Sam)
    participant Dev as /dev (Dave)
    participant Recruiter as /recruiter (Rita)
    
    Note over User: New Product Idea
    
    %% Discovery Phase
    User->>Analyst: Research market opportunity
    activate Analyst
    Analyst->>Analyst: market-researcher sub-agent
    Analyst->>Analyst: competitive-analyzer sub-agent
    Analyst-->>User: Market brief & opportunities
    deactivate Analyst
    
    %% Learning Phase
    User->>Teacher: Learn about domain
    activate Teacher
    Teacher->>Teacher: knowledge-harvester sub-agent
    Teacher->>Teacher: zettelkasten-capture sub-agent
    Teacher-->>User: Domain knowledge captured
    deactivate Teacher
    
    %% User Research Phase
    User->>UX: Create user personas
    activate UX
    UX->>UX: user-researcher sub-agent
    UX-->>User: Personas & journey maps
    deactivate UX
    
    %% Product Definition Phase
    User->>PM: Create PRD from research
    activate PM
    PM->>PM: product-strategist sub-agent
    PM->>PM: document-processor sub-agent
    PM-->>User: Product Requirements Document
    deactivate PM
    
    %% Design Phase
    User->>UX: Create UI mockups from PRD
    activate UX
    UX->>UX: design-creator sub-agent
    UX->>UX: frontend-specifier sub-agent
    UX-->>User: Wireframes & prototypes
    deactivate UX
    
    User->>UX: Validate usability
    activate UX
    UX->>UX: usability-validator sub-agent
    UX-->>User: Usability report
    deactivate UX
    
    %% Architecture Phase
    User->>Architect: Design system architecture
    activate Architect
    Architect->>Architect: system-designer sub-agent
    Architect->>Architect: adr-manager sub-agent
    Architect-->>User: ADRs & system design
    deactivate Architect
    
    %% Planning Phase
    User->>SM: Create user stories
    activate SM
    SM->>SM: story-creator sub-agent
    SM->>SM: story-validator sub-agent
    SM-->>User: User stories with acceptance criteria
    deactivate SM
    
    User->>SM: Plan sprint
    activate SM
    SM->>SM: sprint-planner sub-agent
    SM-->>User: Sprint backlog
    deactivate SM
    
    %% Implementation Phase
    User->>Dev: Implement features
    activate Dev
    Dev->>Dev: Direct implementation
    Dev->>Dev: qa-reviewer sub-agent
    Dev-->>User: Working code with tests
    deactivate Dev
    
    %% Review Phase
    User->>SM: Run retrospective
    activate SM
    SM->>SM: retrospective-facilitator sub-agent
    SM-->>User: Lessons learned
    deactivate SM
    
    %% Tracking Phase
    User->>PM: Track progress
    activate PM
    PM->>PM: execution-tracker sub-agent
    PM-->>User: Metrics & adjustments
    deactivate PM
    
    %% Meta Evolution
    User->>Recruiter: Create new agent for gap
    activate Recruiter
    Recruiter->>Recruiter: agent-planner sub-agent
    Recruiter->>Recruiter: agent-builder sub-agent
    Recruiter->>Recruiter: All validators
    Recruiter-->>User: New agent created
    deactivate Recruiter
```

## 17. Command Handoff Patterns

```mermaid
graph LR
    subgraph "Information Flow Between Commands"
        Ana["/analyst<br/>Market Brief"] --> UX1["/ux<br/>User Research"]
        Ana --> PM["/pm<br/>Product Strategy"]
        
        UX1 --> PM
        PM --> UX2["/ux<br/>Design"]
        PM --> Arch["/architect<br/>Architecture"]
        
        UX2 --> SM["/sm<br/>Stories"]
        Arch --> SM
        PM --> SM
        
        SM --> Dev["/dev<br/>Implementation"]
        UX2 --> Dev
        Arch --> Dev
        
        Dev --> SM2["/sm<br/>Review"]
        SM2 --> PM2["/pm<br/>Track"]
        PM2 --> Ana2["/analyst<br/>Analyze"]
    end
    
    subgraph "Learning & Meta Layer"
        Teacher["/teacher<br/>Domain Learning"] -.-> Ana & UX1 & PM & Arch & Dev
        Recruiter["/recruiter<br/>Agent Creation"] -.-> |"Improves all"| Ana & UX1 & PM & Arch & SM & Dev & Teacher
    end
    
    subgraph "Shared Resources"
        KB["kb-analyst"] -.-> |"Knowledge access"| ALL[All Commands]
        QA["qa-reviewer"] -.-> |"Quality checks"| ALL
        Harvester["knowledge-harvester"] -.-> |"Research"| ALL
        Validator["plan-validator"] -.-> |"Validation"| ALL
    end
```

## 18. Typical Product Development Scenarios

### Scenario 1: New Feature Development

```yaml
Flow:
  1. Analyst:
     - Research competitive features
     - Identify market opportunity
     - Output: Feature opportunity brief
  
  2. Teacher:
     - Learn about technical domain
     - Capture key concepts
     - Output: Knowledge notes
  
  3. UX:
     - Research user needs
     - Create personas
     - Output: User personas, journey maps
  
  4. PM:
     - Synthesize research into PRD
     - Define success metrics
     - Output: Product Requirements Document
  
  5. UX:
     - Design wireframes
     - Create prototypes
     - Validate usability
     - Output: Validated designs
  
  6. Architect:
     - Design technical solution
     - Make architecture decisions
     - Output: ADRs, system design
  
  7. SM:
     - Break down into stories
     - Plan sprints
     - Output: Sprint backlog
  
  8. Dev:
     - Implement features
     - Write tests
     - Output: Working code
  
  9. SM:
     - Review sprint
     - Run retrospective
     - Output: Improvements
  
  10. PM:
      - Track metrics
      - Adjust strategy
      - Output: Product updates
```

### Scenario 2: Technical Debt Reduction

```yaml
Flow:
  1. Dev:
     - Identify technical debt
     - Output: Debt assessment
  
  2. Architect:
     - Review architecture
     - Propose improvements
     - Output: Refactoring plan
  
  3. PM:
     - Prioritize against features
     - Output: Technical debt PRD
  
  4. SM:
     - Create refactoring stories
     - Plan debt reduction sprint
     - Output: Technical backlog
  
  5. Dev:
     - Implement improvements
     - Output: Refactored code
  
  6. Teacher:
     - Document patterns learned
     - Output: Knowledge base updates
```

### Scenario 3: Agent Capability Gap

```yaml
Flow:
  1. Any Command:
     - Identifies missing capability
     - Output: Gap description
  
  2. Recruiter:
     - Plans new agent architecture
     - Output: Agent plan
  
  3. Recruiter:
     - Builds agent files
     - Output: New agent
  
  4. Recruiter:
     - Validates agent
     - Output: Validated agent
  
  5. Teacher:
     - Documents new capability
     - Output: Usage documentation
  
  6. Original Command:
     - Uses new agent
     - Output: Enhanced workflow
```

## 19. Meta-Command: Recruiter Workflow

```mermaid
graph TB
    subgraph "Recruiter: Agent Creation Pipeline"
        Need[Capability Gap Identified] --> Plan[*plan agent-name]
        Plan --> AP[agent-planner<br/>Analyzes requirements]
        AP --> Approve{User Approves?}
        
        Approve -->|Yes| Build[*build]
        Approve -->|No| Refine[Refine requirements]
        Refine --> Plan
        
        Build --> AB[agent-builder<br/>Creates files]
        AB --> Validate[*validate]
        
        Validate --> RV[resource-validator<br/>Check files exist]
        Validate --> CV[constraint-validator<br/>Check balance]
        Validate --> CoV[coherence-verifier<br/>Check integration]
        
        RV & CV & CoV --> Valid{All Pass?}
        
        Valid -->|Yes| Learn[Extract patterns]
        Valid -->|No| Fix[*fix issues]
        Fix --> Build
        
        Learn --> PE[pattern-extractor<br/>Learn patterns]
        PE --> KB[Update knowledge base]
        KB --> Complete[New Agent Ready]
        
        Complete --> Enhance[Enhances all workflows]
    end
    
    style Need fill:#ffebee
    style Complete fill:#e8f5e9
    style Enhance fill:#fff3e0
```

## 20. Teacher: Learning Integration

```mermaid
graph LR
    subgraph "Teacher: Knowledge Management"
        Learn[New Concept] --> Capture[*capture insight]
        Capture --> ZK1[zettelkasten-capture<br/>Atomic note]
        
        ZK1 --> Relate[*relate concepts]
        Relate --> ZK2[zettelkasten-relationship-mapper<br/>Link notes]
        
        ZK2 --> Synthesize[*synthesize notes]
        Synthesize --> ZK3[zettelkasten-synthesizer<br/>Create synthesis]
        
        ZK3 --> KB[Knowledge Base]
    end
    
    subgraph "Learning Detection"
        Gaps[*gaps] --> GA[learning-gap-analyzer<br/>Find gaps]
        Misconceptions[*misconceptions] --> MD[misconception-detector<br/>Find errors]
        
        GA & MD --> Path[*path topic]
        Path --> PP[learning-path-planner<br/>Create path]
        PP --> Learn
    end
    
    subgraph "Knowledge Application"
        KB --> |"Informs"| AllCommands[All Other Commands]
        KB --> |"Examples for"| Dev[Developer]
        KB --> |"Patterns for"| Architect[Architect]
        KB --> |"Insights for"| Analyst[Analyst]
    end
```

## Summary

This comprehensive diagram set shows:

1. **Master Architecture**: How all commands and sub-agents are organized
2. **Context Flow**: How context is prepared and passed between layers
3. **Command-Specific Workflows**: Each command's delegation patterns
4. **Shared Agent Access**: How any command can use shared sub-agents
5. **Decision Flows**: Logic for determining when to delegate
6. **Error Handling**: Recovery patterns for failures

### Key Patterns:

- **Commands are orchestrators**: They parse, delegate, and format responses
- **Sub-agents are specialists**: They do one thing well with minimal tools
- **Context is structured**: Each delegation includes clear context and expectations
- **Shared agents are universal**: Available to all commands for common needs
- **Users see unified responses**: Sub-agents never directly address users

This architecture enables:
- **Modularity**: Easy to add/modify agents
- **Reusability**: Shared agents reduce duplication
- **Maintainability**: Clear boundaries and responsibilities
- **Scalability**: Can grow without increasing complexity
- **Consistency**: Unified patterns across all commands