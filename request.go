package main

import (
	"net/http"
	"time"

	"github.com/apiotrowski312/benchHog/results"
)

// Get - simple function to Get url and collect data
func Get(url string) results.Result {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		return results.CreateResult(time.Since(start), 500)
	}

	defer resp.Body.Close()

	return results.CreateResult(time.Since(start), resp.StatusCode)

}
