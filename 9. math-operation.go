package main

import "fmt"

func main()  {
	// langsung
	var result = 10 + 10
	fmt.Println(result)

	// dengan deklarasi data
	var a = 8
	var b = 5
	var c = 1
	var hasil = a + b + c
	fmt.Println("Hasil penjumlahan a + b + c =", hasil)

	// operasi augmented assignments
	// i = i + 10     disingkat=>    i += 10
	// j = j - 10     disingkat=>    j -= 10
	// k = k * 10     disingkat=>    k *= 10
	// l = l % 10     disingkat=>    l %= 10

	var i = 10
	i += 10
	fmt.Println(i)

	var j = 10
	j -= 10
	fmt.Println(j)

	var k = 10
	k *= 10
	fmt.Println(k)

	var l = 10
	l %= 10
	fmt.Println(l)

	// Unary operation
	// x = x + 1   disingkat=> x++
	// y = y -1    disingkat=> y--
	// -    => negatif
	// !    =>  boleean kebalikan

	var x = 10
	x++
	fmt.Println(x)

	var y = 10
	y--
	fmt.Println(y)

	var negatif = -100
	var positif = +100 // default emang positif
	fmt.Println(negatif)
	fmt.Println(positif)
}
