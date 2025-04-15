package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func atomic_main() {
	var wg sync.WaitGroup
	numOfgoroutines := 10
	ac := &AtomicCounter{}

	for range numOfgoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				ac.increment()
			}
		}()
	}

	wg.Wait()
	fmt.Println(ac.getCount())
}

type AtomicCounter struct {
	count int64
}

func (ac *AtomicCounter) increment() {
	atomic.AddInt64(&ac.count, 1)
}

func (ac *AtomicCounter) getCount() int64 {
	return atomic.LoadInt64(&ac.count)
}
