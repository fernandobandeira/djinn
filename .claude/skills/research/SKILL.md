---
name: research
description: Discover and gather knowledge effectively using KB semantic search and external harvesting. Use when user wants to "search for", "find existing", "what do we know about", "research", "look up", "check if we have", or needs to discover existing knowledge before creating something new.
allowed-tools: Bash, Read, Grep, Glob
---

# Research - Knowledge Discovery Skill

Discover existing knowledge before creating new content. Search the KB semantically, harvest external sources, and synthesize findings.

## Quick Start

1. **Always search KB first** before creating anything
2. Select technique based on what you need
3. Evaluate sources and synthesize findings

## Technique Selection

| Need | Use | Why |
|------|-----|-----|
| Find existing internal knowledge | KB Search | Semantic search across indexed docs |
| Gather external documentation | Knowledge Harvesting | Crawl and index external sources |
| Assess what you found | Source Evaluation | Determine relevance and quality |

**Default**: Always start with KB Search. Only harvest externally if internal knowledge is insufficient.

## Techniques

### KB Search
Semantic search across all indexed project documents using Qdrant vector database.

Read [cookbook/kb-search.md](./cookbook/kb-search.md)

### Knowledge Harvesting
Extract and index knowledge from external sources using crawl4ai.

Read [cookbook/knowledge-harvesting.md](./cookbook/knowledge-harvesting.md)

### Source Evaluation
Assess relevance, quality, and applicability of discovered knowledge.

Read [cookbook/source-evaluation.md](./cookbook/source-evaluation.md)

## Core Principles

1. **Search before create** - Always check what exists first
2. **Internal first** - Prefer existing project knowledge over external
3. **Semantic over keyword** - Use meaning-based queries, not just keywords
4. **Agent context matters** - Use `--agent` flag for optimized results
5. **Read full files** - Search returns chunks; read full docs for context

## KB Commands Reference

```bash
# Semantic search
./.vector_db/kb search "query" --agent [architect|analyst|developer]

# Search specific collection
./.vector_db/kb search "query" --collection [docs|architecture|code|harvested]

# Index new documents
./.vector_db/kb index

# Harvest external content
./.vector_db/harvest --url "URL" --topic "TOPIC" --profile "PROFILE" --agent "AGENT"
```

## When Research Works Best

- Before creating ADRs, patterns, or documentation
- When exploring a new problem space
- When you need prior art or examples
- Before brainstorming (know what's been tried)
- When debugging (search for similar issues)

## The Research Mindset

**Brownfield thinking**: Assume relevant knowledge already exists. Your job is to find it, build on it, and extend it - not recreate it.

Questions to ask:
- "What do we already know about this?"
- "Has this been solved before?"
- "What patterns exist?"
- "What decisions were already made?"
