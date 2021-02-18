package main

import "fmt"

// defer function adalah function yang bisa dijadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
// defer masih bisa di eksekusi walaupun terjadi error di function yang dieksekusi

func logging()  {
	fmt.Println("Selesai memanggil function")
}

func runAplication(value int)  {
	// defer selalu diatas jika function error akan tetap dijalankan
	defer logging()
	fmt.Println("Run Aplication....")
	result := 10 / value
	fmt.Println("Result :", result)
}
func main()  {
	// run pertama benar
	runAplication(10)

	// contoh ketika input function error defer tetap berjalan
	// run ke 2 dengan error
	runAplication(0)
}
