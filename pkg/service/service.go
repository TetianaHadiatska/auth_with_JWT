package service

import (
	auth "AuthWithJWT"
	"AuthWithJWT/pkg/repository"
)

type Authorization interface {
	CreateUser(user auth.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type TodoList interface {
	Create(userId int, list auth.TodoList) (int, error)
	GetAll(userId int) ([]auth.TodoList, error)
	GetById(userId, listId int) (auth.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input auth.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item auth.TodoItem) (int, error)
	GetAll(userId, listId int) ([]auth.TodoItem, error)
	GetById(userId, itemId int) (auth.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input auth.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
