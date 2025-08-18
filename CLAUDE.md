# Project Context for Claude Code

## Knowledge Base System
@.vector_db/KB-INSTRUCTIONS.md

### Custom Commands
- **analyst**: Use `/analyst` to invoke Ana the business analyst
- **architect**: Use `/architect` to invoke Archie the system architect
- **teacher**: Use `/teacher` to invoke Tina the teacher
- **recruiter**: Use `/recruiter` to invoke Rita the recruiter

### Sub-Agents

#### Rita's Sub-Agents (Agent Creation & Analysis)
- **agent-planner**: Plans agent architecture with systematic decomposition
- **agent-builder**: Creates agent files and resource structures
- **pattern-extractor**: Extracts and learns patterns from successful agents
- **architecture-analyst**: Performs deep architectural analysis and improvement suggestions
- **coherence-verifier**: Verifies component coherence across architecture levels
- **constraint-validator**: Validates agent constraint balance during creation
- **resource-validator**: Validates resource references and file existence

#### Teacher's Sub-Agents (Zettelkasten System)
- **zettelkasten-capture**: Captures single atomic notes from learning insights
- **zettelkasten-synthesizer**: Generates synthesis notes from note clusters
- **zettelkasten-relationship-mapper**: Analyzes and establishes note relationships

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
