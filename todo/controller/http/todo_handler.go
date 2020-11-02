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
	res, err := h.todoUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
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

	res, err := h.todoUsecase.GetByID(int64(intID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
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
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := h.todoUsecase.Store(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, err)
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
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err := h.todoUsecase.Update(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
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

	err = h.todoUsecase.Delete(int64(intID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "Success")
	return
}
