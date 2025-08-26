package user

import (
	"context"
	"fmt"

	"github.com/fernandobandeira/djinn/backend/internal/application/dto"
	"github.com/fernandobandeira/djinn/backend/internal/domain/user"
	"log/slog"
)

// CreateUserHandler handles the create user command
type CreateUserHandler struct {
	userService *user.Service
	logger      *slog.Logger
}

// NewCreateUserHandler creates a new create user handler
func NewCreateUserHandler(userService *user.Service, logger *slog.Logger) *CreateUserHandler {
	return &CreateUserHandler{
		userService: userService,
		logger:      logger.With(slog.String("handler", "create_user")),
	}
}

// Handle executes the create user command
// Input validation is guaranteed by the DTO constructor
func (h *CreateUserHandler) Handle(ctx context.Context, input dto.CreateUserInput) (*dto.UserDTO, error) {
	h.logger.Info("Creating user",
		"firebase_uid", input.FirebaseUID,
		"email", input.Email,
	)

	// Input is already validated by DTO constructor
	// Execute domain service
	domainUser, err := h.userService.CreateUser(
		ctx,
		input.FirebaseUID,
		input.Email,
		input.Name,
		input.ProfileImageURL,
	)
	if err != nil {
		h.logger.Error("Failed to create user", "error", err)
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Convert to DTO
	result := dto.ToDTO(domainUser)
	
	h.logger.Info("User created successfully", "user_id", result.ID)
	return result, nil
}

