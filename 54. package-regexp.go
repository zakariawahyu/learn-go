package main

import (
	"fmt"
	"regexp"
)

/**
Package regexp adalah utilitas di golang untuk melakukan pencarian regular expression.
Regular expresiion di golang menggunakan librabry C yang dibuat google bernama RE2
	github.com/google/re2/wiki/Syntax
golang.org/pkg/regexp
 */

/**
Beberapa function di package regexp
regexp.MustCompile(string)       => Membuat regexp
regexp.MatchString(string) bool  => Mengecek apakah regexp match dengan string
regexp.FindAllString(stirng, max)=> Mencari string yang match dengan maximum jumlah hasil
 */
func main()  {
	// striong setelah Z dan sebelum kata boleh a-z dan huruf kecil
	regex := regexp.MustCompile(`z([a-z])k`)

	fmt.Println(regex.MatchString("zak"))
	fmt.Println(regex.MatchString("aAk")) // false karena A huruf besar

	// cari yang match 1
	result := regex.FindAllString("zak zay zok zai zul zbk", 2)
	fmt.Println(result)

}
