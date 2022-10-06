package main

import (
	"encoding/json"
	"fmt"
)


func main() {

	var jsonString string = `
		{
			"full_name":"Zulkarnen",
			"email":"zulkarnen@gmail.com",
			"age":21
		}
	`

	var result map[string]interface{}

	err := json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		fmt.Println("Error decoding: ", err.Error())
		return
	}

	fmt.Println("full_name: ", result["full_name"])
	fmt.Println("email: ", result["email"])
	fmt.Println("age: ", result["age"])
}