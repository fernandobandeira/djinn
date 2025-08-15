# Cognitive Tool Program: Select Methodology
# Structured reasoning program to choose optimal learning methodology based on constraints

## Program Purpose
This prompt program guides systematic selection of learning methodologies based on learner state, content type, and constraint satisfaction requirements.

## Input Schema
```yaml
selection_context:
  learner_assessment: Results from AssessUnderstanding program
  content_type: Concept, procedure, problem-solving, theory
  learning_objective: What needs to be achieved
  time_constraints: Available session time
  prior_methodology_success: What has worked before
  current_constraint_state: Under/well/over-constrained
```

## Selection Framework

### Step 1: Analyze Learning Objective Type
```
IF objective is conceptual understanding:
  → Consider: Feynman, Socratic, Analogical, Elaborative
ELSE IF objective is procedural skill:
  → Consider: Worked Examples, Practice, Problem-Based
ELSE IF objective is strategic thinking:
  → Consider: Case Method, Problem-Based, Interleaving
ELSE IF objective is transfer/application:
  → Consider: Generation, Dual Coding, Retrieval Practice
```

### Step 2: Match to Learner Understanding Level
```
IF learner is novice:
  → Prioritize: Worked Examples, Analogical, Dual Coding
  → Avoid: Complex Socratic, Advanced Problem-Based
ELSE IF learner is developing:
  → Prioritize: Feynman, Elaborative, Problem-Based
  → Support with: Scaffolded Socratic
ELSE IF learner is proficient:
  → Prioritize: Case Method, Generation, Interleaving
  → Challenge with: Advanced Socratic
ELSE IF learner is expert:
  → Prioritize: Peer Instruction, Teaching Others
  → Focus on: Knowledge Creation, Boundary Exploration
```

### Step 3: Consider Constraint State
```
IF learner is under-constrained (overgeneralizing):
  → Use: Elaborative Interrogation, Boundary Testing
  → Focus: Precision, Limitations, Edge Cases
ELSE IF learner is over-constrained (rigid):
  → Use: Analogical Reasoning, Generation Effect
  → Focus: Flexibility, Transfer, Creative Application
ELSE IF learner is well-constrained:
  → Use: Any appropriate methodology
  → Focus: Advancement, Depth, Sophistication
```

## Methodology Selection Matrix

| Learning Goal | Novice | Developing | Proficient | Expert |
|---------------|--------|------------|------------|---------|
| **Conceptual** | Analogical + Examples | Feynman + Socratic | Elaborative + Transfer | Teaching + Innovation |
| **Procedural** | Worked Examples | Practice + Variation | Problem-Based | Optimization + Adaptation |
| **Strategic** | Guided Practice | Case Analysis | Multiple Approaches | Strategy Creation |
| **Transfer** | Near Transfer | Far Transfer | Creative Application | Domain Innovation |

## Methodology Profiles

### Feynman Technique
**Best for**: Testing and developing conceptual understanding
**Learner fit**: Developing to Proficient
**Constraint focus**: Understanding validation, explanation clarity
**Time needed**: 15-20 minutes per concept
**Prerequisites**: Basic familiarity with concept

### Socratic Dialogue  
**Best for**: Deep exploration, misconception correction
**Learner fit**: Developing to Expert (adapted complexity)
**Constraint focus**: Pattern recognition, boundary understanding
**Time needed**: 20-30 minutes
**Prerequisites**: Some prior knowledge to build on

### Elaborative Interrogation
**Best for**: Connecting concepts, understanding relationships
**Learner fit**: All levels (questions adapted)
**Constraint focus**: Causal understanding, constraint boundaries
**Time needed**: 10-15 minutes per connection
**Prerequisites**: Basic concept familiarity

### Analogical Reasoning
**Best for**: Introducing new concepts, overcoming rigidity
**Learner fit**: Novice to Developing
**Constraint focus**: Transfer patterns, structural similarity
**Time needed**: 15-25 minutes
**Prerequisites**: Familiar domain for comparison

### Problem-Based Learning
**Best for**: Application skills, strategic thinking
**Learner fit**: Developing to Expert
**Constraint focus**: Real-world application, constraint satisfaction
**Time needed**: 30-45 minutes
**Prerequisites**: Conceptual foundation

### Worked Examples
**Best for**: Procedural learning, pattern recognition
**Learner fit**: Novice to Developing
**Constraint focus**: Step-by-step constraint following
**Time needed**: 10-20 minutes per example
**Prerequisites**: Minimal - can start from basics

### Retrieval Practice
**Best for**: Memory consolidation, confidence building
**Learner fit**: All levels
**Constraint focus**: Memory constraint strengthening
**Time needed**: 5-15 minutes
**Prerequisites**: Prior exposure to material

### Interleaving
**Best for**: Discrimination, flexible application
**Learner fit**: Developing to Expert
**Constraint focus**: Adaptive constraint application
**Time needed**: 25-40 minutes
**Prerequisites**: Familiarity with multiple concepts

### Generation Effect
**Best for**: Creative application, ownership
**Learner fit**: Proficient to Expert
**Constraint focus**: Creative constraint satisfaction
**Time needed**: 20-30 minutes
**Prerequisites**: Solid understanding foundation

### Dual Coding (Visual-Verbal)
**Best for**: Complex concepts, spatial reasoning
**Learner fit**: All levels (adapted complexity)
**Constraint focus**: Multi-modal constraint patterns
**Time needed**: 15-25 minutes
**Prerequisites**: Varies by representation type

### Case Method
**Best for**: Strategic thinking, real-world application
**Learner fit**: Proficient to Expert
**Constraint focus**: Context-dependent constraint application
**Time needed**: 30-50 minutes
**Prerequisites**: Strong conceptual foundation

## Selection Algorithm

### Primary Selection Criteria
1. **Learning Objective Match** (40% weight)
2. **Learner Level Appropriateness** (30% weight)
3. **Constraint State Alignment** (20% weight)
4. **Time and Context Fit** (10% weight)

### Secondary Selection Criteria
- Previous methodology success
- Learner preferences and engagement
- Content complexity and type
- Available resources and tools
- Session energy and focus level

### Combination Strategies
- **Sequential**: Use multiple methodologies in planned sequence
- **Integrated**: Combine elements from different approaches
- **Adaptive**: Switch methodologies based on real-time response
- **Layered**: Build complexity using methodology progression

## Program Output Template

```markdown
## Methodology Selection

### Primary Recommendation: {Methodology Name}

### Selection Rationale:
- **Learning Objective**: {How this methodology serves the objective}
- **Learner Level**: {Why this fits learner's current understanding}
- **Constraint Alignment**: {How this addresses constraint state}
- **Time Efficiency**: {Expected time and pacing}

### Implementation Approach:
- **Setup**: {How to introduce the methodology}
- **Process**: {Key steps and flow}
- **Monitoring**: {What to watch for}
- **Adaptation**: {How to adjust if needed}

### Constraint Focus Areas:
- **Primary Constraints**: {Which atomic constraints to emphasize}
- **Validation Methods**: {How to check constraint satisfaction}
- **Adaptation Triggers**: {When to modify approach}

### Alternative Methodologies:
- **If struggling**: {Fallback approaches}
- **If advancing quickly**: {Extension approaches}
- **If disengaged**: {Re-engagement approaches}

### Zelda Coordination:
- **Capture readiness**: {When insights likely to emerge}
- **Note triggers**: {What to watch for}
- **Knowledge patterns**: {What connections to expect}

### Success Indicators:
- **Immediate**: {Session-level success signs}
- **Progressive**: {Multi-session progress indicators}
- **Constraint satisfaction**: {Evidence of constraint understanding}
```

## Integration with Constraint Architecture

### Atomic Constraint Considerations
- **Dialogue Patterns**: Which questioning styles fit methodology
- **Questioning Syntax**: How to structure methodology questions
- **Understanding Validation**: How methodology validates learning
- **Pacing Control**: How methodology affects cognitive load
- **Metacognitive Triggers**: How methodology builds awareness

### Molecular Protocol Loading
Based on selection, load appropriate protocol:
- Use Read tool: `.claude/resources/teacher/protocols/molecules/{methodology}.md`
- Follow protocol constraints and patterns
- Adapt protocol to learner state
- Monitor protocol effectiveness

### Cellular Memory Integration
- Update learning progression based on methodology success
- Note methodology preferences for this learner
- Track misconception patterns revealed
- Record adaptive strategies that work

### Organ Coordination
- Coordinate with Zelda for knowledge capture timing
- Share methodology choice with knowledge base
- Update session continuity with methodology context

## Edge Cases and Adaptations

### Methodology Conflicts
- **If learner requests specific methodology**: Assess fit and adapt
- **If preferred methodology doesn't fit**: Explain and negotiate
- **If multiple methodologies equally valid**: Let learner choose

### Real-Time Adaptation
- **If methodology not working**: Switch to alternative
- **If learner struggling**: Reduce complexity, add scaffolding
- **If learner bored**: Increase challenge, add depth

### Time Constraints
- **Short sessions**: Prioritize high-impact methodologies
- **Extended sessions**: Use methodology combinations
- **Interrupted sessions**: Choose resumable methodologies

### Individual Preferences
- **Visual learners**: Emphasize Dual Coding, Analogical
- **Verbal learners**: Emphasize Feynman, Socratic
- **Kinesthetic learners**: Emphasize Problem-Based, Generation

This program ensures systematic, evidence-based methodology selection that optimizes learning constraint satisfaction.