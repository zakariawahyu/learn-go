package main

import "fmt"

func main()  {
	type NoKTP string
	type married bool

	var ktpzaka NoKTP = "331033112506990003"
	var marriedStatus = false
	fmt.Println(ktpzaka)
	fmt.Println(marriedStatus)
}
