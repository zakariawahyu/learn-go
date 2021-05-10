package main

import "fmt"

func main()  {
	// Membuat slice dari array
	// array[low:high]   =  Membuat slice dari array dimulai dari index low samppai index sebelum high
	// array[low:]       =  Membuat slice dari array dimulai dari index low  sampai index akhir array
	// array[:high]      =  Membuat slice dari array dimulai dari index 0 dampai index sebelum high
	// array[:]          =  Membuat slice dari array dimulai dari  index 0 sampai index akhir array

	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 =  months[4:7]
	fmt.Println(slice1)

	// Function slice
	// len(slice)   =  Untuk mendapatkan panjang slice
	// cap(slice)   =  Untuk mendapatkan kapasitas
	// append(slice, data)  =  Membuat slice baru  dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat data array baru
	// make([]TypeData, lenght, capacity)  =  Membuat slice baru tanpa array
	// copy(destionation, source)  =  Menyalin slice dari source ke destination

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// mengubah array dan slice juga ikut keubah
	months[5] = "DIubah"
	fmt.Println(slice1)

	// mengubah slice dan array juga ikut keubah
	slice1[0] = "Mei lagi"
	fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "EKO")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	// Membaut slice baru dari awal tanpa array
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Zakaria"
	newSlice[1] = "Wahyu"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// Copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// Perhatian array dan slice
	iniArray := [5]int{1,2,3,4,5}
	iniSlice := []int{1,3,4,5,}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
