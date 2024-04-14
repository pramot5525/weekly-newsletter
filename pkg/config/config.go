package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	AppPort          string = "APP_PORT"
	PostgresHost     string = "POSTGRES_HOST"
	PostgresPort     string = "POSTGRES_PORT"
	PostgresUser     string = "POSTGRES_USER"
	PostgresPassword string = "POSTGRES_PASSWORD"
	PostgresDatabase string = "POSTGRES_DB"
	EmailHost        string = "EMAIL_HOST"
	EmailPort        string = "EMAIL_PORT"
	EmailSender      string = "EMAIL_SENDER"
	EmailAPIKey      string = "EMAIL_API_KEY"
)

func GetString(key string) string {
	return viper.GetString(key)
}

func Loadenv(file string) error {
	viper.SetConfigType("env")
	viper.SetConfigFile(file)

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("loadenv: %w", err)
	}

	return nil
}
