# Notification Timing & Optimization Guide

**Date:** January 22, 2025  
**Document Type:** Optimization Strategy  
**Status:** Final

## Executive Summary

This guide provides data-driven timing strategies for push notifications, including optimal send times, frequency caps, and contextual triggers to maximize engagement while minimizing fatigue.

## Optimal Send Times Heatmap

### Weekly Engagement Pattern

```
       Mon  Tue  Wed  Thu  Fri  Sat  Sun
6 AM   🟦   🟦   🟦   🟦   🟦   🟨   🟨
9 AM   🟩   🟩   🟩   🟩   🟩   🟦   🟦
12 PM  🟨   🟨   🟨   🟨   🟨   🟩   🟩
3 PM   🟦   🟦   🟦   🟦   🟨   🟩   🟩
6 PM   🟩   🟩   🟩   🟩   🟩   🟨   🟨
9 PM   🟨   🟨   🟨   🟨   🟦   🟦   🟦

Legend: 🟩 High (>40% CTR) | 🟨 Medium (20-40%) | 🟦 Low (<20%)
```

### Key Insights
- **Weekday Peaks:** 9 AM and 6 PM (commute/transition times)
- **Weekend Shift:** 12-3 PM highest (leisure time)
- **Dead Zones:** Late night (10 PM-6 AM), mid-afternoon weekdays

## Receipt Scanning Reminders

### Timing Strategy by Type

#### Real-Time Scanners (Zoe, Alex)
- **Trigger:** Location exit or transaction detection
- **Optimal Window:** 15 minutes post-purchase
- **Follow-up:** 2-3 hours if ignored (once only)
- **Success Rate:** 50% scan rate within first hour

#### Batch Processors (Sarah, Robert)
- **Sarah:** Sunday 5 PM (weekly batch)
- **Robert:** Daily 9 PM (end-of-day routine)
- **Marcus:** On-demand or bulk completion notification
- **Success Rate:** 70% completion when aligned with routine

### Expiration Reminders

```
Timeline for 30-Day Receipt Expiration:
Day 1-13:   No reminder (fresh in memory)
Day 14:     First gentle reminder
Day 23:     "One week left" warning
Day 29:     Final "expires tomorrow" alert
Day 30:     (No notification - too late)
```

### Special Event Timing

#### Point Multiplier Events
- **Announcement:** Friday 3 PM (planning phase)
- **Start:** Saturday 9 AM (weekend begins)
- **Mid-event:** Sunday 12 PM (peak activity)
- **Last Chance:** Sunday 6 PM (urgency)

## Streak & Habit Notifications

### Daily Streak Reminders

#### By Persona Timing
| Persona | Preferred Time | Rationale |
|---------|---------------|-----------|
| Sarah | Skip or 8 PM | Minimal intrusion |
| Zoe | 8-9 PM | Social media wind-down |
| Alex | 7 PM | Post-dinner routine |
| Marcus | None | Dislikes streak concept |
| Robert | 9 PM | End-of-day checklist |

### Milestone Celebrations

```
Immediate Triggers (Real-time):
- First receipt scanned
- Streak milestone reached
- Goal achieved
- Points threshold crossed

Delayed Triggers (Next optimal window):
- If achieved overnight → 8 AM notification
- If during work hours → 6 PM notification
```

### Time Traveler Temporal Alerts

#### Pattern-Based Timing
- **Past Patterns:** Monday mornings (weekly review)
- **Future Warnings:** 3-7 days before event
- **Timeline Shifts:** Within 1 hour of deviation
- **Achievement Approach:** 24-48 hours before

## Frequency Capping Rules

### Global Limits
- **Maximum per day:** 3 notifications
- **Minimum spacing:** 2 hours between non-critical
- **Weekly maximum:** 10-12 (persona dependent)

### Persona-Specific Caps

| Persona | Daily Max | Weekly Max | Preferred Frequency |
|---------|-----------|------------|-------------------|
| Sarah | 1 | 3-4 | Weekly batch |
| Zoe | 3-5 | 15-20 | Multiple daily |
| Alex | 2-3 | 10-12 | Daily with gaps |
| Marcus | 1 | 2-3 | Only critical |
| Robert | 1-2 | 7-8 | Daily summary |

### Smart Batching Rules

```
IF multiple events within 30 minutes:
  → Combine into single notification
  
EXAMPLE:
Instead of:
  - "Receipt processed: +100 points"
  - "Receipt processed: +100 points"
  - "Receipt processed: +100 points"
  
Send:
  - "3 receipts processed: +300 points total!"
```

## Contextual Triggers

### Location-Based (Opt-in)

#### Store Exit Detection
```
Trigger: User leaves retail location
Delay: 5-15 minutes
Message: "Just left [Store]? Scan your receipt for points!"
Limit: Max 2 per day
```

#### Home Arrival
```
Trigger: Arrive home 5-8 PM
Condition: Has unscanned receipts from today
Message: "Welcome home! Empty your wallet, scan receipts"
Limit: Max 3 per week
```

### Transaction-Based

#### Bank Transaction Posted
```
Trigger: New transaction via Plaid
Check: No matching receipt
Wait: 24 hours
Message: "Add receipt for your [Merchant] purchase?"
```

#### Unusual Spending
```
Trigger: Transaction 2x normal for category
Timing: Within 2 hours
Message: "Unusual charge detected. Everything look right?"
```

### Behavioral Patterns

#### Routine Detection
```
IF user always scans Sunday 7 PM:
  → Send reminder Sunday 6:45 PM
  
IF user checks app every morning:
  → Queue insights for 8 AM delivery
```

#### Engagement Decay
```
IF daily user becomes weekly:
  → Send re-engagement Thursday evening
  → Include Time Traveler "miss you" message
  
IF streak broken after 20+ days:
  → Wait 18 hours, send recovery offer
```

## Time Zone & Localization

### Technical Implementation
- Store user timezone on signup
- Update if device timezone changes
- All times in local timezone
- Daylight Saving Time handling

### International Considerations
```
US Users: 12-hour format (9:00 PM)
EU Users: 24-hour format (21:00)
Weekend definition: Locale-specific
Holiday awareness: Skip non-critical
```

## A/B Testing Framework

### Testing Variables

#### Timing Tests
- Morning (8 AM) vs Evening (8 PM) for summaries
- Immediate vs 30-min delay for receipts
- Weekday vs Weekend for challenges

#### Frequency Tests
- Daily vs Every-other-day for Alex
- Single vs Split updates for Robert
- High-frequency vs Batch for Zoe

#### Message Variants
- Time Traveler level (1-5) effectiveness
- Urgency language impact
- Emoji presence/absence

### Success Metrics
- Click-through rate (CTR)
- Action completion rate
- Time to action
- Opt-out rate
- Downstream engagement

## Machine Learning Optimization

### Predictive Timing Model

#### Input Features
- Historical open times
- Day of week patterns
- Recent engagement level
- Persona classification
- Time since last interaction

#### Output
- Optimal send time (hour)
- Engagement probability
- Message type recommendation

### Continuous Learning
```
Week 1-2: Use heatmap defaults
Week 3-4: Incorporate user specifics
Week 5+:  Fully personalized timing
```

## Seasonal Adjustments

### Key Periods

#### New Year (Jan 1-15)
- **Increase:** Goal-setting, motivation
- **Timing:** Morning aspirational messages
- **Frequency:** Higher tolerance

#### Tax Season (Mar-Apr)
- **Focus:** Receipt organization, categorization
- **Timing:** Weekend batch reminders
- **Urgency:** Deadline-driven

#### Summer (Jun-Aug)
- **Decrease:** Vacation mode awareness
- **Timing:** Flexible, less routine
- **Features:** Streak freeze promotions

#### Holidays (Nov-Dec)
- **Increase:** Spending tracking, budget alerts
- **Timing:** Pre-weekend shopping
- **Specials:** Gift budget tracking

## Implementation Checklist

### Technical Requirements
- [ ] Time zone detection and storage
- [ ] Scheduling system with queue
- [ ] Batching logic for multiple events
- [ ] Location services integration
- [ ] Delivery tracking and retry logic

### Analytics Setup
- [ ] CTR tracking by notification type
- [ ] Engagement funnel measurement
- [ ] A/B test infrastructure
- [ ] Persona segmentation tracking
- [ ] ML model training pipeline

### Content Management
- [ ] Template versioning system
- [ ] Dynamic content insertion
- [ ] Localization support
- [ ] Time Traveler personality levels
- [ ] Emergency override capability

## Best Practices Summary

### The Golden Rules

1. **Value Over Volume:** Every notification must provide value
2. **Respect Routines:** Align with user's natural patterns
3. **Smart Batching:** Combine related notifications
4. **Quick Escape:** Easy opt-out always available
5. **Learn & Adapt:** Use data to improve continuously

### Red Flags to Avoid

❌ Sending during sleep hours (11 PM - 6 AM)
❌ Multiple notifications within an hour
❌ Generic timing for all users
❌ Ignoring time zone differences
❌ Not respecting opt-out preferences

### Success Indicators

✅ CTR above 40% average
✅ Opt-out rate below 5%
✅ Engagement increase of 25%+
✅ Positive user feedback
✅ Improved retention metrics

## Appendix: Quick Reference

### Notification Priority Matrix

| Priority | Types | Max Delay | Override Quiet Hours |
|----------|-------|-----------|---------------------|
| Critical | Security, Fraud | None | Yes |
| High | Streak last chance | 1 hour | No |
| Medium | Achievements, Tips | 6 hours | No |
| Low | Social, Marketing | 24 hours | No |

### Default Schedule Template

```
Monday:    9 AM (Week ahead), 6 PM (Tip)
Tuesday:   6 PM (Progress check)
Wednesday: 9 AM (Motivation)
Thursday:  6 PM (Challenge update)
Friday:    3 PM (Weekend prep)
Saturday:  12 PM (Receipt reminder)
Sunday:    5 PM (Week wrap-up)
```