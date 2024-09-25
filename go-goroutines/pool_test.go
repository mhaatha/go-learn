package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	wg := sync.WaitGroup{}
	// Default value
	pool := sync.Pool{
		New: func() interface{} {
			return "Default Value"
		},
	}

	pool.Put("Eko")
	pool.Put("Kurniawan")
	pool.Put("Khannedy")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	wg.Wait()
	fmt.Println("Selesai")
}
