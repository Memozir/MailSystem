package utils

import (
	"os"
)

func getEnvOrDefault(key string, def string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return def
}
