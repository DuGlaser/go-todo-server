package http

import (
	"net/http"
	"strconv"

	"github.com/DuGlaser/go-todo-server/domain"
	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	todoUsecase domain.TodoUsecase
}

func NewTodoHandler(c *gin.Engine, u domain.TodoUsecase) {
	handler := &TodoHandler{
		todoUsecase: u,
	}

	c.GET("/todos", handler.GetAll)

	c.POST("/todo", handler.Create)
	c.PUT("/todo", handler.Update)
	c.GET("/todo/:id", handler.GetByID)
	c.DELETE("/todo/:id", handler.Delete)
}

func (h *TodoHandler) GetAll(c *gin.Context) {
	res, restErr := h.todoUsecase.GetAll()
	if restErr != nil {
		c.JSON(http.StatusInternalServerError, restErr)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func (h *TodoHandler) GetByID(c *gin.Context) {
	intID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	res, restErr := h.todoUsecase.GetByID(int64(intID))
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func (h *TodoHandler) Create(c *gin.Context) {
	var todo domain.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := todo.Validate(); err != nil {
		c.JSON(err.Status, err)
		return
	}

	if restErr := h.todoUsecase.Store(&todo); restErr != nil {
		c.JSON(http.StatusInternalServerError, restErr)
		return
	}

	c.JSON(http.StatusCreated, todo)
	return
}

func (h *TodoHandler) Update(c *gin.Context) {
	var todo domain.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := todo.Validate(); err != nil {
		c.JSON(err.Status, err)
		return
	}

	if restErr := h.todoUsecase.Update(&todo); restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusCreated, todo)
	return
}

func (h *TodoHandler) Delete(c *gin.Context) {
	intID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if restErr := h.todoUsecase.Delete(int64(intID)); restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, "Success")
	return
}
