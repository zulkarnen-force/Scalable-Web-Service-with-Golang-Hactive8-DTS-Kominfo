package main

import (
	"fmt"
	"time"
)

func introduce(name string, c chan string) {
	var result string = fmt.Sprintf("Hello nama saya %s \n", name)

	c <- result
}

func print(c chan string) {
	fmt.Println(<-c)
}

func main() {

	// c := make(chan string)

	// go introduce("Zul", c)
	// go introduce("Aul", c)
	// go introduce("San", c)

	// msg1 := <- c
	// fmt.Println(msg1)

	// msg2 := <- c
	// fmt.Println(msg2)
	
	// msg3 := <- c
	// fmt.Println(msg3)

	// close(c)


	//// Channel with anonymouse function
	// channel := make(chan string) 

	// students := []string{"Ariel", "Mailo", "Indah"}

	// for _, v := range students {
	// 	go func(student string) {
	// 		fmt.Println("Student ", student)
	// 		result := fmt.Sprintf("Hai my name is %s", student)
	// 		channel <- result
	// 	}(v)
	// }

	// for i := 1; i <= 3; i++ {
	// 	print(channel)
	// }

	// close(channel)



	/// Select
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		 
		c1 <- "Hello"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		 
		c2 <- "Salut"
	}()


	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <- c1:
			fmt.Println("Received ", msg1)
		case msg2 := <- c2:
			fmt.Println("Received ", msg2)
		}
	}
	
}
