package main

import (
	"net/http"
	"sync"
	"time"
)

// Get - simple function to Get url and collect waitTime
func Get(url string, limitRatio chan int, wg *sync.WaitGroup, waitTime chan Measurement) {
	defer wg.Done()

	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	wasSuccess := resp.StatusCode >= 200 && resp.StatusCode <= 299

	<-limitRatio
	waitTime <- Measurement{waitTime: time.Since(start), success: wasSuccess}
}
