package resolver

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHealth(t *testing.T) {
	ctx := context.Background()
	
	tests := []struct {
		name        string
		mockSetup   func(*mockPingableDB)
		expected    string
		expectedErr bool
		errMessage  string
	}{
		{
			name: "Healthy database",
			mockSetup: func(mdb *mockPingableDB) {
				mdb.On("Ping", ctx).Return(nil).Once()
			},
			expected:    "healthy",
			expectedErr: false,
		},
		{
			name: "Unhealthy database",
			mockSetup: func(mdb *mockPingableDB) {
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
			mockDB := new(mockPingableDB)
			
			// Setup mock expectations
			if tt.mockSetup != nil {
				tt.mockSetup(mockDB)
			}
			
			// Create logger
			logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelError}))
			
			// Create resolver with type assertion
			// In tests, we use the mock; in production, we use real database.DB
			resolver := &queryResolver{
				Resolver: &Resolver{
					DB:     mockDB,
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
			mockDB.AssertExpectations(t)
		})
	}
}

// Mock database for Health test
type mockPingableDB struct {
	mock.Mock
}

func (m *mockPingableDB) Ping(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}