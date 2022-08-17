package main

import (
	"net/http"
	"todo-list-golang/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/todos", get)
	router.Run("localhost:8080")
}

func get(c *gin.Context) {
	todos := models.GetTodos()
	c.IndentedJSON(http.StatusOK, todos)
}
