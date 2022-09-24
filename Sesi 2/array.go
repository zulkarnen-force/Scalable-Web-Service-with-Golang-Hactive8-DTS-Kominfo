package main

import (
	"fmt"
	"strings"
)

func main() {

	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}
	fmt.Println(numbers)

	var names = [3]string{"Upin", "Ipin", "Mail"}
	fmt.Println(names)

	// melihat panjang array
	fmt.Printf("%#v \n", names)

	names[0] = "Zulkarnen"
	fmt.Println(names)


	// Looping Array

	var fruits = [3]string{"Banana", "Apple", "Manggo"}
	for i, v := range fruits {
		fmt.Printf("Index %d, Value %s \n", i, v)
	}


	fmt.Println(strings.Repeat("#", 25))

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Index %d with Value %s \n", i, fruits[i])
	}


	// multidimensional array
	var balances = [2][3]int{{5,6,7}, {8,9,10}}
	for halo, arr := range balances {
		fmt.Print("Halo ", halo)
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

}