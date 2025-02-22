package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"Apple Mac Pro","price":"20000000"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P001",
		"name":  "Apple Mac Pro",
		"price": "20000000",
	}
	bytes, _ := json.Marshal(product)

	fmt.Println(string(bytes))
}
