package todo

import (
	"net/http"
	"strconv"

	"github.com/DuGlaser/go-todo-server/domain/todo"
	service "github.com/DuGlaser/go-todo-server/service/todo"
	"github.com/gin-gonic/gin"
)

var TodoController todoControllerInterface = &todoController{}

type todoControllerInterface interface {
	GetAll(*gin.Context)
	FindByID(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type todoController struct{}

func (controller *todoController) GetAll(c *gin.Context) {
	result, err := service.TodoService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
	return
}

func (controller *todoController) FindByID(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	result, err := service.TodoService.GetByID(int64(intID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
	return
}

func (controller *todoController) Create(c *gin.Context) {
	var todo todo.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	result, err := service.TodoService.Create(todo.Title, todo.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
	return
}

func (controller *todoController) Update(c *gin.Context) {
	var todo todo.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	result, err := service.TodoService.Update(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
	return
}

func (controller *todoController) Delete(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = service.TodoService.Delete(int64(intID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
	return
}
