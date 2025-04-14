package main

import (
	"fmt"
	"time"
)

// =========== SCHEDULING DELAYED OPERATIONS
func timer() {
	timer := time.NewTimer(2 * time.Second) // non blocking timer starts

	go func() {
		<-timer.C
		fmt.Println("Delayed operation executed")
	}()

	fmt.Println("Waiting...")
	time.Sleep(3 * time.Second) // blocking timer starts
	fmt.Println("End of the program")
}

// func main() {
// 	tm1 := time.NewTimer(4 * time.Second)
// 	tm2 := time.NewTimer(2 * time.Second)

// 	select {
// 	case <-tm1.C:
// 		fmt.Println("Timer 1 executed")
// 	case <-tm2.C:
// 		fmt.Println("Timer 2 executed")
// 	}
// }

// ============= TIMEOUT
// func longRunningOperation() {
// 	for i := range 20 {
// 		fmt.Println(i)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	timeout := time.After(3 * time.Second)
// 	done := make(chan bool)

// 	go func() {
// 		longRunningOperation()
// 		done <- true
// 	}()

// 	select {
// 	case <-timeout:
// 		fmt.Println("operation timed out")
// 	case <-done:
// 		fmt.Println("Operation completed")
// 	}

// }

// ======== BASIC TIMER USE
// func main() {
// 	fmt.Println("Starting app.")
// 	timer := time.NewTimer(2 * time.Second)
// 	fmt.Println("Waiting for timer.c")
// 	stopped := timer.Stop()
// 	if stopped {
// 		fmt.Println("Timer stopped")
// 	}
// 	timer.Reset(time.Second)
// 	fmt.Println("Timer reset")
// 	<-timer.C // blocking in nature
// 	fmt.Println("Timer expired")
// }
