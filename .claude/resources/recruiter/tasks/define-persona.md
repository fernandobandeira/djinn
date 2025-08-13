# Define Persona Task

## Purpose
Create a compelling, consistent personality for the agent that guides all interactions.

## Persona Components

### 1. Identity
Define WHO the agent is:
- Name and nickname
- Role/title
- Expertise areas
- Background/experience implied

### 2. Personality Traits
Define HOW the agent behaves:
- Communication style
- Emotional tone
- Level of formality
- Unique quirks or characteristics

### 3. Core Principles
Define WHAT the agent believes:
- Guiding philosophies
- Non-negotiable standards
- Priority hierarchy
- Decision-making framework

### 4. Interaction Style
Define HOW the agent engages:
- Greeting approach
- Question-asking style
- Response patterns
- Closure methods

## Persona Development Process

### Step 1: Understand the Purpose
Ask the user:
1. What is the agent's primary job?
2. Who will interact with this agent?
3. What feeling should users have?
4. What makes this agent unique?

### Step 2: Create Identity
Based on purpose, establish:
```yaml
identity:
  name: [Full Name]
  nickname: [Short name used in conversation]
  role: [Official title]
  expertise: [List of specialties]
  backstory: [Optional background]
```

### Step 3: Define Style
Choose personality attributes:
```yaml
style:
  formality: [casual|professional|academic]
  energy: [calm|enthusiastic|balanced]
  approach: [directive|facilitative|collaborative]
  traits: [List 3-5 key traits]
```

### Step 4: Establish Principles
Core beliefs that guide behavior:
```yaml
principles:
  - [Principle]: [How it affects behavior]
  - [Principle]: [How it affects behavior]
  - [Principle]: [How it affects behavior]
```

### Step 5: Craft Voice
Examples of how agent speaks:
- Greeting: "Hello! I'm [Name], your [role]..."
- Offering help: "I can help you [action]..."
- Asking questions: "Could you tell me more about..."
- Confirming: "Let me make sure I understand..."
- Closing: "Is there anything else..."

## Persona Patterns by Type

### Facilitator Pattern (like Ana)
```yaml
persona:
  style: Curious, encouraging, creative
  approach: Guide don't generate
  interaction: Ask, listen, build
  principles:
    - User empowerment
    - Creative exploration
    - Iterative refinement
```

### Expert Pattern (like Archie)
```yaml
persona:
  style: Pragmatic, thorough, systematic
  approach: Analyze then recommend
  interaction: Assess, present options, document
  principles:
    - Best practices first
    - Consider trade-offs
    - Document decisions
```

### Instructor Pattern (like Rita)
```yaml
persona:
  style: Methodical, instructive, pattern-aware
  approach: Teach while doing
  interaction: Explain, demonstrate, validate
  principles:
    - Education through action
    - Pattern consistency
    - Quality enforcement
```

### Assistant Pattern
```yaml
persona:
  style: Helpful, efficient, responsive
  approach: Execute requests quickly
  interaction: Acknowledge, perform, report
  principles:
    - Speed and accuracy
    - Clear communication
    - Task completion
```

## Consistency Guidelines

### Maintain Throughout:
1. **Vocabulary**: Use consistent terminology
2. **Tone**: Keep same energy level
3. **Formatting**: Consistent output style
4. **Patterns**: Repeated interaction flows
5. **Boundaries**: Clear about capabilities

### Avoid:
- Personality shifts mid-conversation
- Contradicting established principles
- Breaking character
- Inconsistent formatting
- Mixed communication styles

## Testing Persona Fit

### Questions to Validate:
1. Does the persona match the purpose?
2. Is it distinct from other agents?
3. Will users understand the role?
4. Is the voice consistent?
5. Are principles actionable?

### Test Interactions:
- Initial greeting
- Handling errors
- Complex requests
- Edge cases
- Session closing

## Persona Documentation Template

```markdown
## Persona

### Identity
[Name] is [role description]. [Backstory if relevant].

### Style
- **Communication**: [How they speak]
- **Approach**: [How they work]
- **Energy**: [Their demeanor]

### Core Principles
1. **[Principle]**: [Application]
2. **[Principle]**: [Application]
3. **[Principle]**: [Application]

### Interaction Examples
- Greeting: "[Sample greeting]"
- Question: "[Sample question]"
- Confirmation: "[Sample confirmation]"
```

## Remember
- Persona drives ALL interactions
- Consistency builds trust
- Personality makes agents memorable
- Principles guide decisions
- Voice should be distinctive