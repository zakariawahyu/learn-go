package main

import "fmt"

func main()  {
	name := "a"

	switch name {
	case "Zakaria":
		fmt.Println("Hallo", name)
	case "Wahyu":
		fmt.Println("Hallo", name)
	case "Nur":
		fmt.Println("Hallo,", name)
	default:
		fmt.Println("Hallo boleh kenalan dulu engga,", name)
	}

	// Switch short statement
	switch lenght := len(name); lenght > 5 {
	case true:
		fmt.Println("Nama terllau panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// Switch tanpa kondisi
	// kondisi lagnsung masuk ke case
	// sama saja seperti if else
	panjang := len(name)
	switch {
	case panjang < 5:
		fmt.Println("Nama terllau panjang")
	case panjang > 4:
		fmt.Println("Nama terlalu pendek")
	default:
		fmt.Println("Nama sudah benar")
	}
}
