# Agent Quality Review Checklist

## Overview
Comprehensive review to ensure agent meets quality standards and best practices.

## Pre-Review Information
- **Agent Name**: _______________
- **Agent Type**: [ ] Command | [ ] Sub-Agent
- **Reviewer**: _______________
- **Date**: _______________
- **Version**: _______________

## Core Quality Checks

### Purpose & Scope
- [ ] **Single Responsibility**: Agent has ONE clear purpose
- [ ] **Well-Defined**: Purpose is clearly stated
- [ ] **Unique Value**: Doesn't duplicate existing agents
- [ ] **Appropriate Type**: Correct choice between command/sub-agent

### Structural Quality

#### For Command Agents
- [ ] Located in `.claude/commands/`
- [ ] Follows command file structure
- [ ] Has activation section
- [ ] Commands use `*` prefix
- [ ] Resource files properly referenced
- [ ] Lazy loading implemented

#### For Sub-Agents
- [ ] Located in `.claude/agents/`
- [ ] Valid frontmatter YAML
- [ ] Description starts with trigger phrase
- [ ] Minimal tools selected
- [ ] Appropriate model chosen
- [ ] No interaction loops

### Documentation Quality
- [ ] **Purpose Statement**: Clear and concise
- [ ] **Usage Instructions**: How to activate/use
- [ ] **Examples Provided**: At least one example
- [ ] **Constraints Documented**: Limitations clear
- [ ] **Output Format**: Expected results defined

## Functional Quality

### Behavior Validation
- [ ] **Activation Works**: Agent starts correctly
- [ ] **Commands Function**: All commands work (if applicable)
- [ ] **Delegation Triggers**: Sub-agent delegates properly (if applicable)
- [ ] **Error Handling**: Graceful failure modes
- [ ] **Edge Cases**: Handles unusual inputs

### Resource Quality
- [ ] **Resources Load**: All referenced files exist
- [ ] **Paths Correct**: File paths work
- [ ] **Content Relevant**: Resources match purpose
- [ ] **No Redundancy**: No duplicate resources
- [ ] **Well Organized**: Logical structure

## Best Practices Compliance

### Design Patterns
- [ ] **Pattern Consistent**: Follows established patterns
- [ ] **Navigation Pattern**: Uses tree/ls/grep for discovery
- [ ] **Interactive Design**: Engages user appropriately
- [ ] **Numbered Options**: Uses 1-9 for choices (commands)
- [ ] **Progressive Disclosure**: Information revealed as needed

### Code Quality
- [ ] **Clean Syntax**: No syntax errors
- [ ] **Consistent Formatting**: Uniform style
- [ ] **Readable Code**: Clear and understandable
- [ ] **Comments Appropriate**: Where needed only
- [ ] **No Dead Code**: All code serves purpose

### Performance
- [ ] **Efficient Tools**: Minimal tool usage (sub-agents)
- [ ] **Fast Response**: Reasonable execution time
- [ ] **Context Efficient**: Doesn't waste context window
- [ ] **Lazy Loading**: Resources load only when needed

## User Experience

### Interaction Quality
- [ ] **Clear Communication**: Messages understandable
- [ ] **Helpful Errors**: Error messages guide resolution
- [ ] **Intuitive Flow**: Natural progression
- [ ] **Consistent Voice**: Personality maintained
- [ ] **Appropriate Tone**: Matches purpose

### Usability
- [ ] **Easy to Start**: Clear activation method
- [ ] **Predictable Behavior**: Does what expected
- [ ] **Helpful Output**: Provides value
- [ ] **Clean Exit**: Proper completion/closure
- [ ] **Feedback Provided**: User knows what's happening

## Technical Validation

### Integration
- [ ] **File Locations**: Correct directories used
- [ ] **Tool Access**: Can use specified tools
- [ ] **Resource Access**: Can read resource files
- [ ] **Output Generation**: Creates expected outputs

### Testing Results
- [ ] **Basic Flow Tested**: Primary use case works
- [ ] **Error Cases Tested**: Failures handled
- [ ] **Edge Cases Tested**: Unusual inputs work
- [ ] **Performance Tested**: Acceptable speed
- [ ] **Output Verified**: Results are correct

## Security & Safety

### Security Checks
- [ ] **No Hardcoded Secrets**: No keys/passwords
- [ ] **Safe File Operations**: Proper path handling
- [ ] **Input Validation**: User input sanitized
- [ ] **No Dangerous Commands**: Safe operations only

### Safety Checks
- [ ] **Clear Boundaries**: Limitations stated
- [ ] **Fail Safe**: Errors don't cause damage
- [ ] **Reversible Actions**: Can undo if needed
- [ ] **User Confirmation**: For significant actions

## Comparison with Existing Agents

### Differentiation
- [ ] **Unique Purpose**: Clear distinction from others
- [ ] **No Feature Creep**: Stays in its lane
- [ ] **Complements Others**: Works well with ecosystem
- [ ] **Clear Use Cases**: When to use is obvious

## Final Assessment

### Critical Issues (Must Fix)
- [ ] ⚠️ No critical issues found

List any critical issues:
1. _______________________
2. _______________________

### Major Issues (Should Fix)
- [ ] ⚠️ No major issues found

List any major issues:
1. _______________________
2. _______________________

### Minor Issues (Nice to Fix)
- [ ] 💡 No minor issues found

List any minor issues:
1. _______________________
2. _______________________

## Quality Score

Rate each category (1-5):
- **Purpose Clarity**: ___/5
- **Implementation**: ___/5
- **Documentation**: ___/5
- **User Experience**: ___/5
- **Best Practices**: ___/5

**Overall Score**: ___/25

## Review Decision

[ ] ✅ **Approved** - Ready for use
[ ] 🔄 **Needs Revision** - Address issues and re-review
[ ] ❌ **Rejected** - Major redesign needed

## Recommendations

### Improvements Suggested:
1. _______________________
2. _______________________
3. _______________________

### Commendations:
1. _______________________
2. _______________________

## Follow-Up Actions
- [ ] Issues documented
- [ ] Creator notified
- [ ] Fixes implemented
- [ ] Re-review scheduled (if needed)
- [ ] Documentation updated

## Sign-Off
**Reviewer Signature**: _______________ 
**Date**: _______________
**Next Review Date**: _______________

## Notes
[Additional comments or context]

---

## Quick Review (Abbreviated)

For rapid reviews, use this shortened version:

### Must Have ✅
- [ ] Clear single purpose
- [ ] Works as intended
- [ ] No duplicate functionality
- [ ] Proper file location
- [ ] Basic error handling

### Should Have 📝
- [ ] Good documentation
- [ ] Examples provided
- [ ] Consistent patterns
- [ ] Clean code
- [ ] Efficient implementation

### Nice to Have 💡
- [ ] Comprehensive testing
- [ ] Advanced error handling
- [ ] Performance optimized
- [ ] Extensive examples
- [ ] Perfect formatting

**Quick Decision**: [ ] Pass | [ ] Revise | [ ] Fail