package main

import (
	"context"

	"github.com/labstack/echo/v4"
	pb "github.com/qara-qurt/telegram_plus/user_service/pkg/gen/user"
)

type handler struct {
	user_service pb.UserServiceClient
}

func NewHandler(client pb.UserServiceClient) *handler {
	return &handler{
		user_service: client,
	}
}

func (h *handler) InitRoutes(http *echo.Echo) {
	http.GET("/ping", h.Ping)
	http.GET("/users", h.GetUsers)
}

func (h *handler) Ping(c echo.Context) error {
	return c.String(200, "pong gang shit")
}

func (h *handler) GetUsers(c echo.Context) error {
	users, err := h.user_service.GetUsers(context.Background(), &pb.GetUsersRequests{
		Page:  1,
		Limit: 10,
	})
	if err != nil {
		return c.String(500, err.Error())
	}
	return c.JSON(200, users)
}
