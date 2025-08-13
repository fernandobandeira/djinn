# Agent Creation Checklist

## Pre-Creation Validation

### Requirements Check
- [ ] Clear, single purpose defined
- [ ] Target users identified
- [ ] Use cases documented
- [ ] Success criteria established

### Conflict Check
- [ ] KB searched for similar functionality
- [ ] No duplicate agents exist
- [ ] Command name is unique
- [ ] Scope doesn't overlap existing agents

### Justification
- [ ] Problem clearly stated
- [ ] Existing solutions inadequate
- [ ] New agent adds clear value
- [ ] Maintenance commitment understood

## Persona Development

### Identity
- [ ] Name chosen (short, memorable)
- [ ] Nickname established
- [ ] Icon selected (appropriate emoji)
- [ ] Role clearly defined

### Character
- [ ] Personality traits defined (3-5 adjectives)
- [ ] Interaction style determined
- [ ] Communication tone set
- [ ] Expertise areas listed

### Principles
- [ ] 5-10 core principles defined
- [ ] Each principle has clear description
- [ ] KB integration included
- [ ] Interactive engagement emphasized

## File Structure

### Directory Creation
- [ ] `.claude/commands/[agent].md` created
- [ ] `.claude/resources/[agent]/` created
- [ ] `tasks/` subdirectory created
- [ ] `templates/` subdirectory created
- [ ] `data/` subdirectory created
- [ ] `checklists/` subdirectory created

### Resource Planning
- [ ] Tasks identified and listed
- [ ] Templates needed determined
- [ ] Reference data gathered
- [ ] Checklists required noted

## Command Design

### Core Commands
- [ ] `*help` command included
- [ ] `*status` command included
- [ ] `*exit` command included
- [ ] All commands use `*` prefix

### Functional Commands
- [ ] Primary commands defined
- [ ] Secondary commands planned
- [ ] Support commands included
- [ ] KB commands integrated

### Command Naming
- [ ] Names are clear and descriptive
- [ ] Follow verb-noun pattern where applicable
- [ ] Consistent with agent's domain
- [ ] No conflicts with existing commands

## Knowledge Base Integration

### Search Integration
- [ ] KB searches before all major actions
- [ ] Appropriate collections identified
- [ ] Search patterns documented
- [ ] Cross-collection searches planned

### Index Integration
- [ ] Output locations defined
- [ ] Index commands included
- [ ] Collection mapping correct
- [ ] Update patterns established

## Resource Development

### Tasks
- [ ] Each task has clear purpose
- [ ] Step-by-step workflow defined
- [ ] KB integration points marked
- [ ] Interactive prompts included
- [ ] Output formats specified

### Templates
- [ ] Standard structures created
- [ ] Placeholders defined
- [ ] Instructions included
- [ ] Examples provided

### Data Files
- [ ] Reference information organized
- [ ] Consistent formatting used
- [ ] Examples included
- [ ] Maintainable structure

### Checklists
- [ ] Quality checks defined
- [ ] Validation steps listed
- [ ] Success criteria included
- [ ] Prerequisites documented

## Interaction Design

### Dialogue Patterns
- [ ] Interactive not generative
- [ ] Questions before actions
- [ ] Wait for user responses
- [ ] Build on user input

### Numbered Options
- [ ] Used for all selections
- [ ] Format consistent (1-9, 0)
- [ ] Clear descriptions
- [ ] Logical grouping

### User Experience
- [ ] Clear prompts
- [ ] Helpful guidance
- [ ] Error handling
- [ ] Progress indicators

## Technical Implementation

### Lazy Loading
- [ ] No @ syntax in command file
- [ ] Resources load only when needed
- [ ] Load commands use Read tool
- [ ] File paths correct

### Error Handling
- [ ] File existence checks
- [ ] Input validation
- [ ] Graceful failures
- [ ] Recovery paths

### Performance
- [ ] Minimal resource loading
- [ ] Efficient KB searches
- [ ] Quick response times
- [ ] Scalable approach

## Documentation

### Command File
- [ ] All sections complete
- [ ] YAML properly formatted
- [ ] Resource paths correct
- [ ] Examples included

### User Documentation
- [ ] Purpose clearly stated
- [ ] Commands documented
- [ ] Usage examples provided
- [ ] Common workflows described

### Technical Documentation
- [ ] Architecture documented
- [ ] Dependencies listed
- [ ] Limitations noted
- [ ] Maintenance guide included

## Testing

### Functional Testing
- [ ] All commands tested
- [ ] Resources load correctly
- [ ] KB integration works
- [ ] Outputs generate properly

### User Testing
- [ ] Workflow is intuitive
- [ ] Prompts are clear
- [ ] Options make sense
- [ ] Value is delivered

### Edge Cases
- [ ] Missing files handled
- [ ] Invalid inputs rejected
- [ ] Large datasets managed
- [ ] Concurrent use safe

## Quality Assurance

### Pattern Compliance
- [ ] Follows established patterns
- [ ] Consistent with other agents
- [ ] Best practices applied
- [ ] Anti-patterns avoided

### Code Quality
- [ ] Clean, readable code
- [ ] Proper commenting
- [ ] Logical organization
- [ ] DRY principles

### Maintainability
- [ ] Easy to understand
- [ ] Simple to modify
- [ ] Well documented
- [ ] Version controlled

## Deployment

### Integration
- [ ] CLAUDE.md updated
- [ ] No conflicts with existing agents
- [ ] KB properly configured
- [ ] Dependencies resolved

### Communication
- [ ] Team notified
- [ ] Documentation shared
- [ ] Training provided if needed
- [ ] Feedback channel established

### Monitoring
- [ ] Usage tracking planned
- [ ] Error monitoring setup
- [ ] Performance metrics defined
- [ ] Improvement process established

## Post-Deployment

### Validation
- [ ] Agent works as expected
- [ ] Users understand purpose
- [ ] Value being delivered
- [ ] No negative impacts

### Optimization
- [ ] Performance monitored
- [ ] Feedback collected
- [ ] Improvements identified
- [ ] Updates planned

### Evolution
- [ ] Enhancement path clear
- [ ] Migration strategy defined
- [ ] Deprecation plan ready
- [ ] Documentation maintained

## Final Sign-off

### Critical Checks
- [ ] Single, clear purpose ✓
- [ ] KB integration complete ✓
- [ ] Lazy loading implemented ✓
- [ ] Interactive design confirmed ✓
- [ ] Patterns followed ✓
- [ ] Testing complete ✓
- [ ] Documentation ready ✓
- [ ] Team aligned ✓

### Approval
- [ ] Technical review passed
- [ ] User acceptance confirmed
- [ ] Documentation approved
- [ ] Deployment authorized

## Notes Section
[Space for additional notes, concerns, or special considerations]

---

**Checklist Version**: 1.0
**Last Updated**: [Date]
**Reviewed By**: Rita (Command Recruiter)