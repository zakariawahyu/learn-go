package main

import "fmt"

/**
Type assertions merupakan kemampuan merubah tipe data menjadi tipe data yang di inginkan.
Fitur ini seering sekali digunakan ketika kita bertemu dengan interface kosong, kalau tidak namanya konversio tipe data.
 */

/**
Saat salah menggunakan type assertions maka akan muncul panic di aplikasi kita.
Jika panic tidak di cover, maka otomatis program kita akan berhenti.
Agar lebih aman, sebaiknya kita menggunakan switch expression untuk menggunakan type assertions
 */
func random() interface{}  {
	return "Ups"
}

func main()  {
	result  := random()
	//resultString := result.(int) // int nanti panic
	//fmt.Println(resultString)

	// Type assertions dengan switch untuk meminimalisir jika terjadi panic
	switch value := result.(type) {
		case string:
			fmt.Println("Value :", value, "is string")
		case int:
			fmt.Println("Value :", value, "is integer")
		default:
			fmt.Println("Unknown type")
	}
}
