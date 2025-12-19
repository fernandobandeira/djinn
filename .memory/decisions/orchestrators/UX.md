---
title: UX
type: note
permalink: decisions/orchestrators/ux
---

# UX (Ulysses)

## Core Principle

**User-Centered Design** - Every decision serves user needs first. Users' stated preferences often differ from their actual behavior. Ulysses challenges assumptions with evidence, validates designs against real user needs, and ensures what we build actually helps people.

## Problem

Users need UX support that:
- Gathers real user insights through structured research methods
- Creates personas and journey maps based on evidence, not assumptions
- Validates designs against usability heuristics and accessibility standards
- Produces frontend specifications for implementation handoff
- Challenges design decisions with user behavior data

Without structured UX support:
- Designs are based on assumptions, not evidence
- Usability issues discovered late (expensive to fix)
- Accessibility barriers exclude users
- Frontend specs are incomplete or inconsistent
- User research is ad-hoc or skipped entirely

## Solution

An orchestrator that:
- Uses `user-research` skill for research methodology (10 techniques)
- Uses `devils-advocate` skill for design critique and validation
- Uses `role-playing` skill for stakeholder perspectives
- Uses `ideation` skill for design alternatives
- Produces personas, journey maps, frontend specs
- Generates diagrams directly (no sub-agent - avoids context handoff loss)
- Delegates to `knowledge-harvester` only for heavy external research

**Implementation:** `/ux`

**What Ulysses Does:**
- **User Research** - Journey mapping, interviews, surveys, card sorting
- **Elicitation** - Requirements gathering, user story mapping
- **Design** - Wireframes, prototypes, frontend specifications
- **Validation** - Usability testing, accessibility audits, design critique
- Creates user personas with evidence-based insights
- Documents frontend specs for developer handoff

## Commands

| Command | Purpose |
|---------|---------|
| `*research {topic}` | User research session |
| `*personas` | Create user personas |
| `*journeys` | Map user journeys |
| `*elicit` | Advanced elicitation session |
| `*design {feature}` | Create wireframes/prototypes |
| `*specs` | Generate frontend specification |
| `*diagram {type}` | Generate UX diagram |
| `*validate` | Usability validation |
| `*accessibility` | Accessibility audit |
| `*critique` | Structured design critique |
| `*harvest {topic}` | External UX research |

## Skills Used

- [[User Research]] - Journey mapping, interviews, surveys, card sorting, usability testing, accessibility audits
- [[Devils Advocate]] - Design critique, pre-mortem, heuristic evaluation
- [[Role Playing]] - Six Hats, stakeholder perspectives
- [[Ideation]] - Design alternatives brainstorming

*Ulysses generates diagrams directly (no sub-agent needed - avoids accuracy loss from context handoff).*

## Sub-agents Used

Only for heavy I/O:
- [[Knowledge Harvester]] - External UX trends, accessibility standards, design patterns

## Workflow

Ulysses follows three interconnected flows:

**User Research Flow:**
1. Search KB for existing research
2. Identify research gaps
3. Conduct research (interviews, surveys, observation)
4. Create personas and journey maps
5. Store findings

**Design Flow:**
1. Load user context (personas, journeys)
2. Apply design principles
3. Create wireframes/specs
4. Validate against heuristics

**Validation Flow:**
1. Apply Nielsen's 10 Heuristics
2. Check accessibility (WCAG)
3. Review user flows
4. Document issues with severity

## Storage Structure

| Content | Folder |
|---------|--------|
| User personas | `research/user/` |
| Journey maps | `research/user/` |
| Usability findings | `research/user/` |
| Accessibility audits | `research/user/` |
| Frontend specs | `research/frontend-specs/` |

## Validation Criteria

**Nielsen's 10 Usability Heuristics:**
1. Visibility of system status
2. Match between system and real world
3. User control and freedom
4. Consistency and standards
5. Error prevention
6. Recognition rather than recall
7. Flexibility and efficiency of use
8. Aesthetic and minimalist design
9. Help users recognize/recover from errors
10. Help and documentation

## Integration

### Upstream (UX Consumes)
- [[Analyst]] - Market research, competitive analysis, project briefs

### Downstream (UX Produces)
- [[PM]] - Personas, journey maps, frontend specs
- [[Architect]] - UX requirements, accessibility needs
- Developers - Frontend specifications for implementation

## Templates

Location: `{templates}/ux/` (configurable per [[Templates]] pattern)

| Template | Purpose | Key Sections |
|----------|---------|--------------|
| persona-template.md | User persona documentation | Demographics, Goals, Pain Points, Behaviors, Design Implications |
| frontend-spec-template.md | Frontend specifications | IA, User Flows, Components, Visual Design, Accessibility, Performance |

See [[Templates]] pattern for the platform-agnostic template approach.

## Why It Matters

- **Evidence over assumptions** - Designs grounded in real user behavior
- **Early validation** - Usability issues caught before expensive fixes
- **Inclusive design** - Accessibility built in, not bolted on
- **Clear handoffs** - Frontend specs enable consistent implementation
- **Structured methods** - 10 research techniques for any situation
- **Integrated workflow** - Connects Analyst research to PM requirements

## Relations

- [[Architecture]] - Part of Djinn orchestrator layer
- [[Catalog]] - Listed in component catalog
- [[Orchestrator]] - Ulysses follows this pattern (guides users, uses skills, delegates only for heavy I/O)
- [[Templates]] - Uses UX templates
- [[User Research]] - Primary skill for research methods
