package main

import (
	"fmt"
	"time"
)

func non_blocking() {
	data := make(chan int)
	quit := make(chan bool)

	go func() {
		for i := range 5 {
			data <- i
			time.Sleep(time.Second)
		}
		close(data)
		quit <- false
	}()

	for {
		select {
		case value := <-data:
			fmt.Printf("Data received: %d\n", value)
		case <-quit:
			fmt.Println("Quiting...")
			return
		default:
			fmt.Println("Waiting for data...")
			time.Sleep(time.Millisecond * 500)
		}
	}

}
