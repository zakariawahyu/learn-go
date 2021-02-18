package main

import "fmt"
/**
Biasanya saat kita memberitahu bahwa sebuah funtion mengembalikan value, maka kita hanya mendeklarasikan tipe data return value di funtion.
Namun kita juga bisa membuat variabel secara langsung di tipe data return funtionnya.
 */
func main()  {
	firtName, midleName, lastName := getFullNameku()
	fmt.Println(firtName)
	fmt.Println(midleName)
	fmt.Println(lastName)
}

func getFullNameku()(firstName, midleName, lastName string)  {
	firstName =  "Zakaria"
	midleName =  "Wahyu"
	lastName  =  "NUr"

	return firstName, midleName, lastName
	// atau tinggal ketik return saja
	// return
}

