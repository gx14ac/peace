package config

import (
	"os"
)

type Config struct {
	AppService string
	AppEnv     string
	AppPort    string
	CorsDomain string

	BasicUser string
	BasicPass string

	MaintennanceOngoing string

	DBHost string
	DBPort string
	DBName string
	DBUser string
	DBPass string

	JWTSecret string

	SlackToken string
}

func NewConfig() *Config {
	return &Config{
		AppService: string(os.Getenv("APP_SERVICE")),
		AppEnv:     string(os.Getenv("APP_ENV")),
		AppPort:    string(os.Getenv("APP_PORT")),

		CorsDomain: string(os.Getenv("CORS_DOMAIN")),

		BasicUser: string(os.Getenv("BASIC_USER")),
		BasicPass: string(os.Getenv("BASIC_PASS")),

		DBHost: string(os.Getenv("DB_HOST")),
		DBPort: string(os.Getenv("DB_PORT")),
		DBName: string(os.Getenv("DB_NAME")),
		DBUser: string(os.Getenv("DB_USER")),
		DBPass: string(os.Getenv("DB_PASS")),

		JWTSecret: string(os.Getenv("JWT_SECRET")),
	}
}
