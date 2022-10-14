package main

import (
	"go-jwt/database"
	"go-jwt/router"
	"log"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	err := r.Run(":8080")

	if err != nil {
		log.Fatal(err.Error())		
	}
}