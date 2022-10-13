package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	ID        uint
	FirstName string
	LastName  string
}