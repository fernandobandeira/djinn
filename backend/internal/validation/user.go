package validation

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fernandobandeira/djinn/backend/internal/graph/model"
)

var (
	// Email validation regex based on W3C HTML5 specification
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	
	// Firebase UID format validation
	firebaseUIDRegex = regexp.MustCompile(`^[a-zA-Z0-9]{20,128}$`)
	
	// Name validation - allows letters, spaces, hyphens, apostrophes
	nameRegex = regexp.MustCompile(`^[\p{L}\s\-']{1,255}$`)
)

const (
	MaxNameLength  = 255
	MaxEmailLength = 255
	MaxURLLength   = 2000
	MinNameLength  = 1
)

// ValidateCreateUserInput validates user creation input
func ValidateCreateUserInput(input model.CreateUserInput) error {
	// Validate Firebase UID
	if input.FirebaseUID == "" {
		return fmt.Errorf("firebase UID is required")
	}
	if !firebaseUIDRegex.MatchString(input.FirebaseUID) {
		return fmt.Errorf("invalid firebase UID format")
	}
	
	// Validate email
	if input.Email == "" {
		return fmt.Errorf("email is required")
	}
	if len(input.Email) > MaxEmailLength {
		return fmt.Errorf("email must not exceed %d characters", MaxEmailLength)
	}
	if !emailRegex.MatchString(input.Email) {
		return fmt.Errorf("invalid email format")
	}
	
	// Validate name
	if input.Name == "" {
		return fmt.Errorf("name is required")
	}
	name := strings.TrimSpace(input.Name)
	if len(name) < MinNameLength || len(name) > MaxNameLength {
		return fmt.Errorf("name must be between %d and %d characters", MinNameLength, MaxNameLength)
	}
	if !nameRegex.MatchString(name) {
		return fmt.Errorf("name contains invalid characters")
	}
	
	// Validate profile image URL if provided
	if input.ProfileImageURL != nil && *input.ProfileImageURL != "" {
		if len(*input.ProfileImageURL) > MaxURLLength {
			return fmt.Errorf("profile image URL must not exceed %d characters", MaxURLLength)
		}
		// Basic URL validation - just check it starts with http/https
		url := strings.ToLower(*input.ProfileImageURL)
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			return fmt.Errorf("profile image URL must start with http:// or https://")
		}
	}
	
	return nil
}

// ValidateUpdateUserInput validates user update input
func ValidateUpdateUserInput(input model.UpdateUserInput) error {
	// Validate email if provided
	if input.Email != nil {
		if *input.Email == "" {
			return fmt.Errorf("email cannot be empty")
		}
		if len(*input.Email) > MaxEmailLength {
			return fmt.Errorf("email must not exceed %d characters", MaxEmailLength)
		}
		if !emailRegex.MatchString(*input.Email) {
			return fmt.Errorf("invalid email format")
		}
	}
	
	// Validate name if provided
	if input.Name != nil {
		if *input.Name == "" {
			return fmt.Errorf("name cannot be empty")
		}
		name := strings.TrimSpace(*input.Name)
		if len(name) < MinNameLength || len(name) > MaxNameLength {
			return fmt.Errorf("name must be between %d and %d characters", MinNameLength, MaxNameLength)
		}
		if !nameRegex.MatchString(name) {
			return fmt.Errorf("name contains invalid characters")
		}
	}
	
	// Validate profile image URL if provided
	if input.ProfileImageURL != nil {
		if *input.ProfileImageURL != "" {
			if len(*input.ProfileImageURL) > MaxURLLength {
				return fmt.Errorf("profile image URL must not exceed %d characters", MaxURLLength)
			}
			url := strings.ToLower(*input.ProfileImageURL)
			if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
				return fmt.Errorf("profile image URL must start with http:// or https://")
			}
		}
	}
	
	return nil
}

// SanitizeInput removes potentially dangerous characters from input
func SanitizeInput(input string) string {
	// Remove null bytes
	input = strings.ReplaceAll(input, "\x00", "")
	
	// Trim whitespace
	input = strings.TrimSpace(input)
	
	// Remove control characters
	sanitized := strings.Builder{}
	for _, r := range input {
		if r >= 32 && r != 127 { // Printable characters only
			sanitized.WriteRune(r)
		}
	}
	
	return sanitized.String()
}