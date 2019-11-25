package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/apiotrowski312/benchHog/results"
)

// CommandDetails - TODO
type CommandDetails struct {
	ratio         int
	numOfRequests int
	jsonString    string
	jsonFile      string
	links         []string
}

func main() {

	// Flags
	commandDetails := CommandDetails{}

	flag.IntVar(&commandDetails.numOfRequests, "r", 100, "Provide number of requests")
	flag.StringVar(&commandDetails.jsonString, "body", "", "Json for post request")
	flag.StringVar(&commandDetails.jsonFile, "bodyFile", "", "File (with json usually) for post request (higher priority than string)")
	flag.IntVar(&commandDetails.ratio, "ratio", 10, "Provide ratio")
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Provide method and links")
		os.Exit(1)
	}

	links := make([]string, len(flag.Args())-1)
	for index, link := range flag.Args()[1:] {
		links[index] = parseLink(link)
	}

	commandDetails.links = links

	switch flag.Args()[0] {
	case "help":
		helpCommand()
	case "get":
		getBenchmarkCommand(commandDetails)
	case "post":
		postBenchmarkCommand(commandDetails)
	default:
		helpCommand()
	}
}

func helpCommand() {
	fmt.Printf("Available commands:\n")
	fmt.Printf("help\tshow help\n")
	fmt.Printf("get\t start benchmark with get request\n")   // TODO: descriptions
	fmt.Printf("post\t start benchmark with post request\n") // TODO: descriptions
}

func getBenchmarkCommand(details CommandDetails) {
	get := func(url string) results.Result {
		return Get(url)
	}

	finialMesurments := startBenchmark(get, details)

	results.PrintResults(finialMesurments)
}

func postBenchmarkCommand(details CommandDetails) {

	var jsonBytes []byte

	fmt.Println(details.jsonFile)

	if details.jsonFile != "" {
		fileData, err := ioutil.ReadFile(details.jsonFile)

		if err != nil {
			panic(err)
		}
		jsonBytes = fileData
	} else {
		jsonBytes = []byte(details.jsonString)
	}

	post := func(url string) results.Result {
		return Post(url, "application/json", bytes.NewBuffer(jsonBytes))
	}

	finialMesurments := startBenchmark(post, details)

	results.PrintResults(finialMesurments)
}

func startBenchmark(request func(url string) results.Result, details CommandDetails) chan results.Result {

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
			results <- request(links[rand.Intn(len(links))])
			<-limitRatio
		}()

		showLoader(i, numOfRequests)
	}

	wg.Wait()
	close(limitRatio)
	close(results)

	return results
}
