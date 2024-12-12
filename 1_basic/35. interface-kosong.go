package main

import "fmt"
/**
Golang bukan merupakan pemogrman yang berorientasi objek.
Biasanya dalam pemogrman beroirentasi objek, ada satu data parent di puncak yang bisa dianggap sebagai semua implementasi data yang ada di bahasa pemograman tersebut.
Contoh di java adalah java.lang.Object
Untuk menangani kasus ini, di golang bisa melakukan interface
Interface kosong adalah interface yang tidak memiliki deklarasi method satupun , hal ini membuat secara otomatis semua tipe data yang akan di emplementasinya
 */

// deklatasi interface kosong dalam function
// karena interface kosong bisa return tipe data apa saja, disini contoh int, bool dan string

func Upssh(i int) interface{}  {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Upsshhh"
	}
}

func main()  {
	// pembuatan data pun juga harus interface kosong
	// memasukkan parameter 3
	var data interface{} = Upssh(3)
	fmt.Println(data)
}