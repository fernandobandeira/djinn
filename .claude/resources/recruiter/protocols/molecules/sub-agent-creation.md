# Molecular Protocol: Sub-Agent Creation
# Combines atomic constraints into a complete sub-agent creation workflow

## Protocol Overview
This molecular protocol combines multiple atomic constraints to ensure successful sub-agent creation.

## Atomic Constraints Combined
1. `frontmatter-syntax` - Ensures valid YAML structure
2. `tool-selection` - Optimizes tool choice
3. `delegation-triggers` - Creates effective triggers
4. `file-location` - Places in correct directory

## Protocol Steps

### Step 1: Gather Requirements
```yaml
input_schema:
  agent_name: string
  purpose: string
  trigger_conditions: array
  required_capabilities: array
  model_preference: haiku|sonnet|opus
```

### Step 2: Apply Constraint Stack
1. **Location Constraint**: Determine `.claude/agents/` path
2. **Syntax Constraint**: Generate valid frontmatter
3. **Tool Constraint**: Select minimal tool set
4. **Trigger Constraint**: Craft delegation description

### Step 3: Frontmatter Generation Template
```yaml
---
name: {agent_name}
description: {trigger_keyword} {action_verb} {context}
tools: {comma_separated_tools}
model: {model_choice}
---
```

### Step 4: Constraint Validation Checklist
- [ ] Name is unique (check existing agents)
- [ ] Description contains trigger keywords
- [ ] Tools are minimal for purpose
- [ ] Model is appropriate for complexity
- [ ] File location is correct

### Step 5: System Prompt Structure
```markdown
You are a {role} specialist.

## Core Purpose
{primary_purpose}

## Constraints
- {constraint_1}
- {constraint_2}

## Expected Output
{output_format}
```

## Example Application

### Input:
- Name: "code-security-reviewer"
- Purpose: "Review code for security vulnerabilities"
- Triggers: ["PR creation", "security review request"]
- Capabilities: ["read code", "analyze patterns"]

### Output:
```yaml
---
name: code-security-reviewer
description: IMPORTANT reviews code for security vulnerabilities during PR creation
tools: Read, Grep, Glob
model: sonnet
---

You are a security review specialist.

## Core Purpose
Analyze code changes for security vulnerabilities, focusing on:
- Input validation issues
- Authentication/authorization flaws
- Injection vulnerabilities
- Sensitive data exposure

## Constraints
- Only analyze, never modify code
- Focus on security implications
- Provide actionable recommendations

## Expected Output
Structured security report with:
- Vulnerability severity levels
- Specific line references
- Remediation suggestions
```

## Constraint Diagnostic

### Well-Constrained Indicators:
- Clear trigger conditions
- Minimal but sufficient tools
- Specific, actionable purpose
- Appropriate model for task complexity

### Under-Constrained Warning Signs:
- Vague descriptions
- Missing trigger keywords
- Unclear purpose
- No output format specified

### Over-Constrained Warning Signs:
- Too many specific conditions
- Overly restrictive tool set
- Inflexible output requirements
- Complex rules for simple tasks

## Protocol Validation
After applying this molecular protocol, validate:
1. All atomic constraints satisfied
2. Agent can be triggered appropriately
3. Tools match intended capabilities
4. System prompt is clear and focused
5. File created in correct location