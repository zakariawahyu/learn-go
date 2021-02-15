package other

import "fmt"

// dalam 1 package tidak boleh ada nama atau funtion yang sama
// jika beda package baru bisa menggunakan funtion yang sama
// contoh pada package helper terdapat function sayHello()
func SayHello(name string) {
	fmt.Println("Hello ini "+ name +" sedang coba membuat package helper")
}