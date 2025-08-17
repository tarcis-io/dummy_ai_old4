package config

import (
	"os"
	"testing"
)

// TestNew tests the creation of a new Config instance.
func TestNew(t *testing.T) {
	t.Run("should create a new Config instance with default values", func(t *testing.T) {
		os.Unsetenv(serverAddressKey)
		config, err := New()
		if config == nil {
			t.Error("New() *Config=nil")
		}
		if config != nil && config.ServerAddress != serverAddressDefault {
			t.Errorf("New().ServerAddress got=%s want=%s", config.ServerAddress, serverAddressDefault)
		}
		if err != nil {
			t.Errorf("New() error=%v", err)
		}
	})
	t.Run("should create a new Config instance with custom values", func(t *testing.T) {
		serverAddress := "localhost:8081"
		t.Setenv(serverAddressKey, serverAddress)
		config, err := New()
		if config == nil {
			t.Error("New() *Config=nil")
		}
		if config != nil && config.ServerAddress != serverAddress {
			t.Errorf("New().ServerAddress got=%s want=%s", config.ServerAddress, serverAddress)
		}
		if err != nil {
			t.Errorf("New() error=%v", err)
		}
	})
	t.Run("should return an error if any of the configurations cannot be resolved", func(t *testing.T) {
		serverAddress := "invalid"
		t.Setenv(serverAddressKey, serverAddress)
		config, err := New()
		if config != nil {
			t.Error("New() *Config!=nil")
		}
		if err == nil {
			t.Error("New() error=nil")
		}
	})
}
