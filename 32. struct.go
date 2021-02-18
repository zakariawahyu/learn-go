package main

import "fmt"

/**
Struck adalah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
Struct menggabungkan beberapa tipe data, jika array kan harus 1 tipe data yang sama dan kalau map key dan value tipe datanya juga harus sama
Struct biasanya representasi data dalam program apliaksi yang kita buat
Data struck disimpan dalam field
Sederhananya struct adalah kumpulan field
 */


/**
Struct adalah template data atau prototype
Struct tidak bisa langsung digunakan
Namun kita bisa membuat data/object dari struct yang telah kita buat
 */

// Biasanya penulisan nama struct menggunakan uppercase to lower
// Atau diawali dengan huruf besar terlebih dahulu
type Customer struct {
	Name, Adress string
	Age int
}

func main()  {
	var data Customer
	data.Name = "Zakaria Wahyu"
	data.Adress = "Bandung"
	data.Age = 21

	// print semua data
	fmt.Println(data)

	// print salah satu data
	fmt.Println(data.Name)
	fmt.Println(data.Adress)
	fmt.Println(data.Age)

	// Struct literal atau contoh penggunaan struct yang lain
	cus1 := Customer{
		Name:   "Nur Utomo",
		Adress: "Klaten",
		Age:    21,
	}
	fmt.Println(cus1)

	// Jangan sering menggunakan yang seperti ini, simpel sih tapi bahayanya
	// bahanya jika urutan field struct berubah maka ini akan error dan harus disesuaikan urutanya juga
	cus2 := Customer{"Utomo", "Jabar", 21}
	fmt.Println(cus2)
}
