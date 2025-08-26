package user

import "errors"

// Domain-specific errors for user entity
var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidFirebaseUID = errors.New("invalid firebase UID")
	ErrInvalidEmail       = errors.New("invalid email address")
	ErrInvalidName        = errors.New("invalid user name")
	ErrInvalidUserID      = errors.New("invalid user ID")
)