package main

import (
	"SandBox/config"
	"SandBox/internal/adapters/handler"
	"SandBox/internal/domain/repository"
	"SandBox/internal/domain/service"
	"SandBox/package/server"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	dbConn := server.NewDBConn()

	repositories := repository.NewRepository(dbConn)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	httpServer := &server.Server{}
	err := httpServer.Run(viper.GetString("port"), handlers.InitRoutes())
	if err != nil {
		log.Fatal(err)
	}
}
