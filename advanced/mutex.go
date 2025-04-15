package main

import (
	"fmt"
	"sync"
)

func mutex_main() {
	numOfgoroutines := 10
	var wg sync.WaitGroup
	counter := &counter{}
	// counter := counter{}  --> how is using this different?

	for range numOfgoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
				// counter.count++ --> without mutex; you get wrong results
			}
		}()
	}

	wg.Wait()
	fmt.Println(counter.getCount())
}

type counter struct {
	mu    sync.Mutex
	count int
}

func (c *counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *counter) getCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
