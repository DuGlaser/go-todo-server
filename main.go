package main

import (
	"github.com/DuGlaser/go-todo-server/controller/todo"
	"github.com/DuGlaser/go-todo-server/db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db.TodoDB.Init()
	router.GET("/todos", todo.TodoController.GetAll)
	router.GET("/todos/:id", todo.TodoController.FindByID)
	router.POST("/todos", todo.TodoController.Create)
	router.PUT("/todos", todo.TodoController.Update)
	router.DELETE("/todos/:id", todo.TodoController.Delete)

	router.Run(":8080")
}
