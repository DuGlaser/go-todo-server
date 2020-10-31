package todo

import (
	"github.com/DuGlaser/go-todo-server/domain/todo"
)

var TodoService todoServiceInterface = &todoService{}

type todoServiceInterface interface {
	GetAll() (*todo.Todos, error)
	GetByID(int64) (*todo.Todo, error)
	Create(string, string) (*todo.Todo, error)
	Update(todo.Todo) (*todo.Todo, error)
	Delete(int64) error
}

type todoService struct{}

func (service *todoService) GetAll() (*todo.Todos, error) {
	dao := todo.Todo{}
	result, err := dao.GetAll()
	if err != nil {
		return nil, err
	}

	return &result, err
}

func (service *todoService) GetByID(id int64) (*todo.Todo, error) {
	dao := todo.Todo{
		ID: id,
	}

	if err := dao.GetByID(); err != nil {
		return nil, err
	}

	return &dao, nil
}

func (service *todoService) Create(title, description string) (*todo.Todo, error) {
	dao := todo.Todo{
		Title:       title,
		Description: description,
		Status:      todo.Doing,
	}
	if err := dao.Validate(); err != nil {
		return nil, err
	}

	if err := dao.Save(); err != nil {
		return nil, err
	}

	return &dao, nil
}

func (service *todoService) Update(updateTodo todo.Todo) (*todo.Todo, error) {
	if err := updateTodo.Validate(); err != nil {
		return nil, err
	}

	if err := updateTodo.Update(); err != nil {
		return nil, err
	}

	return &updateTodo, nil
}

func (service *todoService) Delete(id int64) error {
	dao := todo.Todo{
		ID: id,
	}

	if err := dao.Delete(); err != nil {
		return err
	}

	return nil
}
