# Model Selection Guide for Sub-Agents
## August 2025 - Claude 4.1 Era

## Available Models

### Claude Opus 4.1 (claude-opus-4-1-20250805)
**Released**: August 5, 2025
**Pricing**: $15/MTok input, $75/MTok output
**Context**: 200K tokens (1M for Enterprise)
**Max Output**: 32K tokens
**ASL Rating**: ASL-3 (Elevated)

**Best For**:
- Complex architectural refactoring
- Multi-repository debugging
- Research problems requiring deep reasoning
- Algorithm design and optimization
- Multi-step logical problems
- Tasks requiring thousands of steps
- Extended thinking mode for critical decisions

### Claude Sonnet 4 (claude-sonnet-4-20250501)
**Released**: May 2025
**Pricing**: $3/MTok input, $15/MTok output
**Max Output**: 64K tokens
**ASL Rating**: ASL-2 (Standard)

**Best For**:
- Daily coding tasks
- Enterprise workloads
- Customer support
- GitHub issue resolution
- Balance of speed and intelligence
- Most practical coding tasks (72.7% on SWE-bench)

### Claude Haiku 3.5 (claude-haiku-3-5-20241022)
**Released**: October 2024
**Pricing**: $0.25/MTok input, $1.25/MTok output
**Max Output**: 8K tokens
**ASL Rating**: ASL-2 (Standard)

**Best For**:
- Quick responses and brainstorming
- Simple tasks and validations
- Autocompletion
- Real-time interactions
- High-volume processing
- Cost-sensitive operations

## Model Selection Criteria for Sub-Agents

### Critical Planning & Architecture (Use Opus 4.1)
- `agent-planner` - Complex architecture decisions
- `architecture-analyst` - Deep architectural analysis
- `system-designer` - System architecture with multiple options
- `architecture-reviewer` - Architecture improvement opportunities
- `change-coordinator` - Systematic change management

**Rationale**: These require maximum reasoning depth, multi-step analysis, and critical decision-making.

### Complex Analysis & Synthesis (Use Opus 4.1 or Sonnet 4)
- `pattern-extractor` - Learning from successful patterns
- `insight-synthesizer` - Pattern synthesis from data
- `zettelkasten-synthesizer` - Synthesis from note clusters
- `market-researcher` - Comprehensive market research
- `competitive-analyzer` - Competitive positioning

**Rationale**: Require deep analysis but can use Sonnet 4 for speed if cost is a concern.

### Code & Implementation (Use Sonnet 4)
- `agent-builder` - Agent file creation
- `qa-reviewer` - Code review and testing
- `resource-validator` - Resource validation
- `document-processor` - Document processing
- `assessment-generator` - Creating assessments

**Rationale**: Sonnet 4 excels at practical coding with 72.7% SWE-bench score.

### Validation & Checking (Use Haiku 3.5)
- `constraint-validator` - Simple constraint checking
- `coherence-verifier` - Component coherence checks
- `misconception-detector` - Pattern matching in responses
- `learning-gap-analyzer` - Gap identification
- `concept-difficulty-ranker` - Ranking by difficulty

**Rationale**: These are pattern-matching and validation tasks that benefit from speed.

### User Research & Design (Use Sonnet 4)
- `user-researcher` - User research methods
- `design-creator` - Wireframes and prototypes
- `frontend-specifier` - Frontend specifications
- `usability-validator` - Usability validation

**Rationale**: Balance of creativity and practical output.

### Management & Coordination (Use Sonnet 4)
- `product-strategist` - PRDs and strategy
- `stakeholder-coordinator` - Communications
- `execution-tracker` - Sprint tracking
- `kb-analyst` - Knowledge base operations

**Rationale**: Need reliability and consistency for ongoing operations.

### Diagram & Documentation (Use Haiku 3.5 or Sonnet 4)
- `diagram-generator` - Simple diagram generation (Haiku)
- `adr-manager` - ADR management (Sonnet)
- `documentation-generator` - Structured docs (Sonnet)
- `pattern-librarian` - Pattern management (Haiku)

**Rationale**: Haiku for simple generation, Sonnet for complex documentation.

## Decision Matrix

| Complexity | Reasoning Depth | Speed Priority | Cost Priority | Model Choice |
|------------|----------------|----------------|---------------|--------------|
| High       | Critical       | Low            | Low           | Opus 4.1     |
| High       | High           | Medium         | Medium        | Sonnet 4     |
| Medium     | Medium         | High           | Medium        | Sonnet 4     |
| Low        | Low            | Critical       | High          | Haiku 3.5    |
| Simple     | Pattern Match  | Critical       | Critical      | Haiku 3.5    |

## Special Considerations

### When to Use Extended Thinking Mode
Available only for Opus 4.1 and Sonnet 4:
- Architecture planning decisions
- Complex refactoring analysis
- Multi-step problem solving
- Critical path identification
- System design with trade-offs

### When to Upgrade Model Choice
Consider upgrading from default when:
- Task involves multiple repositories
- Requires sustained multi-hour work
- Needs deep reasoning chains
- Involves critical business decisions
- Requires creative problem solving

### When to Downgrade Model Choice
Consider downgrading when:
- Task is highly repetitive
- Pattern matching is sufficient
- Speed is critical
- Budget constraints exist
- Volume processing is needed

## Implementation in Agent Files

```yaml
model:
  provider: anthropic
  id: claude-opus-4-1-20250805  # For critical tasks
  # id: claude-sonnet-4-20250501  # For balanced tasks
  # id: claude-haiku-3-5-20241022  # For simple tasks
  
  thinking_mode: enabled  # Only for Opus/Sonnet
  extended_reasoning: true  # For complex analysis
```

## Cost Optimization Strategy

1. **Default to Sonnet 4** for most sub-agents
2. **Upgrade to Opus 4.1** only for:
   - Initial planning phases
   - Critical architecture decisions
   - Complex multi-step reasoning
3. **Use Haiku 3.5** for:
   - All validation checks
   - Simple pattern matching
   - High-volume operations
4. **Enable thinking mode** selectively:
   - Only for critical decisions
   - During planning phases
   - For complex problem solving

## Monitoring & Adjustment

Track these metrics per sub-agent:
- Task completion rate
- Error frequency
- Token usage
- Response time
- User satisfaction

Adjust model selection based on:
- Consistent failures → Upgrade model
- Consistent success → Consider downgrade
- Speed complaints → Consider Haiku
- Quality issues → Upgrade to Opus