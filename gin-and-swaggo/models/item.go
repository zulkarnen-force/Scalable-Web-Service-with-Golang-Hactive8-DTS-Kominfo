package models

type Item struct {
	ID          int `gorm:"primary key"`
	Description string
	Quantity    int
}