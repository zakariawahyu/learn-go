package main

import (
	"fmt"
	"strconv"
)

/**
Sebelumnya kita sudah belajar cara konversi tipe data, misal dari int32 ke dalam int64
Bagaimana jika kita butuh melakukan konversi yang tipe datanya berbeda? Misal dari int ke string dan sebaliknya
Hal tersebut bisa dilakukan dengan strconv(string conversion)
golang.org/pkg/strconv
 */

/**
Beberapa fucntion di package strconv
kalau awal parse itu dari string ke semua tipe data
strconv.ParseBool(string)     => Mengubah string ke bool
strconv.ParseFloat(string)    => Mengubah string ke float
strconv.ParseInt(string)      => Mengubah string ke int

kalau awal format itu dari semua tipe data ke string
strconv.FormatBool(bool)      => Mengubah bool ke string
strconv.FormatFloat(float, .) => Mengubah float64 ke string
strconv.FormatInt(int, ..)    => Mengubah int64 ke dtring
 */
func main()  {
	boolean, err := strconv.ParseBool("false") // yang bisa di konversi true / false
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error : ", err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println("Error :", err.Error())
	}

	value := strconv.FormatInt(2000000, 10)
	fmt.Println(value)
}
