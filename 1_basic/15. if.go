package main

import "fmt"

func main()  {
	name := "Nur"

	if name == "Zakaria" {
		fmt.Println("Iya ini nama saya", name)
	} else if name == "Wahyu" {
		fmt.Println("Iya ini nama saya", name)
	} else if name == "Nur" {
		fmt.Println("Iya ini nama saya", name)
	} else {
		fmt.Println("Hi "+ name +" ... boleh kenalan dulu ga nih?")
	}


	// if short statement
	if lenght := len(name); lenght > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
