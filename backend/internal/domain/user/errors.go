package user

import (
	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/errors"
)

// Domain-specific error codes
const (
	CodeInvalidFirebaseUID = "USER_INVALID_FIREBASE_UID"
	CodeInvalidEmail       = "USER_INVALID_EMAIL"
	CodeInvalidName        = "USER_INVALID_NAME"
	CodeInvalidProfileURL  = "USER_INVALID_PROFILE_URL"
)

// Domain-specific errors using infrastructure error types
var (
	// Business logic errors
	ErrUserNotFound      = &errors.NotFoundError{Resource: "user"}
	ErrUserAlreadyExists = &errors.ConflictError{Resource: "user", Message: "user already exists"}
	
	// Validation errors - one per field, not per validation type
	ErrInvalidFirebaseUID = &errors.ValidationError{
		Field:   "FirebaseUID",
		Message: "invalid Firebase UID",
		Code:    CodeInvalidFirebaseUID,
	}
	ErrInvalidEmail = &errors.ValidationError{
		Field:   "Email",
		Message: "invalid email address",
		Code:    CodeInvalidEmail,
	}
	ErrInvalidName = &errors.ValidationError{
		Field:   "Name",
		Message: "invalid user name",
		Code:    CodeInvalidName,
	}
	ErrInvalidProfileURL = &errors.ValidationError{
		Field:   "ProfileImageURL",
		Message: "invalid profile image URL",
		Code:    CodeInvalidProfileURL,
	}
)

