package go_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestTick(t *testing.T) {
	ticker := time.Tick(1 * time.Second)

	for tick := range ticker {
		fmt.Println(tick)
	}
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for tick := range ticker.C {
		fmt.Println(tick)
	}
}
