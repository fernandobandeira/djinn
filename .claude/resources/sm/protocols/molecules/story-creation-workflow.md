# Story Creation Workflow

## Purpose
Create comprehensive, development-ready user stories following BMAD methodology with complete context from PRD and Architecture.

## Prerequisites
- PRD exists at `/docs/requirements/prd.md`
- Architecture documents exist in `/docs/architecture/`
- Previous stories (if any) in `/docs/stories/`

## Workflow Steps

### 1. Gather Context
```bash
# Search KB for relevant documents
./.vector_db/kb search "PRD requirements" --collection documentation
./.vector_db/kb search "architecture technical" --collection architecture
./.vector_db/kb search "previous stories" --collection documentation

# Load specific documents
Read /docs/requirements/prd.md
Read /docs/architecture/architecture.md
Read /docs/architecture/tech-stack.md
Read /docs/architecture/api-spec.md
Read /docs/architecture/security.md
Read /docs/architecture/testing-strategy.md
```

### 2. Identify Story from Epic
- Review PRD for epic definitions
- Select next story in sequence (e.g., 2.1, 2.2, etc.)
- Verify dependencies from previous stories

### 3. Define Story Components

#### Story Statement
Use format:
```
**As a** [user role],
**I want** [specific action],
**so that** [business value/benefit].
```

#### Acceptance Criteria
- Minimum 5-7 specific, testable criteria
- Each criterion should be verifiable
- Include UI, API, and data requirements
- Reference specific screens or endpoints

#### Tasks and Subtasks
Break down into implementation tasks:
```markdown
- [ ] Task 1: [High-level task] (AC: [relevant AC numbers])
  - [ ] Subtask 1.1: [Specific implementation step]
  - [ ] Subtask 1.2: [Specific implementation step]
```

### 4. Populate Dev Notes Section

#### Previous Story Insights
- Check `/docs/stories/` for completed stories in same epic
- Extract relevant implementation details
- Note any technical decisions or patterns established

#### Architecture Context
Include specific sections with source references:

**Technology Stack** [Source: /docs/architecture/tech-stack.md]
- List relevant technologies for this story
- Include version numbers
- Note any specific packages needed

**Project Structure** [Source: /docs/architecture/source-tree.md]
- Show exact directory structure for new features
- Include file naming conventions
- List all files to be created/modified

**API Integration** [Source: /docs/architecture/api-spec.md]
- Include endpoint specifications
- Show request/response formats
- Add authentication requirements

**Security Considerations** [Source: /docs/architecture/security.md]
- List security requirements
- Include validation rules
- Note authentication/authorization needs

#### Code Examples
Provide implementation guidance:
```dart
// Example state management
final storyProvider = StateNotifierProvider<StoryNotifier, StoryState>((ref) {
  return StoryNotifier(ref.read(repositoryProvider));
});
```

### 5. Define Testing Requirements

Based on `/docs/architecture/testing-strategy.md`:
- Unit test requirements
- Widget test requirements
- Integration test requirements
- Specific test scenarios for this story

### 6. Quality Checks

Before finalizing:
- [ ] All ACs are testable and specific
- [ ] Tasks cover all ACs
- [ ] Dev notes include all necessary context
- [ ] No need to read PRD/Architecture during development
- [ ] Previous story insights captured
- [ ] Security considerations addressed
- [ ] Testing requirements comprehensive

### 7. Output Story

Save to: `/docs/stories/[epic].[story].story.md`

Example: `/docs/stories/2.1.story.md`

### 8. Index with KB

After creation:
```bash
./.vector_db/kb index --path /docs/stories/
```

## Story Quality Metrics

A well-formed story should:
- Be 500+ lines with complete context
- Include 15+ specific subtasks
- Reference 5+ architecture documents
- Provide 10+ code examples
- Define 10+ test cases
- Have zero ambiguity for developers

## Example Output Structure

See `/docs/stories/2.2.story.md` for reference implementation showing:
- Complete architecture context
- Detailed security section
- Comprehensive testing requirements
- Full project structure
- Code examples for complex parts
- Previous story insights