// Package config provides configuration for the application.
package config

import (
	"fmt"
	"net"
	"os"
)

type (
	// Config represents the configuration for the application.
	Config struct {
		// ServerAddress specifies the host and port on which the server will listen.
		ServerAddress string
	}
)

const (
	// serverAddressKey is the environment variable key for the server address.
	serverAddressKey = "SERVER_ADDRESS"

	// serverAddressDefault is the default value for the server address.
	serverAddressDefault = "0.0.0.0:8080"
)

// New creates and returns a new Config instance by resolving the configurations for the application.
// It returns an error if any of the configurations cannot be resolved.
func New() (*Config, error) {
	serverAddress, err := resolveServerAddress()
	if err != nil {
		return nil, err
	}
	config := &Config{
		ServerAddress: serverAddress,
	}
	return config, nil
}

// resolveServerAddress resolves the server address configuration.
// It returns an error if the configuration cannot be resolved.
func resolveServerAddress() (string, error) {
	serverAddress := getEnv(serverAddressKey, serverAddressDefault)
	_, _, err := net.SplitHostPort(serverAddress)
	if err != nil {
		return "", fmt.Errorf("invalid server address %s=%s error=%w", serverAddressKey, serverAddress, err)
	}
	return serverAddress, nil
}

// getEnv retrieves the value of the environment variable specified by the key.
// It returns the default value if the environment variable is not set.
func getEnv(key, defaultValue string) string {
	value, found := os.LookupEnv(key)
	if found {
		return value
	}
	return defaultValue
}
