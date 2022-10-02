package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	widht, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rectangle) area() float64 {
	return r.height * r.widht
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.widht)
}

func (c circle) volume() float64 {
	return 4 / 3 + math.Pi * math.Pow(c.radius, 3)
}

func print(t string, s shape) {
	fmt.Printf("%s area %v \n", t, s.area())
	fmt.Printf("%s perimeter %v \n", t, s.perimeter())
}

func main() {

	var c1 shape = circle{radius: 5}
	var r1 shape = rectangle{widht: 3, height: 2}

	fmt.Printf("Type of c1: %T \n", c1)
	fmt.Printf("Type of r1: %T \n", r1)

	fmt.Println("Circle area: ", c1.area())
	fmt.Println("Circle perimeter: ", c1.perimeter())

	fmt.Println("Rectangle area: ", r1.area())
	fmt.Println("Rectangle perimeter: ", r1.perimeter())

	print("Circle", c1)
	print("Rectangle", r1)

	// c1.volume() // tidak bisa karena variabel c1 telah jadi type of shape
	// agar bisa balik lagi ke tipe data aslinya, maka perlu 
	// type assertion
	c1.(circle).volume() 

	var value, ok = c1.(circle)

	if ok == true {
		fmt.Printf("Circle value: %+v \n", value)
		fmt.Printf("Circle volume: %f \n", value.volume())
	}
}