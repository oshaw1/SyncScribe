package config

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
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
		ServerPort:    getEnv("SERVER_PORT", "5000"), // Default to 8080 if not set
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

func getAWSCredentials() (string, string) {
	sess, err := session.NewSession()
	if err != nil {
		log.Fatalf("failed to create session: %v", err)
	}

	// Get credentials from the session
	creds, err := sess.Config.Credentials.Get()
	if err != nil {
		log.Fatalf("failed to get credentials: %v", err)
	}

	accessKeyID := creds.AccessKeyID
	secretAccessKey := creds.SecretAccessKey

	return accessKeyID, secretAccessKey
}
