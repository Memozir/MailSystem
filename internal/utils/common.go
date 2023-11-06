package utils

import (
	"os"
)

func GetEnvOrDefault(key string, def string) string {
	value, exists := os.LookupEnv(key)
	if exists || value != "" {
		return value
	}
	return def
}
