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
- **agent-planner**: IMPORTANT plans agent architecture with systematic decomposition when creating new agents
- **agent-builder**: IMPORTANT creates agent files and resource structures during agent creation workflow
- **pattern-extractor**: IMPORTANT extracts and learns patterns from successful agent creations for Rita's knowledge base
- **architecture-analyst**: IMPORTANT performs deep architectural analysis of agents and suggests improvements based on complexity, patterns, and best practices
- **coherence-verifier**: IMPORTANT verifies component coherence across agent architecture levels during creation
- **constraint-validator**: IMPORTANT validates agent constraint balance during Rita's agent creation process
- **resource-validator**: IMPORTANT validates resource references and file existence during agent creation

#### Ana's Sub-Agents (Research & Documentation)
- **market-researcher**: Generates comprehensive market research reports
- **competitive-analyzer**: Performs competitive analysis and positioning assessment
- **documentation-generator**: Creates structured documentation from sessions and findings
- **insight-synthesizer**: Extracts patterns and synthesizes insights from data
- **research-architect**: Designs research methodologies and frameworks

#### Teacher's Sub-Agents (Educational Analysis & Zettelkasten)
- **learning-gap-analyzer**: Analyzes conversations to identify knowledge gaps
- **misconception-detector**: Detects and categorizes learner misconceptions
- **concept-difficulty-ranker**: Ranks concepts by learning difficulty
- **learning-path-planner**: Generates optimal learning sequences
- **assessment-generator**: Creates targeted assessment questions
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
