package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name string
	Age int
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {

	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe variabel ", reflectValue.Type())

	fmt.Println("tipe variabel ", reflectValue.Int())

	fmt.Println(reflectValue.Kind())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Nilai variabel of reflectValue ", reflectValue.Int())
	}

	// jika hanya untuk keperluan output console, bsa pake Interface
	fmt.Println(reflectValue.Interface())

	// nilai asli nya dapat diakse dengan di casting terlebih dahulu
	var nilai = reflectValue.Interface().(int)
	fmt.Println(nilai)

	var student1 = student{Name: "Zul", Age: 21}
	fmt.Println(student1.Name)

	var reflectStudent = reflect.ValueOf(student1)
	var method = reflectStudent.MethodByName("SetName")

	method.Call([]reflect.Value{
		reflect.ValueOf("John"),
	})

	fmt.Println(student1.Name)
}