package main

import (
	"github.com/sirupsen/logrus"

	"github.com/Gunga-D/taskL0/internal/app"
	"github.com/Gunga-D/taskL0/internal/config"
	"github.com/Gunga-D/taskL0/internal/repository/postgres"
	"github.com/Gunga-D/taskL0/internal/routes"
	"github.com/Gunga-D/taskL0/internal/services"
	"github.com/Gunga-D/taskL0/internal/storage"
	"github.com/Gunga-D/taskL0/internal/stream"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	appConfig, err := config.LoadAppConfig("./configs", "app_config")
	if err != nil {
		logrus.Fatalf(err.Error())
	}
	dbConfig, err := config.LoadDBConfig("./configs", "db_config")
	if err != nil {
		logrus.Fatalf(err.Error())
	}
	natStreamingConfig, err := config.LoadNATStreamingConfig("./configs", "natstreaming_config")
	if err != nil {
		logrus.Fatalf(err.Error())
	}

	dbKernel, err := postgres.CreateKernel(dbConfig)
	if err != nil {
		logrus.Fatalf(err.Error())
	}
	ordersTable := postgres.NewPostgresOrdersTable(dbKernel)

	ordersCache := storage.NewOrdersMemory()
	if err := ordersCache.Recover(&ordersTable); err != nil {
		logrus.Fatalf(err.Error())
	}

	ordersService := services.NewOrderService(ordersCache, ordersTable)

	ordersStream := stream.NewOrdersStream(ordersService)
	if err := stream.Init(natStreamingConfig, ordersStream); err != nil {
		logrus.Fatalf(err.Error())
	}

	ordersRoute := routes.NewOrdersRoute(ordersService)
	handler := routes.Init(ordersRoute)

	server, err := app.New(appConfig, handler)
	if err != nil {
		logrus.Fatalf(err.Error())
	}
	if err := server.Run(); err != nil {
		logrus.Fatalf(err.Error())
	}
}
