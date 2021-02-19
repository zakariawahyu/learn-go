package main

import (
	"fmt"
	"math"
)

/**
Package math merupakan package yang berisikan constant dan fungsi matematika
golang.org/pkg/math
 */

/**
Beberapa function package math
math.Round(float64)     => Membulatkan float64 keatas atau kebawah , sesuai dengan yang paling dekat
math.Floor(float64)     => Membulatkan float64 ke bawah
math.Ceil(float64)      => Membulatkan float64 ke atas
math.Max(float64, float64) => Mengembalikan nilai float64 yang paling besar
math.MIN(float64, float64) => Mengembalikan nilai flaot64 yang palin kecil
 */
func main()  {
	fmt.Println(math.Round(1.2))
	fmt.Println(math.Round(1.7))
	fmt.Println(math.Floor(1.7)) // dipaksa dibulatkan ke baawah
	fmt.Println(math.Ceil(1.3)) // dipaksa dibulatkan ke atas
	fmt.Println(math.Max(23, 88))
	fmt.Println(math.Min(121, 12))
}
