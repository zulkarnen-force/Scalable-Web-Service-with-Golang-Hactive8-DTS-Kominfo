package controllers

import (
	"fmt"
	"net/http"
	"practice-build-rest-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DB struct {
	GormDB *gorm.DB
}

func (db *DB) CreatePerson(ctx *gin.Context) {

	var (
		person models.Person
		result gin.H
	)

	firstName := ctx.PostForm("first_name")
	lastName := ctx.PostForm("last_name")

	person = models.Person{FirstName: firstName, LastName: lastName}

	db.GormDB.Create(&person)

	result = gin.H{
		"result": person,
	}

	ctx.JSON(http.StatusOK, result)
}

type Person = models.Person

func (db *DB) GetPerson(c *gin.Context) {
	var results []Person

	db.GormDB.Raw("SELECT * FROM people").Scan(&results)

	fmt.Println(results)

	c.JSON(200, gin.H{
		"results": results,
	})
}