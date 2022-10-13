package database

import (
	"fmt"
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
	configDB := fmt.Sprintf("host=%s user=%s password=%s port=%s, dbname=%s sslmode=disable", 
		HOSTNAME, USER, PASSWORD, DBPORT, DBNAME)
	db, err = gorm.Open(postgres.Open(configDB), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting database: ", err)
	}

	// db.Debug().AutoMigrate(&models.Item{}, &models.Order{})

}


func GetDB() *gorm.DB {
	StartDB()
	return db
}