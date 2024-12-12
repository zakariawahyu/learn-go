package main

import (
	"fmt"
	"os"
)
/**
Golang menyediakan banyak sekali package bawaan, salah satunya adalah package os.
Package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen (bisa semua sistem operasi)
contoh detail lain ada di link : golang.org/pkg/os
 */
func main()  {
	// untuk mengetahui lokasi build
	args := os.Args
	fmt.Println("Argument :", args)

	// untuk mengetahui nama hostname
	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname :", name)
	} else {
		fmt.Println("Error :", err.Error())
	}

	// untuk mengetahui nama os
	fmt.Println(os.Getenv("USERNAME"))
}
