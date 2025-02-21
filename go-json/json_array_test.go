package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArrayComplexEncode(t *testing.T) {
	customer := &Customer{
		FirstName: "Hafidz",
		Addresses: []Address{
			{
				Street:     "Jalan in Aja Dulu",
				Country:    "Indonesia",
				PostalCode: "99999",
			},
			{
				Street:     "Jalan Buntu",
				Country:    "Indonesia",
				PostalCode: "12345",
			},
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Hafidz","MiddleName":"","LastName":"","Age":0,"IsMarried":false,"Hobbies":null,"Addresses":[{"Street":"Jalan in Aja Dulu","Country":"Indonesia","PostalCode":"99999"},{"Street":"Jalan Buntu","Country":"Indonesia","PostalCode":"12345"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexEncode(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jalan in Aja Dulu",
			Country:    "Indonesia",
			PostalCode: "99999",
		},
		{
			Street:     "Jalan Buntu",
			Country:    "Indonesia",
			PostalCode: "12345",
		},
	}

	bytes, err := json.Marshal(addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan in Aja Dulu","Country":"Indonesia","PostalCode":"99999"},{"Street":"Jalan Buntu","Country":"Indonesia","PostalCode":"12345"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}

	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(*addresses)
}
