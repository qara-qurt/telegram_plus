package service

import (
	"github.com/qara-qurt/telegram_plus/user_service/internal/model"
	"github.com/qara-qurt/telegram_plus/user_service/internal/repository"
)

type IUser interface {
	Create(model.UserInfo) (string, error)
	GetUser(uuid string) (model.User, error)
	GetUsers(limit int32, page int32) ([]model.User, error)
	GetUsersByCredentials(text string) ([]model.User, error)
}

type Service struct {
	User IUser
}

func New(repo *repository.Repository) *Service {
	user := NewUser(repo.User)
	return &Service{
		User: user,
	}
}
