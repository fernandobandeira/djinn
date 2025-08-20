# Task: Create Product Requirements Document (PRD)

## Workflow
1. Research Collection
   - Load market research:
     ```bash
     ./.vector_db/kb search "market" --collection documentation
     ./.vector_db/kb search "user needs" --collection documentation
     ```
   - Extract insights from `/docs/analysis/`
   - Review UX research from `/docs/research/user/`

2. PRD Generation
   - Use template: `.claude/resources/pm/templates/prd-template.md`
   - Follow workflow: `.claude/resources/pm/protocols/molecules/prd-creation-workflow.md`

3. Output Management
   - Write PRD to `/docs/requirements/prd.md`
   - Validate with checklist: `.claude/resources/pm/checklists/prd-quality-checklist.md`

4. Knowledge Base Indexing
   ```bash
   ./.vector_db/kb index --path /docs/requirements/prd.md
   ```

## Required Tools
- Ana (Analyst) for market insights
- UX Research team data
- Archie (Architect) for technical feasibility
- Stakeholder input

## Success Criteria
- Comprehensive market analysis
- Clear user needs definition
- Technically feasible requirements
- Stakeholder alignment
- Knowledge base indexed