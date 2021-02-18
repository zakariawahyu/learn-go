package helper

import "fmt"

/**
Package
Package adalah tempat yang bisa digunakan untuk mengorganisir kdoe program yang kita buat di golang.
Dengan menggunakan package, kita bisa merapikan kode program yang kita buat.
Package sendiri sebenarnya hanya direktori folder pada sistem operasi kita
 */

/**
Import
Secara standar , file golang hanya bisa mengakes file golang lainnya dalam satu package yang sama.
Jika ingin mengakses file golang dari package luar, maka bisa menggunakan import
 */

/**
Access Modifier
Di bahasa pemograman lain, biasanya ada kata kunci yang bisa digunakan untuk menentukan access modifier terhadap suatu fundtion atau variabel.
Di golang, untuk menentukan acces modifier cukup dengan merubah penamaan pada nama functionnya.
Jika nama diawali dengan huruf besar, maka artinya bisa di akses ke semua package lain. Jika diawali dengan huruf kecil, artinya tidak bisa diakses di package lain
 */

// tidak bisa diakses di package lain karena penamaan function diawalo dengan huruf kecil
func sayGoodbye(name string)  {
	fmt.Println("Hello ini "+ name +" sedang coba membuat package helper")
}

// bisa diakses di package lain karena penamaan function diawali dengan huruf besar
func SayHello(name string) {
	fmt.Println("Hello ini "+ name +" sedang coba membuat package helper")
}
var version = 1 // tidak bisa diakses di package lain
var Aplication = "Belajar golang" // bisa di akses karena diawali dengan huruf besar
