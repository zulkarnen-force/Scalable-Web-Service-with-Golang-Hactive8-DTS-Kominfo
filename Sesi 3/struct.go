package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	name     string
	age      int
	division string
}


type Person struct {
	name string 
	age int
}

type Mahasiswa struct {
	nim int
	person Person
}

func main() {

	var employee Employee

	employee.name = "Zulkarnen"
	employee.age = 21
	employee.division = "IT"

	fmt.Println(employee.name, employee.age, employee.division)


	var employee1 = Employee{}
	employee1.name = "Airell" 
	employee1.age = 23
	employee1.division = "Curiculum Developer"

	var employee2 = Employee{name: "Abdul Gofar", age: 24, division: "Public Relation"}

	fmt.Printf("Empoloyee1 %+v \n", employee1)
	fmt.Printf("Empoloyee2 %+v \n", employee2)



	// pointer to a struct

	var employee3 = Employee{name: "Bille", age: 40, division: "Support"}
	var employee4 *Employee = &employee3
	// fmt.Println(employee4) // &{Bille 40 Support}
	fmt.Println("Employee3 name ", employee3.name)
	fmt.Println("Employee4 name ", employee4.name)
	
	
	fmt.Println(strings.Repeat("#", 20))
	
	employee4.name = "Nanda"

	fmt.Println("Employee3 name ", employee3.name)
	fmt.Println("Employee4 name ", employee4.name)


	var zulkarnen Mahasiswa = Mahasiswa{}
	zulkarnen.nim = 1900016072
	zulkarnen.person.age = 21
	zulkarnen.person.name = "Zulkarnen"

	fmt.Printf("%+v \n", zulkarnen) // {nim:1900016072 person:{name:Zulkarnen age:21}}

	// anonymouse struct tanpa isi field
	var mhs = struct {
		person Person 
		nim int
	} {}
	mhs.person = Person{name: "Ahmad", age: 22}
	mhs.nim = 190001698


	// anonymouse struct with isi field
	var mhs1 = struct {
		person Person
		nim int
	} {
		person: Person{name: "Berliana", age: 21},
		nim: 1900016045,
	}

	
	fmt.Printf("mhs %+v \n", mhs)
	fmt.Printf("mhs1 %+v \n", mhs1)


	// slice of struct
	var person = []Person{
		{name: "Zulkarnen", age: 21},
		{name: "Dimas", age: 22},
		{name: "Adi", age: 20},
	}

	for _, v := range person {
		fmt.Printf("%+v \n", v)
	}


	var pekerjaKeras = []struct {
		person Person
		id int
	} {
		{person: Person{name: "Hafiz", age: 21}, id: 13213112},
		{person: Person{name: "Dinda", age: 22}, id: 13213114},
		{person: Person{name: "Syahrul", age: 23}, id: 13213113},
	}

	for _, v := range pekerjaKeras {
		fmt.Printf("%+v \n", v)
	}

}
