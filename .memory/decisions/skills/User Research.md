---
title: User Research
type: note
permalink: decisions/skills/user-research
---

# User Research

**Tier:** Domain

## Core Principle

**Watch what users do, not just what they say.** Real needs hide beneath stated wants—research reveals the gap.

## Problem

User needs aren't deeply understood:
- Solutions miss the mark
- Features built that nobody uses
- User pain points remain hidden
- Decisions based on assumptions, not evidence
- Information architecture doesn't match user mental models
- Accessibility barriers exclude users

## Solution

Research techniques organized by phase to reduce cognitive load and match user's research stage.

**Auto-activates when:** User mentions "journey map", "user interviews", "survey design", "usability testing", "user research", "customer discovery", "card sorting", "information architecture", "accessibility audit", "design critique", "prototype testing", "requirements gathering", "user story mapping", "discovery phase", "validation phase", "research planning"

## Structure

```
.claude/skills/user-research/
├── SKILL.md
├── workflow/
│   ├── phase-selection.md
│   └── research-planning.md
└── cookbook/
    ├── discovery/
    ├── validation/
    ├── design/
    └── compliance/
```

---

## Phase 1: Discovery

**Goal:** Understand users deeply before building anything.

### Interview Design

Plan and conduct user interviews that uncover genuine insights. The goal isn't to confirm your ideas—it's to understand users deeply.

#### When to Use

- Early discovery (what problem to solve?)
- Understanding context and motivations
- Exploring "why" behind behaviors
- When you need rich, qualitative data
- **Time:** 30-60 min per interview, 5-10 interviews minimum

#### Process

1. **Define Research Questions** - What do you need to learn? (Not what to ask—what to learn)
2. **Create Interview Guide** - Opening → Warm-up → Core Questions → Wrap-up
3. **Recruit Participants** - 5-8 per segment. Patterns emerge around interview 5.
4. **Conduct Interviews** - Listen 80%, talk 20%. Follow interesting threads.
5. **Analyze** - Affinity mapping, look for patterns, synthesize into insights

#### Common Mistakes

- **Talking too much** - You should talk 20%, listen 80%
- **Leading the witness** - Phrasing that suggests "right" answer
- **Accepting surface answers** - Not digging into "why"
- **Confirming hypotheses** - Seeking validation, not truth

---

### Journey Mapping

Visualize the complete user experience across touchpoints and time. Reveal pain points, opportunities, and emotional moments that data alone can't show.

#### When to Use

- Understanding end-to-end user experience
- Finding pain points across touchpoints
- Aligning team on user perspective
- Identifying improvement opportunities
- **Time:** 1-3 hours

#### Process

1. **Define Scope** - Choose persona, scenario, start/end points
2. **Identify Stages** - Break journey into phases (Awareness → Consideration → Decision → Use)
3. **Map User Actions** - What does user DO at each stage?
4. **Identify Touchpoints** - Where/how do they interact?
5. **Capture Emotions** - Internal monologue, emotional highs/lows
6. **Identify Pain Points** - Where does experience break down?
7. **Find Opportunities** - For each pain point, brainstorm improvements

#### Common Mistakes

- **Inside-out mapping** - Mapping your process, not user's experience
- **Missing emotions** - Facts without feelings misses half the story
- **Assumption-based** - Mapping what you think happens vs. what actually happens
- **No action** - Beautiful map that doesn't drive decisions

---

### Card Sorting

Understand how users mentally organize content and expect to find information. Reveals natural mental models for information architecture.

#### When to Use

- Designing site/app navigation structure
- Organizing feature sets or content areas
- Validating existing information architecture
- **Time:** 20-40 minutes per session

#### Types

| Type | When to Use |
|------|-------------|
| **Open** | Exploring mental models - users create own categories |
| **Closed** | Validating structure - users sort into predefined categories |
| **Hybrid** | Predefined + user can create new |

#### Process

1. **Prepare Cards** - 30-60 cards representing content/features
2. **Run Session** - Give instructions, watch them sort
3. **Probe** - "Why did you put [card] in this group?"
4. **Analyze** - Agreement rate, popular categories, problem cards

#### Common Mistakes

- **Too specific or too general cards**
- **Small sample** - Need 15-30 participants for validity
- **Ignoring conflicts** - Different user types may need different IA

---

## Phase 2: Validation

**Goal:** Test whether your solution works for users.

### Prototype Feedback

Gather rapid feedback on low-fidelity prototypes through task-based testing. Validates interaction design before high-fidelity investment.

#### When to Use

- Validating interaction concepts early
- Testing navigation and flow
- Comparing design alternatives
- **Time:** 15-30 minutes per session

#### Fidelity Levels

| Level | Best For |
|-------|----------|
| Paper | Very early concepts |
| Low-fi | Flow and structure |
| Mid-fi | Interaction patterns |
| High-fi | Final validation |

**Rule:** Test at the lowest fidelity that answers your question.

#### Process

1. **Define What You're Testing** - Prototype, fidelity, core question, tasks
2. **Prepare Task Scenarios** - Realistic scenarios, not instructions
3. **Run Session** - Opening, tasks with think-aloud, follow-ups
4. **Document Findings** - Task results, observations, quotes, issues

#### Common Mistakes

- **Leading questions** - "You found that easy, right?"
- **Testing too late** - When changes are expensive
- **Helping too much** - Rescuing users hides real problems
- **Testing visuals on wireframes** - Low-fi tests structure, not aesthetics

---

### Usability Testing

Systematically evaluate how well users can accomplish tasks with a product. Reveals real problems that assumptions and expert reviews miss.

#### When to Use

- Validating design decisions with real users
- Identifying usability issues before launch
- Benchmarking against competitors
- **Time:** 30-60 minutes per session, 5-10 sessions

#### Core Metrics

| Metric | Definition | Benchmark |
|--------|------------|-----------|
| Task success | % completing without critical error | > 78% |
| Time on task | Seconds to complete | Compare to baseline |
| Errors | Wrong clicks, backtracking | < 2 per task |
| Satisfaction | Post-task rating (1-5) | > 4.0 |
| SUS Score | System Usability Scale | > 68 |

#### Process

1. **Welcome** - Intro, consent, "testing product not you"
2. **Warm-up** - Background questions
3. **Tasks** - Core usability tasks with think-aloud
4. **Debrief** - Overall impressions, SUS
5. **Wrap-up** - Thanks, compensation

#### Common Mistakes

- **Too few participants** - 5 find most issues
- **Leading questions** - "Was that confusing?" → "Tell me about that"
- **Helping too quickly** - Let users struggle to reveal real issues
- **Report without action** - Testing is pointless without fixing

---

### Survey Design

Create surveys that collect reliable, actionable data at scale. Bad surveys generate bad data—worse than no data because they create false confidence.

#### When to Use

- Validating qualitative findings at scale
- Measuring satisfaction, preferences, attitudes
- Prioritizing features or problems
- Benchmarking over time
- **Time:** 2-4 hours to design, days to collect

#### When NOT to Use

| Use Surveys For | Don't Use Surveys For |
|-----------------|----------------------|
| How many, how often | Why (use interviews) |
| Measuring attitudes | Understanding context |
| Validating hypotheses | Discovering unknowns |

#### Survey Length Impact

| Length | Completion Rate |
|--------|-----------------|
| 1-3 min (5-10 questions) | 80%+ |
| 5-7 min (15-20 questions) | 60-70% |
| 10-15 min (30-40 questions) | 40-50% |

#### Common Mistakes

- **Double-barreled questions** - Ask speed and reliability separately
- **Leading questions** - "How would you rate?" not "How much did you enjoy?"
- **Jargon** - Use simple language
- **Too long** - Every additional minute drops completion by 5-10%

---

## Phase 3: Design Support

**Goal:** Shape output with research-informed decisions.

### Design Critique

Systematically evaluate designs to identify what works, what doesn't, and how to improve. Separates personal preference from user-centered feedback.

#### When to Use

- Reviewing design concepts or mockups
- Iterating on existing designs
- Making design decisions with evidence
- **Time:** 30-45 minutes per session

#### Framework: WWW

- **What's Working** - What supports user goals? What feels intuitive?
- **What's Not Working** - Where might users get confused? What's missing?
- **What If** - Alternative approaches? Trade-offs to consider?

#### Process

1. **Set Context** - User, goal, constraints, feedback focus
2. **Present Without Defending** - Walk through, don't justify
3. **Silent Review** - Each person notes observations
4. **Structured Feedback** - What's Working → What's Not → What If
5. **Capture and Prioritize** - Immediate action vs. more exploration

#### Common Mistakes

- **Design by committee** - Trying to incorporate all feedback
- **Preference masquerading as critique** - "I don't like blue" vs. "Blue doesn't have enough contrast"
- **Defending instead of listening**

---

### User Story Mapping

Visualize the user journey and organize features into a prioritized backlog. Connects user goals to features and reveals what to build first.

#### When to Use

- Planning new product or major feature
- Prioritizing a backlog with stakeholder alignment
- Identifying MVP scope
- **Time:** 45-90 minutes

#### Structure

```
USER JOURNEY (Activities)
│ Discover → Evaluate → Sign Up → Onboard → Use
│
└── USER TASKS (Steps under each activity)
    │ Search → Compare → Create account → ...
    │
    └── RELEASE SLICES
        MVP ────────────────────
        V1  ────────────────────
        V2  ────────────────────
```

#### Slicing Criteria

- **MVP:** Can users complete core goal? (Embarrassingly small)
- **V1:** Complete experience (Smooth, not just functional)
- **V2:** Enhanced (Delight, efficiency, edge cases)

#### Common Mistakes

- **System-centric** - Tasks describe what system does, not what users do
- **No prioritization** - Flat map without release slices
- **Skipping validation** - Assuming map matches reality

---

### Requirements Gathering

Systematically capture constraints, requirements, and trade-offs that shape design decisions. Prevents rework and misaligned expectations.

#### When to Use

- Starting a new project or major feature
- Defining technical specifications
- Clarifying stakeholder expectations
- **Time:** 20-30 minutes per stakeholder/area

#### Types of Requirements

| Type | Example |
|------|---------|
| Functional | "Users can filter search results" |
| Non-functional | "Page loads in < 2 seconds" |
| Technical | "Must work in IE11" |
| Business | "Launch by Q2" |
| User | "Must be screen-reader accessible" |

#### MoSCoW Prioritization

| Must Have | Should Have | Could Have | Won't Have |
|-----------|-------------|------------|------------|
| Non-negotiables | Important but not critical | Nice-to-haves | Explicitly out of scope |

#### Common Mistakes

- **Vague requirements** - "Must be user-friendly" (meaningless)
- **Solutions disguised as requirements** - "Must use React" (why?)
- **No prioritization** - Everything is "must have"

---

## Phase 4: Compliance

**Goal:** Ensure the product meets standards.

### Accessibility Audit

Systematically evaluate products against accessibility standards to ensure inclusive design. Catches barriers that exclude users with disabilities.

#### When to Use

- Evaluating existing products for compliance
- Before major releases (accessibility gate)
- After design changes to interactive elements
- **Time:** 2-4 hours for comprehensive audit

#### WCAG Levels

| Level | What It Means |
|-------|---------------|
| A | Minimum accessibility |
| AA | Removes significant barriers (most regulations) |
| AAA | Highest accessibility |

**Most organizations target WCAG 2.1 AA.**

#### The Four Principles (POUR)

- **Perceivable:** Users must perceive content (alt text, captions, contrast)
- **Operable:** Users must operate interface (keyboard accessible, enough time)
- **Understandable:** Users must understand content (readable, predictable)
- **Robust:** Works with assistive technologies (valid HTML, ARIA)

#### Audit Process

1. **Scope Definition** - Pages, standard level, audit type
2. **Automated Testing** - axe DevTools, WAVE, Lighthouse
3. **Manual Testing** - Keyboard, screen reader, visual checks
4. **Document Findings**

#### Common Issues

| Issue | WCAG | Fix |
|-------|------|-----|
| Missing alt text | 1.1.1 | Add descriptive alt |
| Low contrast | 1.4.3 | Increase contrast ratio |
| No focus indicator | 2.4.7 | Add visible focus style |
| Keyboard trap | 2.1.2 | Ensure Esc/Tab exits |

#### Common Mistakes

- **Only automated testing** - Tools miss 40-60% of issues
- **Testing late** - Easier to build accessible from start
- **Ignoring screen reader experience** - The most impactful check
- **No prioritization** - Critical issues should block launch

---

## Core Principles

1. **Observe behavior, not just words** - What people do often differs from what they say
2. **Ask "why" and "show me"** - Surface explanations hide deeper truths
3. **Recruit the right users** - Wrong participants = wrong insights
4. **Separate research from validation** - Don't seek confirmation of your ideas
5. **Document and share** - Insights locked in your head don't help the team

## When User Research Works Best

- Building something new (discovery)
- Existing product isn't performing (diagnosis)
- Prioritizing features (evidence-based decisions)
- Resolving internal debates (let users decide)
- Before major investment (de-risk)

---

## Supported By

- [[Role Playing]] - Journey mapping and persona analysis require stepping into user roles

## Why It Matters

- **Domain need** - [[UX]], [[Analyst]], [[PM]] benefit most
- **Grounds decisions** - Evidence over assumption
- **Reveals hidden needs** - Users often can't articulate wants
- **Reduces waste** - Build what users actually need
- **Inclusive design** - Accessibility audits ensure everyone can use it

## Used By

- [[UX]] - Primary skill for Ulysses; all 10 techniques available
- [[Analyst]] - Understanding user requirements, journey mapping
- [[PM]] - User story validation, requirements gathering

## Relations

- [[Architecture]] - Part of Tier 2 skill layer
- [[Catalog]] - Listed in component index
- [[Skill]] - Follows skill pattern
- [[UX]] - Primary orchestrator using this skill