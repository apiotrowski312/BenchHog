package main

import (
	"sync"
	"time"
)

const (
	ratio            = 5
	numberOfRequests = 25
)

type Measurement struct {
	waitTime time.Duration
	success  bool
}

func main() {

	var wg sync.WaitGroup

	limitRatio := make(chan int, ratio)
	measurment := make(chan Measurement, numberOfRequests)

	wg.Add(numberOfRequests)
	for i := 0; i < numberOfRequests; i++ {
		limitRatio <- 1
		go Get("https://medvocation.develop.thebitbybit.com/api/choice/profession/", limitRatio, &wg, measurment)
	}

	wg.Wait()
	close(limitRatio)
	close(measurment)

	PrintResults(measurment)

}
