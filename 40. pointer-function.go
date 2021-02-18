package main

import "fmt"
/**
Pointer di function
Saat kita membuat parameter di funtion , secara default adalah pass by value. artinya data akan di copy lalu dikirim ke funtion tersebut.
Oleh karena itu, jika kita mengubah data di dalam function maka data asli tidak akan pernah berubah
Hal ini membuat variabel menjadi aman karena tidak bida diubah
Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut, bisa menggunakan pointer di funtionnya
Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator * dalam parameternya
 */

// membuat data struct terlebih dahulu
type KartuPengenal struct {
	Nama, Alamat string
	Umur int
}

// function yang parameter terdapat pointer, diawali dengan *
func changeKartuPengenal(ktp *KartuPengenal)  {
	 ktp.Nama = "Zakaria"
}

func main()  {
	// data asli
	data := KartuPengenal{
		Nama:   "",
		Alamat: "Bandung",
		Umur:   21,
	}
	fmt.Println("Data asli : ",data)

	// pasing pointer
	changeKartuPengenal(&data)
	fmt.Println("Data diubah dengan pointer", data)
}
