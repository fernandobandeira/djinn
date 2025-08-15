# Project Context for Claude Code

## Knowledge Base System
@.vector_db/KB-INSTRUCTIONS.md

### Custom Commands
- **analyst**: Use `/analyst` to invoke Ana the business analyst
- **architect**: Use `/architect` to invoke Archie the system architect
- **teacher**: Use `/teacher` to invoke Tina the teacher
- **recruiter**: Use `/recruiter` to invoke Rita the recruiter

### Sub-Agents
- **zettelkasten-guide**: Zettelkasten knowledge capture specialist (auto-delegated during learning sessions under Teacher)

## Documentation Structure

### Architecture Documentation (`/docs/architecture/`)
- `adrs/` - Architecture Decision Records
- `patterns/` - Reusable architectural patterns
- `standards/` - Coding and architecture standards
- `diagrams/` - Architecture diagrams (Mermaid, PlantUML)
- `reviews/` - Architecture review documents

### Analysis Documentation (`/docs/`)
- `analysis/` - Research and analysis documents
- `brainstorming/` - Brainstorming session results
- `research/` - Market research documents
