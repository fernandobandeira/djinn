# Sub-Agent Creation Checklist

## Pre-Creation
- [ ] Confirmed user wants sub-agent (not command)
- [ ] Checked KB for similar existing agents
- [ ] Fetched latest Claude Code sub-agent documentation
- [ ] Validated use case fits one-shot pattern

## Agent Definition
- [ ] Name is descriptive and kebab-case
- [ ] Description starts with delegation trigger
- [ ] Description clearly indicates when to use
- [ ] Single, focused responsibility defined

## Tool Selection
- [ ] Minimal tool set selected
- [ ] No unnecessary tools included
- [ ] Tools match agent's purpose
- [ ] Considered read-only vs modification needs

## Model Selection
- [ ] Appropriate model chosen for task complexity
- [ ] haiku for simple/fast tasks
- [ ] sonnet for balanced tasks
- [ ] opus only when necessary

## Frontmatter Configuration
- [ ] Valid YAML syntax
- [ ] All required fields present (name, description, tools, model)
- [ ] Color selected for visual distinction
- [ ] No syntax errors

## System Prompt Quality
- [ ] Clear purpose statement
- [ ] IMPORTANT directive included
- [ ] Step-by-step instructions
- [ ] Best practices documented
- [ ] Output format specified
- [ ] Constraints defined

## File Creation
- [ ] Created in `.claude/agents/` directory
- [ ] File name matches agent name
- [ ] `.md` extension used
- [ ] File permissions correct

## Testing
- [ ] Test scenario created
- [ ] Delegation trigger tested
- [ ] Tools accessible
- [ ] Output format correct
- [ ] Error handling works

## Documentation
- [ ] Purpose documented clearly
- [ ] Use cases provided
- [ ] Example input/output included
- [ ] Added to knowledge base

## Post-Creation
- [ ] Indexed in KB
- [ ] Tested with real scenario
- [ ] Delegation works automatically
- [ ] Output meets expectations

## Common Issues to Check
- [ ] Not trying to interact (one-shot only)
- [ ] Not loading unnecessary context
- [ ] Not using more tools than needed
- [ ] Clear when to delegate
- [ ] Distinct from existing agents

## Quality Metrics
- [ ] Single responsibility
- [ ] Clear delegation trigger
- [ ] Minimal tool usage
- [ ] Appropriate model
- [ ] Well-structured output
- [ ] No user interaction
- [ ] Completes autonomously

## Red Flags
- [ ] Trying to do too much
- [ ] Unclear when to use
- [ ] Too many tools
- [ ] Requires interaction
- [ ] Overlaps with existing agent
- [ ] No clear output format
- [ ] Missing IMPORTANT directive

## Final Validation
- [ ] Would Claude know when to delegate?
- [ ] Is the output format clear?
- [ ] Are tools minimal but sufficient?
- [ ] Does it complete in one shot?
- [ ] Is it better than inline execution?