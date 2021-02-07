package main

import "fmt"

// recover function adalah untuk menangkap data panic
// jika ada panic terus kita recover maka aplikadi tidak jadi berhenti
// dengan recover proses panic akan terhenti dan program akan tetap berjalan

func killApp()  {
	// recover untuk menangkap panic
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message :", message)
	}

	fmt.Println("Aplikasi selesai kill")
}

func runApps(error bool)  {
	// namun defer akan tetap dieksekusi
	defer killApp()

	// jika panic kode program dibawahnya tidak dieksekusi
	if error {
		panic("APLIKASI ERROR")
	}

	fmt.Println("Aplikasi berjalan")
}
func main()  {
	// run 1 normal
	runApps(false)

	// run 2 sengaja error
	runApps(true)
	fmt.Println("Recover masih jalan")
}