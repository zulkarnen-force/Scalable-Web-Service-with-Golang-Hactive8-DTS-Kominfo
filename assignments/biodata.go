package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	id int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}


func contains(persons []Person, id int) (bool, Person) {
	for _, person := range persons {
		if person.id == id {
			return true, person
		}
	}
	return false, Person{} 
}

func main() {

	arg, err := strconv.Atoi(os.Args[1])
	
	_ = err

	var persons = []Person{
		{id: 1, Nama: "Zulkarnen", Alamat: "Majalengka", Pekerjaan: "Mahasiswa", Alasan: "Mengembangkan minat, bakat dan mengasah kemampuan Go-lang"},
		{id: 2, Nama: "Abdul", Alamat: "Surabaya", Pekerjaan: "Junior Backend", Alasan: "Menambah pengalaman"},
		{id: 3, Nama: "Berlin", Alamat: "Karawang", Pekerjaan: "Sistem Analis", Alasan: "Belajar Go"},
		{id: 4, Nama: "Nanda", Alamat: "Padang", Pekerjaan: "Data Science", Alasan: "Ingin tahu tentang Go"},
		{id: 5, Nama: "Santi", Alamat: "Blitar", Pekerjaan: "UI UX", Alasan: "Mencoba Go"},
		{id: 6, Nama: "Savira", Alamat: "Pekalongan", Pekerjaan: "Finance", Alasan: "Mengisi waktu luang"},
		{id: 7, Nama: "Kusnadi", Alamat: "Tegal", Pekerjaan: "Administratif", Alasan: "Mengasah skill"},
		{id: 8, Nama: "Gugun", Alamat: "Kebumen", Pekerjaan: "IT Support", Alasan: "Menambah hard skill"},
		{id: 9, Nama: "Rifki", Alamat: "Pemalang", Pekerjaan: "Adm Network", Alasan: "Belajar lanjut tentang go"},
		{id: 10, Nama: "Ega Ramdhani", Alamat: "Magelang", Pekerjaan: "Front end", Alasan: "Membuat web service menggunakan go"},
	}


	var ok, value = contains(persons, arg)
	if ok {
		fmt.Println(value.Nama, value.Alamat, value.Pekerjaan, value.Alasan)
	} else {
		fmt.Printf("Data dari %d tidak dapat ditemukan \n", arg)
	}
	
}
