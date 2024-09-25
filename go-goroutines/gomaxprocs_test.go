package go_goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetThreadNumber(t *testing.T) {
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second)
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU:", totalCPU)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total thread:", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine:", totalGoroutine)

	wg.Wait()
}

func TestGetGomaxprocs(t *testing.T) {
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second)
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU:", totalCPU)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total thread:", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine:", totalGoroutine)

	wg.Wait()
}
