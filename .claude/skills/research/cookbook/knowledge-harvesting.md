# Knowledge Harvesting Cookbook

Extract knowledge from external sources and store it in Basic Memory.

## When to Use

- Internal KB search yields insufficient results
- Need external documentation or examples
- Researching new technologies or patterns
- Gathering competitive intelligence
- Finding troubleshooting solutions

## Harvesting Workflow

1. **Fetch content** using WebFetch
2. **Extract key information** from the response
3. **Store as note** in Basic Memory with proper links

## Basic Harvest

```
# Step 1: Fetch the content
WebFetch(
    url="https://docs.example.com/authentication",
    prompt="Extract the key concepts, patterns, and recommendations for authentication"
)

# Step 2: Store as a note
mcp__basic-memory__write_note(
    title="External: Authentication Patterns from Example.com",
    content="""
## Source
URL: https://docs.example.com/authentication
Harvested: [date]

## Key Concepts
[extracted concepts from WebFetch]

## Patterns
[extracted patterns]

## Recommendations
[extracted recommendations]

## Relations
- [[authentication]] - internal auth work
- [[security-requirements]] - our requirements
""",
    folder="research"
)
```

## Harvesting Patterns

### Technology Assessment
```
# Fetch
WebFetch(
    url="https://newframework.io/docs/overview",
    prompt="Summarize the key features, architecture, and use cases for this framework"
)

# Store
mcp__basic-memory__write_note(
    title="Research: NewFramework Evaluation",
    content="[extracted content with [[links]] to related notes]",
    folder="research"
)
```

### API Documentation
```
# Fetch
WebFetch(
    url="https://api.service.com/docs",
    prompt="Extract API endpoints, authentication methods, and rate limits"
)

# Store
mcp__basic-memory__write_note(
    title="External: Service API Reference",
    content="[extracted API info with [[links]] to our integration notes]",
    folder="research"
)
```

### Pattern Research
```
# Fetch
WebFetch(
    url="https://patterns-site.com/microservices",
    prompt="Extract microservices patterns, when to use each, and trade-offs"
)

# Store
mcp__basic-memory__write_note(
    title="Research: Microservices Patterns",
    content="[extracted patterns with [[links]] to our architecture notes]",
    folder="patterns"
)
```

### Troubleshooting
```
# Search for solutions
WebSearch(query="React hooks memory leak fix")

# Fetch relevant results
WebFetch(
    url="[best result URL]",
    prompt="Extract the problem description, root cause, and solution"
)

# Store solution
mcp__basic-memory__write_note(
    title="Solution: React Hooks Memory Leak",
    content="[problem and solution with [[links]] to related issues]",
    folder="research"
)
```

## Note Structure for Harvested Content

```markdown
## Source
- URL: [original URL]
- Harvested: [date]
- Type: [documentation/article/forum/tutorial]

## Summary
[Brief overview of what was harvested]

## Key Content
[Extracted information organized by topic]

## Applicability
- Relevant to: [[related-project-note]]
- Constraints: [any limitations or context differences]
- Confidence: [high/medium/low]

## Relations
- [[project-note-1]] - how this relates
- [[project-note-2]] - another connection
```

## Best Practices

1. **Search KB first** - Don't harvest if knowledge exists internally
2. **Use focused prompts** - Tell WebFetch exactly what to extract
3. **Add context** - Note how harvested content relates to your project
4. **Link generously** - Connect to existing notes with [[wikilinks]]
5. **Note limitations** - External content may not fit your context exactly
6. **Date everything** - Technical content ages quickly

## Multi-Source Research

For comprehensive research, harvest from multiple sources:

```
# Fetch from multiple sources
source1 = WebFetch(url="[official docs]", prompt="[extraction prompt]")
source2 = WebFetch(url="[tutorial]", prompt="[extraction prompt]")
source3 = WebFetch(url="[comparison article]", prompt="[extraction prompt]")

# Synthesize into one note
mcp__basic-memory__write_note(
    title="Research: [Topic] Comprehensive Review",
    content="""
## Overview
Synthesized from 3 sources on [topic].

## From Official Docs
[key points]

## From Tutorial
[practical examples]

## From Comparison
[trade-offs and alternatives]

## Synthesis
[your conclusions combining all sources]

## Relations
- [[project]] - project context
- [[decision-to-make]] - what this informs
""",
    folder="research"
)
```

## Error Handling

- **URL blocked**: Try alternative sources or use WebSearch first
- **Content too long**: Focus prompt on specific sections
- **No relevant content**: Refine search or try different URLs
- **Conflicting info**: Note the conflict in your research note

## Integration with Other Skills

After harvesting:
- **Source Evaluation** - Assess quality and applicability
- **Strategic Analysis** - If evaluating options
- **Devils Advocate** - Challenge the harvested recommendations
