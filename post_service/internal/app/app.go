package app

import (
	"fmt"
	"net"

	"github.com/qara-qurt/telegram_plus/post_service/internal/api"
	"github.com/qara-qurt/telegram_plus/post_service/internal/config"
	"github.com/qara-qurt/telegram_plus/post_service/internal/repository"
	"github.com/qara-qurt/telegram_plus/post_service/internal/service"
	desk "github.com/qara-qurt/telegram_plus/post_service/pkg/gen/post"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	grpcServer *grpc.Server
	config     *config.Config
}

func NewApp() (*App, error) {
	// Get config
	conf, err := config.New()
	if err != nil {
		return nil, err
	}

	a := &App{
		config: conf,
	}

	// INIT Repository
	repo, err := repository.New(conf.MongoDB)
	if err != nil {
		return nil, err
	}

	service := service.New(repo)
	if err := a.initGRPCServer(service); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initGRPCServer(service *service.Service) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(a.grpcServer)

	// Register controllers
	desk.RegisterPostServiceServer(a.grpcServer, api.New(service.Post))
	return nil
}

func (a *App) RunGRPCServer() error {
	port := fmt.Sprintf(":%s", a.config.Server.Port)
	logrus.Infof("GRPC server is running on %s", port)

	list, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
