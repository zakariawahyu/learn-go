package main

import "fmt"
/**
Biasanay dalam bahasa pemogrman lain, obejct yang belum di inisialisasi maka secara otomatis nilainya adalah null atau nil.
Berbeda di golang tidak merta langsung return null, tapi secara otomatis akan dibuatkan default value nya.
Namun di Golang ada data nill yaitu data kosong.
Nil sendiri hanya dapat digunakan dalam beberapa tipe data, seperti interface, funtion, map, slice, pointer dan channel
 */
func newMap(name string) map[string]string  {
	if name == "" {
		return nil
	}else {
		return map[string]string{
			"name" : name,
		}
	}
}

func main()  {
	data := newMap("")
	fmt.Println(data)

	if data == nil {
		fmt.Println("Data kosong")
	}else {
		fmt.Println(data)
	}

	// ini contoh variabel a data kosong
	// maka di golang akan secara otomatis membuatkan default berdasarkan tipe datanya
	var a int
	fmt.Println(a)
}
