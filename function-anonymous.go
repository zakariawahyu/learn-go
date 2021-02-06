package main

import "fmt"

type Blacklist func(string) bool

func registerUsers(name string, blacklist Blacklist)  {
	if blacklist(name) {
		fmt.Println("Your are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main()  {
	blacklist := func(name string) bool{
		return name == "Admin"
	}

	registerUsers("Admin", blacklist)
	registerUsers("Zakaria", blacklist)
	registerUsers("Rahasia", func(name string) bool {
		return name == "Rahasia"
	})
}
