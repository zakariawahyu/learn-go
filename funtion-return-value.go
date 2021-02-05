package main

import "fmt"

func main()  {
	result := getHello("Zakaria")
	fmt.Println(result)
	// get kosong
	fmt.Println(getHello(""))
}

func getHello(name string) string  {
	if name == "" {
 		return "Hello world kenalan dulu"
	} else {
		return "Hello " + name
	}
}