package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// loads the environment variables from the .env file
func LoadEnvVariables() {
	envFilePath := ".env"

	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
