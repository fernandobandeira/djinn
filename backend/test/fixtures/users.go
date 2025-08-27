package fixtures

import (
	db "github.com/fernandobandeira/djinn/backend/internal/database/generated"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// TestUser represents a test user with common fields
type TestUser struct {
	ID              uuid.UUID
	Email           string
	Name            string
	ProfileImageURL *string
}

// TestUsers provides common test users
var TestUsers = []TestUser{
	{
		ID:              uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		Email:           "user1@test.com",
		Name:            "Test User 1",
		ProfileImageURL: stringPtr("https://example.com/user1.jpg"),
	},
	{
		ID:              uuid.MustParse("00000000-0000-0000-0000-000000000002"),
		Email:           "user2@test.com",
		Name:            "Test User 2",
		ProfileImageURL: nil,
	},
	{
		ID:              uuid.MustParse("00000000-0000-0000-0000-000000000003"),
		Email:           "admin@test.com",
		Name:            "Admin User",
		ProfileImageURL: stringPtr("https://example.com/admin.jpg"),
	},
}

// UserFactory creates a db.User from TestUser
func UserFactory(tu TestUser) db.User {
	return db.User{
		ID:    pgtype.UUID{Bytes: tu.ID, Valid: true},
		Email: tu.Email,
		Name:  tu.Name,
		ProfileImageUrl: pgtype.Text{
			String: stringPtrToString(tu.ProfileImageURL),
			Valid:  tu.ProfileImageURL != nil,
		},
	}
}

// CreateUserParamsFactory creates db.CreateUserParams from TestUser
func CreateUserParamsFactory(tu TestUser) db.CreateUserParams {
	return db.CreateUserParams{
		Email: tu.Email,
		Name:  tu.Name,
		ProfileImageUrl: pgtype.Text{
			String: stringPtrToString(tu.ProfileImageURL),
			Valid:  tu.ProfileImageURL != nil,
		},
	}
}

// Helper functions
func stringPtr(s string) *string {
	return &s
}

func stringPtrToString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

// GetTestUserByEmail returns a test user by email
func GetTestUserByEmail(email string) *TestUser {
	for _, u := range TestUsers {
		if u.Email == email {
			return &u
		}
	}
	return nil
}

// GetTestUserByID returns a test user by ID
func GetTestUserByID(id uuid.UUID) *TestUser {
	for _, u := range TestUsers {
		if u.ID == id {
			return &u
		}
	}
	return nil
}