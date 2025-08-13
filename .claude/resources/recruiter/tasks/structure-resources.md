# Structure Resources Task

## Purpose
Organize agent resources into a clean, maintainable file structure that supports lazy loading and clear separation of concerns.

## Resource Categories

### Tasks (Complex Workflows)
**Purpose**: Multi-step processes that require guidance
**Location**: `.claude/resources/[agent]/tasks/`
**Examples**:
- Brainstorming facilitation
- System design workflow
- Analysis procedures
- Creation processes

### Templates (Output Formats)
**Purpose**: Standardized structures for agent outputs
**Location**: `.claude/resources/[agent]/templates/`
**Examples**:
- Report templates
- Document structures
- Analysis frameworks
- Decision records

### Data (Reference Information)
**Purpose**: Static information agent needs to reference
**Location**: `.claude/resources/[agent]/data/`
**Examples**:
- Best practices
- Pattern libraries
- Technique catalogs
- Domain knowledge

### Checklists (Validation)
**Purpose**: Quality assurance and completeness checks
**Location**: `.claude/resources/[agent]/checklists/`
**Examples**:
- Review criteria
- Completion checks
- Quality metrics
- Prerequisites

## File Structure Creation

### Step 1: Determine Resource Needs
Ask for each category:
1. What multi-step processes will the agent guide?
2. What standardized outputs will it produce?
3. What reference information does it need?
4. What validations should it perform?

### Step 2: Create Directory Structure
```bash
# For command agents
mkdir -p .claude/resources/[agent-name]/tasks
mkdir -p .claude/resources/[agent-name]/templates
mkdir -p .claude/resources/[agent-name]/data
mkdir -p .claude/resources/[agent-name]/checklists

# Verify structure
tree .claude/resources/[agent-name]/
```

### Step 3: Plan Resource Files
Map each need to a specific file:
```yaml
resources:
  tasks:
    - primary_workflow.md
    - secondary_process.md
  templates:
    - output_template.md
    - report_format.md
  data:
    - reference_guide.md
    - patterns.md
  checklists:
    - quality_check.md
    - completion.md
```

## Resource File Guidelines

### Task Files
```markdown
# [Task Name] Task

## Purpose
[What this workflow accomplishes]

## Prerequisites
- [Required conditions]
- [Necessary inputs]

## Steps
1. [First step]
   - Sub-step
   - Sub-step
2. [Second step]
   - Sub-step
   - Sub-step

## Output
[What gets produced]
```

### Template Files
```markdown
# [Template Name] Template

## Purpose
[When to use this template]

## Structure
[Template format with placeholders]

## Example
[Filled example]

## Guidelines
- [Usage note]
- [Formatting rule]
```

### Data Files
```markdown
# [Reference Name]

## Overview
[What this reference contains]

## Categories
### [Category 1]
- [Item]
- [Item]

### [Category 2]
- [Item]
- [Item]

## Usage
[How to apply this information]
```

### Checklist Files
```markdown
# [Checklist Name]

## Purpose
[What this validates]

## Checks
- [ ] [Check item]
- [ ] [Check item]
- [ ] [Check item]

## Success Criteria
[What constitutes completion]
```

## Lazy Loading Implementation

### In Command File
```yaml
# Reference resources but don't load
resource_files:
  tasks:
    main_task: .claude/resources/agent/tasks/main.md
  templates:
    output: .claude/resources/agent/templates/output.md
```

### In Execution
```markdown
When user requests `*command`:
1. THEN load: `.claude/resources/agent/tasks/main.md`
2. Execute task
```

## Organization Patterns

### By Frequency
Most used → least used:
```
tasks/
├── core-workflow.md      # Primary task
├── common-process.md     # Frequent task
└── edge-case.md         # Rare task
```

### By Complexity
Simple → complex:
```
tasks/
├── quick-check.md       # Simple
├── standard-flow.md     # Medium
└── advanced-analysis.md # Complex
```

### By Domain
Group related resources:
```
data/
├── security/
│   ├── vulnerabilities.md
│   └── best-practices.md
└── performance/
    ├── metrics.md
    └── optimization.md
```

## Resource Sizing Guidelines

### Keep Files Focused
- One task per file
- One template per file
- Logical grouping for data
- Single-purpose checklists

### File Size Targets
- Tasks: 100-500 lines
- Templates: 50-200 lines
- Data: 200-1000 lines
- Checklists: 20-100 items

### When to Split
- File exceeds 500 lines
- Multiple unrelated sections
- Different update frequencies
- Distinct user groups

## Maintenance Considerations

### Versioning
- Date in filename if needed: `task-20250112.md`
- Deprecation notices in old files
- Migration guides when changing

### Documentation
- README in resource directory
- Purpose statement in each file
- Cross-references where needed
- Update history if complex

### Testing
- Validate file paths work
- Check lazy loading triggers
- Ensure complete coverage
- Test resource accessibility

## Common Resource Patterns

### Analysis Agent Resources
```
resources/analyst/
├── tasks/
│   ├── brainstorm.md
│   ├── research.md
│   └── synthesize.md
├── templates/
│   ├── report.md
│   └── brief.md
├── data/
│   ├── techniques.md
│   └── frameworks.md
└── checklists/
    └── completeness.md
```

### Architecture Agent Resources
```
resources/architect/
├── tasks/
│   ├── design-system.md
│   └── review-architecture.md
├── templates/
│   ├── adr.md
│   └── diagram.md
├── data/
│   ├── patterns.md
│   └── principles.md
└── checklists/
    └── design-review.md
```

### Utility Agent Resources
```
resources/formatter/
├── tasks/
│   └── format-code.md
├── data/
│   └── style-guides.md
└── checklists/
    └── formatting.md
```

## Quality Checks

Before finalizing:
- [ ] All directories created
- [ ] Resources properly categorized
- [ ] No duplicate functionality
- [ ] Clear file naming
- [ ] Lazy loading implemented
- [ ] Paths verified
- [ ] Documentation complete

## Remember
- Structure enables maintainability
- Categories clarify purpose
- Lazy loading saves context
- Organization aids discovery
- Consistency reduces confusion