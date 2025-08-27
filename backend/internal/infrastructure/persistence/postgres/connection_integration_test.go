//go:build integration
// +build integration

package database_test

import (
	"context"
	"testing"

	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/persistence/postgres/generated"
	"github.com/fernandobandeira/djinn/backend/test/fixtures"
	"github.com/fernandobandeira/djinn/backend/test/integration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// DatabaseIntegrationTestSuite runs database integration tests with real PostgreSQL
type DatabaseIntegrationTestSuite struct {
	suite.Suite
	testDB *integration.TestDatabase
	ctx    context.Context
}

func (s *DatabaseIntegrationTestSuite) SetupSuite() {
	s.ctx = context.Background()
	s.testDB = integration.SetupTestDatabase(s.T())
}

func (s *DatabaseIntegrationTestSuite) SetupTest() {
	// Clear database before each test
	s.testDB.TruncateAllTables(s.T())
}

func (s *DatabaseIntegrationTestSuite) TestConnectionHealth() {
	health, err := s.testDB.DB.Health(s.ctx)
	s.Require().NoError(err)
	s.Equal("healthy", health.Status)
	s.NotEmpty(health.Details)
}

func (s *DatabaseIntegrationTestSuite) TestUserCRUD() {
	queries := s.testDB.DB.GetQueries()

	// Test Create
	createParams := generated.CreateUserParams{
		Email: "integration@test.com",
		Name:  "Integration Test User",
	}
	
	user, err := queries.CreateUser(s.ctx, createParams)
	s.Require().NoError(err)
	s.NotEmpty(user.ID)
	s.Equal(createParams.Email, user.Email)
	s.Equal(createParams.Name, user.Name)

	// Test Get by ID
	foundUser, err := queries.GetUser(s.ctx, user.ID)
	s.Require().NoError(err)
	s.Equal(user.ID, foundUser.ID)
	s.Equal(user.Email, foundUser.Email)

	// Test Get by Email
	foundByEmail, err := queries.GetUserByEmail(s.ctx, user.Email)
	s.Require().NoError(err)
	s.Equal(user.ID, foundByEmail.ID)

	// Test Update
	updateParams := generated.UpdateUserParams{
		ID:   user.ID,
		Name: "Updated Name",
	}
	
	updatedUser, err := queries.UpdateUser(s.ctx, updateParams)
	s.Require().NoError(err)
	s.Equal(updateParams.Name, updatedUser.Name)
	s.Equal(user.Email, updatedUser.Email) // Email should not change

	// Test Delete
	err = queries.DeleteUser(s.ctx, user.ID)
	s.Require().NoError(err)

	// Verify deletion
	_, err = queries.GetUser(s.ctx, user.ID)
	s.Error(err)
	s.Contains(err.Error(), "no rows")
}

func (s *DatabaseIntegrationTestSuite) TestListUsers() {
	queries := s.testDB.DB.GetQueries()

	// Create test users
	for _, testUser := range fixtures.TestUsers[:2] {
		params := fixtures.CreateUserParamsFactory(testUser)
		_, err := queries.CreateUser(s.ctx, params)
		s.Require().NoError(err)
	}

	// Test listing
	users, err := queries.ListUsers(s.ctx)
	s.Require().NoError(err)
	s.Len(users, 2)

	// Verify users are from fixtures
	emails := []string{users[0].Email, users[1].Email}
	s.Contains(emails, fixtures.TestUsers[0].Email)
	s.Contains(emails, fixtures.TestUsers[1].Email)
}

func (s *DatabaseIntegrationTestSuite) TestCountUsers() {
	queries := s.testDB.DB.GetQueries()

	// Initially should be 0
	count, err := queries.CountUsers(s.ctx)
	s.Require().NoError(err)
	s.Equal(int64(0), count)

	// Add users and verify count
	for i, testUser := range fixtures.TestUsers {
		params := fixtures.CreateUserParamsFactory(testUser)
		_, err := queries.CreateUser(s.ctx, params)
		s.Require().NoError(err)

		count, err = queries.CountUsers(s.ctx)
		s.Require().NoError(err)
		s.Equal(int64(i+1), count)
	}
}

func (s *DatabaseIntegrationTestSuite) TestGetUsersByIds() {
	queries := s.testDB.DB.GetQueries()

	// Create test users
	var createdIDs []generated.UUID
	for _, testUser := range fixtures.TestUsers {
		params := fixtures.CreateUserParamsFactory(testUser)
		user, err := queries.CreateUser(s.ctx, params)
		s.Require().NoError(err)
		createdIDs = append(createdIDs, user.ID)
	}

	// Get subset of users
	users, err := queries.GetUsersByIds(s.ctx, createdIDs[:2])
	s.Require().NoError(err)
	s.Len(users, 2)

	// Verify correct users returned
	for _, user := range users {
		s.Contains(createdIDs[:2], user.ID)
	}
}

func TestDatabaseIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseIntegrationTestSuite))
}