package main

import (
	"sync"
	"time"
)

const (
	ratio            = 5
	numberOfRequests = 200
)

func main() {

	var wg sync.WaitGroup

	limitRatio := make(chan int, ratio)
	waitTime := make(chan time.Duration, numberOfRequests)

	wg.Add(numberOfRequests)
	for i := 0; i < numberOfRequests; i++ {
		limitRatio <- 1
		go Get("https://budget.master.thebitbybit.com/nginx_status", limitRatio, &wg, waitTime)
	}

	wg.Wait()
	close(limitRatio)
	close(waitTime)

	Metrics(waitTime)

}
