# Atlas migration configuration for Djinn backend
# https://atlasgo.io/guides/migration-tools

env "local" {
  # URL to the database to migrate
  url = "postgres://djinn_user:djinn_password@localhost:5432/djinn?sslmode=disable"
  
  # Migration directory
  migration {
    dir = "file://migrations"
  }
}

env "test" {
  # Test database configuration
  url = "postgres://djinn_user:djinn_password@localhost:5432/djinn_test?sslmode=disable"
  
  migration {
    dir = "file://migrations"
  }
}

env "production" {
  # Production database configuration (placeholder)
  url = getenv("PRODUCTION_DATABASE_URL")
  
  migration {
    dir = "file://migrations"
  }
  
  # Production-specific settings
  diff {
    # Skip destructive changes in production
    skip {
      drop_schema = true
      drop_table = true
    }
  }
}