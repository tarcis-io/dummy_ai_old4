package config

import (
	"testing"
)

// TestNew tests the creation of a new Config instance.
func TestNew(t *testing.T) {
	testCases := []struct {
		name              string
		envValues         map[string]string
		wantServerAddress string
		wantError         bool
	}{
		{
			name:              "should create a new Config instance with default values",
			envValues:         map[string]string{},
			wantServerAddress: serverAddressDefault,
			wantError:         false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
		})
	}
}
