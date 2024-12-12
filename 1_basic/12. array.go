package main

import "fmt"

func main()  {
	var names[4]string
	names[0] = "Zakaria"
	names[1] = "Wahyu"
	names[2] = "Nur"
	names[3] = "Utomo"

	var wahyu = names[1]
	fmt.Println(names[0])
	fmt.Println(wahyu)

	// menulis secara langsung
	var values = [3]int{
		98,
		95,
		90,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	// Function array
	// len(array)   =  mendapatkan jumlah panjang array
	// array[index] =  mendappatkan data di posisi index yang terkait
	// array[index] =  mengubah data di posisi index yang terkait

	fmt.Println("Panjang array nama = ", len(names))
	fmt.Println("Panjang array values = ", len(values))

	values[2] = 4
	fmt.Println(values)
}
