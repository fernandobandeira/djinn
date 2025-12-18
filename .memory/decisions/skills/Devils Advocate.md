---
title: Devils Advocate
type: note
permalink: decisions/skills/devils-advocate
---

# Devils Advocate

**Tier:** Foundational

## Core Principle

**Break it in the lab, not in production.** Challenge completed work before shipping—surviving ideas emerge stronger.

**Meaningful critique, not nitpicking.** Nothing will be perfect. Focus on issues that actually matter—risks that could cause real harm, not cosmetic imperfections.

## Problem

Confirmation bias leads to blind spots:
- Teams fall in love with their ideas
- Risks get overlooked until too late
- Plans fail due to unconsidered scenarios
- Groupthink suppresses valid concerns

## Solution

Systematic challenge techniques that stress-test ideas before commitment.

**Techniques:**
- **Pre-mortem Analysis** - Imagine failure, work backward to causes
- **Red Team/Blue Team** - Adversarial challenge of proposals (uses [[Role Playing]])
- **Assumption Surfacing** - Explicitly list and challenge assumptions

**Auto-activates when:** User mentions "challenge this", "what could go wrong", "devil's advocate", "stress test", "red team", "poke holes", "review this", "validate", "audit"

## Supported By

- [[Role Playing]] - Red Team/Blue Team requires stepping into attacker/defender roles; effective challenge means genuinely adopting the adversary's perspective

## Critical Quality Gate

This skill is **essential for reviewing completed work**—not just plans:
- **Review implementations** before they ship
- **Validate architectures** after design, before and after build
- **Challenge deliverables** before handoff
- **Audit decisions** that have already been made

**Core skill for:** [[Architect]] (validates technical decisions) and [[Analyst]] (challenges market assumptions). Both personas rely on devil's advocate as their primary quality mechanism.

## Why It Matters

- **Universal need** - Any agent validating ideas benefits
- **Catches problems early** - Before resources committed
- **Legitimizes dissent** - Safe way to raise concerns
- **Improves robustness** - Ideas that survive are stronger
- **Quality gate for completed work** - Most valuable just before shipping

## Anti-Pattern: Nitpicking

Devil's advocate is NOT about:
- Finding every minor flaw
- Perfectionism that blocks progress
- Criticism without constructive purpose
- Death by a thousand paper cuts

Focus on **material risks**—things that could actually derail the project, harm users, or cause significant rework.

## Used By

- [[Analyst]] - Validating research, challenging assumptions
- [[Architect]] - Architecture reviews, risk analysis
- [[Recruiter]] - Validating agent designs

## Relations

- [[Architecture]] - Part of Tier 1 skill layer
- [[Catalog]] - Listed in component index
- [[Skill]] - Follows skill pattern (foundational skill)