package repository

import (
	configs "user_service/internal/config"
	"user_service/internal/model"
	"user_service/internal/repository/postgres"
)

type IUserRepository interface {
	Create(model.UserInfo) (string, error)
	GetUser(uuid string) (model.User, error)
	GetUsers(limit int32, page int32) ([]model.User, error)
	GetUsersByCredentials(text string) ([]model.User, error)
}

type Repository struct {
	User   IUserRepository
	config *configs.Database
}

func New(conf *configs.Database) (*Repository, error) {
	// get pool from database
	postgresDB, err := postgres.NewDatabasePSQL(conf)
	if err != nil {
		return nil, err
	}
	userRepo := postgres.NewUser(postgresDB.DB)
	return &Repository{
		User:   userRepo,
		config: conf,
	}, nil
}
