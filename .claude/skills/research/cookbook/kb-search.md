# KB Search Cookbook

Semantic search across project notes using Basic Memory MCP.

## When to Use

- Before creating any new document
- When exploring what exists on a topic
- When you need prior decisions or patterns
- To find related work before brainstorming

## Basic Search

```
mcp__basic-memory__search_notes(query="your query here")
```

## Search Examples

### Before Creating an ADR
```
mcp__basic-memory__search_notes(query="authentication strategy decision")
```

### Finding Patterns
```
mcp__basic-memory__search_notes(query="error handling patterns")
```

### Research Context
```
mcp__basic-memory__search_notes(query="market analysis competitive landscape")
```

## Understanding Results

Search returns notes with:
- **Permalink** - Unique identifier to read the full note
- **Title** - Note title
- **Preview** - First part of matching content
- **Relations** - Connected notes via [[wikilinks]]

## After Finding Results

1. **Read the full note** - Use `read_note` to get complete content
2. **Follow relations** - Check [[linked]] notes for context
3. **Note the folder** - Indicates document type (decisions/, patterns/, etc.)
4. **Cross-reference** - Search related terms for completeness

## Reading a Note

Once you find a relevant note, read it fully:

```
mcp__basic-memory__read_note(permalink="auth-strategy")
```

## Query Construction Tips

### Be Semantic, Not Keyword-Based
```
# Good - semantic meaning
mcp__basic-memory__search_notes(query="handling errors gracefully in API responses")

# Less effective - just keywords
mcp__basic-memory__search_notes(query="error")
```

### Use Domain Language
```
# For architecture
mcp__basic-memory__search_notes(query="microservices communication patterns")

# For analysis
mcp__basic-memory__search_notes(query="competitive landscape assessment")
```

### Combine Concepts
```
mcp__basic-memory__search_notes(query="authentication security best practices JWT")
```

## Common Search Patterns

### Before Creating an ADR
```
mcp__basic-memory__search_notes(query="[decision topic] decision architecture")
```

### Before Writing Code
```
mcp__basic-memory__search_notes(query="[functionality] implementation pattern")
```

### Before Research/Analysis
```
mcp__basic-memory__search_notes(query="[topic] analysis research findings")
```

### Finding Prior Art
```
mcp__basic-memory__search_notes(query="similar to [concept] pattern example")
```

## Browsing Recent Notes

See what's been added recently:

```
mcp__basic-memory__recent_activity()
```

## Getting Project Overview

Get a canvas view of the project knowledge:

```
mcp__basic-memory__canvas()
```

## Writing New Notes

After research, document findings:

```
mcp__basic-memory__write_note(
    title="Research: User Authentication Options",
    content="""
## Summary
Evaluated three authentication approaches for the API.

## Options Considered
1. JWT with refresh tokens - See [[jwt-research]]
2. Session-based auth - See [[session-auth-research]]
3. OAuth2 - See [[oauth-research]]

## Recommendation
JWT with refresh tokens based on [[security-requirements]] and [[api-design]] constraints.

## Relations
- [[project]] - main project context
- [[auth-strategy]] - final decision
""",
    folder="research"
)
```

## Note Format

Basic Memory notes use this structure:

```markdown
---
title: Note Title
type: note
permalink: note-title
---

## Observations
- Key observation 1
- Key observation 2
- See [[related-note]] for more context

## Relations
- [[project]] - connection to project
- [[other-note]] - related concept
```

## Troubleshooting

**No results?**
1. Try broader query terms
2. Use different phrasing (semantic search)
3. Check if content has been created yet
4. Try `recent_activity()` to see what exists

**Too many irrelevant results?**
1. Be more specific in query
2. Add domain context words
3. Combine multiple concepts

## Integration with Other Techniques

After search:
- **Knowledge Harvesting** - If internal search yields nothing, harvest externally
- **Source Evaluation** - Assess quality of what you found
