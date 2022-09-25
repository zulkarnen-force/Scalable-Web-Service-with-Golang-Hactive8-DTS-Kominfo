package main

import (
	"fmt"
	"strings"
)

func main() {
	/**
	slice seperti array, 
	tidak punya fix length,
	reference type
	*/
	var slice = []string{"satu", "dua", "tiga"}
	fmt.Println(slice)

	var fruits = make([]string, 3) // make(tipe slice, rangeSlice)
	
	fmt.Printf("%#v \n", fruits)

	// fruits = append(fruits, "Banana", "Apple", "Watermelon")


	fmt.Printf("%#v \n", fruits)


	// add element to slice
	// using index
	fruits[0] = "Watermellon"
	fruits[1] = "mango"
	fruits[2] = "Apple"

	fmt.Printf("%#v \n", fruits)


	// using append function
	fruits = append(fruits, "grape", "Kiwi")

	fmt.Printf("%#v \n", fruits)


	// Append with ellips

	fruits1 := []string{"oranges", "grapefruits", "limes"}
	berries := []string{"strawberries", "raspberries", "blueberries"}

	// fruits1 = append(fruits1, berries...)
	// fmt.Printf("%#v \n", fruits1)


	// copy slice
	nn := copy(fruits1, berries);
	fmt.Println(fruits1)
	fmt.Println(berries)
	fmt.Println("Copied ELements =>", nn)


	// Slicing
	// cara lain mendapatkan element slice
	var buahs []string = []string{"apel", "mangga", "durian", "semangka", "manggis"}
	fmt.Println(buahs[1:3]) // [mangga durian]
	fmt.Println(buahs[:4]) // [apel mangga durian semangka]
	fmt.Println(buahs[2:]) // [durian semangka manggis]
	fmt.Println(buahs[:]) // sama seperti buahs[:len(buahs)]


	// Slicing and Append
	sayur := []string{"bayam", "kangkung", "sawi", "kol", "wortel"}
	sayur = append(sayur[:3], "brokoli")
	fmt.Println(sayur) // [bayam kangkung sawi brokoli]


	// Backing array
	// Pada saat buat slice, go otomatis buat array tersembunyi (backing array)
	// untuk menyimpan element slice. Go implement slice sebagai struktur data disebut slice header
	// slice header terdiri dari: alamat memori backing array, panjang slice, dan kapasistas slice
	
	var hewan = []string{"kambing", "sapi", "kuda", "lumba-lumba", "jerapah"}
	var hewan1 = hewan[:3]

	hewan1[0] = "Babi"

	fmt.Println("Hewan => ", hewan)
	fmt.Println("Hewan 1 => ", hewan1)

	// ketika slicing, dia tidak buat backing array baru, pake backing array yang telah dibuat
	// jadi misal ada perubahan, maka perubahan nya akan sama dimana pun slice di simpan



	// CAP FUNCTION
	// Untuk mengetahui kapasitas slice

	var kota = []string{"bandung", "jakarta", "surabaya", "jogja"}
	fmt.Println(cap(kota)) // 4
	fmt.Println(len(kota)) // 4

	fmt.Println(strings.Repeat("#", 20))

	var kota1 = kota[0:3]
	fmt.Println(kota1)
	fmt.Println(cap(kota1)) // 4
	fmt.Println(len(kota1)) // 3

	var kota2 = kota[1:]
	fmt.Println(kota2)
	fmt.Println(cap(kota2)) // 3
	fmt.Println(len(kota2)) // 3



	// create new backing array dapat dengan menggunakan append function
	// dengan membuat backing array baru, element dari slice lain tidak akan ada perubahan jika ada modif
	var cars = []string{"Honda", "BMW", "Mercy", "Range Rover"}
	var cars1 = []string{}

	cars1 = append(cars1, cars[0:2]...)

	cars[0] = "Nissan"

	fmt.Println(cars)
	fmt.Println(cars1)

}