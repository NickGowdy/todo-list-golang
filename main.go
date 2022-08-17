package main

import (
	"net/http"
	"strconv"
	"todo-list-golang/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/todos", get)
	router.GET("/todos/:id", getById)
	router.POST("/todos", post)
	router.Run("localhost:8080")
}

func get(c *gin.Context) {
	todos := models.GetTodos()
	c.IndentedJSON(http.StatusOK, todos)
}

func getById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	}

	todos := models.GetTodos()
	for _, t := range todos {
		if t.Id == id {
			c.IndentedJSON(http.StatusFound, t)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})

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
