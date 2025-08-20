# Project Context for Claude Code

## Date and Time Context
Today's date is provided in the environment context. All agents and sub-agents should:
- Use the current date (format: YYYY-MM-DD) when creating timestamped files
- Use the current year when searching for recent documentation or best practices
- Reference "Today's date: YYYY-MM-DD" from the environment context when needed
- Example: When creating ADRs, use `ADR-YYYYMMDD-{title}.md` format with actual date

## Knowledge Base System
**Documentation**: `.vector_db/KB_ENHANCED_DOCUMENTATION.md`  
**Quick Reference**: `.vector_db/KB-INSTRUCTIONS.md`

### Enhanced KB Features
- **Better Search**: BAAI/bge-large-en-v1.5 embeddings with hybrid BM25+vector search
- **Auto-Indexing**: File watcher automatically indexes changes
- **Agent Context**: Search optimized for specific agents with `--agent` flag
- **Smart Query Processing**: Handles LLM verbosity automatically
- **Document Categorization**: Tracks source (internal/harvested) and creator agent

### Custom Commands
- **analyst**: Use `/analyst` to invoke Ana the business analyst (research, brainstorming, competitive analysis)
- **ux**: Use `/ux` to invoke Ulysses the UX Designer (user research, personas, wireframes, prototypes, design specs)
- **pm**: Use `/pm` to invoke Paul the product manager (PRDs, product strategy, roadmaps, stakeholder coordination)
- **architect**: Use `/architect` to invoke Archie the system architect (system design, ADRs, technical architecture)
- **sm**: Use `/sm` to invoke Sam the scrum master (story creation, sprint planning, backlog management)
- **dev**: Use `/dev` to invoke Dave the developer (implementation, TDD, debugging, code quality)
- **teacher**: Use `/teacher` to invoke Tina the teacher (educational explanations, learning paths, Zettelkasten)
- **recruiter**: Use `/recruiter` to invoke Rita the recruiter (agent creation, architecture analysis, pattern extraction)

### Shared Sub-Agents (Available to All Commands)
- **kb-analyst**: Universal Knowledge Base search and indexing operations across all collections
- **qa-reviewer**: Quality assurance, TDD support, code review, acceptance criteria generation
- **knowledge-harvester**: Intelligent documentation harvesting, external research, pattern discovery
- **plan-validator**: Validates project plans, PRDs, and architectures during planning phase

#### Using Shared Sub-Agents
Commands delegate to shared sub-agents via the Task tool:
```
Task(
  subagent_type="knowledge-harvester",
  description="Research [topic]",
  prompt="Agent context: [your_agent_type]
         Research topic: [specific_topic]
         Scope: [comprehensive|focused|quick_reference]
         Format preference: [documentation|examples|comparison|tutorial]"
)
```

#### knowledge-harvester Usage Examples
- **Architecture Research**: Finding design patterns, ADRs, technology assessments
- **Code Examples**: Library usage, implementation patterns, troubleshooting guides
- **Market Intelligence**: Competitive analysis, industry trends, user feedback
- **Design Patterns**: UI/UX patterns, accessibility standards, design systems
- **Learning Resources**: Tutorials, documentation, educational materials

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
