//go:build integration
// +build integration

package resolver_test

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/fernandobandeira/djinn/backend/internal/graph/generated"
	"github.com/fernandobandeira/djinn/backend/internal/graph/resolver"
	"github.com/fernandobandeira/djinn/backend/test/fixtures"
	"github.com/fernandobandeira/djinn/backend/test/integration"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// GraphQLIntegrationTestSuite tests GraphQL API with real database
type GraphQLIntegrationTestSuite struct {
	suite.Suite
	testDB *integration.TestDatabase
	client *client.Client
}

func (s *GraphQLIntegrationTestSuite) SetupSuite() {
	s.testDB = integration.SetupTestDatabase(s.T())
	
	// Create GraphQL resolver with test database
	r := &resolver.Resolver{
		DB: s.testDB.DB,
	}
	
	// Create GraphQL handler and client
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: r,
	}))
	
	s.client = client.New(srv)
}

func (s *GraphQLIntegrationTestSuite) SetupTest() {
	// Clear database before each test
	s.testDB.TruncateAllTables(s.T())
}

func (s *GraphQLIntegrationTestSuite) TestHealthQuery() {
	var resp struct {
		Health struct {
			Status  string
			Message string
			Details map[string]interface{}
		}
	}

	err := s.client.Post(`
		query {
			health {
				status
				message
				details
			}
		}
	`, &resp)

	s.Require().NoError(err)
	s.Equal("healthy", resp.Health.Status)
	s.NotEmpty(resp.Health.Message)
	s.NotNil(resp.Health.Details)
}

func (s *GraphQLIntegrationTestSuite) TestUserQueries() {
	// Seed test user
	testUser := fixtures.TestUsers[0]
	queries := s.testDB.DB.GetQueries()
	params := fixtures.CreateUserParamsFactory(testUser)
	createdUser, err := queries.CreateUser(context.Background(), params)
	s.Require().NoError(err)

	// Test get user by ID
	var getUserResp struct {
		User struct {
			ID              string
			Email           string
			Name            string
			ProfileImageURL *string
			CreatedAt       string
		}
	}

	query := `
		query GetUser($id: ID!) {
			user(id: $id) {
				id
				email
				name
				profileImageURL
				createdAt
			}
		}
	`

	err = s.client.Post(query, &getUserResp, client.Var("id", createdUser.ID.String()))
	s.Require().NoError(err)
	s.Equal(testUser.Email, getUserResp.User.Email)
	s.Equal(testUser.Name, getUserResp.User.Name)

	// Test list users
	var listUsersResp struct {
		Users []struct {
			ID    string
			Email string
			Name  string
		}
	}

	err = s.client.Post(`
		query {
			users {
				id
				email
				name
			}
		}
	`, &listUsersResp)

	s.Require().NoError(err)
	s.Len(listUsersResp.Users, 1)
	s.Equal(testUser.Email, listUsersResp.Users[0].Email)
}

func (s *GraphQLIntegrationTestSuite) TestUserMutations() {
	// Test create user
	var createResp struct {
		CreateUser struct {
			ID              string
			Email           string
			Name            string
			ProfileImageURL *string
		}
	}

	mutation := `
		mutation CreateUser($input: CreateUserInput!) {
			createUser(input: $input) {
				id
				email
				name
				profileImageURL
			}
		}
	`

	input := map[string]interface{}{
		"email": "graphql@test.com",
		"name":  "GraphQL Test User",
	}

	err := s.client.Post(mutation, &createResp, client.Var("input", input))
	s.Require().NoError(err)
	s.NotEmpty(createResp.CreateUser.ID)
	s.Equal(input["email"], createResp.CreateUser.Email)
	s.Equal(input["name"], createResp.CreateUser.Name)

	// Test update user
	var updateResp struct {
		UpdateUser struct {
			ID    string
			Name  string
			Email string
		}
	}

	updateMutation := `
		mutation UpdateUser($id: ID!, $input: UpdateUserInput!) {
			updateUser(id: $id, input: $input) {
				id
				name
				email
			}
		}
	`

	updateInput := map[string]interface{}{
		"name": "Updated GraphQL User",
	}

	err = s.client.Post(updateMutation, &updateResp, 
		client.Var("id", createResp.CreateUser.ID),
		client.Var("input", updateInput))
	
	s.Require().NoError(err)
	s.Equal(createResp.CreateUser.ID, updateResp.UpdateUser.ID)
	s.Equal("Updated GraphQL User", updateResp.UpdateUser.Name)
	s.Equal(createResp.CreateUser.Email, updateResp.UpdateUser.Email)

	// Test delete user
	var deleteResp struct {
		DeleteUser bool
	}

	deleteMutation := `
		mutation DeleteUser($id: ID!) {
			deleteUser(id: $id)
		}
	`

	err = s.client.Post(deleteMutation, &deleteResp, client.Var("id", createResp.CreateUser.ID))
	s.Require().NoError(err)
	s.True(deleteResp.DeleteUser)

	// Verify user is deleted
	var verifyResp struct {
		User *struct {
			ID string
		}
	}

	err = s.client.Post(`query GetUser($id: ID!) { user(id: $id) { id } }`, 
		&verifyResp, client.Var("id", createResp.CreateUser.ID))
	
	s.Require().NoError(err)
	s.Nil(verifyResp.User)
}

func (s *GraphQLIntegrationTestSuite) TestUserValidation() {
	// Test invalid email
	var resp struct {
		CreateUser *struct {
			ID string
		}
	}

	mutation := `
		mutation CreateUser($input: CreateUserInput!) {
			createUser(input: $input) {
				id
			}
		}
	`

	invalidInput := map[string]interface{}{
		"email": "invalid-email",
		"name":  "Test User",
	}

	err := s.client.Post(mutation, &resp, client.Var("input", invalidInput))
	s.Error(err)
	s.Contains(err.Error(), "Email must be a valid email address")

	// Test empty name
	emptyNameInput := map[string]interface{}{
		"email": "valid@test.com",
		"name":  "",
	}

	err = s.client.Post(mutation, &resp, client.Var("input", emptyNameInput))
	s.Error(err)
	s.Contains(err.Error(), "Name is required")
}

func TestGraphQLIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(GraphQLIntegrationTestSuite))
}