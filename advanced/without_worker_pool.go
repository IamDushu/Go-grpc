package main

// 20 seconds -> 10 jobs
// 40 seconds -> 20 jobs
// type ticket struct {
// 	personId   int
// 	numOfSeats int
// 	cost       int
// }

// func ticketProcessor(reqId int, ticketRequests *[]ticket, ticketResults *[]ticket) {
// 	for _, val := range *ticketRequests {
// 		fmt.Println("Processing request from ", reqId)
// 		// simulating processing
// 		time.Sleep(2 * time.Second)
// 		*ticketResults = append(*ticketResults, val)
// 	}
// }

// func main() {
// 	fmt.Println("Time start: ", time.Now())
// 	numOfRequests := 20
// 	price := 10
// 	ch1 := []ticket{}
// 	ch2 := []ticket{}

// 	for val := range numOfRequests {
// 		ch1 = append(ch1, ticket{personId: val + 1, numOfSeats: (val + 1) * 2, cost: price * ((val + 1) * 2)})
// 	}

// 	ticketProcessor(1, &ch1, &ch2)

// 	for i := range numOfRequests {
// 		fmt.Printf("Ticket Booked: %v \n", ch2[i])
// 	}
// 	fmt.Println("Time End: ", time.Now())
// }
