# KB Search Cookbook

Semantic search across indexed project documents using Qdrant vector database.

## When to Use

- Before creating any new document
- When exploring what exists on a topic
- When you need prior decisions or patterns
- To find related work before brainstorming

## Basic Search

```bash
./.vector_db/kb search "your query here"
```

## Agent-Optimized Search

Different agents get results optimized for their needs:

```bash
# For architecture work - prioritizes ADRs, patterns, technical docs
./.vector_db/kb search "query" --agent architect

# For analysis work - prioritizes research, strategy docs
./.vector_db/kb search "query" --agent analyst

# For development work - prioritizes code, implementation docs
./.vector_db/kb search "query" --agent developer
```

## Collection-Specific Search

Target specific document types:

```bash
# Architecture decisions and patterns
./.vector_db/kb search "query" --collection architecture

# General documentation
./.vector_db/kb search "query" --collection docs

# Source code
./.vector_db/kb search "query" --collection code

# External harvested content
./.vector_db/kb search "query" --collection harvested
```

## Understanding Results

```
Found 5 results:

1. Collection: architecture
   Score: 0.7004
   File: /docs/architecture/adrs/ADR-20250120-error-handling.md:16
   Source: internal
   Preview: Error handling strategy using...
```

**Key fields:**
- **Score**: 0.0-1.0, higher = more relevant (0.7+ is highly relevant)
- **File:N**: Path with chunk index (`:16` means chunk 16)
- **Source**: `internal` (your docs) vs `harvested` (external)
- **Preview**: First 200 chars of matching chunk

## After Finding Results

1. **Read the full file** - chunks are excerpts, get full context
2. **Note the chunk index** - helps locate specific section
3. **Check source type** - internal docs are more authoritative
4. **Cross-reference** - search related terms for completeness

## Query Construction Tips

### Be Semantic, Not Keyword-Based
```bash
# Good - semantic meaning
./.vector_db/kb search "handling errors gracefully"

# Less effective - just keywords
./.vector_db/kb search "error"
```

### Use Domain Language
```bash
# For architecture
./.vector_db/kb search "microservices communication patterns"

# For analysis
./.vector_db/kb search "competitive landscape assessment"
```

### Combine Concepts
```bash
./.vector_db/kb search "authentication security best practices"
```

## Common Search Patterns

### Before Creating an ADR
```bash
./.vector_db/kb search "[decision topic]" --agent architect --collection architecture
```

### Before Writing Code
```bash
./.vector_db/kb search "[functionality] implementation" --agent developer
```

### Before Research/Analysis
```bash
./.vector_db/kb search "[topic] analysis research" --agent analyst
```

### Finding Prior Art
```bash
./.vector_db/kb search "similar to [concept]" --collection architecture
```

## Indexing New Content

After creating documents, index them:

```bash
# Smart index (only changed files)
./.vector_db/kb index

# Index specific path
./.vector_db/kb index --path /docs/architecture/

# Force complete re-index
./.vector_db/kb index --force
```

## Troubleshooting

**No results?**
1. Try broader query terms
2. Remove collection filter
3. Try different agent context
4. Check if content is indexed: `./.vector_db/kb stats`

**Too many irrelevant results?**
1. Add collection filter
2. Use agent context
3. Be more specific in query
4. Limit results: `--limit 5`
