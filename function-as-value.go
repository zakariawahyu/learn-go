package main

import "fmt"

func main()  {
	// function value
	// function adalah first class citizen (tidak di anak tirikan)
	// function juga merupakan tipe data dan bisa disimpan didalam variable

	// menyimpan ke variabel
	sayHello := welcomeGuys
	// set nilai
	result := sayHello("Zakaria")
	fmt.Println(result)
}

func welcomeGuys(name string)string  {
	return "Hallo selamat datang " + name
}
