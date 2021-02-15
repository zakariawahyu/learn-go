package main

import (
	"fmt"
	"learn-go/helper"
)


// Memanggil funtion dari luar package, maka dari itu harus import terlebih dahulu packagenya
func main()  {
	helper.SayHello("Zakaria")

	// maksa untuk memanggil function luar package, dengan nama function diawali huruf kecil
	// hasilnya tidak bisa dipanggil
	//helper.sayGoodbye("All")

	fmt.Println(helper.Aplication)
}
