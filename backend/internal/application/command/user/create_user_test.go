package user

import (
	"context"
	"errors"
	"log/slog"
	"testing"
	"time"

	"github.com/fernandobandeira/djinn/backend/internal/application/dto"
	"github.com/fernandobandeira/djinn/backend/internal/domain/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// MockUserService implements a mock for the user service
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(ctx context.Context, firebaseUID, email, name string, profileImageURL *string) (*user.User, error) {
	args := m.Called(ctx, firebaseUID, email, name, profileImageURL)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *MockUserService) GetUser(ctx context.Context, id uuid.UUID) (*user.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *MockUserService) GetUserByFirebaseUID(ctx context.Context, firebaseUID string) (*user.User, error) {
	args := m.Called(ctx, firebaseUID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *MockUserService) UpdateUser(ctx context.Context, id uuid.UUID, email, name string, profileImageURL *string) (*user.User, error) {
	args := m.Called(ctx, id, email, name, profileImageURL)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *MockUserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreateUserHandler_Handle(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(nil, nil))
	
	tests := []struct {
		name      string
		input     dto.CreateUserInput
		setupMock func(*MockUserService)
		want      *dto.UserDTO
		wantErr   bool
		errMsg    string
	}{
		{
			name: "successful user creation",
			input: dto.CreateUserInput{
				FirebaseUID:     "firebase123456789012345678901234567890",
				Email:           "test@example.com",
				Name:            "Test User",
				ProfileImageURL: stringPtr("https://example.com/avatar.jpg"),
			},
			setupMock: func(mockService *MockUserService) {
				expectedUser := &user.User{
					ID:              uuid.New(),
					FirebaseUID:     "firebase123456789012345678901234567890",
					Email:           "test@example.com",
					Name:            "Test User",
					ProfileImageURL: stringPtr("https://example.com/avatar.jpg"),
					CreatedAt:       time.Now(),
					UpdatedAt:       time.Now(),
				}
				
				mockService.On("CreateUser", 
					mock.Anything, 
					"firebase123456789012345678901234567890",
					"test@example.com",
					"Test User",
					stringPtr("https://example.com/avatar.jpg"),
				).Return(expectedUser, nil)
			},
			wantErr: false,
		},
		{
			name: "successful user creation without profile image",
			input: dto.CreateUserInput{
				FirebaseUID: "firebase456789012345678901234567890",
				Email:       "test2@example.com",
				Name:        "Test User 2",
			},
			setupMock: func(mockService *MockUserService) {
				expectedUser := &user.User{
					ID:          uuid.New(),
					FirebaseUID: "firebase456789012345678901234567890",
					Email:       "test2@example.com",
					Name:        "Test User 2",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}
				
				mockService.On("CreateUser",
					mock.Anything,
					"firebase456789012345678901234567890",
					"test2@example.com",
					"Test User 2",
					(*string)(nil),
				).Return(expectedUser, nil)
			},
			wantErr: false,
		},
		{
			name: "domain service error",
			input: dto.CreateUserInput{
				FirebaseUID: "firebase789012345678901234567890",
				Email:       "test3@example.com",
				Name:        "Test User 3",
			},
			setupMock: func(mockService *MockUserService) {
				mockService.On("CreateUser",
					mock.Anything,
					"firebase789012345678901234567890",
					"test3@example.com",
					"Test User 3",
					(*string)(nil),
				).Return(nil, errors.New("domain error"))
			},
			wantErr: true,
			errMsg:  "failed to create user",
		},
		{
			name: "validation error from domain",
			input: dto.CreateUserInput{
				FirebaseUID: "invalid",
				Email:       "invalid-email",
				Name:        "",
			},
			setupMock: func(mockService *MockUserService) {
				mockService.On("CreateUser",
					mock.Anything,
					"invalid",
					"invalid-email",
					"",
					(*string)(nil),
				).Return(nil, user.ErrInvalidFirebaseUID)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &MockUserService{}
			handler := NewCreateUserHandler(mockService, logger)
			
			tt.setupMock(mockService)
			
			result, err := handler.Handle(context.Background(), tt.input)
			
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.input.FirebaseUID, result.FirebaseUID)
				assert.Equal(t, tt.input.Email, result.Email)
				assert.Equal(t, tt.input.Name, result.Name)
				assert.NotEmpty(t, result.ID)
				assert.NotZero(t, result.CreatedAt)
				assert.NotZero(t, result.UpdatedAt)
				
				if tt.input.ProfileImageURL != nil {
					require.NotNil(t, result.ProfileImageURL)
					assert.Equal(t, *tt.input.ProfileImageURL, *result.ProfileImageURL)
				}
			}
			
			mockService.AssertExpectations(t)
		})
	}
}

func TestNewCreateUserHandler(t *testing.T) {
	mockService := &MockUserService{}
	logger := slog.New(slog.NewTextHandler(nil, nil))
	
	handler := NewCreateUserHandler(mockService, logger)
	
	assert.NotNil(t, handler)
	assert.Equal(t, mockService, handler.userService)
	assert.NotNil(t, handler.logger)
}

// Helper function
func stringPtr(s string) *string {
	return &s
}