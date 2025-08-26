package dataloader

import (
	"context"
	"fmt"
	"time"

	"github.com/fernandobandeira/djinn/backend/internal/database"
	db "github.com/fernandobandeira/djinn/backend/internal/database/generated"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

//go:generate go run github.com/vektah/dataloaden UserLoader string *github.com/fernandobandeira/djinn/backend/internal/database/generated.User

type UserReader struct {
	db *database.DB
}

func NewUserReader(database *database.DB) *UserReader {
	return &UserReader{
		db: database,
	}
}

func (u *UserReader) GetUsers(ctx context.Context, ids []string) ([]*db.User, []error) {
	usersByID := make(map[string]*db.User, len(ids))
	errors := make([]error, len(ids))
	
	// Convert string IDs to UUIDs
	uuids := make([]pgtype.UUID, 0, len(ids))
	idToIndex := make(map[string]int)
	
	for i, id := range ids {
		userID, err := uuid.Parse(id)
		if err != nil {
			errors[i] = fmt.Errorf("invalid user ID: %w", err)
			continue
		}
		
		pgUUID := pgtype.UUID{
			Bytes: userID,
			Valid: true,
		}
		uuids = append(uuids, pgUUID)
		idToIndex[id] = i
	}
	
	// Batch fetch users from database
	if len(uuids) > 0 {
		users, err := u.db.Queries.GetUsersByIDs(ctx, uuids)
		if err != nil {
			// If batch query fails, set error for all requested IDs
			for i := range ids {
				errors[i] = fmt.Errorf("failed to fetch users: %w", err)
			}
			return make([]*db.User, len(ids)), errors
		}
		
		// Map users by their ID
		for _, user := range users {
			if user.ID.Valid {
				id, err := uuid.FromBytes(user.ID.Bytes[:])
				if err == nil {
					userCopy := user // Create a copy to avoid pointer issues
					usersByID[id.String()] = &userCopy
				}
			}
		}
	}
	
	// Build result slice maintaining the order of requested IDs
	result := make([]*db.User, len(ids))
	for i, id := range ids {
		if user, ok := usersByID[id]; ok {
			result[i] = user
		} else if errors[i] == nil {
			// User not found and no parse error
			errors[i] = fmt.Errorf("user not found: %s", id)
		}
	}
	
	return result, errors
}

// LoaderMiddleware creates a new DataLoader middleware
func LoaderMiddleware(database *database.DB, next func(context.Context) context.Context) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		reader := NewUserReader(database)
		loader := NewUserLoader(UserLoaderConfig{
			// Batch function to load users
			Fetch: func(ids []string) ([]*db.User, []error) {
				return reader.GetUsers(ctx, ids)
			},
			// Wait duration before executing batch
			Wait: 1 * time.Millisecond,
			// Maximum batch size
			MaxBatch: 100,
		})
		
		ctx = context.WithValue(ctx, userLoaderKey, loader)
		return next(ctx)
	}
}

type contextKey string

const userLoaderKey contextKey = "userLoader"

// GetUserLoader returns the UserLoader from context
func GetUserLoader(ctx context.Context) (*UserLoader, error) {
	loader, ok := ctx.Value(userLoaderKey).(*UserLoader)
	if !ok {
		return nil, fmt.Errorf("UserLoader not found in context")
	}
	return loader, nil
}

// LoadUser loads a single user by ID using DataLoader
func LoadUser(ctx context.Context, id string) (*db.User, error) {
	loader, err := GetUserLoader(ctx)
	if err != nil {
		return nil, err
	}
	return loader.Load(id)
}

// LoadUsers loads multiple users by IDs using DataLoader
func LoadUsers(ctx context.Context, ids []string) ([]*db.User, []error) {
	loader, err := GetUserLoader(ctx)
	if err != nil {
		errors := make([]error, len(ids))
		for i := range errors {
			errors[i] = err
		}
		return make([]*db.User, len(ids)), errors
	}
	return loader.LoadAll(ids)
}