package user

import (
	"context"
	"fmt"

	"github.com/fernandobandeira/djinn/backend/internal/application/dto"
	"github.com/fernandobandeira/djinn/backend/internal/domain/user"
	"github.com/google/uuid"
	"log/slog"
)

// UpdateUserHandler handles the update user command
type UpdateUserHandler struct {
	userService *user.Service
	logger      *slog.Logger
}

// NewUpdateUserHandler creates a new update user handler
func NewUpdateUserHandler(userService *user.Service, logger *slog.Logger) *UpdateUserHandler {
	return &UpdateUserHandler{
		userService: userService,
		logger:      logger.With(slog.String("handler", "update_user")),
	}
}

// Handle executes the update user command
// Input validation is guaranteed by the DTO constructor
func (h *UpdateUserHandler) Handle(ctx context.Context, id string, input dto.UpdateUserInput) (*dto.UserDTO, error) {
	h.logger.Info("Updating user", "user_id", id)

	// Parse UUID
	userID, err := uuid.Parse(id)
	if err != nil {
		h.logger.Warn("Invalid user ID", "id", id, "error", err)
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	// Input is already validated by DTO constructor
	// Prepare update values
	var email, name string
	if input.Email != nil {
		email = *input.Email
	}
	if input.Name != nil {
		name = *input.Name
	}

	// Execute domain service
	domainUser, err := h.userService.UpdateUser(
		ctx,
		userID,
		email,
		name,
		input.ProfileImageURL,
	)
	if err != nil {
		h.logger.Error("Failed to update user", "user_id", id, "error", err)
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	// Convert to DTO
	result := dto.ToDTO(domainUser)
	
	h.logger.Info("User updated successfully", "user_id", result.ID)
	return result, nil
}

