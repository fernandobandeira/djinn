# Thinking Mode Triggers for Critical Tasks
## Strategic Use of Extended Reasoning in Claude 4.1

## Overview
Claude 4.1 Opus and Sonnet support extended thinking modes that dramatically improve accuracy for complex tasks at the cost of increased token usage. This guide defines when and how to trigger these modes for critical sub-agent operations.

## Available Thinking Modes

### 1. Standard Thinking (`think`)
**Trigger Keywords**: `think`, `reason`, `analyze`
**Token Impact**: ~2-3x normal usage
**Best For**: 
- Single-step complex decisions
- Quick architectural reviews
- Validation logic

### 2. Deep Thinking (`deepthink`)
**Trigger Keywords**: `deepthink`, `deep analysis`, `thorough reasoning`
**Token Impact**: ~5-10x normal usage
**Best For**:
- Multi-step planning
- Complex refactoring
- System design

### 3. Ultra Thinking (`ultrathink`)
**Trigger Keywords**: `ultrathink`, `maximum reasoning`, `critical analysis`
**Token Impact**: ~10-20x normal usage
**Best For**:
- Mission-critical decisions
- Complete system architecture
- Complex problem decomposition

## Implementation in Orchestrators

### How to Trigger Thinking Mode

Orchestrators (like Rita) should modify their Task tool prompts:

```python
# Standard invocation
prompt = "Analyze the requirements for new agent"

# With thinking mode trigger
prompt = "think carefully: Analyze the requirements for new agent"
prompt = "deepthink: Plan the complete architecture for distributed system"
prompt = "ultrathink: Decompose this monolithic application into microservices"
```

### In Agent Definitions

```yaml
---
type: subagent
name: agent-planner
description: IMPORTANT plans agent architecture with systematic decomposition (Supports: think, deepthink)
model:
  provider: anthropic
  id: claude-opus-4-1-20250805
  thinking_mode: adaptive  # Responds to trigger keywords
---
```

## Critical Scenarios Requiring Thinking Mode

### Always Use Thinking Mode For:

#### 1. Agent Planning (`agent-planner`)
```python
# Rita's invocation
prompt = "deepthink: Plan architecture for agent '{name}' with requirements: {requirements}"
```
**Why**: Architecture decisions have cascading effects

#### 2. Architecture Analysis (`architecture-analyst`)
```python
prompt = "think: Analyze architecture complexity and identify decomposition opportunities"
```
**Why**: Needs to balance multiple architectural concerns

#### 3. System Design (`system-designer`)
```python
prompt = "ultrathink: Design complete system with multiple architecture options"
```
**Why**: Requires exploring multiple solution paths

#### 4. Change Coordination (`change-coordinator`)
```python
prompt = "deepthink: Plan migration from current to target architecture"
```
**Why**: Must consider all dependencies and risks

#### 5. Pattern Extraction (`pattern-extractor`)
```python
prompt = "think: Extract reusable patterns from successful implementations"
```
**Why**: Pattern recognition requires deep analysis

### Conditionally Use Thinking Mode For:

#### 1. Validation Tasks
- Use standard mode for simple checks
- Use `think` for complex constraint validation
- Example: `constraint-validator` with score < 7.0

#### 2. Synthesis Tasks
- Use `think` for small datasets (< 5 sources)
- Use `deepthink` for large datasets (> 10 sources)
- Example: `zettelkasten-synthesizer` with many notes

#### 3. Research Tasks
- Use standard for known domains
- Use `think` for novel areas
- Use `deepthink` for competitive analysis

## Orchestrator Implementation Guide

### For Rita (Recruiter)

```python
class RitaOrchestrator:
    def invoke_critical_subagent(self, agent_type, task, criticality):
        """Invoke sub-agent with appropriate thinking mode"""
        
        thinking_triggers = {
            'low': '',
            'medium': 'think: ',
            'high': 'deepthink: ',
            'critical': 'ultrathink: '
        }
        
        criticality_map = {
            'agent-planner': 'high',
            'architecture-analyst': 'high',
            'system-designer': 'critical',
            'agent-builder': 'medium',
            'constraint-validator': 'medium',
            'resource-validator': 'low',
            'pattern-extractor': 'high'
        }
        
        trigger = thinking_triggers[criticality_map.get(agent_type, 'low')]
        
        return Task(
            subagent_type=agent_type,
            description=f"Critical task with {criticality} priority",
            prompt=f"{trigger}{task}"
        )
```

### For Command Orchestrators

Add to orchestrator instructions:
```markdown
## Thinking Mode Usage

When invoking critical sub-agents, prepend thinking triggers:

1. **Planning Tasks**: Use "deepthink: " prefix
2. **Analysis Tasks**: Use "think: " prefix  
3. **Validation Tasks**: Use standard (no prefix)
4. **Emergency Fixes**: Use "ultrathink: " prefix

Example:
- Normal: "Validate the agent configuration"
- Critical: "deepthink: Plan the agent decomposition strategy"
```

## Cost-Benefit Analysis

### When Thinking Mode is Worth It

| Scenario | Token Cost | Benefit | Use Thinking? |
|----------|------------|---------|---------------|
| Initial architecture planning | High | Prevents rework | ✅ Yes (deepthink) |
| Simple validation | Low | Minimal | ❌ No |
| Complex refactoring | Very High | Ensures correctness | ✅ Yes (think) |
| Pattern extraction | Medium | Improves future agents | ✅ Yes (think) |
| Resource checking | Low | Binary result | ❌ No |
| System design | Very High | Multiple options | ✅ Yes (ultrathink) |
| Bug diagnosis | High | Finds root cause | ✅ Yes (think) |

### Token Usage Guidelines

```yaml
token_budgets:
  standard_task: 1000-2000
  think_task: 3000-6000
  deepthink_task: 10000-20000
  ultrathink_task: 20000-50000
  
recommended_limits:
  per_session_think_calls: 5
  per_session_deepthink_calls: 2
  per_session_ultrathink_calls: 1
```

## Monitoring and Optimization

### Track Thinking Mode Effectiveness

```yaml
metrics_to_track:
  - task_success_rate:
      without_thinking: baseline
      with_think: improvement_percentage
      with_deepthink: improvement_percentage
  
  - token_efficiency:
      tokens_per_success: ratio
      cost_per_outcome: dollar_value
  
  - time_savings:
      rework_avoided: hours
      errors_prevented: count
```

### Optimization Rules

1. **Start Conservative**: Begin with standard mode
2. **Escalate on Failure**: If task fails, retry with thinking
3. **Learn Patterns**: Track which tasks benefit most
4. **Set Budgets**: Limit thinking mode calls per session
5. **Review Weekly**: Adjust triggers based on metrics

## Implementation Checklist

### For Orchestrators
- [ ] Add thinking trigger logic to Task invocations
- [ ] Map sub-agents to default criticality levels
- [ ] Implement token budget tracking
- [ ] Add user notification for thinking mode usage
- [ ] Create override mechanism for urgent tasks

### For Sub-Agents
- [ ] Document thinking mode support in description
- [ ] Add `thinking_mode: adaptive` to model config
- [ ] Include trigger examples in instructions
- [ ] Define token budget expectations
- [ ] Specify when thinking is beneficial

### For Knowledge Base
- [ ] Index thinking mode patterns
- [ ] Track successful thinking mode usage
- [ ] Build corpus of effective triggers
- [ ] Document cost-benefit examples
- [ ] Create optimization recommendations

## Example Integration

### Rita Invoking Agent-Planner with Thinking

```python
# User request: Create complex distributed agent system

# Rita's internal logic
def plan_complex_agent(name, requirements):
    # Determine complexity
    complexity = assess_complexity(requirements)
    
    # Select thinking mode
    if complexity > 8:
        trigger = "ultrathink: "
        notify_user("Using extended reasoning for complex architecture...")
    elif complexity > 6:
        trigger = "deepthink: "
        notify_user("Engaging deep analysis mode...")
    else:
        trigger = "think: "
    
    # Invoke with thinking
    prompt = f"""{trigger}Plan architecture for '{name}':
    
    Requirements:
    {requirements}
    
    Consider:
    - Decomposition into sub-agents
    - Optimal model selection (Opus/Sonnet/Haiku)
    - Tool requirements
    - Integration patterns
    - Constraint balance (8.0-8.5)
    
    Provide comprehensive architecture plan.
    """
    
    result = Task(
        subagent_type="agent-planner",
        description="Complex architecture planning",
        prompt=prompt
    )
    
    return result
```

## Remember

- Thinking modes are **powerful but expensive**
- Use strategically for **critical decisions**
- **Track effectiveness** to optimize usage
- **Start simple**, escalate as needed
- **Document patterns** for future reference
- Balance **accuracy vs cost** based on task importance