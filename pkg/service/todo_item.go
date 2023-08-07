package service

import (
	auth "AuthWithJWT"
	"AuthWithJWT/pkg/repository"
)

type TodoItemService struct {
	itemRepo repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(itemRepo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{
		itemRepo: itemRepo,
		listRepo: listRepo,
	}
}

func (s *TodoItemService) Create(userId, listId int, item auth.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.itemRepo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]auth.TodoItem, error) {
	return s.itemRepo.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (auth.TodoItem, error) {
	return s.itemRepo.GetById(userId, itemId)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.itemRepo.Delete(userId, itemId)
}

func (s *TodoItemService) Update(userId, itemId int, input auth.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.itemRepo.Update(userId, itemId, input)
}
