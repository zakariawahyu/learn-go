package main

import "fmt"

func main()  {
	name := "Zakaria"
	counter := 0

	// bisa ngakses data disekitarnya
	// jika datanya kerubah maka akan kerubah di closure
	increment := func() {
		name = "Wahyu"
		fmt.Println("Incremenet")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
