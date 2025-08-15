package config

import (
	"fmt"
	"net"
	"os"
)

type (
	Config struct {
		ServerAddress string
	}
)

const (
	serverAddressKey     = "SERVER_ADDRESS"
	serverAddressDefault = "0.0.0.0:8080"
)

func resolveServerAddress() (string, error) {
	serverAddress := getEnv(serverAddressKey, serverAddressDefault)
	_, _, err := net.SplitHostPort(serverAddress)
	if err != nil {
		return "", fmt.Errorf("invalid server address %s=%s error=%w", serverAddressKey, serverAddress, err)
	}
	return serverAddress, nil
}

func getEnv(key, defaultValue string) string {
	value, found := os.LookupEnv(key)
	if found {
		return value
	}
	return defaultValue
}
