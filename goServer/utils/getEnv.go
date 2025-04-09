package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvNames string

const (
	SERVER_PORT EnvNames = "SERVER_PORT"
	DB_NAME     EnvNames = "DB_NAME"
)

func GetEnv(envName EnvNames) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln(err)
	}

	value := os.Getenv(string(envName))

	if value == "" {
		log.Fatalln(envName, " is empty. Are you sure its configured?")
	}

	return value
}
