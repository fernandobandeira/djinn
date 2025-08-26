package resolver

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/fernandobandeira/djinn/backend/internal/database"
	"github.com/fernandobandeira/djinn/backend/internal/database/db"
	"github.com/fernandobandeira/djinn/backend/internal/database/mocks"
	"github.com/fernandobandeira/djinn/backend/internal/graph/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	userID := uuid.New()
	now := time.Now()
	
	tests := []struct {
		name        string
		input       model.CreateUserInput
		mockSetup   func(*mocks.MockQuerier)
		expectedErr bool
		errMessage  string
	}{
		{
			name: "Successful user creation",
			input: model.CreateUserInput{
				FirebaseUID:     "firebase123456789012345678",
				Email:          "test@example.com",
				Name:           "Test User",
				ProfileImageURL: stringPtr("https://example.com/image.jpg"),
			},
			mockSetup: func(mockQuerier *mocks.MockQuerier) {
				expectedParams := db.CreateUserParams{
					FirebaseUid: "firebase123456789012345678",
					Email:       "test@example.com",
					Name:        "Test User",
					ProfileImageUrl: pgtype.Text{
						String: "https://example.com/image.jpg",
						Valid:  true,
					},
				}
				
				mockQuerier.EXPECT().CreateUser(ctx, expectedParams).Return(
					db.User{
						ID: pgtype.UUID{
							Bytes: userID,
							Valid: true,
						},
						FirebaseUid:     "firebase123456789012345678",
						Email:           "test@example.com",
						Name:            "Test User",
						ProfileImageUrl: pgtype.Text{String: "https://example.com/image.jpg", Valid: true},
						CreatedAt:       pgtype.Timestamptz{Time: now, Valid: true},
						UpdatedAt:       pgtype.Timestamptz{Time: now, Valid: true},
					},
					nil,
				).Once()
			},
			expectedErr: false,
		},
		{
			name: "User creation without profile image",
			input: model.CreateUserInput{
				FirebaseUID: "firebase123456789012345678",
				Email:      "test@example.com",
				Name:       "Test User",
			},
			mockSetup: func(mockQuerier *mocks.MockQuerier) {
				expectedParams := db.CreateUserParams{
					FirebaseUid:     "firebase123456789012345678",
					Email:           "test@example.com",
					Name:            "Test User",
					ProfileImageUrl: pgtype.Text{Valid: false},
				}
				
				mockQuerier.EXPECT().CreateUser(ctx, expectedParams).Return(
					db.User{
						ID: pgtype.UUID{
							Bytes: userID,
							Valid: true,
						},
						FirebaseUid:     "firebase123456789012345678",
						Email:           "test@example.com",
						Name:            "Test User",
						ProfileImageUrl: pgtype.Text{Valid: false},
						CreatedAt:       pgtype.Timestamptz{Time: now, Valid: true},
						UpdatedAt:       pgtype.Timestamptz{Time: now, Valid: true},
					},
					nil,
				).Once()
			},
			expectedErr: false,
		},
		{
			name: "Invalid input validation error",
			input: model.CreateUserInput{
				FirebaseUID: "invalid!",
				Email:      "test@example.com",
				Name:       "Test User",
			},
			mockSetup: func(mockQuerier *mocks.MockQuerier) {
				// No mock setup needed as validation fails before DB call
			},
			expectedErr: true,
			errMessage:  "FirebaseUID must be a valid Firebase UID",
		},
		{
			name: "Database error",
			input: model.CreateUserInput{
				FirebaseUID: "firebase123456789012345678",
				Email:      "test@example.com",
				Name:       "Test User",
			},
			mockSetup: func(mockQuerier *mocks.MockQuerier) {
				mockQuerier.EXPECT().CreateUser(ctx, mock.Anything).Return(
					db.User{},
					errors.New("database error"),
				).Once()
			},
			expectedErr: true,
			errMessage:  "database error",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock
			mockQuerier := new(mocks.MockQuerier)
			
			// Setup mock expectations
			if tt.mockSetup != nil {
				tt.mockSetup(mockQuerier)
			}
			
			// Create logger
			logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelError}))
			
			// Create resolver with mocks
			resolver := &mutationResolver{
				Resolver: NewResolverWithInterfaces(mockQuerier, logger),
			}
			
			// Execute
			user, err := resolver.CreateUser(ctx, tt.input)
			
			// Assert
			if tt.expectedErr {
				assert.Error(t, err)
				if tt.errMessage != "" {
					assert.Contains(t, err.Error(), tt.errMessage)
				}
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.input.Email, user.Email)
				assert.Equal(t, tt.input.Name, user.Name)
			}
			
			// Verify mock expectations
			mockQuerier.AssertExpectations(t)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	ctx := context.Background()
	userID := uuid.New()
	now := time.Now()
	
	tests := []struct {
		name        string
		id          string
		input       model.UpdateUserInput
		mockSetup   func(*mocks.MockQuerier)
		expectedErr bool
		errMessage  string
	}{
		{
			name: "Successful update all fields",
			id:   userID.String(),
			input: model.UpdateUserInput{
				Email:           stringPtr("new@example.com"),
				Name:            stringPtr("New Name"),
				ProfileImageURL: stringPtr("https://example.com/new.jpg"),
			},
			mockSetup: func(mockQuerier *mocks.MockQuerier) {
				pgUUID := pgtype.UUID{Bytes: userID, Valid: true}
				
				// Mock GetUser to return existing user
				mockQuerier.EXPECT().GetUser(ctx, pgUUID).Return(
					db.User{
						ID:              pgUUID,
						FirebaseUid:     "firebase123456789012345678",
						Email:           "old@example.com",
						Name:            "Old Name",
						ProfileImageUrl: pgtype.Text{String: "https://example.com/old.jpg", Valid: true},
						CreatedAt:       pgtype.Timestamptz{Time: now, Valid: true},
						UpdatedAt:       pgtype.Timestamptz{Time: now, Valid: true},
					},
					nil,
				).Once()
				
				// Mock UpdateUser
				expectedParams := db.UpdateUserParams{
					ID:    pgUUID,
					Email: "new@example.com",
					Name:  "New Name",
					ProfileImageUrl: pgtype.Text{
						String: "https://example.com/new.jpg",
						Valid:  true,
					},
				}
				
				mockQuerier.EXPECT().UpdateUser(ctx, expectedParams).Return(
					db.User{
						ID:              pgUUID,
						FirebaseUid:     "firebase123456789012345678",
						Email:           "new@example.com",
						Name:            "New Name",
						ProfileImageUrl: pgtype.Text{String: "https://example.com/new.jpg", Valid: true},
						CreatedAt:       pgtype.Timestamptz{Time: now, Valid: true},
						UpdatedAt:       pgtype.Timestamptz{Time: now, Valid: true},
					},
					nil,
				).Once()
			},
			expectedErr: false,
		},
		{
			name: "Invalid UUID",
			id:   "invalid-uuid",
			input: model.UpdateUserInput{
				Name: stringPtr("New Name"),
			},
			mockSetup:   func(mockQuerier *mocks.MockQuerier) {},
			expectedErr: true,
			errMessage:  "invalid UUID",
		},
		{
			name: "Validation error",
			id:   userID.String(),
			input: model.UpdateUserInput{
				Email: stringPtr("invalid-email"),
			},
			mockSetup:   func(mockQuerier *mocks.MockQuerier) {},
			expectedErr: true,
			errMessage:  "must be a valid email",
		},
		{
			name: "User not found",
			id:   userID.String(),
			input: model.UpdateUserInput{
				Name: stringPtr("New Name"),
			},
			mockSetup: func(mockQuerier *mocks.MockQuerier) {
				pgUUID := pgtype.UUID{Bytes: userID, Valid: true}
				
				mockQuerier.EXPECT().GetUser(ctx, pgUUID).Return(
					db.User{},
					errors.New("no rows in result set"),
				).Once()
			},
			expectedErr: true,
			errMessage:  "user not found",
		},
		{
			name: "Partial update - only email",
			id:   userID.String(),
			input: model.UpdateUserInput{
				Email: stringPtr("newemail@example.com"),
			},
			mockSetup: func(mockQuerier *mocks.MockQuerier) {
				pgUUID := pgtype.UUID{Bytes: userID, Valid: true}
				
				// Mock GetUser to return existing user
				mockQuerier.EXPECT().GetUser(ctx, pgUUID).Return(
					db.User{
						ID:              pgUUID,
						FirebaseUid:     "firebase123456789012345678",
						Email:           "old@example.com",
						Name:            "Existing Name",
						ProfileImageUrl: pgtype.Text{String: "https://example.com/existing.jpg", Valid: true},
						CreatedAt:       pgtype.Timestamptz{Time: now, Valid: true},
						UpdatedAt:       pgtype.Timestamptz{Time: now, Valid: true},
					},
					nil,
				).Once()
				
				// Mock UpdateUser with only email changed
				expectedParams := db.UpdateUserParams{
					ID:              pgUUID,
					Email:           "newemail@example.com",
					Name:            "Existing Name",
					ProfileImageUrl: pgtype.Text{String: "https://example.com/existing.jpg", Valid: true},
				}
				
				mockQuerier.EXPECT().UpdateUser(ctx, expectedParams).Return(
					db.User{
						ID:              pgUUID,
						FirebaseUid:     "firebase123456789012345678",
						Email:           "newemail@example.com",
						Name:            "Existing Name",
						ProfileImageUrl: pgtype.Text{String: "https://example.com/existing.jpg", Valid: true},
						CreatedAt:       pgtype.Timestamptz{Time: now, Valid: true},
						UpdatedAt:       pgtype.Timestamptz{Time: now, Valid: true},
					},
					nil,
				).Once()
			},
			expectedErr: false,
		},
		{
			name: "Update fails at database level",
			id:   userID.String(),
			input: model.UpdateUserInput{
				Name: stringPtr("New Name"),
			},
			mockSetup: func(mockQuerier *mocks.MockQuerier) {
				pgUUID := pgtype.UUID{Bytes: userID, Valid: true}
				
				// Mock GetUser to return existing user
				mockQuerier.EXPECT().GetUser(ctx, pgUUID).Return(
					db.User{
						ID:              pgUUID,
						FirebaseUid:     "firebase123456789012345678",
						Email:           "existing@example.com",
						Name:            "Existing Name",
						ProfileImageUrl: pgtype.Text{Valid: false},
						CreatedAt:       pgtype.Timestamptz{Time: now, Valid: true},
						UpdatedAt:       pgtype.Timestamptz{Time: now, Valid: true},
					},
					nil,
				).Once()
				
				// Mock UpdateUser to fail
				mockQuerier.EXPECT().UpdateUser(ctx, mock.Anything).Return(
					db.User{},
					errors.New("constraint violation"),
				).Once()
			},
			expectedErr: true,
			errMessage:  "failed to update user",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock
			mockQuerier := new(mocks.MockQuerier)
			
			// Setup mock expectations
			if tt.mockSetup != nil {
				tt.mockSetup(mockQuerier)
			}
			
			// Create logger
			logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelError}))
			
			// Create resolver with mocks
			resolver := &mutationResolver{
				Resolver: NewResolverWithInterfaces(mockQuerier, logger),
			}
			
			// Execute
			user, err := resolver.UpdateUser(ctx, tt.id, tt.input)
			
			// Assert
			if tt.expectedErr {
				assert.Error(t, err)
				if tt.errMessage != "" {
					assert.Contains(t, err.Error(), tt.errMessage)
				}
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
			}
			
			// Verify mock expectations
			mockQuerier.AssertExpectations(t)
		})
	}
}

func TestGetUser(t *testing.T) {
	ctx := context.Background()
	userID := uuid.New()
	now := time.Now()
	
	tests := []struct {
		name        string
		id          string
		mockSetup   func(*mocks.MockQuerier)
		expectedErr bool
		errMessage  string
	}{
		{
			name: "Successful user retrieval",
			id:   userID.String(),
			mockSetup: func(mockQuerier *mocks.MockQuerier) {
				pgUUID := pgtype.UUID{
					Bytes: userID,
					Valid: true,
				}
				
				mockQuerier.EXPECT().GetUser(ctx, pgUUID).Return(
					db.User{
						ID:              pgUUID,
						FirebaseUid:     "firebase123456789012345678",
						Email:           "test@example.com",
						Name:            "Test User",
						ProfileImageUrl: pgtype.Text{String: "https://example.com/image.jpg", Valid: true},
						CreatedAt:       pgtype.Timestamptz{Time: now, Valid: true},
						UpdatedAt:       pgtype.Timestamptz{Time: now, Valid: true},
					},
					nil,
				).Once()
			},
			expectedErr: false,
		},
		{
			name: "Invalid UUID",
			id:   "invalid-uuid",
			mockSetup: func(mockQuerier *mocks.MockQuerier) {
			},
			expectedErr: true,
			errMessage:  "invalid UUID",
		},
		{
			name: "User not found",
			id:   userID.String(),
			mockSetup: func(mockQuerier *mocks.MockQuerier) {
				pgUUID := pgtype.UUID{
					Bytes: userID,
					Valid: true,
				}
				
				mockQuerier.EXPECT().GetUser(ctx, pgUUID).Return(
					db.User{},
					errors.New("no rows in result set"),
				).Once()
			},
			expectedErr: true,
			errMessage:  "user not found",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock
			mockQuerier := new(mocks.MockQuerier)
			
			// Setup mock expectations
			if tt.mockSetup != nil {
				tt.mockSetup(mockQuerier)
			}
			
			// Create logger
			logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelError}))
			
			// Create resolver with mocks
			resolver := &queryResolver{
				Resolver: NewResolverWithInterfaces(mockQuerier, logger),
			}
			
			// Execute
			user, err := resolver.User(ctx, tt.id)
			
			// Assert
			if tt.expectedErr {
				assert.Error(t, err)
				if tt.errMessage != "" {
					assert.Contains(t, err.Error(), tt.errMessage)
				}
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				// Convert UUID for comparison
				actualID := uuid.UUID(user.ID.Bytes).String()
				assert.Equal(t, userID.String(), actualID)
			}
			
			// Verify mock expectations
			mockQuerier.AssertExpectations(t)
		})
	}
}

func TestUserByFirebaseUID(t *testing.T) {
	ctx := context.Background()
	userID := uuid.New()
	firebaseUID := "firebase123456789012345678"
	now := time.Now()
	
	tests := []struct {
		name        string
		uid         string
		mockSetup   func(*mocks.MockQuerier)
		expectedErr bool
		errMessage  string
	}{
		{
			name: "Successful user retrieval by Firebase UID",
			uid:  firebaseUID,
			mockSetup: func(mockQuerier *mocks.MockQuerier) {
				mockQuerier.EXPECT().GetUserByFirebaseUID(ctx, firebaseUID).Return(
					db.User{
						ID: pgtype.UUID{
							Bytes: userID,
							Valid: true,
						},
						FirebaseUid:     firebaseUID,
						Email:           "test@example.com",
						Name:            "Test User",
						ProfileImageUrl: pgtype.Text{String: "https://example.com/image.jpg", Valid: true},
						CreatedAt:       pgtype.Timestamptz{Time: now, Valid: true},
						UpdatedAt:       pgtype.Timestamptz{Time: now, Valid: true},
					},
					nil,
				).Once()
			},
			expectedErr: false,
		},
		{
			name: "User not found",
			uid:  "nonexistent",
			mockSetup: func(mockQuerier *mocks.MockQuerier) {
				mockQuerier.EXPECT().GetUserByFirebaseUID(ctx, "nonexistent").Return(
					db.User{},
					errors.New("no rows in result set"),
				).Once()
			},
			expectedErr: true,
			errMessage:  "user not found",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock
			mockQuerier := new(mocks.MockQuerier)
			
			// Setup mock expectations
			if tt.mockSetup != nil {
				tt.mockSetup(mockQuerier)
			}
			
			// Create logger
			logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelError}))
			
			// Create resolver with mocks
			resolver := &queryResolver{
				Resolver: NewResolverWithInterfaces(mockQuerier, logger),
			}
			
			// Execute
			user, err := resolver.UserByFirebaseUID(ctx, tt.uid)
			
			// Assert
			if tt.expectedErr {
				assert.Error(t, err)
				if tt.errMessage != "" {
					assert.Contains(t, err.Error(), tt.errMessage)
				}
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, firebaseUID, user.FirebaseUid)
			}
			
			// Verify mock expectations
			mockQuerier.AssertExpectations(t)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	ctx := context.Background()
	userID := uuid.New()
	
	tests := []struct {
		name        string
		id          string
		mockSetup   func(*mocks.MockQuerier)
		expectedErr bool
		errMessage  string
	}{
		{
			name: "Successful user deletion",
			id:   userID.String(),
			mockSetup: func(mockQuerier *mocks.MockQuerier) {
				pgUUID := pgtype.UUID{
					Bytes: userID,
					Valid: true,
				}
				
				mockQuerier.EXPECT().DeleteUser(ctx, pgUUID).Return(nil).Once()
			},
			expectedErr: false,
		},
		{
			name: "Invalid UUID",
			id:   "invalid-uuid",
			mockSetup: func(mockQuerier *mocks.MockQuerier) {
			},
			expectedErr: true,
			errMessage:  "invalid UUID",
		},
		{
			name: "User not found",
			id:   userID.String(),
			mockSetup: func(mockQuerier *mocks.MockQuerier) {
				pgUUID := pgtype.UUID{
					Bytes: userID,
					Valid: true,
				}
				
				mockQuerier.EXPECT().DeleteUser(ctx, pgUUID).Return(
					errors.New("no rows affected"),
				).Once()
			},
			expectedErr: true,
			errMessage:  "failed to delete",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock
			mockQuerier := new(mocks.MockQuerier)
			
			// Setup mock expectations
			if tt.mockSetup != nil {
				tt.mockSetup(mockQuerier)
			}
			
			// Create logger
			logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelError}))
			
			// Create resolver with mocks
			resolver := &mutationResolver{
				Resolver: NewResolverWithInterfaces(mockQuerier, logger),
			}
			
			// Execute
			success, err := resolver.DeleteUser(ctx, tt.id)
			
			// Assert
			if tt.expectedErr {
				assert.Error(t, err)
				if tt.errMessage != "" {
					assert.Contains(t, err.Error(), tt.errMessage)
				}
				assert.False(t, success)
			} else {
				assert.NoError(t, err)
				assert.True(t, success)
			}
			
			// Verify mock expectations
			mockQuerier.AssertExpectations(t)
		})
	}
}

func TestHealth(t *testing.T) {
	ctx := context.Background()
	
	tests := []struct {
		name        string
		mockSetup   func(*mockDB)
		expected    string
		expectedErr bool
		errMessage  string
	}{
		{
			name: "Healthy database",
			mockSetup: func(mdb *mockDB) {
				mdb.On("Ping", ctx).Return(nil).Once()
			},
			expected:    "healthy",
			expectedErr: false,
		},
		{
			name: "Unhealthy database",
			mockSetup: func(mdb *mockDB) {
				mdb.On("Ping", ctx).Return(errors.New("connection refused")).Once()
			},
			expected:    "unhealthy",
			expectedErr: true,
			errMessage:  "database connection failed",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock database
			mockDatabase := new(mockDB)
			
			// Setup mock expectations
			if tt.mockSetup != nil {
				tt.mockSetup(mockDatabase)
			}
			
			// Create logger
			logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelError}))
			
			// Create resolver
			resolver := &queryResolver{
				Resolver: &Resolver{
					DB:     mockDatabase,
					Logger: logger,
				},
			}
			
			// Execute
			result, err := resolver.Health(ctx)
			
			// Assert
			assert.Equal(t, tt.expected, result)
			if tt.expectedErr {
				assert.Error(t, err)
				if tt.errMessage != "" {
					assert.Contains(t, err.Error(), tt.errMessage)
				}
			} else {
				assert.NoError(t, err)
			}
			
			// Verify mock expectations
			mockDatabase.AssertExpectations(t)
		})
	}
}

func TestMe(t *testing.T) {
	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelError}))
	
	// Create resolver
	resolver := &queryResolver{
		Resolver: &Resolver{
			Logger: logger,
		},
	}
	
	// Execute
	user, err := resolver.Me(ctx)
	
	// Assert - Me is not implemented yet
	assert.Nil(t, user)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "authentication not yet implemented")
}

func TestUserID(t *testing.T) {
	ctx := context.Background()
	userID := uuid.New()
	
	tests := []struct {
		name        string
		user        *db.User
		expected    string
		expectedErr bool
		errMessage  string
	}{
		{
			name: "Valid user ID",
			user: &db.User{
				ID: pgtype.UUID{
					Bytes: userID,
					Valid: true,
				},
			},
			expected:    userID.String(),
			expectedErr: false,
		},
		{
			name: "Invalid user ID",
			user: &db.User{
				ID: pgtype.UUID{
					Valid: false,
				},
			},
			expectedErr: true,
			errMessage:  "invalid user ID",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resolver := &userResolver{}
			
			// Execute
			result, err := resolver.ID(ctx, tt.user)
			
			// Assert
			if tt.expectedErr {
				assert.Error(t, err)
				if tt.errMessage != "" {
					assert.Contains(t, err.Error(), tt.errMessage)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestUserProfileImageURL(t *testing.T) {
	ctx := context.Background()
	imageURL := "https://example.com/image.jpg"
	
	tests := []struct {
		name     string
		user     *db.User
		expected *string
	}{
		{
			name: "User with profile image",
			user: &db.User{
				ProfileImageUrl: pgtype.Text{
					String: imageURL,
					Valid:  true,
				},
			},
			expected: &imageURL,
		},
		{
			name: "User without profile image",
			user: &db.User{
				ProfileImageUrl: pgtype.Text{
					Valid: false,
				},
			},
			expected: nil,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resolver := &userResolver{}
			
			// Execute
			result, err := resolver.ProfileImageURL(ctx, tt.user)
			
			// Assert
			assert.NoError(t, err)
			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *tt.expected, *result)
			}
		})
	}
}

func TestUserCreatedAt(t *testing.T) {
	ctx := context.Background()
	now := time.Now()
	
	user := &db.User{
		CreatedAt: pgtype.Timestamptz{
			Time:  now,
			Valid: true,
		},
	}
	
	resolver := &userResolver{}
	
	// Execute
	result, err := resolver.CreatedAt(ctx, user)
	
	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, now, *result)
}

func TestUserUpdatedAt(t *testing.T) {
	ctx := context.Background()
	now := time.Now()
	
	user := &db.User{
		UpdatedAt: pgtype.Timestamptz{
			Time:  now,
			Valid: true,
		},
	}
	
	resolver := &userResolver{}
	
	// Execute
	result, err := resolver.UpdatedAt(ctx, user)
	
	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, now, *result)
}

func TestResolverMethods(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelError}))
	
	resolver := &Resolver{
		Logger: logger,
	}
	
	// Test Mutation method
	mutation := resolver.Mutation()
	assert.NotNil(t, mutation)
	assert.IsType(t, &mutationResolver{}, mutation)
	
	// Test Query method
	query := resolver.Query()
	assert.NotNil(t, query)
	assert.IsType(t, &queryResolver{}, query)
	
	// Test User method
	user := resolver.User()
	assert.NotNil(t, user)
	assert.IsType(t, &userResolver{}, user)
}

func TestNewResolver(t *testing.T) {
	// Create a mock database
	db := &database.DB{
		Queries: &db.Queries{},
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelError}))
	
	// Test NewResolver
	resolver := NewResolver(db, logger)
	
	assert.NotNil(t, resolver)
	assert.NotNil(t, resolver.DB)
	assert.NotNil(t, resolver.Queries)
	assert.NotNil(t, resolver.Logger)
}

func TestDatabaseWrapper(t *testing.T) {
	// Skip this test as it requires a real database connection
	// The wrapper is tested through other tests that mock the interface
	t.Skip("Requires database connection - wrapper tested through mocked interfaces")
}

// Mock database for Health test
type mockDB struct {
	mock.Mock
}

func (m *mockDB) GetQueries() QueriesInterface {
	args := m.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(QueriesInterface)
}

func (m *mockDB) Ping(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// Helper function
func stringPtr(s string) *string {
	return &s
}