package main

import "fmt"

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

