package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadENV() error {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" || goEnv == "dev" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}

	return nil
}
