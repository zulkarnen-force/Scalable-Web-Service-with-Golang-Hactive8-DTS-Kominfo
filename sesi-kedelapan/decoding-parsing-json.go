package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Fullname string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {

	var jsonString string = `
		{
			"full_name":"Zulkarnen",
			"email":"zulkarnen@gmail.com",
			"age":21
		}
	`

	var result Employee

	err := json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		fmt.Println("Error decoding: ", err.Error())
		return
	}

	fmt.Println("full_name: ", result.Fullname)
	fmt.Println("email: ", result.Email)
	fmt.Println("age: ", result.Age)
}