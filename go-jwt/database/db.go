package database

import (
	"fmt"
	"go-jwt/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	HOST     = "localhost"
	USER     = "postgres"
	PASSWORD = "root"
	PORT     = "5432"
	DB_NAME  = "simple_api"
	db       *gorm.DB
	err		error
)


func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		HOST, USER, PASSWORD, DB_NAME, PORT)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection to database successfully")

	db.Debug().AutoMigrate(models.User{}, models.Product{})

	
}


func GetDB() *gorm.DB {
	return db
}