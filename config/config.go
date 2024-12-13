package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type DBConfig struct {
	ConnectionString string `envconfig:"PRICEINDEX_DB_URL"`
	MigrationsDir    string `envconfig:"PRICEINDEX_DB_MIGRATIONS_DIR"`
}

type Config struct {
	DB DBConfig
}

func LoadConfig() *Config {
	cfg := &Config{}

	configFile := godotenv.Load(".env")
	if configFile != nil {
		log.Fatalln("Error loading .env file")
	}

	if err := envconfig.Process("", cfg); err != nil {
		log.Fatalln("Error processing environment variables")
	}

	return cfg
}
