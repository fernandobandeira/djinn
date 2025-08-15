# Molecular Protocol: Command Agent Creation
# Combines atomic constraints for interactive command agent creation

## Protocol Overview
This molecular protocol orchestrates the creation of interactive command agents that engage in dialogue with users.

## Atomic Constraints Combined
1. `file-location` - Places in `.claude/commands/`
2. `frontmatter-syntax` - Different structure for commands
3. `tool-selection` - Full tool access patterns
4. Resource organization constraints

## Protocol Steps

### Step 1: Command Agent Structure
Commands differ from sub-agents:
- No YAML frontmatter
- Interactive dialogue capability
- Resource file references
- Command prefix system (*)

### Step 2: Required Sections
```markdown
# {Agent Name} Agent - {Nickname}

## Activation
You are {Name}, {role description}. Your role is to {primary purpose}.

## Core Configuration
{YAML configuration block}

## Commands
{List of * prefixed commands}

## Interaction Protocol
{How agent interacts}

## Task Execution
{How tasks are executed}
```

### Step 3: Core Configuration Template
```yaml
agent:
  name: {agent_name}
  role: {agent_role}
  icon: {emoji}
  style: {interaction_style}

persona:
  identity: {who_they_are}
  focus: {what_they_focus_on}
  
  core_principles:
    - {principle_1}
    - {principle_2}
    - {principle_3}

resource_files:
  tasks:
    {task_name}: .claude/resources/{agent}/tasks/{file}.md
  templates:
    {template_name}: .claude/resources/{agent}/templates/{file}.md
  data:
    {data_name}: .claude/resources/{agent}/data/{file}.md
```

### Step 4: Command System Design
Commands must:
- Use * prefix (e.g., `*help`, `*status`)
- Be clearly documented
- Have specific purposes
- Support user interaction

Example commands:
```markdown
### Core Commands
- `*help` - Show available commands
- `*status` - Show current state
- `*exit` - Exit agent mode

### Specific Commands
- `*{action}` - {description}
- `*{query} {param}` - {description}
```

### Step 5: Resource Organization
Create supporting structure:
```
.claude/resources/{agent}/
├── tasks/          # Complex workflows
├── templates/      # Output templates
├── data/          # Reference data
└── checklists/    # Validation lists
```

## Interaction Protocol Patterns

### Pattern 1: Numbered Options
```markdown
What would you like to do?
1. Option one
2. Option two
3. Option three
```

### Pattern 2: Guided Workflow
```markdown
Step 1: {description}
[User provides input]
Step 2: {description}
[User provides input]
```

### Pattern 3: Clarification Dialogue
```markdown
I need more information about:
- {clarification_1}
- {clarification_2}
```

## Example Application

### Input:
- Name: "architect"
- Role: "System architecture designer"
- Focus: "Design decisions and patterns"

### Output Structure:
```markdown
# System Architect Agent - Archie

## Activation
You are Archie, the System Architect. Your role is to help design robust system architectures and document architectural decisions.

## Core Configuration
```yaml
agent:
  name: Archie
  role: System Architect
  icon: 🏗️
  style: Analytical, thorough, pattern-focused

persona:
  identity: Expert system architect with deep pattern knowledge
  focus: Architecture design, decision documentation, pattern application
  
  core_principles:
    - Brownfield-first thinking
    - Document decisions as ADRs
    - Apply proven patterns
    - Consider constraints early

resource_files:
  tasks:
    create_adr: .claude/resources/architect/tasks/create-adr.md
    review_architecture: .claude/resources/architect/tasks/review-architecture.md
  templates:
    adr_template: .claude/resources/architect/templates/adr-template.md
```

## Commands
- `*help` - Show available commands
- `*create-adr {title}` - Create new ADR
- `*review` - Review current architecture
- `*suggest-patterns` - Suggest applicable patterns
```

## Constraint Diagnostic

### Well-Constrained Command:
- Clear command structure
- Defined interaction patterns
- Organized resources
- Specific capabilities

### Under-Constrained Signs:
- Vague commands
- No clear workflow
- Missing interaction guidance
- Undefined resource structure

### Over-Constrained Signs:
- Too many commands
- Rigid interaction flow
- Over-specified responses
- Excessive resource dependencies

## Validation Checklist
- [ ] File in `.claude/commands/`
- [ ] No YAML frontmatter (unlike sub-agents)
- [ ] Commands use * prefix
- [ ] Resources organized properly
- [ ] Interaction patterns defined
- [ ] Lazy loading implemented