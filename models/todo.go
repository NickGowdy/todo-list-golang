package models

type todo struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsComplete  bool   `json:"isComplete"`
}

func GetTodos() []todo {
	todos := []todo{
		{Id: 1, Name: "Laundry", Description: "", IsComplete: false},
		{Id: 2, Name: "Dog", Description: "walk the dog", IsComplete: false},
	}

	return todos
}
