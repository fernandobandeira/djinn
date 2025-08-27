package user

import (
	"context"
	"fmt"

	"github.com/fernandobandeira/djinn/backend/internal/domain/user"
	"github.com/google/uuid"
	"log/slog"
)

// DeleteUserHandler handles the delete user command
type DeleteUserHandler struct {
	userService user.ServiceInterface
	logger      *slog.Logger
}

// NewDeleteUserHandler creates a new delete user handler
func NewDeleteUserHandler(userService user.ServiceInterface, logger *slog.Logger) *DeleteUserHandler {
	return &DeleteUserHandler{
		userService: userService,
		logger:      logger.With(slog.String("handler", "delete_user")),
	}
}

// Handle executes the delete user command
func (h *DeleteUserHandler) Handle(ctx context.Context, id string) error {
	h.logger.Info("Deleting user", "user_id", id)

	// Parse UUID
	userID, err := uuid.Parse(id)
	if err != nil {
		h.logger.Warn("Invalid user ID", "id", id, "error", err)
		return fmt.Errorf("invalid user ID: %w", err)
	}

	// Execute domain service
	if err := h.userService.DeleteUser(ctx, userID); err != nil {
		h.logger.Error("Failed to delete user", "user_id", id, "error", err)
		return fmt.Errorf("failed to delete user: %w", err)
	}

	h.logger.Info("User deleted successfully", "user_id", id)
	return nil
}