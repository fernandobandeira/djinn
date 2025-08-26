package validation

import (
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestValidateUserCreation(t *testing.T) {
	tests := []struct {
		name            string
		firebaseUID     string
		email           string
		userName        string
		profileImageURL *string
		wantErr         bool
		errMsg          string
	}{
		{
			name:            "Valid user input",
			firebaseUID:     "validFirebaseUID123456789012",
			email:           "test@example.com",
			userName:        "Test User",
			profileImageURL: stringPtr("https://example.com/image.jpg"),
			wantErr:         false,
		},
		{
			name:            "Valid user without profile image",
			firebaseUID:     "validFirebaseUID123456789012",
			email:           "test@example.com",
			userName:        "Test User",
			profileImageURL: nil,
			wantErr:         false,
		},
		{
			name:        "Missing Firebase UID",
			firebaseUID: "",
			email:       "test@example.com",
			userName:    "Test User",
			wantErr:     true,
			errMsg:      "FirebaseUID is required",
		},
		{
			name:        "Invalid Firebase UID",
			firebaseUID: "short",
			email:       "test@example.com",
			userName:    "Test User",
			wantErr:     true,
			errMsg:      "FirebaseUID must be a valid Firebase UID",
		},
		{
			name:        "Missing email",
			firebaseUID: "validFirebaseUID123456789012",
			email:       "",
			userName:    "Test User",
			wantErr:     true,
			errMsg:      "Email is required",
		},
		{
			name:        "Invalid email",
			firebaseUID: "validFirebaseUID123456789012",
			email:       "invalid-email",
			userName:    "Test User",
			wantErr:     true,
			errMsg:      "Email must be a valid email address",
		},
		{
			name:        "Email too long",
			firebaseUID: "validFirebaseUID123456789012",
			email:       strings.Repeat("a", 247) + "@test.com",
			userName:    "Test User",
			wantErr:     true,
			errMsg:      "Email must not exceed 255 characters",
		},
		{
			name:        "Missing name",
			firebaseUID: "validFirebaseUID123456789012",
			email:       "test@example.com",
			userName:    "",
			wantErr:     true,
			errMsg:      "Name is required",
		},
		{
			name:        "Empty name",
			firebaseUID: "validFirebaseUID123456789012",
			email:       "test@example.com",
			userName:    "",
			wantErr:     true,
			errMsg:      "Name is required",
		},
		{
			name:        "Name too long",
			firebaseUID: "validFirebaseUID123456789012",
			email:       "test@example.com",
			userName:    strings.Repeat("a", 256),
			wantErr:     true,
			errMsg:      "Name must not exceed 255 characters",
		},
		{
			name:            "Invalid profile image URL",
			firebaseUID:     "validFirebaseUID123456789012",
			email:           "test@example.com",
			userName:        "Test User",
			profileImageURL: stringPtr("not-a-url"),
			wantErr:         true,
			errMsg:          "ProfileImageURL must be a valid URL",
		},
		{
			name:            "Profile image URL too long",
			firebaseUID:     "validFirebaseUID123456789012",
			email:           "test@example.com",
			userName:        "Test User",
			profileImageURL: stringPtr("https://example.com/" + strings.Repeat("a", 1981)),
			wantErr:         true,
			errMsg:          "ProfileImageURL must not exceed 2000 characters",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUserCreation(
				tt.firebaseUID,
				tt.email,
				tt.userName,
				tt.profileImageURL,
			)
			
			if tt.wantErr {
				assert.Error(t, err, "Expected validation error")
				if tt.errMsg != "" {
					assert.EqualError(t, err, tt.errMsg)
				}
			} else {
				assert.NoError(t, err, "Expected no validation error")
			}
		})
	}
}

func TestValidateUserUpdate(t *testing.T) {
	tests := []struct {
		name            string
		email           *string
		userName        *string
		profileImageURL *string
		wantErr         bool
		errMsg          string
	}{
		{
			name:            "Valid update all fields",
			email:           stringPtr("newemail@example.com"),
			userName:        stringPtr("New Name"),
			profileImageURL: stringPtr("https://example.com/new.jpg"),
			wantErr:         false,
		},
		{
			name:     "Valid update email only",
			email:    stringPtr("newemail@example.com"),
			wantErr:  false,
		},
		{
			name:     "Valid update name only",
			userName: stringPtr("New Name"),
			wantErr:  false,
		},
		{
			name:    "Empty update (all nil)",
			wantErr: false,
		},
		{
			name:    "Invalid email",
			email:   stringPtr("invalid-email"),
			wantErr: true,
			errMsg:  "Email must be a valid email address",
		},
		{
			name:    "Email too long",
			email:   stringPtr(strings.Repeat("a", 247) + "@test.com"),
			wantErr: true,
			errMsg:  "Email must not exceed 255 characters",
		},
		{
			name:     "Empty name (not allowed)",
			userName: stringPtr(""),
			wantErr:  true,
			errMsg:   "Name must be at least 1 characters",
		},
		{
			name:     "Name too long",
			userName: stringPtr(strings.Repeat("a", 256)),
			wantErr:  true,
			errMsg:   "Name must not exceed 255 characters",
		},
		{
			name:            "Invalid profile image URL",
			profileImageURL: stringPtr("not-a-url"),
			wantErr:         true,
			errMsg:          "ProfileImageURL must be a valid URL",
		},
		{
			name:            "Profile image URL too long",
			profileImageURL: stringPtr("https://example.com/" + strings.Repeat("a", 1981)),
			wantErr:         true,
			errMsg:          "ProfileImageURL must not exceed 2000 characters",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUserUpdate(
				tt.email,
				tt.userName,
				tt.profileImageURL,
			)
			
			if tt.wantErr {
				assert.Error(t, err, "Expected validation error")
				if tt.errMsg != "" {
					assert.EqualError(t, err, tt.errMsg)
				}
			} else {
				assert.NoError(t, err, "Expected no validation error")
			}
		})
	}
}

func TestFormatValidationError(t *testing.T) {
	v := GetValidator()
	
	// Create a struct that will generate specific validation errors
	type TestStruct struct {
		Required  string `validate:"required"`
		Email     string `validate:"email"`
		MinLength string `validate:"min=5"`
		MaxLength string `validate:"max=10"`
		URL       string `validate:"url"`
		Firebase  string `validate:"firebaseuid"`
	}
	
	tests := []struct {
		name          string
		input         TestStruct
		expectedError string
	}{
		{
			name:          "Required field error",
			input:         TestStruct{Email: "test@test.com"},
			expectedError: "Required is required",
		},
		{
			name:          "Email format error",
			input:         TestStruct{Required: "test", Email: "invalid"},
			expectedError: "Email must be a valid email address",
		},
		{
			name:          "Min length error",
			input:         TestStruct{Required: "test", Email: "test@test.com", MinLength: "abc"},
			expectedError: "MinLength must be at least 5 characters",
		},
		{
			name:          "Max length error",
			input:         TestStruct{Required: "test", Email: "test@test.com", MinLength: "abcdef", MaxLength: "12345678901"},
			expectedError: "MaxLength must not exceed 10 characters",
		},
		{
			name:          "URL format error",
			input:         TestStruct{Required: "test", Email: "test@test.com", MinLength: "abcdef", MaxLength: "1234567890", URL: "not-a-url"},
			expectedError: "URL must be a valid URL",
		},
		{
			name:          "Firebase UID error",
			input:         TestStruct{Required: "test", Email: "test@test.com", MinLength: "abcdef", MaxLength: "1234567890", URL: "https://test.com", Firebase: "invalid"},
			expectedError: "Firebase must be a valid Firebase UID",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v.Struct(tt.input)
			assert.Error(t, err)
			
			if validationErrors, ok := err.(validator.ValidationErrors); ok {
				for _, e := range validationErrors {
					formattedErr := formatValidationError(e)
					assert.EqualError(t, formattedErr, tt.expectedError)
					break // Only test the first error
				}
			}
		})
	}
}

func TestGetValidator(t *testing.T) {
	v1 := GetValidator()
	v2 := GetValidator()
	
	// Should return the same instance (singleton)
	assert.Equal(t, v1, v2, "GetValidator should return a singleton instance")
	assert.NotNil(t, v1, "Validator should not be nil")
}

func TestValidateFirebaseUID(t *testing.T) {
	v := GetValidator()
	
	type TestStruct struct {
		UID string `validate:"firebaseuid"`
	}
	
	tests := []struct {
		name    string
		uid     string
		wantErr bool
	}{
		{
			name:    "Valid Firebase UID",
			uid:     "abcdefghijklmnopqrstuvwx",
			wantErr: false,
		},
		{
			name:    "Valid longer Firebase UID",
			uid:     "abcdefghijklmnopqrstuvwxyz1234567890ABCDEF",
			wantErr: false,
		},
		{
			name:    "Too short UID",
			uid:     "short",
			wantErr: true,
		},
		{
			name:    "Too long UID",
			uid:     strings.Repeat("a", 129),
			wantErr: true,
		},
		{
			name:    "Invalid characters",
			uid:     "abcdefghijklmnopqrstuvwx!@#",
			wantErr: true,
		},
		{
			name:    "Empty UID",
			uid:     "",
			wantErr: true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := TestStruct{UID: tt.uid}
			err := v.Struct(s)
			
			if tt.wantErr {
				assert.Error(t, err, "Expected validation error for UID: %s", tt.uid)
			} else {
				assert.NoError(t, err, "Expected no validation error for UID: %s", tt.uid)
			}
		})
	}
}

func TestRegisterCustomValidations(t *testing.T) {
	// This test ensures custom validations are registered
	v := validator.New(validator.WithRequiredStructEnabled())
	registerCustomValidations(v)
	
	// Test that the firebaseuid validation is registered
	type TestStruct struct {
		UID string `validate:"firebaseuid"`
	}
	
	// Valid UID should pass
	validTest := TestStruct{UID: "abcdefghijklmnopqrstuvwx"}
	err := v.Struct(validTest)
	assert.NoError(t, err, "Valid Firebase UID should not produce an error")
	
	// Invalid UID should fail
	invalidTest := TestStruct{UID: "invalid!"}
	err = v.Struct(invalidTest)
	assert.Error(t, err, "Invalid Firebase UID should produce an error")
}

// Helper function
func stringPtr(s string) *string {
	return &s
}