package api

import (
	"context"

	"github.com/qara-qurt/telegram_plus/post_service/internal/service"
	desk "github.com/qara-qurt/telegram_plus/post_service/pkg/gen/post"
)

type Implementation struct {
	desk.UnimplementedPostServiceServer
	postService service.IPost
}

func New(service service.IPost) *Implementation {
	return &Implementation{
		postService: service,
	}
}

func (i *Implementation) GetPost(ctx context.Context, req *desk.GetPostRequest) (*desk.GetPostsResponse, error) {
	return nil, nil
}
