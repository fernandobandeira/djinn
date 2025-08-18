# Generate AI Frontend Prompt Task

## Purpose
Generate comprehensive, optimized prompts for AI-driven frontend development tools (v0, Lovable, Cursor, etc.) based on UX specifications and architecture documents.

## Prerequisites
- Completed UI/UX Specification (`/docs/ux/design/ui-spec.md`)
- Frontend Architecture (`/docs/architecture/frontend-architecture.md`) or Full Stack Architecture
- Design System documentation if available
- Component library specifications

## Core Prompting Principles

### 1. Structured Framework
Every prompt must follow this four-part structure:

1. **High-Level Goal**: Clear, concise objective summary
2. **Detailed Instructions**: Granular, numbered steps
3. **Code Examples & Constraints**: Concrete examples and boundaries
4. **Strict Scope Definition**: What to modify and what to leave untouched

### 2. Key Strategies
- **Be Explicit**: AI needs complete context and details
- **Mobile-First**: Always describe mobile layouts first, then tablet/desktop
- **Iterate**: Build complex UIs component by component
- **Provide Context**: Include tech stack, existing patterns, design tokens

## Workflow Steps

### Step 1: Gather Context
```markdown
**Project Information:**
- Project Name: [from PRD/Brief]
- Tech Stack: [from architecture]
- UI Framework: [React/Vue/Svelte/etc.]
- Styling: [Tailwind/CSS Modules/Styled Components]
- Component Library: [if any]
- Design System: [link or description]
```

### Step 2: Define Visual Style
```markdown
**Visual Context:**
- Design Files: [Figma/Sketch links]
- Color Palette: [primary, secondary, accent, semantic colors]
- Typography: [font families, scale]
- Spacing System: [8px grid, etc.]
- Overall Aesthetic: [minimal/corporate/playful]
```

### Step 3: Structure Component Request

#### For Single Component:
```markdown
## Component: [ComponentName]

**Goal:** Create a [description] component with [key features]

**Detailed Instructions:**
1. Create new file: `components/[ComponentName].[ext]`
2. Import required dependencies: [list]
3. Define props interface/types:
   - prop1: type - description
   - prop2: type - description
4. Implement component structure:
   - Mobile layout (< 768px): [description]
   - Tablet layout (768-1024px): [description]  
   - Desktop layout (> 1024px): [description]
5. Add state management for: [interactions]
6. Implement accessibility:
   - ARIA labels for [elements]
   - Keyboard navigation for [interactions]
   - Focus management for [flows]
7. Add error states: [list scenarios]
8. Include loading states: [describe]

**Code Context:**
```typescript
// Existing types/interfaces to use:
interface User {
  id: string;
  name: string;
  email: string;
}

// API endpoint:
POST /api/[endpoint]
Request: { ... }
Response: { ... }
```

**Constraints:**
- Use existing [DesignToken] system
- Follow [ComponentLibrary] patterns
- Do NOT modify [ExistingComponent]
- Do NOT change global styles
- Maximum bundle size: [limit]

**Scope:**
- CREATE: `components/[ComponentName].[ext]`
- UPDATE: `pages/[page].[ext]` to import and use component
- IGNORE: All other files
```

#### For Full Page/Flow:
```markdown
## Page: [PageName]

**Goal:** Create [page description] with [key features]

**Information Architecture:**
- Header: [components]
- Main Content:
  - Section 1: [description]
  - Section 2: [description]
- Sidebar/Secondary: [if applicable]
- Footer: [components]

**User Flow:**
1. Entry: User arrives from [source]
2. Initial State: [description]
3. Interactions:
   - Action 1: [trigger] → [result]
   - Action 2: [trigger] → [result]
4. Success State: [description]
5. Error Handling: [scenarios]

[Rest follows component structure above]
```

### Step 4: Add Advanced Requirements

```markdown
**Performance Requirements:**
- Lazy load: [components/images]
- Code split: [boundaries]
- Optimize: [specific areas]

**Accessibility Requirements:**
- Screen reader: [announcements]
- Keyboard: [navigation pattern]
- Focus: [management strategy]

**SEO Requirements:**
- Meta tags: [list]
- Structured data: [if applicable]
- Open Graph: [properties]

**Analytics Requirements:**
- Track events: [list with names]
- Page views: [configuration]
- Custom dimensions: [if any]
```

### Step 5: Format Final Prompt

```markdown
# AI Frontend Generation Prompt

## Project Context
[Include gathered context from Step 1]

## Visual Design
[Include style guide from Step 2]

## Component/Page Specification
[Include structured request from Step 3]

## Technical Requirements
[Include advanced requirements from Step 4]

## File Structure
```
src/
  components/
    [ComponentName]/
      index.[ext]
      styles.[ext]
      types.[ext]
      [ComponentName].test.[ext]
  pages/
    [page].[ext]
  styles/
    [relevant files]
```

## Quality Checklist
Before accepting generated code, verify:
- [ ] Responsive across all breakpoints
- [ ] Accessible (ARIA, keyboard, focus)
- [ ] Error states handled
- [ ] Loading states included
- [ ] Type-safe (if TypeScript)
- [ ] Follows project conventions
- [ ] Performance optimized
- [ ] SEO friendly (if applicable)

## Iteration Instructions
After initial generation:
1. Test responsiveness at each breakpoint
2. Verify accessibility with screen reader
3. Check all interactive states
4. Validate API integration
5. Review with design team

## Important Notes
- Generated code requires human review
- Test thoroughly before production
- Ensure security best practices
- Validate against design system
- Check browser compatibility
```

## Output Format

Save the generated prompt as:
- `/docs/ux/prompts/[feature]-ai-prompt.md`

Include metadata:
```yaml
---
generated: [date]
tool: [v0/lovable/cursor]
feature: [feature name]
status: draft|reviewed|approved
---
```

## Success Criteria
- Prompt produces 80%+ accurate initial code
- Minimal iterations needed for refinement
- Code follows project conventions
- Accessible and responsive by default
- Performance requirements met

## Common Pitfalls to Avoid
- Vague descriptions → Be specific
- Missing context → Include all dependencies
- No constraints → Define boundaries clearly
- Single attempt → Plan for iteration
- Desktop-first → Always mobile-first
- Ignoring errors → Include all states
- No accessibility → Build in from start