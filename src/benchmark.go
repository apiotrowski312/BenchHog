package main

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"sync"
	"time"

	"github.com/apiotrowski312/benchHog/src/helpers"
	"github.com/apiotrowski312/benchHog/src/results"
)

func getBenchmarkCommand(details CommandDetails) {
	get := func(url string) results.Result {
		return Get(url)
	}

	finialMesurments := startBenchmark(get, details)

	results.PrintResults(finialMesurments)
}

func postBenchmarkCommand(details CommandDetails) {

	var jsonBytes []byte

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

		helpers.ShowLoader(i, numOfRequests)
	}

	wg.Wait()
	close(limitRatio)
	close(results)

	return results
}
