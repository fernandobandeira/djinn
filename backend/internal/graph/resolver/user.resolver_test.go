package resolver

import (
	"testing"

	"github.com/fernandobandeira/djinn/backend/internal/application/dto"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser_DTOValidation(t *testing.T) {
	// These tests verify that the GraphQL resolver DTOs are created properly
	// Note: Full validation happens in the domain layer, not in DTOs
	
	tests := []struct {
		name        string
		input       dto.CreateUserInput
		expectedErr bool
		errMessage  string
	}{
		{
			name: "Valid user creation with profile image",
			input: dto.CreateUserInput{
				FirebaseUID:     "firebase123456789012345678",
				Email:          "test@example.com",
				Name:           "Test User",
				ProfileImageURL: stringPtr("https://example.com/image.jpg"),
			},
			expectedErr: false,
		},
		{
			name: "Valid user creation without profile image",
			input: dto.CreateUserInput{
				FirebaseUID: "firebase123456789012345678",
				Email:      "test@example.com",
				Name:       "Test User",
			},
			expectedErr: false,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test DTO creation and validation
			_, err := dto.NewCreateUserInput(
				tt.input.FirebaseUID,
				tt.input.Email,
				tt.input.Name,
				tt.input.ProfileImageURL,
			)
			
			// Assert validation behavior
			if tt.expectedErr {
				assert.Error(t, err)
				if err != nil && tt.errMessage != "" {
					assert.Contains(t, err.Error(), tt.errMessage)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateUser_DTOValidation(t *testing.T) {
	// Test that update DTOs properly validate through their constructor
	
	tests := []struct {
		name        string
		input       dto.UpdateUserInput
		expectedErr bool
		errMessage  string
	}{
		{
			name: "Valid update all fields",
			input: dto.UpdateUserInput{
				Email:           stringPtr("new@example.com"),
				Name:            stringPtr("New Name"),
				ProfileImageURL: stringPtr("https://example.com/new.jpg"),
			},
			expectedErr: false,
		},
		{
			name: "Valid update email only",
			input: dto.UpdateUserInput{
				Email: stringPtr("new@example.com"),
			},
			expectedErr: false,
		},
		{
			name: "Valid update name only",
			input: dto.UpdateUserInput{
				Name: stringPtr("New Name"),
			},
			expectedErr: false,
		},
		{
			name: "Empty update (all nil)",
			input: dto.UpdateUserInput{},
			expectedErr: false,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test DTO creation and validation
			_, err := dto.NewUpdateUserInput(
				tt.input.Email,
				tt.input.Name,
				tt.input.ProfileImageURL,
			)
			
			// Assert validation behavior
			if tt.expectedErr {
				assert.Error(t, err)
				if err != nil && tt.errMessage != "" {
					assert.Contains(t, err.Error(), tt.errMessage)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// Helper function
func stringPtr(s string) *string {
	return &s
}