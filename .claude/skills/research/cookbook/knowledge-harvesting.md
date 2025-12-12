# Knowledge Harvesting Cookbook

Extract and index knowledge from external sources using crawl4ai.

## When to Use

- Internal KB search yields insufficient results
- Need external documentation or examples
- Researching new technologies or patterns
- Gathering competitive intelligence
- Finding troubleshooting solutions

## Basic Harvest

```bash
./.vector_db/harvest \
    --url "https://example.com/docs" \
    --topic "Topic Name" \
    --profile "deep_research" \
    --agent "architect"
```

## Harvesting Profiles

| Profile | Depth | Timeout | Best For |
|---------|-------|---------|----------|
| `quick_reference` | 1 page | 30s | Single doc page, quick lookups |
| `deep_research` | 5+ pages | 120s | Comprehensive documentation |
| `code_examples` | 3 pages | 60s | Implementation examples |
| `api_documentation` | 4 pages | 90s | API endpoints, schemas |
| `troubleshooting` | 20 results | 60s | Error solutions, forum posts |

## Profile Selection

### Quick Reference
Single page extraction for immediate answers.
```bash
./.vector_db/harvest --url "URL" --topic "API Reference" --profile "quick_reference" --agent "developer"
```

### Deep Research
Multi-page comprehensive crawling.
```bash
./.vector_db/harvest --url "URL" --topic "Framework Architecture" --profile "deep_research" --agent "architect"
```

### Code Examples
Focus on extracting code blocks.
```bash
./.vector_db/harvest --url "URL" --topic "React Hooks Examples" --profile "code_examples" --agent "developer"
```

### API Documentation
Structured API endpoint extraction.
```bash
./.vector_db/harvest --url "URL" --topic "REST API Docs" --profile "api_documentation" --agent "developer"
```

### Troubleshooting
Problem-solution pairs from forums.
```bash
./.vector_db/harvest --url "URL" --topic "Error Resolution" --profile "troubleshooting" --agent "developer"
```

## Output Location

Harvested content saves to `/harvested/` with subdirectories:
- `documentation/` - Official docs, guides
- `research/` - Deep research results
- `web/` - General web content
- `troubleshooting/` - Problem-solution pairs

Files include metadata frontmatter:
```yaml
---
source: https://original-url.com
harvested_at: 2025-01-15T10:30:00Z
profile: deep_research
agent_context: architect
topic: Topic Name
---
```

## Automatic Indexing

Harvested content is **automatically indexed** in the KB under the `harvested` collection. Search it with:

```bash
./.vector_db/kb search "query" --collection harvested
```

## Common Harvesting Patterns

### Technology Assessment
```bash
./.vector_db/harvest \
    --url "https://newframework.io/docs" \
    --topic "NewFramework Evaluation" \
    --profile "deep_research" \
    --agent "architect"
```

### Library Documentation
```bash
./.vector_db/harvest \
    --url "https://library-docs.com/api" \
    --topic "Library API Reference" \
    --profile "api_documentation" \
    --agent "developer"
```

### Pattern Research
```bash
./.vector_db/harvest \
    --url "https://patterns-site.com/microservices" \
    --topic "Microservices Patterns" \
    --profile "deep_research" \
    --agent "architect"
```

### Debug/Troubleshoot
```bash
./.vector_db/harvest \
    --url "https://stackoverflow.com/questions/tagged/react-hooks" \
    --topic "React Hooks Issues" \
    --profile "troubleshooting" \
    --agent "developer"
```

## Best Practices

1. **Search KB first** - Don't harvest if knowledge exists internally
2. **Use appropriate profile** - Match profile to content type
3. **Set agent context** - Helps with relevance scoring
4. **Be specific with topics** - Descriptive names aid retrieval
5. **Review harvested content** - Verify quality and relevance

## Crawl4ai Features

The harvester uses crawl4ai which provides:
- **6x faster** than traditional crawlers
- **JavaScript rendering** for dynamic sites
- **Adaptive depth** - stops when sufficient info gathered
- **Stealth mode** - handles protected sites
- **Smart extraction** - preserves formatting, tables, code

## Limitations

- Respects robots.txt
- Some sites may block crawling
- Rate limited to avoid overload
- Large sites may timeout
- Quality varies by source
