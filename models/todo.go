package models

type Todo struct {
	Id         int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Value      string `gorm:"type:varchar(255);" json:"value"`
	IsComplete bool   `gorm:"type:bool;default:0" json:"isComplete"`
}
