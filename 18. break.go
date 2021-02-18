package main

import "fmt"

func main()  {
	// break untuk menghentikan seluruh perulangan
	for i := 0; i <10; i++ {
		// jika i = 5 maka perulangan akan berhenti
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke", i)
	}
}
