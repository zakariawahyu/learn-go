package main

import "fmt"

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
