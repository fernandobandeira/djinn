# Djinn - AI-Powered Development Assistant

A comprehensive development assistant with semantic knowledge base, intelligent agents, and powerful search capabilities.

## Quick Start

### Prerequisites
- Docker and Docker Compose
- Python 3.12+ with `uv` package manager
- Git

### Setup

1. **Start the services**
   ```bash
   docker-compose up -d
   ```
   This starts Qdrant vector database for the knowledge base.

2. **Install Python dependencies**
   ```bash
   cd .vector_db
   uv venv
   uv pip install qdrant-client sentence-transformers tiktoken
   ```

3. **Index your documentation**
   ```bash
   # Index the docs folder
   ./.vector_db/kb index --path docs
   
   # Index everything (may take a while)
   ./.vector_db/kb index
   ```

## Knowledge Base Usage

The KB uses Qdrant for fast semantic search with no file locking issues.

### Search
```bash
# Basic search
./.vector_db/kb search "your query"

# Agent-specific search (optimized for different roles)
./.vector_db/kb search "architecture patterns" --agent architect
./.vector_db/kb search "user stories" --agent product_manager
./.vector_db/kb search "learning concepts" --agent teacher

# Search specific collection
./.vector_db/kb search "query" --collection architecture

# Limit results
./.vector_db/kb search "query" --limit 10
```

### Index Management
```bash
# View statistics
./.vector_db/kb stats

# Index specific directory
./.vector_db/kb index --path docs/architecture

# Force re-index everything
./.vector_db/kb index --force

# Clear and rebuild
./.vector_db/kb clear
./.vector_db/kb index
```

## Custom Commands (Claude Code)

Use these commands in Claude Code for specialized assistance:

- `/analyst` - Business analysis and research
- `/ux` - UX design and user research  
- `/pm` - Product management and PRDs
- `/architect` - System design and architecture
- `/sm` - Scrum master and story creation
- `/dev` - Development and implementation
- `/teacher` - Educational explanations
- `/recruiter` - Agent creation and patterns

## Architecture

### Components
- **Qdrant Server**: Vector database running in Docker (port 6333)
- **Infinity**: High-performance embedding server with GPU acceleration (port 8080)
- **Knowledge Base**: Semantic search with BAAI/bge-large embeddings (1024 dimensions)
- **Agent System**: Specialized AI agents for different roles
- **Collections**: Organized knowledge (docs, architecture, code, zettelkasten, harvested)

### Performance
- **No file locking**: Unlimited parallel searches via Qdrant server mode
- **Ultra-fast embeddings**: ~50ms via Infinity server (vs 3-7s loading model each time)
- **Fast searches**: ~1 second total response time
- **Smart indexing**: Only re-indexes changed files (MD5 hash tracking)
- **Agent optimization**: Context-aware result ranking
- **GPU accelerated**: Infinity uses NVIDIA GPU for embeddings
- **Automatic fallback**: Falls back to local model if Infinity is unavailable

## Configuration

Environment variables in `.env`:
```bash
# Qdrant server URL
QDRANT_URL=http://localhost:6333
```

## Troubleshooting

### Qdrant Connection Issues
```bash
# Check if Qdrant is running
docker-compose ps

# View Qdrant logs
docker-compose logs qdrant

# Restart Qdrant
docker-compose restart qdrant
```

### KB Search Issues
```bash
# Check KB statistics
./.vector_db/kb stats

# Re-index if needed
./.vector_db/kb index --force
```

## Data Storage

- **Qdrant data**: `./data/qdrant/` (persisted across restarts)
- **Document hashes**: `./.vector_db/qdrant_hashes.json`
- **Python environment**: `./.vector_db/.venv/`

## Development

### Adding New Documents
1. Add documents to appropriate directories
2. Run `./.vector_db/kb index --path <dir>`
3. Search will immediately include new content

### Parallel Search Example
```bash
# These can all run simultaneously without blocking
./.vector_db/kb search "architecture" &
./.vector_db/kb search "requirements" &
./.vector_db/kb search "patterns" &
wait
```

## License

[Your License Here]