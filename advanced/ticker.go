package main

import (
	"fmt"
	"time"
)

func ticker() {
	tick := time.NewTicker(time.Second)
	timeout := time.After(5 * time.Second)

	defer tick.Stop()

	for {
		select {
		case val := <-tick.C:
			fmt.Println("Ticking..", val)
		case <-timeout:
			fmt.Println("Stopped ticker")
			return
		}
	}
}

// func main() {
// 	ticker := time.NewTicker(2 * time.Second)
// 	defer ticker.Stop()

// 	// i := 1
// 	// for range 5 {
// 	// 	i *= 2
// 	// 	fmt.Println(i)
// 	// }

// 	i := 1
// 	for val := range ticker.C {
// 		i *= 2
// 		fmt.Println(val)
// 	}
// }
