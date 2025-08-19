# Harvested Knowledge Repository

This directory contains automatically harvested content from the knowledge-harvester agent.

## Directory Structure

- `web/` - General web content and articles
- `research/` - Deep research on technologies and patterns
- `documentation/` - Official documentation and API references
- `troubleshooting/` - Problem-solution pairs from forums

## Automatic Processing

When content is harvested:
1. Saved here with timestamp and source metadata
2. Automatically indexed in ChromaDB
3. GraphRAG extracts entities and relationships
4. Available for all agents to search

## File Naming Convention

Files are named: `{yyyy-mm-dd}_{profile}_{topic_slug}_{hash}.md`

Example: `2024-01-19_deep_research_react_hooks_a3f2b1.md`

## Metadata Format

Each file includes:
```yaml
---
source: URL or source identifier
harvested_at: ISO timestamp
profile: harvesting profile used
agent_context: requesting agent
confidence: extraction confidence score
---
```

## Collections Mapping

- Files in `documentation/` → `documentation` collection
- Files in `research/` → `research` collection  
- Files in `troubleshooting/` → `troubleshooting` collection
- Others → `harvested` collection