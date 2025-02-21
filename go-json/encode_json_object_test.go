package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	IsMarried  bool
	Hobbies    []string
	Addresses  []Address
}

func TestEncodeJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Muhammad",
		MiddleName: "Hafidz",
		LastName:   "Athaya",
		Age:        20,
		IsMarried:  false,
		Hobbies:    []string{"Coding", "Reading", "Gaming"},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
