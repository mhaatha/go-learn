package go_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	// Main Parent
	contextA := context.Background()

	// contextB and contextC are children (child) of contextA
	contextB := context.WithValue(contextA, "B", "A's Child-1")
	contextC := context.WithValue(contextA, "C", "A's Child-2")

	contextD := context.WithValue(contextB, "D", "B's Child-1")
	contextE := context.WithValue(contextB, "E", "B's Child-2")

	contextF := context.WithValue(contextC, "F", "C's Child-1")

	fmt.Println(contextA) // context.Background
	fmt.Println(contextB) // context.Background.WithValue(type string, val A's Child-1)
	fmt.Println(contextC) // context.Background.WithValue(type string, val A's Child-2)
	fmt.Println(contextD) // context.Background.WithValue(type string, val A's Child-1).WithValue(type string, val B's Child-1)
	fmt.Println(contextE) // context.Background.WithValue(type string, val A's Child-1).WithValue(type string, val B's Child-2)
	fmt.Println(contextF) // context.Background.WithValue(type string, val A's Child-2).WithValue(type string, val C's Child-1)

	// Get Value
	fmt.Println(contextC.Value("C")) // Receive value from contextC
	fmt.Println(contextD.Value("B")) // Receive value from contextD's parent: contextB value
	fmt.Println(contextE.Value("C")) // nil
	fmt.Println(contextA.Value("C")) // nil
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(time.Second) // Simulate slow process
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter:", n)
		if n == 10 {
			break
		}
	}

	// Send cancel signal to context
	cancel()

	// To close the goroutine, we stop the program for 2 second
	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter:", n)
	}

	// To close the goroutine, we stop the program for 2 second
	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter:", n)
	}

	// To close the goroutine, we stop the program for 2 second
	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
}

// ChatGPT Context Practice
func FetchData(ctx context.Context, url string) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(1 * time.Second)
				fmt.Println(url)
			}
		}
	}()
}

func TestFetchData(t *testing.T) {
	t.Logf("Total Goroutine before fetch: %d", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	FetchData(ctx, "https://github.com/mhaatha")

	time.Sleep(6 * time.Second)

	cancel()

	time.Sleep(2 * time.Second)

	t.Logf("Total Goroutine: %d", runtime.NumGoroutine())
}
