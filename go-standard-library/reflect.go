package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readField(value any) {
	valueType := reflect.TypeOf(value)

	fmt.Println("Type Name", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		structField := valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)
		fmt.Println("Tag required:", structField.Tag.Get("required"))
		fmt.Println("Tag max:", structField.Tag.Get("max"))
	}
}

func isValid(value any) bool {
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			if data == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	readField(Sample{"Eko"})
	readField(Person{"Eko", "Jakarta", "eko@example.com"})

	person := Person{"Eko", "Jakarta", "test"}
	fmt.Println(isValid(person))
}
