package database

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConnection(t *testing.T) {
	t.Run("Invalid database URL", func(t *testing.T) {
		config := Config{
			URL:             "invalid://url",
			MaxConnections:  10,
			MinConnections:  2,
			MaxConnLifetime: time.Hour,
		}
		
		db, err := NewConnection(config)
		assert.Error(t, err)
		assert.Nil(t, db)
	})
	
	// Note: Testing actual connection requires a real database or test container
	// This is typically done in integration tests
}

func TestDB_Close(t *testing.T) {
	// Create mock database
	mockDB, _, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()
	
	// Note: Full testing requires actual pgxpool which is difficult to mock
	// This demonstrates the testing approach
}

func TestDB_GetQueries(t *testing.T) {
	// This would be tested with integration tests as it requires actual database
	// Here we can test that the method exists and returns expected type
	config := Config{
		URL:             "postgres://test:test@localhost:5432/test",
		MaxConnections:  10,
		MinConnections:  2,
		MaxConnLifetime: time.Hour,
	}
	
	// Note: This will fail without actual database, included for structure
	t.Skip("Requires database connection")
	
	db, err := NewConnection(config)
	if err != nil {
		t.Skip("Database not available")
	}
	defer db.Close()
	
	queries := db.Queries
	assert.NotNil(t, queries)
}

func TestDB_GetPool(t *testing.T) {
	// Similar to above, requires actual connection
	t.Skip("Requires database connection")
	
	config := Config{
		URL:             "postgres://test:test@localhost:5432/test",
		MaxConnections:  10,
		MinConnections:  2,
		MaxConnLifetime: time.Hour,
	}
	
	db, err := NewConnection(config)
	if err != nil {
		t.Skip("Database not available")
	}
	defer db.Close()
	
	pool := db.GetPool()
	assert.NotNil(t, pool)
}

func TestDB_Health(t *testing.T) {
	t.Skip("Requires database connection")
	
	config := Config{
		URL:             "postgres://test:test@localhost:5432/test",
		MaxConnections:  10,
		MinConnections:  2,
		MaxConnLifetime: time.Hour,
	}
	
	db, err := NewConnection(config)
	if err != nil {
		t.Skip("Database not available")
	}
	defer db.Close()
	
	ctx := context.Background()
	err = db.Health(ctx)
	assert.NoError(t, err)
}