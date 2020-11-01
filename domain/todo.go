package domain

import (
	"errors"
	"strings"
)

type Status string

const (
	Doing = Status("Doing")
	Done  = Status("Done")
)

type Todo struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}

type Todos []Todo

func (t *Todo) Validate() error {
	t.Title = strings.TrimSpace(t.Title)
	t.Description = strings.TrimSpace(t.Description)

	if t.Title == "" {
		return errors.New("title is empty")
	}

	if t.Description == "" {
		return errors.New("description is empty")
	}

	if !(t.Status == Doing || t.Status == Done) {
		return errors.New("invalid status")
	}

	return nil
}

type TodoUsecase interface {
	GetAll() (Todos, error)
	GetByID(int64) (Todo, error)
	Store(*Todo) error
	Update(*Todo) error
	Delete(int64) error
}

type TodoRepository interface {
	GetAll() (Todos, error)
	GetByID(int64) (Todo, error)
	Store(*Todo) error
	Update(*Todo) error
	Delete(int64) error
}
