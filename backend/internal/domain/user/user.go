package user

import (
	"fmt"
	"time"

	infraErrors "github.com/fernandobandeira/djinn/backend/internal/infrastructure/errors"
	"github.com/fernandobandeira/djinn/backend/internal/validation"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// User represents the core user entity in the domain layer
// The struct tags define the validation rules that ensure the entity is always in a valid state
type User struct {
	ID              uuid.UUID `validate:"required"`
	FirebaseUID     string    `validate:"required,firebaseuid"`
	Email           string    `validate:"required,email,max=255"`
	Name            string    `validate:"required,min=1,max=255"`
	ProfileImageURL *string   `validate:"omitempty,url,max=2000"`
	CreatedAt       time.Time `validate:"required"`
	UpdatedAt       time.Time `validate:"required"`
}

// NewUser creates a new user entity with validation
func NewUser(firebaseUID, email, name string, profileImageURL *string) (*User, error) {
	user := &User{
		ID:              uuid.New(),
		FirebaseUID:     firebaseUID,
		Email:           email,
		Name:            name,
		ProfileImageURL: profileImageURL,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	// Validate the entity using the domain rules defined in struct tags
	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

// Update applies updates to the user entity with validation
func (u *User) Update(email, name string, profileImageURL *string) error {
	// Apply changes
	if email != "" {
		u.Email = email
	}
	if name != "" {
		u.Name = name
	}
	if profileImageURL != nil {
		u.ProfileImageURL = profileImageURL
	}
	u.UpdatedAt = time.Now()
	
	// Validate the entity is still in a valid state after updates
	if err := u.Validate(); err != nil {
		return fmt.Errorf("invalid update: %w", err)
	}
	
	return nil
}

// Validate ensures the user entity is in a valid state using the domain rules
func (u *User) Validate() error {
	v := validation.GetValidator()
	
	if err := v.Struct(u); err != nil {
		// Convert validator errors to infrastructure validation errors
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrors {
				field := e.Field()
				
				// Return field-specific errors wrapped with validator's message
				var baseErr error
				switch field {
				case "FirebaseUID":
					baseErr = ErrInvalidFirebaseUID
				case "Email":
					baseErr = ErrInvalidEmail
				case "Name":
					baseErr = ErrInvalidName
				case "ProfileImageURL":
					baseErr = ErrInvalidProfileURL
				default:
					// For unexpected fields, create a generic validation error
					return &infraErrors.ValidationError{
						Field:   field,
						Message: e.Error(), // Use validator's error message
						Code:    fmt.Sprintf("USER_INVALID_%s", field),
					}
				}
				
				// Wrap the base error with the validator's descriptive message
				return fmt.Errorf("%s: %w", e.Error(), baseErr)
			}
		}
		return fmt.Errorf("validation failed: %w", err)
	}
	
	return nil
}