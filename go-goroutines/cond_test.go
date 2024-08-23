package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var wg = sync.WaitGroup{}

func WaitCondition(value int) {
	defer wg.Done()

	cond.L.Lock()
	cond.Wait()

	fmt.Println("Done", value)

	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			cond.Signal()
		}
	}()

	wg.Wait()
}
