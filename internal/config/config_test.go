package config

import (
	"testing"
)

// TestNew verifies that the New function correctly initializes a Config instance.
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
		{
			name:              "should create a new Config instance with custom server address: localhost:8081",
			envValues:         map[string]string{serverAddressKey: "localhost:8081"},
			wantServerAddress: "localhost:8081",
			wantError:         false,
		},
		{
			name:              "should create a new Config instance with custom server address: 127.0.0.1:3000",
			envValues:         map[string]string{serverAddressKey: "127.0.0.1:3000"},
			wantServerAddress: "127.0.0.1:3000",
			wantError:         false,
		},
		{
			name:              "should return an error if the server address cannot be resolved: empty",
			envValues:         map[string]string{serverAddressKey: ""},
			wantServerAddress: "",
			wantError:         true,
		},
		{
			name:              "should return an error if the server address cannot be resolved: invalid",
			envValues:         map[string]string{serverAddressKey: "invalid"},
			wantServerAddress: "",
			wantError:         true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			for key, value := range testCase.envValues {
				t.Setenv(key, value)
			}
			config, err := New()
			if testCase.wantError {
				if err == nil {
					t.Fatal("New() error=nil")
				}
				if config != nil {
					t.Fatal("New() *Config!=nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("New() error=%v", err)
			}
			if config == nil {
				t.Fatal("New() *Config=nil")
			}
			if config.ServerAddress != testCase.wantServerAddress {
				t.Errorf("New().ServerAddress got=%s want=%s", config.ServerAddress, testCase.wantServerAddress)
			}
		})
	}
}
