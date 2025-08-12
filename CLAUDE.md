# Project Context for Claude Code

## Project Overview
This project includes a unified knowledge base system with vector database support for intelligent code understanding and architectural decision-making.

## Knowledge Base System

### Unified Knowledge Base Tool
**Simple CLI**: `./kb` - A wrapper script for the knowledge base
- Located at `.vector_db/kb.py` with virtual environment
- Indexes documentation in `/docs` folder (safety restricted)
- Stores vector data in `.vector_db/` (hidden directory)
- Intelligent parsing and semantic search
- Multi-collection organization (architecture, documentation, etc.)
- Special handling for ADRs and patterns

### Vector Database
- **Storage**: `.vector_db/` - Hidden directory at project root
- **Technology**: ChromaDB with sentence-transformers (all-MiniLM-L6-v2)
- **Collections**: architecture, documentation, code, config, api, tests, scripts

### Custom Commands
- **architect**: Use `/architect` to invoke the technical architect assistant

## Documentation Structure

### Architecture Documentation (`/docs/architecture/`)
- `adrs/` - Architecture Decision Records
- `patterns/` - Reusable architectural patterns
- `standards/` - Coding and architecture standards
- `diagrams/` - Architecture diagrams (Mermaid, PlantUML)
- `templates/` - Document templates
- `runbooks/` - Operational runbooks (create as needed)

### Available Templates
All templates in `/docs/architecture/templates/`:
- `adr-template.md` - Architecture Decision Records
- `pattern-template.md` - Architectural patterns
- `rfc-template.md` - Request for Comments
- `runbook-template.md` - Operational runbooks

## Using the Knowledge Base

### Quick Start
```bash
# Index documentation
./kb index

# Search for something
./kb search "authentication"

# View statistics
./kb stats
```

### Indexing Options
```bash
# Index all docs (default)
./kb index

# Index only architecture docs
./kb index --mode arch

# Index only documentation
./kb index --mode docs

# Force re-index all files
./kb index --force

# Index specific path (must be in /docs)
./kb index --path ./docs/architecture
```

### Search Options
```bash
# Basic search
./kb search "query"

# Search specific collection
./kb search "database" --collection architecture

# Get more results
./kb search "pattern" --limit 10

# Filter by file type
./kb search "ADR" --type .md
```

### Management
```bash
# View statistics
./kb stats

# Clear specific collection
./kb clear --collection architecture --confirm

# Clear entire knowledge base
./kb clear --confirm
```

## Development Workflow

### Before Writing Code
1. Search the knowledge base for similar implementations
2. Check existing patterns: `./kb search "pattern name"`
3. Review architecture decisions: `./kb search "decision" --collection architecture`

### Before Making Architectural Changes
1. Search existing ADRs and patterns
2. Consult the architect assistant: `/architect analyze [problem]`
3. Document decisions using templates
4. Update the knowledge base after changes

### After Making Changes
1. Re-index documentation: `./kb index --force`
2. Update documentation if needed
3. The knowledge base will track changes via file hashing

## Creating Documentation

### Architecture Decision Record (ADR)
```bash
# Copy template
cp docs/architecture/templates/adr-template.md \
   docs/architecture/adrs/ADR-$(date +%Y%m%d)-decision-name.md

# Edit the ADR
# Then re-index
./kb index --mode arch
```

### Design Document
```bash
cp docs/architecture/templates/design-template.md \
   docs/architecture/designs/component-name.md
```

### Pattern Documentation
```bash
cp docs/architecture/templates/pattern-template.md \
   docs/architecture/patterns/pattern-name.md
```

## Advanced Features

### Multi-Collection Search
The knowledge base automatically categorizes content:
- **architecture**: ADRs, designs, patterns
- **documentation**: README files, docs
- **code**: Source code with extracted docstrings
- **config**: Configuration files
- **api**: API specifications
- **tests**: Test files
- **scripts**: Utility scripts

### Intelligent Code Parsing
- Extracts Python docstrings and comment blocks
- Parses JavaScript/TypeScript JSDoc comments
- Chunks code intelligently for better retrieval
- Maintains context with overlapping chunks

### Change Detection
- Uses MD5 hashing to detect file changes
- Only re-indexes modified files
- Preserves unchanged embeddings for efficiency

## Project-Specific Context
<!-- Add your specific project details here -->
[Add information about your tech stack, constraints, team preferences, etc.]