package main

import "fmt"

func main() {

	var first uint8 = 89
	var second int8 = -12

	fmt.Printf("tipe data first %T\n", first)
	fmt.Printf("tipe data second %T\n", second)


	// float

	var decimalNumber float32 = 3.68

	fmt.Printf("decimal number %f\n", decimalNumber) // decimal number 3.680000
	fmt.Printf("decimal number %.3f\n", decimalNumber) // decimal number 3.680


	var condition bool = true
	fmt.Printf("is it permitted? %t \n", condition)


	// String
	var message string = "Halo"
	var messages string = `Hallo nama saya "Zulkarnen".
		Dan saya juga ada manusia
		`
	fmt.Println(message)
	fmt.Println(messages)


}