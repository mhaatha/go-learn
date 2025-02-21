package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSONObject(t *testing.T) {
	jsonString := `{"FirstName":"Muhammad","MiddleName":"Hafidz","LastName":"Athaya","Age":20,"IsMarried":false,"Hobbies":["Coding","Reading","Gaming"]}
`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.IsMarried)
	fmt.Println(customer.Hobbies)
}
