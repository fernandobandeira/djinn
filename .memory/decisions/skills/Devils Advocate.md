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

**Auto-activates when:** User mentions "challenge this", "what could go wrong", "devil's advocate", "stress test", "red team", "poke holes", "review this", "validate", "audit"

---

## Techniques

### Pre-mortem Analysis

Imagine the project has failed completely, then work backward to identify what went wrong. By vividly imagining failure, you surface risks that optimism blinds you to.

**Post-mortems** analyze failure after it happens—useful but too late. **Pre-mortems** create psychological permission to voice concerns. Instead of asking "what could go wrong?" (which triggers defensiveness), you say "it failed—why?" This reframes criticism as analysis.

Research shows pre-mortems increase ability to identify reasons for future outcomes by 30%.

#### When to Use

- Before starting significant projects
- At major decision points
- When team is overly confident
- Before committing significant resources
- **Time:** 20-40 minutes

#### Process

**Step 1: Set the Scene** (2 min)

Transport yourself to the future where the project has failed.

> "It's [date 6-12 months out]. The project has failed completely. Not partially—completely. It's being shut down as a total failure. We're here to understand why."

Key mindset shifts:
- Failure is a fact, not a possibility
- We're analyzing, not preventing (yet)
- All concerns are valid data

**Step 2: Individual Failure Brainstorm** (10 min)

Each person writes reasons for failure. Prompt: "What went wrong? Write as many reasons as you can think of. Be specific."

Categories to consider:
- **Technical failures:** Technology didn't work, integration broke, performance issues
- **Team failures:** Lost key people, skill gaps, communication breakdown
- **Resource failures:** Ran out of money, time, or people
- **Market failures:** Wrong timing, customers didn't want it, competition won
- **Execution failures:** Scope creep, missed deadlines, quality issues
- **External failures:** Regulatory changes, economic shifts, dependencies failed

Write complete statements:
- Not: "timeline"
- But: "We underestimated integration complexity and missed the launch window"

**Step 3: Share and Expand** (10 min)

Collect all failure reasons without filtering.

Rules:
- No debating whether something is "realistic"
- Add to others' ideas ("Yes, and that probably led to...")
- Encourage the uncomfortable ones
- The concerns people are hesitant to voice are often most valuable

Expansion prompts:
- "What else could have gone wrong?"
- "What's the failure mode nobody wants to mention?"
- "What did we assume that turned out to be wrong?"

**Step 4: Cluster and Prioritize** (10 min)

Group failure modes by theme and assess:

| Failure Mode | Likelihood | Impact | Detection | Priority |
|--------------|------------|--------|-----------|----------|
| [Description] | H/M/L | H/M/L | Easy/Hard | 1-5 |

Priority scoring:
- High likelihood + High impact + Hard to detect = P1
- Low likelihood + Low impact + Easy to detect = P5

**Step 5: Create Prevention Plan** (10 min)

For top priority failure modes, define prevention:

For each P1/P2 failure mode:
- **Early warning signs:** How would we know this is happening?
- **Prevention actions:** What can we do now to reduce likelihood?
- **Contingency plan:** If it happens anyway, what's our response?
- **Owner:** Who's responsible for monitoring and prevention?

#### Example

**Project:** Launching new mobile app in 6 months

**Failure Statement:** "It's January. The app launch was a disaster. We're pulling it from the store. What happened?"

**Failure Reasons Generated:**
1. "Performance was terrible on older devices—we only tested on new phones"
2. "Apple rejected us twice for privacy policy issues, delaying launch"
3. "Backend team was pulled to support the main product crisis in October"
4. "We launched during the holiday freeze when nobody could fix bugs"
5. "Key features were cut at the last minute, leaving a hollow product"
6. "QA was compressed to two weeks and missed critical bugs"
7. "We assumed the existing API would work but it couldn't handle mobile patterns"
8. "The designer quit in month 3 and we never found a replacement"

**Priority Assessment:**

| Failure Mode | Likelihood | Impact | Priority |
|--------------|------------|--------|----------|
| Backend team pulled away | High | High | P1 |
| Performance on old devices | Medium | High | P1 |
| QA compression | High | Medium | P2 |
| App store rejection | Medium | Medium | P2 |
| Holiday launch timing | Low | High | P2 |

**Prevention Plan** (for P1: Backend team pulled):
- **Early warning:** Any escalation on main product
- **Prevention:** Get explicit commitment from leadership on team allocation
- **Contingency:** Identify backup engineers, document architecture heavily
- **Owner:** Project lead, checkpoint at each sprint

#### Common Mistakes

- **Not committing to failure** - Hedging with "might" instead of "did"
- **Focusing only on external factors** - "The market changed" vs. "We didn't adapt"
- **Stopping at surface reasons** - "Timeline slipped" vs. "We underestimated X because Y"
- **Skipping the prevention plan** - Analysis without action is just anxiety
- **Only doing once** - Risks change as projects evolve

#### Tips

- **Make failure vivid** - Don't say "the project might fail." Say "The project has failed. The team is demoralized. The company lost $2M. Your boss is asking what happened."
- **Include embarrassing failures** - "We failed because we ignored obvious warning signs everyone saw but nobody mentioned"
- **Time pressure helps** - Rapid generation surfaces intuitive concerns before the rational mind filters them
- **Revisit periodically** - Pre-mortems at project start, then again at major milestones as new risks emerge

---

### Red Team / Blue Team

Adversarial analysis where one side attacks and the other defends. Borrowed from military and security practices, this technique finds vulnerabilities through simulated opposition.

#### When to Use

- Testing security of systems or processes
- Validating business strategies against competition
- Stress-testing plans before launch
- Finding holes in arguments or proposals
- **Time:** 30-60 minutes

#### The Roles

**Red Team (Attackers)**

Mission: Find every way this could fail, be exploited, or break.

Mindset:
- "I want this to fail"
- "What's the weakest point?"
- "How would a smart adversary attack this?"
- "What assumptions can I violate?"

Attack vectors:
- Technical vulnerabilities
- Process gaps
- Human factors (social engineering, errors)
- Resource constraints
- External threats
- Edge cases and exceptions

**Blue Team (Defenders)**

Mission: Justify decisions, strengthen weak points, prove resilience.

Mindset:
- "Why is this the right approach?"
- "How do we handle that attack?"
- "What safeguards exist?"
- "Where do we need to add protection?"

Defense strategies:
- Explain existing mitigations
- Propose new safeguards
- Accept valid vulnerabilities
- Prioritize fixes by severity

#### Process

**Round 1: Red Team Attack** (15-20 min)

Red team generates attacks without filtering:

Attack categories:
1. **Direct attacks:** How do we break the core functionality?
2. **Indirect attacks:** What adjacent systems can we exploit?
3. **Resource attacks:** Can we overwhelm, exhaust, or starve it?
4. **Social attacks:** How do we manipulate the humans involved?
5. **Timing attacks:** When is it most vulnerable?

Output: List of attack vectors with severity assessment

| Attack | Vector | Severity | Likelihood |
|--------|--------|----------|------------|
| [Description] | [How] | H/M/L | H/M/L |

**Round 2: Blue Team Defense** (15-20 min)

Blue team responds to each attack:

For each attack:
- Acknowledge or dispute the vulnerability
- Explain existing mitigations (if any)
- Propose additional safeguards
- Estimate cost to fix vs. cost of exploitation

Output: Defense assessment

| Attack | Current Mitigation | Proposed Fix | Priority |
|--------|-------------------|--------------|----------|
| [Attack] | [Existing defense] | [New safeguard] | P1/P2/P3 |

**Round 3: Escalation** (10-15 min)

Red team tries to break the defenses:
- "Your mitigation assumes X—what if X fails?"
- "That safeguard has its own vulnerabilities..."
- "You can't afford to implement that defense..."

Blue team responds:
- Strengthen position with deeper justification
- Accept valid criticisms and escalate to action items
- Identify acceptable risk levels

**Round 4: Synthesis** (10 min)

Together, create action plan:
1. **Critical fixes:** Must address before proceeding
2. **Important improvements:** Address within defined timeframe
3. **Accepted risks:** Documented with rationale
4. **Monitoring:** How to detect if attacks occur

#### Example

**Subject:** New customer data API

**Red Team Attacks:**
1. "Rate limiting is too generous—we can scrape all customer data in hours"
2. "API key in URL parameters means it ends up in logs"
3. "No field-level permissions—all or nothing access"
4. "Error messages reveal internal structure"
5. "No audit logging—breaches go undetected"

**Blue Team Defense:**
1. "Fair point—will reduce rate limits and add anomaly detection"
2. "Acknowledged—will move to header-based authentication"
3. "Intentional for v1—adding field-level permissions in Q2"
4. "Will implement generic error messages in production"
5. "Audit logging exists but isn't comprehensive—expanding scope"

**Red Team Escalation:**
- "Q2 is too late for field-level permissions—you're launching to enterprise customers in Q1"
- "Rate limiting + anomaly detection still allows slow exfiltration over weeks"

**Final Actions:**
- P0: Fix authentication method before launch
- P0: Accelerate field-level permissions to launch
- P1: Implement comprehensive audit logging
- P2: Add slow-exfiltration detection
- Accepted: Generic error messages can wait for v1.1

#### Solo Red/Blue Teaming

When working alone:
1. **Write red attacks first** - Be ruthless, list everything
2. **Step away briefly** - Clear your head
3. **Return as blue team** - Defend genuinely
4. **Alternate rounds** - Keep going until attacks are weak
5. **Synthesize** - Create action list

Written format helps maintain separation between attacker/defender mindsets.

#### Common Mistakes

- **Pulling punches** - Red team goes easy to avoid conflict
- **Taking it personally** - Blue team gets defensive about their work
- **Stopping at obvious attacks** - Real adversaries are creative
- **Defending everything** - Some attacks aren't worth mitigating
- **No follow-through** - Findings collected but never acted on

#### Tips

**For Red Team:**
- Be creative—think like different types of attackers
- Go meta—attack the defense mechanisms themselves
- Consider insiders—not all threats are external
- Time matters—attack during deployment, maintenance, transitions
- Chain attacks—small vulnerabilities combine into big ones

**For Blue Team:**
- Don't be defensive (emotionally)—valid attacks help you
- Admit unknowns—"We don't have mitigation for that yet"
- Quantify risk—not all vulnerabilities need fixing
- Propose, don't promise—solutions need validation too

---

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