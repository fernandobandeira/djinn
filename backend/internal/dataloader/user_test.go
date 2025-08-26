package dataloader

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/fernandobandeira/djinn/backend/internal/database/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// MockDB is a mock database for testing
type MockDB struct {
	mock.Mock
}

func (m *MockDB) GetUsersByIDs(ctx context.Context, ids []uuid.UUID) ([]db.User, error) {
	args := m.Called(ctx, ids)
	return args.Get(0).([]db.User), args.Error(1)
}

func TestNewUserLoader(t *testing.T) {
	mockDB := new(MockDB)
	
	loader := NewUserLoader(mockDB)
	assert.NotNil(t, loader)
}

func TestUserLoader_Load(t *testing.T) {
	mockDB := new(MockDB)
	loader := NewUserLoader(mockDB)
	
	ctx := context.Background()
	userID := uuid.New()
	now := time.Now()
	
	t.Run("Successful load", func(t *testing.T) {
		mockDB.ExpectedCalls = nil
		
		expectedUser := db.User{
			ID:          userID,
			FirebaseUid: "firebase123",
			Email:       "test@example.com",
			Name:        "Test User",
			CreatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
			UpdatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
		}
		
		mockDB.On("GetUsersByIDs", ctx, []uuid.UUID{userID}).Return(
			[]db.User{expectedUser},
			nil,
		).Once()
		
		user, err := loader.Load(ctx, userID.String())
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, userID, user.ID)
		assert.Equal(t, "test@example.com", user.Email)
		
		// Wait for batch to process
		time.Sleep(10 * time.Millisecond)
		mockDB.AssertExpectations(t)
	})
	
	t.Run("Invalid UUID", func(t *testing.T) {
		user, err := loader.Load(ctx, "invalid-uuid")
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "invalid UUID")
	})
	
	t.Run("User not found", func(t *testing.T) {
		mockDB.ExpectedCalls = nil
		
		notFoundID := uuid.New()
		mockDB.On("GetUsersByIDs", ctx, []uuid.UUID{notFoundID}).Return(
			[]db.User{},
			nil,
		).Once()
		
		user, err := loader.Load(ctx, notFoundID.String())
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "user not found")
		
		// Wait for batch to process
		time.Sleep(10 * time.Millisecond)
		mockDB.AssertExpectations(t)
	})
	
	t.Run("Database error", func(t *testing.T) {
		mockDB.ExpectedCalls = nil
		
		errorID := uuid.New()
		mockDB.On("GetUsersByIDs", ctx, []uuid.UUID{errorID}).Return(
			[]db.User{},
			errors.New("database error"),
		).Once()
		
		user, err := loader.Load(ctx, errorID.String())
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "database error")
		
		// Wait for batch to process
		time.Sleep(10 * time.Millisecond)
		mockDB.AssertExpectations(t)
	})
}

func TestUserLoader_LoadAll(t *testing.T) {
	mockDB := new(MockDB)
	loader := NewUserLoader(mockDB)
	
	ctx := context.Background()
	userID1 := uuid.New()
	userID2 := uuid.New()
	now := time.Now()
	
	t.Run("Successful load multiple", func(t *testing.T) {
		mockDB.ExpectedCalls = nil
		
		expectedUsers := []db.User{
			{
				ID:          userID1,
				FirebaseUid: "firebase1",
				Email:       "test1@example.com",
				Name:        "Test User 1",
				CreatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
				UpdatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
			},
			{
				ID:          userID2,
				FirebaseUid: "firebase2",
				Email:       "test2@example.com",
				Name:        "Test User 2",
				CreatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
				UpdatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
			},
		}
		
		mockDB.On("GetUsersByIDs", ctx, mock.MatchedBy(func(ids []uuid.UUID) bool {
			return len(ids) == 2
		})).Return(expectedUsers, nil).Once()
		
		users, errs := loader.LoadAll(ctx, []string{userID1.String(), userID2.String()})
		
		// Wait for batch to process
		time.Sleep(10 * time.Millisecond)
		
		assert.Len(t, users, 2)
		assert.Len(t, errs, 2)
		assert.NoError(t, errs[0])
		assert.NoError(t, errs[1])
		assert.NotNil(t, users[0])
		assert.NotNil(t, users[1])
		
		mockDB.AssertExpectations(t)
	})
	
	t.Run("Mixed valid and invalid UUIDs", func(t *testing.T) {
		users, errs := loader.LoadAll(ctx, []string{userID1.String(), "invalid-uuid"})
		
		assert.Len(t, users, 2)
		assert.Len(t, errs, 2)
		
		// First might succeed if cached or batched
		// Second should fail
		assert.Error(t, errs[1])
		assert.Contains(t, errs[1].Error(), "invalid UUID")
	})
	
	t.Run("Empty input", func(t *testing.T) {
		users, errs := loader.LoadAll(ctx, []string{})
		assert.Empty(t, users)
		assert.Empty(t, errs)
	})
}

func TestGetUserLoader(t *testing.T) {
	mockDB := new(MockDB)
	loader := NewUserLoader(mockDB)
	
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
	mockDB := new(MockDB)
	loader := NewUserLoader(mockDB)
	
	userID := uuid.New()
	now := time.Now()
	
	t.Run("Successful load", func(t *testing.T) {
		mockDB.ExpectedCalls = nil
		
		ctx := context.WithValue(context.Background(), userLoaderKey, loader)
		
		expectedUser := db.User{
			ID:          userID,
			FirebaseUid: "firebase123",
			Email:       "test@example.com",
			Name:        "Test User",
			CreatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
			UpdatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
		}
		
		mockDB.On("GetUsersByIDs", ctx, []uuid.UUID{userID}).Return(
			[]db.User{expectedUser},
			nil,
		).Once()
		
		user, err := LoadUser(ctx, userID.String())
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, userID, user.ID)
		
		// Wait for batch to process
		time.Sleep(10 * time.Millisecond)
		mockDB.AssertExpectations(t)
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
	mockDB := new(MockDB)
	loader := NewUserLoader(mockDB)
	
	userID1 := uuid.New()
	userID2 := uuid.New()
	now := time.Now()
	
	t.Run("Successful load multiple", func(t *testing.T) {
		mockDB.ExpectedCalls = nil
		
		ctx := context.WithValue(context.Background(), userLoaderKey, loader)
		
		expectedUsers := []db.User{
			{
				ID:          userID1,
				FirebaseUid: "firebase1",
				Email:       "test1@example.com",
				Name:        "Test User 1",
				CreatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
				UpdatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
			},
			{
				ID:          userID2,
				FirebaseUid: "firebase2",
				Email:       "test2@example.com",
				Name:        "Test User 2",
				CreatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
				UpdatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
			},
		}
		
		mockDB.On("GetUsersByIDs", ctx, mock.MatchedBy(func(ids []uuid.UUID) bool {
			return len(ids) == 2
		})).Return(expectedUsers, nil).Once()
		
		users, errs := LoadUsers(ctx, []string{userID1.String(), userID2.String()})
		
		// Wait for batch to process
		time.Sleep(10 * time.Millisecond)
		
		assert.Len(t, users, 2)
		assert.Len(t, errs, 2)
		assert.NoError(t, errs[0])
		assert.NoError(t, errs[1])
		
		mockDB.AssertExpectations(t)
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

func TestUserLoader_BatchFunction(t *testing.T) {
	mockDB := new(MockDB)
	loader := NewUserLoader(mockDB)
	
	ctx := context.Background()
	userID1 := uuid.New()
	userID2 := uuid.New()
	userID3 := uuid.New()
	now := time.Now()
	
	t.Run("Batch function processes correctly", func(t *testing.T) {
		mockDB.ExpectedCalls = nil
		
		keys := []string{userID1.String(), userID2.String(), userID3.String()}
		
		expectedUsers := []db.User{
			{
				ID:          userID1,
				FirebaseUid: "firebase1",
				Email:       "test1@example.com",
				Name:        "Test User 1",
				CreatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
				UpdatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
			},
			// userID2 is missing (simulating not found)
			{
				ID:          userID3,
				FirebaseUid: "firebase3",
				Email:       "test3@example.com",
				Name:        "Test User 3",
				CreatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
				UpdatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
			},
		}
		
		mockDB.On("GetUsersByIDs", ctx, mock.MatchedBy(func(ids []uuid.UUID) bool {
			return len(ids) == 3
		})).Return(expectedUsers, nil).Once()
		
		results, errs := loader.batchFunc(ctx, keys)
		
		require.Len(t, results, 3)
		require.Len(t, errs, 3)
		
		// First user found
		assert.NotNil(t, results[0])
		assert.NoError(t, errs[0])
		assert.Equal(t, userID1, results[0].ID)
		
		// Second user not found
		assert.Nil(t, results[1])
		assert.Error(t, errs[1])
		assert.Contains(t, errs[1].Error(), "user not found")
		
		// Third user found
		assert.NotNil(t, results[2])
		assert.NoError(t, errs[2])
		assert.Equal(t, userID3, results[2].ID)
		
		mockDB.AssertExpectations(t)
	})
	
	t.Run("Batch function handles database error", func(t *testing.T) {
		mockDB.ExpectedCalls = nil
		
		keys := []string{userID1.String()}
		
		mockDB.On("GetUsersByIDs", ctx, mock.Anything).Return(
			[]db.User{},
			errors.New("database connection failed"),
		).Once()
		
		results, errs := loader.batchFunc(ctx, keys)
		
		require.Len(t, results, 1)
		require.Len(t, errs, 1)
		
		assert.Nil(t, results[0])
		assert.Error(t, errs[0])
		assert.Contains(t, errs[0].Error(), "database connection failed")
		
		mockDB.AssertExpectations(t)
	})
	
	t.Run("Batch function handles invalid UUID", func(t *testing.T) {
		keys := []string{"invalid-uuid", userID1.String()}
		
		results, errs := loader.batchFunc(ctx, keys)
		
		require.Len(t, results, 2)
		require.Len(t, errs, 2)
		
		// First key is invalid
		assert.Nil(t, results[0])
		assert.Error(t, errs[0])
		assert.Contains(t, errs[0].Error(), "invalid UUID")
		
		// Second key might be processed depending on implementation
	})
}