package controllers

import (
	"go-jwt/database"
	"go-jwt/helpers"
	"go-jwt/models"
	"net/http"

	jwt "github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	contentType := helpers.GetContentType(c)

	Product := models.Product{}
	// u64, err := uintstrconv.ParseUint(string(userData["id"].(float)), 10, 32)

	idString := userData["id"].(float64)
	userID := uint(idString)


	if contentType == appJson {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID

	err := db.Debug().Create(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "status bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": Product,
	})
}