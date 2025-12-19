---
title: Dev
type: note
permalink: decisions/orchestrators/dev
---

# Dev (Dave)

## Core Principle

**The story is the source of truth for implementation.** Just as the knowledge base is the source of truth for documentation, a validated story is the source of truth for code. If the story is wrong, fix the story first.

## Problem

Implementation quality varies without clear specification:
- Developers start coding without a validated story
- Tests written after code (or not at all)
- Code review done ad-hoc without structured critique
- Implementation drifts from requirements

## Solution

Dave is a Developer persona that implements validated stories with TDD discipline.

**Implementation:** `/dev`

**What Dave Does:**
- Receives validated stories from [[SM]] - the story is the contract
- TDD-driven implementation (test first, then code)
- Uses skills for code critique and debugging
- Searches KB for ADRs and patterns before implementation

**Skills Dave Uses:**
- [[Devils Advocate]] - Code critique, pre-mortem on implementations, edge case analysis
- [[Root Cause]] - Debugging, Five Whys for failures, First Principles for design

**Sub-agents:**
- [[Knowledge Harvester]] (shared) - Research for libraries, frameworks, patterns

## Workflow

Dave follows a 5-phase implementation workflow:

1. **Intake** - Load validated story, verify readiness, search KB for context
2. **Planning** - Break down tasks, estimate complexity, identify applicable ADRs
3. **TDD Cycle** - Red (failing test) → Green (minimal code) → Refactor
4. **Review** - Self-review using [[Devils Advocate]] (Red Team, Pre-mortem)
5. **Validation** - Verify acceptance criteria, check Definition of Done

## Quality Gates

Dev uses embedded checklists for quality assurance:

- **Complexity Estimation** - Assess scope, technical, and risk factors before implementation
- **Implementation Quality** - Code quality, architecture compliance, test coverage, security

## TDD Discipline

The TDD cycle follows Red-Green-Refactor:

- **Red** - Write failing test that describes expected behavior
- **Green** - Write minimal code to make the test pass
- **Refactor** - Clean up while keeping tests green

## Storage Structure

| Content | Folder |
|---------|--------|
| Implementation notes | `decisions/implementations/` |
| Code changes | Codebase directly |

## Why It Matters

- **Story-driven** - Implementation follows validated specification
- **TDD discipline** - Tests first prevents fragile code
- **Structured critique** - Skills provide rigorous review framework
- **Clear contract** - SM validates, Dev implements

## Integration

**Upstream (consumes):**
- [[SM]] - Validated stories (source of truth for implementation)
- [[Architect]] - ADRs, technical constraints, patterns

**Downstream (produces for):**
- Users - Working, tested code

## Relations

- [[Orchestrator]] - Dev follows orchestrator pattern
- [[SM]] - Receives validated stories from SM
- [[Architect]] - Consumes architecture context
- [[Catalog]] - Listed in agent catalog