---
title: Project
type: note
permalink: project
tags:
- djinn
- overview
- vision
---

# Djinn

AI agent personas that help you brainstorm, research, analyze, and build.

## The Problem

AI assistance is powerful but unstructured. Most interactions are ad-hoc: prompt, respond, forget. No framework for sustained success.

- **No structure** - AI assistance is inconsistent without clear patterns for what it should do
- **No standards** - Teams reinvent workflows instead of building on proven patterns
- **No methodology** - Proven thinking techniques (SCAMPER, Five Whys) aren't integrated
- **No memory** - Insights disappear; your 100th conversation isn't smarter than your first
- **No boundaries** - AI tries to be everything instead of focused expertise

## The Solution

A framework with three layers:

| Layer | Purpose | Examples |
|-------|---------|----------|
| **Orchestrators** | WHO - Role personas with boundaries | Analyst, Architect, Recruiter |
| **Skills** | HOW - Thinking techniques that auto-activate | SCAMPER, Five Whys, Six Hats |
| **Sub-agents** | ISOLATE - Heavy I/O without polluting context | market-researcher, knowledge-harvester |

Invoke an orchestrator and the AI becomes that persona - thinking with that role's perspective, using relevant skills, delegating research to sub-agents while keeping reasoning clean.

## Vision

Make AI interactions compound instead of starting fresh every session.

## Philosophy

**You stay in the loop.** Djinn isn't about vibe coding or blind delegation. AI accelerates your work - drafting research, applying frameworks, generating options - but you remain the decision maker.

- **AI generates, you review** - Research, plans, and decisions are drafts until you validate them
- **Shared knowledge base** - Memory serves both you and AI; keep it accurate
- **Structured collaboration** - Personas and skills make AI predictable, not autonomous

The goal is faster, better work - not less involvement.

**Docs-First Development.** Your memory is the source of truth. Implementation is ephemeral.

This inverts the typical "code is truth" pattern. Memory contains decisions, principles, and patterns - the source of truth. Implementation (code, prompts, configs) is derived from memory and can be regenerated.

The docs define WHAT and WHY; implementation defines HOW.

**Human owns the docs.** AI generates drafts to memory. You review them in tools like Obsidian. You refine, correct, and approve. Once reviewed, memory becomes the trusted foundation that guides all future work.

See [[Memory]] pattern for the full docs-first philosophy.

## Goals

- **Extensible standards** - Patterns anyone can build on
- **Integrated methodology** - Proven thinking techniques baked in
- **Persistent knowledge** - Decisions and insights accumulate over time
- **User as reviewer** - AI proposes, you decide

## Relations

- [[Architecture]] - Conceptual design and decision frameworks
- [[Claude Code Guide]] - Installation and usage for Claude Code
- [[Memory]] - Docs-first philosophy
- [[decisions/implementations/Claude Code]] - How we implement Djinn on Claude Code
