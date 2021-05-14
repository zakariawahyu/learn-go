package main

import  (
	"fmt"
	"time"
)

/**
Package time adalah package yang berisikan funsionalitas untuk management waktu di golang
golang.org/pkg/time
 */

/**
Beberapa function time yang bisa digunakan
time.Now()     			   => Untuk mendapatkan waktu saat ini
time.Date(     			   => Untuk membuat waktu
time.Parse(layout, string)  => Untuk memparsing waktu dari string
 */
func main()  {
	now := time.Now()
	fmt.Println(now.Day()) // untuk mendapatkan hari
	fmt.Println(now.Month()) // untuk mendapatkan bulan
	fmt.Println(now.Year()) // untuk mendapatkan tahun
	fmt.Println(now.Hour()) // untuk mendapatkan jam
	fmt.Println(now.Minute()) // untuk mendapatkan menit
	fmt.Println(now.Second()) // untuk mendapatkan detik
	fmt.Println("Detail waktu :", now.Local())

	// membuat waktu sendiri
	utc := time.Date(2021, 10, 10, 10, 10 ,10, 10, time.UTC)
	fmt.Println(utc)

	// layout pada golang berbeda dengan yang lain, tidak seperti Y-M-D
	// Tapi ada format sendiri dengan
	layout := "2006-01-02"
	parse, err := time.Parse(layout, "2020-06-25")
	if err == nil {
		fmt.Println(parse)
	} else {
		fmt.Println(err.Error())
	}
}
