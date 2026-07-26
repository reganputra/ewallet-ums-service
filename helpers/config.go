package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Env = map[string]string{}

func SetupConfig() {
	var err error
	Env, err = godotenv.Read(".env")
	if err != nil {
		log.Println("[config] .env file not found, relying on OS environment variables")
	}
}

func GetEnv(key, val string) string {
	// 1st priority: OS env var (set by Docker / docker-compose)
	if result := os.Getenv(key); result != "" {
		return result
	}
	// 2nd priority: value from .env file
	if result := Env[key]; result != "" {
		return result
	}
	// 3rd priority: provided default
	return val
}
