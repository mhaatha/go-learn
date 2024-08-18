package go_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func DisplayNumber(number int) {
	fmt.Println("Dipslay", number)
}

func TestManyGoroutines(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Upss")

	time.Sleep(time.Second)
}
