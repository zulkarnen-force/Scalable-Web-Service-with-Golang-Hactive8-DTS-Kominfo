package main

import "belajar-go/helpers"

func main() {

	helpers.Greet()
	// helpers.Greet()

	var person = helpers.Person{
		Name: "Zulkarnen",
		Age: 21,
		Address: "Majalengka",
	}
	person.Invokedgreet()

}