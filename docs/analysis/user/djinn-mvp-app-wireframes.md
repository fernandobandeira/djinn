# Djinn MVP Main App Wireframes
*Version 1.0 - Core Application Functionality*

Generated: 2025-01-21
Status: MVP App Specification
Platform: Mobile (iOS/Android)

## Overview
This document contains the main application screens (11-18) from the Djinn personal finance app MVP, following the onboarding flow. These screens represent the core user experience after setup is complete.

**Tagline**: "Your Money Wishes Granted"

## Brand & Visual Identity
- **Primary Color**: Emerald Green (#10B981)
- **Secondary**: Purple accents (#8B5CF6)
- **Mascot**: Magic lamp with animated liquid (no separate genie)
- **Design Balance**: 70% clean white UI, 20% gradients, 10% liquid glass effects
- **Typography**: Inter (primary), Playfair Display (headers)

## Core Application Screens

### 11. Main Dashboard (Regular)
```
┌─────────────────────────────────┐
│ 🏦 All Accounts    [Lamp] 600pts│ ← Points balance visible
│                         🔥 3     │ ← Streak counter
├─────────────────────────────────┤
│                                 │
│  Good morning, Sarah! ☀️       │
│                                 │
│  ┌────────────────────────┐    │
│  │ Total Balance          │    │ ← Primary card
│  │ $12,847.23            │    │
│  │ ↑ $523 from last month │    │ ← Green for positive
│  └────────────────────────┘    │
│                                 │
│  This Month's Spending          │
│  ┌────────────────────────┐    │
│  │ 📊 $2,147 of $3,000    │    │ ← Budget progress
│  │ ████████░░░░ 71%       │    │ ← Emerald bar
│  │                        │    │
│  │ Top Categories:         │    │
│  │ 🍔 Food: $487          │    │
│  │ 🚗 Transport: $324     │    │
│  │ 🏠 Bills: $892         │    │
│  └────────────────────────┘    │
│                                 │
│  Quick Actions                  │
│  ┌──────┐ ┌──────┐ ┌──────┐  │
│  │ ✨  │ │ 📷  │ │ 💰  │  │ ← Action buttons
│  │ Wish │ │Receipt│ │Budget│  │
│  └──────┘ └──────┘ └──────┘  │
│                                 │
├─────────────────────────────────┤
│ [🏠] [💳] [✨] [📊] [👤]      │ ← Tab bar
└─────────────────────────────────┘
```

**Key Features**:
- Points balance always visible
- Streak counter for engagement
- Personalized greeting with time of day
- Account aggregation selector
- Quick access to primary actions
- Visual budget progress indicator

### 12. AI Wishes Screen
```
┌─────────────────────────────────┐
│ ← Back      Points: 1,250 🪙   │
│                                 │
│      "Ask Your Genie ✨"        │
│                                 │
│    [Lamp Idle Animation]        │
│                                 │
│  ┌────────────────────────┐    │
│  │ Type your question...   │    │ ← Natural language input
│  └────────────────────────┘    │
│                                 │
│  Popular Questions:             │
│  ┌────────────────────────┐    │
│  │ "What am I wasting      │    │ ← Direct questions
│  │  money on?"             │    │
│  │ 💎 500 points           │    │ ← Credit cost
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ "How much did I spend   │    │
│  │  on Uber this month?"   │    │
│  │ 💎 500 points           │    │
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ "Build me a budget to   │    │
│  │  save $500/month"       │    │
│  │ 🪙 2000 credits         │    │ ← Premium wish
│  └────────────────────────┘    │
│                                 │
│  Recent Wishes:                 │
│  • "Show my subscriptions" ✓    │
│  • "Coffee spending trend" ✓    │
│                                 │
│  [Get More Credits] ←───────    │ ← Link to store
└─────────────────────────────────┘
```

**Key Features**:
- Natural language AI interface
- Clear point costs per query
- Popular question templates
- Recent query history
- Direct link to purchase points

### 13. Wish Processing
```
┌─────────────────────────────────┐
│                                 │
│    [Lamp Working Animation]     │ ← Liquid churning fast
│      ✨ ✨ ✨ ✨              │ ← Sparkles emerging
│                                 │
│    "Analyzing your spending"    │
│         "patterns..."           │
│                                 │
│    [Progress Bar ████░░░]      │ ← Animated fill
│                                 │
│   "Looking at 847 transactions" │
│   "across 3 accounts..."        │
│                                 │
│                                 │
│         [Cancel]                │
└─────────────────────────────────┘
```

**Key Features**:
- Engaging lamp animation during processing
- Progress indication
- Transparency about data being analyzed
- Cancel option for long queries

### 14. Wish Results
```
┌─────────────────────────────────┐
│ ← Back                Share →   │
│                                 │
│    [Lamp Success State]         │ ← Satisfied glow
│                                 │
│   "Here's How to Save $50"      │
│                                 │
│  ┌────────────────────────┐    │
│  │ 1. Coffee: $47/month    │    │ ← Insight cards
│  │    Make coffee at home  │    │
│  │    3 days/week = $28    │    │
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ 2. Subscriptions: $23   │    │
│  │    Cancel Hulu (unused) │    │
│  │    Last watched: 47 days│    │
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ 3. Dining: $156/month   │    │
│  │    One less dinner out   │    │
│  │    Average meal: $43     │    │
│  └────────────────────────┘    │
│                                 │
│  Total Savings: $94/month 🎉   │
│                                 │
│   [Ask Another Wish] ←──────    │
└─────────────────────────────────┘
```

**Key Features**:
- Actionable insights with specific amounts
- Clear savings opportunities
- Evidence-based recommendations
- Share functionality for social engagement
- Total impact summary

### 15. Transactions Screen
```
┌─────────────────────────────────┐
│ ← Back         Filter: All ▼    │
│                                 │
│        Transactions             │
│                                 │
│  Today                          │
│  ┌────────────────────────┐    │
│  │ Starbucks              │    │
│  │ Coffee & Snacks  ☕    │    │
│  │ -$12.47         10:32am│    │
│  │ [📷 Add Receipt]       │    │ ← CTA on each transaction
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ Shell Gas Station      │    │
│  │ Transportation  ⛽      │    │
│  │ -$45.23         8:15am │    │
│  │ [📷 Add Receipt]       │    │
│  └────────────────────────┘    │
│                                 │
│  Yesterday                      │
│  ┌────────────────────────┐    │
│  │ Netflix               │    │
│  │ Subscription  📺  [🔄] │    │ ← Recurring indicator
│  │ -$15.99               │    │
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ Whole Foods           │    │
│  │ Groceries  🛒         │    │
│  │ -$127.84              │    │
│  │ ✅ Receipt attached    │    │ ← Already has receipt
│  └────────────────────────┘    │
│                                 │
├─────────────────────────────────┤
│ [🏠] [💳] [✨] [📊] [👤]      │
└─────────────────────────────────┘
```

**Key Features**:
- Chronological transaction list
- Category icons for quick scanning
- Receipt attachment CTAs
- Recurring transaction indicators
- Filter capabilities

### 15A. Adding Receipt Photo
```
┌─────────────────────────────────┐
│ ← Cancel                        │
│                                 │
│      Add Receipt Photo          │
│        Starbucks               │
│                                 │
│  ┌────────────────────────┐    │
│  │                        │    │
│  │     [Camera View]      │    │ ← Live camera
│  │                        │    │
│  │    [Receipt Guide]     │    │ ← Overlay guide
│  │     ┌──────────┐       │    │
│  │     │          │       │    │
│  │     │  Align   │       │    │
│  │     │ Receipt  │       │    │
│  │     │   Here   │       │    │
│  │     └──────────┘       │    │
│  │                        │    │
│  └────────────────────────┘    │
│                                 │
│  Tips for best results:         │
│  • Good lighting               │
│  • Flat receipt                │
│  • All text visible            │
│                                 │
│  [📷 Capture] ←─────────        │ ← Main action
│  [📁 Choose from Gallery]       │ ← Alternative
│                                 │
└─────────────────────────────────┘
```

**Key Features**:
- Camera integration with guide overlay
- Clear capture instructions
- Gallery option for existing photos
- Transaction context (Starbucks)

### 15B. Receipt Captured (Matched)
```
┌─────────────────────────────────┐
│ ← Retake              Save →    │
│                                 │
│      Receipt Captured!          │
│                                 │
│  ┌────────────────────────┐    │
│  │ [Receipt Image Preview] │    │
│  │                        │    │
│  │  STARBUCKS             │    │
│  │  ---------------       │    │
│  │  Venti Latte    $6.45  │    │
│  │  Cake Pop       $3.50  │    │
│  │  Banana         $2.52  │    │
│  │  ---------------       │    │
│  │  Total:        $12.47  │    │
│  │                        │    │
│  └────────────────────────┘    │
│                                 │
│  ✅ Matched: Starbucks $12.47   │ ← Transaction match
│  ✨ +50 Points earned!          │
│                                 │
│  ℹ️ Receipt matched to your bank │
│    transaction from today       │ ← Verification message
│                                 │
│  Note: Full item details coming │
│  in Phase 2 with OCR!           │
│                                 │
│   [Save Receipt] ←──────        │
│                                 │
└─────────────────────────────────┘
```

**Key Features**:
- Receipt preview with basic OCR
- Transaction matching confirmation
- Points reward for successful match
- Clear next steps

### 15C. Receipt Captured (No Match)
```
┌─────────────────────────────────┐
│ ← Retake           Save Anyway →│
│                                 │
│      Receipt Captured          │
│                                 │
│  ┌────────────────────────┐    │
│  │ [Receipt Image Preview] │    │
│  │                        │    │
│  │  TARGET                │    │
│  │  ---------------       │    │
│  │  Items...              │    │
│  │  ---------------       │    │
│  │  Total:        $47.23  │    │
│  │                        │    │
│  └────────────────────────┘    │
│                                 │
│  ⚠️ No matching transaction      │ ← Warning
│                                 │
│  We couldn't find a Target      │
│  transaction for $47.23 in      │
│  your connected accounts.       │
│                                 │
│  Possible reasons:              │
│  • Transaction hasn't posted    │
│  • Used different payment       │
│  • Receipt from another day     │
│                                 │
│  No points awarded unless       │
│  matched to bank transaction    │ ← Clear policy
│                                 │
│   [Try Different Receipt]       │
│   [Save Without Points]         │
│                                 │
└─────────────────────────────────┘
```

**Key Features**:
- Clear mismatch explanation
- Possible reasons for no match
- Options to retry or save anyway
- Transparent point policy

### 16. Budget Screen
```
┌─────────────────────────────────┐
│ ← Back              Edit →      │
│                                 │
│        March Budget             │
│     $1,853 of $3,000           │
│                                 │
│  ┌────────────────────────┐    │
│  │ Needs (50%) • $1,500   │    │
│  │ ████████████░ $1,247   │    │ ← Within budget (green)
│  │ 🏠 Rent, utilities     │    │
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ Wants (30%) • $900     │    │
│  │ ██████░░░░░░ $487     │    │ ← Under budget
│  │ 🎮 Entertainment, dining│    │
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ Savings (20%) • $600   │    │
│  │ ███░░░░░░░░░ $119     │    │ ← Needs attention (yellow)
│  │ 💰 Emergency + goals   │    │
│  └────────────────────────┘    │
│                                 │
│  Alerts:                        │
│  ⚠️ Dining out 40% over target │
│  ✅ Bills on track             │
│  📈 Consider increasing savings │
│                                 │
│  [Adjust Budget] ←──────────    │
├─────────────────────────────────┤
│ [🏠] [💳] [✨] [📊] [👤]      │
└─────────────────────────────────┘
```

**Key Features**:
- 50/30/20 rule implementation
- Visual progress bars with color coding
- Category breakdowns with emojis
- Proactive alerts and recommendations
- Edit capability for adjustments

### 17. Points Store
```
┌─────────────────────────────────┐
│ ← Back      💎 1,250 pts        │
│                                 │
│      Get More Points            │
│                                 │
│  Best Value                     │
│  ┌────────────────────────┐    │
│  │ 💰 5,000 Points        │    │
│  │ Most Popular!          │    │
│  │ $9.99                  │    │ ← Best deal
│  │ [Buy Now]              │    │
│  └────────────────────────┘    │
│                                 │
│  Point Packs                    │
│  ┌────────────────────────┐    │
│  │ 💎 1,000 Points       │    │
│  │ $2.99                  │    │
│  │ [Buy]                  │    │
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ 💎 2,500 Points       │    │
│  │ $5.99 (Save 20%)       │    │
│  │ [Buy]                  │    │
│  └────────────────────────┘    │
│                                 │
│  How Points Work:               │
│  • 500 pts = 1 basic wish       │
│  • 2000 pts = 1 advanced wish   │
│  • Earn 50 pts per receipt      │
│  • Monthly bonus: 1000 pts      │
│                                 │
├─────────────────────────────────┤
│ [🏠] [💳] [✨] [📊] [👤]      │
└─────────────────────────────────┘
```

**Key Features**:
- Clear pricing tiers with savings
- Best value highlighting
- Point usage explanation
- Earning opportunities outlined
- Current balance visible

### 18. Profile & Settings
```
┌─────────────────────────────────┐
│ ← Back                          │
│                                 │
│        Your Profile             │
│                                 │
│  ┌────────────────────────┐    │
│  │ Sarah Martinez         │    │
│  │ Active Subscriber ✓    │    │ ← Subscription status
│  │                        │    │
│  │ 💎 1,850 Points        │    │
│  │ 🔥 3 day streak        │    │
│  └────────────────────────┘    │
│                                 │
│  Account                        │
│  • Connected Banks        [3] > │
│  • Notifications          On >  │
│  • Privacy & Security       >   │
│  • Export Data              >   │
│                                 │
│  Points & Rewards               │
│  • Points History          >    │
│  • Achievements            >    │
│  • Referral Program        >    │
│                                 │
│  Subscription                   │
│  • Manage Subscription     >    │ ← Opens App Store
│  • Restore Purchase        >    │
│  • Cancel Subscription     >    │
│                                 │
│  Support                        │
│  • Help Center             >    │
│  • Contact Support         >    │
│  • Rate App               >     │
│                                 │
│  [Sign Out]                     │
│                                 │
├─────────────────────────────────┤
│ [🏠] [💳] [✨] [📊] [👤]      │
└─────────────────────────────────┘
```

**Key Features**:
- Profile summary with key metrics
- Grouped settings by category
- Direct subscription management
- Privacy controls accessible
- Support options visible

### 2C. Subscription Expired (Paywall)
```
┌─────────────────────────────────┐
│                                 │
│    [Lamp Dimmed Animation]     │ ← Lamp looks sad
│         😔                     │
│                                 │
│    "Your Trial Has Ended"      │
│                                 │
│  To continue using Djinn:      │
│                                 │
│  ✓ 5,000 points monthly    │
│  ✓ Unlimited accounts          │
│  ✓ Unlimited receipt scans     │
│  ✓ All premium features        │
│                                 │
│  Only $7.99/month              │
│                                 │
│  ╔═══════════════════════════╗  │
│  ║ 🍎 Subscribe Now         ║  │ ← Primary CTA
│  ╚═══════════════════════════╝  │
│                                 │
│  ╔═══════════════════════════╗  │
│  ║ [G] Subscribe Now        ║  │
│  ╚═══════════════════════════╝  │
│                                 │
│  [Restore Purchase]            │ ← For reinstalls
│                                 │
│  ─────────────────────────     │
│                                 │
│  Questions? Contact support    │
│                                 │
└─────────────────────────────────┘
```

**Key Features**:
- Emotional lamp state (dimmed/sad)
- Clear value proposition
- Platform-specific subscription buttons
- Restore option for existing subscribers
- Support contact option

## Component Specifications

### Navigation Bar
- **Height**: 56px
- **Background**: White with subtle bottom border
- **Elements**: Account selector, Lamp icon with points, Streak counter
- **Typography**: Inter 14px medium

### Tab Bar
- **Height**: 60px
- **Icons**: 
  - 🏠 Dashboard (home)
  - 💳 Transactions (credit card)
  - ✨ Wishes (sparkles) - CENTER with emerald accent
  - 📊 Budget (chart)
  - 👤 Profile (person)
- **Active State**: Emerald green with label
- **Inactive**: Gray #6B7280

### Cards
- **Background**: White
- **Border**: 1px solid #E5E7EB
- **Border Radius**: 16px
- **Padding**: 16px
- **Shadow**: 0 1px 3px rgba(0,0,0,0.1)

### Buttons
```css
/* Primary Button */
.btn-primary {
  background: linear-gradient(135deg, #10B981, #6EE7B7);
  color: white;
  border-radius: 12px;
  padding: 16px 24px;
  font-weight: 600;
  min-height: 48px;
}

/* Secondary Button */
.btn-secondary {
  background: white;
  color: #10B981;
  border: 1px solid #10B981;
  border-radius: 12px;
  padding: 16px 24px;
  min-height: 48px;
}

/* Floating Action Button */
.fab {
  background: linear-gradient(135deg, #10B981, #6EE7B7);
  border-radius: 50%;
  width: 56px;
  height: 56px;
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
}
```

## Lamp Animation States

### Idle
- Gentle liquid swirling
- Soft emerald glow
- Occasional small sparkle

### Thinking/Processing
- Faster liquid movement
- Brighter pulsing glow
- Stream of sparkles from spout

### Success
- Liquid settles into satisfied pattern
- Golden sparkles burst
- Brief scale animation (1.0 → 1.1 → 1.0)

### Error/Sad
- Liquid slows down
- Brief red tint or dimmed state
- Gentle shake animation

## Color Palette
```css
/* Primary */
--emerald-500: #10B981;
--emerald-400: #34D399;
--emerald-600: #059669;

/* Secondary */
--purple-500: #8B5CF6;
--purple-400: #A78BFA;

/* Neutral */
--gray-900: #111827;
--gray-700: #374151;
--gray-500: #6B7280;
--gray-300: #D1D5DB;
--gray-100: #F3F4F6;

/* Semantic */
--success: #10B981;
--warning: #F59E0B;
--error: #EF4444;
--info: #3B82F6;
```

## Typography
```css
/* Headers */
.h1 {
  font-family: 'Playfair Display', serif;
  font-size: 28px;
  font-weight: 700;
  color: #111827;
}

.h2 {
  font-family: 'Inter', sans-serif;
  font-size: 20px;
  font-weight: 600;
  color: #374151;
}

/* Body */
.body {
  font-family: 'Inter', sans-serif;
  font-size: 16px;
  font-weight: 400;
  color: #6B7280;
}

.body-small {
  font-family: 'Inter', sans-serif;
  font-size: 14px;
  font-weight: 400;
  color: #9CA3AF;
}
```

## Accessibility Requirements
- **Touch Targets**: Minimum 48x48px
- **Color Contrast**: 4.5:1 minimum for body text
- **Focus States**: Visible emerald outline
- **Screen Reader**: All interactive elements labeled
- **Reduced Motion**: Respect system preference

## Implementation Notes

### MVP Priorities (Week 1-6)
1. **Dashboard**: Account aggregation and overview
2. **AI Wishes**: Natural language queries with point system
3. **Transactions**: List with receipt attachment
4. **Budget Tracking**: 50/30/20 rule implementation
5. **Points System**: Store and earning mechanics
6. **Profile**: Settings and subscription management

### Phase 2 Features (Month 3-6)
1. **Advanced Analytics**: Predictive insights and trends
2. **Social Features**: Lamp Circles and challenges
3. **Achievement System**: Badges, levels, and rewards
4. **Enhanced OCR**: Full receipt itemization
5. **Export Features**: PDF reports and tax documents

### Technical Stack
- **Frontend**: React Native
- **State**: Redux + Ferry/HiveStore (offline-first)
- **AI**: DeepSeek R1 via OpenRouter
- **Banking**: Plaid API
- **Analytics**: PostHog
- **Receipt Processing**: Basic OCR (Phase 1), Advanced ML (Phase 2)

## Success Metrics
- **Daily Active Users**: >30%
- **Weekly Wish Usage**: >70% of users
- **Receipt Upload Rate**: >20% of transactions
- **Budget Interaction**: >50% weekly
- **Points Purchase**: >15% of users

## User Flow Summary

### Primary User Journey
Dashboard → AI Wish → View Results → Add Receipt → Earn Points → Check Budget

### Engagement Loop
1. Daily login (streak maintenance)
2. Check spending (dashboard)
3. Ask AI question (wish)
4. Upload receipts (points)
5. Review budget (insights)
6. Share achievements (social)

### Monetization Flow
Free Trial → Point Usage → Point Purchase → Subscription Renewal

---

*This document contains the main app functionality. See `djinn-mvp-onboarding-wireframes.md` for the onboarding experience.*