# Djinn MVP Onboarding & Tutorial Wireframes
*Version 2.0 - Persona-Optimized Onboarding Flow*

Generated: 2025-01-21
Status: MVP Onboarding Specification
Platform: Mobile (iOS/Android)

## Overview
This document contains the onboarding, setup, and tutorial screens (1-10A) from the Djinn personal finance app MVP. Each screen includes detailed persona reactions and recommendations based on our three primary user segments.

## Key Changes in Version 2.0
1. **Screen 2**: Added privacy messaging and beginner-friendly badge
2. **Screen 4**: Replaced goals with persona/life stage selection
3. **Screen 4B (NEW)**: Personalized goal suggestions based on persona
4. **Screen 6**: Demo adapts to selected persona
5. **Screen 7**: Added ROI/savings messaging to justify cost
6. **Screen 8**: Enhanced security messaging
7. **Screen 10A**: Tutorial varies by persona type

## Authentication Strategy
- **OAuth-Only**: No email/password option
- **Two-Layer Security**: OAuth for account creation + biometric for daily access
- **Biometric Support**: FaceID, TouchID, device-specific fallbacks

## Brand & Visual Identity
- **Primary Color**: Emerald Green (#10B981)
- **Secondary**: Purple accents (#8B5CF6)
- **Mascot**: Magic lamp with animated liquid (no separate genie)
- **Design Balance**: 70% clean white UI, 20% gradients, 10% liquid glass effects
- **Typography**: Inter (primary), Playfair Display (headers)

## Onboarding Flow Analysis

### 1. Splash Screen
```
┌─────────────────────────────────┐
│                                 │
│                                 │
│     [Animated Magic Lamp]       │ ← Emerald liquid swirling
│         ✨ ✨ ✨               │ ← Sparkles from spout
│                                 │
│          Djinn                  │ ← Playfair Display
│   Your Money Wishes             │ ← Inter, #6B7280
│         Granted                 │
│                                 │
│    [Loading Progress Bar]       │ ← Emerald gradient
│                                 │
└─────────────────────────────────┘
```

**Animation**: 
- Lamp liquid gently swirls during 2-3 second load
- Soft emerald glow pulses
- Sparkles emit from spout

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Professional appearance builds trust
- ✅ Clean, non-intrusive loading
- ⚠️ May want to skip animation if returning user
- **Recommendation**: Add "Skip" after 1 second for returning users

**Zoe (Digital Native, 25)**
- ✅ Loves the magical animation
- ✅ Engaging visual hook
- 💡 Might want more personality
- **Recommendation**: Consider seasonal variations or achievements affecting lamp appearance

**Alex (Financial Freshman, 19)**
- ✅ Not intimidating, feels approachable
- ✅ Game-like quality reduces anxiety
- 💡 Might not understand "genie" metaphor immediately
- **Recommendation**: Ensure value prop is clear in next screen

### 2. Welcome/Sign In Screen
```
┌─────────────────────────────────┐
│                                 │
│    [Large Lamp Animation]       │ ← Hero lamp, liquid moving
│        ✨ ✨ ✨                │
│                                 │
│          Djinn                  │ ← Playfair Display, large
│   Your Money Wishes             │ ← Inter, medium
│         Granted                 │
│                                 │
│  ┌────────────────────────┐    │
│  │ 🔒 Your data stays yours│    │ ← NEW: Privacy badge
│  └────────────────────────┘    │
│                                 │
│  "Track spending, earn rewards" │
│    "and get AI insights"        │
│                                 │
│  ╔═══════════════════════════╗  │
│  ║  🍎  Continue with Apple  ║  │ 
│  ╚═══════════════════════════╝  │
│                                 │
│  ╔═══════════════════════════╗  │
│  ║  [G] Continue with Google ║  │
│  ╚═══════════════════════════╝  │
│                                 │
│  ✨ Perfect for beginners       │ ← NEW: Beginner-friendly
│     No financial expertise needed│
│                                 │
│  By continuing, you agree to    │
│  Terms of Service & Privacy     │
│                                 │
└─────────────────────────────────┘
```

**Changes Made:**
- ✅ Added privacy badge for Sarah
- ✅ Added beginner-friendly messaging for Alex
- ✅ Kept OAuth-only for security

#### Updated Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ "Your data stays yours" badge addresses her #1 concern
- ✅ OAuth-only provides excellent security
- ✅ No password management or leak risk
- ✅ Privacy messaging builds immediate trust
- **Result**: Major barrier removed, likely to continue

**Zoe (Digital Native, 25)**
- ✅ Instant recognition of social login (familiar)
- ✅ "Earn rewards" catches attention
- ✅ Clean, modern design appeals to her aesthetic
- 💡 Still wants social proof/features
- **Recommendation**: Add user count for social validation

**Alex (Financial Freshman, 19)**
- ✅ "Perfect for beginners" directly addresses his anxiety
- ✅ "No financial expertise needed" is reassuring
- ✅ Simple two-button choice reduces overwhelm
- ✅ Familiar login methods from other apps
- **Result**: Anxiety significantly reduced, confident to proceed

### 2B. Returning User - Biometric Login
```
┌─────────────────────────────────┐
│                                 │
│                                 │
│    [Small Lamp Animation]       │ ← Gentle idle animation
│                                 │
│       Welcome back!             │
│      Sarah Martinez             │ ← Name from previous login
│                                 │
│    [FaceID/TouchID Icon]        │ ← Device-specific icon
│         ⟳ Scanning              │ ← Animated state
│                                 │
│   "Use Face ID to continue"     │ ← Device-specific text
│                                 │
│                                 │
│  ╔═══════════════════════════╗  │
│  ║  🍎  Continue with Apple  ║  │ ← Same OAuth buttons
│  ╚═══════════════════════════╝  │    as fallback
│                                 │
│  ╔═══════════════════════════╗  │
│  ║  [G] Continue with Google ║  │ ← If biometric fails
│  ╚═══════════════════════════╝  │
│                                 │
└─────────────────────────────────┘
```

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Biometric adds security layer she values
- ✅ Quick access without password exposure
- ✅ OAuth fallback maintains security
- **Recommendation**: Add "Switch account" option for multiple profiles

**Zoe (Digital Native, 25)**
- ✅ Fast re-entry keeps engagement high
- ✅ FaceID feels modern and premium
- **Recommendation**: Add streak indicator "🔥 Welcome back! Day 3"

**Alex (Financial Freshman, 19)**
- ✅ No passwords to remember
- ✅ Same as other apps they use
- **Recommendation**: Clear indication of what happens if biometric fails

### 3. OAuth Permission Screen (System)
```
┌─────────────────────────────────┐
│     [Google/Apple Dialog]       │
│                                 │
│  "Djinn" wants to use          │
│  "google.com" to sign in        │
│                                 │
│  This allows the app to:        │
│  • See your email address      │
│  • See your basic profile      │
│                                 │
│  [Cancel]     [Continue]        │ ← System buttons
│                                 │
└─────────────────────────────────┘
```

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ System-level dialog builds trust (not Djinn asking directly)
- ✅ Clear permissions listed (email, basic profile only)
- ✅ Familiar OAuth flow from other apps
- **Result**: Comfortable with standard permissions

**Zoe (Digital Native, 25)**
- ✅ Standard OAuth flow, seen it hundreds of times
- ✅ Quick tap "Continue" without hesitation
- ✅ Trusts Google/Apple more than random app
- **Result**: No friction, continues immediately

**Alex (Financial Freshman, 19)**
- ✅ System dialog feels safe and official
- ✅ Only basic permissions requested
- ✅ Same flow as signing into games, social apps
- **Result**: Familiar process, no anxiety

### 4. Who Are You?
```
┌─────────────────────────────────┐
│                                 │
│     Tell Us About You ✨        │
│                                 │
│    [Lamp Curious Animation]     │
│                                 │
│  This helps us personalize      │
│  your experience                │
│                                 │
│  Which best describes you?      │
│                                 │
│  ┌────────────────────────┐    │
│  │ 🏠 Managing a Household │    │
│  │ Budgets, bills, family  │    │
│  └────────────────────────┘    │
│                                 │
│  ┌────────────────────────┐    │
│  │ 🎓 Student              │    │
│  │ Learning to budget      │    │
│  └────────────────────────┘    │
│                                 │
│  ┌────────────────────────┐    │
│  │ 🌱 Just Beginning       │    │
│  │ First job, new to money │    │
│  └────────────────────────┘    │
│                                 │
│  ┌────────────────────────┐    │
│  │ 💼 Growing Wealth       │    │
│  │ Investing, optimizing   │    │
│  └────────────────────────┘    │
│                                 │
│   [Continue] ←──────────        │
│   You can change this anytime   │ ← Reassurance
│                                 │
└─────────────────────────────────┘
```

**Why This Works:**
- Life stages are clearer than abstract goals
- Each option has a brief description
- Maps directly to our personas:
  - 🏠 Household → Sarah (Privacy-First Professional)
  - 🎓 Student → Alex (Financial Freshman)
  - 🌱 Beginning → Zoe (Digital Native starting career)
  - 💼 Growing → Marcus (Tech-Savvy optimizer)

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ "Managing a Household" immediately resonates
- ✅ "Budgets, bills, family" matches her exact needs
- ✅ Clear, professional categorization
- ✅ "You can change this anytime" reduces commitment fear
- **Result**: Confident selection, feels understood

**Zoe (Digital Native, 25)**
- ✅ "Just Beginning" feels appropriate for first job
- ✅ "First job, new to money" is exactly her situation
- ✅ No intimidating financial jargon
- ✅ Options feel achievable, not overwhelming
- **Result**: Sees herself clearly in the options

**Alex (Financial Freshman, 19)**
- ✅ "Student" category removes all confusion
- ✅ "Learning to budget" is exactly what he wants
- ✅ No need to understand complex financial goals
- ✅ Descriptive text explains what each means
- **Result**: Decision paralysis completely eliminated

### 4B. Personalized Goals (NEW SCREEN)
```
┌─────────────────────────────────┐
│                                 │
│   Great! Here's what we        │
│   recommend for [students]:     │ ← Dynamic based on selection
│                                 │
│    [Lamp Helper Animation]      │
│                                 │
│  Your suggested goals:          │
│                                 │
│  FOR STUDENTS:                  │
│  ☑️ Track spending              │ ← Pre-selected
│  "See where your money goes"    │
│  ☑️ Build emergency fund        │ ← Pre-selected
│  "Start with just $500"         │
│  ☐ Manage subscriptions        │
│  "Find services you forgot"     │
│                                 │
│  FOR HOUSEHOLD:                 │
│  ☑️ Optimize family budget      │
│  "Save $200+ monthly"           │
│  ☑️ Track tax deductions        │
│  "Never miss a write-off"       │
│  ☐ Plan major purchases        │
│  "Vacation, renovation, car"    │
│                                 │
│  FOR BEGINNERS:                 │
│  ☑️ Learn smart spending        │
│  "Build good habits early"      │
│  ☑️ Start saving               │
│  "Even $20/month helps"         │
│  ☐ Understand your money       │
│  "Simple insights, no jargon"   │
│                                 │
│   [Continue with selections]    │
│   [Choose different goals]      │
│                                 │
└─────────────────────────────────┘
```

**Personalization Benefits:**
- Pre-selects appropriate goals
- Uses language that resonates with each persona
- Sets realistic expectations

#### Updated Persona Reactions:

**Sarah (Household Manager)**
- ✅ Goals focus on family finances
- ✅ "Tax deductions" highly relevant
- ✅ Dollar amounts are meaningful ($200+ monthly)
- **Recommendation**: Default notifications OFF

**Zoe (Just Beginning)**
- ✅ Goals are achievable and non-intimidating
- ✅ "Build good habits early" resonates
- ✅ Can check multiple goals
- **Recommendation**: Show social features

**Alex (Student)**
- ✅ Realistic goals for student budget
- ✅ "Start with just $500" feels achievable
- ✅ No pressure with "Even $20/month helps"
- **Recommendation**: Emphasize learning aspects

### 5. Biometric Setup
```
┌─────────────────────────────────┐
│                                 │
│     Secure Your Account         │
│                                 │
│    [Lamp Shield Animation]      │ ← Protective glow
│                                 │
│  "Enable quick & secure access" │
│   "with your fingerprint or"    │
│         "Face ID"               │
│                                 │
│    [FaceID/Fingerprint Icon]    │ ← Platform specific
│                                 │
│  Benefits:                      │
│  ✓ Bank-level security         │
│  ✓ Quick daily access          │
│  ✓ No passwords to remember    │
│                                 │
│   [Enable Biometric] ←──────    │ ← Primary CTA
│                                 │
│   [Maybe Later]                 │ ← Skip option
│                                 │
└─────────────────────────────────┘
```

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ "Bank-level security" resonates strongly
- ✅ Clear benefits explanation
- ✅ Can skip if uncomfortable
- **Recommendation**: Add "Biometric data never leaves your device"

**Zoe (Digital Native, 25)**
- ✅ Quick setup like Instagram/TikTok
- ✅ Visual shield animation adds trust
- **Recommendation**: Emphasize convenience over security

**Alex (Financial Freshman, 19)**
- ✅ "No passwords" is appealing
- ⚠️ "Bank-level" might seem serious/scary
- **Recommendations**:
  - Use simpler language: "Keep your money info safe"
  - Larger "Maybe Later" button to reduce pressure

### 6. Try Djinn First (Demo - Persona Adaptive)
```
┌─────────────────────────────────┐
│                        Skip →    │
│                                 │
│     "See What Djinn Can Do"     │
│     [For Students Like You]     │ ← Dynamic subtitle
│                                 │
│    [Lamp Demo Animation]        │ ← Sparkly demo state
│                                 │
│  Let's show you how it works    │
│  with sample [student] data:    │ ← Persona-specific
│                                 │
│  STUDENT VERSION:               │
│  ┌────────────────────────┐    │
│  │ "Why am I always broke  │    │
│  │  by Thursday?"          │    │
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ "How much do I spend    │    │
│  │  on coffee?"            │    │
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ "Can I afford spring    │    │
│  │  break trip?"           │    │
│  └────────────────────────┘    │
│                                 │
│  HOUSEHOLD VERSION:             │
│  ┌────────────────────────┐    │
│  │ "Where can we cut       │    │
│  │  family expenses?"      │    │
│  └────────────────────────┘    │
│  ┌────────────────────────┐    │
│  │ "Track kids' spending"  │    │
│  └────────────────────────┘    │
│                                 │
│  Your real data will be even    │
│  more powerful! 💫              │
│                                 │
│   [Try Demo] ←──────────────    │ ← Primary CTA
│         • ○ ○                   │ ← Step 1 of 3
└─────────────────────────────────┘
```

**Note**: The demo questions shown would dynamically change based on the persona selected in Screen 4:
- **🎓 Students**: "Why broke by Thursday?", "Coffee spending?", "Spring break trip?"
- **🏠 Household**: "Cut family expenses?", "Track kids' spending?", "Optimize groceries?"
- **🌱 Beginning**: "Build first budget?", "Emergency fund?", "Understand spending?"
- **💼 Growing**: "Optimize investments?", "Tax strategies?", "Maximize rewards?"

#### Updated Persona Reactions:

**Sarah (Household Manager)**
- ✅ "For Households Like You" feels personalized
- ✅ Family-focused questions resonate immediately
- ✅ Demo with fake data maintains privacy
- ✅ "Cut family expenses" directly addresses her pain
- **Result**: High engagement, sees clear family value

**Zoe (Digital Native, 25)**
- ✅ Interactive demo is engaging
- ✅ Instant gratification from demo
- 💡 Wants more exciting questions
- **Recommendations**:
  - Add "Find money you're wasting" 
  - Include "Discover hidden fees"
  - Show points earned in demo

**Alex (Financial Freshman, 19)**
- ✅ Demo reduces commitment anxiety
- ✅ Questions are understandable
- ✅ "How can I save" is exactly their need
- **Recommendations**:
  - Make "Try Demo" bigger than "Skip"
  - Add "2 minute demo" time expectation
  - Include beginner-focused question

### 6A. Demo Results (Persona Adaptive)
```
┌─────────────────────────────────┐
│ ← Back                          │
│                                 │
│    [Lamp Success Animation]     │
│                                 │
│  STUDENT DEMO RESULTS:          │
│                                 │
│  You're spending most on:       │
│  • Food & Dining - $187         │ ← Realistic student amounts
│  • Entertainment - $76          │
│  • Transportation - $34         │
│                                 │
│  💡 You could save $30/month    │
│     by meal prepping twice      │
│     per week                    │
│                                 │
│  That's a new video game        │ ← Relatable savings
│  every month! 🎮                │
│                                 │
│  HOUSEHOLD DEMO RESULTS:        │
│                                 │
│  Family spending breakdown:     │
│  • Groceries - $847             │ ← Realistic family amounts
│  • Kids Activities - $234       │
│  • Utilities - $156             │
│                                 │
│  💡 You could save $150/month   │
│     by switching providers and  │
│     meal planning               │
│                                 │
│   [Start Free Trial] ←──────    │
│   [View Pricing]                │
│         ○ • ○                   │ ← Step 2 of 3
└─────────────────────────────────┘
```

#### Updated Persona Reactions:

**Sarah (Household Manager)**
- ✅ Family-focused spending breakdown resonates
- ✅ $847 groceries is realistic for family
- ✅ "Kids Activities" category is spot-on
- ✅ $150/month savings is substantial
- **Result**: Clear ROI demonstration for families

**Zoe (Digital Native)**
- ✅ Would see "Beginning" version with smaller amounts
- ✅ Gamification elements (achievements) appealing
- ✅ Career-starter focused categories
- 💡 Still wants social comparison features
- **Result**: Sees value for her life stage

**Alex (Student)**
- ✅ $187 food budget is realistic for student
- ✅ "Video game every month" is perfect framing
- ✅ $30/month feels achievable, not overwhelming
- ✅ Entertainment category resonates
- **Result**: High confidence this will work for him

### 7. Start Your Free Trial
```
┌─────────────────────────────────┐
│ ← Back                          │
│                                 │
│    [Lamp Premium Animation]     │
│                                 │
│        "Try Djinn FREE"         │
│      "7 Days, Zero Risk"        │
│                                 │
│  ┌────────────────────────┐    │
│  │ 💰 You'll likely save   │    │ ← Honest messaging
│  │    more than $7.99/month│    │
│  │                         │    │
│  │ Making this a smart     │    │
│  │ investment in yourself  │    │
│  └────────────────────────┘    │
│                                 │
│  What's included:               │
│  ✓ 5,000 points monthly        │
│  ✓ Find hidden money leaks      │ ← Value-focused
│  ✓ Cancel forgotten subscriptions│
│  ✓ Optimize your spending       │
│  ✓ Build wealth automatically   │
│                                 │
│  After trial: $7.99/month       │
│                                 │
│  🎯 Most users find hidden       │ ← Value proposition
│     money worth much more!      │
│                                 │
│  ╔═══════════════════════════╗  │
│  ║ 🍎 Start Saving Now      ║  │ ← Action-oriented
│  ╚═══════════════════════════╝  │
│                                 │
│  ╔═══════════════════════════╗  │
│  ║ [G] Start Saving Now     ║  │
│  ╚═══════════════════════════╝  │
│                                 │
│  Cancel anytime. No questions.  │
│         ○ ○ •                   │ ← Step 3 of 3
└─────────────────────────────────┘
```

**Key Changes:**
- Honest value messaging ("You'll likely save more than $7.99")
- Reframes as "investment" not "cost"
- Changes CTA from "Start Trial" to "Start Saving"
- Focus on hidden money discovery vs specific amounts

#### Updated Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Honest messaging builds trust ("You'll likely save more")
- ✅ "Smart investment" reframing resonates
- ✅ No fake specific claims (appreciated)
- ✅ "Cancel anytime. No questions" is reassuring
- ✅ Platform billing maintains security
- **Result**: Honest approach increases conversion likelihood

**Zoe (Digital Native, 25)**
- ✅ "Start Saving Now" CTA is action-oriented
- ⚠️ Still needs better value justification for Gen Z budget
- ✅ "Hidden money" discovery sounds exciting
- ✅ 7-day trial reduces commitment anxiety
- **Recommendations**:
  - Add peer comparison: "Join 50,000+ users saving money"
  - Show social proof of success

**Alex (Financial Freshman, 19)**
- 🔴 $7.99/month still significant for student budget
- ✅ Honest "likely save more" is believable vs fake claims
- ✅ "Smart investment in yourself" reframes positively
- ✅ Simplified features list is less intimidating
- **Critical Need**: Student pricing option ($3.99) to convert this segment
- **Result**: Honest approach helps, but price remains barrier

### 7A. Subscription Confirmation (Platform)
```
┌─────────────────────────────────┐
│    [Apple/Google Dialog]        │ ← Platform UI
│                                 │
│      Subscribe to Djinn         │
│                                 │
│  Free for 7 days                │
│  Then $7.99/month               │
│                                 │
│  [Fingerprint/FaceID Icon]      │
│  Confirm with Touch ID          │
│                                 │
│  ─────────────────────      │
│                                 │
│  [Cancel]     [Subscribe]       │ ← Platform buttons
│                                 │
└─────────────────────────────────┘
```

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ System-level confirmation builds trust
- ✅ Platform billing (no Djinn access to payment)
- ✅ Familiar App Store/Play Store flow
- **Result**: Comfortable with platform transaction

**Zoe (Digital Native, 25)**
- ✅ Quick biometric confirmation
- ✅ Same flow as other app subscriptions
- ✅ No new payment method needed
- **Result**: Frictionless conversion

**Alex (Financial Freshman, 19)**
- ✅ Biometric feels secure but simple
- ⚠️ Still concerned about monthly cost
- ✅ Can cancel through familiar App Store
- **Result**: Proceeds but with cost anxiety

### 8. Connect Your Banks
```
┌─────────────────────────────────┐
│                        Skip →    │
│                                 │
│   "Connect Your Accounts"       │
│                                 │
│    [Lamp with Shield Icon]      │ ← Trust visual
│                                 │
│  ┌────────────────────────┐    │
│  │ 🔒 Bank-Level Security │    │
│  │                        │    │
│  │ ✓ We NEVER store your  │    │ ← NEW: Clear security
│  │   bank password         │    │
│  │                        │    │
│  │ ✓ Read-only access     │    │
│  │   (can't move money)   │    │
│  │                        │    │
│  │ ✓ 256-bit encryption   │    │
│  │   (same as your bank)  │    │
│  │                        │    │
│  │ ✓ Disconnect anytime   │    │
│  └────────────────────────┘    │
│                                 │
│  2 million users trust us       │ ← Social proof
│  Powered by Plaid (used by      │
│  Venmo, Robinhood, Coinbase)    │ ← Familiar names
│                                 │
│  💎 Earn 200 bonus points!      │
│                                 │
│   [Connect Securely] ←──────    │ ← Security-focused CTA
│   [Skip for now]                │
│                                 │
└─────────────────────────────────┘
```

#### Updated Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ "We NEVER store your password" addresses key concern
- ✅ "Can't move money" reassurance is critical
- ✅ 256-bit encryption matches bank security
- ✅ Familiar company names (Venmo, Robinhood) build trust
- **Result**: Security messaging convinces her to connect

**Zoe (Digital Native, 25)**
- ✅ "2 million users" social proof works
- ✅ 200 bonus points immediate incentive
- ✅ Recognizes Plaid from other apps
- ✅ 30-second setup expectation
- **Result**: Connects without hesitation

**Alex (Financial Freshman, 19)**
- ✅ "Parents can't see this" would be reassuring
- ⚠️ Still nervous but security messaging helps
- ✅ Can skip reduces pressure
- ✅ "Start with just one account" would help
- **Result**: Likely connects after initial hesitation

### 8A. Plaid Link Widget (Native SDK)
```
┌─────────────────────────────────┐
│ [X] Close        Plaid          │ ← Plaid's UI
│                                 │
│     Select your bank            │
│                                 │
│  [🔍 Search institutions...]    │
│                                 │
│  ┌────────────────────────┐    │
│  │ • Chase                │    │
│  │ • Bank of America      │    │
│  │ • Wells Fargo          │    │
│  │ • Capital One          │    │
│  │ • US Bank              │    │
│  └────────────────────────┘    │
│                                 │
└─────────────────────────────────┘
```

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Plaid branding reinforces trust
- ✅ Familiar bank names visible
- ✅ Clean, professional interface
- ✅ Can close anytime (X button)
- **Result**: Proceeds with confidence

**Zoe (Digital Native, 25)**
- ✅ Quick bank selection process
- ✅ Search functionality efficient
- ✅ Same as linking to other apps
- **Result**: Selects bank quickly

**Alex (Financial Freshman, 19)**
- ✅ Sees major bank names (reassuring)
- ⚠️ May only see one bank they recognize
- ✅ Search helps if smaller credit union
- **Result**: Finds their bank and proceeds

### 8B. After Bank Connection
```
┌─────────────────────────────────┐
│                                 │
│    [Lamp Working Animation]     │ ← Processing
│                                 │
│    "Analyzing Your Accounts"    │
│                                 │
│  ✅ Connected 3 accounts        │
│  ⏳ Fetching transactions...    │
│  ⏳ Finding subscriptions...    │
│  ⏳ Calculating insights...     │
│                                 │
│  This takes 10-30 seconds       │
│                                 │
│  💎 +200 Points earned!         │
│                                 │
│  We'll notify you when ready    │
│  (usually within 1 minute)      │
│                                 │
│   [Continue to Dashboard] ←──   │
│                                 │
└─────────────────────────────────┘
```

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Transparent processing status
- ✅ Clear timeframe (10-30 seconds)
- ✅ Can see what's being analyzed
- ✅ Points reward for connection
- **Result**: Patient during processing, trusts the system

**Zoe (Digital Native, 25)**
- ✅ Immediate feedback (+200 points)
- ✅ Progress indicators familiar
- ✅ 1-minute notification promise
- **Result**: Satisfied with quick processing

**Alex (Financial Freshman, 19)**
- ✅ Clear explanation of what's happening
- ✅ Not overwhelming with too much data
- ✅ Points earned feels rewarding
- ⚠️ Might worry about "analyzing accounts"
- **Result**: Nervous but reassured by transparency

### 9. Welcome Bonus
```
┌─────────────────────────────────┐
│                                 │
│    [Lamp Success Animation]     │ ← Satisfied swirl
│       ✨ 🎉 ✨                │ ← Celebration
│                                 │
│      "Welcome Bonus!"           │
│    "You earned 600 Points"      │
│                                 │
│  ┌────────────────────────┐    │
│  │  500 pts - Welcome Gift │    │
│  │  100 pts - First Wish   │    │
│  │ ──────────────────  │    │
│  │  600 total Lamp Points  │    │
│  └────────────────────────┘    │
│                                 │
│  "That's enough for another"    │
│        "AI Wish! 🎯"           │
│                                 │
│   [Start Using Djinn] ←─────    │
│                                 │ ← Onboarding complete
└─────────────────────────────────┘
```

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Points system is transparent
- ✅ Clear breakdown of earnings
- ✅ Immediate value demonstration
- **Result**: Confident in product value

**Zoe (Digital Native, 25)**
- ✅ Gamification immediately apparent
- ✅ "Another AI Wish" creates excitement
- ✅ Visual celebration feels rewarding
- **Result**: Highly engaged, ready to explore

**Alex (Financial Freshman, 19)**
- ✅ Clear point explanation with simple math
- ✅ "Welcome Gift" feels friendly
- ✅ Immediate reward builds confidence
- **Result**: Feels supported and successful

### 10. Dashboard - First Visit Tutorial
```
┌─────────────────────────────────┐
│ 🏦 All Accounts    [Lamp] 600pts│
│                         🔥 1     │ ← New user
├─────────────────────────────────┤
│                                 │
│  Welcome, Sarah! 👋            │
│                                 │
│  [Tutorial Overlay]             │ ← Semi-transparent
│  ┌────────────────────────┐    │
│  │   📷 Pro Tip!          │    │
│  │                        │    │
│  │ Add receipt photos to  │    │
│  │ transactions for:      │    │
│  │                        │    │
│  │ • Better AI insights   │    │
│  │ • 50 points per receipt│    │
│  │ • Tax deduction finder │    │
│  │                        │    │
│  │ [Show Me How] [Skip]   │    │
│  └────────────────────────┘    │
│                                 │
│  [Blurred dashboard behind]     │
│                                 │
├─────────────────────────────────┤
│ [🏠] [💳] [✨] [📊] [👤]      │
└─────────────────────────────────┘
```

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Tax deduction finder is valuable
- ✅ Can skip tutorial
- ✅ Points balance visible
- **Recommendations**:
  - Add "Receipts stored securely"
  - Mention data retention policy
  - Allow tutorial replay later

**Zoe (Digital Native, 25)**
- ✅ Gamification visible (points, streak)
- ✅ Clear point earning opportunity
- 💡 Wants more social features visible
- **Recommendations**:
  - Add "Friends earning 2x more"
  - Show leaderboard teaser
  - Include achievement preview

**Alex (Financial Freshman, 19)**
- ✅ Friendly welcome message
- ✅ Simple first tip
- ⚠️ "Tax deduction" not relevant
- **Recommendations**:
  - Adjust tip for students
  - Focus on "Track what you buy"
  - Simpler point explanation

### 10A. Receipt Tutorial
```
┌─────────────────────────────────┐
│ ← Back                          │
│                                 │
│    "Make AI Wishes Smarter"     │
│                                 │
│    [Animated Demo]              │
│  ┌────────────────────────┐    │
│  │ Starbucks               │    │
│  │ -$12.47                │    │ ← Sample transaction
│  │ [📷 Add Receipt]       │    │ ← Button highlights
│  └────────────────────────┘    │
│                                 │
│  When you add receipts:         │
│                                 │
│  ✓ AI knows WHAT you bought    │
│    not just WHERE              │
│                                 │
│  ✓ Finds patterns like "always │
│    buying expensive drinks"     │
│                                 │
│  ✓ Suggests specific savings   │
│    "Switch to regular coffee"   │
│                                 │
│  💎 Earn 50 points per receipt! │
│                                 │
│  ℹ️ Points awarded when receipt  │
│    matches your bank transaction│
│                                 │
│   [Got It!] ←───────────        │
└─────────────────────────────────┘
```

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Clear value in detailed tracking
- ✅ Specific savings suggestions
- ✅ Understands matching requirement
- **Recommendations**:
  - Add OCR accuracy disclaimer
  - Mention receipt data privacy

**Zoe (Digital Native, 25)**
- ✅ Gamification clear (50 points)
- ✅ Specific patterns are insightful
- ✅ Starbucks example is relatable
- **Recommendations**:
  - Add "Compete with friends"
  - Show "Unlock achievements"
  - Include social sharing option

**Alex (Financial Freshman, 19)**
- ✅ Simple, clear example
- ✅ Starbucks is relatable
- ✅ "Switch to regular" is actionable
- **Recommendations**:
  - Use student-relevant examples
  - Simplify to "Save money on coffee"
  - Show weekly/monthly savings

## Onboarding Flow Summary

### Complete Persona-Adaptive Flow
1. **Splash** → 2. **Welcome/Sign In** (with privacy badge) → 3. **OAuth Permission** → 4. **Who Are You?** (persona selection) → 4B. **Personalized Goals** → 5. **Biometric Setup** → 6. **Persona-Adaptive Demo** → 6A. **Persona-Adaptive Results** → 7. **Free Trial** (honest messaging) → 7A. **Subscription Confirmation** → 8. **Connect Banks** (enhanced security) → 8A. **Plaid Widget** → 8B. **Bank Processing** → 9. **Welcome Bonus** → 10. **Dashboard Tutorial**

### Key Persona-Adaptive Features
- **Screen 4**: Life stage selection instead of confusing financial goals
- **Screen 4B**: Personalized goal recommendations based on selected persona
- **Screen 6**: Demo questions adapt to persona (student vs household vs beginning)
- **Screen 6A**: Demo results show realistic amounts per persona type
- **All screens**: Enhanced security messaging and honest value propositions

### Skip Points
- Biometric Setup (Maybe Later)
- Demo (Skip)
- Connect Banks (Skip for now) 
- Dashboard Tutorial (Skip)

### Implemented Persona Optimizations

#### ✅ For Sarah (Household Manager - Privacy-First):
1. ✅ Privacy messaging throughout ("Your data stays yours")
2. ✅ Enhanced security features ("We NEVER store passwords")
3. ✅ Family-focused demo questions and results
4. ✅ Honest value propositions build trust
5. ✅ Platform billing maintains security

#### ✅ For Zoe (Just Beginning - Digital Native):
1. ✅ Gamification visible (points, streaks, achievements)
2. ✅ Career-starter focused content
3. ✅ Modern, engaging interface design
4. ✅ Social proof messaging ("2 million users")
5. ✅ Action-oriented CTAs ("Start Saving Now")

#### ✅ For Alex (Student - Financial Freshman):
1. ✅ Simplified language throughout
2. ✅ Student-focused demo content (realistic amounts)
3. ✅ Beginner-friendly messaging ("Perfect for beginners")
4. 🔴 Student pricing option still needed ($3.99 vs $7.99)
5. ✅ Reduced financial jargon and complexity

## Key Achievements & Remaining Recommendations

### ✅ Successfully Implemented
1. **✅ Persona-First Approach**: Life stage selection drives personalized experience
2. **✅ Progressive Disclosure**: Screen 4B shows relevant goals per persona
3. **✅ Trust Signals**: Privacy badges, security messaging, social proof
4. **✅ Honest Value Propositions**: Removed fake claims, honest messaging
5. **✅ Enhanced Security**: Clear encryption, read-only access messaging
6. **✅ Adaptive Content**: Demo questions and results change per persona

### 🔴 Critical Remaining Need
1. **Student Pricing**: $3.99/month option for Alex segment (major conversion blocker)
2. **Social Features**: Peer comparisons for Zoe segment
3. **Privacy Controls**: Granular notification settings for Sarah segment

### Onboarding Metrics to Track
- Screen completion rates
- Skip button usage by persona
- Time per screen
- Drop-off points
- Trial conversion by path taken

---

*This document focuses on the onboarding experience. See `djinn-mvp-app-wireframes.md` for main app functionality.*