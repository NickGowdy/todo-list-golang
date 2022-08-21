package migrations

import (
	"todo-list-golang/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SeedDatabase() {
	dsn := "host=localhost user=user password=password dbname=todo_list port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Todo{})

	// Create
	db.Create(&models.Todo{Name: "My first todo", Description: "This is the first todo I've created", IsComplete: false})
}
