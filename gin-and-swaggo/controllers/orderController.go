package controllers

import (
	"fmt"
	"gin-and-swaggo/database"
	"gin-and-swaggo/models"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)


func CreateOrder(ctx *gin.Context) {


	var order models.Order

	db := database.GetDB()

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":http.StatusBadRequest,
			"message":"body request required",
		})

	
		return
	}  

	fmt.Println(order)

	newOrder := models.Order{CustomerName: order.CustomerName, ItemsID: order.ItemsID, OrderAt: time.Now()}

	if err := db.Create(&newOrder).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":http.StatusInternalServerError,
			"message":"internal server error",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message":"successfully created order data",
			"datas":newOrder,
		})
		return
	}


}

// GetOrders godoc
// @Summary      Show an orders
// @Description  get orders data
// @Tags         orders
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Order
// @Router       /orders [get]
func GetOrders(ctx *gin.Context) {

	db := database.GetDB()

	var orders []models.Order

	err := db.Raw("SELECT * FROM orders").Scan(&orders).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"message": "order data not found",
		})
		return
	}

	
	ctx.JSON(http.StatusOK, gin.H{
		"datas":orders,
	})

}
