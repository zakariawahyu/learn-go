package main

// menambahkan _ jika tidak ada function yang di panggil dalam import package
import _"learn-go/database"
/**
Package initialization
Saat kita membuat package, kita bisa membuat sebuah function yang akan diakses ketika package kita diakses.
Ini sangat cocok ketika contohnya, jika package kita berisi funtion-funtion untuk berkomunikasi dengan database, kita membuat function iniliasisasinya untuk membuka koneksi ke database.
Untuk membuat sebuah funtion yang diakses otomatis ketika package diakses, kita cukup membuat function dengan nama init.
 */

/**
Blank identifier
Kadang kita hanya ingin menjalankan init function di package tanpa harus mengeksekusi salah satu funtion yang berada di package tersebut.
Secara default, Golang akan komplain ketika ada package yang di import namun tidak digunakan.
Untuk menamngani hal tersebut , kita bisa menggunakan blank identifier (_) sebelum namam package ketika melakukan import
 */
func main()  {
	// contoh menjalankan init dengan memanggil function maka tidak akan error
	//result := database.GetDatabase()
	//fmt.Println(result)

	// jika kita tidak memanggil funtion dari package yang di import maka akan error
	// mengatasinya menambahkan _
	//import _"learn-go/database"
}
