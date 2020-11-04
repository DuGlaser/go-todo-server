package http_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DuGlaser/go-todo-server/domain"
	"github.com/DuGlaser/go-todo-server/domain/mocks"
	controller "github.com/DuGlaser/go-todo-server/todo/controller/http"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	UnexpectedErr = domain.NewInternalServerError("Unexpected")
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func performGetRequest(r http.Handler, path string) *httptest.ResponseRecorder {
	return performRequest(r, http.MethodGet, path)
}

func performDeleteRequest(r http.Handler, path string) *httptest.ResponseRecorder {
	return performRequest(r, http.MethodDelete, path)
}

func performPutRequest(r http.Handler, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(http.MethodPut, path, body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func performPostRequest(r http.Handler, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(http.MethodPost, path, body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestGetAll(t *testing.T) {
	todoUsecase := new(mocks.TodoUsecase)
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
		router := gin.Default()
		todoUsecase.On("GetAll").Return(mockTodos, nil).Once()

		controller.NewTodoHandler(router, todoUsecase)

		res := performGetRequest(router, "/todos")

		var resTodos domain.Todos
		json.Unmarshal([]byte(res.Body.String()), &resTodos)

		assert.EqualValues(t, http.StatusOK, res.Code)
		assert.EqualValues(t, mockTodos, resTodos)
	})

	t.Run("error-GetAll-failed", func(t *testing.T) {
		router := gin.Default()
		todoUsecase.On("GetAll").Return(domain.Todos{}, UnexpectedErr).Once()

		controller.NewTodoHandler(router, todoUsecase)

		res := performGetRequest(router, "/todos")

		var restErr domain.RestErr
		json.Unmarshal([]byte(res.Body.String()), &restErr)

		assert.EqualValues(t, http.StatusInternalServerError, res.Code)
		assert.EqualValues(t, *UnexpectedErr, restErr)
	})
}

func TestGetByID(t *testing.T) {
	todoUsecase := new(mocks.TodoUsecase)
	mockTodo := domain.Todo{
		ID:          1,
		Title:       "test_title_1",
		Description: "test_description_1",
		Status:      "Doing",
	}

	t.Run("success", func(t *testing.T) {
		router := gin.Default()
		todoUsecase.On("GetByID", mock.AnythingOfType("int64")).Return(mockTodo, nil).Once()

		controller.NewTodoHandler(router, todoUsecase)

		res := performGetRequest(router, "/todo/1")

		var resTodo domain.Todo
		json.Unmarshal([]byte(res.Body.String()), &resTodo)

		assert.EqualValues(t, http.StatusOK, res.Code)
		assert.EqualValues(t, mockTodo, resTodo)
	})

	t.Run("error-GetByID-failed", func(t *testing.T) {
		router := gin.Default()
		todoUsecase.On("GetByID", mock.AnythingOfType("int64")).Return(domain.Todo{}, UnexpectedErr).Once()

		controller.NewTodoHandler(router, todoUsecase)

		res := performGetRequest(router, "/todo/1")

		var restErr domain.RestErr
		json.Unmarshal([]byte(res.Body.String()), &restErr)

		assert.EqualValues(t, http.StatusInternalServerError, res.Code)
		assert.EqualValues(t, *UnexpectedErr, restErr)
	})

	t.Run("error-invalid-id-failed", func(t *testing.T) {
		router := gin.Default()

		controller.NewTodoHandler(router, todoUsecase)
		res := performGetRequest(router, "/todo/invalid_id")

		assert.EqualValues(t, http.StatusInternalServerError, res.Code)
	})
}

func TestCreate(t *testing.T) {
	todoUsecase := new(mocks.TodoUsecase)
	mockTodo := domain.Todo{
		ID:          1,
		Title:       "test_title",
		Description: "test_description",
		Status:      "Doing",
	}

	mockInvalidStatusTodo := domain.Todo{
		ID:          1,
		Title:       "test_title",
		Description: "test_description",
		Status:      "Invalid",
	}

	mockInvalidIDTodo := []byte(`{"id":"1"}`)

	t.Run("success", func(t *testing.T) {
		router := gin.Default()
		todoUsecase.On("Store", mock.Anything).Return(nil).Once()

		controller.NewTodoHandler(router, todoUsecase)

		data, _ := json.Marshal(mockTodo)
		res := performPostRequest(router, "/todo", bytes.NewBuffer(data))

		var resTodo domain.Todo
		json.Unmarshal([]byte(res.Body.String()), &resTodo)

		assert.EqualValues(t, http.StatusCreated, res.Code)
		assert.EqualValues(t, mockTodo, resTodo)
	})

	t.Run("error-validate-failed", func(t *testing.T) {
		router := gin.Default()

		controller.NewTodoHandler(router, todoUsecase)

		data, _ := json.Marshal(mockInvalidStatusTodo)
		res := performPostRequest(router, "/todo", bytes.NewBuffer(data))

		assert.EqualValues(t, http.StatusBadRequest, res.Code)
	})

	t.Run("error-store-failed", func(t *testing.T) {
		router := gin.Default()
		todoUsecase.On("Store", mock.Anything).Return(UnexpectedErr).Once()

		controller.NewTodoHandler(router, todoUsecase)

		data, _ := json.Marshal(mockTodo)
		res := performPostRequest(router, "/todo", bytes.NewBuffer(data))

		assert.EqualValues(t, http.StatusInternalServerError, res.Code)
	})

	t.Run("error-bind-json-failed", func(t *testing.T) {
		router := gin.Default()

		controller.NewTodoHandler(router, todoUsecase)

		data, _ := json.Marshal(mockInvalidIDTodo)
		res := performPostRequest(router, "/todo", bytes.NewBuffer(data))

		assert.EqualValues(t, http.StatusInternalServerError, res.Code)
	})
}

func TestUpdate(t *testing.T) {
	todoUsecase := new(mocks.TodoUsecase)
	mockTodo := domain.Todo{
		ID:          1,
		Title:       "test_title",
		Description: "test_description",
		Status:      "Doing",
	}

	mockInvalidStatusTodo := domain.Todo{
		ID:          1,
		Title:       "test_title",
		Description: "test_description",
		Status:      "Invalid",
	}

	mockInvalidIDTodo := []byte(`{"id":"1"}`)

	t.Run("success", func(t *testing.T) {
		router := gin.Default()
		todoUsecase.On("Update", mock.Anything).Return(nil).Once()

		controller.NewTodoHandler(router, todoUsecase)

		data, _ := json.Marshal(mockTodo)
		res := performPutRequest(router, "/todo", bytes.NewBuffer(data))

		var resTodo domain.Todo
		json.Unmarshal([]byte(res.Body.String()), &resTodo)

		assert.EqualValues(t, http.StatusCreated, res.Code)
		assert.EqualValues(t, mockTodo, resTodo)
	})

	t.Run("error-validate-failed", func(t *testing.T) {
		router := gin.Default()

		controller.NewTodoHandler(router, todoUsecase)

		data, _ := json.Marshal(mockInvalidStatusTodo)
		res := performPutRequest(router, "/todo", bytes.NewBuffer(data))

		assert.EqualValues(t, http.StatusBadRequest, res.Code)
	})

	t.Run("error-update-failed", func(t *testing.T) {
		router := gin.Default()
		todoUsecase.On("Update", mock.Anything).Return(UnexpectedErr).Once()

		controller.NewTodoHandler(router, todoUsecase)

		data, _ := json.Marshal(mockTodo)
		res := performPutRequest(router, "/todo", bytes.NewBuffer(data))

		assert.EqualValues(t, http.StatusInternalServerError, res.Code)
	})

	t.Run("error-bind-json-failed", func(t *testing.T) {
		router := gin.Default()

		controller.NewTodoHandler(router, todoUsecase)

		data, _ := json.Marshal(mockInvalidIDTodo)
		res := performPutRequest(router, "/todo", bytes.NewBuffer(data))

		assert.EqualValues(t, http.StatusInternalServerError, res.Code)
	})
}

func TestDelete(t *testing.T) {
	todoUsecase := new(mocks.TodoUsecase)

	t.Run("success", func(t *testing.T) {
		router := gin.Default()
		todoUsecase.On("Delete", mock.Anything).Return(nil).Once()

		controller.NewTodoHandler(router, todoUsecase)
		res := performDeleteRequest(router, "/todo/1")

		assert.EqualValues(t, http.StatusOK, res.Code)
	})

	t.Run("error-delete-failed", func(t *testing.T) {
		router := gin.Default()
		todoUsecase.On("Delete", mock.Anything).Return(UnexpectedErr).Once()

		controller.NewTodoHandler(router, todoUsecase)
		res := performDeleteRequest(router, "/todo/1")

		assert.EqualValues(t, http.StatusInternalServerError, res.Code)
	})

	t.Run("error-invalid-id-failed", func(t *testing.T) {
		router := gin.Default()

		controller.NewTodoHandler(router, todoUsecase)
		res := performDeleteRequest(router, "/todo/invalid_id")

		assert.EqualValues(t, http.StatusInternalServerError, res.Code)
	})
}
