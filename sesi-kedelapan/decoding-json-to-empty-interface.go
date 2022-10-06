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

	var temp interface{}

	err := json.Unmarshal([]byte(jsonString), &temp)

	if err != nil {
		fmt.Println("Error decoding: ", err.Error())
		return
	}

	var result = temp.(map[string]interface{})

	fmt.Println("full_name: ", result["full_name"])
	fmt.Println("email: ", result["email"])
	fmt.Println("age: ", result["age"])
}