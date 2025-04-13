package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
		close(ch1)
	}()

	go func() {
		ch2 <- 2
		close(ch2)
	}()

	for {
		select {
		case value, ok := <-ch1:
			if !ok {
				fmt.Println("Channel 1 closed")
				return //we return not break, this also gets us outside of for loop
			}
			fmt.Printf("Received value from channel 1 = %d\n", value)

		case value, ok := <-ch2:
			if !ok {
				fmt.Println("Channel 2 closed")
				return
			}
			fmt.Printf("Received value from channel 2 = %d\n", value)

		case <-time.After(time.Second * 3):
			fmt.Println("No channels ready")
		}
	}

}
