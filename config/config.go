package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func Config(key string) string {
	err := godotenv.Load("/tmp/.env-go-api")
	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err)
	}
	return os.Getenv(key)
}
