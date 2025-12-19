---
title: Root Cause
type: note
permalink: decisions/skills/root-cause
---

# Root Cause

**Tier:** Universal

## Core Principle

**Solve the right problem first.** Surface symptoms hide deeper causes—dig until you find what's actually broken.

## Problem

Surface-level problem solving misses real issues. Symptoms get treated, not causes:
- Quick fixes address symptoms, problems recur
- Teams jump to solutions before understanding the problem
- Underlying needs remain unmet
- Effort wasted on wrong solutions

## Solution

Techniques to find fundamental causes and underlying needs.

**Auto-activates when:** User mentions "why", "root cause", "first principles", "real problem", "underlying"

---

## Techniques

### Five Whys

Reach root cause by repeatedly asking "why?" Originally from Toyota's manufacturing process, this technique cuts through symptoms to fundamental causes.

#### When to Use

- Understanding why a problem occurs
- Problem keeps recurring despite fixes
- Want to avoid band-aid solutions
- Need to explain a situation clearly
- **Time:** 10-20 minutes

#### Process

**Step 1: State the Problem Clearly**

Define what's happening in specific, observable terms.

Bad problem statements:
- "Performance is bad" (vague)
- "Users are unhappy" (subjective)
- "The system doesn't work" (undefined)

Good problem statements:
- "Page load time exceeds 5 seconds for 30% of users"
- "Support tickets doubled in the last month"
- "Deployment failed 3 times this week"

**Step 2: Ask "Why?" Repeatedly**

For each answer, ask why again. Continue until you reach a cause you can act on.

Format:
```
Problem: [Observable problem]
Why 1: [First-level cause] → Why?
Why 2: [Deeper cause] → Why?
Why 3: [Deeper still] → Why?
Why 4: [Getting fundamental] → Why?
Why 5: [Root cause - something you can address]
```

Tips for asking why:
- Ask with genuine curiosity, not accusation
- Accept "I don't know" as a signal to investigate
- Follow the chain of causation, not blame
- Sometimes you need more than 5, sometimes fewer

**Step 3: Verify the Root Cause**

Test whether you've found the actual root:
- If we fix this, does the original problem go away?
- Can we go deeper? (If yes, keep going)
- Is this actionable? (Root causes should be fixable)
- Does this explain other related problems?

**Step 4: Define Action**

Once at root cause, define specific actions:
- What needs to change?
- Who owns making it happen?
- How will we know it worked?
- How do we prevent recurrence?

#### Example: The Classic

**Problem:** The Jefferson Memorial is deteriorating faster than other monuments.

- Why? Because it gets washed more frequently.
- Why? Because it has more bird droppings.
- Why? Because there are more birds around it.
- Why? Because there are more spiders for birds to eat.
- Why? Because there are more midges (spider food) around the memorial.
- Why? Because the lights are turned on at dusk, attracting midges.

**Root cause:** Timing of lights
**Solution:** Turn on lights later
**Result:** Midge population dropped, as did birds, spiders, droppings, and washing frequency.

#### Example: Software Development

**Problem:** Users abandoning signup flow

- Why are users abandoning? Form takes too long to complete.
- Why does it take too long? We ask for 15 fields of information.
- Why do we ask for so much? We thought we needed it for personalization.
- Why do we think we need it all upfront? We never tested progressive profiling.
- Why haven't we tested it? We assumed signup had to be complete before access.

**Root cause:** Untested assumption about when to collect data
**Solution:** Test progressive profiling—collect basics at signup, rest during usage
**Result:** Signup completion increased 40%

#### Branching Five Whys

Sometimes causes branch. Follow each branch.

```
Problem: Deployment failed
├── Why? Tests failed
│   ├── Why? Test environment differs from production
│   │   └── Why? Environment setup is manual
│   │       └── Root: No infrastructure-as-code
│   └── Why? New test was flaky
│       └── Why? Test depends on timing
│           └── Root: Missing async handling
└── Why? Deployment script error
    └── Why? Script wasn't updated for new service
        └── Root: No deployment checklist
```

Multiple root causes are common. Address the most impactful first.

#### Workshop Format (30 min)

For teams:
1. **Define problem** (5 min) - Get specific, observable problem statement
2. **Individual why chains** (10 min) - Each person writes their own chain
3. **Share and discuss** (10 min) - Compare chains, identify common roots
4. **Action planning** (5 min) - Define fixes for top root causes

#### Common Mistakes

- **Stopping too early** - Surface causes feel like root causes
- **Accepting blame as cause** - "Human error" is a symptom, not a root
- **Single chain thinking** - Missing that multiple factors contribute
- **No verification** - Assuming root cause without testing
- **No action** - Finding root cause but not fixing it
- **Asking "who" instead of "why"** - This is about systems, not people

#### Tips

- Keep asking—the first "why" is almost never the root cause
- Look for systems—root causes are usually systemic (process, tooling, communication) not individual
- Accept multiple causes—real problems rarely have single causes
- Verify, don't assume—"I think this is the root cause" should lead to investigation, not immediate action

---

### First Principles Decomposition

Strip away assumptions and conventions to find fundamental truths, then rebuild solutions from basics. Instead of reasoning by analogy ("how do others do it?"), reason from foundational truths ("what must be true?").

#### When to Use

- Current approaches feel fundamentally wrong
- Industry "best practices" aren't working for you
- Want to innovate rather than iterate
- Challenging "that's just how it's done"
- **Time:** 30-60 minutes for thorough analysis

#### Process

**Step 1: State What You're Examining**

Identify the solution, approach, or assumption to decompose.

Examples:
- "We need to hire more developers to go faster"
- "Enterprise software requires a sales team"
- "Mobile apps need to be native for good performance"
- "We must raise prices to increase revenue"

**Step 2: List All Assumptions**

Write every assumption embedded in the current approach.

Questions to surface assumptions:
- "Why do we believe this is necessary?"
- "What are we taking for granted?"
- "What would an outsider find strange about this?"
- "What would we have to explain to someone who knows nothing?"

Example decomposition:
Statement: "We need to hire more developers to go faster"

Assumptions:
1. More people = faster progress
2. Bottleneck is developer capacity
3. New developers will be productive quickly
4. Current developers are working on the right things
5. The work requires developers specifically
6. Going faster is the right goal
7. Hiring is the only way to add capacity

**Step 3: Challenge Each Assumption**

For each assumption, ask: "Is this actually true? What evidence supports it?"

| Assumption | Evidence For | Evidence Against | Verdict |
|------------|--------------|------------------|---------|
| [Assumption] | [Support] | [Counter] | True/False/Uncertain |

Types of challenges:
- **Empirical:** Is there data showing this is true?
- **Logical:** Does this necessarily follow?
- **Historical:** Has this always been true? Will it remain true?
- **Contextual:** Is this true in our specific situation?

**Step 4: Identify Fundamental Truths**

What must be true regardless of assumptions?

The bedrock questions:
- "What are the physics-level constraints here?"
- "What cannot be changed no matter what?"
- "If we started from scratch with unlimited resources, what would still be true?"

Example fundamentals (for software development):
- Code must be written (or generated)
- Requirements must be understood
- Code must be deployed to be useful
- Users must be able to access the software
- Bugs will exist and must be found

Note: Very few things are truly fundamental. Most "constraints" are actually assumptions.

**Step 5: Rebuild from Fundamentals**

Given only the fundamental truths, what solutions emerge?

Rebuilding questions:
- "Knowing only the fundamentals, how would we solve this?"
- "What's the simplest solution that satisfies the fundamentals?"
- "What becomes possible without the false assumptions?"

Example rebuilding:

Fundamental: "Code must be written or generated"
- What if we generate more code? (AI, code gen tools)
- What if we reduce code needed? (no-code tools, libraries)
- What if we buy instead of build? (existing solutions)

New solutions:
- Improve requirements clarity (might be the real bottleneck)
- Adopt more automation/generation
- Buy/integrate instead of building
- Remove features to go faster
- Change what "faster" means

#### Example: Electric Vehicles (Tesla's approach)

**Conventional wisdom:** "Electric cars must be expensive because batteries are expensive"

**Assumption decomposition:**
- Battery packs cost $X per kWh (fact at time)
- Battery packs are made of cells, modules, cooling, wiring (fact)
- Battery costs are fixed (ASSUMPTION)
- We must buy batteries from suppliers (ASSUMPTION)
- Battery design is optimized (ASSUMPTION)

**Fundamental truths:**
- Batteries need certain raw materials
- Raw material costs set theoretical floor
- Physics of energy storage is fixed

**Gap analysis:**
- Spot price of raw materials = ~$80/kWh theoretical minimum
- Actual battery cost = ~$600/kWh
- Gap = manufacturing, margins, inefficiency (not physics)

**Rebuilt approach:**
- Vertical integration to remove margins
- Manufacturing innovation to reduce labor
- Design optimization to reduce materials
- Result: Dramatic cost reduction over time

#### Common Mistakes

- **Calling preferences "fundamentals"** - "Users want pretty UI" is not fundamental
- **Stopping at industry standard** - "That's how everyone does it" is assumption, not truth
- **Ignoring real constraints** - Some things (physics, law, math) actually can't change
- **Analysis paralysis** - First principles is for insight, not infinite regression
- **Throwing away valid conventions** - Sometimes the conventional approach is right

#### Tips

- **Question experts** - Experts often know "how" without questioning "why." They may be the most trapped by convention.
- **Look for "everybody knows"** - When something is obvious to everyone, it's often an unexamined assumption.
- **Find the physics** - Actual laws of physics are fundamental. Business "laws" rarely are.
- **Test with "why 5 times"** - If you can't explain why something must be true after 5 whys, it's probably an assumption.
- **Start over mentally** - "If we were starting this company/product/team today, would we do it this way?"

---

### Jobs-to-be-Done (JTBD)

Understand what users are really trying to accomplish, not what they say they want. People don't buy products—they hire them to do jobs. Focus on the job, not the solution.

**The insight:** People don't want a quarter-inch drill. They don't even want a quarter-inch hole. They want to hang a picture. Maybe they want to make their home feel personal. The "job" is at the level of the outcome they care about.

**The framework:**
> When [situation], I want to [motivation], so I can [outcome].

#### When to Use

- Understanding user needs for product development
- Figuring out why users adopt (or don't adopt) solutions
- Finding innovation opportunities
- Escaping feature-comparison thinking
- **Time:** 20-40 minutes

#### Process

**Step 1: Identify the Job**

Look beyond the product to what users are trying to accomplish.

Questions to find the job:
- "What were you trying to accomplish when you [used our product]?"
- "What would you have done if this solution didn't exist?"
- "What does success look like?"
- "What happens after you complete this task?"

Moving up the job ladder:

| Level | Example |
|-------|---------|
| Task | Send an email |
| Activity | Communicate with colleague |
| Job | Stay aligned on project |
| Outcome | Successfully complete project |
| Aspiration | Be valued at work |

Usually the job is 2-3 levels above the task.

**Step 2: Map the Job's Dimensions**

Jobs have functional, emotional, and social dimensions.

**Functional job:** What task are they trying to complete?
- "Get from point A to point B"
- "Organize my thoughts"
- "Track expenses"

**Emotional job:** How do they want to feel?
- "Feel confident I won't forget anything"
- "Feel in control of my finances"
- "Feel less anxious about the presentation"

**Social job:** How do they want to be perceived?
- "Look prepared to my boss"
- "Seem sophisticated to friends"
- "Appear professional to clients"

**Complete job statement:**
> When [preparing for a presentation], I want to [organize my key points] so I can [deliver confidently] and [appear knowledgeable to my audience].

**Step 3: Understand Hiring Criteria**

What makes someone "hire" or "fire" a solution for this job?

**Hiring triggers** (why they start using something):
- Situation changed (new job, new problem)
- Old solution stopped working
- Discovered new possibility
- Pain became unbearable

**Firing triggers** (why they stop):
- Job changed or went away
- Found better solution
- Too much friction
- Emotional need unmet

**Switching costs** (why they stay with current solution):
- Learning curve of new solution
- Data/history in current solution
- Habits and familiarity
- Social pressure

**Step 4: Identify Job Competitors**

Competition isn't just similar products—it's anything people hire for the same job.

Example: Job = "feel more awake in the morning"
- Coffee (product category you might think of)
- Energy drinks
- Exercise
- Cold shower
- Going to bed earlier
- Stimulating news/social media
- Doing nothing (accepting tiredness)

Questions:
- "What did you use before our product?"
- "What would you do if our product disappeared?"
- "What else do you use for this purpose?"

**Step 5: Define Success Metrics**

How does the user measure whether the job is done well?

- **Speed:** How quickly can I get this done?
- **Quality:** How well is the outcome achieved?
- **Reliability:** Can I count on this working?
- **Convenience:** How much effort does it take?
- **Cost:** What does it cost me (money, time, social)?
- **Experience:** How does it feel to do this?

#### Example: Note-Taking App

**Initial view:** Users want note-taking features (folders, tags, search)

**JTBD Analysis:**

**Job:** "When I encounter interesting information, I want to capture and organize it so I can find and use it later when relevant."

**Dimensions:**
- Functional: Capture quickly, find later
- Emotional: Feel secure that good ideas aren't lost
- Social: Appear well-prepared and knowledgeable

**Competitors:**
- Other note apps
- Bookmarking
- Screenshots
- Sending to self via email
- Memory (doing nothing)
- Telling someone else

**Success metrics:**
- Can I capture without interrupting my flow?
- Can I find it when I need it (not just when I look)?
- Do my notes actually get used or just accumulate?

**Insights:**
- "Find it when I need it" suggests surfacing, not just search
- "Without interrupting flow" means speed matters more than organization
- "Actually get used" means retrieval context matters
- Real competition is "memory + doing nothing"—very low friction

#### Common Mistakes

- **Confusing task with job** - "Send email" is task; "stay aligned with team" is job
- **Ignoring emotional/social** - Functional job alone misses why people really choose
- **Narrow competitive view** - Only comparing to "similar" products
- **Assuming you know the job** - Actually talk to users, don't guess
- **Static job definition** - Jobs evolve with context and capability

#### Tips

- **Talk about specific instances** - Not "how do you usually..." but "tell me about the last time you..."
- **Watch for workarounds** - What awkward things do people do? Workarounds reveal unmet jobs.
- **Listen for struggling moments** - "The hard part is..." and "I wish I could..." reveal job friction.
- **The job is stable, solutions change** - "Get from A to B" hasn't changed. Horse → car → Uber → autonomous vehicles.
- **Don't ask what they want** - Ask what they're trying to accomplish. Users are experts on problems, not solutions.

---

## Supported By

- [[Role Playing]] - JTBD requires stepping into user's shoes to understand what job they're "hiring" the product for

- [[User Research]] - JTBD needs deep understanding of user context, motivations, and struggles

- [[Devils Advocate]] - Embedded throughout:
  - **Five Whys:** Each "why" challenges the previous answer
  - **First Principles:** Challenges assumptions to break down to fundamentals

## Why It Matters

- **Universal need** - Any agent analyzing problems benefits
- **Prevents wasted effort** - Solving right problem the first time
- **Reveals hidden needs** - JTBD surfaces what users actually want
- **Foundational thinking** - First principles enables novel solutions

## Used By

- [[Analyst]] - Problem analysis, requirements elicitation
- [[Architect]] - Understanding architectural constraints
- [[Recruiter]] - Understanding what agent user really needs

## Relations

- [[Architecture]] - Part of Tier 1 skill layer
- [[Catalog]] - Listed in component index
- [[Skill]] - Follows skill pattern