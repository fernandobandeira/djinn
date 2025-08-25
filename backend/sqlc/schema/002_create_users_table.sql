-- Create users table for Djinn personal finance application
-- Migration: 002_create_users_table
-- Created: 2025-08-25
-- Dependencies: 001_setup_uuid_extension

-- Users table with UUIDv7 primary keys for optimal performance
CREATE TABLE users (
    -- Primary key using UUIDv7 for timestamp-ordered performance
    id UUID PRIMARY KEY DEFAULT uuid_generate_v7(),
    
    -- Firebase Auth integration
    firebase_uid VARCHAR(128) NOT NULL UNIQUE,
    
    -- User profile information
    email VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    profile_image_url TEXT,
    
    -- Audit timestamps
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    -- Constraints
    CONSTRAINT users_email_check CHECK (email ~ '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
    CONSTRAINT users_name_check CHECK (LENGTH(TRIM(name)) > 0)
);

-- Performance indexes
CREATE INDEX idx_users_firebase_uid ON users(firebase_uid);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_created_at ON users(created_at);

-- Updated timestamp trigger function
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Apply updated_at trigger to users table
CREATE TRIGGER users_updated_at_trigger
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();