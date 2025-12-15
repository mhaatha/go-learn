package main

import (
	"encoding/json"
	"fmt"

	"github.com/mhaatha/go-learn/advance-json-struct-tags/challenge"
)

type Product struct {
	ID           string `json:"product_id"`
	Name         string `json:"name"`
	Description  string `json:"desc,omitempty"`
	InternalCode string `json:"-"`

	Stock    *int  `json:"stock,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
}

func main() {
	// zeroStock := 0
	// active := false

	// p1 := Product{
	// 	ID:           "P1",
	// 	Name:         "Laptop",
	// 	Description:  "Gaming Laptop",
	// 	InternalCode: "SECRET_123",
	// 	Stock:        &zeroStock,
	// 	IsActive:     &active,
	// }

	// b1, _ := json.Marshal(p1)
	// fmt.Println("Skenario 1 (lengkap):", string(b1))

	// p2 := Product{
	// 	ID:   "P2",
	// 	Name: "Mouse",
	// }

	// b2, _ := json.Marshal(p2)
	// fmt.Println("Skenario 2 (minimal):", string(b2))

	db1 := challenge.NewDB("localhost", "secret")
	b1, _ := json.Marshal(db1)
	fmt.Println("Default:", string(b1))

	db2 := challenge.NewDB("localhost", "secret", challenge.WithMaxIdleTime(0))
	b2, _ := json.Marshal(db2)
	fmt.Println("With Zero:", string(b2))
}
