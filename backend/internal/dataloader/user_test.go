package dataloader

import (
	"context"
	"testing"
	"time"

	"github.com/fernandobandeira/djinn/backend/internal/database"
	"github.com/fernandobandeira/djinn/backend/internal/database/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// MockQueries implements a mock for testing
type MockQueries struct {
	getUsersByIDsCalls int
	users             []db.User
}

func (m *MockQueries) GetUsersByIDs(ctx context.Context, ids []pgtype.UUID) ([]db.User, error) {
	m.getUsersByIDsCalls++
	return m.users, nil
}

func TestUserDataLoader(t *testing.T) {
	ctx := context.Background()
	
	// Create test users
	user1ID := uuid.New()
	user2ID := uuid.New()
	
	testUsers := []db.User{
		{
			ID: pgtype.UUID{Bytes: user1ID, Valid: true},
			FirebaseUid: "firebase1",
			Email: "user1@example.com",
			Name: "User One",
			CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
			UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		},
		{
			ID: pgtype.UUID{Bytes: user2ID, Valid: true},
			FirebaseUid: "firebase2", 
			Email: "user2@example.com",
			Name: "User Two",
			CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
			UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		},
	}
	
	// Create mock database
	mockQueries := &MockQueries{users: testUsers}
	mockDB := &database.DB{
		Queries: mockQueries,
	}
	
	// Create reader and loader
	reader := NewUserReader(mockDB)
	loader := NewUserLoader(UserLoaderConfig{
		Fetch: func(ids []string) ([]*db.User, []error) {
			return reader.GetUsers(ctx, ids)
		},
		Wait:     1 * time.Millisecond,
		MaxBatch: 100,
	})
	
	// Test loading single user
	t.Run("LoadSingleUser", func(t *testing.T) {
		user, err := loader.Load(user1ID.String())
		if err != nil {
			t.Fatalf("Failed to load user: %v", err)
		}
		if user == nil {
			t.Fatal("User should not be nil")
		}
		if user.Email != "user1@example.com" {
			t.Errorf("Expected email user1@example.com, got %s", user.Email)
		}
	})
	
	// Test batch loading
	t.Run("BatchLoadUsers", func(t *testing.T) {
		// Reset call count
		mockQueries.getUsersByIDsCalls = 0
		
		// Create new loader for this test
		loader := NewUserLoader(UserLoaderConfig{
			Fetch: func(ids []string) ([]*db.User, []error) {
				return reader.GetUsers(ctx, ids)
			},
			Wait:     10 * time.Millisecond,
			MaxBatch: 100,
		})
		
		// Request multiple users concurrently
		done := make(chan bool, 2)
		
		go func() {
			loader.Load(user1ID.String())
			done <- true
		}()
		
		go func() {
			loader.Load(user2ID.String())
			done <- true
		}()
		
		// Wait for both requests
		<-done
		<-done
		
		// Wait for batching
		time.Sleep(15 * time.Millisecond)
		
		// Check that only one batch query was made
		if mockQueries.getUsersByIDsCalls != 1 {
			t.Errorf("Expected 1 batch query, got %d", mockQueries.getUsersByIDsCalls)
		}
	})
	
	// Test caching
	t.Run("CachingUsers", func(t *testing.T) {
		mockQueries.getUsersByIDsCalls = 0
		
		// Load user twice
		loader.Load(user1ID.String())
		loader.Load(user1ID.String())
		
		time.Sleep(5 * time.Millisecond)
		
		// Should only make one query due to caching
		if mockQueries.getUsersByIDsCalls != 1 {
			t.Errorf("Expected 1 query due to caching, got %d", mockQueries.getUsersByIDsCalls)
		}
	})
}

// TestDataLoaderMiddleware tests the middleware integration
func TestDataLoaderMiddleware(t *testing.T) {
	// Create mock database
	mockDB := &database.DB{}
	
	// Create context with middleware
	ctx := context.Background()
	ctx = LoaderMiddleware(mockDB, func(ctx context.Context) context.Context {
		return ctx
	})(ctx)
	
	// Verify loader is in context
	loader := GetUserLoader(ctx)
	if loader == nil {
		t.Fatal("DataLoader should be in context")
	}
}