package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvValues(env_value_name string) string {

	err := godotenv.Load()

	if err != nil {

		log.Fatal("Error while loading .env file")

	}

	return os.Getenv(env_value_name)
}
