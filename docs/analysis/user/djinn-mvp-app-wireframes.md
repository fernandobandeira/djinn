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

### 11. Main Dashboard (Multi-Account View)
```
┌─────────────────────────────────┐
│ 🏦 All Accounts ▼  [Lamp] 600pts│ ← Account selector
│                         🔥 3     │ ← Streak counter
├─────────────────────────────────┤
│                                 │
│  Good morning, Sarah! ☀️       │
│                                 │
│  ┌────────────────────────┐    │
│  │ Total Balance          │    │ ← All accounts total
│  │ $12,847.23            │    │
│  │ ↑ $523 from last month │    │ ← Green for positive
│  └────────────────────────┘    │
│                                 │
│  Your Accounts:                │
│  ┌────────────────────────┐    │
│  │ 🏦 Chase Checking      │    │ ← Institution shown
│  │ ****4521              │    │ ← Masked account
│  │ $4,892.15             │    │
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ 💳 Chase Sapphire      │    │ ← Card name
│  │ ****7823              │    │
│  │ -$1,247 / $5,000      │    │ ← Balance/Limit
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ 🏦 Ally Savings        │    │
│  │ ****9012              │    │
│  │ $9,202.97             │    │
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
- Multi-account view with institutions
- Account numbers masked for security
- Credit card utilization shown
- Account selector dropdown
- Quick access to primary actions
- Visual budget progress indicator

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ All family accounts with institutions clearly shown
- ✅ Account numbers masked for privacy
- ✅ Can see credit utilization at a glance
- **Result**: Complete financial picture with security

**Zoe (Digital Native, 25)**
- ✅ Multiple cards tracked separately
- ✅ Institution names help organize
- ✅ Credit limit prevents overspending
- **Result**: Manages multiple cards effectively

**Alex (Financial Freshman, 19)**
- ✅ Sees checking vs credit card clearly
- ✅ Credit utilization helps learn limits
- ✅ Simple institution labels
- **Result**: Learns responsible credit usage

### 11A. Account Selector Dropdown
```
┌─────────────────────────────────┐
│ Select Accounts to View         │
├─────────────────────────────────┤
│                                 │
│  [✓] All Accounts              │ ← Selected
│      Total: $12,847.23         │
│                                 │
│  ─────────────────────────     │
│                                 │
│  [ ] 🏦 Chase                   │ ← Institution group
│      ├ [✓] Checking ****4521   │
│      │     $4,892.15           │
│      └ [✓] Sapphire ****7823   │
│           -$1,247 / $5,000     │
│                                 │
│  [ ] 🏦 Ally Bank               │
│      └ [✓] Savings ****9012    │
│           $9,202.97            │
│                                 │
│  [ ] 💳 American Express        │
│      └ [ ] Platinum ****1234   │
│           -$523 / $10,000      │
│                                 │
│  [✓] 💵 Cash Transactions       │ ← Include cash
│      This month: $234.50       │
│                                 │
│  ─────────────────────────     │
│                                 │
│  [Apply Filter]                │
│                                 │
└─────────────────────────────────┘
```

**Key Features**:
- Multi-select account filtering
- Group by institution
- Shows balances for each account
- Cash transactions toggle
- Quick "All Accounts" option

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
│  │ 🪙 2000 points          │    │ ← Premium wish
│  └────────────────────────┘    │
│                                 │
│  Recent Wishes:                 │
│  • "Show my subscriptions" ✓    │
│  • "Coffee spending trend" ✓    │
│                                 │
│  [Get More Points] ←────────    │ ← Link to store
└─────────────────────────────────┘
```

**Key Features**:
- Natural language AI interface
- Clear point costs per query
- Popular question templates
- Recent query history
- Direct link to purchase points

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Natural language interface feels intuitive
- ✅ Clear point costs shown upfront
- ✅ Query history for expense report insights
- **Result**: Trusts the transparent pricing model

**Zoe (Digital Native, 25)**
- ✅ Popular questions match her needs perfectly
- ✅ Recent wishes show AI learning patterns
- ✅ Direct link to buy more points
- **Result**: Engaged with AI assistant daily

**Alex (Financial Freshman, 19)**
- ✅ Can ask questions in plain English
- ✅ "What am I wasting money on?" resonates
- ⚠️ 2000 points for budget seems expensive
- **Recommendation**: Offer starter questions at 250 points

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

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Shows exact number of transactions analyzed
- ✅ Transparency about which accounts accessed
- ✅ Cancel option provides control
- **Result**: Trusts the data processing transparency

**Zoe (Digital Native, 25)**
- ✅ Engaging lamp animation during wait
- ✅ Sparkles make processing feel magical
- ✅ Progress bar shows it's actually working
- **Result**: Entertained while waiting

**Alex (Financial Freshman, 19)**
- ✅ Cool animations hold attention
- ✅ Cancel button if impatient
- ⚠️ "847 transactions" might feel overwhelming
- **Recommendation**: Use friendlier language like "Looking through your purchases"

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
│  │    Chase Sapphire: $38  │    │ ← Card breakdown
│  │    Cash: $9            │    │
│  │    Make at home 3x/week │    │
│  │    Save: $28           │    │
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ 2. Subscriptions: $23   │    │
│  │    Cancel Hulu (unused) │    │
│  │    On: Amex ****1234    │    │ ← Shows which card
│  │    Last watched: 47 days│    │
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ 3. Dining: $156/month   │    │
│  │    Multiple cards used: │    │
│  │    • Sapphire: $89 (57%)│    │ ← Rewards card
│  │    • Checking: $45 (29%)│    │
│  │    • Cash: $22 (14%)    │    │
│  │    Skip 1 dinner = $43  │    │
│  └────────────────────────┘    │
│                                 │
│  Total Savings: $94/month 🎉   │
│                                 │
│   [Ask Another Wish] ←──────    │
└─────────────────────────────────┘
```

**Key Features**:
- Actionable insights with specific amounts
- Shows which accounts/cards involved
- Breakdown by payment method
- Clear savings opportunities
- Evidence-based recommendations
- Share functionality for social engagement
- Total impact summary

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Specific dollar amounts for each saving
- ✅ Actionable steps ("3 days/week")
- ✅ Evidence-based recommendations
- **Result**: Implements suggestions immediately

**Zoe (Digital Native, 25)**
- ✅ Share button for social bragging
- ✅ Catches forgotten subscriptions
- ✅ Total savings shown prominently
- **Result**: Shares wins on Instagram stories

**Alex (Financial Freshman, 19)**
- ✅ Average meal cost makes sense
- ✅ Simple actions to take
- ✅ Celebration emoji for wins
- **Result**: Finally understands spending patterns

### 15. Transactions Screen (With Receipt Status)
```
┌─────────────────────────────────┐
│ ← Back         Filter: All ▼    │
│                                 │
│        Transactions             │
│                                 │
│  [💵 Add Cash] [📷 Receipt]     │ ← Quick actions
│                                 │
│  Today                          │
│  ┌────────────────────────┐    │
│  │ Starbucks              │    │
│  │ Coffee & Snacks  ☕    │    │
│  │ Chase Sapphire ****7823│    │ ← Card used
│  │ -$12.47         10:32am│    │
│  │ [📷 Add Receipt]       │    │ ← CTA on each transaction
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ Shell Gas Station      │    │
│  │ Transportation  ⛽      │    │
│  │ Chase Checking ****4521│    │ ← Debit used
│  │ -$45.23         8:15am │    │
│  │ [📷 Add Receipt]       │    │
│  └────────────────────────┘    │
│                                 │
│  Yesterday                      │
│  ┌────────────────────────┐    │
│  │ Netflix               │    │
│  │ Subscription  📺  [🔄] │    │ ← Recurring indicator
│  │ Amex ****1234          │    │
│  │ -$15.99               │    │
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ Whole Foods           │    │
│  │ Groceries  🛒         │    │
│  │ Chase Sapphire ****7823│    │
│  │ -$127.84              │    │
│  │ ✅ Receipt matched     │    │ ← Receipt status shown
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ Amazon                │    │
│  │ Shopping  📦          │    │
│  │ Amex ****1234          │    │
│  │ -$89.99               │    │
│  │ 🔄 Receipt pending    │    │ ← Pending match status
│  └────────────────────────┘    │
│                                 │
│  Floating action button:        │
│         [➕]                    │ ← FAB for quick add
│                                 │
├─────────────────────────────────┤
│ [🏠] [💳] [✨] [📊] [👤]      │
└─────────────────────────────────┘
```

**Key Features**:
- Chronological transaction list
- Shows account/card used for each transaction
- Quick buttons for cash and receipt entry
- Category icons for quick scanning
- Receipt attachment CTAs
- Recurring transaction indicators
- Filter capabilities (All, Cards, Cash)

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Can see which card was used for each purchase
- ✅ Cash tracking for complete picture
- ✅ Receipt attachment for tax documentation
- **Result**: Complete expense tracking across all payment methods

**Zoe (Digital Native, 25)**
- ✅ Quick cash entry button for split bills
- ✅ Card details help track rewards
- ✅ Filter by payment type
- **Result**: Never misses tracking any expense

**Alex (Financial Freshman, 19)**
- ✅ Cash button for campus purchases
- ✅ Sees which card was charged
- ✅ Manual entry for unbanked spending
- **Result**: Learns to track everything

### 15A. Adding Receipt Photo (Works Offline)
```
┌─────────────────────────────────┐
│ ← Cancel                        │
│                                 │
│      Add Receipt Photo          │
│                                 │
│  How did you pay?               │ ← Payment method
│  ┌────────────────────────┐    │
│  │ 💳 Chase Sapphire     │    │ ← Default last used
│  │    ****7823     [✓]    │    │
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ 🏦 Chase Checking     │    │
│  │    ****4521     [ ]    │    │
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ 💵 Cash               │    │
│  │    For tax records    │    │ ← Still useful
│  │    No points    [ ]    │    │ ← Clear expectation
│  └────────────────────────┘    │
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
│  [📷 Capture] ←─────────        │ ← Main action
│  [📁 Choose from Gallery]       │ ← Alternative
│                                 │
└─────────────────────────────────┘
```

**Key Features**:
- Payment method selector upfront
- Defaults to last used card
- Cash option for non-trackable purchases
- Camera integration with guide overlay
- Works offline - syncs when connected

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Can select correct card for tax categorization
- ✅ Cash option for untraceable purchases
- ✅ Works without internet connection
- **Result**: Accurate expense tracking for taxes

**Zoe (Digital Native, 25)**
- ✅ Quick card selection speeds up process
- ✅ Remembers last used payment method
- ✅ Camera guide makes it foolproof
- **Result**: Faster, more accurate receipt capture

**Alex (Financial Freshman, 19)**
- ✅ Learns which card was used for what
- ✅ Cash option for campus purchases
- ✅ Simple selection interface
- **Result**: Better understanding of payment methods

### 15B. Receipt Captured (Saved for Processing)
```
┌─────────────────────────────────┐
│ ← Retake              Done →    │
│                                 │
│      Receipt Saved! ✓           │
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
│  🔄 Processing Receipt           │ ← Processing status
│                                 │
│  We're checking for matching    │
│  transactions in your accounts. │
│                                 │
│  • If matched: +50 points       │
│  • Usually takes: 1-5 minutes   │
│  • You'll be notified           │
│                                 │
│  📱 Keep uploading! Even if the │
│     transaction hasn't posted   │
│     yet, we'll match it later. │
│                                 │
│   [Upload Another]   [Done]     │
│                                 │
└─────────────────────────────────┘
```

**Key Features**:
- Receipt saved immediately
- Async processing explained
- Timeline expectations (1-5 minutes)
- Points pending until matched
- Encourages continuous uploading

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Receipt saved immediately for records
- ✅ Clear processing timeline
- ✅ Async matching works in background
- **Result**: Batch uploads all receipts quickly

**Zoe (Digital Native, 25)**
- ✅ Can keep uploading without waiting
- ✅ Notification coming for points
- ✅ Upload flow is fast like stories
- **Result**: Uploads receipts rapid-fire

**Alex (Financial Freshman, 19)**
- ✅ Simple "saved" confirmation
- ✅ Understanding of 1-5 minute wait
- ✅ Can upload before transaction posts
- **Result**: Uploads right after purchase

### 15C. Receipt Status - Matched
```
┌─────────────────────────────────┐
│        Notification             │
├─────────────────────────────────┤
│                                 │
│    ✅ Receipt Matched!          │
│                                 │
│  Your Starbucks receipt was     │
│  matched to a transaction       │
│                                 │
│  ┌────────────────────────┐    │
│  │ Starbucks              │    │
│  │ Today, 10:32am         │    │
│  │ $12.47                 │    │
│  │                        │    │
│  │ ✨ +50 points earned!  │    │
│  └────────────────────────┘    │
│                                 │
│  Total points: 1,300            │
│                                 │
│  [View Details]  [Dismiss]      │
│                                 │
└─────────────────────────────────┘
```

**Key Features**:
- Notification when async match completes
- Points awarded after successful match
- Links to transaction details
- Shows updated point balance

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Clear confirmation of successful match
- ✅ Points automatically added
- ✅ Can verify transaction details
- **Result**: Trusts the matching system

**Zoe (Digital Native, 25)**
- ✅ Push notification creates engagement
- ✅ Points feel like rewards
- ✅ Quick dismiss option
- **Result**: Excited when matches happen

**Alex (Financial Freshman, 19)**
- ✅ Simple success notification
- ✅ Free points earned
- ✅ Clear connection to purchase
- **Result**: Motivated to keep uploading

### 15D. Manual Cash Transaction Entry
```
┌─────────────────────────────────┐
│ ← Cancel                Save →  │
│                                 │
│    Add Cash Transaction         │
│                                 │
│  Have the receipt?              │
│  ┌────────────────────────┐    │
│  │ 📷 Yes, scan receipt   │    │ ← Redirects to camera
│  │    (Easier & faster)   │    │
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ ✏️ No, enter manually  │    │ ← Shows form below
│  │    [Selected]          │    │
│  └────────────────────────┘    │
│                                 │
│  ─────────────────────────     │
│                                 │
│  ┌────────────────────────┐    │
│  │ Amount:                │    │
│  │ $[     ]               │    │ ← Numeric input
│  └────────────────────────┘    │
│                                 │
│  ┌────────────────────────┐    │
│  │ Where:                 │    │
│  │ [Merchant name...   ]  │    │ ← Text input
│  └────────────────────────┘    │
│                                 │
│  Category:                     │
│  ┌────────────────────────┐    │
│  │ 🍔 Food & Dining   [✓] │    │ ← Quick select
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ 🚗 Transportation   [ ] │    │
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ 🛒 Shopping         [ ] │    │
│  └────────────────────────┘    │
│                                 │
│  ┌────────────────────────┡    │
│  │ Notes (optional):      │    │
│  │ [                   ]  │    │
│  └────────────────────────┘    │
│                                 │
│  💡 Cash purchases won't earn   │
│     points but help track      │
│     your complete spending     │
│                                 │
└─────────────────────────────────┘
```

**Key Features**:
- Receipt check upfront to guide user
- Redirects to camera if receipt available
- Manual entry only if no receipt
- Category selection for budgeting
- No points for cash (clearly stated)
- Helps complete spending picture

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Smart routing based on receipt availability
- ✅ Receipt capture for tax documentation
- ✅ Manual fallback when needed
- **Result**: Efficient cash expense tracking

**Zoe (Digital Native, 25)**
- ✅ Guided to fastest input method
- ✅ Receipt photo is quicker than typing
- ✅ Category icons make it visual
- **Result**: Uses receipt capture whenever possible

**Alex (Financial Freshman, 19)**
- ✅ Clear choice between methods
- ✅ Learns receipts are valuable
- ✅ Simple manual entry as backup
- **Result**: Starts keeping receipts

### 15E. Cash Transaction in List
```
┌─────────────────────────────────┐
│ ← Back         Filter: All ▼    │
│                                 │
│        Transactions             │
│                                 │
│  Today                          │
│  ┌────────────────────────┐    │
│  │ Campus Bookstore       │    │
│  │ School Supplies  📚    │    │
│  │ 💵 Cash                │    │ ← Cash indicator
│  │ -$47.23         2:15pm │    │
│  │ Manual entry           │    │ ← Entry type
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ Food Truck            │    │
│  │ Food & Dining  🌮     │    │
│  │ 💵 Cash                │    │
│  │ -$12.00         12:30pm│    │
│  │ ✅ Receipt attached    │    │ ← Has receipt
│  └────────────────────────┘    │
│                                 │
├─────────────────────────────────┤
│ [🏠] [💳] [✨] [📊] [👤]      │
└─────────────────────────────────┘
```

**Key Features**:
- Cash transactions clearly marked
- Shows manual vs receipt entry
- Included in spending totals
- Filtered view options

### 15F. Receipt Status - Expired
```
┌─────────────────────────────────┐
│      Receipt Management         │
├─────────────────────────────────┤
│                                 │
│    ⏰ Receipt Expired           │
│                                 │
│  We couldn't match this receipt │
│  after 30 days:                │
│                                 │
│  ┌────────────────────────┐    │
│  │ CVS Pharmacy           │    │
│  │ Uploaded: Nov 15       │    │
│  │ Amount: $23.47         │    │
│  │                        │    │
│  │ Status: No match found │    │
│  └────────────────────────┘    │
│                                 │
│  Possible reasons:              │
│  • Used cash payment            │
│  • Different card used          │
│  • Transaction didn't sync      │
│                                 │
│  Receipt will be removed from   │
│  pending queue but kept in      │
│  your receipt history.          │
│                                 │
│  [View History]  [Got it]       │
│                                 │
└─────────────────────────────────┘
```

**Key Features**:
- Receipts expire after 30 days unmatched
- Clear explanation of expiration
- Kept in history for records
- Possible reasons for no match

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Receipts kept in history for taxes
- ✅ Clear 30-day timeline
- ✅ Explanation of why no match
- **Result**: Understands system limits

**Zoe (Digital Native, 25)**
- ⚠️ Might be annoyed by no points
- ✅ Clear reasons provided
- ✅ Can still view in history
- **Recommendation**: Offer 5 sympathy points

**Alex (Financial Freshman, 19)**
- ✅ Learns about cash vs card tracking
- ✅ 30 days is generous window
- ✅ Simple explanation
- **Result**: Better understands system

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

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ 50/30/20 rule for family budgeting
- ✅ Proactive alerts for course correction
- ✅ Edit capability for adjustments
- **Result**: Maintains household budget effectively

**Zoe (Digital Native, 25)**
- ✅ Visual progress bars with color coding
- ✅ Instant understanding of status
- ✅ Specific overspending alerts
- **Result**: Makes immediate spending adjustments

**Alex (Financial Freshman, 19)**
- ✅ Simple Needs/Wants/Savings breakdown
- ✅ Percentages easier than dollar amounts
- ⚠️ Savings might feel impossible
- **Recommendation**: Celebrate any savings amount

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

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Clear pricing tiers with savings shown
- ✅ Multiple earning opportunities listed
- ✅ Transparent point usage explanation
- **Result**: Buys bulk pack for family use

**Zoe (Digital Native, 25)**
- ✅ "Most Popular" badge influences choice
- ✅ Best value clearly highlighted
- ✅ Monthly bonus points shown
- **Result**: Purchases 5,000 point pack

**Alex (Financial Freshman, 19)**
- ✅ Small $2.99 pack affordable
- ✅ Can earn points without buying
- ⚠️ Might avoid purchasing entirely
- **Recommendation**: Offer student discount or starter pack

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

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Privacy & Security prominently placed
- ✅ Data export for record keeping
- ✅ Clear subscription management
- **Result**: Feels in control of data and billing

**Zoe (Digital Native, 25)**
- ✅ Achievements section for badges
- ✅ Referral program for social sharing
- ✅ Points history tracking
- **Result**: Explores every gamification feature

**Alex (Financial Freshman, 19)**
- ✅ Simple App Store subscription flow
- ✅ Help center easily accessible
- ✅ No complex billing pages
- **Result**: Comfortable with settings

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

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ $7.99 reasonable for family management
- ✅ Unlimited accounts included
- ✅ Restore purchase option available
- **Result**: Subscribes for household use

**Zoe (Digital Native, 25)**
- ✅ 5,000 monthly points compelling
- ✅ Sad lamp creates emotional response
- ✅ Platform-specific buttons familiar
- **Result**: Subscribes to maintain streak

**Alex (Financial Freshman, 19)**
- ⚠️ $7.99 feels expensive for student
- ✅ Clear value proposition shown
- ⚠️ No student discount visible
- **Critical Recommendation**: Add "Student? Tap here for 50% off" link

## Receipt Processing Flow

### Async Receipt Matching System
All receipts are processed asynchronously - there are NO instant matches:

1. **Upload** → Receipt saved immediately (works offline)
2. **Processing** → System checks for matches (1-5 minutes)
3. **Outcomes**:
   - **Matched with existing transaction** → Points awarded, notification sent
   - **Saved for future transaction** → Waits for transaction to post (1-3 days)
   - **Expired after 30 days** → Moved to history, no points

### Receipt Status Indicators
- 📷 **Add Receipt** - No receipt uploaded yet
- 🔄 **Processing** - Receipt uploaded, checking for match
- ⏳ **Pending** - Receipt waiting for transaction to post
- ✅ **Matched** - Receipt matched, points awarded
- ⏰ **Expired** - No match found after 30 days

### Points Policy
- **50 points** awarded only when receipt successfully matches
- **No points** for:
  - Unmatched receipts
  - Cash transactions
  - Expired receipts
  - Duplicate receipts

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

## Persona-Specific Optimizations

### Sarah (Privacy-First Professional)
- **Key Motivators**: Security, family financial management, tax organization
- **Critical Features**: Multi-account aggregation, receipt storage, data export
- **Conversion Triggers**: Privacy controls, transparent data handling
- **Retention Factors**: Time-saving automation, expense report generation

### Zoe (Digital Native)
- **Key Motivators**: Gamification, social sharing, instant insights
- **Critical Features**: Points/streaks, AI wishes, achievement badges
- **Conversion Triggers**: Social proof, FOMO on features, bulk point deals
- **Retention Factors**: Continuous engagement mechanics, shareable wins

### Alex (Financial Freshman)
- **Key Motivators**: Learning, simplicity, affordability
- **Critical Features**: Natural language AI, visual budgets, educational tips
- **Conversion Triggers**: Student pricing ($3.99), beginner-friendly language
- **Retention Factors**: Progressive learning, celebrating small wins

---

*This document contains the main app functionality. See `djinn-mvp-onboarding-wireframes.md` for the onboarding experience.*