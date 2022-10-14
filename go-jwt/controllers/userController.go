package controllers

import (
	"fmt"
	"go-jwt/database"
	"go-jwt/helpers"
	"go-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJson = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	User := models.User{}

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"bad request",
			"message":err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":User.ID,
		"email":User.Email,
		"full_name":User.FullName,
	})
}



func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	User := models.User{}
	password := ""

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error


	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":"unauthorized",
			"message":"invalid email/password",
		})
		return
	}

	isValidPassword := helpers.ComparePassword([]byte(User.Password), []byte(password))

	if !isValidPassword {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":"unauthorized",
			"message":"invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)
	fmt.Println("Token => ", token)

	c.JSON(http.StatusOK, gin.H{
		"token": token,

	})
}