package env

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"template/src/utils/logs"
)

// LoadEnv loads environment variables from the .env file (if present).
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Warn(".env file not found, relying on actual environment variables", nil)
	}
}

// GetEnv retrieves an environment variable and validates whether it's optional or required.
// If a required variable is missing, it logs a critical error.
// If an optional variable is missing, it logs a warning.
func GetEnv(key string, required bool) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		if required {
			err := fmt.Errorf("environment variable %s is required but not set", key)
			log.Error(err, nil)
			return "", err
		} else {
			log.Warn(fmt.Sprintf("environment variable %s is optional but not set", key), nil)
			return "", nil
		}
	}
	return value, nil
}

// GetEnvOrDefault retrieves an environment variable or returns the default value if not found.
// If a required variable is missing, it logs a critical error.
func GetEnvOrDefault(key string, defaultValue string) string {
	value, err := GetEnv(key, false)
	if err != nil || value == "" {
		return defaultValue
	}
	return value
}

// CheckRequiredEnv checks a list of required environment variables and returns an error if any are missing.
func CheckRequiredEnv(keys []string) error {
	var missingVars []string
	for _, key := range keys {
		_, err := GetEnv(key, true)
		if err != nil {
			missingVars = append(missingVars, key)
		}
	}

	if len(missingVars) > 0 {
		return errors.New(fmt.Sprintf("Missing required environment variables: %v", missingVars))
	}
	return nil
}
