package main

import (
	"github.com/DuGlaser/go-todo-server/db"
	"github.com/DuGlaser/go-todo-server/todo/controller/http"
	"github.com/DuGlaser/go-todo-server/todo/repository/mysql"
	"github.com/DuGlaser/go-todo-server/todo/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	todoRepo := mysql.NewMysqlTodoRepository(db.NewDB())
	todoUsecase := usecase.NewTodoUsecase(todoRepo)
	http.NewTodoHandler(router, todoUsecase)

	router.Run(":8080")
}
