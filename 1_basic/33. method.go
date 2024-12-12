package main

import "fmt"

/**
Struct adalah tipe data seperti tipe data yang lainnya, bisa digunakan sebagai paramaeter function.
Bisa menambahkan method kedalam struct, seakan-akan struct mempunyai function
method adalah function 
 */

type Pelanggan struct {
	Nama, Alamat string
	Umur int
}
// struct method penulisan salam kurung sebelum nama function
func (pelanggan Pelanggan) sayHello() {
	fmt.Println("Hello, My name is", pelanggan.Nama)
}

func (a Pelanggan) sayHuuu(name string) {
	fmt.Println("Huuuu from", a.Nama, "to", name)
}
func main()  {
	// set struct dan pemanggilan struct
	zaka := Pelanggan{Nama: "Zakaria"}
	zaka.sayHello()
	zaka.sayHuuu("Wahyu")
}
