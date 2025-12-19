---
title: Checklists
type: note
permalink: patterns/checklists
---

# Checklists Pattern

## Core Principle

**Checklists belong embedded in the entity that uses them.** Unlike templates (which are platform-agnostic and shared), checklists are operational verification tools tightly coupled to specific workflows.

## Checklists vs Templates

| Aspect | Templates | Checklists |
|--------|-----------|------------|
| **Purpose** | Create artifacts (PRDs, ADRs, etc.) | Verify work during workflows |
| **Location** | `templates/` (separate, reusable) | Embedded in orchestrator/command |
| **Scope** | Platform-agnostic, shared | Workflow-specific, single owner |
| **Usage** | Fill structure to create | Check items during execution |

## Why Embed Checklists?

1. **Single source of truth** - Checklist lives where it's used
2. **No context switching** - Orchestrator has everything it needs inline
3. **Clear ownership** - Each workflow owns its verification criteria
4. **Easier maintenance** - Update workflow and checklist together
5. **Reduced complexity** - No separate resources directory to manage

## Where to Embed

| If checklist is used by... | Embed in... |
|---------------------------|-------------|
| Single orchestrator/command | That orchestrator file |
| Multiple orchestrators | Shared skill (used by all) |
| Workflow step | The workflow section that invokes it |

**Most common**: Orchestrators contain their checklists inline since checklists are typically workflow-specific.

## Embedding Structure

Embed checklists as a dedicated section in the orchestrator:

```markdown
## Checklists

### {Checklist Name}

#### {Section}
- [ ] Item 1
- [ ] Item 2

#### {Another Section}
- [ ] Item 3
- [ ] Item 4
```

Or inline within a workflow that uses it:

```markdown
### *review-something

1. Gather context
2. **Run checklist:**
   - [ ] Check A verified
   - [ ] Check B verified
   - [ ] Check C verified
3. Present findings
```

## Anti-Patterns

❌ **Separate resources directory** - Creates indirection, harder to maintain
❌ **Checklist shared across many orchestrators** - Should be a skill instead
❌ **Checklist without owner** - Every checklist needs exactly one home

## Migration Path

If you have checklists in separate files:

1. Identify the single orchestrator/command that uses each checklist
2. Copy checklist content into that orchestrator's `## Checklists` section
3. Remove the separate file
4. Update any references

If multiple orchestrators share a checklist:
- Evaluate if it should be a skill (teaches HOW to verify)
- Or embed in the primary user and let others reference that orchestrator

## Relations

- [[Skill]] - Skills teach thinking; if verification needs structured reasoning, make it a skill
- [[Orchestrator]] - Most checklists belong in orchestrators
- [[Templates]] - Templates are separate (platform-agnostic), checklists are embedded