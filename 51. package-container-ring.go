package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

/**
Package container-ring adalah implementasi struktur data circular list
Circilar list adalah struktur data berbentuk ring atau cincin, dimana akhir element akan kembali ke awal element(HEAD)
golang.org/pkg/container/ring
 */
func main()  {
	data := ring.New(5)
	// contoh membuat manual
	data.Value = "Zakaria"

	data2 := data.Next()
	data2.Value = "Nur"

	data3 := data.Prev()
	data3.Value = "Wahyu"

	data.Do(func(value interface{}) {
		fmt.Println(value)
	})

	// lebih simoel
	angka := ring.New(6)
	for i := 0; i <= angka.Len(); i++ {
		angka.Value = "Angka ke" + strconv.FormatInt(int64(i), 10)
		angka = angka.Next()
	}

	angka.Do(func(angkaval interface{}) {
		fmt.Println(angkaval)
	})

}
