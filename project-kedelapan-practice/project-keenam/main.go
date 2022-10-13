package main

import (
	"fmt"
	"project-keenam/routers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"project-keenam/models"
)

// @title          API Employee Hacktiv8 Practice
// @version        1.0
// @description    Just for Practice Go Programming Language
// @termsOfService http://swagger.io/terms/

// @contact.name  Employee API CRUD
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:8080
// @BasePath /

var (
	host = "localhost"
	port = 5432
	user = "postgres"
	dbname = "hacktiv8-chap8"
	db *gorm.DB
	err error
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user,dbname)

	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}


	db.Debug().AutoMigrate(models.Employee{})
	routers.StartServer(db).Run()
}
