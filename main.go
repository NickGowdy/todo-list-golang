package main

import (
	"net/http"
	"todo-list-golang/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/todos", get)
	router.POST("/todos", post)
	router.Run("localhost:8080")
}

func get(c *gin.Context) {
	todos := models.GetTodos()
	c.IndentedJSON(http.StatusOK, todos)
}

func post(c *gin.Context) {
	var newTodo models.Todo
	var todos = models.GetTodos()

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	_ = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}
