package main

import (
	"context"
	"fmt"
	"time"
)

func contextExample() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	ctx = context.WithValue(ctx, "requestId", "1234")

	go doWork(ctx)

	time.Sleep(3 * time.Second)

	requestID := ctx.Value("requestId")
	if requestID != nil {
		fmt.Println("Request ID: ", requestID)
	} else {
		fmt.Println("No Request ID found.")
	}
}

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work Cancelled: ", ctx.Err())
			return
		default:
			fmt.Println("Working...")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
