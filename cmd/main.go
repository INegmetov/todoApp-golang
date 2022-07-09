package main

import (
	"log"

	"github.com/inegmetov/todoApp-golang"
	"github.com/inegmetov/todoApp-golang/pkg/handler"
	"github.com/inegmetov/todoApp-golang/pkg/repository"
	"github.com/inegmetov/todoApp-golang/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if error := initConfig(); error != nil {
		log.Fatalf("error initializing config: %v", error.Error())
	}

	repos := repository.NewRepository()
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)
	srv := new(todoApp.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {

		log.Fatal("error occurred while running the server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
