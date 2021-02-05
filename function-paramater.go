package main

import "fmt"

func main()  {
	sayHelloTo("Zakaria", "Wahyu", 21)
	sayHelloTo("Nur", "Utomo", 21)
}

func sayHelloTo(firtName string, lastName string, Age int)  {
	fmt.Println("Hello", firtName, lastName, "dan umur", Age)
}
