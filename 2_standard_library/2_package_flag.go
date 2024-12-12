package main

import (
	"flag"
	"fmt"
)

/**
Package flag berisikan functionalitas untuk memparsing comand line argument
contoh : go run package-flag.go -host=MR.ROBOT -user=Zakaria -password=123
detail link : golang.org/pkg/flag
 */
func main()  {
	host := flag.String("host", "MR.ROBOT", "Put your hostname")
	user := flag.String("user", "Zakaria", "Put your username")
	pass := flag.String("password", "inipas", "Put your password")

	// wajib dipanggil untuk parsing data
	flag.Parse()

	// flag termasuk pointer maka harus dipanggil *
	fmt.Println("Hostname : ",*host)
	fmt.Println("User :", *user)
	fmt.Println("Password", *pass)

}
