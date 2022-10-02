package main

import "fmt"

func main() {
	var randomValue interface{}

	randomValue = 100

	randomValue = "String"

	randomValue = 3.14

	randomValue = []int{1, 2, 3, 4}

	// _ = randomValue
	fmt.Println(randomValue)

	var v interface{}

	v = 9

	// v = v * 9 // perkalian hanya bisa diopersikan oleh nilai kongkrit, tidak bisa empty interface
	
	if value, ok := v.(int); ok == true {
		v = value * 9
	}


	rs := []interface{}{1, "Airlee", true, 2, "Ananda", false}

	rm := map[string]interface{}{
		"Nama":"Zulkarnen",
		"Status": true,
		"Age": 19,
	}

	_, _ = rs, rm
}