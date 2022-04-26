package configs

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
)

var configs Settings

type Settings struct {
	DbHost string `env:"DB_HOST"`
	DbUser string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbName string `env:"DB_NAME"`
	DbPort string `env:"DB_PORT"`
	DbSSLMode string `env:"DB_SSL_MODE" envDefault:"disable"`
	DbConnectionString string
	AppDefaultHost string `env:"APP_DEFAULT_HOST" envDefault:"0.0.0.0"`
	AppDefaultPort string `env:"APP_DEFAULT_PORT" envDefault:"8081"`
	AppBasePath string `env:"BASE_PATH"`
}

func init() {
	ParseEnv()
}

func ParseEnv() {
	log.Println("[INFO] - Config the envs")

	if err := env.Parse(&configs); err != nil {
		log.Fatalf("Unexpected error when trying to parse the envs")
	}

	configs.DbConnectionString = fmt.Sprintf("host=%s user=%s password=%s " +
		"dbname=%s port=%s sslmode=%s", configs.DbHost, configs.DbUser, configs.DbPassword, configs.DbName,
		configs.DbPort, configs.DbSSLMode)
}

func LoadEnvFromFile() {
	godotenv.Load("./.env-dev")
	ParseEnv()
}

func GetConfig() Settings {
	return configs
}
