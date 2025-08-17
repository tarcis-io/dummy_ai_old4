package config

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("should create a new Config with default values", func(t *testing.T) {
		os.Unsetenv(serverAddressKey)
		config, err := New()
		if err != nil {
		}
		if config == nil {
		}
		if config.ServerAddress != serverAddressDefault {
		}
	})
}
