package validation

import (
	"testing"

	"github.com/fernandobandeira/djinn/backend/internal/graph/model"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestValidateCreateUserInput(t *testing.T) {
	tests := []struct {
		name    string
		input   model.CreateUserInput
		wantErr bool
		errMsg  string
	}{
		{
			name: "Valid user input",
			input: model.CreateUserInput{
				FirebaseUID:     "validFirebaseUID123456789012",
				Email:          "user@example.com",
				Name:           "John Doe",
				ProfileImageURL: stringPtr("https://example.com/image.jpg"),
			},
			wantErr: false,
		},
		{
			name: "Valid user without profile image",
			input: model.CreateUserInput{
				FirebaseUID: "validFirebaseUID123456789012",
				Email:      "user@example.com",
				Name:       "John Doe",
			},
			wantErr: false,
		},
		{
			name: "Missing Firebase UID",
			input: model.CreateUserInput{
				Email: "user@example.com",
				Name:  "John Doe",
			},
			wantErr: true,
			errMsg:  "FirebaseUID is required",
		},
		{
			name: "Invalid Firebase UID",
			input: model.CreateUserInput{
				FirebaseUID: "invalid-uid!",
				Email:      "user@example.com",
				Name:       "John Doe",
			},
			wantErr: true,
			errMsg:  "FirebaseUID must be a valid Firebase UID",
		},
		{
			name: "Missing email",
			input: model.CreateUserInput{
				FirebaseUID: "validFirebaseUID123456789012",
				Name:       "John Doe",
			},
			wantErr: true,
			errMsg:  "Email is required",
		},
		{
			name: "Invalid email",
			input: model.CreateUserInput{
				FirebaseUID: "validFirebaseUID123456789012",
				Email:      "invalid-email",
				Name:       "John Doe",
			},
			wantErr: true,
			errMsg:  "Email must be a valid email address",
		},
		{
			name: "Email too long",
			input: model.CreateUserInput{
				FirebaseUID: "validFirebaseUID123456789012",
				Email:      "test" + string(make([]byte, 245, 245)) + "@example.com", // 260+ characters total
				Name:       "John Doe",
			},
			wantErr: true,
			errMsg:  "Email must be a valid email address", // Validator checks format before length
		},
		{
			name: "Missing name",
			input: model.CreateUserInput{
				FirebaseUID: "validFirebaseUID123456789012",
				Email:      "user@example.com",
			},
			wantErr: true,
			errMsg:  "Name is required",
		},
		{
			name: "Empty name",
			input: model.CreateUserInput{
				FirebaseUID: "validFirebaseUID123456789012",
				Email:      "user@example.com",
				Name:       "",
			},
			wantErr: true,
			errMsg:  "Name is required",
		},
		{
			name: "Name too long",
			input: model.CreateUserInput{
				FirebaseUID: "validFirebaseUID123456789012",
				Email:      "user@example.com",
				Name:       string(make([]byte, 256)),
			},
			wantErr: true,
			errMsg:  "Name must not exceed 255 characters",
		},
		{
			name: "Invalid profile image URL",
			input: model.CreateUserInput{
				FirebaseUID:     "validFirebaseUID123456789012",
				Email:          "user@example.com",
				Name:           "John Doe",
				ProfileImageURL: stringPtr("not-a-url"),
			},
			wantErr: true,
			errMsg:  "ProfileImageURL must be a valid URL",
		},
		{
			name: "Profile image URL too long",
			input: model.CreateUserInput{
				FirebaseUID:     "validFirebaseUID123456789012",
				Email:          "user@example.com",
				Name:           "John Doe",
				ProfileImageURL: stringPtr("https://example.com/" + string(make([]byte, 1980, 1980))),
			},
			wantErr: true,
			errMsg:  "ProfileImageURL must be a valid URL", // Validator checks format before length
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCreateUserInput(tt.input)
			
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

func TestValidateUpdateUserInput(t *testing.T) {
	tests := []struct {
		name    string
		input   model.UpdateUserInput
		wantErr bool
		errMsg  string
	}{
		{
			name: "Valid update all fields",
			input: model.UpdateUserInput{
				Email:           stringPtr("newemail@example.com"),
				Name:            stringPtr("New Name"),
				ProfileImageURL: stringPtr("https://example.com/newimage.jpg"),
			},
			wantErr: false,
		},
		{
			name: "Valid update email only",
			input: model.UpdateUserInput{
				Email: stringPtr("newemail@example.com"),
			},
			wantErr: false,
		},
		{
			name: "Valid update name only",
			input: model.UpdateUserInput{
				Name: stringPtr("New Name"),
			},
			wantErr: false,
		},
		{
			name: "Empty update (all nil)",
			input: model.UpdateUserInput{},
			wantErr: false,
		},
		{
			name: "Invalid email",
			input: model.UpdateUserInput{
				Email: stringPtr("invalid-email"),
			},
			wantErr: true,
			errMsg:  "Email must be a valid email address",
		},
		{
			name: "Email too long",
			input: model.UpdateUserInput{
				Email: stringPtr("test" + string(make([]byte, 245, 245)) + "@example.com"),
			},
			wantErr: true,
			errMsg:  "Email must be a valid email address", // Validator checks format before length
		},
		{
			name: "Empty name (not allowed)",
			input: model.UpdateUserInput{
				Name: stringPtr(""),
			},
			wantErr: true,
			errMsg:  "Name must be at least 1 characters",
		},
		{
			name: "Name too long",
			input: model.UpdateUserInput{
				Name: stringPtr(string(make([]byte, 256))),
			},
			wantErr: true,
			errMsg:  "Name must not exceed 255 characters",
		},
		{
			name: "Invalid profile image URL",
			input: model.UpdateUserInput{
				ProfileImageURL: stringPtr("not-a-url"),
			},
			wantErr: true,
			errMsg:  "ProfileImageURL must be a valid URL",
		},
		{
			name: "Profile image URL too long",
			input: model.UpdateUserInput{
				ProfileImageURL: stringPtr("https://example.com/" + string(make([]byte, 1980, 1980))),
			},
			wantErr: true,
			errMsg:  "ProfileImageURL must be a valid URL", // Validator checks format before length
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUpdateUserInput(tt.input)
			
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
	
	// Test required field error
	type RequiredTest struct {
		Field string `validate:"required"`
	}
	err := v.Struct(RequiredTest{})
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			formatted := formatValidationError(e)
			assert.EqualError(t, formatted, "Field is required")
			break
		}
	}
	
	// Test email validation error
	type EmailTest struct {
		Email string `validate:"email"`
	}
	err = v.Struct(EmailTest{Email: "invalid"})
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			formatted := formatValidationError(e)
			assert.EqualError(t, formatted, "Email must be a valid email address")
			break
		}
	}
	
	// Test min length error
	type MinTest struct {
		Name string `validate:"min=3"`
	}
	err = v.Struct(MinTest{Name: "ab"})
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			formatted := formatValidationError(e)
			assert.EqualError(t, formatted, "Name must be at least 3 characters")
			break
		}
	}
	
	// Test max length error
	type MaxTest struct {
		Name string `validate:"max=5"`
	}
	err = v.Struct(MaxTest{Name: "toolong"})
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			formatted := formatValidationError(e)
			assert.EqualError(t, formatted, "Name must not exceed 5 characters")
			break
		}
	}
	
	// Test URL validation error
	type URLTest struct {
		URL string `validate:"url"`
	}
	err = v.Struct(URLTest{URL: "not-a-url"})
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			formatted := formatValidationError(e)
			assert.EqualError(t, formatted, "URL must be a valid URL")
			break
		}
	}
}

// Helper function
func stringPtr(s string) *string {
	return &s
}