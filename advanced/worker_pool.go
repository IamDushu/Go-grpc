package main

import (
	"fmt"
	"time"
)

// 9 seconds -> 10 jobs
// 14 seconds -> 20 jobs
type ticket struct {
	personId   int
	numOfSeats int
	cost       int
}

func ticketProcessor(reqId int, ticketRequests <-chan ticket, ticketResults chan<- ticket) {
	for val := range ticketRequests {
		fmt.Println("Processing request from ", reqId)
		// simulating processing
		time.Sleep(2 * time.Second)
		ticketResults <- val
	}
}

func main() {
	fmt.Println("Time start: ", time.Now())
	numOfRequests := 10
	price := 10
	ch1 := make(chan ticket, numOfRequests)
	ch2 := make(chan ticket, numOfRequests)

	for val := range numOfRequests {
		ch1 <- ticket{personId: val + 1, numOfSeats: (val + 1) * 2, cost: price * ((val + 1) * 2)}
	}
	close(ch1)

	for i := range 3 {
		go ticketProcessor(i, ch1, ch2)
	}

	for range numOfRequests {
		fmt.Printf("Ticket Booked: %v \n", <-ch2)
	}
	fmt.Println("Time End: ", time.Now())
}

// === BASIC WORKER POOL PATTERN
// func worker(id int, tasks <-chan int, results chan<- int) {
// 	for val := range tasks {
// 		total := val * 2
// 		fmt.Printf("Worker %d finished task by adding value: %d\n", id, total)
// 		// simulate long work
// 		time.Sleep(time.Second)
// 		results <- total
// 	}
// 	// Can't close hear as other workers would need it
// 	// close(results)
// }

// func main() {
// 	numOfWorkers := 3
// 	numOfTasks := 10
// 	ch1 := make(chan int, numOfTasks)
// 	ch2 := make(chan int, numOfTasks)

// 	for val := range numOfWorkers {
// 		go worker(val, ch1, ch2)
// 	}

// 	for val := range numOfTasks {
// 		ch1 <- val
// 	}
// 	close(ch1)

// 	for range numOfTasks {
// 		result := <-ch2
// 		fmt.Println("Result:", result)
// 	}
// }
