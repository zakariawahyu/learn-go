package main

import "fmt"

func main()  {
	// Operasi boolean
	// &&   => AND
	// ||   => OR
	// !    => kebalikan

	var ujian = 80
	var absensi = 80

	var lulusUjian = ujian >= 80 // true
	var lulusAbsensi = absensi >= 80 // true

	var lulus = lulusUjian && lulusAbsensi // true dan true
	fmt.Println(lulus) // true

	//atau langsung
	fmt.Println(ujian >= 80 && absensi >= 80)
}
