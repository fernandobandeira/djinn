# Common Patterns for Agent Creation

## Pattern 1: Knowledge Base First
Always check what exists before creating anything new.

### Implementation:
```markdown
When user requests `*[command]`:
1. **FIRST search knowledge base**:
   - `./.vector_db/kb search "[topic]" --collection documentation`
   - `./.vector_db/kb search "[related]" --collection architecture`
2. Review findings
3. THEN proceed with task
```

### Why It Works:
- Builds on existing knowledge
- Avoids duplication
- Maintains consistency
- Leverages past decisions

## Pattern 2: Lazy Loading
Load resources only when actually needed, not at startup.

### Implementation:
```markdown
### Resource Loading Protocol
Only load resources when specific commands are invoked:
- Do NOT preload all files
- Load task files only when that task is requested
- Use Read tool to load files: Read .claude/resources/[agent]/...
```

### Why It Works:
- Conserves context window
- Faster startup
- More efficient
- Cleaner command file

## Pattern 3: Interactive Dialogue
Engage with users, don't generate at them.

### Implementation:
```markdown
1. Ask focused question
2. Wait for user response
3. Build on their answer
4. Guide to solution
```

### Why It Works:
- User feels involved
- Better outcomes
- Catches misunderstandings
- Builds trust

## Pattern 4: Numbered Options
Present choices in a consistent, scannable format.

### Implementation:
```markdown
Please select an option:
1. [First option]
2. [Second option]
3. [Third option]
...
9. [Ninth option]
0. Proceed/Cancel

Your choice (0-9):
```

### Why It Works:
- Easy to scan
- Quick selection
- Consistent UX
- Clear navigation

## Pattern 5: Brownfield Awareness
Always assume the system exists and build on it.

### Implementation:
```markdown
1. Check existing system state
2. Identify what needs change
3. Plan incremental evolution
4. Document changes as ADRs
5. Provide migration path
```

### Why It Works:
- Realistic approach
- Preserves what works
- Reduces risk
- Enables rollback

## Pattern 6: Document Everything
Every significant output should be captured and indexed.

### Implementation:
```markdown
After completing [task]:
1. Save to appropriate location
2. Index in knowledge base:
   - `./.vector_db/kb index --path ./docs/[location]/`
3. Update relevant documentation
```

### Why It Works:
- Knowledge accumulates
- Team learning
- Future reference
- Audit trail

## Pattern 7: Clear Separation of Concerns
Each resource type has a specific purpose.

### Structure:
```
tasks/      → Complex multi-step workflows
templates/  → Standardized output formats
data/       → Reference information
checklists/ → Validation and review
```

### Why It Works:
- Easy to find things
- Clear responsibilities
- Maintainable
- Scalable

## Pattern 8: Progressive Disclosure
Start simple, add detail as needed.

### Implementation:
```markdown
1. High-level options first
2. Drill down based on selection
3. Provide detail when relevant
4. Keep context focused
```

### Why It Works:
- Not overwhelming
- User-driven depth
- Efficient interaction
- Clear navigation

## Pattern 9: Consistent Commands
Standard commands across all agents.

### Core Set:
```markdown
- `*help` - Show available commands
- `*status` - Show current state
- `*exit` - Leave agent mode
- `*kb-search` - Search knowledge base
- `*kb-index` - Index outputs
```

### Why It Works:
- Predictable UX
- Easy to learn
- Muscle memory
- Less documentation

## Pattern 10: Error Recovery
Graceful handling of problems.

### Implementation:
```markdown
If error occurs:
1. Clear error message
2. Explain what went wrong
3. Suggest how to fix
4. Provide alternative path
```

### Why It Works:
- User not stuck
- Learning opportunity
- Maintains flow
- Builds confidence

## Anti-Patterns to Avoid

### Anti-Pattern 1: Everything Agent
Trying to do too much in one agent.

**Problem**: Confusing, hard to maintain
**Solution**: Single responsibility principle

### Anti-Pattern 2: Blind Generation
Creating without checking existing work.

**Problem**: Duplication, inconsistency
**Solution**: Always search KB first

### Anti-Pattern 3: Resource Dump
Loading everything at startup.

**Problem**: Wastes context, slow
**Solution**: Lazy loading pattern

### Anti-Pattern 4: Monologue Mode
Talking at users without interaction.

**Problem**: Poor outcomes, frustration
**Solution**: Interactive dialogue

### Anti-Pattern 5: Mystery Commands
Unclear what commands do.

**Problem**: Users don't know when to use
**Solution**: Clear, descriptive names

## Pattern Combinations

### Analysis Pattern (Ana Style):
```
KB Search → Interactive Facilitation → Multiple Techniques → Document Output → Index Results
```

### Architecture Pattern (Archie Style):
```
KB Search → Understand Current → Present Options → Document Decision → Create ADR → Index
```

### Creation Pattern (Rita Style):
```
Validate Need → Check Conflicts → Define Purpose → Create Structure → Test → Document
```

## Pattern Evolution

### Adding to Patterns:
1. Identify repeated successful approach
2. Document the pattern
3. Explain why it works
4. Show implementation
5. Add to this guide

### Modifying Patterns:
1. Document reason for change
2. Update all agents using pattern
3. Test thoroughly
4. Communicate change

### Deprecating Patterns:
1. Mark as deprecated
2. Explain why obsolete
3. Provide alternative
4. Update agents gradually

## Remember
- Patterns are proven solutions
- Consistency matters
- Patterns can evolve
- Document new patterns
- Share learnings