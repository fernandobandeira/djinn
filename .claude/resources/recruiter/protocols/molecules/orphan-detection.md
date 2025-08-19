# Orphan Detection Protocol for Shared Agents

## Purpose
Prevent shared sub-agents from becoming orphaned (unreferenced by any command agent) and ensure proper integration tracking across the agent ecosystem.

## Validation Rules

### 1. Orphan Detection Rule
**Trigger**: On creation or modification of any shared agent
**Check**: Verify at least one command agent references the shared agent
**Implementation**:
```python
def check_orphaned_agents():
    shared_agents = glob(".claude/agents/shared/*.md")
    command_agents = glob(".claude/commands/*.md")
    
    for shared_agent in shared_agents:
        agent_name = extract_agent_name(shared_agent)
        referenced = False
        
        for command in command_agents:
            if contains_reference(command, agent_name):
                referenced = True
                break
        
        if not referenced:
            return {
                "status": "fail",
                "orphaned": agent_name,
                "recommendation": f"Add {agent_name} to relevant command agents"
            }
    
    return {"status": "pass"}
```

### 2. Integration Documentation Rule
**Trigger**: On shared agent creation
**Check**: Verify integration protocols section exists
**Sections Required**:
- When to use this agent
- Example Task() invocations
- Expected return format
- Agent context requirements

### 3. Usage Tracking Rule
**Trigger**: Weekly or on-demand
**Check**: Track actual Task() invocations
**Logging Format**:
```yaml
timestamp: ISO-8601
invoker: command_agent_name
target: shared_agent_name
context: agent_context_provided
success: boolean
```

### 4. Discovery Protocol
**Trigger**: On new shared agent creation
**Actions**:
1. Update CLAUDE.md shared agents section
2. Identify relevant command agents
3. Generate integration suggestions
4. Create integration tasks

### 5. Reference Validation Rule
**Trigger**: On command agent modification
**Check**: Verify Task() calls use valid shared agents
**Implementation**:
```python
def validate_task_references(command_file):
    task_calls = extract_task_calls(command_file)
    valid_agents = get_shared_agent_names()
    
    for task in task_calls:
        if task.subagent_type not in valid_agents:
            return {
                "status": "fail", 
                "invalid_reference": task.subagent_type,
                "in_file": command_file
            }
    
    return {"status": "pass"}
```

## Integration Checklist

When adding a new shared agent:

- [ ] Agent file created in `.claude/agents/shared/`
- [ ] Added to CLAUDE.md shared agents section
- [ ] Integration examples documented
- [ ] At least 3 command agents updated with references
- [ ] Task() delegation examples provided
- [ ] Usage tracking enabled
- [ ] Validation tests pass

## Enforcement Levels

### Critical (Block Deployment)
- Orphaned agent with no references
- Missing integration documentation
- Invalid Task() references

### Warning (Allow with Notice)
- Low usage statistics (< 5 invocations/week)
- Missing from some relevant commands
- Incomplete integration examples

### Info (Track Only)
- Usage patterns and statistics
- Integration opportunities
- Performance metrics

## Validation Commands

### Check for Orphans
```bash
rita *validate-orphans
```

### Generate Integration Report
```bash
rita *integration-report
```

### Track Usage Statistics
```bash
rita *usage-stats [agent-name]
```

## Success Metrics

- **Zero orphaned agents** in production
- **100% documentation coverage** for shared agents
- **>80% command coverage** for universally useful agents
- **Weekly usage** for all shared agents
- **<24hr integration time** for new shared agents

## Notes

### Future Enhancements
1. **Auto-integration**: Automatically suggest and apply integrations
2. **Usage prediction**: ML-based prediction of which commands need an agent
3. **Performance tracking**: Monitor delegation success rates
4. **Deprecation workflow**: Safely remove unused agents

### Integration Patterns
1. **Universal agents** (kb-analyst, qa-reviewer): All commands
2. **Domain agents** (knowledge-harvester): Most commands
3. **Specialized agents** (plan-validator): Specific commands
4. **Helper agents**: Task-specific sub-agents

This protocol ensures no shared agent becomes orphaned and maintains a healthy, well-integrated agent ecosystem.