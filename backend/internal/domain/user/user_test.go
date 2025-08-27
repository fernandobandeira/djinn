package user

import (
	"errors"
	"testing"
	"time"

	infraErrors "github.com/fernandobandeira/djinn/backend/internal/infrastructure/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name            string
		firebaseUID     string
		email           string
		userName        string
		profileImageURL *string
		wantErr         bool
		expectedErr     error
		checkValidation bool
	}{
		{
			name:        "Valid user",
			firebaseUID: "firebase_123",
			email:       "test@example.com",
			userName:    "Test User",
			wantErr:     false,
		},
		{
			name:            "Valid user with profile image",
			firebaseUID:     "firebase_123",
			email:           "test@example.com",
			userName:        "Test User",
			profileImageURL: stringPtr("https://example.com/avatar.jpg"),
			wantErr:         false,
		},
		{
			name:            "Empty Firebase UID",
			firebaseUID:     "",
			email:           "test@example.com",
			userName:        "Test User",
			wantErr:         true,
			checkValidation: true,
		},
		{
			name:            "Invalid email",
			firebaseUID:     "firebase_123",
			email:           "not-an-email",
			userName:        "Test User",
			wantErr:         true,
			checkValidation: true,
		},
		{
			name:            "Empty name",
			firebaseUID:     "firebase_123",
			email:           "test@example.com",
			userName:        "",
			wantErr:         true,
			checkValidation: true,
		},
		{
			name:            "Invalid profile URL",
			firebaseUID:     "firebase_123",
			email:           "test@example.com",
			userName:        "Test User",
			profileImageURL: stringPtr("not-a-url"),
			wantErr:         true,
			checkValidation: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := NewUser(tt.firebaseUID, tt.email, tt.userName, tt.profileImageURL)
			
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, user)
				
				// Check that validation errors are infrastructure ValidationError
				if tt.checkValidation {
					var validationErr *infraErrors.ValidationError
					assert.True(t, errors.As(err, &validationErr))
					assert.NotEmpty(t, validationErr.Field)
					assert.NotEmpty(t, validationErr.Message)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.firebaseUID, user.FirebaseUID)
				assert.Equal(t, tt.email, user.Email)
				assert.Equal(t, tt.userName, user.Name)
				assert.Equal(t, tt.profileImageURL, user.ProfileImageURL)
				assert.NotZero(t, user.ID)
				assert.NotZero(t, user.CreatedAt)
				assert.NotZero(t, user.UpdatedAt)
			}
		})
	}
}

func TestUser_Update(t *testing.T) {
	// Create a valid user first
	user, err := NewUser("firebase_123", "test@example.com", "Test User", nil)
	assert.NoError(t, err)
	
	tests := []struct {
		name            string
		email           string
		userName        string
		profileImageURL *string
		wantErr         bool
	}{
		{
			name:     "Valid update - all fields",
			email:    "newemail@example.com",
			userName: "New Name",
			wantErr:  false,
		},
		{
			name:     "Valid update - only email",
			email:    "another@example.com",
			userName: "",
			wantErr:  false,
		},
		{
			name:     "Invalid email",
			email:    "not-an-email",
			userName: "Valid Name",
			wantErr:  true,
		},
		{
			name:            "Invalid profile URL",
			email:           "valid@example.com",
			userName:        "Valid Name",
			profileImageURL: stringPtr("invalid-url"),
			wantErr:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the user for this test
			testUser := *user
			originalUpdatedAt := testUser.UpdatedAt
			
			err := testUser.Update(tt.email, tt.userName, tt.profileImageURL)
			
			if tt.wantErr {
				assert.Error(t, err)
				// Check for validation error
				var validationErr *infraErrors.ValidationError
				if errors.As(err, &validationErr) {
					assert.NotEmpty(t, validationErr.Field)
				}
			} else {
				assert.NoError(t, err)
				if tt.email != "" {
					assert.Equal(t, tt.email, testUser.Email)
				}
				if tt.userName != "" {
					assert.Equal(t, tt.userName, testUser.Name)
				}
				assert.NotEqual(t, originalUpdatedAt, testUser.UpdatedAt)
			}
		})
	}
}

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr bool
	}{
		{
			name: "Valid user",
			user: User{
				ID:          testUUID(),
				FirebaseUID: "firebase_123",
				Email:       "test@example.com",
				Name:        "Test User",
				CreatedAt:   testTime(),
				UpdatedAt:   testTime(),
			},
			wantErr: false,
		},
		{
			name: "Missing Firebase UID",
			user: User{
				ID:        testUUID(),
				Email:     "test@example.com",
				Name:      "Test User",
				CreatedAt: testTime(),
				UpdatedAt: testTime(),
			},
			wantErr: true,
		},
		{
			name: "Invalid email",
			user: User{
				ID:          testUUID(),
				FirebaseUID: "firebase_123",
				Email:       "invalid",
				Name:        "Test User",
				CreatedAt:   testTime(),
				UpdatedAt:   testTime(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			if tt.wantErr {
				assert.Error(t, err)
				// Check that it returns our custom ValidationError
				var validationErr *infraErrors.ValidationError
				if errors.As(err, &validationErr) {
					assert.NotEmpty(t, validationErr.Field)
					assert.NotEmpty(t, validationErr.Message)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// Helper functions
func stringPtr(s string) *string {
	return &s
}

func testUUID() uuid.UUID {
	return uuid.New()
}

func testTime() time.Time {
	return time.Now()
}