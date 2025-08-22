# Knowledge Base Instructions for AI Agents

## Overview
The knowledge base uses **Qdrant** vector database with **Infinity** embedding server for ultra-fast semantic search.

### 🚀 Architecture
- **Qdrant Server** - Vector database running in Docker (no file locking)
- **Infinity Server** - GPU-accelerated embedding service (keeps model in memory)
- **Automatic Fallback** - Uses local model if Infinity is unavailable

### ⚡ Performance Features
- **~1 second searches** - Down from 7+ seconds with local model loading
- **Unlimited parallel access** - No file locking with server mode
- **GPU Accelerated** - Infinity uses NVIDIA GPU for embeddings (~50ms)
- **Semantic Search** - Finds conceptually related content using BAAI/bge-large-en-v1.5
- **1024-dimensional embeddings** - Best quality for superior understanding
- **Incremental Indexing** - Only reindexes changed files (MD5 hash tracking)
- **Agent-optimized** - Use `--agent architect|teacher|developer|analyst` for context

## Quick Start

### Search
```bash
# Basic semantic search
./.vector_db/kb search "error handling"

# Specify agent context for better results
./.vector_db/kb search "architecture patterns" --agent architect
./.vector_db/kb search "learning concepts" --agent teacher
./.vector_db/kb search "implementation" --agent developer
./.vector_db/kb search "market research" --agent analyst
```

#### Search Output Example
```
📊 Found 5 results:

1. Collection: architecture
   Score: 0.7004
   File: /home/fernando/git/djinn/docs/architecture/adrs/ADR-20250120-error-handling-logging-strategy.md:16
   Source: internal
   Preview: {
        return fmt.Errorf("failed to get user %s from database: %w", userID, err)
    }
    
    if err := validateUser(user); err != nil {
        return fmt.Errorf("user %s validation failed: %w",...

2. Collection: harvested
   Score: 0.6874
   File: /home/fernando/git/djinn/harvested/research/2025-08-19_deep_research_go_cryptography.md:151
   Source: harvested
   Preview: we can examine for more information about the error...
```

**Key Information in Results:**
- **File Path**: Full path to the source document
- **Chunk Index**: Number after `:` shows which part of the file (e.g., `:16` is chunk 16)
- **Source**: `internal` (your docs) vs `harvested` (external research)
- **Score**: Semantic similarity score (higher = more relevant)
- **Preview**: First 200 characters of the matching chunk

#### Reading Full Context
The search shows relevant chunks, but **you should read the full file for complete context**:
```bash
# Read the full file from search result
cat /home/fernando/git/djinn/docs/architecture/adrs/ADR-20250120-error-handling-logging-strategy.md

# Or open in editor
vim /home/fernando/git/djinn/docs/architecture/adrs/ADR-20250120-error-handling-logging-strategy.md
```

### Index Documents
```bash
# Index everything (smart - only changed files)
./.vector_db/kb index

# Index specific path
./.vector_db/kb index --path ./docs/architecture

# Force complete re-index
./.vector_db/kb index --force
```

### Check Status
```bash
# View KB health and statistics
./.vector_db/kb status
```

### No Server Required!
Qdrant runs in **local mode** - no server process needed. Just use the `kb` command directly.

## Collections

| Collection | Content | Example Search |
|------------|---------|----------------|
| `docs` | General documentation, guides | `./.vector_db/kb search "requirements" --collection docs` |
| `architecture` | ADRs, patterns, designs | `./.vector_db/kb search "ADR" --collection architecture` |
| `zettelkasten` | Learning notes, insights | `./.vector_db/kb search "concept" --collection zettelkasten` |
| `code` | Source code files | `./.vector_db/kb search "function" --collection code` |
| `harvested` | External web content | `./.vector_db/kb search "best practices" --collection harvested` |

## Agent-Specific Search Optimization

### For Architects (Archie)
```bash
./.vector_db/kb search "system design" --agent architect
# Prioritizes: architecture, documentation, code collections
# Boosts: "design", "pattern", "adr", "system" terms
```

### For Teachers (Tina)
```bash
./.vector_db/kb search "learning concept" --agent teacher
# Prioritizes: zettelkasten, documentation collections
# Boosts: "learning", "concept", "explanation" terms
```

### For Developers (Dave)
```bash
./.vector_db/kb search "implementation" --agent developer
# Prioritizes: code, tests, api collections
# Boosts: "implementation", "function", "test" terms
```

### For Analysts (Ana)
```bash
./.vector_db/kb search "market analysis" --agent analyst
# Prioritizes: documentation, harvested collections
# Boosts: "analysis", "research", "market" terms
```

### For Product Managers (Paul)
```bash
./.vector_db/kb search "user requirements" --agent product
# Prioritizes: documentation, architecture collections
# Boosts: "requirement", "user", "feature" terms
```

## Search Best Practices

### 1. Always Search Before Creating
```bash
# Before creating an ADR
./.vector_db/kb search "[decision topic]" --agent architect

# Before brainstorming
./.vector_db/kb search "[feature area]" --agent analyst

# Before writing code
./.vector_db/kb search "[functionality]" --agent developer
```

### 2. Use Agent Context
The `--agent` flag optimizes search for your specific role:
- Query preprocessing tailored to agent type
- Collection prioritization based on agent needs  
- Term boosting for agent-relevant concepts
- Result ranking optimized for agent context

### 3. Cross-Reference Collections
```bash
# Find business context for technical decisions
./.vector_db/kb search "payment" --collection documentation
./.vector_db/kb search "payment" --collection architecture
./.vector_db/kb search "payment" --collection code
```

### 4. Use Specific Collections When Known
```bash
# If you know it's an ADR
./.vector_db/kb search "database selection" --collection architecture --agent architect

# If you know it's a learning note
./.vector_db/kb search "recursion" --collection zettelkasten --agent teacher
```

## Indexing

### Smart Incremental Indexing
The KB automatically detects changed files using MD5 hashes:
- **New files** → Indexed automatically
- **Modified files** → Re-indexed automatically
- **Unchanged files** → Skipped for speed
- **Deleted files** → Removed from index

No need to manually track what needs indexing!

### Manual Indexing
```bash
# Index everything (smart - only changed files)
./.vector_db/kb index

# Index specific paths
./.vector_db/kb index --path ./docs/architecture/adrs/
./.vector_db/kb index --path ./zettelkasten/permanent/
```

## Performance

| Operation | Qdrant | Old System |
|-----------|--------|------------|
| Index 100 docs | ~5 seconds | 30-60 seconds |
| Search query | 10-50ms | 500-2000ms |
| Memory usage | ~200MB | 2-4GB |
| GPU required | No | Preferred |

## How It Works

### Semantic Search
Unlike keyword search, Qdrant understands **meaning**:
- "error handling" finds docs about exceptions, logging, monitoring
- "user authentication" finds docs about login, OAuth, sessions
- "performance" finds docs about optimization, caching, scalability

### Relevance Scoring
The KB applies intelligent scoring to prioritize high-value documentation:

**Boosted Content** (Higher Priority):
- `architecture/adrs/` → +30% (Architecture Decision Records)
- `architecture/patterns/` → +25% (Architectural Patterns)
- `/analysis/` → +20% (Business Analysis)
- `/strategy/` → +20% (Strategic Documents)
- Other `architecture/` → +15% (Technical Docs)

**Reduced Priority**:
- `/research/` → -10% (Raw research data)
- `/harvested/` → -15% (External content)

Agent-specific boosts are also applied when using `--agent` flag.

### Vector Embeddings
Text is converted to numerical vectors that capture semantic meaning:
```
"testing strategy" → [0.12, -0.34, 0.56, ...] (1024 dimensions)
```
Similar concepts have similar vectors, enabling semantic search.

## Troubleshooting

### No Results Found
1. Check if documents are indexed: `./.vector_db/kb stats`
2. Try without collection filter: `./.vector_db/kb search "query"`
3. Try different agent context: `./.vector_db/kb search "query" --agent architect`

### Concurrent Access Error
Qdrant local mode doesn't support concurrent access. If you see:
```
❌ Error: Qdrant storage is being used by another process
```
This means indexing is running. Wait a moment and try again. The KB will retry automatically for searches.

### Indexing Issues
- Ensure files have supported extensions (.md, .py, .js, etc.)
- Check file isn't in excluded directories (node_modules, .git, etc.)
- Try force reindex: `./.vector_db/kb index --force`

### Memory Issues
Qdrant uses minimal memory (~200MB). If issues persist:
```bash
./.vector_db/kb clear && ./.vector_db/kb index
```

## File Types Indexed

Automatically indexes:
- **Markdown**: `.md`
- **Code**: `.py`, `.js`, `.ts`, `.jsx`, `.tsx`, `.go`, `.rs`
- **Config**: `.yaml`, `.yml`, `.json`
- **Text**: `.txt`

Excluded directories: `.git`, `node_modules`, `__pycache__`, `.venv`, `build`, `dist`

## Remember

- **Search first, create second** - Always check existing knowledge
- **Use agent context** - Better results with `--agent` flag
- **Index regularly** - Keep knowledge current
- **Cross-reference** - Check multiple collections for complete context
- **Let auto-indexer run** - Real-time updates improve search quality

The KB provides instant, semantic search across all your project knowledge, making it easy for LLMs to find relevant context quickly and accurately.