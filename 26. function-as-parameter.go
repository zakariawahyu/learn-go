package main

import "fmt"
/**
Funstion as parameter
Funtion tidak hanya bisa kita simpan di dalam variabel sebagai value.
Namun juga bisa kita gunakan sebagai parameter untuk function lain
 */

// function type declaration

type Filter = func(string)string
func sayHElloWithFilter1(name string, filter Filter)  {
	namedFilter := filter(name)
	fmt.Println("Hello", namedFilter)
}

func sayHElloWithFilter2(name string, filter func(string)string)  {
	namedFilter := filter(name)
	fmt.Println("Hello", namedFilter)
}

func spamFilter(name string)string  {
	if name == "Gblk" {
		return "..."
	} else {
		return name
	}
}

func main()  {
	sayHElloWithFilter1("Zakaria", spamFilter)
	sayHElloWithFilter2("Wahyu", spamFilter)
}
