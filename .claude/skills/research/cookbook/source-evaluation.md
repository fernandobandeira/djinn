# Source Evaluation Cookbook

Assess relevance, quality, and applicability of discovered knowledge.

## When to Use

- After KB search returns results
- After harvesting external content
- When synthesizing multiple sources
- Before citing or building on findings

## Evaluation Framework

### 1. Relevance Assessment

**Score interpretation:**
| Score | Meaning | Action |
|-------|---------|--------|
| 0.8+ | Highly relevant | Read full document, likely useful |
| 0.6-0.8 | Moderately relevant | Skim for specific sections |
| 0.4-0.6 | Tangentially related | Quick scan, may have insights |
| < 0.4 | Low relevance | Skip unless desperate |

**Questions to ask:**
- Does this directly address my need?
- Is the context similar to my situation?
- Are the assumptions still valid?

### 2. Source Type Priority

| Source Type | Priority | Rationale |
|-------------|----------|-----------|
| ADRs (internal) | Highest | Project decisions, must follow |
| Patterns (internal) | High | Established project approaches |
| Documentation (internal) | High | Current project state |
| Code (internal) | Medium | Implementation examples |
| Harvested (external) | Lower | May not fit project context |
| General web | Lowest | Unverified, may be outdated |

**Internal vs External:**
- Internal sources understand project context
- External sources need adaptation
- Always prefer internal when available

### 3. Freshness Check

**For internal docs:**
- Check file modification date
- Look for version/date in frontmatter
- Verify against current codebase

**For external docs:**
- Check publication/update date
- Verify library/framework version matches
- Look for deprecation notices

### 4. Authority Assessment

**High authority indicators:**
- Official documentation
- Established project ADRs
- Well-known pattern catalogs
- Peer-reviewed content

**Lower authority indicators:**
- Blog posts (variable quality)
- Forum answers (may be outdated)
- Generated/AI content
- Anonymous sources

## Synthesis Process

### Step 1: Gather
Collect all relevant search results:
```bash
./.vector_db/kb search "topic" --limit 10
```

### Step 2: Categorize
Group by source type:
- Internal ADRs/patterns
- Internal documentation
- External harvested
- Code examples

### Step 3: Prioritize
Rank by:
1. Relevance score
2. Source authority
3. Freshness
4. Applicability

### Step 4: Reconcile
Handle conflicts:
- Internal trumps external
- Recent trumps old (usually)
- ADRs are authoritative decisions
- Document contradictions

### Step 5: Synthesize
Combine findings:
- Extract key insights
- Note applicable patterns
- Identify gaps
- Formulate recommendations

## Red Flags

**Discard or verify if:**
- Score below 0.4
- Source older than 2 years (tech topics)
- Contradicts project ADRs
- From unknown/untrusted source
- Specific to different tech stack
- Contains obvious errors

## Output Template

After evaluation, summarize:

```markdown
## Research Summary: [Topic]

### Key Findings
1. [Finding from highest-priority source]
2. [Finding from second source]
3. [Finding from third source]

### Sources Used
- [Source 1]: Score X, [internal/external], [date]
- [Source 2]: Score X, [internal/external], [date]

### Gaps Identified
- [What wasn't found]
- [What needs external research]

### Recommendations
- [What to do next based on findings]
```

## Decision Matrix

| Situation | Action |
|-----------|--------|
| Found relevant ADR | Follow it, reference in new work |
| Found relevant pattern | Adapt and apply |
| Multiple conflicting sources | Prefer internal, recent, authoritative |
| No relevant results | Harvest externally or create new |
| External-only results | Evaluate fit, adapt carefully |

## Integration with Other Skills

After evaluation, you may need:
- `ideation` - Generate solutions based on findings
- `root-cause` - Dig deeper into problems found
- `strategic-analysis` - Assess options discovered
- `devils-advocate` - Challenge the findings
