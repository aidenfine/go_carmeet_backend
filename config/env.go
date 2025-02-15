package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	GO_ENV                 string
	JWTSecret              string
	JWTExpirationInSeconds int64
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		GO_ENV:                 getEnv("GO_ENV", "dev"),
		JWTSecret:              getEnv("JWTSecret", "no"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}

func LoadEnv() error {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" || goEnv == "dev" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}
