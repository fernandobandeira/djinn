# Create User Story Task

## Workflow
1. Load PRD Context
   - Source: /docs/requirements/prd.md
   - Extract relevant epic and requirements

2. Architecture Alignment
   - Review: /docs/architecture/
   - Ensure technical feasibility
   - Map to system constraints

3. Story Generation
   - Template: .claude/resources/sm/templates/story-template.yaml
   - Output Path: /docs/stories/{epic}.{story}.md
   
4. Quality Assurance
   - Reviewer: qa-reviewer
   - Generate Acceptance Criteria
   - Validate against checklist

5. Post-Creation Actions
   - Index with KB
   - Update Backlog
   - Notify Team

## Implementation
```bash
# Pseudocode workflow
load_prd
extract_requirements
validate_architecture
generate_story
review_story
index_story
update_backlog
```

## Validation Checks
- [ ] PRD context loaded
- [ ] Architecture constraints met
- [ ] Story template followed
- [ ] Acceptance criteria generated
- [ ] KB indexing completed
- [ ] Backlog updated