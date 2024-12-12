package main

import (
	"errors"
	"fmt"
)

/**
Ga-lang memiliki interface yang digunakan sebagai kontrak untuk membuat error, nama interface nya adalah error
Untuk membuat error tidak perlu manual
Golang sudah menyediakan library untuk membuat helper secara mudah yang terdapat di package errors
 */

func Pembagian(nilai int, pembagi int) (int, error)  {
	if pembagi == 0 {
		return 0, errors.New("Error, pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main()  {
	var contohError = errors.New("Upss something wrong")
	fmt.Println(contohError.Error())

	// contoh fucntion pembagian error
	hasil, err := Pembagian(10, 2)
	if err == nil {
		fmt.Println("Hasil dari pembagian adalah :", hasil)
	} else {
		fmt.Println("Warning Error ! ", err.Error())
	}
}
