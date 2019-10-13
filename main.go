package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/apiotrowski312/benchHog/results"
)

func main() {

	//

	if len(os.Args) < 2 {
		fmt.Println("Provide link to site")
		os.Exit(1)
	}
	link := os.Args[1]

	if !strings.HasPrefix(link, "http") {
		link = "http://" + link

	}
	// Parse link (add http and e)

	// Flags
	var numberOfRequests int
	var ratio int

	description := "Provide number of requests"
	flag.IntVar(&numberOfRequests, "request", 100, description)
	flag.IntVar(&numberOfRequests, "r", 100, description+" (shorthand)")

	description = "Provide ratio"
	flag.IntVar(&ratio, "ratio", 10, description)

	flag.Parse()

	var wg sync.WaitGroup

	limitRatio := make(chan int, ratio)
	measurment := make(chan results.Measurement, numberOfRequests)

	wg.Add(numberOfRequests)
	for i := 0; i < numberOfRequests; i++ {
		limitRatio <- 1
		go Get(link, limitRatio, &wg, measurment)
	}

	wg.Wait()
	close(limitRatio)
	close(measurment)

	results.PrintResults(measurment)

}
