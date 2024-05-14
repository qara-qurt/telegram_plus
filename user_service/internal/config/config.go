package configs

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	Server   Server
	Database Database
}

type Server struct {
	Port string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	SSLMode  string
	DBName   string
}

func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func New() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return &Config{
		Server: Server{
			Port: GetEnv("SERVER_PORT", "5001"),
		},
		Database: Database{
			Host:     GetEnv("DB_HOST", "localhost"),
			Port:     GetEnv("DB_PORT", "5432"),
			User:     GetEnv("DB_USER", "postgres"),
			Password: GetEnv("DB_PASSWORD", "password"),
			DBName:   GetEnv("DB_NAME", "postgres"),
			SSLMode:  GetEnv("SSL_MODE", "disable"),
		},
	}, nil
}
