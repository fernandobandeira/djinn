# Cognitive Tool Program: Fetch Documentation
# Structured program to retrieve and parse Claude Code documentation

## Program Purpose
Retrieves the latest Claude Code documentation and extracts relevant patterns for agent creation.

## Documentation Sources

### Primary Documentation
```yaml
base_url: https://docs.anthropic.com/en/docs/claude-code
pages:
  sub_agents: /sub-agents
  commands: /commands  
  mcp: /mcp
  syntax: /syntax
  best_practices: /best-practices
```

### Retrieval Strategy

#### For Sub-Agents
1. Fetch sub-agents documentation
2. Extract frontmatter syntax rules
3. Identify keyword patterns (IMPORTANT, Proactively)
4. Parse tool specification format
5. Note model selection guidelines

#### For Commands
1. Fetch commands documentation
2. Extract command structure patterns
3. Identify interaction protocols
4. Parse resource organization
5. Note dialogue patterns

## Extraction Patterns

### Frontmatter Rules
```yaml
pattern_extraction:
  - YAML delimiter requirements
  - Required vs optional fields
  - Tool list format (comma-separated)
  - Model options and defaults
  - Description keyword importance
```

### Delegation Triggers
```yaml
keyword_patterns:
  strong:
    - "IMPORTANT"
    - "MUST BE USED"
  proactive:
    - "Proactively"
    - "Use proactively"
  conditional:
    - "when"
    - "after"
    - "during"
```

## Processing Logic

### Step 1: Determine Documentation Needs
```
IF creating_sub_agent:
  FETCH sub-agents page
  EXTRACT frontmatter syntax
  EXTRACT delegation patterns
ELSE IF creating_command:
  FETCH commands page
  EXTRACT structure patterns
  EXTRACT interaction patterns
```

### Step 2: Parse Syntax Rules
```
FOR each documentation_page:
  IDENTIFY syntax requirements
  EXTRACT example patterns
  NOTE special directives
  CAPTURE best practices
```

### Step 3: Store Patterns
```
patterns_found:
  - Store in working memory
  - Apply to current creation
  - Document in Rita's KB if novel
```

## Output Schema
```yaml
documentation_extracted:
  syntax_rules:
    frontmatter: {}
    tools: {}
    models: {}
  delegation_patterns:
    keywords: []
    contexts: []
  best_practices: []
  examples: []
```

## Integration with Rita
1. Use before any agent creation
2. Validate against latest syntax
3. Update Rita's KB with new patterns
4. Flag deprecated patterns

## Error Handling
- If docs unavailable: Use cached patterns from Rita's KB
- If syntax changed: Alert user to updates
- If examples found: Store as successful patterns

This cognitive tool ensures Rita always uses the latest Claude Code conventions and syntax.