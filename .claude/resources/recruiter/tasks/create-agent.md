# Create Agent Task

## Purpose
Guide the complete creation of a new Claude Code agent/command, following established patterns and best practices from successful agents like Ana and Archie.

## Agent Creation Workflow

### Phase 1: Requirements Gathering

#### Initial Questions:
1. **Agent Name & Identity**
   - What is the agent's name? (e.g., Ana, Archie)
   - What's their nickname/full title?
   - What icon represents them? (emoji)

2. **Purpose & Scope**
   - What specific problem does this agent solve?
   - What is their primary responsibility?
   - What are they NOT responsible for?
   - Who will use this agent?

3. **Capabilities Needed**
   - What tasks will the agent perform?
   - What decisions will they make?
   - What documents will they create?
   - What systems will they interact with?

4. **Knowledge Base Usage**
   - What will they search for?
   - What will they index?
   - Which collections are relevant?

### Phase 2: Validation

#### Check for Conflicts:
```bash
# Search for existing similar agents
./.vector_db/kb search "[agent purpose]" --collection documentation
./.vector_db/kb search "command [name]"

# Check command availability
ls .claude/commands/ | grep -i [name]
```

#### Validate Use Case:
- Is this truly needed?
- Could existing agents handle this?
- Is the scope well-defined?
- Will users understand when to use it?

### Phase 3: Persona Definition

#### Core Persona Elements:
```yaml
persona:
  identity: [One sentence description]
  role: [Official title/function]
  style: [Adjectives describing interaction style]
  focus: [Primary areas of expertise]
  
  core_principles:
    - [Principle 1]: [Description]
    - [Principle 2]: [Description]
    - [Continue for 5-10 principles]
```

#### Interaction Style:
- Formal vs Casual
- Directive vs Facilitative
- Expert vs Collaborative
- Verbose vs Concise

### Phase 4: Command Design

#### Command Structure:
```markdown
### Core Commands
- `*help` - Show available commands
- `*status` - Show current context
- `*exit` - Exit agent mode

### Primary Commands
- `*[main-action]` - [Description]
- `*[secondary-action]` - [Description]

### Support Commands
- `*kb-search` - Search knowledge base
- `*kb-index` - Index outputs
```

#### Command Naming:
- Use verb-noun pattern (e.g., create-brief)
- Keep consistent with agent's domain
- Ensure clarity of purpose
- Avoid conflicts with existing commands

### Phase 5: Resource Planning

#### Determine Needed Resources:

**Tasks** (Complex Workflows):
- Multi-step processes
- Interactive workflows
- Guided procedures
- Complex decision trees

**Templates** (Output Formats):
- Document templates
- Report structures
- Analysis formats
- Standardized outputs

**Data** (Reference Information):
- Technique libraries
- Pattern collections
- Best practices
- Domain knowledge

**Checklists** (Validation):
- Quality reviews
- Completion checks
- Validation steps
- Prerequisites

### Phase 6: File Structure Creation

#### Create Directory Structure:
```bash
# Create agent directories
mkdir -p .claude/resources/[agent-name]/{tasks,templates,data,checklists}
```

#### File Organization:
```
.claude/
├── commands/
│   └── [agent-name].md         # Main command file
└── resources/
    └── [agent-name]/
        ├── tasks/
        │   ├── [task1].md       # Complex workflow
        │   └── [task2].md       # Another workflow
        ├── templates/
        │   ├── [template1].md   # Output template
        │   └── [template2].md   # Another template
        ├── data/
        │   └── [reference].md   # Reference data
        └── checklists/
            └── [checklist].md   # Review checklist
```

### Phase 7: Command File Creation

#### Command File Structure:
```markdown
# [Agent Name] Agent - [Nickname]

## Activation
You are [Name], [role description]. Your role is to [primary purpose].

## Core Configuration
```yaml
agent:
  name: [Name]
  role: [Role]
  icon: [Emoji]
  style: [Style descriptors]

persona:
  [Persona definition]

resource_files:
  tasks:
    [task_name]: .claude/resources/[agent]/tasks/[file].md
  templates:
    [template_name]: .claude/resources/[agent]/templates/[file].md
  data:
    [data_name]: .claude/resources/[agent]/data/[file].md
  checklists:
    [checklist_name]: .claude/resources/[agent]/checklists/[file].md
```

## Commands
[Command list with * prefix]

## Interaction Protocol
[How agent interacts]

## Task Execution
[How tasks are executed with lazy loading]

## Knowledge Base Integration
[KB search and index patterns]

## Remember
[Key behavioral reminders]
```

### Phase 8: Resource File Creation

#### For Each Task:
1. Define clear purpose
2. Create step-by-step workflow
3. Include KB integration points
4. Add interaction prompts
5. Document outputs

#### For Each Template:
1. Create standard structure
2. Include placeholders
3. Add instructions
4. Provide examples

#### For Each Data File:
1. Organize reference information
2. Use consistent formatting
3. Include examples
4. Keep maintainable

### Phase 9: Knowledge Base Integration

#### Search Patterns:
```markdown
### Before [Action]
When user requests `*[command]`:
1. **FIRST search knowledge base**:
   - `./.vector_db/kb search "[topic]" --collection [collection]`
   - `./.vector_db/kb search "[related]" --collection [collection]`
2. THEN load: `.claude/resources/[agent]/[resource]`
3. [Continue with task]
```

#### Index Patterns:
```markdown
### After [Action]
After completing [task]:
1. Save output to appropriate location
2. Index in knowledge base:
   - `./.vector_db/kb index --path [output path]`
3. Update statistics:
   - `./.vector_db/kb stats`
```

### Phase 10: Testing & Validation

#### Test Checklist:
- [ ] Command file loads without errors
- [ ] All commands work as expected
- [ ] Resources load when needed (not before)
- [ ] KB search executes properly
- [ ] KB indexing works correctly
- [ ] Interactive dialogue maintained
- [ ] Numbered options displayed
- [ ] Outputs generated correctly

#### Validation Questions:
1. Does the agent solve the intended problem?
2. Is the workflow intuitive?
3. Are outputs useful and complete?
4. Is KB integration working?
5. Does it follow established patterns?

### Phase 11: Documentation

#### Update CLAUDE.md:
```markdown
### Custom Commands
- **[agent]**: Use `/[agent]` to invoke [description]
```

#### Create README:
Document in agent's resource folder:
- Purpose and capabilities
- Command reference
- Usage examples
- Best practices

### Phase 12: Final Review

#### Quality Checklist:
- [ ] Single, clear purpose
- [ ] No duplicate functionality
- [ ] Proper file structure
- [ ] KB integration complete
- [ ] Lazy loading implemented
- [ ] Interactive design
- [ ] Numbered options used
- [ ] Brownfield awareness
- [ ] Well-documented
- [ ] Tested thoroughly

## Common Patterns to Follow

### Pattern 1: Lazy Loading
```markdown
When user requests `*[command]`:
1. [Do KB search first]
2. THEN load: `[resource path]`
3. [Execute task]
```

### Pattern 2: KB First
```markdown
1. **FIRST search knowledge base**:
   - Check what exists
   - Find related patterns
2. Then proceed with task
```

### Pattern 3: Numbered Options
```markdown
Please select an option:
1. [Option 1]
2. [Option 2]
3. [Option 3]

Your choice (1-3):
```

### Pattern 4: Interactive Dialogue
- Ask questions
- Wait for responses
- Build on answers
- Confirm understanding

## Anti-Patterns to Avoid

### ❌ Loading Everything
Don't use @ syntax in command file
Don't load all resources at activation

### ❌ Monolithic Agents
Don't combine multiple responsibilities
Don't create "do everything" agents

### ❌ Ignoring KB
Don't skip KB searches
Don't forget to index outputs

### ❌ Poor Interaction
Don't monologue at users
Don't skip numbered options
Don't assume without asking

## Example Creation

### Good Example:
```markdown
# Creating a "reviewer" agent
1. Check KB for existing review functionality
2. Define focused purpose: code review
3. Create interactive review workflow
4. Use templates for review output
5. Index reviews in KB
6. Test with actual code
```

### Bad Example:
```markdown
# Creating a "super-developer" agent
❌ Combines coding, reviewing, testing, deploying
❌ No clear boundaries
❌ Too many responsibilities
❌ Users won't know when to use it
```

## Success Criteria

An agent is successful when:
1. Users immediately understand its purpose
2. It solves a specific problem well
3. It integrates seamlessly with KB
4. It follows established patterns
5. It maintains consistent interaction
6. It produces valuable outputs
7. It's maintainable and extensible

## Remember
- Every agent should have a clear, single purpose
- Always check KB before creating anything
- Follow patterns from successful agents
- Test thoroughly before considering complete
- Document everything for future maintenance