package go_goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	wg := sync.WaitGroup{}

	for i := 1; i <= 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("Counter =", x)
}
