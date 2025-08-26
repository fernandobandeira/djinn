package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// UserCreateRequest represents the validation structure for user creation
type UserCreateRequest struct {
	FirebaseUID     string  `validate:"required,firebaseuid"`
	Email          string  `validate:"required,email,max=255"`
	Name           string  `validate:"required,min=1,max=255"`
	ProfileImageURL *string `validate:"omitempty,url,max=2000"`
}

// UserUpdateRequest represents the validation structure for user updates
type UserUpdateRequest struct {
	Email           *string `validate:"omitempty,email,max=255"`
	Name            *string `validate:"omitempty,min=1,max=255"`
	ProfileImageURL *string `validate:"omitempty,url,max=2000"`
}

// ValidateUserCreation validates any user creation input (GraphQL or Application layer)
func ValidateUserCreation(firebaseUID string, email string, name string, profileImageURL *string) error {
	v := GetValidator()
	
	req := UserCreateRequest{
		FirebaseUID:     firebaseUID,
		Email:          email,
		Name:           name,
		ProfileImageURL: profileImageURL,
	}
	
	if err := v.Struct(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			// Return first validation error in a user-friendly format
			for _, e := range validationErrors {
				return formatValidationError(e)
			}
		}
		return fmt.Errorf("validation failed: %w", err)
	}
	
	return nil
}

// ValidateUserUpdate validates any user update input (GraphQL or Application layer)
func ValidateUserUpdate(email *string, name *string, profileImageURL *string) error {
	v := GetValidator()
	
	req := UserUpdateRequest{
		Email:           email,
		Name:            name,
		ProfileImageURL: profileImageURL,
	}
	
	if err := v.Struct(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			// Return first validation error in a user-friendly format
			for _, e := range validationErrors {
				return formatValidationError(e)
			}
		}
		return fmt.Errorf("validation failed: %w", err)
	}
	
	return nil
}

// formatValidationError formats a validation error for user-friendly display
func formatValidationError(e validator.FieldError) error {
	field := e.Field()
	tag := e.Tag()
	
	switch tag {
	case "required":
		return fmt.Errorf("%s is required", field)
	case "email":
		return fmt.Errorf("%s must be a valid email address", field)
	case "min":
		return fmt.Errorf("%s must be at least %s characters", field, e.Param())
	case "max":
		return fmt.Errorf("%s must not exceed %s characters", field, e.Param())
	case "url":
		return fmt.Errorf("%s must be a valid URL", field)
	case "firebaseuid":
		return fmt.Errorf("%s must be a valid Firebase UID", field)
	default:
		return fmt.Errorf("%s failed validation: %s", field, tag)
	}
}