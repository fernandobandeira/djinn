# Djinn

Agent architecture project with reusable skills and shared sub-agents.

## Knowledge Base System

**MANDATORY**: All agents must use the `research` skill for project discovery and search.

### KB-First Workflow
1. **Invoke `research` skill** for any discovery task
2. Search KB before creating anything
3. Use agent context (`--agent architect|analyst|developer`) for optimized results
4. Read full files after finding relevant chunks

### KB Commands
```bash
# Semantic search
./.vector_db/kb search "query" --agent architect

# Index documents after creation
./.vector_db/kb index

# Harvest external content
./.vector_db/harvest --url "URL" --topic "TOPIC" --profile "PROFILE" --agent "AGENT"
```

### Search Before Create
- Before creating an ADR → search existing ADRs
- Before creating a pattern → search existing patterns
- Before brainstorming → search prior sessions

See `.vector_db/KB-INSTRUCTIONS.md` for full documentation.

## Skills Architecture

Skills teach HOW to think. They're organized in tiers:

### Tier 1 (Universal)
- `research` - **Use first** - KB search, harvesting, source evaluation
- `root-cause` - Five Whys, First Principles, Jobs-to-be-Done
- `ideation` - SCAMPER, Walt Disney Method, Reverse Brainstorming
- `devils-advocate` - Red Team/Blue Team, Pre-mortem Analysis
- `role-playing` - Six Thinking Hats, Stakeholder Roundtable
- `teacher` - Socratic Dialogue, Feynman Technique

### Tier 2 (Domain)
- `strategic-analysis` - SWOT, Porter's Five Forces, Scenario Planning
- `user-research` - Journey Mapping, Interview Design
- `systems-thinking` - Systems Mapping, Feedback Loops

## Shared Sub-Agents

Delegate execution work via Task tool to these in `.claude/agents/shared/`:

### Research & Analysis
- `market-researcher` - Market research reports
- `competitive-analyzer` - Competitive analysis
- `research-architect` - Research methodology design
- `insight-synthesizer` - Pattern extraction from data

### Documentation & Visualization
- `documentation-generator` - Structured documentation
- `diagram-generator` - Mermaid/PlantUML technical diagrams

## Key Principles

1. **Skills teach HOW to think. Sub-agents DO work.**
2. **Research First**: Always search before creating (use `research` skill)
3. **Brownfield**: Build on existing knowledge, never recreate

## Agent Creation

Use the `agent-recruitment` skill to create new agents. Check `.claude/skills/agent-recruitment/decision-frameworks/reusability-assessment.md` to determine if something should be a skill, shared sub-agent, or agent-specific.
