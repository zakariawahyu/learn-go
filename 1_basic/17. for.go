package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan counter =", counter)
		counter++
	}

	// for dengan statement
	for nilai := 1; nilai <= 10; nilai++ {
		fmt.Println("Perulangan nilai ke", nilai)
	}

	// for range
	slice := []string{"Zakaria", "Wahyu", "Nur", "Utomo"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//cara cepat for range
	for index, name := range slice{
		fmt.Println("Index", index, "=", name)
	}

	// mengganti index dengan _ agar bisa mencetak value saja
	// _ artinya variabel tidak dibutuhkan
	for _, name := range slice{
		fmt.Println(name)
	}

	person := make(map[string]string)
	person["name"] = "Zakaria"
	person["title"] = "Programmer"

	for key, value := range person{
		fmt.Println("Key =", key, "Val =",value)
	}
}