package main

import "fmt"

func main()  {
	// Map adalah tipe data kumpulan key dan value, dimana kata kuncinya bersifat unik alias tidak boleh sama

	// membuat map baru
	person := map[string]string{
		"name" : "Zakaria",
		"addres" : "Bandung",
	}

	// menambahkan data
	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["addres"])

	// Function map
	// len    				=>  mengambil jumlah data dalam map
	// map[key]  			=>  mengambil data dengan key
	// map[key] = value		=>  mengubah atau menambah data map baru dengan key
	// make(map[TypeKey]TypeValue)   => membuat map baru
	// delete(map, key)     =>  menghapus data di map dengan key

	book := make(map[string]string)

	book["title"] = "Belajar Go-Lang"
	book["author"] = "Zakaria"
	book["ups"] = "Salah"

	fmt.Println("Sebelum hapus =", book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println("Setelah hapus =", book)
	fmt.Println(len(book))
}
