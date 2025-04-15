package main

import (
	"fmt"
	"sync"
	"time"
)

type ticke struct {
	personId   int
	numOfSeats int
	cost       int
}

func wait_group() {
	fmt.Println("Time start: ", time.Now())
	var wg sync.WaitGroup
	numOfRequests := 10
	price := 10
	ch1 := make(chan ticke, numOfRequests)
	ch2 := make(chan ticke, numOfRequests)

	for val := range numOfRequests {
		ch1 <- ticke{personId: val + 1, numOfSeats: (val + 1) * 2, cost: price * ((val + 1) * 2)}
	}
	close(ch1)

	for i := range 3 {
		wg.Add(1)
		go tickeProcessor(&wg, i, ch1, ch2)
	}

	go func() {
		wg.Wait()
		close(ch2)
	}()

	for result := range ch2 {
		fmt.Printf("Ticket Booked: %v \n", result)
	}
	fmt.Println("Time End: ", time.Now())
}

func tickeProcessor(wg *sync.WaitGroup, reqId int, ticketRequests <-chan ticke, ticketResults chan<- ticke) {
	defer wg.Done()
	for val := range ticketRequests {
		fmt.Println("Processing request from ", reqId)
		// simulating processing
		time.Sleep(2 * time.Second)
		ticketResults <- val
	}
}
