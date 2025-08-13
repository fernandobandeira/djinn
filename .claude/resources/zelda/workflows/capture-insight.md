# Capture Insight Workflow

## Trigger Recognition
Watch for these learning moments:
- "Aha!" moments of sudden understanding
- Connections between previously unrelated concepts
- Successful problem-solving patterns
- Corrected misconceptions
- Useful analogies or mental models
- Generalizable principles from specific examples

## Capture Process

### Step 1: Identify the Atomic Idea
Ask yourself:
- What is the single core insight?
- Can it be understood independently?
- Is it specific enough to be useful?
- Is it general enough to connect to other ideas?

### Step 2: Check for Existing Notes
```bash
# Search for related concepts
grep -r "[concept keywords]" zettelkasten/

# Look for similar patterns
grep -r "[pattern description]" zettelkasten/permanent/
```

### Step 3: Create the Note

#### For New Insights
1. Generate unique ID: `YYYYMMDD-HHMMSS`
2. Choose clear, searchable title
3. Write self-contained explanation
4. Add concrete example
5. List initial connections

#### For Extensions
1. Reference original note
2. Explain what's new or different
3. Update original with forward link
4. Consider if this changes other notes

### Step 4: Establish Connections

#### Backward Links (This builds on...)
- Find foundational concepts
- Add "Builds on:" section
- Update those notes with forward links

#### Lateral Links (This relates to...)
- Find similar patterns in other domains
- Add "Related to:" section
- Create bidirectional links

#### Forward Links (This leads to...)
- Consider implications
- Add "Enables:" section
- Mark areas for future exploration

### Step 5: Enrich the Note

#### Add Examples
- Concrete application
- Code snippet
- Real-world scenario
- Visual representation

#### Add Questions
- What's still unclear?
- What assumptions are we making?
- How might this fail?
- Where else might this apply?

### Step 6: Integrate into System

#### Update Relevant Hubs
- Add to topic hub if exists
- Create new hub if cluster emerging
- Update learning path if applicable

#### Schedule for Review
- Add to spaced repetition queue
- Set initial review for tomorrow
- Mark maturity level (seedling)

## Quality Checklist

### Before Saving
- [ ] Single, clear idea?
- [ ] Self-contained explanation?
- [ ] Meaningful title?
- [ ] At least one connection?
- [ ] Includes example?
- [ ] Poses further question?
- [ ] Tagged appropriately?

### After Creation
- [ ] Backlinks created?
- [ ] Hub updated?
- [ ] Review scheduled?
- [ ] Related notes updated?

## Common Patterns

### Pattern: Generalization
When you see a specific example that represents a broader principle:
1. Create note for the principle
2. Link to specific example
3. Search for other examples
4. Create pattern note if multiple found

### Pattern: Contradiction
When new information conflicts with existing notes:
1. Create note for new perspective
2. Create comparison note
3. Update original note with "Contrasts with:"
4. Consider synthesis possibility

### Pattern: Synthesis
When multiple notes combine to form new understanding:
1. Create synthesis note
2. Link to all component notes
3. Explain the emergence
4. Update components with "Contributes to:"

## Automation Helpers

### Quick Capture Script
```bash
#!/bin/bash
# Quick capture for Zettelkasten

TIMESTAMP=$(date +%Y%m%d-%H%M%S)
TITLE="$1"
INSIGHT="$2"

cat > "zettelkasten/permanent/${TIMESTAMP}-${TITLE// /-}.md" << EOF
# ${TIMESTAMP} ${TITLE}

## Core Insight
${INSIGHT}

## Context
- Source: Current learning session
- Date: $(date +%Y-%m-%d)

## Connections
- Related to: [[pending]]

## Tags
#captured #pending-review
EOF

echo "Created note: ${TIMESTAMP}-${TITLE// /-}.md"
```

### Link Finder Script
```bash
#!/bin/bash
# Find potential connections

KEYWORDS="$1"
echo "Searching for connections to: ${KEYWORDS}"

# Search in permanent notes
echo "=== Permanent Notes ==="
grep -l "${KEYWORDS}" zettelkasten/permanent/*.md | head -10

# Search in hubs
echo "=== Hub Notes ==="
grep -l "${KEYWORDS}" zettelkasten/hubs/*.md | head -5
```

## Remember
- Capture immediately when insight occurs
- Don't overthink - can refine later
- Focus on atomic ideas
- Always create at least one connection
- Trust the process - patterns emerge naturally