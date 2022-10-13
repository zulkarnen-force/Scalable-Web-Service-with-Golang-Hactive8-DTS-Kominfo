package main

import (
	"practice-build-rest-api/config"
	"practice-build-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	gormDB := &controllers.DB{GormDB: db}

	router := gin.Default()

	router.POST("/person", gormDB.CreatePerson)
	router.GET("/person", gormDB.GetPerson)

	router.Run(":8080")
}