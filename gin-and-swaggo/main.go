package main

import (
	"gin-and-swaggo/routes"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server order.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  zulkarnen1900016072@webmail.uad.ac.id

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth
func main() {
	routes.StartServer().Run(":8080")
}