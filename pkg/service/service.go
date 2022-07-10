package service

import (
	"github.com/inegmetov/todoApp-golang"
	"github.com/inegmetov/todoApp-golang/pkg/repository"
)

type Authorization interface {
	CreateUser(todoApp.User) (int, error)
	GenerateToken( username, password string ) (string, error)
}
type TodoList interface {
}
type TodoItem interface {
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
