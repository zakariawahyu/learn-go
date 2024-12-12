package main

import (
	"container/list"
	"fmt"
)

/**
Package container/list adalah implementasi struktur data double linked list di golang
golang.org/pkg/container/list
 */

/**
Double linked list adalah data yang berurutan dan saling terkait data setelah atau sebelumnya

 */
func main()  {
	data := list.New()

	//untuk menambahkan data ke belakang dengan PushBack
	data.PushBack("Wahyu")
	data.PushBack("Nur")
	data.PushBack("Utomo")

	//untuk menambahkan data ke depan dengan PushFront
	data.PushFront("Zakaria")

	// untuk mengambil data yang pertama
	depan := data.Front().Value
	fmt.Println(depan)
	//mengambil data setelah dari depan
	// mengambil data setelah tingal tabahin neext
	depanNeext := data.Front().Next().Value
	fmt.Println(depanNeext)

	// untuk mengambil data paling belakang
	belakang := data.Back().Value
	fmt.Println(belakang)
	// mengambil data sebelumnya dari belakang
	// mengambil data sebelum tinggal tambahkan prev
	belakangPrev := data.Back().Prev().Prev().Value
	fmt.Println(belakangPrev)

	// Mengambil semua data dari depan
	i := 1
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println("Data ke", i ,":",e.Value)
		i++
	}

	// Mengambil semua data dari belakang
	j := 4
	for elementt := data.Back(); elementt != nil; elementt = elementt.Prev() {
		fmt.Println("Data ke", j ,":", elementt.Value)
		j--
	}
}
