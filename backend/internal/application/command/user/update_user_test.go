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
	queryUser "github.com/fernandobandeira/djinn/backend/internal/application/query/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateUserHandler_Handle(t *testing.T) {
	tests := []struct {
		name      string
		userID    string
		input     dto.UpdateUserInput
		setupMock func(*queryUser.MockUserService, uuid.UUID)
		wantErr   bool
		errMsg    string
	}{
		{
			name:   "successful update all fields",
			userID: uuid.New().String(),
			input: dto.UpdateUserInput{
				Email:           stringPtr("updated@example.com"),
				Name:            stringPtr("Updated User"),
				ProfileImageURL: stringPtr("https://example.com/new-avatar.jpg"),
			},
			setupMock: func(mockService *queryUser.MockUserService, userID uuid.UUID) {
				updatedUser := &user.User{
					ID:              userID,
					FirebaseUID:     "firebase123",
					Email:           "updated@example.com",
					Name:            "Updated User",
					ProfileImageURL: stringPtr("https://example.com/new-avatar.jpg"),
					CreatedAt:       time.Now(),
					UpdatedAt:       time.Now(),
				}
				
				mockService.On("UpdateUser", mock.Anything, userID, 
					stringPtr("updated@example.com"),
					stringPtr("Updated User"),
					stringPtr("https://example.com/new-avatar.jpg"),
				).Return(updatedUser, nil)
			},
			wantErr: false,
		},
		{
			name:   "successful partial update - email only",
			userID: uuid.New().String(),
			input: dto.UpdateUserInput{
				Email: stringPtr("partial@example.com"),
			},
			setupMock: func(mockService *queryUser.MockUserService, userID uuid.UUID) {
				updatedUser := &user.User{
					ID:          userID,
					FirebaseUID: "firebase456",
					Email:       "partial@example.com",
					Name:        "Original Name",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}
				
				mockService.On("UpdateUser", mock.Anything, userID,
					stringPtr("partial@example.com"), (*string)(nil), (*string)(nil),
				).Return(updatedUser, nil)
			},
			wantErr: false,
		},
		{
			name:   "successful partial update - name only",
			userID: uuid.New().String(),
			input: dto.UpdateUserInput{
				Name: stringPtr("New Name Only"),
			},
			setupMock: func(mockService *queryUser.MockUserService, userID uuid.UUID) {
				updatedUser := &user.User{
					ID:          userID,
					FirebaseUID: "firebase789",
					Email:       "original@example.com",
					Name:        "New Name Only",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}
				
				mockService.On("UpdateUser", mock.Anything, userID,
					(*string)(nil), stringPtr("New Name Only"), (*string)(nil),
				).Return(updatedUser, nil)
			},
			wantErr: false,
		},
		{
			name:   "invalid user ID format",
			userID: "invalid-uuid",
			input: dto.UpdateUserInput{
				Email: stringPtr("test@example.com"),
			},
			setupMock: func(mockService *queryUser.MockUserService, userID uuid.UUID) {
				// No mock setup needed for validation failure
			},
			wantErr: true,
			errMsg:  "invalid user ID",
		},
		{
			name:   "user not found",
			userID: uuid.New().String(),
			input: dto.UpdateUserInput{
				Email: stringPtr("notfound@example.com"),
			},
			setupMock: func(mockService *queryUser.MockUserService, userID uuid.UUID) {
				mockService.On("UpdateUser", mock.Anything, userID,
					stringPtr("notfound@example.com"), (*string)(nil), (*string)(nil),
				).Return(nil, user.ErrUserNotFound)
			},
			wantErr: true,
			errMsg:  "user not found",
		},
		{
			name:   "service error",
			userID: uuid.New().String(),
			input: dto.UpdateUserInput{
				Name: stringPtr("Service Error User"),
			},
			setupMock: func(mockService *queryUser.MockUserService, userID uuid.UUID) {
				mockService.On("UpdateUser", mock.Anything, userID,
					(*string)(nil), stringPtr("Service Error User"), (*string)(nil),
				).Return(nil, errors.New("service error"))
			},
			wantErr: true,
			errMsg:  "service error",
		},
		{
			name:   "context cancellation handling",
			userID: uuid.New().String(),
			input: dto.UpdateUserInput{
				Email: stringPtr("context@example.com"),
			},
			setupMock: func(mockService *queryUser.MockUserService, userID uuid.UUID) {
				mockService.On("UpdateUser", mock.Anything, userID,
					stringPtr("context@example.com"), (*string)(nil), (*string)(nil),
				).Return(nil, context.Canceled)
			},
			wantErr: true,
			errMsg:  "context canceled",
		},
		{
			name:   "boundary value testing - empty strings",
			userID: uuid.New().String(),
			input: dto.UpdateUserInput{
				Email: stringPtr(""),
				Name:  stringPtr(""),
			},
			setupMock: func(mockService *queryUser.MockUserService, userID uuid.UUID) {
				mockService.On("UpdateUser", mock.Anything, userID,
					stringPtr(""), stringPtr(""), (*string)(nil),
				).Return(nil, user.ErrInvalidEmail)
			},
			wantErr: true,
			errMsg:  "invalid email",
		},
		{
			name:   "unicode character handling",
			userID: uuid.New().String(),
			input: dto.UpdateUserInput{
				Name: stringPtr("José Müller 测试"),
			},
			setupMock: func(mockService *queryUser.MockUserService, userID uuid.UUID) {
				updatedUser := &user.User{
					ID:          userID,
					FirebaseUID: "firebase123",
					Email:       "original@example.com",
					Name:        "José Müller 测试",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}
				mockService.On("UpdateUser", mock.Anything, userID,
					(*string)(nil), stringPtr("José Müller 测试"), (*string)(nil),
				).Return(updatedUser, nil)
			},
			wantErr: false,
		},
		{
			name:   "malformed profile image URL",
			userID: uuid.New().String(),
			input: dto.UpdateUserInput{
				ProfileImageURL: stringPtr("not-a-valid-url"),
			},
			setupMock: func(mockService *queryUser.MockUserService, userID uuid.UUID) {
				mockService.On("UpdateUser", mock.Anything, userID,
					(*string)(nil), (*string)(nil), stringPtr("not-a-valid-url"),
				).Return(nil, user.ErrInvalidProfileURL)
			},
			wantErr: true,
			errMsg:  "invalid profile image URL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(queryUser.MockUserService)
			logger := slog.New(slog.NewTextHandler(io.Discard, nil))
			handler := NewUpdateUserHandler(mockService, logger)
			
			var userID uuid.UUID
			if tt.userID != "invalid-uuid" {
				userID, _ = uuid.Parse(tt.userID)
			}
			
			tt.setupMock(mockService, userID)
			
			result, err := handler.Handle(context.Background(), tt.userID, tt.input)
			
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
				// Enhanced error assertion using errors.Is for domain errors
				if tt.errMsg == "user not found" {
					assert.ErrorIs(t, err, user.ErrUserNotFound)
				} else if tt.errMsg == "context canceled" {
					assert.ErrorIs(t, err, context.Canceled)
				} else if tt.errMsg == "invalid email" {
					assert.ErrorIs(t, err, user.ErrInvalidEmail)
				} else if tt.errMsg == "invalid profile image URL" {
					assert.ErrorIs(t, err, user.ErrInvalidProfileURL)
				}
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				if tt.input.Email != nil {
					assert.Equal(t, *tt.input.Email, result.Email)
				}
				if tt.input.Name != nil {
					assert.Equal(t, *tt.input.Name, result.Name)
				}
			}
			
			mockService.AssertExpectations(t)
		})
	}
}

func TestNewUpdateUserHandler(t *testing.T) {
	mockService := new(queryUser.MockUserService)
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	handler := NewUpdateUserHandler(mockService, logger)
	
	assert.NotNil(t, handler)
	assert.Equal(t, mockService, handler.userService)
}