# Create Architecture Decision Record Task

## Purpose
Guide the creation of lightweight, effective Architecture Decision Records (ADRs) that capture significant architectural decisions with their context and rationale.

## ADR Creation Process

### Step 1: Determine ADR Necessity
Ask yourself:
- Is this architecturally significant?
- Will it affect system structure or quality?
- Will future teams wonder "why did they do it this way?"
- Does it involve trade-offs worth documenting?

If yes to any → Create an ADR

### Step 2: Gather Context

#### Questions to Answer:
1. **What problem are we solving?**
   - Current pain points
   - Technical debt being addressed
   - New requirements being met

2. **What constraints exist?**
   - Technical limitations
   - Budget/time constraints
   - Team expertise
   - Existing system dependencies

3. **Who are the stakeholders?**
   - Who's affected by this decision?
   - Who needs to approve it?
   - Who will implement it?

4. **What's the urgency?**
   - Why decide now vs later?
   - What happens if we don't decide?

### Step 3: Explore Alternatives

For each viable option, document:
- **Description**: What is this approach?
- **Pros**: Benefits and strengths
- **Cons**: Drawbacks and weaknesses
- **Effort**: Implementation complexity
- **Risk**: What could go wrong?
- **Cost**: Financial and resource implications

Present at least 2-3 alternatives, including:
- The recommended approach
- The "do nothing" option (if applicable)
- At least one alternative considered

### Step 4: Make & Document Decision

#### Decision Criteria:
- Alignment with architectural principles
- Technical feasibility
- Team capability
- Cost-benefit ratio
- Risk tolerance
- Future flexibility

#### Document Format:
Use the lightweight template at `/docs/architecture/templates/adr-template.md`

### Step 5: ADR File Creation

#### Naming Convention:
```
ADR-YYYYMMDD-descriptive-title.md
```

Examples:
- `ADR-20240115-database-selection.md`
- `ADR-20240116-authentication-strategy.md`
- `ADR-20240117-api-versioning-approach.md`

#### File Location:
```
/docs/architecture/adrs/
```

### Step 6: ADR Content Structure

```markdown
# ADR-[YYYYMMDD]: [Clear, Descriptive Title]

## Status
[Proposed | Accepted | Deprecated | Superseded by ADR-XXX]

## Context
[2-3 paragraphs explaining:]
- The problem or opportunity
- Current situation
- Forces at play (constraints, requirements)
- Why this decision is needed now

## Decision
[1-2 paragraphs stating:]
- What we will do
- Key aspects of the solution
- High-level implementation approach

## Consequences

### Positive
- [Benefit 1]: [Brief explanation]
- [Benefit 2]: [Brief explanation]
- [Benefit 3]: [Brief explanation]

### Negative
- [Trade-off 1]: [What we're giving up]
- [Trade-off 2]: [Complexity added]
- [Trade-off 3]: [Limitation accepted]

### Risks
- [Risk 1]: [What could go wrong] → [Mitigation]
- [Risk 2]: [What could go wrong] → [Mitigation]

## Alternatives Considered

### [Alternative 1 Name]
- **Description**: [What this approach would do]
- **Pros**: [Key benefits]
- **Cons**: [Key drawbacks]
- **Rejection Reason**: [Why not chosen]

### [Alternative 2 Name]
- **Description**: [What this approach would do]
- **Pros**: [Key benefits]
- **Cons**: [Key drawbacks]
- **Rejection Reason**: [Why not chosen]

## Implementation Notes
<!-- Optional section for implementation guidance -->
- Key implementation steps
- Migration strategy (if applicable)
- Rollback plan
- Success metrics

## References
<!-- Optional section for related materials -->
- Related ADRs: [ADR-XXX]
- External resources: [Links]
- Similar implementations: [Examples]
```

### Step 7: Review Process

#### Before Marking "Accepted":
1. **Peer Review**: Have 1-2 team members review
2. **Stakeholder Input**: Confirm with affected parties
3. **Feasibility Check**: Validate with implementers
4. **Final Review**: Ensure clarity and completeness

#### Review Checklist:
- [ ] Context is clear to someone unfamiliar with the problem
- [ ] Decision directly addresses the stated problem
- [ ] Trade-offs are honestly documented
- [ ] Alternatives show genuine consideration
- [ ] Language is clear and jargon-free where possible
- [ ] Document is 1-2 pages maximum

### Step 8: Post-Creation Actions

1. **Update Status**: Start as "Proposed", move to "Accepted" after review
2. **Communicate**: Share with team via appropriate channels
3. **Index in KB**: Run `./.vector_db/kb index --path /docs/architecture/adrs/`
4. **Link References**: Update related documentation
5. **Track Implementation**: Create tickets if needed

## Common ADR Topics

### Technical Decisions:
- Database technology selection
- Framework/library choices
- API design patterns
- Authentication/authorization approach
- Caching strategy
- Message queue selection

### Architectural Patterns:
- Microservices vs Monolith
- Synchronous vs Asynchronous
- Event-driven architecture
- CQRS implementation
- Service mesh adoption

### Operational Decisions:
- Deployment strategy
- Monitoring approach
- Logging architecture
- Disaster recovery plan
- Scaling strategy

### Process Decisions:
- Development workflow
- Testing strategy
- Code review process
- Documentation standards

## Best Practices

### DO:
- Keep it concise (1-2 pages max)
- Write for future readers
- Be honest about trade-offs
- Use concrete examples
- Date everything
- Link related decisions

### DON'T:
- Combine multiple decisions
- Use excessive technical jargon
- Hide negative consequences
- Skip the alternatives section
- Modify accepted ADRs (supersede instead)
- Delete old ADRs

## ADR Lifecycle

### States:
1. **Proposed**: Under discussion
2. **Accepted**: Decision made, implementation can begin
3. **Deprecated**: No longer relevant but kept for history
4. **Superseded**: Replaced by a newer ADR

### Superseding ADRs:
When replacing an existing ADR:
1. Create new ADR with reference to old
2. Update old ADR status to "Superseded by ADR-XXX"
3. Explain why the change is needed
4. Keep both files for history

## Knowledge Base Integration

### Before Creating:
```bash
# Search for related decisions
./.vector_db/kb search "topic" --collection architecture

# Check existing ADRs
ls /docs/architecture/adrs/ | grep -i "keyword"
```

### After Creating:
```bash
# Index the new ADR
./.vector_db/kb index --path /docs/architecture/adrs/ADR-YYYYMMDD-title.md

# Verify indexing
./.vector_db/kb search "decision topic"
```

## Example ADR Creation Session

```
User: We need to decide on a caching strategy

Archie: I'll help you create an ADR for the caching strategy decision. Let me gather some context first:

1. What performance problems are you trying to solve with caching?
2. What types of data need caching (user sessions, API responses, computed values)?
3. What are your constraints (budget, complexity, team expertise)?
4. What's your current architecture (single server, distributed, cloud)?

[User provides context]

Based on your needs, I see three viable options:
1. Redis (Recommended) - Distributed, feature-rich
2. In-memory caching - Simple, no infrastructure
3. CDN caching - Good for static content

Let me create the ADR documenting this decision...

[Creates ADR-20240115-caching-strategy.md]
```

## Remember
- ADRs are for significant decisions only
- Keep them lightweight but complete
- Focus on the "why" not just the "what"
- Write for developers who join the team later
- Honest trade-offs build trust
- A decision made is better than perpetual analysis