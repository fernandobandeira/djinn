# Djinn Manual CSV Upload Flow - Wireframes & Journey Maps
*Version 1.0 - Non-Plaid Bank Support*

Generated: 2025-01-21
Status: Manual Upload Feature Specification
Platform: Mobile (iOS/Android)

## Overview
This document details the complete user experience for customers whose banks are not supported by Plaid. It includes manual account creation, CSV file upload, transaction matching, and ongoing management flows.

## Brand Consistency
Maintains existing Djinn design system:
- **Primary Color**: Emerald Green (#10B981)
- **Secondary**: Purple accents (#8B5CF6)
- **Mascot**: Magic lamp with animated liquid
- **Typography**: Inter (primary), Playfair Display (headers)

## User Journey Map - Manual Banking Flow

### Journey Overview
```
Discovery → Account Setup → CSV Upload → Transaction Matching → Ongoing Management
    ↓            ↓             ↓               ↓                    ↓
Bank not    Manual entry   File import   Smart matching      Regular updates
supported   of accounts    & parsing     & validation        & reconciliation
```

### Emotional Journey
```
😟 Frustrated → 😐 Reassured → 🤔 Cautious → 😊 Relieved → 😎 Confident
"My bank      "I can still   "Is this      "It matched   "This works
isn't listed"  use Djinn"     secure?"      correctly!"   great!"
```

## Entry Points & Detection

### 8C. Bank Not Found (During Plaid Connection)
```
┌─────────────────────────────────┐
│ ← Back          Plaid           │
│                                 │
│     Bank Not Found 😔           │
│                                 │
│  We couldn't find your bank in │
│  our automated system.         │
│                                 │
│  ┌────────────────────────┐    │
│  │ 🏦 [Search result]      │    │
│  │ "Community Credit Union"│    │
│  │ Not available           │    │
│  └────────────────────────┘    │
│                                 │
│  Don't worry! You can still    │
│  use Djinn with:              │
│                                 │
│  ╔═══════════════════════════╗  │
│  ║ 📄 Upload Bank Statements║  │ ← Primary CTA
│  ╚═══════════════════════════╝  │
│                                 │
│  • Add accounts manually       │
│  • Upload CSV/PDF statements   │
│  • Same insights & features    │
│  • Earn 150 bonus points! 💎   │ ← Incentive
│                                 │
│  [Try Another Bank]            │
│  [Skip for Now]                │
│                                 │
└─────────────────────────────────┘
```

**Key Features**:
- Empathetic messaging ("Don't worry!")
- Clear alternative path
- Bonus points to offset inconvenience
- Option to search for another bank

### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Manual upload might feel MORE secure
- ✅ Control over what data is shared
- ✅ 150 bonus points compensates for effort
- **Result**: Likely prefers manual option

**Zoe (Digital Native, 25)**
- 😤 Frustrated by lack of automation
- ✅ Points incentive helps
- ⚠️ Might abandon if too complex
- **Recommendation**: Emphasize "2-minute setup"

**Alex (Financial Freshman, 19)**
- 😟 Worried about complexity
- ✅ Clear instructions help
- ✅ "Same features" is reassuring
- **Result**: Proceeds cautiously

## Manual Account Setup Flow

### 11. Manual Account Dashboard
```
┌─────────────────────────────────┐
│ ← Back                          │
│                                 │
│    Set Up Your Accounts         │
│     Manually 📝                 │
│                                 │
│  [Lamp Helper Animation]        │
│                                 │
│  Let's add your accounts.       │
│  You can upload statements      │
│  after setup.                   │
│                                 │
│  Your Accounts (0)              │
│  ┌────────────────────────┐    │
│  │ No accounts yet         │    │
│  │                        │    │
│  │ [+ Add First Account]  │    │ ← Large, friendly CTA
│  │                        │    │
│  └────────────────────────┘    │
│                                 │
│  What you'll need:              │
│  • Bank/institution name        │
│  • Account nickname             │
│  • Account type                 │
│  • Last 4 digits               │
│  • Current balance (optional)   │
│                                 │
│  💡 Tip: Start with your main   │
│     checking account            │
│                                 │
│  [Watch 30-sec Tutorial] 📹     │
│                                 │
└─────────────────────────────────┘
```

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Manual setup feels more secure than auto-connect
- ✅ Clear list of required information
- ✅ No sensitive data like passwords needed
- **Result**: Prefers this over Plaid connection

**Zoe (Digital Native, 25)**
- ⚠️ Manual setup feels like work
- ✅ Video tutorial option helps
- ✅ Clear, simple requirements list
- **Result**: Proceeds but wishes it was automatic

**Alex (Financial Freshman, 19)**
- ✅ Step-by-step approach not overwhelming
- ✅ "Start with checking" tip is helpful
- ✅ Video tutorial reduces anxiety
- **Result**: Feels capable of completing setup

### 11A. Account List (After Adding)
```
┌─────────────────────────────────┐
│ ← Back           Add More →     │
│                                 │
│    Your Manual Accounts         │
│         3 Connected             │
│                                 │
│  Bank Accounts (2)              │
│  ┌────────────────────────┐    │
│  │ 🏦 Main Checking        │    │
│  │ Community Credit Union  │    │ ← Institution shown
│  │ •••• 1234              │    │
│  │ Balance: $2,847.50      │    │
│  │ [Upload CSV] [Edit]     │    │
│  ├────────────────────────┤    │
│  │ 🏦 Family Savings       │    │
│  │ Community Credit Union  │    │ ← Same institution
│  │ •••• 5678              │    │
│  │ Balance: $10,234.00     │    │
│  │ [Upload CSV] [Edit]     │    │
│  └────────────────────────┘    │
│                                 │
│  Credit Cards (1)               │
│  ┌────────────────────────┐    │
│  │ 💳 Sapphire Reserve     │    │
│  │ Chase Bank             │    │ ← Card issuer shown
│  │ •••• 9012              │    │
│  │ Balance: $1,234.00      │    │
│  │ Limit: $10,000          │    │
│  │ [Upload CSV] [Edit]     │    │
│  └────────────────────────┘    │
│                                 │
│  [+ Add Bank Account]           │
│  [+ Add Credit Card]            │
│                                 │
│  [Continue to Upload CSVs] →    │
│                                 │
└─────────────────────────────────┘
```

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ All accounts organized with institutions clearly shown
- ✅ Can see complete financial picture at a glance
- ✅ Edit options give her control over data
- **Result**: Feels organized and in control

**Zoe (Digital Native, 25)**
- ✅ Quick upload buttons on each account
- ✅ Multiple cards tracked separately
- ✅ Visual organization by type
- **Result**: Ready to start uploading CSVs

**Alex (Financial Freshman, 19)**
- ✅ Clear separation of checking vs credit
- ✅ Masked numbers feel secure
- ✅ Simple add more buttons
- **Result**: Understands account structure

### 12. Add Account Manually
```
┌─────────────────────────────────┐
│ ← Cancel              Save →     │
│                                 │
│      Add Bank Account           │
│                                 │
│  Bank/Credit Union *             │
│  ┌────────────────────────┐    │
│  │ 🔍 Start typing...      │    │ ← Auto-complete search
│  └────────────────────────┘    │
│                                 │
│  Suggestions:                   │
│  • Community Credit Union       │ ← From our DB
│  • Community First CU           │
│  • Community Bank               │
│  [Can't find it? Enter manually]│
│                                 │
│  Account Nickname *              │
│  ┌────────────────────────┐    │
│  │ Main Checking          │    │ ← User's label
│  └────────────────────────┘    │
│                                 │
│  Account Type *                 │
│  ┌────────────────────────┐    │
│  │ ▼ Checking             │    │ ← Dropdown
│  └────────────────────────┘    │
│  • Checking                     │
│  • Savings                      │
│  • Money Market                 │
│  • Investment                   │
│  • Loan                         │
│                                 │
│  Account Number (last 4)         │
│  ┌────────────────────────┐    │
│  │ •••• 1234              │    │ ← Optional but helpful
│  └────────────────────────┘    │
│  💡 Helps match your transactions│
│                                 │
│  Current Balance (optional)     │
│  ┌────────────────────────┐    │
│  │ $ 2,847.50             │    │
│  └────────────────────────┘    │
│                                 │
│  🔒 No passwords or full account │
│     numbers needed              │
│                                 │
│  [Save Account] ←──────────     │
│                                 │
└─────────────────────────────────┘
```

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Type-ahead search prevents typos
- ✅ No routing number needed upfront
- ✅ Only last 4 digits for security
- **Result**: Trusts the minimal data requirement

**Zoe (Digital Native, 25)**
- ✅ Auto-complete speeds up entry
- ✅ Recognizes bank in suggestions
- ⚠️ Still wishes it was automatic
- **Result**: Completes quickly with auto-complete

**Alex (Financial Freshman, 19)**
- ✅ Suggestions help spell bank name
- ✅ Optional fields reduce pressure
- ✅ Security message reassuring
- **Result**: Successfully adds first account

### 12A. Institution Verification (During CSV Upload)
```
┌─────────────────────────────────┐
│                                 │
│   Verifying Institution 🔍      │
│                                 │
│  [Lamp Detective Animation]     │
│                                 │
│  Your CSV contains routing      │
│  number: ****6789              │
│                                 │
│  This belongs to:               │
│  ┌────────────────────────┐    │
│  │ Community Federal CU    │    │ ← From routing lookup
│  └────────────────────────┘    │
│                                 │
│  But your account is set as:    │
│  ┌────────────────────────┐    │
│  │ Community Credit Union  │    │ ← What user entered
│  └────────────────────────┘    │
│                                 │
│  These appear to be the same.   │
│  Should we update the name?     │
│                                 │
│  [Yes, Update] [Keep Mine]      │
│                                 │
│  💡 This helps match future     │
│     transactions better         │
│                                 │
└─────────────────────────────────┘
```

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ System verifies data accuracy
- ✅ Can keep her original entry
- ✅ Routing verification adds trust
- **Result**: Appreciates data validation

**Zoe (Digital Native, 25)**
- ✅ Auto-detection feels smart
- ✅ Quick decision to update
- ✅ One-click correction
- **Result**: Updates to correct name

**Alex (Financial Freshman, 19)**
- ✅ Learns official bank name
- ✅ Simple yes/no choice
- ✅ Explanation helps understanding
- **Result**: Updates for accuracy

### 13. Credit Card Addition
```
┌─────────────────────────────────┐
│ ← Cancel              Save →     │
│                                 │
│      Add Credit Card            │
│                                 │
│  Issuing Bank *                 │
│  ┌────────────────────────┐    │
│  │ Chase Bank            │    │ ← Full institution name
│  └────────────────────────┘    │
│  💡 Bank that issued the card   │
│                                 │
│  Card Nickname *                │
│  ┌────────────────────────┐    │
│  │ Sapphire Reserve      │    │ ← User's label
│  └────────────────────────┘    │
│                                 │
│  Card Network *                 │
│  ┌────────────────────────┐    │
│  │ ▼ Visa                │    │ ← Network type
│  └────────────────────────┘    │
│  • Visa                         │
│  • Mastercard                   │
│  • American Express             │
│  • Discover                     │
│                                 │
│  Card Number (last 4) *         │
│  ┌────────────────────────┐    │
│  │ •••• 5678              │    │ ← For identification
│  └────────────────────────┘    │
│                                 │
│  Current Balance                │
│  ┌────────────────────────┐    │
│  │ $ 1,234.00             │    │
│  └────────────────────────┘    │
│                                 │
│  Credit Limit                   │
│  ┌────────────────────────┐    │
│  │ $ 10,000               │    │
│  └────────────────────────┘    │
│                                 │
│  Statement Close Day            │
│  ┌────────────────────────┐    │
│  │ ▼ 15th of month       │    │ ← Dropdown 1-31
│  └────────────────────────┘    │
│                                 │
│  Payment Due Day                │
│  ┌────────────────────────┐    │
│  │ ▼ 10th of month       │    │ ← For reminders
│  └────────────────────────┘    │
│                                 │
│  [Save Card] ←──────────────    │
│                                 │
└─────────────────────────────────┘
```

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Full control over card details
- ✅ Statement dates help tax tracking
- ✅ No full card number required
- **Result**: Adds all family cards

**Zoe (Digital Native, 25)**
- ✅ Network selection for rewards tracking
- ✅ Credit limit helps budget
- ✅ Payment reminders useful
- **Result**: Tracks multiple cards effectively

**Alex (Financial Freshman, 19)**
- ✅ Learns about statement cycles
- ✅ Credit limit visibility helpful
- ✅ Simple dropdown selections
- **Result**: Better understands credit cards

## CSV Upload Flow

### 14. Account Detail View (Entry to CSV Upload)
```
┌─────────────────────────────────┐
│ ← Back              Edit →      │
│                                 │
│    Main Checking                │
│    Community Credit Union       │
│    •••• 1234                   │
│                                 │
│  Current Balance: $2,847.50     │
│  Last Updated: 3 days ago       │
│                                 │
│  Recent Transactions:           │
│  ┌────────────────────────┐    │
│  │ Jan 18 | Starbucks     │    │
│  │ -$12.47                │    │
│  ├────────────────────────┤    │
│  │ Jan 17 | Shell Gas     │    │
│  │ -$45.23                │    │
│  └────────────────────────┘    │
│  [View All Transactions]        │
│                                 │
│  ╔═══════════════════════════╗  │
│  ║ 📄 Upload New Statement  ║  │ ← Primary action
│  ╚═══════════════════════════╝  │
│                                 │
│  ⚠️ It's been 3 days since your │
│     last update. Upload a new   │
│     CSV to see latest activity  │
│                                 │
│  [Download CSV Help] 📹          │
│                                 │
└─────────────────────────────────┘
```

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Account-specific view prevents errors
- ✅ Recent transactions shown for verification
- ✅ Update reminder is helpful not pushy
- **Result**: Uploads CSV with confidence

**Zoe (Digital Native, 25)**
- ✅ Primary action button is prominent
- ✅ 3-day reminder creates urgency
- ✅ Video help available if needed
- **Result**: Immediately uploads new statement

**Alex (Financial Freshman, 19)**
- ✅ Clear which account being updated
- ✅ Balance shown for reference
- ✅ Help video reduces anxiety
- **Result**: Learns to maintain accounts

### 14A. Select Account for Import
```
┌─────────────────────────────────┐
│ ← Back                          │
│                                 │
│   Select Account for Import     │
│                                 │
│  Which account is this CSV for? │
│                                 │
│  Bank Accounts:                 │
│  ┌────────────────────────┐    │
│  │ 🏦 Main Checking        │    │
│  │ Community Credit Union  │    │
│  │ •••• 1234     [Select] │    │ ← Choose this
│  ├────────────────────────┤    │
│  │ 🏦 Family Savings       │    │
│  │ Community Credit Union  │    │
│  │ •••• 5678     [Select] │    │
│  └────────────────────────┘    │
│                                 │
│  Credit Cards:                  │
│  ┌────────────────────────┐    │
│  │ 💳 Sapphire Reserve     │    │
│  │ Chase Bank             │    │
│  │ •••• 9012     [Select] │    │
│  └────────────────────────┘    │
│                                 │
│  💡 Make sure to download the   │
│     CSV from the same account   │
│     you're selecting here       │
│                                 │
│  [+ Add New Account First]      │
│                                 │
└─────────────────────────────────┘
```

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Must select specific account first
- ✅ All accounts visible for selection
- ✅ Institution names prevent confusion
- **Result**: Selects correct account confidently

**Zoe (Digital Native, 25)**
- ⚠️ Extra step feels unnecessary
- ✅ Quick selection buttons
- ✅ Can add new account if missing
- **Result**: Quickly selects and continues

**Alex (Financial Freshman, 19)**
- ✅ Clear which account for which CSV
- ✅ Visual separation helps choose
- ✅ Reminder to match CSV to account
- **Result**: Avoids mixing up accounts

### 15. CSV Upload Instructions (Account-Specific)
```
┌─────────────────────────────────┐
│ ← Back                          │
│                                 │
│   Upload Statement for:         │
│   Main Checking (****1234)      │ ← Selected account
│   Community Credit Union        │
│                                 │
│  ┌────────────────────────┐    │
│  │                        │    │
│  │    📄                  │    │ ← Drop zone
│  │                        │    │
│  │  Drag & drop CSV here  │    │
│  │         or             │    │
│  │  [Browse Files]        │    │
│  │                        │    │
│  └────────────────────────┘    │
│                                 │
│  Supported formats:             │
│  • CSV (.csv)                   │
│  • Excel (.xls, .xlsx)          │
│  • QIF (.qif)                   │
│  • OFX (.ofx, .qfx)            │
│                                 │
│  Date Range (optional):         │
│  ┌──────────┐ to ┌──────────┐  │
│  │ Jan 1    │    │ Today    │  │
│  └──────────┘    └──────────┘  │
│                                 │
│  ⚠️ Make sure this CSV is from  │
│     Main Checking at Community  │
│     Credit Union                │
│                                 │
│  [Watch How-To Video] 📹        │
│                                 │
└─────────────────────────────────┘
```

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Account clearly identified at top
- ✅ Multiple format support
- ✅ Security message about encryption
- **Result**: Uploads with trust in security

**Zoe (Digital Native, 25)**
- ✅ Drag and drop is familiar
- ✅ Date range saves time
- ⚠️ Manual process still feels slow
- **Result**: Drags file and continues

**Alex (Financial Freshman, 19)**
- ✅ Warning prevents wrong file upload
- ✅ Video help available
- ✅ Browse option if drag fails
- **Result**: Successfully uploads first CSV

### 16. CSV Preview & Mapping
```
┌─────────────────────────────────┐
│ ← Back            Continue →    │
│                                 │
│   Review Your CSV Import        │
│                                 │
│  ✅ File: checking_jan2025.csv  │
│  ✅ 147 transactions found      │
│                                 │
│  Help us understand your format:│
│                                 │
│  Date Column:                   │
│  ┌────────────────────────┐    │
│  │ ▼ Transaction Date    │    │ ← Auto-detected
│  └────────────────────────┘    │
│  Preview: "01/15/2025"          │
│                                 │
│  Description Column:            │
│  ┌────────────────────────┐    │
│  │ ▼ Description         │    │
│  └────────────────────────┘    │
│  Preview: "STARBUCKS #1234"     │
│                                 │
│  Amount Column:                 │
│  ┌────────────────────────┐    │
│  │ ▼ Amount              │    │
│  └────────────────────────┘    │
│  Preview: "-12.47"              │
│                                 │
│  ☐ Amounts are positive for    │
│    withdrawals (some banks)     │ ← Toggle if needed
│                                 │
│  Sample Transactions:           │
│  ┌────────────────────────┐    │
│  │ Jan 15 | Starbucks     │    │
│  │ -$12.47                │    │
│  ├────────────────────────┤    │
│  │ Jan 14 | Shell Gas     │    │
│  │ -$45.23                │    │
│  ├────────────────────────┤    │
│  │ Jan 13 | Paycheck      │    │
│  │ +$2,500.00             │    │
│  └────────────────────────┘    │
│                                 │
│  [Looks Good!] ←────────────    │
│  [Fix Mapping]                  │
│                                 │
└─────────────────────────────────┘
```

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Can verify data before import
- ✅ Sample transactions for validation
- ✅ Fix mapping if needed
- **Result**: Carefully reviews before confirming

**Zoe (Digital Native, 25)**
- ✅ Auto-detection usually works
- ✅ Quick preview of transactions
- ✅ "Looks Good" for speed
- **Result**: Confirms quickly if correct

**Alex (Financial Freshman, 19)**
- ✅ Visual preview helps understand
- ✅ Auto-detection reduces confusion
- ✅ Can fix if something looks wrong
- **Result**: Learns how CSVs work

### 17. Processing & Categorization
```
┌─────────────────────────────────┐
│                                 │
│   [Lamp Working Animation]      │ ← Fast spinning
│       ✨ ✨ ✨ ✨              │
│                                 │
│   "Analyzing Transactions"      │
│                                 │
│  ┌────────────────────────┐    │
│  │ ✓ Parsing CSV...       │    │ ← Completed
│  │ ✓ Detecting patterns...│    │
│  │ ⟳ Categorizing...      │    │ ← In progress
│  │   87 of 147             │    │
│  │ ░ Finding duplicates... │    │ ← Upcoming
│  │ ░ Calculating insights..│    │
│  └────────────────────────┘    │
│                                 │
│  Smart Detection:               │
│  • 12 subscriptions found       │
│  • 3 potential duplicates       │
│  • 23 uncategorized             │
│                                 │
│  This usually takes 30-60       │
│  seconds for first import       │
│                                 │
│  [Cancel]                       │
│                                 │
└─────────────────────────────────┘
```

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Transparent about what's happening
- ✅ Shows smart detection features
- ✅ Can cancel if concerned
- **Result**: Appreciates transparency

**Zoe (Digital Native, 25)**
- ✅ Progress indicators familiar
- ✅ Subscription detection valuable
- ⚠️ 30-60 seconds feels long
- **Result**: Waits but slightly impatient

**Alex (Financial Freshman, 19)**
- ✅ Clear progress indication
- ✅ Interesting to see subscriptions found
- ✅ Time estimate helps patience
- **Result**: Watches progress with interest

## Transaction Matching & Review

### 18. Import Summary & Review
```
┌─────────────────────────────────┐
│ ← Back              Finish →    │
│                                 │
│   Import Complete! 🎉           │
│                                 │
│  [Lamp Success Animation]       │
│                                 │
│  Summary:                       │
│  ✅ 144 transactions imported   │
│  ⚠️ 3 possible duplicates       │
│  ❓ 23 need categories          │
│                                 │
│  Possible Duplicates:           │
│  ┌────────────────────────┐    │
│  │ Jan 10 | Netflix       │    │
│  │ -$15.99                │    │
│  │ [Keep] [Skip]          │    │ ← Quick actions
│  ├────────────────────────┤    │
│  │ Jan 10 | Netflix       │    │
│  │ -$15.99                │    │
│  │ [Keep] [Skip]          │    │
│  └────────────────────────┘    │
│                                 │
│  Review Uncategorized:          │
│  ┌────────────────────────┐    │
│  │ AMZN*2K4Y8 | -$47.23   │    │
│  │ [▼ Select Category]    │    │
│  ├────────────────────────┤    │
│  │ PAYPAL*SELLER | -$23.00│    │
│  │ [▼ Select Category]    │    │
│  └────────────────────────┘    │
│  [Review All 23 →]             │
│                                 │
│  💎 +150 Points earned!         │
│                                 │
│  [Finish Import] ←──────────    │
│  [Review & Edit]                │
│                                 │
└─────────────────────────────────┘
```

### 19. Quick Categorization
```
┌─────────────────────────────────┐
│ ← Back           Save All →     │
│                                 │
│   Categorize Transactions       │
│   23 remaining                  │
│                                 │
│  [Progress Bar ████░░░░░]       │
│                                 │
│  AMZN*2K4Y8                     │
│  January 12 • -$47.23           │
│                                 │
│  ┌────────────────────────┐    │
│  │ 🛍️ Shopping            │    │ ← Quick select
│  │ 🍔 Food & Dining       │    │    categories
│  │ 🚗 Transportation      │    │
│  │ 🏠 Bills & Utilities   │    │
│  │ 💊 Health & Medical    │    │
│  │ 🎮 Entertainment       │    │
│  │ [More Categories ▼]    │    │
│  └────────────────────────┘    │
│                                 │
│  ☐ Apply to all Amazon          │ ← Smart suggestion
│    transactions (12 found)      │
│                                 │
│  Merchant Name (optional):      │
│  ┌────────────────────────┐    │
│  │ Amazon                 │    │ ← Auto-filled
│  └────────────────────────┘    │
│                                 │
│  [Skip] [Previous] [Next →]     │
│                                 │
│  💡 AI will learn and improve   │
│     categorization over time    │
│                                 │
└─────────────────────────────────┘
```

## Ongoing Management

### 20. Manual Update Reminder
```
┌─────────────────────────────────┐
│ 🏦 Manual Accounts [Lamp] 1,850 │
│                         🔥 7     │
├─────────────────────────────────┤
│                                 │
│  ⚠️ Account Update Needed       │ ← Yellow banner
│                                 │
│  Your manual accounts haven't   │
│  been updated in 14 days        │
│                                 │
│  ┌────────────────────────┐    │
│  │ Main Checking          │    │
│  │ Last updated: Jan 7    │    │
│  │ [Upload New CSV]       │    │
│  ├────────────────────────┤    │
│  │ Chase Sapphire         │    │
│  │ Last updated: Jan 5    │    │
│  │ [Upload Statement]     │    │
│  └────────────────────────┘    │
│                                 │
│  ╔═══════════════════════════╗  │
│  ║ 📄 Update All Accounts   ║  │
│  ╚═══════════════════════════╝  │
│                                 │
│  💎 Earn 50 points per update   │
│                                 │
│  [Remind Me Later]              │
│  [Set Update Schedule →]        │
│                                 │
├─────────────────────────────────┤
│ [🏠] [💳] [✨] [📊] [👤]      │
└─────────────────────────────────┘
```

### 21. Update Schedule Settings
```
┌─────────────────────────────────┐
│ ← Back              Save →       │
│                                 │
│   CSV Upload Schedule           │
│                                 │
│  [Lamp Settings Animation]      │
│                                 │
│  How often do you want          │
│  update reminders?              │
│                                 │
│  ○ Weekly (Mondays)             │
│  ● Every 2 weeks                │ ← Default
│  ○ Monthly (1st of month)       │
│  ○ Monthly (15th of month)      │
│  ○ Never (I'll update manually) │
│                                 │
│  Notification Settings:          │
│  ☑️ Push notification           │
│  ☑️ Email reminder              │
│  ☐ In-app banner only          │
│                                 │
│  Best Practice:                 │
│  Upload statements every 2      │
│  weeks for accurate insights    │
│  and to maximize points! 💎     │
│                                 │
│  Current Streak: 3 updates 🔥   │
│  Keep it going!                 │
│                                 │
│  [Save Schedule] ←──────────    │
│                                 │
└─────────────────────────────────┘
```

### 22. Bulk CSV Upload
```
┌─────────────────────────────────┐
│ ← Back                          │
│                                 │
│   Update Multiple Accounts      │
│                                 │
│  [Lamp Multi Animation]         │
│                                 │
│  Select files for each account: │
│                                 │
│  ┌────────────────────────┐    │
│  │ Main Checking          │    │
│  │ ┌──────────────────┐   │    │
│  │ │ 📄 checking.csv  │   │    │ ← File attached
│  │ │ 147 transactions │   │    │
│  │ └──────────────────┘   │    │
│  └────────────────────────┘    │
│                                 │
│  ┌────────────────────────┐    │
│  │ Chase Sapphire         │    │
│  │ [+ Add CSV]            │    │ ← Empty state
│  └────────────────────────┘    │
│                                 │
│  ┌────────────────────────┐    │
│  │ Savings Account        │    │
│  │ [+ Add CSV]            │    │
│  └────────────────────────┘    │
│                                 │
│  Or drag multiple files here:   │
│  ┌────────────────────────┐    │
│  │     Drop Zone          │    │
│  │  We'll match them to   │    │
│  │  accounts automatically │    │
│  └────────────────────────┘    │
│                                 │
│  [Process All Files] ←──────    │
│                                 │
│  💎 Bulk upload = 200 bonus pts │
│                                 │
└─────────────────────────────────┘
```

#### Persona Reactions:

**Sarah (Privacy-First, 34)**
- ✅ Can update all accounts at once
- ✅ Auto-matching to accounts
- ✅ 200 bonus points valuable
- **Result**: Uses bulk upload monthly

**Zoe (Digital Native, 25)**
- ✅ Bulk upload saves time
- ✅ Drag multiple files easy
- ✅ Big point bonus motivating
- **Result**: Always bulk uploads

**Alex (Financial Freshman, 19)**
- ✅ Visual account matching
- ✅ Can add files one by one
- ✅ Extra points for efficiency
- **Result**: Learns to batch updates

## Error States & Recovery

### 23. CSV Format Error
```
┌─────────────────────────────────┐
│ ← Back                          │
│                                 │
│   ⚠️ CSV Format Issue           │
│                                 │
│  [Lamp Confused Animation]      │
│                                 │
│  We're having trouble reading   │
│  your file. Common issues:      │
│                                 │
│  ❌ Date format unclear         │
│  Your file: "1/15/25"           │
│  Expected: "01/15/2025"         │
│                                 │
│  Quick Fixes:                   │
│  ┌────────────────────────┐    │
│  │ Date Format:           │    │
│  │ [▼ MM/DD/YYYY]        │    │ ← Dropdown
│  └────────────────────────┘    │
│                                 │
│  ┌────────────────────────┐    │
│  │ ☐ First row is headers │    │ ← Common issue
│  └────────────────────────┘    │
│                                 │
│  ┌────────────────────────┐    │
│  │ Currency Symbol:       │    │
│  │ [▼ $ (USD)]           │    │
│  └────────────────────────┘    │
│                                 │
│  [Try Again] ←──────────────    │
│  [Upload Different File]        │
│  [Get Help]                     │
│                                 │
└─────────────────────────────────┘
```

### 24. Reconciliation Helper
```
┌─────────────────────────────────┐
│ ← Back              Done →      │
│                                 │
│   Balance Reconciliation        │
│                                 │
│  Your CSV shows a different     │
│  balance than expected:         │
│                                 │
│  Account: Main Checking         │
│  ┌────────────────────────┐    │
│  │ You entered: $2,847.50 │    │
│  │ CSV shows:  $2,647.50  │    │
│  │ Difference: -$200.00   │    │ ← Red text
│  └────────────────────────┘    │
│                                 │
│  Possible reasons:              │
│  • Pending transactions         │
│  • Different statement dates    │
│  • Bank fees not included       │
│                                 │
│  What would you like to do?     │
│                                 │
│  ○ Use CSV balance ($2,647.50)  │
│  ● Keep my balance ($2,847.50)  │
│  ○ Enter new balance manually   │
│                                 │
│  ☐ Don't ask again for small    │
│    differences (<$50)           │
│                                 │
│  [Continue] ←───────────────    │
│                                 │
└─────────────────────────────────┘
```

## Success States & Rewards

### 25. First Manual Import Success
```
┌─────────────────────────────────┐
│                                 │
│   [Lamp Celebration Animation]  │
│      ✨ 🎊 ✨ 🎊 ✨           │
│                                 │
│    Manual Import Master! 🏆     │
│                                 │
│  You've successfully imported   │
│  your first CSV file!           │
│                                 │
│  Rewards Earned:                │
│  ┌────────────────────────┐    │
│  │ 💎 150 pts - CSV import│    │
│  │ 🏆 50 pts - First timer│    │
│  │ 🎯 100 pts - Completion│    │
│  │ ────────────────────   │    │
│  │ 300 points total!      │    │
│  └────────────────────────┘    │
│                                 │
│  Your data is now ready for     │
│  AI insights and analysis!      │
│                                 │
│  Next Steps:                    │
│  • Ask your first AI wish       │
│  • Set up other accounts        │
│  • Schedule regular updates     │
│                                 │
│  [Explore Djinn] ←──────────    │
│                                 │
└─────────────────────────────────┘
```

## Feature Integration

### 26. AI Wishes with Manual Data
```
┌─────────────────────────────────┐
│ ← Back      Points: 1,550 🪙    │
│                                 │
│      AI Insights Available      │
│                                 │
│  Based on your manual imports:  │
│                                 │
│  Popular Questions:             │
│  ┌────────────────────────┐    │
│  │ "Analyze my January     │    │
│  │  spending patterns"     │    │
│  │ 💎 500 points           │    │
│  └────────────────────────┘    │
│                                 │
│  ┌────────────────────────┐    │
│  │ "Find subscriptions in  │    │
│  │  my CSV data"           │    │
│  │ 💎 500 points           │    │
│  └────────────────────────┘    │
│                                 │
│  ⚠️ Note: Manual accounts may   │
│  have less real-time data       │
│  than connected accounts        │
│                                 │
│  💡 Upload CSVs regularly for   │
│     best insights (every 2 wks) │
│                                 │
│  [Type your question...]        │
│                                 │
└─────────────────────────────────┘
```

## Comparison Table: Plaid vs Manual

### 27. Account Connection Options
```
┌─────────────────────────────────┐
│ ← Back                          │
│                                 │
│   Choose Connection Method      │
│                                 │
│  ┌────────────────────────┐    │
│  │ 🔗 Automatic (Plaid)   │    │
│  │ ─────────────────────  │    │
│  │ ✓ Real-time updates    │    │
│  │ ✓ No manual work       │    │
│  │ ✓ 3,000+ banks         │    │
│  │ ✓ Instant setup        │    │
│  │                        │    │
│  │ Best for: Most users   │    │
│  └────────────────────────┘    │
│                                 │
│  ┌────────────────────────┐    │
│  │ 📄 Manual Upload       │    │
│  │ ─────────────────────  │    │
│  │ ✓ Any bank works       │    │
│  │ ✓ You control data     │    │
│  │ ✓ No passwords shared  │    │
│  │ ⚠️ Manual updates       │    │
│  │                        │    │
│  │ Best for: Privacy-first│    │
│  │ or unsupported banks   │    │
│  └────────────────────────┘    │
│                                 │
│  Not sure? Try automatic first. │
│  You can always add manual      │
│  accounts later!                │
│                                 │
│  [Try Automatic] [Go Manual]    │
│                                 │
└─────────────────────────────────┘
```

## Technical Specifications

### Bank Institution Database
- **Pre-populated Database**: 10,000+ US financial institutions
- **Auto-complete Search**: Type-ahead with fuzzy matching
- **Routing Number Mapping**: Backend stores routing → institution mapping
- **CSV Pattern Recognition**: Detect bank from CSV headers/format
- **Learning System**: Improve detection over time from user confirmations

### Smart Institution & Routing Strategy
1. **Manual Account Creation First**: User creates account with institution name
2. **Type-ahead Institution Search**: Help users select from known banks
3. **CSV Upload to Existing Account**: User assigns CSV to created account
4. **Routing Verification**: If CSV contains routing, verify it matches
5. **Institution Correction**: Offer to update institution name if mismatch
6. **Learning System**: Remember corrections for better matching

### CSV Processing Requirements
- **Supported Formats**: CSV, XLS, XLSX, QIF, OFX, QFX
- **Max File Size**: 10MB per file
- **Date Formats**: Auto-detect common formats, manual override available
- **Encoding**: UTF-8, UTF-16, ASCII auto-detection
- **Column Mapping**: Smart detection with manual override
- **Duplicate Detection**: Transaction matching within 7-day window

### Security & Privacy
- **Local Processing First**: Initial parsing happens on device
- **Encryption**: All uploads use TLS 1.3
- **Data Retention**: CSVs deleted after processing (transactions kept)
- **No Bank Credentials**: Never ask for passwords or routing numbers
- **User Control**: Can delete all data anytime

### Points & Gamification
- **First Manual Account**: 100 points
- **First CSV Upload**: 150 points
- **Regular Updates**: 50 points per CSV
- **Bulk Upload**: 200 points bonus
- **Categorization**: 1 point per transaction categorized
- **Update Streak**: 10 points per week maintained

## Persona-Specific Optimizations

### Sarah (Privacy-First Professional)
- ✅ Complete control over data sharing
- ✅ No bank passwords required
- ✅ Can review before importing
- ✅ Local processing emphasis
- **Messaging**: "Your data, your control"

### Zoe (Digital Native)
- ⚠️ Manual process less appealing
- ✅ Bulk upload for efficiency
- ✅ Points incentives throughout
- ✅ Progress bars and animations
- **Messaging**: "Quick 2-minute setup"

### Alex (Financial Freshman)
- ✅ Step-by-step guidance
- ✅ Video tutorials available
- ✅ Simple language throughout
- ✅ Error recovery assistance
- **Messaging**: "We'll help you every step"

## Success Metrics

### Conversion Metrics
- **Plaid Failure → Manual Success**: Target 70% conversion
- **First CSV Upload Success**: Target 85% completion
- **Regular Update Adoption**: Target 60% bi-weekly updates
- **Error Recovery Rate**: Target 90% successful retry

### Engagement Metrics
- **Time to First Upload**: Target <5 minutes
- **Categorization Accuracy**: Target 80% auto-correct
- **Update Streak Maintenance**: Target 40% maintain 4+ weeks
- **Points Earned per User**: Target 500+ monthly

## Implementation Priorities

### Phase 1 (MVP)
1. Manual account creation
2. Basic CSV upload & parsing
3. Simple transaction import
4. Basic categorization
5. Manual balance entry

### Phase 2 (Enhancement)
1. PDF statement OCR
2. Bulk file upload
3. Smart duplicate detection
4. Advanced categorization AI
5. Automated reconciliation

### Phase 3 (Optimization)
1. Bank-specific templates
2. Historical data import
3. Advanced matching algorithms
4. Predictive categorization
5. Cross-account reconciliation

## Key Differentiators

### vs. Competitors
- **Mint**: Doesn't support manual well
- **YNAB**: Too complex for beginners
- **PocketGuard**: Limited manual features
- **Djinn Advantage**: Seamless manual/auto hybrid with gamification

### Unique Value Props
1. **No Bank Left Behind**: Any bank works
2. **Privacy First Option**: Complete control
3. **Gamified Updates**: Points for maintenance
4. **Smart Learning**: AI improves over time
5. **Hybrid Approach**: Mix manual and automatic

---

*End of Manual CSV Upload Flow Documentation*