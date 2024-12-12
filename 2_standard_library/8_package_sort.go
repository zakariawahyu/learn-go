package main

import (
	"fmt"
	"sort"
)

/**
Package sort adalah sebuah package yang berisikan utilitas unuk proses pengurutan
Agar data kita bisa diurutkan, kita juga harus mengimplementasikan kontrak di interdace sort.interface
golang.org/pkg/sort
 */

type User struct {
	Name string
	Age int
}

type UserSlice []User

// untuk mengembalikan berapa jumlah data

func (value UserSlice) Len() int  {
	return len(value)
}

// untuk mengecek apakah data i lebih kecil atau besar daripada 9

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

// untuk menukar posisi index i dan j

func (value UserSlice) Swap(i, j int)  {
	value[i], value[j] = value[j], value[i]
}
func main()  {
	users := []User{
		{"Zakaria", 20},
		{"Wahyu", 19},
		{"Nur", 18},
		{"utomo", 45},
	}

	fmt.Println("Sebelum sort : ",users)
	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
