package main

import (
	"fmt"
)

func main() {

	var currentYear = 2022

	if age := currentYear - 2000; age > 17 {
		fmt.Println("Kamu bisa daftar SIM")
	} else {
		fmt.Println("Kamu tidak bisa daftar SIM")
	}



	var score = 80

	switch score {
	case 80:
		fmt.Println("Perfecto")
	case 70:
		fmt.Println("Awesome")
	default:
		fmt.Println("not bad")
	}

	// Switch with Relational Conditions

	switch {
	case score == 80:
		fmt.Println("Sempurna")
	case (score < 80) && (score > 30):
		fmt.Println("Tidak Buruk")
	}


	// Switch (Fallthrough keyword)
	// ketika ingin lanjut cek case selanjutnya, meski case sebelumnya terpenuhi

	score = 6
	switch {
	case score == 8:
		fmt.Println("Perfect") 
	case (score < 8) && (score > 3):
		fmt.Println("Not bad")
		fallthrough
	case score < 5:
		fmt.Println("It is okay, but please study harder")
	default:
		fmt.Println("study harder")
		fmt.Println("you don't have a good score yet")
	} 
	

	/*
	Nested Conditions
	*/

	score = 10

	if score > 7 {
		switch score {
		case 7:
			fmt.Println("perfect!")
		default:
			fmt.Println("Nice!")
		}
	} else {
		if score == 5 {
			fmt.Println("Not bad")
			
		} else if score == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
		}
		
		if score == 3 {
			fmt.Println("trying harder!")
		}
	}
}