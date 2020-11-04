package usecase

import "github.com/DuGlaser/go-todo-server/domain"

type todoUsecase struct {
	todoRepo domain.TodoRepository
}

func NewTodoUsecase(tr domain.TodoRepository) domain.TodoUsecase {
	return &todoUsecase{
		todoRepo: tr,
	}
}

func (u *todoUsecase) GetAll() (domain.Todos, *domain.RestErr) {
	todos, err := u.todoRepo.GetAll()
	if err != nil {
		return domain.Todos{}, err
	}

	return todos, err
}

func (u *todoUsecase) GetByID(id int64) (domain.Todo, *domain.RestErr) {
	todo, err := u.todoRepo.GetByID(id)
	if err != nil {
		return domain.Todo{}, err
	}

	return todo, nil
}

func (u *todoUsecase) Store(t *domain.Todo) *domain.RestErr {
	err := u.todoRepo.Store(t)

	return err
}

func (u *todoUsecase) Update(t *domain.Todo) *domain.RestErr {
	err := u.todoRepo.Update(t)

	return err
}

func (u *todoUsecase) Delete(id int64) *domain.RestErr {
	err := u.todoRepo.Delete(id)

	return err
}
