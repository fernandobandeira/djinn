# Agent Creation Patterns Knowledge Base

This directory contains Rita's accumulated knowledge about agent creation patterns, stored in `/docs` to enable vector DB indexing.

## Purpose
Rita uses this knowledge base to:
- Learn from successful agent creations
- Avoid known failure patterns
- Apply constraint architecture principles
- Continuously improve agent quality

## Collections

### successful-agents/
Patterns from well-functioning agents with balanced constraints

### syntax-examples/
Valid Claude Code frontmatter and syntax patterns

### delegation-patterns/
Effective trigger keywords and delegation strategies

### tool-combinations/
Minimal tool sets for different agent purposes

### failure-patterns/
Anti-patterns and common mistakes to avoid

### constraint-examples/
Examples of under/over/well-constrained agents

## Vector DB Integration

This knowledge is indexed in the main documentation collection:

```bash
# Index Rita's patterns
./.vector_db/kb index --path ./docs/agent-patterns/

# Search for patterns
./.vector_db/kb search "agent creation patterns"
```

## Constraint Architecture Application

Rita applies the atoms→molecules→cells→organs pattern:
- **Atoms**: Individual syntax rules and constraints
- **Molecules**: Combined protocols for agent creation
- **Cells**: This accumulated pattern knowledge
- **Organs**: Sub-agents for validation and learning

## Learning Loop

1. Create agent
2. Validate constraints
3. Extract successful patterns
4. Document in this KB
5. Index in vector DB
6. Apply to next creation

This creates continuous improvement in agent quality.