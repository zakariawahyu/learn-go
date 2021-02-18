package main

import "fmt"

func main()  {
	fmt.Println("Zakaria")
	fmt.Println("Zakaria Wahyu")
	fmt.Println("Zakaria Wahyu Nur Utomo")

	// menghitung jumlah string
	fmt.Println(len("Zakaria"))

	var namaku string = "Zakaria"
	fmt.Println("Jumlah string", namaku, "adalah =", len(namaku))

	// mengambil string ke 0
	fmt.Println("Zakaria Wahyu"[0])
	// mengambim string ke 1
	fmt.Println("Zakaria"[1])
}
