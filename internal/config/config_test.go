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
		{
			name: "should create a new Config instance with custom server address: localhost:8081",
			envValues: map[string]string{
				serverAddressKey: "localhost:8081",
			},
			wantServerAddress: "localhost:8081",
			wantError:         false,
		},
		{
			name: "should create a new Config instance with custom server address: 127.0.0.1:3000",
			envValues: map[string]string{
				serverAddressKey: "127.0.0.1:3000",
			},
			wantServerAddress: "127.0.0.1:3000",
			wantError:         false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			for key, value := range testCase.envValues {
				t.Setenv(key, value)
			}
			config, err := New()
			if err == nil && testCase.wantError {
				t.Fatal("New() error=nil")
			}
			if err != nil && !testCase.wantError {
				t.Fatalf("New() error=%v", err)
			}
			if testCase.wantError {
				return
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
