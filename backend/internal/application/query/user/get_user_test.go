package user

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/fernandobandeira/djinn/backend/internal/application/dto"
	"github.com/fernandobandeira/djinn/backend/internal/domain/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserHandler_Handle(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	userID := uuid.New()
	
	tests := []struct {
		name      string
		userID    string
		setupMock func(*MockUserService)
		want      *dto.UserDTO
		wantErr   bool
		errMsg    string
	}{
		{
			name:   "successful user retrieval",
			userID: userID.String(),
			setupMock: func(mockService *MockUserService) {
				expectedUser := &user.User{
					ID:              userID,
					FirebaseUID:     "firebase123456789012345678901234567890",
					Email:           "test@example.com",
					Name:            "Test User",
					ProfileImageURL: stringPtr("https://example.com/avatar.jpg"),
					CreatedAt:       time.Now(),
					UpdatedAt:       time.Now(),
				}
				
				mockService.On("GetUser", mock.Anything, userID).Return(expectedUser, nil)
			},
			wantErr: false,
		},
		{
			name:   "invalid user ID format",
			userID: "invalid-uuid",
			setupMock: func(mockService *MockUserService) {
				// No mock setup needed as parsing will fail first
			},
			wantErr: true,
			errMsg:  "invalid user ID",
		},
		{
			name:   "user not found",
			userID: userID.String(),
			setupMock: func(mockService *MockUserService) {
				mockService.On("GetUser", mock.Anything, userID).Return(nil, user.ErrUserNotFound)
			},
			wantErr: true,
		},
		{
			name:   "service error",
			userID: userID.String(),
			setupMock: func(mockService *MockUserService) {
				mockService.On("GetUser", mock.Anything, userID).Return(nil, errors.New("service error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &MockUserService{}
			handler := NewGetUserHandler(mockService, logger)
			
			tt.setupMock(mockService)
			
			result, err := handler.Handle(context.Background(), tt.userID)
			
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, userID.String(), result.ID)
				assert.Equal(t, "firebase123456789012345678901234567890", result.FirebaseUID)
				assert.Equal(t, "test@example.com", result.Email)
				assert.Equal(t, "Test User", result.Name)
				assert.NotNil(t, result.ProfileImageURL)
				assert.Equal(t, "https://example.com/avatar.jpg", *result.ProfileImageURL)
			}
			
			mockService.AssertExpectations(t)
		})
	}
}

func TestGetUserByFirebaseUIDHandler_Handle(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	userID := uuid.New()
	firebaseUID := "firebase123456789012345678901234567890"
	
	tests := []struct {
		name        string
		firebaseUID string
		setupMock   func(*MockUserService)
		want        *dto.UserDTO
		wantErr     bool
		errMsg      string
	}{
		{
			name:        "successful user retrieval by Firebase UID",
			firebaseUID: firebaseUID,
			setupMock: func(mockService *MockUserService) {
				expectedUser := &user.User{
					ID:          userID,
					FirebaseUID: firebaseUID,
					Email:       "test@example.com",
					Name:        "Test User",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}
				
				mockService.On("GetUserByFirebaseUID", mock.Anything, firebaseUID).Return(expectedUser, nil)
			},
			wantErr: false,
		},
		{
			name:        "user not found by Firebase UID",
			firebaseUID: "nonexistent",
			setupMock: func(mockService *MockUserService) {
				mockService.On("GetUserByFirebaseUID", mock.Anything, "nonexistent").Return(nil, user.ErrUserNotFound)
			},
			wantErr: true,
		},
		{
			name:        "service error",
			firebaseUID: firebaseUID,
			setupMock: func(mockService *MockUserService) {
				mockService.On("GetUserByFirebaseUID", mock.Anything, firebaseUID).Return(nil, errors.New("service error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &MockUserService{}
			handler := NewGetUserByFirebaseUIDHandler(mockService, logger)
			
			tt.setupMock(mockService)
			
			result, err := handler.Handle(context.Background(), tt.firebaseUID)
			
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, userID.String(), result.ID)
				assert.Equal(t, firebaseUID, result.FirebaseUID)
				assert.Equal(t, "test@example.com", result.Email)
				assert.Equal(t, "Test User", result.Name)
			}
			
			mockService.AssertExpectations(t)
		})
	}
}

func TestNewGetUserHandler(t *testing.T) {
	mockService := &MockUserService{}
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	
	handler := NewGetUserHandler(mockService, logger)
	
	assert.NotNil(t, handler)
	assert.Equal(t, mockService, handler.userService)
	assert.NotNil(t, handler.logger)
}

func TestNewGetUserByFirebaseUIDHandler(t *testing.T) {
	mockService := &MockUserService{}
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	
	handler := NewGetUserByFirebaseUIDHandler(mockService, logger)
	
	assert.NotNil(t, handler)
	assert.Equal(t, mockService, handler.userService)
	assert.NotNil(t, handler.logger)
}

// Helper function
func stringPtr(s string) *string {
	return &s
}