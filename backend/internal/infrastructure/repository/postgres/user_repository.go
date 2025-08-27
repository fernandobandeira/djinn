package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/persistence/postgres/generated"
	domainUser "github.com/fernandobandeira/djinn/backend/internal/domain/user"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"log/slog"
)

// UserRepository is the PostgreSQL implementation of the user repository
type UserRepository struct {
	queries generated.Querier
	logger  *slog.Logger
}

// NewUserRepository creates a new PostgreSQL user repository
func NewUserRepository(queries generated.Querier, logger *slog.Logger) *UserRepository {
	return &UserRepository{
		queries: queries,
		logger:  logger.With(slog.String("repository", "user")),
	}
}

// Create persists a new user to the database
func (r *UserRepository) Create(ctx context.Context, user *domainUser.User) error {
	params := generated.CreateUserParams{
		FirebaseUid: user.FirebaseUID,
		Email:       user.Email,
		Name:        user.Name,
	}

	if user.ProfileImageURL != nil {
		params.ProfileImageUrl = pgtype.Text{
			String: *user.ProfileImageURL,
			Valid:  true,
		}
	}

	dbUser, err := r.queries.CreateUser(ctx, params)
	if err != nil {
		r.logger.Error("Failed to create user in database", "error", err)
		return fmt.Errorf("failed to create user: %w", err)
	}

	// Update the domain entity with the generated ID from database
	user.ID = uuid.UUID(dbUser.ID.Bytes)
	user.CreatedAt = dbUser.CreatedAt.Time
	user.UpdatedAt = dbUser.UpdatedAt.Time

	return nil
}

// GetByID retrieves a user by their ID
func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*domainUser.User, error) {
	pgUUID := pgtype.UUID{
		Bytes: id,
		Valid: true,
	}

	dbUser, err := r.queries.GetUser(ctx, pgUUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || errors.Is(err, pgx.ErrNoRows) {
			return nil, domainUser.ErrUserNotFound
		}
		r.logger.Error("Failed to get user from database", "user_id", id, "error", err)
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return r.toDomainUser(&dbUser), nil
}

// GetByFirebaseUID retrieves a user by their Firebase UID
func (r *UserRepository) GetByFirebaseUID(ctx context.Context, firebaseUID string) (*domainUser.User, error) {
	dbUser, err := r.queries.GetUserByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || errors.Is(err, pgx.ErrNoRows) {
			return nil, domainUser.ErrUserNotFound
		}
		r.logger.Error("Failed to get user by Firebase UID", "firebase_uid", firebaseUID, "error", err)
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return r.toDomainUser(&dbUser), nil
}

// Update persists changes to an existing user
func (r *UserRepository) Update(ctx context.Context, user *domainUser.User) error {
	pgUUID := pgtype.UUID{
		Bytes: user.ID,
		Valid: true,
	}

	params := generated.UpdateUserParams{
		ID:    pgUUID,
		Email: user.Email,
		Name:  user.Name,
	}

	if user.ProfileImageURL != nil {
		params.ProfileImageUrl = pgtype.Text{
			String: *user.ProfileImageURL,
			Valid:  true,
		}
	} else {
		params.ProfileImageUrl = pgtype.Text{Valid: false}
	}

	dbUser, err := r.queries.UpdateUser(ctx, params)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || errors.Is(err, pgx.ErrNoRows) {
			return domainUser.ErrUserNotFound
		}
		r.logger.Error("Failed to update user in database", "user_id", user.ID, "error", err)
		return fmt.Errorf("failed to update user: %w", err)
	}

	// Update timestamp from database
	user.UpdatedAt = dbUser.UpdatedAt.Time

	return nil
}

// Delete removes a user from the database
func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	pgUUID := pgtype.UUID{
		Bytes: id,
		Valid: true,
	}

	err := r.queries.DeleteUser(ctx, pgUUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || errors.Is(err, pgx.ErrNoRows) {
			return domainUser.ErrUserNotFound
		}
		r.logger.Error("Failed to delete user from database", "user_id", id, "error", err)
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

// Exists checks if a user exists by Firebase UID
func (r *UserRepository) Exists(ctx context.Context, firebaseUID string) (bool, error) {
	_, err := r.queries.GetUserByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		// Check for both sql.ErrNoRows and pgx.ErrNoRows
		if errors.Is(err, sql.ErrNoRows) || errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		r.logger.Error("Failed to check user existence", "firebase_uid", firebaseUID, "error", err)
		return false, fmt.Errorf("failed to check user existence: %w", err)
	}

	return true, nil
}

// toDomainUser converts a database user to a domain user
func (r *UserRepository) toDomainUser(dbUser *generated.User) *domainUser.User {
	user := &domainUser.User{
		ID:          uuid.UUID(dbUser.ID.Bytes),
		FirebaseUID: dbUser.FirebaseUid,
		Email:       dbUser.Email,
		Name:        dbUser.Name,
		CreatedAt:   dbUser.CreatedAt.Time,
		UpdatedAt:   dbUser.UpdatedAt.Time,
	}

	if dbUser.ProfileImageUrl.Valid {
		user.ProfileImageURL = &dbUser.ProfileImageUrl.String
	}

	return user
}