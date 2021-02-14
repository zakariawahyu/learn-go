package main

import "fmt"
// Past by value
/**
Secara default di go-lang semua variable itu di passing by value bukan by reference.
Artinya, jika kita mengirim sebuah value ke dalam function, method atau variabel lain, sebenarnya yang dikirim itu duplicate dari value aslinya
 */

// Pointer
/**
Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data value utama yang sudah ada.
Sederhananya, dengan pointer kita bisa membuat pass by reference karena di golang sendiri secara default pass by value atau duplikasi.

Operator &
Untuk membuat sebuah variabel dengan pointer ke variabel lain, kita bisa gunakan operator & dan diikuti nama variabelnya.

Operator bintang *
Saat kita merubah variabel  pointer, maka yang berubah hanya variabel tersebut
Semua variabel yang mengacu ke data yang sama tidak akan pernah berubah
Jika ingin mengubah seluruh variabel yang mengacu ke data tersebut, kita bisa menggunakan operator *

Membuat pointer dengan data kosong
Golang memiliki function new yang bisa digunakan untuk membuat pointer
Namun function new hanya mengembalikan pointer ke data kosong
 */
type Address struct {
	Kota, Propinsi, Negara string
	KodePos int
}

func main()  {
	// Contoh pass by value
	address1 := Address{"Klaten", "Jawa Tengah", "Indonesia", 57465}
	address2 := address1
	// ubah data
	address2.Kota = "Boyolali"
	address2.KodePos = 1231223
	// Address 1 bentuk awal
	fmt.Println("Alamat 1 :", address1)
	// Address 2 hasil duplicate address 1
	// Ketika address 2 diubah, maka tidak akan pengaruh di address 1 karena address 2 merupakan duplicate addres 1
	fmt.Println("Alaat 2 :", address2)


	// Contoh pointer atau pass by reference dengan operator &
	address3 := Address{"Solo", "Jawa Tengah", "Indonesia", 123445}
	// pointer dengan menambahkan & pada depan nama variabel
	address4 := &address3
	// ubah data
	address4.Kota = "Bandung"
	address4.Propinsi = "Jawa Barat"
	// karena menggunakan pointer jika adress 4 diubah maka addres 3 juga ikut ke ubah datanya
	fmt.Println("Alamat 3 :", address3)
	fmt.Println("Alamat 4 :", address4)

	// Contoh pointer atau pass by reference dengan paramater *
	address5 := Address{"Jakarta selatan", "DKI Jakarta", "Indonesia", 576354}
	address6 := &address5
	// membuat data struct baru
	// dengan operator * memaksa untuk mengikuti data utama yang baru dibuat
	*address6 = Address{"Surabaya", "Jawa Timur", "Indonesia", 1234}
	// addres 5 dipaksa mengikuti addres 6
	fmt.Println("Alamat 5 :", address5)
	fmt.Println("Alamat 6 :", address6)

	// Function new atau pointer kosong
	alamat1 := new(Address)
	alamat1.Kota = "Bandung"
	fmt.Println("Alamat kosong :", alamat1)

}
