# Database Schema Management Workflow

## Overview
We use a **SQL-first** approach where both Atlas (migrations) and sqlc (code generation) share the same schema source.

## Workflow Diagram
```
schema.sql (Source of Truth)
    ├── Atlas → Generates migrations → Apply to database
    └── sqlc  → Generates Go code → Type-safe queries
```

## How It Works

### 1. sqlc Process:
- **Reads**: `schema.sql` + `queries/*.sql`
- **Analyzes**: Table structures, column types, relationships
- **Generates**: 
  - `User` struct from `users` table
  - `Account` struct from `accounts` table
  - Type-safe query methods

### 2. Atlas Process:
- **Reads**: `schema.sql`
- **Compares**: Current database state vs desired schema
- **Generates**: Migration files to sync database

## The Workflow

### Step 1: Modify Schema
Edit `schema.sql` with your desired database structure.

### Step 2: Generate Migration
```bash
# Atlas compares schema.sql to current database and creates migration
atlas migrate diff <migration_name> \
  --to file://schema.sql \
  --dev-url "docker://postgres/16/dev?search_path=public"
```

### Step 3: Apply Migration
```bash
# Update the actual database
atlas migrate apply --env local
```

### Step 4: Generate Go Code
```bash
# sqlc reads the same schema.sql and generates types
sqlc generate
```

## Why This Works Well

1. **Single Source of Truth**: One `schema.sql` file defines everything
2. **Type Safety**: sqlc knows exact table structure from schema
3. **Migration Safety**: Atlas tracks what's been applied
4. **No Manual Sync**: Both tools read the same file

## Example

When `schema.sql` has:
```sql
CREATE TABLE users (
    id UUID PRIMARY KEY,
    email VARCHAR(255),
    name VARCHAR(255)
);
```

sqlc generates:
```go
type User struct {
    ID    uuid.UUID
    Email string
    Name  string
}
```

And Atlas knows to create the table in the database!