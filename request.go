package main

import (
	"net/http"
	"time"

	"github.com/apiotrowski312/benchHog/results"
)

// Get - simple function to Get url and collect data
func Get(url string) results.Measurement {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		return results.CreateMeasurment(time.Since(start), false)
	}

	defer resp.Body.Close()
	wasSuccess := resp.StatusCode >= 200 && resp.StatusCode <= 299

	return results.CreateMeasurment(time.Since(start), wasSuccess)

}
