package main

import "fmt"

func main()  {
	// digunakan untuk menghentikan perulangan yang berjalan dan langsung melanjutkan ke perulangan berikutnya
	for i := 0; i <= 10; i++ {
		// jika i genap maka akan di stop dan lanjut lagi perulangan berikutnya dan menampilkan ganjil
		if i % 2 == 0{
			continue
		}

		fmt.Println("Perulangan ke", i)
	}
}
