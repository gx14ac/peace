package config

import (
	"os"
)

type Config struct {
	AppService string
	AppEnv     string
	AppPort    string

	BasicUser string
	BasicPass string

	MaintennanceOngoing string

	DBHost string
	DBPort string
	DBName string
	DBUser string
	DBPass string

	SlackToken string
}

func New() Config {
	return Config{
		AppService:          string(os.Getenv("APP_SERVICE")),
		AppEnv:              string(os.Getenv("APP_ENV")),
		AppPort:             string(os.Getenv("APP_PORT")),
		MaintennanceOngoing: string(os.Getenv("MAINTENANCE_ONGOING")),

		DBHost: string(os.Getenv("DB_HOST")),
		DBPort: string(os.Getenv("DB_PORT")),
		DBName: string(os.Getenv("DB_NAME")),
		DBUser: string(os.Getenv("DB_USER")),
		DBPass: string(os.Getenv("DB_PASS")),
	}
}
