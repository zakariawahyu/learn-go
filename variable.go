package main

import "fmt"

func main()  {
	//macam-macam deklarasi pada go
	var name string

	name = "Zakaria"
	fmt.Println(name)

	name = "Zakaria Wahyu"
	fmt.Println(name)

	// deklarasi tanpa menyebutkan tipe datanya
	var nameKu = "Zakaria Wahyu Nur Utomo"
	fmt.Println(nameKu)

	var umur = 21
	fmt.Println(umur)

	// penulisan langsung tanpa harus nulis var dan tipe data
	tempatlahir := "Klaten"
	fmt.Println(tempatlahir)

	// penulisan juga bisa seperti array (multiple variable)
	var (
		kerja = "Haruka"
		posisi = "Backend Developer"
	)

	fmt.Println(kerja)

	kerja = "backend"

	fmt.Println(kerja)
	fmt.Println(posisi)
}
