package main

import (
	"fmt"
	"reflect"
)

/**
Dalam bahasa pemograman, biasanya ada fitur reflection yang dimana kita bisa meilhat struktur kode kita pada saat aplikasi sedang berjalan.
Hal ini juga bisa dilakukan di golang dengan package reflection.
golang.org/pkg/reflect/
 */

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Novel struct {
	Name string `required:"true"`
	Description string `required:"true"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}

func main()  {
	sample := Sample{
		"Zakaria",
	}

	sampleType := reflect.TypeOf(sample)
	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))

	fmt.Println(IsValid(sample))

	// tidak ada validasi karena tidak required di struct
	novel := Novel{"Soekaro presiden", "Dia adalah presiden pertama indonesia"}
	fmt.Println(IsValid(novel))
}