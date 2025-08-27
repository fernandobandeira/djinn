package user_test

import (
	"context"
	"errors"
	"log/slog"
	"testing"

	"github.com/fernandobandeira/djinn/backend/internal/domain/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository is a mock implementation of user.Repository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(ctx context.Context, u *user.User) error {
	args := m.Called(ctx, u)
	return args.Error(0)
}

func (m *MockRepository) GetByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *MockRepository) GetByFirebaseUID(ctx context.Context, firebaseUID string) (*user.User, error) {
	args := m.Called(ctx, firebaseUID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *MockRepository) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *MockRepository) List(ctx context.Context) ([]*user.User, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*user.User), args.Error(1)
}

func (m *MockRepository) Update(ctx context.Context, u *user.User) error {
	args := m.Called(ctx, u)
	return args.Error(0)
}

func (m *MockRepository) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockRepository) Exists(ctx context.Context, firebaseUID string) (bool, error) {
	args := m.Called(ctx, firebaseUID)
	return args.Bool(0), args.Error(1)
}

func TestService_CreateUser(t *testing.T) {
	tests := []struct {
		name            string
		firebaseUID     string
		email          string
		userName        string
		profileImageURL *string
		mockSetup      func(*MockRepository)
		expectedErr    bool
		errMessage     string
	}{
		{
			name:        "Successfully create user",
			firebaseUID: "firebase123456789012345678",
			email:      "test@example.com",
			userName:   "Test User",
			mockSetup: func(m *MockRepository) {
				m.On("Exists", mock.Anything, "firebase123456789012345678").Return(false, nil)
				m.On("Create", mock.Anything, mock.MatchedBy(func(u *user.User) bool {
					return u.FirebaseUID == "firebase123456789012345678" &&
						u.Email == "test@example.com" &&
						u.Name == "Test User"
				})).Return(nil)
			},
			expectedErr: false,
		},
		{
			name:            "Successfully create user with profile image",
			firebaseUID:     "firebase123456789012345678",
			email:          "test@example.com",
			userName:       "Test User",
			profileImageURL: stringPtr("https://example.com/image.jpg"),
			mockSetup: func(m *MockRepository) {
				m.On("Exists", mock.Anything, "firebase123456789012345678").Return(false, nil)
				m.On("Create", mock.Anything, mock.MatchedBy(func(u *user.User) bool {
					return u.ProfileImageURL != nil && *u.ProfileImageURL == "https://example.com/image.jpg"
				})).Return(nil)
			},
			expectedErr: false,
		},
		{
			name:        "User already exists",
			firebaseUID: "firebase123456789012345678",
			email:      "test@example.com",
			userName:   "Test User",
			mockSetup: func(m *MockRepository) {
				m.On("Exists", mock.Anything, "firebase123456789012345678").Return(true, nil)
			},
			expectedErr: true,
			errMessage:  "user already exists",
		},
		{
			name:        "Repository error on exists check",
			firebaseUID: "firebase123456789012345678",
			email:      "test@example.com",
			userName:   "Test User",
			mockSetup: func(m *MockRepository) {
				m.On("Exists", mock.Anything, "firebase123456789012345678").Return(false, errors.New("db error"))
			},
			expectedErr: true,
			errMessage:  "failed to check user existence",
		},
		{
			name:        "Repository error on create",
			firebaseUID: "firebase123456789012345678",
			email:      "test@example.com",
			userName:   "Test User",
			mockSetup: func(m *MockRepository) {
				m.On("Exists", mock.Anything, "firebase123456789012345678").Return(false, nil)
				m.On("Create", mock.Anything, mock.Anything).Return(errors.New("create failed"))
			},
			expectedErr: true,
			errMessage:  "failed to create user",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockRepo := new(MockRepository)
			if tt.mockSetup != nil {
				tt.mockSetup(mockRepo)
			}

			logger := slog.Default()
			service := user.NewService(mockRepo, logger)
			ctx := context.Background()

			// Execute
			result, err := service.CreateUser(ctx, tt.firebaseUID, tt.email, tt.userName, tt.profileImageURL)

			// Validate
			if tt.expectedErr {
				assert.Error(t, err)
				if tt.errMessage != "" {
					assert.Contains(t, err.Error(), tt.errMessage)
				}
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.firebaseUID, result.FirebaseUID)
				assert.Equal(t, tt.email, result.Email)
				assert.Equal(t, tt.userName, result.Name)
				if tt.profileImageURL != nil {
					assert.Equal(t, tt.profileImageURL, result.ProfileImageURL)
				}
			}

			// Verify mock expectations
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestService_GetUser(t *testing.T) {
	testID := uuid.New()

	tests := []struct {
		name        string
		userID      uuid.UUID
		mockSetup   func(*MockRepository)
		expectedErr bool
	}{
		{
			name:   "Successfully get user",
			userID: testID,
			mockSetup: func(m *MockRepository) {
				testUser, _ := user.NewUser("firebase123", "test@example.com", "Test User", nil)
				testUser.ID = testID
				m.On("GetByID", mock.Anything, testID).Return(testUser, nil)
			},
			expectedErr: false,
		},
		{
			name:   "User not found",
			userID: testID,
			mockSetup: func(m *MockRepository) {
				m.On("GetByID", mock.Anything, testID).Return(nil, user.ErrUserNotFound)
			},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockRepo := new(MockRepository)
			if tt.mockSetup != nil {
				tt.mockSetup(mockRepo)
			}

			logger := slog.Default()
			service := user.NewService(mockRepo, logger)
			ctx := context.Background()

			// Execute
			result, err := service.GetUser(ctx, tt.userID)

			// Validate
			if tt.expectedErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.userID, result.ID)
			}

			// Verify mock expectations
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestService_UpdateUser(t *testing.T) {
	testID := uuid.New()

	tests := []struct {
		name            string
		userID          uuid.UUID
		email          string
		userName       string
		profileImageURL *string
		mockSetup      func(*MockRepository)
		expectedErr    bool
	}{
		{
			name:     "Successfully update user",
			userID:   testID,
			email:    "updated@example.com",
			userName: "Updated User",
			mockSetup: func(m *MockRepository) {
				existingUser, _ := user.NewUser("firebase123", "old@example.com", "Old User", nil)
				existingUser.ID = testID
				m.On("GetByID", mock.Anything, testID).Return(existingUser, nil)
				m.On("Update", mock.Anything, mock.MatchedBy(func(u *user.User) bool {
					return u.Email == "updated@example.com" && u.Name == "Updated User"
				})).Return(nil)
			},
			expectedErr: false,
		},
		{
			name:     "User not found",
			userID:   testID,
			email:    "updated@example.com",
			userName: "Updated User",
			mockSetup: func(m *MockRepository) {
				m.On("GetByID", mock.Anything, testID).Return(nil, user.ErrUserNotFound)
			},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockRepo := new(MockRepository)
			if tt.mockSetup != nil {
				tt.mockSetup(mockRepo)
			}

			logger := slog.Default()
			service := user.NewService(mockRepo, logger)
			ctx := context.Background()

			// Execute
			result, err := service.UpdateUser(ctx, tt.userID, &tt.email, &tt.userName, tt.profileImageURL)

			// Validate
			if tt.expectedErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.email, result.Email)
				assert.Equal(t, tt.userName, result.Name)
			}

			// Verify mock expectations
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestService_DeleteUser(t *testing.T) {
	testID := uuid.New()

	tests := []struct {
		name        string
		userID      uuid.UUID
		mockSetup   func(*MockRepository)
		expectedErr bool
	}{
		{
			name:   "Successfully delete user",
			userID: testID,
			mockSetup: func(m *MockRepository) {
				testUser, _ := user.NewUser("firebase123", "test@example.com", "Test User", nil)
				testUser.ID = testID
				m.On("GetByID", mock.Anything, testID).Return(testUser, nil)
				m.On("Delete", mock.Anything, testID).Return(nil)
			},
			expectedErr: false,
		},
		{
			name:   "User not found",
			userID: testID,
			mockSetup: func(m *MockRepository) {
				m.On("GetByID", mock.Anything, testID).Return(nil, user.ErrUserNotFound)
			},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockRepo := new(MockRepository)
			if tt.mockSetup != nil {
				tt.mockSetup(mockRepo)
			}

			logger := slog.Default()
			service := user.NewService(mockRepo, logger)
			ctx := context.Background()

			// Execute
			err := service.DeleteUser(ctx, tt.userID)

			// Validate
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// Verify mock expectations
			mockRepo.AssertExpectations(t)
		})
	}
}

func stringPtr(s string) *string {
	return &s
}