package main

import "fmt"

func main()  {
	sayHello()
	sayHello()
	sayHello()
	for i := 0; i <= 10; i++ {
		sayHello()
	}
}

func sayHello()  {
	fmt.Println("Hello World")
}
