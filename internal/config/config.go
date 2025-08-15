package config

import (
	"os"
)

type (
	Config struct {
		ServerAddress string
	}
)

func getEnv(key, defaultValue string) string {
	value, found := os.LookupEnv(key)
	if found {
		return value
	}
	return defaultValue
}
