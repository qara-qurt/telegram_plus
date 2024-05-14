package app

import (
	"context"
	"fmt"
	"net"
	user "user_service/internal/api"
	configs "user_service/internal/config"
	"user_service/internal/repository"
	"user_service/internal/service"
	desk "user_service/pkg/gen/user"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	grpcServer *grpc.Server
	config     *configs.Config
}

func NewApp(ctx context.Context) (*App, error) {
	// Get config
	conf, err := configs.New()
	if err != nil {
		logrus.Errorf("config cannot initialized, %v", err)
		return nil, err
	}
	a := &App{
		config: conf,
	}

	// INIT Repository
	repo, err := repository.New(&conf.Database)
	if err != nil {
		logrus.Errorf("error with repository initialize, %v", err)
		return nil, err
	}
	// INIT Service
	serv := service.New(repo)
	// INIT GRPC Server
	err = a.initGRPCServer(serv)
	if err != nil {
		logrus.Errorf("error with initialize grpc server, %v", err)
		return nil, err
	}
	return a, nil
}

func (a *App) initGRPCServer(service *service.Service) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(a.grpcServer)

	// Register controllers
	desk.RegisterUserServiceServer(a.grpcServer, user.New(service.User))
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
