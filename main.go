package main

import (
	"sync"
)

const (
	ratio            = 25
	numberOfRequests = 250
)

func main() {

	var wg sync.WaitGroup

	limitRatio := make(chan int, ratio)
	measurment := make(chan Measurement, numberOfRequests)

	wg.Add(numberOfRequests)
	for i := 0; i < numberOfRequests; i++ {
		limitRatio <- 1
		go Get("URL", limitRatio, &wg, measurment)
	}

	wg.Wait()
	close(limitRatio)
	close(measurment)

	PrintResults(measurment)

}
