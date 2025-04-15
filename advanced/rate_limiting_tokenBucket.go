package main

import (
	"fmt"
	"time"
)

func rate_limit_main() {
	ratelimiter := NewRateLimiter(5, time.Second)

	for range 10 {
		select {
		case <-ratelimiter.tokens:
			fmt.Println("Request Allowed")
		default:
			fmt.Println("Request Denied")
		}
		time.Sleep(time.Millisecond * 200)
	}
}

type RateLimiter struct {
	tokens     chan struct{}
	refillTime time.Duration
}

func NewRateLimiter(rateLimit int, refillTime time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens:     make(chan struct{}, rateLimit),
		refillTime: refillTime,
	}

	for range rateLimit {
		rl.tokens <- struct{}{}
	}

	go func() {
		rl.refillTokens()
	}()

	return rl
}

func (rl *RateLimiter) refillTokens() {
	ticker := time.NewTicker(rl.refillTime)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			select {
			case rl.tokens <- struct{}{}:
			default:
			}
		}
	}
	close(rl.tokens)
}
