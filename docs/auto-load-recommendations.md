# Auto-Load Recommendations for Command Agents

## Critical Resources That Should Use @ Syntax

### 1. Teacher (`/teacher`)
**Recommendation**: Auto-load Zettelkasten coordination
```markdown
@.claude/resources/teacher/cognitive-tools/programs/CoordinateWithZettelkasten.md
```
**Reason**: Teacher always needs to coordinate with Zettelkasten for knowledge capture.

### 2. Architect (`/architect`)
**Recommendation**: Auto-load task routing for delegation
```markdown
@.claude/resources/architect/orchestration/task-routing.yaml
```
**Reason**: Architect constantly delegates to sub-agents and needs routing rules.

### 3. PM (`/pm`)
**Recommendation**: Auto-load elicitation methods
```markdown
@.claude/resources/pm/data/elicitation-methods.md
```
**Reason**: PM uses elicitation in almost every interaction.

### 4. SM (`/sm`)
**Recommendation**: Auto-load story creation workflow
```markdown
@.claude/resources/sm/protocols/molecules/story-creation-workflow.md
@.claude/resources/sm/data/output-paths.yaml
```
**Reason**: Story creation is SM's primary function and needs consistent workflow.

### 5. Dev (`/dev`)
**Recommendation**: Auto-load implementation workflow
```markdown
@.claude/resources/dev/protocols/molecules/implementation-workflow.md
@.claude/resources/dev/orchestration/workflow-state.yaml
```
**Reason**: Dev needs workflow state management for all implementations.

### 6. Analyst (`/analyst`)
**Recommendation**: Auto-load elicitation methods (shared with PM)
```markdown
@.claude/resources/analyst/data/elicitation-methods.md
```
**Reason**: Analyst uses advanced elicitation throughout research.

### 7. UX (`/ux`)
**Recommendation**: Auto-load advanced elicitation protocol
```markdown
@.claude/resources/ux/protocols/advanced-elicitation.md
```
**Reason**: UX research requires consistent elicitation methodology.

## Implementation Pattern

### Before (Using THEN load):
```bash
## Resource Loading Protocol
When executing commands:
```bash
THEN load .claude/resources/command/critical-resource.md
```

### After (Using @ auto-load):
```bash
## Resource Loading Protocol
**AUTO-LOADED ON ACTIVATION:**
@.claude/resources/command/critical-resource.md

Additional resources loaded contextually:
```bash
THEN load .claude/resources/command/contextual-resource.md
```

## Guidelines for @ Usage

1. **Use @ for resources that are ALWAYS needed**
   - Core workflows
   - Delegation mappings
   - Orchestration rules
   - Shared methodologies

2. **Continue using THEN load for contextual resources**
   - Command-specific templates
   - Situational checklists
   - Optional cognitive tools

3. **Benefits of @ auto-loading**
   - Ensures critical resources are always available
   - Reduces cognitive overhead
   - Maintains consistency
   - Prevents forgetting to load essential files

## Priority Order for Implementation

1. **HIGH**: Recruiter (DONE) - Already uses @ for decomposition mapping
2. **HIGH**: Teacher - Needs Zettelkasten coordination
3. **HIGH**: SM - Story creation workflow is critical
4. **MEDIUM**: Dev - Implementation workflow for consistency
5. **MEDIUM**: Architect - Task routing for delegation
6. **LOW**: PM, Analyst, UX - Elicitation methods are important but contextual

## Notes

- The `@` syntax appears to be processed at agent activation
- Only use for truly essential resources
- Keep auto-loaded resources minimal to avoid bloat
- Most resources should remain contextually loaded with THEN load