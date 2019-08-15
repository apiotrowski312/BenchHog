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
	http.Get(url)

	<-limitRatio
	waitTime <- Measurement{waitTime: time.Since(start), success: true}
}
