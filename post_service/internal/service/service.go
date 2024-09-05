package service

import "github.com/qara-qurt/telegram_plus/post_service/internal/repository"

type IPost interface {
}

type Service struct {
	Post IPost
}

func New(repo *repository.Repository) *Service {
	return &Service{
		Post: NewPost(repo.Post),
	}
}
