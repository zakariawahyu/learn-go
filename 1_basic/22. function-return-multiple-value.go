package main

import "fmt"
/**
Returning multiple value
Funtion tidak hanya dapat emngembalikan satu value, tapi juga bisa banyak atau multiple value.
Untuk memberitahu jika funtion mengembalikan multiple value, kita harus menuis tipe data return valuenya di funtion.
 */
func main()  {
	firtName, lastName := getFullName()
	fmt.Println(firtName, lastName)

	//menghiraukan value dikasih _
	//firtName, _ := getFullName()
	//fmt.Println(firtName)
}

func getFullName() (string, string)  {
	return "Zakaria", "Wahyu"
}
