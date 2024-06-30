package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestWithCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(
		context.Background(),
	)

	go printUntilCancel(ctx)

	time.Sleep(5 * time.Second)

	cancel()
	// t.Error(errors.New("force error"))
}

func printUntilCancel(ctx context.Context) {
	count := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Cancel signal received, exiting")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Printf("Printing until cancel, number = %d \n", count)
			count += 1
		}
	}
}
