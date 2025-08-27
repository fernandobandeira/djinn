package dto

import (
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