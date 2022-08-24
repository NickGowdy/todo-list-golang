package models

type Todo struct {
	Id         int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Value      string `gorm:"type:varchar(255);" json:"value"`
	IsComplete bool   `gorm:"type:bool;default:0" json:"isComplete"`
}

func GetTodos() []Todo {
	todos := []Todo{
		// {Id: 1, Value: "Laundry", IsComplete: false},
		// {Id: 2, Value: "Dog", IsComplete: false},
	}

	return todos
}
