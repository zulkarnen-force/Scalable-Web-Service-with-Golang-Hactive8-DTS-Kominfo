package main

import "fmt"

func main() {
	
	var name string = "Zulkarnen"
	var age int = 21;

	fmt.Println("My name is: ", name)
	fmt.Println("And my age is: ", age)

	var jenisKelamin string
	
	jenisKelamin = "pria"
	age = age + 1

	fmt.Println("Dan saya adalah: ", jenisKelamin)
	fmt.Println("And my age is, dan tahun depan akan bertambah menjadi: ", age)

	// variabel without data type
	var newName = "Sambo"
	var newAge = 40


	fmt.Printf("%T, %T \n", newName, newAge); // string, int

	/**
		short declaration
		tanpa keyword var, wajib isi value dari nama variabel.
	**/

	nama := "Zulkarnen"
	umur := 21

	fmt.Printf("%T, %T \n", nama, umur) // string, int

	/**
	Multiple declaration
	**/

	var student1, student2, student3 string = "satu", "dua", "tiga"

	var first, second, third int

	first, second, third = 1, 2, 3

	fmt.Println(student1, student2, student3)
	fmt.Println(first, second, third)

	/**
		Multiple declaration with without data type and using shord declaration
	**/

	var mahasiswa1, semester, angkatan = "Zulkarnen", 7, 2019
	satu, dua, tiga := 1, 2, 3

	fmt.Println(mahasiswa1, semester, angkatan)
	fmt.Println(satu, dua, tiga)


	/**
		Underscore variabel
	*/

	nim, ipk := 1900016072, 3.68
	_, _ = nim, ipk


	// fmt.Printf() usage
	nama = "Zulkarnen"
	umur = 22
	fmt.Printf("Nama saya adalah %s dan umur saya adalah %d", nama, umur)

}