package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	godotenv.Load(".env")
	db = Setup()
	router := SetupRouter()
	router.Run()
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/todos", get)
	router.GET("/todos/:id", getById)
	router.PUT("/todos/:id", put)
	router.POST("/todos", post)
	router.DELETE("/todos/:id", delete)
	return router
}

func get(c *gin.Context) {
	var todos []Todo
	db.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func getById(c *gin.Context) {
	var todo Todo
	if err := db.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func post(c *gin.Context) {
	var newTodo Todo
	if err := c.BindJSON(&newTodo); err != nil {
		log.Fatal(err)
		return
	}

	todo := Todo{Value: newTodo.Value, IsComplete: newTodo.IsComplete}
	db.Create(&todo)
	c.JSON(http.StatusOK, todo)
}

func put(c *gin.Context) {
	var todo Todo
	if err := db.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var updatedTodo Todo
	if err := c.BindJSON(&updatedTodo); err != nil {
		log.Fatal(err)
		return
	}

	db.Updates(updatedTodo)
	c.JSON(http.StatusOK, updatedTodo)
}

func delete(c *gin.Context) {
	var todo Todo
	if err := db.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(todo)
	c.JSON(http.StatusOK, true)
}
