package main

import (
	"fmt"
	"strconv"
)

func channel_sync() {

	data := make(chan string)

	go func() {
		for i := range 4 {
			data <- "Hello " + strconv.Itoa(i)
		}
		close(data)
	}()

	// for range data {
	// 	fmt.Println(<-data)
	// }

	for value := range data {
		fmt.Println(value)
	}
}
