package usecase_test

import (
	"testing"

	"github.com/DuGlaser/go-todo-server/domain"
	"github.com/DuGlaser/go-todo-server/domain/mocks"
	"github.com/DuGlaser/go-todo-server/todo/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	mockTodos := make(domain.Todos, 0)
	mockTodo1 := domain.Todo{
		ID:          1,
		Title:       "test_title_1",
		Description: "test_description_1",
		Status:      "Doing",
	}
	mockTodo2 := domain.Todo{
		ID:          2,
		Title:       "test_title_2",
		Description: "test_description_2",
		Status:      "Doing",
	}

	mockTodos = append(mockTodos, mockTodo1)
	mockTodos = append(mockTodos, mockTodo2)

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("GetAll").Return(mockTodos, nil).Once()
		u := usecase.NewTodoUsecase(mockTodoRepo)

		res, restErr := u.GetAll()

		assert.Nil(t, restErr)
		assert.EqualValues(t, mockTodos[0], res[0])
		assert.EqualValues(t, mockTodos[1], res[1])
	})

	t.Run("error-failed", func(t *testing.T) {
		mockTodoRepo.On("GetAll").Return(domain.Todos{}, domain.NewInternalServerError("Unexpected")).Once()
		u := usecase.NewTodoUsecase(mockTodoRepo)

		res, restErr := u.GetAll()

		assert.NotNil(t, restErr)
		assert.EqualValues(t, domain.Todos{}, res)
	})
}

func TestGetByID(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	mockTodo := domain.Todo{
		ID:          1,
		Title:       "test_title",
		Description: "test_description",
		Status:      "Doing",
	}

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("GetByID", mock.AnythingOfType("int64")).Return(mockTodo, nil).Once()
		u := usecase.NewTodoUsecase(mockTodoRepo)

		res, restErr := u.GetByID(mockTodo.ID)

		assert.Nil(t, restErr)
		assert.EqualValues(t, res.ID, mockTodo.ID)
		assert.EqualValues(t, res.Title, mockTodo.Title)
		assert.EqualValues(t, res.Description, mockTodo.Description)
		assert.EqualValues(t, res.Status, mockTodo.Status)

		mockTodoRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockTodoRepo.On("GetByID", mock.AnythingOfType("int64")).Return(domain.Todo{}, domain.NewInternalServerError("Unexpected")).Once()
		u := usecase.NewTodoUsecase(mockTodoRepo)

		res, restErr := u.GetByID(mockTodo.ID)

		assert.NotNil(t, restErr)
		assert.EqualValues(t, domain.Todo{}, res)
	})
}

func TestStore(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	mockTodo := domain.Todo{
		Title:       "test_title",
		Description: "test_description",
		Status:      "Doing",
	}

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("Store", mock.Anything).Return(nil).Once()
		u := usecase.NewTodoUsecase(mockTodoRepo)

		restErr := u.Store(&mockTodo)

		assert.Nil(t, restErr)
		mockTodoRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockTodoRepo.On("Store", mock.Anything).Return(domain.NewInternalServerError("Unexpected")).Once()
		u := usecase.NewTodoUsecase(mockTodoRepo)

		restErr := u.Store(&mockTodo)

		assert.NotNil(t, restErr)
		mockTodoRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("Delete", mock.AnythingOfType("int64")).Return(nil).Once()
		u := usecase.NewTodoUsecase(mockTodoRepo)

		restErr := u.Delete(1)

		assert.Nil(t, restErr)
		mockTodoRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockTodoRepo.On("Delete", mock.AnythingOfType("int64")).Return(domain.NewInternalServerError("Unexpected")).Once()
		u := usecase.NewTodoUsecase(mockTodoRepo)

		restErr := u.Delete(1)

		assert.NotNil(t, restErr)
		mockTodoRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	mockTodo := domain.Todo{
		Title:       "test_title",
		Description: "test_description",
		Status:      "Doing",
	}

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("Update", mock.Anything).Return(nil).Once()
		u := usecase.NewTodoUsecase(mockTodoRepo)

		restErr := u.Update(&mockTodo)

		assert.Nil(t, restErr)
		mockTodoRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockTodoRepo.On("Update", mock.Anything).Return(domain.NewInternalServerError("Unexpected")).Once()
		u := usecase.NewTodoUsecase(mockTodoRepo)

		restErr := u.Update(&mockTodo)

		assert.NotNil(t, restErr)
		mockTodoRepo.AssertExpectations(t)
	})
}
