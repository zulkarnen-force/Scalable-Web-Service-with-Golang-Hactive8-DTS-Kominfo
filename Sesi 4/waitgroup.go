package main

import (
	"fmt"
	"sync"
)

func printFruits(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("index => %d, fruit => %s, \n", index, fruit)
	wg.Done()
}

func main() {

	fruits := []string{"Apple", "Manggo", "Durian", "Rambutan"}

	var wg sync.WaitGroup

	for index, fruit := range fruits {
		wg.Add(1)
		go printFruits(index, fruit, &wg)
	}

	wg.Wait()

}