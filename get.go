package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/apiotrowski312/benchHog/results"
)

// Get - simple function to Get url and collect data
func Get(url string, limitRatio chan int, wg *sync.WaitGroup, data chan results.Measurement) {
	defer wg.Done()

	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		data <- results.CreateMeasurment(time.Since(start), false)
	} else {
		defer resp.Body.Close()
		wasSuccess := resp.StatusCode >= 200 && resp.StatusCode <= 299

		data <- results.CreateMeasurment(time.Since(start), wasSuccess)
	}

	<-limitRatio
}
