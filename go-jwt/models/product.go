package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GormModel
	Title string    `gorm:"not null" json:"title" valid:"required~Your Title is required"`
	Description    string `json:"description" form:"description" valid:"required~Your of your product Description is required"`
	UserID uint    
	User *User
}

func (p Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil 
	return
}


func (p Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil 
	return
}