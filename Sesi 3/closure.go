package main

import (
	"fmt"
	"strings"
)

func main() {

	var evenNumbers = func(numbers ...int) []int {

		var result []int

		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}

		return result
	}

	var numbers = []int{1, 2, 3, 4, 234, 23, 12, 454, 123, 54}
	fmt.Println(evenNumbers(numbers...))


	// IIFE
	// immediately-invoked function exspresion
	var isPalindrome = func (str string) bool {
		var temp string = ""

		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}

		return temp == str
	}("katak")

	fmt.Println(isPalindrome)


	var students = []string{"Zulkarnen", "Ariel", "Billie", "Hamzah", "Upin"}
	var find = findStudent(students)
	fmt.Println(find("Hamzah"))


	// callbacks

	var findOdd = findOddNumbers(numbers, func(number int) bool {
		return number % 2 == 0
	})

	fmt.Println("Total odd numbers:", findOdd)

}


// Closure sebagai return value
func findStudent(students []string) func (string) string {

	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = i
				break
			}
		}

		if student == "" {
			return fmt.Sprintf("%s does'nt exist!", s)
		}
		
		return fmt.Sprintf("We found %s at position %d", s, position+1)

	}



}

type isOddNum func(int) bool

// Callbacks
func findOddNumbers (numbers []int, callback isOddNum) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}

	return totalOddNumbers
}
