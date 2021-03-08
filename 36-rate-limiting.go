package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter // limit to 1 request every 200 ms
		fmt.Println("request ", req, time.Now())
	}

	// Allow short bursts of request while maintaining
	// overall rate limit by buffering

	burstyLimiter := make(chan time.Time, 3)

	// Fill up the channel to represent allowed bursting
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
