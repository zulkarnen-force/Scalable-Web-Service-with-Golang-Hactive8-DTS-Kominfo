package main

import "fmt"

func main() {

	const PI = 3.14
	const A = 10
	const B = 5

	fmt.Println(A + B)
	fmt.Println(A + B * 2)
	fmt.Println(A - B)
	fmt.Println(A * B)
	fmt.Println(A / B)
	fmt.Println(A / B)
	fmt.Println(A / B)


	// Relational Operators

	var firstCondition bool =  2 < 3
	var secondCondition bool = "joey" == "Joey"
	var thirdCondition bool = 23 != 2.3
	var fourthCondition bool = 11 >= 11

	fmt.Println("First Condition: ", firstCondition)
	fmt.Println("second Condition: ", secondCondition)
	fmt.Println("third Condition: ", thirdCondition)
	fmt.Println("fourth Condition: ", fourthCondition)


	// Logical Operators
	var right bool = true
	var wrong bool = false

	var rightAndWrong = right && wrong
	fmt.Printf("right && wrong \t (%t)\n", rightAndWrong)

	var rightOrWrong bool = right || wrong
	fmt.Printf("right || wrong \t (%t)\n", rightOrWrong)

	var wrongReverse = !wrong
	fmt.Printf("!wrong \t (%t)\n", wrongReverse)

	

}