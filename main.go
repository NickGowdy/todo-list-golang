package main

import (
	"log"
	"net/http"
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
	var todos []models.Todo
	db.Find(&todos)

	c.JSON(http.StatusOK, gin.H{"data": todos})
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

	db, err := db.Initialize()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	var newTodo models.Todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todo := models.Todo{Value: newTodo.Value, IsComplete: newTodo.IsComplete}
	db.Create(&todo)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func put(c *gin.Context) {
	db, err := db.Initialize()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	var todo models.Todo
	if err := db.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var updatedTodo models.Todo

	if err := c.BindJSON(&updatedTodo); err != nil {
		return
	}

	db.Updates(updatedTodo)

	c.JSON(http.StatusOK, gin.H{"data": updatedTodo})
}

func delete(c *gin.Context) {
	db, err := db.Initialize()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	var todo models.Todo
	if err := db.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(todo)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
