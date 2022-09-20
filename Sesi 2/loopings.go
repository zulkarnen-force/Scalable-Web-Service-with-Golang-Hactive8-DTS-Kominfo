package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("Angka ", i)
	}

	i := 0

	for i < 5 {
		fmt.Println("Angka ", i)
		i++
	}

	j := 0

	for {
		fmt.Println("Angka ", j)
		j++
		if j == 3 {
			break
		}
	}



	// break and continue
	for i := 1; i <= 10; i++ {
		
		if i % 2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Bilangan Ganjil: ", i)
	}


	// nested loop

	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}


	// Label
	outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("Perulangan ke-", i+1)
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}

}