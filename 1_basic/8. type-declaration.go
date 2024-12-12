package main

import "fmt"
/**
Type declaration adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada.
Type declaration biasanya digunakan untuk membuat alias terhadao tipe data yang sudah ada, dengan tujuan agar lebih mudah dimengeri
 */
func main()  {
	type NoKTP string
	type married bool

	var ktpzaka NoKTP = "331033112506990003"
	var menikah married = true
	var marriedStatus = false
	fmt.Println(ktpzaka)
	fmt.Println(marriedStatus)
	fmt.Println(menikah)
}
