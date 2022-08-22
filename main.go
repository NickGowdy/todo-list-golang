package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"todo-list-golang/db"
	"todo-list-golang/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	dbUser, dbPassword, dbName :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")
	database, err := db.Initialize(dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	defer database.Conn.Close()

	router := gin.Default()
	router.GET("/todos", get)
	router.GET("/todos/:id", getById)
	router.PUT("/todos/:id", put)
	router.POST("/todos", post)
	router.DELETE("/todos/:id", delete)
	router.Run()
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

func put(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	}

	var updatedTodo models.Todo

	todos := models.GetTodos()
	for _, t := range todos {
		if t.Id == id {
			t.Name = updatedTodo.Name
			t.Description = updatedTodo.Description
			t.IsComplete = updatedTodo.IsComplete
			c.IndentedJSON(http.StatusOK, updatedTodo)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	}

	todos := models.GetTodos()

	for i := len(todos) - 1; i >= 0; i-- {
		if todos[i].Id == id {
			todos = append(todos[:i], todos[i+1:]...)
		}
	}

	c.IndentedJSON(http.StatusOK, todos)

	// if len(todos) < length {
	// 	c.IndentedJSON(http.StatusOK, todos)
	// }

	// c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
}
