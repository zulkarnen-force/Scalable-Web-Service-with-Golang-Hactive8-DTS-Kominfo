package main

import "fmt"

func main() {

	var a uint8 = 10
	var b byte

	a = b

	_ = a

	type second = byte

	var threeSecond second = 3
	fmt.Println(threeSecond)

}