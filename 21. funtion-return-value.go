package main

import "fmt"
/**
Function bisa membalikkan data atau return.
Untuk memberitahu bahwa funtion mengembalikan data, kita harus menuliskan tipe data kembalian dari function tersebut.
Jika fucntion tersebut kita deklarasikan dengan tipe data pengembalian, maka wajib di dalam function untuk menggembalikan tipe datana
Untuk mengembalikan data dari funtion, ktia bisa menggunakan kata kunci return, diikuti dengan datanya
 */
func main()  {
	result := getHello("Zakaria")
	fmt.Println(result)
	// get kosong
	fmt.Println(getHello(""))
}

func getHello(name string) string  {
	if name == "" {
 		return "Hello world kenalan dulu"
	} else {
		return "Hello " + name
	}
}