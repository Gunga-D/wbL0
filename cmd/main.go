package main

import (
	"github.com/sirupsen/logrus"

	"github.com/Gunga-D/taskL0/internal/app"
	"github.com/Gunga-D/taskL0/internal/config"
	"github.com/Gunga-D/taskL0/internal/routes"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	appConfig, err := config.LoadAppConfig("./config")
	if err != nil {
		logrus.Fatalf(err.Error())
	}
	dbConfig, err := config.LoadDBConfig("./config")
	if err != nil {
		logrus.Fatalf(err.Error())
	}

	handler := routes.Init()

	server, err := app.New(appConfig, handler)
	if err != nil {
		logrus.Fatalf(err.Error())
	}
	if err := server.Run(); err != nil {
		logrus.Fatalf(err.Error())
	}
}
