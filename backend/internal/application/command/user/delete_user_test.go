package user

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"

	"github.com/fernandobandeira/djinn/backend/internal/domain/user"
	queryUser "github.com/fernandobandeira/djinn/backend/internal/application/query/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteUserHandler_Handle(t *testing.T) {
	tests := []struct {
		name      string
		userID    string
		setupMock func(*queryUser.MockUserService, uuid.UUID)
		wantErr   bool
		errMsg    string
	}{
		{
			name:   "successful deletion",
			userID: uuid.New().String(),
			setupMock: func(mockService *queryUser.MockUserService, userID uuid.UUID) {
				mockService.On("DeleteUser", mock.Anything, userID).Return(nil)
			},
			wantErr: false,
		},
		{
			name:   "invalid user ID format",
			userID: "invalid-uuid",
			setupMock: func(mockService *queryUser.MockUserService, userID uuid.UUID) {
				// No mock setup needed for validation failure
			},
			wantErr: true,
			errMsg:  "invalid user ID",
		},
		{
			name:   "user not found",
			userID: uuid.New().String(),
			setupMock: func(mockService *queryUser.MockUserService, userID uuid.UUID) {
				mockService.On("DeleteUser", mock.Anything, userID).Return(user.ErrUserNotFound)
			},
			wantErr: true,
			errMsg:  "user not found",
		},
		{
			name:   "service error",
			userID: uuid.New().String(),
			setupMock: func(mockService *queryUser.MockUserService, userID uuid.UUID) {
				mockService.On("DeleteUser", mock.Anything, userID).Return(errors.New("database connection failed"))
			},
			wantErr: true,
			errMsg:  "database connection failed",
		},
		{
			name:   "context cancellation during deletion",
			userID: uuid.New().String(),
			setupMock: func(mockService *queryUser.MockUserService, userID uuid.UUID) {
				mockService.On("DeleteUser", mock.Anything, userID).Return(context.Canceled)
			},
			wantErr: true,
			errMsg:  "context canceled",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(queryUser.MockUserService)
			logger := slog.New(slog.NewTextHandler(io.Discard, nil))
			handler := NewDeleteUserHandler(mockService, logger)
			
			var userID uuid.UUID
			if tt.userID != "invalid-uuid" {
				userID, _ = uuid.Parse(tt.userID)
			}
			
			tt.setupMock(mockService, userID)
			
			err := handler.Handle(context.Background(), tt.userID)
			
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
				}
			} else {
				assert.NoError(t, err)
			}
			
			mockService.AssertExpectations(t)
		})
	}
}

func TestNewDeleteUserHandler(t *testing.T) {
	mockService := new(queryUser.MockUserService)
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	handler := NewDeleteUserHandler(mockService, logger)
	
	assert.NotNil(t, handler)
	assert.Equal(t, mockService, handler.userService)
}