package main

import (
	"basic-auth/handlers"
	"basic-auth/models"
	"fmt"
	"net/http"
)

func init() {
	models.Students = append(models.Students, &models.Student{Id: "1", Name: "Zulkarnen", Grade: 10})
	models.Students = append(models.Students, &models.Student{Id: "2", Name: "Abdul", Grade: 10})
	models.Students = append(models.Students, &models.Student{Id: "3", Name: "Hanif", Grade: 10})
}


func main() {
	http.HandleFunc("/student", handlers.StudentHandler)

	server := new(http.Server)
	server.Addr = ":9000"

	fmt.Printf("%+v", &models.Students)
	fmt.Println("Server started on http://localhost:9000")
	server.ListenAndServe()

}
