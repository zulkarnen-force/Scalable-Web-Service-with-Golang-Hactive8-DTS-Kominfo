package routers

import (
	"project-keenam/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	_ "project-keenam/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartServer(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLFiles("views/template.html")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	var controller = controllers.InDB{
		DB: db,
	}

	router.POST("/employees", controller.CreateEmployee)
	router.GET("/employees", controller.GetAllEmployees)
	router.GET("/employees/views", controller.GetAllEmployeesWithView)
	router.GET("/employees/:id", controller.ShowEmployee)
	router.DELETE("/employees/:id", controller.DeleteEmployee)
	router.PUT("/employees/:id", controller.UpdateEmployee)

	return router
}