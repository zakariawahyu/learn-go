package main

import "fmt"

/**
Interface adalah tipe data abstract, tidak memiliki implementasi langsung.
Sebuah interface biasanya berisikan definisi-definisi method
Biasanya interface digunakan sebagai kontrak
 */

// deklarasi interface terlebih dahulu
// berisikan method atau function
type HasName interface {
	getName() string
}

func sayName(hasName HasName)  {
	fmt.Println("Hello", hasName.getName())
}
// deklarasi struct
type Person struct {
	Name string
}

type Animal struct {
	Name string
}

// contoh penggunaan interface
func (animal Animal) getName() string  {
	return animal.Name
}

func (person Person) getName() string{
	return person.Name
}

func main()  {
	var zaka = Person{
		Name: "Zakaria",
	}
	sayName(zaka)

	var kucing = Animal{
		Name: "Kucing",
	}
	sayName(kucing)
}
