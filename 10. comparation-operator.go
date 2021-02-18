package main

import "fmt"

func main()  {
	// Operator pembanding
	// >     =>  lebih dari
	// <     =>  kurang dari
	// >=    =>  lebih dari sama dengan
	// <=    =>  kurang dari sama dengan
	// ==    =>  sama dengan
	// !=    =>  tidak sama dengan

	var name1 = "Zaka"
	var name2 = "Zaka"

	var result bool = name1 == name2
	fmt.Println(result)

	var value1 = 100
	var value2 = 200
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
	fmt.Println(value1 >= value2)
	fmt.Println(value1 <= value2)

}
