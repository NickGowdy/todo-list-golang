package main

import (
	"log"
	"net/http"
	"strconv"
	"todo-list-golang/db"
	"todo-list-golang/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	db, err := db.Initialize()
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	defer db.DB()

	router := gin.Default()
	router.GET("/todos", get)
	router.GET("/todos/:id", getById)
	router.PUT("/todos/:id", put)
	router.POST("/todos", post)
	router.DELETE("/todos/:id", delete)
	router.Run()
}

func get(c *gin.Context) {
	db, err := db.Initialize()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	todos := db.Find(&models.Todo{})
	c.IndentedJSON(http.StatusOK, todos)
}

func getById(c *gin.Context) {
	var todo models.Todo

	db, err := db.Initialize()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	if err := db.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
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
