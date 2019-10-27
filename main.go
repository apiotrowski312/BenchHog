package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/apiotrowski312/benchHog/results"
)

type CommandDetails struct {
	requestType   string
	ratio         int
	numOfRequests int
	links         []string
}

func main() {

	// Flags
	commandDetails := CommandDetails{}

	flag.IntVar(&commandDetails.numOfRequests, "r", 100, "Provide number of requests")
	flag.IntVar(&commandDetails.ratio, "ratio", 10, "Provide ratio")
	flag.Parse()

	links := make([]string, len(os.Args)-2)
	for index, link := range os.Args[2:] {
		links[index] = parseLink(link)
	}

	commandDetails.links = links

	switch os.Args[1] {
	case "help":
		helpCommand()
	case "get":
		getBenchmarkCommand(commandDetails)
	default:
		helpCommand()
	}
}

func helpCommand() {
	fmt.Printf("Available commands:\n")
	fmt.Printf("help\twill show this help\n")
	fmt.Printf("get\twill get things\n") // TODO: descriptions
}

func getBenchmarkCommand(details CommandDetails) {
	details.requestType = "get"
	finialMesurments := startBenchmark(details)

	results.PrintResults(finialMesurments)
}

func startBenchmark(details CommandDetails) chan results.Result {

	ratio, numOfRequests, links := details.ratio, details.numOfRequests, details.links

	var wg sync.WaitGroup
	limitRatio := make(chan int, ratio)
	results := make(chan results.Result, numOfRequests)

	rand.Seed(time.Now().Unix())

	wg.Add(numOfRequests)
	for i := 0; i < numOfRequests; i++ {
		limitRatio <- 1
		go func() {
			defer wg.Done()
			results <- Get(links[rand.Intn(len(links))])
			<-limitRatio
		}()

		showLoader(i, numOfRequests)
	}

	fmt.Println()

	wg.Wait()
	close(limitRatio)
	close(results)

	return results
}
