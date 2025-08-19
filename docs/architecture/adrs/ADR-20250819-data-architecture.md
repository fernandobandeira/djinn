# ADR-20250819: Data Architecture and Schema Design

## Status
Accepted

## Context
The Djinn personal finance application requires a robust data architecture to handle:
- Multi-user financial data with strict isolation
- Import and categorization of existing bank transactions (read-only financial data)
- High precision monetary calculations without floating-point errors
- Receipt itemization for granular spending insights (killer feature)
- Transaction-receipt correlation for detailed expense tracking
- Offline-first mobile synchronization
- Basic privacy compliance (data export, deletion)
- Scale to 100K+ users with millions of transactions

### Constraints
- PostgreSQL 16+ as the primary database (decided in tech stack ADR)
- GraphQL API requiring efficient query patterns
- Mobile clients need offline support with eventual consistency
- OAuth-only authentication (Google/Apple Sign-In)
- Import-only transactions (no financial transaction creation)
- MVP timeline requires pragmatic choices over perfection

## Decision

### Core Data Model

#### 1. Entity Hierarchy
```
Users (root tenant)
├── Institutions (banks, credit unions, brokerages)
│   └── Accounts (checking, savings, credit cards, investments)
├── Categories (hierarchical, user-customizable)
├── Budgets (monthly/weekly/custom periods)
├── Rules (auto-categorization)
├── Receipts (images + OCR data)
└── Merchants (shared but user-enrichable)

Institutions (shared + user-customizable)
├── System institutions (pre-populated: Chase, BofA, etc.)
├── User custom institutions (manual entries)
└── Plaid institutions (auto-imported)

Transactions (imported financial records)
├── Bank/Card transactions (from Plaid/manual import)
├── Receipt correlation (bi-directional matching)
└── Receipt Items (granular purchase details)
```

#### Credit Card Complexity Model

Credit cards operate on **two distinct levels** that we must track:

##### 1. Transaction Level (Individual Purchases)
- **Each purchase** creates a transaction (increases balance owed)
- **Receipts match** to individual credit card transactions (NOT to statements)
- **Categories apply** to each transaction for budgeting
- **Same as debit cards** from a tracking perspective

##### 2. Statement/Liability Level (Monthly Bills)
- **Monthly statement** aggregates all transactions in billing cycle
- **Minimum payment** calculated by issuer
- **Due date** for payment (critical for notifications)
- **Interest charges** if carrying balance
- **APR variations** (purchase, cash advance, balance transfer)

##### 3. Payment Flow (Critical Distinction)

```
Individual CC Purchases          vs.        CC Payment from Checking
├── $50 Restaurant                          ├── $500 payment
├── $100 Gas                                ├── Applied to STATEMENT balance
├── $75 Groceries                           ├── NOT linked to specific purchases
└── Each has receipt potential              └── Could be minimum, full, or partial
```

**Key Insights:**
- **CC Payments reduce statement balance**, not specific transactions
- **Minimum payment** calculated by issuer (e.g., 2% of balance + interest)
- **Payment allocation** follows issuer rules (highest APR first, etc.)
- **You can't track** which specific purchase a payment "covers"

**Payment Transaction Modeling:**
```sql
Checking Account Transaction:
- amount: -$500 (debit from checking)
- type: 'payment'
- transfer_account_id: [credit_card_account_id]
- description: "Payment to Chase Visa"

Credit Card Account Transaction:
- amount: +$500 (credit to reduce balance)
- type: 'payment'
- transfer_account_id: [checking_account_id]
- description: "Payment received"
- is_statement_payment: true  -- Key flag!
```

##### 4. Data Sources & Manual Entry

**Automated Import:**
- **Plaid Transactions API**: Individual purchases and payments
- **Plaid Liabilities API**: Statement data, APR, minimum payments
- **CSV/OFX Import**: Bank statement downloads

**Manual Entry Support:**
- **Manual Credit Card Account**: User creates card with institution, credit limit
- **Manual Transactions**: User enters individual purchases (amount, merchant, date)
- **Manual Statement Info**: User enters minimum payment, due date each month
- **Manual Payment Recording**: User records when they pay the card
- **Receipt Upload**: Works same as automated transactions

**Use Cases for Manual Entry:**
- Store credit cards (Target RedCard, Macy's, etc.)
- Unsupported regional banks
- International credit cards
- Users preferring privacy (no account linking)
- Historical data before Plaid connection

#### Receipt-Transaction Matching Strategy

The system supports **bi-directional matching** since receipts and transactions arrive independently:

1. **Receipt First**: User uploads receipt → OCR processing → Wait for matching transaction
2. **Transaction First**: Bank import → Look for existing unmatched receipts
3. **Matching Algorithm**:
   - **Amount matching** (40% weight): Exact match or within 1%
   - **Date matching** (30% weight): Same day preferred, ±3 days acceptable
   - **Merchant matching** (30% weight): Fuzzy text matching using pg_trgm
4. **Confidence Levels & User Interaction**:
   - **≥80%**: Auto-match and link silently
   - **50-79%**: Suggest match, require user confirmation
     - Show side-by-side comparison
     - User can accept, reject, or select different transaction
   - **<50%**: Keep unmatched, user manually selects from transaction list
     - Show nearby transactions by date
     - Allow skipping (receipt remains unmatched)
5. **Match History**: All potential matches stored for audit and ML training

**Critical Relationships**:
- **Receipts** → Individual CC transactions (for itemization and categorization)
- **CC Payments** → Statement balance (NOT individual transactions)
- **Statement** → Aggregates all transactions in billing period
- **Minimum Payment** → Calculated by issuer, not tied to specific purchases

#### 2. PostgreSQL Schema Design

##### Core Tables

```sql
-- Users table (root of all multi-tenancy)
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),  -- UUIDv7 when available
    firebase_uid TEXT UNIQUE NOT NULL,              -- Firebase Auth UID
    email TEXT NOT NULL,                            -- From OAuth provider
    auth_provider TEXT NOT NULL,                    -- 'google' or 'apple'
    display_name TEXT,
    avatar_url TEXT,
    subscription_tier TEXT DEFAULT 'free',           -- free, premium, premium_plus
    subscription_expires_at TIMESTAMPTZ,
    settings JSONB DEFAULT '{}',                    -- User preferences
    metadata JSONB DEFAULT '{}',                    -- Analytics, features flags
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,                         -- Soft delete for privacy
    
    CONSTRAINT valid_auth_provider CHECK (auth_provider IN ('google', 'apple')),
    CONSTRAINT valid_subscription CHECK (subscription_tier IN ('free', 'premium', 'premium_plus'))
);

CREATE INDEX idx_users_firebase_uid ON users(firebase_uid) WHERE deleted_at IS NULL;
CREATE INDEX idx_users_email ON users(email) WHERE deleted_at IS NULL;

-- Financial Institutions (shared across users)
CREATE TABLE institutions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    plaid_institution_id TEXT UNIQUE,               -- Plaid's institution ID
    name TEXT NOT NULL,
    display_name TEXT NOT NULL,
    logo_url TEXT,
    website_url TEXT,
    routing_numbers TEXT[],                         -- Array of routing numbers
    primary_color TEXT,                             -- Brand color for UI
    country_code CHAR(2) DEFAULT 'US',
    is_system BOOLEAN DEFAULT FALSE,                -- Pre-populated vs user-created
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_institutions_plaid ON institutions(plaid_institution_id) WHERE plaid_institution_id IS NOT NULL;
CREATE INDEX idx_institutions_name ON institutions(lower(name));

-- User institution overrides (similar to merchant overrides)
CREATE TABLE user_institution_settings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    institution_id UUID NOT NULL REFERENCES institutions(id) ON DELETE CASCADE,
    custom_name TEXT,                               -- User's nickname for institution
    custom_color TEXT,                              -- User's preferred color
    notes TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    CONSTRAINT unique_user_institution UNIQUE(user_id, institution_id)
);

-- Financial Accounts (now linked to institutions)
CREATE TABLE accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    institution_id UUID NOT NULL REFERENCES institutions(id) ON DELETE RESTRICT,
    plaid_account_id TEXT,                          -- External ID from Plaid
    account_type TEXT NOT NULL,                     -- checking, savings, credit, investment
    account_name TEXT NOT NULL,                     -- User's name for account
    account_number_mask TEXT,                       -- Last 4 digits only
    currency CHAR(3) NOT NULL DEFAULT 'USD',
    current_balance_minor BIGINT NOT NULL DEFAULT 0, -- In cents/minor units
    available_balance_minor BIGINT,
    credit_limit_minor BIGINT,                      -- For credit cards
    is_active BOOLEAN DEFAULT TRUE,
    is_manual BOOLEAN DEFAULT FALSE,                -- Manual vs connected account
    sync_status TEXT DEFAULT 'pending',             -- pending, syncing, ready, error
    last_synced_at TIMESTAMPTZ,
    metadata JSONB DEFAULT '{}',                    -- Plaid metadata, features, etc.
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    CONSTRAINT valid_account_type CHECK (account_type IN ('checking', 'savings', 'credit', 'investment', 'loan', 'other')),
    CONSTRAINT valid_currency CHECK (currency ~ '^[A-Z]{3}$'),
    CONSTRAINT valid_sync_status CHECK (sync_status IN ('pending', 'syncing', 'ready', 'error'))
);

CREATE INDEX idx_accounts_user_id ON accounts(user_id) WHERE is_active = TRUE;
CREATE INDEX idx_accounts_institution ON accounts(institution_id);
CREATE INDEX idx_accounts_plaid_id ON accounts(plaid_account_id) WHERE plaid_account_id IS NOT NULL;
CREATE INDEX idx_accounts_type ON accounts(account_type, user_id);

-- Credit Card Liabilities (from Plaid Liabilities API)
CREATE TABLE credit_card_liabilities (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    -- APR Information
    apr_percentage_purchase DECIMAL(5,2),           -- Purchase APR
    apr_percentage_cash_advance DECIMAL(5,2),       -- Cash advance APR
    apr_percentage_balance_transfer DECIMAL(5,2),   -- Balance transfer APR
    balance_subject_to_apr_minor BIGINT,            -- Amount subject to interest
    interest_charge_amount_minor BIGINT,            -- Last interest charge
    -- Statement Information
    last_statement_date DATE,
    last_statement_balance_minor BIGINT,
    statement_close_date DATE,                      -- Next statement close
    -- Payment Information
    minimum_payment_amount_minor BIGINT,
    next_payment_due_date DATE,
    last_payment_date DATE,
    last_payment_amount_minor BIGINT,
    is_overdue BOOLEAN DEFAULT FALSE,
    days_overdue INTEGER DEFAULT 0,
    -- Metadata
    plaid_updated_at TIMESTAMPTZ,                   -- When Plaid last updated
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    CONSTRAINT unique_account_liability UNIQUE(account_id)
);

CREATE INDEX idx_liabilities_account ON credit_card_liabilities(account_id);
CREATE INDEX idx_liabilities_due_date ON credit_card_liabilities(next_payment_due_date) 
    WHERE next_payment_due_date IS NOT NULL;

-- Categories (hierarchical)
CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,  -- NULL for system categories
    parent_id UUID REFERENCES categories(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    icon TEXT,
    color TEXT,
    is_system BOOLEAN DEFAULT FALSE,                -- Built-in categories
    is_income BOOLEAN DEFAULT FALSE,
    sort_order INTEGER DEFAULT 0,
    rules JSONB DEFAULT '[]',                       -- Auto-categorization rules
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    CONSTRAINT unique_user_category UNIQUE(user_id, name, parent_id),
    CONSTRAINT no_deep_nesting CHECK (
        -- Prevent more than 3 levels of nesting
        NOT EXISTS (
            SELECT 1 FROM categories c1 
            JOIN categories c2 ON c1.parent_id = c2.id 
            WHERE c2.parent_id = categories.id
        )
    )
);

CREATE INDEX idx_categories_user_id ON categories(user_id);
CREATE INDEX idx_categories_parent_id ON categories(parent_id);
CREATE INDEX idx_categories_system ON categories(is_system) WHERE is_system = TRUE;

-- Transactions (imported from banks/cards)
CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    external_id TEXT,                               -- Plaid transaction ID (for deduplication)
    transaction_type TEXT NOT NULL,                 -- purchase, payment, transfer, income, fee, interest
    amount_minor BIGINT NOT NULL,                   -- In cents/minor units (negative for debits)
    currency CHAR(3) NOT NULL DEFAULT 'USD',
    description TEXT NOT NULL,
    merchant_id UUID REFERENCES merchants(id),
    merchant_name_raw TEXT,                         -- Original merchant name from bank
    category_id UUID REFERENCES categories(id),
    -- Transfer/Payment fields
    transfer_account_id UUID REFERENCES accounts(id), -- For transfers between accounts
    transfer_transaction_id UUID REFERENCES transactions(id), -- Linked transaction on other account
    is_statement_payment BOOLEAN DEFAULT FALSE,     -- True for CC statement payments (not linked to purchases)
    -- Timing fields
    occurred_at TIMESTAMPTZ NOT NULL,
    pending BOOLEAN DEFAULT FALSE,                  -- Pending vs posted
    -- Receipt matching
    has_receipt BOOLEAN DEFAULT FALSE,              -- Quick lookup for receipt existence
    -- Import tracking
    import_source TEXT NOT NULL,                    -- 'plaid', 'manual', 'csv'
    notes TEXT,
    metadata JSONB DEFAULT '{}',                    -- Additional data from import
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    CONSTRAINT valid_transaction_type CHECK (transaction_type IN ('purchase', 'payment', 'transfer', 'income', 'fee', 'adjustment', 'interest')),
    CONSTRAINT valid_currency CHECK (currency ~ '^[A-Z]{3}$'),
    CONSTRAINT valid_import_source CHECK (import_source IN ('plaid', 'manual', 'csv', 'ofx'))
);

CREATE INDEX idx_transactions_user_date ON transactions(user_id, occurred_at DESC);
CREATE INDEX idx_transactions_account ON transactions(account_id, occurred_at DESC);
CREATE INDEX idx_transactions_category ON transactions(category_id) WHERE category_id IS NOT NULL;
CREATE INDEX idx_transactions_merchant ON transactions(merchant_id) WHERE merchant_id IS NOT NULL;
CREATE INDEX idx_transactions_no_receipt ON transactions(user_id, occurred_at) WHERE has_receipt = FALSE;
CREATE INDEX idx_transactions_matching ON transactions(user_id, amount_minor, occurred_at::date) WHERE has_receipt = FALSE;

-- Prevent duplicate imports from Plaid
CREATE UNIQUE INDEX idx_prevent_duplicate_imports 
    ON transactions(account_id, external_id) 
    WHERE external_id IS NOT NULL;

-- Merchants (shared across users but enrichable)
CREATE TABLE merchants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    canonical_name TEXT NOT NULL UNIQUE,           -- Normalized name
    display_name TEXT NOT NULL,
    logo_url TEXT,
    website_url TEXT,
    default_category_id UUID REFERENCES categories(id),
    merchant_type TEXT,
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_merchants_canonical ON merchants(canonical_name);

-- User merchant overrides
CREATE TABLE user_merchant_overrides (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    merchant_id UUID NOT NULL REFERENCES merchants(id) ON DELETE CASCADE,
    custom_name TEXT,
    custom_category_id UUID REFERENCES categories(id),
    always_ignore BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    CONSTRAINT unique_user_merchant UNIQUE(user_id, merchant_id)
);

-- Receipts and OCR data (can exist before or after transactions)
CREATE TABLE receipts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    transaction_id UUID REFERENCES transactions(id) ON DELETE SET NULL,  -- Nullable, matched later
    image_url TEXT NOT NULL,                       -- S3/GCS URL
    thumbnail_url TEXT,
    ocr_status TEXT DEFAULT 'pending',             -- pending, processing, completed, failed
    ocr_provider TEXT,                              -- azure, google, textract
    ocr_raw_response JSONB,                        -- Raw OCR response
    merchant_name TEXT,
    merchant_address TEXT,
    receipt_date DATE,
    receipt_time TIME,                             -- For precise matching
    total_amount_minor BIGINT,
    tax_amount_minor BIGINT,
    currency CHAR(3) DEFAULT 'USD',
    confidence_score DECIMAL(3,2),                 -- 0.00 to 1.00
    match_status TEXT DEFAULT 'unmatched',         -- unmatched, auto_matched, manual_matched
    match_confidence DECIMAL(3,2),                 -- Confidence of transaction match
    user_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    processed_at TIMESTAMPTZ,
    
    CONSTRAINT valid_ocr_status CHECK (ocr_status IN ('pending', 'processing', 'completed', 'failed')),
    CONSTRAINT valid_match_status CHECK (match_status IN ('unmatched', 'auto_matched', 'manual_matched')),
    CONSTRAINT valid_confidence CHECK (confidence_score >= 0 AND confidence_score <= 1)
);

CREATE INDEX idx_receipts_user_id ON receipts(user_id);
CREATE INDEX idx_receipts_transaction_id ON receipts(transaction_id) WHERE transaction_id IS NOT NULL;
CREATE INDEX idx_receipts_unmatched ON receipts(user_id, receipt_date) WHERE match_status = 'unmatched';
CREATE INDEX idx_receipts_matching ON receipts(user_id, total_amount_minor, receipt_date) WHERE match_status = 'unmatched';

-- Receipt line items (the killer feature)
CREATE TABLE receipt_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    receipt_id UUID NOT NULL REFERENCES receipts(id) ON DELETE CASCADE,
    line_number INTEGER NOT NULL,
    product_name TEXT NOT NULL,
    product_code TEXT,                             -- UPC/SKU if detected
    quantity DECIMAL(10,3) DEFAULT 1,
    unit_price_minor BIGINT,
    total_price_minor BIGINT NOT NULL,
    category_id UUID REFERENCES categories(id),
    is_taxable BOOLEAN DEFAULT FALSE,
    confidence_score DECIMAL(3,2),
    user_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    CONSTRAINT valid_line_number CHECK (line_number > 0),
    CONSTRAINT valid_quantity CHECK (quantity > 0)
);

CREATE INDEX idx_receipt_items_receipt_id ON receipt_items(receipt_id);
CREATE INDEX idx_receipt_items_category_id ON receipt_items(category_id) WHERE category_id IS NOT NULL;

-- Receipt-Transaction Matching Table (for matching candidates and history)
CREATE TABLE receipt_transaction_matches (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    receipt_id UUID NOT NULL REFERENCES receipts(id) ON DELETE CASCADE,
    transaction_id UUID NOT NULL REFERENCES transactions(id) ON DELETE CASCADE,
    match_score DECIMAL(3,2) NOT NULL,             -- 0.00 to 1.00
    match_factors JSONB NOT NULL,                  -- What matched: amount, date, merchant, etc.
    match_type TEXT NOT NULL,                      -- 'auto', 'manual', 'suggested'
    is_active BOOLEAN DEFAULT FALSE,               -- Currently selected match
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by TEXT,                               -- 'system' or user_id
    
    CONSTRAINT valid_match_type CHECK (match_type IN ('auto', 'manual', 'suggested')),
    CONSTRAINT valid_match_score CHECK (match_score >= 0 AND match_score <= 1),
    CONSTRAINT unique_active_match UNIQUE(receipt_id, is_active) WHERE is_active = TRUE
);

CREATE INDEX idx_match_receipt ON receipt_transaction_matches(receipt_id);
CREATE INDEX idx_match_transaction ON receipt_transaction_matches(transaction_id);
CREATE INDEX idx_match_active ON receipt_transaction_matches(receipt_id) WHERE is_active = TRUE;

-- Budgets
CREATE TABLE budgets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    budget_type TEXT NOT NULL,                     -- category, total, envelope
    period_type TEXT NOT NULL,                     -- monthly, weekly, custom
    amount_minor BIGINT NOT NULL,
    currency CHAR(3) DEFAULT 'USD',
    category_id UUID REFERENCES categories(id),    -- For category budgets
    start_date DATE NOT NULL,
    end_date DATE,                                 -- NULL for recurring
    is_active BOOLEAN DEFAULT TRUE,
    rollover_enabled BOOLEAN DEFAULT FALSE,        -- Carry over unused amounts
    alert_enabled BOOLEAN DEFAULT TRUE,
    alert_threshold_percent INTEGER DEFAULT 80,
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    CONSTRAINT valid_budget_type CHECK (budget_type IN ('category', 'total', 'envelope')),
    CONSTRAINT valid_period_type CHECK (period_type IN ('daily', 'weekly', 'monthly', 'yearly', 'custom')),
    CONSTRAINT positive_amount CHECK (amount_minor > 0),
    CONSTRAINT valid_threshold CHECK (alert_threshold_percent BETWEEN 0 AND 100)
);

CREATE INDEX idx_budgets_user_id ON budgets(user_id) WHERE is_active = TRUE;
CREATE INDEX idx_budgets_category_id ON budgets(category_id) WHERE category_id IS NOT NULL;
CREATE INDEX idx_budgets_dates ON budgets(start_date, end_date);

-- Budget tracking snapshots (for historical tracking)
CREATE TABLE budget_periods (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    budget_id UUID NOT NULL REFERENCES budgets(id) ON DELETE CASCADE,
    period_start DATE NOT NULL,
    period_end DATE NOT NULL,
    allocated_amount_minor BIGINT NOT NULL,
    spent_amount_minor BIGINT DEFAULT 0,
    rollover_amount_minor BIGINT DEFAULT 0,        -- From previous period
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    CONSTRAINT unique_budget_period UNIQUE(budget_id, period_start),
    CONSTRAINT valid_period CHECK (period_end > period_start)
);

CREATE INDEX idx_budget_periods_budget_id ON budget_periods(budget_id);
CREATE INDEX idx_budget_periods_dates ON budget_periods(period_start, period_end);

-- Categorization Rules
CREATE TABLE categorization_rules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    rule_name TEXT NOT NULL,
    priority INTEGER DEFAULT 100,
    is_active BOOLEAN DEFAULT TRUE,
    conditions JSONB NOT NULL,                     -- Rule conditions
    category_id UUID NOT NULL REFERENCES categories(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    CONSTRAINT valid_priority CHECK (priority >= 0)
);

CREATE INDEX idx_rules_user_id ON categorization_rules(user_id) WHERE is_active = TRUE;
CREATE INDEX idx_rules_priority ON categorization_rules(user_id, priority DESC);

-- Sync status tracking for offline support
CREATE TABLE sync_queue (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    device_id TEXT NOT NULL,
    operation_type TEXT NOT NULL,                  -- create, update, delete
    entity_type TEXT NOT NULL,                     -- transaction, category, etc
    entity_id UUID NOT NULL,
    payload JSONB NOT NULL,
    sync_status TEXT DEFAULT 'pending',            -- pending, processing, completed, failed
    retry_count INTEGER DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    processed_at TIMESTAMPTZ,
    
    CONSTRAINT valid_operation CHECK (operation_type IN ('create', 'update', 'delete')),
    CONSTRAINT valid_sync_status CHECK (sync_status IN ('pending', 'processing', 'completed', 'failed'))
);

CREATE INDEX idx_sync_queue_user_device ON sync_queue(user_id, device_id) WHERE sync_status = 'pending';
CREATE INDEX idx_sync_queue_status ON sync_queue(sync_status) WHERE sync_status IN ('pending', 'processing');
```

#### 3. Data Integrity Rules (Application Layer)

##### Receipt-Transaction Matching Service
```go
// Matching service in Go - no database triggers
type ReceiptMatchingService struct {
    db *sql.DB
}

// Match transaction to receipts - called after inserting transaction
func (m *ReceiptMatchingService) MatchTransactionToReceipts(tx *Transaction) error {
    // Only match purchase transactions
    if tx.Type != "purchase" {
        return nil
    }
    
    // Find potential receipt matches
    rows, err := m.db.Query(`
        SELECT id, total_amount_minor, receipt_date, merchant_name
        FROM receipts 
        WHERE user_id = $1
          AND match_status = 'unmatched'
          AND receipt_date BETWEEN $2 AND $3
          AND ABS(total_amount_minor - $4) < ($4 * 0.01)
    `, tx.UserID, 
       tx.OccurredAt.AddDate(0, 0, -3),
       tx.OccurredAt.AddDate(0, 0, 1),
       math.Abs(float64(tx.AmountMinor)))
    
    if err != nil {
        return err
    }
    defer rows.Close()
    
    for rows.Next() {
        var receipt Receipt
        rows.Scan(&receipt.ID, &receipt.TotalAmountMinor, &receipt.Date, &receipt.MerchantName)
        
        // Calculate match score
        score := m.calculateMatchScore(tx, &receipt)
        
        // Record match candidate
        _, err = m.db.Exec(`
            INSERT INTO receipt_transaction_matches 
            (receipt_id, transaction_id, match_score, match_factors, match_type, is_active)
            VALUES ($1, $2, $3, $4, $5, $6)
        `, receipt.ID, tx.ID, score, 
           m.getMatchFactors(tx, &receipt),
           m.getMatchType(score),
           score >= 0.80)
        
        // Auto-match if high confidence
        if score >= 0.80 {
            m.applyMatch(receipt.ID, tx.ID, score)
            break // Stop after first high-confidence match
        }
    }
    
    return nil
}

// Match receipt to transactions - called after OCR completes
func (m *ReceiptMatchingService) MatchReceiptToTransactions(receipt *Receipt) error {
    // Only process completed OCR receipts
    if receipt.OCRStatus != "completed" || receipt.TotalAmountMinor == nil {
        return nil
    }
    
    // Find potential transaction matches
    // Similar logic to above but reversed
    // ...
    
    return nil
}

func (m *ReceiptMatchingService) calculateMatchScore(tx *Transaction, receipt *Receipt) float64 {
    score := 0.0
    
    // Amount match (40% weight)
    if receipt.TotalAmountMinor == int(math.Abs(float64(tx.AmountMinor))) {
        score += 0.40
    } else if math.Abs(float64(receipt.TotalAmountMinor-int(math.Abs(float64(tx.AmountMinor))))) < 100 {
        score += 0.30
    }
    
    // Date match (30% weight)
    if receipt.Date.Equal(tx.OccurredAt.Truncate(24*time.Hour)) {
        score += 0.30
    } else if math.Abs(receipt.Date.Sub(tx.OccurredAt).Hours()) <= 24 {
        score += 0.20
    }
    
    // Merchant similarity (30% weight)
    if receipt.MerchantName != "" && tx.MerchantNameRaw != "" {
        similarity := calculateStringSimilarity(receipt.MerchantName, tx.MerchantNameRaw)
        if similarity > 0.7 {
            score += 0.30
        } else if similarity > 0.5 {
            score += 0.15
        }
    }
    
    return score
}

-- Enable fuzzy text matching
CREATE EXTENSION IF NOT EXISTS pg_trgm;

##### Institution Management (Application Layer)
```go
// Institution service - find or create institution when importing from Plaid
type InstitutionService struct {
    db *sql.DB
}

func (s *InstitutionService) FindOrCreateInstitution(
    plaidID string,
    name string,
    userID *string,
) (string, error) {
    var institutionID string
    
    // Check if Plaid institution exists
    if plaidID != "" {
        err := s.db.QueryRow(`
            SELECT id FROM institutions WHERE plaid_institution_id = $1
        `, plaidID).Scan(&institutionID)
        
        if err == nil {
            return institutionID, nil
        }
    }
    
    // Check for system institution by name (fuzzy match)
    err := s.db.QueryRow(`
        SELECT id FROM institutions
        WHERE is_system = TRUE
          AND similarity(lower(name), lower($1)) > 0.8
        ORDER BY similarity(lower(name), lower($1)) DESC
        LIMIT 1
    `, name).Scan(&institutionID)
    
    if err == nil {
        return institutionID, nil
    }
    
    // Create new institution
    err = s.db.QueryRow(`
        INSERT INTO institutions (plaid_institution_id, name, display_name, is_system)
        VALUES ($1, $2, $2, FALSE)
        RETURNING id
    `, plaidID, name).Scan(&institutionID)
    
    return institutionID, err
}
$$ LANGUAGE plpgsql;

-- Pre-populate common institutions (including store cards)
INSERT INTO institutions (name, display_name, logo_url, website_url, is_system) VALUES
    -- Major Banks
    ('Chase Bank', 'Chase', 'https://logos/chase.png', 'https://chase.com', TRUE),
    ('Bank of America', 'Bank of America', 'https://logos/bofa.png', 'https://bankofamerica.com', TRUE),
    ('Wells Fargo', 'Wells Fargo', 'https://logos/wellsfargo.png', 'https://wellsfargo.com', TRUE),
    ('Citi', 'Citi', 'https://logos/citi.png', 'https://citi.com', TRUE),
    ('Capital One', 'Capital One', 'https://logos/capitalone.png', 'https://capitalone.com', TRUE),
    -- Credit Card Companies
    ('American Express', 'American Express', 'https://logos/amex.png', 'https://americanexpress.com', TRUE),
    ('Discover', 'Discover', 'https://logos/discover.png', 'https://discover.com', TRUE),
    -- Store Cards
    ('Target', 'Target RedCard', 'https://logos/target.png', 'https://target.com', TRUE),
    ('Amazon', 'Amazon Store Card', 'https://logos/amazon.png', 'https://amazon.com', TRUE),
    ('Walmart', 'Walmart Card', 'https://logos/walmart.png', 'https://walmart.com', TRUE),
    -- Investment
    ('Charles Schwab', 'Schwab', 'https://logos/schwab.png', 'https://schwab.com', TRUE),
    ('Fidelity', 'Fidelity', 'https://logos/fidelity.png', 'https://fidelity.com', TRUE),
    ('Vanguard', 'Vanguard', 'https://logos/vanguard.png', 'https://vanguard.com', TRUE)
ON CONFLICT DO NOTHING;

-- Manual credit card statement update
CREATE OR REPLACE FUNCTION update_manual_credit_statement(
    p_account_id UUID,
    p_statement_date DATE,
    p_balance_minor BIGINT,
    p_minimum_payment_minor BIGINT,
    p_due_date DATE
) error {
BEGIN
    INSERT INTO credit_card_liabilities (
        account_id,
        last_statement_date,
        last_statement_balance_minor,
        minimum_payment_amount_minor,
        next_payment_due_date,
        updated_at
    ) VALUES (
        p_account_id,
        p_statement_date,
        p_balance_minor,
        p_minimum_payment_minor,
        p_due_date,
        NOW()
    )
    ON CONFLICT (account_id) DO UPDATE SET
        last_statement_date = EXCLUDED.last_statement_date,
        last_statement_balance_minor = EXCLUDED.last_statement_balance_minor,
        minimum_payment_amount_minor = EXCLUDED.minimum_payment_amount_minor,
        next_payment_due_date = EXCLUDED.next_payment_due_date,
        updated_at = NOW();
        
    -- Update account balance
    UPDATE accounts 
    SET current_balance_minor = p_balance_minor,
        updated_at = NOW()
    WHERE id = p_account_id;
END;
}
```

##### Business Rules Enforcement
```sql
-- Update account balance from transactions
func (s *BalanceService) UpdateAccountBalance(accountID string) error {
DECLARE
    v_balance BIGINT;
    v_account_type TEXT;
BEGIN
    -- Get account type
    SELECT account_type INTO v_account_type
    FROM accounts
    WHERE id = COALESCE(NEW.account_id, OLD.account_id);
    
    -- Calculate balance differently for credit cards vs bank accounts
    IF v_account_type = 'credit' THEN
        -- Credit card: sum of all transactions (purchases increase debt, payments decrease it)
        SELECT COALESCE(SUM(
            CASE 
                WHEN transaction_type IN ('purchase', 'fee', 'interest') THEN ABS(amount_minor)  -- Increases balance
                WHEN transaction_type = 'payment' THEN -ABS(amount_minor)  -- Decreases balance
                ELSE amount_minor
            END
        ), 0) INTO v_balance
        FROM transactions
        WHERE account_id = COALESCE(NEW.account_id, OLD.account_id)
          AND NOT pending;
    ELSE
        -- Bank account: simple sum (deposits positive, withdrawals negative)
        SELECT COALESCE(SUM(amount_minor), 0) INTO v_balance
        FROM transactions
        WHERE account_id = COALESCE(NEW.account_id, OLD.account_id)
          AND NOT pending;
    END IF;
    
    UPDATE accounts 
    SET current_balance_minor = v_balance,
        updated_at = NOW()
    WHERE id = COALESCE(NEW.account_id, OLD.account_id);
    
    RETURN COALESCE(NEW, OLD);
END;
$$ LANGUAGE plpgsql;

// Called after transaction operations
    AFTER INSERT OR UPDATE OR DELETE ON transactions
    FOR EACH ROW EXECUTE FUNCTION update_account_balance();

-- Track budget spending (notification only, no blocking)
func (s *BudgetService) CheckBudgetLimit(tx *Transaction) error {
DECLARE
    v_spent BIGINT;
    v_budget BIGINT;
BEGIN
    -- Only check for purchase transactions
    IF NEW.transaction_type != 'purchase' THEN
        RETURN NEW;
    END IF;
    
    -- Get current period spending
    SELECT COALESCE(SUM(ABS(amount_minor)), 0) INTO v_spent
    FROM transactions
    WHERE category_id = NEW.category_id
      AND user_id = NEW.user_id
      AND transaction_type = 'purchase'
      AND DATE_TRUNC('month', occurred_at) = DATE_TRUNC('month', NEW.occurred_at);
    
    -- Get budget limit
    SELECT b.amount_minor INTO v_budget
    FROM budgets b
    WHERE b.category_id = NEW.category_id
      AND b.user_id = NEW.user_id
      AND b.is_active = TRUE
      AND NEW.occurred_at BETWEEN b.start_date AND COALESCE(b.end_date, '9999-12-31');
    
    -- Log budget alert if exceeded
    IF v_budget IS NOT NULL AND v_spent > v_budget THEN
        INSERT INTO budget_alerts (user_id, budget_id, amount_exceeded, created_at)
        SELECT NEW.user_id, b.id, v_spent - v_budget, NOW()
        FROM budgets b
        WHERE b.category_id = NEW.category_id
          AND b.user_id = NEW.user_id
          AND b.is_active = TRUE;
    END IF;
    
    RETURN NEW;
END;
}
```

#### 4. GraphQL Schema

```graphql
# Core Types
type User {
  id: ID!
  email: String!
  displayName: String
  avatarUrl: String
  subscriptionTier: SubscriptionTier!
  institutions: [Institution!]!
  accounts: [Account!]!
  categories: [Category!]!
  budgets: [Budget!]!
  transactions(
    first: Int
    after: String
    filter: TransactionFilter
  ): TransactionConnection!
  receipts(
    first: Int
    after: String
    filter: ReceiptFilter
  ): ReceiptConnection!
  stats: UserStats!
  createdAt: DateTime!
}

type Institution {
  id: ID!
  name: String!
  displayName: String!
  logoUrl: String
  websiteUrl: String
  isSystem: Boolean!
  accounts: [Account!]!
  # User customizations
  customName: String
  customColor: String
  totalBalance: Money!
  accountCount: Int!
}

type Account {
  id: ID!
  institution: Institution!
  accountType: AccountType!
  accountName: String!
  accountNumberMask: String
  currency: String!
  currentBalance: Money!
  availableBalance: Money
  creditLimit: Money  # For credit cards
  # Status fields
  isActive: Boolean!
  isManual: Boolean!
  syncStatus: SyncStatus!
  lastSyncedAt: DateTime
  # Credit card liability details (null for non-credit accounts)
  creditCardDetails: CreditCardLiability
  transactions(
    first: Int
    after: String
    filter: TransactionFilter
  ): TransactionConnection!
}

type CreditCardLiability {
  # APR Information
  aprPurchase: Float
  aprCashAdvance: Float
  aprBalanceTransfer: Float
  balanceSubjectToApr: Money
  interestChargeAmount: Money
  # Statement Information
  lastStatementDate: Date
  lastStatementBalance: Money
  statementCloseDate: Date
  # Payment Information
  minimumPaymentAmount: Money!
  nextPaymentDueDate: Date!
  lastPaymentDate: Date
  lastPaymentAmount: Money
  isOverdue: Boolean!
  daysOverdue: Int
  # Calculated fields
  utilizationPercentage: Float  # balance / credit limit
  paymentStatus: PaymentStatus!
}

type Transaction {
  id: ID!
  type: TransactionType!
  description: String!
  merchant: Merchant
  amount: Money!
  account: Account!
  category: Category
  occurredAt: DateTime!
  receipt: Receipt
  receiptItems: [ReceiptItem!]!  # Direct access to itemized details
  pending: Boolean!
  importSource: ImportSource!
  notes: String
  metadata: JSON
}

type Category {
  id: ID!
  name: String!
  parent: Category
  children: [Category!]!
  icon: String
  color: String
  isSystem: Boolean!
  isIncome: Boolean!
  monthlyAverage: Money
  rules: [CategorizationRule!]!
}

type Receipt {
  id: ID!
  imageUrl: String!
  thumbnailUrl: String
  ocrStatus: OCRStatus!
  merchantName: String
  merchantAddress: String
  receiptDate: Date
  receiptTime: Time
  totalAmount: Money
  taxAmount: Money
  items: [ReceiptItem!]!
  confidence: Float
  matchStatus: MatchStatus!
  matchConfidence: Float
  transaction: Transaction  # Current matched transaction
  suggestedMatches: [TransactionMatch!]!  # Potential matches
  userVerified: Boolean!
}

type TransactionMatch {
  transaction: Transaction!
  matchScore: Float!
  matchFactors: JSON!
  matchType: MatchType!
}

type ReceiptItem {
  id: ID!
  lineNumber: Int!
  productName: String!
  productCode: String
  quantity: Float!
  unitPrice: Money
  totalPrice: Money!
  category: Category
  isTaxable: Boolean!
  confidence: Float
  userVerified: Boolean!
}

type Budget {
  id: ID!
  name: String!
  type: BudgetType!
  period: BudgetPeriod!
  amount: Money!
  category: Category
  startDate: Date!
  endDate: Date
  isActive: Boolean!
  currentPeriod: BudgetPeriodSnapshot
  alertEnabled: Boolean!
  alertThreshold: Int!
}

type BudgetPeriodSnapshot {
  periodStart: Date!
  periodEnd: Date!
  allocated: Money!
  spent: Money!
  remaining: Money!
  percentUsed: Float!
  projectedEndAmount: Money
  isOnTrack: Boolean!
}

# Value Objects
type Money {
  amount: Float!  # Converted from minor units for display
  currency: String!
  formatted: String!  # e.g., "$1,234.56"
}

# Enums
enum SubscriptionTier {
  FREE
  PREMIUM
  PREMIUM_PLUS
}

enum AccountType {
  CHECKING
  SAVINGS
  CREDIT
  INVESTMENT
  LOAN
  OTHER
}

enum TransactionType {
  PURCHASE
  TRANSFER
  INCOME
  FEE
  ADJUSTMENT
  INTEREST
}

enum SyncStatus {
  PENDING
  SYNCING
  READY
  ERROR
}

enum OCRStatus {
  PENDING
  PROCESSING
  COMPLETED
  FAILED
}

enum ImportSource {
  PLAID
  MANUAL
  CSV
  OFX
}

enum MatchStatus {
  UNMATCHED
  AUTO_MATCHED
  MANUAL_MATCHED
}

enum MatchType {
  AUTO
  MANUAL
  SUGGESTED
}

enum BudgetType {
  CATEGORY
  TOTAL
  ENVELOPE
}

enum BudgetPeriod {
  DAILY
  WEEKLY
  MONTHLY
  YEARLY
  CUSTOM
}

# Input Types
input TransactionFilter {
  accountIds: [ID!]
  categoryIds: [ID!]
  merchantIds: [ID!]
  dateFrom: DateTime
  dateTo: DateTime
  amountMin: Float
  amountMax: Float
  searchTerm: String
  hasReceipt: Boolean
}

input CreateManualTransactionInput {
  accountId: ID!
  amount: Float!
  type: TransactionType!
  description: String!
  merchantId: ID
  categoryId: ID
  occurredAt: DateTime!
  notes: String
  idempotencyKey: ID!
}

input ImportTransactionsInput {
  accountId: ID!
  source: ImportSource!
  transactions: [TransactionData!]!
  idempotencyKey: ID!
}

input UpdateTransactionInput {
  id: ID!
  categoryId: ID
  description: String
  notes: String
}

# Mutations
type Mutation {
  # Institutions
  createCustomInstitution(input: CreateInstitutionInput!): Institution!
  updateInstitutionSettings(input: UpdateInstitutionSettingsInput!): Institution!
  
  # Accounts
  connectPlaidAccount(input: ConnectPlaidAccountInput!): [Account!]!  # May create multiple accounts
  createManualAccount(input: CreateManualAccountInput!): Account!
  updateAccount(input: UpdateAccountInput!): Account!
  syncAccount(id: ID!): Account!
  disconnectAccount(id: ID!): Boolean!
  
  # Credit Card Management
  updateCreditCardStatement(input: UpdateCreditCardStatementInput!): CreditCardLiability!
  recordCreditCardPayment(input: RecordPaymentInput!): Transaction!
  
  # Transactions
  createManualTransaction(input: CreateManualTransactionInput!): Transaction!
  importTransactions(input: ImportTransactionsInput!): [Transaction!]!
  updateTransaction(input: UpdateTransactionInput!): Transaction!
  deleteTransaction(id: ID!): Boolean!
  
  # Receipts
  uploadReceipt(transactionId: ID, image: Upload!): Receipt!
  matchReceiptToTransaction(receiptId: ID!, transactionId: ID!): Receipt!
  unmatchReceipt(receiptId: ID!): Receipt!
  acceptSuggestedMatch(receiptId: ID!, transactionId: ID!): Receipt!
  verifyReceiptItem(itemId: ID!, verified: Boolean!): ReceiptItem!
  
  # Categories
  createCategory(input: CreateCategoryInput!): Category!
  updateCategory(input: UpdateCategoryInput!): Category!
  deleteCategory(id: ID!): Boolean!
  
  # Budgets
  createBudget(input: CreateBudgetInput!): Budget!
  updateBudget(input: UpdateBudgetInput!): Budget!
  deleteBudget(id: ID!): Boolean!
}

# Queries
type Query {
  # User
  me: User!
  
  # Institutions
  institutions: [Institution!]!
  institution(id: ID!): Institution
  searchInstitutions(query: String!): [Institution!]!
  
  # Accounts
  account(id: ID!): Account
  accounts(filter: AccountFilter): [Account!]!
  accountsByInstitution(institutionId: ID!): [Account!]!
  
  # Transactions
  transaction(id: ID!): Transaction
  transactions(
    first: Int = 20
    after: String
    filter: TransactionFilter
  ): TransactionConnection!
  
  # Categories
  categories(includeSystem: Boolean = true): [Category!]!
  categoryTree: [Category!]!
  
  # Budgets
  budget(id: ID!): Budget
  budgets(active: Boolean = true): [Budget!]!
  budgetSummary(period: BudgetPeriod!): BudgetSummary!
  
  # Receipts
  receipt(id: ID!): Receipt
  receipts(
    first: Int = 20
    after: String
    filter: ReceiptFilter
  ): ReceiptConnection!
  
  # Analytics
  spendingByCategory(
    period: AnalyticsPeriod!
    accountIds: [ID!]
  ): [CategorySpending!]!
  
  cashFlow(
    period: AnalyticsPeriod!
    accountIds: [ID!]
  ): CashFlowAnalysis!
}

# Subscriptions
type Subscription {
  accountSyncStatus(accountId: ID!): SyncStatus!
  receiptProcessingStatus(receiptId: ID!): OCRStatus!
  budgetAlert(budgetId: ID!): BudgetAlert!
}
```

#### 5. Data Privacy & Security

##### Simplified Security Model
```sql
-- PostgreSQL Row-Level Security for multi-tenancy
-- No email encryption needed (OAuth providers handle authentication)
-- Focus on financial data protection

-- Enable RLS on all user-scoped tables
ALTER TABLE accounts ENABLE ROW LEVEL SECURITY;
ALTER TABLE transactions ENABLE ROW LEVEL SECURITY;
ALTER TABLE categories ENABLE ROW LEVEL SECURITY;
ALTER TABLE budgets ENABLE ROW LEVEL SECURITY;
ALTER TABLE receipts ENABLE ROW LEVEL SECURITY;

-- Simple user isolation policy
CREATE POLICY user_isolation ON transactions
    FOR ALL
    USING (user_id = current_setting('app.current_user')::uuid);

-- Apply same policy to all tables
DO $$
DECLARE
    t TEXT;
BEGIN
    FOR t IN SELECT tablename FROM pg_tables 
    WHERE schemaname = 'public' 
    AND tablename IN ('accounts', 'categories', 'budgets', 'receipts')
    LOOP
        EXECUTE format('CREATE POLICY user_isolation ON %I FOR ALL USING (user_id = current_setting(''app.current_user'')::uuid)', t);
    END LOOP;
END $$;
```

##### Implementation in API Layer
```go
// Set user context for RLS in Go/gqlgen middleware
func SetUserContext(ctx context.Context, db *sql.DB, userID string) error {
    _, err := db.ExecContext(ctx, "SET LOCAL app.current_user = $1", userID)
    return err
}

-- Audit logging
CREATE TABLE audit_log (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    event_type TEXT NOT NULL,
    entity_type TEXT,
    entity_id UUID,
    old_values JSONB,
    new_values JSONB,
    ip_address INET,
    user_agent TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_audit_user ON audit_log(user_id, created_at DESC);
CREATE INDEX idx_audit_entity ON audit_log(entity_type, entity_id);
```

##### Basic Privacy Compliance (MVP)
```sql
-- Simple user data export
CREATE VIEW user_data_export AS
SELECT 
    u.id,
    u.email,
    u.display_name,
    json_agg(DISTINCT a.*) as accounts,
    json_agg(DISTINCT t.*) as transactions,
    json_agg(DISTINCT c.*) as categories,
    json_agg(DISTINCT r.*) as receipts
FROM users u
LEFT JOIN accounts a ON a.user_id = u.id
LEFT JOIN transactions t ON t.user_id = u.id
LEFT JOIN categories c ON c.user_id = u.id
LEFT JOIN receipts r ON r.user_id = u.id
GROUP BY u.id;

-- Simple account deletion (soft delete)
CREATE OR REPLACE FUNCTION delete_user_account(p_user_id UUID)
error {
BEGIN
    -- Soft delete user
    UPDATE users 
    SET email = concat('deleted_', id, '@deleted.com'),
        display_name = 'Deleted User',
        firebase_uid = concat('deleted_', id),
        deleted_at = NOW()
    WHERE id = p_user_id;
    
    -- Cascade will handle related data due to ON DELETE CASCADE
    -- Or keep data for analytics by anonymizing:
    UPDATE transactions 
    SET description = 'Anonymized',
        notes = NULL,
        metadata = '{}'
    WHERE user_id = p_user_id;
    
    RETURN TRUE;
END;
$$ LANGUAGE plpgsql SECURITY DEFINER;
```

## Consequences

### Positive
- **Simplicity**: Straightforward transaction model for imported data
- **Performance**: UUIDv7 provides time-ordered keys improving index performance
- **Scalability**: Simple schema scales easily to millions of transactions
- **Privacy**: Row-level security ensures complete user data isolation
- **Flexibility**: JSONB fields allow schema evolution without migrations
- **Deduplication**: External IDs prevent duplicate imports automatically
- **Receipt Correlation**: Clean relationship between transactions and itemized receipts
- **Type Safety**: GraphQL schema provides strong typing across stack
- **OAuth Simplicity**: No email verification or password management needed

### Negative
- **Limited Financial Features**: No support for creating transactions (import-only)
- **Receipt Dependency**: Full value requires OCR processing for receipts
- **No Transfer Tracking**: Simplified model makes internal transfers less explicit
- **Basic Compliance**: MVP-level privacy features only

### Risks
- **Import Accuracy**: Dependent on quality of bank/Plaid data
- **Receipt Matching**: Correlating receipts to transactions may require manual intervention
- **Privacy Evolution**: May need to add GDPR/CCPA features post-MVP
- **RLS Performance**: Row-level security adds small overhead to queries

## Alternatives Considered

### Option A: Double-Entry Ledger System
- **Description**: Full accounting system with journals and entries
- **Pros**: Complete financial integrity, supports transaction creation
- **Cons**: Over-engineered for import-only system, complex implementation
- **Reason for not choosing**: Unnecessary complexity for read-only transaction data

### Option B: NoSQL (MongoDB)
- **Description**: Document database for flexibility
- **Pros**: Easier schema evolution, better for unstructured receipt data
- **Cons**: No ACID guarantees, complex aggregations, poor GraphQL integration
- **Reason for not choosing**: Financial data requires strong consistency

### Option C: Email/Password Authentication
- **Description**: Traditional email verification and password management
- **Pros**: Familiar pattern, works offline, no OAuth dependency
- **Cons**: Security burden, email verification complexity, password resets
- **Reason for not choosing**: OAuth providers handle security better, simpler UX

## Implementation Notes

### Migration Strategy
1. Create schema with all tables and constraints
2. Populate system categories and default data
3. Set up RLS policies and audit triggers
4. Configure database backup and replication
5. Load test with simulated data (1M+ transactions)

### Rollback Plan
- All DDL changes wrapped in transactions
- Backup before any production migration
- Blue-green deployment for zero-downtime updates
- Feature flags for gradual rollout

### Testing Approach
- Unit tests for all database functions
- Integration tests for GraphQL resolvers
- Property-based testing for financial calculations
- Load testing for concurrent user scenarios

### Monitoring and Success Metrics
- Query performance (p95 < 100ms)
- Sync conflict rate (< 0.1%)
- Data integrity violations (0 tolerance)
- Storage growth rate tracking

## References
- [PostgreSQL 16 Documentation](https://www.postgresql.org/docs/16/)
- [Double-Entry Bookkeeping in Postgres](https://www.moderntreasury.com/journal/accounting-for-developers-part-i)
- [GraphQL Cursor Pagination](https://relay.dev/graphql/connections.htm)
- [Row Level Security Best Practices](https://www.postgresql.org/docs/current/ddl-rowsecurity.html)
- ADR-20250812: Personal Finance Tech Stack Selection
- [UUIDv7 Specification](https://www.ietf.org/rfc/rfc9562.html#name-uuid-version-7)

## Decision Makers
- Author: Archie (System Architect)
- Reviewers: [Pending]
- Approver: [Pending]
- Date: 2025-08-19

## Revision History
- 2025-08-19: Initial comprehensive draft