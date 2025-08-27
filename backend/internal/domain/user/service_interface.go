package user

import (
	"context"

	"github.com/google/uuid"
)

// ServiceInterface defines the interface for user service operations
// This allows application layer to depend on abstraction rather than concrete implementation
type ServiceInterface interface {
	CreateUser(ctx context.Context, firebaseUID, email, name string, profileImageURL *string) (*User, error)
	GetUser(ctx context.Context, id uuid.UUID) (*User, error)
	GetUserByFirebaseUID(ctx context.Context, firebaseUID string) (*User, error)
	UpdateUser(ctx context.Context, id uuid.UUID, email *string, name *string, profileImageURL *string) (*User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}