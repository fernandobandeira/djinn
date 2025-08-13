# Agent Creation Best Practices

## Lessons from Successful Agents

### From Ana the Analyst
- **Interactive Facilitation**: Never generate for users, guide them
- **Multiple Techniques**: Offer variety in approaches
- **Numbered Options**: Always use 1-9 for selections
- **Document Everything**: Capture all outputs for KB
- **Elicitation Excellence**: Offer refinement after each section

### From Archie the Architect  
- **Brownfield First**: System always exists, check what's there
- **ADR Everything**: Document every significant decision
- **Multiple Options**: Present 2-3 approaches with trade-offs
- **Pattern Awareness**: Reference and build on patterns
- **Migration Paths**: Always provide forward and rollback paths

## Core Design Principles

### 1. Single Responsibility Principle
Each agent should have ONE clear purpose:
- ✅ Ana: Business analysis and research
- ✅ Archie: System architecture and design
- ❌ "SuperAgent": Does everything (too broad)

### 2. Knowledge Base Integration

#### Always Search First:
```bash
# Before any task
./.vector_db/kb search "[topic]" --collection documentation
./.vector_db/kb search "[existing work]" --collection architecture
```

#### Always Index After:
```bash
# After creating outputs
./.vector_db/kb index --path ./docs/[output-location]/
```

#### Collection Mapping:
- `architecture`: ADRs, system designs, patterns
- `documentation`: Research, brainstorms, briefs
- `code`: Implementations, functions, classes
- `api`: Endpoints, contracts, specifications
- `config`: Settings, environment configs
- `tests`: Test cases, test patterns

### 3. Resource Organization

#### Directory Structure:
```
.claude/
├── commands/
│   └── [agent].md           # ONE command file
└── resources/
    └── [agent]/
        ├── tasks/           # Complex multi-step workflows
        ├── templates/       # Output document templates
        ├── data/           # Reference information
        └── checklists/     # Validation and review
```

#### When to Use Each Type:

**Tasks**: Multi-step processes
- Brainstorming facilitation
- System design workflow
- Document creation process
- Analysis procedures

**Templates**: Standardized outputs
- Report structures
- Document formats
- Analysis frameworks
- Decision records

**Data**: Reference material
- Technique libraries
- Pattern catalogs
- Best practices
- Domain knowledge

**Checklists**: Validation
- Quality reviews
- Completion checks
- Prerequisites
- Success criteria

### 4. Lazy Loading Pattern

#### ❌ Wrong Way (Loads Everything):
```markdown
## Task Execution
@.claude/resources/agent/tasks/task1.md
@.claude/resources/agent/data/data1.md
```

#### ✅ Right Way (Loads When Needed):
```markdown
When user requests `*command`:
1. THEN load: `.claude/resources/agent/tasks/task1.md`
2. Execute task
```

### 5. Interactive Design

#### Always Engage:
- Ask clarifying questions
- Wait for user responses  
- Build on their input
- Confirm understanding

#### Never Monologue:
- ❌ Generating complete documents without input
- ❌ Making assumptions without confirmation
- ❌ Proceeding without user engagement

### 6. Numbered Options Protocol

#### Standard Format:
```
Please select an option:
1. [First option]
2. [Second option]
3. [Third option]
...
9. [Ninth option]
0. [Proceed/Skip/Cancel]

Your choice (0-9):
```

#### When to Use:
- Technique selection
- Approach choices
- Refinement options
- Navigation decisions

### 7. Command Naming Conventions

#### Command Prefixes:
All commands use `*` prefix:
- `*help` - Always included
- `*status` - Show current state
- `*exit` - Leave agent mode

#### Naming Patterns:
- Verb-noun: `*create-brief`, `*analyze-market`
- Single word for simple: `*brainstorm`, `*research`
- Descriptive: Clear purpose from name alone

### 8. Persona Development

#### Essential Elements:
```yaml
persona:
  identity: [One-line description]
  role: [Official function]
  style: [Personality traits]
  focus: [Areas of expertise]
  core_principles:
    - [Principle]: [Explanation]
```

#### Personality Consistency:
- Ana: Facilitative, curious, creative
- Archie: Pragmatic, thorough, systematic
- Rita: Methodical, instructive, pattern-aware

### 9. Error Handling

#### Graceful Failures:
- Check if files exist before loading
- Validate inputs before processing
- Provide helpful error messages
- Suggest corrections

#### Example:
```markdown
If file not found:
- Inform user clearly
- Suggest alternatives
- Provide recovery path
```

### 10. Documentation Standards

#### In-Code Documentation:
- Clear section headers
- Purpose statements
- Step-by-step instructions
- Examples where helpful

#### User Documentation:
- Command reference
- Usage examples
- Common workflows
- Troubleshooting

## Common Pitfalls to Avoid

### 1. Scope Creep
- **Problem**: Agent tries to do too much
- **Solution**: Focus on single responsibility
- **Example**: Don't combine analyst + architect

### 2. KB Ignorance  
- **Problem**: Not checking existing work
- **Solution**: Always search before creating
- **Example**: Check for existing patterns before designing

### 3. Resource Overload
- **Problem**: Loading everything at start
- **Solution**: Lazy loading pattern
- **Example**: Load task only when command invoked

### 4. Poor Interaction
- **Problem**: Monologuing at users
- **Solution**: Interactive dialogue
- **Example**: Ask, wait, respond, build

### 5. Inconsistent Patterns
- **Problem**: Each agent works differently
- **Solution**: Follow established patterns
- **Example**: Always use numbered options

### 6. Missing KB Integration
- **Problem**: Outputs not indexed
- **Solution**: Index after every creation
- **Example**: Index brainstorm results

### 7. Weak Personas
- **Problem**: Generic, undefined personality
- **Solution**: Strong, consistent character
- **Example**: Ana's facilitative style

### 8. No Migration Path
- **Problem**: Can't evolve or rollback
- **Solution**: Plan for change
- **Example**: ADRs with supersede capability

## Quality Metrics

### Good Agent Indicators:
- Users understand purpose immediately
- Clear when to use vs other agents
- Consistent interaction patterns
- Valuable, indexed outputs
- Follows established patterns
- Maintainable structure
- Testable workflows

### Bad Agent Indicators:
- Unclear purpose or scope
- Overlaps with existing agents
- Inconsistent interaction
- No KB integration
- Monolithic design
- Hard to maintain
- Untested workflows

## Testing Checklist

### Functional Testing:
- [ ] All commands work
- [ ] Resources load properly
- [ ] KB search executes
- [ ] KB indexing works
- [ ] Outputs generate correctly

### User Experience Testing:
- [ ] Clear purpose
- [ ] Intuitive workflow
- [ ] Helpful interactions
- [ ] Valuable outputs
- [ ] Consistent patterns

### Integration Testing:
- [ ] Works with KB
- [ ] Fits with other agents
- [ ] Follows conventions
- [ ] Updates CLAUDE.md

## Evolution Strategy

### Adding Features:
1. Check if it fits agent's purpose
2. Consider creating new agent instead
3. Follow existing patterns
4. Update documentation
5. Test thoroughly

### Deprecating Features:
1. Mark as deprecated in docs
2. Provide migration path
3. Update agent to warn users
4. Remove after grace period

### Major Changes:
1. Create ADR for decision
2. Plan migration strategy
3. Update incrementally
4. Maintain backward compatibility
5. Document thoroughly

## Success Patterns

### Pattern: Knowledge-Driven Design
```
1. Search KB for context
2. Build on what exists
3. Document decisions
4. Index new knowledge
```

### Pattern: Interactive Facilitation
```
1. Ask focused questions
2. Wait for user input
3. Build on responses
4. Guide to solution
```

### Pattern: Incremental Evolution
```
1. Start simple
2. Add complexity as needed
3. Document each change
4. Maintain compatibility
```

### Pattern: Consistent Experience
```
1. Follow UI patterns
2. Use standard commands
3. Maintain personality
4. Deliver expected outputs
```

## Remember
- Patterns exist for good reasons
- Consistency improves usability
- KB integration is mandatory
- Interactive > Generative
- Document everything
- Test before shipping
- Evolution > Revolution