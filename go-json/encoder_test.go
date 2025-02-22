package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("customer_output.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Muhammad",
		MiddleName: "Hafidz",
		LastName:   "Athaya",
	}
	_ = encoder.Encode(&customer)

	fmt.Println(encoder)
}
