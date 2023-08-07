package repository

import (
	auth "AuthWithJWT"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user auth.User) (int, error)
	GetUser(username, password string) (auth.User, error)
}

type TodoList interface {
	Create(userId int, list auth.TodoList) (int, error)
	GetAll(userId int) ([]auth.TodoList, error)
	GetById(userId, listId int) (auth.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input auth.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item auth.TodoItem) (int, error)
	GetAll(userId, listId int) ([]auth.TodoItem, error)
	GetById(userId, itemId int) (auth.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input auth.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
