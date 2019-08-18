package main

import (
	"sync"

	"github.com/apiotrowski312/benchHog/results"
)

const (
	ratio            = 25
	numberOfRequests = 250
)

func main() {

	var wg sync.WaitGroup

	limitRatio := make(chan int, ratio)
	measurment := make(chan results.Measurement, numberOfRequests)

	wg.Add(numberOfRequests)
	for i := 0; i < numberOfRequests; i++ {
		limitRatio <- 1
		go Get("http://onet.pl", limitRatio, &wg, measurment)
	}

	wg.Wait()
	close(limitRatio)
	close(measurment)

	results.PrintResults(measurment)

}
