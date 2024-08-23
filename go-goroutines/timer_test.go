package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestAfterFunc(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	time.AfterFunc(5*time.Second, func() {
		defer wg.Done()

		fmt.Println(time.Now())
	})

	fmt.Println(time.Now())

	wg.Wait()
}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println("Time Now:", time.Now())

	afterFiveSeconds := <-channel
	fmt.Println("Time After:", afterFiveSeconds)
}

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("Time Now:", time.Now())

	afterFiveSeconds := <-timer.C
	fmt.Println("Time Now:", afterFiveSeconds)
}
