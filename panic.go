package main

import "fmt"

// panic funtion bisa digunakan untuk menghrntikan program
// panic funtion biasanya dipanggil ketika terjadi error pada saat program dijalankan
// saat panic funtion dipanggil maka program akan berhenti, namun defer funtion akan tetap dieksekusi

func endApp()  {
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool)  {
	// namun defer akan tetap dieksekusi
	defer endApp()

	// jika panic kode program dibawahnya tidak dieksekusi
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}
func main()  {
	// run 1 normal
	runApp(false)

	// run 2 sengaja error
	runApp(true)
}
