# DataLoader Implementation

## Overview
DataLoader is a utility to batch and cache database queries to prevent N+1 query problems in GraphQL resolvers.

## Features
- **Batching**: Combines multiple individual database queries into a single batch query
- **Caching**: Caches results within a single request to avoid duplicate queries
- **Type Safety**: Generated code provides type-safe loaders for each entity

## Usage

### In Resolvers
When you need to load a user by ID, use the DataLoader from context instead of direct database queries:

```go
// Instead of:
user, err := r.DB.Queries.GetUser(ctx, userID)

// Use:
user, err := dataloader.LoadUser(ctx, userIDString)
```

### Batch Loading
The DataLoader automatically batches multiple requests made within a short time window (1ms by default):

```go
// These requests will be batched into a single database query
user1, err1 := dataloader.LoadUser(ctx, "uuid1")
user2, err2 := dataloader.LoadUser(ctx, "uuid2")
user3, err3 := dataloader.LoadUser(ctx, "uuid3")
```

## Configuration
- **Wait Duration**: 1ms - Time to wait before executing a batch
- **Max Batch Size**: 100 - Maximum number of items in a single batch query

## Adding New DataLoaders
To add a DataLoader for a new entity:

1. Create a new file in this package (e.g., `account.go`)
2. Add the go:generate directive:
   ```go
   //go:generate go run github.com/vektah/dataloaden AccountLoader string *github.com/fernandobandeira/djinn/backend/internal/database/db.Account
   ```
3. Implement the reader with batch query
4. Add middleware integration
5. Run `go generate ./...`

## Performance Benefits
- Reduces database round trips
- Prevents N+1 query problems
- Improves GraphQL API response times
- Reduces database load