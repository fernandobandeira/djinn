# Spaced Repetition Workflow

## Overview
Implement scientifically-proven spaced repetition intervals to maximize long-term retention of Zettelkasten notes.

## Review Intervals

### Standard Schedule (SM-2 Based)
- **New Note**: Review after 1 day
- **After 1st Review**: 
  - Success → 3 days
  - Struggle → 1 day
- **After 2nd Review**:
  - Success → 7 days
  - Struggle → 3 days
- **After 3rd Review**:
  - Success → 14 days
  - Struggle → 7 days
- **After 4th Review**:
  - Success → 30 days
  - Struggle → 14 days
- **After 5th Review**:
  - Success → 90 days
  - Struggle → 30 days

## Review Process

### Step 1: Daily Review Queue
```bash
# Find notes due for review today
find zettelkasten/reviews/ -name "*.md" -mtime +0 -mtime -1

# Or check review metadata
grep -l "review_date: $(date +%Y-%m-%d)" zettelkasten/**/*.md
```

### Step 2: Review Methods

#### Method 1: Free Recall
1. Read only the note title
2. Try to recall:
   - Core insight
   - Key connections
   - Main example
3. Check accuracy
4. Rate difficulty (1-5)

#### Method 2: Elaboration
1. Read the note
2. Explain in different words
3. Generate new example
4. Find new connection
5. Update note if enhanced

#### Method 3: Application
1. Read the note
2. Apply to current problem
3. Test understanding
4. Document application
5. Create new note if insight emerges

### Step 3: Update Review Metadata

#### Success Case (Rating 4-5)
```yaml
last_reviewed: [today]
review_count: [increment]
next_review: [next interval]
ease_factor: [increase by 0.1]
understanding: [growing|established]
```

#### Struggle Case (Rating 1-3)
```yaml
last_reviewed: [today]
review_count: [increment]
next_review: [shorter interval]
ease_factor: [decrease by 0.2]
understanding: [needs work]
needs_elaboration: true
```

## Review Types

### Quick Review (2-3 minutes)
- Read title and try to recall
- Scan for accuracy
- Update review date
- Move to next note

### Standard Review (5-7 minutes)
- Read completely
- Test understanding
- Add new connection if found
- Update example if better one exists
- Adjust metadata

### Deep Review (10-15 minutes)
- Full elaboration
- Search for new connections
- Update related notes
- Consider synthesis opportunities
- Refine expression for clarity

## Automated Review System

### Review Queue Generator
```python
#!/usr/bin/env python3
import os
import yaml
from datetime import datetime, timedelta

def generate_review_queue():
    today = datetime.now().date()
    queue = []
    
    for root, dirs, files in os.walk('zettelkasten'):
        for file in files:
            if file.endswith('.md'):
                filepath = os.path.join(root, file)
                with open(filepath, 'r') as f:
                    content = f.read()
                    # Extract metadata
                    if '---' in content:
                        meta = yaml.safe_load(
                            content.split('---')[1]
                        )
                        if 'next_review' in meta:
                            review_date = datetime.strptime(
                                meta['next_review'], 
                                '%Y-%m-%d'
                            ).date()
                            if review_date <= today:
                                queue.append({
                                    'file': filepath,
                                    'overdue': (today - review_date).days
                                })
    
    # Sort by how overdue
    queue.sort(key=lambda x: x['overdue'], reverse=True)
    return queue

# Generate today's queue
queue = generate_review_queue()
print(f"📚 Today's Review Queue: {len(queue)} notes")
for item in queue[:10]:  # Show top 10
    print(f"  - {item['file']} (overdue: {item['overdue']} days)")
```

### Review Performance Tracker
```python
def track_review_performance(note_id, rating):
    """Update intervals based on performance"""
    intervals = {
        1: [1, 1, 2, 3, 5, 8],      # Very hard
        2: [1, 2, 3, 5, 8, 13],      # Hard
        3: [1, 3, 7, 14, 30, 60],    # Medium
        4: [1, 3, 7, 14, 30, 90],    # Easy
        5: [2, 7, 14, 30, 90, 180]   # Very easy
    }
    
    # Get current review count
    # Select appropriate interval
    # Update note metadata
    # Schedule next review
```

## Review Strategies

### For Concept Notes
1. Explain the concept without looking
2. Give three examples
3. Name two connections
4. Identify one application
5. Rate understanding

### For Process Notes
1. List the steps from memory
2. Explain why each step matters
3. Identify potential failures
4. Describe optimization opportunities
5. Rate confidence

### For Connection Notes
1. Explain the relationship
2. Provide evidence for connection
3. Consider counter-examples
4. Explore implications
5. Rate insight depth

## Gamification Elements

### Streak Tracking
- Daily review streak counter
- Weekly review goals
- Monthly mastery targets

### Understanding Levels
- 🌱 Seedling (0-2 reviews)
- 🌿 Sprout (3-5 reviews)
- 🌲 Growing (6-10 reviews)
- 🌳 Established (11-20 reviews)
- 🏛️ Permanent (20+ reviews)

### Achievement Milestones
- First Week: Review 7 days straight
- Knowledge Builder: 100 notes reviewed
- Connection Master: 50 new connections found
- Synthesis Expert: 10 synthesis notes created

## Review Session Structure

### Morning Review (15 minutes)
1. **Warm-up** (2 min): Quick scan of 3 easy notes
2. **Challenge** (8 min): Deep review of 2 difficult notes
3. **Connections** (3 min): Find links between reviewed notes
4. **Planning** (2 min): Note areas needing exploration

### Evening Review (10 minutes)
1. **Recall** (5 min): Test memory of morning's notes
2. **New** (3 min): Review today's captured notes
3. **Schedule** (2 min): Set tomorrow's review queue

## Quality Metrics

### Review Quality Indicators
- Can explain without reading
- Can generate new examples
- Can identify non-obvious connections
- Can apply to current problems
- Can teach to others

### Note Maturity Indicators
- Stable understanding across reviews
- Rich connection network
- Multiple examples and applications
- Clear, refined expression
- Generates new insights

## Remember
- Consistency beats intensity
- Trust the intervals
- Struggle is part of learning
- Connections matter more than memorization
- Understanding deepens with each review