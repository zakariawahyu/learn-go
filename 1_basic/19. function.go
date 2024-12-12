package main

import "fmt"
/**
Function adalah sebuah blok kode yang sengaja dibuat dalam program agar bisa digunakan berulang-ulang.
Cara membuat function sangat sederhana, hanya dengan membuat kata kunci func lalu diikuti dengan nama functionnya dan blok kode isi functionnya.
Setelah membuat function, kita bisa mengeksekusi fucntion tersebut dengan memanggilnya menggunakan kata kunci nama functionnnya diikuti tanda kurung buka dan tutup ()
 */
func main()  {
	sayHello()
	sayHello()
	sayHello()
	for i := 0; i <= 10; i++ {
		sayHello()
	}
}

func sayHello()  {
	fmt.Println("Hello World")
}
