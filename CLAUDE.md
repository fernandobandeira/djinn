# Project Context for Claude Code

## Knowledge Base System
.vector_db/KB-INSTRUCTIONS.md

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
