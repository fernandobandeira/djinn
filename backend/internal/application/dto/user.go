package dto

import (
	"fmt"
	"time"

	"github.com/fernandobandeira/djinn/backend/internal/domain/user"
	"github.com/fernandobandeira/djinn/backend/internal/validation"
	"github.com/google/uuid"
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
// This is an immutable value object that validates on construction
type CreateUserInput struct {
	FirebaseUID     string  `json:"firebaseUid"`
	Email           string  `json:"email"`
	Name            string  `json:"name"`
	ProfileImageURL *string `json:"profileImageUrl,omitempty"`
}

// NewCreateUserInput creates and validates a new CreateUserInput
func NewCreateUserInput(firebaseUID, email, name string, profileImageURL *string) (*CreateUserInput, error) {
	// Validate the input using the existing validation logic
	if err := validation.ValidateUserCreation(firebaseUID, email, name, profileImageURL); err != nil {
		return nil, fmt.Errorf("invalid user creation input: %w", err)
	}

	return &CreateUserInput{
		FirebaseUID:     firebaseUID,
		Email:           email,
		Name:            name,
		ProfileImageURL: profileImageURL,
	}, nil
}

// UpdateUserInput represents the input for updating a user
// This is an immutable value object that validates on construction
type UpdateUserInput struct {
	Email           *string `json:"email,omitempty"`
	Name            *string `json:"name,omitempty"`
	ProfileImageURL *string `json:"profileImageUrl,omitempty"`
}

// NewUpdateUserInput creates and validates a new UpdateUserInput
func NewUpdateUserInput(email, name, profileImageURL *string) (*UpdateUserInput, error) {
	// Validate the input using the existing validation logic
	if err := validation.ValidateUserUpdate(email, name, profileImageURL); err != nil {
		return nil, fmt.Errorf("invalid user update input: %w", err)
	}

	return &UpdateUserInput{
		Email:           email,
		Name:            name,
		ProfileImageURL: profileImageURL,
	}, nil
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

// FromDTO converts a DTO to a domain user
func FromDTO(dto *UserDTO) (*user.User, error) {
	id, err := uuid.Parse(dto.ID)
	if err != nil {
		return nil, err
	}

	return &user.User{
		ID:              id,
		FirebaseUID:     dto.FirebaseUID,
		Email:           dto.Email,
		Name:            dto.Name,
		ProfileImageURL: dto.ProfileImageURL,
		CreatedAt:       dto.CreatedAt,
		UpdatedAt:       dto.UpdatedAt,
	}, nil
}