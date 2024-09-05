package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Server  Server
	MongoDB MongoDB
}

type Server struct {
	Port string
}

type MongoDB struct {
	Host       string
	Port       string
	User       string
	Password   string
	SSLMode    string
	DBName     string
	AuthSource string
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
			Port: GetEnv("SERVER_PORT", "5002"),
		},
		MongoDB: MongoDB{
			Host:       GetEnv("MONGO_HOST", "localhost"),
			Port:       GetEnv("MONGO_PORT", "27017"),
			User:       GetEnv("MONGO_USERNAME", "root"),
			Password:   GetEnv("MONGO_PASSWORD", "qaraqurt"),
			DBName:     GetEnv("MONGO_DATABASE", "post_service"),
			AuthSource: GetEnv("MONGO_AUTH_SOURCE", "admin"),
		},
	}, nil
}
