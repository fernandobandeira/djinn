# Project Context for Claude Code

## Date and Time Context
Today's date is provided in the environment context. All agents and sub-agents should:
- Use the current date (format: YYYY-MM-DD) when creating timestamped files
- Use the current year when searching for recent documentation or best practices
- Reference "Today's date: YYYY-MM-DD" from the environment context when needed
- Example: When creating ADRs, use `ADR-YYYYMMDD-{title}.md` format with actual date

## Knowledge Base System
`.vector_db/KB-INSTRUCTIONS.md`

### MANDATORY: Knowledge Base Usage for All Agents

**CRITICAL**: All command agents MUST use the kb-analyst sub-agent for project discovery and search operations. Direct file system commands (find, grep, ls) should NEVER be used for searching project information.

#### Required Initial Workflow for Commands
When any command agent is activated and asked to analyze a project, they MUST:
1. FIRST delegate to kb-analyst for initial project discovery
2. Use kb-analyst for all subsequent searches
3. Never use raw Bash commands for file discovery (find, grep, ls)
4. Use Read/Glob/Grep tools only AFTER kb-analyst provides specific file paths

#### Example: Proper Project Analysis Start
```python
# CORRECT - Using kb-analyst
Task(
  subagent_type="kb-analyst",
  description="Initial project discovery",
  prompt="Agent context: [agent_type]
         Search query: project overview requirements documentation
         Collection focus: documentation, architecture, zettelkasten
         Detail level: comprehensive"
)

# WRONG - Using raw commands
Bash("find . -name '*.md' | grep -E 'readme|design'")
```

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

## Core Development Principles

### Brownfield Development Principle

**CRITICAL: Always Brownfield Approach**
- The project is ALWAYS considered a brownfield environment
- Mandatory workflow for ALL agents and commands:
  1. FIRST discover existing systems and resources
  2. ALWAYS build upon existing work
  3. NEVER recreate what already exists
  4. FOLLOW existing patterns and conventions
  5. EXTEND current architectures, do not replace them

#### Brownfield Workflow Implications
- kb-analyst discovery is MANDATORY for every project task
- Comprehensive existing system understanding precedes any new work
- Reuse and extend existing code, documentation, and patterns
- Minimize redundancy and maximize incremental improvement

### Documentation Structure

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