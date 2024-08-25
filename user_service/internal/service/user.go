package service

import (
	"github.com/qara-qurt/telegram_plus/user_service/internal/model"
	"github.com/qara-qurt/telegram_plus/user_service/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	repo repository.IUserRepository
}

func NewUser(repo repository.IUserRepository) *User {
	return &User{
		repo: repo,
	}
}

func (u *User) Create(user model.UserInfo) (string, error) {
	hashPassword, err := hashPassword(user.Password)
	if err != nil {
		return "", err
	}

	user.Password = hashPassword
	return u.repo.Create(user)
}

func (u *User) GetUser(uuid string) (model.User, error) {
	return u.repo.GetUser(uuid)
}

func (u *User) GetUsers(limit int32, page int32) ([]model.User, error) {
	return u.repo.GetUsers(limit, page)
}

func (u *User) GetUsersByCredentials(text string) ([]model.User, error) {
	return u.repo.GetUsersByCredentials(text)
}

// Other functions
func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
