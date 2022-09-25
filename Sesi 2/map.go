package main

import "fmt"

func main() {

	// var person map[string]string // declare map

	// person = map[string]string{} // initial

	// person["name"] = "Zulkarnen"
	// person["nim"] = "1900016072"
	// person["city"] = "Majalengka"

	// fmt.Println(person["name"])
	// fmt.Println(person["nim"])
	// fmt.Println(person["city"])

	// var person map[string]string = map[string]string{
	// 	"name": "zulkarnen",
	// 	"nim":  "1900016072",
	// 	"city": "Majalengka",
	// }

	// for k, v := range person {
	// 	fmt.Println(k, "=>", v)
	// }

	// // Delete Map
	// fmt.Println("Before delete name ", person)
	// delete(person, "name")
	// fmt.Println("After delete name ", person)

	// Detect element on map

	var person map[string]string = map[string]string{
		"name": "zulkarnen",
		"nim":  "1900016072",
		"city": "Majalengka",
	}

	value, exists := person["name"]

	if exists {
		fmt.Println("Name is exists on person with value ", value)
	} else {
		fmt.Println("Sorry, no have prop name on person")
	}

	delete(person, "name")

	value, exists = person["name"]

	if exists {
		fmt.Println("Name is exists on person with value ", value)
	} else {
		fmt.Println("Sorry, no have prop name on person")
	}


	// slice with map

	var people = []map[string]string{
		{"name":"Zulkarnen", "age":"90"},
		{"name":"Abdul", "age":"50"},
		{"name":"Khadir", "age":"80"},
	}

	for i, person := range people {
		fmt.Printf("index : %d, name %s, age %s \n", i, person["name"], person["age"])
	}
}