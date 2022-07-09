package service

import "github.com/inegmetov/todoApp-golang/pkg/repository"

type Authorization interface {
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
	return &Service{}
}
