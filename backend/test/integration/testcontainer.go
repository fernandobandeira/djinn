package integration

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/persistence/postgres"
	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/persistence/postgres/generated"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	postgrescontainer "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

// TestDatabase represents a test database container
type TestDatabase struct {
	Container *postgrescontainer.PostgresContainer
	DB        *postgres.DB
	URL       string
}

// SetupTestDatabase creates a PostgreSQL container for testing
func SetupTestDatabase(t *testing.T) *TestDatabase {
	ctx := context.Background()

	// Find the migrations directory
	migrationsPath, err := filepath.Abs("../../migrations")
	if err != nil {
		// Try alternative path if running from different directory
		migrationsPath, err = filepath.Abs("migrations")
		require.NoError(t, err, "Failed to find migrations directory")
	}

	// Create PostgreSQL container
	container, err := postgrescontainer.Run(ctx,
		"postgres:16-alpine",
		postgrescontainer.WithDatabase("djinn_test"),
		postgrescontainer.WithUsername("postgres"),
		postgrescontainer.WithPassword("postgres"),
		postgrescontainer.WithInitScripts(filepath.Join(migrationsPath, "*.up.sql")),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30*time.Second),
		),
	)
	require.NoError(t, err, "Failed to start PostgreSQL container")

	// Get connection string
	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err, "Failed to get connection string")

	// Connect to the database
	config := postgres.Config{
		URL:             connStr,
		MaxConnections:  10,
		MinConnections:  2,
		MaxConnLifetime: 5 * time.Minute,
	}
	db, err := postgres.NewConnection(config)
	require.NoError(t, err, "Failed to connect to test database")

	// Verify connection
	err = db.Health(ctx)
	require.NoError(t, err, "Failed to check database health")

	t.Cleanup(func() {
		db.Close()  // Close doesn't return an error
		if err := container.Terminate(ctx); err != nil {
			t.Logf("Failed to terminate container: %v", err)
		}
	})

	return &TestDatabase{
		Container: container,
		DB:        db,
		URL:       connStr,
	}
}

// TruncateAllTables clears all data from tables for test isolation
func (td *TestDatabase) TruncateAllTables(t *testing.T) {
	ctx := context.Background()
	
	// List of tables to truncate (add more as needed)
	tables := []string{"users"}
	
	for _, table := range tables {
		query := fmt.Sprintf("TRUNCATE TABLE %s CASCADE", table)
		_, err := td.DB.GetPool().Exec(ctx, query)
		require.NoError(t, err, "Failed to truncate table %s", table)
	}
}

// SeedTestData adds test data to the database
func (td *TestDatabase) SeedTestData(t *testing.T) {
	ctx := context.Background()
	queries := td.DB.Queries
	
	// Seed users from fixtures
	// This can be expanded with actual fixture data
	_, err := queries.CreateUser(ctx, generated.CreateUserParams{
		Email: "test@example.com",
		Name:  "Test User",
	})
	require.NoError(t, err, "Failed to seed test user")
}