package challenge

import (
	"context"
	"time"
)

func CheckPaymentStatus(ctx context.Context) (string, error) {
	resultChan := make(chan string, 1)

	go func() {
		time.Sleep(3000 * time.Millisecond)
		resultChan <- "Payment Success"
	}()

	select {
	case msg := <-resultChan:
		return msg, nil

	case <-ctx.Done():
		return "", ctx.Err()
	}
}
