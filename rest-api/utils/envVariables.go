package utils

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from .env file
func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        logger := GetLogger()
        logger.Warn().Msg("Error loading .env file, using system environment variables")
    }
}

// GetEnv retrieves an environment variable or returns a default value if not found
func GetEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}

// GetEnvAsInt retrieves an environment variable as an integer
func GetEnvAsInt(key string, defaultValue int) int {
    valueStr := GetEnv(key, "")
    if valueStr == "" {
        return defaultValue
    }
    
    value, err := strconv.Atoi(valueStr)
    if err != nil {
        return defaultValue
    }
    
    return value
}

// GetEnvAsBool retrieves an environment variable as a boolean
func GetEnvAsBool(key string, defaultValue bool) bool {
    valueStr := GetEnv(key, "")
    if valueStr == "" {
        return defaultValue
    }
    
    value, err := strconv.ParseBool(valueStr)
    if err != nil {
        return defaultValue
    }
    
    return value
}