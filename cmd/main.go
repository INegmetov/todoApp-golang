package main

import (
	"log"

	"github.com/inegmetov/todoApp-golang"
	"github.com/inegmetov/todoApp-golang/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todoApp.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal("error occurred while running the server: %s", err.Error())
	}
}
