# Agent Folder Structure Requirements

## Mandatory Folder Organization

All agents and resources MUST follow this strict folder structure to maintain system coherence and enable proper discovery.

### Command Agents
**Location**: `.claude/commands/`
```
.claude/commands/{command_name}.md
```

**Examples**:
- `.claude/commands/analyst.md` - Ana the business analyst
- `.claude/commands/architect.md` - Archie the architect
- `.claude/commands/pm.md` - Paul the product manager
- `.claude/commands/ux.md` - Ulysses the UX designer

### Sub-Agents
**Location**: `.claude/agents/{parent_command}/`
```
.claude/agents/{parent_command}/{subagent_name}.md
```

**Examples**:
- `.claude/agents/analyst/market-researcher.md`
- `.claude/agents/analyst/competitive-analyzer.md`
- `.claude/agents/pm/product-strategist.md`
- `.claude/agents/architect/adr-manager.md`

**CRITICAL**: Sub-agents MUST be placed in a folder named after their parent command. NEVER place sub-agents directly in `.claude/agents/` root.

### Shared Agents
**Location**: `.claude/agents/shared/`
```
.claude/agents/shared/{agent_name}.md
```

**Examples**:
- `.claude/agents/shared/kb-analyst.md` - Knowledge base operations
- `.claude/agents/shared/qa-reviewer.md` - Quality assurance

Shared agents are used by multiple commands and don't belong to a specific parent.

### Resources
**Location**: `.claude/resources/{parent_command}/`
```
.claude/resources/{parent_command}/{type}/{file}
```

**Resource Types**:
- `templates/` - Document and code templates
- `tasks/` - Workflow and task definitions
- `data/` - Reference data and configurations
- `checklists/` - Validation and quality checklists
- `protocols/` - Interaction and integration protocols
- `cognitive-tools/` - Decision-making tools
- `constraints/` - Constraint definitions

**Examples**:
```
.claude/resources/pm/templates/prd-template.md
.claude/resources/pm/tasks/create-prd.md
.claude/resources/pm/data/elicitation-methods.md
.claude/resources/pm/checklists/pm-validation-checklist.md
.claude/resources/analyst/protocols/molecules/market-research-workflow.md
```

## Folder Creation Rules

### When Creating a New Command
1. Create command file: `.claude/commands/{name}.md`
2. Create sub-agent folder: `.claude/agents/{name}/`
3. Create resource folder: `.claude/resources/{name}/`
4. Create resource subfolders as needed

### When Creating a Sub-Agent
1. Identify parent command
2. Create/verify parent folder exists: `.claude/agents/{parent}/`
3. Create sub-agent file: `.claude/agents/{parent}/{agent}.md`
4. Create resources if needed: `.claude/resources/{parent}/{type}/`

### When Creating Shared Agents
1. Verify agent is truly shared (used by 2+ commands)
2. Create in shared folder: `.claude/agents/shared/{agent}.md`
3. Document which commands use this agent

## Validation Requirements

### Folder Structure Validation
- Commands MUST exist in `.claude/commands/`
- Sub-agents MUST be in parent folders
- No orphaned agents in `.claude/agents/` root
- Resources match agent organization

### Path Consistency
- All references use relative paths from project root
- Paths in command files match actual file locations
- Resource loading paths are correct

### Naming Conventions
- Use kebab-case for all files: `market-researcher.md`
- Folder names match command names exactly
- Resource types use standard names (templates, tasks, data, etc.)

## Current Valid Structure

```
.claude/
├── commands/           # Command orchestrators
│   ├── analyst.md
│   ├── architect.md
│   ├── dev.md
│   ├── pm.md
│   ├── recruiter.md
│   ├── sm.md
│   ├── teacher.md
│   └── ux.md
│
├── agents/            # Sub-agents organized by parent
│   ├── analyst/       # Ana's sub-agents
│   │   ├── market-researcher.md
│   │   ├── competitive-analyzer.md
│   │   ├── documentation-generator.md
│   │   ├── insight-synthesizer.md
│   │   └── research-architect.md
│   │
│   ├── architect/     # Archie's sub-agents
│   │   ├── adr-manager.md
│   │   ├── architecture-reviewer.md
│   │   ├── diagram-generator.md
│   │   ├── pattern-librarian.md
│   │   └── system-designer.md
│   │
│   ├── pm/           # Paul's sub-agents
│   │   ├── product-strategist.md
│   │   ├── stakeholder-coordinator.md
│   │   ├── change-coordinator.md
│   │   ├── document-processor.md
│   │   └── execution-tracker.md
│   │
│   ├── recruiter/    # Rita's sub-agents
│   │   ├── agent-planner.md
│   │   ├── agent-builder.md
│   │   ├── architecture-analyst.md
│   │   ├── constraint-validator.md
│   │   ├── coherence-verifier.md
│   │   ├── pattern-extractor.md
│   │   └── resource-validator.md
│   │
│   ├── teacher/      # Tina's sub-agents
│   │   ├── assessment-generator.md
│   │   ├── concept-difficulty-ranker.md
│   │   ├── learning-path-planner.md
│   │   ├── misconception-detector.md
│   │   ├── zettelkasten-capture.md
│   │   ├── zettelkasten-relationship-mapper.md
│   │   └── zettelkasten-synthesizer.md
│   │
│   ├── ux/          # Ulysses's sub-agents (empty - to be created)
│   │
│   └── shared/      # Shared across commands
│       ├── kb-analyst.md
│       └── qa-reviewer.md
│
└── resources/       # Resources organized by parent
    ├── analyst/     # Ana's resources
    │   ├── cognitive-tools/
    │   ├── data/
    │   ├── protocols/
    │   ├── tasks/
    │   └── templates/
    │
    ├── architect/   # Archie's resources  
    │   ├── checklists/
    │   ├── cognitive-tools/
    │   ├── data/
    │   ├── orchestration/
    │   ├── standards/
    │   ├── system-design/
    │   ├── tasks/
    │   └── templates/
    │
    ├── pm/         # Paul's resources
    │   ├── checklists/
    │   ├── cognitive-tools/
    │   ├── data/
    │   ├── protocols/
    │   ├── tasks/
    │   └── templates/
    │
    ├── recruiter/  # Rita's resources
    │   ├── cells/
    │   ├── cognitive-tools/
    │   ├── constraints/
    │   ├── data/
    │   ├── diagnostics/
    │   └── protocols/
    │
    ├── shared/     # Shared resources
    │   ├── constraint-learning-system.md
    │   ├── multi-phase-verification-framework.md
    │   ├── patterns.md
    │   └── search-strategies.yaml
    │
    ├── sm/         # Sam's resources
    │   ├── checklists/
    │   ├── cognitive-tools/
    │   ├── data/
    │   ├── protocols/
    │   ├── tasks/
    │   └── templates/
    │
    ├── teacher/    # Tina's resources
    │   ├── algorithms/
    │   ├── cells/
    │   ├── cognitive-tools/
    │   ├── constraints/
    │   ├── diagnostics/
    │   ├── protocols/
    │   ├── zettelkasten/      # Zettelkasten sub-agent resources
    │   └── zettelkasten-guide/
    │
    └── ux/         # Ulysses's resources
        ├── guidelines/
        ├── protocols/
        ├── tasks/
        └── templates/
```

## Enforcement in Rita's Workflow

Rita (Recruiter) MUST enforce this structure:

1. **Planning Phase**: Agent-planner determines correct folder location
2. **Building Phase**: Agent-builder creates files in correct folders
3. **Validation Phase**: Resource-validator checks folder structure
4. **Fix Phase**: Correct any misplaced files before completion

## Common Violations to Prevent

❌ **NEVER DO**:
- Place sub-agents directly in `.claude/agents/`
- Mix command and sub-agent files
- Create resources in wrong parent folder
- Use inconsistent naming conventions
- Create duplicate agents in multiple locations

✅ **ALWAYS DO**:
- Verify parent folder exists before creating sub-agents
- Use consistent kebab-case naming
- Match folder names to command names
- Document shared agent usage
- Validate folder structure after creation