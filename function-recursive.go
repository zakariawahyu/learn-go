package main

import "fmt"

func factorialLoop(value int) int  {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factoeialRevursive(value int) int  {
	if value == 1 {
		return 1
	} else {
		return value * factoeialRevursive(value-1)
	}
}
func main()  {
	// recursive function adalah function yang memanggil dirinya sendiri
	// contoh kasus lebih mudah menggunakan recursive adaalah factorial

	loop := factorialLoop(5)
	fmt.Println(loop)
	fmt.Println(5 * 4 * 3 * 2 * 1)

	recursive := factoeialRevursive(6)
	fmt.Println(recursive)
	fmt.Println(6 * 5 * 4 * 3 * 2 * 1)
}
