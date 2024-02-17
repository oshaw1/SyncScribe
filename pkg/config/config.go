package config

import (
	"os"
)

// Config represents the application configuration
type Config struct {
	ServerPort    string
	AWSRegion     string
	DynamoDBTable string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	return &Config{
		ServerPort:    getEnv("SERVER_PORT", "8080"), // Default to 8080 if not set
		AWSRegion:     os.Getenv("AWS_REGION"),       // AWS SDK automatically uses this
		DynamoDBTable: os.Getenv("DYNAMODB_TABLE"),
	}, nil
}

// getEnv is a helper function to read an environment variable or return a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
