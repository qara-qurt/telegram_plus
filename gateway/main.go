package main

import (
	"log"

	"github.com/labstack/echo/v4"
	pb "github.com/qara-qurt/telegram_plus/user_service/pkg/gen/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	PORT_ADDR = ":3000"
	GPRC_ADDR = "localhost:8080"
)

func main() {
	conn, err := grpc.NewClient(GPRC_ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	user_service := pb.NewUserServiceClient(conn)

	e := echo.New()
	handler := NewHandler(user_service)
	// Register all routes
	handler.InitRoutes(e)

	// Start server
	if err := e.Start(PORT_ADDR); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
