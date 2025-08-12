# Knowledge Base System Documentation

## Overview
The project uses a vector-based knowledge base system for intelligent semantic search across documentation. This enables the architect assistant and developers to quickly find relevant information based on meaning, not just keywords.

## Installation
The system is already installed with dependencies in `.vector_db/.venv/`:
- ChromaDB 1.0.16 - Vector database
- Sentence-Transformers 5.1.0 - Embedding generation
- All-MiniLM-L6-v2 model - Balanced speed/quality

## Usage

### Quick Command Reference
```bash
# Index all documentation
./kb index

# Search for concepts
./kb search "microservices architecture"

# View what's indexed
./kb stats

# Get help
./kb --help
```

### Indexing Modes
- `./kb index` - Index all docs (default)
- `./kb index --mode arch` - Only architecture docs
- `./kb index --mode docs` - Only documentation
- `./kb index --force` - Force re-index everything

### Search Features
- Semantic search understands meaning and context
- Returns relevance scores for each result
- Can filter by collection or file type
- Shows document metadata (ADR status, titles, etc.)

## Architecture

### Components
1. **kb wrapper** (`/kb`) - Simple bash script for CLI access
2. **kb.py** (`.vector_db/kb.py`) - Main Python application
3. **ChromaDB** (`.vector_db/chromadb/`) - Persistent vector storage
4. **Virtual Environment** (`.vector_db/.venv/`) - Isolated dependencies

### Collections
- `architecture` - ADRs, patterns, standards
- `documentation` - README files, docs, guides
- `code` - Source code (if enabled)
- `config` - Configuration files (if enabled)

### Safety Features
- **Path Restriction**: Only indexes `/docs` folder by default
- **Hidden Folders**: Automatically excludes `.git`, `.venv`, etc.
- **File Size Limit**: Skips files over 10MB
- **Gitignore**: Keeps vector database out of version control

## How It Works

### Indexing Process
1. Scans `/docs` folder for markdown files
2. Parses documents to extract metadata (titles, status, etc.)
3. Chunks large documents with overlap for context
4. Generates embeddings using sentence-transformers
5. Stores vectors in ChromaDB with metadata

### Search Process
1. Converts query to embedding vector
2. Finds nearest neighbors in vector space
3. Returns results ranked by similarity
4. Includes metadata for context

### Change Detection
- Uses MD5 hashing to detect file changes
- Only re-indexes modified files
- Preserves embeddings for unchanged content

## Integration with Architect Assistant

The `/architect` command automatically uses the knowledge base to:
- Search for similar problems and solutions
- Reference existing ADRs and patterns
- Build on past architectural decisions
- Maintain consistency across documentation

## Maintenance

### Updating the Index
```bash
# After adding new documentation
./kb index

# After major changes
./kb index --force
```

### Monitoring Usage
```bash
# Check statistics
./kb stats

# See what's indexed
ls -la .vector_db/
```

### Troubleshooting
```bash
# Clear and rebuild if issues
./kb clear --confirm
./kb index --force

# Check virtual environment
.vector_db/.venv/bin/python --version
```

## Best Practices

1. **Index Regularly**: Run `./kb index` after adding documentation
2. **Use Semantic Queries**: Search for concepts, not exact text
3. **Check ADR Status**: Review status in search results
4. **Keep Docs Organized**: Use proper folder structure
5. **Write Clear Titles**: First # heading becomes searchable title

## Technical Details

### Embedding Model
- **Model**: all-MiniLM-L6-v2
- **Dimensions**: 384
- **Performance**: ~2000 docs/minute
- **Quality**: Good for technical documentation

### Storage
- **Location**: `.vector_db/chromadb/`
- **Format**: DuckDB + Parquet files
- **Persistence**: Survives restarts
- **Size**: ~1KB per document chunk

### Chunking Strategy
- **Size**: 1000 characters default
- **Overlap**: 200 characters
- **Preserves**: Paragraph boundaries
- **Special**: Architecture docs get smaller chunks

## Future Enhancements

Potential improvements to consider:
- Fine-tune embeddings on architecture terminology
- Add code search capabilities
- Integrate with IDE plugins
- Create web UI for search
- Add incremental indexing via file watchers