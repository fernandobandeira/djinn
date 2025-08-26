package user

import (
	"time"

	"github.com/google/uuid"
)

// User represents the core user entity in the domain layer
type User struct {
	ID              uuid.UUID
	FirebaseUID     string
	Email           string
	Name            string
	ProfileImageURL *string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// NewUser creates a new user entity with validation
func NewUser(firebaseUID, email, name string, profileImageURL *string) (*User, error) {
	// Domain validation rules
	if firebaseUID == "" {
		return nil, ErrInvalidFirebaseUID
	}
	if email == "" {
		return nil, ErrInvalidEmail
	}
	if name == "" {
		return nil, ErrInvalidName
	}

	return &User{
		ID:              uuid.New(),
		FirebaseUID:     firebaseUID,
		Email:           email,
		Name:            name,
		ProfileImageURL: profileImageURL,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}, nil
}

// Update applies updates to the user entity
func (u *User) Update(email, name string, profileImageURL *string) error {
	if email != "" {
		u.Email = email
	}
	if name != "" {
		u.Name = name
	}
	if profileImageURL != nil {
		u.ProfileImageURL = profileImageURL
	}
	u.UpdatedAt = time.Now()
	return nil
}

// Validate ensures the user entity is in a valid state
func (u *User) Validate() error {
	if u.FirebaseUID == "" {
		return ErrInvalidFirebaseUID
	}
	if u.Email == "" {
		return ErrInvalidEmail
	}
	if u.Name == "" {
		return ErrInvalidName
	}
	return nil
}