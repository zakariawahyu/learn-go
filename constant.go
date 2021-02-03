package main

import "fmt"

func main()  {
	// variable yang nilainya tetap atau tidak dapat dirubah lagi
	const firstName = "Zakaria Wahyu"
	const lastName = "Nur Utomo"
	const value = 100

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	// deklarasi multiple variable
	const (
		umur = 21
		kota = "Bandung"
		kerja = "HarukaEDU"
	)

	fmt.Println(umur)
	fmt.Println(kota)
	fmt.Println(kerja)
}
