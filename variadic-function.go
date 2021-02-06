package main

import "fmt"

func main()  {
	// parameter  yang berada  diposisi terkahir, memiliki kemampuan untuk  dijasikan varargs (variabel argument)
	// varargs artinya data bisa menerima 1 input lebih dari satu, atau anggap saja semacam array
	// perbedaan dengan array, kalau array harus membuat arraynya terlebih dahulu dan kalau varagrs bisa langsung mengrim datanya lebih dari satu

	total := sumAll(10, 10, 10, 10)
	fmt.Println(total)

	slice := []int{1,1,1,1}
	total = sumAll(slice...)
	fmt.Println(total)
}

func sumAll(number ...int) int  {
	total := 0
	for _, value := range number {
		total += value
	}
	return total
}