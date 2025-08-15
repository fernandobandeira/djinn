# Cognitive Tool Program: Determine Agent Type
# Structured reasoning program to classify whether to create a sub-agent or command

## Program Purpose
This prompt program guides the classification of agent requests into either sub-agents (one-shot) or command agents (interactive).

## Input Schema
```yaml
user_request:
  description: What the user wants the agent to do
  interaction_needs: Does it require back-and-forth?
  output_type: Single deliverable or ongoing dialogue?
  trigger_context: When should it activate?
```

## Classification Logic

### Step 1: Analyze Interaction Pattern
```
IF user_request requires multiple exchanges with user:
  → COMMAND AGENT
ELSE IF user_request is single task with final output:
  → SUB-AGENT
```

### Step 2: Evaluate Output Type
```
IF output is single report/result/action:
  → SUB-AGENT
ELSE IF output is dialogue/guidance/facilitation:
  → COMMAND AGENT
```

### Step 3: Consider Trigger Context
```
IF should activate automatically on condition:
  → SUB-AGENT (with IMPORTANT/Proactively)
ELSE IF user explicitly invokes with /:
  → COMMAND AGENT
```

## Decision Matrix

| Criteria | Sub-Agent | Command Agent |
|----------|-----------|---------------|
| **Interaction** | None after trigger | Multiple exchanges |
| **Output** | Single deliverable | Ongoing dialogue |
| **Invocation** | Automatic (IMPORTANT) | Manual (/command) |
| **State** | Stateless | Maintains context |
| **Examples** | Code reviewer, Test runner | Analyst, Teacher |

## Classification Examples

### Example 1: Code Reviewer
- **Request**: "Review code for security issues"
- **Interaction**: None needed
- **Output**: Security report
- **Decision**: SUB-AGENT
- **Rationale**: Single task, no interaction, delivers report

### Example 2: Brainstorming Assistant
- **Request**: "Help brainstorm feature ideas"
- **Interaction**: Back-and-forth needed
- **Output**: Iterative refinement
- **Decision**: COMMAND AGENT
- **Rationale**: Requires dialogue, multiple exchanges

### Example 3: Test Runner
- **Request**: "Run tests when code changes"
- **Interaction**: None needed
- **Output**: Test results and fixes
- **Decision**: SUB-AGENT
- **Rationale**: Automatic trigger, single output

### Example 4: Architecture Designer
- **Request**: "Guide through system design"
- **Interaction**: Step-by-step guidance
- **Output**: Progressive refinement
- **Decision**: COMMAND AGENT
- **Rationale**: Interactive guidance needed

## Program Output Template

```markdown
## Agent Type Determination

### Analysis:
- Interaction Pattern: {none|dialogue}
- Output Type: {single|ongoing}
- Trigger Method: {automatic|manual}

### Decision: {SUB-AGENT|COMMAND}

### Rationale:
{explanation of why this type was chosen}

### Recommended Configuration:
- Location: {.claude/agents/|.claude/commands/}
- Trigger: {IMPORTANT keyword|/command}
- Tools: {minimal set needed}
- Model: {haiku|sonnet|opus}
```

## Edge Cases

### Hybrid Behaviors
Some agents might seem to fit both categories:
- **Solution**: Choose based on PRIMARY use case
- **Example**: A reviewer that could be automatic OR interactive
  - If mainly automatic → SUB-AGENT
  - If mainly interactive → COMMAND with sub-agent helper

### Complex Workflows
Multi-step processes might seem interactive:
- **Test**: Does each step need user input?
  - Yes → COMMAND AGENT
  - No → SUB-AGENT with complex logic

## Integration with Rita

When Rita uses this program:
1. Parse user's agent request
2. Apply classification logic
3. Guide user to appropriate creation path
4. Store decision pattern in Rita's KB

This program ensures consistent, correct agent type selection based on actual use patterns rather than assumptions.