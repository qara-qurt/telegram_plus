package main

import (
	"context"
	app2 "user_service/internal/app"

	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	app, err := app2.NewApp(ctx)
	if err != nil {
		logrus.Fatal(err)
	}
	if err := app.RunGRPCServer(); err != nil {
		logrus.Fatal(err)
	}

}
