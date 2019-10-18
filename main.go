package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/apiotrowski312/benchHog/results"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Provide link to site")
		os.Exit(1)
	}

	links := make([]string, len(os.Args)-1)
	for index, link := range os.Args[1:] {
		links[index] = parseLink(link)
	}

	// Flags
	var numberOfRequests int
	var ratio int

	description := "Provide number of requests"
	flag.IntVar(&numberOfRequests, "request", 100, description)
	flag.IntVar(&numberOfRequests, "r", 100, description+" (shorthand)")

	description = "Provide ratio"
	flag.IntVar(&ratio, "ratio", 10, description)
	flag.Parse()

	// Rest of code
	finialMesurments := startBenchmark(numberOfRequests, ratio, links)

	results.PrintResults(finialMesurments)

}

func startBenchmark(numberOfRequests int, ratio int, links []string) chan results.Measurement {
	var wg sync.WaitGroup
	limitRatio := make(chan int, ratio)
	measurment := make(chan results.Measurement, numberOfRequests)

	rand.Seed(time.Now().Unix())

	wg.Add(numberOfRequests)
	for i := 0; i < numberOfRequests; i++ {
		limitRatio <- 1
		go func() {
			defer wg.Done()
			measurment <- Get(links[rand.Intn(len(links))])
			<-limitRatio
		}()
	}

	wg.Wait()
	close(limitRatio)
	close(measurment)

	return measurment
}

func parseLink(link string) string {
	if !strings.HasPrefix(link, "http") {
		return "http://" + link
	}
	return link
}
