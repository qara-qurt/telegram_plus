package repository

import (
	"github.com/qara-qurt/telegram_plus/post_service/internal/config"
	"github.com/qara-qurt/telegram_plus/post_service/internal/repository/mongo"
)

type IPostRepository interface {
}

type Repository struct {
	Post IPostRepository
}

func New(conf config.MongoDB) (*Repository, error) {
	mongodb, err := mongo.NewMongo(conf)
	if err != nil {
		return nil, err
	}

	postRepo := mongo.NewPost(mongodb.Mongodb)

	return &Repository{
		Post: postRepo,
	}, nil
}
