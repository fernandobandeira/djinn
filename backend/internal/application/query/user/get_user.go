package user

import (
	"context"
	"fmt"

	"github.com/fernandobandeira/djinn/backend/internal/application/dto"
	"github.com/fernandobandeira/djinn/backend/internal/domain/user"
	"github.com/google/uuid"
	"log/slog"
)

// GetUserHandler handles the get user query
type GetUserHandler struct {
	userService user.ServiceInterface
	logger      *slog.Logger
}

// NewGetUserHandler creates a new get user handler
func NewGetUserHandler(userService user.ServiceInterface, logger *slog.Logger) *GetUserHandler {
	return &GetUserHandler{
		userService: userService,
		logger:      logger.With(slog.String("handler", "get_user")),
	}
}

// Handle executes the get user query
func (h *GetUserHandler) Handle(ctx context.Context, id string) (*dto.UserDTO, error) {
	h.logger.Debug("Getting user", "user_id", id)

	// Parse UUID
	userID, err := uuid.Parse(id)
	if err != nil {
		h.logger.Warn("Invalid user ID", "id", id, "error", err)
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	// Execute domain service
	domainUser, err := h.userService.GetUser(ctx, userID)
	if err != nil {
		h.logger.Error("Failed to get user", "user_id", id, "error", err)
		return nil, err
	}

	// Convert to DTO
	return dto.ToDTO(domainUser), nil
}

// GetUserByFirebaseUIDHandler handles the get user by Firebase UID query
type GetUserByFirebaseUIDHandler struct {
	userService user.ServiceInterface
	logger      *slog.Logger
}

// NewGetUserByFirebaseUIDHandler creates a new handler
func NewGetUserByFirebaseUIDHandler(userService user.ServiceInterface, logger *slog.Logger) *GetUserByFirebaseUIDHandler {
	return &GetUserByFirebaseUIDHandler{
		userService: userService,
		logger:      logger.With(slog.String("handler", "get_user_by_firebase_uid")),
	}
}

// Handle executes the get user by Firebase UID query
func (h *GetUserByFirebaseUIDHandler) Handle(ctx context.Context, firebaseUID string) (*dto.UserDTO, error) {
	h.logger.Debug("Getting user by Firebase UID", "firebase_uid", firebaseUID)

	// Execute domain service
	domainUser, err := h.userService.GetUserByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		h.logger.Error("Failed to get user by Firebase UID", 
			"firebase_uid", firebaseUID, 
			"error", err)
		return nil, err
	}

	// Convert to DTO
	return dto.ToDTO(domainUser), nil
}