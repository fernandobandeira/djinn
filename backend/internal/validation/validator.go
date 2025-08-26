package validation

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
	once     sync.Once
)

// GetValidator returns a singleton instance of the validator
func GetValidator() *validator.Validate {
	once.Do(func() {
		validate = validator.New(validator.WithRequiredStructEnabled())
		
		// Register custom validations
		registerCustomValidations(validate)
	})
	return validate
}

// registerCustomValidations registers all custom validation rules
func registerCustomValidations(v *validator.Validate) {
	// Register Firebase UID validation
	_ = v.RegisterValidation("firebaseuid", validateFirebaseUID)
}

// validateFirebaseUID validates Firebase UID format
func validateFirebaseUID(fl validator.FieldLevel) bool {
	uid := fl.Field().String()
	// Firebase UIDs are typically 28 characters but can vary
	// They contain alphanumeric characters
	if len(uid) < 20 || len(uid) > 128 {
		return false
	}
	for _, r := range uid {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')) {
			return false
		}
	}
	return true
}