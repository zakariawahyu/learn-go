package main

import "fmt"
/**
Sebelumnya setiap membuat funtion, kita akan selalu memberikan sebuah nama funtion tersebut.
Namnun terkadang ada kalanya lebih mudah membuat funtion secara langsung di variabel atau parameter tanpa harus membuat funtion terlebih dahulu.
Hal tersebut dinamakan anonymous function atau funtion tanpa nama
 */

type Blacklist func(string) bool

func registerUsers(name string, blacklist Blacklist)  {
	if blacklist(name) {
		fmt.Println("Your are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main()  {
	// anonymous funtion
	blacklist := func(name string) bool{
		return name == "Admin"
	}

	registerUsers("Admin", blacklist)
	registerUsers("Zakaria", blacklist)
	registerUsers("Rahasia", func(name string) bool {
		return name == "Rahasia"
	})
}
