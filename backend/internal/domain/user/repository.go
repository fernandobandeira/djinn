package user

import (
	"context"

	"github.com/google/uuid"
)

// Repository defines the interface for user persistence
// This is a domain interface - implementations belong in the adapter layer
type Repository interface {
	// Create persists a new user
	Create(ctx context.Context, user *User) error
	
	// GetByID retrieves a user by their ID
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	
	// GetByFirebaseUID retrieves a user by their Firebase UID
	GetByFirebaseUID(ctx context.Context, firebaseUID string) (*User, error)
	
	// Update persists changes to an existing user
	Update(ctx context.Context, user *User) error
	
	// Delete removes a user
	Delete(ctx context.Context, id uuid.UUID) error
	
	// Exists checks if a user exists by Firebase UID
	Exists(ctx context.Context, firebaseUID string) (bool, error)
}