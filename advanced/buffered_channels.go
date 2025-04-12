package main

import (
	"fmt"
	"time"
)

func buffered() {
	ch := make(chan int, 2)

	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println(<-ch)
		time.Sleep(time.Second * 2)
		fmt.Println("Ho ho")
	}()

	fmt.Println("before 1")
	ch <- 1
	fmt.Println("before 2")
	ch <- 2
	fmt.Println("before 3")
	ch <- 3
}
