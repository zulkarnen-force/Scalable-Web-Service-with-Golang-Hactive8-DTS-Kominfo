package database

import (
	"fmt"
	"learning-gorm/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	HOSTNAME = "localhost"
	USER     = "postgres"
	PASSWORD = "root"
	DBPORT   = "5432"
	DBNAME   = "learning_gorm"

	db *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s port=%s, dbname=%s sslmode=disable", 
			HOSTNAME, USER, PASSWORD, DBPORT, DBNAME)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting database:", err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Product{})
}


func GetDB() *gorm.DB {
	return db
}

