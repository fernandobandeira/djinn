package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetValidator(t *testing.T) {
	// Test singleton pattern
	v1 := GetValidator()
	v2 := GetValidator()
	
	assert.NotNil(t, v1, "First validator instance should not be nil")
	assert.NotNil(t, v2, "Second validator instance should not be nil")
	assert.Same(t, v1, v2, "GetValidator should return the same instance")
}

func TestValidateFirebaseUID(t *testing.T) {
	tests := []struct {
		name  string
		uid   string
		valid bool
	}{
		{
			name:  "Valid Firebase UID",
			uid:   "abcdefghijklmnopqrstuvwxyz12",
			valid: true,
		},
		{
			name:  "Valid longer Firebase UID",
			uid:   "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
			valid: true,
		},
		{
			name:  "Too short UID",
			uid:   "short",
			valid: false,
		},
		{
			name:  "Too long UID",
			uid:   "a" + string(make([]byte, 128)), // 129 characters
			valid: false,
		},
		{
			name:  "Invalid characters",
			uid:   "abc-def_ghi@jkl!mnopqrstuvwx",
			valid: false,
		},
		{
			name:  "Empty UID",
			uid:   "",
			valid: false,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			type testStruct struct {
				UID string `validate:"firebaseuid"`
			}
			
			v := GetValidator()
			err := v.Struct(testStruct{UID: tt.uid})
			
			if tt.valid {
				assert.NoError(t, err, "Expected valid UID but got error")
			} else {
				assert.Error(t, err, "Expected invalid UID but got no error")
			}
		})
	}
}

func TestRegisterCustomValidations(t *testing.T) {
	// Test that custom validations are registered
	v := GetValidator()
	require.NotNil(t, v)
	
	// Test Firebase UID validation is registered
	type testStruct struct {
		UID string `validate:"firebaseuid"`
	}
	
	// Valid case
	err := v.Struct(testStruct{UID: "validFirebaseUID123456789012"})
	assert.NoError(t, err)
	
	// Invalid case
	err = v.Struct(testStruct{UID: "invalid-uid!"})
	assert.Error(t, err)
}