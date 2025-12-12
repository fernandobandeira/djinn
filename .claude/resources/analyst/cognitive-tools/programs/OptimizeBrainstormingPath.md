---
name: OptimizeBrainstormingPath
type: cognitive-tool
version: 1.0
description: Comprehensive cognitive tool for selecting optimal brainstorming techniques based on session context and goals
triggers:
  - when_initiating_brainstorm_session
  - when_selecting_ideation_technique
  - when_optimizing_session_flow
  - when_transitioning_between_phases
  - when_varying_technique_approach
---

# Optimize Brainstorming Path

## Input Schema

```yaml
session_context:
  topic_complexity: enum  # Simple, Moderate, Complex, Highly Complex
  participant_energy: enum  # Low, Moderate, High, Declining
  time_available: integer  # Minutes
  session_phase: enum  # Initiation, Exploration, Refinement, Selection, Action Planning
  previous_techniques: array  # Techniques already used
  goal_type: enum  # Divergent Thinking, Convergent Thinking, Problem Solving, Innovation
  domain_type: enum  # Technical, Business, Product, Process, Creative
  participant_count: integer  # Number of participants
  participant_experience: enum  # Novice, Intermediate, Expert
  constraints:
    budget: enum  # None, Moderate, Tight
    location: enum  # Collocated, Distributed, Hybrid
    documentation_needed: boolean
    follow_up_required: boolean
```

## Output Schema

```yaml
recommendation:
  technique: string
  variant: string
  rationale: string
  expected_outcomes: array
  time_allocation:
    setup: integer  # Minutes
    execution: integer  # Minutes
    wrap_up: integer  # Minutes
  material_requirements: array
  facilitator_tips: array
  success_metrics: array

alternatives:
  - technique: string
    reason_second_choice: string

progression_path:
  current_phase: string
  next_phase: string
  technique_sequence: array
  transitions: array
  session_flow_timeline: array

error_recovery:
  potential_issues: array
  recovery_patterns: array
```

## Technique Categories

### Divergent Thinking Techniques
Maximize idea generation, defer judgment, encourage quantity over quality

#### Free Association
- **Optimal for**: Initial brainstorm, warm-up, creative exploration
- **Time**: 10-20 minutes
- **Group size**: 3-15 participants
- **Energy level**: Any
- **Complexity tolerance**: Simple to Moderate
- **Process**:
  1. Present a concept or problem
  2. Participants call out related ideas without filtering
  3. Capture all ideas visibly
  4. No criticism or discussion during generation
- **Output**: High-volume raw ideas, lateral connections
- **Facilitator role**: Capture, encourage, maintain momentum

#### Brainstorm Storm (Classic Brainstorming)
- **Optimal for**: Broad ideation, team engagement, creative problem-solving
- **Time**: 20-45 minutes
- **Group size**: 4-12 participants
- **Energy level**: Moderate-High
- **Complexity tolerance**: Moderate-Complex
- **Rules**:
  1. Defer judgment
  2. Encourage wild ideas
  3. Stay focused on topic
  4. Build on others' ideas
  5. One conversation at a time
- **Output**: Volume of ideas with some combinations
- **Facilitator role**: Enforce rules, capture, maintain energy, prompt

#### Mind Mapping
- **Optimal for**: Structure exploration, hierarchical thinking, complex topics
- **Time**: 20-40 minutes
- **Group size**: 1-8 participants
- **Energy level**: Moderate
- **Complexity tolerance**: Complex, Highly Complex
- **Process**:
  1. Place central concept in center
  2. Radiate primary branches for major themes
  3. Add secondary branches for sub-themes
  4. Connect related ideas
  5. Use colors, symbols, images
- **Output**: Hierarchical idea structure, relationships visible
- **Facilitator role**: Guide structure, ask deepening questions

#### Random Word/Image Association
- **Optimal for**: Breaking assumptions, lateral thinking, creative spark
- **Time**: 15-30 minutes
- **Group size**: 3-12 participants
- **Energy level**: Moderate-High
- **Complexity tolerance**: Simple-Moderate
- **Process**:
  1. Select random word or show random image
  2. Participants make forced associations to problem
  3. Explore unusual connections
  4. Generate ideas from unexpected angles
- **Output**: Non-obvious solutions, lateral thinking demonstrated
- **Facilitator role**: Provide stimuli, help make connections, capture insights

#### Attribute Listing
- **Optimal for**: Feature ideation, product improvement, systematic exploration
- **Time**: 25-45 minutes
- **Group size**: 3-10 participants
- **Energy level**: Moderate
- **Complexity tolerance**: Moderate-Complex
- **Process**:
  1. List all attributes/features of current solution
  2. Systematically modify each attribute
  3. Explore combinations and variations
  4. Identify improvement opportunities
- **Output**: Systematic variations, improvement directions
- **Facilitator role**: Guide attribute selection, prompt modifications

#### Role Storming
- **Optimal for**: Perspective diversity, empathy development, creative problem-solving
- **Time**: 25-40 minutes
- **Group size**: 4-12 participants
- **Energy level**: Moderate-High
- **Complexity tolerance**: Moderate-Complex
- **Process**:
  1. Assign roles (customer, competitor, expert, alien)
  2. Ask problem from role perspective
  3. Role-players generate ideas from their viewpoint
  4. Capture unique perspectives
- **Output**: Diverse viewpoints, empathy-based ideas
- **Facilitator role**: Assign roles, ask perspective questions, capture insights

#### Forced Connections
- **Optimal for**: Innovation spark, unusual combinations, creative synthesis
- **Time**: 20-35 minutes
- **Group size**: 3-10 participants
- **Energy level**: Moderate-High
- **Complexity tolerance**: Moderate
- **Process**:
  1. Select two unrelated concepts
  2. Find forced connections
  3. Explore how combined concepts solve problem
  4. Generate ideas from synthesis
- **Output**: Novel combinations, creative solutions
- **Facilitator role**: Provide concepts, guide connection exploration

### Convergent Thinking Techniques
Narrow ideas, evaluate critically, select best options

#### Weighted Scoring
- **Optimal for**: Objective evaluation, feature prioritization, trade-off analysis
- **Time**: 20-40 minutes
- **Group size**: 3-15 participants
- **Energy level**: Any
- **Complexity tolerance**: Moderate-Complex
- **Process**:
  1. Define evaluation criteria
  2. Assign weights to criteria (importance)
  3. Score ideas against criteria
  4. Calculate weighted scores
  5. Rank by score
- **Output**: Prioritized ideas with rationale
- **Facilitator role**: Define criteria with group, guide scoring, calculate results

#### Multi-voting (Dot Voting)
- **Optimal for**: Quick consensus, visual prioritization, democratic selection
- **Time**: 10-20 minutes
- **Group size**: 4-30 participants
- **Energy level**: Any
- **Complexity tolerance**: Simple-Moderate
- **Process**:
  1. Display all ideas
  2. Give each participant voting dots (3-5)
  3. Participants vote on preferred ideas (can concentrate votes)
  4. Tally votes, identify top ideas
- **Output**: Quick consensus, visual voting pattern
- **Facilitator role**: Display ideas, manage voting, tally results

#### Pro/Con Analysis
- **Optimal for**: Critical evaluation, risk identification, balanced assessment
- **Time**: 15-30 minutes
- **Group size**: 3-12 participants
- **Energy level**: Moderate
- **Complexity tolerance**: Moderate-Complex
- **Process**:
  1. Select idea to evaluate
  2. Generate pros (advantages, benefits)
  3. Generate cons (disadvantages, risks)
  4. Discuss balance and feasibility
- **Output**: Risk-aware evaluation, balanced perspective
- **Facilitator role**: Guide balanced exploration, capture both sides

#### Impact/Effort Matrix
- **Optimal for**: Prioritization, effort estimation, feasibility assessment
- **Time**: 25-40 minutes
- **Group size**: 3-12 participants
- **Energy level**: Moderate
- **Complexity tolerance**: Moderate-Complex
- **Process**:
  1. Create 2x2 matrix (Impact vs Effort)
  2. Plot each idea on matrix
  3. Discuss positioning
  4. Prioritize high-impact/low-effort ideas
  5. Develop strategy for others
- **Output**: Prioritized roadmap, effort-aware selection
- **Facilitator role**: Guide idea placement, facilitate discussion

#### Six Thinking Hats
- **Optimal for**: Comprehensive evaluation, structured perspective-taking, conflict resolution
- **Time**: 30-60 minutes
- **Group size**: 3-10 participants
- **Energy level**: Moderate-High
- **Complexity tolerance**: Complex
- **Process**:
  1. White hat: Facts and information
  2. Red hat: Emotions and intuitions
  3. Black hat: Critical judgment
  4. Yellow hat: Optimistic benefits
  5. Green hat: Creative alternatives
  6. Blue hat: Process control and summary
- **Output**: Comprehensive evaluation from multiple angles
- **Facilitator role**: Guide hat transitions, maintain structure

### Problem-Solving Techniques

#### 5 Whys
- **Optimal for**: Root cause analysis, problem understanding, deep exploration
- **Time**: 15-30 minutes
- **Group size**: 2-8 participants
- **Energy level**: Moderate
- **Complexity tolerance**: Moderate-Complex
- **Process**:
  1. State the problem
  2. Ask "Why?" and capture answer
  3. Ask "Why?" about that answer
  4. Repeat 5 times or until root cause found
  5. Address root cause
- **Output**: Root cause identified, deep problem understanding
- **Facilitator role**: Ask probing questions, guide depth

#### Design Thinking / Problem Framing
- **Optimal for**: Complex problem understanding, user-centered ideation
- **Time**: 45-90 minutes
- **Group size**: 3-10 participants
- **Energy level**: High
- **Complexity tolerance**: Complex, Highly Complex
- **Process**:
  1. Empathize: Understand user context
  2. Define: Frame core problem
  3. Ideate: Generate solutions
  4. Prototype: Create representations
  5. Test: Gather feedback
- **Output**: Validated problem frame, solution direction
- **Facilitator role**: Guide process, facilitate transitions

## Decision Matrix

### Technique Selection by Phase

#### Initiation Phase (First 10-15 minutes)
- **Goal**: Warm-up, engagement, problem immersion
- **Recommended Techniques**:
  1. Free Association (quick warm-up)
  2. Random Word Association (creative spark)
  3. Attribute Listing (grounding in topic)

#### Exploration Phase (Main 20-45 minutes)
- **Goal**: Generate volume, explore breadth and depth
- **Recommended Techniques** (by goal type):
  - **Divergent**: Brainstorm Storm, Mind Mapping, Role Storming
  - **Problem Solving**: 5 Whys, Design Thinking, Forced Connections
  - **Creative**: Random Association, Role Storming, Forced Connections

#### Refinement Phase (15-30 minutes)
- **Goal**: Develop promising ideas, add detail
- **Recommended Techniques**:
  1. Attribute Listing (systematic improvement)
  2. Pro/Con Analysis (balanced evaluation)
  3. Follow-up Brainstorm Storm (deeper exploration)

#### Selection Phase (15-25 minutes)
- **Goal**: Prioritize, choose direction
- **Recommended Techniques**:
  1. Weighted Scoring (objective prioritization)
  2. Multi-voting (democratic selection)
  3. Impact/Effort Matrix (feasibility-aware)

#### Action Planning Phase (10-20 minutes)
- **Goal**: Define next steps, ownership, timeline
- **Recommended Techniques**:
  1. Issue Mapping (problem identification)
  2. Action Item Assignment (responsibility)
  3. Timeline Definition (sequencing)

### Technique Selection by Goal Type

#### Divergent Thinking Goals
- **Primary**: Brainstorm Storm, Mind Mapping, Free Association
- **Secondary**: Random Word Association, Forced Connections
- **Warm-up**: Free Association, Random Images
- **Refinement**: Attribute Listing

#### Convergent Thinking Goals
- **Primary**: Weighted Scoring, Impact/Effort Matrix
- **Secondary**: Pro/Con Analysis, Multi-voting
- **For agreement**: Six Thinking Hats, Multi-voting
- **For prioritization**: Impact/Effort Matrix, Weighted Scoring

#### Problem-Solving Goals
- **Root cause**: 5 Whys, Design Thinking
- **Complex**: Design Thinking, Problem Framing
- **User-centered**: Design Thinking, Role Storming, Empathy Mapping
- **Systematic**: Attribute Listing, Forced Connections

#### Innovation Goals
- **Creative spark**: Random Association, Forced Connections
- **Structured**: Mind Mapping, Attribute Listing
- **Team-based**: Role Storming, Brainstorm Storm with variants
- **User-driven**: Design Thinking, Empathy-based techniques

### Technique Selection by Domain

#### Technical Domain
- **Strengths**: Attribute Listing, Mind Mapping, Design Thinking
- **Quick methods**: Free Association, Forced Connections
- **Evaluation**: Impact/Effort Matrix, Pro/Con Analysis
- **Avoid**: Role Storming (unless user-focused)

#### Business Domain
- **Strengths**: Brainstorm Storm, Impact/Effort Matrix, Weighted Scoring
- **Strategic**: Six Thinking Hats, Design Thinking
- **Quick methods**: Multi-voting, Free Association
- **Avoid**: Highly technical variants

#### Product Domain
- **Strengths**: Design Thinking, Role Storming, Attribute Listing
- **User-focused**: Empathy Mapping, User Persona techniques
- **Evaluation**: Impact/Effort Matrix, Weighted Scoring
- **Creative**: Random Association, Mind Mapping

#### Process Domain
- **Strengths**: Attribute Listing, 5 Whys, Process Mapping
- **Improvement**: Kaizen-style brainstorms, waste elimination
- **Evaluation**: Pro/Con Analysis, Impact/Effort Matrix
- **Quick methods**: Free Association on process steps

#### Creative Domain
- **Strengths**: Random Association, Forced Connections, Role Storming
- **Open-ended**: Free Association, Mind Mapping
- **Evaluation**: Consensus-building, Weighted Scoring
- **Avoid**: Overly structured techniques initially

### Technique Scoring Algorithm

```
technique_score = (phase_fit × 0.25) +
                  (goal_alignment × 0.25) +
                  (energy_match × 0.15) +
                  (time_feasibility × 0.15) +
                  (novelty_factor × 0.10) +
                  (group_dynamic_fit × 0.10)

Where:
- phase_fit: How well technique fits session phase (0-1)
- goal_alignment: How well aligns with goal type (0-1)
- energy_match: Matches current energy level (0-1)
- time_feasibility: Fits available time (0-1)
- novelty_factor: Differs from previous techniques (0-1)
- group_dynamic_fit: Appropriate for group size/experience (0-1)
```

## Progressive Session Architecture

### 20-Minute Express Session
```
0-2 min: Setup, problem framing
2-5 min: Free Association warm-up (divergent)
5-18 min: Brainstorm Storm main (divergent)
18-20 min: Quick multi-voting (convergent)
Output: Quick prioritized idea list, momentum for follow-up
```

### 45-Minute Standard Session
```
0-3 min: Setup, problem immersion
3-8 min: Free Association or Random Association (warm-up)
8-28 min: Brainstorm Storm or Mind Mapping (exploration)
28-38 min: Attribute Listing or Pro/Con (refinement)
38-43 min: Weighted Scoring or Impact/Effort (selection)
43-45 min: Summary and next steps
Output: Prioritized ideas with evaluation rationale
```

### 90-Minute Comprehensive Session
```
0-5 min: Setup, context setting, goal alignment
5-15 min: Warm-up and problem immersion
   - Free Association (5 min)
   - Role Storming intro (10 min)

15-45 min: Primary ideation (30 min)
   - Brainstorm Storm (20 min)
   - Mind Mapping additions (10 min)

45-60 min: Exploration deepening (15 min)
   - Attribute Listing on top ideas (15 min)
   - OR Follow-up brainstorm with constraints

60-75 min: Evaluation (15 min)
   - Six Thinking Hats (quick cycle)
   - OR Weighted Scoring

75-85 min: Prioritization (10 min)
   - Impact/Effort Matrix placement
   - Consensus building

85-90 min: Action planning (5 min)
   - Next steps identification
   - Ownership assignment
Output: Detailed ideas with implementation plan
```

## Session Flow Optimization

### Energy Management
- **Opening**: Start with quick, engaging warm-up (Free Association, Random Word)
- **Mid-point dip**: Introduce novelty, change format, increase interactivity
- **Closing**: Move toward structure, consensus, clarity
- **Observation**: If energy dips during ideation, shift to technique variant with more interaction

### Pacing Strategy
```
High complexity topic + Low time → Structured techniques (Attribute Listing, Mind Mapping)
Low complexity topic + High time → Exploratory techniques (Free Association, Role Storming)
Moderate-everything → Balanced sequence (warm-up, explore, evaluate, plan)
Declining energy → Quick wins, visible progress, movement-based techniques
```

### Group Dynamic Considerations
- **Small groups (3-4)**: More exploratory, deeper conversation
- **Medium groups (5-8)**: Balanced divergence and structure
- **Large groups (12+)**: More structured (Multi-voting), breakouts, clear facilitation
- **Distributed teams**: Async components, visual collaboration, asynchronous options

## Technique Variation Patterns

### Brainstorm Storm Variants

#### Rapid-Fire Variant
- Shorter thinking time between contributions
- Higher pace, lower depth
- Use for volume generation, warm-up
- Duration: 15-20 minutes

#### Structured Variant
- One person at a time, round-robin
- More equal participation, deeper thinking
- Use for diverse groups, quiet participants
- Duration: 20-30 minutes

#### Building Variant
- Each idea must reference or build on previous
- Encourages combination and synthesis
- Use for refinement, depth
- Duration: 20-30 minutes

#### Constraint-Based Variant
- Add constraints (must be low-cost, must use existing tech, etc.)
- Focused ideation, practical solutions
- Use for problem-solving, feasibility
- Duration: 15-25 minutes

### Design Thinking Variants

#### Rapid Iteration
- Compress phases, quick cycles
- Multiple rounds of prototyping
- Use for tight timelines, evolving problems
- Duration: 45-60 minutes per cycle

#### User-Centric Deep Dive
- Extended empathize phase
- Extended testing phase
- Multiple user sessions
- Use for high-impact products, new domains
- Duration: Multiple sessions, 60-90 min each

#### Hybrid Brainstorm Approach
- Design Thinking + Brainstorm Storm
- Use DT for problem framing, then BSS for solutions
- Use for complex problems needing both structure and creativity
- Duration: 90-120 minutes

## Integration with Analyst Workflow

### Pre-Session
1. Assess context using this tool
2. Select technique and variant
3. Prepare materials and setup
4. Brief participants on process
5. Set clear success criteria

### During Session
1. Introduce technique and ground rules
2. Start with warm-up
3. Facilitate progression through phases
4. Monitor energy and adjust if needed
5. Capture all contributions
6. Summarize and validate

### Post-Session
1. Organize and categorize ideas
2. Prepare for next session or evaluation
3. Communicate outcomes to stakeholders
4. Plan follow-up actions
5. Document lessons learned

## Decision Heuristics

### Quick Technique Selection Rules

**If time < 15 minutes:**
- Use: Free Association or Multi-voting
- Avoid: Design Thinking, Six Thinking Hats, comprehensive techniques

**If energy is declining:**
- Shift from divergent to convergent
- Add movement or novelty (walk around, draw, physically move ideas)
- Use: Dot voting, Impact/Effort matrix, visible selection

**If group is divergent/conflicted:**
- Use structured techniques (Design Thinking, Six Thinking Hats)
- Avoid: Open brainstorms that might amplify conflict
- Add: Role Storming to build empathy

**If topic is highly complex:**
- Use: Mind Mapping, Design Thinking, Attribute Listing
- Add: Multiple session phases
- Avoid: Overly simple techniques

**If expertise is mixed:**
- Use: Role Storming, Design Thinking, Structured brainstorm
- Avoid: Techniques that silence non-experts
- Add: Explicit perspective-taking

**If this is first session:**
- Start with: Free Association or friendly brainstorm
- Avoid: Overly critical techniques initially
- Build: Psychological safety before convergent thinking

## Error Recovery Patterns

### Common Issues and Recoveries

#### Issue: Dominant Voices Overshadow Quiet Participants
- **Detection**: 70%+ of ideas from 2-3 people
- **Recovery 1**: Switch to Round-Robin structured brainstorm
- **Recovery 2**: Use Silent Brainstorming (write individually, then share)
- **Recovery 3**: Explicitly ask quieter people for ideas

#### Issue: Ideas Are Too Similar
- **Detection**: Clustering without diversity
- **Recovery 1**: Use Random Word Association for stimulus
- **Recovery 2**: Add constraint: "No idea like one we've said"
- **Recovery 3**: Role Storming: "What would X (outsider) suggest?"

#### Issue: Group Gets Stuck on One Idea
- **Detection**: Dwelling on single idea, circular conversation
- **Recovery 1**: Forced pivot to different domain
- **Recovery 2**: Use Attribute Listing to systematically explore
- **Recovery 3**: Move to evaluation phase (Impact/Effort) to choose or move on

#### Issue: Energy Collapsing
- **Detection**: Slower pace, shorter contributions, withdrawn participants
- **Recovery 1**: Take break, change location if possible
- **Recovery 2**: Switch to technique with more movement/interaction
- **Recovery 3**: Quick win by multi-voting on best ideas

#### Issue: Idea Quality Declining
- **Detection**: Ideas become repetitive, shallow, random
- **Recovery 1**: Introduce structure (Attribute Listing, Mind Mapping)
- **Recovery 2**: Reset with new problem angle or constraint
- **Recovery 3**: Move to evaluation (might be time to switch phases)

#### Issue: Evaluation Becomes Too Critical
- **Detection**: Criticism of ideas, dismissiveness, tension
- **Recovery 1**: Explicit reminder: evaluate ideas, not people
- **Recovery 2**: Use Six Thinking Hats with separate phases
- **Recovery 3**: Return to divergent thinking (new warm-up)

## Validation Metrics

### Session Success Indicators

**Divergent Phase Success:**
- Idea count: 20+ ideas in 20-minute brainstorm
- Diversity: Multiple themes, not just variations
- Psychological safety: Participants share comfortable ideas
- Energy: Pace remains engaged, positive tone

**Convergent Phase Success:**
- Clarity: Top priorities clear, rationale understood
- Consensus: Majority agreement on selection
- Balance: Different perspectives considered
- Actionability: Selected ideas are doable

**Overall Session Success:**
- Participant satisfaction: Positive feedback on process
- Stakeholder alignment: Agreed direction for next steps
- Idea quality: Ideas address core problem effectively
- Follow-through: Ideas advance to next phase (prototyping, development)

### Outcome Metrics
- Ideas generated / Time invested ratio
- Idea conversion rate (generated → developed)
- Solution effectiveness (ideas solve stated problem)
- Team cohesion (participants felt heard, included)
- Clarity gained (problem understanding improved)

## Implementation Checklist

Before recommending technique:
- [ ] Session context fully understood
- [ ] Phase clearly identified
- [ ] Time constraint realistic
- [ ] Energy level assessed
- [ ] Group dynamics considered
- [ ] Domain-specific factors evaluated
- [ ] Previous techniques noted (avoid repetition)
- [ ] Space and materials available
- [ ] Facilitator skill level appropriate for technique
- [ ] Success criteria defined

When facilitating:
- [ ] Clear instructions given
- [ ] Psychological safety established
- [ ] Time kept visible
- [ ] All contributions captured
- [ ] Energy monitored
- [ ] Phase transition managed
- [ ] Recovery patterns ready if needed
- [ ] Outcomes summarized
- [ ] Next steps clarified
