package main

import "fmt"

func channelDirection() {
	data := make(chan int)
	producer(data)
	consumer(data)
}

func producer(ch chan<- int) {
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()
}

func consumer(ch <-chan int) {
	for value := range ch {
		fmt.Printf("Received Data: %d\n", value)
	}
}
