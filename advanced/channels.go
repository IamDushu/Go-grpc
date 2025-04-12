package main

import (
	"fmt"
	"time"
)

func main() {
	greeting := make(chan string)

	go func() {
		fmt.Println("Inside goroutine")
		greeting <- "Hello"
	}()

	time.Sleep(time.Second * 5)
	fmt.Println("Before receive")
	greet := <-greeting //blocking
	fmt.Println("After receive")

	fmt.Println(greet)
}
