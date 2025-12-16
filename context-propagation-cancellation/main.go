package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mhaatha/go-learn/context-propagation-cancellation/challenge"
)

func SlowOperation(ctx context.Context) (string, error) {
	resultChan := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		resultChan <- "Data Berhasil Diambil!"
	}()

	select {
	case res := <-resultChan:
		return res, nil

	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func main() {
	fmt.Println("--- Test 1: Timeout 1000ms ---")

	ctx1, cancel1 := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel1()

	res, err := challenge.CheckPaymentStatus(ctx1)
	if err != nil {
		fmt.Println("Hasil:", err)
		fmt.Println("-> Transaksi digagalkan!")
	} else {
		fmt.Println("Hasil:", res)
	}

	fmt.Println()

	fmt.Println("--- Test 2: Timeout 4000ms ---")

	ctx2, cancel2 := context.WithTimeout(context.Background(), 4000*time.Millisecond)
	defer cancel2()

	res2, err2 := challenge.CheckPaymentStatus(ctx2)
	if err2 != nil {
		fmt.Println("Hasil:", err2)
	} else {
		fmt.Println("Hasil:", res2)
	}
}
