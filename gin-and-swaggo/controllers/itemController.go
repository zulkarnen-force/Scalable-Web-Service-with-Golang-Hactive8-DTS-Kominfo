package controllers

import (
	"fmt"
	"gin-and-swaggo/database"
	"gin-and-swaggo/models"
)


func CreateItem(desc string, qty int) {

	db := database.GetDB()

	Item := models.Item{
		Quantity: qty,
		Description: desc,
	}

	err := db.Create(&Item).Error

	if err != nil {
		fmt.Println("Error creating user data ", err)
		return
	}

	fmt.Println("New user data", Item)
}
