package main

import (
	"io"
	"net/http"
	"time"

	"github.com/apiotrowski312/benchHog/results"
)

// Get - simple function to Get url and collect data
func Get(url string) results.Result {
	get := func() (*http.Response, error) {
		return http.Get(url)
	}

	return request(get)
}

// Post - simple function to Get url and collect data
func Post(url string, byteType string, bytes io.Reader) results.Result {
	post := func() (*http.Response, error) {
		return http.Post(url, byteType, bytes)
	}

	return request(post)
}

func request(method func() (*http.Response, error)) results.Result {
	start := time.Now()
	resp, err := method()

	if err != nil {
		return results.CreateResult(time.Since(start), 500)
	}

	defer resp.Body.Close()

	return results.CreateResult(time.Since(start), resp.StatusCode)

}
