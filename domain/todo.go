package domain

import (
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

func (t *Todo) Validate() *RestErr {
	t.Title = strings.TrimSpace(t.Title)
	t.Description = strings.TrimSpace(t.Description)

	if t.Title == "" {
		return NewBadRequest("title is empty")
	}

	if t.Description == "" {
		return NewBadRequest("description is empty")
	}

	if !(t.Status == Doing || t.Status == Done) {
		return NewBadRequest("invalid status")
	}

	return nil
}

type TodoUsecase interface {
	GetAll() (Todos, *RestErr)
	GetByID(int64) (Todo, *RestErr)
	Store(*Todo) *RestErr
	Update(*Todo) *RestErr
	Delete(int64) *RestErr
}

type TodoRepository interface {
	GetAll() (Todos, *RestErr)
	GetByID(int64) (Todo, *RestErr)
	Store(*Todo) *RestErr
	Update(*Todo) *RestErr
	Delete(int64) *RestErr
}
