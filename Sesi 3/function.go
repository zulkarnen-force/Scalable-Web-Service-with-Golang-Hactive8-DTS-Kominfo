package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	var names = []string{"Zul", "Karnen"}
	fmt.Println(greet("Selamat Pagi", names))

	var diameter float64 = 10
	var area, circum = calculate(diameter)
	fmt.Println(area, circum)

	resultPrint := print("Zulkarnen", "abdul rosad", "Zulfa R")
	fmt.Println(resultPrint)

	var numbers []int = []int{1, 2, 3, 4}
	var hasil = sum(numbers...)
	fmt.Println(hasil) // 10

	// favFoods := []string{"Ayam Bakar", "Mie", "Telor", "Satai"} // error, gak bisa pake slice
	profile("Zulkarnen", "Ayam Bakar", "Mie", "Telor", "Satai")
}


func greet(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")
	var result string = fmt.Sprintf("%s %s",msg, joinStr)
	return result
}


// // multiple return
// func calculate(d float64) (float64 float64) {

// 	var area float64 = math.Pi * math.Pow(d/2, 2)
// 	var circumference = math.Pi * d
// 	return area, circumference

// }


// Predefined return value
func calculate(d float64) (area float64, circumference float64) {

	 area  = math.Pi * math.Pow(d/2, 2)
	 
	 circumference = math.Pi * d
	
	 return 

}


// variadic function
// dapat menerima arguments yang tidak terbatas
func print(names ...string) []map[string]string {

	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i)
		temp := map[string]string{
			key:v,
		}

		result = append(result, temp)
	}

	return result

}


// kita juga dapat memasukan slice pada parameter variadic function
func sum(numbers ...int) int {
	var result int

	for _, v := range numbers {
		result += v	
	}

	return result
}


func profile(name string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ", ")

	fmt.Println("Hello there! I'm", name)
	fmt.Println("I really love to eat", mergeFavFoods)
}