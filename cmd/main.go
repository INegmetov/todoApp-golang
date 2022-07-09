package main

import (
	"log"

	"github.com/inegmetov/todoApp-golang"
	"github.com/inegmetov/todoApp-golang/pkg/handler"
	"github.com/inegmetov/todoApp-golang/pkg/repository"
	"github.com/inegmetov/todoApp-golang/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)
	srv := new(todoApp.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal("error occurred while running the server: %s", err.Error())
	}
}
