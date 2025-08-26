package resolver

// This file contains thin GraphQL resolvers for User that delegate to application services

import (
	"context"
	"fmt"
	"time"

	"github.com/fernandobandeira/djinn/backend/internal/application/dto"
	"github.com/fernandobandeira/djinn/backend/internal/database/db"
	"github.com/fernandobandeira/djinn/backend/internal/graph/generated"
	"github.com/fernandobandeira/djinn/backend/internal/graph/model"
	"github.com/fernandobandeira/djinn/backend/internal/middleware"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// CreateUser is the resolver for the createUser field - thin layer delegating to command handler
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*db.User, error) {
	requestID := middleware.GetRequestID(ctx)
	r.Logger.Debug("GraphQL: CreateUser mutation", "request_id", requestID)

	// Create and validate DTO from GraphQL input
	appInput, err := dto.NewCreateUserInput(
		input.FirebaseUID,
		input.Email,
		input.Name,
		input.ProfileImageURL,
	)
	if err != nil {
		return nil, err
	}

	// Delegate to command handler
	userDTO, err := r.CreateUserHandler.Handle(ctx, *appInput)
	if err != nil {
		return nil, err
	}

	// Convert DTO to database model for GraphQL response
	return dtoToDBUser(userDTO), nil
}

// UpdateUser is the resolver for the updateUser field - thin layer delegating to command handler
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UpdateUserInput) (*db.User, error) {
	requestID := middleware.GetRequestID(ctx)
	r.Logger.Debug("GraphQL: UpdateUser mutation", "user_id", id, "request_id", requestID)

	// Create and validate DTO from GraphQL input
	appInput, err := dto.NewUpdateUserInput(
		input.Email,
		input.Name,
		input.ProfileImageURL,
	)
	if err != nil {
		return nil, err
	}

	// Delegate to command handler
	userDTO, err := r.UpdateUserHandler.Handle(ctx, id, *appInput)
	if err != nil {
		return nil, err
	}

	// Convert DTO to database model for GraphQL response
	return dtoToDBUser(userDTO), nil
}

// DeleteUser is the resolver for the deleteUser field - thin layer delegating to command handler
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	requestID := middleware.GetRequestID(ctx)
	r.Logger.Debug("GraphQL: DeleteUser mutation", "user_id", id, "request_id", requestID)

	// Delegate to command handler
	err := r.DeleteUserHandler.Handle(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

// User is the resolver for the user field - thin layer delegating to query handler
func (r *queryResolver) User(ctx context.Context, id string) (*db.User, error) {
	requestID := middleware.GetRequestID(ctx)
	r.Logger.Debug("GraphQL: User query", "user_id", id, "request_id", requestID)

	// Delegate to query handler
	userDTO, err := r.GetUserHandler.Handle(ctx, id)
	if err != nil {
		return nil, err
	}

	// Convert DTO to database model for GraphQL response
	return dtoToDBUser(userDTO), nil
}

// UserByFirebaseUID is the resolver for the userByFirebaseUid field - thin layer delegating to query handler
func (r *queryResolver) UserByFirebaseUID(ctx context.Context, firebaseUID string) (*db.User, error) {
	requestID := middleware.GetRequestID(ctx)
	r.Logger.Debug("GraphQL: UserByFirebaseUID query", "firebase_uid", firebaseUID, "request_id", requestID)

	// Delegate to query handler
	userDTO, err := r.GetUserByFirebaseUIDHandler.Handle(ctx, firebaseUID)
	if err != nil {
		return nil, err
	}

	// Convert DTO to database model for GraphQL response
	return dtoToDBUser(userDTO), nil
}

// Me is the resolver for the me field
func (r *queryResolver) Me(ctx context.Context) (*db.User, error) {
	// This will be implemented after authentication is added
	// For now, return an error
	return nil, fmt.Errorf("authentication not yet implemented")
}

// Field resolvers for User type

// ID is the resolver for the id field
func (r *userResolver) ID(ctx context.Context, obj *db.User) (string, error) {
	if !obj.ID.Valid {
		return "", fmt.Errorf("invalid user ID")
	}
	// Convert pgtype.UUID bytes to UUID string
	id, err := uuid.FromBytes(obj.ID.Bytes[:])
	if err != nil {
		return "", fmt.Errorf("failed to convert UUID: %w", err)
	}
	return id.String(), nil
}

// ProfileImageURL is the resolver for the profileImageUrl field
func (r *userResolver) ProfileImageURL(ctx context.Context, obj *db.User) (*string, error) {
	if obj.ProfileImageUrl.Valid {
		return &obj.ProfileImageUrl.String, nil
	}
	return nil, nil
}

// CreatedAt is the resolver for the createdAt field
func (r *userResolver) CreatedAt(ctx context.Context, obj *db.User) (*time.Time, error) {
	createdAt := obj.CreatedAt.Time
	return &createdAt, nil
}

// UpdatedAt is the resolver for the updatedAt field
func (r *userResolver) UpdatedAt(ctx context.Context, obj *db.User) (*time.Time, error) {
	updatedAt := obj.UpdatedAt.Time
	return &updatedAt, nil
}

// User returns generated.UserResolver implementation
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }

// Helper function to convert DTO to database model
// This is needed because GraphQL still expects db.User type
// In a future refactor, we could change the GraphQL schema to use DTOs directly
func dtoToDBUser(dto *dto.UserDTO) *db.User {
	userID, _ := uuid.Parse(dto.ID)
	
	user := &db.User{
		ID: pgtype.UUID{
			Bytes: userID,
			Valid: true,
		},
		FirebaseUid: dto.FirebaseUID,
		Email:       dto.Email,
		Name:        dto.Name,
		CreatedAt: pgtype.Timestamptz{
			Time:  dto.CreatedAt,
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamptz{
			Time:  dto.UpdatedAt,
			Valid: true,
		},
	}

	if dto.ProfileImageURL != nil {
		user.ProfileImageUrl = pgtype.Text{
			String: *dto.ProfileImageURL,
			Valid:  true,
		}
	}

	return user
}