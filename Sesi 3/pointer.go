package main

import (
	"fmt"
	"strings"
)

func main() {

	var firstNumber int = 4
	var secondNumber *int = &firstNumber

	fmt.Println("first number (value) ", firstNumber)
	fmt.Println("first number (memori address) ", &firstNumber)

	fmt.Println("second number (value) ", *secondNumber)
	fmt.Println("second number (memori address) ", secondNumber)


	// changing value through pointer
	var firstPerson string = "Zul" 
	var secondPerson *string = &firstPerson
	
	fmt.Println("firstPerson (value) : ", firstPerson)
	fmt.Println("firstPerson (memori address) : ", &firstPerson)
	fmt.Println("secondPerson (value) : ", *secondPerson)
	fmt.Println("secondPerson (value) : ", secondPerson)

	fmt.Println("\n", strings.Repeat("#", 30), "\n")

	// changing
	*secondPerson = "Doe"

	fmt.Println("firstPerson (value) : ", firstPerson)
	fmt.Println("firstPerson (memori address) : ", &firstPerson)
	fmt.Println("secondPerson (value) : ", *secondPerson)
	fmt.Println("secondPerson (value) : ", secondPerson)


	// Pointer as a parameter
	var a int = 10

	fmt.Println("Before ", a)
	changeValue(&a)
	fmt.Println("After ", a)
	
}

func changeValue(number *int) {
	*number = 20
}