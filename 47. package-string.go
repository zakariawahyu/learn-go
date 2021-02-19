package main

import (
	"fmt"
	"strings"
)

/**
Package string adalah package yang berisikan function-funtion untuk memanipulasi tipe data string
golang.org/pkg/strings
 */

/**
Beberapa function di package string
strings.Trim(string, cutset)          => Memotong cutset di awal dan akhir string
strings.ToLower(string)               => Membuat semua karakter string menjadi lower case
strings.ToUpper(string)               => Membuat semua karakter string menjadi upper case
strings.Split(string, separator)      => Memotong string berdasarkan separator
strings.Constains(string, search)     => Mengecek string apakah string mengandung string lain
strings.ReplaceAll(string, from, to)  => Mengubah semua string dari from ke to
 */

func main()  {
	nama := "Zakaria Wahyu "
	fmt.Println(strings.Contains(nama, "Wahyu")) // ada kata Wahyu maka true
	fmt.Println(strings.Contains(nama, "Utomo")) // tida kada kata Utomo maka false
	fmt.Println(strings.Split(nama, "a"))
	fmt.Println(strings.ToLower(nama))
	fmt.Println(strings.ToUpper(nama))
	fmt.Println(strings.ToTitle(nama))
	fmt.Println(strings.Trim(nama, " ")) // menghilangkan spasi di kanan dan kiri string
	fmt.Println(strings.ReplaceAll(nama, "Wahyu", "Nur"))
}
