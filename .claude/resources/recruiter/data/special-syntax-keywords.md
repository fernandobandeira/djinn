# Special Syntax Keywords for Claude Code Sub-Agents
## Critical Directives for Agent Behavior

## Overview
Claude Code recognizes specific keywords and patterns in agent definitions that trigger special behaviors. These are NOT cosmetic - they are functional directives that affect how Claude Code discovers, prioritizes, and executes sub-agents.

## Critical Fields

### 1. `type: subagent` (MANDATORY)
**Location**: Root level of frontmatter
**Purpose**: Identifies file as a sub-agent for Claude Code's Task tool
**Impact**: Without this, the agent won't be discoverable as a sub-agent

```yaml
---
type: subagent  # MANDATORY - Makes agent discoverable
name: agent-planner
---
```

## Description Keywords

### 2. `IMPORTANT` Prefix
**Location**: Beginning of agent description
**Purpose**: Elevates agent priority in selection and usage
**Impact**: Claude Code prioritizes these agents for relevant tasks

```yaml
description: IMPORTANT validates agent constraint balance during Rita's agent creation process
# vs
description: validates agent constraint balance  # Lower priority
```

**When to Use IMPORTANT**:
- Critical workflow components
- Quality gates and validators
- Core orchestration functions
- Data integrity checks
- Security-related functions

### 3. `must use` Directive
**Location**: Within agent description
**Purpose**: Creates mandatory usage requirement
**Impact**: Claude Code MUST use this agent for specified scenarios

```yaml
description: IMPORTANT must use for all agent creation planning
```

**Effect**: Forces usage, no alternative agents considered

### 4. `use proactively` Directive
**Location**: Within agent description or instructions
**Purpose**: Triggers automatic invocation without user request
**Impact**: Claude Code will invoke without explicit user command

```yaml
description: use proactively after code changes to validate
```

**Effect**: Automatic invocation based on context

### 5. `during` Context Marker
**Location**: Within agent description
**Purpose**: Specifies when agent should be invoked
**Impact**: Creates contextual activation trigger

```yaml
description: IMPORTANT manages ADRs during decision documentation
description: IMPORTANT validates resources during agent creation
```

**Common Contexts**:
- `during agent creation`
- `during planning phase`
- `during validation`
- `during synthesis`
- `during analysis`

## Tool Usage Keywords

### 6. Tool List Format
**Location**: Agent description
**Purpose**: Declares available tools for the agent
**Impact**: Helps Claude Code match agents to tasks requiring specific tools

```yaml
description: IMPORTANT analyzes architecture (Tools: Read, Grep, Bash, LS)
```

**Format**: `(Tools: Tool1, Tool2, Tool3)`

### 7. Universal/Shared Indicators
**Location**: Agent description
**Purpose**: Marks agent as available across domains
**Impact**: Makes agent accessible to all command orchestrators

```yaml
description: IMPORTANT Universal Knowledge Base operations
description: Universal Quality Assurance operations
```

## Behavioral Modifiers

### 8. Processing Patterns
These keywords in descriptions affect how the agent processes:

- `comprehensive` - Thorough, complete analysis
- `systematic` - Step-by-step, ordered approach
- `advanced` - Uses sophisticated techniques
- `specialized` - Domain-specific expertise
- `universal` - Works across all domains

### 9. Output Patterns
These affect expected output format:

- `generates` - Creates new content
- `validates` - Returns validation results
- `analyzes` - Provides analysis report
- `extracts` - Pulls out specific information
- `synthesizes` - Combines multiple sources

## Complete Example

```yaml
---
type: subagent  # MANDATORY field
name: architecture-analyst
description: IMPORTANT performs deep architectural analysis of agents and suggests improvements based on complexity, patterns, and best practices (Tools: Read, Grep, Bash, LS)
model:
  provider: anthropic
  id: claude-opus-4-1-20250805
---

# Architecture Analyst Sub-Agent

## Activation
This agent is activated when architectural analysis is needed, especially during agent creation and review workflows.

### Proactive Usage
use proactively when:
- Agent complexity exceeds thresholds
- Validation failures occur
- Performance issues detected

### Mandatory Usage
must use for:
- Pre-deployment architecture review
- Complexity assessment
- Pattern identification
```

## Discovery Patterns

Claude Code discovers agents through:

1. **Type Matching**: `type: subagent` in frontmatter
2. **Keyword Scanning**: Searches descriptions for task keywords
3. **Tool Matching**: Matches required tools to agent capabilities
4. **Context Matching**: `during` contexts align with current workflow
5. **Priority Sorting**: `IMPORTANT` agents rank higher

## Best Practices

### Do's
- ✅ Always include `type: subagent` in frontmatter
- ✅ Use `IMPORTANT` for critical workflow components
- ✅ Specify tools in parentheses format
- ✅ Include `during` context for workflow placement
- ✅ Use `must use` for mandatory steps
- ✅ Add `use proactively` for automatic triggers

### Don'ts
- ❌ Don't use IMPORTANT for every agent
- ❌ Don't forget the `type: subagent` field
- ❌ Don't mix tool formats
- ❌ Don't use conflicting directives
- ❌ Don't create circular dependencies

## Validation Checklist

For agent-builder and validators to check:

```yaml
validation:
  mandatory:
    - type: subagent  # Must exist
    - name: string    # Must have name
    - description: string  # Must have description
  
  recommended:
    - IMPORTANT prefix (for critical agents)
    - Tools list in (Tools: X, Y, Z) format
    - during context marker
    - Clear action verb (validates, generates, etc.)
  
  optional:
    - must use directive
    - use proactively directive
    - Universal marker
```

## Impact on Agent Selection

### Priority Hierarchy
1. Agents with `must use` for current context
2. Agents with `IMPORTANT` + matching context
3. Agents with `IMPORTANT` + matching tools
4. Agents with matching context
5. Agents with matching tools
6. Other available agents

### Automatic Invocation
Agents trigger automatically when:
- `use proactively` + trigger condition met
- `must use` + context matches
- Workflow requirement + agent available

## Integration with Thinking Mode

For critical agents, combine with thinking triggers:
```yaml
description: IMPORTANT must use for critical planning (trigger: think)
```

This signals orchestrators to invoke with thinking mode for better accuracy.

## Remember
These keywords are **functional directives**, not documentation. They directly affect:
- Agent discovery
- Selection priority  
- Automatic invocation
- Workflow integration
- Quality gates

Always validate keyword usage during agent creation to ensure proper system integration.