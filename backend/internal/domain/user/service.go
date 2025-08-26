package user

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

// Service encapsulates user domain logic
type Service struct {
	repo   Repository
	logger *slog.Logger
}

// NewService creates a new user domain service
func NewService(repo Repository, logger *slog.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger.With(slog.String("service", "user")),
	}
}

// CreateUser handles the business logic for creating a new user
func (s *Service) CreateUser(ctx context.Context, firebaseUID, email, name string, profileImageURL *string) (*User, error) {
	s.logger.Debug("Creating user",
		"firebase_uid", firebaseUID,
		"email", email,
	)

	// Check if user already exists
	exists, err := s.repo.Exists(ctx, firebaseUID)
	if err != nil {
		return nil, fmt.Errorf("failed to check user existence: %w", err)
	}
	if exists {
		return nil, ErrUserAlreadyExists
	}

	// Create domain entity
	user, err := NewUser(firebaseUID, email, name, profileImageURL)
	if err != nil {
		s.logger.Warn("Invalid user data", "error", err)
		return nil, err
	}

	// Persist to repository
	if err := s.repo.Create(ctx, user); err != nil {
		s.logger.Error("Failed to create user", "error", err)
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	s.logger.Info("User created successfully", "user_id", user.ID)
	return user, nil
}

// GetUser retrieves a user by ID
func (s *Service) GetUser(ctx context.Context, id uuid.UUID) (*User, error) {
	s.logger.Debug("Getting user", "user_id", id)

	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("Failed to get user", "user_id", id, "error", err)
		return nil, err
	}

	return user, nil
}

// GetUserByFirebaseUID retrieves a user by Firebase UID
func (s *Service) GetUserByFirebaseUID(ctx context.Context, firebaseUID string) (*User, error) {
	s.logger.Debug("Getting user by Firebase UID", "firebase_uid", firebaseUID)

	if firebaseUID == "" {
		return nil, ErrInvalidFirebaseUID
	}

	user, err := s.repo.GetByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		s.logger.Error("Failed to get user by Firebase UID", 
			"firebase_uid", firebaseUID, 
			"error", err)
		return nil, err
	}

	return user, nil
}

// UpdateUser updates an existing user
func (s *Service) UpdateUser(ctx context.Context, id uuid.UUID, email, name string, profileImageURL *string) (*User, error) {
	s.logger.Debug("Updating user", "user_id", id)

	// Retrieve existing user
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("Failed to find user for update", "user_id", id, "error", err)
		return nil, err
	}

	// Apply updates to domain entity
	if err := user.Update(email, name, profileImageURL); err != nil {
		s.logger.Warn("Invalid update data", "error", err)
		return nil, err
	}

	// Persist changes
	if err := s.repo.Update(ctx, user); err != nil {
		s.logger.Error("Failed to update user", "user_id", id, "error", err)
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	s.logger.Info("User updated successfully", "user_id", id)
	return user, nil
}

// DeleteUser removes a user
func (s *Service) DeleteUser(ctx context.Context, id uuid.UUID) error {
	s.logger.Debug("Deleting user", "user_id", id)

	// Check if user exists
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("Failed to find user for deletion", "user_id", id, "error", err)
		return err
	}

	// Delete user
	if err := s.repo.Delete(ctx, id); err != nil {
		s.logger.Error("Failed to delete user", "user_id", id, "error", err)
		return fmt.Errorf("failed to delete user: %w", err)
	}

	s.logger.Info("User deleted successfully", "user_id", id)
	return nil
}