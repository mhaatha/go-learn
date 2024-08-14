package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")

	if result != "Hello Eko" {
		panic("Result is not 'Hello Eko'")
	}
}

func TestHelloWorldKhannedy(t *testing.T) {
	result := HelloWorld("Khannedy")

	if result != "Hello Khannedy" {
		panic("Result is not 'Hello Khannedy'")
	}
}
