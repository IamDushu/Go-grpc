package main

import "fmt"

func closing_channels() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producerFunc(ch1)
	go filter(ch1, ch2)

	for val := range ch2 {
		fmt.Printf("Value received: %d\n", val)
	}
}

func producerFunc(ch chan<- int) {
	for value := range 5 {
		ch <- value
	}
	close(ch)
}

func filter(in <-chan int, out chan<- int) {
	for val := range in {
		if val%2 == 0 {
			out <- val
		}
	}
	close(out)
}
