package dto

import (
	"testing"
	"time"

	"github.com/fernandobandeira/djinn/backend/internal/domain/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCreateUserInput(t *testing.T) {
	tests := []struct {
		name            string
		firebaseUID     string
		email           string
		userName        string
		profileImageURL *string
		wantErr         bool
		errMessage      string
	}{
		{
			name:        "valid input with profile image",
			firebaseUID: "firebase123456789012345678901234567890",
			email:       "test@example.com",
			userName:    "Test User",
			profileImageURL: stringPtr("https://example.com/avatar.jpg"),
			wantErr:     false,
		},
		{
			name:        "valid input without profile image",
			firebaseUID: "firebase123456789012345678901234567890",
			email:       "test@example.com",
			userName:    "Test User",
			wantErr:     false,
		},
		{
			name:        "empty Firebase UID",
			firebaseUID: "",
			email:       "test@example.com",
			userName:    "Test User",
			wantErr:     true,
			errMessage:  "FirebaseUID is required",
		},
		{
			name:        "empty email",
			firebaseUID: "firebase123456789012345678901234567890",
			email:       "",
			userName:    "Test User",
			wantErr:     true,
			errMessage:  "email is required",
		},
		{
			name:        "empty name",
			firebaseUID: "firebase123456789012345678901234567890",
			email:       "test@example.com",
			userName:    "",
			wantErr:     true,
			errMessage:  "name is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input, err := NewCreateUserInput(tt.firebaseUID, tt.email, tt.userName, tt.profileImageURL)
			
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMessage != "" {
					assert.Contains(t, err.Error(), tt.errMessage)
				}
				assert.Nil(t, input)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, input)
				assert.Equal(t, tt.firebaseUID, input.FirebaseUID)
				assert.Equal(t, tt.email, input.Email)
				assert.Equal(t, tt.userName, input.Name)
				assert.Equal(t, tt.profileImageURL, input.ProfileImageURL)
			}
		})
	}
}

func TestNewUpdateUserInput(t *testing.T) {
	tests := []struct {
		name            string
		email           *string
		userName        *string
		profileImageURL *string
		wantErr         bool
		errMessage      string
	}{
		{
			name:            "valid update with all fields",
			email:           stringPtr("updated@example.com"),
			userName:        stringPtr("Updated User"),
			profileImageURL: stringPtr("https://example.com/new-avatar.jpg"),
			wantErr:         false,
		},
		{
			name:     "valid update with only email",
			email:    stringPtr("updated@example.com"),
			wantErr:  false,
		},
		{
			name:     "valid update with only name",
			userName: stringPtr("Updated User"),
			wantErr:  false,
		},
		{
			name:            "valid update with only profile image",
			profileImageURL: stringPtr("https://example.com/new-avatar.jpg"),
			wantErr:         false,
		},
		{
			name:    "valid update with no fields (should be allowed at DTO level)",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input, err := NewUpdateUserInput(tt.email, tt.userName, tt.profileImageURL)
			
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMessage != "" {
					assert.Contains(t, err.Error(), tt.errMessage)
				}
				assert.Nil(t, input)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, input)
				assert.Equal(t, tt.email, input.Email)
				assert.Equal(t, tt.userName, input.Name)
				assert.Equal(t, tt.profileImageURL, input.ProfileImageURL)
			}
		})
	}
}

func TestToDTO(t *testing.T) {
	userID := uuid.New()
	createdAt := time.Now()
	updatedAt := time.Now().Add(time.Hour)

	tests := []struct {
		name       string
		domainUser *user.User
		want       *UserDTO
	}{
		{
			name: "user with profile image",
			domainUser: &user.User{
				ID:              userID,
				FirebaseUID:     "firebase123456789012345678901234567890",
				Email:           "test@example.com",
				Name:            "Test User",
				ProfileImageURL: stringPtr("https://example.com/avatar.jpg"),
				CreatedAt:       createdAt,
				UpdatedAt:       updatedAt,
			},
			want: &UserDTO{
				ID:              userID.String(),
				FirebaseUID:     "firebase123456789012345678901234567890",
				Email:           "test@example.com",
				Name:            "Test User",
				ProfileImageURL: stringPtr("https://example.com/avatar.jpg"),
				CreatedAt:       createdAt,
				UpdatedAt:       updatedAt,
			},
		},
		{
			name: "user without profile image",
			domainUser: &user.User{
				ID:          userID,
				FirebaseUID: "firebase456789012345678901234567890",
				Email:       "test2@example.com",
				Name:        "Test User 2",
				CreatedAt:   createdAt,
				UpdatedAt:   updatedAt,
			},
			want: &UserDTO{
				ID:              userID.String(),
				FirebaseUID:     "firebase456789012345678901234567890",
				Email:           "test2@example.com",
				Name:            "Test User 2",
				ProfileImageURL: nil,
				CreatedAt:       createdAt,
				UpdatedAt:       updatedAt,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToDTO(tt.domainUser)
			
			assert.NotNil(t, result)
			assert.Equal(t, tt.want.ID, result.ID)
			assert.Equal(t, tt.want.FirebaseUID, result.FirebaseUID)
			assert.Equal(t, tt.want.Email, result.Email)
			assert.Equal(t, tt.want.Name, result.Name)
			assert.Equal(t, tt.want.CreatedAt.Unix(), result.CreatedAt.Unix())
			assert.Equal(t, tt.want.UpdatedAt.Unix(), result.UpdatedAt.Unix())
			
			if tt.want.ProfileImageURL != nil {
				require.NotNil(t, result.ProfileImageURL)
				assert.Equal(t, *tt.want.ProfileImageURL, *result.ProfileImageURL)
			} else {
				assert.Nil(t, result.ProfileImageURL)
			}
		})
	}
}

func TestCreateUserInput_Validation(t *testing.T) {
	// These are basic DTO-level validations
	// Full validation happens in the domain layer
	
	t.Run("struct creation", func(t *testing.T) {
		input := CreateUserInput{
			FirebaseUID:     "firebase123456789012345678901234567890",
			Email:           "test@example.com",
			Name:            "Test User",
			ProfileImageURL: stringPtr("https://example.com/avatar.jpg"),
		}
		
		assert.Equal(t, "firebase123456789012345678901234567890", input.FirebaseUID)
		assert.Equal(t, "test@example.com", input.Email)
		assert.Equal(t, "Test User", input.Name)
		assert.NotNil(t, input.ProfileImageURL)
		assert.Equal(t, "https://example.com/avatar.jpg", *input.ProfileImageURL)
	})
}

func TestUpdateUserInput_Validation(t *testing.T) {
	t.Run("struct creation", func(t *testing.T) {
		input := UpdateUserInput{
			Email:           stringPtr("updated@example.com"),
			Name:            stringPtr("Updated User"),
			ProfileImageURL: stringPtr("https://example.com/new-avatar.jpg"),
		}
		
		assert.NotNil(t, input.Email)
		assert.Equal(t, "updated@example.com", *input.Email)
		assert.NotNil(t, input.Name)
		assert.Equal(t, "Updated User", *input.Name)
		assert.NotNil(t, input.ProfileImageURL)
		assert.Equal(t, "https://example.com/new-avatar.jpg", *input.ProfileImageURL)
	})
	
	t.Run("nil values allowed", func(t *testing.T) {
		input := UpdateUserInput{}
		
		assert.Nil(t, input.Email)
		assert.Nil(t, input.Name)
		assert.Nil(t, input.ProfileImageURL)
	})
}

// Helper function
func stringPtr(s string) *string {
	return &s
}