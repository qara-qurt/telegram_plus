package service

import "github.com/qara-qurt/telegram_plus/post_service/internal/repository"

type Post struct {
	repo repository.IPostRepository
}

func NewPost(repo repository.IPostRepository) *Post {
	return &Post{
		repo: repo,
	}
}
