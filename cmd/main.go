package main

import (
	"clother"
	"clother/pkg/handler"
	"clother/pkg/repository"
	"clother/pkg/service"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrus.SetReportCaller(true)

	db, err := repository.NewConnection(os.Getenv("APP_DSN"))
	if err != nil {
		logrus.Fatal(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	engine := handlers.InitRoutes()

	srv := clother.Server{
		Engine: engine,
	}
	err = srv.Start(os.Getenv("APP_ADDR"))
	if err != nil {
		logrus.Fatal(err)
	}
}
