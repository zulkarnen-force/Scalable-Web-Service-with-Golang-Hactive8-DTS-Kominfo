package config

import (
	"fmt"
	"practice-build-rest-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
  
var (
	db *gorm.DB
	err error
)

func DBInit() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/learn_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})


	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connection MYSQL")
	}

	db.Debug().AutoMigrate(&models.Person{})
	
	return db

}
