package config

import (
	"os"
	"testing"
)

// TestNew tests the creation of a new Config instance.
func TestNew(t *testing.T) {
	t.Run("should create a new Config with default values", func(t *testing.T) {
		os.Unsetenv(serverAddressKey)
		config, err := New()
		if err != nil {
			t.Fatalf("New() error=%v", err)
		}
		if config == nil {
			t.Fatal("New() *Config=nil")
		}
		if config.ServerAddress != serverAddressDefault {
			t.Errorf("New().ServerAddress got=%s want=%s", config.ServerAddress, serverAddressDefault)
		}
	})
}
