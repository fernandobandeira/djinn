package dataloader

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/persistence/postgres"
	db "github.com/fernandobandeira/djinn/backend/internal/infrastructure/persistence/postgres/generated"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUserLoader(t *testing.T) {
	loader := NewUserLoader(UserLoaderConfig{
		Fetch: func(keys []string) ([]*db.User, []error) {
			return make([]*db.User, len(keys)), make([]error, len(keys))
		},
		Wait:     1 * time.Millisecond,
		MaxBatch: 100,
	})
	assert.NotNil(t, loader)
}

func TestUserLoader_Load(t *testing.T) {
	userID := uuid.New()
	now := time.Now()
	
	t.Run("Successful load", func(t *testing.T) {
		pgUUID := pgtype.UUID{
			Bytes: userID,
			Valid: true,
		}
		
		expectedUser := db.User{
			ID:          pgUUID,
			FirebaseUid: "firebase123",
			Email:       "test@example.com",
			Name:        "Test User",
			CreatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
			UpdatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
		}
		
		loader := NewUserLoader(UserLoaderConfig{
			Fetch: func(keys []string) ([]*db.User, []error) {
				if len(keys) > 0 && keys[0] == userID.String() {
					return []*db.User{&expectedUser}, []error{nil}
				}
				return []*db.User{nil}, []error{errors.New("not found")}
			},
			Wait:     1 * time.Millisecond,
			MaxBatch: 100,
		})
		
		user, err := loader.Load(userID.String())
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, pgUUID, user.ID)
		assert.Equal(t, "test@example.com", user.Email)
	})
	
	t.Run("User not found", func(t *testing.T) {
		notFoundID := uuid.New()
		
		loader := NewUserLoader(UserLoaderConfig{
			Fetch: func(keys []string) ([]*db.User, []error) {
				errs := make([]error, len(keys))
				for i := range keys {
					errs[i] = errors.New("user not found: " + keys[i])
				}
				return make([]*db.User, len(keys)), errs
			},
			Wait:     1 * time.Millisecond,
			MaxBatch: 100,
		})
		
		user, err := loader.Load(notFoundID.String())
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "user not found")
	})
	
	t.Run("Database error", func(t *testing.T) {
		errorID := uuid.New()
		
		loader := NewUserLoader(UserLoaderConfig{
			Fetch: func(keys []string) ([]*db.User, []error) {
				errs := make([]error, len(keys))
				for i := range keys {
					errs[i] = errors.New("database error")
				}
				return make([]*db.User, len(keys)), errs
			},
			Wait:     1 * time.Millisecond,
			MaxBatch: 100,
		})
		
		user, err := loader.Load(errorID.String())
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "database error")
	})
}

func TestUserLoader_LoadAll(t *testing.T) {
	userID1 := uuid.New()
	userID2 := uuid.New()
	now := time.Now()
	
	t.Run("Successful load multiple", func(t *testing.T) {
		pgUUID1 := pgtype.UUID{
			Bytes: userID1,
			Valid: true,
		}
		pgUUID2 := pgtype.UUID{
			Bytes: userID2,
			Valid: true,
		}
		
		expectedUsers := []*db.User{
			{
				ID:          pgUUID1,
				FirebaseUid: "firebase1",
				Email:       "test1@example.com",
				Name:        "Test User 1",
				CreatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
				UpdatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
			},
			{
				ID:          pgUUID2,
				FirebaseUid: "firebase2",
				Email:       "test2@example.com",
				Name:        "Test User 2",
				CreatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
				UpdatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
			},
		}
		
		loader := NewUserLoader(UserLoaderConfig{
			Fetch: func(keys []string) ([]*db.User, []error) {
				results := make([]*db.User, len(keys))
				errs := make([]error, len(keys))
				for i, key := range keys {
					if key == userID1.String() {
						results[i] = expectedUsers[0]
					} else if key == userID2.String() {
						results[i] = expectedUsers[1]
					} else {
						errs[i] = errors.New("not found")
					}
				}
				return results, errs
			},
			Wait:     1 * time.Millisecond,
			MaxBatch: 100,
		})
		
		users, errs := loader.LoadAll([]string{userID1.String(), userID2.String()})
		
		assert.Len(t, users, 2)
		assert.Len(t, errs, 2)
		assert.NoError(t, errs[0])
		assert.NoError(t, errs[1])
		assert.NotNil(t, users[0])
		assert.NotNil(t, users[1])
	})
	
	t.Run("Empty input", func(t *testing.T) {
		loader := NewUserLoader(UserLoaderConfig{
			Fetch: func(keys []string) ([]*db.User, []error) {
				return make([]*db.User, len(keys)), make([]error, len(keys))
			},
			Wait:     1 * time.Millisecond,
			MaxBatch: 100,
		})
		
		users, errs := loader.LoadAll([]string{})
		assert.Empty(t, users)
		assert.Empty(t, errs)
	})
}

func TestGetUserLoader(t *testing.T) {
	loader := NewUserLoader(UserLoaderConfig{
		Fetch: func(keys []string) ([]*db.User, []error) {
			return make([]*db.User, len(keys)), make([]error, len(keys))
		},
		Wait:     1 * time.Millisecond,
		MaxBatch: 100,
	})
	
	t.Run("Loader in context", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), userLoaderKey, loader)
		
		retrievedLoader, err := GetUserLoader(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, retrievedLoader)
		assert.Same(t, loader, retrievedLoader)
	})
	
	t.Run("Loader not in context", func(t *testing.T) {
		ctx := context.Background()
		
		retrievedLoader, err := GetUserLoader(ctx)
		assert.Error(t, err)
		assert.Nil(t, retrievedLoader)
		assert.EqualError(t, err, "UserLoader not found in context")
	})
	
	t.Run("Wrong type in context", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), userLoaderKey, "not a loader")
		
		retrievedLoader, err := GetUserLoader(ctx)
		assert.Error(t, err)
		assert.Nil(t, retrievedLoader)
	})
}

func TestLoadUser(t *testing.T) {
	userID := uuid.New()
	now := time.Now()
	
	t.Run("Successful load", func(t *testing.T) {
		pgUUID := pgtype.UUID{
			Bytes: userID,
			Valid: true,
		}
		
		expectedUser := db.User{
			ID:          pgUUID,
			FirebaseUid: "firebase123",
			Email:       "test@example.com",
			Name:        "Test User",
			CreatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
			UpdatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
		}
		
		loader := NewUserLoader(UserLoaderConfig{
			Fetch: func(keys []string) ([]*db.User, []error) {
				if len(keys) > 0 && keys[0] == userID.String() {
					return []*db.User{&expectedUser}, []error{nil}
				}
				return []*db.User{nil}, []error{errors.New("not found")}
			},
			Wait:     1 * time.Millisecond,
			MaxBatch: 100,
		})
		
		ctx := context.WithValue(context.Background(), userLoaderKey, loader)
		
		user, err := LoadUser(ctx, userID.String())
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, pgUUID, user.ID)
	})
	
	t.Run("No loader in context", func(t *testing.T) {
		ctx := context.Background()
		
		user, err := LoadUser(ctx, userID.String())
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.EqualError(t, err, "UserLoader not found in context")
	})
}

func TestLoadUsers(t *testing.T) {
	userID1 := uuid.New()
	userID2 := uuid.New()
	now := time.Now()
	
	t.Run("Successful load multiple", func(t *testing.T) {
		pgUUID1 := pgtype.UUID{
			Bytes: userID1,
			Valid: true,
		}
		pgUUID2 := pgtype.UUID{
			Bytes: userID2,
			Valid: true,
		}
		
		expectedUsers := []*db.User{
			{
				ID:          pgUUID1,
				FirebaseUid: "firebase1",
				Email:       "test1@example.com",
				Name:        "Test User 1",
				CreatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
				UpdatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
			},
			{
				ID:          pgUUID2,
				FirebaseUid: "firebase2",
				Email:       "test2@example.com",
				Name:        "Test User 2",
				CreatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
				UpdatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
			},
		}
		
		loader := NewUserLoader(UserLoaderConfig{
			Fetch: func(keys []string) ([]*db.User, []error) {
				results := make([]*db.User, len(keys))
				errs := make([]error, len(keys))
				for i, key := range keys {
					if key == userID1.String() {
						results[i] = expectedUsers[0]
					} else if key == userID2.String() {
						results[i] = expectedUsers[1]
					} else {
						errs[i] = errors.New("not found")
					}
				}
				return results, errs
			},
			Wait:     1 * time.Millisecond,
			MaxBatch: 100,
		})
		
		ctx := context.WithValue(context.Background(), userLoaderKey, loader)
		
		users, errs := LoadUsers(ctx, []string{userID1.String(), userID2.String()})
		
		assert.Len(t, users, 2)
		assert.Len(t, errs, 2)
		assert.NoError(t, errs[0])
		assert.NoError(t, errs[1])
	})
	
	t.Run("No loader in context", func(t *testing.T) {
		ctx := context.Background()
		
		ids := []string{userID1.String(), userID2.String()}
		users, errs := LoadUsers(ctx, ids)
		
		assert.Len(t, users, 2)
		assert.Len(t, errs, 2)
		assert.Nil(t, users[0])
		assert.Nil(t, users[1])
		assert.Error(t, errs[0])
		assert.Error(t, errs[1])
		assert.EqualError(t, errs[0], "UserLoader not found in context")
		assert.EqualError(t, errs[1], "UserLoader not found in context")
	})
}

func TestUserReader_GetUsers(t *testing.T) {
	ctx := context.Background()
	
	t.Run("Successful batch fetch", func(t *testing.T) {
		// This test would require a real database connection or mocked queries
		// Since we're testing the reader directly, we'd need to mock the db.Queries
		// For now, we'll skip the implementation test and focus on unit tests
		t.Skip("Requires database mock setup")
	})
	
	t.Run("Invalid UUID handling", func(t *testing.T) {
		reader := &UserReader{db: &postgres.DB{}}
		
		ids := []string{"invalid-uuid", "also-invalid"}
		users, errs := reader.GetUsers(ctx, ids)
		
		require.Len(t, users, 2)
		require.Len(t, errs, 2)
		
		assert.Nil(t, users[0])
		assert.Nil(t, users[1])
		assert.Error(t, errs[0])
		assert.Error(t, errs[1])
		assert.Contains(t, errs[0].Error(), "invalid user ID")
		assert.Contains(t, errs[1].Error(), "invalid user ID")
	})
}

func TestLoaderMiddleware(t *testing.T) {
	db := &postgres.DB{}
	
	nextCalled := false
	next := func(ctx context.Context) context.Context {
		nextCalled = true
		return ctx
	}
	
	middleware := LoaderMiddleware(db, next)
	ctx := middleware(context.Background())
	
	assert.True(t, nextCalled, "Next function should be called")
	
	loader, err := GetUserLoader(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, loader)
}