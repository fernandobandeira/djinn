# Task Template

## Standard Task Structure

```markdown
# [Task Name] Task

## Purpose
[Single sentence describing what this task accomplishes]

## Context
[Optional: Background information or why this task exists]

## Prerequisites
- [ ] [Required condition or resource]
- [ ] [Required knowledge or tool]
- [ ] [Required input or state]

## Inputs
- **[Input Name]**: [Description and format]
- **[Input Name]**: [Description and format]

## Process

### Step 1: [Action Verb + Object]
[Detailed description of what to do]

**Sub-steps:**
1. [Specific action]
2. [Specific action]
3. [Specific action]

**Example:**
```
[Code or command example if applicable]
```

**Expected Result:**
[What should happen after this step]

### Step 2: [Action Verb + Object]
[Detailed description of what to do]

**Decision Point:**
- If [condition], then [action]
- If [condition], then [alternative action]

### Step 3: [Action Verb + Object]
[Detailed description of what to do]

**Common Issues:**
- **Problem**: [Description]
  **Solution**: [How to fix]

## Outputs
- **[Output Name]**: [Description and format]
- **[Output Name]**: [Description and format]

## Success Criteria
- [ ] [Measurable outcome]
- [ ] [Quality metric]
- [ ] [Completion indicator]

## Error Handling
| Error | Cause | Resolution |
|-------|-------|------------|
| [Error message/type] | [Why it happens] | [How to fix] |
| [Error message/type] | [Why it happens] | [How to fix] |

## Best Practices
- [Recommendation for optimal execution]
- [Common pitfall to avoid]
- [Efficiency tip]

## Related Tasks
- [Link to prerequisite task]
- [Link to follow-up task]
- [Link to alternative approach]

## Examples

### Example 1: [Scenario Name]
**Input:**
[Sample input data]

**Process:**
[Step-by-step walkthrough]

**Output:**
[Expected result]

### Example 2: [Edge Case]
**Input:**
[Unusual input]

**Process:**
[How to handle]

**Output:**
[Expected result]

## Notes
[Optional: Additional context, warnings, or tips]

## Version History
- [Date]: [Change description]
- [Date]: Initial version
```

## Task Types and Variations

### Workflow Task Template
```markdown
# [Workflow Name] Workflow

## Purpose
Guide user through [process name] from start to finish

## Workflow Diagram
```
[Start] → [Step 1] → [Step 2] → [Decision] → [Step 3] → [End]
                            ↓
                      [Alternative Path]
```

## Stages

### Stage 1: [Preparation]
- Duration: [Estimated time]
- Owner: [Role responsible]
- Actions: [List of actions]

### Stage 2: [Execution]
...

### Stage 3: [Validation]
...

## Checkpoints
- [ ] After Stage 1: [Validation]
- [ ] After Stage 2: [Validation]
- [ ] Final: [Validation]
```

### Analysis Task Template
```markdown
# [Analysis Type] Analysis Task

## Purpose
Analyze [subject] to determine [outcome]

## Analysis Framework

### Data Collection
1. Gather [data type]
2. Collect [data type]
3. Compile [data type]

### Analysis Method
- **Technique**: [Method name]
- **Tools**: [Required tools]
- **Duration**: [Time estimate]

### Evaluation Criteria
- [Metric 1]: [Threshold]
- [Metric 2]: [Threshold]

## Interpretation Guide
| Finding | Meaning | Action |
|---------|---------|--------|
| [Result] | [Interpretation] | [Next step] |

## Report Template
See: [Link to report template]
```

### Creation Task Template
```markdown
# Create [Asset Type] Task

## Purpose
Create a new [asset] following established patterns

## Creation Process

### Step 1: Gather Requirements
Questions to ask:
1. [Question]
2. [Question]

### Step 2: Check Existing
Commands to run:
```bash
[Discovery commands]
```

### Step 3: Generate
Template to follow:
[Link to template]

### Step 4: Validate
Checklist to verify:
[Link to checklist]

### Step 5: Deploy
Location: [Where to save]
Naming: [Convention to follow]

## Quality Standards
- Must include: [Requirements]
- Should follow: [Guidelines]
- Must not: [Restrictions]
```

## Usage Guidelines

### When to Create a Task
- Process has 3+ steps
- Multiple decision points
- Requires specific order
- Common repetitive workflow
- Complex coordination needed

### Task Granularity
- **Too Big**: Entire project lifecycle
- **Just Right**: Specific workflow/process
- **Too Small**: Single command execution

### Naming Conventions
- Use verb-noun format: `create-agent.md`
- Be specific: `analyze-security.md` not `analyze.md`
- Include type if ambiguous: `review-code-task.md`

## Remember
- Tasks guide execution
- Include decision points
- Provide examples
- Handle errors gracefully
- Link related resources
- Keep steps actionable
- Test the workflow