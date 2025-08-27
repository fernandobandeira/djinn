package dto

import (
	"fmt"
	"time"

	"github.com/fernandobandeira/djinn/backend/internal/domain/user"
)

// UserDTO is a data transfer object for user data
type UserDTO struct {
	ID              string     `json:"id"`
	FirebaseUID     string     `json:"firebaseUid"`
	Email           string     `json:"email"`
	Name            string     `json:"name"`
	ProfileImageURL *string    `json:"profileImageUrl,omitempty"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`
}

// CreateUserInput represents the input for creating a user
// Simple data carrier - validation happens in the domain layer
type CreateUserInput struct {
	FirebaseUID     string  `json:"firebaseUid"`
	Email           string  `json:"email"`
	Name            string  `json:"name"`
	ProfileImageURL *string `json:"profileImageUrl,omitempty"`
}

// UpdateUserInput represents the input for updating a user
// Simple data carrier - validation happens in the domain layer
type UpdateUserInput struct {
	Email           *string `json:"email,omitempty"`
	Name            *string `json:"name,omitempty"`
	ProfileImageURL *string `json:"profileImageUrl,omitempty"`
}

// NewCreateUserInput creates and validates a CreateUserInput
func NewCreateUserInput(firebaseUID, email, name string, profileImageURL *string) (*CreateUserInput, error) {
	input := &CreateUserInput{
		FirebaseUID:     firebaseUID,
		Email:           email,
		Name:            name,
		ProfileImageURL: profileImageURL,
	}

	// Basic validation - full validation happens in domain layer
	if firebaseUID == "" {
		return nil, fmt.Errorf("FirebaseUID is required")
	}
	if email == "" {
		return nil, fmt.Errorf("email is required")
	}
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}

	// Basic validation - detailed validation happens in domain layer
	// Just basic checks here to provide fast feedback

	return input, nil
}

// NewUpdateUserInput creates and validates an UpdateUserInput
func NewUpdateUserInput(email, name *string, profileImageURL *string) (*UpdateUserInput, error) {
	input := &UpdateUserInput{
		Email:           email,
		Name:            name,
		ProfileImageURL: profileImageURL,
	}

	// Basic validation - detailed validation happens in domain layer

	return input, nil
}

// ToDTO converts a domain user to a DTO
func ToDTO(u *user.User) *UserDTO {
	return &UserDTO{
		ID:              u.ID.String(),
		FirebaseUID:     u.FirebaseUID,
		Email:           u.Email,
		Name:            u.Name,
		ProfileImageURL: u.ProfileImageURL,
		CreatedAt:       u.CreatedAt,
		UpdatedAt:       u.UpdatedAt,
	}
}