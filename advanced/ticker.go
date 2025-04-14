package main

import (
	"fmt"
	"time"
)

func ticker() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	i := 1
	for range 5 {
		i *= 2
		fmt.Println(i)
	}
}
