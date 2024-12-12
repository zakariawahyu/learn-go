package main

import "fmt"
/**
Saat membaut function, kadang-kadang kita membutuhkan data dari luar atau kita sebut dengan parameter.
Kita bisa menambahkan parameter di function, lebih dari satu.
Jika kita menambahkan parameter di funtionnya, maka ketika memanggil function tersebut wajib memasukkan data/value parameternya.
 */
func main()  {
	sayHelloTo("Zakaria", "Wahyu", 21)
	sayHelloTo("Nur", "Utomo", 21)
}

func sayHelloTo(firtName string, lastName string, Age int)  {
	fmt.Println("Hello", firtName, lastName, "dan umur", Age)
}
