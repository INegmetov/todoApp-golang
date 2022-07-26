package service

import (
	"github.com/inegmetov/todoApp-golang"
	"github.com/inegmetov/todoApp-golang/pkg/repository"
)

type Authorization interface {
	CreateUser(todoApp.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type TodoList interface {
	Create(userId int, list todoApp.TodoList) (int, error)
	GetAll(userId int) ([]todoApp.TodoList, error)
	GetById(userId, listId int) (todoApp.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todoApp.UpdateListInput) error
}
type TodoItem interface {
	Create(userId, listId int, item todoApp.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todoApp.TodoItem, error)
	GetById(userId, itemId int) (todoApp.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todoApp.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
