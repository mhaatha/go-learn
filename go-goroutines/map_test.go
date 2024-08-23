package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, value int, wg *sync.WaitGroup) {
	defer wg.Done()

	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go AddToMap(data, i, wg)
	}

	wg.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return false
	})
}
