package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Id          int    `gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsComplete  bool   `json:"isComplete"`
}

func GetTodos() []Todo {
	todos := []Todo{
		{Id: 1, Name: "Laundry", Description: "", IsComplete: false},
		{Id: 2, Name: "Dog", Description: "walk the dog", IsComplete: false},
	}

	return todos
}
