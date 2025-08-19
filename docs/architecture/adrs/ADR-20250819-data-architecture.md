# ADR-20250819: Data Architecture and Schema Design

## Status
Accepted

## Context
The Djinn personal finance application requires a robust data architecture to handle:
- Multi-user financial data with strict isolation
- Complex financial entities (accounts, transactions, budgets, receipts with items)
- High precision monetary calculations without floating-point errors
- Receipt itemization for granular spending insights
- Double-entry bookkeeping for data integrity
- Offline-first mobile synchronization
- GDPR/CCPA compliance for personal financial data
- Scale to 100K+ users with millions of transactions

### Constraints
- PostgreSQL 16+ as the primary database (decided in tech stack ADR)
- GraphQL API requiring efficient query patterns
- Mobile clients need offline support with eventual consistency
- Financial regulations require audit trails and data retention
- PII must be encrypted and properly isolated
- MVP timeline requires pragmatic choices over perfection

## Decision

### Core Data Model

#### 1. Entity Hierarchy
```
Users (root tenant)
├── Accounts (checking, savings, credit cards, investments)
├── Categories (hierarchical, user-customizable)
├── Budgets (monthly/weekly/custom periods)
├── Rules (auto-categorization, alerts)
├── Receipts (images + OCR data)
└── Merchants (shared but user-enrichable)

Transactions (financial events)
├── Ledger Journal (atomic financial events)
├── Ledger Entries (double-entry records)
├── Receipt Items (itemized purchases)
└── Attachments (receipts, documents)
```

#### 2. PostgreSQL Schema Design

##### Core Tables

```sql
-- Users table (root of all multi-tenancy)
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),  -- UUIDv7 when available
    firebase_uid TEXT UNIQUE NOT NULL,              -- Firebase Auth UID
    email TEXT NOT NULL,
    email_verified BOOLEAN DEFAULT FALSE,
    display_name TEXT,
    avatar_url TEXT,
    subscription_tier TEXT DEFAULT 'free',           -- free, premium, premium_plus
    subscription_expires_at TIMESTAMPTZ,
    settings JSONB DEFAULT '{}',                    -- User preferences
    metadata JSONB DEFAULT '{}',                    -- Analytics, features flags
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,                         -- Soft delete for GDPR
    
    CONSTRAINT valid_email CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}$'),
    CONSTRAINT valid_subscription CHECK (subscription_tier IN ('free', 'premium', 'premium_plus'))
);

CREATE INDEX idx_users_firebase_uid ON users(firebase_uid) WHERE deleted_at IS NULL;
CREATE INDEX idx_users_email ON users(email) WHERE deleted_at IS NULL;

-- Financial Accounts
CREATE TABLE accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    plaid_account_id TEXT,                          -- External ID from Plaid
    account_type TEXT NOT NULL,                     -- checking, savings, credit, investment
    institution_name TEXT,
    account_name TEXT NOT NULL,
    account_number_mask TEXT,                       -- Last 4 digits only
    currency CHAR(3) NOT NULL DEFAULT 'USD',
    current_balance_minor BIGINT NOT NULL DEFAULT 0, -- In cents/minor units
    available_balance_minor BIGINT,
    credit_limit_minor BIGINT,                      -- For credit cards
    is_active BOOLEAN DEFAULT TRUE,
    is_manual BOOLEAN DEFAULT FALSE,                -- Manual vs connected account
    sync_status TEXT DEFAULT 'pending',             -- pending, syncing, ready, error
    last_synced_at TIMESTAMPTZ,
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    CONSTRAINT valid_account_type CHECK (account_type IN ('checking', 'savings', 'credit', 'investment', 'loan', 'other')),
    CONSTRAINT valid_currency CHECK (currency ~ '^[A-Z]{3}$'),
    CONSTRAINT valid_sync_status CHECK (sync_status IN ('pending', 'syncing', 'ready', 'error'))
);

CREATE INDEX idx_accounts_user_id ON accounts(user_id) WHERE is_active = TRUE;
CREATE INDEX idx_accounts_plaid_id ON accounts(plaid_account_id) WHERE plaid_account_id IS NOT NULL;

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

-- Double-Entry Ledger Journal (groups related entries)
CREATE TABLE ledger_journals (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    transaction_type TEXT NOT NULL,                 -- purchase, transfer, income, adjustment
    description TEXT NOT NULL,
    merchant_id UUID REFERENCES merchants(id),
    occurred_at TIMESTAMPTZ NOT NULL,
    idempotency_key UUID UNIQUE,                   -- Client-generated for deduplication
    external_id TEXT,                              -- Plaid transaction ID
    metadata JSONB DEFAULT '{}',                   -- Additional transaction data
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    CONSTRAINT valid_transaction_type CHECK (transaction_type IN ('purchase', 'transfer', 'income', 'fee', 'adjustment', 'interest'))
);

CREATE INDEX idx_journals_user_date ON ledger_journals(user_id, occurred_at DESC);
CREATE INDEX idx_journals_external_id ON ledger_journals(external_id) WHERE external_id IS NOT NULL;
CREATE INDEX idx_journals_idempotency ON ledger_journals(idempotency_key) WHERE idempotency_key IS NOT NULL;

-- Double-Entry Ledger Entries (must balance to zero per journal)
CREATE TABLE ledger_entries (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    journal_id UUID NOT NULL REFERENCES ledger_journals(id) ON DELETE CASCADE,
    account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE RESTRICT,
    entry_type TEXT NOT NULL,                      -- debit or credit
    amount_minor BIGINT NOT NULL,                  -- Always positive, type determines sign
    currency CHAR(3) NOT NULL DEFAULT 'USD',
    category_id UUID REFERENCES categories(id),
    notes TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    CONSTRAINT valid_entry_type CHECK (entry_type IN ('debit', 'credit')),
    CONSTRAINT positive_amount CHECK (amount_minor > 0),
    CONSTRAINT valid_currency CHECK (currency ~ '^[A-Z]{3}$')
);

CREATE INDEX idx_entries_journal_id ON ledger_entries(journal_id);
CREATE INDEX idx_entries_account_id ON ledger_entries(account_id);
CREATE INDEX idx_entries_category_id ON ledger_entries(category_id) WHERE category_id IS NOT NULL;

-- Enforce double-entry balance
CREATE OR REPLACE FUNCTION check_journal_balance() RETURNS TRIGGER AS $$
BEGIN
    IF (
        SELECT SUM(CASE WHEN entry_type = 'debit' THEN amount_minor ELSE -amount_minor END)
        FROM ledger_entries
        WHERE journal_id = NEW.journal_id
    ) != 0 THEN
        RAISE EXCEPTION 'Journal entries must balance to zero';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE CONSTRAINT TRIGGER enforce_journal_balance
    AFTER INSERT OR UPDATE ON ledger_entries
    DEFERRABLE INITIALLY DEFERRED
    FOR EACH ROW EXECUTE FUNCTION check_journal_balance();

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

-- Receipts and OCR data
CREATE TABLE receipts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    journal_id UUID REFERENCES ledger_journals(id) ON DELETE CASCADE,
    image_url TEXT NOT NULL,                       -- S3/GCS URL
    thumbnail_url TEXT,
    ocr_status TEXT DEFAULT 'pending',             -- pending, processing, completed, failed
    ocr_provider TEXT,                              -- azure, google, textract
    ocr_raw_response JSONB,                        -- Raw OCR response
    merchant_name TEXT,
    merchant_address TEXT,
    receipt_date DATE,
    total_amount_minor BIGINT,
    tax_amount_minor BIGINT,
    currency CHAR(3) DEFAULT 'USD',
    confidence_score DECIMAL(3,2),                 -- 0.00 to 1.00
    user_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    processed_at TIMESTAMPTZ,
    
    CONSTRAINT valid_ocr_status CHECK (ocr_status IN ('pending', 'processing', 'completed', 'failed')),
    CONSTRAINT valid_confidence CHECK (confidence_score >= 0 AND confidence_score <= 1)
);

CREATE INDEX idx_receipts_user_id ON receipts(user_id);
CREATE INDEX idx_receipts_journal_id ON receipts(journal_id) WHERE journal_id IS NOT NULL;
CREATE INDEX idx_receipts_status ON receipts(ocr_status) WHERE ocr_status IN ('pending', 'processing');

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

#### 3. Data Integrity Rules

##### Business Rules Enforcement
```sql
-- Ensure account balances match ledger entries
CREATE OR REPLACE FUNCTION update_account_balance() RETURNS TRIGGER AS $$
DECLARE
    v_balance BIGINT;
BEGIN
    SELECT COALESCE(SUM(
        CASE 
            WHEN le.entry_type = 'debit' AND a.account_type IN ('checking', 'savings') THEN le.amount_minor
            WHEN le.entry_type = 'credit' AND a.account_type IN ('checking', 'savings') THEN -le.amount_minor
            WHEN le.entry_type = 'credit' AND a.account_type = 'credit' THEN le.amount_minor
            WHEN le.entry_type = 'debit' AND a.account_type = 'credit' THEN -le.amount_minor
            ELSE 0
        END
    ), 0) INTO v_balance
    FROM ledger_entries le
    JOIN accounts a ON a.id = le.account_id
    WHERE le.account_id = NEW.account_id;
    
    UPDATE accounts 
    SET current_balance_minor = v_balance,
        updated_at = NOW()
    WHERE id = NEW.account_id;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_balance_on_entry
    AFTER INSERT OR UPDATE OR DELETE ON ledger_entries
    FOR EACH ROW EXECUTE FUNCTION update_account_balance();

-- Prevent budget overspending (optional enforcement)
CREATE OR REPLACE FUNCTION check_budget_limit() RETURNS TRIGGER AS $$
DECLARE
    v_spent BIGINT;
    v_budget BIGINT;
BEGIN
    -- Only check for purchase transactions
    IF NEW.transaction_type != 'purchase' THEN
        RETURN NEW;
    END IF;
    
    -- Get current period spending
    SELECT COALESCE(SUM(le.amount_minor), 0) INTO v_spent
    FROM ledger_entries le
    JOIN ledger_journals lj ON lj.id = le.journal_id
    WHERE le.category_id = NEW.category_id
      AND lj.user_id = NEW.user_id
      AND DATE_TRUNC('month', lj.occurred_at) = DATE_TRUNC('month', NEW.occurred_at);
    
    -- Get budget limit
    SELECT b.amount_minor INTO v_budget
    FROM budgets b
    WHERE b.category_id = NEW.category_id
      AND b.user_id = NEW.user_id
      AND b.is_active = TRUE
      AND NEW.occurred_at BETWEEN b.start_date AND COALESCE(b.end_date, '9999-12-31');
    
    -- This could either RAISE EXCEPTION or just log a warning
    IF v_budget IS NOT NULL AND (v_spent + NEW.amount_minor) > v_budget THEN
        -- Log budget exceeded event instead of blocking
        INSERT INTO audit_log (user_id, event_type, details)
        VALUES (NEW.user_id, 'budget_exceeded', jsonb_build_object(
            'category_id', NEW.category_id,
            'budget_amount', v_budget,
            'new_total', v_spent + NEW.amount_minor
        ));
    END IF;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
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

type Account {
  id: ID!
  accountType: AccountType!
  institutionName: String
  accountName: String!
  accountNumberMask: String
  currency: String!
  currentBalance: Money!
  availableBalance: Money
  creditLimit: Money
  isActive: Boolean!
  isManual: Boolean!
  syncStatus: SyncStatus!
  lastSyncedAt: DateTime
  transactions(
    first: Int
    after: String
    filter: TransactionFilter
  ): TransactionConnection!
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
  entries: [LedgerEntry!]!
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
  totalAmount: Money
  taxAmount: Money
  items: [ReceiptItem!]!
  confidence: Float
  userVerified: Boolean!
  transaction: Transaction
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

input CreateTransactionInput {
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

input UpdateTransactionInput {
  id: ID!
  categoryId: ID
  description: String
  notes: String
}

# Mutations
type Mutation {
  # Transactions
  createTransaction(input: CreateTransactionInput!): Transaction!
  updateTransaction(input: UpdateTransactionInput!): Transaction!
  deleteTransaction(id: ID!): Boolean!
  
  # Receipts
  uploadReceipt(transactionId: ID, image: Upload!): Receipt!
  verifyReceiptItem(itemId: ID!, verified: Boolean!): ReceiptItem!
  
  # Categories
  createCategory(input: CreateCategoryInput!): Category!
  updateCategory(input: UpdateCategoryInput!): Category!
  deleteCategory(id: ID!): Boolean!
  
  # Budgets
  createBudget(input: CreateBudgetInput!): Budget!
  updateBudget(input: UpdateBudgetInput!): Budget!
  deleteBudget(id: ID!): Boolean!
  
  # Accounts
  connectAccount(input: ConnectAccountInput!): Account!
  updateAccount(input: UpdateAccountInput!): Account!
  syncAccount(id: ID!): Account!
  disconnectAccount(id: ID!): Boolean!
}

# Queries
type Query {
  # User
  me: User!
  
  # Accounts
  account(id: ID!): Account
  accounts(filter: AccountFilter): [Account!]!
  
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

##### PII Handling
```sql
-- Encryption at rest via PostgreSQL TDE (Transparent Data Encryption)
-- Configure pgcrypto for sensitive fields

-- Create encrypted email storage
CREATE EXTENSION IF NOT EXISTS pgcrypto;

ALTER TABLE users 
  ADD COLUMN email_encrypted BYTEA,
  ADD COLUMN email_hash TEXT GENERATED ALWAYS AS (
    encode(digest(lower(email), 'sha256'), 'hex')
  ) STORED;

-- Function to encrypt PII
CREATE OR REPLACE FUNCTION encrypt_pii(plain_text TEXT, key_id TEXT) 
RETURNS BYTEA AS $$
DECLARE
    encryption_key BYTEA;
BEGIN
    -- In production, retrieve key from AWS KMS or similar
    SELECT key_data INTO encryption_key 
    FROM encryption_keys 
    WHERE id = key_id AND is_active = true;
    
    RETURN pgp_sym_encrypt(plain_text, encryption_key::text);
END;
$$ LANGUAGE plpgsql SECURITY DEFINER;
```

##### Data Isolation
```sql
-- Row Level Security for multi-tenancy
ALTER TABLE accounts ENABLE ROW LEVEL SECURITY;
ALTER TABLE ledger_journals ENABLE ROW LEVEL SECURITY;
ALTER TABLE categories ENABLE ROW LEVEL SECURITY;
ALTER TABLE budgets ENABLE ROW LEVEL SECURITY;
ALTER TABLE receipts ENABLE ROW LEVEL SECURITY;

-- Policy for user data isolation
CREATE POLICY user_isolation ON accounts
    FOR ALL
    USING (user_id = current_setting('app.current_user')::uuid);

CREATE POLICY user_isolation ON ledger_journals
    FOR ALL
    USING (user_id = current_setting('app.current_user')::uuid);

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

##### GDPR/CCPA Compliance
```sql
-- User data export
CREATE OR REPLACE FUNCTION export_user_data(p_user_id UUID)
RETURNS JSONB AS $$
DECLARE
    result JSONB;
BEGIN
    SELECT jsonb_build_object(
        'user', row_to_json(u.*),
        'accounts', jsonb_agg(DISTINCT a.*),
        'transactions', jsonb_agg(DISTINCT t.*),
        'categories', jsonb_agg(DISTINCT c.*),
        'budgets', jsonb_agg(DISTINCT b.*),
        'receipts', jsonb_agg(DISTINCT r.*)
    ) INTO result
    FROM users u
    LEFT JOIN accounts a ON a.user_id = u.id
    LEFT JOIN ledger_journals t ON t.user_id = u.id
    LEFT JOIN categories c ON c.user_id = u.id
    LEFT JOIN budgets b ON b.user_id = u.id
    LEFT JOIN receipts r ON r.user_id = u.id
    WHERE u.id = p_user_id
    GROUP BY u.id;
    
    RETURN result;
END;
$$ LANGUAGE plpgsql SECURITY DEFINER;

-- User data deletion (right to be forgotten)
CREATE OR REPLACE FUNCTION delete_user_data(p_user_id UUID)
RETURNS BOOLEAN AS $$
BEGIN
    -- Soft delete with anonymization
    UPDATE users 
    SET email = concat('deleted_', id, '@deleted.com'),
        display_name = 'Deleted User',
        firebase_uid = concat('deleted_', id),
        deleted_at = NOW()
    WHERE id = p_user_id;
    
    -- Remove PII from related tables
    DELETE FROM receipts WHERE user_id = p_user_id;
    DELETE FROM receipt_items WHERE receipt_id IN (
        SELECT id FROM receipts WHERE user_id = p_user_id
    );
    
    -- Anonymize financial data (keep for regulatory requirements)
    UPDATE ledger_journals 
    SET description = 'Anonymized',
        metadata = '{}'
    WHERE user_id = p_user_id;
    
    RETURN TRUE;
END;
$$ LANGUAGE plpgsql SECURITY DEFINER;
```

## Consequences

### Positive
- **Data Integrity**: Double-entry bookkeeping ensures financial accuracy
- **Performance**: UUIDv7 provides time-ordered keys improving index performance
- **Scalability**: Partitioning strategy ready for millions of transactions
- **Privacy**: Row-level security ensures complete user data isolation
- **Flexibility**: JSONB fields allow schema evolution without migrations
- **Audit Trail**: Complete history with soft deletes and audit logging
- **Offline Support**: Sync queue enables robust offline-first mobile experience
- **Type Safety**: GraphQL schema provides strong typing across stack

### Negative
- **Complexity**: Double-entry ledger adds complexity vs simple transaction table
- **Storage**: Storing both journal and entries increases storage requirements
- **Learning Curve**: Team needs to understand double-entry accounting principles
- **Migration Effort**: Complex schema requires careful data migration planning

### Risks
- **Performance**: Complex joins between ledger tables may need optimization
- **Consistency**: Offline sync conflicts need careful resolution strategies
- **Encryption Overhead**: PII encryption may impact query performance
- **Schema Evolution**: Some changes may require complex migrations

## Alternatives Considered

### Option A: Single Transaction Table
- **Description**: Simple transactions table without double-entry
- **Pros**: Simpler to implement, easier to understand
- **Cons**: No built-in integrity checks, harder to track transfers
- **Reason for not choosing**: Financial integrity is critical for user trust

### Option B: NoSQL (MongoDB)
- **Description**: Document database for flexibility
- **Pros**: Easier schema evolution, better for unstructured receipt data
- **Cons**: No ACID guarantees, complex aggregations, poor GraphQL integration
- **Reason for not choosing**: Financial data requires strong consistency

### Option C: Event Sourcing
- **Description**: Store all changes as events
- **Pros**: Complete audit trail, temporal queries, event replay
- **Cons**: Complex implementation, eventual consistency, steep learning curve
- **Reason for not choosing**: Over-engineered for MVP timeline

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