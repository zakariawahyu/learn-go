package main

import "fmt"

/**
Walaupun method akan menempel di struct , tapi sebenernya data struct yang di akses di dalam method adalah pass by value
Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu duplicate ketika memanggil
 */

// membaut struct
type Man struct {
	Name string
}

// method yang didalamnya ada pointer
// penulisan menggunakan operator * sebagai tanda kalau itu pointer
func (man *Man) Married()  {
	man.Name = "Mr. " + man.Name
}

func main()  {
	// deklarasi data utamanya
	data := Man{
		"Zakaria",
	}
	// panggil method
	data.Married()
	fmt.Println(data)
}
