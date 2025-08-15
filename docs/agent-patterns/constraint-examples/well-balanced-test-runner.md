# Well-Balanced Agent Example: Test Runner

## Agent Definition
```yaml
---
name: test-runner
description: IMPORTANT runs tests and fixes failures after code changes
tools: Read, Bash
model: sonnet
---
```

## Constraint Analysis

### Diagnosis: Well-Balanced (Score: 8/10)

#### Strengths
1. **Clear Trigger**: "IMPORTANT" ensures proactive delegation
2. **Specific Action**: "runs tests and fixes failures"
3. **Contextual Activation**: "after code changes"
4. **Minimal Tools**: Only Read and Bash needed
5. **Appropriate Model**: Sonnet for moderate complexity

#### Balance Indicators
- ✅ Not too vague: Clear purpose stated
- ✅ Not too rigid: Works with any test framework
- ✅ Flexible execution: Can adapt to different test outputs
- ✅ Clear boundaries: Test execution and fixing
- ✅ Predictable behavior: Consistent test running

## Why It's Well-Balanced

### Avoids Under-Constraining
- Has specific trigger keyword (IMPORTANT)
- Defines clear action (run tests, fix failures)
- Specifies activation context (code changes)

### Avoids Over-Constraining
- Doesn't specify exact test commands
- Flexible about test frameworks
- Adaptable to different project structures
- No rigid output format requirements

## Pattern Application

Use this pattern when creating agents that:
- Need automatic activation
- Perform clear, bounded tasks
- Require minimal tools
- Should adapt to project context
- Have predictable but flexible behavior

## Learning Points
1. Start with IMPORTANT for proactive behavior
2. Use action verbs (runs, fixes)
3. Specify context without over-specifying
4. Select minimal sufficient tools
5. Choose model based on task complexity