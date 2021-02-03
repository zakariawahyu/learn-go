package main

import "fmt"

func main()  {
	var kata1 bool = true
	var kata2 bool = false

	if kata1 == true {
		fmt.Println("Iya anda benar ini angka1")
	} else if kata2 == true {
		fmt.Println("Kurang benar")
	}

	fmt.Println("Benar =", true)
	fmt.Println("Salah =", false)


}
