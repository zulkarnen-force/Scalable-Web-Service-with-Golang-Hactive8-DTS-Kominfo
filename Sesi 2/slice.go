package main

import "fmt"

func main() {
	/**
	slice seperti array, 
	tidak punya fix length,
	reference type
	*/
	var slice = []string{"satu", "dua", "tiga"}
	fmt.Println(slice)

	var fruits = make([]string, 3) // make(tipe slice, rangeSlice)
	
	fmt.Printf("%#v \n", fruits)

	// fruits = append(fruits, "Banana", "Apple", "Watermelon")


	fmt.Printf("%#v \n", fruits)


	// add element to slice
	// using index
	fruits[0] = "Watermellon"
	fruits[1] = "mango"
	fruits[2] = "Apple"

	fmt.Printf("%#v \n", fruits)


	// using append function
	fruits = append(fruits, "grape", "Kiwi")

	fmt.Printf("%#v \n", fruits)


	// Append with ellips

	fruits1 := []string{"oranges", "grapefruits", "limes"}
	berries := []string{"strawberries", "raspberries", "blueberries"}

	fruits1 = append(fruits1, berries...)
	fmt.Printf("%#v \n", fruits1)

}