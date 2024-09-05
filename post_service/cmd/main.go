package main

import (
	app2 "github.com/qara-qurt/telegram_plus/post_service/internal/app"
	"github.com/sirupsen/logrus"
)

func main() {
	app, err := app2.NewApp()
	if err != nil {
		logrus.Fatal(err)
	}
	if err := app.RunGRPCServer(); err != nil {
		logrus.Fatal(err)
	}
}
