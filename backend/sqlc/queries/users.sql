-- User CRUD operations for Djinn personal finance application
-- Generated queries will be type-safe with pgx/v5

-- name: GetUser :one
-- Get a user by their ID
SELECT id, firebase_uid, email, name, profile_image_url, created_at, updated_at
FROM users
WHERE id = $1;

-- name: GetUserByFirebaseUID :one
-- Get a user by their Firebase UID for authentication
SELECT id, firebase_uid, email, name, profile_image_url, created_at, updated_at
FROM users
WHERE firebase_uid = $1;

-- name: GetUserByEmail :one
-- Get a user by their email address
SELECT id, firebase_uid, email, name, profile_image_url, created_at, updated_at
FROM users
WHERE email = $1;

-- name: ListUsers :many
-- List users with pagination support
SELECT id, firebase_uid, email, name, profile_image_url, created_at, updated_at
FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CountUsers :one
-- Count total number of users for pagination
SELECT COUNT(*) FROM users;

-- name: CreateUser :one
-- Create a new user account
INSERT INTO users (firebase_uid, email, name, profile_image_url)
VALUES ($1, $2, $3, $4)
RETURNING id, firebase_uid, email, name, profile_image_url, created_at, updated_at;

-- name: UpdateUser :one
-- Update user profile information
UPDATE users
SET 
    email = COALESCE($2, email),
    name = COALESCE($3, name),
    profile_image_url = COALESCE($4, profile_image_url)
WHERE id = $1
RETURNING id, firebase_uid, email, name, profile_image_url, created_at, updated_at;

-- name: UpdateUserByFirebaseUID :one
-- Update user profile by Firebase UID (for auth flows)
UPDATE users
SET 
    email = COALESCE($2, email),
    name = COALESCE($3, name),
    profile_image_url = COALESCE($4, profile_image_url)
WHERE firebase_uid = $1
RETURNING id, firebase_uid, email, name, profile_image_url, created_at, updated_at;

-- name: DeleteUser :exec
-- Delete a user account (soft delete could be implemented later)
DELETE FROM users
WHERE id = $1;

-- name: DeleteUserByFirebaseUID :exec
-- Delete a user by Firebase UID
DELETE FROM users
WHERE firebase_uid = $1;

-- name: UserExists :one
-- Check if a user exists by ID
SELECT EXISTS(SELECT 1 FROM users WHERE id = $1);

-- name: UserExistsByFirebaseUID :one
-- Check if a user exists by Firebase UID
SELECT EXISTS(SELECT 1 FROM users WHERE firebase_uid = $1);

-- name: UserExistsByEmail :one
-- Check if a user exists by email
SELECT EXISTS(SELECT 1 FROM users WHERE email = $1);

-- name: GetRecentUsers :many
-- Get recently created users (for admin/analytics)
SELECT id, firebase_uid, email, name, profile_image_url, created_at, updated_at
FROM users
WHERE created_at >= $1
ORDER BY created_at DESC
LIMIT $2;