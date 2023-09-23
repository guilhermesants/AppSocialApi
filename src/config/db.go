package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment string
	Port        string
	Database    *Database
}

type Database struct {
	Host     string
	Port     string
	User     string
	DB       string
	Password string
	Driver   string
}

func NewConfig() (*Config, error) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &Config{
		Environment: os.Getenv("ENVIRONMENT"),
		Port:        os.Getenv("PORT"),
		Database: &Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			DB:       os.Getenv("DB_NAME"),
			Password: os.Getenv("DB_PASS"),
			Driver:   os.Getenv("DB_DRIVER"),
		},
	}, nil
}
